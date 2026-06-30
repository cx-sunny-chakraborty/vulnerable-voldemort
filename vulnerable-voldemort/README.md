# vulnerable-voldemort 🐍

> *"There is no good and evil, there is only insecure code and those too lazy to scan it."*

A deliberately vulnerable Go application — themed around He-Who-Must-Not-Be-Named — built **for security-scanning demos only**.

## ⚠️ Warning
This repository is **intentionally insecure**. Every handler contains a well-known, documented vulnerability pattern so that a SAST scan produces a predictable set of findings (great for demoing Feedback Apps, issue creation, and triage). 

**Do not deploy this. Do not reuse any of this code.** It exists to be flagged, not to run anything real.

## What's inside (the Dark Arts on display)
| Spell | Vulnerability | CWE |
|-------|---------------|-----|
| Hardcoded credentials | Secrets in source | CWE-798 |
| Find a wizard | SQL Injection | CWE-89 |
| Send an owl | OS Command Injection | CWE-78 |
| Restricted Section | Path Traversal | CWE-22 |
| Marauder's hashing | Weak crypto (MD5) | CWE-327 |
| Portkey token | Insecure randomness | CWE-338 |
| Greet a house | Reflected XSS | CWE-79 |
| Fetch a prophecy | SSRF | CWE-918 |
| Reveal secrets | Information exposure | CWE-200 |
| Ministry server | Cleartext transmission | CWE-319 |
| Destroy a horcrux | SQL Injection (Sprintf) | CWE-89 |
| Disapparate | Open redirect | CWE-601 |
| Cast a spell | Template injection | CWE-79 |
| Chamber of Secrets | Improper auth check | CWE-703 |

## Run (if you must)
```
go run .
```
The Ministry will fall and listen on `:8080`. Then point Checkmarx at it and watch the issues appear.
