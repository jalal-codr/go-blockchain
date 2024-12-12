package main

import (
	"fmt"
	"go-blockchain/blockchain"
	"go-blockchain/controllers"
	"go-blockchain/routes"
)

func main() {

	go blockchain.InitDB()
	fmt.Println("....Database initialized....")

	go controllers.CreateBlockchain()
	fmt.Println("....Blockchain created....")

	routes.StartServer()

}
