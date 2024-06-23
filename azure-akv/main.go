package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"github.com/joho/godotenv"
)

var (
	// resourceGroupName = "rg-empsecure-landingzone-weu-01"
	// vaultName         = "vikivault01"
	secretName = "DBPASSWORD"
	vaultURI   = "https://vikivault01.vault.azure.net/"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	clientID := os.Getenv("AZURE_CLIENT_ID")
	clientSecret := os.Getenv("AZURE_CLIENT_SECRET")
	tenantID := os.Getenv("AZURE_TENANT_ID")

	if subscriptionID == "" || clientID == "" || clientSecret == "" || tenantID == "" {
		log.Fatalf("Environment variables AZURE_SUBSCRIPTION_ID, AZURE_CLIENT_ID, AZURE_CLIENT_SECRET, or AZURE_TENANT_ID are not set")
	}

	/*cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Failed to create client secret credential: %v", err)
	}*/

	cred, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, nil)
	if err != nil {
		log.Fatalf("Failed to create client secret credential: %v", err)
	}

	ctx := context.Background()
	secretClient, err := azsecrets.NewClient(vaultURI, cred, nil)
	if err != nil {
		log.Fatalf("failed to create secret client: %v", err)
	}

	// Get a secret. An empty string version gets the latest version of the secret.
	version := ""
	resp, err := secretClient.GetSecret(ctx, secretName, version, nil)
	if err != nil {
		log.Fatalf("failed to get the secret: %v", err)
	}

	fmt.Printf("secretValue: %s\n", *resp.Value)
}
