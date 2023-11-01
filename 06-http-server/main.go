package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
}

func main() {
	http.HandleFunc("/", FindCep)
	http.ListenAndServe(":8080", nil)
}

func FindCep(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cepResponse, err := GetCep(cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cepResponse)
}

func GetCep(cep string) (*CepResponse, error) {
	response, err := http.Get("https://viacep.com.br/ws/" + cep + "/json")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var result CepResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
