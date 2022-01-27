package BusinessLogic

import (
	"fmt"
)

func Welcome(user Utente) int { //si passa una copia di UserInfo per capire se l'utente è già loggato o meno
	var choice = 0

	fmt.Println("******************************************************")
	fmt.Println("\tSISTEMA DI ASSICURAZIONE MEDICA\t")
	fmt.Println("******************************************************")

	if user.Token == "" {
		fmt.Println("\nSelezionare il tipo di operazione che si desidera eseguire:")
		fmt.Println("1. Login Amministratore")
		fmt.Println("2. Login Cliente")
		fmt.Println("3. Registrazione Cliente")
	} else { // l'utente si è loggato precedentemente
		fmt.Println("Utente: ", user.Email)
		fmt.Println("\nSelezionare il tipo di operazione che si desidera eseguire:")
		if user.Tipo == false { //ovvero l'utente è un amministratore
			choice = 3
			fmt.Println("1. Visualizza dati medici clienti")
			fmt.Println("2. Visualizza lista polizze")
			fmt.Println("3. Visualizza lista spese mediche sostenute")
			fmt.Println("4. Gestisci richieste sottoscrizione polizza")
		} else { //l'utente è un cliente
			choice = 7
			fmt.Println("1. Crea nuova richiesta di sottoscrizione polizza assicurativa")
			fmt.Println("2. Visualizza stato polizze assicurative")
			fmt.Println("3. Inserisci spese mediche sostenute")
			fmt.Println("4. Visualizza spese mediche sostenute")
		}
	}
	fmt.Println("\nInserisci l'operazione:")
	var value int
	fmt.Scanf("%d\n", &value)
	choice += value
	return choice
}
