package bot

import (
	"github.com/bwmarrin/discordgo"
	c "myTest/src/bot/commands"
	"os"
	"regexp"
	"strconv"
)

func Commande(s *discordgo.Session, m *discordgo.MessageCreate) {

	listeCommande := []string{
		"ping",
		"CommandeListe",
		"Fibonnacci",
		"NewChanelText",
		"NewChanelVoice"}

	bID := os.Getenv("TOKEN_DISCORD")
	if m.Author.ID == bID {
		return
	}

	regxForComplexCommande := regexp.MustCompile("^(!\\S+) (.*)$")
	regxForSimpleCommande := regexp.MustCompile("^!\\S+$")

	if regxForSimpleCommande.MatchString(m.Content) {
		com := m.Content[1:]
		switch com {
		case "ping":
			s.ChannelMessageSend(m.ChannelID, c.Ping())
		case "CommandeListe":
			liste := ""
			for index, command := range listeCommande {
				liste += "[" + strconv.Itoa(index) + "] " + command + "\n"
			}
			s.ChannelMessageSend(m.ChannelID, "liste (prefix = !) : \n"+liste)
		}
	}

	if regxForComplexCommande.MatchString(m.Content) {
		groups := regxForComplexCommande.FindStringSubmatch(m.Content)
		com := groups[1][1:]
		param := groups[2]
		switch com {
		case "Fibonnacci":
			number, err := strconv.Atoi(param)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, ":boom: Il faut entre un nombre pour la commande Fibonnacci")
				return
			}
			s.ChannelMessageSend(m.ChannelID, strconv.Itoa(c.Fibonacci(number)))
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

}
