package routes

import (
	"net/http"
	"webApp/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/remove", controllers.Remove)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/getUser", controllers.GetUsers)
}
