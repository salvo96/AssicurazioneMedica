package BusinessLogic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (cliente *Cliente) RegistrazioneUtente(codice_fiscale, nome, cognome, data_nascita, via, cap, citta, provincia, email, password string) {
	registerData := map[string]string{"codice_fiscale": codice_fiscale, "nome": nome, "cognome": cognome, "data_nascita": data_nascita, "via": via, "cap": cap, "citta": citta, "provincia": provincia, "email": email, "password": password} //creo mappa (struttura dati) con dati per registrazione Cliente
	json_data, err := json.Marshal(registerData)                                                                                                                                                                                              //conversione in formato JSON
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post(RegistrazioneCliente, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200: // la registrazione è avvenuta correttamente
		body, err := ioutil.ReadAll(resp.Body) // response body is []byte
		if err != nil {
			log.Fatal(err)
		}
		var result ResponseMessage
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Can not unmarshal JSON")
		}
		fmt.Println("MESSAGGIO DAL SERVER: ", result.Message)
	case 400:
		fmt.Println("Codice Fiscale e/o email già presenti nel sistema")
	default:
		fmt.Println("Errore non conosciuto: ", resp.StatusCode)
	}
}

func RegistrazioneMessage() (string, string, string, string, string, string, string, string, string, string) {
	fmt.Println("\nInserisci di seguito i dati per effettuare la registrazione: ")
	codice_fiscale := ReadStringFormat("Codice Fiscale: ")
	nome := ReadStringFormat("Nome: ")
	cognome := ReadStringFormat("Cognome: ")
	data_nascita := ReadStringFormat("Data di nascita (formato: YYYY-MM-DD): ")
	via := ReadStringFormat("Via: ")
	cap_ := ReadStringFormat("CAP: ")
	citta := ReadStringFormat("Città: ")
	provincia := ReadStringFormat("Provincia: ")
	email := ReadStringFormat("Email: ")
	password := ReadStringFormat("Password: ")
	return codice_fiscale, nome, cognome, data_nascita, via, cap_, citta, provincia, email, password
}
