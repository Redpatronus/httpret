package service

// Browser
type SvcBr struct {
	Enabled bool `json: enabled`
}

// IpInfo
type SvcIi struct {
	Enabled bool   `json: enabled`
	ApiKey  string `json: apikey`
}

// VirusTotal
type SvcVt struct {
	Enabled bool   `json: enabled`
	ApiKey  string `json: apikey`
}

// Sentry
type SvcSt struct {
	Enabled bool   `json: enabled`
	Dsn     string `json: dsn`
}

// Sqreen
type SvcSq struct {
	Enabled bool   `json: enabled`
}

type Svc struct {
	Listen     string `json: listen`
	VirusTotal *SvcVt `json: virustotal`
	IpInfo     *SvcIi `json: ipinfo`
	Browser    *SvcBr `json: browser`
	Sentry     *SvcSt `json: sentry`
	Sqreen	   *SvcSq `json: sqreen`
}
