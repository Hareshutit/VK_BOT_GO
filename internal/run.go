package internal

import (
	"VK_BOT_GO/internal/delivery/http/vk"
	"flag"
	"log"
	"net/http"
	"strconv"
)

func Run() {
	var str int // Строка, которую должен вернуть сервер
	flag.IntVar(&str, "c", 0, "check server")
	flag.Parse()

	mux := http.NewServeMux()
	server := vk.New()
	mux.HandleFunc("/", server.Check(strconv.Itoa(str)))

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
