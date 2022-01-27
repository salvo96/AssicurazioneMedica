from AssicurazioneMedica.models import Amministratore, Cliente
from pytz import UTC
from datetime import datetime

def controllaAutenticazioneCliente(token):
    try:
        cliente = Cliente.objects.get(token=token)
        deadline = cliente.session_valid_until
        ora_attuale = UTC.localize(datetime.now())
        if ora_attuale > deadline:  #è scaduto il token, non devo consentire il login
            return False
        else:   #l'ora attuale è minore dell'ora di scadenza del token -> posso consentire il login
            return True
    except(Cliente.DoesNotExist):   #non ho trovato nessun Cliente con questo token -> non devo consentire il login
        return False

def controllaAutenticazioneAmministratore(token):
    try:
        amministratore = Amministratore.objects.get(token=token)
        deadline = amministratore.session_valid_until
        ora_attuale = UTC.localize(datetime.now())
        if ora_attuale > deadline:  #è scaduto il token, non devo consentire il login
            return False
        else:   #l'ora attuale è minore dell'ora di scadenza del token -> posso consentire il login
            return True
    except(Amministratore.DoesNotExist):   #non ho trovato nessun Amministratore con questo token -> non devo consentire il login
        return False
