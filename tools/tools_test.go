package tools

import (
	"fmt"
	"net"
	"testing"

	"github.com/fff-chain/go-fff/crypto"
	"github.com/fff-chain/go-fff/p2p/enode"
)

func TestEnodeSync(t *testing.T) {
	key, _ := crypto.HexToECDSA("c11f1be8baea65bd18bb81f2c7213deabbef6360f9eb24c470a36577e38b8062")
	address := crypto.PubkeyToAddress(key.PublicKey)
	fmt.Println(address)
	s := enode.NewV4(&key.PublicKey, net.IP{127, 0, 0, 1}, 30300, 0)
	fmt.Println(s, "s")
}
