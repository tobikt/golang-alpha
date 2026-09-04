# AGENTS.md

# SaaS Backend – Hermes Agent Instructions

This repository is both a realistic SaaS backend and a hands-on Go learning project.

The goal is not only working code. Changes should remain understandable, idiomatic, secure, testable, and useful for learning Go and backend engineering.

The primary developer is a senior Python engineer and cloud architect who is learning Go.
Communicate at senior engineering depth, but teach Go-specific syntax, idioms, tooling,
and runtime behavior explicitly. Do not reteach generic backend or cloud concepts unless
the current decision depends on them.

## 1. Mission

Build a production-oriented example SaaS backend in Go as a **modular monolith**.

The domain will grow toward:

- users
- organizations
- memberships and roles
- projects
- authentication
- API keys
- audit logs
- background jobs
- mail
- billing
- observability
- security hardening
- CI/CD and deployment

Implement incrementally according to `plan.md`. Do not jump directly to the final architecture.

## 2. Source of Truth

Use this precedence:

1. explicit user instruction
2. active ticket/task
3. `plan.md`
4. this `AGENTS.md`
5. existing local conventions
6. general Go conventions

`plan.md` is the roadmap and ticket source of truth.

Never mark work complete unless its acceptance criteria are actually satisfied.

## 3. Architecture

Use one Go backend repository and one Go module.

Primary request flow:

```text
HTTP
  ↓
Handler
  ↓
Service
  ↓
Repository
  ↓
sqlc / PostgreSQL
```

Background flow:

```text
API
 ↓
Job Producer
 ↓
Queue / PostgreSQL
 ↓
Worker
```

Do not introduce microservices, Kubernetes, Kafka, Redis, Elasticsearch, GraphQL, CQRS, event sourcing, or other infrastructure without a concrete requirement or ticket.

## 4. Repository Layout

Expected direction:

```text
cmd/
├── api/
└── worker/

internal/
├── users/
├── organizations/
├── memberships/
├── projects/
├── auth/
├── apikeys/
├── billing/
├── audit/
├── jobs/
└── platform/
    ├── config/
    ├── database/
    ├── httpserver/
    ├── logger/
    ├── telemetry/
    ├── mail/
    └── clock/

db/
├── migrations/
├── queries/
└── sqlc.yaml

api/
docs/
scripts/
test/

compose.yaml
Dockerfile
Makefile
go.mod
go.sum
plan.md
README.md
```

Do not create empty architecture for hypothetical features.

## 5. Go Module and Dependencies

Use the root `go.mod`. Do not create nested Go modules.

Read the actual module path from `go.mod`; never assume it.

Prefer the standard library when it solves the problem well.

Before adding a dependency, ask:

- Is it maintained?
- Does stdlib already solve this?
- Does it materially simplify the code?
- Does it create unnecessary framework lock-in?

After dependency changes, normally run:

```bash
go mod tidy
```

## 6. Naming

Follow idiomatic Go.

Packages are short, lowercase, and descriptive:

```text
users
billing
database
httpserver
```

Avoid vague package names:

```text
common
utils
helpers
misc
```

Use Go casing:

```go
databaseURL
projectID
HTTPAddr
DatabaseURL
APIKey
NewService
NewPostgres
```

Use `ID`, `HTTP`, `URL`, `API`, `JSON`, `SQL`, `UUID` casing.

Export only what another package needs. A capitalized identifier is part of the package API.

Prefer `users.Service` over redundant names such as `users.UserService` when the package already provides context.

## 7. Package Design

Organize business code primarily by domain:

```text
internal/users/
internal/projects/
internal/organizations/
```

not global layer folders such as:

```text
internal/handlers/
internal/services/
internal/repositories/
internal/models/
```

A domain package may contain:

```text
model.go
service.go
repository.go
postgres.go
handler.go
errors.go
```

Keep related behavior together. Avoid excessive subpackages.

## 8. Dependency Direction

`cmd/...` is the composition root and may know concrete implementations.

Business services must not initialize infrastructure themselves.

Avoid:

- DB connection creation inside services
- `os.Getenv` throughout business packages
- global database handles
- global services
- provider SDKs imported everywhere

Pass dependencies explicitly through constructors.

```go
type Service struct {
    repo Repository
}

func NewService(repo Repository) *Service {
    return &Service{repo: repo}
}
```

