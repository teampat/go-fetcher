package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Repos struct {
	Name string `json:"Name"`
	URL  string `json:"url"`
}

func main() {
	var repos []Repos
	res, _ := http.Get("http://www.mocky.io/v2/5b52fd522f0000510d3bb683")
	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &repos)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	for _, repo := range repos {
		fmt.Printf("Repo name: %s \n", repo.Name)
		fmt.Printf("URL: %s \n", repo.URL)
	}
}
