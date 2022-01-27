package BusinessLogic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Polizza []struct {
	ID                 int     `json:"id"`
	ClienteID          int     `json:"cliente_id"`
	DataSottoscrizione string  `json:"data_sottoscrizione"`
	Scadenza           string  `json:"scadenza"`
	PremioAssicurativo float32 `json:"premio_assicurativo"`
}

type ResponsePolizze struct {
	Polizze_attive  Polizza `json:"polizze_attive"`
	Polizze_scadute Polizza `json:"polizze_scadute"`
}

func (cliente Cliente) PolizzeCliente() (Polizza, Polizza) {
	emailData := map[string]string{"email": cliente.Email} //creo mappa (struttura dati) con dati per la richiesta lista polizze Cliente
	json_data, err := json.Marshal(emailData)              //conversione in formato JSON
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{}
	req, err := http.NewRequest("POST", PolizzeCliente, bytes.NewBuffer(json_data))
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
		var result ResponsePolizze
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Can not unmarshal JSON")
		}
		return result.Polizze_attive, result.Polizze_scadute
	case 404:
		return nil, nil
	default:
		fmt.Println("Errore non conosciuto: ", resp.StatusCode)
		return nil, nil
	}
}

func StampaListaPolizze(polizze_attive, polizze_scadute Polizza) {
	PolizzePrint(polizze_attive, "Lista polizze attive: ")
	PolizzePrint(polizze_scadute, "Lista polizze scadute: ")
}

func PolizzePrint(polizze Polizza, testo string) bool {
	fmt.Println(testo)
	if length := len(polizze); length != 0 {
		for i := 0; i < length; i++ {
			fmt.Println("ID: ", polizze[i].ID, " Data sottoscrizione: ", polizze[i].DataSottoscrizione, " Data scadenza: ", polizze[i].Scadenza, " Premio Assicurativo: ", polizze[i].PremioAssicurativo, "\n")
		}
		return true
	} else {
		fmt.Println("Nessuna polizza presente\n")
		return false
	}
}
