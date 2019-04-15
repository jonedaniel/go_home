package main

import (
	"bufio"
	"encoding/json"
	"github.com/bitly/go-nsq"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/joeshaw/envdecode"
	"github.com/matryer/go-oauth/oauth"
	"gopkg.in/mgo.v2"
)

var (
	authClient *oauth.Client
	creds      *oauth.Credentials
)

type poll struct {
	Options []string
}
type tweet struct {
	Text string
}

func main() {
	var ts struct {
		ConsumerKey    string `env:"SP_TWITTER_KEY,required"`
		ConsumerSecret string `env:"SP_TWITTER_SECRET,required"`
		AccessToken    string `env:"SP_TWITTER_ACCESSTOKEN,required"`
		AccessSecret   string `env:"SP_TWITTER_ACCESSSECRET,required"`
		FlexibleHost   string `env:"SP_FLEXIBLE_HOST,required"`
	}
	if err := envdecode.Decode(&ts); err != nil {
		log.Fatalln(err)
	}
	client := &http.Client{
		Transport: &http.Transport{
			Dial: Dial,
		},
	}
	creds = &oauth.Credentials{
		Token:  ts.AccessToken,
		Secret: ts.AccessSecret,
	}
	authClient = &oauth.Client{
		Credentials: oauth.Credentials{
			Token:  ts.ConsumerKey,
			Secret: ts.ConsumerSecret,
		},
	}
	twitterStopChan := make(chan struct{}, 1)
	publisherStopChan := make(chan struct{}, 1)
	stop := false
	signalChan := make(chan os.Signal, 1)
	go func() {
		<-signalChan
		stop = true
		log.Println("Stopping...")
		closeConn()
	}()
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	votes := make(chan string) // chan for votes
	go func() {
		pub, err := nsq.NewProducer(ts.FlexibleHost+":4150", nsq.NewConfig())
		if err != nil {
			log.Fatalln("connect nsq error:", err)
		}
		for vote := range votes {
			pub.Publish("votes", []byte(vote)) // publish vote
		}
		log.Println("Publisher: Stopping")
		pub.Stop()
		log.Println("Publisher: Stopped")
		publisherStopChan <- struct{}{}
	}()
	go func() {
		defer func() {
			twitterStopChan <- struct{}{}
		}()
		for {
			if stop {
				log.Println("Twitter: Stopped")
				return
			}
			time.Sleep(2 * time.Second) // calm
			var options []string
			db, err := mgo.Dial(ts.FlexibleHost)
			if err != nil {
				log.Fatalln("connect mongo error:", err)
			}
			iter := db.DB("ballots").C("polls").Find(nil).Iter()
			var p poll
			for iter.Next(&p) {
				options = append(options, p.Options...)
			}
			iter.Close()
			db.Close()

			log.Println("ballots polls: ", options)
			hashtags := make([]string, len(options))
			for i := range options {
				hashtags[i] = "#" + strings.ToLower(options[i])
			}

			form := url.Values{"track": {strings.Join(hashtags, ",")}}
			formEnc := form.Encode()

			u, _ := url.Parse("https://stream.twitter.com/1.1/statuses/filter.json")
			req, err := http.NewRequest("POST", u.String(), strings.NewReader(formEnc))
			if err != nil {
				log.Println("creating filter request failed:", err)
			}
			req.Header.Set("Authorization", authClient.AuthorizationHeader(creds, "POST", u, form))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Set("Content-Length", strconv.Itoa(len(formEnc)))

			resp, err := client.Do(req)
			if err != nil {
				log.Println("Error getting response:", err)
				continue
			}
			if resp.StatusCode != http.StatusOK {
				// this is a nice way to see what the error actually is:
				s := bufio.NewScanner(resp.Body)
				s.Scan()
				log.Println(s.Text())
				log.Println(hashtags)
				log.Println("StatusCode =", resp.StatusCode)
				continue
			}

			reader = resp.Body
			decoder := json.NewDecoder(reader)
			for {
				var t tweet
				if err := decoder.Decode(&t); err == nil {
					for _, option := range options {
						if strings.Contains(
							strings.ToLower(t.Text),
							strings.ToLower(option),
						) {
							log.Println("vote:", option)
							votes <- option
						}
					}
				} else {
					break
				}
			}

		}

	}()

	// update by forcing the connection to close
	go func() {
		for {
			time.Sleep(1 * time.Minute)
			closeConn()
			if stop {
				break
			}
		}
	}()

	<-twitterStopChan // important to avoid panic
	close(votes)
	<-publisherStopChan

}
