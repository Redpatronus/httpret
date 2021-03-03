package service

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	vt "github.com/VirusTotal/vt-go"
	sentry "github.com/getsentry/sentry-go"
	echo "github.com/labstack/echo/v4"
)

type VtLastAnalysisResult struct {
	Category    string `json: category`
	Engine_Name string `json: engine_name`
	Method      string `json: method`
	Result      string `json: result`
}

type VtLastAnalysisResults map[string]VtLastAnalysisResult

type VtLastAnalysisStats struct {
	Harmless   int `json: harmless`
	Malicious  int `json: malicious`
	Suspicious int `json: suspicious`
	Timeout    int `json: timeout`
	Undetected int `json: undetected`
}

type VtAttributes struct {
	Last_Analysis_Results      VtLastAnalysisResults `json: last_analysis_results`
	Last_Analysis_Stats        VtLastAnalysisStats   `json: last_analysis_stats`
	Network                    string                `json: network`
	Reputation                 int                   `json: reputation`
	Regional_Internet_Registry string                `json: regional_internet_registry`
	Asn                        int                   `json: asn`
	Country                    string                `json: country`
	As_Owner                   string                `json: as_owner`
	Continent                  string                `json: continent`
}

type Vt struct {
	Attributes VtAttributes `json: attributes`
}

func getVtIp(remoteAddr string, apiKey string) *Vt {
	client := vt.NewClient(apiKey)
	ret := &Vt{}

	url, err := client.GetObject(vt.URL("ip_addresses/%s", remoteAddr))
	if err != nil {
		log.Fatal(err)
	}

	a, err := json.Marshal(url)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(a, &ret)
	if err != nil {
		log.Fatal(err)
	}

	return ret
}

/*
GetVirusTotalDetails (echo.Context)
-> app.GET("/api/v1/virustotal", svc.GetVirusTotalDetails)
*/
func (s *Svc) GetVirusTotalDetails(c echo.Context) error {
	remoteAddr := c.QueryParam("ip")
	if remoteAddr == "" {
		remoteAddr = c.Request().Header.Get("Forwarded")
	}

	vt := getIpInfo(remoteAddr, s.VirusTotal.ApiKey)
	if vt.Error != nil {
		sentry.CaptureException(vt.Error)
		sentry.Flush(time.Second * 5)

		return c.JSON(http.StatusInternalServerError, vt)
	}

	return c.JSON(http.StatusOK, vt)
}
