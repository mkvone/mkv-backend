package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/bech32"
)

func PrettyPrint(data interface{}) {
	b, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Println("Error: ", err)
	}
	fmt.Println(string(b))
}
func ExtractPrefixFromAddress(address string) string {
	parts := strings.Split(address, "valoper")
	if len(parts) > 0 {
		return parts[0] + "valcons"
	}
	return ""
}

func EncodeValidatorAddress(pubKeyBase64, valaddr string) (string, error) {
	prefix := ExtractPrefixFromAddress(valaddr)
	if prefix == "" {
		return "", fmt.Errorf("failed to extract prefix from address")
	}
	if len(pubKeyBase64) == 0 {
		return "", fmt.Errorf("pubKeyBase64 is empty")
	}
	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKeyBase64)
	if err != nil {
		return "", fmt.Errorf("base64 decoding error: %v", err)
	}

	hash := sha256.Sum256(pubKeyBytes)
	truncatedHash := hash[:20]

	address, err := bech32.ConvertAndEncode(prefix, truncatedHash)
	if err != nil {
		return "", fmt.Errorf("bech32 encoding error: %v", err)
	}

	return address, nil
}
