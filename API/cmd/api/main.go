package api

import (
	bootstrap2 "api_project/API/cmd/api/bootstrap"
	"log"
)

func main() {
	if err := bootstrap2.Run(); err != nil {
		log.Fatal(err)
	}
}
