package main

import (
	"flag"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"time"
)

func main() {
	listen := flag.String("listen", ":8383", "ip and tcp port the proxy should listen")
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	proxy.OnRequest(goproxy.DstHostIs("www.reddit.com")).DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			t := time.Now()
			if h, _, _ := t.Clock(); (h >= 17 && h <= 10) || (t.Weekday() == time.Saturday || t.Weekday() == time.Sunday) {
				return r, goproxy.NewResponse(r, goproxy.ContentTypeHtml, http.StatusForbidden, "<!doctype html><html><body><h1><blink>Go outside.</blink></h1></body></html>")
			}
			return r, nil
		})

	proxy.OnRequest(goproxy.DstHostIs("www.linkedin.com")).DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			t := time.Now()
			if h, _, _ := t.Clock(); (h >= 17 && h <= 10) || (t.Weekday() == time.Saturday || t.Weekday() == time.Sunday) {
				return r, goproxy.NewResponse(r, goproxy.ContentTypeText, http.StatusForbidden, "who cares?!")
			}
			return r, nil
		})

	log.Fatalln(http.ListenAndServe(*listen, proxy))
}
