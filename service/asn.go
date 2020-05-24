package service

import (
	"net"
	"time"

	"github.com/ammario/ipisp"
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

func GetAsn(remoteAddr string) *Asn {
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
