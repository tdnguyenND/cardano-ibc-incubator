package helpers

import (
	"encoding/hex"
	"errors"
	"github.com/fxamacker/cbor/v2"
	"golang.org/x/crypto/sha3"
	"math"
	"strconv"
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

func GetConnectionGetDatumDetail(datum string) (*ConnectionDatum, error) {
	datum = datum[2:]
	var vOutput ConnectionDatum
	datumBytes, err := hex.DecodeString(datum)
	if err != nil {
		return nil, err
	}
	err = cbor.Unmarshal(datumBytes, &vOutput)
	if err != nil {
		return nil, err
	}
	return &vOutput, nil
}
