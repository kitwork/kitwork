# KitWork – Engine Serverless & Ngăn xếp Workflows YAML-Native với sức mạnh Golang

**KitWork là một nền tảng cực nhẹ, hiệu năng cao, xây dựng trên Golang, giúp mọi người, không chỉ các lập trình viên, tự động hóa workflow, chạy các hàm serverless, xây dựng API và sinh code Golang gốc, tất cả được định nghĩa bằng YAML.**


## Tầm nhìn

KitWork hướng đến trở thành **meta-engine tối ưu** cho:

* **Tự động hóa:** thực thi workflow, hàm serverless, lập lịch action
* **Phát triển Full-Stack:** backend API, frontend web, ứng dụng mobile, ứng dụng native
* **Sinh code:** sinh code Golang gốc từ YAML, giảm boilerplate
* **Dễ tiếp cận:** đơn giản cho mọi người, thay thế các công cụ phức tạp như n8n hay framework cồng kềnh

**Triết lý:** biến phức tạp thành đơn giản bằng **một nguồn YAML duy nhất làm trung tâm**.


## Mục tiêu ban đầu (MVP)

### Core Engine

* Thực thi **workflow và action định nghĩa bằng YAML**
* Hỗ trợ các action: `fetch / http`, `script`, `cmd / command`, `for`, `sendmail`, `save`, `check`, `return`, `chrome / chromedp`
* Cron scheduler cho các action định kỳ
* Quản lý secrets, proxy, và file watcher tự động

### Serverless Runtime

* Chạy hàm khi có request hoặc theo cron
* Cung cấp endpoint API đơn giản để gọi action từ bên ngoài

### JS Sandbox Runtime (v8go)

* Thực thi an toàn các script JavaScript trong action
* Cho phép logic phức tạp vẫn giữ YAML làm định nghĩa trung tâm


## Giai đoạn mở rộng

### Full-Stack Workflows

* Backend API hoàn toàn định nghĩa bằng YAML
* Frontend web sinh tự động từ template YAML
* UI động sinh ra từ YAML

### Nền tảng No-Code / Low-Code

* GUI trực quan cho workflow, API và action
* Drag & drop component, xuất chuẩn YAML

### Sinh code nâng cao

* YAML → Golang native code (backend, API, CLI)
* YAML → HTML/CSS/JS (web app)
* Mở rộng cho runtime mobile

### Plugin & Extension

* Thêm loại action, runtime engine, hoặc module tùy chỉnh
* Kết nối với dịch vụ bên thứ ba hoặc cloud provider


## Triết lý & Điểm khác biệt

* **Single Source of Truth:** YAML tập trung workflow, API, web và code native
* **Đơn giản trước tiên:** thiết kế tối giản, dễ học
* **Hiệu năng:** chạy nhanh nhờ Golang, hỗ trợ serverless, async
* **Mở rộng:** hỗ trợ full-stack, ứng dụng native, hệ sinh thái plugin, deploy cloud


## Lộ trình dài hạn

| Giai đoạn | Mục tiêu                 | Mô tả                                                 |
| --------- | ------------------------ | ----------------------------------------------------- |
| 1         | Core Engine & Serverless | Action, cron, JS sandbox, secrets, file watcher       |
| 2         | Full-Stack & API         | Backend API, template web động, routing action        |
| 3         | No-Code / Low-Code       | GUI workflow editor, drag & drop, YAML export         |
| 4         | Native Code Generation   | YAML → Golang → binary apps, web apps, runtime mobile |
| 5         | Cloud & Scaling          | Deploy serverless, multi-tenant, auto-scaling         |
| 6         | Plugin Ecosystem         | Extension, tích hợp bên thứ 3, marketplace            |


## Ví dụ Workflow (YAML Actions)

```yaml
cron: 
  name: "example api"
  schedules:
    - "0 2 * * *"
    - "0 3 * * *"

  actions:
    - script:
        name: "Fetch giá vàng"
        run: ""       # path hoặc inline script
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
              name: "Chi tiết sản phẩm"
              url: ""  
              method: "GET"
              as: "productDetail"
              timeout: 200

    - fetch:
        name: "Giá vàng hàng ngày"
        url: "https://www.pnj.com.vn/site/gia-vang?r=1709798169672"
        method: "GET"
        as: "price"
        timeout: 200
```

