package app

import (
	"fmt"
	"net/http"
)



func (s *Server) routes(){
	s.router.HandleFunc("/health-check", s.handleHealthCheck())
}

func (s *Server) handleHealthCheck() http.HandlerFunc{
	fmt.Println("\nhealth-check endpoint RUNNING")

	return func(w http.ResponseWriter, r *http.Request){
		fmt.Println("\nhealth-check hit")
		_, err := w.Write([]byte("HEALTH CHECK OK"))
		if err != nil {
			fmt.Println("Error occured during health check: ", err)
			InternalErr.ServeHttp(w, r)
		}
	}
}