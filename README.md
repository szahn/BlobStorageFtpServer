# Go Blob Storage FTP Server

AWS, Go, Docker Blob Storage File Server

How do you expose static files in the cloud? Learn the architecture of containerized FTP Server built in GoLang. The FTP server will feature a Docker container, blob storage API access to AWS S3 and Cognito  for OAuth authentication.

# Why FTP?

FTP is considered a legacy, insecure protocol. [SFTP](https://www.digitalocean.com/community/tutorials/how-to-use-sftp-to-securely-transfer-files-with-a-remote-server) supersedes FTP. However, it is the easiest way to provide AWS blob storage access to the public without requiring additional dependencies. Most organizations know FTP pretty well.

# Why Blob Storage?

Secure, scaleable, store any kind of data, does not require (NFS) file system knowledge.

# Architecture

- [AWS S3 Bucket](https://s3.console.aws.amazon.com/s3/home?region=us-east-1): hosts files in blob storage
- [AWS Cognito](https://console.aws.amazon.com/cognito/users/?region=us-east-1#/?_k=av6nyb): IAM - identity access management, OAuth, user pools
- [Go](https://golang.org/): Server-side, compiled language
- [Docker](https://docs.docker.com/engine/reference/builder/): Container technology providing namespace isolation

# Dependencies
- [AWS Command Line](https://aws.amazon.com/cli/)

# Packages

- [Go FTP Server](https://goftp.io)
- [AWS S3 Go Lib](https://godoc.org/github.com/aws/aws-sdk-go/service/s3)
- [AWS Cognito Lib](https://godoc.org/github.com/aws/aws-sdk-go/service/cognitoidentityprovider)


# AWS Environment Setup

Run `aws configure` to setup the default credentials

# S3 Bucket

Create an S3 Bucket to store FTP files

# Cognito User Pool

Setup a user pool and app client to authenticate users with. Be sure to use the same region as the default credentials are set to (example: us-east-1).

## Testing via FTP

```bash
ftp -p 127.0.0.1 3000
ls
get c.pdf
cp c.pdf d.pdf
put d.pdf
ls
```

## Testing via Telnet

```bash
telnet 127.0.0.1 3000
USER admin
PASS password
```

# Docs

- [Godocs](https://godoc.org)


# Learning Go

- [Go Tour](https://tour.golang.org/welcome/1)
- [Go by Example](https://gobyexample.com/)
- [Go OCR server example](https://github.com/szahn/OCR-Engine)