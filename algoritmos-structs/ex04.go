//Crie uma struct chamada Playlist com os campos "nome" e "músicas". O campo "músicas" deve ser
//um slice de structs, cada uma representando uma música com os campos "título", "artista"
//e "duração". Escreva uma função que receba uma Playlist como parâmetro e imprima
//o nome da playlist, o tempo total da playlist e as informações de cada música.

package main

import (
	"fmt"
)

type Playlist struct {
	Nome    string
	Musicas []InfoMusicas
}

type InfoMusicas struct {
	Titulo  string
	Artista string
	Duracao int
}

func main() {
	PlayList1 := Playlist{
		Nome: "PlayList - Avicii",
		Musicas: []InfoMusicas{
			{
				Titulo:  "The Nights",
				Artista: "Avicii",
				Duracao: 440,
			},
			{
				Titulo:  "Hey Brother",
				Artista: "Avicii",
				Duracao: 370,
			},
		},
	}

	PlayList2 := Playlist{
		Nome: "PlayList - Livros",
		Musicas: []InfoMusicas{
			{
				Titulo:  "48 Leis do Poder",
				Artista: "Robert Grenee",
				Duracao: 356,
			},
			{
				Titulo:  "Work 4h/Week",
				Artista: "Tim Ferriss",
				Duracao: 298,
			},
		},
	}

	PlaylistInfo(PlayList1)
	PlaylistInfo(PlayList2)
}

func PlaylistInfo(playlist Playlist) {
	var totalDuration int

	fmt.Println("Nome da PlayList: ", playlist.Nome)

	for _, musica := range playlist.Musicas {
		totalDuration += musica.Duracao
	}

	minutes := totalDuration / 60
	seconds := totalDuration % 60

	fmt.Printf("Duração total: %d:%.2d\n", minutes, seconds)
	fmt.Println("Informações: ", playlist.Musicas)
	fmt.Println("")
}
