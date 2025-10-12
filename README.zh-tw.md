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

這行用於測試。

[English](./README.md) | [简体中文](./README.zh-cn.md) | [繁體中文](./README.zh-tw.md) | [日本語](./README.ja-JP.md)

## 目的

這個專案的目標是提供最簡單、最快速、最省心的方式，讓你建置自託管的 Git 服務。

由於 Gitea 是以 Go 撰寫，它能在 Go 支援的 **所有** 平台與架構上運行，包括 Linux、macOS 與 Windows 的 x86、amd64、ARM 以及 PowerPC 等架構。自 2016 年 11 月自 [Gogs](https://gogs.io) [分支](https://blog.gitea.com/welcome-to-gitea/) 以來，專案已經歷許多變革。

若想線上體驗，可造訪 [demo.gitea.com](https://demo.gitea.com)。

想使用免費的 Gitea 服務（倉庫數量有限），請前往 [gitea.com](https://gitea.com/user/login)。

若想快速部署專屬的 Gitea 雲端實例，可在 [cloud.gitea.com](https://cloud.gitea.com) 開始免費試用。

## 文件

你可以在官方的 [文件網站](https://docs.gitea.com/) 上找到完整的資料。

文件涵蓋安裝、管理、使用、開發、貢獻指南等內容，協助你快速上手並深入探索所有功能。

若你有建議或想直接貢獻，歡迎造訪 [文件倉庫](https://gitea.com/gitea/docs)。

## 建置

請在原始碼樹的根目錄執行：

    TAGS="bindata" make build

若需要 SQLite 支援：

    TAGS="bindata sqlite sqlite_unlock_notify" make build

`build` 目標包含兩個子目標：

- `make backend` 需要 [Go Stable](https://go.dev/dl/)，所需版本記載於 [go.mod](/go.mod)。
- `make frontend` 需要 [Node.js LTS](https://nodejs.org/en/download/) 或更新版本，以及 [pnpm](https://pnpm.io/installation)。

建置時需要網際網路連線以下載 go 與 npm 相依套件。若使用附帶預先建置前端檔案的官方原始碼壓縮包，`frontend` 目標不會被觸發，因此無須安裝 Node.js。

更多資訊：https://docs.gitea.com/installation/install-from-source

## 使用

建置完成後，預設會在原始碼樹根目錄產生名為 `gitea` 的可執行檔。請使用下列指令啟動：

    ./gitea web

> [!注意]
> 想使用我們的 API 嗎？我們提供了附有 [文件](https://docs.gitea.com/api) 的實驗性支援。

## 貢獻

建議的工作流程：Fork -> Patch -> Push -> Pull Request

> [!注意]
>
> 1. **在開始提交 Pull Request 之前，務必閱讀 [貢獻者指南](CONTRIBUTING.md)。**
> 2. 若你發現專案中的安全性問題，請來信 **security@gitea.io**。感謝協助！

## 翻譯

[![Crowdin](https://badges.crowdin.net/gitea/localized.svg)](https://translate.gitea.com)

翻譯工作透過 [Crowdin](https://translate.gitea.com) 進行。若想加入新的語言，請聯絡 Crowdin 專案的管理員開啟該語言。

你也可以建立 issue 要求新增語言，或到 Discord 的 #translation 頻道尋求協助。若需要上下文或發現翻譯問題，可以在字串下留言或於 Discord 提問。文件中也有一節介紹翻譯流程，目前內容仍在補充中，歡迎提供問題讓我們持續完善。

更多資訊請參閱 [文件](https://docs.gitea.com/contributing/localization)。

## 官方與第三方專案

我們提供官方的 [go-sdk](https://gitea.com/gitea/go-sdk)、名為 [tea](https://gitea.com/gitea/tea) 的 CLI 工具，以及 Gitea Action 的 [action runner](https://gitea.com/gitea/act_runner)。

我們在 [gitea/awesome-gitea](https://gitea.com/gitea/awesome-gitea) 維護 Gitea 相關專案清單，你可以在那裡找到更多第三方專案，例如 SDK、外掛、佈景主題等。

## 交流

[![](https://img.shields.io/discord/322538954119184384.svg?logo=discord&logoColor=white&label=Discord&color=5865F2)](https://discord.gg/Gitea "Join the Discord chat at https://discord.gg/Gitea")

如果文件沒有回答你的問題，歡迎加入我們的 [Discord 伺服器](https://discord.gg/Gitea)，或在 [Discourse 論壇](https://forum.gitea.com/) 發文討論。

## 作者

- [維護者](https://github.com/orgs/go-gitea/people)
- [貢獻者](https://github.com/go-gitea/gitea/graphs/contributors)
- [翻譯者](options/locale/TRANSLATORS)

## 支持者

感謝所有支持者！🙏 [[成為支持者](https://opencollective.com/gitea#backer)]

<a href="https://opencollective.com/gitea#backers" target="_blank"><img src="https://opencollective.com/gitea/backers.svg?width=890"></a>

## 贊助商

歡迎成為贊助商支持這個專案。你的商標會顯示在此並連結到你的網站。[[成為贊助商](https://opencollective.com/gitea#sponsor)]

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

## 常見問題

**Gitea 怎麼讀？**

Gitea 的發音是 [/ɡɪ’ti:/](https://youtu.be/EM71-2uDAoY)，就像「gi-tea」，其中 g 發硬音。

**為什麼這個專案沒有托管在 Gitea 實例上？**

我們正在[努力](https://github.com/go-gitea/gitea/issues/1029)。

**在哪裡可以找到安全性修補？**

請在 [發佈日誌](https://github.com/go-gitea/gitea/releases) 或 [變更日誌](https://github.com/go-gitea/gitea/blob/main/CHANGELOG.md) 中搜尋關鍵字 `SECURITY`。

## 授權條款

本專案採用 MIT 授權條款。
完整內容請參閱 [LICENSE](https://github.com/go-gitea/gitea/blob/main/LICENSE)。

## 更多資訊

<details>
<summary>想快速了解介面嗎？請往下看！</summary>

### 登入／註冊頁面

![Login](https://dl.gitea.com/screenshots/login.png)
![Register](https://dl.gitea.com/screenshots/register.png)

### 使用者儀表板

![Home](https://dl.gitea.com/screenshots/home.png)
![Issues](https://dl.gitea.com/screenshots/issues.png)
![Pull Requests](https://dl.gitea.com/screenshots/pull_requests.png)
![Milestones](https://dl.gitea.com/screenshots/milestones.png)

### 使用者檔案

![Profile](https://dl.gitea.com/screenshots/user_profile.png)

### 探索

![Repos](https://dl.gitea.com/screenshots/explore_repos.png)
![Users](https://dl.gitea.com/screenshots/explore_users.png)
![Orgs](https://dl.gitea.com/screenshots/explore_orgs.png)

### 倉庫

![Home](https://dl.gitea.com/screenshots/repo_home.png)
![Commits](https://dl.gitea.com/screenshots/repo_commits.png)
![Branches](https://dl.gitea.com/screenshots/repo_branches.png)
![Labels](https://dl.gitea.com/screenshots/repo_labels.png)
![Milestones](https://dl.gitea.com/screenshots/repo_milestones.png)
![Releases](https://dl.gitea.com/screenshots/repo_releases.png)
![Tags](https://dl.gitea.com/screenshots/repo_tags.png)

#### 倉庫問題

![List](https://dl.gitea.com/screenshots/repo_issues.png)
![Issue](https://dl.gitea.com/screenshots/repo_issue.png)

#### 倉庫拉取請求

![List](https://dl.gitea.com/screenshots/repo_pull_requests.png)
![Pull Request](https://dl.gitea.com/screenshots/repo_pull_request.png)
![File](https://dl.gitea.com/screenshots/repo_pull_request_file.png)
![Commits](https://dl.gitea.com/screenshots/repo_pull_request_commits.png)

#### 倉庫動作

![List](https://dl.gitea.com/screenshots/repo_actions.png)
![Details](https://dl.gitea.com/screenshots/repo_actions_run.png)

#### 倉庫動態

![Activity](https://dl.gitea.com/screenshots/repo_activity.png)
![Contributors](https://dl.gitea.com/screenshots/repo_contributors.png)
![Code Frequency](https://dl.gitea.com/screenshots/repo_code_frequency.png)
![Recent Commits](https://dl.gitea.com/screenshots/repo_recent_commits.png)

### 組織

![Home](https://dl.gitea.com/screenshots/org_home.png)

</details>
