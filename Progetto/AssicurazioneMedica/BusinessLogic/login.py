from AssicurazioneMedica.models import Cliente, Amministratore
import secrets
import datetime
from AssicurazioneMedica.BusinessLogic import constant

class Autentica:
    def __init__(self):
        self.utente = None
    def autentica(self):
        if self.utente.password == self.login.getPassword():
            #fai operazioni di generazione 'secrets', salvataggio nel db
            self.utente.token = secrets.token_urlsafe(constant.SESSION_TOKEN_LENGTH)    #viene generato il token per quell'utente
            deadline = datetime.datetime.now() + datetime.timedelta(minutes=constant.SESSION_LAST)  #viene settata la sessione
            self.utente.session_valid_until = deadline.strftime("%Y-%m-%d %H:%M:%S.%f")
            self.utente.save()
            return {"token": self.utente.token, "session_lifetime": self.utente.session_valid_until} #restituisco un dizionario con coppia token:deadline_token
        else:
            raise WrongPasswordException("Password non corretta!")

class AutenticaCliente(Autentica):
    def __init__(self, login):
        self.login = login
        self.utente = Cliente.objects.get(email = self.login.getEmail())

class AutenticaAmministratore(Autentica):
    def __init__(self, login):
        self.login = login
        self.utente = Amministratore.objects.get(email=self.login.getEmail())

class WrongPasswordException(Exception): #eccezione definita per password errata
    def __init__(self, value):
        self.value = value
    def __str__(self):
        return self.value