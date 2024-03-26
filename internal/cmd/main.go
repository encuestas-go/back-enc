package main

import (
	"fmt"

	"github.com/encuestas-go/back-enc/internal/database"
)

func main() {
	Build()

	surveyDB := database.ConnectToDB()
	//Solo para que no de error de que no se usa la variable
	fmt.Println(surveyDB)

}
