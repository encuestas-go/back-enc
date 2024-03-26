package main

import (
	"github.com/encuestas-go/back-enc/internal/database"
)

func main() {
	Build()

	surveyDB := database.ConnectToDB()

}
