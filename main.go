package main

import (
    "fmt"
    "mpc-ecdsa/keygen"
    "mpc-ecdsa/sign"
)

func main() {
    // 1. Key Generation (Threshold)
    pid, err := keygen.GenerateThresholdKey()
    if err != nil {
        panic(err)
    }

    // 2. Sign a message
    message := "Hello MPC ECDSA"
    sig, err := sign.SignMessage(pid, message)
    if err != nil {
        panic(err)
    }

    // 3. Output
    fmt.Println("✅ Final Signature:", sig)
}
