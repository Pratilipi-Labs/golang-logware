package requestlogger

import (
    "github.com/labstack/echo"
    "strings"
    "fmt"
    "encoding/json"
)


//TODO: Add support to skip logging for certain apis
//TODO: Framework agnostic middleware
//TODO: Configurable log structure


type RequestLogger struct {
    Level,ServiceName string
}

type logStructure struct {
	Service        string `json:"service"`
	RequestId      string `json:"requestId"`
	Url            string `json:"url"`
	Level          string `json:"level"`
	HttpStatusCode int    `json:"httpStatusCode"`
}

func (logger *RequestLogger) Log(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}

        if logger.Level == "INFO" || (logger.Level == "ERROR" && c.Response().Status >= 400) {
            go logger.printRequestLogs(c)
		}
		return nil
	}
}

func (logger *RequestLogger) printRequestLogs(c echo.Context) {
    ls := logStructure{logger.ServiceName, c.Request().Header.Get("request-id") , c.Request().URL.String(), strings.ToLower(logger.Level), c.Response().Status}
    b, _ := json.Marshal(ls)
    fmt.Println(string(b))
}


