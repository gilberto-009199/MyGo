package main


import (
	"fmt"
	"time"
)

// iotas e declarando mais de um cosntante na mesma declaração

const (
	// Vai de 0 ate ..... * 10 para todas as cosntantes a baixo
	CONSTANTE0 = iota * 10
	CONSTANTE1 
	CONSTANTE2 
	CONSTANTE3 
)

func main(){

	fmt.Println(" ## Constante IOTA ##")
	fmt.Println(CONSTANTE0,CONSTANTE1,CONSTANTE2,CONSTANTE3);

	// if(){}else{}
	if 1 > 0 {
        fmt.Println(" 1 é maior que 0!")
    } else {
        fmt.Println(" 0 é maior que 1!")
	}

	// if(){}else if(){}
	if valor0 := 9; valor0 < 0 {
        fmt.Println(valor0, "is negative")
    } else if valor0 < 10 {
        fmt.Println(valor0, " e 1 digito")
    } else {
        fmt.Println(valor0, " tem multiplos digitos")
    }
	// Switch

	i := 2
    fmt.Print(" Esse valor ", i, " é ")
    switch i {
    case 1:
        fmt.Println("um")
    case 2:
        fmt.Println("dois")
    case 3:
        fmt.Println("trés")
	}
	
	// For 
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum) // 10 (1+2+3+4)
	// While
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n) // 8 (1*2*2*2)

	/* For infinite
	 *   sum = 1;
	 *   for {
	 *	 	 sum++ // repete para sempre
	 *	 }
	 */

	// For in range
	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
	// Break e continue 
	sum = 0;
	for i := 1; i < 5; i++ {
    	if i%2 != 0 { // vai para o proximo loop
    	    continue;
    	}
    	sum += i
	}
	fmt.Println(sum) // 6 (2+4)

	// Não exsite ternario mais tem um gambiarra com map kkkk
	valorDoTernario := map[bool]int{true: 1, false: 0} [5 > 4]
	fmt.Println(valorDoTernario)

	// Sleep na thread de excecução
	fmt.Printf("Unix Time: %v\n", time.Now().Unix())

    time.Sleep(2 * time.Second)

    fmt.Printf("Unix Time: %v\n", time.Now().Unix())
}