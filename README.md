# ORCHADYN Open

Public distribution repository for the ORCHADYN Enterprise Capability Planning
Compiler.

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
