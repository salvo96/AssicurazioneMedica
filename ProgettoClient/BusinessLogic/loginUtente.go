package BusinessLogic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type responseLogin struct {
	Token           string `json:"token"`
	SessionLifetime string `json:"session_lifetime"`
}

//funzione che definisce le operazioni di Login dell'Utente: Cliente o Amministratore

func (utente *Utente) LoginUtente(email string, password string, endpoint string) bool { //questo metodo si occupa di effettuare la richiesta e farsi restituire il JSON con i dati
	loginData := map[string]string{"email": email, "password": password} //creo mappa (struttura dati) con email e password dell'Utente
	json_data, err := json.Marshal(loginData)                            //li converto in formato JSON
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200: // l'utente esiste e la password è corretta
		body, err := ioutil.ReadAll(resp.Body) // response body is []byte
		if err != nil {
			log.Fatal(err)
		}
		var result responseLogin
		if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}
		utente.Token = result.Token
		utente.Session_valid_until = result.SessionLifetime
		utente.Email = email
		utente.Password = password
		utente.Endpoint = endpoint
		return true
	case 404: // l'utente non esiste
		fmt.Println("L'Utente non è presente nel sistema!\n")
	case 400:
		fmt.Println("La password inserita non è corretta")
	default:
		fmt.Println("Errore non conosciuto: ", resp.StatusCode)
	}
	return false
}

func LoginMessage() (string, string) {
	fmt.Println("\nInserisci di seguito i tuoi dati personali per effettuare il login: ")
	email := ReadStringFormat("Email: ")
	password := ReadStringFormat("Password: ")
	return email, password
}
