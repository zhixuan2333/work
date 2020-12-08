package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
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
	channelS string
	channelT string
)

// Root Send msg to root
func Root(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("main.html")
		if err != nil {
			log.Printf("root failed: %e", err)
		}
		t.Execute(w, nil)
	}
}

// sendlinemsg to send msg to user
func sendlinemsg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
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
	err := linesendmsg(userIDs, msg)

	if err {
		Resp(w, 200, "OK")
	} else {
		Resp(w, -1, "failed")
	}

}

// endpoint is webhook
func endpoint(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "POST" {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal("Get resp body failed: ", err)
		}
		decoded, _ := base64.StdEncoding.DecodeString(r.Header.Get("X-Line-Signature"))

		if CheckMAC(body, decoded) {
			Resp(w, 200, "OK")
		} else {
			Resp(w, -1, "Unauthorized access")
		}

	} else {

		t, err := template.ParseFiles("main.html")
		if err != nil {
			log.Printf("root failed: %e", err)
		}
		t.Execute(w, nil)
	}

}

// linsendmsg to send line msg
func linesendmsg(userIDs []string, msg string) bool {
	bot, err := linebot.New(channelS, channelT)

	if err != nil {
		log.Printf("Create a line bot failed: %s\n", err.Error())
		return false
	}

	if _, err := bot.Multicast(userIDs, linebot.NewTextMessage(msg)).Do(); err != nil {
		log.Printf("Send massage failed: %s\n", err.Error())
		return false
	}

	return true

}

// Resp rsp json msg
func Resp(w http.ResponseWriter, statusCode int, reason string) {
	var rsp linepost

	rsp.Success = false

	t := time.Now().Format("2006-01-02 15:04:05")
	if statusCode == 200 {
		rsp.Success = true
		rsp.Timestamp = t
		rsp.StatusCode = statusCode
		rsp.Reason = reason
		rsp.Detail = "200"
	} else {

		rsp.Success = false
		rsp.Timestamp = t
		rsp.StatusCode = statusCode
		rsp.Reason = reason
		rsp.Detail = "-1"
	}

	js, err := json.Marshal(rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

// CheckMAC to check hash
func CheckMAC(message, messageMAC []byte) bool {
	mac := hmac.New(sha256.New, []byte(channelS))
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Root)
	mux.HandleFunc("/line", sendlinemsg)
	mux.HandleFunc("/hoge", endpoint)

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("load config file failed: %s\n", err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "12345"
	}

	channelS = os.Getenv("channelS")
	if channelS == "" {
		channelS = viper.GetString("channel.secret")
	}

	channelT = os.Getenv("channelT")
	if channelT == "" {
		channelT = viper.GetString("channel.token")
	}

	http.ListenAndServe(":"+port, mux)

}
