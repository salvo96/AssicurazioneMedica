from AssicurazioneMedica.models import Cliente, Persona
from AssicurazioneMedica.BusinessLogic.constant import messages

class Registra:
    def __init__(self, registrazione):
        self.registrazione = registrazione

    def registra(self): #se siamo arrivati qui significa che il cliente non è presente sul DB e possiamo effettuare la registrazione
        persona = Persona(codice_fiscale = self.registrazione.getCodiceFiscale(),
                        nome = self.registrazione.getNome(),
                        cognome = self.registrazione.getCognome(),
                        data_nascita = self.registrazione.getDataNascita(),
                        via = self.registrazione.getVia(),
                        cap = self.registrazione.getCap(),
                        citta = self.registrazione.getCitta(),
                        provincia = self.registrazione.getProvincia())
        persona.save()
        cliente = Cliente(email = self.registrazione.getEmail(),
                          password = self.registrazione.getPassword(),
                          persona_id = self.registrazione.getCodiceFiscale())
        cliente.save()
        return {'message': messages['REGISTRATION_SUCCESSFUL_MESSAGE']}


