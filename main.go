package main

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	errRequestFailed = errors.New("request failed")
)

type requestResult struct {
	url string
	status string
}

func main() {
	// Code about commit #1.1 ~ #1.3

	// account := accounts.CreateAccount("Jay")
	// account.Deposit(1000)
	// account.Deposit(3000)
	// account.Withdraw(2000)
	// err := account.Withdraw(5000)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(account.CheckBalance())
	// account.ChangeOwner("Hwang")
	// fmt.Println(account.CheckOwner())

	// Code about commit #1.4 ~ #1.5

	// dictionary := dictionarys.Dictionary{}
	// exword := "hello"
	// dictionary.Add(exword, "greeting")
	// result := dictionary.Update("xdd", "first")
	// if result != nil {
	// 	fmt.Println(result)
	// }
	// word, _ := dictionary.Search(exword)
	// fmt.Println(word)
	// fmt.Println(dictionary)
	// result2 := dictionary.Delete("second")
	// if result2 != nil {
	// 	fmt.Println(result2)
	// }
	// fmt.Println(dictionary)
	// }


	// Code about commit #2.1 ~ #2.2
	
	results := make(map[string]string)
	c := make(chan requestResult)

	urls := [] string {
		"https://www.google.com/",
		"https://www.youtube.com/",
		"https://www.airbnb.com/",
		"https://www.amazon.com/",
		"https://www.instagram.com/",
		"https://www.soundcloud.com/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i:=0; i<len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL (url string, c chan<- requestResult) {
	// fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		c <- requestResult{url: url, status: "FAILED"}
	} else {
		c <- requestResult{url: url, status: "OK"}
	}
}
