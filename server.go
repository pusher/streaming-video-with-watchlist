package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pusher/pusher-http-go/v5"
)

var teams = map[string][]string{
	"group-Dave": {"Dave", "Lisa", "Phil", "Britney", "Phil"},
	"group-Lora": {"Lora", "Frank", "Peter", "Janine"},
}

type user struct {
	Username string `json:"username" xml:"username" form:"username" query:"username"`
	Email    string `json:"email" xml:"email" form:"email" query:"email"`
}

var loggedInUser user
var client = pusher.Client{}

func isUserLoggedIn(rw http.ResponseWriter, req *http.Request) {
	if loggedInUser.Username != "" && loggedInUser.Email != "" {
		json.NewEncoder(rw).Encode(loggedInUser)
	} else {
		json.NewEncoder(rw).Encode("false")
	}
}

func NewUser(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &loggedInUser)

	if err != nil {
		panic(err)
	}

	json.NewEncoder(rw).Encode(loggedInUser)
}

func pusherUserAuth(res http.ResponseWriter, req *http.Request) {
	params, _ := ioutil.ReadAll(req.Body)

	userData := map[string]interface{}{
		"id":        loggedInUser.Username,
		"email":     loggedInUser.Email,
		"watchlist": getUserWatchlist(loggedInUser.Username),
	}

	response, err := client.AuthenticateUser(params, userData)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(res, string(response))
}

func getUserWatchlist(userName string) []string {
	for key := range teams {
		for _, v := range teams[key] {
			if v == userName {
				return teams[key]
			}
		}
	}
	return []string{}
}

func main() {
	err := godotenv.Load(`.env`)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client = pusher.Client{
		AppID:   os.Getenv("PUSHER_APP_ID"),
		Key:     os.Getenv("PUSHER_APP_KEY"),
		Secret:  os.Getenv("PUSHER_APP_SECRET"),
		Cluster: os.Getenv("PUSHER_APP_CLUSTER"),
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/isLoggedIn", isUserLoggedIn)
	http.HandleFunc("/new/user", NewUser)
	http.HandleFunc("/pusher/user-auth", pusherUserAuth)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
