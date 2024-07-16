package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/soulwill1/estudos/api_products/internal/infra/akafka"
	"github.com/soulwill1/estudos/api_products/internal/infra/repository"
	"github.com/soulwill1/estudos/api_products/internal/infra/web"
	"github.com/soulwill1/estudos/api_products/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306/products)")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewProductRepositoryMysql(db)
	createProductUseCase := usecase.NewCreatProductUseCase(repository)
	listProductUsecase := usecase.NewListProductUseCase(repository)

	productHandlers := web.NewProductHandlers(createProductUseCase, listProductUsecase)

	r := chi.NewRouter()
	r.Post("/products", productHandlers.CreateProductHandler)
	r.Get("/products", productHandlers.ListProductsHandler)

	go http.ListenAndServe(":8000", r)
	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			//logar o erro
		}

		_, err = createProductUseCase.Execute(dto)
	}
}