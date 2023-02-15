package controller

import (
	"net/http"

	"github.com/artamananda/artanymous/app/model"
)

func (api *API) IndexPage(w http.ResponseWriter, r *http.Request) {
	tmpl := api.BaseViewPath()
	qData, err := api.messageRepo.ReadMessage()
	if err != nil {
		qData = []model.ViewMessage{}
	}
	data := []map[string]string{}
	for i := range qData {
		data = append(data, map[string]string{})
		data[i]["created_at"] = qData[len(qData)-1-i].CreatedAt.Format("2006-01-02  [15:04:05]")
		data[i]["question"] = qData[len(qData)-1-i].Question
	}
	tmpl.ExecuteTemplate(w, "message.html", data)
}
