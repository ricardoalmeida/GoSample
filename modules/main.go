package main

import (
	"github/ricardoalmeida/webservice/controllers"
	"net/http"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
