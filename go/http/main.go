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
	body, _ := ioutil.ReadAll(r.Body)
	decoded, _ := base64.StdEncoding.DecodeString(r.Header.Get("X-Line-Signature"))
	status := CheckMAC(body, decoded)

	if status {
		Resp(w, 200, "OK")
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

// linepost is rep json
type linepost struct {
	Success    bool   `json:"success"`
	Timestamp  string `json:"timestamp"`
	StatusCode int    `json:"statusCode"`
	Reason     string `json:"reason"`
	Detail     string `json:"detail"`
}

// Resp rsp json msg
func Resp(w http.ResponseWriter, statusCode int, reason string) {
	var rsp linepost

	t := time.Now().Format("2006-01-02 15:04:05")
	if statusCode == 200 {
		rsp = linepost{true, t, statusCode, reason, "200"}
	} else {
		rsp = linepost{false, t, -1, "failed", "-1"}
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
	mux.HandleFunc("/line", sendlinemsg)
	mux.HandleFunc("/hoge", endpoint)

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("config file failed: %s\n", err.Error())
		os.Exit(1)
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
	callbackURL = os.Getenv("CallbackURL")
	http.ListenAndServe(":"+port, mux)

}
