//Usando as mesmas structs do exercício anterior, escreva uma função que receba um título de música
//como parâmetro e retorne uma lista das Playlists que possuem esse título.

package main

import "fmt"

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
		Nome: "PlayLis-Avicii",
		Musicas: []InfoMusicas{
			{
				Titulo:  "The Nights",
				Artista: "Avicii",
				Duracao: 440,
			},
			{
				Titulo:  "abc",
				Artista: "Avicii",
				Duracao: 370,
			},
		},
	}

	PlayList2 := Playlist{
		Nome: "PlayList-Livros",
		Musicas: []InfoMusicas{
			{
				Titulo:  "48 Leis do Poder",
				Artista: "Robert Grenee",
				Duracao: 356,
			},
			{
				Titulo:  "abc",
				Artista: "Tim Ferriss",
				Duracao: 298,
			},
		},
	}

	PlayListList := []Playlist{PlayList1, PlayList2}

	//mecanismo de busca de playlists do spotify por musicas
	fmt.Print(HaveSameMusics("abc", PlayListList))
}

func HaveSameMusics(titulo string, playlists []Playlist) []string {
	var result []string

	for _, playlist := range playlists {
		for _, musicas := range playlist.Musicas {
			if titulo == musicas.Titulo {
				result = append(result, playlist.Nome)
			}
		}
	}

	return result
}
