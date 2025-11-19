# KitWork – Serverless Engine & YAML-Native Workflows Stack powered by Golang

**KitWork is an ultra-lightweight, high-performance platform built on Golang, designed to help everyone, not just developers, automate workflows, run serverless functions, build APIs, and generate native Golang code, all defined using YAML.**

## Vision

KitWork aims to become the **ultimate meta-engine** for:

* **Automation:** executing workflows, serverless functions, and action scheduling
* **Full-Stack Development:** backend APIs, web frontends, mobile apps, native applications
* **Code Generation:** producing native Golang code from YAML, reducing boilerplate
* **Accessibility:** simple enough for everyone, replacing complex tools like n8n or cumbersome frameworks

**Philosophy:** transform complexity into simplicity using a **single YAML source as the central hub**.

## Initial Goals (MVP)

### Core Engine

* Execute **YAML-defined workflows** and actions
* Support actions: `fetch / http`, `script`, `cmd / command`, `for`, `sendmail`, `save`, `check`, `return`, `chrome / chromedp`
* Cron scheduler for periodic actions
* Manage secrets, proxies, and file watching automatically

### Serverless Runtime

* Run functions triggered by requests or cron schedules
* Provide a simple API endpoint to invoke actions externally

### JS Sandbox Runtime (v8go)

* Safely execute JavaScript scripts within actions
* Enable complex logic while keeping YAML as the central definition

## Expansion Phase

### Full-Stack Workflows

* Backend APIs fully defined in YAML
* Web frontend generated from YAML templates
* Dynamic UI generation from YAML

### No-Code / Low-Code Platform

* Visual GUI editor for workflows, APIs, and actions
* Drag & drop components with standard YAML export

### Advanced Code Generation

* YAML → Golang native code (backend actions, APIs, CLI)
* YAML → HTML/CSS/JS (web apps)
* Extendable for mobile runtime

### Plugin & Extensions

* Add custom action types, runtime engines, or modules
* Connect to third-party services or cloud providers

## Philosophy & Key Differentiators

* **Single Source of Truth:** YAML centralizes workflow, API, web, and native code
* **Simplicity First:** minimalistic design, easy to learn
* **Performance:** powered by Golang for fast actions, serverless functions, and APIs
* **Extensibility:** supports full-stack expansion, native apps, plugin ecosystem, and cloud deployment

## Long-Term Roadmap

| Phase | Goal                     | Description                                           |
| ----- | ------------------------ | ----------------------------------------------------- |
| 1     | Core Engine & Serverless | Actions, cron, JS sandbox, secrets, file watcher      |
| 2     | Full-Stack & API         | Backend API, dynamic web templates, action routing    |
| 3     | No-Code / Low-Code       | GUI workflow editor, drag & drop, YAML export         |
| 4     | Native Code Generation   | YAML → Golang → binary apps, web apps, mobile runtime |
| 5     | Cloud & Scaling          | Serverless deployment, multi-tenant, auto-scaling     |
| 6     | Plugin Ecosystem         | Extensions, 3rd-party integrations, marketplace       |

## Example Workflow (YAML Actions)

```yaml
cron: 
  name: "example api"
  schedules:
    - "0 2 * * *"
    - "0 3 * * *"

  actions:
    - script:
        name: "Fetch gold prices"
        run: ""       # path or inline script
        as: "products"
        timeout: 200
        success:
          - save:
              content: "{{ products }}"
              filename: "hello.json"

    - foreach:
        range: "{{ products }}"
        as: "product"
        async: true
        actions:
          - fetch:
              name: "Fetch product details"
              url: ""  
              method: "GET"
              as: "productDetail"
              timeout: 200

    - fetch:
        name: "Daily gold price"
        url: "https://www.pnj.com.vn/site/gia-vang?r=1709798169672"
        method: "GET"
        as: "price"
        timeout: 200
```

**Highlights:**

* Multiple cron schedules for the same workflow
* Script actions with inline or file-based logic
* Asynchronous `foreach` over items
* Fetch actions with success/error handling, saving, or switching flows
* Chromedp integration possible for dynamic web scraping

## Special with Chromedp Automation

KitWork comes with a **powerful web automation capability** thanks to **Chromedp**, allowing you to control a headless Chrome/Chromium browser directly from YAML. This is an **advanced feature** that enables automation of complex web tasks that standard APIs cannot handle.

