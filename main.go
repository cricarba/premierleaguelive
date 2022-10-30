package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	var w1 string
	fmt.Println("Match Id: ")
	fmt.Scanln(&w1)
	fmt.Println("Init:.. ")
	teller := time.Tick(30 * time.Second)
	finish := time.After(3 * time.Hour)
	for {
		select {
		case <-teller:
			postTweet(w1)
		case <-finish:
			fmt.Println("Finish!")
			return
		}
	}

}

func postTweet(fixture string) {

	creds := Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}
	//fmt.Printf("%+v\n", creds)

	clientTwitter, err := getClient(&creds)
	//fmt.Printf("%+v\n", clientTwitter)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}
	apiUrl := os.Getenv("URL_API_PL")
	origin := os.Getenv("ORIGIN_SECRET")
	fmt.Println(apiUrl)
	fmt.Println(origin)
	url := apiUrl + fixture + "/textstream/EN?pageSize=3&sort=desc"
	fmt.Println(url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("origin", origin)
	fmt.Println("Fecth Data:.. ")
	response, _ := client.Do(req)

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))

	var match Match
	json.Unmarshal(responseData, &match)

	hashTag := "#" + match.Fixture.Teams[0].Team.Club.Abbr + match.Fixture.Teams[1].Team.Club.Abbr
	score := strconv.FormatInt(match.Fixture.Teams[0].Score, 10) + "-" + strconv.FormatInt(match.Fixture.Teams[1].Score, 10)

	//tweet := "{hashTag.Text} \n/n⚽ {teamHome.Text} {score.Text} {teamAway.Text} \n/n🕕 {timeMatch}'  \n/n🎙️ {message} \n/n#PremierLeague #PL"
	//tweetTemplate := "%s \n \n⚽ %s %s %s \n \n🕕 %s  \n/n🎙️ %s \n \n#PremierLeague #PL"
	for i := 2; i >= 0; i-- {

		tweet := hashTag + " \n \n ⚽ " + match.Fixture.Teams[0].Team.ShortName + " " + score + " " + match.Fixture.Teams[1].Team.ShortName + " \n\n🕕 " + match.Events.Content[i].Time.Label + "' \n\n🎙️ " + match.Events.Content[i].Text + " \n\n #PremierLeague #PL"
		_, _, err := clientTwitter.Statuses.Update(tweet, nil)
		if err != nil {
			log.Println(err)
		}
		log.Printf("%+v/n", tweet)
	}
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"/n", values...)
}
