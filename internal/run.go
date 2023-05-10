package internal

import (
	"VK_BOT_GO/internal/delivery/http/vk"
	"flag"
	"log"
	"net/http"
)

func Run() {
	var str string // Строка, которую должен вернуть сервер
	flag.StringVar(&str, "c", "", "check server")
	flag.Parse()

	mux := http.NewServeMux()
	server := vk.New()
	mux.HandleFunc("/", server.Check(str))

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
