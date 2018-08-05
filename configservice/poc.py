import boto3
import os
import json
from bson import json_util
AWS_DEFAULT_REGION = 'us-west-1'
AWS_ACCESS_KEY_ID = 'access-key-id'
AWS_SECRET_ACCESS_KEY = 'secret-access-key'

def set_aws_config():
    # do the credential setting here
    #os.environ['AWS_ACCESS_KEY_ID'] = AWS_ACCESS_KEY_ID
    #os.environ['AWS_SECRET_ACCESS_KEY'] = AWS_SECRET_ACCESS_KEY
    pass

def get_config_client(region):
    set_aws_config()
    session = boto3.session.Session()
    client = session.client('config', region_name=region)
    return client

def get_config_rules(client):
    next_token = ""
    config_rules = []
    while True:
        response = client.describe_config_rules()
        if isRequestSuccessful(response):
            if response['ConfigRules']:
                config_rules.extend(response['ConfigRules'])
        next_token = response.get("NextToken")
        if not next_token:
            break
    return config_rules

def get_compliance_details(client, rule_name):
    next_token = ""
    compliance_details = []
    while True:
        response = client \
            .get_compliance_details_by_config_rule(ConfigRuleName=rule_name,
                                                   NextToken=next_token)
        if isRequestSuccessful(response):
            if response["EvaluationResults"]:
                compliance_details.extend(response["EvaluationResults"])
        next_token = response.get("NextToken")
        if not next_token:
            break
    return compliance_details

def get_compliance_summary(client):
    response = client \
        .get_compliance_summary_by_config_rule()
    if isRequestSuccessful(response):
        #print(response)
        #print(json.dumps(response, default=json_util.default))
        return response["ComplianceSummary"]

def isRequestSuccessful(response):
    return response["ResponseMetadata"]["HTTPStatusCode"] == 200

def main():
    # get config rules
    client = None
    regions = [AWS_DEFAULT_REGION]

    config_rules = []
    config_rules_details = []

    for region in regions:
        client = get_config_client(region)
        if client:
            config_rules = get_config_rules(client)
            # check compliance of each rule
            for config_rule in config_rules:
                config_rule_details = get_compliance_details(client, config_rule["ConfigRuleName"])
                config_rules_details.extend(config_rule_details)
                #config_rule_summary = get_compliance_summary(client)
                #config_rules_summary.append(config_rule_summary)

    print(json.dumps(config_rules, default=json_util.default))
    print(json.dumps(config_rules_details, default=json_util.default))



if __name__=='__main__':
    main()