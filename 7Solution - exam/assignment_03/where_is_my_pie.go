package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type BeefSummary struct {
	Beef map[string]int `json:"beef"`
}

func fetchBaconIpsumData() (string, error) {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if len(body) == 0 {
		return "", fmt.Errorf("empty response body")
	}

	return string(body), nil
}

func summarizeBeefData(text string) BeefSummary {
	meatTypes := []string{
		"fatback", "t-bone", "pastrami", "pork", "meatloaf", "jowl", "enim", "bresaola",
	}

	text = strings.ToLower(text)
	text = regexp.MustCompile(`[^a-z0-9\s-]`).ReplaceAllString(text, "")

	beefSummary := BeefSummary{Beef: make(map[string]int)}

	for _, word := range strings.Fields(text) {
		for _, meat := range meatTypes {
			if word == meat {
				beefSummary.Beef[meat]++
			}
		}
	}

	return beefSummary
}

func main() {
	data, err := fetchBaconIpsumData()
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	summary := summarizeBeefData(data)

	jsonData, err := json.MarshalIndent(summary, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}
