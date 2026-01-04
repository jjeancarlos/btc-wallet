# 🛡️ Cold Storage Security Guide

Generating a seed is only the first step. How you handle that seed determines the safety of your funds.

This guide outlines the **Operational Security (OpSec)** best practices for using `btc-wallet`.

## 1. The Environment (Physical Security)

Before running the software, secure your physical space:
* **No Cameras:** Cover all webcams on the device. Ensure no CCTV or smartphone cameras are pointing at your screen.
* **No Smart Devices:** Remove Amazon Echo, Google Home, or Siri-enabled devices from the room (they have microphones).
* **Privacy:** Close windows and curtains. Ensure you are alone.

## 2. The Operating System (The "Clean" Machine)

Never generate a cold wallet on your day-to-day Windows or macOS installation, as they may have hidden malware or keyloggers.

### 🥇 Gold Standard: Tails OS (Recommended)
[Tails](https://tails.boum.org/) is a live operating system that runs from a USB stick and leaves no trace on the computer (it wipes RAM on shutdown).
1.  Flash Tails to a USB drive.
2.  Boot your computer from the USB.
3.  **Disable Networking** immediately (or physically remove the Wi-Fi card/Ethernet cable).
4.  Run `btc-wallet`.

### 🥈 Silver Standard: Qubes OS
[Qubes OS](https://www.qubes-os.org/) isolates programs in virtual machines. Use a completely offline AppVM (Vault) to run the binary.

### 🥉 Bronze Standard: Ubuntu Live USB
Similar to Tails, but ensure you do not connect to Wi-Fi during the session.

## 3. The Generation Process (Air-Gap)

**"Air-gapped"** means the computer is physically isolated from any network.

1.  Download `btc-wallet` on an online computer.
2.  Verify the GPG signature and Hash (see README).
3.  Copy the binary to a USB drive.
4.  Move the USB to the **Offline Machine**.
5.  Run the tool, write down the seed, and verify the addresses.
6.  **Shutdown:** Turn off the computer completely to clear the RAM.

## 4. Storing the Seed Phrase

Your 24 words are your money. If you lose them, the money is gone forever.

* **Paper:** Use acid-free paper and a pencil (ink fades over time).
* **Metal (Best):** Use stainless steel plates (Cryptosteel, Coldcard, etc) to protect against fire and flood.
* **Redundancy:** Make 2 copies. Store them in geographically separate secure locations.
* **NEVER:**
    * Never take a photo of the seed.
    * Never type the seed into a computer connected to the internet.
    * Never save it in a password manager or cloud storage.
    * Never say the words out loud (microphones).

## 5. The Verification Test (Don't skip this!)

Before sending large amounts of Bitcoin to your new wallet:

1.  Generate the wallet.
2.  Write down the seed.
3.  Send a **very small amount** (e.g., $5 worth of BTC) to the generated address.
4.  **Wipe everything** (close the tool, restart the PC).
5.  Try to **restore** the wallet using a different software (like Electrum or Sparrow Wallet) on an offline machine using the seed you wrote down.
6.  If you see your funds, your backup is correct. Now you can transfer the rest.

---
*Stay paranoid. Stay safe.*