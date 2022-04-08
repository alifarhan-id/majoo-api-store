package controllers

import "github.com/alifarhan1230/majoo-api-store/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//laporan route
	s.Router.HandleFunc("/laporan/{id}", middlewares.SetMiddlewareJSON(s.GetLaporan)).Methods("GET")
}
