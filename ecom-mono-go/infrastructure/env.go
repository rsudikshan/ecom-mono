package infrastructure

import (
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

type Env struct {
	SERVER_PORT 	string
	DB_USER     	string
	DB_PASSWORD     string 
	DB_SERVER_PORT  string
	DB_NAME         string
	EMAIL_PORT		string
	EMAIL_SENDER    string
	EMAIL_SENDER_URL string
	EMAIL_APP_PASSWORD string
	AUTH_TOKEN_KEY_SECRET string
	EMAIL_VERIFICATION_TOKEN_VALIDITY_TIME int
	JWT_EXPIRATION_TIME int
	REFRESH_TOKEN_EXPIRATION_TIME int
	DUMMY_USER_EMAIL string
	DUMMY_VENDOR_EMAIL string
	DUMMY_USERNAME string
	DUMMY_PASSWORD string
	SUPER_ADMIN_EMAIL string
	SUPER_ADMIN_USERNAME string
	SUPER_ADMIN_PASSWORD string
}

func NewEnv() *Env {
	if err:=godotenv.Load(); err!=nil{
		log.Fatal(err.Error())
	}

	evtvt,_ :=  strconv.Atoi(os.Getenv("EMAIL_VERIFICATION_TOKEN_VALIDITY_TIME"))
	jet,_ := strconv.Atoi(os.Getenv("JWT_EXPIRATION_TIME"))
	rtet,_ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRATION_TIME"))

	return &Env{
		SERVER_PORT: os.Getenv("SERVER_PORT"),
		DB_USER:  	 os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_SERVER_PORT: os.Getenv("DB_SERVER_PORT"),
		DB_NAME: 		os.Getenv("DB_NAME"),	
		EMAIL_SENDER:   os.Getenv("EMAIL_SENDER"),
		EMAIL_SENDER_URL: os.Getenv("EMAIL_SENDER_URL"),	
		EMAIL_APP_PASSWORD: os.Getenv("EMAIL_APP_PASSWORD"),
		EMAIL_PORT: os.Getenv("EMAIL_PORT"),
		AUTH_TOKEN_KEY_SECRET: os.Getenv("AUTH_TOKEN_KEY_SECRET"),
		EMAIL_VERIFICATION_TOKEN_VALIDITY_TIME: evtvt,
		JWT_EXPIRATION_TIME: jet,
		REFRESH_TOKEN_EXPIRATION_TIME: rtet,
		DUMMY_USER_EMAIL: os.Getenv("DUMMY_USER_EMAIL"),
		DUMMY_VENDOR_EMAIL: os.Getenv("DUMMY_VENDOR_EMAIL"),
		DUMMY_USERNAME:  os.Getenv("DUMMY_USERNAME"),
		DUMMY_PASSWORD:  os.Getenv("DUMMY_PASSWORD"),
		SUPER_ADMIN_EMAIL: os.Getenv("SUPER_ADMIN_EMAIL"),
		SUPER_ADMIN_USERNAME: os.Getenv("SUPER_ADMIN_USERNAME"),
		SUPER_ADMIN_PASSWORD: os.Getenv("SUPER_ADMIN_PASSWORD"),
	}
}