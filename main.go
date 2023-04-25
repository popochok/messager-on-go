package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Message struct {
	UserName    string
	MessageText string
	TimeStamp   string
}

func main() {
	r := chi.NewRouter()
	fmt.Print("Service is running")
	messages := make([]Message, 0)

	r.Post("/api/Messager", func(w http.ResponseWriter, r *http.Request) {
		mes := Message{}
		ram, _ := ioutil.ReadAll((r.Body))
		json.Unmarshal(ram, &mes)
		messages = append(messages, mes)
	})

	r.Get("/api/Messager/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		id_num, _ := strconv.Atoi(id)
		if id_num >= 0 && id_num < len(messages) {
			msg := messages[id_num]
			data, _ := json.Marshal(msg)
			w.Write((data))
		}
	})
	http.ListenAndServe(":8080", r)
}
