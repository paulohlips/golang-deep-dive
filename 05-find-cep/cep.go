package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while was searching CEP: %v", err)
		}

		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while was reading req.body: %v", err)
		}

		var response CepResponse
		err = json.Unmarshal(res, &response)
		println(string(res))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while was unmarshalling response: %v", err)
		}

	}
}
