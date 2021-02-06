package handlers

import (
	"fmt"
	"log"
	"time"

	discord "github.com/bwmarrin/discordgo"
)

// MessageCreate handler
func MessageCreate(s *discord.Session, m *discord.MessageCreate) {
	fmt.Printf("%20s %20s %20s > %s\n", m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)

	switch m.Content {
	case "<@&807574022463815701> " + "hello":
		msg := "<@" + m.Author.ID + "> Hello World!"
		sendMessage(s, m.ChannelID, msg)
	}

}

func sendMessage(s *discord.Session, channelID string, msg string) {
	_, err := s.ChannelMessageSend(channelID, msg)

	log.Println(">>> " + msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}
