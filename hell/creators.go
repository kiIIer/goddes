package hell

import (
	"crypto/tls"
	"github.com/valyala/fasthttp"
	"net"
	url2 "net/url"
)

func createRequest(url *url2.URL, method string) *fasthttp.Request {
	req := fasthttp.AcquireRequest()

	req.Header.SetHost(url.Host)
	req.Header.SetMethod(method)
	req.URI().SetScheme(url.Scheme)
	req.SetRequestURI(url.RequestURI())
	req.SetBodyString("")

	return req
}

type countingConn struct {
	net.Conn
	bytesRead, bytesWritten *int
}

func createClient(url *url2.URL, gophersCount int) *fasthttp.HostClient {

	var dialFunc = func(
		bytesRead, bytesWritten *int,
	) func(string) (net.Conn, error) {
		return func(address string) (net.Conn, error) {
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return nil, err
			}

			wrappedConn := &countingConn{
				Conn:         conn,
				bytesRead:    bytesRead,
				bytesWritten: bytesWritten,
			}

			return wrappedConn, nil
		}
	}
	var bytes int
	bytes = 0
	client := &fasthttp.HostClient{
		Addr:                          url.Host,
		IsTLS:                         url.Scheme == "https",
		ReadTimeout:                   2000000000,
		WriteTimeout:                  2000000000,
		DisableHeaderNamesNormalizing: true,
		TLSConfig: &tls.Config{
			Rand:                        nil,
			Time:                        nil,
			Certificates:                nil,
			GetCertificate:              nil,
			GetClientCertificate:        nil,
			GetConfigForClient:          nil,
			VerifyPeerCertificate:       nil,
			VerifyConnection:            nil,
			RootCAs:                     nil,
			NextProtos:                  nil,
			ServerName:                  "",
			ClientAuth:                  tls.NoClientCert,
			ClientCAs:                   nil,
			InsecureSkipVerify:          false,
			CipherSuites:                nil,
			PreferServerCipherSuites:    false,
			SessionTicketsDisabled:      false,
			ClientSessionCache:          nil,
			MinVersion:                  0,
			MaxVersion:                  0,
			CurvePreferences:            nil,
			DynamicRecordSizingDisabled: false,
			Renegotiation:               tls.RenegotiateNever,
			KeyLogWriter:                nil,
		},
		Dial:     dialFunc(&bytes, &bytes),
		MaxConns: gophersCount,
	}

	return client

}
