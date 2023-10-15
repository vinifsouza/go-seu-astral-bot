package main

import (
	"fmt"
	"go-seu-astral-bot/commands"
	"go-seu-astral-bot/sounds"
	"go-seu-astral-bot/translate"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	buffer     = make([][]byte, 0)
	configs    Configs
	loopCount  = 0
	_translate translate.Translate
)

const (
	CMDStart = "!sa start"
	CMDCount = "!sa count"
)

func main() {
	configs = LoadConfigs()
	_translate = translate.Load(configs["APP_LANGUAGE"])

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

	fmt.Println(_translate["START_MESSAGE"])

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func handleCommands(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.SessionID {
		return
	}

	if strings.HasPrefix(message.Content, CMDStart) {
		commands.Start(session, message, buffer, &loopCount)
	}

	if strings.HasPrefix(message.Content, CMDCount) {
		commands.Count(session, message, &loopCount)
	}
}
