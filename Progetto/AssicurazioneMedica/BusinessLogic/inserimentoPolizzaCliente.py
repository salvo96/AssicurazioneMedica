from AssicurazioneMedica.models import Polizza, RichiesteSottoscrizione, Cliente
from AssicurazioneMedica.BusinessLogic.constant import messages

class InserimentoPolizzaCliente:
    def __init__(self, polizzaData):
        self.polizzaData = polizzaData

    def inserisciPolizza(self):
        cliente = Cliente.objects.get(id=self.polizzaData.getClienteId())
        polizza = Polizza(cliente = cliente, data_sottoscrizione = self.polizzaData.getDataSottoscrizione(), scadenza = self.polizzaData.getScadenza(), premio_assicurativo = self.polizzaData.getPremioAssicurativo())
        polizza.save()
        richiestaSottoscrizioneId = self.polizzaData.getRichiestaSottoscrizioneId()
        richiestaSottoscrizione = RichiesteSottoscrizione.objects.get(id=richiestaSottoscrizioneId)
        richiestaSottoscrizione.evasa = 1
        richiestaSottoscrizione.save()

        return {"message": messages['BILL_SUCCESSFUL_MESSAGE']}
