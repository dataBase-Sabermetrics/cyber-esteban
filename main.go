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

var discord *discordgo.Session

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
  
  err = sess.Open()
  if err != nil {
    log.Fatal(err)
  }
  return sess
}

func homePage(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Cyber Esteban is running 🚀")
}

func activityMessage(w http.ResponseWriter, r *http.Request) {
  channelID := os.Getenv("CHANNELID")
  if r.Method != http.MethodPost {
    http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
    return
  }
  body, err := io.ReadAll(r.Body)
  if err != nil {
    http.Error(w, "Error reading request body", http.StatusBadRequest)
    return
  }

  message := string(body)
  
  _, err = discord.ChannelMessageSend(channelID, message)
  if err != nil {
    http.Error(w, "Error sending Discord message", http.StatusInternalServerError)
    return
  }
  
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "Received and sent to Discord: %s", message)
}

func handleRequests() {
  http.HandleFunc("/", homePage)
  http.HandleFunc("/api", activityMessage)

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
