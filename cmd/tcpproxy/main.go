package main

import "github.com/google/tcpproxy"
import "log"

func main() {
	var p tcpproxy.Proxy

	p.AddHTTPHostRoute(":80", "upspin.cosmocr.at", tcpproxy.To("upspin-backend.default.svc.cluster.local:80"))
	p.AddSNIRoute(":443", "upspin.cosmocr.at", tcpproxy.To("upspin-backend.default.svc.cluster.local:8443"))

	p.AddHTTPHostRoute(":80", "135ppw.nyc", tcpproxy.To("cosmocrat-backend.default.svc.cluster.local:80"))
	p.AddSNIRoute(":443", "135ppw.nyc", tcpproxy.To("cosmocrat-backend.default.svc.cluster.local:8443"))

	p.AddHTTPHostRoute(":80", "135ppw.org", tcpproxy.To("cosmocrat-backend.default.svc.cluster.local:80"))
	p.AddSNIRoute(":443", "135ppw.org", tcpproxy.To("cosmocrat-backend.default.svc.cluster.local:8443"))

	p.AddRoute(":8008", tcpproxy.To("ssb-pub.scuttlebutt.svc.cluster.local:8008"))
	log.Fatal(p.Run())
}
