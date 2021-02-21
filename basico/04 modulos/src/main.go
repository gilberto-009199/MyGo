package main

import (
	"Myproject/src/model" // será necessario usar o "go mod init Myproject" no terminal
	"Myproject/src/util" // será necessario usar o "go mod init Myproject" no terminal
    "fmt"
)

func main() {
    fmt.Printf("Main Function");

    var employee = model.Usuario{
        Uuid:"546545-5645152-54865-i3724278",
        Nome:"Gabriel",
	}
	

	fmt.Println(employee.Nome);
	
	fmt.Println(util.Sum(8,16))
	fmt.Println(util.PI)

}