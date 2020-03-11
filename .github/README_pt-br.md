<h1 align="center">
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/gb.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_ua.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/ua.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_ru.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/ru.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_zh.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/cn.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_ko.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/kr.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_pt-br.md">
        <img height="25px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/br.svg">
    </a>
    <br>
    Goxygen
    <a href="https://github.com/Shpota/goxygen/actions?query=workflow%3Abuild">
        <img src="https://github.com/Shpota/goxygen/workflows/build/badge.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/releases">
        <img src="https://img.shields.io/badge/version-v0.2.1-green">
    </a>
    <a href="https://gitter.im/goxygen/community">
        <img src="https://badges.gitter.im/goxygen/community.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/pulls">
        <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square">
    </a>
</h1>

<img src="../templates/react.webapp/src/logo.svg" align="right" width="230px" alt="goxygen logo">

**Crie um projeto Web com Go, Angular/React/Vue, e MongoDB em segundos.**

Goxygen visa poupar seu tempo ao montar um novo projeto. Ele 
cria o esqueleto de uma aplicação com todas as configurações 
prontas para você. Você pode começar a implementar sua lógica 
de negócios imediatamente. Goxygen gera o código back end Go, 
conecta ele com os componentes front end, fornece o Dockerfile
para a aplicação e cria o arquivo docker-compose para rapidamente
executar em ambientes de desenvolvimento e produção.



## Como utilizar
Você precisa ter Go 1.11 ou mais recente na sua máquina.
```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-app
```
Isso gera um projeto no diretório `my-app`.

Por padrão, um React-based project é gerado. Você pode escolher 
entre Angular, React e Vue passando `angular`, `react` e `vue` 
para a flag `--frontend`. Por exemplo:

```go
go run github.com/shpota/goxygen init --frontend vue my-app
```

O projeto gerado está pronto para ser executado com `docker-compose`:
```sh
cd my-app
docker-compose up
```

Após o build ser concluído, a aplicação está disponível em 
http://localhost:8080. 

Você pode encontrar mais detalhes de como trabalhar com o projeto
gerado em seu arquivo README.

![Showcase](showcase.gif)

## Estrutura do projeto gerado (React-based app)

    my-app
    ├── server                   # Go project files
    │   ├── db                   # MongoDB communications
    │   ├── model                # domain objects
    │   ├── web                  # REST APIs, web server
    │   ├── server.go            # O ponto inicial do server
    │   └── go.mod               # dependências do server
    ├── webapp                    
    │   ├── public               # ícones, arquivos estáticos e index.html
    │   ├── src                       
    │   │   ├── App.js           # O principal componente React
    │   │   ├── App.css          # Estilização do componente App
    │   │   ├── index.js         # Ponto de entrada da aplicação          
    │   │   └── index.css        # global styles
    │   ├── package.json         # dependências do front end
    │   ├── .env.development     # API endpoint para ambiente de desenvolvimento
    │   └── .env.production      # API endpoint para ambiente de produção
    ├── Dockerfile               # faz o build do back end e o front end juntos
    ├── docker-compose.yml       # deploy em abiente de produção
    ├── docker-compose-dev.yml   # executa um MongoDB local para fins de desenvolvimento
    ├── init-db.js               # cria um MongoDB collection com dados de teste
    ├── .dockerignore            # especifica arquivos ignorados no Docker build
    ├── .gitignore
    └── README.md                # guia sobre como usar o repositório gerado

Arquivos como testes de unidade ou sample components não estão incluídos 
aqui pela simplicidade.

## Dependências

Goxygen gera uma estrutura básica de um projeto e não o força usar um 
conjunto específico de ferramentas. E por isso não traz dependências 
desnecessárias para o seu projeto. Ele usa apenas 
[mongo-go-driver](https://github.com/mongodb/mongo-go-driver) no back end 
e [axios](https://github.com/axios/axios) nos projetos React e Vue. 
Projetos Angular usam apenas bibliotecas específicas do Angular.

## Como contribuir

Se você encontrou um bug ou tem alguma idéia de como melhorar o projeto,
[abra uma issue](https://github.com/Shpota/goxygen/issues)
e nós iremos consertar o mais rápido possível. Você também pode enviar
suas alterações via Pull Request. Dê um fork no repositório, faça
alterações, nos envie um pull request que nós iremos analisar em breve. 
Nós também temos um
[Gitter chat](https://gitter.im/goxygen/community) onde discutimos todas
as alterações.

### Créditos
O logo do Goxygen foi criado por [Egon Elbre](https://twitter.com/egonelbre).
