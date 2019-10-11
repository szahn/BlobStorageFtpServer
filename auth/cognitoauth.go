package auth

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CognitoAuthOpts struct {
	ClientId string
}

type CognitoAuth struct {
	clientId string
	session *session.Session
}

func NewCognitoAuth(opts *CognitoAuthOpts) (*CognitoAuth, error) {

	session, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})

	if err != nil {
		fmt.Println("Error creating session", err)
		return nil, err
	}

	fmt.Printf("Using client %s\n", opts.ClientId)

	return &CognitoAuth{
		clientId: opts.ClientId,
		session: session,
	}, nil;
}

func (auth *CognitoAuth) CheckPasswd(username string, password string) (bool, error) {
	fmt.Println("AWS CognitoAuth CheckPasswd")


	params := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		ClientId: aws.String(auth.clientId),
		AuthParameters: map[string]*string{
		  "USERNAME": aws.String(username),
		  "PASSWORD": aws.String(password),
		},
	  }

	cip := cognitoidentityprovider.New(auth.session)
	_, err := cip.InitiateAuth(params)

	if err != nil {
		fmt.Printf("Error authenticating: %v\n", err)
		return false, nil
	}

	return true, nil
 }