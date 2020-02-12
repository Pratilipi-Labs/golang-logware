# golang-logware
## RequestLogger
Request logging middleware for golang web services. Use this library to log the requests hitting the web service. It is advised to use **Request-Id** in request headers to trace the spread of requests in microservice environment.

## How to use?
```package main

import (
    "net/http"
    "github.com/labstack/echo"
    "https://github.com/Pratilipi-Labs/golang-logware"
)

func main() {
    e := echo.New()

    logger := requestLogger.RequestLogger{"ERROR","test"}

    e.Use(logger.Log)

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })  

    e.GET("/400",func(c echo.Context) error {
        return c.String(http.StatusBadRequest,"Bad request")
    })  
    e.Logger.Fatal(e.Start(":8080"))
}
```

### Log example:
The log prints in json string format.  
```{"level":"error","service":"test","message":"Bad request","uri":"/bad?a=1","responseCode":400,"requestId":"23545"}```

## TODO
Add support to skip logging for certain apis
Framework agnostic middleware
Configurable log structure
