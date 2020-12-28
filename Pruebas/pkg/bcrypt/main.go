package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `clave123`
	encPw, err := bcrypt.GenerateFromPassword([]byte(s), 4)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("El hash es:", string(encPw))

	claveLogin := `clave123`
	err = bcrypt.CompareHashAndPassword(encPw, []byte(claveLogin))
	if err != nil {
		fmt.Println("Nombre de usuario o contraseña no son correctos")
	}

	fmt.Println("Eres el puto amo y estás logueado :* ")

}
