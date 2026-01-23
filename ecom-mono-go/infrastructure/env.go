package infrastructure

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	SERVER_PORT string
}

func NewEnv() *Env {
	if err:=godotenv.Load(); err!=nil{
		log.Fatal(err.Error())
	}

	return &Env{
		SERVER_PORT: os.Getenv("SERVER_PORT"),
	}
}