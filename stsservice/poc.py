import boto3

AWS_DEFAULT_REGION = 'us-west-1'
AWS_ACCESS_KEY_ID = 'access-key-id'
AWS_SECRET_ACCESS_KEY = 'secret-access-key'

def set_aws_config():
    # do the credential setting here
    #os.environ['AWS_ACCESS_KEY_ID'] = AWS_ACCESS_KEY_ID
    #os.environ['AWS_SECRET_ACCESS_KEY'] = AWS_SECRET_ACCESS_KEY
    pass

def get_sts_client():
    set_aws_config()
    session = boto3.session.Session()
    client = session.client('sts')
    return client

def main():
    sts_client = get_sts_client()
    session_token = sts_client.get_session_token()
    print(session_token)

if __name__ == '__main__':
    main()
