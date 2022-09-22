package app

import (
	// "encoding/json"
	"net/http"
)

type Server struct {
	router *http.ServeMux
}

func NewServer() *Server{
	s := &Server{
		http.NewServeMux(),
	}
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	s.router.ServeHTTP(w, r)
}

// func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
// 	w.WriteHeader(status)
// 	w.Header().Set("Content-Type", "application/json")

// 	if data != nil {
// 		if err := json.NewEncoder(w).Encode(data); err != nil {
// 			InternalErr.ServeHttp(w, r)
// 			fmt.Println("Error: ", err)
// 		}
// 	}
// }

// func (s *Server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {

// 	w.Header().Set("Content-Type", "application/json")
// 	if r.Body == nil {
// 		fmt.Println("body is nil:", r)
// 		return errors.New("body of request is nil")
// 	}
// 	return json.NewDecoder(r.Body).Decode(v)
// }

// func (s *Server) handleUnexpectedError(w http.ResponseWriter, req *http.Request) {
// 	if r := recover(); r != nil {
// 		fmt.Println("Error while handling error: ", r)
// 		InternalErr.ServeHttp(w, req)
// 	}
// }