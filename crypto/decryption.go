package crypto

import (
	"os"

	"github.com/tuneinsight/lattigo/v5/core/rlwe"
	"github.com/tuneinsight/lattigo/v5/he/heint"
)

func DecryptData(ctFilePath, skFilePath string, numValues int) []uint64 {
	params, err := heint.NewParametersFromLiteral(heint.ParametersLiteral{
		LogN:             14,
		LogQ:             []int{55, 45, 45, 45, 45, 45, 45, 45},
		LogP:             []int{61},
		PlaintextModulus: 0x10001,
	})
	if err != nil {
		panic(err)
	}

	serializedSk, err := os.ReadFile(skFilePath)
	if err != nil {
		panic(err)
	}

	sk := &rlwe.SecretKey{}
	err = sk.UnmarshalBinary(serializedSk)
	if err != nil {
		panic(err)
	}

	serializedCt, err := os.ReadFile(ctFilePath)
	if err != nil {
		panic(err)
	}

	ct := new(rlwe.Ciphertext)
	err = ct.UnmarshalBinary(serializedCt)
	if err != nil {
		panic(err)
	}

	dec := rlwe.NewDecryptor(params, sk)
	ecd := heint.NewEncoder(params)

	pt := dec.DecryptNew(ct)
	have := make([]uint64, params.MaxSlots())
	if err = ecd.Decode(pt, have); err != nil {
		panic(err)
	}

    if numValues > params.MaxSlots() || numValues <= 0 {
        numValues = params.MaxSlots()
    }

    return have[:numValues]
}
