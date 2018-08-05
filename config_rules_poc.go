package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/aws/aws-sdk-go/service/configservice/configserviceiface"
	"github.com/aws/aws-sdk-go/aws"
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

func getConfigClient(region string) configserviceiface.ConfigServiceAPI {
	sess := session.Must(session.NewSession())
	client := configservice.New(sess, aws.NewConfig().WithRegion(region))
	return client
}

func getConfigRules(client configserviceiface.ConfigServiceAPI) (
	[]*configservice.ConfigRule, error) {
		configRules := make([]*configservice.ConfigRule, 0)
		nextToken := ""
		for {
			output, err := client.DescribeConfigRules(&configservice.DescribeConfigRulesInput{
				NextToken: &nextToken,
			})
			if err != nil {
				return configRules, err
			}
			configRules = append(configRules, output.ConfigRules...)
			if output.NextToken == nil {
				break
			}
			nextToken = *output.NextToken
		}
		return configRules, nil
}

func getComplianceDetails(client configserviceiface.ConfigServiceAPI, rule_name string) (
	[]*configservice.EvaluationResult, error) {
		complianceDetails := make([]*configservice.EvaluationResult, 0)
		nextToken := ""
		for {
			output, err := client.GetComplianceDetailsByConfigRule(
				&configservice.GetComplianceDetailsByConfigRuleInput{
					ConfigRuleName: &rule_name,
					NextToken: &nextToken,
				})
			if err != nil {
				return complianceDetails, err
			}
			complianceDetails = append(complianceDetails, output.EvaluationResults...)
			if output.NextToken == nil {
				break
			}
			nextToken = *output.NextToken
		}
		return complianceDetails, nil

}

func main() {
	setAwsConfig()
	regions := [1]string{AwsDefaultRegion}
	configRulesResult := make([]*configservice.ConfigRule, 0)
	complianceDetailsResult := make([]*configservice.EvaluationResult, 0)
	for _, region := range regions {
		client := getConfigClient(region)
		configRules, err := getConfigRules(client)
		if err != nil {
			fmt.Println(err)
			continue
		}
		configRulesResult = append(configRulesResult, configRules...)
		for _, configRule := range configRules {
			evaluationResults, err := getComplianceDetails(client, *configRule.ConfigRuleName)
			if err != nil {
				fmt.Println(err)
				continue
			}
			complianceDetailsResult = append(complianceDetailsResult, evaluationResults...)
		}
	}
	fmt.Println(configRulesResult)
	fmt.Println(complianceDetailsResult)
}