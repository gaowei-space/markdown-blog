package app

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/urfave/cli/v2"
)

// NewProxy takes target host and creates a reverse proxy
func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		modifyRequest(req)
	}

	proxy.ErrorHandler = errorHandler()
	return proxy, nil
}

func modifyRequest(req *http.Request) {
	req.Header.Set("X-Proxy", "Simple-Reverse-Proxy")
}

func errorHandler() func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, req *http.Request, err error) {
		fmt.Printf("Got error while modifying response: %v \n", err)
	}
}

// ProxyRequestHandler handles the http request using proxy
func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

func setProxy() {
	// initialize a reverse proxy and pass the actual backend server url here
	proxy, err := NewProxy("http://127.0.0.1:" + Port)
	if err != nil {
		log.Panic(err)
	}

	// handle all requests to your server using the proxy
	http.HandleFunc("/", ProxyRequestHandler(proxy))
}

func startHttp() {
	http.ListenAndServe(Listener.Host+":80", nil)
}

func startHttps() {
	if Listener.Cert.CertFile == "" {
		return
	}

	if Listener.Cert.KeyFile == "" {
		return
	}

	var err error

	_, err = os.Stat(Listener.Cert.CertFile)
	if err != nil {
		return
	}

	_, err = os.Stat(Listener.Cert.KeyFile)
	if err != nil {
		return
	}

	http.ListenAndServeTLS(Listener.Host+":443", Listener.Cert.CertFile, Listener.Cert.KeyFile, nil)
}

func RunListener(ctx *cli.Context) {
	if !Listener.Open {
		return
	}

	setProxy()

	go startHttp()
	go startHttps()
}
