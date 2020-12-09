package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	writer      http.ResponseWriter
}

func CreateDefaultResponse(w http.ResponseWriter) Response {
	return Response{Status: http.StatusOK, writer: w, contentType: "application/json"}
}

func (this *Response) NotFound() {
	this.Status = http.StatusNotFound
	this.Data = nil
	this.Message = "Resources Not Found."
}

func (this *Response) UnprocessableEntity() {
	this.Status = http.StatusUnprocessableEntity
	this.Data = nil
	this.Message = "UnprocessableEntity"
}

func SendUnprocessableEntity(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.UnprocessableEntity()
	response.Send()
}

func SendNotFound(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotFound()
	response.Send()
}

func (this *Response) Send() {
	this.writer.Header().Set("Content-type", this.contentType)
	this.writer.WriteHeader(this.Status)

	outpout, _ := json.Marshal(&this)
	fmt.Fprintf(this.writer, string(outpout))
}

func SendData(w http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(w)
	response.Data = data
	response.Send()
}
