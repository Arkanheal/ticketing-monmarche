<div align="center">

# Ticketing (mon-marche.fr)
</div>

Temps de réalisation : ~4h30

## Stack

1. Golang avec [Fiber](https://gofiber.io/). Suite à la consultation de divers benchmark
1. PostgreSQL
1. WIP: Docker, docker-compose

## WIP

1. Finalisation de l'installation via Docker
1. Mise en place d'un outil de migration de base type [golang-migrate](https://github.com/golang-migrate/migrate)
1. Optimisations du traitement du payload
1. Test Unitaire + test de charges
1. Ajout de middleware CORS/CSRF/Idempotency
1. Ajout d'un middleware de metrics
1. Log vers un logfile
1. Prise en charge de plusieurs `.env` (`.env.prod`, `.env.dev`)
1. Ajout d'un loadbalancing
1. Auth
1. HTTPS

## Installation (WIP)

#### Préréquis
* Go
* PostgreSQL
* Une base ticket (un exemple de configuration de la base est présente dans `.env`)

```
git clone https://github.com/Arkanheal/ticketing-monmarche.git
cd ticketing-monmarche
go mod download
go build -o build/ticket .
./build/ticket
```

Possibilité d'envoi de requêtes sur `:3000.ticket`, le payload est attendu en `text/plain` suivant le format renseigné dans l'exemple
