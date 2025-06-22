# Altcha Solver

A Go library for solving Altcha proof-of-work challenges.

## Installation

```bash
go get github.com/267H/altcha-solver
```

## How it works

Altcha uses a proof-of-work challenge where you must find a number that when concatenated with a salt produces a specific SHA-256 hash.

1. **Request a challenge** from an Altcha endpoint:
   ```
   GET https://api.example.com/altcha
   ```

2. **Challenge response**:
   ```json
   {
     "algorithm": "SHA-256",
     "challenge": "506e536c83c509fd6a1ed417f7ee6f3d10180e08e3aef762bb96a6e5ea14a91d",
     "maxNumber": 50000,
     "salt": "03f85730982b393ac97e033e",
     "signature": "0d69195b4a9194dae596306847a991fbdc798a67f916f0ed1778ccccefffb469"
   }
   ```

3. **Solve the challenge** by finding a number where:
   ```
   SHA256(salt + number) = challenge
   ```

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/267H/altcha-solver"
)

func main() {
	challenge := altcha.Challenge{
		Algorithm: "SHA-256",
		Challenge: "506e536c83c509fd6a1ed417f7ee6f3d10180e08e3aef762bb96a6e5ea14a91d",
		MaxNumber: 50000,
		Salt:      "03f85730982b393ac97e033e",
		Signature: "0d69195b4a9194dae596306847a991fbdc798a67f916f0ed1778ccccefffb469",
	}

	solution, err := altcha.SolveAltchaChallenge(challenge)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solution)
}
```

## Returns

The solver returns a base64-encoded JSON string containing the solution:

```json
{
  "algorithm": "SHA-256",
  "challenge": "506e536c83c509fd6a1ed417f7ee6f3d10180e08e3aef762bb96a6e5ea14a91d",
  "number": 48597,
  "salt": "03f85730982b393ac97e033e",
  "signature": "0d69195b4a9194dae596306847a991fbdc798a67f916f0ed1778ccccefffb469",
  "took": 65
}
```

The base64-encoded string can be used to access protected endpoints.