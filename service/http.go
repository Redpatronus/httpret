package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Data struct {
	RemoteAddr string `json: remote_addr`
	// Port             int64    `json: port`
	Proto   string   `json: proto`
	Asn     *Asn     `json: asn`
	IpInfo  *IpInfo  `json: ipinfo`
	Vt      *Vt      `json: vt`
	Browser *Browser `json: browser`
}

type HttpSt struct {
	Data  *Data `json: data`
	Error error `json: error`
}

func RemoteAddrSplit(w http.ResponseWriter, r *http.Request) (string, int64, error) {
	remoteAddr := strings.Split(r.RemoteAddr, ":")
	port, err := strconv.ParseInt(remoteAddr[1], 10, 64)

	return remoteAddr[0], port, err
}

func (s *Svc) HttpRet(w http.ResponseWriter, r *http.Request) {
	log.Println("ra:[" + r.Header.Get("Forwarded") + "] ua:[" + r.UserAgent() + "] m:[" + r.Method + "]")

	var remoteAddr string
	keys, ok := r.URL.Query()["ip"]
	if !ok || len(keys[0]) < 1 {
		// remoteAddr, _, _ = RemoteAddrSplit(w, r)
		remoteAddr = r.Header.Get("Forwarded")
	} else {
		remoteAddr = keys[0]
	}

	var ii *IpInfo
	var vt *Vt
	var br *Browser

	if s.IpInfo.Enabled == true {
		ii = s.GetIPInfo(remoteAddr)
	}

	if s.VirusTotal.Enabled == true {
		vt = s.GetVtIp(remoteAddr)
	}

	if s.Browser.Enabled == true {
		br = s.GetBrowserDetails(r)
	}

	as := GetAsn(remoteAddr)

	htst := &HttpSt{
		Data: &Data{
			RemoteAddr: remoteAddr,
			// Port:             port,
			Proto:   r.Proto,
			Asn:     as,
			IpInfo:  ii,
			Vt:      vt,
			Browser: br,
		},
		Error: nil,
	}

	res, _ := json.Marshal(htst)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", string(res))
}
