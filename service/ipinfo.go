package service

import (
	"net"

	"github.com/ipinfo/go-ipinfo/ipinfo"
)

type IpInfo struct {
	Data  *ipinfo.Info `json: data`
	Error error        `json: error`
}

func (s *Svc) GetIPInfo(remoteAddr string) *IpInfo {
	authTransport := ipinfo.AuthTransport{Token: s.IpInfo.ApiKey}
	httpClient := authTransport.Client()
	client := ipinfo.NewClient(httpClient)

	info, err := client.GetInfo(net.ParseIP(remoteAddr))

	return &IpInfo{
		Data:  info,
		Error: err,
	}
}
