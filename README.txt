---------------------------------------------------README.txt---------------------------------------------------
ISTRUZIONI PER L'INSTALLAZIONE:
SERVER: Folder Progetto
1)Installare MySQL Server: https://dev.mysql.com/downloads/mysql/
	1.1)Configurare MySQL Server con dati 'root', 'password' e porta '3306'
	1.2)Creare un database con nome 'assicurazione' da associare all'applicazione
	1.3)Avviare il servizio associato al server
2)Installare Python
	2.1) Installare il framework "Django": $ python -m pip install Django
3)Configurare il framework per farlo lavorare con il server
	3.1) pip install pymysql
	3.2) Settare i parametri in Progetto/Progetto/settings.py:
		'ENGINE': 'django.db.backends.mysql',
        	'NAME': 'assicurazione',
        	'USER': 'root',
        	'PASSWORD': 'password',
        	'HOST': '127.0.0.1',
        	'PORT': '3306',
4)Avviare la creazione delle tabelle:
	4.1) $ python manage.py makemigrations AssicurazioneMedica
	     $ python manage.py migrate
5)Installare "Django REST Framework"
	5.1) pip install djangorestframework
6)Installare "pytz"
	6.1) pip install pytz
7)Installare "rpy2":
	7.1)pip install rpy2

CLIENT: Folder ProgettoClient
Il Client non richiede alcuna installazione ma solo l'avvio dell'eseguibile una volta avviato il server