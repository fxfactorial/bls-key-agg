package main

import (
	"fmt"

	"github.com/harmony-one/bls/ffi/go/bls"
	b "github.com/harmony-one/harmony/crypto/bls"
)

func green(s string) {
	fmt.Printf("\x1b[32m%s\x1b[0m \n", s)
}

func red(s string) {
	fmt.Printf("\x1b[31m%s\x1b[0m \n", s)
}

func p() {
	p := "hello world"
	fmt.Println("Using the following inside the backticks as signing input:\t", "`"+p+"`")
	b1, b2, b3 := b.RandPrivateKey(), b.RandPrivateKey(), b.RandPrivateKey()
	var (
		sigs       []*bls.Sign
		publicKeys []*bls.PublicKey
	)
	// NOTE Order of public keys matters, the mask is
	// and implementation detail in harmony
	privateKeys := []*bls.SecretKey{b1, b2, b3}
	for _, k := range privateKeys {
		publicKeys = append(publicKeys, k.GetPublicKey())
	}
	mask, _ := b.NewMask(publicKeys, nil)
	for i, key := range privateKeys {

		// mask.SetKey(publicKeys[i], false)
		// sigs = append(sigs, &bls.Sign{})

		if i == 2 {
			shouldBe := publicKeys[i]
			mask.SetKey(shouldBe, true)
			other := privateKeys[i-1]
			otherP := other.GetPublicKey()
			sigs = append(sigs, other.Sign(p))
			red("this voter voter should be signing \n" +
				shouldBe.SerializeToHexStr() + "\nbut instead this voter signed\n" +
				otherP.SerializeToHexStr(),
			)
			break
		}

		green("this voter does sign->\n\t" + publicKeys[i].SerializeToHexStr())
		sigs = append(sigs, key.Sign(p))
		mask.SetKey(publicKeys[i], true)

	}

	// Can only touch aggregatepublic after the SetKey functions happened
	agKey := mask.AggregatePublic
	agSig := b.AggregateSig(sigs)

	green("\nThe aggregate public key: \n\t" + agKey.SerializeToHexStr())
	green("\nThe signature aggregated: \n\t" + agSig.SerializeToHexStr())

	if agSig.Verify(agKey, p) {
		green("aggregate key does verify the aggregate signature")
	} else {
		red("\n\nthe aggregate signature cannot be verified by the aggregate key")

	}

}

func main() {
	p()
}
