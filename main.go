package main

import (
	"api/src/config"
	"api/src/router"

	//	"crypto/rand"
	//	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

// FUNÇÃO PARA GERAR UM TOKEN ALEATÓRIO PARA GUARDAR NO ARQUIVO ENV
//func init() {
//	chave := make([]byte, 64)
//	if _, erro := rand.Read(chave); erro != nil {
//		log.Fatal(erro)
//	}
//	stringBase64 := base64.StdEncoding.EncodeToString(chave)
//	fmt.Println(stringBase64)
//}

func main() {
	config.Carregar()
	r := router.Gerar()

	//fmt.Println(config.SecretKey)

	fmt.Printf("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
