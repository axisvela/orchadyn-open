# ORCHADYN Open

> Make business architecture executable without turning planning into another runtime.

## Product Introduction

ORCHADYN is an Enterprise Capability Planning Compiler for Executable Business
Architecture. It transforms business goals, governed capabilities,
organizational authority, operational context, and constraints into plans that
can be verified before execution and projected to heterogeneous runtimes.

Business architecture is the source language. Agent graphs, durable workflows,
robot missions, and other execution systems are compilation targets. ORCHADYN
does not execute the work itself, own the capability registry, or replace
identity and operational-state authorities.

## Problems It Solves

Enterprise automation usually begins at the execution layer: agents, tools,
workflows, and runtime state. This leaves a planning gap between a business
goal and a governed execution decision. Teams repeatedly encode capability
choices, budgets, organization boundaries, policy constraints, and evidence
requirements in application-specific workflows.

ORCHADYN addresses that gap by making Goal, Capability, Authority, Constraint,
and Plan first-class planning semantics. It separates a required capability
from its runtime implementation, so a plan can be governed, explained, revised,
and projected without being locked to a single agent framework, workflow engine,
vendor, or deployment environment.

It also makes delegation accountable: a plan can retain who authorized an
action, the authority and budget passed downstream, the evidence required for
completion, and why other planning alternatives were rejected.

## Open Repository

This repository contains Apache-2.0 licensed documentation, SDKs, examples,
public API definitions, and binary release manifests. The AGPL planning kernel
is maintained in [ORCHADYN](https://github.com/axisrobo/ORCHADYN).

## Contents

- `api/`: versioned public HTTP API contracts.
- `examples/`: API and SDK usage examples.
- `docs/`: installation, architecture, and integration guidance.
- `releases/`: checksums and notices for published binaries.

## Compatibility

Open artifacts use only published ORCHADYN contracts. Examples must not import
Core internals or copy Core source. Enterprise capabilities are documented only
through their public compatibility contracts; their implementation is in
[ORCHADYN-ee](https://github.com/axisrobo/ORCHADYN-ee).

## License

Apache-2.0. See [LICENSE](LICENSE).
