<h1 align="center">
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/gb.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_zh.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/cn.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_ua.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/ua.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_ko.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/kr.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_pt-br.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/br.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_by.md">
        <img height="20px" src="https://raw.githubusercontent.com/Shpota/goxygen/main/.github/flag-by.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_fr.md">
        <img height="25px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/fr.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_es.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/es.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_jp.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/jp.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_id.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/id.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_he.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/il.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_tr.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/tr.svg">
    </a>
    <br>
    Goxygen
    <a href="https://github.com/Shpota/goxygen/actions?query=workflow%3Abuild">
        <img src="https://github.com/Shpota/goxygen/workflows/build/badge.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/releases">
        <img src="https://img.shields.io/github/v/tag/shpota/goxygen?color=green&label=version">
    </a>
    <a href="https://gitter.im/goxygen/community">
        <img src="https://badges.gitter.im/goxygen/community.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/pulls">
        <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg">
    </a>
</h1>

<img src="../templates/vue.webapp/src/assets/logo.svg" align="right" width="230px" alt="goxygen logo">

**Générer un projet web avec Go et Angular/React/Vue en quelques secondes.**

Goxygen vise à vous faire gagner du temps lors de la mise en place d'un nouveau projet. Il
crée un squelette d'application avec une configuration complète par défaut.
Vous pouvez commencer immédiatement à implémenter votre logique métier.
Goxygen génère du code Backend en Go, le connecte aux composants Frontend, fournit un
Dockerfile pour l'application et crée des fichiers Docker-compose pour une exécution pratique
dans des environnements de développement et de production.

<table>
    <thead>
    <tr align="center">
        <td colspan=4><b>Technoligies supportées</b></td>
    </tr>
    </thead>
    <tbody>
    <tr align="center">
        <td align="center">Front End</td>
        <td>Angular</td>
        <td>React</td>
        <td>Vue</td>
    </tr>
    <tr align="center">
        <td>Back End</td>
        <td colspan=3>Go</td>
    </tr>
    <tr align="center">
        <td>Base de données</td>
        <td>MongoDB</td>
        <td>MySQL</td>
        <td>PostgreSQL</td>
    </tr>
    </tbody>
</table>

## Requirements
Vous devez disposer de la version 1.11 ou supérieure de Go sur votre machine.

The `GO111MODULE` environment variable has to be set to `auto`
for the generation logic to work. It is a default for Go
versions up to 1.15. For Go 1.16, you need to set it explicitly:
```
export GO111MODULE=auto
```

## Guide d'utilisation

```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-app
```

Ces commandes génèrent un projet dans le dossier `my-app`.

Par défaut, React et MongoDB seront utilisés. Vous pouvez
choisir divers frameworks frontend ou bases de données en
utilisant les flags `--frontend` et `--db`. Par exemple, cette
commande va créer un projet avec Vue et PostreSQL:

```go
go run github.com/shpota/goxygen init --frontend vue --db postgres my-app
```

Le flag `--frontend` accepte les valeurs `angular`, `react` and `vue`.
Le flag `--db` accepte les valeurs `mongo`, `mysql` and `postgres`.

Le projet généré est prêt à l'utilisation avec `docker-compose`:

```sh
cd my-app
docker-compose up
```

Aprés la fin du build, l'application est
accessible sur http://localhost:8080.

Vous trouverez plus de détails sur l'utilisation
du projet généré dans son fichier README.

![Présentation](showcase.gif)

## Structure d'un projet généré (application React/MongoDB)

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
sont database driver côté backend et [axios](https://github.com/axios/axios) côté
frontend les projets React et Vue. Les projets Angular utilisent leurs propres librairies.

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
