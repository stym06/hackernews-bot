package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/fatih/color"
)

type HackerStory struct {
	Title string `json: "title"`
	Url   string `json: "url"`
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	arg := os.Args[1]

	rand.Seed(time.Now().UnixNano())
	topStoriesUrl := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/%sstories.json?print=pretty", arg)

	resp, err := http.Get(topStoriesUrl)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var topStories []int64
	err = json.Unmarshal(body, &topStories)
	if err != nil {
		panic(err)
	}

	randomStoryId := topStories[rand.Intn(len(topStories))]

	randomStoryUrl := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", randomStoryId)

	resp, err = http.Get(randomStoryUrl)
	if err != nil {
		panic(err)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var hackerStory HackerStory
	err = json.Unmarshal(body, &hackerStory)
	if err != nil {
		panic(err)
	}

	color.Green(hackerStory.Title)
	color.White(hackerStory.Url)

	openbrowser(hackerStory.Url)

}
