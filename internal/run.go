package internal

import (
	"VK_BOT_GO/internal/delivery/http/vk"
	"log"
	"net/http"
	"os"
)

func Run() {
	var str string // Строка, которую должен вернуть сервер
	str = os.Getenv("STR")

	mux := http.NewServeMux()
	server := vk.New()
	mux.HandleFunc("/", server.Check(str))

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
