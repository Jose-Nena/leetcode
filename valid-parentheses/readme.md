Leetcode
20. Valid Parentheses
Easy

--> Minha implementação
--> Implementação YouTuber:
https://www.youtube.com/watch?v=cNEfSzs1ijs&list=PLPHf0lj0nCz9dtE3isR_R8r7ohVNpH_Tb&index=22

--------------------------------------------------------------------------------------------
Observações:

Na função leftParenthesesStack(parentheses string, leftStack []string) ([]string, bool), tive que retornar o slice leftStack modificado, pois descobri que quando se usa a função append() para inserir um elemento, aparentemente o compilador cria uma cópia deste slice, e o original(fora da função, no bloco chamador) não é modificado. Não testei se o mesmo comportamento acontece quando elimino um ou mais elemento do slice passado como argumento. 
P.S. -  Mais tarde, testei fazendo eliminação do último elemento do slice dentro da função chamada, e o slice no bloco chamador (fora desta função chamada) não foi modificado por esta ação. As únicas ações que vi funcionar foram aquelas que modificam o valor dos elementos do slice dentro da função chamada. Fora estas mudanças, não identifiquei outras que afetem o slice original (passado como argumento para a função que o modifica). Em outras palavras, o slice pode até ser passado com referência, mas isso não se aplica em todos os casos.
Outra dificuldade que encontrei, foi pra deletar o último elemento de um slice passado como ponteiro pra uma função. Deu erro, pois a própria IDE apontou que não é possível fazer slice de um ponteiro pra um slice, ou algo do tipo. Quando puder, seria interessante verificar isso.


Isso funciona:

package main

import "fmt"

func mudaSlice(slice []string) {
	slice[0] = "b"
	slice = append(slice, "c")
	fmt.Print(slice)
}

func teste() {
	stack := []string{"a"}
	fmt.Print(stack)
	mudaSlice(stack)
	fmt.Print(stack)
}

func main() {
	//slice := []string{"a", "a"}
	teste()
}

Output:
[a][b c][b]