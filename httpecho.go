// Package aleshkin предназначен для выполнения заданий ItUniver
package aleshkin

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Функция SolutionHttpEcho выполняет возврат тела POST запроса
func SolutionHttpEcho(writer http.ResponseWriter, request *http.Request) {
	var response []byte
	var err error
	if request.Method == "POST" {
		response, err = ioutil.ReadAll(request.Body)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		response = []byte("Этот HTTP эхо сервис работает только с POST запросами")
	}
	_, err = writer.Write([]byte(response))
	if err != nil {
		log.Fatal(err)
	}
}
