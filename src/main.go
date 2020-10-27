package main

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"myTest/src/bot"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	env := os.Getenv("ENV")

	if env == "dev" || env == "" {
		_ = godotenv.Load(".env")
	}

	println(env)

	token := os.Getenv("TOKEN_DISCORD")

	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		panic(errors.New("Erreur dans la création du bot"))
	}

	u, err := dg.User("@me")

	if err != nil {
		panic(errors.New("Erreur dans la recuperation de l'utilisateur"))
	}

	_ = os.Setenv("BOT_ID", u.ID)

	err = dg.Open()

	dg.AddHandler(bot.Commande)

	if err != nil {
		panic(errors.New("Erreur dans Connection du bot"))
	}

	fmt.Println("Bot ", u.ID, "is running !")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	_ = dg.Close()
	fmt.Println("Bot ", u.ID, "is stoping !")

}
