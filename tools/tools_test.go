package tools

import (
	"fmt"
	"net"
	"testing"

	"github.com/fff-chain/go-fff/crypto"
	"github.com/fff-chain/go-fff/p2p/enode"
)

func TestEnodeSync(t *testing.T) {
	key, _ := crypto.HexToECDSA("ffa3ae7c376602fd606409755b6b9aa2f6318995de90b17192fbe5c7db0b6741")
	address := crypto.PubkeyToAddress(key.PublicKey)
	fmt.Println(address)
	s := enode.NewV4(&key.PublicKey, net.IP{127, 0, 0, 1}, 30300, 0)
	fmt.Println(s, "s")
}
