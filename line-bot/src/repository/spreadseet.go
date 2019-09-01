package repository

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Home struct {
	Date  string `json:"date"`
	ID    string `json:"id"`
	Value string `json:"value"`
}

func GetHomeList(id string) []Home {
	url := "https://script.google.com/macros/s/AKfycbzw_hkJwGt9NgCMOD06OaWZctk_2qWnH3RAKlNWpPlcbIVgnYiA/exec?id=" + id
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return []Home{}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Home{}
	}
	homeList := make([]Home, 0)
	err = json.Unmarshal(body, &homeList)
	if err != nil {
		return []Home{}
	}
	defer resp.Body.Close()
	return homeList
}

func SetHome(home *Home) {
	url := "https://script.google.com/macros/s/AKfycbzw_hkJwGt9NgCMOD06OaWZctk_2qWnH3RAKlNWpPlcbIVgnYiA/exec"
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	j, err := json.Marshal(home)
	if err != nil {
		log.Println(err)
		return
	}
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(j),
	)
	req.Header.Set("Content-Type", "application/json")
	_,err =client.Do(req)
  log.Println(err)
}
