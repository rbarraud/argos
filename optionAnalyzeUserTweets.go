package main

import (
	"fmt"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
)

func optionAnalyzeUserTweets(client *twitter.Client, username string) {

	//get tweets
	tweets := getTweets(client, username, iterationsCount)

	if len(tweets) == 0 {
		fmt.Println("User @" + username + " does not have tweets")
		return
	}

	analyzeUserAccount(client, username)

	//now analyze words and dates
	fmt.Println("Word frequency (more than " + strconv.Itoa(minNumWords) + " times):")
	words := analyzeWords(tweets)

	fmt.Println("Hashtags used (more than " + strconv.Itoa(minNumHashtag) + " times): ")
	hashtags := analyzeHashtags(tweets, words)
	printSortedMapStringInt(hashtags, minNumHashtag)
	fmt.Println("")

	fmt.Println("Interactions with other users (more than " + strconv.Itoa(minNumUserInteractions) + " times): ")
	userInteractions := analyzeUserInteractions(tweets, words)
	printSortedMapStringInt(userInteractions, minNumUserInteractions)
	fmt.Println("")

	analyzeDates(tweets)
	fmt.Println("")
	fmt.Println("Devices:")
	sources := analyzeSource(tweets)
	for k, v := range sources {
		fmt.Print("\x1b[32;1m") //cyan
		fmt.Print(k + ": ")
		fmt.Print("\x1b[0m") //defaultColor
		fmt.Println(strconv.Itoa(v) + "tw	")
	}

	fmt.Println(" ")
	fmt.Print("first tweet analyzed: ")
	fmt.Println(tweets[len(tweets)-1].CreatedAt)
	fmt.Print("last tweet analyzed: ")
	fmt.Println(tweets[0].CreatedAt)

	fmt.Println(" ")
	fmt.Println("Total of " + strconv.Itoa(len(tweets)) + " tweets analyzed")
	fmt.Println(" ")
	fmt.Print("User @")
	c.Cyan(username)
	//analyzeUserAccount(client, username)
	fmt.Println(" analysis finished")
}
