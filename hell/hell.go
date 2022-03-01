package hell

import (
	"github.com/valyala/fasthttp"
	url2 "net/url"
	"sync"
	"time"
)

type Hell struct {
	RawUrl       string
	GophersCount int
	Method       string
}

func rain(request *fasthttp.Request, client *fasthttp.HostClient) {
	client.Do(request, nil)
}

func gopher(r chan struct{}, wg *sync.WaitGroup, request *fasthttp.Request, client *fasthttp.HostClient) {
	<-r
	defer wg.Done()

	time.Sleep(1 * time.Second)

	for {
		rain(request, client)
	}
}

func prepareGophers(r chan struct{}, wg *sync.WaitGroup, request *fasthttp.Request, client *fasthttp.HostClient) {
	println("Preparing Gophers")

	for i := 0; i < client.MaxConns; i++ {
		go gopher(r, wg, request, client)
	}

	println("Gophers ready")

}

func (b *Hell) Start() {
	url, _ := url2.Parse(b.RawUrl)

	request := createRequest(url, b.Method)
	client := createClient(url, b.GophersCount)

	var wg sync.WaitGroup
	r := make(chan struct{})

	wg.Add(b.GophersCount)

	prepareGophers(r, &wg, request, client)

	println("Start the Hell on Earth")
	close(r)
	wg.Wait()
}
