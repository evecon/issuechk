package main

import (
	"fmt"
	"github.com/google/go-github/github"
)

func main() {
	client := github.NewClient(nil)

	for i := 0; i < 100; i++ {
		issue, _, err := client.Issues.Get("go-swagger", "go-swagger", i)
		if err != nil {
			fmt.Errorf("Issues.Get returned error: %v", err)
		} else {
			fmt.Println(*issue.Number, " ", *issue.State)
		}
	}
}
