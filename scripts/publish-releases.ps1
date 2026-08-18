[CmdletBinding()]
param(
    [string]$From = "v1.23.0",
    [string]$To = "v1.57.0",
    [string]$Repo = "axisrobo/ORCHADYN-open",
    [switch]$DryRun
)

# Promotes staged release manifests to GitHub Release assets. GitHub requires
# unique asset names, so each platform's files are staged under platform-
# prefixed names before upload. Versions that already have a GitHub release
# are skipped. Requires an authenticated `gh` CLI.
#
# Usage after the release manifests are staged in releases/:
#   .\scripts\publish-releases.ps1 -From v1.55.0 -To v1.57.0

$ErrorActionPreference = "Stop"
$repoRoot = Split-Path -Parent $PSScriptRoot

function Get-Attestation {
    param([string]$Version)
    $notice = Get-ChildItem (Join-Path $repoRoot "releases/$Version") -Recurse -Filter NOTICE -ErrorAction SilentlyContinue | Select-Object -First 1
    if ($notice) {
        return [regex]::Match((Get-Content -Raw $notice.FullName), 'commit ([0-9a-f]+)').Groups[1].Value
    }
    return ""
}

$fromVersion = [version]$From.TrimStart("v")
$toVersion = [version]$To.TrimStart("v")
$pending = @()
foreach ($dir in (Get-ChildItem (Join-Path $repoRoot "releases") -Directory)) {
    if ($dir.Name -notmatch '^v(\d+)\.(\d+)\.(\d+)$') {
        continue
    }
    $version = [version]$dir.Name.TrimStart("v")
    if ($version -lt $fromVersion -or $version -gt $toVersion) {
        continue
    }
    if (-not (Test-Path (Join-Path $dir.FullName "sbom.json"))) {
        continue
    }
    $null = cmd /c "gh release view $($dir.Name) --repo $Repo --json tagName 2>nul"
    if ($LASTEXITCODE -ne 0) {
        $pending += $dir.Name
    }
}
$pending = $pending | Sort-Object { [version]$_.TrimStart("v") }
if ($pending.Count -eq 0) {
    Write-Host "No unreleased manifests between $From and $To."
    return
}
Write-Host "Publishing $($pending.Count) releases: $($pending -join ', ')"

$stage = Join-Path ([System.IO.Path]::GetTempPath()) "orchadyn-open-release-assets"
foreach ($version in $pending) {
    $base = Join-Path $repoRoot "releases/$version"
    $stageDir = Join-Path $stage $version
    Remove-Item -Recurse -Force $stageDir -ErrorAction SilentlyContinue
    New-Item -ItemType Directory -Force $stageDir | Out-Null
    Get-ChildItem $base -Recurse -File | ForEach-Object {
        $rel = $_.FullName.Substring($base.Length).TrimStart('\')
        Copy-Item $_.FullName (Join-Path $stageDir ($rel -replace '[\\/]', '-')) -Force
    }
    $commit = Get-Attestation $version
    if ($DryRun) {
        Write-Host "[dry-run] would create $version with $((Get-ChildItem $stageDir -File).Count) assets"
        continue
    }
    Write-Host "Creating release $version..."
    gh release create $version --repo $Repo --title "ORCHADYN $version" --notes "Staged manifest $version. Source: ORCHADYN core commit $commit. Binaries are AGPL-3.0-or-later." 2>$null
    if ($LASTEXITCODE -ne 0) {
        throw "release create failed for $version"
    }
    $uploadArgs = @("release", "upload", $version, "--repo", $Repo)
    $uploadArgs += (Get-ChildItem $stageDir -File | Select-Object -ExpandProperty FullName)
    gh @uploadArgs 2>$null
    if ($LASTEXITCODE -ne 0) {
        throw "release upload failed for $version"
    }
    Write-Host "Published $version."
}
Remove-Item -Recurse -Force $stage -ErrorAction SilentlyContinue
Write-Host "Publish complete: $($pending.Count) releases."
