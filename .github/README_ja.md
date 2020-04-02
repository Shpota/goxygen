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

**Go、Angular / React / Vue、MongoDB を使用するWebプロジェクトを生成します。**

Goxygen は、新しいプロジェクトを始める際の時間の節約を目的としており、すべて
の設定が行われたアプリケーションの雛形を作成します。ビジネスロジックの実装をす
ぐに開始できます。Goxygen はバックエンドの Go 言語のコードを生成し、フロントエ
ンドコンポーネントに接続します。加えて、アプリケーション用の Dockerfile を提供
し、開発および本番環境での実行に便利な docker-compose ファイルを作成します。

## 使用方法
Go 1.11 以上が必要です。
```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-app
```
`my-app` プロジェクトを `my-app` フォルダに生成します。

デフォルトでは、React ベースのプロジェクトが生成されますが、`--frontend` フラグに
`angular`、`react`、`vue` を渡すことで、Angular、React、Vue から選択できます。
例:

```go
go run github.com/shpota/goxygen init --frontend vue my-app
```

生成されたプロジェクトは、`docker-compse` で実行する準備が整っています:
```sh
cd my-app
docker-compose up
```
ビルドが完了後、アプリケーションは http://localhost:8080 でアクセスできるように
なります。

生成されたプロジェクトの操作方法の詳細については、READMEファイルを参照してくだ
さい。

![Showcase](showcase.gif)

## 生成されるプロジェクトの構造 (Reactベースアプリケーション)

    my-app
    ├── server                   # Go プロジェクトファイル
    │   ├── db                   # MongoDB communications
    │   ├── model                # domain objects
    │   ├── web                  # REST APIs, web サーバ
    │   ├── server.go            # サーバの開始点
    │   └── go.mod               # サーバ依存関係
    ├── webapp                    
    │   ├── public               # アイコン、静的ファイル、index.html
    │   ├── src                       
    │   │   ├── App.js           # 主要な React コンポーネント
    │   │   ├── App.css          # アプリコンポーネント固有のスタイル
    │   │   ├── index.js         # アプリケーションのエントリポイント
    │   │   └── index.css        # 全体のスタイル
    │   ├── package.json         # フロントエンド依存関係
    │   ├── .env.development     # 開発環境の API エンドポイント保持
    │   └── .env.production      # プロダクション環境の API エンドポイント
    ├── Dockerfile               # バックエンド・フロントエンドを共に構築します
    ├── docker-compose.yml       # プロダクション環境の配置記述
    ├── docker-compose-dev.yml   # 開発に必要なローカルの MongoDB を実行する
    ├── init-db.js               # テストデータで MongoDB コレクションを作成します
    ├── .dockerignore            # Docker のビルドで無視するファイルを指定します
    ├── .gitignore
    └── README.md                # 生成したリポジトリの使用方法に関するガイド

簡単にするため、ここには単体テストやサンプルコンポーネントなどの
ファイルは含まれていません。

## 依存関係

Goxygen は、プロジェクトの基本的な構造を生成し、特定のツール群を使用する
ように強制しません。それが、プロジェクトに不要な依存関係をもたらさない
理由です。バックエンド側では
[mongo-go-driver](https://github.com/mongodb/mongo-go-driver) を
React 及び Vue プロジェクトでは、[axios](https://github.com/axios/axios) 
のみを使用します。Angular プロジェクトでは、Angular 固有のライブラリのみ
を使用します。

## How to contribute

If you found a bug or have an idea on how to improve the project
[open an issue](https://github.com/Shpota/goxygen/issues)
and we will fix it as soon as possible. You can also propose your
changes via a Pull Request. Fork the repository, make changes, send
us a pull request and we'll review it shortly. We also have a
[Gitter chat](https://gitter.im/goxygen/community) where we discuss
all the changes.

### Credits
Goxygen's logo was created by [Egon Elbre](https://twitter.com/egonelbre).
