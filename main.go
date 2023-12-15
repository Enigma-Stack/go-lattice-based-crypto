package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "go-lattice-based-crypto/crypto"
)

func main() {
    generateKeysCmd := flag.NewFlagSet("generate-keys", flag.ExitOnError)
    encryptCmd := flag.NewFlagSet("encrypt", flag.ExitOnError)
    decryptCmd := flag.NewFlagSet("decrypt", flag.ExitOnError)

    // Flags for generate-keys command
    skFilePath := generateKeysCmd.String("sk", "", "Path to save the secret key (required)")
    pkFilePath := generateKeysCmd.String("pk", "", "Path to save the public key (required)")

    // Flags for encrypt command
    encryptValues := encryptCmd.String("values", "", "Values to encrypt (comma-separated, required)")
    encryptPkFilePath := encryptCmd.String("pk", "", "Path to the secret key (required)")
    ctFilePath := encryptCmd.String("ct", "", "Path to save the ciphertext (required)")

    // Flags for decrypt command
    decryptCtFilePath := decryptCmd.String("ct", "", "Path to the ciphertext (required)")
    decryptSkFilePath := decryptCmd.String("sk", "", "Path to the secret key (required)")
    decryptNumValues := decryptCmd.Int("numvalues", 0, "Number of values that were encrypted (required)")

    if len(os.Args) < 2 {
        printUsage()
        os.Exit(1)
    }

    switch os.Args[1] {
    case "generate-keys":
        handleCommand(generateKeysCmd, func() {
            validateFilePaths(*skFilePath, *pkFilePath)
            crypto.GenerateKeyPair(*skFilePath, *pkFilePath)
        })
    case "encrypt":
        handleCommand(encryptCmd, func() {
            validateFilePaths(*encryptPkFilePath, *ctFilePath)
            if *encryptValues == "" {
                log.Fatal("Values to encrypt not provided")
            }
            // values := utils.ParseValues(*encryptValues)
            crypto.EncryptData(*encryptValues, *encryptPkFilePath, *ctFilePath)
        })
    case "decrypt":
        handleCommand(decryptCmd, func() {
            validateFilePaths(*decryptCtFilePath, *decryptSkFilePath)
            if *decryptNumValues <= 0 {
                log.Fatal("Invalid number of values for decryption")
            }
            decryptedValues := crypto.DecryptData(*decryptCtFilePath, *decryptSkFilePath, *decryptNumValues)
            fmt.Println("Decrypted Values:", decryptedValues)
        })
    default:
        printUsage()
        os.Exit(1)
    }
}

func handleCommand(f *flag.FlagSet, action func()) {
    err := f.Parse(os.Args[2:])
    if err != nil {
        log.Fatalf("Error parsing flags: %v", err)
    }
    action()
}

func validateFilePaths(paths ...string) {
    for _, path := range paths {
        if path == "" {
            log.Fatal("Required file path not provided")
        }
    }
}

func printUsage() {
    fmt.Println("Usage: <command> [options]")
    fmt.Println("Commands: generate-keys, encrypt, decrypt")
    fmt.Println("Use -h flag with any command to see its options")
}
