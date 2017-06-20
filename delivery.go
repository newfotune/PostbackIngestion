package main

import (
	"log"
	"os"
	"time"

	"encoding/json"

	"strings"

	"gopkg.in/redis.v3"
)

type (
	//Request - Struct to hold the incoming request data
	Request struct {
		Method string        `json:"method"`
		URL    string        `json:"url"`
		Data   []RequestData `json:"data"`
	}

	RequestData struct {
		Mascot   string `json:"mascot"`
		Location string `json:"location"`
		Bar      string `json:"bar"`
	}
)

var ticker *time.Ticker

func main() {

	client, err := createClient("localhost:6379", "evenkingscry!", 0)
	if err != nil {
		panicAndExit(err)
	}
	file := getLogFile("requests.log")

	logger := log.New(file, "", log.Ltime)

	ticker = time.NewTicker(time.Second)
	go func() {

		for range ticker.C {
			keys, err := client.Keys("*").Result()
			if err != nil {
				panicAndExit(err)
			}
			go processData(client, logger, keys)
		}
	}()

	for {
		//time.Sleep(time.Millisecond * 3500)
		//ticker.Stop()
	}

}

/*
	Process all request data stored in redis, and delete the processed keys from redis.
	@param client: The redis client
	@param logger: The logger for writing response to output file.
	@param []keys: A slice of all keys in redis.
*/
func processData(client *redis.Client, logger *log.Logger, keys []string) {
	for _, key := range keys {
		value, err := client.Get(key).Result()
		if err != nil {
			logger.Printf("Failure on key: %s. %v\n", key, err)
			continue
		}

		_, err = client.Del(key).Result()
		if err != nil {
			panicAndExit(err)
		}

		request := Request{}
		err = json.Unmarshal([]byte(value), &request)
		if err != nil {
			panicAndExit(err)
		}
		processAndLogRequest(logger, request)
	}

	return
}

/*
	processAndLogRequest: Process an individual request and logs the data to the "requests.log" file
	@param request: Request data gotten from redis
	@param logger: The logger for writing response to output file.
*/
func processAndLogRequest(logger *log.Logger, request Request) {
	var mascot = "mascot"
	var location = "location"
	var bar = "bar"
	for _, data := range request.Data {
		url := request.URL

		if len(data.Mascot) == 0 {
			data.Mascot = mascot
		}
		if len(data.Location) == 0 {
			data.Location = location
		}
		if len(data.Bar) == 0 {
			data.Bar = bar
		}
		url = strings.Replace(url, "{mascot}", data.Mascot, 1)

		url = strings.Replace(url, "{location}", data.Location, 1)
		url = strings.Replace(url, "{bar}", data.Bar, 1)

		logger.Printf("%s %s ", request.Method, url)
	}

}

func getLogFile(fileName string) *os.File {
	//read write access, create it it doesnt exist, append data
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panicAndExit(err)
	}
	return file
}

func createClient(address string, password string, dbNumber int64) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       dbNumber, // use default DB
	})

	_, err := client.Ping().Result()

	if err != nil {
		panicAndExit(err)
	}
	return client, nil
}

func panicAndExit(err error) {
	log.Printf("%v\n", err)
	ticker.Stop() //stop time ticker
	os.Exit(-1)
}
