package main

import (
	"embed"
	"os"

	"github.com/artamananda/artanymous/app/controller"
	"github.com/artamananda/artanymous/app/model"
	"github.com/artamananda/artanymous/app/repository"
	"github.com/artamananda/artanymous/config"
	_ "github.com/jackc/pgx"
)

//go:embed app/view/*
var Resources embed.FS

func main() {
	os.Setenv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/test_db_arta")
	db := config.NewDB()
	conn, err := db.Connect()
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&model.Message{})
	messageHandle := repository.NewMessageRepo(conn)

	mainAPI := controller.NewAPI(messageHandle, Resources)
	mainAPI.Start()
}
