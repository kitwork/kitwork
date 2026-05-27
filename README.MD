# 🚀 Kitwork — The Industrial-Grade Logic Operating System

> **"Logic is the soul of machines. Emotion creates civilization."**

**An ultra-lightweight, sovereign logic execution engine powered by Go. Run APIs, cron schedules, templates, and background tasks on a custom stack-based bytecode virtual machine with nanosecond precision.**

---

## ⚡ Performance Benchmarks

Kitwork is engineered for **Nanosecond-Precision Sovereign Execution**. We bypass traditional runtime reflection, utilizing zero-copy mapping and intensive context pooling to run scripts with bare-metal efficiency.

The following load-test benchmarks were executed concurrently against a single local Kitwork node using `k6`:

| Metric / Scenario | Value / Result | Architectural Benefit |
| :--- | :--- | :--- |
| **🚀 VM Core Clock** | **~14,100,000 ops/s** | Raw Bytecode Execution Speed |
| **⏱️ Logic Latency** | **70ns** | Pure Stack-Based Precision |
| **📈 Concurrent Throughput** | **12,726.19 requests/sec (RPS)** | Native Go Trie Router & Thread-Safe Engine |
| **⚡ Response Latency (p50)** | **1.16 ms** (median) / **90 µs** (avg JSON) | Zero-Allocation Response Builder |
| **✅ Success Rate (HTTP 200)** | **100.00%** (0 / 127,292 failed) | Highly Robust Connection Pooling |
| **⚙️ Cold Boot Time** | **< 10ms** | Instant-scale Serverless Bootstrap |
| **🔋 GC Memory Pressure** | **Near Zero** | Aggressive Context & VM stack pooling |

---

## 💎 Key Features

*   **Custom Stack-Based VM**: A dedicated bytecode engine optimized for zero-GC pressure and constant-time variable access.
*   **Zero-Copy Routing Trie**: Matches routes natively in Go, bypassing Javascript overhead entirely for path resolution.
*   **Thread-Safe Static Caching**: Uses a custom lock manager (`sync.RWMutex` map per file path) to serve pre-compressed cached responses dynamically without disk conflicts.
*   **Fluent Zero-Allocation ORM**: Database connection builder executing queries in nanoseconds—nearly 20x faster than heavyweight ORMs.
*   **Compliant VietQR Generator**: Built-in EMVCo compliance check and bank mapping for generating custom SVGs dynamically.
*   **Headless Web Capture**: Spawn headless chromium browser contexts on-the-fly (`chromedp` integration) to snap screenshots and crawl dynamic DOM content.
*   **Parallel Concurrency**: Non-blocking asynchronous threads running parallel tasks via Go's lightweight scheduler.

---

## 🏁 Standalone Quick Start

Get your sovereign engine running locally in under 60 seconds (Standalone Mode without Tenants):

### 1. Install & Run
Ensure you have Go installed on your system:
```bash
git clone https://github.com/kitwork/kitwork.git
cd kitwork
go run .
```

### 2. Standard Folder Structure
Define your routes and system logic directly at the root level of your directory:
```text
kitwork/
├── config.kitwork.yml      # Global configuration (port, DB config)
├── app.kitwork.js          # Main application script (VM router and logic)
├── main.go                 # Engine entry bootstrap
└── views/                  # HTML templates and views
```

### 3. Global Config (`config.kitwork.yml`)
Create a configuration file at the root:
```yaml
port: 8085
root: "."
hot_reload: true
database:
  type: "postgres"
  user: "postgres"
  password: "your_password"
  name: "postgres"
  host: "localhost"
  port: 5432
  max_open: 50
  max_idle: 10
```

### 4. Write Main Logic (`app.kitwork.js`)
Create your main entry script at the root:
```javascript
import { router, database, log } from "kitwork"

const db = database.connection()

// Simple JSON API
router.get("/api/hello").handle((req, res) => {
    return res.json({
        status: "active",
        engine: "Kitwork Sovereign VM",
        timestamp: Date().toISOString()
    })
})

// Database Query API
router.get("/api/users").handle((req, res) => {
    const list = db.table("user").list(10)
    return res.json({ success: true, users: list })
})
```

---

## 📖 API Handbook & Snippets

### 1. Zero-Copy Routing & Cache
```javascript
// Static file endpoint
router.get("/favicon.ico").file("/assets/favicon.ico");

// Serving asset directory
router.get("/assets/*").directory("./assets/*");

// In-Memory cache (thread-safe sync.RWMutex)
router.get("/api/gold").cache("5s").handle((req, res) => {
    return res.json({ price: 2300 })
});

// Disk-backed static router cache
router.get("/api/cached").static("10s").handle((req, res) => {
    return res.text("Static Cached at: " + Date().toISOString())
});
```

### 2. Fluent Database ORM
```javascript
const db = database.connection()

// Find record by primary key (default: username="grace")
const user = db.table("user").find("username", "grace")

// Fetch first record in table
const firstUser = db.table("user").first()

// Fetch record list with limit
const users = db.table("user").list(5)

// Aggregate count queries
const totalUsers = db.table("user").count()

// Update records
db.table("user").where("id", 1).update({ status: "inactive" })
```

### 3. Compliant VietQR dynamic SVG
```javascript
import { napas, qrcode } from "kitwork"

router.get("/qrcode/napas").handle((req, res) => {
    const payment = napas
        .bank("vcb", "1234567890") // bank code, account
        .amount("50000")
        .info("Payment Description")

    const svgContent = qrcode
        .napas(payment)
        .template("circular") // QR cell design template
        .padding(2)
        .cell({ color: "#0f172a", size: 0.75 }) // custom colors & contrasts
        .svg()

    return res.svg(svgContent)
})
```

### 4. Headless Web Capture
```javascript
import { chromedp } from "kitwork"

router.get("/screenshot").handle((req, res) => {
    const urlStr = req.query("url").text() || "https://github.com"
    const pngBytes = chromedp.capture(urlStr, {
        width: 1280,
        height: 720
    })
    return res.image(pngBytes)
})
```

### 5. Goroutine Concurrency Worker
```javascript
import { go, log } from "kitwork"

router.get("/background").handle((req, res) => {
    log.Print("Starting worker...")
    
    // Spawn non-blocking thread
    go(() => {
        log.Print("Worker running parallel...")
        // run async actions here
        log.Print("Worker completed successfully!")
    })

    return res.text("Background task started!")
})
```

---

## ✒️ Author's Note

> *"While the world is busy using AI to automate everything, I choose to breathe a soul into every line of code. I write code like essays, like unspoken confessions. I use AI not to replace myself, but as a mirror to reflect my own inner world. I expose this system to the world simply because it is beautiful, crazy, and dreamy."*
> 
> — **Huỳnh Nhân Quốc**, *Indie-Stack Developer* ❤️

---

Released under Apache 2.0 License.
Support development: [Sponsor Kitwork](https://github.com/sponsors/huynhnhanquoc)
