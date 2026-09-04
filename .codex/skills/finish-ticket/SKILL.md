---
name: finish-ticket
description: Verify completed ticket work against its acceptance criteria, run the required Go and Graphify checks, review the diff, and prepare a commit and pull request handoff. Use when implementation is ready for review; do not use to start unrelated work or merge a pull request.
---

# Finish Ticket

Turn completed local work into an evidence-backed review handoff.

## Workflow

1. Identify the ticket from the argument, current branch, `plan.md`, or linked
   GitHub issue. If identification is ambiguous, ask for the ticket ID.
2. Read the complete acceptance criteria and inspect `git status`, the diff, and
   untracked files. Preserve unrelated user changes and report scope drift.
3. Check each acceptance criterion against observable code, tests, migrations,
   and documentation. Never mark incomplete work complete.
4. Run the repository's actual checks. For relevant Go changes, normally run:

   ```text
   gofmt -w .
   go vet ./...
   go test ./...
   ```

   Run `go test -race ./...` for concurrency-sensitive work. Regenerate sqlc
   output only from source files and only when the ticket requires it.
5. After code or long-lived documentation changes, run `graphify update .`.
   Treat generated Graphify changes as expected, but review their scope.
6. Reinspect the diff for secrets, generated-file mistakes, weakened tests,
   tenant-isolation risks, unintended files, and changes outside the ticket.
7. Prepare, but do not silently perform, a commit named
   `<TICKET-ID>: <imperative summary>` and a pull request titled
   `<TICKET-ID>: <outcome>`. Populate `.github/pull_request_template.md` with
   actual checks and `Closes #<issue>` when an issue number is known.

Local commits, pushes, GitHub status changes, pull-request creation, review
submission, and merge are separate mutations. Execute only those explicitly
requested or confirmed by the user. Prefer a draft pull request until the work
and self-review are complete; never merge with failing required checks.

End with acceptance-criteria status, checks actually run and their results,
remaining risks or deferred work, plus the proposed commit and PR metadata.
