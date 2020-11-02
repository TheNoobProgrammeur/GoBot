package commands

import (
	"github.com/bwmarrin/discordgo"
	"myTest/src/models"
	"myTest/src/services"
)

func Ping(user discordgo.User) string {
	updateUser(user.ID, user.Username)
	return "Pong"
}

func updateUser(idUser string, username string)   {
	dbservice := services.New()
	user := models.User{}

	dbservice.GetDB().Where("discord_number = ?",idUser).First(&user)
	if user.Name == "" {
		user = models.User{
			DiscordNumber: idUser,
			Name: username,

		}
	}
	ping := models.Ping{}
	user.PingList = append(user.PingList,ping )
	dbservice.GetDB().Save(&user)

}