### Key Capabilities

* **Navigate & Interact:** open websites, click buttons, fill forms, scroll pages
* **Data Extraction:** extract data from the DOM, innerText, HTML, or JSON from dynamic pages
* **JavaScript Execution:** run custom scripts directly on the page
* **Screenshots & PDF:** capture screenshots or export web pages as PDF
* **Workflow Integration:** seamlessly combine with other actions (`fetch`, `script`, `save`) in cron jobs or async `foreach` loops
* **Web Automation Beyond APIs:** Chromedp allows scraping and interacting with dynamic pages

### Available Chromedp Steps

The steps (`steps`) available in YAML workflows when using Chromedp:

* `wait` – wait for a selector to appear on the page
* `click` – click a specific element
* `fill` – input a value into an input or textarea
* `evaluate` – run JavaScript on the page and store the result
* `screenshot` – capture a screenshot of the whole page or a specific selector
* `navigate` – navigate to a new URL
* `scroll` – scroll the page to a position or selector
* `select` – choose a value in a dropdown or select element
* `extract` – extract an attribute value from an element

### Why It’s Special

* **Web Automation Beyond APIs:** scrape data, perform UI testing, or interact with dynamic pages that APIs do not support
* **Full YAML Integration:** all automation steps are defined in YAML without writing Go code directly
* **Async & Scalable:** combine with cron and `foreach` to handle large-scale web tasks in parallel
* **Reusable & Maintainable:** YAML workflows can be shared, version-controlled, and easily edited without requiring professional developers


## Corner & Diary

* An action is triggered by request, command, schedule, listen, activate, or event (input …)
* An action only returns a result of **success** or **error** (like in Golang) and handles the cases accordingly
* Routers handle only essential HTTP methods: **get**, **post**, **put**, **delete**, and are **guarded**
* Alias path mapping: `{id}` → param (`:id` in GoFiber, path does not support `:`), `{$}` → `*`

```yaml
/router
  /api/  
    /{$}
      get.yaml # API page not found
      post.yaml
      update.yaml  
      delete.yaml

    guard.yaml
    /post 
      /{id}
      get.yaml
      post.yaml
      update.yaml  
      delete.yaml

    /account 
      guard.yaml
      get.yaml
      post.yaml
      update.yaml  
      delete.yaml

  /dashboard
    /{$}
      get.yaml # dashboard page not found

    /post 
        /{id}
        get.yaml # can use check action inside YAML file

    guard.yaml
    get.yaml # cms.html  

  /{$}
      get.yaml # page not found

  get.yaml # index.html
```

* `return` exits the action immediately
* Database operations are simple: select, create, update, delete

```yaml
select: 
  db: "postgres"
  from: "table"
  where: ...
  offset: 0
  limit: 50
```

* Use embedded commands for security (fixed, immutable). Core can also read files dynamically for flexibility
* Scripts default to **JavaScript**, executed by **v8go**
* Variable expansion (`var`) may be supported in the future
* Mobile apps compiled from Golang, supporting native, OpenGL/Vulkan, HTML (including WASM or Reactive integration)
* Fast, self-hosted load balancing / proxy deployment
* YAML is extendable, e.g., table format
* Use **Golang templates** when compiling a future system from YAML to binary, with or without templates
* Build a system that can run **without a database**, using YAML as a data source, optionally inheriting database queries
* Core manages an **entire distributed machine**, with nodes identified by node ID
* Future plans include **UI automation**, similar to n8n drag-and-drop, supporting workflows or code paths via interface
* Game development is not the main focus, but minimal support may exist
* Designed as a **dynamic JAM-stack server architecture**
* Two key features: **shortest path cache** and **full-text search**, implemented iteratively with non-repetitive caching

##  Author

**Huỳnh Nhân Quốc** ❤️ Lập trình viên Indie-stack mộng mơ 

* KitModule: [@kitmodule](https://github.com/kitmodule)
* KitWork: [@kitwork](https://github.com/kitwork)

Released under [MIT License](https://github.com/kitwork/kitwork/blob/master/LICENSE)

**KitWork empowers anyone to automate, develop full-stack applications, and generate native Golang code from a single YAML source, providing a simple, fast, and extensible platform for modern development.**

> I didn’t create a new programming language. I created a new way of programming that allows humans and AI to work together seamlessly.
