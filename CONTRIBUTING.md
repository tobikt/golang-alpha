# Contributing

This repository is developed ticket by ticket. `plan.md` is the roadmap and
source of truth; GitHub Issues are the executable work items derived from it.

## Ticket lifecycle

1. Pick one ticket whose dependencies are complete.
2. Start it with `$start-ticket <ticket-id or issue-number>`.
3. Implement only the ticket scope and its acceptance criteria.
4. Finish it with `$finish-ticket <ticket-id>`.
5. Push the branch and open a draft pull request.
6. Resolve CI and review findings, then mark the pull request ready.
7. Merge only after all required checks pass. `Closes #<issue>` closes the
   linked issue when the pull request is merged.

Changing a GitHub Issue or Project status is an external write. The local
skills prepare it, but perform it only after explicit confirmation.

## Naming

Use the ticket ID in branches, commits, and pull requests:

```text
Branch: feat/GO-040-user-model
        fix/AUTH-021-reject-expired-token
        chore/OPS-010-go-ci

Commit: GO-040: add user domain model
PR:     GO-040: Add user domain model
```

Prefer `feat`, `fix`, `test`, `docs`, `refactor`, or `chore` as the branch
prefix. Keep one ticket per branch unless the tickets form a deliberately small,
tightly related group.

## Before opening a pull request

Run the checks that apply to the change:

```powershell
gofmt -w .
go vet ./...
go test ./...
```

For concurrency-sensitive changes, also run:

```powershell
go test -race ./...
```

After changing project code or long-lived documentation, update the local
knowledge graph:

```powershell
graphify update .
```

Do not claim a check passed unless it was actually run. Document unavailable or
failing checks in the pull request.

## Pull requests

Open the pull request as a draft while implementation or self-review is still
in progress. Complete the repository pull request template and include:

- the intended outcome and ticket/issue link;
- in-scope changes and explicit exclusions;
- observable acceptance criteria;
- commands that were actually run;
- migration, security, and operational implications;
- the Go learning objective and, where useful, its Python comparison.

Keep changes reviewable. Do not mix speculative refactors or later roadmap work
into the ticket.

## CI and security checks

Pull requests run formatting, `go vet`, tests, CodeQL, and dependency review.
The default branch also receives the Go checks and CodeQL analysis. A failing
required check must be fixed or explicitly understood; never weaken tests merely
to make CI green.
