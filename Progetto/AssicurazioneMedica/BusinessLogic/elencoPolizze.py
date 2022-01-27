from AssicurazioneMedica.models import Cliente, Polizza
from datetime import date

class PolizzeCliente:

    def __init__(self, email):
        self.email = email

    def ricercaPolizze(self):
        cliente = Cliente.objects.get(email = self.email.getEmail())
        polizze = Polizza.objects.filter(cliente_id=cliente.id)
        if not polizze:
            raise EmptySetException("Polizze non presenti")
        else:   #se per quel dato cliente le polizze sono presenti all'interno del database, crea delle liste saparate di polizze sulla base della data di scadenza
            listaPolizzeScadute = [polizza for polizza in polizze if polizza.scadenza < date.today()]   # con list comprehension
            listaPolizzeAttive = list()
            for polizza in filter(lambda polizza: polizza.scadenza > date.today(), polizze):    # con espressione lambda
                listaPolizzeAttive.append(polizza)
            return listaPolizzeAttive, listaPolizzeScadute


class EmptySetException(Exception):  # eccezione definita per elenco polizze vuoto
    def __init__(self, value):
        self.value = value

    def __str__(self):
        return self.value