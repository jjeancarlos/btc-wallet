# btc-wallet (Cold Storage Generator)

A minimalist, security-focused, **offline** CLI tool written in Go to generate Bitcoin wallets.

It strictly follows **BIP-39** (Mnemonic), **BIP-32** (HD Wallets), and **BIP-84** (Native SegWit) standards, using cryptographically strong entropy provided by the operating system via Go’s `crypto/rand` package.

> ⚠️ **SECURITY WARNING**
>
> This tool is designed to be run on an **OFFLINE (air-gapped)** computer.
> **Do not generate real funds on a machine connected to the internet.**
>
> 👉 **[Read the Security & OpSec Guide](GUIDE.md)** to learn how to prepare a secure environment
> (Tails OS, air-gapped machines, physical backups, etc).
>
> ⚠️ **Version Notice**
>
> Releases **prior to v1.1.0** do **NOT** support BIP-39 passphrases and may display
> incorrect master fingerprint information.
>
> **Do not use older releases to generate wallets intended to hold real funds.**
> Always use the latest release and verify signatures before execution.

---

## 🔒 Security Features

* **Zero External APIs**  
  No data is ever sent to the network.

* **Internet Connectivity Warning**  
  Detects common outbound network connectivity and warns the user before proceeding.

* **Best-effort Memory Wiping**  
  Attempts to overwrite sensitive byte slices (entropy, seed, and passphrase buffers)
  immediately after use.  
  Due to Go runtime behavior and garbage collection semantics, complete memory erasure
  **cannot be guaranteed** and must not be relied upon as the sole protection mechanism.

* **Strong Entropy**  
  Uses **256-bit (32 bytes)** entropy generated via Go’s `crypto/rand`,
  backed by the operating system’s CSPRNG.

* **Optional BIP-39 Passphrase Support**  
  Supports an additional BIP-39 passphrase to protect against physical seed exposure
  and provide plausible deniability.

* **Bitcoin Standards Compliance**
  * **BIP-39**: 24-word mnemonic seed phrase
  * **BIP-32**: Hierarchical Deterministic wallet structure
  * **BIP-84**: Native SegWit addresses (`bc1q...`) for modern wallet compatibility and lower fees

---

### ⚠️ Important: BIP-39 Passphrase Compatibility

If you choose to use a BIP-39 passphrase, be aware that:

- A mnemonic seed **with** a passphrase and the same seed **without** a passphrase
  generate **completely different wallets**
- Restoring the 24-word seed **without the correct passphrase** will result in
  a valid but **empty wallet**

Some wallet software and hardware devices:
- Do not clearly prompt for a BIP-39 passphrase during restoration
- Label it as “optional”, “advanced”, or “25th word”
- Default to **no passphrase** if the user does not explicitly enter one

**Always ensure that the wallet used for recovery fully supports BIP-39 passphrases
and that the passphrase is entered exactly as originally used.**

Failure to do so does **not** mean your funds are lost — it usually means the wallet
was restored incorrectly.


## 🛠 Dependencies

This project relies only on well-known, battle-tested libraries commonly used
in the Bitcoin ecosystem:

* [`github.com/btcsuite/btcd`](https://github.com/btcsuite/btcd)  
  Reference Bitcoin protocol implementation in Go.

* [`github.com/tyler-smith/go-bip39`](https://github.com/tyler-smith/go-bip39)  
  Industry-standard BIP-39 mnemonic generation library.

* [`golang.org/x/term`](https://pkg.go.dev/golang.org/x/term)  
  Secure terminal input used to read passphrases without echoing.

---

## 🚀 How to Build & Run

### Prerequisites

* Go **1.20+**

---

### 1. Clone the repository

```bash
git clone https://github.com/jjeancarlos/btc-wallet.git
```

```bash
cd btc-wallet
````

---

### 2. Verify Dependency Integrity (Optional but Recommended)

Inspect `go.sum` to ensure dependencies have not been tampered with:

```bash
cat go.sum
```

---

### 3. Build the Binary

**Linux (recommended for Live USBs such as Tails or Ubuntu Live):**

```bash
GOOS=linux GOARCH=amd64 go build -o btc-wallet main.go
```

**Windows:**

```bash
GOOS=windows GOARCH=amd64 go build -o btc-wallet.exe main.go
```

**macOS (Apple Silicon):**

```bash
GOOS=darwin GOARCH=arm64 go build -o btc-wallet main.go
```

---

### 4. Run (Offline)

Transfer the compiled binary to your offline machine via USB and run:

```bash
./btc-wallet
```

## 🛡️ Verifying Authenticity (Don’t Trust, Verify)

To ensure the binaries were not tampered with and were produced by the original
developer, follow these steps.

### 1. Download the release files

From the [Releases page](https://github.com/jjeancarlos/btc-wallet/releases), download:

* The binary for your OS
* `SHA256SUMS.txt`
* `SHA256SUMS.txt.asc`

---

### 2. Verify the Checksum (Integrity)

**Linux / macOS:**

```bash
shasum -a 256 -c SHA256SUMS.txt --ignore-missing
```

Expected output:

```
btc-wallet-linux-amd64: OK
```

**Windows (PowerShell):**

```powershell
Get-FileHash .\btc-wallet-windows-amd64.exe -Algorithm SHA256
```

Manually compare the output with the hash listed in `SHA256SUMS.txt`.

---

### 3. Verify the Signature (Authenticity)

Import the developer’s public GPG key:

```bash
gpg --keyserver keyserver.ubuntu.com --recv-keys FE125B66F25875A56B129BE8FB9E2F51656B1941
```

Verify the signature:

```bash
gpg --verify SHA256SUMS.txt.asc SHA256SUMS.txt
```

Expected output:

```
Good signature from "jjeancarlos <jeanpastebin@gmail.com>"
```


## 📝 Usage Flow

1. The tool checks for active internet connectivity.
2. If offline (or explicitly approved), it generates **256-bit entropy**.
3. The following information is displayed:

   * **24-word BIP-39 mnemonic seed** (write it down securely)
   * **Optional BIP-39 passphrase prompt**
   * **First receiving address** (Native SegWit)
   * **Master fingerprint** (BIP-32 compliant)
4. On exit, the tool clears the terminal and attempts best-effort memory cleanup.


## 🎯 Threat Model

This tool is designed to protect against:

* Remote attackers
* Network-based data exfiltration
* Accidental key reuse

It **does not** protect against:

* Compromised firmware, BIOS, or hardware
* Malicious Go compiler or runtime
* Physical attackers observing the key generation process


## ⚠️ Virtual Machines

Generating wallets inside virtual machines is discouraged unless the VM provides
strong entropy guarantees and has **no shared clipboard, folders, or devices**
enabled.


## 🐛 Reporting Vulnerabilities

If you discover a security vulnerability, please **do not open a public issue**.

Refer to the [Security Policy](SECURITY.md) for instructions on responsible
disclosure using PGP encryption.

## 📄 License

This project is licensed under the MIT License.
See the [LICENSE](LICENSE) file for details.

**Disclaimer:**
This software is provided *“as is”*, without warranty of any kind.
The user is solely responsible for the secure storage of generated keys and backups.