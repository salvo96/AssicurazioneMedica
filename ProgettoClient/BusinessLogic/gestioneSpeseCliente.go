package BusinessLogic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Spesa struct {
	Descrizione_spesa string `json:"descrizione_spesa"`
	Data              string `json:"data"`
	Importo           string `json:"importo"`
	Polizza_id        string `json:"polizza_id"`
}

type SpesaNum struct { //struct con dati Spesa di tipo numerico
	Descrizione_spesa string  `json:"descrizione_spesa"`
	Data              string  `json:"data"`
	Importo           float32 `json:"importo"`
	Polizza_id        int     `json:"polizza_id"`
}

func (cliente Cliente) InserimentoSpeseMedicheCliente(listaSpese []Spesa) {
	json_data, err := json.Marshal(listaSpese) //viene generato direttamente il JSON della lista di struct annotata                                                                                                                                                                                               //conversione in formato JSON
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{}
	req, err := http.NewRequest("POST", InserimentoSpeseCliente, bytes.NewBuffer(json_data))
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
	case 200: // l'inserimento delle spese è avvenuto correttamente
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

func InserimentoSpeseMessage(polizze Polizza) []Spesa { //passo come parametro la lista delle polizze attive del Cliente
	var polizza_id string
	inserimento := true
	var listaSpese []Spesa //questa è la lista delle spese da inserire nel DB
	if len(polizze) > 0 {
		for inserimento {
			fmt.Println("Inserisci di seguito i dati delle spese mediche sostenute: ")
			descrizione_spesa := ReadStringFormat("Descrizione spesa: ")
			data := ReadStringFormat("Data (formato: YYYY-MM-DD): ")
			importo := ReadStringFormat("Importo: ")
			PolizzePrint(polizze, "Selezionare l'ID di una polizza a cui associare la spesa: ")
			fmt.Println("ID: ")
			fmt.Scanf("%s\n", &polizza_id)
			spesa := &Spesa{descrizione_spesa, data, importo, polizza_id}
			listaSpese = append(listaSpese, *spesa)
			res := ReadStringFormat("Inserire altra spesa? (0 = no, 1 = si): ")
			if res == "0" {
				inserimento = false
			}
		}
		return listaSpese
	} else {
		fmt.Println("Nessuna polizza disponibile. Sottoscrivere prima una polizza.")
		return nil
	}
}

func (cliente Cliente) SpeseMedicheCliente() []SpesaNum {
	emailData := map[string]string{"email": cliente.Email} //creo mappa (struttura dati) con dati per la richiesta lista polizze Cliente
	json_data, err := json.Marshal(emailData)              //conversione in formato JSON
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{}
	req, err := http.NewRequest("POST", SpeseCliente, bytes.NewBuffer(json_data))
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
	case 200: // la richiesta della lista di spese è andata a buon fine
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var result []SpesaNum
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Can not unmarshal JSON")
		}
		return result
	case 404:
		return nil
	default:
		fmt.Println("Errore non conosciuto: ", resp.StatusCode)
		return nil
	}

}

func StampaListaSpese(listaSpese []SpesaNum) {
	fmt.Println("Lista spese sostenute dal cliente: ")
	if length := len(listaSpese); length != 0 {
		for i := 0; i < length; i++ {
			fmt.Println("Descrizione spesa: ", listaSpese[i].Descrizione_spesa, " Data: ", listaSpese[i].Data, " Importo: ", listaSpese[i].Importo, " ID Polizza di appartenenza: ", listaSpese[i].Polizza_id)
		}
	} else {
		fmt.Println("Nessuna spesa presente\n")
	}
}
