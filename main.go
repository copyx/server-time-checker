package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		printUsages()
		return
	}

	url := os.Args[1]
	fmt.Println(url)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		start := time.Now().UnixMilli()

		resp, err := http.Head(url)
		if err != nil {
			fmt.Println(err)
		}

		latency := time.Now().UnixMilli() - start
		respDates := resp.Header["Date"]
		if len(respDates) < 1 {
			fmt.Println("There is no date header value in response")
			os.Exit(1)
		}
		localTime := time.Now().Format(time.RFC1123Z)
		serverTime, err := time.Parse(time.RFC1123, respDates[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Local: %s, Server: %s, Latency: %dms\n", localTime, serverTime.Format(time.RFC1123Z), latency)
	}
}

func printUsages() {
	usages := [...]string{
		"server-time-checker is a CLI command to check the server time of target http server.",
		"It will check the server time every seconds infinitely by default",
		"",
		"Usages:",
		"",
		"\tserver-time-checker <http-server-url>",
	}

	for _, v := range usages {
		fmt.Println(v)
	}
}
