package service

import (
	//	"encoding/json"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type Header map[string]string

type Browser struct {
	Data map[string]string
}

func processRequestHeaders(r *http.Request) *Browser {
	var b *Browser = &Browser{}
	b.Data = make(map[string]string)
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			b.Data[name] = value
		}
	}
	return b
}

/*
GetBrowserDetails (echo.Context)
-> app.GET("/api/v1/browser", svc.GetBrowserDetails)
*/
func (s *Svc) GetBrowserDetails(c echo.Context) error {
	browserDetails := processRequestHeaders(c.Request())
	return c.JSON(http.StatusOK, browserDetails)
}
