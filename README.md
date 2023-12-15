# Go Lattice-Based Crypto

Go Lattice-Based Crypto is a Go package providing an implementation for key generation, encryption, and decryption using lattice-based cryptographic methods. The implementation leverages the Lattigo library for homomorphic encryption.

## Features

- **Key Generation:** Generate public and secret keys using lattice-based cryptographic methods.
- **Encryption:** Encrypt data using the generated public key.
- **Decryption:** Decrypt data using the corresponding secret key.

## Requirements

- Go programming environment
- Lattigo v5 library

## Installation

Before installing the package, ensure that you have a working Go environment. See the [Go installation guide](https://golang.org/doc/install) for details.

To install the Go Lattice-Based Crypto package, run the following command:

```bash
go get -u github.com/Enigma-Stack/go-lattice-based-crypto
```

## Usage
The package provides the following commands:

#### Generate Keys
To generate a public and secret key pair:
```bash
go run main.go generate-keys -sk [path-to-secret-key-file] -pk [path-to-public-key-file]
```

#### Encrypt Data
To encrypt data:

```bash
go run main.go encrypt -values [comma-separated-values] -pk [path-to-public-key-file] -ct [path-to-ciphertext-file]
```

#### Decrypt Data
To decrypt data:
```bash
go run main.go decrypt -ct [path-to-ciphertext-file] -sk [path-to-secret-key-file] -numvalues [number-of-values]
```

## Examples
#### Key Generation
```bash
go run main.go generate-keys -sk secret.key -pk public.key
```

#### Encryption
```bash
go run main.go encrypt -values 123,456,789 -pk public.key -ct data.ct
```

#### Decryption
```bash
go run main.go decrypt -ct data.ct -sk secret.key -numvalues 3
```

## Disclaimer
This project is for educational purposes only. Please ensure to understand the underlying cryptographic principles before using it for security-critical applications.