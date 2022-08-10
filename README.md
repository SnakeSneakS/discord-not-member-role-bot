# discord-not-member-role-bot
- list-up and add role to "not-member" who doesn't have MEMBER_ROLES. 
- 名前が悪すぎるけど、DiscordでMEMBER_ROLES(ロール1、ロール2、ロール3、…)のいずれも持っていないユーザーを列挙してロール「NOT_MEMBER_ROLE」を付与するだけの簡単なbotです

1. access [here](https://discord.com/developers/applications) and create bot. 
2. create & fill in .env file. Please refer to .env.template file.
    1. to get DISCORD_GUILD_ID, ROLE_ID, etc, turn on Discord > UserSettings > Details > DeveloperMode.
3. invite bot to your discord server
    - permissions: 268437504 
        - Manage Roles
        - Send Messages
        - 
4. ```go run main.go```
5. run commands in your discord server
   - find_not_members: list up not_member members
   - add_role_to_not_members: add role to not_member members




