# Security Policy

We take the security of `btc-wallet` seriously. Given the nature of this software (handling private keys), we encourage security researchers to report vulnerabilities to us in a responsible manner.

## Supported Versions

Only the latest stable release is supported.

| Version | Supported          |
| ------- | ------------------ |
| 1.0.x   | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

**DO NOT** create a public GitHub issue for security vulnerabilities. Publicly disclosing a vulnerability could put user funds at risk before a patch is available.

### How to Report

Please send an email to **jeanpastebin@gmail.com**.

### 🔐 Encrypted Communication (Recommended)

Since reports may contain sensitive exploits, we highly recommend encrypting your email using our PGP/GPG Key.

* **Key ID:** `FE125B66F25875A56B129BE8FB9E2F51656B1941`
* **Fingerprint:** `FE12 5B66 F258 75A5 6B12 9BE8 FB9E 2F51 656B 1941`
* **Public Key Server:** [keyserver.ubuntu.com](https://keyserver.ubuntu.com/pks/lookup?search=0xFE125B66F25875A56B129BE8FB9E2F51656B1941&op=index)

### What to Include

1.  A description of the vulnerability.
2.  Steps to reproduce the issue (PoC - Proof of Concept).
3.  Potential impact of the vulnerability.

### Our Response Process

1.  We will acknowledge receipt of your report within **48 hours**.
2.  We will investigate the issue and determine its severity.
3.  We will communicate a timeline for the fix.
4.  Once the fix is released, we will publicly acknowledge your contribution (unless you prefer to remain anonymous).

## Out of Scope

The following are generally considered out of scope, unless they lead to direct compromise of the seed generation:

* Vulnerabilities requiring physical access to the user's device (Evil Maid attacks).
* Social engineering attacks against the user.
* Denial of Service (DoS) attacks.
* Issues on the user's operating system (malware, keyloggers).

---

**Thank you for helping keep the Bitcoin ecosystem safe.**