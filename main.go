package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	// Clear screen at start
	clearScreen()

	// ---------------------------------------------------------
	// 1. INTERNET CHECKING (NEW)
	// ---------------------------------------------------------
	if isOnline() {
		fmt.Println("\033[31m!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println(" CRITICAL DANGER: INTERNET CONNECTION DETECTED!")
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\033[0m")
		fmt.Println("")
		fmt.Println("This program is designed to run in an OFFLINE environment.")
		fmt.Println("Generating keys while connected exposes you to malware/loggers.")
		fmt.Println("")
		fmt.Println("Recommendation: DISCONNECT THE NETWORK CABLE / WIFI NOW.")
		fmt.Println("")
		fmt.Print("Do you want to continue despite the risk? (type 'YES' to proceed): ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		if strings.TrimSpace(strings.ToUpper(input)) != "YES" {
			fmt.Println("Aborting for safety.")
			return
		}
	} else {
		fmt.Println("\033[32m✔ No internet connection detected (Ideal).\033[0m")
	}

	fmt.Println("\nStarting Cold Wallet generator...")
	time.Sleep(2 * time.Second)
	clearScreen()

	// ---------------------------------------------------------
	// 2. GENERATE ENTROPY & MNEMONIC (BIP-39)
	// ---------------------------------------------------------

	// Entropy: 32 bytes (256 bits)
	entropy := make([]byte, 32)
	_, err := rand.Read(entropy)
	if err != nil {
		panic("Critical failure in system random number generator.")
	}

	// Mnemonic (BIP-39)
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		panic(err)
	}

	// WIPE 1: Raw entropy is no longer needed, wipe immediately.
	zeroize(entropy)

	// Generate Binary Seed (BIP-39)
	// Note: go-bip39 returns []byte, which is great for wiping later.
	seed := bip39.NewSeed(mnemonic, "")

	// ---------------------------------------------------------
	// 3. KEY DERIVATION (BIP-84)
	// ---------------------------------------------------------

	params := &chaincfg.MainNetParams

	masterKey, err := hdkeychain.NewMaster(seed, params)
	if err != nil {
		panic(err)
	}

	// WIPE 2: Master seed is already loaded into key memory, wipe original array.
	zeroize(seed)

	// Path: m/84'/0'/0'/0/0
	purposeKey, _ := masterKey.Derive(hdkeychain.HardenedKeyStart + 84)
	coinKey, _ := purposeKey.Derive(hdkeychain.HardenedKeyStart + 0)
	accountKey, _ := coinKey.Derive(hdkeychain.HardenedKeyStart + 0)
	changeKey, _ := accountKey.Derive(0)
	addressKey, _ := changeKey.Derive(0)

	// Bech32 Address
	pubKeyHash, _ := addressKey.Address(params)
	witnessAddr, _ := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash.ScriptAddress(), params)

	// ---------------------------------------------------------
	// 4. DISPLAY AND CLEANUP
	// ---------------------------------------------------------
	clearScreen()
	fmt.Println("========================================")
	fmt.Println("      BITCOIN WALLET (BIP-84)           ")
	fmt.Println("========================================")
	fmt.Println("")

	// Display Mnemonic
	words := strings.Split(mnemonic, " ")
	for i, word := range words {
		fmt.Printf("%2d. %-10s ", i+1, word)
		if (i+1)%4 == 0 {
			fmt.Println("")
		}
	}

	fmt.Println("\n----------------------------------------")
	fmt.Println("ADDRESS (Receive BTC here):")
	fmt.Printf("\033[33m%s\033[0m\n", witnessAddr.EncodeAddress())
	fmt.Println("----------------------------------------")

	// Fingerprint
	fp, _ := masterKey.Neuter()
	fmt.Printf("Master Fingerprint: %08x\n", fp.ParentFingerprint())

	fmt.Println("\n⚠️  WRITE EVERYTHING DOWN ON PAPER NOW.")
	fmt.Println("Press ENTER to wipe memory and exit.")

	bufio.NewReader(os.Stdin).ReadString('\n')

	// WIPE 3: Attempt to wipe mnemonic from screen and force GC
	clearScreen()

	// Zero out accessible local variables (Best Effort in Go)
	// Note: Strings in Go are immutable, so 'mnemonic' and 'words' will stay
	// in Heap until the Garbage Collector passes. We force GC here.
	mnemonic = ""
	words = nil
	masterKey = nil
	runtime.GC()

	fmt.Println("✔ Memory wiped (best effort).")
	fmt.Println("✔ Terminal cleared.")
	fmt.Println("Goodbye.")
}

// ---------------------------------------------------------
// HELPER FUNCTIONS
// ---------------------------------------------------------

// isOnline attempts to connect to Google DNS (8.8.8.8:53) to check internet.
// Uses a short 2s timeout to avoid hanging if offline.
func isOnline() bool {
	conn, err := net.DialTimeout("tcp", "8.8.8.8:53", 2*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// zeroize overwrites a byte slice with zeros to protect memory
func zeroize(data []byte) {
	for i := range data {
		data[i] = 0
	}
}

// clearScreen visually clears the terminal
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}