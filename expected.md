# KitWork – A single-YAML engine that changes how the world writes software

**KitWork is an ultra-lightweight Golang engine that lets you define and run entire systems — workflows, serverless functions, REST/GraphQL APIs, microservices, and real-browser web data collection — using just one or a few YAML files.**

Everything runs as a single tiny Go binary.  
No Node.js. No Docker layers. No runtime dependencies. Just pure speed.

<img src="https://kitwork.io/hero-demo.gif" alt="KitWork demo" />

### What you can build with a single YAML file

| Goal                                                  | Lines of YAML | Files needed |
|-------------------------------------------------------|---------------|--------------|
| Daily PNJ gold price → Telegram                       | ~25           | 1            |
| Bitcoin price API with 60 s cache                     | ~18           | 1            |
| Shopee price monitor → Zalo alert when < $400         | ~35           | 1            |
| Full backend + admin dashboard                        | < 100 + templates | 1 folder |
| Deploy the entire system                              | One binary    | Zero runtime |

### 6 killer features no tool combines today

| Feature                            | How KitWork does it                                            | Anyone else? |
|------------------------------------|----------------------------------------------------------------|--------------|
| True YAML-first no/low-code        | All logic in YAML — non-devs can build real apps               | Not yet      |
| Full Go web framework              | Routes, auth, DB, templates — auto-generated from folders+YAML| No           |
| Native serverless runtime          | Every YAML file runs on cron or HTTP trigger                   | Vercel isn’t this light |
| Real Chrome automation built-in    | Chromedp 100 % controlled by YAML (Shopee, Tiki, banks…)       | Extra nodes required elsewhere |
| Single-binary deployment           | One executable runs everywhere (Windows / Linux / macOS)       | Nobody yet   |
| File system = instant router       | `/api/user/get.yaml` → GET /api/user instantly                 | Closest: Next.js App Router |

### Project structure = your entire app

```text
my-app/
├── kitwork.yaml          # global config, secrets, proxy
├── cron/
│   └── daily-gold.yaml   # runs every morning
├── api/
│   ├── price.yaml        → GET  /api/price
│   ├── login.yaml        → POST /api/login + guard
│   └── {$}.yaml          → custom 404
├── dashboard/
│   └── index.yaml        → HTML + Go templates
└── data/                 # auto-saved results
```

One command:
```bash
kitwork run
```
→ API + cron + dashboard + browser automation — all in one tiny process.

### Real example: Daily PNJ gold price collector (~25 lines)

```yaml
# cron/daily-gold.yaml
name: Daily PNJ gold price
schedules:
  - "0 8 * * *"

actions:
  - chrome:
      headless: true
      steps:
        - navigate: "https://pnj.com.vn/shop/gia-vang/"
        - wait: "table.gia-vang"
        - evaluate: |
            return Array.from(document.querySelectorAll('table.gia-vang tr'))
              .slice(1)
              .map(r => {
                const c = r.querySelectorAll('td');
                return {type: c[0].innerText.trim(), buy: c[1].innerText.trim(), sell: c[2].innerText.trim()};
              });
          as: prices

  - save:
      content: "{{ prices | json }}"
      filename: "data/gold-{{now.Format \"20060102\"}}.json"

  - fetch:
      url: "https://api.telegram.org/botXXX/sendMessage"
      method: POST
      json:
        chat_id: "YYY"
        text: "PNJ Gold today:\n{{prices | prettify}}"
```

### Quick comparison (November 2025)

| Feature                          | n8n / Make   | Windmill    | Temporal   | KitWork       |
|----------------------------------|--------------|-------------|------------|---------------|
| Core language                    | Node.js      | Go + DB     | Go/Java    | 100% Golang   |
| Workflow definition              | GUI          | YAML+code   | SDK        | YAML only     |
| Runs on 512 MB VPS               | Hard         | OK          | OK         | Extremely light |
| Real Chrome automation           | Extra nodes  | Code        | Code       | Native in YAML |
| Single binary deployment         | No           | No          | No         | Yes           |
| Zero-database mode               | Not recommended | Possible | No      | Fully supported |
| File = route                     | No           | No          | No         | Yes           |

### Why KitWork exists

To let anyone:
- Automate anything with pure YAML  
- Build production backends as fast as landing pages  
- Collect data from the hardest websites without pain  
- Self-host forever — zero SaaS lock-in  
- Deploy with one tiny binary that just works

**Website:** https://kitwork.io  
**GitHub:** https://github.com/kitwork/kitwork  
**License:** MIT

Alpha goes public in the next few days.

Want to be among the very first humans to run it?  
Drop a 🔥 below — I’ll DM you the alpha binary tonight.

KitWork – turning the most complex things into the simplest ones.
