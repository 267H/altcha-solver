package altcha

import (
	"encoding/base64"
	"testing"

	"github.com/mailru/easyjson"
)

var knownChallenge = Challenge{
	Algorithm: "SHA-256",
	Challenge: "506e536c83c509fd6a1ed417f7ee6f3d10180e08e3aef762bb96a6e5ea14a91d",
	MaxNumber: 50000,
	Salt:      "03f85730982b393ac97e033e",
	Signature: "0d69195b4a9194dae596306847a991fbdc798a67f916f0ed1778ccccefffb469",
}

func TestSolveAltchaChallenge_KnownChallenge(t *testing.T) {
	result, err := SolveAltchaChallenge(knownChallenge)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if result == "" {
		t.Fatal("Expected non-empty result")
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(result)
	if err != nil {
		t.Fatalf("Failed to decode base64 result: %v", err)
	}

	var solution Solution
	if err := easyjson.Unmarshal(decodedBytes, &solution); err != nil {
		t.Fatalf("Failed to unmarshal solution: %v", err)
	}

	if solution.Algorithm != knownChallenge.Algorithm {
		t.Errorf("Expected algorithm %s, got %s", knownChallenge.Algorithm, solution.Algorithm)
	}

	if solution.Challenge != knownChallenge.Challenge {
		t.Errorf("Expected challenge %s, got %s", knownChallenge.Challenge, solution.Challenge)
	}

	if solution.Salt != knownChallenge.Salt {
		t.Errorf("Expected salt %s, got %s", knownChallenge.Salt, solution.Salt)
	}

	if solution.Signature != knownChallenge.Signature {
		t.Errorf("Expected signature %s, got %s", knownChallenge.Signature, solution.Signature)
	}

	if solution.Number != 48597 {
		t.Errorf("Expected number 48597, got %d", solution.Number)
	}

	if solution.Took < 40 || solution.Took > 90 {
		t.Errorf("Expected took between 40-90, got %d", solution.Took)
	}
}

func BenchmarkSolveAltchaChallenge(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := SolveAltchaChallenge(knownChallenge)
		if err != nil {
			b.Fatalf("Benchmark failed: %v", err)
		}
	}
}
