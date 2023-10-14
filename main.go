package main

import (
	"fmt"
	"go-seu-astral-bot/commands"
	"go-seu-astral-bot/sounds"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	buffer            = make([][]byte, 0)
	DISCORD_BOT_TOKEN string
	configs           Configs
)

const (
	CMDStart = "!sa start"
	CMDCount = "!sa count"
)

func main() {
	configs = LoadConfigs()

	sess, err := discordgo.New("Bot " + configs["DISCORD_BOT_TOKEN"])
	if err != nil {
		fmt.Println("Error creating Discord Session: ", err)
		return
	}

	err = sounds.Load(&buffer)
	if err != nil {
		fmt.Println("Error loading sound: ", err)
		return
	}

	sess.AddHandler(handleCommands)

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer sess.Close()

	fmt.Println("O bot para tocar Seu Astral em loop foi inciado!")

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func handleCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.SessionID {
		return
	}

	if strings.HasPrefix(m.Content, CMDStart) {
		commands.Start(s, m, buffer)
	}
}
