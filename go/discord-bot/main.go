package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	discord "github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"github.com/zhixuan666/discord-bot/handlers"
)

var (
	// Token Discord
	Token string
	// BotName ID
	BotName = "<@!807573034289463316>"
	stopBot = make(chan bool)
	dg      *discord.Session
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("load config file failed: %s\n", err.Error())
	}
	Token = os.Getenv("BOT_TOKEN")
	if Token == "" {
		Token = viper.GetString("token")
	}
	dg, err := discord.New("Bot " + Token)
	if err != nil {
		log.Printf("Create discord bot failed: %v\n", err)
		return
	}

	dg.AddHandler(handlers.MessageCreate)

	err = dg.Open()
	if err != nil {
		log.Printf("Open discord bot failed: %v\n", err)
		return
	}
	defer dg.Close()
	fmt.Println("Discord Bot is running now. Press CTRL-C to exit.")

	mux := http.NewServeMux()
	mux.HandleFunc("/", Root)
	mux.HandleFunc("/Sendmsg", Sendmsg)

	port := os.Getenv("PORT")
	if port == "" {
		port = "12345"
	}

	http.ListenAndServe(":"+port, mux)
}

// Root send message to root
func Root(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("main.html")
		if err != nil {
			log.Printf("root failed: %e", err)
		}
		t.Execute(w, nil)
	} else {
		w.Write([]byte("Error"))
	}
}

// Sendmsg send message to discord
func Sendmsg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")

	var token, msg string
	for k, v := range r.Form {
		if k == "token" {
			token = strings.Join(v, "")
		} else if k == "msg" {
			msg = strings.Join(v, "")
		}

	}
	fmt.Printf("%v %v", token, msg)
	if token != "token" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "token not good"}`))
		return
	}

	sendMessage("807573961378234391", "msg")

	w.Write([]byte(fmt.Sprintf(`{"msg": %v }`, msg)))
}

func sendMessage(channelID string, msg string) {
	_, err := dg.ChannelMessageSend(channelID, msg)

	log.Println(">>> " + msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}
