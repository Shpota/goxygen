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
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/fr.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_es.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/es.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_jp.md">
        <img height="25px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/jp.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_id.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/id.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_he.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/il.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_tr.md">
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

<img src="../templates/react.webapp/src/logo.svg" align="right" width="230px" alt="goxygen logo">

**Go、Angular / React / Vue を用いた Web プロジェクトを生成します。**

Goxygen は、新しいプロジェクトを始める際の時間の節約を目的としており、すべての
設定が済まされたアプリケーションの雛形を作成します。ビジネスロジックの実装をす
ぐに開始できます。Goxygen はバックエンドの Go 言語のコードを生成し、それをフロ
ントエンドコンポーネントと関連づけます。加えて、アプリケーション用の Dockerfile
を提供し、開発および本番環境での実行に便利な docker-compose ファイルを作成しま
す。

<table>
    <thead>
    <tr align="center">
        <td colspan=4><b>対応技術</b></td>
    </tr>
    </thead>
    <tbody>
    <tr align="center">
        <td align="center">フロントエンド</td>
        <td>Angular</td>
        <td>React</td>
        <td>Vue</td>
    </tr>
    <tr align="center">
        <td>バックエンド</td>
        <td colspan=3>Go</td>
    </tr>
    <tr align="center">
        <td>データベース</td>
        <td>MongoDB</td>
        <td>MySQL</td>
        <td>PostgreSQL</td>
    </tr>
    </tbody>
</table>

## Requirements
Go 1.11 以上が必要です。

The `GO111MODULE` environment variable has to be set to `auto`
for the generation logic to work. It is a default for Go
versions up to 1.15. For Go 1.16, you need to set it explicitly:
```
export GO111MODULE=auto
```

## 使用方法
```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-app
```
`my-app` プロジェクトを `my-app` フォルダに生成します。

デフォルトでは、React 及び MongoDB を使用します。
`--frontend` と `--db` フラグを使って異なる
フロントエンドフレームワークとデータベースを選べます。
例えば、次のコマンドは Vue と PostgreSQL を含むプロジェクト作成します:

```go
go run github.com/shpota/goxygen init --frontend vue --db postgres my-app
```

`--frontend` フラグは、`angular` や `react`、`vue` を受け取ります。
`--db` フラグは、`mongo` や `mysql`、`postgres` を受け取ります。

生成されたプロジェクトは、`docker-compse` で実行する準備が整っています:
```sh
cd my-app
docker-compose up
```
ビルドが完了後、アプリケーションは http://localhost:8080 でアクセスできるように
なります。

生成されたプロジェクトの操作方法に関する詳細については、READMEファイルを参照し
てください。

![Showcase](showcase.gif)

## 生成されるプロジェクトの構造 (React/MongoDBベースアプリケーション)

    my-app
    ├── server                   # Go プロジェクトファイル
    │   ├── db                   # MongoDB との通信
    │   ├── model                # ドメインオブジェクト
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

単純にするため、ここには単体テストやサンプルコンポーネントなどのファイルは含ま
れていません。

## 依存関係

Goxygen は、プロジェクトの基本的な構造を生成し、特定のツール群を使用するように
強制しません。それが、プロジェクトに不要な依存関係をもたらさない理由です。バッ
クエンド側では database driver を、
React 及び Vue プロジェクトでは、[axios](https://github.com/axios/axios) のみを
使用します。Angular プロジェクトでは、Angular 固有のライブラリのみを使用します。

## 貢献の仕方

バグを発見した場合、またはプロジェクトを改善する方法についてアイディアがある場
合 [open an issue](https://github.com/Shpota/goxygen/issues)、できるだけ早く修
正します。Pull Request を通して変更を提案することもできます。リポジトリを Fork
し、変更を加え、Pull Request を送信してください。すぐに確認をします。また、あら
ゆる変更点について話し合う [Gitter chat](https://gitter.im/goxygen/community)
もあります。

### クレジット
Goxygen のロゴは、[Egon Elbre](https://twitter.com/egonelbre) によって作成され
ました。
