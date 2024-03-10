# fifth task

## Usage/Examples

```
 go run ./cmd/api/
```

## Features

- Start server that running on 3000 port
- registerion
- authorization 

## API Endpoints

### Register

```http
POST /api/v1/register
```

### Login

```http
POST /api/v1/login
```

## Web Routes

### Register
```http
GET /register
```

### Login
```http
GET /login
```

## Authors

- [Elaman Ismagulov](https://t.me/Double_power)

## Additional

i used json Marshal and Unmarshal in pkg/httpjson package
and you can check the /internal/auth/transport/http/handlers.go  Rgister handler