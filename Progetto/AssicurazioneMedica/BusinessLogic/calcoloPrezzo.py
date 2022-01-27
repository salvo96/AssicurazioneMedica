from AssicurazioneMedica.models import RichiesteSottoscrizione
import rpy2.robjects as robjects
import rpy2.robjects.numpy2ri as rpyn
from dateutil.relativedelta import relativedelta
from datetime import date
from AssicurazioneMedica.BusinessLogic.constant import messages

class CalcolatorePrezzo:

    def __init__(self, richiestaStimaPrezzo):
        self.richiestaStimaPrezzo = richiestaStimaPrezzo

    def leggiDatiMedici(self):
        #leggo i dati medici associati al cliente con richiesta identificata da "id"
        dati_medici = RichiesteSottoscrizione.objects.get(id=self.richiestaStimaPrezzo.getId()).dati_medici

        #Età: age
        data_nascita = dati_medici.cliente.persona.data_nascita
        oggi = date.today()
        age = relativedelta(oggi, data_nascita).years

        #Sesso: sex
        if dati_medici.sesso == "M":
            sex = 1
        else:
            sex = 2

        #Indice di massa corporea: bmi
        bmi = dati_medici.peso / ((dati_medici.altezza ** 2) / 10 ** 4)

        #Numero figli: children
        children = dati_medici.numero_figli

        #Regione: region
        region = self.richiestaStimaPrezzo.getRegion()

        #Fumatore: smoker
        smoker = dati_medici.fumatore

        return age, sex, bmi, children, region, smoker

    def stimaPrezzoPolizza(self):   #metodo per il calcolo del prezzo stimato della polizza mediante modello di regressione lineare in R
        r = robjects.r
        r['source']('pricePredictor.R')
        pricePredict = robjects.globalenv['pricePredict']
        age, sex, bmi, children, region, smoker = self.leggiDatiMedici()
        result = pricePredict(age, sex, bmi, children, region, smoker)
        prezzo = round(rpyn.rpy2py(result)[0])
        return {'message': messages['REQUEST_PRICE_ESTIMATION_SUCCESSFUL'],
                'prezzo': prezzo}