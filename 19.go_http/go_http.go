package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func MyHttp() {
	response, err := http.Get("https://ifconfig.io/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
func httpPut() {
	postData := strings.NewReader(`{"some": "json"}`)
	response, err := http.Post("https://httpbin.org/post", "application/json", postData)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
func http1() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://ifconfig.io/", nil)
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(request)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}

func http2() {
	debug := os.Getenv("DEBUG")
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://ifconfig.io/", nil)
	if err != nil {
		log.Fatal(err)
	}
	if debug == "1" {
		debugRequest, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", debugRequest)
	}
	response, err := client.Do(request)
	defer response.Body.Close()
	if debug == "1" {
		debugResponse, err := httputil.DumpResponse(response, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", debugResponse)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}

func main() {
	http2()
	// http1()
	// httpPut()
	// MyHttp()
}
