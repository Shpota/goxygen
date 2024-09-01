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
        <img height="25px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/br.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_by.md">
        <img height="20px" src="https://raw.githubusercontent.com/Shpota/goxygen/main/.github/flag-by.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_fr.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/fr.svg">
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

**Crie um projeto Web com Go e Angular/React/Vue em segundos.**

Goxygen visa poupar seu tempo ao montar um novo projeto. Ele 
cria o esqueleto de uma aplicação com todas as configurações 
prontas para você. Você pode começar a implementar sua lógica 
de negócios imediatamente. Goxygen gera o código back end em Go, 
conecta ele com os componentes front end, fornece o Dockerfile
para a aplicação e cria o arquivo docker-compose para rapidamente
executar em ambientes de desenvolvimento e produção.

<table>
    <thead>
    <tr align="center">
        <td colspan=4><b>Tecnologias suportadas</b></td>
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
        <td>Database</td>
        <td>MongoDB</td>
        <td>MySQL</td>
        <td>PostgreSQL</td>
    </tr>
    </tbody>
</table>

## Requisitos
Você precisa de Go 1.11 ou mais recente na sua máquina.

## Como utilizar

Go 1.17 e posterior:
```go
go run github.com/shpota/goxygen@latest init my-app
```

<details>
  <summary>Versões anteriores de Go</summary>

### Go 1.16

Configure a variável de ambiente `GO111MODULE` como `auto`.
```
export GO111MODULE=auto
```
Rode os comandos
```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-app
```

### Go 1.11 - 1.15

Rode os comandos
```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-app
```
</details>

Isso gera um projeto no diretório `my-app`.

Por padrão, será usado React e MongoDB. Você pode selecionar
um framework de front end e um banco de dados diferentes usando
os argumentos `--frontend` e `--db`. Por exemplo, este comando
criará um projeto usando Vue e PostgreSQL:

```go
go run github.com/shpota/goxygen init --frontend vue --db postgres my-app
```

O argumento `--frontend` aceita as opções `angular`, `react` e `vue`.
O argumento `--db` aceita as opções `mongo`, `mysql` e `postgres`.

O projeto gerado está pronto para ser executado com `docker-compose`:
```sh
cd my-app
docker compose up
```

Após o build ser concluído, a aplicação ficará disponível em 
http://localhost:8080. 

Você pode encontrar mais detalhes de como utilizar o projeto
gerado em seu arquivo README.

![Showcase](showcase.gif)

## Estrutura do projeto gerado (Exemplo com React/MongoDB)

    my-app
    ├── server                   # Arquivos do projeto Go
    │   ├── db                   # Comunicação com o MongoDB
    │   ├── model                # Objetos de domínio
    │   ├── web                  # REST APIs, servidor web
    │   ├── server.go            # O ponto de acesso do servidor
    │   └── go.mod               # dependências do servidor
    ├── webapp                    
    │   ├── public               # ícones, arquivos estáticos e index.html
    │   ├── src                       
    │   │   ├── App.js           # O principal componente React
    │   │   ├── App.css          # Estilização do componente App
    │   │   ├── index.js         # Ponto de entrada da aplicação          
    │   │   └── index.css        # Estilos globais
    │   ├── package.json         # dependências do front end
    │   ├── .env.development     # API endpoint para ambiente de desenvolvimento
    │   └── .env.production      # API endpoint para ambiente de produção
    ├── Dockerfile               # faz o build do back end e o front end juntos
    ├── docker-compose.yml       # deploy em ambiente de produção
    ├── docker-compose-dev.yml   # executa um MongoDB local para fins de desenvolvimento
    ├── init-db.js               # cria um MongoDB collection com dados de teste
    ├── .dockerignore            # especifica arquivos ignorados no Docker build
    ├── .gitignore
    └── README.md                # guia sobre como usar o repositório gerado

Arquivos como testes unitários ou componentes de exemplo não estão incluídos 
aqui para fins de simplicidade.

## Dependências

Goxygen gera uma estrutura básica de um projeto e não o força usar um 
conjunto específico de ferramentas. E por isso não traz dependências 
desnecessárias para o seu projeto. Ele usa apenas
um driver de banco de dados no back end
e [axios](https://github.com/axios/axios) nos projetos React e Vue.
Projetos Angular usam apenas bibliotecas específicas do Angular.

## Como contribuir

Se você encontrou um bug ou tem alguma ideia de como melhorar o projeto,
[abra um issue](https://github.com/Shpota/goxygen/issues)
e nós iremos consertar o mais rápido possível. Você também pode enviar
suas alterações via Pull Request. Dê um fork no repositório, faça
alterações, nos envie um pull request que nós iremos analisar em breve. 
Nós também temos um
[Gitter chat](https://gitter.im/goxygen/community) onde discutimos todas
as alterações.

### Créditos
O logo do Goxygen foi criado por [Egon Elbre](https://twitter.com/egonelbre).
