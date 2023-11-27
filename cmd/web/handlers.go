package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	publicKey := os.Getenv("STRIPE_KEY")

	td := &TemplateData{
		StripePublickKey: publicKey,
	}

	if err := app.renderTemplate(w, r, "terminal", td); err != nil {
		app.errorLog.Println(err)
	}
}
