package api

import (
	bootstrap2 "github.com/CastilloXavier/api_project_go/API/cmd/api/bootstrap"
	"log"
)

func main() {
	if err := bootstrap2.Run(); err != nil {
		log.Fatal(err)
	}
}
