
# notes-service

A small Go REST API for taking notes.  
Built as part of a backend-systems demo

## Features
- CRUD `/notes` with JWT auth
- In-memory or PostgreSQL storage (flag-switchable)
- Middleware: logging, graceful shutdown
- Async task demo using a goroutine worker pool
- Unit tests & benchmarks (`go test ./...`, `go test -bench . ./internal/...`)

## Quick start
```bash
# clone & run
git clone https://github.com/<you>/notes-service.git
cd notes-service
go run ./cmd/server
```

### Environment variables
| name          | default      | purpose                |
|---------------|--------------|------------------------|
| `PORT`        | `8080`       | HTTP port              |
| `JWT_SECRET`  | `secret123`  | HMAC key for JWT       |

## API
| method | endpoint    | description           |
|--------|-------------|-----------------------|
| GET    | /healthz    | liveness check        |
| POST   | /signup     | create user           |
| POST   | /login      | obtain JWT token      |
| GET    | /notes      | list user notes       |
| POST   | /notes      | create new note       |

> `Authorization: Bearer <token>` required on `/notes`.

---

## Benchmarks

Simulated heavy task: `time.Sleep(100 * time.Millisecond)`  
Machine: Ryzen 7 5800H, Go 1.22

| Benchmark              | ns/op    |
|------------------------|----------|
| Serial (n*n)           | **0.68** |
| Pool (WaitGroup)       | 6,353,625 |
| Pool (Chan-only)       | 6,320,522 |

Counters

| Benchmark         |  ns/op |
|-------------------|--------|
| Counter-Atomic    |   13   |
| Counter-Mutex     |   52   |

---

## Roadmap
- [x] Core CRUD
- [x] JWT auth
- [x] Worker-pool demo
- [ ] Redis task queue (Week 3)
- [ ] Dockerfile + CI (Week 6)
- [ ] Fly.io deploy (Week 7)

---

## Dev
```bash
go vet ./...
golangci-lint run
```

