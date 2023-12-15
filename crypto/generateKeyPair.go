package crypto

import (
    "github.com/tuneinsight/lattigo/v5/core/rlwe"
    "github.com/tuneinsight/lattigo/v5/he/heint"
    "os"
)

func GenerateKeyPair(skFilePath, pkFilePath string) {
    params, err := heint.NewParametersFromLiteral(heint.ParametersLiteral{
        LogN:             14,
        LogQ:             []int{55, 45, 45, 45, 45, 45, 45, 45},
        LogP:             []int{61},
        PlaintextModulus: 0x10001,
    })
    if err != nil {
        panic(err)
    }

    kgen := rlwe.NewKeyGenerator(params)

    // Generate Secret Key
    sk := kgen.GenSecretKeyNew()

    // Serialize and save the Secret Key
    serializedSk, err := sk.MarshalBinary()
    if err != nil {
        panic(err)
    }
    err = os.WriteFile(skFilePath, serializedSk, 0600)
    if err != nil {
        panic(err)
    }

    // Generate Public Key
    pk := kgen.GenPublicKeyNew(sk)

    // Serialize and save the Public Key
    serializedPk, err := pk.MarshalBinary()
    if err != nil {
        panic(err)
    }
    err = os.WriteFile(pkFilePath, serializedPk, 0600)
    if err != nil {
        panic(err)
    }
}
