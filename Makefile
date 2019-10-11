AWS_PATH := $(HOME)

init:
	@go get goftp.io/server
	@go get -u github.com/aws/aws-sdk-go/...

docker_build:
	@docker build . -t goftp

# Run docker, mounting the aws credentials folder, and passing in cognito client id

docker_run:
	@docker run -p 3000:3000 \
		-e "CLIENT_ID=btr36m7q3h76m1cirjjkh408h" \
		-e "BUCKET_NAME=zahnsoftware-go-ftp-host" \
		-v $(AWS_PATH)/.aws:/root/.aws \
		goftp
