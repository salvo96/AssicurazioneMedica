from rest_framework import serializers
from AssicurazioneMedica.models import Cliente, Polizza, Persona, DatiMedici, RichiesteSottoscrizione

## Serializzazione per operazioni di Login Cliente + Login Amministratore
class Login:
    def __init__(self, email, password):
        self.email = email
        self.password = password
    def getEmail(self):
        return self.email
    def getPassword(self):
        return self.password

class LoginSerializer(serializers.Serializer):
    email = serializers.EmailField()
    password = serializers.CharField(max_length = 20)

    def create(self, validated_data):
        return Login(**validated_data)

    def update(self, instance, validated_data):
        instance.email = validated_data.get('email', instance.email)
        instance.password = validated_data.get('password', instance.password)
        return instance

##Serializzazione per operazioni di registrazione Cliente
class ClienteRegistrazione:
    def __init__(self, codice_fiscale, nome, cognome, data_nascita, via, cap, citta, provincia, email, password):
        self.codice_fiscale = codice_fiscale
        self.nome = nome
        self.cognome = cognome
        self.data_nascita = data_nascita
        self.via = via
        self.cap = cap
        self.citta = citta
        self.provincia = provincia
        self.email = email
        self.password = password

    def getEmail(self):
        return self.email

    def getCodiceFiscale(self):
        return self.codice_fiscale

    def getNome(self):
        return self.nome

    def getCognome(self):
        return self.cognome

    def getDataNascita(self):
        return self.data_nascita

    def getVia(self):
        return self.via

    def getCap(self):
        return self.cap

    def getCitta(self):
        return self.citta

    def getProvincia(self):
        return self.provincia

    def getPassword(self):
        return self.password

class ClienteSerializer(serializers.Serializer):
    codice_fiscale = serializers.CharField(max_length = 16)
    nome = serializers.CharField(max_length = 25)
    cognome = serializers.CharField(max_length = 25)
    data_nascita = serializers.DateField()
    via = serializers.CharField(max_length = 30)
    cap = serializers.CharField(max_length = 5)
    citta = serializers.CharField(max_length = 25)
    provincia = serializers.CharField(max_length = 2)
    email = serializers.EmailField()
    password = serializers.CharField(max_length = 20)

    def create(self, validated_data):
        return ClienteRegistrazione(**validated_data)

#serializzazione per operazioni di richiesta sottoscrizione da parte dell'utente
class RichiestaSottoscrizione:
    def __init__(self, email, sesso, peso, altezza, numero_figli, fumatore):
        self.email = email
        self.sesso = sesso
        self.peso = peso
        self.altezza = altezza
        self.numero_figli = numero_figli
        self.fumatore = fumatore

    def getEmail(self):
        return self.email

    def getSesso(self):
        return self.sesso

    def getPeso(self):
        return self.peso

    def getAltezza(self):
        return self.altezza

    def getNumeroFigli(self):
        return self.numero_figli

    def getFumatore(self):
        return self.fumatore


class RichiestaSottoscrizioneSerializer(serializers.Serializer):
    email = serializers.EmailField()
    sesso = serializers.CharField(max_length=1)
    peso = serializers.FloatField()
    altezza = serializers.FloatField()
    numero_figli = serializers.IntegerField()
    fumatore = serializers.BooleanField()

    def create(self, validated_data):
        return RichiestaSottoscrizione(**validated_data)

class Email:
    def __init__(self, email):
        self.email = email

    def getEmail(self):
        return self.email

class EmailSerializer(serializers.Serializer):
    email = serializers.EmailField()

    def create(self, validated_data):
        return Email(**validated_data)

class PolizzaSerializer(serializers.ModelSerializer):
    cliente_id = serializers.IntegerField(read_only=False)
    class Meta:
        model = Polizza
        fields = ['id', 'cliente_id', 'data_sottoscrizione', 'scadenza', 'premio_assicurativo']

class SpeseMediche:
    def __init__(self, descrizione_spesa, data, importo, polizza_id):
        self.descrizione_spesa = descrizione_spesa
        self.data = data
        self.importo = importo
        self.polizza_id = polizza_id

    def getDescrizioneSpesa(self):
        return self.descrizione_spesa

    def getData(self):
        return self.data

    def getImporto(self):
        return self.importo

    def getPolizzaId(self):
        return self.polizza_id

class SpeseMedicheSerializer(serializers.Serializer):
    descrizione_spesa = serializers.CharField(max_length=150)
    data = serializers.DateField()
    importo = serializers.FloatField()
    polizza_id = serializers.IntegerField()

    def create(self, validated_data):
        return SpeseMediche(**validated_data)

#Serializers per operazioni Amministratore
class APersonaSerializer(serializers.ModelSerializer):
    class Meta:
        model = Persona
        fields = ['codice_fiscale', 'nome', 'cognome', 'data_nascita', 'via', 'cap', 'citta', 'provincia']


class AClienteSerializer(serializers.ModelSerializer):
    persona = APersonaSerializer()
    class Meta:
        model = Cliente
        fields = ['id', 'persona', 'data_iscrizione']

class ADatiMediciSerializer(serializers.ModelSerializer):
    cliente = AClienteSerializer()
    class Meta:
        model = DatiMedici
        fields = ['cliente', 'id', 'sesso', 'peso', 'altezza', 'numero_figli', 'fumatore']

class RichiesteSottoscrizioneSerializer(serializers.ModelSerializer):
    class Meta:
        model = RichiesteSottoscrizione
        fields = ['id', 'data', 'evasa', 'prezzoProposto', 'dati_medici']

class RichiestaStimaPrezzo:
    def __init__(self, id, region):
        self.id = id
        self.region = region

    def getId(self):
        return self.id

    def getRegion(self):
        return self.region

class RichiestaStimaPrezzoSerializer(serializers.Serializer):
    id = serializers.IntegerField()
    region = serializers.IntegerField()

    def create(self, validated_data):
        return RichiestaStimaPrezzo(**validated_data)

class PolizzaCliente:
    def __init__(self, cliente_id, data_sottoscrizione, scadenza, premio_assicurativo, richiesta_sottoscrizione_id):
        self.cliente_id = cliente_id
        self.data_sottoscrizione = data_sottoscrizione
        self.scadenza = scadenza
        self.premio_assicurativo = premio_assicurativo
        self.richiesta_sottoscrizione_id = richiesta_sottoscrizione_id

    def getClienteId(self):
        return self.cliente_id

    def getDataSottoscrizione(self):
        return self.data_sottoscrizione

    def getScadenza(self):
        return self.scadenza

    def getPremioAssicurativo(self):
        return self.premio_assicurativo

    def getRichiestaSottoscrizioneId(self):
        return self.richiesta_sottoscrizione_id

class PolizzaClienteSerializer(serializers.Serializer):
    cliente_id = serializers.IntegerField()
    data_sottoscrizione = serializers.DateField()
    scadenza = serializers.DateField()
    premio_assicurativo = serializers.FloatField()
    richiesta_sottoscrizione_id = serializers.IntegerField()

    def create(self, validated_data):
        return PolizzaCliente(**validated_data)