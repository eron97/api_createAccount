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

func createAccountHandler(w http.ResponseWriter, r *http.Request) {
	// lógica para lidar com a criação de uma conta.
	// Pode ler dados JSON de solicitação e persistir os dados
	// Enviar uma resposta de confirmação
}
