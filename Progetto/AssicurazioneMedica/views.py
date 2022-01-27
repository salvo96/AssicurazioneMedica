from django.http import HttpResponse, JsonResponse
from rest_framework.decorators import api_view
from AssicurazioneMedica.models import Cliente, Amministratore, Persona, SpeseMedicheCliente, Polizza, DatiMedici, RichiesteSottoscrizione
from AssicurazioneMedica.serializers import LoginSerializer, ClienteSerializer, RichiestaSottoscrizioneSerializer, EmailSerializer, PolizzaSerializer, SpeseMedicheSerializer, ADatiMediciSerializer, RichiesteSottoscrizioneSerializer, RichiestaStimaPrezzoSerializer, PolizzaClienteSerializer
from AssicurazioneMedica.BusinessLogic import login, registrazione, richiestaPolizza, elencoPolizze, elencoSpese, calcoloPrezzo, inserimentoPolizzaCliente, controllaAutenticazione
from AssicurazioneMedica.BusinessLogic.constant import messages

# Operazioni Cliente
@api_view(['POST'])
def clienteLogin(request):
    if request.method == 'POST':
        serializer = LoginSerializer(data = request.data)
        if serializer.is_valid():
            authData = serializer.save()
            #Logica del programma
            try:
                autenticaCliente = login.AutenticaCliente(authData)
                authGrant = autenticaCliente.autentica()
                return JsonResponse(authGrant)
            except Cliente.DoesNotExist:    #ovvero il Cliente non è stato trovato sul Database
                return HttpResponse(status=404)
            except login.WrongPasswordException: #la password fornita è errata
                return HttpResponse(status=400)

@api_view(['POST'])
def clienteRegistrazione(request):
    if request.method == 'POST':
        serializer = ClienteSerializer(data=request.data)
        if serializer.is_valid():
            regData = serializer.save()
            #Logica del programma
            try:    #verifica che l'email non sia già presente
                Cliente.objects.get(email=regData.getEmail())
            except Cliente.DoesNotExist:
                try:    #verifica che il cf non sia già presente
                    Persona.objects.get(codice_fiscale=regData.getCodiceFiscale())
                except Persona.DoesNotExist:
                    registraCliente = registrazione.Registra(regData)
                    regGrant = registraCliente.registra()
                    return JsonResponse(regGrant)
            return HttpResponse(status=400)

@api_view(['POST'])
def richiestaSottoscrizione(request):
    if controllaAutenticazione.controllaAutenticazioneCliente(request.headers.get('Authorization')):
        if request.method == 'POST':
            serializer = RichiestaSottoscrizioneSerializer(data=request.data)
            if serializer.is_valid():
                reqData = serializer.save()
                #Logica del programma
                reqPolizza = richiestaPolizza.RichiestaPolizza(reqData)
                reqGrant = reqPolizza.richiedi()
                return JsonResponse(reqGrant)
    else:
        return HttpResponse(status=401)  #il cliente non si è loggato e non ha autorizzazione ad accedere a tale risorsa oppure la sessione è scaduta oppure il token non è valido

@api_view(['POST'])
def elencoPolizzeCliente(request):
    if controllaAutenticazione.controllaAutenticazioneCliente(request.headers.get('Authorization')):
        if request.method == 'POST':
            serializer = EmailSerializer(data=request.data)
            if serializer.is_valid():
                emailData = serializer.save()
                #Logica del programma
                try:
                    polizzeCliente = elencoPolizze.PolizzeCliente(emailData)
                    listaPolizzeAttive, listaPolizzeScadute = polizzeCliente.ricercaPolizze()
                    serializerAttive = PolizzaSerializer(listaPolizzeAttive, many=True)
                    serializerScadute = PolizzaSerializer(listaPolizzeScadute, many=True)
                    return JsonResponse({'polizze_attive': serializerAttive.data, 'polizze_scadute': serializerScadute.data})
                except (elencoPolizze.EmptySetException):   #non sono presenti polizze per l'utente
                    return HttpResponse(status=404)
    else:
        return HttpResponse(status=401)  # il cliente non si è loggato e non ha autorizzazione ad accedere a tale risorsa oppure la sessione è scaduta oppure il token non è valido

@api_view(['POST'])
def inserimentoSpeseMediche(request):
    if controllaAutenticazione.controllaAutenticazioneCliente(request.headers.get('Authorization')):
        if request.method == 'POST':
            serializer = SpeseMedicheSerializer(data=request.data, many=True)   # devo poter gestire l'inserimento di più spese in una volta
            if serializer.is_valid():
                speseMediche = serializer.save()
                for spesa in speseMediche:
                    SpeseMedicheCliente(polizza=Polizza.objects.get(id = spesa.getPolizzaId()),
                                        descrizione_spesa=spesa.getDescrizioneSpesa(),
                                        data=spesa.getData(),
                                        importo=spesa.getImporto()).save()
                return JsonResponse({'message': messages['CHARGE_INSERT_SUCCESSFUL_MESSAGE']})
    else:
        return HttpResponse(status=401)  # il cliente non si è loggato e non ha autorizzazione ad accedere a tale risorsa oppure la sessione è scaduta oppure il token non è valido

