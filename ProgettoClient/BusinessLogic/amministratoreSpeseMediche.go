package BusinessLogic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (amministratore Amministratore) AmministratoreSpeseMediche() []SpesaNum {
	client := http.Client{}
	req, err := http.NewRequest("GET", AmministratoreSpeseMediche, nil)
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
	case 200: // la richiesta della lista di spese mediche dei Clienti è andata a buon fine
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var result []SpesaNum
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Can not unmarshal JSON")
		}
		return result
	case 404: //non sono presenti spese mediche sostenute dei Clienti
		fmt.Println("Non sono presenti spese mediche sostenute dai Clienti")
		return nil
	default:
		fmt.Println("Errore non conosciuto: ", resp.StatusCode)
		return nil
	}
}

func StampaListaSpeseMedicheClienti(listaSpese []SpesaNum, polizze Polizza, datiMedici []DatiMedici) {
	fmt.Println("Spese mediche presenti nel sistema: ")
	var codice_fiscale, nome, cognome, dataSottoscrizione, scadenza string
	var premio float32
	for _, spesa := range listaSpese {
		for _, polizza := range polizze {
			if spesa.Polizza_id == polizza.ID {
				for _, dati := range datiMedici {
					if polizza.ClienteID == dati.Cliente.Id {
						codice_fiscale, nome, cognome = dati.Cliente.Persona.Codice_fiscale, dati.Cliente.Persona.Nome, dati.Cliente.Persona.Cognome
						break
					}
				}
				dataSottoscrizione, scadenza, premio = polizza.DataSottoscrizione, polizza.Scadenza, polizza.PremioAssicurativo
				break
			}
		}
		fmt.Println("Descrizione spesa: ", spesa.Descrizione_spesa, " Sostenuta il: ", spesa.Data, " Costo: ", spesa.Importo, "€")
		fmt.Println("Relativa alla polizza: ", spesa.Polizza_id, "sottoscritta il: ", dataSottoscrizione, " con scadenza: ", scadenza, " e premio: ", premio, "€")
		fmt.Println("Sostenuta da: ", nome, " ", cognome, " Codice Fiscale: ", codice_fiscale, "\n")
	}
}
