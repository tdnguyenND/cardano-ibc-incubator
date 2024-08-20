package helpers

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/sha3"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

// Fraction struct to hold the numerator and denominator
type Fraction struct {
	Numerator   uint64
	Denominator uint64
}

// Function to compute the greatest common divisor (GCD)
func gcd(a, b uint64) uint64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Function to convert a float to a Fraction
func FloatToFraction(f float64) Fraction {
	const precision = 1e-9 // precision for floating point comparison

	// Handle special cases for zero
	if f == 0.0 {
		return Fraction{0, 1}
	}

	// Determine the sign of the fraction
	sign := int64(1)
	if f < 0 {
		sign = -1
		f = -f
	}

	// Initialize the numerator and denominator
	numerator := uint64(f)
	denominator := uint64(1)

	// Adjust until the fractional part is within the precision range
	for math.Abs(f-float64(numerator)/float64(denominator)) > precision {
		denominator *= 10
		numerator = uint64(f * float64(denominator))
	}

	// Reduce the fraction by dividing by the GCD
	g := gcd(numerator, denominator)
	numerator /= g
	denominator /= g

	return Fraction{uint64(sign) * numerator, denominator}
}

type AuthToken struct {
	PolicyId string
	Name     string
}

func GenerateTokenName(baseToken AuthToken, prefix string, postfix int64) (string, error) {
	if postfix < 0 {
		return "", errors.New("sequence must be unsigned integer")
	}
	postfixHex := ConvertString2Hex(strconv.FormatInt(postfix, 10))
	if len(postfixHex) > 16 {
		return "", errors.New("postfix size > 8 bytes")
	}
	baseTokenPart := HashSha3_256(baseToken.PolicyId + baseToken.Name)[:40]
	prefixPart := HashSha3_256(prefix)[:8]
	fullName := baseTokenPart + prefixPart + postfixHex
	return fullName, nil
}
func ConvertString2Hex(str string) string {
	if str == "" {
		return ""
	}
	return hex.EncodeToString([]byte(str))
}
func HashSha3_256(data string) string {
	dataBytes, _ := hex.DecodeString(data)
	hash := sha3.Sum256(dataBytes)
	return hex.EncodeToString(hash[:])
}

func GetChainHandler() (*AutoGenerated, error) {
	jsonFile, err := os.Open("./examples/demo/configs/chains/chain_handler.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var chainHandler AutoGenerated
	err = json.Unmarshal(byteValue, &chainHandler)
	if err != nil {
		return nil, err
	}
	return &chainHandler, nil
}

func GeneratePaginationKey(key PaginationKeyDto) []byte {
	keyByteArray, err := json.Marshal(key)
	if err != nil {
		return nil
	}
	return keyByteArray
}

func DecodePaginationKey(key []byte) (uint64, error) {
	res := &PaginationKeyDto{}
	err := json.Unmarshal(key, res)
	if err != nil {
		return 0, err
	}
	return res.Offset, nil
}

func GetEntityIdFromTokenName(tokenName string, baseToken AuthToken, prefix string) string {
	baseTokenPart := HashSha3_256(baseToken.PolicyId + baseToken.Name)[:40]
	prefixPart := HashSha3_256(prefix)[:8]
	fullName := baseTokenPart + prefixPart
	if !strings.Contains(tokenName, fullName) {
		return ""
	}

	idHex := strings.ReplaceAll(tokenName, fullName, "")
	id, err := hex.DecodeString(idHex)
	if err != nil {
		return ""
	}
	return string(id)
}
