package network

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
	"testing"
)

func TestWebSocket(t *testing.T) {
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/ws", myws)
	if err := http.ListenAndServe(":5000", router); err != nil {
		fmt.Println("err: ", err)
		return
	}
}

var h = hub{
	c: make(map[*connection]bool),
	u: make(chan *connection),
	b: make(chan []byte),
	r: make(chan *connection),
}

type hub struct {
	c map[*connection]bool
	b chan []byte
	r chan *connection
	u chan *connection
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.r:
			h.c[c] = true
			c.data.Ip = c.ws.RemoteAddr().String()
			c.data.Type = "handshake"
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			c.sc <- data_b
		case c := <-h.u:
			if _, ok := h.c[c]; ok {
				delete(h.c, c)
				close(c.sc)
			}
		case data := <-h.b:
			for c := range h.c {
				select {
				case c.sc <- data:
				default:
					delete(h.c, c)
					close(c.sc)
				}
			}
		}
	}
}

type connection struct {
	ws   *websocket.Conn
	sc   chan []byte
	data *Data
}

type Data struct {
	Ip       string   `json:"ip"`
	User     string   `json:"user"`
	Form     string   `json:"form"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	UserList []string `json:"userList"`
}

var wu = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func myws(w http.ResponseWriter, r *http.Request) {
	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	c := &connection{sc: make(chan []byte, 1024), ws: ws, data: &Data{}}
	h.r <- c
	go c.writer()
	defer func() {

	}()
}

func (c *connection) writer() {
	for msg := range c.sc {
		c.ws.WriteMessage(websocket.TextMessage, msg)
	}

	c.ws.Close()
}

var user_list = []string{}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			h.r <- c
			break
		}

		json.Unmarshal(message, &c.data)

		switch c.data.Type {
		case "login":
			c.data.User = c.data.Content
			c.data.Form = c.data.User
			user_list = append(user_list, c.data.User)
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
		case "user":
			c.data.Type = "user"
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
		case "logout":
			c.data.Type = "logout"
			user_list = del(user_list, c.data.User)
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
			h.r <- c
		default:
			fmt.Println("===defalut====")
		}

	}
}

func del(slice []string, user string) []string {
	count := len(slice)
	if count == 0 {
		return slice
	}

	if count == 1 && slice[0] == user {
		return []string{}
	}

	var n_slice = []string{}
	for i := range slice {
		if slice[i] == user && i == count {
			return slice[:count]
		} else if slice[i] == user {
			n_slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	fmt.Println(n_slice)
	return n_slice
}
