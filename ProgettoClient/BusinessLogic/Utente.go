package BusinessLogic

type Utente struct { //oggetto Utente che generalizza il Cliente e l'Amministratore
	Email, Password, Persona_id, Session_valid_until, Token, Endpoint string
	Tipo                                                              bool //indica il tipo di utente: 0 => Amministratore, 1 => Cliente
	UserMethods
}

type Cliente struct { //oggetto cliente che eredita dall'oggetto Utente
	Utente
	Data_iscrizione string
}

type Amministratore struct {
	Utente
	note string
}

// ------------------------------------------------------------------------------------------------------------
// sezione contenente le interfacce dei vari metodi che devono avere il Cliente e l'Amministratore

type UserMethods interface { //interfaccia con metodi che devono essere implementati dall'oggetto Utente
	LoginUtente(email string, password string, endpoint string) bool
	RegistrazioneUtente(codice_fiscale, nome, cognome, data_nascita, via, cap, citta, provincia, email, password string)
	SottoscrizioneCliente(sesso, peso, altezza, numero_figli, fumatore string)
	PolizzeCliente() (Polizza, Polizza)
	InserimentoSpeseMedicheCliente(listaSpese []Spesa)
	SpeseMedicheCliente() []SpesaNum
	AmministratoreDatiMedici() []DatiMedici
	AmministratorePolizze() Polizza
	AmministratoreSpeseMediche() []SpesaNum
	AmministratoreSottoscrizioniClienti() []RichiestaSottoscrizione
	AmministratoreStimaPrezzoPolizza(id_richiesta, region int) float32
	AmministratoreInserimentoPolizza(cliente_id int, data_sottoscrizione string, scadenza string, premio_assicurativo float32, richiesta_sottoscrizione_id int)
}
