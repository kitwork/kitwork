# KitWork — Automate Workflows Effortlessly

**Automate your workflows with KitWork, a lightweight, fast, and flexible engine for cloud or self-hosted environments**

## What is KitWork?

KitWork is a high-performance platform for developers and teams who want to automate workflows, run serverless functions, build APIs, and manage systems. Everything is defined in simple `.work` files. No heavy frameworks, no complicated setup, just fast, lightweight, and flexible automation

## Key Features

**Effortless Automation**: Run workflows, cron jobs, and serverless functions with minimal setup

**.work-Native**: Define workflows, actions, and microservices in readable `.work` files

**Fast and Lightweight**: Built with Golang for high-performance execution

**Flexible Deployment**: Cloud-ready or self-hosted, scale as needed

**JS and Chromedp Support**: Run dynamic scripts, automate web tasks, scrape data, and interact with web interfaces

**Open Source**: Fully transparent, extendable, and community-friendly

## Example Workflow (`.work`)

```work
name: "get BTC to Telegram"
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
      as: "telegram"
      headers:
        content-type: "application/json"
      body:
        chat_id: "-5077473015"
        text: "KitWork BOT1\n BTC/USDT: {{ btc.price | dollar }}"
      error:
        - log: Failed to send Telegram message

error:
  - log: An error occurred
success:
  - log: BTC/USDT successfully sent to Telegram
```

## Why KitWork?

Turn complex workflows into readable `.work` files. Automate backend processes and web interactions. Self-host and scale without relying on third-party SaaS. Make workflows readable for both humans and automation engines

## Get Started

Clone the repository and run your first workflow

```bash
git clone https://github.com/kitwork/kitwork
cd kitwork
go run .
```

Watch your workflows execute automatically

## About the Author

**Huỳnh Nhân Quốc** ❤️ Dreamy Indie-stack Developer

KitWork is open-source to help anyone automate, build, and scale workflows easily. Support development [Sponsor KitWork](https://github.com/sponsors/huynhnhanquoc)

Released under Apache 2.0 License [LICENSE](https://github.com/kitwork/kitwork/blob/master/LICENSE)

**Automate workflows effortlessly with KitWork, fast, lightweight, and flexible for any environment**

2025.11.29
