package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/aws/aws-sdk-go/service/configservice/configserviceiface"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
)

const (
	AwsDefaultRegion = "us-west-1"
	AwsAccessKeyId = "access-key-id"
	AwsSecretAccessKey = "secret-access-key"
)

func setAwsConfig() {
	// do the credential setting here
	//os.Setenv("AWS_ACCESS_KEY_ID", AwsAccessKeyId)
	//os.Setenv("AWS_SECRET_ACCESS_KEY", AwsSecretAccessKey)
}

func getConfigClient(region string) configserviceiface.ConfigServiceAPI {
	sess := session.Must(session.NewSession())
	client := configservice.New(sess, aws.NewConfig().WithRegion(region))
	return client
}

func getConfigRules(client configserviceiface.ConfigServiceAPI) (
	[]*configservice.ConfigRule, error) {
	output, err := client.DescribeConfigRules(&configservice.DescribeConfigRulesInput{
	})
	return output.ConfigRules, err
}

func getComplianceDetails(client configserviceiface.ConfigServiceAPI, rule_name string) (
	[]*configservice.EvaluationResult, error) {
	output, err := client.GetComplianceDetailsByConfigRule(
		&configservice.GetComplianceDetailsByConfigRuleInput{
			ConfigRuleName: &rule_name,
		})
	return output.EvaluationResults, err
}

func main() {
	setAwsConfig()
	regions := [1]string{AwsDefaultRegion}
	configRulesResult := make([]*configservice.ConfigRule, 0)
	configRulesDetailsResult := make([]*configservice.EvaluationResult, 0)
	for _, region := range regions {
		client := getConfigClient(region)
		configRules, err := getConfigRules(client)
		configRulesResult = append(configRulesResult, configRules...)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _,configRule := range configRules {
			evaluationResults, err := getComplianceDetails(client, *configRule.ConfigRuleName)
			if err != nil {
				fmt.Println(err)
				continue
			}
			configRulesDetailsResult = append(configRulesDetailsResult, evaluationResults...)
		}
	}
	fmt.Println(configRulesResult)
	fmt.Println(configRulesDetailsResult)
}