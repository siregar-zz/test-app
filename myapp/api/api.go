package api

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"myapp/db"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func CheckCounter(w http.ResponseWriter, r *http.Request) {
	num := db.CounterProc()
	fmt.Fprintf(w, "Halaman telah dibuka sebanyak %+v kali.", num)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	users := db.Users{}
	
	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := db.Insert(users); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	users, err := db.GetOne(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	respondWithJson(w, http.StatusOK, users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if err := db.Remove(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func AddBarang(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	keranjang := db.Keranjang{}
	params := mux.Vars(r)
	id := params["id"]
	
	if err := json.NewDecoder(r.Body).Decode(&keranjang); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := db.TambahKeranjang(id, keranjang); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, keranjang)
}

func DelBarang(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	iduser := params["iduser"]
	idbarang := params["idbarang"]
	if err := db.HapusKeranjang(iduser, idbarang); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	all, err := db.GetAllUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, all)
}


