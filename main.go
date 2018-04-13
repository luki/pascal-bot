// Pascal
// A discord calculator

package main

import (
  "github.com/joho/godotenv"
  "github.com/bwmarrin/discordgo"

  "fmt"
  "log"
  "strings"

  "os"
  "os/signal"
  "syscall"

  "./logic"
)

func main() {
  err := godotenv.Load()
  if err != nil {
   log.Fatal("Error loading .env file")
  }

  fmt.Printf("Hi, I'm Pascal and am up and running!\n")

  token := os.Getenv("TOKEN")

  sess, err := discordgo.New(fmt.Sprintf("Bot %s", token))
  if err != nil {
    fmt.Printf("Couldn't set up a session.")
    return
  }

  // Handler Setup

  sess.AddHandler(ready)
  sess.AddHandler(messageEvent)

  err = sess.Open()
  if err != nil { fmt.Println("Couldn't open the session") }

  fmt.Println("Pascal is now running.\nPress CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
  <-sc

  sess.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
  s.UpdateStatus(0, "calculation of (simple) equations!")
}

func messageEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
  if strings.HasPrefix(m.Content, "!pascal") {

    if strings.HasPrefix(m.Content, "!pascal intro") {
      _, _ = s.ChannelMessageSend(m.ChannelID, "I'm Pascal; I'll do equations soon!")
      return
    }

    res, err := logic.GetCalculation(strings.TrimSuffix(m.Content, "!pascal"))

    if err != nil {
      s.ChannelMessageSend(m.ChannelID, "Your equation doesn't make sense (to me)!")
      return
    }

    s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%f", res))
  }

}
