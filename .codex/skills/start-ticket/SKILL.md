---
name: start-ticket
description: Start a ticket from plan.md or GitHub safely by validating its scope and dependencies, inspecting relevant repository context, and creating a correctly named local branch. Use when beginning a new implementation ticket; do not use to finish, push, or merge work.
---

# Start Ticket

Prepare one ticket for focused implementation.

## Inputs

Accept a ticket ID such as `DB-010`, a GitHub issue number or URL, and an
optional branch prefix. If no identifier is provided, ask for exactly one.

## Workflow

1. Read `AGENTS.md`, the ticket in `plan.md`, `go.mod`, nearby code, and nearby
   tests. For a GitHub issue, retrieve it read-only and reconcile it with
   `plan.md`; report material conflicts instead of silently choosing one.
2. Use the existing Graphify graph first when available to locate relevant
   packages and relationships. Do not rebuild the graph merely to start work.
3. Summarize the outcome, scope, exclusions, acceptance criteria, dependencies,
   verification, and Go learning objective. Stop if an incomplete dependency
   makes implementation unsafe or would change the intended architecture.
4. Inspect the current branch and `git status --short`. Never switch branches,
   pull, or create a branch over uncommitted work without the user's explicit
   direction. Do not stash, discard, or rewrite user changes.
5. Start from an up-to-date `main`. Network access and changes to the local main
   branch must remain visible; request authorization if the environment requires
   it. Use fast-forward-only pull behavior.
6. Derive a branch name using `<type>/<TICKET-ID>-<short-slug>`. Choose `feat`,
   `fix`, `test`, `docs`, `refactor`, or `chore` from the actual ticket. Keep
   the type and slug lowercase and preserve the uppercase ticket ID.
7. For a GitHub issue, first run `gh issue develop --list` and reuse an existing
   linked development branch when appropriate. Otherwise create and link the
   branch natively with:

   ```text
   gh issue develop <issue> --repo <owner/repo> --base main \
     --name <branch-name> --checkout
   ```

   This remote branch creation is an external write. State the repository,
   issue, base, and branch and obtain explicit confirmation immediately before
   running it. Do not replace this with an issue comment: a comment is not a
   GitHub Development-branch link. If a branch with that name already exists
   but is not linked, do not delete or recreate it merely to add the link;
   report the limitation and let the pull request link the issue instead.
8. For a ticket that has no GitHub issue, create a local branch with
   `git switch -c <branch-name>` only after the prior checks succeed.
9. Verify the resulting relationship: for GitHub issues, run
   `gh issue develop --list <issue> --repo <owner/repo>`; for local-only tickets,
   verify the current branch. Do not claim the issue is linked without this
   check.
10. Present the smallest complete implementation plan and the commands that will
   verify it. Do not implement unrelated roadmap work.

Creating a linked development branch, changing an Issue, or changing a GitHub
Project status is an external write. Perform each only after explicit
confirmation. Offer to mark the ticket `In Progress`, but do it only after that
confirmation. This skill does not commit, push, open a pull request, or merge.

End with the selected ticket, created/current branch, dependency status, and
next implementation step.