## 9. Interfaces

Do not create interfaces preemptively.

Use them when they create a real boundary, such as persistence, an external provider, or a useful testing seam.

Prefer small consumer-defined interfaces.

Avoid `IUserService` or `IRepository` naming.

Do not build giant repository interfaces spanning unrelated domains.

## 10. Constructors and Options

Use conventional constructors when initialization is useful:

```go
NewService(...)
NewRepository(...)
NewPostgres(...)
NewServer(...)
```

Use direct parameters for a small number of required dependencies.

Use config structs for many related values.

Use functional options only when they materially improve a complex/public API.

## 11. Configuration

Configuration belongs under:

```text
internal/platform/config
```

Load configuration from environment variables.

Never commit secrets.

`.env.example` may document keys but must not contain real credentials.

Business packages receive configuration/dependencies; they should not repeatedly call `os.Getenv`.

Validate required configuration at startup and fail early.

## 12. PostgreSQL

PostgreSQL is the primary database.

Use `pgx/v5` and `pgxpool`.

Create one pool during startup, pass it where needed, and close it during shutdown.

Do not create a pool per request or store it globally.

Use `context.Context` for DB work.

Wrap infrastructure errors with context while preserving the cause:

```go
return fmt.Errorf("create project: %w", err)
```

Never expose raw SQL/driver errors to HTTP clients.

## 13. Migrations

Schema changes must be versioned migrations under:

```text
db/migrations/
```

Every schema change should consider:

- forward migration
- rollback implications
- constraints
- indexes
- foreign keys
- existing data
- production safety

Do not manually mutate shared/production schemas outside migrations.

Do not edit migrations that may already have been applied; add another migration.

## 14. SQL and sqlc

Prefer explicit SQL over a large ORM.

Queries live under:

```text
db/queries/
```

Use `sqlc` for type-safe generated Go DB code.

Never hand-edit generated code.

Business logic does not belong in generated query code.

Repository code owns persistence behavior and translations needed by the domain.

## 15. Transactions

Use transactions when a business operation must be atomic.

Examples:

- organization creation + owner membership
- invariant-preserving membership changes
- billing state transitions
- idempotent webhook processing

Transaction boundaries should follow business operations, not arbitrary layers.

Do not wrap every DB call in a transaction by default.

## 16. HTTP API

Use lightweight Go HTTP infrastructure.

Routes are versioned:

```text
/v1/...
```

Handlers should:

1. parse transport input
2. perform transport validation
3. call a service
4. map domain errors to HTTP
5. write the response

Handlers should not contain substantial business logic.

Services should not know HTTP status codes.

Repositories should not know HTTP.

## 17. JSON

Use `encoding/json` unless measurements justify another library.

Use explicit request/response structs.

Do not expose DB structs as API contracts merely because fields currently match.

For PATCH operations, distinguish omitted, zero, and null values with pointers or another deliberate representation.

## 18. Errors

Errors are values.

Use:

```go
errors.Is
errors.As
fmt.Errorf("context: %w", err)
```

when appropriate.

Domain errors should be distinguishable from infrastructure errors.

Examples:

```text
ErrUserNotFound
ErrEmailAlreadyExists
ErrForbidden
ErrProjectArchived
```

Do not expose SQL errors, stack traces, secrets, or provider internals to clients.

Do not use panic for normal runtime errors.

## 19. Context

Pass `context.Context` as the first parameter for request-scoped or cancellable work.

```go
func (s *Service) CreateProject(ctx context.Context, ...) (...)
```

Do not store context in long-lived structs.

Do not use context as a generic dependency bag.

Respect cancellation and timeouts.

## 20. Concurrency

Do not add goroutines unless concurrency is required.

Every goroutine must have a clear owner and shutdown path.

Ask:

- who owns it?
- how does it stop?
- where do errors go?
- can it leak?

Prefer context cancellation.

For concurrency-sensitive changes run:

```bash
go test -race ./...
```

when practical.

## 21. Logging

Prefer structured logging with standard-library `log/slog` unless the project explicitly chooses another logger.

Useful context includes:

```text
request_id
user_id
organization_id
method
path
duration
job_id
```

Never log passwords, auth tokens, full API keys, session secrets, or unreviewed sensitive payloads.

Avoid logging the same error redundantly at every layer.

## 22. Authentication and Authorization

Keep them separate.

Authentication answers: **who is the caller?**

