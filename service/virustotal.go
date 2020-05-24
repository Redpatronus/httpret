package service

import (
	"encoding/json"
	"log"

	vt "github.com/VirusTotal/vt-go"
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

func (s *Svc) GetVtIp(remoteAddr string) *Vt {
	client := vt.NewClient(s.VirusTotal.Apikey)
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
