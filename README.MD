# **KitWork — The Lightweight Work~~flow~~ Engine for your Services**

[![Go](https://img.shields.io/badge/go-1.22-00ADD8?logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/license-Apache_2.0-blue.svg)](LICENSE)

KitWork is a lightweight, high-performance workflow engine that helps developers automate, orchestrate, and connect services with clean, predictable workflows.

It works perfectly for cron jobs, backend automation, browser automation, scraping, serverless scripting, and microservice orchestration — whether on the cloud or fully self-hosted.

> **n8n + cron + Puppeteer + Airflow — but self-hosted, <20MB, no Docker, no Node.js.**

KitWork treats automation as a structured system: easy to define, effortless to maintain, and powerful enough to scale.


## 🚀 Get Started

### **1. Run from source**

```bash
git clone https://github.com/kitwork/kitwork.git
cd kitwork
go run .
```

### **2. Or download a prebuilt binary**

#### Linux x64

```bash
curl -L https://github.com/kitwork/kitwork/releases/latest/download/kitwork-linux-amd64 -o kitwork
chmod +x kitwork
./kitwork
```

#### Windows

Download `kitwork-win.exe` from the Releases page and run:

```bash
./kitwork-win.exe
```

#### macOS (Intel)

```bash
curl -L https://github.com/kitwork/kitwork/releases/latest/download/kitwork-darwin-amd64 -o kitwork
chmod +x kitwork
./kitwork
```


### **3. Create your first workflow**

Create a file named `hello.work`:

```yaml
name: "Hello World"

schedules:
  - every: 5s

work:
  - log: "Hello from KitWork!"
```

Start KitWork — it automatically discovers and runs `.work` files.


## 🔧 Key Features

### ⚡ Lightweight & Fast

* Single binary (~12–18 MB depending on platform)
* No dependencies
* Drop → run
* Built with Go: fast, portable, reliable

Perfect for servers, VPS, containers, or tiny self-hosted setups.


### 📄 `.work` — a human-first workflow format

KitWork uses a YAML-like declarative language:

* readable
* simple
* zero boilerplate
* easy to debug


### 🌀 Flexible Orchestration

Supports:

* Cron schedules
* Webhooks
* API calls
* Conditional branches
* Chained workflows
* Retry logic
* Error/success handlers
* Secrets system


### 🧩 Browser Automation & Scripts

**First-class JavaScript & browser automation**:

* **V8** — secure JS runtime
* **Chromedp** — headless browser

Perfect for:

* scraping
* bot flows
* autofill
* browser-based checks
* automation pipelines


### 🔗 HTTP, Fetching, JSON & API Tasks

Easily call APIs, transform JSON, pass data between workflow steps.


### 🗄 SQL & Data Interaction (PostgreSQL, MySQL, SQLite)

Query or modify your database directly inside `.work` files.

Example:

```yaml
- sql:
    as: users
    select:
      - "email"
      - "username"
    from: user
    where:
      is_active: true
    limit: 30
```


## 🧑‍💻 What You Can Build With KitWork

* Automated reports, emails, reminders
* Multi-step backend workflows
* Scraping & data pipelines
* Browser bots (Chromedp)
* Serverless-style functions (self-hosted)
* Structured cron task replacements
* Microservice orchestration
* Internal automation for teams
* Mini API endpoints
* Self-hosted automation platforms

Perfect for indie hackers, solo devs, SaaS builders, and automation-heavy teams.


## 🧪 Example: Fetch BTC/USDT & Send to Telegram

```yaml
name: "Get BTC to Telegram"

schedules:
  - every: 5s

work:
  - fetch:
      name: "Fetch BTC Price"
      url: "https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT"
      method: "GET"
      as: btc
      timeout: 1000
      error:
        - log: Failed to fetch BTC

  - log: BTC/USDT is currently {{ btc.price | dollar }}

  - fetch:
      name: "Send to Telegram"
      url: "https://api.telegram.org/bot{{ secret.telegram.key }}/sendMessage"
      method: "POST"
      headers:
        content-type: "application/json"
      body:
        chat_id: "-5077473015"
        text: "KitWork BOT\nBTC/USDT: {{ btc.price | dollar }}"
      error:
        - log: Failed to send Telegram message

error:
  - log: An error occurred

success:
  - log: BTC/USDT successfully sent to Telegram
```


## 🏗 Architecture Overview

KitWork is powered by:

* **Golang** — performance & simplicity
* **Fiber** (optional) — fast HTTP API
* **V8go** — isolated JavaScript runtime
* **Chromedp** — browser automation
* **Declarative `.work` format** — structured and predictable

Works as:

* Workflow engine
* Task runner
* Script executor
* Automation hub


## 🎯 Why KitWork Exists

Existing automation platforms are often:

* heavy
* cloud-locked
* expensive
* complicated
* hard to self-host

KitWork provides a **lightweight, self-hosted, developer-first alternative**.

Small scripts → structured workflows → production-ready automation.


## 🧩 Who Uses KitWork?

* Indie developers launching SaaS
* Teams building internal automation
* DevOps replacing messy cron jobs
* Browser automation creators
* Data engineering enthusiasts
* Self-hosting users
* Developers who love simplicity


## 📜 License

Released under **Apache 2.0 License**.
Free for personal, commercial, and enterprise use.


## ❤️ Made for Developers

> **“Automation should be powerful, predictable, and fun to build.”**


## 🧡 Author

**[Huỳnh Nhân Quốc](https://github.com/huynhnhanquoc)— Dreamy Indie-stack Developer**
KitWork is open-source to help anyone automate, build, and scale workflows effortlessly.

Support development → **[Sponsor KitWork](https://github.com/sponsors/huynhnhanquoc)**


## 🌟 Tagline

**Automate workflows effortlessly with KitWork — fast, lightweight, flexible everywhere.**

**2025.12.10**

