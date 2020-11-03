package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/spf13/viper"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	io.WriteString(w, "Hello, world!\n")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Path)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()

		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		io.WriteString(w, "helloworld")
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello World!") //这个写入到w的是输出到客户端的
}

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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/", echoHandler)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/c", sayhelloName)
	mux.HandleFunc("/line", sendlinemsg)
	port := os.Getenv("PORT")
	if port == "" {
		port = "12345"
	}
	http.ListenAndServe(":"+port, mux)

}
