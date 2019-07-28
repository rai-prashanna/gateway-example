package app

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"log"
)

func RunApp() {
	// New functionality written in Go
	log.Printf("Starting...")
	someFunc()

}


func someFunc() {
	target, err := url.Parse("http://localhost:8081")
	log.Printf("forwarding to -> %s %s\n", target.Scheme, target.Host)

	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(target)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// https://stackoverflow.com/questions/38016477/reverse-proxy-does-not-work
		// https://forum.golangbridge.org/t/explain-how-reverse-proxy-work/6492/7
		// https://stackoverflow.com/questions/34745654/golang-reverseproxy-with-apache2-sni-hostname-error

		req.Host = req.URL.Host // if you remove this line the request will fail... I want to debug why.

		proxy.ServeHTTP(w, req)
	})

	err = http.ListenAndServe(":8989", nil)
	if err != nil {
		panic(err)
	}
}


func someFunc1(w http.ResponseWriter, r *http.Request) {

	// change the request host to match the target
	r.Host = "you.host.here:port"
	u, _ := url.Parse("http://you.host.here:port/some/path/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	// You can optionally capture/wrap the transport if that's necessary (for
	// instance, if the transport has been replaced by middleware). Example:
	// proxy.Transport = &myTransport{proxy.Transport}
	proxy.Transport = &myTransport{}

	proxy.ServeHTTP(w, r)
}
type myTransport struct {
	// Uncomment this if you want to capture the transport
	// CapturedTransport http.RoundTripper
}

func (t *myTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	response, err := http.DefaultTransport.RoundTrip(request)
	// or, if you captured the transport
	// response, err := t.CapturedTransport.RoundTrip(request)

	// The httputil package provides a DumpResponse() func that will copy the
	// contents of the body into a []byte and return it. It also wraps it in an
	// ioutil.NopCloser and sets up the response to be passed on to the client.
	body, err := httputil.DumpResponse(response, true)
	if err != nil {
		// copying the response body did not work
		return nil, err
	}

	// You may want to check the Content-Type header to decide how to deal with
	// the body. In this case, we're assuming it's text.
	log.Print(string(body))

	return response, err
}