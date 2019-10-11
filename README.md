# Go Blob Storage FTP Server

AWS, Go, Docker Blob Storage File Server

How do you expose static files in the cloud? Learn the architecture of containerized FTP Server built in GoLang. The FTP server will feature a Docker container, blob storage API access to AWS S3 and Cognito  for OAuth authentication.

# AWS Environment Setup

Run `aws configure` to setup the default credentials

# S3 Bucket

Create an S3 Bucket to store FTP files

# Cognito User Pool

Setup a user pool and app client to authenticate users with. Be sure to use the same region as the default credentials are set to (example: us-east-1).

## Testing via FTP

```bash
ftp -p 127.0.0.1 3000
```

## Testing via Telnet

```bash
telnet 127.0.0.1 3000
USER admin
PASS password
```

# Docs

- [Godocs](https://godoc.org)