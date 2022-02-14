package main

import "fmt"

type User struct {
	nombre string
	email  string
	activo bool
}

type Student struct {
	user   User
	codigo string
}

//Relacion de uno a muchos
type Curso struct {
	titulo string
	videos []Video
}

type Video struct {
	titulo string
	curso  Curso
}

func main() {

	//Relacion de uno a uno
	/*victor := User{
		nombre: "Victor",
		email:  "victor05@hotmail.com",
		activo: true,
	}
	juan := User{
		nombre: "Juan",
		email:  "juan32@hotmail.com",
		activo: true,
	}

	alexStudent := Student{
		user:   victor,
		codigo: "001ard",
	}

	fmt.Println(juan)
	fmt.Println(alexStudent)
	*/
	video1 := Video{titulo: "01-Introduccion"}
	video2 := Video{titulo: "02-Instalacion"}

	curso := Curso{
		titulo: "curso de GO",
		videos: []Video{video1, video2},
	}
	video1.curso = curso
	video2.curso = curso
	fmt.Println(video1.curso.titulo)
	for _, video := range curso.videos {
		fmt.Println(video.titulo)
	}
}
