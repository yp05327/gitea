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

这行是用于测试的。

[English](./README.md) | [简体中文](./README.zh-cn.md) | [繁體中文](./README.zh-tw.md) | [日本語](./README.ja-JP.md)

## 目的

这个项目的目标是提供最简单、最快速、最省心的方式来搭建自托管的 Git 服务。

由于 Gitea 使用 Go 编写，它可以在 Go 支持的 **所有** 平台和架构上运行，包括 Linux、macOS 与 Windows 的 x86、amd64、ARM 以及 PowerPC 架构。该项目自 2016 年 11 月从 [Gogs](https://gogs.io) [分叉](https://blog.gitea.com/welcome-to-gitea/) 后，已经历了大量更新。

如需体验在线演示，可访问 [demo.gitea.com](https://demo.gitea.com)。

想要使用免费的 Gitea 服务（仓库数量有限制），请访问 [gitea.com](https://gitea.com/user/login)。

要快速部署专属的 Gitea 云实例，可以在 [cloud.gitea.com](https://cloud.gitea.com) 开启免费试用。

## 文档

您可以在我们的官方 [文档网站](https://docs.gitea.com/) 上查阅完整资料。

文档涵盖安装、管理、使用、开发、贡献指南等内容，帮助你快速上手并深入探索全部功能。

如果你有建议或希望参与贡献，可以访问 [文档仓库](https://gitea.com/gitea/docs)。

## 构建

在源码树根目录运行：

    TAGS="bindata" make build

如需 SQLite 支持：

    TAGS="bindata sqlite sqlite_unlock_notify" make build

`build` 目标包含两个子目标：

- `make backend` 需要 [Go Stable](https://go.dev/dl/)，所需版本定义在 [go.mod](/go.mod)。
- `make frontend` 需要 [Node.js LTS](https://nodejs.org/en/download/) 或更高版本以及 [pnpm](https://pnpm.io/installation)。

构建时需要互联网连接以下载 go 和 npm 依赖。从包含预构建前端文件的官方源码包构建时不会执行 `frontend` 目标，因此可以在没有 Node.js 的情况下完成构建。

更多信息：https://docs.gitea.com/installation/install-from-source

## 使用

构建完成后，默认会在源码树根目录生成名为 `gitea` 的二进制文件。运行命令：

    ./gitea web

> [!注意]
> 想要使用我们的 API？我们提供了带有 [文档](https://docs.gitea.com/api) 的实验性支持。

## 贡献

推荐的工作流程是：Fork -> Patch -> Push -> Pull Request

> [!注意]
>
> 1. **开始提交 Pull Request 之前，请务必阅读 [贡献者指南](CONTRIBUTING.md)。**
> 2. 如果你发现项目中的安全漏洞，请发送邮件至 **security@gitea.io**。感谢配合！

## 翻译

[![Crowdin](https://badges.crowdin.net/gitea/localized.svg)](https://translate.gitea.com)

翻译工作通过 [Crowdin](https://translate.gitea.com) 进行。如果你希望添加新的语言，请联系 Crowdin 项目中的管理员启用对应语言。

你也可以创建 issue 请求添加语言，或在 Discord 的 #translation 频道寻求帮助。如果需要上下文或发现翻译问题，可以在字符串下留言或在 Discord 提问。关于翻译的常见问题，文档中有专门章节，目前内容不多，我们期待随着问题积累逐步完善。

更多信息请参阅 [文档](https://docs.gitea.com/contributing/localization)。

## 官方与第三方项目

我们提供官方的 [go-sdk](https://gitea.com/gitea/go-sdk)、名为 [tea](https://gitea.com/gitea/tea) 的 CLI 工具，以及用于 Gitea Action 的 [action runner](https://gitea.com/gitea/act_runner)。

我们在 [gitea/awesome-gitea](https://gitea.com/gitea/awesome-gitea) 维护了一份 Gitea 相关项目清单，你可以在其中发现更多第三方项目，包括 SDK、插件、主题等。

## 交流

[![](https://img.shields.io/discord/322538954119184384.svg?logo=discord&logoColor=white&label=Discord&color=5865F2)](https://discord.gg/Gitea "Join the Discord chat at https://discord.gg/Gitea")

若文档未覆盖你的问题，可以加入我们的 [Discord 服务器](https://discord.gg/Gitea) 联系我们，或在 [Discourse 论坛](https://forum.gitea.com/) 发帖。

## 作者

- [维护者](https://github.com/orgs/go-gitea/people)
- [贡献者](https://github.com/go-gitea/gitea/graphs/contributors)
- [翻译者](options/locale/TRANSLATORS)

## 支持者

感谢所有的支持者！ 🙏 [[成为支持者](https://opencollective.com/gitea#backer)]

<a href="https://opencollective.com/gitea#backers" target="_blank"><img src="https://opencollective.com/gitea/backers.svg?width=890"></a>

## 赞助商

通过成为赞助商支持该项目。你的 Logo 会显示在此处并链接到你的网站。[[成为赞助商](https://opencollective.com/gitea#sponsor)]

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

## 常见问题

**Gitea 怎么读？**

Gitea 的发音是 [/ɡɪ’ti:/](https://youtu.be/EM71-2uDAoY)，就像发 “gi-tea”，其中 g 发硬音。

**为什么这个项目没有托管在 Gitea 实例上？**

我们正在[努力](https://github.com/go-gitea/gitea/issues/1029)。

**在哪里可以找到安全补丁？**

在 [发布日志](https://github.com/go-gitea/gitea/releases) 或 [变更日志](https://github.com/go-gitea/gitea/blob/main/CHANGELOG.md) 中搜索关键字 `SECURITY` 即可找到相关安全补丁。

## 许可证

本项目采用 MIT 许可证。
完整许可文本请参阅 [LICENSE](https://github.com/go-gitea/gitea/blob/main/LICENSE)。

## 更多信息

<details>
<summary>想快速了解界面？来看这里！</summary>

### 登录/注册页面

![Login](https://dl.gitea.com/screenshots/login.png)
![Register](https://dl.gitea.com/screenshots/register.png)

### 用户仪表板

![Home](https://dl.gitea.com/screenshots/home.png)
![Issues](https://dl.gitea.com/screenshots/issues.png)
![Pull Requests](https://dl.gitea.com/screenshots/pull_requests.png)
![Milestones](https://dl.gitea.com/screenshots/milestones.png)

### 用户资料

![Profile](https://dl.gitea.com/screenshots/user_profile.png)

### 探索

![Repos](https://dl.gitea.com/screenshots/explore_repos.png)
![Users](https://dl.gitea.com/screenshots/explore_users.png)
![Orgs](https://dl.gitea.com/screenshots/explore_orgs.png)

### 仓库

![Home](https://dl.gitea.com/screenshots/repo_home.png)
![Commits](https://dl.gitea.com/screenshots/repo_commits.png)
![Branches](https://dl.gitea.com/screenshots/repo_branches.png)
![Labels](https://dl.gitea.com/screenshots/repo_labels.png)
![Milestones](https://dl.gitea.com/screenshots/repo_milestones.png)
![Releases](https://dl.gitea.com/screenshots/repo_releases.png)
![Tags](https://dl.gitea.com/screenshots/repo_tags.png)

#### 仓库问题

![List](https://dl.gitea.com/screenshots/repo_issues.png)
![Issue](https://dl.gitea.com/screenshots/repo_issue.png)

#### 仓库拉取请求

![List](https://dl.gitea.com/screenshots/repo_pull_requests.png)
![Pull Request](https://dl.gitea.com/screenshots/repo_pull_request.png)
![File](https://dl.gitea.com/screenshots/repo_pull_request_file.png)
![Commits](https://dl.gitea.com/screenshots/repo_pull_request_commits.png)

#### 仓库操作

![List](https://dl.gitea.com/screenshots/repo_actions.png)
![Details](https://dl.gitea.com/screenshots/repo_actions_run.png)

#### 仓库活动

![Activity](https://dl.gitea.com/screenshots/repo_activity.png)
![Contributors](https://dl.gitea.com/screenshots/repo_contributors.png)
![Code Frequency](https://dl.gitea.com/screenshots/repo_code_frequency.png)
![Recent Commits](https://dl.gitea.com/screenshots/repo_recent_commits.png)

### 组织

![Home](https://dl.gitea.com/screenshots/org_home.png)

</details>
