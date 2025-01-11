package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func fetchJSON(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching URL: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	return body, nil
}

// it have only 2 choice by 1 layer and 1 way
// just compare from that 2 choice don't mind all of layer just care for index
//Then it fall down shape like triangle

func maxPathSum(triangle [][]int) int {
	// used short hand valiable for reduce line number
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += max(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

// Convention : Function shouldn't in logic code
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	url := "https://raw.githubusercontent.com/7-solutions/backend-challenge/main/files/hard.json"
	body, err := fetchJSON(url)
	if err != nil {
		log.Fatalf("Failed to fetch JSON: %v", err)
	}
	var triangle [][]int
	err = json.Unmarshal(body, &triangle)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}
	result := maxPathSum(triangle)
	fmt.Printf("The maximum path sum is: %d\n", result)
}
