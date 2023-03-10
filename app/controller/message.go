package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/artamananda/artanymous/app/model"
	"github.com/artamananda/artanymous/config"
)

func (api *API) AddMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}
	var message model.Message
	message.Question = r.FormValue("question")

	err := api.messageRepo.AddMessage(message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(model.SuccessResponse{Msg: "Message Added"})
	http.Redirect(w, r, "/", http.StatusFound)
}

func (api *API) ReadMessage(w http.ResponseWriter, r *http.Request) {
	res, err := api.messageRepo.ReadMessage()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	if len(res) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Messsage not found!"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (api *API) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	err = api.messageRepo.DeleteMessage(uint(idInt))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Msg: "Message Delete Success"})
}

func (api *API) ResetMessage(w http.ResponseWriter, r *http.Request) {
	db := config.NewDB()
	conn, err := db.Connect()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = db.Reset(conn, "messages")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Msg: "Message Reset Success"})
}
