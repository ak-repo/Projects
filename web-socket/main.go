// Go server simplified for one-to-one chat
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type Message struct {
	From     string `json:"from"`
	To       string `json:"to,omitempty"`
	Content  string `json:"content"`
	UserType string `json:"user_type"`
}

type Server struct {
	mu    sync.Mutex
	users map[string]*websocket.Conn
	admin *websocket.Conn
}

func NewServer() *Server {
	return &Server{
		users: make(map[string]*websocket.Conn),
	}
}

func (s *Server) handleUser(ws *websocket.Conn) {
	userID := uuid.New().String()
	s.mu.Lock()
	s.users[userID] = ws
	s.mu.Unlock()

	// Send user their ID
	ws.Write([]byte(fmt.Sprintf(`{"from":"server","content":"Your ID: %s","user_type":"server"}`, userID)))

	fmt.Println("New user connected:", userID)

	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			continue
		}

		msg := Message{
			From:     userID,
			Content:  string(buf[:n]),
			UserType: "user",
		}
		s.sendToAdmin(msg)
	}

	// Cleanup on disconnect
	s.mu.Lock()
	delete(s.users, userID)
	s.mu.Unlock()
}

func (s *Server) handleAdmin(ws *websocket.Conn) {
	s.mu.Lock()
	s.admin = ws
	s.mu.Unlock()

	buf := make([]byte, 2048)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			continue
		}

		var msg Message
		if err := json.Unmarshal(buf[:n], &msg); err != nil {
			fmt.Println(err)
			continue
		}

		s.sendToUser(msg)
	}

	s.mu.Lock()
	s.admin = nil
	s.mu.Unlock()
}

func (s *Server) sendToAdmin(msg Message) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.admin == nil {
		return
	}
	data, _ := json.Marshal(msg)
	s.admin.Write(data)
}

func (s *Server) sendToUser(msg Message) {
	s.mu.Lock()
	defer s.mu.Unlock()
	userConn, ok := s.users[msg.To]
	if !ok {
		fmt.Println("User not found:", msg.To)
		return
	}
	data, _ := json.Marshal(msg)
	userConn.Write(data)
}

func main() {
	server := NewServer()
	http.Handle("/ws/user", websocket.Handler(server.handleUser))
	http.Handle("/ws/admin", websocket.Handler(server.handleAdmin))
	fmt.Println("Server running on :3000")
	http.ListenAndServe(":3000", nil)
}
