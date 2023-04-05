# HRBot - GoLang library for HighRise Bot API
HRBot is a GoLang library for HighRise Game API that provides an easy way to interact with the game using the official web API provided by the game developers. With this library, you can automate game actions and create custom tools to enhance your developer experience.

This library is built on the official documentation provided by HighRise, which can be found at https://highrise.game/support/web-api. It provides functions for authentication, retrieving game data, and performing actions in the game, such as sending messages, managing contacts, and more.

Installation
To use HRBot, you need to have GoLang installed on your computer. You can download and install GoLang from the official website at https://golang.org/dl/.

To install HRBot, run the following command:

`go get github.com/AltairInglorious/hrbot`

This will download and install the HRBot library and its dependencies.

Usage
To use HRBot in your GoLang project, import the library:

`import "github.com/AltairInglorious/hrbot"`

To use the library, you need to create a new instance of the `Bot` struct. You can do this by calling the `NewBot` function:

```
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	bot "github.com/AltairInglorious/hrbot"
)

func main() {
	bot, err := bot.NewBot(handler)
	if err != nil {
		log.Fatalf("error creating bot: %v", err)
	}
	defer bot.Close()

	log.Printf("🚀 Bot started")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM)
	signal.Notify(signalChan, syscall.SIGINT)

	<-signalChan

	log.Printf("🔽 Bot stopped")
}

func handler(bot *bot.Bot, msg string) {
	log.Printf("📨 Message received: %s", msg)
}
```

The `NewBot` function takes a single argument, which is a function that will be called whenever the bot receives a message. The function takes two arguments: the bot instance and the message received. The function should return a string, which will be sent back to the sender.

To work with the bot, you need set this environment variables:

`BOT_TOKEN` - your bot token

`ROOM_ID` - the ID of the room where the bot will be active


# Contributing
Contributions to HRBot are welcome! If you find a bug, have a feature request, or want to contribute code, please open an issue or pull request on the GitHub repository at https://github.com/AltairInglorious/hrbot.