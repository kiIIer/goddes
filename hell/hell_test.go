package hell

import (
	url2 "net/url"
	"testing"
)

func BenchmarkRain(b *testing.B) {
	url, _ := url2.Parse("http://localhost:80/")
	gophersCount := 1

	client := createClient(url, gophersCount)
	request := createRequest(url, "GET")

	for i := 0; i < b.N; i++ {
		rain(request, client)
	}
}
