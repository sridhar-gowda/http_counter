package main

import (
	"fmt"
	"http_counter/config"
	"http_counter/counter"
	"http_counter/model"
	"http_counter/utils"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

// Request store handler
type RequestHandler struct {
	counter *counter.HitCounter
	cfg *config.Config
}

// Handles each request
func (rh *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {


	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	rh.counter.CheckRequest(ip)
	fmt.Println(rh.counter.Limiter.Allowed)
	if rh.counter.Limiter.Allowed == false {
		http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := &model.CountResponse{
		Msg:   "REQUEST REGISTERED",
	}
	res.ToJSON(w)
}

func main() {

	// Config
	cfgFile, err := os.Open("./config/config.yaml")
	if err != nil {
		log.Fatal("Unable to read the config file")
	}
	cfg := config.New(cfgFile)
	cfgData, err := cfg.Get()
	if err != nil {
		log.Fatal(err)
	}
	limiter := counter.NewLimiter(cfgData.LimitSize,cfgData.LimitRate)
	dataFile := utils.GetFile(cfgData.DataFile)
	c := counter.NewCounter(cfgData,limiter)
	c.Initialize(dataFile)

	handler := &RequestHandler{counter: c,cfg: cfgData}
	onExit(dataFile, handler)

	mux := http.NewServeMux()
    mux.Handle(cfgData.CountUrl,handler)


	hostPort := fmt.Sprintf("%s:%d", cfgData.Host, cfgData.Port)
	log.Printf("Listening on Port %d\n", cfgData.Port)
	if err := http.ListenAndServe(hostPort, handler); err != nil {
		log.Fatal(err)
	}
}

// Store the state of counter before exit
func onExit(file *os.File, requestHandler *RequestHandler) {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	go func() {
		<-exit
		file.Truncate(0)
		err := requestHandler.counter.Save(file)
		if err != nil {
			log.Println("Could not save on exit")
		}
		log.Println("Saved in the file on exit")
		os.Exit(1)
	}()
}
