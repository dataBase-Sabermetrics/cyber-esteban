package main

import (
  "fmt"
  "log"
  "os"
  "io"
  "os/signal"
  "syscall"
  "github.com/bwmarrin/discordgo"
  "github.com/joho/godotenv"
  "net/http"
)

var (
  discord *discordgo.Session
  channelID = os.Getenv("CHANNELID")
)

func startDiscord() *discordgo.Session {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  discordKey := os.Getenv("DISCORD_KEY")
  sess, err := discordgo.New("Bot " + discordKey)
  if err != nil {
    log.Fatal(err)
  }

  sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
  
  sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
    if m.Author.ID == s.State.User.ID {
      return
    }
    if m.Content == "hello" {
      s.ChannelMessageSend(m.ChannelID, "Hello, "+m.Author.Username)
    }
  })

  err = sess.Open()
  if err != nil {
    log.Fatal(err)
  }
  return sess
}

func homePage(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Cyber Esteban is running 🚀")
}

func newGame(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
    return
  }
  body, err := io.ReadAll(r.Body)
  if err != nil {
    http.Error(w, "Error reading request body", http.StatusBadRequest)
    return
  }
  gameMessage := string(body)
  
  _, err = discord.ChannelMessageSend(channelID, gameMessage)
  if err != nil {
    http.Error(w, "Error sending Discord message", http.StatusInternalServerError)
    return
  }
  
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "Received and sent to Discord: %s", gameMessage)
}

func handleRequests() {
  http.HandleFunc("/", homePage)
  http.HandleFunc("/api/game", newGame)

  fmt.Println("Starting HTTP server on :8080...")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Printf("HTTP server error: %v\n", err)
  }
}

func main() {
  discord = startDiscord()
  defer discord.Close()

  fmt.Println("Bot is running")

  go handleRequests()

  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
  <-sc
}
