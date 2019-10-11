package main

import (
	"fmt"
	"goftp.io/server"
	"blob_storage_ftp/drivers"
	"blob_storage_ftp/auth"
	"os"
)

func main(){
	fmt.Println("Start FTP Server")

	bucketName := os.Getenv("BUCKET_NAME")

	// Initialize AWS Storage Driver
	driverOpts := &drivers.AWSS3DriverFactory{
		BucketName: bucketName,
	}

	clientId:= os.Getenv("CLIENT_ID")

	// Initialize AWS Cognito Auth Provider
	auth, err := auth.NewCognitoAuth(&auth.CognitoAuthOpts{ClientId: clientId})
	if err != nil {
		fmt.Printf("FTP Error %v\n", err)
	}

	// Configure FTP Server
	opts := &server.ServerOpts{
		Factory: driverOpts,
		Auth: auth,
		Name: "Blob Storage FTP Server",
		WelcomeMessage: "Welcome to Blob Storage FTP Server",
		Hostname: "::",
		Port: 3000,
	}

	server := server.NewServer(opts)

	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("FTP Error %v\n", err)
	}

}