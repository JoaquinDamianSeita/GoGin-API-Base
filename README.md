# GoGin-API-Base

run with:

``` bash
go run .
```

Generate wire file:
``` bash
wire gen GoGin-API-Base/config
```

.env:

``` go
PORT=8080
# Application
APPLICATION_NAME=GoGin-API-Base

# Database
DB_DSN="host=HOST user=USER password=PASSWORD dbname=DBNAME port=PORT"

# Logging
LOG_LEVEL=DEBUG
```

Live Reload Golang Development With Gin:

``` bash
gin --appPort 3000 --port 8080 --immediate
```
