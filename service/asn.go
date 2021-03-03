package service

import (
	"net"
	"net/http"
	"time"

	"github.com/ammario/ipisp"
	sentry "github.com/getsentry/sentry-go"
	echo "github.com/labstack/echo/v4"
)

type AsnData struct {
	Asn         ipisp.ASN `json: asn`
	Country     string    `json: country`
	Registry    string    `json: registry`
	Range       string    `json: range`
	AllocatedAt time.Time `json: allocated_at`
}

type Asn struct {
	Data  *AsnData `json: data`
	Error error    `json: error`
}

func getAsn(remoteAddr string) *Asn {
	client, err := ipisp.NewDNSClient()
	if err != nil {
		return &Asn{
			Data:  nil,
			Error: err,
		}
	}
	defer client.Close()

	resp, err := client.LookupIP(net.ParseIP(remoteAddr))
	if err != nil {
		return &Asn{
			Data:  nil,
			Error: err,
		}
	}

	return &Asn{
		Data: &AsnData{
			Asn:         resp.ASN,
			Country:     resp.Country,
			Registry:    resp.Registry,
			Range:       resp.Range.String(),
			AllocatedAt: resp.AllocatedAt,
		},
		Error: nil,
	}
}

/*
GetAsnDetails (echo.Context)
-> app.GET("/api/v1/asn", svc.GetAsnDetails())
*/
func (s *Svc) GetAsnDetails(c echo.Context) error {
	remoteAddr := c.QueryParam("ip")
	if remoteAddr == "" {
		remoteAddr = c.Request().Header.Get("Forwarded")
	}

	asn := getAsn(remoteAddr)
	if asn.Error != nil {
		sentry.CaptureException(asn.Error)
		sentry.Flush(time.Second * 5)

		return c.JSON(http.StatusInternalServerError, asn)
	}

	return c.JSON(http.StatusOK, asn)
}
