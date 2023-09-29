package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Definir o manipulador (handler) ao endpoint /createaccount
	http.HandleFunc("/createaccount", createAccountHandler)

	// Porta em que o servidor HTTP irá ouvir
	port := ":8080"

	// Inicia o servidor HTTP
	fmt.Printf("Servidor escutando em %s ... \n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}

}
