package BusinessLogic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (amministratore Amministratore) AmministratorePolizze() Polizza {
	client := http.Client{}
	req, err := http.NewRequest("GET", AmministratorePolizze, nil)
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
	case 200: // la richiesta della lista di polizze dei Clienti è andata a buon fine
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var result Polizza
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Can not unmarshal JSON")
		}
		return result
	case 404: //non sono presenti polizze dei Clienti
		fmt.Println("Non sono presenti polizze per i Clienti")
		return nil
	default:
		fmt.Println("Errore non conosciuto: ", resp.StatusCode)
		return nil
	}
}

func StampaListaPolizzeClienti(listaPolizze Polizza, datiMedici []DatiMedici) {
	fmt.Println("Polizze nel sistema: ")
	var codice_fiscale, nome, cognome string
	for i := 0; i < len(listaPolizze); i++ {
		for j := 0; j < len(datiMedici); j++ {
			if datiMedici[j].Cliente.Id == listaPolizze[i].ClienteID {
				codice_fiscale, nome, cognome = datiMedici[j].Cliente.Persona.Codice_fiscale, datiMedici[j].Cliente.Persona.Nome, datiMedici[j].Cliente.Persona.Cognome
				break
			}
		}
		fmt.Println("ID polizza: ", listaPolizze[i].ID, " Codice Fiscale: ", codice_fiscale, " Nome e Cognome: ", nome, " ", cognome, " Data sottoscrizione: ", listaPolizze[i].DataSottoscrizione, " Data scadenza: ", listaPolizze[i].Scadenza, " Premio: ", listaPolizze[i].PremioAssicurativo, "€")
	}
}
