package main

import (
	"golang.org/x/oauth2/google"
	"flag"
	"context"
	"io/ioutil"
)

var (
	ServiceAccount string
	Scope string
)

func main () {
	flag.StringVar(&ServiceAccount, "service_account", "./service_account.json", "Path to service account")
	flag.StringVar(&Scope, "scope", "https://www.googleapis.com/auth/dialogflow", "Scope")
	flag.Parse()

	data, err := ioutil.ReadFile(ServiceAccount)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	creds, err := google.CredentialsFromJSON(ctx, data, Scope)
	if err != nil {
		panic(err)
	}
	
	token, err := creds.TokenSource.Token()
	if err != nil {
		panic(err)
	}

	println("Service Account: " + ServiceAccount)
	println("Access Token: " + token.AccessToken)
	println("Token Type: " + token.TokenType)
	println("Refresh Token: " + token.RefreshToken)
	println("Expiry: " + token.Expiry.String())
}