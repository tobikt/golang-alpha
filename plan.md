# SaaS Backend Learning Project – Master Plan

> Ziel: Schritt für Schritt ein produktionsnahes SaaS-Backend in Go bauen und dabei Go, Backend-Architektur, PostgreSQL, Testing, Security, Observability und Deployment praktisch lernen.
>
> Dieser Plan ist absichtlich ticket-orientiert aufgebaut, damit die einzelnen Punkte später direkt in GitHub Issues, Jira, Linear oder ein anderes Ticketsystem übernommen werden können.

---

# 1. Projektziel

Wir bauen ein Beispiel-SaaS als **modularen Monolithen**.

Das Beispielprodukt ist ein einfaches **Multi-Tenant Project SaaS**:

- Benutzer können Accounts haben.
- Benutzer können Organisationen anlegen.
- Benutzer können Mitglied mehrerer Organisationen sein.
- Organisationen haben Rollen und Berechtigungen.
- Organisationen können Projekte anlegen.
- Projekte können archiviert werden.
- Organisationen können unterschiedliche Pläne besitzen.
- Billing wird später über Stripe angebunden.
- Nutzer können API Keys erzeugen.
- Änderungen werden in Audit Logs geschrieben.
- E-Mails und andere langsame Aufgaben laufen über Background Jobs.
- Die Anwendung stellt eine versionierte HTTP API bereit.

Das Ziel ist nicht, ein fertiges kommerzielles Produkt zu bauen, sondern eine realistische Referenzarchitektur, an der wir Go systematisch lernen können.

---

# 2. Zielarchitektur

```text
Client
  |
  v
HTTP API
  |
  v
Handler
  |
  v
Service / Business Logic
  |
  v
Repository
  |
  v
PostgreSQL
```

Zusätzliche Komponenten:

```text
                   +------------------+
                   |   HTTP Client    |
                   +--------+---------+
                            |
                            v
                   +------------------+
                   |     Go API       |
                   +--------+---------+
                            |
             +--------------+--------------+
             |              |              |
             v              v              v
        PostgreSQL      Job Queue      External APIs
                            |          Stripe / Mail
                            v
                         Worker
```

---

# 3. Architekturprinzipien

- Ein Go-Modul für das gesamte Backend.
- Modularer Monolith statt Microservices.
- Business-Code wird nach Fachmodulen organisiert.
- Technische Infrastruktur liegt unter `internal/platform`.
- Dependencies werden explizit über Konstruktoren übergeben.
- Keine globalen Services.
- Keine globale Datenbankvariable.
- Interfaces werden nur dort eingeführt, wo sie einen echten Nutzen haben.
- Interfaces werden bevorzugt vom Consumer definiert.
- PostgreSQL ist die primäre Datenbank.
- SQL wird bewusst gelernt und genutzt.
- `sqlc` generiert typsicheren Go-Code aus SQL.
- Externe Anbieter werden hinter kleinen Adaptern gekapselt.
- HTTP Handler enthalten möglichst wenig Business-Logik.
- Services enthalten Business-Regeln.
- Repository-Code kümmert sich um Persistenz.
- Konfiguration kommt aus Environment Variables.
- Secrets werden niemals ins Repository committed.
- Migrationen sind Teil des Codes.
- Tests werden parallel zur Anwendung aufgebaut.
- Observability wird nicht erst ganz am Ende betrachtet.

---

# 4. Zielstruktur des Repositories

```text
saas-backend/
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── worker/
│       └── main.go
│
├── internal/
│   ├── users/
│   ├── organizations/
│   ├── memberships/
│   ├── projects/
│   ├── auth/
│   ├── apikeys/
│   ├── billing/
│   ├── audit/
│   ├── jobs/
│   └── platform/
│       ├── config/
│       ├── database/
│       ├── httpserver/
│       ├── logger/
│       ├── telemetry/
│       ├── mail/
│       └── clock/
│
├── db/
│   ├── migrations/
│   ├── queries/
│   └── sqlc.yaml
│
├── api/
│   └── openapi.yaml
│
├── docs/
│   ├── architecture.md
│   ├── development.md
│   ├── deployment.md
│   └── decisions/
│
├── scripts/
├── test/
│
├── .env.example
├── .gitignore
├── compose.yaml
├── Dockerfile
├── Makefile
├── go.mod
├── go.sum
├── plan.md
└── README.md
```

---

# 5. Ticket-Konvention

Jedes Ticket bekommt:

- **ID**
- **Titel**
- **Ziel**
- **Umsetzung**
- **Akzeptanzkriterien**
- **Go-Lernziel**
- **Abhängigkeiten**

Ticket-ID Schema:

```text
FOUND-xxx   Foundation
DB-xxx      Database
USER-xxx    Users
ORG-xxx     Organizations
MEM-xxx     Memberships
PROJ-xxx    Projects
AUTH-xxx    Authentication
API-xxx     HTTP API
TEST-xxx    Testing
JOB-xxx     Background Jobs
BILL-xxx    Billing
OBS-xxx     Observability
SEC-xxx     Security
OPS-xxx     Operations / Deployment
DOC-xxx     Documentation
```

# 9. Was wir bewusst zunächst NICHT machen

Um das Projekt verständlich zu halten:

- kein Kubernetes
- keine Microservices
- kein Kafka
- kein Event Sourcing
- kein CQRS Framework
- keine komplizierte Dependency Injection Library
- kein großes ORM
- kein GraphQL zum Start
- kein Redis ohne konkreten Bedarf
- keine separate Search Engine
- keine selbstgebaute verteilte Infrastruktur

Diese Dinge können später bewusst als Lernmodule ergänzt werden.
