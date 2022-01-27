package BusinessLogic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type DatiPersona struct {
	Codice_fiscale string `json:"codice_fiscale"`
	Nome           string `json:"nome"`
	Cognome        string `json:"cognome"`
	Data_nascita   string `json:"data_nascita"`
	Via            string `json:"via"`
	Cap            string `json:"cap"`
	Citta          string `json:"citta"`
	Provincia      string `json:"provincia"`
}

type DatiCliente struct {
	Id              int         `json:"id"`
	Persona         DatiPersona `json:"persona"`
	Data_Iscrizione string      `json:"data_iscrizione"`
}

type DatiMedici struct {
	Cliente      DatiCliente `json:"cliente"`
	Id           int         `json:"id"`
	Sesso        string      `json:"sesso"`
	Peso         float32     `json:"peso"`
	Altezza      float32     `json:"altezza"`
	Numero_figli int         `json:"numero_figli"`
	Fumatore     bool        `json:"fumatore"`
}

func (amministratore Amministratore) AmministratoreDatiMedici() []DatiMedici {
	client := http.Client{}
	req, err := http.NewRequest("GET", AmministratoreDatiMedici, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", amministratore.Token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200: // la richiesta della lista di spese è andata a buon fine
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var result []DatiMedici
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Can not unmarshal JSON")
		}
		return result
	case 404: //non sono presenti dati medici di utenti
		fmt.Println("Non sono presenti dati medici")
		return nil
	default:
		fmt.Println("Errore non conosciuto: ", resp.StatusCode)
		return nil
	}
}

func StampaListaDatiMedici(listaSpese []DatiMedici) {
	fmt.Println("Dati medici presenti nel sistema: ")
	for i := 0; i < len(listaSpese); i++ {
		fmt.Println(listaSpese[i].Cliente.Persona.Codice_fiscale, " ", listaSpese[i].Cliente.Persona.Nome, " ", listaSpese[i].Cliente.Persona.Cognome, " ", listaSpese[i].Cliente.Persona.Data_nascita)
		fmt.Println("Sesso: ", listaSpese[i].Sesso,
			" Statura: ", listaSpese[i].Altezza,
			" cm Peso:", listaSpese[i].Peso,
			" kg Numero Figli: ", listaSpese[i].Numero_figli,
			" Fumatore: ",
			func(cond bool) string { //funzione anonima per la determinazione della condizione di fumatore a video
				if cond {
					return "SI"
				} else {
					return "NO"
				}
			}(listaSpese[i].Fumatore), "\n")
	}
}
