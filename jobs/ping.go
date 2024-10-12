package jobs

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func PingMicroService() {
	url := fmt.Sprintf("http://localhost:%s/ping", os.Getenv("PORT"))

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
}
