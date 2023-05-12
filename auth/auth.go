package auth

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Auth(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		http.Error(writer, "Request failed", http.StatusInternalServerError)
		return
	}
	log.Println(string(body))
	fmt.Fprint(writer, string(body))
}
