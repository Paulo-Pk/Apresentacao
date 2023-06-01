package main

import (
	"fmt"
	"time"
)

func typeEffect(text string) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(50 * time.Millisecond)
	}
}
func main() {
	message := "Olá, sou o Paulo e seja bem-vindo à este Humilde Perfil onde há um pouquinho da minha trajetória profissional, conquistas e paixões... Estou comprometido em buscar diáriamente o aprendizado como também a excelência e colaborar em projetos desafiadores. Convido você a explorar meu perfil e descobrir como podemos conectar nossos interesses e objetivos. Juntos construir oportunidades de sucesso!!"
	typeEffect(message)
}
