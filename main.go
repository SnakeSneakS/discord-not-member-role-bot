package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/snakesneaks/discord-not-member-role-bot/config"
	"github.com/snakesneaks/discord-not-member-role-bot/service/command"
)

func main() {
	s, err := discordgo.New("Bot " + config.Env.DiscordBot.TOKEN)
	if err != nil {
		fmt.Println("error creating Discord session:", err)
		return
	}
	defer s.Close()

	s.Identify.Intents = discordgo.IntentsAll

	err = s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	//add commands
	log.Println("Adding commands...")
	commands := command.NewCommands()
	log.Printf("Commands: \n")

	for _, command := range commands {
		log.Printf(" - %s: %s", command.Name, command.Description)
	}

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, config.Env.DiscordBot.GUILD_ID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	//add handlers
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		commandHandlers := command.NewCommandHandler()
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	//gracefully shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	//remove commands
	if config.Env.App.REMOVE_COMMANDS {
		log.Println("Removing commands...")

		registeredCommands, err := s.ApplicationCommands(s.State.User.ID, config.Env.DiscordBot.GUILD_ID)
		if err != nil {
			log.Fatalf("Could not fetch registered commands: %v", err)
		}

		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, config.Env.DiscordBot.GUILD_ID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Gracefully shutting down.")

}
