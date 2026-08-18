# Release Manifests

Published binary releases include checksums, the originating Core version, and
all applicable license notices. Every release's `SHA256SUMS` files are signed
with the ORCHADYN release signing key; since v0.5.0 the release SBOM is signed
alongside them.

## Signature Verification

The ORCHADYN release signing key:

- Fingerprint: `8304 0C4F 09F4 EA5A 0D5F 3161 4E34 E2EF 54FB 0822`
- Public key: [`v1.58.0/ORCHADYN-RELEASE-SIGNING-KEY.asc`](v1.58.0/ORCHADYN-RELEASE-SIGNING-KEY.asc)

Verify the SBOM and a platform's checksums for a release (example: v1.58.0):

```text
gpg --import releases/v1.58.0/ORCHADYN-RELEASE-SIGNING-KEY.asc
gpg --verify releases/v1.58.0/sbom.json.asc releases/v1.58.0/sbom.json
gpg --verify releases/v1.58.0/windows-amd64/SHA256SUMS.asc releases/v1.58.0/windows-amd64/SHA256SUMS
```

## v0.1.0

- Manifest: [`v0.1.0/`](v0.1.0/)
- Binaries: [`orchadyn-api`](https://github.com/axisrobo/orchadyn-open/releases/tag/v0.1.0)
  and [`orchadyn-mcp`](https://github.com/axisrobo/orchadyn-open/releases/tag/v0.1.0)
  for linux, darwin, and windows (amd64 + arm64).
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `6f15520`. Binaries are AGPL-3.0-or-later.

## v0.2.0

- Manifest: [`v0.2.0/windows-amd64/`](v0.2.0/windows-amd64/)
- Binaries: `orchadyn-api.exe`, `orchadyn-mcp.exe`, and
  `orchadyn-migrate.exe` for Windows/amd64.
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `8487670` (`v0.2.0`). Binaries are AGPL-3.0-or-later.

## v0.3.0

- Manifest: [`v0.3.0/`](v0.3.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `af5b540` (`v0.3.0`). Binaries are AGPL-3.0-or-later.

## v0.4.0

- Manifest: [`v0.4.0/`](v0.4.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `53186e3` (`v0.4.0`). Binaries are AGPL-3.0-or-later.

## v0.5.0

- Manifest: [`v0.5.0/`](v0.5.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v0.5.0/sbom.json`](v0.5.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `e5ccf88` (`v0.5.0`). Binaries are AGPL-3.0-or-later.

## v0.6.0

- Manifest: [`v0.6.0/`](v0.6.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v0.6.0/sbom.json`](v0.6.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `71ac6fb` (`v0.6.0`). Binaries are AGPL-3.0-or-later.

## v0.7.0

- Manifest: [`v0.7.0/`](v0.7.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v0.7.0/sbom.json`](v0.7.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `3dfcd3b` (`v0.7.0`). Binaries are AGPL-3.0-or-later.

## v0.8.0

- Manifest: [`v0.8.0/`](v0.8.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v0.8.0/sbom.json`](v0.8.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `cfaf231` (`v0.8.0`). Binaries are AGPL-3.0-or-later.

## v0.9.0

- Manifest: [`v0.9.0/`](v0.9.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v0.9.0/sbom.json`](v0.9.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `7e69112` (`v0.9.0`). Binaries are AGPL-3.0-or-later.

## v1.0.0

- Manifest: [`v1.0.0/`](v1.0.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.0.0/sbom.json`](v1.0.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `b208e6a` (`v1.0.0`). Binaries are AGPL-3.0-or-later.

## v1.1.0

- Manifest: [`v1.1.0/`](v1.1.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.1.0/sbom.json`](v1.1.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `afa8de1` (`v1.1.0`). Binaries are AGPL-3.0-or-later.
## v1.2.0

- Manifest: [`v1.2.0/`](v1.2.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.2.0/sbom.json`](v1.2.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `f033c66` (`v1.2.0`). Binaries are AGPL-3.0-or-later.

## v1.3.0

- Manifest: [`v1.3.0/`](v1.3.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.3.0/sbom.json`](v1.3.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `d58da7d` (`v1.3.0`). Binaries are AGPL-3.0-or-later.

## v1.4.0

- Manifest: [`v1.4.0/`](v1.4.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.4.0/sbom.json`](v1.4.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `61059fa` (`v1.4.0`). Binaries are AGPL-3.0-or-later.

## v1.5.0

- Manifest: [`v1.5.0/`](v1.5.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.5.0/sbom.json`](v1.5.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `b0f409f` (`v1.5.0`). Binaries are AGPL-3.0-or-later.

## v1.6.0

- Manifest: [`v1.6.0/`](v1.6.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.6.0/sbom.json`](v1.6.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `82d62e1` (`v1.6.0`). Binaries are AGPL-3.0-or-later.

## v1.7.0

- Manifest: [`v1.7.0/`](v1.7.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.7.0/sbom.json`](v1.7.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `74eb599` (`v1.7.0`). Binaries are AGPL-3.0-or-later.

## v1.8.0

- Manifest: [`v1.8.0/`](v1.8.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.8.0/sbom.json`](v1.8.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `8ef992f` (`v1.8.0`). Binaries are AGPL-3.0-or-later.

## v1.9.0

- Manifest: [`v1.9.0/`](v1.9.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.9.0/sbom.json`](v1.9.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `4fc3a3c` (`v1.9.0`). Binaries are AGPL-3.0-or-later.

## v1.10.0

- Manifest: [`v1.10.0/`](v1.10.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.10.0/sbom.json`](v1.10.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `0081ab5` (`v1.10.0`). Binaries are AGPL-3.0-or-later.

## v1.11.0

- Manifest: [`v1.11.0/`](v1.11.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.11.0/sbom.json`](v1.11.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `de9a414` (`v1.11.0`). Binaries are AGPL-3.0-or-later.

## v1.12.0

- Manifest: [`v1.12.0/`](v1.12.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.12.0/sbom.json`](v1.12.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `92a79f9` (`v1.12.0`). Binaries are AGPL-3.0-or-later.

## v1.13.0

- Manifest: [`v1.13.0/`](v1.13.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.13.0/sbom.json`](v1.13.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `8b61c08` (`v1.13.0`). Binaries are AGPL-3.0-or-later.

## v1.14.0

- Manifest: [`v1.14.0/`](v1.14.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.14.0/sbom.json`](v1.14.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `0d73268` (`v1.14.0`). Binaries are AGPL-3.0-or-later.

## v1.15.0

- Manifest: [`v1.15.0/`](v1.15.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.15.0/sbom.json`](v1.15.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `4612d4b` (`v1.15.0`). Binaries are AGPL-3.0-or-later.

## v1.16.0

- Manifest: [`v1.16.0/`](v1.16.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.16.0/sbom.json`](v1.16.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `0347410` (`v1.16.0`). Binaries are AGPL-3.0-or-later.

## v1.17.0

- Manifest: [`v1.17.0/`](v1.17.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.17.0/sbom.json`](v1.17.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `522dbbb` (`v1.17.0`). Binaries are AGPL-3.0-or-later.

## v1.18.0

- Manifest: [`v1.18.0/`](v1.18.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.18.0/sbom.json`](v1.18.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `308e494` (`v1.18.0`). Binaries are AGPL-3.0-or-later.

## v1.19.0

- Manifest: [`v1.19.0/`](v1.19.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.19.0/sbom.json`](v1.19.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `c84e016` (`v1.19.0`). Binaries are AGPL-3.0-or-later.

## v1.20.0

- Manifest: [`v1.20.0/`](v1.20.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.20.0/sbom.json`](v1.20.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `6ec1141` (`v1.20.0`). Binaries are AGPL-3.0-or-later.

## v1.21.0

- Manifest: [`v1.21.0/`](v1.21.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.21.0/sbom.json`](v1.21.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `724683e` (`v1.21.0`). Binaries are AGPL-3.0-or-later.

## v1.22.0

- Manifest: [`v1.22.0/`](v1.22.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.22.0/sbom.json`](v1.22.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `89d3a72` (`v1.22.0`). Binaries are AGPL-3.0-or-later.

## v1.23.0

- Manifest: [`v1.23.0/`](v1.23.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.23.0/sbom.json`](v1.23.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `1effa6e` (`v1.23.0`). Binaries are AGPL-3.0-or-later.

## v1.24.0

- Manifest: [`v1.24.0/`](v1.24.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.24.0/sbom.json`](v1.24.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `e25fc21` (`v1.24.0`). Binaries are AGPL-3.0-or-later.

## v1.25.0

- Manifest: [`v1.25.0/`](v1.25.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.25.0/sbom.json`](v1.25.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `0c551d0` (`v1.25.0`). Binaries are AGPL-3.0-or-later.

## v1.26.0

- Manifest: [`v1.26.0/`](v1.26.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.26.0/sbom.json`](v1.26.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `5490908` (`v1.26.0`). Binaries are AGPL-3.0-or-later.

## v1.27.0

- Manifest: [`v1.27.0/`](v1.27.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.27.0/sbom.json`](v1.27.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `3563343` (`v1.27.0`). Binaries are AGPL-3.0-or-later.

## v1.28.0

- Manifest: [`v1.28.0/`](v1.28.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.28.0/sbom.json`](v1.28.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `3c727e7` (`v1.28.0`). Binaries are AGPL-3.0-or-later.

## v1.29.0

- Manifest: [`v1.29.0/`](v1.29.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.29.0/sbom.json`](v1.29.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `fa827e5` (`v1.29.0`). Binaries are AGPL-3.0-or-later.

## v1.30.0

- Manifest: [`v1.30.0/`](v1.30.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.30.0/sbom.json`](v1.30.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `12326a3` (`v1.30.0`). Binaries are AGPL-3.0-or-later.

## v1.31.0

- Manifest: [`v1.31.0/`](v1.31.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.31.0/sbom.json`](v1.31.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `d994003` (`v1.31.0`). Binaries are AGPL-3.0-or-later.

## v1.32.0

- Manifest: [`v1.32.0/`](v1.32.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.32.0/sbom.json`](v1.32.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `2293440` (`v1.32.0`). Binaries are AGPL-3.0-or-later.

## v1.33.0

- Manifest: [`v1.33.0/`](v1.33.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.33.0/sbom.json`](v1.33.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `f1d1aca` (`v1.33.0`). Binaries are AGPL-3.0-or-later.

## v1.34.0

- Manifest: [`v1.34.0/`](v1.34.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.34.0/sbom.json`](v1.34.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `9706512` (`v1.34.0`). Binaries are AGPL-3.0-or-later.

## v1.35.0

- Manifest: [`v1.35.0/`](v1.35.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.35.0/sbom.json`](v1.35.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `e66b087` (`v1.35.0`). Binaries are AGPL-3.0-or-later.

## v1.36.0

- Manifest: [`v1.36.0/`](v1.36.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.36.0/sbom.json`](v1.36.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `0194bac` (`v1.36.0`). Binaries are AGPL-3.0-or-later.

## v1.37.0

- Manifest: [`v1.37.0/`](v1.37.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.37.0/sbom.json`](v1.37.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `4181b33` (`v1.37.0`). Binaries are AGPL-3.0-or-later.

## v1.38.0

- Manifest: [`v1.38.0/`](v1.38.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.38.0/sbom.json`](v1.38.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `32b032c` (`v1.38.0`). Binaries are AGPL-3.0-or-later.

## v1.39.0

- Manifest: [`v1.39.0/`](v1.39.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.39.0/sbom.json`](v1.39.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `158658f` (`v1.39.0`). Binaries are AGPL-3.0-or-later.

## v1.40.0

- Manifest: [`v1.40.0/`](v1.40.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.40.0/sbom.json`](v1.40.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `68daca2` (`v1.40.0`). Binaries are AGPL-3.0-or-later.

## v1.41.0

- Manifest: [`v1.41.0/`](v1.41.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.41.0/sbom.json`](v1.41.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `0115a56` (`v1.41.0`). Binaries are AGPL-3.0-or-later.

## v1.42.0

- Manifest: [`v1.42.0/`](v1.42.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.42.0/sbom.json`](v1.42.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `2638b67` (`v1.42.0`). Binaries are AGPL-3.0-or-later.

## v1.43.0

- Manifest: [`v1.43.0/`](v1.43.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.43.0/sbom.json`](v1.43.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `8493e09` (`v1.43.0`). Binaries are AGPL-3.0-or-later.

## v1.44.0

- Manifest: [`v1.44.0/`](v1.44.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.44.0/sbom.json`](v1.44.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `469baad` (`v1.44.0`). Binaries are AGPL-3.0-or-later.

## v1.45.0

- Manifest: [`v1.45.0/`](v1.45.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.45.0/sbom.json`](v1.45.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `796912f` (`v1.45.0`). Binaries are AGPL-3.0-or-later.

## v1.46.0

- Manifest: [`v1.46.0/`](v1.46.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.46.0/sbom.json`](v1.46.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `fdc4914` (`v1.46.0`). Binaries are AGPL-3.0-or-later.

## v1.47.0

- Manifest: [`v1.47.0/`](v1.47.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.47.0/sbom.json`](v1.47.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `fefe33d` (`v1.47.0`). Binaries are AGPL-3.0-or-later.

## v1.48.0

- Manifest: [`v1.48.0/`](v1.48.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.48.0/sbom.json`](v1.48.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `a3425a9` (`v1.48.0`). Binaries are AGPL-3.0-or-later.

## v1.49.0

- Manifest: [`v1.49.0/`](v1.49.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.49.0/sbom.json`](v1.49.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `30b9475` (`v1.49.0`). Binaries are AGPL-3.0-or-later.

## v1.50.0

- Manifest: [`v1.50.0/`](v1.50.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.50.0/sbom.json`](v1.50.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `957929e` (`v1.50.0`). Binaries are AGPL-3.0-or-later.

## v1.51.0

- Manifest: [`v1.51.0/`](v1.51.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.51.0/sbom.json`](v1.51.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `3a704f7` (`v1.51.0`). Binaries are AGPL-3.0-or-later.

## v1.52.0

- Manifest: [`v1.52.0/`](v1.52.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.52.0/sbom.json`](v1.52.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `c15ce5b` (`v1.52.0`). Binaries are AGPL-3.0-or-later.

## v1.53.0

- Manifest: [`v1.53.0/`](v1.53.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.53.0/sbom.json`](v1.53.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `88c9479` (`v1.53.0`). Binaries are AGPL-3.0-or-later.

## v1.54.0

- Manifest: [`v1.54.0/`](v1.54.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.54.0/sbom.json`](v1.54.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `189d10c` (`v1.54.0`). Binaries are AGPL-3.0-or-later.

## v1.55.0

- Manifest: [`v1.55.0/`](v1.55.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.55.0/sbom.json`](v1.55.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `4d9b2ed` (`v1.55.0`). Binaries are AGPL-3.0-or-later.

## v1.56.0

- Manifest: [`v1.56.0/`](v1.56.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.56.0/sbom.json`](v1.56.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `f8ce70f` (`v1.56.0`). Binaries are AGPL-3.0-or-later.

## v1.57.0

- Manifest: [`v1.57.0/`](v1.57.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.57.0/sbom.json`](v1.57.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `af094ba` (`v1.57.0`). Binaries are AGPL-3.0-or-later.


## v1.58.0

- Manifest: [`v1.58.0/`](v1.58.0/)
- Binaries: `orchadyn-api`, `orchadyn-mcp`, and `orchadyn-migrate` for
  Windows/amd64, linux (amd64 + arm64), and darwin (amd64 + arm64).
- SBOM: [`v1.58.0/sbom.json`](v1.58.0/sbom.json)
- Source: [ORCHADYN core](https://github.com/axisrobo/ORCHADYN) at commit
  `b1ba175` (`v1.58.0`). Binaries are AGPL-3.0-or-later.

