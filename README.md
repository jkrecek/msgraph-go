#msgraph-go
Library for communication with Microsoft Graph API.

Library support
---------------
Currently supported are requests to:
* basic user info
* calendars

More request support may come soon. All exposed methods are subject to change.
 
Usage
-----
To install the libary into your project, add msgraph-go inti your `GOPATH`:
 
```sh
$ go get github.com/jkrecek/msgraph-go
```

Then, in your go project, you can create graph client:
 
 ```go
 import "github.com/jkrecek/msgraph-go"

 var exampleOAuthConfig = &oauth2.Config{
    ClientID: "YOUR_CLIENT_ID",
    RedirectURL: "https://localhost",
    Scopes: []string{
        "User.Read",
    },
    Endpoint: oauth2.Endpoint{
        AuthURL: "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
        TokenURL: "https://login.microsoftonline.com/common/oauth2/v2.0/token",
    },
}

exampleOAuthToken := &oauth2.Token{
    AccessToken: "YOUR_ACCESS_TOKEN",
    RefreshToken: "YOUR_REFRESH_TOKEN",
    Expiry: time.Now().Add(time.Hour), 
}
 
 
 client := graph.NewClient(exampleOAuthConfig, exampleOAuthToken)
 ```

Both oauth2.Config and oauth2.Token are required to successfully authorize outgoing requests.

To make further requests, you can call method directly on client, for example:

```go
...

client := graph.NewClient(exampleOAuthConfig, exampleOAuthToken)

me, err := client.GetMe()
if err != nil {
    log.Println(err)
}

_ = me.Id                   // Returns user id
_ = me.UserPrincipalName    // Returns user principal name, usually email address

