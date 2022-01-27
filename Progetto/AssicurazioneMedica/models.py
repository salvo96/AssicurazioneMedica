from datetime import datetime
from django.db import models

# Qui definisco i modelli per il database del sistema

class Persona(models.Model):        #Tabella che contiene le info delle persone registrate al sistema (Amministratori o Clienti o entrambi)
    codice_fiscale = models.CharField(max_length=16, primary_key=True, unique=True) #il CF è unico per singola persona ed è l'identificativo della persona(id)
    nome = models.CharField(max_length=25)
    cognome = models.CharField(max_length=25)
    data_nascita = models.DateField()
    via = models.CharField(max_length = 30)
    cap = models.CharField(max_length = 5)
    citta = models.CharField(max_length = 25)
    provincia = models.CharField(max_length = 2)

class LoginData(models.Model):  #classe astratta per le informazioni comuni a Cliente e Amministratore
    email = models.CharField(max_length=40, unique=True)
    password = models.CharField(max_length=20)
    token = models.CharField(max_length=22, default="")
    session_valid_until = models.DateTimeField(default=datetime.now)

    class Meta:
        abstract = True

class Cliente(LoginData):
    persona = models.ForeignKey(Persona, on_delete = models.CASCADE)  #se cancello la persona automaticamente cancello il cliente
    data_iscrizione = models.DateField(auto_now_add=True)   #la data viene inserita quando il dato viene immesso nel database

class Amministratore(LoginData):
    persona = models.ForeignKey(Persona, on_delete = models.CASCADE)
    note = models.CharField(max_length = 100)

class DatiMedici(models.Model):
    cliente = models.ForeignKey(Cliente, on_delete=models.CASCADE)
    sesso = models.CharField(max_length=1)
    peso = models.FloatField()
    altezza = models.FloatField()
    numero_figli = models.IntegerField()
    fumatore = models.BooleanField(default=False)

class RichiesteSottoscrizione(models.Model):
    data = models.DateField(auto_now_add=True)
    evasa = models.BooleanField(default = False)
    prezzoProposto = models.FloatField()
    dati_medici = models.ForeignKey(DatiMedici, on_delete=models.CASCADE)

class Polizza(models.Model):
    cliente = models.ForeignKey(Cliente, on_delete=models.CASCADE)
    data_sottoscrizione = models.DateField()
    scadenza = models.DateField()
    premio_assicurativo = models.FloatField()

class SpeseMedicheCliente(models.Model):
    polizza = models.ForeignKey(Polizza, on_delete=models.CASCADE)
    descrizione_spesa = models.CharField(max_length=150)
    data = models.DateField()
    importo = models.FloatField()