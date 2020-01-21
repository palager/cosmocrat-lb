package main

import "github.com/google/tcpproxy"
import "log"

func main() {
	var p tcpproxy.Proxy

	p.AddHTTPHostRoute(":80", "upspin.cosmocr.at", tcpproxy.To("upspin-backend.default.svc.cluster.local:80"))
	p.AddSNIRoute(":443", "upspin.cosmocr.at", tcpproxy.To("upspin-backend.default.svc.cluster.local:443"))

	p.AddHTTPHostRoute(":80", "135ppw.nyc", tcpproxy.To("cosmocrat-backend.default.svc.cluster.local:80"))
	p.AddSNIRoute(":443", "135ppw.nyc", tcpproxy.To("cosmocrat-backend.default.svc.cluster.local:443"))

	p.AddHTTPHostRoute(":80", "135ppw.org", tcpproxy.To("cosmocrat-backend.default.svc.cluster.local:80"))
	p.AddSNIRoute(":443", "135ppw.org", tcpproxy.To("cosmocrat-backend.default.svc.cluster.local:443"))

	p.AddRoute(":8008", tcpproxy.To("ssb-pub.scuttlebutt.svc.cluster.local:8008"))
	p.AddHTTPHostRoute(":80", "ssb.cosmocr.at", tcpproxy.To("ssb-viewer.scuttlebutt.svc.cluster.local:8080"))
	p.AddSNIRoute(":443", "ssb.cosmocr.at", tcpproxy.To("ssb-viewer.scuttlebutt.svc.cluster.local:8443"))

	p.AddHTTPHostRoute(":80", "cosmocr.at", tcpproxy.To("cosmocrat-backend.default.svc.cluster.local:80"))
	p.AddSNIRoute(":443", "cosmocr.at", tcpproxy.To("cosmocrat-backend.default.svc.cluster.local:443"))
	log.Fatal(p.Run())
}
