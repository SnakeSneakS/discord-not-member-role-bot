package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/snakesneaks/discord-not-member-role-bot/config"
	"github.com/snakesneaks/discord-not-member-role-bot/service/command/members"
	"github.com/snakesneaks/discord-not-member-role-bot/service/response"
)

func NewCommandHandler() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"find_not_members":        notMembersDisplayCommand,
		"add_role_to_not_members": notMembersAddRoleCommand,
	}
	return commandHandlers
}

func NewCommands() []*discordgo.ApplicationCommand {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "find_not_members",
			Description: "find_not_members",
		},
		{
			Name:        "add_role_to_not_members",
			Description: "add_role_to_not_members",
		},
	}
	return commands
}

//notMembersDisplayCommand - 会員じゃない人を取得・表示
func notMembersDisplayCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	content := "listing up not member members...\n"
	response.Create(s, i.Interaction, content)

	not_member_members, err := members.FindNotMembers(s, i)
	if err != nil {
		content += "error: " + err.Error() + "\n"
		content += "stopped!!\n"
		response.Edit(s, i.Interaction, content)
		return
	}

	for _, m := range not_member_members {
		//content += m.User.Mention()
		content += m.User.Username + ", "
	}
	response.Edit(s, i.Interaction, content)
}

func notMembersAddRoleCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	content := "adding role to not members...\n"
	response.Create(s, i.Interaction, content)

	not_member_members, err := members.FindNotMembers(s, i)
	if err != nil {
		content += "error: " + err.Error() + "\n"
		content += "stopped!!\n"
		response.Edit(s, i.Interaction, content)
		return
	}

	err = nil
	for _, m := range not_member_members {
		err = s.GuildMemberRoleAdd(i.GuildID, m.User.ID, config.Env.Workspace.NOT_MEMBER_ROLE)
		if err != nil {
			content += "error: " + err.Error() + "\n"
			content += "stopped!!\n"
			response.Edit(s, i.Interaction, content)
			return
		}
	}
	content += "success!!\n"
	response.Edit(s, i.Interaction, content)
}
