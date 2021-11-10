package plookup

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr/kzg"
)

func TestLookupVector(t *testing.T) {

	lookupVector := make(Table, 8)
	fvector := make(Table, 7)
	for i := 0; i < 8; i++ {
		lookupVector[i].SetUint64(uint64(2 * i))
	}
	for i := 0; i < 7; i++ {
		fvector[i].Set(&lookupVector[(4*i+1)%8])
	}

	srs, err := kzg.NewSRS(64, big.NewInt(13))
	if err != nil {
		t.Fatal(err)
	}

	// correct proof vector
	{
		proof, err := ProveLookupVector(srs, fvector, lookupVector)
		if err != nil {
			t.Fatal(err)
		}

		err = VerifyLookupVector(srs, proof)
		if err != nil {
			t.Fatal(err)
		}
	}

	// wrong proofs vector
	{
		fvector[0].SetRandom()

		proof, err := ProveLookupVector(srs, fvector, lookupVector)
		if err != nil {
			t.Fatal(err)
		}

		err = VerifyLookupVector(srs, proof)
		if err == nil {
			t.Fatal(err)
		}
	}

}

func TestLookupTable(t *testing.T) {

	srs, err := kzg.NewSRS(64, big.NewInt(13))
	if err != nil {
		t.Fatal(err)
	}

	lookupTable := make([]Table, 3)
	fTable := make([]Table, 3)
	for i := 0; i < 3; i++ {
		lookupTable[i] = make(Table, 8)
		fTable[i] = make(Table, 7)
		for j := 0; j < 8; j++ {
			lookupTable[i][j].SetUint64(uint64(2*i + j))
		}
		for j := 0; j < 7; j++ {
			fTable[i][j].Set(&lookupTable[i][(4*j+1)%8])
		}
	}

	// correct proof
	{
		proof, err := ProveLookupTables(srs, fTable, lookupTable)
		if err != nil {
			t.Fatal(err)
		}

		err = VerifyLookupTables(srs, proof)
		if err != nil {
			t.Fatal(err)
		}
	}

	// wrong proof
	{
		fTable[0][0].SetRandom()
		proof, err := ProveLookupTables(srs, fTable, lookupTable)
		if err != nil {
			t.Fatal(err)
		}

		err = VerifyLookupTables(srs, proof)
		if err == nil {
			t.Fatal(err)
		}
	}

}
