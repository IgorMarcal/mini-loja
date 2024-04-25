package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"webApp/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "index", todosProdutos)
}

// func BuscaPorId(w http.ResponseWriter, r *http.Request, id int) {
// 	idConvertido, err := strconv.Atoi(id)

// 	if err != nil {
// 		log.Println("Erro na conversao do preco", err)
// 	}

// 	product := models.BuscaPorId(id)
// 	log.Print(product)
// 	temp.ExecuteTemplate(w, "edit", product)
// }

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		panic("Formulario apresentou erros. Favor contactar suporte")
	}
	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco := r.FormValue("preco")
	quantidade := r.FormValue("quantidade")

	precoConvertido, err := strconv.ParseFloat(preco, 64)

	if err != nil {
		log.Println("Erro na conversao do preco", err)
	}

	quantidadeConvertido, err := strconv.Atoi(quantidade)

	if err != nil {
		log.Println("Erro na conversao de quantidade", err)
	}

	models.InsereProduto(nome, descricao, precoConvertido, quantidadeConvertido)

	http.Redirect(w, r, "/", 301)

}

func Remove(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		log.Println("Erro ao remover. ID nao encontrado!")
	}

	idConvertido, err := strconv.Atoi(id)

	if err != nil {
		log.Println("Erro na conversao do ID", err)
	}

	models.RemoveProduto(idConvertido)

	http.Redirect(w, r, "/", 301)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		idConvertido, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Erro na conversao do preco", err)
		}

		produto := models.BuscaPorId(idConvertido)
		temp.ExecuteTemplate(w, "edit", produto)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		if id == "" {
			log.Println("Erro ao editar. ID nao encontrado!")
		}

		idConvertido, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Erro na conversao do ID", err)
		}

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversao do preco", err)
		}

		quantidadeConvertido, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversao de quantidade", err)
		}

		models.EditaProduto(nome, descricao, precoConvertido, quantidadeConvertido, idConvertido)

		http.Redirect(w, r, "/", 301)
	}

}
