package test

import (
	"github.com/fff-chain/go-fff/global_config"
	"github.com/fff-chain/go-fff/global_config/wallet"
	"log"
	"testing"
)

func TestTransfer(t *testing.T) {
	wl := wallet.Init("http://127.0.0.1:8545", "http://127.0.0.1:8545")
	hash := wl.Transfer("20a442166fda1598b32d85d0cf5b30fe5867bf23c943f8584013eab0ad49a88c", "0x49E354E5149eAC047d74879595F97C320b53D66B", global_config.EthToWei(1), "", 0)
	log.Println(hash)
}
