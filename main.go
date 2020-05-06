package main

import (
	"fmt"

	"github.com/harmony-one/bls/ffi/go/bls"
	b "github.com/harmony-one/harmony/crypto/bls"
)

func p() {
	p := "hello world"
	b1, b2, b3 := b.RandPrivateKey(), b.RandPrivateKey(), b.RandPrivateKey()
	b1p, b2p, _ := b1.GetPublicKey(), b2.GetPublicKey(), b3.GetPublicKey()
	sigs := []*bls.Sign{b1.Sign(p), b2.Sign(p), b2.Sign(p)}
	agSig := b.AggregateSig(sigs)
	publicKeys := []*bls.PublicKey{b1p, b2p, b2p}
	mask, _ := b.NewMask(publicKeys, nil)

	for _, key := range publicKeys {
		fmt.Println("signer", key.SerializeToHexStr())
		mask.SetKey(key, true)
	}

	// Can only touch aggregatepublic after the SetKey functions happened
	agKey := mask.AggregatePublic

	if agSig.Verify(agKey, p) {
		fmt.Println("does verify")
	} else {
		fmt.Println("does not verify")
	}

	fmt.Println("agg sig:", agSig.SerializeToHexStr())

	for i, key := range mask.GetPubKeyFromMask(true) {
		switch signed, err := mask.IndexEnabled(i); true {
		case err != nil:
			continue
		case signed:
			fmt.Println("signed-", key.SerializeToHexStr())
		default:
			fmt.Println("shouldnt happen")
		}
	}

}

func main() {
	p()
}
