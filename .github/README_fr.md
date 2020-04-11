<h1 align="center">
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/gb.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_zh.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/cn.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_ua.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/ua.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_ru.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/ru.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_ko.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/kr.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_pt-br.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/br.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_by.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/by.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_fr.md">
        <img height="25px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/fr.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_es.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/es.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_jp.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/jp.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_id.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/id.svg">
    </a>
    <br>
    Goxygen
    <a href="https://github.com/Shpota/goxygen/actions?query=workflow%3Abuild">
        <img src="https://github.com/Shpota/goxygen/workflows/build/badge.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/releases">
        <img src="https://img.shields.io/badge/version-v0.2.2-green">
    </a>
    <a href="https://gitter.im/goxygen/community">
        <img src="https://badges.gitter.im/goxygen/community.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/pulls">
        <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg">
    </a>
</h1>

<img src="../templates/react.webapp/src/logo.svg" align="right" width="230px" alt="goxygen logo">

**Générer un projet web avec Go, Angular/React/Vue, et MongoDB en quelques secondes.**

Goxygen vise à vous faire gagner du temps lors de la mise en place d'un nouveau projet. Il
crée un squelette d'application avec une configuration complète par défaut.
Vous pouvez commencer immédiatement à implémenter votre logique métier.
Goxygen génère du code Backend en Go, le connecte aux composants Frontend, fournit un 
Dockerfile pour l'application et crée des fichiers Docker-compose pour une exécution pratique 
dans des environnements de développement et de production.

## Guide d'utilisation

Vous devez disposer de la version 1.11 ou supérieure de Go sur votre machine.
```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-app
```
Ces commandes génèrent un projet dans le dossier `my-app`.

Par défaut, le projet est généré avec un frontend React. Vous pouvez choisir parmis
Angular, React ou Vue en passant les arguments `angular`, `react` ou `vue` au flag `--frontend`. 
Par exemple:

```go
go run github.com/shpota/goxygen init --frontend vue my-app
```

Le projet généré est prêt à l'utilisation avec `docker-compose`:
```sh
cd my-app
docker-compose up
```

Aprés la fin du build, l'application est accessible sur http://localhost:8080.

Vous trouverez plus de détails sur l'utilisation du projet généré dans son fichier README.

![Présentation](showcase.gif)

## Structure d'un projet généré (application React)

    my-app
    ├── server                   # fichier du projet Go
    │   ├── db                   # communications avec MongoDB
    │   ├── model                # objets du domaine
    │   ├── web                  # REST APIs, serveur web
    │   ├── server.go            # point d'entrée du serveur
    │   └── go.mod               # dépendances du serveur
    ├── webapp                    
    │   ├── public               # icones, fichiers statiques, et index.html
    │   ├── src                       
    │   │   ├── App.js           # le composant React principale
    │   │   ├── App.css          # style spécifiques au composant App
    │   │   ├── index.js         # le point d'entrée de l'application          
    │   │   └── index.css        # styles globaux
    │   ├── package.json         # dépendances du frontend
    │   ├── .env.development     # endpoint de l'API pour l'environnement de développement
    │   └── .env.production      # endpoint de l'API pour l'environnement de production
    ├── Dockerfile               # build du backend et du frontend
    ├── docker-compose.yml       # description des déploiements dans un environnement de production
    ├── docker-compose-dev.yml   # lance une instance de MongoDB en local pour développement
    ├── init-db.js               # crée une collection MongoDB avec des données de test
    ├── .dockerignore            # spécifie les fichiers ignorés lors des builds docker
    ├── .gitignore
    └── README.md                # guide d'utilisation du projet généré

Les fichiers de tests unitaires ou d'exemples de composants ne sont pas inclus
pour des raisons de simplicité.

## Dépendances

Goxygen génére un projet avec une structure basique et ne vous force pas à utiliser des 
outils spécifiques. C'est pour cela que les seules dépendances utilisées dans le projet 
sont [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) côté backend et 
[axios](https://github.com/axios/axios) côté fronted les projets React et Vue. Les projets 
Angular utilisent leurs propres librairies.

## Comment contribuer

Si vous trouvez un bug ou avez une idée sur comment améliorer le projet
[créez une issue](https://github.com/Shpota/goxygen/issues)
et nous corrigerons le problème le plus rapidement possible. Vous pouvez aussi
proposer vos propres changement avec des Pull Request. Faites un Fork du dépôt, 
effectuer vos modifications, envoyez-nous un pull request et nous le consulterons dans 
les plus brefs délais. Nous disposons aussi d'un [chat Gitter](https://gitter.im/goxygen/community) 
où nous discutons tous les changements.

### Crédits

Le logo de Goxygen à été crée par [Egon Elbre](https://twitter.com/egonelbre).
