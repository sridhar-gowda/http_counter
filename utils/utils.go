package utils

import (
	"log"
	"net/http"
	"os"
	"time"
)

func GetFile(fileName string) *os.File {
	var file *os.File
	_, err := os.Stat(fileName)
	if err != nil {
		file, err = createFile(fileName)
	} else {
		file = openFile(fileName)
	}
	return file
}

func openFile(fileName string) *os.File {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("Unable to open file")
	}
	return file
}

func createFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Unable to create file")
	}
	return file, err
}

func GetLatestWindow(now int64, size int) int64 {
	return now-int64(size)
}

func GetCurrentTime()int64{
	return time.Now().Unix()
}

func GetIpFromHeader( req *http.Request)string{
	return req.Header.Get("dummy-ip")
}
