package main

import (
	"os"

	"github.com/Emyrk/fortnitediscord/gobot"
)

func main() {
	token := os.Getenv("DISCORD_TOKEN")
	bot := gobot.NewBot(token, []string{"Emyrks", "LopDropFlop"})
	bot.Connect()
	bot.Run()
}
