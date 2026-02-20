# Project Agents Guidelines

## Tooling

This project uses `mise` for dependency management and task execution.

- **Install**: `mise install`
- **CI**: `mise run ci` (lint, test, build)
- **Lint**: `mise run lint` (uses `trunk`)
- **Test**: `mise run test`
- **Build**: `mise run build`
- **Codegen**: `mise run codegen`

## Conventions

- **Linting**: All code must pass `trunk check`.
- **Testing**: All code must pass `go test ./...`.
- **CI**: The `ci` task is the source of truth for CI checks.
- **Dependencies**: Manage tools via `mise.toml`.
