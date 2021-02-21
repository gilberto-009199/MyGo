package main

import (
	"fmt"
)

/* Criando Funções 
 *
 * func <nome da função>( parametro <tipo do parametro> ) <Tipo de retorno>{
 *	 // Conteudo da função
 * }
 *
 */
func sum( val0 float64, val1 float64) float64{
	return val0 + val1;
}

// É possivel deixar para depois o tipo do parametro para o ultimo 
func sum2( val0 /* <Deixado o tipo para o ultimo parametro> */, val1 float64) float64 {
	return val0 + val1;
}
// Multiplos resultados
func somaUm(val0, val1 float64) (float64, float64) {
	return val0 + 1, val1 +1; 
}
// Multiplos resultados com nomes
func somaUm2(val0, val1 float64) ( v0, v1 float64) {
	v0 = val0 + 1;
	v1 = val1 +1
	return // Vai montar o return 
}
//Criando um strut 
type Retangulo struct {
	largura int
	altura int
};

type Endereco struct {
	uf string
	bairro string
	cidade string
	n int
	logradouro,	complemento string
};
type Usuario struct {
	uuid string
	nome string
	email string
	endereco Endereco
};

func main(){
	fmt.Println(" SUM : ");
	fmt.Printf("  %v + %v = %.2f\n", 4, 7 , sum(4,7));

	retan := Retangulo{10, 6}

	fmt.Println(retan);


	user := Usuario{};
	user.uuid = "3bff3d56-0dfa-474f-9260-0a22b11db21c";
	user.nome = "gil";
	user.endereco = Endereco{ uf:"SP",
							  bairro:"JD. Silveira",
							  cidade:"Barueri",
							  n:446,
							  logradouro:"Rua Brigadeiro Jordão",
							  complemento:" Escadão",
							};
	fmt.Println(user);
	user = Usuario{ uuid:"b68d4c6b-819d-4c75-8817-6bc524d9e324",
					 nome:"Guilherme",
					 email:"guilherme@gmail.com",
					 endereco: Endereco{
							uf:"MG",
							bairro:"Montes Claros",
							cidade:"Itacarambi",
							n:569,
							logradouro:"Rua Almeida Freire",
							complemento:" 2 casas depois da Padaria",
					 	},
					};
	fmt.Println(user);

}