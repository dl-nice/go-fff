package tools

import (
	"fmt"
	"net"
	"testing"

	"github.com/fff-chain/go-fff/crypto"
	"github.com/fff-chain/go-fff/p2p/enode"
)

// 0647bca0fe72da1acad51538796f9333f0ea41b49669af3b33a31f59059a51c8
// bdb9c1b42ebeb76eb2418621ea40b15e30340b677dd003d9a20f20dbf9431841
// enode://76d26bd2f784468e4e0924a46eacbba710ae0dd37e389fad39f68fd63526cdb5a1d695484ad1d8b1302b1e46567232d23a6527552f7330062b2aa6790360a6e0@87.118.86.56:30303?
func TestEnodeSync(t *testing.T) {
	key, _ := crypto.HexToECDSA("a001e32e78602d2796282c647b4225601cb6399fb0cb353e26e6b877d79c8b97")
	address := crypto.PubkeyToAddress(key.PublicKey)
	fmt.Println(address)
	s := enode.NewV4(&key.PublicKey, net.IP{127, 0, 0, 1}, 30300, 0)
	fmt.Println(s, "s")
}
