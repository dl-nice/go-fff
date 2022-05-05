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

	bl,err:=d.TransactionReceipt(context.Background(),common.HexToHash("0x24b9db5436a3e0e2ccde6e3c91dac2e8e7a05834836ed5b4bdaaf990b8ba5536"))


	log.Println(bl, err)

	return

	var currSubmit = 0

	cl2, _ := storage.NewBSCValidatorSetTransactor(common.HexToAddress("0x0000000000000000000000000000000000001000"), d)
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

	bf, err := d.TransactionReceipt(context.Background(), common.HexToHash("0x673b92c09099ed42284baafb5370cc900fc1c328b234a631deef2eafa5f22036"))

	log.Println(bf.Status, err)

}
