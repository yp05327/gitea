# Gitea

[![](https://github.com/go-gitea/gitea/actions/workflows/release-nightly.yml/badge.svg?branch=main)](https://github.com/go-gitea/gitea/actions/workflows/release-nightly.yml?query=branch%3Amain "Release Nightly")
[![](https://img.shields.io/discord/322538954119184384.svg?logo=discord&logoColor=white&label=Discord&color=5865F2)](https://discord.gg/Gitea "Join the Discord chat at https://discord.gg/Gitea")
[![](https://goreportcard.com/badge/code.gitea.io/gitea)](https://goreportcard.com/report/code.gitea.io/gitea "Go Report Card")
[![](https://pkg.go.dev/badge/code.gitea.io/gitea?status.svg)](https://pkg.go.dev/code.gitea.io/gitea "GoDoc")
[![](https://img.shields.io/github/release/go-gitea/gitea.svg)](https://github.com/go-gitea/gitea/releases/latest "GitHub release")
[![](https://www.codetriage.com/go-gitea/gitea/badges/users.svg)](https://www.codetriage.com/go-gitea/gitea "Help Contribute to Open Source")
[![](https://opencollective.com/gitea/tiers/backers/badge.svg?label=backers&color=brightgreen)](https://opencollective.com/gitea "Become a backer/sponsor of gitea")
[![](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT "License: MIT")
[![Contribute with Gitpod](https://img.shields.io/badge/Contribute%20with-Gitpod-908a85?logo=gitpod&color=green)](https://gitpod.io/#https://github.com/go-gitea/gitea)
[![](https://badges.crowdin.net/gitea/localized.svg)](https://translate.gitea.com "Crowdin")

この行はテスト用に追加されています。

[English](./README.md) | [简体中文](./README.zh-cn.md) | [繁體中文](./README.zh-tw.md) | [日本語](./README.ja-JP.md)

## 目的

このプロジェクトの目標は、自前で Git サービスを立ち上げる作業をできるだけ簡単・高速・ストレスなく行えるようにすることです。

Gitea は Go で書かれているため、Go がサポートする **すべて** のプラットフォームとアーキテクチャ（Linux、macOS、Windows の x86、amd64、ARM、PowerPC など）で動作します。このプロジェクトは 2016 年 11 月に [Gogs](https://gogs.io) から [フォーク](https://blog.gitea.com/welcome-to-gitea/) されて以来、大きく進化してきました。

オンラインデモは [demo.gitea.com](https://demo.gitea.com) でご覧いただけます。

リポジトリ数に制限のある無償の Gitea サービスを利用したい場合は [gitea.com](https://gitea.com/user/login) をご利用ください。

Gitea Cloud で専用の Gitea インスタンスを素早くデプロイしたい場合は、[cloud.gitea.com](https://cloud.gitea.com) から無料トライアルを開始できます。

## ドキュメント

公式の [ドキュメントサイト](https://docs.gitea.com/) には必要な情報がまとまっています。

インストール、管理、使い方、開発、コントリビュートガイドなど、すべての機能を効果的に理解するための内容を取りそろえています。

提案がある方やドキュメントに貢献したい方は、[ドキュメントリポジトリ](https://gitea.com/gitea/docs) をご覧ください。

## ビルド

ソースツリーのルートで次を実行します:

    TAGS="bindata" make build

SQLite が必要な場合:

    TAGS="bindata sqlite sqlite_unlock_notify" make build

`build` ターゲットは次の 2 つのサブターゲットに分かれています:

- `make backend`: [Go Stable](https://go.dev/dl/) が必要です。必要なバージョンは [go.mod](/go.mod) に記載されています。
- `make frontend`: [Node.js LTS](https://nodejs.org/en/download/) 以降と [pnpm](https://pnpm.io/installation) が必要です。

Go モジュールと npm モジュールを取得するためにインターネット接続が必要です。事前にビルド済みのフロントエンドファイルを含む公式ソースのアーカイブからビルドする場合、`frontend` ターゲットは実行されないため Node.js を用意する必要はありません。

詳細: https://docs.gitea.com/installation/install-from-source

## 使い方

ビルド後、デフォルトではソースツリーのルートに `gitea` という実行ファイルが生成されます。実行するには次のコマンドを使用します:

    ./gitea web

> [!NOTE]
> API を利用したい方のために、[ドキュメント](https://docs.gitea.com/api) 付きの実験的サポートを提供しています。

## コントリビューション

推奨されるワークフロー: Fork -> Patch -> Push -> Pull Request

> [!NOTE]
>
> 1. **Pull Request に取り組む前に必ず [貢献ガイド](CONTRIBUTING.md) をお読みください。**
> 2. プロジェクトで脆弱性を見つけた場合は、**security@gitea.io** まで非公開でご連絡ください。ご協力ありがとうございます！

## 翻訳

[![Crowdin](https://badges.crowdin.net/gitea/localized.svg)](https://translate.gitea.com)

翻訳は [Crowdin](https://translate.gitea.com) を通じて行っています。新しい言語を追加したい場合は、Crowdin プロジェクトのマネージャーに連絡して言語を有効化してもらってください。

言語追加の要望は issue を作成するか、Discord の #translation チャンネルで相談することもできます。文脈が必要なときや翻訳の問題を見つけたときは、文字列にコメントしたり Discord で質問してください。翻訳に関する一般的な質問についてはドキュメントの専用セクションをご覧ください。まだ情報は多くありませんが、寄せられた質問に応じて充実させていく予定です。

詳しくは [ドキュメント](https://docs.gitea.com/contributing/localization) をご確認ください。

## 公式およびサードパーティープロジェクト

Gitea には公式の [go-sdk](https://gitea.com/gitea/go-sdk)、CLI ツールの [tea](https://gitea.com/gitea/tea)、そして Gitea Action 用の [action runner](https://gitea.com/gitea/act_runner) が用意されています。

さらに Gitea 関連プロジェクトをまとめたリスト [gitea/awesome-gitea](https://gitea.com/gitea/awesome-gitea) を維持しており、SDK、プラグイン、テーマなど多彩なサードパーティープロジェクトを見つけられます。

## コミュニケーション

[![](https://img.shields.io/discord/322538954119184384.svg?logo=discord&logoColor=white&label=Discord&color=5865F2)](https://discord.gg/Gitea "Join the Discord chat at https://discord.gg/Gitea")

ドキュメントで疑問が解決しない場合は、[Discord サーバー](https://discord.gg/Gitea) に参加するか、[Discourse フォーラム](https://forum.gitea.com/) に投稿してください。

## 作者

- [メンテナー](https://github.com/orgs/go-gitea/people)
- [コントリビューター](https://github.com/go-gitea/gitea/graphs/contributors)
- [翻訳者](options/locale/TRANSLATORS)

## 支援者

すべての支援者に感謝します！🙏 [[支援する](https://opencollective.com/gitea#backer)]

<a href="https://opencollective.com/gitea#backers" target="_blank"><img src="https://opencollective.com/gitea/backers.svg?width=890"></a>

## スポンサー

スポンサーとしてプロジェクトを支援してください。あなたのロゴがここに表示され、ウェブサイトへのリンクが掲載されます。[[スポンサーになる](https://opencollective.com/gitea#sponsor)]

<a href="https://opencollective.com/gitea/sponsor/0/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/0/avatar.svg"></a>
<a href="https://opencollective.com/gitea/sponsor/1/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/1/avatar.svg"></a>
<a href="https://opencollective.com/gitea/sponsor/2/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/2/avatar.svg"></a>
<a href="https://opencollective.com/gitea/sponsor/3/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/3/avatar.svg"></a>
<a href="https://opencollective.com/gitea/sponsor/4/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/4/avatar.svg"></a>
<a href="https://opencollective.com/gitea/sponsor/5/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/5/avatar.svg"></a>
<a href="https://opencollective.com/gitea/sponsor/6/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/6/avatar.svg"></a>
<a href="https://opencollective.com/gitea/sponsor/7/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/7/avatar.svg"></a>
<a href="https://opencollective.com/gitea/sponsor/8/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/8/avatar.svg"></a>
<a href="https://opencollective.com/gitea/sponsor/9/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/9/avatar.svg"></a>

## よくある質問

**Gitea はどのように発音しますか？**

Gitea の発音は [/ɡɪ’ti:/](https://youtu.be/EM71-2uDAoY) で、「gi-tea」のように g を強く発音します。

**なぜこのプロジェクトは Gitea インスタンス上でホストされていないのですか？**

現在[対応中です](https://github.com/go-gitea/gitea/issues/1029)。

**セキュリティパッチはどこで確認できますか？**

[リリースノート](https://github.com/go-gitea/gitea/releases) または [変更履歴](https://github.com/go-gitea/gitea/blob/main/CHANGELOG.md) で `SECURITY` を検索してください。

## ライセンス

このプロジェクトは MIT ライセンスで提供されています。
全文は [LICENSE](https://github.com/go-gitea/gitea/blob/main/LICENSE) をご覧ください。

## さらに詳しく

<details>
<summary>インターフェースの概要を知りたい方はこちら！</summary>

### ログイン／登録ページ

![Login](https://dl.gitea.com/screenshots/login.png)
![Register](https://dl.gitea.com/screenshots/register.png)

### ユーザーダッシュボード

![Home](https://dl.gitea.com/screenshots/home.png)
![Issues](https://dl.gitea.com/screenshots/issues.png)
![Pull Requests](https://dl.gitea.com/screenshots/pull_requests.png)
![Milestones](https://dl.gitea.com/screenshots/milestones.png)

### ユーザープロフィール

![Profile](https://dl.gitea.com/screenshots/user_profile.png)

### 探索

![Repos](https://dl.gitea.com/screenshots/explore_repos.png)
![Users](https://dl.gitea.com/screenshots/explore_users.png)
![Orgs](https://dl.gitea.com/screenshots/explore_orgs.png)

### リポジトリ

![Home](https://dl.gitea.com/screenshots/repo_home.png)
![Commits](https://dl.gitea.com/screenshots/repo_commits.png)
![Branches](https://dl.gitea.com/screenshots/repo_branches.png)
![Labels](https://dl.gitea.com/screenshots/repo_labels.png)
![Milestones](https://dl.gitea.com/screenshots/repo_milestones.png)
![Releases](https://dl.gitea.com/screenshots/repo_releases.png)
![Tags](https://dl.gitea.com/screenshots/repo_tags.png)

#### リポジトリの課題

![List](https://dl.gitea.com/screenshots/repo_issues.png)
![Issue](https://dl.gitea.com/screenshots/repo_issue.png)

#### リポジトリのプルリクエスト

![List](https://dl.gitea.com/screenshots/repo_pull_requests.png)
![Pull Request](https://dl.gitea.com/screenshots/repo_pull_request.png)
![File](https://dl.gitea.com/screenshots/repo_pull_request_file.png)
![Commits](https://dl.gitea.com/screenshots/repo_pull_request_commits.png)

#### リポジトリアクション

![List](https://dl.gitea.com/screenshots/repo_actions.png)
![Details](https://dl.gitea.com/screenshots/repo_actions_run.png)

#### リポジトリのアクティビティ

![Activity](https://dl.gitea.com/screenshots/repo_activity.png)
![Contributors](https://dl.gitea.com/screenshots/repo_contributors.png)
![Code Frequency](https://dl.gitea.com/screenshots/repo_code_frequency.png)
![Recent Commits](https://dl.gitea.com/screenshots/repo_recent_commits.png)

### 組織

![Home](https://dl.gitea.com/screenshots/org_home.png)

</details>
