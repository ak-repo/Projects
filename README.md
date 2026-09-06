# Engineering Workspace

This workspace contains backend, frontend, systems, tooling, DevOps, and
future-facing experiments. The top-level folders describe ownership and
purpose rather than only programming language.

## Structure

```text
implementations/
├── engineering/
│   ├── backend/
│   │   ├── services/
│   │   ├── api-protocols/
│   │   ├── databases/
│   │   ├── messaging/
│   │   ├── caching/
│   │   ├── storage/
│   │   ├── realtime/
│   │   └── security/
│   ├── frontend/react/
│   ├── languages/{go,rust,typescript}/
│   ├── foundations/
│   └── tools/
├── devops/
│   ├── containers/
│   ├── ci-cd/
│   ├── cloud/
│   ├── observability/
│   ├── automation/
│   └── local-development/
├── future/
│   ├── products/
│   ├── prototypes/
│   ├── ai/
│   ├── robotics/
│   └── research/
├── shared/
│   ├── contracts/
│   ├── libraries/
│   ├── documentation/
│   └── templates/
└── archive/
```

## Current Projects

| Project | Location | Focus |
| --- | --- | --- |
| Cron order worker API | `engineering/backend/services/` | Go service, jobs, orders |
| GraphQL REST demo | `engineering/backend/api-protocols/` | API design |
| gRPC demo | `engineering/backend/api-protocols/` | Protocol buffers and gRPC |
| Kafka | `engineering/backend/messaging/` | Event streaming |
| MySQL | `engineering/backend/databases/mysql/` | Relational database |
| Elasticsearch demo | `engineering/backend/databases/` | Search |
| Redis demo | `engineering/backend/caching/` | Caching |
| S3 demo | `engineering/backend/storage/` | Object storage |
| WebSocket | `engineering/backend/realtime/` | Real-time communication |
| OAuth | `engineering/backend/security/` | Authentication and authorization |
| Concurrency | `engineering/foundations/` | Go concurrency concepts |
| Go project skeleton | `engineering/tools/project-templates/` | Go project starter |

## Placement Rules

- Classify projects by their primary responsibility before their language.
- Keep each project self-contained, including its migrations, web files,
  Docker files, and configuration examples.
- Put reusable Go, Rust, or TypeScript libraries and language exercises under
  `engineering/languages/`.
- Put reusable templates, contracts, and libraries used by multiple projects
  under `shared/`.
- Put deployment, CI/CD, cloud, monitoring, and environment automation under
  `devops/`.
- Put new applications, prototypes, AI work, robotics, and research under
  `future/`.
- Move obsolete experiments to `archive/` instead of deleting them.

## Dependency Direction

The intended flow is:

```text
shared/       <- engineering/ <- future/
                     ^              ^
                     └── devops/ ───┘
```

`future/` may consume capabilities from `engineering/`, `shared/`, and
`devops/`. Engineering projects should not depend on future projects. A
project-specific deployment file may remain inside that project; shared
deployment components belong in `devops/`.

## Naming

- Use lowercase kebab-case for directories, such as `api-protocols`.
- Use `README.md` for project documentation.
- Keep each project independently buildable and testable.
