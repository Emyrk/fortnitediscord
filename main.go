package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Emyrk/fortnitediscord/gobot"
)

func main() {
	token := os.Getenv("DISCORD_TOKEN")
	bot := gobot.NewBot(token, []string{"Emyrks", "LopDropFlop", "GuyWhoDoesThings", "r0bd0g364", "ScaRe TacticS23", "Bad Assassins_YT"})
	bot.Connect()

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		bot.Close()
		os.Exit(1)
	}()

	bot.Run()
}
