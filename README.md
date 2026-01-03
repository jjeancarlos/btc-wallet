# btc-wallet (Cold Storage Generator)

A minimalist, high-security, offline CLI tool written in Go to generate Bitcoin wallets.

It strictly follows **BIP-39** (Mnemonic) and **BIP-84** (Native SegWit) standards, using cryptographically strong entropy from the operating system (`crypto/rand`).

> ⚠️ **SECURITY WARNING**: This tool is designed to be run on an **OFFLINE (Air-gapped)** computer. Do not generate real funds on a machine connected to the internet.

## 🔒 Security Features

* **Zero Dependencies on External APIs**: No data is ever sent to the network.
* **Internet Kill Switch**: Detects active connections and warns the user before proceeding.
* **Memory Wiping**: Attempts to zeroize entropy and seed variables from memory immediately after use.
* **Strong Entropy**: Uses 256-bit (32 bytes) entropy via Go's `crypto/rand` (OS CSPRNG).
* **Standard Compliance**:
    * **BIP-39**: 24-word mnemonic phrase.
    * **BIP-84**: Modern Native SegWit addresses (`bc1q...`) for lower fees.
    * **BIP-32**: Hierarchical Deterministic wallet structure.

## 🛠 Dependencies

This project uses battle-tested libraries standard in the Bitcoin ecosystem:

* [`github.com/btcsuite/btcd`](https://github.com/btcsuite/btcd): The reference implementation of Bitcoin in Go.
* [`github.com/tyler-smith/go-bip39`](https://github.com/tyler-smith/go-bip39): Industry standard for mnemonic generation.

## 🚀 How to Build & Run

### Prerequisites
* Go 1.20+ installed.

### 1. Clone the repository
```bash
git clone https://github.com/jjeancarlos/btc-wallet.git
```
```bash
cd btc-wallet
```

### 2. Verify Integrity (Optional but Recommended)

Check `go.sum` to ensure dependencies haven't been tampered with.

```bash
cat go.sum

```

### 3. Build the Binary

**For Linux (Recommended for Live USBs like Tails/Ubuntu):**

```bash
GOOS=linux GOARCH=amd64 go build -o btc-wallet main.go

```

**For Windows:**

```bash
GOOS=windows GOARCH=amd64 go build -o btc-wallet.exe main.go

```

**For macOS (Apple Silicon):**

```bash
GOOS=darwin GOARCH=arm64 go build -o btc-wallet main.go

```

### 4. Run (Offline)

Move the binary to your offline machine via USB and run:

```bash
./btc-wallet

```

## 🛡️ Verifying Authenticity (Don't Trust, Verify)

To guarantee that the binaries you downloaded have not been tampered with and were created by the original developer, follow these steps:

### 1. Download the files

Download the binary for your OS, plus `SHA256SUMS.txt` and `SHA256SUMS.txt.asc` from the [suspicious link removed].

### 2. Verify the Checksum (Integrity)

Open your terminal in the download folder and run:

**Linux / macOS:**

```bash
# This command checks if the file matches the hash list
shasum -a 256 -c SHA256SUMS.txt --ignore-missing

```

*Expected Output:* `btc-wallet-linux-amd64: OK`

**Windows (PowerShell):**

```powershell
Get-FileHash .\btc-wallet-windows-amd64.exe -Algorithm SHA256
# Compare the output hash manually with the one in SHA256SUMS.txt

```

### 3. Verify the Signature (Authenticity)

This ensures the `SHA256SUMS.txt` file was signed by the developer's GPG Key.

First, import the developer's public key (if you haven't already):

```bash
gpg --keyserver keyserver.ubuntu.com --recv-keys FE125B66F25875A56B129BE8FB9E2F51656B1941

```

Then verify the signature:

```bash
gpg --verify SHA256SUMS.txt.asc SHA256SUMS.txt

```

*Expected Output:* `Good signature from "jjeancarlos <jeanpastebin@gmail.com>"`

## 📝 Usage Flow

1. The tool checks for internet connection.
2. If offline, it generates 256-bit entropy.
3. It displays:
* **24-Word Seed Phrase** (Write this down!)
* **First Receiving Address** (Native SegWit)
* **Master Fingerprint**


4. Upon exit, it attempts to clear the screen and wipe memory variables.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**Disclaimer:** This software is provided "as is", without warranty of any kind. The user is solely responsible for the safe storage of the generated keys.

```

```