@api_view(['POST'])
def letturaSpeseMediche(request):
    if controllaAutenticazione.controllaAutenticazioneCliente(request.headers.get('Authorization')):
        if request.method == 'POST':
            serializer = EmailSerializer(data=request.data)
            if serializer.is_valid():
                emailData = serializer.save()
                #logica del programma
                try:
                    speseCliente = elencoSpese.SpeseCliente(emailData)
                    listaSpese = speseCliente.ricercaSpese()
                    speseGrant = SpeseMedicheSerializer(listaSpese, many=True)
                    return JsonResponse(speseGrant.data, safe=False)
                except (elencoPolizze.EmptySetException, Cliente.DoesNotExist):
                    return HttpResponse(status=404)
    else:
        return HttpResponse(status=401)  # il cliente non si è loggato e non ha autorizzazione ad accedere a tale risorsa oppure la sessione è scaduta oppure il token non è valido

#Operazioni Amministratore
@api_view(['POST'])
def amministratoreLogin(request):
    if request.method == 'POST':
        serializer = LoginSerializer(data = request.data)
        if serializer.is_valid():
            authData = serializer.save()
            #Logica del programma
            try:
                autenticaAmministratore = login.AutenticaAmministratore(authData)
                authGrant = autenticaAmministratore.autentica()
                return JsonResponse(authGrant)
            except Amministratore.DoesNotExist:    #ovvero l'Amministratore non è stato trovato sul Database
                return HttpResponse(status=404)
            except login.WrongPasswordException: #la password fornita è errata
                return HttpResponse(status=400)

@api_view(['GET'])
def datiMediciClienti(request):
    if controllaAutenticazione.controllaAutenticazioneAmministratore(request.headers.get('Authorization')):
        if request.method == 'GET':
            try:
                datiMedici = DatiMedici.objects.all()
                serializer = ADatiMediciSerializer(datiMedici, many = True)
                return JsonResponse(serializer.data, safe=False)
            except DatiMedici.DoesNotExist:
                return HttpResponse(status=404)
    else:
        return HttpResponse(status=401) #l'amministratore non si è loggato e non ha autorizzazione ad accedere a tale risorsa oppure la sessione è scaduta oppure il token non è valido

@api_view(['GET'])
def polizzeClienti(request):
    if controllaAutenticazione.controllaAutenticazioneAmministratore(request.headers.get('Authorization')):
        if request.method == 'GET':
            try:
                polizze = Polizza.objects.all()
                serializer = PolizzaSerializer(polizze, many = True)
                return JsonResponse(serializer.data, safe = False)
            except Polizza.DoesNotExist:
                return HttpResponse(status=404)
    else:
        return HttpResponse(status=401) #l'amministratore non si è loggato e non ha autorizzazione ad accedere a tale risorsa oppure la sessione è scaduta oppure il token non è valido

@api_view(['GET'])
def speseMedicheClienti(request):
    if controllaAutenticazione.controllaAutenticazioneAmministratore(request.headers.get('Authorization')):
        if request.method == 'GET':
            try:
                speseMediche = SpeseMedicheCliente.objects.all()
                serializer = SpeseMedicheSerializer(speseMediche, many = True)
                return JsonResponse(serializer.data, safe=False)
            except SpeseMedicheCliente.DoesNotExist:
                return HttpResponse(status=404)
    else:
        return HttpResponse(status=401) #l'amministratore non si è loggato e non ha autorizzazione ad accedere a tale risorsa oppure la sessione è scaduta oppure il token non è valido

@api_view(['GET'])
def richiesteSottoscrizioneClienti(request):
    if controllaAutenticazione.controllaAutenticazioneAmministratore(request.headers.get('Authorization')):
        if request.method == 'GET':
            try:
                richieste = RichiesteSottoscrizione.objects.all()
                serializer = RichiesteSottoscrizioneSerializer(richieste, many = True)
                return JsonResponse(serializer.data, safe=False)
            except RichiesteSottoscrizione.DoesNotExist:
                return HttpResponse(status=404)
    else:
        return HttpResponse(status=401) #l'amministratore non si è loggato e non ha autorizzazione ad accedere a tale risorsa oppure la sessione è scaduta oppure il token non è valido

@api_view(['POST'])
def calcoloStimaPrezzo(request):
    if controllaAutenticazione.controllaAutenticazioneAmministratore(request.headers.get('Authorization')):
        if request.method == 'POST':
            serializer = RichiestaStimaPrezzoSerializer(data = request.data)
            if serializer.is_valid():
                richiestaStimaPrezzo = serializer.save()
                #Logica del programma
                try:
                    calcolatorePrezzo = calcoloPrezzo.CalcolatorePrezzo(richiestaStimaPrezzo)
                    reqGrant = calcolatorePrezzo.stimaPrezzoPolizza()
                    return JsonResponse(reqGrant)
                except (RichiesteSottoscrizione.DoesNotExist):
                    return HttpResponse(status=404)
    else:
        return HttpResponse(status=401) #l'amministratore non si è loggato e non ha autorizzazione ad accedere a tale risorsa oppure la sessione è scaduta oppure il token non è valido

@api_view(['POST'])
def inserimentoPolizza(request):    #funzione per l'inserimento della Polizza approvata da parte dell'Amministratore nel sistema
    if controllaAutenticazione.controllaAutenticazioneAmministratore(request.headers.get('Authorization')):
        if request.method == 'POST':
            serializer = PolizzaClienteSerializer(data = request.data)
            if serializer.is_valid():
                polizzaData = serializer.save()
                # Logica del programma
                inserimentoPolizza = inserimentoPolizzaCliente.InserimentoPolizzaCliente(polizzaData)
                reqGrant = inserimentoPolizza.inserisciPolizza()
                return JsonResponse(reqGrant)
    else:
        return HttpResponse(status=401)  # l'amministratore non si è loggato e non ha autorizzazione ad accedere a tale risorsa oppure la sessione è scaduta oppure il token non è valido
