package handler

import (
	"encoding/json"
	"fmt"
	"gomultithreading/internal/models"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func fetchViaCep(cep string, ch chan models.ViaCep) {
	defer close(ch)

	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		fmt.Println("error while requesting viacep: " + err.Error())
		return
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("error while reading viacep: " + err.Error())
		return
	}

	var viacep models.ViaCep
	err = json.Unmarshal(res, &viacep)
	if err != nil {
		fmt.Println("error while Unmarshal viacep: " + err.Error())
		return
	}

	ch <- viacep
}

func fetchBrasilApi(cep string, ch chan models.ApiCEP) {
	defer close(ch)

	req, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		fmt.Println("error while requesting brasilapi: " + err.Error())
		return
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("error while reading brasilapi: " + err.Error())
		return
	}

	var brasilapi models.ApiCEP
	err = json.Unmarshal(res, &brasilapi)
	if err != nil {
		fmt.Println("error while Unmarshal brasilapi: " + err.Error())
		return
	}

	ch <- brasilapi
}

func SearchCepHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cep := vars["cep"]

	if len(cep) != 8 && len(cep) <= 0 {
		http.Error(w, "CEP must be 8 valid numbers", http.StatusBadRequest)
		return
	}

	if _, err := strconv.Atoi(cep); err != nil {
		http.Error(w, "CEP must be a valid number", http.StatusBadRequest)
		return
	}

	chVia := make(chan models.ViaCep)
	chBrasil := make(chan models.ApiCEP)

	go fetchViaCep(cep, chVia)
	go fetchBrasilApi(cep, chBrasil)

	var message string

	select {
	case msgVia := <-chVia:
		if msgVia.Logradouro == "" || msgVia.Bairro == "" || msgVia.Localidade == "" || msgVia.Uf == "" {
			message := fmt.Sprintf("error while Unmarshal ViaCep: %s, %s, %s, %s some fields are empty", msgVia.Logradouro, msgVia.Bairro, msgVia.Localidade, msgVia.Uf)
			http.Error(w, message, http.StatusInternalServerError)
			return
		}
		message = fmt.Sprintf("Information Sending by: ViaCEP, Endereco: %s, %s, %s, %s - %s", msgVia.Logradouro, msgVia.Bairro, msgVia.Localidade, msgVia.Uf, msgVia.Cep)
	case msgBrasil := <-chBrasil:
		if msgBrasil.Street == "" || msgBrasil.Neighborhood == "" || msgBrasil.City == "" || msgBrasil.State == "" {
			message := fmt.Sprintf("error while Unmarshal BrasilAPI: %s, %s, %s, %s some fields are empty ", msgBrasil.Street, msgBrasil.Neighborhood, msgBrasil.City, msgBrasil.State)
			http.Error(w, message, http.StatusInternalServerError)
			return
		}
		message = fmt.Sprintf("Information Sending by: BrasilAPI, Endereco: %s, %s, %s, %s - %s", msgBrasil.Street, msgBrasil.Neighborhood, msgBrasil.City, msgBrasil.State, msgBrasil.Cep)
	case <-time.After(time.Second):
		http.Error(w, "Request timeout", http.StatusGatewayTimeout)
		return
	}

	fmt.Println(message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
