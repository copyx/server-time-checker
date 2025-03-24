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
		select {
		case <-ticker.C:
			resp, err := http.Head(url)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Local:", time.Now().Format(time.RFC1123), ", Server:", resp.Header["Date"])
		}
	}
}

func printUsages() {
	var usages = [...]string{
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
