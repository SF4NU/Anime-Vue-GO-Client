# [Vai Al Sito](https://sf4nu.github.io/Anime-Vue-GO-Client/)

![Immagine 2024-04-12 124229](https://github.com/SF4NU/Anime-Vue-GO-Client/assets/129513838/a9acbbaa-9d8f-4488-acb7-ba468801d3c9)

## Panoramica

Anime-Vue-GO-Client è un'applicazione web creata per gli appassionati di anime e manga, offrendo un'esperienza senza interruzioni per scoprire, monitorare e gestire i propri titoli preferiti. Sfruttando l'API di Kitsu per dati completi su anime e manga, questa piattaforma consente agli utenti di esplorare una vasta libreria di contenuti.

## Funzionalità

- **Registrazione e Autenticazione Utenti**: Utilizzando un backend robusto scritto in Go, supportato dal framework Fiber e GORM, Anime-Vue-GO-Client offre un'autenticazione e registrazione utente sicure. L'autenticazione basata su JWT garantisce l'integrità dei dati e la privacy degli utenti.
![Immagine 2024-04-12 124348](https://github.com/SF4NU/Anime-Vue-GO-Client/assets/129513838/3de4eb55-c507-40e7-8476-8360d4eac7a7)

- **Liste Personalizzate**: Gli utenti possono creare liste personalizzate per tenere traccia di anime e manga che hanno visto o intendono vedere.
![Immagine 2024-04-12 124440](https://github.com/SF4NU/Anime-Vue-GO-Client/assets/129513838/6c553cbe-86a0-461e-942a-b82319743433)

- **Gestione del Database**: Anime-Vue-GO-Client utilizza PostgreSQL come sistema di gestione del database.

## Installazione

Per configurare Anime-Vue-GO-Client localmente, seguire questi passaggi:

1. Clonare il repository: `git clone https://github.com/username/Anime-Vue-GO-Client.git`
2. Navigare nella directory del progetto: `cd Anime-Vue-GO-Client`
3. Installare le dipendenze per il frontend e il backend:
   - Frontend: `cd Client && npm install`
   - Backend: `cd Server && go mod tidy`
4. Configurare il database:
   - Creare un database PostgreSQL e configurare le impostazioni di connessione nel file di configurazione del backend.
   - Eseguire le migrazioni del database per creare le tabelle necessarie.
5. Avviare il server backend: `go run main.go` (o compilare ed eseguire il binario)
6. Avviare il server di sviluppo frontend: `npm run dev`

## Cosa ho utilizzato

- [Kitsu API](https://kitsu.docs.apiary.io/): Fornisce dati completi su anime e manga.
- [Fiber](https://github.com/gofiber/fiber): Framework web GO.
- [Vue.js](https://vuejs.org/): Il framework  Vue.
- [GORM](https://gorm.io/): La libreria ORM per Golang.
