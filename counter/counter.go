package counter

import (
	"encoding/json"
	"http_counter/config"
	"http_counter/model"
	"http_counter/utils"
	"io"
	"sync"
	"time"
)

type HitCounter struct {
	Visitors map[string][]int64
	Mu       sync.RWMutex
	Limiter  *Limiter
	Cfg      *config.Config
}

func NewCounter(config *config.Config, limiter *Limiter) *HitCounter {
	v := make(map[string][]int64)
	return &HitCounter{Visitors: v, Cfg: config, Limiter: limiter}

}

func (c *HitCounter) Initialize(r io.Reader) error {
	var storeData model.DbStore
	err := storeData.ReadJSON(r)
	if err == nil {
		c.Visitors = model.ToCustomers(&storeData)
	}
	return err
}

func (c *HitCounter) Save(w io.Writer) error {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	if len(c.Visitors) > 0 {
		dbStore := model.ToDbStore(c.Visitors)
		err := json.NewEncoder(w).Encode(dbStore)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *HitCounter) CheckRequest(ip string) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	now := time.Now().Unix()
	c.CheckLimit(ip)
	if c.Limiter.Allowed == false {
		return
	}
	data, exists := c.Visitors[ip]
	if !exists {
		var t []int64
		c.Visitors[ip] = append(t, now)
	} else {
		timeStamps := c.Visitors[ip]
		c.Visitors[ip] = append(timeStamps, now)
	}
	c.updateCount(now, ip, data)
}

func (c *HitCounter) updateCount(now int64, ip string, timeStamps []int64) int {
	var total int
	// Last 60 seconds
	lastMin := utils.GetLatestWindow(now, c.Cfg.Reset)
	for i := len(timeStamps) - 1; i >= 0; i-- {
		if timeStamps[i] > lastMin {
			total++
		} else {
			// Remove entries which are one minute old
			timeStamps = timeStamps[i+1:]
			c.Visitors[ip] = timeStamps
			break
		}
	}
	return total
}

func (c *HitCounter) CheckLimit(ip string) {
	now := utils.GetCurrentTime()
	timeStamps, exists := c.Visitors[ip]
	if !exists || len(timeStamps) == 0 {
		c.Limiter.Allowed = true
		return
	}

	window := utils.GetLatestWindow(now, c.Limiter.Rate)
	var windowCount int
	for i := len(timeStamps) - 1; i >= 0; i-- {
		if timeStamps[i] > window {
			windowCount += 1
		} else {
			break
		}
	}
	if windowCount > c.Limiter.Size {
		c.Limiter.Allowed = false
	} else {
		c.Limiter.Allowed = true
	}
}
