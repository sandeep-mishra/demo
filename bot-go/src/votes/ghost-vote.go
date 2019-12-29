package votes

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"util"
)

func CastGhostVote() {
	fmt.Println("casting vote")
	var urlStr string = vote_base_url + "/castVote"
	quoteId := strconv.Itoa(rand.Intn(16))
	strconv.Itoa(97)
	formData := url.Values{
		"quoteId": {quoteId},
	}
	makePostRequest(urlStr, formData)
}

func ListQuotes() {
	fmt.Println("listing quotes")
	var urlStr string = quote_base_url + "/list"
	makeGetRequest(urlStr)

}

func TallyVotes() {
	fmt.Println("tally votes'")
	var urlStr string = vote_base_url + "/tallyVote"

	makeGetRequest(urlStr)

}

func makeGetRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

func makePostRequest(url string, formData url.Values) {

	resp, err := http.PostForm(url, formData)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))

}

var vote_base_url string
var quote_base_url string

func init() {
	vote_base_url = "http://" + util.Cfg.Frontend.Host + ":" + util.Cfg.Frontend.Port + "/api/vote"
	quote_base_url = "http://" + util.Cfg.Frontend.Host + ":" + util.Cfg.Frontend.Port + "/api/quote"
}