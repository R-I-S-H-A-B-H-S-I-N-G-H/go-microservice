package error_util

import "log"

func Handle(message string, error error) {
	if error == nil {
		return
	}
	log.Fatal(message, error)
}
