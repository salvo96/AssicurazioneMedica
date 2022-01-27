#Funzione per l'allenamento di un semplice modello di Regressione Lineare in grado di predire il prezzo del premio
#assicurativo basato sui dati medici del cliente

pricePredict <- function(age, sex, bmi, children, region, smoker){
  suppressMessages(suppressWarnings(library(corrplot)))
  suppressMessages(suppressWarnings(library(caret)))
  suppressMessages(suppressWarnings(library(tibble)))
  #<----------------------------DEFINIZIONE funzioni, oggetti e metodi--------------------------------->
  readDataset <- function(){  #funzione per la lettura e l'analisi iniziale del dataset
    dataset <- read.csv("insurance.csv")
    min_charge <- min(dataset$charges)
    max_charge <- max(dataset$charges)

    #visualizzazione dataset -> numero righe, numero features, valori e statistiche
    #glimpse(dataset)
    #summary(dataset)

    #Verifico la presenza di valori NA e li rimuovo
    na <- complete.cases(dataset)
    dataset <- dataset[na, ]

     #converto la colonna 'smoker' in valori true/false, quindi in logical
    dataset$smoker[dataset$smoker=="yes"] <- "T"
    dataset$smoker[dataset$smoker=="no"] <- "F"
    dataset$smoker <- as.logical(dataset$smoker)

    #converto la colonna 'sex' in valori "1","2"
    dataset$sex[dataset$sex=="male"] <- "1"
    dataset$sex[dataset$sex=="female"] <- "2"

    #converto la colonna 'region' in valori "1","2","3","4"
    dataset$region[dataset$region=="northeast"] <- "1"
    dataset$region[dataset$region=="northwest"] <- "2"
    dataset$region[dataset$region=="southeast"] <- "3"
    dataset$region[dataset$region=="southwest"] <- "4"

    #infine converto le colonne non logical in numeric
    x <- !sapply(dataset, is.logical) #utilizzo la sapply per eseguire la medesima operazione su più dati
    dataset[ , x] <- as.data.frame(apply(dataset[ , x], 2, as.numeric))

    return(list(dataset, min_charge, max_charge))

  }

  modelTrain <- function(dataset){ #funzione che effettua il training del modello e lo restituisce
    #Prima operazione: caricamento dataset assicurazioni
    insurance <- dataset #dataset letto

    #Data Pre-processing: una serie di step per analizzare il dataset e verificare che sia adatto per la costruzione del modello

    #Analisi correlazione delle features: features con alta correlazione sono approssimabili reciprocamente
    #si considera matrice di covarianza
    #res <- cor(insurance)
    #corrplot(res) #stampo la matrice di covarianza e osservo che non posso eliminare nessuna feature

    #Verfica relazione di linearità tra label = charges e bmi, age
    #plot(charges ~ age, data=insurance)#relazione di linearità più chiara
    #plot(charges ~ bmi, data=insurance)

    #Normalizzazione del dataset
    preproc <- preProcess(insurance[,c(1:4,6:7)], method=c("range"))
    norm <- predict(preproc, insurance[,c(1:4,6:7)])
    smoker <- insurance$smoker
    insurance_norm <- cbind(norm, smoker) #dataset completo e normalizzato

    #train del modello di regressione lineare
    insurance_norm.charges.lm <- lm(charges ~ age + sex + bmi + children + region + smoker , data = insurance_norm)
    #summary(insurance_norm.charges.lm)

    return(insurance_norm.charges.lm)
  }

  #Creo classe in S4 rappresentante l'oggetto usato per la predizione del prezzo del premio
  PricePredictor <- setClass(
    "PricePredictor",
          slots=c(age="numeric",
                  sex="numeric",
                  bmi="numeric",
                  children="numeric",
                  region="numeric",
                  smoker="logical",
                  predictedPrice="numeric"

          ))

  #di seguito viene settato il metodo per il calcolo del prezzo
  setGeneric("prediction", def=function(object, dataset, min_charge, max_charge){
    standardGeneric("prediction")
  })

  setMethod("prediction", signature = "PricePredictor", definition = function(object, dataset, min_charge, max_charge){ #funzione che effettua il training del modello e lo restituisce
    modello <- modelTrain(dataset)
    new_df <- data.frame(age=object@age, sex=object@sex, bmi=object@bmi, children=object@children, region=object@region, smoker=object@smoker)
    price_norm <- predict(modello, new_df)
    object@predictedPrice <- (max_charge - min_charge) * price_norm + min_charge
    return(object)
  })

  #<----------------------esecuzione della funzione--------------->
  values <- readDataset()
  dataset <- values[[1]] #lettura del dataset da file
  min_charge <- values[[2]]
  max_charge <- values[[3]]
  smoker <- as.logical(smoker)  #il campo smoker viene reso logical

  data <- data.frame(age=age, sex=sex, bmi=bmi, children=children, region=region) #creiamo un dataframe di una riga con i dati usati per la predizione

  preproc <- preProcess(dataset[,c(1:4,6)], method=c("range"))  #calcoliamo i valori per la normalizzazione
  norm <- predict(preproc, data[,c(1:5)]) #applichiamo la normalizzazione
  data_norm <- cbind(norm, smoker)  #dati normalizzati da poter fornire in input al modello

  model <- PricePredictor(age=data_norm$age, sex=data_norm$sex, bmi=data_norm$bmi, children=data_norm$children, region=data_norm$region, smoker=data_norm$smoker)
  predictor <- prediction(model, dataset, min_charge, max_charge)

  return (predictor@predictedPrice)
}
