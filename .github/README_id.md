<h1 align="center">
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README.md">
        <img height="25px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/gb.svg">
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
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/fr.svg">
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

**Menyiapkan web projek menggunakan Go, Angular/React/Vue, dan MongoDB dalam hitungan detik.**

Goxygen bertujuan untuk menghemat waktu Anda dalam menyiapkan sebuah proyek web baru. Goxygen menciptakan kerangka aplikasi dengan semua konfigurasi yang telah disiapkan untuk Anda. Jadi Anda dapat langsung segera memulai menerapkan logika bisnis pada proyek Anda.
Goxygen menghasilkan kode Go yang berfungsi sebagai back-end, menghubungkannya dengan komponen front-end
, menyediakan Dockerfile untuk aplikasi docker dan membuat file docker-compose agar mudah dijalankan dalam fase pengembangan dan fase produksi.

## Cara penggunaan
Anda harus mempunyai Go versi 1.11 atau yang lebih pada mesin Anda.

```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-app
```
Ini akan meng-generate sebuah projek di dalam `my-app` folder. 

Secara default, proyek dengan React-based akan di dihasiklan. Anda juga dapat menggunakan framework lain mulai dari Angular, React dan Vue dengan mengambahkan `angular`, `react` dan 
`vue` pada parameter `--frontend`. Misalnya:

```go
go run github.com/shpota/goxygen init --frontend vue my-app
```

Proyek yang telah di generate dapat dijalankan dengan perintah `docker-compose`:
```sh
cd my-app
docker-compose up
```
Setelah proses build selesai, aplikasi dapat diakses
pada http://localhost:8080.

Anda dapat menemukan detail lebih lanjut tentang bagaimana cara bekerja dengan proyek dengan membuka file README.

![Showcase](showcase.gif)

## Struktur proyek yang dihasilkan (React-based app)

    my-app
    ├── server                   # Proyek file go
    │   ├── db                   # Komunikasi dengan MongoDB
    │   ├── model                # domain objek
    │   ├── web                  # REST APIs, web server
    │   ├── server.go            # titik awal dari server
    │   └── go.mod               # server dependensi
    ├── webapp                    
    │   ├── public               # icons, static file, dan index.html
    │   ├── src                       
    │   │   ├── App.js           # titik awal komponen React
    │   │   ├── App.css          # App component-specific style
    │   │   ├── index.js         # titik awal dari the application          
    │   │   └── index.css        # style global
    │   ├── package.json         # front end dependensi
    │   ├── .env.development     # memegang API endpoint untuk fase pengembangan
    │   └── .env.production      # API endpoint untuk fase produksi
    ├── Dockerfile               # membangun back end dan front-end secara bersamaan.
    ├── docker-compose.yml       # prod environemnt deployment deskriptor
    ├── docker-compose-dev.yml   # menjalankan local MongoDB untuk kepentingan pengembangan
    ├── init-db.js               # membuat koleksi MongoDB dengan data uji
    ├── .dockerignore            # menentukan file yang diabaikan dalam pembuatan Docker
    ├── .gitignore
    └── README.md                # panduan tentang cara menggunakan repo yang dihasilkan

File seperti unit test atau komponen sampel tidak termasuk di sini
untuk kesederhanaan.

## Dependensi

Goxygen menghasilkan struktur dasar suatu proyek dan tidak memaksa Anda
untuk menggunakan seperangkat alat tertentu. Itu sebabnya Goxygen tidak membawa dependensi yang tidak dibutuhkan pada proyek Anda. Hanya menggunakan
[mongo-go-driver](https://github.com/mongodb/mongo-go-driver) pada
back end side dan [axios](https://github.com/axios/axios) pada projek React
dan Vue. Angular proyek hanya menggunakan Angular library tertentu.

## Bagaimana cara berkontribusi

Jika Anda menemukan bug atau punya ide tentang cara meningkatkan proyek
[open an issue](https://github.com/Shpota/goxygen/issues)
dan kami akan memperbaikinya sesegera mungkin. Anda juga dapat mengusulkan
perubahan melalui Pull Request. Fork repositori, buat perubahan, dan kirim
kami pull request dan kami akan segera memeriksanya. Kami juga punya
[Gitter chat](https://gitter.im/goxygen/community) dimana kita bahas
semua perubahan.

### Penghargaan
Goxygen's logo dibuat oleh [Egon Elbre](https://twitter.com/egonelbre).