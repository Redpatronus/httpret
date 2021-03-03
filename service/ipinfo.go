package service

import (
	"net"
	"net/http"
	"time"

	sentry "github.com/getsentry/sentry-go"
	"github.com/ipinfo/go-ipinfo/ipinfo"
	echo "github.com/labstack/echo/v4"
)

type IpInfo struct {
	Data  *ipinfo.Info `json: data`
	Error error        `json: error`
}

func getIpInfo(remoteAddr string, apiKey string) *IpInfo {
	authTransport := ipinfo.AuthTransport{Token: apiKey}
	httpClient := authTransport.Client()
	client := ipinfo.NewClient(httpClient)

	info, err := client.GetInfo(net.ParseIP(remoteAddr))

	return &IpInfo{
		Data:  info,
		Error: err,
	}
}

func (s *Svc) GetIPInfo(c echo.Context) error {
	remoteAddr := c.QueryParam("ip")
	if remoteAddr == "" {
		remoteAddr = c.Request().Header.Get("Forwarded")
	}

	ii := getIpInfo(remoteAddr, s.IpInfo.ApiKey)
	if ii.Error != nil {
		sentry.CaptureException(ii.Error)
		sentry.Flush(time.Second * 5)

		return c.JSON(http.StatusInternalServerError, ii)
	}

	return c.JSON(http.StatusOK, ii)
}
