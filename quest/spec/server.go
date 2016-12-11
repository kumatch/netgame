package spec

import (
	"net/http"
)

var (
	listenAddress = ":18080"
)

func runServer(d *Deliver, v *Verifier) {
	handler := newServerHandler(d, v)
	http.HandleFunc("/", handler.serveMain)
	http.HandleFunc("/ws", handler.serveWebSocket)
	go handler.deliverSpecSheet()
	http.ListenAndServe(listenAddress, nil)
}
