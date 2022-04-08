package controllers

import (
	"net/http"

	"github.com/alifarhan1230/majoo-api-store/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {

	responses.JSON(w, http.StatusOK, responses.HomeResponse{StatusCode: 200, Success: true, Message: "Welcome to services majoo api store"})

}
