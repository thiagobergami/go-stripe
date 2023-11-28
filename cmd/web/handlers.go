package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func (app *application) PaymentSucceeded(w http.ResponseWriter, r *http.Request) {

	//handle the parse form
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	cardHolder := r.Form.Get("cardholder_name")
	paymentIntent := r.Form.Get("payment_intent")
	paymentMethod := r.Form.Get("payment_method")
	paymentAmount := r.Form.Get("payment_amount")
	paymentCurrency := r.Form.Get("payment_currency")
	email := r.Form.Get("cardholder_email")

	data := make(map[string]interface{})
	data["cardholder"] = cardHolder
	data["email"] = email
	data["pi"] = paymentIntent
	data["pm"] = paymentMethod
	data["pa"] = paymentAmount
	data["pc"] = paymentCurrency

	if err := app.renderTemplate(w, r, "succeeded", &TemplateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	stringMap := make(map[string]string)
	stringMap["publishable_key"] = app.config.stripe.key

	publicKey := os.Getenv("STRIPE_KEY")

	td := &TemplateData{
		StringMap:        stringMap,
		StripePublickKey: publicKey,
	}

	if err := app.renderTemplate(w, r, "terminal", td); err != nil {
		app.errorLog.Println(err)
	}
}
