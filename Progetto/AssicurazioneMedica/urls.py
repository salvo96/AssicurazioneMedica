from django.urls import path
from . import views

app_name = 'AssicurazioneMedica'
urlpatterns = [
    path('clienteLogin/', views.clienteLogin),
    path('amministratoreLogin/', views.amministratoreLogin),
    path('clienteRegistrazione/', views.clienteRegistrazione),
    path('richiestaSottoscrizione/', views.richiestaSottoscrizione),
    path('elencoPolizzeCliente/', views.elencoPolizzeCliente),
    path('inserimentoSpese/', views.inserimentoSpeseMediche),
    path('letturaSpese/', views.letturaSpeseMediche),
    path('datiMedici/', views.datiMediciClienti),
    path('polizzeClienti/', views.polizzeClienti),
    path('speseMedicheClienti/', views.speseMedicheClienti),
    path('richiesteSottoscrizioneClienti/', views.richiesteSottoscrizioneClienti),
    path('stimaPrezzoPolizza/', views.calcoloStimaPrezzo),
    path('inserimentoPolizza/', views.inserimentoPolizza)
]