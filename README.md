# Deprecation Notice

Google built API endpoint to generate the OAuth Access Tokens for Service Accounts: https://cloud.google.com/iam/credentials/reference/rest/v1/projects.serviceAccounts/generateAccessToken

---

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

## Using the Token

After you obtain an access token, you can use the token to make calls to a Google API on behalf of a given service account. To do this, include the access token in a request to the API by including either an access_token query parameter or an Authorization: Bearer HTTP header. When possible, the HTTP header is preferable, because query strings tend to be visible in server logs. In most cases you can use a client library to set up your calls to Google APIs (for example, when [calling the Drive Files API](https://developers.google.com/drive/api/v2/reference/#Files)).

You can try out all the Google APIs and view their scopes at the [OAuth 2.0 Playground](https://developers.google.com/oauthplayground/).

**HTTP GET examples**

A call to the [drive.files](https://developers.google.com/drive/api/v2/reference/files/list) endpoint (the Drive Files API) using the Authorization: Bearer HTTP header might look like the following. Note that you need to specify your own access token:

```http
GET /drive/v2/files HTTP/1.1
Authorization: Bearer <access_token>
Host: www.googleapis.com/
```

Here is a call to the same API for the authenticated user using the access_token query string parameter:

```http
GET https://www.googleapis.com/drive/v2/files?access_token=<access_token>
```

**curl examples**

You can test these commands with the curl command-line application. Here's an example that uses the HTTP header option (preferred):

```sh
curl -H "Authorization: Bearer <access_token>" https://www.googleapis.com/drive/v2/files
```

Or, alternatively, the query string parameter option:

```sh
curl https://www.googleapis.com/drive/v2/files?access_token=<access_token>
```

**When access tokens expire**

Access tokens issued by the Google OAuth 2.0 Authorization Server expire one hour after they are issued. When an access token expires, then the application should generate another JWT, sign it, and request another access token.

You can find all the available the scopes under https://developers.google.com/identity/protocols/googlescopes

## Thank you for using the utility!
