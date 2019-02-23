# ServiceBear

ServiceBear is a small CLI-Utility to generate OAuth Access Tokens from Google's Service Accounts. It uses exactly the same authenthication flow, as described in https://developers.google.com/identity/protocols/OAuth2ServiceAccount

![](https://developers.google.com/accounts/images/serviceaccount.png)

## Installation

Grab the latest binary release from the [Releases](https://github.com/mishushakov/servicebear/releases) page

Binaries for Mac are with `-darwin` suffix

Binaries for Linux (amd64) are with `-linux` suffix

Binaries for Windows (x64) are with `-win` suffix

## Usage

```sh
Getting Help

./servicebear --help
Usage of ./servicebear:
  -scope string
        Scope (default "https://www.googleapis.com/auth/dialogflow")
  -service_account string
        Path to service account (default "./service_account.json")

Generating the Token from Service Account

./servicebear -service_account="./../example/service_account.json" -scope="https://www.googleapis.com/auth/bigquery"

Service Account: ./../example/service_account.json
Access Token: ya29.c.Elq5BkwdrMPMQWU7LCUmdv_YslHSLEp1HMtHeVAmSh86fGfrsbfdSbkUQ0qaoqz4iR0vr-c86ckk6UWPd8cvRYnwfY6Ogh2PiGlC5ZvX2jV66i87PqpA
Token Type: Bearer
Refresh Token: 
Expiry: 2019-02-23 05:03:27.489464 +0100 CET m=+3600.437215906
```

You can find all the available the scopes under https://developers.google.com/identity/protocols/googlescopes

## Thank you for using the utility!