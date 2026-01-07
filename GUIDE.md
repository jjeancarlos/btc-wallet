# 🛡️ Cold Storage Security Guide

Generating a seed is only the first step.
**How you handle that seed determines the safety of your funds.**

This guide outlines **Operational Security (OpSec)** best practices for safely using `btc-wallet`.

---

## 1. The Environment (Physical Security)

Before running the software, secure your physical environment:

* **No Cameras**
  Cover all webcams on the device. Ensure no CCTV or smartphone cameras can see your screen.

* **No Smart Devices**
  Remove Amazon Echo, Google Home, Siri-enabled devices, or anything with microphones from the room.

* **Privacy**
  Close windows and curtains. Ensure you are alone and not being observed.

---

## 2. The Operating System (The “Clean” Machine)

Never generate a cold wallet on your daily-use Windows or macOS installation.
These systems may contain malware, keyloggers, or unknown background services.

### 🥇 Gold Standard: Tails OS (Recommended)

[Tails](https://tails.boum.org/) is a live operating system that runs from a USB stick and leaves no trace on the computer (RAM is wiped on shutdown).

1. Flash Tails to a USB drive.
2. Boot the computer from the USB.
3. **Disable networking immediately** (or physically remove Wi-Fi and Ethernet).
4. Run `btc-wallet`.

---

### 🥈 Silver Standard: Qubes OS

[Qubes OS](https://www.qubes-os.org/) isolates applications into virtual machines.

* Use a **fully offline AppVM (Vault)**.
* Ensure no clipboard, file sharing, or devices are attached.

---

### 🥉 Bronze Standard: Ubuntu Live USB

An Ubuntu Live USB is acceptable if used carefully:

* Never connect to Wi-Fi.
* Do not install additional software.
* Shut down immediately after wallet generation.

---

## 3. The Generation Process (Air-Gap)

An **air-gapped** computer is physically isolated from all networks.

1. Download `btc-wallet` on an **online computer**.
2. Verify the GPG signature and checksum (see README).
3. Copy the binary to a USB drive.
4. Move the USB to the **offline machine**.
5. Run the tool, write down the seed, and verify the generated address.
6. **Shutdown the computer completely** to clear RAM.

---

## 4. Storing the Seed Phrase

Your 24-word seed (and optional passphrase) **is your money**.
If you lose it, your funds are lost forever.

### Recommended Storage Methods

* **Paper**
  Use acid-free paper and a pencil (ink can fade over time).

* **Metal (Best)**
  Stainless steel backups (e.g., Cryptosteel, Coldcard, Seedplate) protect against fire and water damage.

* **Redundancy**
  Make at least **two copies** and store them in geographically separate, secure locations.

### ❌ Never Do This

* Never take photos of the seed.
* Never type the seed into an internet-connected computer.
* Never store it in cloud services or password managers.
* Never say the words out loud (microphones exist everywhere).

---

## 5. The Verification Test (Do Not Skip)

Before sending significant funds:

1. Generate the wallet.
2. Write down the seed (and passphrase, if used).
3. Send a **small test amount** (e.g., ~$5 worth of BTC).
4. Close the tool and **reboot or power off** the machine.
5. Restore the wallet using **different software** (Electrum, Sparrow) on an offline system.
6. Confirm the funds appear.

If successful, your backup is valid.
Only then should you transfer larger amounts.

---

*Stay paranoid. Stay safe.*