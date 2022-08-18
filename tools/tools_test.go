package tools

import (
	"fmt"
	"net"
	"testing"

	"github.com/fff-chain/go-fff/crypto"
	"github.com/fff-chain/go-fff/p2p/enode"
)

func TestEnodeSync(t *testing.T) {
	key, _ := crypto.HexToECDSA("4fb97c16fb9d1e5f801ad71eb79dd36ca36ac1dde6c47e7079875a0bead73bfb")
	address := crypto.PubkeyToAddress(key.PublicKey)
	fmt.Println(address)
	s := enode.NewV4(&key.PublicKey, net.IP{127, 0, 0, 1}, 30300, 0)
	fmt.Println(s, "s")
}
