package config

import (
	"log"

	env "github.com/Netflix/go-env"
)

type Environment struct {
	App struct {
		REMOVE_COMMANDS bool `env:"REMOVE_COMMANDS"`
	}

	DiscordBot struct {
		TOKEN     string `env:"DISCORD_BOT_TOKEN"`
		CLIENT_ID string `env:"DISCORD_CLIENT_ID"`
		GUILD_ID  string `env:"DISCORD_GUILD_ID"`
	}

	Workspace struct {
		MEMBER_ROLES    string `env:"MEMBER_ROLES"`    // roles which member has. comma separated like "role-id,role_id,role_id,etc"
		NOT_MEMBER_ROLE string `env:"NOT_MEMBER_ROLE"` // not member role id
	}

	Extras env.EnvSet
}

func GetEnvironment() Environment {
	var environment Environment
	es, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
	}
	environment.Extras = es
	//log.Print(es)
	return environment
}

var Env Environment = GetEnvironment()
