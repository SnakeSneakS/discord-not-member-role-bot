package members

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/snakesneaks/discord-not-member-role-bot/config"
	"github.com/snakesneaks/discord-not-member-role-bot/core"
)

func FindNotMembers(s *discordgo.Session, i *discordgo.InteractionCreate) ([]*discordgo.Member, error) {
	g, err := s.State.Guild(i.GuildID)
	if err != nil {
		log.Printf("error: %v", err)
		return []*discordgo.Member{}, err
	}

	//log.Print(g.Members)
	member_roles := strings.Split(config.Env.Workspace.MEMBER_ROLES, ",")
	log.Print("member_roles: ", member_roles)

	not_member_members := make([]*discordgo.Member, 0)

	for _, m := range g.Members {
		c := core.CountSameString(m.Roles, member_roles)
		if c == 0 {
			not_member_members = append(not_member_members, m)
			log.Printf("not_member: %s\n", m.User.Username)
		} else {
			log.Printf("member: %s\n", m.User.Username)
		}
	}

	return not_member_members, nil
}
