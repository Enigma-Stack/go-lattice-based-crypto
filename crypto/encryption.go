package crypto

import (
	"fmt"
	"go-lattice-based-crypto/utils"
	"log"
	"os"

	"github.com/tuneinsight/lattigo/v5/core/rlwe"
	"github.com/tuneinsight/lattigo/v5/he/heint"
)

// EncryptData encrypts a slice of uint64 values and saves the ciphertext to a file.
func EncryptData(valuesStr, pkFilePath, ctFilePath string) {
	params, err := heint.NewParametersFromLiteral(heint.ParametersLiteral{
		LogN:             14,
		LogQ:             []int{55, 45, 45, 45, 45, 45, 45, 45},
		LogP:             []int{61},
		PlaintextModulus: 0x10001,
	})
	if err != nil {
		log.Fatalf("Error creating parameters: %v", err)
	}

    serializedPk, err := os.ReadFile(pkFilePath)
    if err != nil {
        panic(err)  // Or better error handling
    }

    pk := &rlwe.PublicKey{}
    if err = pk.UnmarshalBinary(serializedPk); err != nil {
        panic(err)  // Or better error handling
    }

	ecd := heint.NewEncoder(params)
	enc := rlwe.NewEncryptor(params, pk)

	values := utils.ParseValues(valuesStr)
	if len(values) > params.MaxSlots() {
		log.Printf("Number of input values exceeds MaxSlots (%d). Truncating...\n", params.MaxSlots())
		values = values[:params.MaxSlots()]
	}

	pt := heint.NewPlaintext(params, params.MaxLevel())
	if err = ecd.Encode(values, pt); err != nil {
		log.Fatalf("Error encoding values: %v", err)
	}

	ct, err := enc.EncryptNew(pt)
	if err != nil {
		log.Fatalf("Error encrypting data: %v", err)
	}

	serializedCt, err := ct.MarshalBinary()
	if err != nil {
		log.Fatalf("Error serializing ciphertext: %v", err)
	}

	err = os.WriteFile(ctFilePath, serializedCt, 0600)
	if err != nil {
		log.Fatalf("Error writing ciphertext to file: %v", err)
	}

	fmt.Printf("Ciphertext saved to: %s\n", ctFilePath)
}
