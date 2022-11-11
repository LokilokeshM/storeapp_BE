package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/storeapp/cmd/product/domain"
	"github.com/storeapp/cmd/product/service"
)

func Start() {
	r := mux.NewRouter()
	client := &http.Client{}
	//wiring
	var d = Handlers{
		Service: service.NewService(domain.NewHttpRepository(client)),
	} // define routes

	fmt.Println("Started Service")
	port := "8080"
	fmt.Println(port)
	r.HandleFunc("/getCollectList", d.GetCollectList).Methods(http.MethodGet)
	r.HandleFunc("/getCollectionList", d.GetCollectionList).Methods(http.MethodGet)

	r.HandleFunc("/addNewProduct", d.AddNewProduct).Methods(http.MethodPost)
	r.HandleFunc("/getProductList", d.GetProductList).Methods(http.MethodGet)
	// r.HandleFunc("/getProduct", d.GetProduct).Methods(http.MethodPost)
	// r.HandleFunc("/getProductVariant", d.GetProductVariant).Methods(http.MethodGet)
	// r.HandleFunc("/getProductVariantList", d.GetProductVariantList).Methods(http.MethodGet)
	r.HandleFunc("/getProductImageList", d.GetProductImageList).Methods(http.MethodGet)
	// r.HandleFunc("/addProductImage", d.AddProductImage).Methods(http.MethodPost)
	r.HandleFunc("/updateProductImage", d.UpdateProductImage).Methods(http.MethodPut)
	// port := os.Getenv("PORT")
	// fmt.Print(port)
	// if port == "" {
	// 	port = "8080"
	// }

	log.Fatal(http.ListenAndServe(":"+port, r))
}
