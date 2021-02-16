package main
// Rune é apelido para int32
import (
	"fmt"
	"log"
	"strconv"
)

// Constante #define MINHA_COSNTANTE 10
const MINHA_CONSTANTE = 10;

func main(){

	/* Printf
	 *	Formatação
	 *		\n => pula linha
	 *		\t => pula um Tab no texto
	 *  Concatenação
	 		 %v	 => Conteudo da variavel seja string, int ou float
			 %b  => Valor em binario 
			 %d  => Valor em decimal
			 %#x => Valor em hexadecimal
			 %0.2f => Valor em float com 2 casas decimais
 	 */
	fmt.Println("##  CONSTANTE  ##");
	fmt.Printf(" Tipo : %T \n  Valor: %v \n",MINHA_CONSTANTE, MINHA_CONSTANTE);
	/* Criação de variavel com o Marmota 
	*	Para criar uma variavel e obrigatorio fazer a inicialização dela com um valor
	*  pois ó Go necessita disso para determinar o tipo da variavel
	*  		nomeDaVariavel 							:= 										0;
	*	  <Nome da variavel> <Operador de Atribuição Marmota, ou Normal (:=,=) >	<Valor da variavel>		  
	*/
	nomeDaVariavel := 1; // Define uma variavel segundo o dado atribuito
	
	/* Criação de variavel definindo o tamanho de memoria
	 *	
	 * 	var    nomeDaVariavel	 uint					=	 						0
	 *	var  <Nome da variavel> <Tipo> <Operador de Atribuição Normal (=) >	<Valor da variavel>		  
	 *							 uint8        8-bit	 inteiros (0 to 255)
	 *							 uint16       16-bit inteiros (0 to 65535)
	 *							 uint32       32-bit inteiros (0 to 4294967295)
	 *							 uint64       64-bit inteiros (0 to 18446744073709551615)
	 *							 int8         8-bit	 inteiros (-128 to 127)
	 *							 int16        16-bit inteiros (-32768 to 32767)
	 *							 int32        32-bit inteiros (-2147483648 to 2147483647)
	 *							 int64        64-bit inteiros (-9223372036854775808 to 9223372036854775807)
	 *							 float8       8-bit	 inteiros (-128 to 127)
	 *							 float16      16-bit inteiros (-32768 to 32767)
	 *							 float32      32-bit inteiros (-2147483648 to 2147483647)
	 *							 float64      64-bit inteiros (-9223372036854775808 to 9223372036854775807)
	 *
	 *	@Atenção!  Cuidado para não Exceder o limite de cada variavel
	 *
	 */
	var variavelPreDefinida uint = 10;
	fmt.Println("##  Numeros  ##");
	fmt.Printf(" Tipo : %T \n  Valor: %v \n",nomeDaVariavel, nomeDaVariavel);
	fmt.Printf(" Tipo : %T \n  Valor: %v \n",variavelPreDefinida, variavelPreDefinida);

	/*  Tipo de variaveis Numericos
	 *		Int		1
	 *		float	-459.67
	 *	Go tem dois tipos de dados numéricos que são 
	 *	distinguidos pela natureza estática ou dinâmica
	 * 	dos seus tamanhos
	 *	1° Tipo 
	 *		val0 int32 = idependente da arquitetura( Tamanho cosntante para diversos sistemas)
	 *		val0 int64 = idependente da arquitetura( Tamanho cosntante para diversos sistemas)
	 *	2° Tipo
	 *		val0 uint =	varia entre inteiros de 32 ou 64 bits dependendo da implementação(Compilação para a plataforma)
	 *		val0 int  =	varia entre inteiros de 32 ou 64 bits dependendo da implementação(Compilação para a plataforma)
	 *   @Atenção!! Você pode sim usar um int64 ou uint64 e compilar para uma arquitetura de 32 bits
	 *	no entanto a alocação de memoria fara no minimo o processador ter que demorar o dobro para realizar
	 *	operações com ele, poisa CPU vai ter que gstar mais ciclos para mover os dados pela arquitetura
	 *    Se vc deixar como um int na hora da compilação ele ira definir com base na arquitetura de build
	 */
	
	// Variaveis Boolenas
	valorBoleana :=  true;
	valorBoleanaExpressao := 15 < 10;
	fmt.Println("##  Boleanas  ##");
	fmt.Printf(" Tipo : %T \n  Valor: %v \n",valorBoleana, valorBoleana);
	fmt.Printf(" Tipo : %T \n  Valor: %v \n",valorBoleanaExpressao, valorBoleanaExpressao);

	/* Variaveis String
	 *  Dependendo de como vc criaa  String ela pode ser:
	 *		1°	string bruta		=> ` Hellow "World" `
	 *			 print(` Hellow "World" \n,Go`) =>
	 *				 Console =>	
	 *						Hellow "World"\n, Go
	 *
	 *		2°	string interpretada	=> " Hellow 'World' "
	 *			 print(" Hellow 'World' \n,Go") =>
	 *			 	Console =>	
	 *						Hellow 'World'
	 *			 			, Go 
	 */
	 texto := ` Hellow "World" \n,Go`;
	 textoLiteral := " Hellow 'World' \n\t,Go";

	fmt.Println("##  Strings  ##");
	fmt.Printf(" Tipo : %T \n  Valor: %v \n", texto, texto);
	fmt.Printf(" Tipo : %T \n  Valor: %v \n", textoLiteral, textoLiteral);

	/* 	A linguagem Go já vem totalmente compatível com os caracteres UTF-8,
	 * sem qualquer configuração especial
	 */
	TextoComIdeograma := "Hello, 世界";
	fmt.Printf(" Tipo : %T \n  Valor: %v \n", TextoComIdeograma, TextoComIdeograma);

	/*	# Matrizes
	 *   As matrizes tems eu tamanho predefinido e são declaradas dessa forma:
	 *		frutas := [5]string{"abacaxi", "amora", "banana", "caqui", "carambola"}
	 */
	frutas := [5]string{"abacaxi", "amora", "banana",  "caqui", "carambola"}
	fmt.Println("##  matrizes  ##");
	fmt.Printf(" Tipo : %T \n  Valor: %v \n", frutas, frutas);

	/*	# Fatias
	 *   As fatias tem seu tamanho não predefinido e são declaradas dessa forma:
	 *		cha := []string{"Chá verde", "Chá preto", "Chá de camomila",  "Chás de erva cidreira", "Chás de hortelã"}
	 *		cha = append(cha, "Chá de boldo"); // Adicionado o Chá de Boldo		 
	 */
	cha := []string{"Chá verde", "Chá preto", "Chá de camomila",  "Chás de erva cidreira", "Chás de hortelã"}
	cha = append(cha,"Chá de boldo"); // Adicionado o Chá de Boldo

	fmt.Println("##  Fatias  ##");
	fmt.Printf(" Tipo : %T \n  Valor: %v \n", cha, cha);

	/*	# Mapas
	 *   Os mapas são como um hash ou dicionario usando Chave e valor para armazenar a informação:
	 *		usuario := map[string]string{"id": "1", "nome": "gil", "dtNasc": "1999-06-02", "location": "brazil"}
	 *		usuario["nome"] => 
	 *			 Console =>
	 *			 	Gil
	 */
	usuario := map[string]string{"id": "1", "nome": "gil", "dtNasc": "1999-06-02", "location": "brazil"}
	fmt.Println("##  Map(Mapas)  ##");
	fmt.Printf(" Tipo : %T \n  Valor: %v \n", usuario, usuario);
	fmt.Printf(" Tipo : %T \n  Valor: %v \n", usuario["nome"], usuario["nome"]);

	// # Conversions
	b := 125.56
	c := 390.9
	fmt.Println("##  Conversoes  ##");
	fmt.Printf(" Tipo : %T \n  Valor: %v \n", int(b), int(b));
	fmt.Printf(" Tipo : %T \n  Valor: %v \n", int(c), int(c));
	// 		Concatenando 
	user := "Sammy"
    linhas := 50

	fmt.Println("Parabens, " + user + "! Você excreveu " + strconv.Itoa(linhas) + " linhas de codigo.")
	// 		Calculando a partir de um String https://pkg.go.dev/strconv

	linhas_ontem := "50"
    linhas_hoje := "108"

	// linhas_ontem - linhas_hoje => Erro de operação 
	// Tenta converter "50" para => 50
	ontem, err := strconv.Atoi(linhas_ontem)
	// Verifica se ocorreu um erro na conversao apra inteiro
    if err != nil {
        log.Fatal(err)
    }
	// Tenta converter "108" para => 108
	hoje, err := strconv.Atoi(linhas_hoje)
	// Verifica se ocorreu um erro na conversao apra inteiro
    if err != nil {
        log.Fatal(err)
    }
    var linhasAhMaisHoje int = hoje - ontem

    fmt.Println(linhasAhMaisHoje)

}