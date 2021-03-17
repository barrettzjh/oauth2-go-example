# oauth2-go-example
oauth2 example

# Run APP
```bee run```
or
```go run main.go```

# API Test
```
不携带token访问资源时
HTTP GET
127.0.0.1:8080/
Status 403
invalid access token


获取clientid和clientsecret
HTTP GET
127.0.0.1:8080/credentials
response
{
    "client_id": "54b052bb",
    "client_secret": "8fe0a1c8"
}


获取Token
HTTP GET
127.0.0.1:8080/token?grant_type=client_credentials&client_id=54b052bb&client_secret=8fe0a1c8&scope=all
response
{
    "access_token": "ZDQWMWU1ZMYTMTRKZS0ZZWY3LTK0OTUTYTC0YJJJZTCXMTG5",
    "expires_in": 7200,
    "scope": "all",
    "token_type": "Bearer"
}

携带token访问资源
HTTP GET
127.0.0.1:8080/?access_token=ZDQWMWU1ZMYTMTRKZS0ZZWY3LTK0OTUTYTC0YJJJZTCXMTG5

```