Authorization answers: **may the caller do this?**

Do not trust client-supplied ownership or organization claims.

Tenant access must be enforced by queries and business rules.

API key secrets must eventually be stored as hashes, not plaintext.

## 23. Multi-Tenancy

Tenant isolation is a security requirement.

Organization-owned resources should normally be queried with organization scope.

Prefer:

```sql
WHERE id = $1
  AND organization_id = $2
```

over querying by a resource ID alone and relying on scattered ownership checks.

Include negative cross-tenant tests.

## 24. Testing

Use Go's standard `testing` package by default.

Use `*_test.go`.

Prefer table-driven tests when multiple cases exercise the same behavior.

Test service business rules directly.

Use small fakes/stubs where useful; do not mock mechanically.

Use `httptest` for handler tests.

Repository queries should eventually have integration tests against real PostgreSQL.

Cover:

- happy path
- validation
- domain-rule failures
- not found
- dependency failures
- tenant isolation

Do not delete or weaken tests just to make a change pass.

## 25. Learning Mode

This repo is intentionally used to learn Go.

Prefer understandable code over clever abstractions.

When a task naturally introduces a Go concept, briefly explain:

- what it is
- why it is used
- what problem it solves
- one common mistake

Distinguish Go language conventions from project architecture.

Do not dump a large implementation without explaining the important new concepts.

Do not over-explain concepts the user already understands.

Use Python analogies selectively. Always identify where the analogy stops being valid.
In particular, guard against importing Python habits into Go: exception-style error flow,
inheritance-heavy designs, broad framework abstractions, implicit dependency wiring, and
unnecessary dynamic behavior.

For meaningful changes, the handoff should briefly state:

- the Go concept introduced
- the closest useful Python comparison, when helpful
- the important semantic difference
- one common mistake to avoid

When useful, connect work to relevant `GO-xxx` tickets in `plan.md`.

## 26. Ticket Workflow

Work on one ticket or a small tightly related group at a time.

Before implementation:

1. read the ticket in `plan.md`
2. inspect relevant code and tests
3. inspect dependencies
4. identify the smallest complete change

During implementation:

1. make focused changes
2. preserve package boundaries
3. add/update tests
4. format
5. run targeted checks

After implementation:

1. verify acceptance criteria
2. summarize changes
3. explain important Go concepts
4. state commands/tests actually run
5. mention deferred work
6. update ticket status when requested

Do not silently implement unrelated roadmap items.

When deriving GitHub issues from `plan.md`, every implementation issue must be independently
actionable and include:

- problem/context and intended outcome
- in-scope and explicitly out-of-scope work
- concrete acceptance criteria phrased as observable results
- dependencies and the milestone/phase
- verification commands or manual checks
- a focused Go learning objective, including a Python comparison when useful
- security, migration, operational, or documentation notes when relevant

Preserve ticket IDs from `plan.md` in issue titles. Do not create remote issues, labels,
milestones, or project-board items unless the user explicitly requests that external write.

## 27. Definition of Done

For relevant code tickets, normally verify:

```bash
go fmt ./...
go vet ./...
go test ./...
```

Also require:

- code compiles
- acceptance criteria satisfied
- errors handled
- new business logic tested
- no secrets committed
- generated code current
- idiomatic naming
- justified dependencies
- docs updated when behavior/setup changes

Never claim a check was run if it was not.

## 28. Generated Files

Never manually edit generated files.

For sqlc:

1. modify schema/query source
2. regenerate
3. compile/test

If generation tooling is unavailable, say so rather than patching generated output.

## 29. Docker

Docker Compose is for reproducible local infrastructure.

PostgreSQL currently runs through Compose.

Do not destroy persistent volumes unless explicitly required.

Treat data-deleting commands with extra care.

Production containers should later use multi-stage builds and a minimal runtime image.

## 30. Security

Never:

- commit real secrets
- hardcode production credentials
- store plaintext passwords
- store plaintext API key secrets
- concatenate untrusted values into SQL
- return internal stack traces
- disable production TLS verification
- weaken authorization for convenience
- expose cross-tenant data

Use parameterized SQL.

Validate webhook signatures.

Add rate limits, body limits, HTTP timeouts, and vulnerability scanning when their roadmap tickets are reached.

## 31. External Providers

Keep provider SDKs behind the owning package boundary.

Example:

