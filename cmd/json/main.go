package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Sites struct {
	Items []Site `json:"items"`
}

type Site struct {
	Name string `json:"name"`
}

func main() {
	const url = "https://api.stackexchange.com/2.3/sites"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var sites Sites
	json.Unmarshal(body, &sites)

	for _, site := range sites.Items {
		fmt.Println(site.Name)
	}
}
