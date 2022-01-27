package BusinessLogic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type response struct {
	ResponseMessage
	Prezzo float32 `json:"prezzo"`
}

func (amministratore Amministratore) AmministratoreStimaPrezzoPolizza(id_richiesta, region int) float32 {
	stimaPrezzoData := map[string]int{"id": id_richiesta, "region": region} //creo mappa (struttura dati) con id richiesta di sottoscrizione e regione del Cliente
	json_data, err := json.Marshal(stimaPrezzoData)                         //li converto in formato JSON
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Calcolo prezzo in corso...")
	client := http.Client{}
	req, err := http.NewRequest("POST", AmministratoreStimaPrezzoPolizza, bytes.NewBuffer(json_data))
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
	case 200: // la richiesta di sottoscrizione esiste
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var result response
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Can not unmarshal JSON")
		}
		fmt.Println(result.Message, ": ", result.Prezzo, " €")
		return result.Prezzo
	case 404: // la richiesta di sottoscrizione non esiste
		fmt.Println("L'Utente non è presente nel sistema!\n")
		return -1
	default:
		fmt.Println("Errore non conosciuto: ", resp.StatusCode)
		return -1
	}
}

func SceltaRichiestaMessage(listaSottoscrizioni []RichiestaSottoscrizione) (int, int, float32) {
	var id_richiesta, region int
	var prezzo float32
	scegli := true
	for scegli {
		fmt.Println("\nScegli la richiesta di sottoscrizione che vuoi gestire: ")
		fmt.Println("Id richiesta: ")
		fmt.Scanf("%d\n", &id_richiesta)
		fmt.Println("\nScegli la regione di residenza del cliente: \n" +
			"1)Nord-est\n" +
			"2)Nord-ovest\n" +
			"3)Sud-est\n" +
			"4)Sud-ovest\n")
		fmt.Println("Regione di residenza: ")
		fmt.Scanf("%d\n", &region)
		for _, richiesta := range listaSottoscrizioni {
			if richiesta.Id == id_richiesta && richiesta.Evasa == false {
				prezzo = richiesta.Prezzoproposto
				scegli = false
				break
			}
		}
		if scegli {
			fmt.Println("Richiesta scelta errata\n")
		}
	}
	return id_richiesta, region, prezzo
}
