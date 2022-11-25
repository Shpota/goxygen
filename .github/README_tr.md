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
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/by.svg">
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
        <img height="25px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/tr.svg">
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

<img src="../templates/react.webapp/src/logo.svg" align="right" width="230px" alt="goxygen logo">

**Go ve Angular, React veya Vue ile bir web projesi oluşturun.**

Goxygen, yeni bir proje oluştururken sizin zamandan tasarruf etmenize odaklanır. Sizin için bütün konfigürasyonları yapılmış bir uygulama iskeleti oluşturur. İş mantığınızı hemen uygulamaya başlayabilirsiniz. Goxygen, back-end Go kodu oluşturur, front-end tarafında componentlere bağlar, uygulama için bir Dockerfile sağlar ve production ortamında rahat çalışma için docker-compose dosyalarını oluşturur.

<table>
    <thead>
    <tr align="center">
        <td colspan=4><b>Desteklenen Teknolojiler</b></td>
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

## Gereksinimler

Makinenizde Go 1.11 veya daha yenisine sahip olmanız gerekmektedir.

`GO111MODULE` ortam değişkeni üretim mantığının çalışması için `auto` olarak ayarlanmalıdır. Bu ayar, Go 1.15'e kadar olan sürümlerde varsayılan ayardır. Ancak, Go 1.16 için, bunu açıkça ayarlamanız gerekmektedir:

```
export GO111MODULE=auto
```

## Nasıl Kullanılır

Çalıştır:

```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-app

```

`my-app` klasöründe bir proje oluşturur.

Varsayılan olarak React ve MongoDB'yi kullanacaktır. Farklı bir front-end framework'ü ve veritabanını `--frontend` ve `--db` flag'lerini kullanarak seçebilirsin. Örneğin bu komut Vue ve PostgreSQL ile bir proje oluşturacaktır:

```go
go run github.com/shpota/goxygen init --frontend vue --db postgres my-app
```

`--frontend` flag'i `angular`, `react` ve `vue` değerlerini kabul eder.
`--db` flag'i `mongo`, `mysql` ve `postgres` değerlerini kabul eder.

Oluşturulan proje `docker-compose` komutu ile çalıştırmaya hazır olacaktır:

```sh
cd my-app
docker-compose up
```

Build tamamlandıktan sonra, uygulama http://localhost:8080 adresinden erişilebilir olacaktır.

Oluşturulan projeyle nasıl çalışılacağı hakkında daha fazla ayrıntıyı README dosyasında bulabilirsiniz.

![Showcase](showcase.gif)

## Oluşturulmuş bir projenin yapısı (Örnek: React/MongoDB)

    my-app
    ├── server                   # Go proje dosyaları
    │   ├── db                   # MongoDB iletişimleri
    │   ├── model                # domain nesneleri
    │   ├── web                  # REST APIs, web server
    │   ├── server.go            # Sunucunun başlama noktası
    │   └── go.mod               # Sunucu bağımlılıkları
    ├── webapp
    │   ├── public               # ikonlar, sabit dosyalar, ve index.html
    │   ├── src
    │   │   ├── App.js           # Ana React komponenti
    │   │   ├── App.css          # Uygulama bileşenine özgü stiller
    │   │   ├── index.js         # Uygulamanın giriş noktası
    │   │   └── index.css        # global stiller
    │   ├── package.json         # front end bağımlılıkları
    │   ├── .env.development     # Geliştirme ortamı için API Endpointlerini barındırır.
    │   └── .env.production      # Production ortamı için API endpoint'leri
    ├── Dockerfile               # Back-end ve front-end'i birlikte oluşturur.
    ├── docker-compose.yml       # Üretim ortamı dağıtım tanımlayıcısı
    ├── docker-compose-dev.yml   # Geliştirme ihtiyaçları için yerel MongoDB çalıştırır.
    ├── init-db.js               # Test verileriyle bir MongoDB koleksiyonu oluşturur.
    ├── .dockerignore            # Docker yapılarında yok sayılan dosyaları belirtir.
    ├── .gitignore
    └── README.md                # Oluşturulan repo'nun nasıl kullanılacağına dair kılavuz.

Birim testleri veya örnek komponentler gibi dosyalar, basitlik açısından buraya dahil edilmemiştir.

## Bağımlılıklar

Goxygen, bir projenin temel yapısını oluşturur ve sizi belirli bir takım araçları kullanmaya zorlamaz. Bu yüzden projenize gereksiz bağımlılıklar getirmez. React ve Vue projelerinde yalnızca arka uç tarafında bir veritabanı sürücüsü ve [axios] (https://github.com/axios/axios) kullanır. Angular projeleri yalnızca Angular'a özgü kütüphaneleri kullanır.

## Nasıl katkıda bulunulur

Bir hata bulduysanız veya projeyi nasıl iyileştireceğiniz konusunda bir fikriniz varsa [issue oluşturun](https://github.com/Shpota/goxygen/issues) ve mümkün olan en kısa süre içerisinde düzelteceğiz. Ayrıca Pull Request ile önerebilirsin. Repository'i fork'layın, değişikleri yapın ve bize bir Pull Request gönderin ve kısa süre içerisinde inceleyeceğiz. Ayrıca tüm değişikleri tartıştığımız [Gitter chat](https://gitter.im/goxygen/community) var.

## Jenerik

Goxygen'in logosu [Egon Elbre](https://twitter.com/egonelbre) tarafından yaratılmıştır.
