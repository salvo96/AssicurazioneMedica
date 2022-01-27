package BusinessLogic

import (
	"time"
)

func ManageSession(utente *Utente) {
	layout := "2006-01-02 15:04:05.000000" //layout stringa di tempo
	for {
		deadline, _ := time.ParseInLocation(layout, utente.Session_valid_until, time.Local) //istante di scadenza della sessione
		currentTime := time.Now()                                                           //tempo attuale
		differenza := deadline.Sub(currentTime)
		time.Sleep(differenza)                                             //tempo di sospensione della goroutine
		utente.LoginUtente(utente.Email, utente.Password, utente.Endpoint) //rieffettuo un nuovo login: aggiorno la sessione, aggiornando la validità e il token
	}
}
