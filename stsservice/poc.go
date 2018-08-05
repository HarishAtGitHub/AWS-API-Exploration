package main

import (
"github.com/aws/aws-sdk-go/aws/session"
"github.com/aws/aws-sdk-go/service/sts"
"github.com/aws/aws-sdk-go/service/sts/stsiface"
"fmt"
)

const (
	AwsDefaultRegion   = "us-west-1"
	AwsAccessKeyId     = "access-key-id"
	AwsSecretAccessKey = "secret-access-key"
)

func setAwsConfig() {
	// do the credential setting here
	//os.Setenv("AWS_ACCESS_KEY_ID", AwsAccessKeyId)
	//os.Setenv("AWS_SECRET_ACCESS_KEY", AwsSecretAccessKey)
}

func getStsClient() stsiface.STSAPI {
	sess := session.Must(session.NewSession())
	client := sts.New(sess)
	return client
}

func getSessionToken(client stsiface.STSAPI) (*sts.GetSessionTokenOutput, error) {
	sessionToken, err := client.GetSessionToken(&sts.GetSessionTokenInput{
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return sessionToken, nil
}

func main() {
	client := getStsClient()
	sessionToken, err := getSessionToken(client)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sessionToken)
}