package spec

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"html/template"
)

const (
	wait           = 10 * time.Second
	pongWait       = 30 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 2048
)

type jsonMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

func newSpecSheetMessage(sheet *Sheet) *jsonMessage {
	return &jsonMessage{
		Type:    "spec",
		Payload: sheet,
	}
}

func newVerifierResultMessage(r bool) *jsonMessage {
	return &jsonMessage{Type: "verify", Payload: r}
}

func newKeepMessage() *jsonMessage {
	return &jsonMessage{Type: "keep", Payload: true}
}

type mainPageParameter struct {
	Host string
}

type connection struct {
	ws *websocket.Conn
}

type serverHandler struct {
	deliver     *Deliver
	verifier    *Verifier
	sheet       *Sheet
	connections map[*websocket.Conn]*connection
}

func newServerHandler(d *Deliver, v *Verifier) *serverHandler {
	return &serverHandler{
		deliver:     d,
		verifier:    v,
		connections: make(map[*websocket.Conn]*connection),
	}
}

func (h *serverHandler) addConnection(ws *websocket.Conn) {
	h.connections[ws] = &connection{ws}
}

func (h *serverHandler) removeConnection(ws *websocket.Conn) {
	delete(h.connections, ws)
}

func (h *serverHandler) cleanup(ws *websocket.Conn) {
	h.removeConnection(ws)
	ws.Close()
}

func (h *serverHandler) reader(ws *websocket.Conn) {
	defer h.cleanup(ws)

	for {
		var message jsonMessage
		err := ws.ReadJSON(&message)
		if err != nil {
			if err, ok := err.(*websocket.CloseError); ok && err.Code == websocket.CloseAbnormalClosure {
				return
			}
			//log.Println("websocket.ReadJSON err:", err)
			return
		}

		switch message.Type {
		case "verify":
			h.verifier.Request()
			r := <-h.verifier.ReceiveResult()
			err := ws.WriteJSON(newVerifierResultMessage(r))
			if err != nil {
				log.Println("write error:", err)
			}
		case "keep":
			// nothing.
		default:
			log.Println("unknown message type")
		}
	}
}

var upgrader = websocket.Upgrader{}

func (h *serverHandler) serveWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		//log.Println("upgrade error:", err)
		return
	}
	ws.SetReadLimit(maxMessageSize)
	// send current spec sheet.
	ws.WriteJSON(newSpecSheetMessage(h.sheet))

	h.addConnection(ws)
	h.reader(ws)
}

func (h *serverHandler) serveMain(w http.ResponseWriter, r *http.Request) {
	homeHTML, err := Asset("static/index.html")
	if err != nil {
		panic(err)
	}

	homeTemplate, err := template.New("index").Parse(string(homeHTML)) 
	if err != nil {
		panic(err)
	}

	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	params := &mainPageParameter{
		Host: r.Host,
	}

	homeTemplate.Execute(w, params)
}

func (h *serverHandler) deliverSpecSheet() {
	ticker := time.NewTimer(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case sheet := <-h.deliver.ReceiveSheet():
			h.sheet = sheet
			message := newSpecSheetMessage(sheet)
			for ws := range h.connections {
				ws.WriteJSON(message)
			}
		case <-ticker.C:
			message := newKeepMessage()
			for ws := range h.connections {
				err := ws.WriteJSON(message)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
