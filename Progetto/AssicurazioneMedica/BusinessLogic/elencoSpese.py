from AssicurazioneMedica.BusinessLogic.elencoPolizze import PolizzeCliente, EmptySetException
from AssicurazioneMedica.models import SpeseMedicheCliente

class SpeseCliente(PolizzeCliente):
    def ricercaSpese(self):
        self.elencoPolizze = self.ricercaPolizze()   #mi faccio restituire la tupla contente le due liste di polizze per quel dato utente
        listaPolizzeId = [polizza.id for lista in self.elencoPolizze for polizza in lista] # creo una lista contente gli
        # id delle polizze dell'utente da poter utilizzare per effettuare il filtraggio delle spese mediche dell'utente
        speseMediche = SpeseMedicheCliente.objects.filter(polizza_id__in=listaPolizzeId)
        if not speseMediche:
            raise EmptySetException("Polizze non presenti")
        return speseMediche
