# fourth task

## Usage/Examples

```
 go run ./cmd/api/
```

## Features

- Start server that running on 3000 port
- "registerion" 

## Routes

- path: /api/v1/auth/register method: POST data: name, email, password
 
## Authors

- [Elaman Ismagulov](https://t.me/Double_power)

## Additional

i used json Marshal and Unmarshal in pkg/httpjson package
and you can check the /internal/auth/transport/http/handlers.go  Rgister handler