package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "github.com/bwmarrin/discordgo"
    "strings"
)

var token = ""

func main() {
    // Create a new Discord session using the provided bot token.
    dg, err := discordgo.New("Bot " + token)
    if err != nil {
        fmt.Println("Error creating Discord session: ", err)
        return
    }
    //dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent
    dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuildMembers | discordgo.IntentsMessageContent
    // Register the messageCreate func as a callback for the messageCreate events.
    dg.AddHandler(messageCreate)
    
    // Open a websocket connection to Discord and begin listening.
    err = dg.Open()
    if err != nil {
        fmt.Println("Error opening connection: ", err)
        return
    }

    fmt.Println("Bot is running. Press CTRL+C to exit.")

    // Wait for a signal to terminate the bot (CTRL+C).
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc

    // Close the connection when the bot stops.
    dg.Close()
}

// This function will be called every time a new message is created on any channel the bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    // Ignore messages from the bot itself
    if m.Author.ID == s.State.User.ID {
        return
    }

    channel, err := s.Channel(m.ChannelID)
    if err != nil {
        fmt.Println("Error retrieving channel: ", err)
        return
    }

    // Log the received message and its length
    fmt.Printf("Received message: '%s' (Length: %d)\n", m.Content, len(m.Content))

    
    messageContent := strings.TrimSpace(m.Content)
    if channel.Type == discordgo.ChannelTypeDM {
        fmt.Printf("DM Received from %s: '%s' (Length: %d)\n", m.Author.Username, m.Content, len(m.Content))
        return
    }
    fmt.Printf("Trimmed message: '%s' (Length: %d)\n", messageContent, len(messageContent))
    fmt.Printf("Raw message bytes: %v\n", []byte(m.Content))

    if strings.TrimSpace(strings.ToLower(m.Content)) == "ping" {
        s.ChannelMessageSend(m.ChannelID, "Pong!")
    }else if strings.TrimSpace(strings.ToLower(m.Content)) == "hi" {
        s.ChannelMessageSend(m.ChannelID, "Welcome to the channel again")
    } else {
        fmt.Println("Message received, but no reply: ", m.Content)
    }

    
}
