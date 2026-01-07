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
	"golang.org/x/term"
)

func main() {
	clearScreen()

	// ---------------------------------------------------------
	// 1. CONNECTIVITY CHECK (WARNING ONLY)
	// ---------------------------------------------------------
	if isOnline() {
		fmt.Println("\033[31m!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println(" WARNING: INTERNET CONNECTION DETECTED!")
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\033[0m")
		fmt.Println("")
		fmt.Println("This tool is designed for OFFLINE (air-gapped) environments.")
		fmt.Println("Generating keys on an online machine increases risk.")
		fmt.Println("")
		fmt.Print("Type 'YES' to continue anyway: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		if strings.TrimSpace(strings.ToUpper(input)) != "YES" {
			fmt.Println("Aborting for safety.")
			return
		}
	} else {
		fmt.Println("\033[32m✔ No internet connection detected (ideal).\033[0m")
	}

	time.Sleep(2 * time.Second)
	clearScreen()

	// ---------------------------------------------------------
	// 2. ENTROPY & MNEMONIC (BIP-39)
	// ---------------------------------------------------------
	entropy := make([]byte, 32) // 256 bits
	_, err := rand.Read(entropy)
	if err != nil {
		panic("failed to read from system CSPRNG")
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		panic(err)
	}

	// Wipe raw entropy ASAP
	zeroize(entropy)

	// ---------------------------------------------------------
	// 3. OPTIONAL BIP-39 PASSPHRASE
	// ---------------------------------------------------------
	fmt.Print("Optional BIP-39 Passphrase (leave empty for none): ")
	passphraseBytes, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()

	passphrase := string(passphraseBytes)

	seed := bip39.NewSeed(mnemonic, passphrase)

	// Wipe passphrase material
	zeroize(passphraseBytes)
	passphrase = ""

	// ---------------------------------------------------------
	// 4. HD WALLET DERIVATION (BIP-84)
	// ---------------------------------------------------------
	params := &chaincfg.MainNetParams

	masterKey, err := hdkeychain.NewMaster(seed, params)
	if err != nil {
		panic(err)
	}

	// Wipe seed ASAP
	zeroize(seed)

	// Path: m/84'/0'/0'/0/0
	purposeKey, err := masterKey.Derive(hdkeychain.HardenedKeyStart + 84)
	if err != nil {
		panic(err)
	}

	coinKey, err := purposeKey.Derive(hdkeychain.HardenedKeyStart + 0)
	if err != nil {
		panic(err)
	}

	accountKey, err := coinKey.Derive(hdkeychain.HardenedKeyStart + 0)
	if err != nil {
		panic(err)
	}

	changeKey, err := accountKey.Derive(0)
	if err != nil {
		panic(err)
	}

	addressKey, err := changeKey.Derive(0)
	if err != nil {
		panic(err)
	}

	// ---------------------------------------------------------
	// 5. ADDRESS GENERATION (Native SegWit)
	// ---------------------------------------------------------
	pubKeyHash, err := addressKey.Address(params)
	if err != nil {
		panic(err)
	}

	witnessAddr, err := btcutil.NewAddressWitnessPubKeyHash(
		pubKeyHash.ScriptAddress(),
		params,
	)
	if err != nil {
		panic(err)
	}

	// ---------------------------------------------------------
	// 6. MASTER FINGERPRINT (BIP-32 CORRECT)
	// ---------------------------------------------------------
	masterPub, err := masterKey.Neuter()
	if err != nil {
		panic(err)
	}

	ecPub, err := masterPub.ECPubKey()
	if err != nil {
		panic(err)
	}

	fp := btcutil.Hash160(ecPub.SerializeCompressed())[:4]

	// ---------------------------------------------------------
	// 7. DISPLAY RESULTS
	// ---------------------------------------------------------
	clearScreen()
	fmt.Println("========================================")
	fmt.Println("   BITCOIN COLD WALLET (BIP-84)          ")
	fmt.Println("========================================")
	fmt.Println("")

	words := strings.Split(mnemonic, " ")
	for i, word := range words {
		fmt.Printf("%2d. %-10s ", i+1, word)
		if (i+1)%4 == 0 {
			fmt.Println("")
		}
	}

	fmt.Println("\n----------------------------------------")
	fmt.Println("RECEIVE ADDRESS (Native SegWit):")
	fmt.Printf("\033[33m%s\033[0m\n", witnessAddr.EncodeAddress())
	fmt.Println("----------------------------------------")

	fmt.Printf("Master Fingerprint: %02x%02x%02x%02x\n",
		fp[0], fp[1], fp[2], fp[3],
	)

	fmt.Println("\n⚠️  WRITE EVERYTHING DOWN ON PAPER.")
	fmt.Println("Press ENTER to wipe memory and exit.")

	bufio.NewReader(os.Stdin).ReadString('\n')

	// ---------------------------------------------------------
	// 8. BEST-EFFORT MEMORY CLEANUP
	// ---------------------------------------------------------
	clearScreen()

	mnemonic = ""
	words = nil
	masterKey = nil
	runtime.GC()

	fmt.Println("✔ Memory cleanup attempted (best effort).")
	fmt.Println("✔ Terminal cleared.")
	fmt.Println("Goodbye.")
}

// ---------------------------------------------------------
// HELPERS
// ---------------------------------------------------------

func isOnline() bool {
	conn, err := net.DialTimeout("tcp", "8.8.8.8:53", 2*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func zeroize(b []byte) {
	for i := range b {
		b[i] = 0
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
