package jobs

import (
	"log"
	"net/http"
	"os"
)

func PingMicroService() {
	url := os.Getenv("PING_URL)")

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
}
