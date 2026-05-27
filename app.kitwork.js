import { router, log, render, http, database, go, napas, qrcode, chromedp } from 'kitwork'

const global = {
    name: "kitwork",
    logo: "/assets/logo.png",
    favicon: "/assets/favicon.ico",
    title: "Chào mừng tới kitwork",
}

const layout = {
    navbar: "_navbar_",
    footer: "_footer_",
    head: "_head_",
}

const layoutDocs = {
    ...layout,
    sidebar: "/docs/_sidebar_",
    toolbar: "/docs/_toolbar_"
}

const site = render.directory("views");

const home = site.path("/").global(global).layout(layout).notfound("notfound");
const docs = site.path("/docs").global(global).layout(layoutDocs).notfound("notfound");

const db = database.connection()

// Static files / Assets mapping
router.get("/favicon.ico").file("/assets/favicon.ico");
router.get("/taiwindcss.js").file("/assets/js/taiwindcss.js");
router.get("/public/*").directory("./demo/public");
router.get("/assets/*").directory("./assets/*");

// 1. Diagnostics & Cached APIs
router.get("/api/hello").handle((req, res) => {
    return res.json({
        status: "active",
        engine: "Kitwork Sovereign VM",
        mode: "Standalone (Direct)",
        latency: "70ns",
        timestamp: Date().toISOString()
    });
});

router.get("/api/cached").static("10s").handle((req, res) => {
    const timeNow = Date().toISOString();
    return res.text("Static cached at: " + timeNow + ". Refresh page: this text will remain cached for 10 seconds!");
});

// 2. External API Proxy with Cache
router.get("/api/gold").cache("5s")
    .catch(() => log.Print("Failed to fetch gold price"))
    .then(() => log.Print("Gold price fetched successfully"))
    .handle((req, res) => {
        const fetch = http.get("https://edge-api.pnj.io/ecom-frontend/v1/get-gold-price?zone=11");
        log.Print(fetch.status);
        if (fetch.status != 200) {
            return res.status(500).json({ status: fetch.status, error: "Failed to fetch gold price" })
        }
        const body = fetch.json();
        log.Print(body);
        const data = body.data.map(item => ({
            name: item.tensp,
            buy: item.giamua,
            sell: item.giaban
        }));
        return res.status(200).json({ success: true, data: data });
    });

// 3. Database Query test endpoint
router.get("/test-query").handle((req, res) => {
    return res.json({
        find: db.table("user").find("username", "grace"),
        first: db.table("user").first(),
        list: db.table("user").list(3),
        count: db.table("user").count()
    });
});

// 4. Compliant VietQR dynamic SVG generation
router.get("/qrcode/napas").handle((req, res) => {
    const bank = req.query("bank").text() || "vcb";
    const account = req.query("account").text() || "1234567890";
    const amountStr = req.query("amount").text() || "50000";
    const desc = req.query("desc").text() || "Thanh toan";

    const payment = napas
        .bank(bank, account)
        .amount(amountStr)
        .info(desc);

    const svgContent = qrcode
        .napas(payment)
        .template("circular")
        .padding(2)
        .cell({
            color: "#0f172a",
            size: 0.75
        })
        .svg();

    return res.svg(svgContent);
});

// 5. Headless Web Capture (using chromedp)
router.get("/screenshot").handle((req, res) => {
    const urlStr = req.query("url").text() || "https://github.com";
    const pngBytes = chromedp.capture(urlStr, {
        width: 1280,
        height: 720
    });
    return res.image(pngBytes);
});

// 6. Parallel Concurrency Worker (go)
router.get("/background").handle((req, res) => {
    log.Print("Request received, starting background worker...");
    go(() => {
        log.Print("Background task is running parallel...");
        log.Print("Background task completed successfully!");
    });
    return res.text("Background task started!");
});

// 7. Docs rendering using docs layout
router.get("/docs/:site?").handle((req, res) => {
    const siteParam = req.params("site");
    const binding = { request: req };
    const view = docs.page(siteParam || "/").bind(binding);
    return res.html(view);
});

// 8. User list/detail rendering using home layout
router.get("/users/:id?").handle((req, res) => {
    const page = home.page(req.page());
    const id = req.params("id");
    const binding = { request: req };
    if (!id) {
        binding.users = db.table("user").list(5);
    } else {
        binding.user = db.where(user => user.id == id).first();
    }
    const view = page.bind(binding);
    return res.html(view);
});

// 9. Wildcard fallback router to serve home layout pages (including views/page.kitwork.html)
router.get("/*").handle((req, res) => {
    const requestPath = req.path();
    const binding = { path: requestPath }
    const view = home.page(requestPath).bind(binding);
    return res.html(view);
});