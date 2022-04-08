package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/alifarhan1230/majoo-api-store/models"
	"github.com/alifarhan1230/majoo-api-store/responses"
)

func (server *Server) GetLaporan(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	transactions := models.Transaction{}

	trxReceived, err := transactions.FindLaporanByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, trxReceived)
}
