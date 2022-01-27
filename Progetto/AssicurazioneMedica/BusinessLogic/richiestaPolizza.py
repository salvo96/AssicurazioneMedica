from AssicurazioneMedica.BusinessLogic.elencoSpese import SpeseCliente
from AssicurazioneMedica.models import DatiMedici, Cliente, RichiesteSottoscrizione
from AssicurazioneMedica.BusinessLogic.constant import messages
from AssicurazioneMedica.serializers import Email
from AssicurazioneMedica.BusinessLogic.elencoPolizze import EmptySetException
from statistics import mean

class RichiestaPolizza(SpeseCliente):
    def __init__(self, richiesta):
        self.richiesta = richiesta
        super().__init__(Email(self.richiesta.getEmail()))

    def richiedi(self):
        # genero il prezzo base per l'utente tenendo conto delle spese sostenute, se sostenute
        try:
            spese = self.ricercaSpese()
            prezzo = round(mean(spese.values_list('importo', flat=True)))
        except(EmptySetException):
            prezzo = 0
        id_cliente = Cliente.objects.get(email=self.richiesta.getEmail())
        #inserire dati richiesta nel DB
        try:
            datiMedici = DatiMedici.objects.get(cliente=id_cliente)
            datiMedici.sesso = self.richiesta.getSesso()
            datiMedici.peso = self.richiesta.getPeso()
            datiMedici.altezza = self.richiesta.getAltezza()
            datiMedici.numero_figli = self.richiesta.getNumeroFigli()
            datiMedici.fumatore = self.richiesta.getFumatore()
        except (DatiMedici.DoesNotExist):
            datiMedici = DatiMedici(cliente = id_cliente,
                                                          sesso = self.richiesta.getSesso(),
                                                          peso = self.richiesta.getPeso(),
                                                          altezza = self.richiesta.getAltezza(),
                                                          numero_figli = self.richiesta.getNumeroFigli(),
                                                          fumatore = self.richiesta.getFumatore())
        datiMedici.save()
        richiestaSottoscrizione = RichiesteSottoscrizione(dati_medici_id = datiMedici.id, prezzoProposto = prezzo)
        richiestaSottoscrizione.save()
        return {'message': messages['REQUEST_SUCCESSFUL_MESSAGE'],
                'prezzo': str(prezzo)+" €"}