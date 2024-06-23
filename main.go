package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	subscriptionID    string
	TenantID          string
	clientID          string
	clientSecret      string
	location          = "westeurope"
	resourceGroupName = "rg-empsecure-landingzone-weu-01"
	vaultName         = "vikivault01"
	secretName        = "vikisecret"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	subscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")
	TenantID = os.Getenv("AZURE_TENANT_ID")
	clientID = os.Getenv("AZURE_CLIENT_ID")
	clientSecret = os.Getenv("AZURE_CLIENT_SECRET")

	if len(subscriptionID) == 0 || len(TenantID) == 0 || len(clientID) == 0 || len(clientSecret) == 0 {
		log.Fatal("AZURE_SUBSCRIPTION_ID or any of related env vars are not set.")
	}

}
