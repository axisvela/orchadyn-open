#!/usr/bin/env python3
import os
import re
import subprocess
import sys
from pathlib import Path

blocked = {
    item.strip().lower()
    for item in os.environ.get("PRAXO_DISALLOWED_EMAIL_DOMAINS", "").split(",")
    if item.strip()
}
if not blocked:
    print("PRAXO_DISALLOWED_EMAIL_DOMAINS is not configured", file=sys.stderr)
    sys.exit(1)

email_re = re.compile(r"[A-Za-z0-9._%+\-]+@([A-Za-z0-9.\-]+\.[A-Za-z]{2,})")
binary_ext = {
    ".exe", ".dll", ".so", ".dylib", ".a", ".o", ".class", ".jar",
    ".png", ".jpg", ".jpeg", ".gif", ".ico", ".pdf", ".zip", ".gz",
    ".woff", ".woff2", ".ttf", ".eot", ".pyc", ".bin",
}

errors = []

history = subprocess.run(
    ["git", "log", "--all", "--format=%H%x09%ae%x09%ce"],
    capture_output=True, text=True, check=True,
).stdout
for line in history.splitlines():
    fields = line.split("\t")
    if len(fields) != 3:
        continue
    sha = fields[0]
    for email in fields[1:]:
        if "@" not in email:
            continue
        domain = email.rsplit("@", 1)[1].lower()
        if domain in blocked:
            errors.append(f"{sha}: commit identity uses blocked email domain {domain}")

tracked = subprocess.run(
    ["git", "ls-files", "-z"], capture_output=True, check=True,
).stdout
for raw in tracked.split(b"\0"):
    if not raw:
        continue
    path = Path(raw.decode("utf-8", "surrogateescape"))
    if path.suffix.lower() in binary_ext:
        continue
    try:
        if path.stat().st_size > 2 * 1024 * 1024:
            continue
        text = path.read_text(encoding="utf-8", errors="ignore")
    except OSError:
        continue
    for match in email_re.finditer(text):
        domain = match.group(1).lower()
        if domain in blocked:
            line = text.count("\n", 0, match.start()) + 1
            errors.append(f"{path}:{line}: email domain {domain} is blocked")

if errors:
    for error in errors:
        print(f"::error::{error}")
    sys.exit(1)
print("privacy check passed")
