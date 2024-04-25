package models

import (
	"webApp/db"
)

type Produto struct {
	Nome           string
	Descricao      string
	Preco          float64
	Quantidade, Id int
}

func BuscaTodosOsProdutos() []Produto {

	db := db.ConectaComBancoDeDados()
	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos

}

func BuscaPorId(idConvertido int) Produto {

	db := db.ConectaComBancoDeDados()

	row := db.QueryRow("SELECT * FROM produtos WHERE id = $1", idConvertido)
	p := Produto{}

	err := row.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	return p
}

func InsereProduto(nome, descricao string, precoConvertido float64, quantidadeConvertido int) {

	db := db.ConectaComBancoDeDados()
	insereDadosBanco, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosBanco.Exec(nome, descricao, precoConvertido, quantidadeConvertido)
	defer db.Close()

}

func RemoveProduto(id int) {

	db := db.ConectaComBancoDeDados()
	insereDadosBanco, err := db.Prepare("delete from produtos where id = ($1)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosBanco.Exec(id)
	defer db.Close()

}

func EditaProduto(nome, descricao string, precoConvertido float64, quantidadeConvertido, id int) {

	db := db.ConectaComBancoDeDados()
	insereDadosBanco, err := db.Prepare("update produtos set nome = ($1), descricao = ($2), preco = ($3), quantidade = ($4) where id = ($5)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosBanco.Exec(nome, descricao, precoConvertido, quantidadeConvertido, id)

	defer db.Close()

}
