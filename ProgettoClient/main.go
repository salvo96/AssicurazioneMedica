package main

import (
	"ProgettoClient/BusinessLogic"
	"fmt"
)

func main() {
	esegui := true
	utente := &BusinessLogic.Utente{}
	amministratore := new(BusinessLogic.Amministratore)
	cliente := new(BusinessLogic.Cliente)

	for esegui {
		//Funzione per la stampa a video del menu
		choice := BusinessLogic.Welcome(*utente)
		switch {
		case choice == 1, choice == 2: //login Utente: Amministratore o Cliente
			email, password := BusinessLogic.LoginMessage()
			var endpoint string
			if choice == 1 { //login Amministratore
				utente.Tipo = false
				endpoint = BusinessLogic.LoginAmministratore
			} else if choice == 2 { //login Cliente
				utente.Tipo = true
				endpoint = BusinessLogic.LoginCliente
			}
			if utente.LoginUtente(email, password, endpoint) {
				fmt.Println("Login Utente avvenuto correttamente!")
				go BusinessLogic.ManageSession(utente) //avvio goroutine che gestisce riconnessione all'atto della fine della sessione
				if choice == 1 {
					amministratore.Utente = *utente
				} else if choice == 2 {
					cliente.Utente = *utente
				}
			}
		case choice == 3: //registrazione Cliente
			codice_fiscale, nome, cognome, data_nascita, via, cap_, citta, provincia, email, password := BusinessLogic.RegistrazioneMessage()
			cliente := new(BusinessLogic.Cliente)
			cliente.RegistrazioneUtente(codice_fiscale, nome, cognome, data_nascita, via, cap_, citta, provincia, email, password)
		case choice == 4: //visualizzazione dei dati medici dei Clienti
			datiMedici := amministratore.AmministratoreDatiMedici()
			BusinessLogic.StampaListaDatiMedici(datiMedici)
		case choice == 5: //visualizzazione della lista delle polizze dei Clienti
			listaPolizze := amministratore.AmministratorePolizze()
			BusinessLogic.StampaListaPolizzeClienti(listaPolizze, amministratore.AmministratoreDatiMedici())
		case choice == 6: //visualizzazione della lista delle spese mediche sostenute
			listaSpese := amministratore.AmministratoreSpeseMediche()
			BusinessLogic.StampaListaSpeseMedicheClienti(listaSpese, amministratore.AmministratorePolizze(), amministratore.AmministratoreDatiMedici())
		case choice == 7: // gestione richieste sottoscrizione Clienti

			//STEP 1: si mostra a video la lista delle richieste di sottoscrizione
			listaSottoscrizioni := amministratore.AmministratoreSottoscrizioniClienti()
			datiMedici := amministratore.AmministratoreDatiMedici()
			sottoscrizioni := BusinessLogic.StampaListaSottoscrizioniClienti(listaSottoscrizioni, datiMedici)
			if !sottoscrizioni {
				break
			}
			//STEP 2: Si effettua il calcolo della stima del prezzo della polizza
			id_richiesta, region, prezzo_base := BusinessLogic.SceltaRichiestaMessage(listaSottoscrizioni)
			prezzo := amministratore.AmministratoreStimaPrezzoPolizza(id_richiesta, region)

			//STEP 3:
			prezzo_totale := BusinessLogic.SceltaPrezzo(prezzo, prezzo_base)
			cliente_id, data_sottoscrizione, scadenza := BusinessLogic.InfoPolizza(datiMedici, listaSottoscrizioni, id_richiesta)
			amministratore.AmministratoreInserimentoPolizza(cliente_id, data_sottoscrizione, scadenza, prezzo_totale, id_richiesta)
		case choice == 8: //creazione nuova richiesta di sottoscrizione polizza assicurativa
			sesso, peso, altezza, numero_figli, fumatore := BusinessLogic.SottoscrizioneMessage()
			cliente.SottoscrizioneCliente(sesso, peso, altezza, numero_figli, fumatore)
		case choice == 9: //visualizzazione stato polizze assicurative del cliente
			polizze_attive, polizze_scadute := cliente.PolizzeCliente()
			BusinessLogic.StampaListaPolizze(polizze_attive, polizze_scadute)
		case choice == 10: //inserimento spese mediche sostenute
			polizze_attive, _ := cliente.PolizzeCliente()
			spese := BusinessLogic.InserimentoSpeseMessage(polizze_attive)
			if spese != nil {
				cliente.InserimentoSpeseMedicheCliente(spese)
			}
		case choice == 11: //visualizzazione spese mediche sostenute
			spese := cliente.SpeseMedicheCliente()
			BusinessLogic.StampaListaSpese(spese)
		}
	}
}
