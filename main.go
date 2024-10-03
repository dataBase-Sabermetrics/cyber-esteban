package main

import (
	"fmt"
	"log"
	"os"
  "os/signal"
  "syscall"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	discordKey := os.Getenv("DISCORD_KEY")

  sess, err := discordgo.New("Bot "+discordKey)
  if err != nil {
    log.Fatal(err)
  }

  sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
    if m.Author.ID == s.State.User.ID {
      return
    }

    if m.Content == "hello" {
      s.ChannelMessageSend(m.ChannelID, "Hello, "+m.Author.Username)
    }
  })

  sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

  err = sess.Open()
  if err != nil {
    log.Fatal(err)
  }
  defer sess.Close()

  fmt.Println("Bot is running")

  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
  <-sc

}
