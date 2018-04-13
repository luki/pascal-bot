// Pascal
// A discord calculator

package main

import (
  "github.com/bwmarrin/discordgo"
  "fmt"
  "strings"

  "os"
  "os/signal"
  "syscall"
)

// WARNING: Remove tokens!!!!

func main() {
  fmt.Printf("Hi, I'm Pascal and am up and running!\n")

  token := "NDM0NDExMjQ3NDIxODgyMzg5.DbKTZQ.vT5N4EHhNnl6EP2oUAROnXqCXOY"

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
  s.UpdateStatus(0, "!psc")
}

func messageEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
  if strings.HasPrefix(m.Content, "!pascal") {

    if strings.HasPrefix(m.Content, "!pascal intro") {
      _, _ = s.ChannelMessageSend(m.ChannelID, "I'm Pascal!")
    }

    _, _ = s.ChannelMessageSend(m.ChannelID, "")
  }

}
