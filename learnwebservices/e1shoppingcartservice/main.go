package main

import (
	"context"
	svc "e1shoppingcartservice/services"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	cs := svc.CreateCustomerService()
	ps := svc.CreateProductService()
	scs := svc.CreateShoppingCartService()

	go func() {
		cs.ListenAndServe()
	}()

	go func() {
		ps.ListenAndServe()
	}()

	go func() {
		scs.ListenAndServe()
	}()

	time.Sleep(1 * time.Second)
	res, err := http.Get("http://localhost:5005/api/carts")
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response: ", string(data))

	fmt.Println("Services started, press <Enter> to shutdown")
	fmt.Scanln()

	cs.Shutdown(context.Background())
	ps.Shutdown(context.Background())
	scs.Shutdown(context.Background())

	fmt.Println("Services stopped")
}
