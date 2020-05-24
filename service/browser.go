package service

import (
	"net/http"
)

type Header map[string]string

type Browser struct {
	Data map[string]string
}

func (s *Svc) GetBrowserDetails(r *http.Request) *Browser {
	var b *Browser = &Browser{}
	b.Data = make(map[string]string)
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			b.Data[name] = value
		}
	}
	return b
}