```text
billing
  ↓
PaymentProvider
  ↑
Stripe adapter
```

Do not import provider SDKs throughout unrelated packages.

Do not abstract hypothetical provider swaps unless the boundary already has practical value.

## 32. Avoid Overengineering

Prefer the smallest clean solution for the current requirement.

Do not create ceremonial layers such as:

```text
domain/
application/
usecase/
ports/
adapters/
factories/
managers/
utils/
```

unless they have a clear responsibility.

Do not turn a simple endpoint into many interfaces/files without a reason.

Explicit, boring Go is often good Go.

## 33. Refactoring

Refactor for an actual pain point:

- duplication
- confusing ownership
- difficult testing
- leaky dependencies
- repeated error handling
- repeated transaction logic

Avoid unrelated broad refactors during feature tickets.

Keep diffs reviewable and preserve behavior with tests.

## 34. Documentation

`README.md` is the fast entry point.

Use `docs/` for deeper documentation.

Long-lived architecture decisions belong under:

```text
docs/decisions/
```

OpenAPI belongs under:

```text
api/
```

not general docs.

## 35. Commands and Environment

Never assume a tool or Make target exists. Inspect the repository first.

Common Go commands:

```bash
go run ./cmd/api
go build ./cmd/api
go test ./...
go fmt ./...
go vet ./...
go mod tidy
```

Migration, sqlc, Docker, Makefile, and CI commands must follow the actual files present.

The project may be developed on Windows/PowerShell. Do not assume Bash-specific syntax when giving commands.

Repository code should remain platform-neutral unless a requirement says otherwise.

## 36. Agent Operating Rules

Before changing code:

- inspect relevant files
- inspect `go.mod`
- inspect `plan.md`
- inspect nearby tests
- understand the package boundary

Do not guess file contents.

Do not invent dependencies, ports, environment variables, Make targets, migration versions, or package paths.

Do not rewrite working code merely to satisfy stylistic preference.

Prefer task completion over speculative redesign.

For significant architectural changes, explain the reason before making the change.

## 37. Collaboration Style

The user is learning Go while building this backend.

Act like a senior engineer pairing with a developer.

For meaningful changes, explain:

- what changed
- why it belongs there
- the important Go concept
- how to verify it

Recommend one approach when there is a clear best fit and explain the tradeoff.

Preserve momentum rather than redesigning everything.

## 38. Never Fake Progress

Never:

- mark a ticket done without implementing it
- claim tests passed without running them
- claim a migration succeeded without verifying it when execution is required
- claim the API starts without verification when verification is available
- hide known failures
- silently skip acceptance criteria

Partial completion is acceptable when stated clearly.

False completion is not.

## 39. Current State

At the time this file was authored, the foundation already included:

- Git repository and `main` branch
- Go module and base folders
- runnable HTTP API
- `/health`
- environment config
- PostgreSQL via Docker Compose
- Go PostgreSQL connectivity with `pgxpool`

Treat `plan.md` as authoritative because this section can become stale.

The next roadmap area is migrations and sqlc, followed by the first complete vertical business slice.

## 40. Guiding Principle

Optimize in this order:

```text
correctness
  ↓
clarity
  ↓
idiomatic Go
  ↓
security
  ↓
testability
  ↓
maintainability
  ↓
measured performance
```

Build the simplest thing that is correct today and leaves a clean path for tomorrow.

## graphify

This project has a knowledge graph at graphify-out/ with god nodes, community structure, and cross-file relationships.

When the user types `/graphify`, use the installed graphify skill or instructions before doing anything else.

Rules:
- For codebase questions, first run `graphify query "<question>"` when graphify-out/graph.json exists. Use `graphify path "<A>" "<B>"` for relationships and `graphify explain "<concept>"` for focused concepts. These return a scoped subgraph, usually much smaller than GRAPH_REPORT.md or raw grep output.
- Dirty graphify-out/ files are expected after hooks or incremental updates; dirty graph files are not a reason to skip graphify. Only skip graphify if the task is about stale or incorrect graph output, or the user explicitly says not to use it.
- If graphify-out/wiki/index.md exists, use it for broad navigation instead of raw source browsing.
- Read graphify-out/GRAPH_REPORT.md only for broad architecture review or when query/path/explain do not surface enough context.
- After modifying code, run `graphify update .` to keep the graph current (AST-only, no API cost).
