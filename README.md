# Go Utilities Toolkit

This repository contains reusable and modular Go utilities for solving common development sub-problems. Each utility is placed in a logically grouped package and is designed to be plug-and-play into your project.

---

## 🗂️ Project Structure

```plaintext
.
├── README.md
├── go.mod
├── go.sum
├── main.go
├── utils
│   ├── cache
│   │   └── cache.go
│   ├── csvutils
│   │   ├── csv_to_json.go
│   │   └── json_to_csv.go
│   ├── excel
│   │   └── reader.go
│   ├── fileutils
│   │   ├── exists.go
│   │   └── zip.go
│   ├── gcs
│   │   └── uploader.go
│   ├── httpclient
│   │   ├── json_post.go
│   │   └── retry.go
│   ├── markdown
│   │   └── to_html.go
│   ├── pdfgen
│   │   └── generate_pdf.go
│   ├── redisutils
│   │   └── redis.go
│   ├── retryutils
│   │   └── backoff.go
│   ├── stringutils
│   │   ├── email_validator.go
│   │   ├── phone_validator.go
│   │   ├── slugify.go
│   │   └── truncate.go
│   ├── timeutils
│   │   ├── formatter.go
│   │   ├── time_ago.go
│   │   └── cron.go
│   └── web
│       ├── html_parser.go
│       └── url_checker.go
└── tests
    └── handler_test.go
```

---

## 🧪 Example Usage

```go
package main

import (
	"fmt"
	"go_utilities_toolkit/utils/stringutils"
)

func main() {
	email := "test@example.com"
	if stringutils.IsValidEmail(email) {
		fmt.Println("Email is valid")
	} else {
		fmt.Println("Invalid email")
	}
}
```

---

## 📦 Modules Included
- File & CSV/Excel Utilities
- HTTP Client Wrappers
- Retry Logic with Backoff
- PDF Generator
- Cloud Storage (GCS) Uploader
- Redis Utility
- JWT and Token Helpers
- Time Formatting Utilities
- Markdown to HTML
- CLI Input Helpers
- URL & HTML Parsers
- Caching Layer (in-memory)

---

## 📌 Setup

```bash
git clone https://github.com/your-username/go-utilities-toolkit.git
cd go-utilities-toolkit
go mod tidy
go run main.go
```

---

## 🧠 Contributions
Feel free to open an issue or PR for improvements or new utilities.

---

## 📝 License
MIT

---

Want a version of this toolkit structured as a Go CLI app or microservices scaffold? Let me know!