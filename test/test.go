package main

import (
	"fmt"
	"http_counter/config"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)


// A HTTP CLIENT ON LOOP JUST FOR TESTING PURPOSE.

func main() {

	for {
		cfgFile, err := os.Open("./config/config.yaml")
		if err != nil {
			log.Fatal("Unable to read the config file")
		}
		cfg := config.New(cfgFile)
		cfgData, err := cfg.Get()
		if err != nil {
			log.Fatal(err)
		}


		URL := fmt.Sprintf("http://%s:%d/%s", cfgData.Host, cfgData.Port, cfgData.CountUrl)
		request, err := http.NewRequest("GET", URL, nil)
		if err != nil {
			log.Fatal("Unable to create request")
		}

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			 log.Println(err)
		} else {
			defer response.Body.Close()
			res, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal("Unable to read response")
			}
			log.Print(string(res))
		}

		time.Sleep(time.Duration(1000) * time.Millisecond)
	}

}
