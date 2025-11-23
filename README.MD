# KitWork — Automate Kit Workflows Effortlessly

**Automate kit workflows effortlessly with a lightweight, fast, and flexible engine ready for cloud or self-hosted setups.**


## What is KitWork?

KitWork is a high-performance platform built for developers and teams who want to automate workflows, run serverless functions, build APIs, and manage systems. Everything is defined in simple YAML files. No heavy frameworks, no complicated setup, just lightweight, fast, and flexible automation.


## Key Features

**Effortless Automation:** Run workflows, cron jobs, and serverless functions with minimal setup.

**YAML-Native:** Define workflows, actions, and microservices in simple YAML files.

**Fast and Lightweight:** Powered by Golang for high-performance execution.

**Flexible Deployment:** Cloud-ready or fully self-hosted, scale as needed.

**JS and Chromedp Support:** Execute dynamic scripts safely, automate web tasks, scrape data, and interact with web interfaces.

**Open Source:** Fully transparent, extendable, and ready for community.


## Example Workflow (YAML)

```yaml
cron:  # or empty
   name: "example api"
   schedules:
    - daily: 15:00


   work:
   - fetch:
         name: "Lấy giá BTC"
         url: "https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT"
         method: "GET"
         as: btc
         timeout: 1000 
         success: 
          - fetch:
              name: "Gửi đến Telegram"
              url: "https://api.telegram.org/bot{{ .Key.Telegram }}/sendMessage"
              method: "POST"
              as: "telegram"
              headers: 
                 content-type: "application/json"
              body: 
                chat_id: "-5077473015"
                text: "Kitwork send Telegram {{ btc.price }}"
              success: 
              error: 

         error: 

   success: 
     - log: "Gửi thành công {{ btc | json }}"
   error: 
      - log: "Đã có lỗi xảy ra"
```


## Why KitWork?

Turn complex workflows into readable, maintainable YAML. Automate both backend processes and web interactions. Self-host and scale without relying on third-party SaaS. Bridge human and machine understanding by making workflows readable by both humans and automation engines.


## Get Started

Clone the repository and run your first workflow:

```bash
git clone https://github.com/kitwork/kitwork
cd kitwork
go run .
```
Then watch your workflows execute automatically.

## About the Author

**Huỳnh Nhân Quốc** ❤️ Dreamy Indie-stack Developer

KitWork is open-source to empower anyone to automate, build, and scale workflows effortlessly. Support development [Sponsor KitWork](https://github.com/sponsors/huynhnhanquoc)

Released under Apache 2.0 License – [LICENSE](https://github.com/kitwork/kitwork/blob/master/LICENSE)

**Automate kit workflows effortlessly with KitWork, fast, lightweight, and flexible for any environment.**

2025.11.23