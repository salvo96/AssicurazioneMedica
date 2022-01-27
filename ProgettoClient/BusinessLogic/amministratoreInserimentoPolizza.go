package BusinessLogic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (amministratore Amministratore) AmministratoreInserimentoPolizza(cliente_id int, data_sottoscrizione string, scadenza string, premio_assicurativo float32, richiesta_sottoscrizione_id int) {
	polizzaData := struct { //faccio uso di struct anonima per creare il JSON da inviare con la richiesta
		Cliente_id                  int     `json:"cliente_id"`
		Data_sottoscrizione         string  `json:"data_sottoscrizione"`
		Scadenza                    string  `json:"scadenza"`
		Premio_assicurativo         float32 `json:"premio_assicurativo"`
		Richiesta_sottoscrizione_id int     `json:"richiesta_sottoscrizione_id"`
	}{
		Cliente_id:                  cliente_id,
		Data_sottoscrizione:         data_sottoscrizione,
		Scadenza:                    scadenza,
		Premio_assicurativo:         premio_assicurativo,
		Richiesta_sottoscrizione_id: richiesta_sottoscrizione_id,
	}
	json_data, err := json.Marshal(polizzaData) //conversione in formato JSON
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{}
	req, err := http.NewRequest("POST", AmministratoreInserimentoPolizza, bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", amministratore.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200: // l'inserimento è avvenuto correttamente
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var result ResponseMessage
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Can not unmarshal JSON")
		}
		fmt.Println("MESSAGGIO DAL SERVER: ", result.Message)
	default:
		fmt.Println("Errore non conosciuto: ", resp.StatusCode)
	}
}

func SceltaPrezzo(prezzo_stimato, prezzo_proposto float32) float32 {
	fmt.Println("Riepilogo prezzo polizza calcolato:")
	fmt.Println("Prezzo base polizza: ", prezzo_proposto, " €")
	fmt.Println("Prezzo stimato mediante algoritmo: ", prezzo_stimato, " €")
	prezzo_totale := prezzo_stimato + prezzo_proposto
	var value = -1
	for value == -1 {
		fmt.Println("Prezzo totale: ", prezzo_totale, " €\n")
		fmt.Println("Accettare?(0=no/1=si): ")
		fmt.Scanf("%d\n", &value)
		if value == 1 {
			continue
		} else if value == 0 {
			fmt.Println("Inserire nuovo prezzo: ")
			fmt.Scanf("%f\n", &prezzo_totale)
		} else {
			value = -1
		}
	}
	return prezzo_totale
}

func InfoPolizza(datiMedici []DatiMedici, listaSottoscrizioni []RichiestaSottoscrizione, id_richiesta int) (int, string, string) {
	var cliente_id int
	for _, sottoscrizione := range listaSottoscrizioni {
		if sottoscrizione.Id == id_richiesta {
			for _, dato := range datiMedici {
				if dato.Id == sottoscrizione.DatiMediciId {
					cliente_id = dato.Cliente.Id
					break
				}
			}
			break
		}
	}
	fmt.Println("Inserire di seguito i dati per la polizza: ")
	data_sottoscrizione := ReadStringFormat("Data sottoscrizione(yyyy-mm-dd): ")
	scadenza := ReadStringFormat("Data scadenza(yyyy-mm-dd): ")

	return cliente_id, data_sottoscrizione, scadenza
}
