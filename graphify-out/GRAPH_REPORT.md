# Graph Report - golang-alpha  (2026-09-04)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 91 nodes · 97 edges · 8 communities (7 shown, 1 thin omitted)
- Extraction: 85% EXTRACTED · 14% INFERRED · 1% AMBIGUOUS · INFERRED: 14 edges (avg confidence: 0.91)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `9d0fb6bf`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- Graphify
- SaaS Backend Learning Project
- Contributing
- Finish Ticket
- Load
- pull_request_template.md
- Start Ticket
- kuekelheim.de/golang-aplpha

## God Nodes (most connected - your core abstractions)
1. `Graphify` - 13 edges
2. `Contributing` - 11 edges
3. `Finish Ticket` - 8 edges
4. `Start Ticket` - 8 edges
5. `Load()` - 4 edges
6. `NewPostgres()` - 4 edges
7. `Target Architecture` - 4 edges
8. `SaaS Backend Learning Project` - 4 edges
9. `Ticket Lifecycle` - 4 edges
10. `Required Checks Policy` - 4 edges

## Surprising Connections (you probably didn't know these)
- `Ticket Branch Naming` --semantically_similar_to--> `Contributing`  [INFERRED] [semantically similar]
  .codex/skills/start-ticket/SKILL.md → CONTRIBUTING.md
- `Handler Service Repository SQL Flow` --semantically_similar_to--> `Target Architecture`  [INFERRED] [semantically similar]
  AGENTS.md → plan.md
- `Modular Monolith` --semantically_similar_to--> `Target Architecture`  [INFERRED] [semantically similar]
  AGENTS.md → plan.md
- `Go Learning Mode` --semantically_similar_to--> `Go Mentorship for a Senior Python Engineer`  [INFERRED] [semantically similar]
  AGENTS.md → SOUL.md
- `Simple Explicit Engineering` --semantically_similar_to--> `Modular Monolith`  [INFERRED] [semantically similar]
  SOUL.md → AGENTS.md

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Ticket to Review Lifecycle** — codex_skills_start_ticket_skill_start_ticket, contributing_ticket_lifecycle, codex_skills_finish_ticket_skill_finish_ticket, contributing_draft_pull_request_workflow [EXTRACTED 1.00]
- **Actionable Ticket Definition** — github_issue_template_implementation_implementation_ticket, github_issue_template_implementation_observable_acceptance_criteria, github_issue_template_implementation_go_learning_objective, codex_skills_start_ticket_skill_ticket_scope_validation [INFERRED 0.85]
- **Graphify Operational Lifecycle** — _codex_skills_graphify_references_add_watch_url_ingestion, _codex_skills_graphify_references_add_watch_folder_watch, _codex_skills_graphify_references_update_incremental_graph_update, _codex_skills_graphify_references_hooks_graph_maintenance_hooks [INFERRED 0.85]
- **Go Learning Engineering Guidance** — agents_go_learning_mode, soul_go_mentorship, soul_simple_explicit_engineering [INFERRED 0.95]
- **Pull Request Quality Gate** — github_pull_request_template_verification_checklist, github_workflows_ci_go_ci, github_workflows_codeql_codeql, github_workflows_dependency_review_dependency_review [INFERRED 0.95]
- **SaaS Modular Monolith Architecture** — agents_modular_monolith, agents_layered_request_flow, plan_target_architecture, agents_postgresql_persistence [INFERRED 0.95]

## Communities (8 total, 1 thin omitted)

### Community 0 - "Graphify"
Cohesion: 0.13
Nodes (17): Folder Watch, URL Ingestion, Agent-Crawlable Wiki, Graph Exports, Confidence Audit Trail, Semantic Extraction Specification, Cross-Repository Graph Merge, Graph Maintenance Hooks (+9 more)

### Community 1 - "SaaS Backend Learning Project"
Cohesion: 0.12
Nodes (17): Go Learning Mode, Handler Service Repository SQL Flow, Modular Monolith, PostgreSQL Persistence, Tenant Isolation, PostgreSQL Data Volume, PostgreSQL Health Check, PostgreSQL 17 Alpine Service (+9 more)

### Community 2 - "Contributing"
Cohesion: 0.16
Nodes (14): Before opening a pull request, CI and security checks, Contributing, Draft Pull Request Workflow, Naming, Pull requests, Required Checks Policy, Ticket Lifecycle (+6 more)

### Community 3 - "Finish Ticket"
Cohesion: 0.17
Nodes (12): Finish Ticket Agent Interface, Acceptance Criteria Verification, Finish Ticket, Evidence-Backed Review Handoff, Separate Mutation Authorization, Workflow, Go Learning Objective, Implementation Ticket Template (+4 more)

### Community 4 - "Load"
Cohesion: 0.24
Nodes (7): main(), Config, context.Context, github.com/jackc/pgx/v5/pgxpool.Pool, Load(), LoadFromEnvOrDefault(), NewPostgres()

### Community 5 - "pull_request_template.md"
Cohesion: 0.20
Nodes (9): Akzeptanzkriterien, Ergebnis, Go-Lernziel, In Scope, Out of Scope, Review-Hinweise, Sicherheit, Migration und Betrieb, Umfang (+1 more)

### Community 6 - "Start Ticket"
Cohesion: 0.25
Nodes (7): Start Ticket Agent Interface, Ticket Branch Naming, Inputs, Start Ticket, Ticket Scope and Dependency Validation, Uncommitted Work Protection, Workflow

## Ambiguous Edges - Review These
- `SaaS Backend Learning Project` → `golang-alpha`  [AMBIGUOUS]
  README.md · relation: conceptually_related_to

## Knowledge Gaps
- **36 isolated node(s):** `Before opening a pull request`, `CI and security checks`, `Naming`, `Pull requests`, `Workflow` (+31 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 47 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **1 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **What is the exact relationship between `SaaS Backend Learning Project` and `golang-alpha`?**
  _Edge tagged AMBIGUOUS (relation: conceptually_related_to) - confidence is low._
- **Why does `Ticket Lifecycle` connect `Contributing` to `Finish Ticket`, `Start Ticket`?**
  _High betweenness centrality (0.073) - this node is a cross-community bridge._
- **Why does `Finish Ticket` connect `Finish Ticket` to `Contributing`?**
  _High betweenness centrality (0.072) - this node is a cross-community bridge._
- **Why does `Contributing` connect `Contributing` to `Start Ticket`?**
  _High betweenness centrality (0.065) - this node is a cross-community bridge._
- **What connects `Before opening a pull request`, `CI and security checks`, `Naming` to the rest of the system?**
  _36 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Graphify` be split into smaller, more focused modules?**
  _Cohesion score 0.1323529411764706 - nodes in this community are weakly interconnected._
- **Should `SaaS Backend Learning Project` be split into smaller, more focused modules?**
  _Cohesion score 0.125 - nodes in this community are weakly interconnected._