**Điểm nổi bật:**

* Hỗ trợ nhiều cron schedule cho cùng workflow
* Script action inline hoặc file
* Async `foreach` xử lý nhiều item
* Fetch action với success/error handling, save hoặc switch flow
* Có thể tích hợp Chromedp để scrape web động


## Đặc biệt với Chromedp Automation

KitWork tích hợp **Chromedp** cho khả năng tự động hóa web mạnh mẽ, điều khiển headless Chrome/Chromium trực tiếp từ YAML. Đây là tính năng **nâng cao**, giúp xử lý các tác vụ web phức tạp mà API thường không làm được.

### Tính năng chính

* **Điều hướng & tương tác:** mở web, click button, fill form, scroll page
* **Trích xuất dữ liệu:** DOM, innerText, HTML hoặc JSON
* **Thực thi JavaScript:** chạy script trực tiếp trên page
* **Chụp ảnh & PDF:** screenshot hoặc xuất page ra PDF
* **Tích hợp workflow:** kết hợp với fetch, script, save trong cron hoặc async foreach
* **Web Automation Beyond APIs:** scrape, test UI, thao tác page động

### Các bước Chromedp có thể dùng trong YAML

* `wait` – chờ selector xuất hiện
* `click` – click element
* `fill` – nhập giá trị vào input/textarea
* `evaluate` – chạy JS trên page và lưu kết quả
* `screenshot` – chụp toàn page hoặc selector
* `navigate` – đi tới URL mới
* `scroll` – cuộn page đến vị trí hoặc selector
* `select` – chọn giá trị dropdown/select
* `extract` – trích attribute từ element


## Góc nhà & Nhật ký

* Action được trigger bởi request, command, schedule, listen, activate, hoặc event (input …)
* Action chỉ trả kết quả **success** hoặc **error** (như Golang) và xử lý theo case
* Router chỉ handle HTTP methods cơ bản: **get, post, put, delete** và có bảo vệ
* Alias path mapping: `{id}` → param (`:id` trong GoFiber), `{$}` → `*`

```yaml
/router
  /api/  
    /{$}
      get.yaml # API không tìm thấy
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
```

* `return` thoát action ngay lập tức
* DB operations đơn giản: select, create, update, delete

```yaml
select: 
  db: "postgres"
  from: "table"
  where: ...
  offset: 0
  limit: 50
```

* Embedded commands cho bảo mật (fixed, immutable)
* Scripts mặc định là **JavaScript**, chạy bằng **v8go**
* Có thể mở rộng hỗ trợ biến (`var`) trong tương lai
* Mobile apps biên dịch từ Golang, hỗ trợ native, OpenGL/Vulkan, HTML (WASM hoặc Reactive)
* Load balancing / proxy tự host, nhanh
* YAML extendable (ví dụ: bảng, template)
* Core quản lý **cả hệ thống phân tán**, node định danh bằng node ID
* Tương lai: UI automation, kiểu drag-and-drop, workflow/code path qua interface
* Không tập trung phát triển game, nhưng có hỗ trợ tối thiểu
* Thiết kế như **dynamic JAM-stack server architecture**
* Hai tính năng chính: **shortest path cache** và **full-text search**, triển khai lặp mà không trùng cache


##  Tác giả

**Huỳnh Nhân Quốc** ❤️ Nhà phát triển indie-stack mộng mơ

* KitModule: [@kitmodule](https://github.com/kitmodule)
* KitWork: [@kitwork](https://github.com/kitwork)

Được phát hành theo [MIT License](https://github.com/kitwork/engine/blob/master/LICENSE)

**KitWork giúp mọi người tự động hóa, phát triển full-stack app và sinh code Golang gốc từ một nguồn YAML duy nhất, cung cấp nền tảng đơn giản, nhanh và mở rộng cho phát triển hiện đại.**

> Tôi không tạo ra ngôn ngữ lập trình mới. Tôi tạo ra một cách lập trình mới, giúp con người và AI làm việc cùng nhau một cách liền mạch.

