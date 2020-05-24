package service

type SvcBr struct {
	Enabled bool `json: enabled`
}

type SvcIi struct {
	Enabled bool   `json: enabled`
	ApiKey  string `json: apikey`
}

type SvcVt struct {
	Enabled bool   `json: enabled`
	Apikey  string `json: apikey`
}

type Svc struct {
	Listen     string `json: listen`
	VirusTotal *SvcVt `json: virustotal`
	IpInfo     *SvcIi `json: ipinfo`
	Browser    *SvcBr `json: browser`
}
