package bot

import (
	"github.com/bwmarrin/discordgo"
	c "myTest/src/bot/commands"
	"os"
	"regexp"
	"strconv"
)

func Commande(s *discordgo.Session, m *discordgo.MessageCreate) {

	bID := os.Getenv("TOKEN_DISCORD")
	if m.Author.ID == bID {
		return
	}

	regxForComplexCommand := regexp.MustCompile("^(!\\S+) (.*)$")
	regxForSimpleCommand := regexp.MustCompile("^!\\S+$")

	if regxForSimpleCommand.MatchString(m.Content) {
		SimpleCommand(s,m)
	}

	if regxForComplexCommand.MatchString(m.Content) {
		ComplexCommand(s,m,regxForComplexCommand)
	}
}
func SimpleCommand(s *discordgo.Session, m *discordgo.MessageCreate)  {

	listCommand := []string{
		"ping",
		"listCommand",
		"Fibonacci",
		"NewChanelText",
		"NewChanelVoice"}

	com := m.Content[1:]
	switch com {
	case "listCommand":
		list := ""
		for index, command := range listCommand {
			list += "[" + strconv.Itoa(index) + "] " + command + "\n"
		}
		s.ChannelMessageSend(m.ChannelID, "lite (prefix = !) : \n"+list)
	case "ping":
		s.ChannelMessageSend(m.ChannelID, c.Ping(*m.Author))
	}
}
func ComplexCommand(s *discordgo.Session, m *discordgo.MessageCreate,regxForComplexCommand *regexp.Regexp)  {
	groups := regxForComplexCommand.FindStringSubmatch(m.Content)
	com := groups[1][1:]
	param := groups[2]
	switch com {
	case "Fibonacci":
		number, err := strconv.Atoi(param)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, ":boom: Il faut entre un nombre pour la commande Fibonacci")
			return
		}

		number64 := uint64 (number)
		if number64 >= 93 {
			s.ChannelMessageSend(m.ChannelID, "Desolé c'est trop gros .... plaigné vous a postgresql ")
			return
		}

		s.ChannelMessageSend(m.ChannelID, strconv.FormatUint(c.Fibonacci(number64), 10))
	case "NewChanelText":
		s.ChannelMessageSend(m.ChannelID, "Creation du chanel textuel : "+param)
		_, err := s.GuildChannelCreate(m.GuildID, param, 0)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Erreur dans la création du chanel")
			return
		}
		s.ChannelMessageSend(m.ChannelID, "Chanel créé")
	case "NewChanelVoice":
		s.ChannelMessageSend(m.ChannelID, "Creation du chanel vocal : "+param)
		_, err := s.GuildChannelCreate(m.GuildID, param, 2)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Erreur dans la création du chanel")
			return
		}
		s.ChannelMessageSend(m.ChannelID, "Chanel créé")
	case "Joke":
		if param == "-h" {
			s.ChannelMessageSend(m.ChannelID, "Liste des theme de blague : Any|Programming|Miscellaneous|Darck|Pun")
			return
		}
		jocke, err := c.Jocke(param)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}
		s.ChannelMessageSend(m.ChannelID, "My jock : \n"+jocke)
	}
}