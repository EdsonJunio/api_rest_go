package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

var usuarios []Usuario

func main() {
	r := mux.NewRouter()

	usuarios = append(usuarios, Usuario{ID: 1, Nome: "João Silva", Email: "joao@exemplo.com"})
	usuarios = append(usuarios, Usuario{ID: 2, Nome: "Maria Souza", Email: "maria@exemplo.com"})

	r.HandleFunc("/usuarios", listarUsuarios).Methods("GET")
	r.HandleFunc("/usuarios/{id}", buscarUsuario).Methods("GET")
	r.HandleFunc("/usuarios", criarUsuario).Methods("POST")
	r.HandleFunc("/usuarios/{id}", atualizarUsuario).Methods("PUT")
	r.HandleFunc("/usuarios/{id}", deletarUsuario).Methods("DELETE")

	log.Println("API rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func listarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

func buscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, value := range usuarios {
		if value.ID == id {
			json.NewEncoder(w).Encode(value)
		}
	}
}

func criarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var novoUsuario Usuario
	_ = json.NewDecoder(r.Body).Decode(&novoUsuario)
	novoUsuario.ID = len(usuarios)
	add := append(usuarios, novoUsuario)
	json.NewEncoder(w).Encode(add)

}

func atualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, value := range usuarios {
		if value.ID == id {
			var updatedUsuario Usuario
			_ = json.NewDecoder(r.Body).Decode(&updatedUsuario)
			updatedUsuario.ID = id
			usuarios[index] = updatedUsuario
			json.NewEncoder(w).Encode(updatedUsuario)

		}
	}

}

func deletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, usuario := range usuarios {
		if usuario.ID == id {
			usuarios = append(usuarios[:index], usuarios[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}
