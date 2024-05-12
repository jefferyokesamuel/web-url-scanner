package main

import (
	"fmt"
	"net/http"
)

func checkURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is down\n", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s -> status code is %d\n", url, resp.StatusCode)
	}

}

func main() {

}