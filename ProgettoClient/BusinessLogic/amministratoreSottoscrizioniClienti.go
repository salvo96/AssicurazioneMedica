package BusinessLogic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RichiestaSottoscrizione struct {
	Id             int     `json:"id"`
	Data           string  `json:"data"`
	Evasa          bool    `json:"evasa"`
	Prezzoproposto float32 `json:"prezzoproposto"`
	DatiMediciId   int     `json:"dati_medici"`
}

func (amministratore Amministratore) AmministratoreSottoscrizioniClienti() []RichiestaSottoscrizione {
	client := http.Client{}
	req, err := http.NewRequest("GET", AmministratoreSottoscrizioniClienti, nil)
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
	case 200: // la richiesta della lista di sottoscrizioni inoltrate dai Clienti è andata a buon fine
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var result []RichiestaSottoscrizione
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Can not unmarshal JSON")
		}
		return result
	case 404: //non sono presenti richieste di sottoscrizione
		fmt.Println("Non sono presenti richieste di sottoscrizione Polizza inoltrate dai Clienti")
		return nil
	default:
		fmt.Println("Errore non conosciuto: ", resp.StatusCode)
		return nil
	}
}

func StampaListaSottoscrizioniClienti(listaSottoscrizioni []RichiestaSottoscrizione, listaDatiMedici []DatiMedici) bool {
	fmt.Println("Richieste di sottoscrizione Polizza Assicurativa presenti nel sistema: ")
	var sottoscrizioniPresenti = false
	for _, richiesta := range listaSottoscrizioni {
		if !richiesta.Evasa {
			sottoscrizioniPresenti = true
			for _, datiMedici := range listaDatiMedici {
				if richiesta.DatiMediciId == datiMedici.Id {
					fmt.Println("ID: ", richiesta.Id, " Data inoltro: ", richiesta.Data, " Prezzo base: ", richiesta.Prezzoproposto, "€")
					fmt.Println("Dati residenza cliente: \n", "CAP: ", datiMedici.Cliente.Persona.Cap, " Città: ", datiMedici.Cliente.Persona.Citta, " Provincia: ", datiMedici.Cliente.Persona.Provincia, "\n")
				}
			}
		}
	}
	if !sottoscrizioniPresenti {
		fmt.Println("NESSUNA")
	}
	return sottoscrizioniPresenti
}
