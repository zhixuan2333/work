package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/spf13/viper"
)

var (
	channelS    string
	channelT    string
	callbackURL string
	token       string
)

func sendlinemsg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	var userIDs []string
	var msg string
	for k, v := range r.Form {
		if k == "uid" {
			u := strings.Join(v, "")
			userIDs = strings.Split(u, ",")
		} else if k == "msg" {
			msg = strings.Join(v, "")
		}

	}
	fmt.Println(userIDs)
	fmt.Println(msg)
	linesendmsg(userIDs, msg)
}

func linesendmsg(userIDs []string, msg string) {
	channelS, channelT := lineconf()
	bot, err := linebot.New(channelS, channelT)
	if err != nil {
		log.Printf("Create a line bot failed: %s\n", err.Error())
	}
	if _, err := bot.Multicast(userIDs, linebot.NewTextMessage(msg)).Do(); err != nil {
		log.Printf("Send massage failed: %s\n", err.Error())
	}
}

func lineconf() (string, string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("config file failed: %s\n", err.Error())
		os.Exit(1)
	}
	channelS := viper.GetString("channel.secret")
	channelT := viper.GetString("channel.token")
	return channelS, channelT
}

type linepost struct {
	Success    bool   `json:"success"`
	Timestamp  string `json:"timestamp"`
	StatusCode int    `json:"statusCode"`
	Reason     string `json:"reason"`
	Detail     string `json:"detail"`
}

func endpoint(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	fmt.Println(w, "Header全部数据:", header)
	t := time.Now().Format("2006-01-02 15:04:05")
	rsp := linepost{true, t, 200, "OK", "200"}
	js, err := json.Marshal(rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	body, _ := ioutil.ReadAll(r.Body)
	decoded, _ := base64.StdEncoding.DecodeString(r.Header.Get("X-Line-Signature"))
	status := CheckMAC(body, decoded)
	fmt.Println(body)
	fmt.Println(decoded)
	fmt.Println(status)
}

func CheckMAC(message, messageMAC []byte) bool {
	mac := hmac.New(sha256.New, []byte(channelS))
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/line", sendlinemsg)
	mux.HandleFunc("/hoge", endpoint)
	port := os.Getenv("PORT")
	if port == "" {
		port = "12345"
	}
	channelS = os.Getenv("ClientID")
	channelT = os.Getenv("ClientSecret")
	callbackURL = os.Getenv("CallbackURL")
	http.ListenAndServe(":"+port, mux)
}
