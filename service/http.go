package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
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
	Data  *Data  `json: data`
	Error string `json: error`
}

func isIpv4(host string) bool {
	return net.ParseIP(host) != nil
}

func RemoteAddrSplit(w http.ResponseWriter, r *http.Request) (string, int64, error) {
	remoteAddr := strings.Split(r.RemoteAddr, ":")
	port, err := strconv.ParseInt(remoteAddr[1], 10, 64)

	return remoteAddr[0], port, err
}

func (s *Svc) HttpRet(w http.ResponseWriter, r *http.Request) {
	log.Println("ra:[" + r.Header.Get("Forwarded") + "] ua:[" + r.UserAgent() + "] m:[" + r.Method + "]")
	w.Header().Add("Content-Type", "application/json")

	var all bool

	htst := &HttpSt{
		Data: &Data{
			RemoteAddr: "",
			// Port:             port,
			Proto:   r.Proto,
			Asn:     nil,
			IpInfo:  nil,
			Vt:      nil,
			Browser: nil,
		},
		Error: "",
	}

	keys, ok := r.URL.Query()["ip"]
	if !ok || len(keys[0]) < 1 {
		// remoteAddr, _, _ = RemoteAddrSplit(w, r)
		htst.Data.RemoteAddr = r.Header.Get("Forwarded")
	} else {
		htst.Data.RemoteAddr = keys[0]
	}

	keys, ok = r.URL.Query()["all"]
	if !ok || len(keys[0]) < 1 {
		// remoteAddr, _, _ = RemoteAddrSplit(w, r)
		all = false
	} else {
		var err error
		all, err = strconv.ParseBool(keys[0])
		if err != nil {
			htst.Error = "Could not parse GET['all'] parameter"
			res, _ := json.Marshal(htst)
			fmt.Fprintf(w, "%s\n", string(res))
			return
		}
	}

	if !isIpv4(htst.Data.RemoteAddr) {
		htst.Error = "Could not parse IP from request"
		res, _ := json.Marshal(htst)
		fmt.Fprintf(w, "%s\n", string(res))
		return
	}

	if s.IpInfo.Enabled == true {
		htst.Data.IpInfo = s.GetIPInfo(htst.Data.RemoteAddr)
	}

	if all {
		if s.VirusTotal.Enabled == true {
			htst.Data.Vt = s.GetVtIp(htst.Data.RemoteAddr)
		}

		if s.Browser.Enabled == true {
			htst.Data.Browser = s.GetBrowserDetails(r)
		}

		htst.Data.Asn = GetAsn(htst.Data.RemoteAddr)
	}

	res, _ := json.Marshal(htst)
	fmt.Fprintf(w, "%s\n", string(res))
}
