package api

/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>

*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/grahamplata/sixers/config"
	"github.com/grahamplata/sixers/utils"
)

// urlBuilder takes the query parameters and returns a string to be used to retrieve data from the api
func urlBuilder(r Request) string {
	url := fmt.Sprintf("%s/?seasons[]=%s,%s&postseason=False&team_ids[]=%v&per_page=100",
		config.APIURL, r.Year, utils.DecrementString(r.Year), config.TeamID)
	return url
}

// Fetch ...
func Fetch(r Request) Response {
	var responseObject Response

	url := urlBuilder(r)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &responseObject)

	return responseObject
}
