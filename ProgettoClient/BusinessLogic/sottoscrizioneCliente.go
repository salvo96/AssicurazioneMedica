package BusinessLogic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ResponseMessagePrezzo struct {
	ResponseMessage
	Prezzo string `json:"prezzo"`
}

func (cliente Cliente) SottoscrizioneCliente(sesso, peso, altezza, numero_figli, fumatore string) {
	sottoscrizioneData := map[string]string{"email": cliente.Email, "sesso": sesso, "peso": peso, "altezza": altezza, "numero_figli": numero_figli, "fumatore": fumatore} //creo mappa (struttura dati) con dati per creazione sottoscrizione
	json_data, err := json.Marshal(sottoscrizioneData)                                                                                                                    //conversione in formato JSON
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{}
	req, err := http.NewRequest("POST", SottoscrizioneCliente, bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", cliente.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200: // la richiesta di sottoscrizione è stata processata correttamente
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var result ResponseMessagePrezzo
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Can not unmarshal JSON")
		}
		fmt.Println("MESSAGGIO DAL SERVER: ", result.Message, result.Prezzo)
	default:
		fmt.Println("Errore non conosciuto: ", resp.StatusCode)
	}
}

func SottoscrizioneMessage() (string, string, string, string, string) {
	fmt.Println("\nInserisci di seguito i dati per effettuare una nuova richiesta di sottoscrizione: ")
	sesso := ReadStringFormat("Sesso: ")
	peso := ReadStringFormat("Peso: ")
	altezza := ReadStringFormat("Altezza: ")
	numero_figli := ReadStringFormat("Numero figli: ")
	var fumatore = func() string { //funzione anonima che converte 0 in 'false' e qualunque altro valore in 'true'
		res := ReadStringFormat("Fumatore? (0 = no, 1 = si): ")
		if res == "0" {
			return "False"
		} else if res == "1" {
			return "True"
		}
		return res
	}
	return sesso, peso, altezza, numero_figli, fumatore()
}
