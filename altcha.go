package altcha

import (
	"bytes"
	"crypto/sha256"
	"encoding"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/mailru/easyjson"
	"math/rand"
	"strconv"
)

type Challenge struct {
	Algorithm string `json:"algorithm"`
	Challenge string `json:"challenge"`
	MaxNumber int    `json:"maxNumber"`
	Salt      string `json:"salt"`
	Signature string `json:"signature"`
}

type Solution struct {
	Algorithm string `json:"algorithm"`
	Challenge string `json:"challenge"`
	Number    int    `json:"number"`
	Salt      string `json:"salt"`
	Signature string `json:"signature"`
	Took      int64  `json:"took"`
}

func SolveAltchaChallenge(challenge Challenge) (string, error) {
	targetHash, err := hex.DecodeString(challenge.Challenge)
	if err != nil {
		return "", fmt.Errorf("invalid challenge hash: %v", err)
	}

	saltBytes := []byte(challenge.Salt)

	hasher := sha256.New()
	hasher.Write(saltBytes)
	saltState, _ := hasher.(encoding.BinaryMarshaler).MarshalBinary()

	buf := make([]byte, 0, len(saltBytes)+20)
	buf = append(buf, saltBytes...)
	saltLen := len(saltBytes)

	var hashResult [sha256.Size]byte

	for i := 0; i <= challenge.MaxNumber; i++ {
		err := hasher.(encoding.BinaryUnmarshaler).UnmarshalBinary(saltState)
		if err != nil {
			return "", err
		}

		buf = buf[:saltLen]
		buf = strconv.AppendInt(buf, int64(i), 10)
		hasher.Write(buf[saltLen:])
		hasher.Sum(hashResult[:0])

		if bytes.Equal(hashResult[:], targetHash) {
			randomTook := rand.Int63n(51) + 40
			solution := &Solution{
				Algorithm: challenge.Algorithm,
				Challenge: challenge.Challenge,
				Number:    i,
				Salt:      challenge.Salt,
				Signature: challenge.Signature,
				Took:      randomTook,
			}

			jsonData, err := easyjson.Marshal(solution)
			if err != nil {
				return "", fmt.Errorf("error marshaling solution: %v", err)
			}

			return base64.StdEncoding.EncodeToString(jsonData), nil
		}
	}

	return "", fmt.Errorf("no solution found within maxNumber %d", challenge.MaxNumber)
}
