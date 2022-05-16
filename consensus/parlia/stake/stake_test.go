package stake

import (
	"context"
	"github.com/fff-chain/go-fff/accounts/abi/bind"
	"github.com/fff-chain/go-fff/common"
	"github.com/fff-chain/go-fff/consensus/parlia/storage"
	"github.com/fff-chain/go-fff/crypto"
	"github.com/fff-chain/go-fff/ethclient"
	"log"
	"math/big"
	"testing"
	"time"
)

func TestDeploySafeMath(t *testing.T) {

	d, _ := ethclient.Dial("http://127.0.0.1:8545")

	cl3, _ := storage.NewBSCValidatorSetCaller(common.HexToAddress("FFF3Psbq3enwAmwXGa2QejWFdd9AwV1rczE6w1GPzs6WTPmJpKbmWiBrcX"), d)


	va,_:=cl3.GetValidators(nil)

	log.Println(va)


	return


	var currSubmit = 0

	cl2, _ := storage.NewBSCValidatorSetTransactor(common.HexToAddress("FFF3Psbq3enwAmwXGa2QejWFdd9AwV1rczE6w1GPzs6WTPmJpKbmWiBrcX"), d)
	nonce, err := d.PendingNonceAt(context.Background(), common.HexToAddress("0x5Eba2ee915B919b1da7E39906ccE78af9eF26869"))
	if err != nil {
		log.Println("节点异常")
	}
	for {
		pri := "6d952b8d68a0eeba71fa309d45aadacc4e7a3104f9d1df91b56be468790bbf39"
		priKeyECD, err := crypto.HexToECDSA(pri)

		if err != nil {
			log.Println("私钥异常")

		}

		chianId, err := d.ChainID(context.Background())
		if err != nil {
			log.Println("节点chianId异常", err)

		}
		param, err := bind.NewKeyedTransactorWithChainID(priKeyECD, chianId)
		if err != nil {
			log.Println("节点Transactor异常", err)

		}
		param.GasPrice,_= d.SuggestGasPrice(context.Background())

		param.Nonce = big.NewInt(int64(nonce))

		log.Println(param)

		trans, err := cl2.CreateAddress(param,big.NewInt(1))

		if err != nil {
			log.Println(err)
			continue
		}

		log.Println(trans.Hash(), err)
		currSubmit++
		nonce++

		if currSubmit > 0{
			break
		}
	}
	time.Sleep(time.Second*300)

	TestDeploySafeMath(t)

}
func TestNewTransferHelperCaller2(t *testing.T) {
	log.Println(storage.BSCValidatorSetABI)

}
func TestDeployTransferHelper(t *testing.T) {


}
func TestNewTransferHelperCaller(t *testing.T) {
	d, _ := ethclient.Dial("http://127.0.0.1:8545")

	bf, err := d.PendingBalanceAt(context.Background(), common.HexToAddress("FFF5tXp4dJJWemxB4RPMmnVAcanbnsz8MH5Vi3Uf1caR8wtiKtW5nUsGmr"))

	log.Println(bf, err)

}
