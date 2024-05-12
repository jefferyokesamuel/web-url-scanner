package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"ioutil"
	"strings"
)

func checkURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is down\n", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s -> status code is %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1]
			file += ".txt"

			fmt.Printf("Writing URL body into %s\n", file)

			err = os.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func main() {
	urls := []string{"https://golang.org", "https://github.com", "https://medium.com"}

	for _, url := range urls{
		checkURL(url)
		fmt.Println(strings.Repeat("*", 20))
	}
}