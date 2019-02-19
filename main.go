package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get(fmt.Sprintf("http://lapi.transitchicago.com/api/1.0/ttarrivals.aspx?key=%s&mapid=40380", KEY))
	if err != nil {
		fmt.Println(err);
		return
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		responseBody, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(responseBody))
	}
}
