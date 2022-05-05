// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package storage

import (
	ethereum "github.com/fff-chain/go-fff"
	"github.com/fff-chain/go-fff/accounts/abi"
	"github.com/fff-chain/go-fff/accounts/abi/bind"
	"github.com/fff-chain/go-fff/common"
	"github.com/fff-chain/go-fff/core/types"
	"github.com/fff-chain/go-fff/event"
	"math/big"
	"strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BSCValidatorSetStakeInfo is an auto generated low-level Go binding around an user-defined struct.
type BSCValidatorSetStakeInfo struct {
	StakeAddress common.Address
	StakeCount   *big.Int
}

// BSCValidatorSetABI is the input ABI used to generate the binding from.
const BSCValidatorSetABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"DistributeBlockRewardFunc\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"UserStakeFFF\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"UserUnStakeFFF\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"batchTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"batchTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"batchTransferLowerFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deprecatedDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"addresspayable\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"directTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"addresspayable\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"directTransferFail\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"failReasonWithStr\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"feeBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"paramChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"systemTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"unexpectedPackage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"validatorDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"validatorEmptyJailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"validatorFelony\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"validatorJailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"validatorMisdemeanor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"validatorSetUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BIND_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURN_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURN_RATIO_SCALE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CODE_OK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"CreateAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DUSTY_INCOMING\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"DistributeBlockReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_FAIL_CHECK_VALIDATORS\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_FAIL_DECODE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_LEN_OF_VAL_MISMATCH\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_RELAYFEE_TOO_LARGE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_UNKNOWN_PACKAGE_TYPE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXPIRE_TIME_SECOND_GAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetAllStakeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"stakeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeCount\",\"type\":\"uint256\"}],\"internalType\":\"structBSCValidatorSet.StakeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INCENTIVIZE_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_BURN_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_VALIDATORSET_BYTES\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"JAIL_MESSAGE_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NUM_OF_VALIDATORS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STAKE_FFF_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM_REWARD_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"StakeFFF\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_MANAGER_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UnStakeFFF\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORS_UPDATE_MESSAGE_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bscChainID\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnRatioInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currStakeFFF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currStakePeople\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentValidatorSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"feeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BBCFeeAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"votingPower\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"incoming\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"currentValidatorSetMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expireTimeSecondGap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"felony\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getIncoming\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleAckPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleFailAckPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleSynPackage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"responsePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxStakeFFFCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxStakePeopleCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miniStakeFFFCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"misdemeanor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"myBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfJailed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeInfoIndexMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakeInfoMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInComing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"updateParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BSCValidatorSetFuncSigs maps the 4-byte function signature to its string representation.
var BSCValidatorSetFuncSigs = map[string]string{
	"3dffc387": "BIND_CHANNELID()",
	"fccc2813": "BURN_ADDRESS()",
	"3de0f0d8": "BURN_RATIO_SCALE()",
	"ab51bb96": "CODE_OK()",
	"51e80672": "CROSS_CHAIN_CONTRACT_ADDR()",
	"2463cd09": "CreateAddress(uint256)",
	"d86222d5": "DUSTY_INCOMING()",
	"50403540": "DistributeBlockReward(uint256)",
	"81650b62": "ERROR_FAIL_CHECK_VALIDATORS()",
	"0bee7a67": "ERROR_FAIL_DECODE()",
	"5d77156c": "ERROR_LEN_OF_VAL_MISMATCH()",
	"219f22d5": "ERROR_RELAYFEE_TOO_LARGE()",
	"eda5868c": "ERROR_UNKNOWN_PACKAGE_TYPE()",
	"853230aa": "EXPIRE_TIME_SECOND_GAP()",
	"96713da9": "GOV_CHANNELID()",
	"9dc09262": "GOV_HUB_ADDR()",
	"54b1a671": "GetAllStakeInfo()",
	"6e47b482": "INCENTIVIZE_ADDR()",
	"78dfed4a": "INIT_BURN_RATIO()",
	"a5422d5c": "INIT_VALIDATORSET_BYTES()",
	"bf9f4995": "JAIL_MESSAGE_TYPE()",
	"dc927faf": "LIGHT_CLIENT_ADDR()",
	"e086c7b1": "MAX_NUM_OF_VALIDATORS()",
	"8a10febe": "MAX_STAKE_FFF_COUNT()",
	"aaf5eb68": "PRECISION()",
	"a1a11bf5": "RELAYERHUB_CONTRACT_ADDR()",
	"7942fd05": "SLASH_CHANNELID()",
	"43756e5c": "SLASH_CONTRACT_ADDR()",
	"4bf6c882": "STAKING_CHANNELID()",
	"c81b1662": "SYSTEM_REWARD_ADDR()",
	"0db6cccb": "StakeFFF()",
	"fd6a6879": "TOKEN_HUB_ADDR()",
	"75d47a0a": "TOKEN_MANAGER_ADDR()",
	"70fd5bad": "TRANSFER_IN_CHANNELID()",
	"fc3e5908": "TRANSFER_OUT_CHANNELID()",
	"30c62a13": "UnStakeFFF()",
	"5667515a": "VALIDATORS_UPDATE_MESSAGE_TYPE()",
	"f9a2bbc7": "VALIDATOR_CONTRACT_ADDR()",
	"a78abc16": "alreadyInit()",
	"493279b1": "bscChainID()",
	"5192c82c": "burnRatio()",
	"152ad3b8": "burnRatioInitialized()",
	"101a875f": "currStakeFFF()",
	"93a101f3": "currStakePeople()",
	"6969a25c": "currentValidatorSet(uint256)",
	"ad3c9da6": "currentValidatorSetMap(address)",
	"f340fa01": "deposit(address)",
	"86249882": "expireTimeSecondGap()",
	"35409f7f": "felony(address)",
	"565c56b3": "getIncoming(address)",
	"b7ab4db5": "getValidators()",
	"831d65d1": "handleAckPackage(uint8,bytes)",
	"c8509d81": "handleFailAckPackage(uint8,bytes)",
	"1182b875": "handleSynPackage(uint8,bytes)",
	"e1c7392a": "init()",
	"b2ccf3a4": "maxStakeFFFCount()",
	"f78ffcbc": "maxStakePeopleCount()",
	"1faea20b": "miniStakeFFFCount()",
	"eb57e202": "misdemeanor(address)",
	"c9116b69": "myBalance()",
	"daacdb66": "numOfJailed()",
	"8da5cb5b": "owner()",
	"7dedf367": "stakeInfoIndexMap(uint256)",
	"c5f30db5": "stakeInfoMap(address)",
	"1ff18069": "totalInComing()",
	"ac431751": "updateParam(string,bytes)",
}

// BSCValidatorSetBin is the compiled bytecode used for deploying new contracts.
var BSCValidatorSetBin = "0x608060405234801561001057600080fd5b50614b1c806100206000396000f3fe6080604052600436106103d95760003560e01c806386249882116101fd578063c5f30db511610118578063e1c7392a116100ab578063f78ffcbc1161007a578063f78ffcbc146109a2578063f9a2bbc7146109b7578063fc3e5908146109cc578063fccc2813146109e1578063fd6a6879146109f6576103d9565b8063e1c7392a14610945578063eb57e2021461095a578063eda5868c1461097a578063f340fa011461098f576103d9565b8063d86222d5116100e7578063d86222d5146108f1578063daacdb6614610906578063dc927faf1461091b578063e086c7b114610930576103d9565b8063c5f30db5146108a7578063c81b1662146108c7578063c8509d8114610714578063c9116b69146108dc576103d9565b8063a78abc1611610190578063ad3c9da61161015f578063ad3c9da614610850578063b2ccf3a414610870578063b7ab4db514610885578063bf9f499514610520576103d9565b8063a78abc16146107f1578063aaf5eb6814610806578063ab51bb961461081b578063ac43175114610830576103d9565b806396713da9116101cc57806396713da91461079d5780639dc09262146107b2578063a1a11bf5146107c7578063a5422d5c146107dc576103d9565b806386249882146107495780638a10febe1461075e5780638da5cb5b1461077357806393a101f314610788576103d9565b806350403540116102f85780636e47b4821161028b5780637942fd051161025a5780637942fd05146106ca5780637dedf367146106df57806381650b62146106ff578063831d65d114610714578063853230aa14610734576103d9565b80636e47b4821461067657806370fd5bad1461068b57806375d47a0a146106a057806378dfed4a146106b5576103d9565b8063565c56b3116102c7578063565c56b3146105fa5780635667515a1461061a5780635d77156c1461062f5780636969a25c14610644576103d9565b8063504035401461059b5780635192c82c146105ae57806351e80672146105c357806354b1a671146105d8576103d9565b80632463cd09116103705780633dffc3871161033f5780633dffc3871461052057806343756e5c14610542578063493279b1146105645780634bf6c88214610586576103d9565b80632463cd09146104c357806330c62a13146104e357806335409f7f146104eb5780633de0f0d81461050b576103d9565b8063152ad3b8116103ac578063152ad3b8146104625780631faea20b146104845780631ff1806914610499578063219f22d5146104ae576103d9565b80630bee7a67146103de5780630db6cccb14610409578063101a875f146104135780631182b87514610435575b600080fd5b3480156103ea57600080fd5b506103f3610a0b565b60405161040091906149e3565b60405180910390f35b610411610a10565b005b34801561041f57600080fd5b50610428610a1c565b60405161040091906149cc565b34801561044157600080fd5b50610455610450366004614122565b610a22565b60405161040091906143a3565b34801561046e57600080fd5b50610477610bb2565b6040516104009190614398565b34801561049057600080fd5b50610428610bbb565b3480156104a557600080fd5b50610428610bc7565b3480156104ba57600080fd5b506103f3610bcd565b3480156104cf57600080fd5b506104116104de3660046140f2565b610bd2565b610411610c51565b3480156104f757600080fd5b50610411610506366004614043565b610d1f565b34801561051757600080fd5b50610428611016565b34801561052c57600080fd5b5061053561101c565b60405161040091906149f4565b34801561054e57600080fd5b50610557611021565b6040516104009190614244565b34801561057057600080fd5b50610579611027565b60405161040091906149bd565b34801561059257600080fd5b5061053561102c565b6104116105a93660046140f2565b611031565b3480156105ba57600080fd5b5061042861114d565b3480156105cf57600080fd5b50610557611153565b3480156105e457600080fd5b506105ed611159565b6040516104009190614340565b34801561060657600080fd5b50610428610615366004614043565b61122e565b34801561062657600080fd5b50610535611281565b34801561063b57600080fd5b506103f3611286565b34801561065057600080fd5b5061066461065f3660046140f2565b61128b565b60405161040096959493929190614271565b34801561068257600080fd5b506105576112ef565b34801561069757600080fd5b506105356112f5565b3480156106ac57600080fd5b506105576112fa565b3480156106c157600080fd5b50610428611281565b3480156106d657600080fd5b50610535611300565b3480156106eb57600080fd5b506105576106fa3660046140f2565b611305565b34801561070b57600080fd5b506103f3611320565b34801561072057600080fd5b5061041161072f366004614122565b611325565b34801561074057600080fd5b50610428611386565b34801561075557600080fd5b5061042861138c565b34801561076a57600080fd5b50610428611392565b34801561077f57600080fd5b5061055761139d565b34801561079457600080fd5b506104286113ac565b3480156107a957600080fd5b506105356113b2565b3480156107be57600080fd5b506105576113b7565b3480156107d357600080fd5b506105576113bd565b3480156107e857600080fd5b506104556113c3565b3480156107fd57600080fd5b506104776113df565b34801561081257600080fd5b506104286113e8565b34801561082757600080fd5b506103f3611281565b34801561083c57600080fd5b5061041161084b36600461408a565b6113f1565b34801561085c57600080fd5b5061042861086b366004614043565b611692565b34801561087c57600080fd5b506104286116a4565b34801561089157600080fd5b5061089a6116b2565b60405161040091906142b3565b3480156108b357600080fd5b506104286108c2366004614043565b6117d8565b3480156108d357600080fd5b506105576117ea565b3480156108e857600080fd5b506104286117f0565b3480156108fd57600080fd5b506104286117f4565b34801561091257600080fd5b50610428611800565b34801561092757600080fd5b50610557611806565b34801561093c57600080fd5b5061042861180c565b34801561095157600080fd5b50610411611811565b34801561096657600080fd5b50610411610975366004614043565b6119bd565b34801561098657600080fd5b506103f3611b6e565b61041161099d366004614043565b611b73565b3480156109ae57600080fd5b50610428611e04565b3480156109c357600080fd5b50610557611e0b565b3480156109d857600080fd5b50610535611e11565b3480156109ed57600080fd5b50610557611e16565b348015610a0257600080fd5b50610557611e1c565b606481565b610a1a3334611e22565b565b60025481565b60005460609060ff16610a505760405162461bcd60e51b8152600401610a4790614456565b60405180910390fd5b3361200014610a715760405162461bcd60e51b8152600401610a479061486e565b610a79613f5f565b6000610aba85858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611fa692505050565b9150915080610ad657610acd60646120ff565b92505050610bab565b815160009060ff16610af657610aef8360200151612160565b9050610b77565b825160ff1660011415610b7357826020015151600114610b4d577f70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2604051610b3d90614634565b60405180910390a1506067610b6e565b610aef8360200151600081518110610b6157fe5b6020026020010151612d5d565b610b77565b5060655b63ffffffff8116610b9c5750506040805160008152602081019091529150610bab9050565b610ba5816120ff565b93505050505b9392505050565b600c5460ff1681565b670de0b6b3a764000081565b60065481565b606881565b60005460ff16610bf45760405162461bcd60e51b8152600401610a4790614456565b60005b81811015610c4d57600081424340604051602001610c179392919061422e565b6040516020818303038152906040528051906020012060001c9050610c4481670de0b6b3a7640000611e22565b50600101610bf7565b5050565b33600081815260086020526040902054610c7d5760405162461bcd60e51b8152600401610a47906144c4565b3360009081526008602052604080822054905190916001600160a01b0384169183156108fc0291849190818181858888f19350505050158015610cc4573d6000803e3d6000fd5b503360008181526008602052604080822080546002805491909103905591909155517f433979128f828c623a175588cd9821124640d4353dd48d812d2b7273f1b76c3f91610d13918490614258565b60405180910390a15050565b3361100114610d405760405162461bcd60e51b8152600401610a4790614974565b6001600160a01b03811660009081526007602052604090205480610d645750611013565b600181039050600060048281548110610d7957fe5b600091825260209091206004918202016003015490549091506000190180610dc757600060048481548110610daa57fe5b906000526020600020906004020160030181905550505050611013565b836001600160a01b03167f3b6f9ef90462b512a1293ecec018670bf7b7f1876fb727590a8a6d7643130a7083604051610e0091906149cc565b60405180910390a26001600160a01b038416600090815260076020526040812055600454600019018314610f4c57600480546000198101908110610e4057fe5b906000526020600020906004020160048481548110610e5b57fe5b6000918252602082208354600492830290910180546001600160a01b03199081166001600160a01b0393841617825560018087015481840180548416918616919091179055600280880180549185018054909416919095161780835584546001600160401b03600160a01b91829004160267ffffffffffffffff60a01b1990911617808355935460ff600160e01b918290041615150260ff60e01b19909416939093179055600394850154940193909355805492860192600792919087908110610f2157fe5b600091825260208083206004909202909101546001600160a01b031683528201929092526040019020555b6004805480610f5757fe5b60008281526020812060046000199093019283020180546001600160a01b0319908116825560018201805490911690556002810180546001600160e81b03191690556003018190559155818381610faa57fe5b049050801561100e5760045460005b8181101561100b578260048281548110610fcf57fe5b9060005260206000209060040201600301540160048281548110610fef57fe5b6000918252602090912060036004909202010155600101610fb9565b50505b505050505b50565b61271081565b600181565b61100181565b606081565b600881565b3341146110505760405162461bcd60e51b8152600401610a47906148bd565b60005460ff166110725760405162461bcd60e51b8152600401610a4790614456565b60005b60015481101561110e576000818152600960209081526040808320546001600160a01b0316808452600890925282205490916110bc6110b48387612ed5565b600254612f18565b90506110c88382612f5a565b6040516001600160a01b0384169082156108fc029083906000818181858888f193505050501580156110fe573d6000803e3d6000fd5b5050600190920191506110759050565b507f77affbec109ff6a3609c802050db936c55bc48fb23db5c88f4a4818922e65239600154826040516111429291906149d5565b60405180910390a150565b600b5481565b61200081565b60608060015460405190808252806020026020018201604052801561119857816020015b611185613f77565b81526020019060019003908161117d5790505b50905060005b600154811015611228576000818152600960209081526040808320546001600160a01b0316808452600890925290912054835182908590859081106111df57fe5b6020026020010151600001906001600160a01b031690816001600160a01b0316815250508084848151811061121057fe5b6020908102919091018101510152505060010161119e565b50905090565b6001600160a01b0381166000908152600760205260408120548061125657600091505061127c565b6004600182038154811061126657fe5b9060005260206000209060040201600301549150505b919050565b600081565b606781565b6004818154811061129857fe5b600091825260209091206004909102018054600182015460028301546003909301546001600160a01b0392831694509082169291821691600160a01b81046001600160401b031691600160e01b90910460ff169086565b61100581565b600281565b61100881565b600b81565b6009602052600090815260409020546001600160a01b031681565b606681565b33612000146113465760405162461bcd60e51b8152600401610a479061486e565b7f41ce201247b6ceb957dcdb217d0b8acb50b9ea0e12af9af4f5e7f3890210160583838360405161137993929190614a02565b60405180910390a1505050565b61045f81565b60055481565b6603f2a69c702cdf81565b6003546001600160a01b031681565b60015481565b600981565b61100781565b61100681565b604051806080016040528060478152602001614aa06047913981565b60005460ff1681565b6402540be40081565b60005460ff166114135760405162461bcd60e51b8152600401610a4790614456565b33611007146114345760405162461bcd60e51b8152600401610a47906146ba565b61149e84848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080518082019091526013815272065787069726554696d655365636f6e6447617606c1b60208201529150612fec9050565b1561153b57602081146114c35760405162461bcd60e51b8152600401610a47906147e7565b604080516020601f84018190048102820181019092528281526000916115019185858083850183828082843760009201919091525061304592505050565b9050606481101580156115175750620186a08111155b6115335760405162461bcd60e51b8152600401610a4790614593565b60055561164f565b61159b84848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040805180820190915260098152686275726e526174696f60b81b60208201529150612fec9050565b1561163757602081146115c05760405162461bcd60e51b8152600401610a47906143e8565b604080516020601f84018190048102820181019092528281526000916115fe9185858083850183828082843760009201919091525061304592505050565b90506127108111156116225760405162461bcd60e51b8152600401610a4790614520565b600b55600c805460ff1916600117905561164f565b60405162461bcd60e51b8152600401610a479061494d565b7f6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a8484848460405161168494939291906143b6565b60405180910390a150505050565b60076020526000908152604090205481565b69152d02c7e14af680000081565b6004546060906000805b8281101561170357600481815481106116d157fe5b9060005260206000209060040201600201601c9054906101000a900460ff166116fb576001909101905b6001016116bc565b50606081604051908082528060200260200182016040528015611730578160200160208202803683370190505b50600092509050815b838110156117d0576004818154811061174e57fe5b9060005260206000209060040201600201601c9054906101000a900460ff166117c8576004818154811061177e57fe5b600091825260209091206004909102015482516001600160a01b03909116908390859081106117a957fe5b6001600160a01b03909216602092830291909101909101526001909201915b600101611739565b509250505090565b60086020526000908152604090205481565b61100281565b4790565b67016345785d8a000081565b600a5481565b61100381565b602981565b60005460ff16156118345760405162461bcd60e51b8152600401610a4790614749565b61183c613f5f565b600061185f604051806080016040528060478152602001614aa060479139611fa6565b91509150806118805760405162461bcd60e51b8152600401610a479061482d565b60005b8260200151518110156119a5576004836020015182815181106118a257fe5b6020908102919091018101518254600181810185556000948552838520835160049093020180546001600160a01b039384166001600160a01b03199182161782558486015182840180549186169183169190911790556040850151600283018054606088015160808901511515600160e01b0260ff60e01b196001600160401b03909216600160a01b0267ffffffffffffffff60a01b199590991692909516919091179290921695909517161790925560a090920151600390910155908501518051918401926007929091908590811061197857fe5b602090810291909101810151516001600160a01b0316825281019190915260400160002055600101611883565b505061045f600555506000805460ff19166001179055565b33611001146119de5760405162461bcd60e51b8152600401610a4790614974565b6001600160a01b03811660009081526007602052604090205480611a025750611013565b600181039050600060048281548110611a1757fe5b9060005260206000209060040201600301549050600060048381548110611a3a57fe5b90600052602060002090600402016003018190555060006001600480549050039050836001600160a01b03167f8cd4e147d8af98a9e3b6724021b8bf6aed2e5dac71c38f2dce8161b82585b25d83604051611a9591906149cc565b60405180910390a280611aaa57505050611013565b6000818381611ab557fe5b049050801561100e5760005b84811015611b13578160048281548110611ad757fe5b9060005260206000209060040201600301540160048281548110611af757fe5b6000918252602090912060036004909202010155600101611ac1565b50600454600185015b8181101561100b578260048281548110611b3257fe5b9060005260206000209060040201600301540160048281548110611b5257fe5b6000918252602090912060036004909202010155600101611b1c565b606581565b334114611b925760405162461bcd60e51b8152600401610a47906148bd565b60005460ff16611bb45760405162461bcd60e51b8152600401610a4790614456565b60003411611bd45760405162461bcd60e51b8152600401610a47906145da565b6001600160a01b038116600090815260076020526040812054600c5434929060ff1615611c005750600b545b600083118015611c105750600081115b15611cbd576000611c39612710611c2d868563ffffffff612ed516565b9063ffffffff612f1816565b90508015611cbb5760405161dead9082156108fc029083906000818181858888f19350505050158015611c70573d6000803e3d6000fd5b507f627059660ea01c4733a328effb2294d2f86905bf806da763a89cee254de8bee581604051611ca091906149cc565b60405180910390a1611cb8848263ffffffff61304a16565b93505b505b8115611dbc57600060046001840381548110611cd557fe5b9060005260206000209060040201905080600201601c9054906101000a900460ff1615611d4257846001600160a01b03167ff177e5d6c5764d79c32883ed824111d9b13f5668cf6ab1cc12dd36791dd955b485604051611d3591906149cc565b60405180910390a2611db6565b600654611d55908563ffffffff61308c16565b6006556003810154611d6d908563ffffffff61308c16565b60038201556040516001600160a01b038616907f93a090ecc682c002995fad3c85b30c5651d7fd29b0be5da9d784a3302aedc05590611dad9087906149cc565b60405180910390a25b50611dfe565b836001600160a01b03167ff177e5d6c5764d79c32883ed824111d9b13f5668cf6ab1cc12dd36791dd955b484604051611df591906149cc565b60405180910390a25b50505050565b620186a081565b61100081565b600381565b61dead81565b61100481565b80670de0b6b3a76400008110801590611e6057506001600160a01b03831660009081526008602052604090205469152d02c7e14af680000090820111155b611e7c5760405162461bcd60e51b8152600401610a4790614780565b6001600160a01b038316600090815260086020526040902054158015611eaa5750620186a060015460010111155b80611ecc57506001600160a01b03831660009081526008602052604090205415155b611ee85760405162461bcd60e51b8152600401610a47906144e9565b6001600160a01b038316600090815260086020526040902054611f4d5760018054600090815260096020908152604080832080546001600160a01b0319166001600160a01b038916908117909155835260089091529020829055805481019055611f6c565b6001600160a01b03831660009081526008602052604090208054820190555b60028054820190556040517f26a5c0f8e27c8ffa9a8414a94d6c0137a2ec1446b6925835dc0cf5ba74d2967f906113799085908490614258565b611fae613f5f565b6000611fb8613f5f565b611fc0613f8e565b611fd1611fcc866130b1565b6130d6565b90506000805b611fe083613120565b156120f1578061200557611ffb611ff684613141565b61318f565b60ff1684526120e9565b80600114156120e457606061202161201c85613141565b613211565b9050805160405190808252806020026020018201604052801561205e57816020015b61204b613fae565b8152602001906001900390816120435790505b50602086015260005b81518110156120d957612078613fae565b600061209684848151811061208957fe5b60200260200101516132e2565b91509150806120b3578760009950995050505050505050506120fa565b81886020015184815181106120c457fe5b60209081029190910101525050600101612067565b5060019250506120e9565b6120f1565b600101611fd7565b50919350909150505b915091565b604080516001808252818301909252606091829190816020015b606081526020019060019003908161211957905050905061213f8363ffffffff166133bf565b8160008151811061214c57fe5b6020026020010181905250610bab816133d2565b600080606061216e8461345c565b91509150816121b9577f70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2816040516121a691906143a3565b60405180910390a160669250505061127c565b600080805b6004548110156122365767016345785d8a0000600482815481106121de57fe5b906000526020600020906004020160030154106122005760019092019161222e565b60006004828154811061220f57fe5b906000526020600020906004020160030154111561222e576001909101905b6001016121be565b50606082604051908082528060200260200182016040528015612263578160200160208202803683370190505b509050606083604051908082528060200260200182016040528015612292578160200160208202803683370190505b5090506060846040519080825280602002602001820160405280156122c1578160200160208202803683370190505b5090506060856040519080825280602002602001820160405280156122f0578160200160208202803683370190505b5090506000606086604051908082528060200260200182016040528015612321578160200160208202803683370190505b509050606087604051908082528060200260200182016040528015612350578160200160208202803683370190505b509050600098506000975060608d905060006110046001600160a01b031663149d14d96040518163ffffffff1660e01b815260040160206040518083038186803b15801561239d57600080fd5b505afa1580156123b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123d5919061410a565b905067016345785d8a0000811115612432577f70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb260405161241490614708565b60405180910390a160689d505050505050505050505050505061127c565b60005b6004548110156126a55767016345785d8a00006004828154811061245557fe5b906000526020600020906004020160030154106125db576004818154811061247957fe5b906000526020600020906004020160020160009054906101000a90046001600160a01b03168a8d815181106124aa57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505060006402540be400600483815481106124df57fe5b906000526020600020906004020160030154816124f857fe5b066004838154811061250657fe5b90600052602060002090600402016003015403905061252e838261304a90919063ffffffff16565b8a8e8151811061253a57fe5b6020026020010181815250506004828154811061255357fe5b906000526020600020906004020160020160009054906101000a90046001600160a01b0316888e8151811061258457fe5b60200260200101906001600160a01b031690816001600160a01b03168152505081898e815181106125b157fe5b60209081029190910101526125cc878263ffffffff61308c16565b6001909d019c965061269d9050565b6000600482815481106125ea57fe5b906000526020600020906004020160030154111561269d576004818154811061260f57fe5b906000526020600020906004020160010160009054906101000a90046001600160a01b0316858c8151811061264057fe5b60200260200101906001600160a01b031690816001600160a01b0316815250506004818154811061266d57fe5b906000526020600020906004020160030154848c8151811061268b57fe5b60209081029190910101526001909a01995b600101612435565b506000851561291b576005546040516303702b2960e51b815261100491636e0565209189916126df918f918f918e914201906004016142c6565b6020604051808303818588803b1580156126f857600080fd5b505af193505050508015612729575060408051601f3d908101601f191682019092526127269181019061406a565b60015b6128a0576040516000815260443d1015612745575060006127e0565b60046000803e60005160e01c6308c379a081146127665760009150506127e0565b60043d036004833e81513d60248201116001600160401b0382111715612791576000925050506127e0565b80830180516001600160401b038111156127b25760009450505050506127e0565b8060208301013d86018111156127d0576000955050505050506127e0565b601f01601f191660405250925050505b806127eb575061282d565b60019150867fa7cdeed7d0db45e3219a6e5d60838824c16f1d39991fcfe3f963029c844bf2808260405161281f91906143a3565b60405180910390a25061289b565b3d808015612857576040519150601f19603f3d011682016040523d82523d6000602084013e61285c565b606091505b5060019150867fbfa884552dd8921b6ce90bfe906952ae5b3b29be0cc1a951d4f62697635a3a458260405161289191906143a3565b60405180910390a2505b61291b565b80156128e2577fa217d08e65f80c73121cd9db834d81652d544bfbf452f6d04922b16c90a37b70876040516128d591906149cc565b60405180910390a1612919565b867fa7cdeed7d0db45e3219a6e5d60838824c16f1d39991fcfe3f963029c844bf2806040516129109061441f565b60405180910390a25b505b8015612af85760005b8851811015612af657600089828151811061293b57fe5b6020026020010151905060006004828154811061295457fe5b906000526020600020906004020160010160009054906101000a90046001600160a01b03166001600160a01b03166108fc6004848154811061299257fe5b9060005260206000209060040201600301549081150290604051600060405180830381858888f1935050505090508015612a5b57600482815481106129d357fe5b906000526020600020906004020160010160009054906101000a90046001600160a01b03166001600160a01b03167f6c61d60f69a7beb3e1c80db7f39f37b208537cbb19da3174511b477812b2fc7d60048481548110612a2f57fe5b906000526020600020906004020160030154604051612a4e91906149cc565b60405180910390a2612aec565b60048281548110612a6857fe5b906000526020600020906004020160010160009054906101000a90046001600160a01b03166001600160a01b03167f25d0ce7d2f0cec669a8c17efe49d195c13455bb8872b65fa610ac7f53fe4ca7d60048481548110612ac457fe5b906000526020600020906004020160030154604051612ae391906149cc565b60405180910390a25b5050600101612924565b505b845115612c425760005b8551811015612c40576000868281518110612b1957fe5b60200260200101516001600160a01b03166108fc878481518110612b3957fe5b60200260200101519081150290604051600060405180830381858888f1935050505090508015612bcf57868281518110612b6f57fe5b60200260200101516001600160a01b03167f6c61d60f69a7beb3e1c80db7f39f37b208537cbb19da3174511b477812b2fc7d878481518110612bad57fe5b6020026020010151604051612bc291906149cc565b60405180910390a2612c37565b868281518110612bdb57fe5b60200260200101516001600160a01b03167f25d0ce7d2f0cec669a8c17efe49d195c13455bb8872b65fa610ac7f53fe4ca7d878481518110612c1957fe5b6020026020010151604051612c2e91906149cc565b60405180910390a25b50600101612b02565b505b4715612caf577f6ecc855f9440a9282c90913bbc91619fd44f5ec0b462af28d127b116f130aa4d47604051612c7791906149cc565b60405180910390a1604051611002904780156108fc02916000818181858888f19350505050158015612cad573d6000803e3d6000fd5b505b60006006819055600a55825115612cc957612cc98361353e565b6110016001600160a01b031663fc4333cd6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015612d0657600080fd5b505af1158015612d1a573d6000803e3d6000fd5b50506040517fedd8d7296956dd970ab4de3f2fc03be2b0ffc615d20cd4c72c6e44f928630ebf925060009150a15060009f9e505050505050505050505050505050565b80516001600160a01b0316600090815260076020526040812054801580612daf575060046001820381548110612d8f57fe5b9060005260206000209060040201600201601c9054906101000a900460ff165b15612df55782516040516001600160a01b03909116907fe209c46bebf57cf265d5d9009a00870e256d9150f3ed5281ab9d9eb3cec6e4be90600090a2600091505061127c565b600454600a54600019820111801590612e4b5784516040516001600160a01b03909116907fe209c46bebf57cf265d5d9009a00870e256d9150f3ed5281ab9d9eb3cec6e4be90600090a26000935050505061127c565b600a80546001908101909155600480546000198601908110612e6957fe5b6000918252602082206002600490920201018054921515600160e01b0260ff60e01b199093169290921790915585516040516001600160a01b03909116917ff226e7d8f547ff903d9d419cf5f54e0d7d07efa9584135a53a057c5f1f27f49a91a2506000949350505050565b600082612ee457506000612f12565b82820282848281612ef157fe5b0414612f0f5760405162461bcd60e51b8152600401610a4790614679565b90505b92915050565b6000612f0f83836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250613a00565b604080516000808252602082019092526001600160a01b038416908390604051612f849190614212565b60006040518083038185875af1925050503d8060008114612fc1576040519150601f19603f3d011682016040523d82523d6000602084013e612fc6565b606091505b5050905080612fe75760405162461bcd60e51b8152600401610a479061490a565b505050565b600081604051602001612fff9190614212565b60405160208183030381529060405280519060200120836040516020016130269190614212565b6040516020818303038152906040528051906020012014905092915050565b015190565b6000612f0f83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250613a37565b600082820183811015612f0f5760405162461bcd60e51b8152600401610a479061448d565b6130b9613fe3565b506040805180820190915281518152602082810190820152919050565b6130de613f8e565b6130e782613a63565b6130f057600080fd5b60006130ff8360200151613a9d565b60208085015160408051808201909152868152920190820152915050919050565b600061312a613fe3565b505080518051602091820151919092015191011190565b613149613fe3565b61315282613120565b61315b57600080fd5b6020820151600061316b82613b00565b80830160209586015260408051808201909152908152938401919091525090919050565b8051600090158015906131a457508151602110155b6131ad57600080fd5b60006131bc8360200151613a9d565b905080836000015110156131e25760405162461bcd60e51b8152600401610a47906147b0565b82516020808501518301805192849003929183101561320857826020036101000a820491505b50949350505050565b606061321c82613a63565b61322557600080fd5b600061323083613be1565b905060608160405190808252806020026020018201604052801561326e57816020015b61325b613fe3565b8152602001906001900390816132535790505b50905060006132808560200151613a9d565b60208601510190506000805b848110156132d75761329d83613b00565b91506040518060400160405280838152602001848152508482815181106132c057fe5b60209081029190910101529181019160010161328c565b509195945050505050565b6132ea613fae565b60006132f4613fae565b6132fc613f8e565b613305856130d6565b90506000805b61331483613120565b156120f1578061333f5761332f61332a84613141565b613c3d565b6001600160a01b031684526133b7565b80600114156133675761335461332a84613141565b6001600160a01b031660208501526133b7565b806002141561338f5761337c61332a84613141565b6001600160a01b031660408501526133b7565b80600314156120e4576133a4611ff684613141565b6001600160401b03166060850152600191505b60010161330b565b6060612f126133cd83613c57565b613d3d565b60608151600014156133f3575060408051600081526020810190915261127c565b60608260008151811061340257fe5b602002602001015190506000600190505b8351811015613443576134398285838151811061342c57fe5b6020026020010151613d8f565b9150600101613413565b50610bab613456825160c060ff16613e0c565b82613d8f565b6000606060298351111561348e576000604051806060016040528060298152602001614a4c60299139915091506120fa565b60005b83518110156135245760005b8181101561351b578481815181106134b157fe5b6020026020010151600001516001600160a01b03168583815181106134d257fe5b6020026020010151600001516001600160a01b031614156135135760006040518060600160405280602b8152602001614a75602b91399350935050506120fa565b60010161349d565b50600101613491565b505060408051602081019091526000815260019150915091565b600454815160005b8281101561365b576001613558613fae565b6004838154811061356557fe5b600091825260208083206040805160c08101825260049490940290910180546001600160a01b0390811685526001820154811693850193909352600281015492831691840191909152600160a01b82046001600160401b03166060840152600160e01b90910460ff16151560808301526003015460a082015291505b8481101561362f578681815181106135f557fe5b6020026020010151600001516001600160a01b031682600001516001600160a01b03161415613627576000925061362f565b6001016135e1565b5081156136515780516001600160a01b03166000908152600760205260408120555b5050600101613546565b50808211156136d057805b828110156136ce57600480548061367957fe5b60008281526020812060046000199093019283020180546001600160a01b03199081168255600182810180549092169091556002820180546001600160e81b0319169055600390910191909155915501613666565b505b60008183106136df57816136e1565b825b905060005b818110156138db576137938582815181106136fd57fe5b60200260200101516004838154811061371257fe5b60009182526020918290206040805160c08101825260049390930290910180546001600160a01b0390811684526001820154811694840194909452600281015493841691830191909152600160a01b83046001600160401b03166060830152600160e01b90920460ff161515608082015260039091015460a0820152613ede565b6138ae5780600101600760008784815181106137ab57fe5b6020026020010151600001516001600160a01b03166001600160a01b03168152602001908152602001600020819055508481815181106137e757fe5b6020026020010151600482815481106137fc57fe5b6000918252602091829020835160049092020180546001600160a01b039283166001600160a01b0319918216178255928401516001820180549184169185169190911790556040840151600282018054606087015160808801511515600160e01b0260ff60e01b196001600160401b03909216600160a01b0267ffffffffffffffff60a01b1995909716929097169190911792909216939093171692909217905560a0909101516003909101556138d3565b6000600482815481106138bd57fe5b9060005260206000209060040201600301819055505b6001016136e6565b5082821115611dfe57825b8281101561100e5760048582815181106138fc57fe5b6020908102919091018101518254600181810185556000948552838520835160049093020180546001600160a01b039384166001600160a01b03199182161782559484015181830180549185169187169190911790556040840151600282018054606087015160808801511515600160e01b0260ff60e01b196001600160401b03909216600160a01b0267ffffffffffffffff60a01b199590981692909916919091179290921694909417169490941790915560a0909101516003909201919091558651908301916007918890859081106139d357fe5b602090810291909101810151516001600160a01b03168252810191909152604001600020556001016138e6565b60008183613a215760405162461bcd60e51b8152600401610a4791906143a3565b506000838581613a2d57fe5b0495945050505050565b60008184841115613a5b5760405162461bcd60e51b8152600401610a4791906143a3565b505050900390565b8051600090613a745750600061127c565b6020820151805160001a9060c0821015613a935760009250505061127c565b5060019392505050565b8051600090811a6080811015613ab757600091505061127c565b60b8811080613ad2575060c08110801590613ad2575060f881105b15613ae157600191505061127c565b60c0811015613af55760b51901905061127c565b60f51901905061127c565b80516000908190811a6080811015613b1b5760019150613bda565b60b8811015613b3057607e1981019150613bda565b60c0811015613b8157600060b78203600186019550806020036101000a865104915060018101820193505080831015613b7b5760405162461bcd60e51b8152600401610a4790614609565b50613bda565b60f8811015613b965760be1981019150613bda565b600060f78203600186019550806020036101000a865104915060018101820193505080831015613bd85760405162461bcd60e51b8152600401610a4790614609565b505b5092915050565b8051600090613bf25750600061127c565b60008090506000613c068460200151613a9d565b602085015185519181019250015b80821015613c3457613c2582613b00565b60019093019290910190613c14565b50909392505050565b8051600090601514613c4e57600080fd5b612f128261318f565b604080516020808252818301909252606091829190602082018180368337505050602081018490529050600067ffffffffffffffff198416613c9b57506018613cbf565b6fffffffffffffffffffffffffffffffff198416613cbb57506010613cbf565b5060005b6020811015613cf557818181518110613cd457fe5b01602001516001600160f81b03191615613ced57613cf5565b600101613cbf565b60008160200390506060816040519080825280601f01601f191660200182016040528015613d2a576020820181803683370190505b5080830196909652508452509192915050565b606081516001148015613d6f5750607f60f81b82600081518110613d5d57fe5b01602001516001600160f81b03191611155b15613d7b57508061127c565b612f12613d8d8351608060ff16613e0c565b835b6060806040519050835180825260208201818101602087015b81831015613dc0578051835260209283019201613da8565b50855184518101855292509050808201602086015b81831015613ded578051835260209283019201613dd5565b508651929092011591909101601f01601f191660405250905092915050565b6060680100000000000000008310613e365760405162461bcd60e51b8152600401610a479061456b565b60408051600180825281830190925260609160208201818036833701905050905060378411613e905782840160f81b81600081518110613e7257fe5b60200101906001600160f81b031916908160001a9053509050612f12565b6060613e9b85613c57565b90508381510160370160f81b82600081518110613eb457fe5b60200101906001600160f81b031916908160001a905350613ed58282613d8f565b95945050505050565b805182516000916001600160a01b039182169116148015613f18575081602001516001600160a01b031683602001516001600160a01b0316145b8015613f3d575081604001516001600160a01b031683604001516001600160a01b0316145b8015612f0f5750506060908101519101516001600160401b0390811691161490565b60408051808201909152600081526060602082015290565b604080518082019091526000808252602082015290565b6040518060400160405280613fa1613fe3565b8152602001600081525090565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b604051806040016040528060008152602001600081525090565b60008083601f84011261400e578182fd5b5081356001600160401b03811115614024578182fd5b60208301915083602082850101111561403c57600080fd5b9250929050565b600060208284031215614054578081fd5b81356001600160a01b0381168114612f0f578182fd5b60006020828403121561407b578081fd5b81518015158114612f0f578182fd5b6000806000806040858703121561409f578283fd5b84356001600160401b03808211156140b5578485fd5b6140c188838901613ffd565b909650945060208701359150808211156140d9578384fd5b506140e687828801613ffd565b95989497509550505050565b600060208284031215614103578081fd5b5035919050565b60006020828403121561411b578081fd5b5051919050565b600080600060408486031215614136578283fd5b833560ff81168114614146578384fd5b925060208401356001600160401b03811115614160578283fd5b61416c86828701613ffd565b9497909650939450505050565b6000815180845260208085019450808401835b838110156141b15781516001600160a01b03168752958201959082019060010161418c565b509495945050505050565b60008284528282602086013780602084860101526020601f19601f85011685010190509392505050565b600081518084526141fe816020860160208601614a1f565b601f01601f19169290920160200192915050565b60008251614224818460208701614a1f565b9190910192915050565b9283526020830191909152604082015260600190565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b03968716815294861660208601529290941660408401526001600160401b03166060830152911515608082015260a081019190915260c00190565b600060208252612f0f6020830184614179565b6000608082526142d96080830187614179565b828103602084810191909152865180835287820192820190845b8181101561430f578451835293830193918301916001016142f3565b505084810360408601526143238188614179565b93505050506001600160401b038316606083015295945050505050565b602080825282518282018190526000919060409081850190868401855b8281101561438b57815180516001600160a01b0316855286015186850152928401929085019060010161435d565b5091979650505050505050565b901515815260200190565b600060208252612f0f60208301846141e6565b6000604082526143ca6040830186886141bc565b82810360208401526143dd8185876141bc565b979650505050505050565b6020808252601c908201527f6c656e677468206f66206275726e526174696f206d69736d6174636800000000604082015260600190565b6020808252601b908201527f6261746368207472616e736665722072657475726e2066616c73650000000000604082015260600190565b60208082526019908201527f74686520636f6e7472616374206e6f7420696e69742079657400000000000000604082015260600190565b6020808252601b908201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604082015260600190565b6020808252600b908201526a1b5d5cdd081cdd185ad95960aa1b604082015260600190565b60208082526019908201527f7374616b652070656f706c65206f7574206f662072616e676500000000000000604082015260600190565b6020808252602b908201527f746865206275726e526174696f206d757374206265206e6f206772656174657260408201526a0207468616e2031303030360ac1b606082015260800190565b6020808252600e908201526d696e70757420746f6f206c6f6e6760901b604082015260600190565b60208082526027908201527f7468652065787069726554696d655365636f6e64476170206973206f7574206f604082015266662072616e676560c81b606082015260800190565b6020808252601590820152746465706f7369742076616c7565206973207a65726f60581b604082015260600190565b6020808252601190820152706164646974696f6e206f766572666c6f7760781b604082015260600190565b60208082526025908201527f6c656e677468206f66206a61696c2076616c696461746f7273206d757374206260408201526465206f6e6560d81b606082015260800190565b60208082526021908201527f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f6040820152607760f81b606082015260800190565b6020808252602e908201527f746865206d6573736167652073656e646572206d75737420626520676f76657260408201526d1b985b98d94818dbdb9d1c9858dd60921b606082015260800190565b60208082526021908201527f666565206973206c6172676572207468616e2044555354595f494e434f4d494e6040820152604760f81b606082015260800190565b60208082526019908201527f74686520636f6e747261637420616c726561647920696e697400000000000000604082015260600190565b6020808252601690820152751cdd185ad94818dbdd5b9d081b5d5cdd081d985b1a5960521b604082015260600190565b6020808252601a908201527f6c656e677468206973206c657373207468616e206f6666736574000000000000604082015260600190565b60208082526026908201527f6c656e677468206f662065787069726554696d655365636f6e64476170206d696040820152650e6dac2e8c6d60d31b606082015260800190565b60208082526021908201527f6661696c656420746f20706172736520696e69742076616c696461746f7253656040820152601d60fa1b606082015260800190565b6020808252602f908201527f746865206d6573736167652073656e646572206d7573742062652063726f737360408201526e0818da185a5b8818dbdb9d1c9858dd608a1b606082015260800190565b6020808252602d908201527f746865206d6573736167652073656e646572206d75737420626520746865206260408201526c3637b1b590383937b23ab1b2b960991b606082015260800190565b60208082526023908201527f5472616e7366657248656c7065723a205452585f5452414e534645525f46414960408201526213115160ea1b606082015260800190565b6020808252600d908201526c756e6b6e6f776e20706172616d60981b604082015260600190565b60208082526029908201527f746865206d6573736167652073656e646572206d75737420626520736c6173686040820152680818dbdb9d1c9858dd60ba1b606082015260800190565b61ffff91909116815260200190565b90815260200190565b918252602082015260400190565b63ffffffff91909116815260200190565b60ff91909116815260200190565b600060ff8516825260406020830152613ed56040830184866141bc565b60005b83811015614a3a578181015183820152602001614a22565b83811115611dfe575050600091015256fe746865206e756d626572206f662076616c696461746f72732065786365656420746865206c696d69746475706c696361746520636f6e73656e7375732061646472657373206f662076616c696461746f72536574f84580f842f8409485245585de1768e6d40b46c5b4d0c17785ab7a9e9485245585de1768e6d40b46c5b4d0c17785ab7a9e9485245585de1768e6d40b46c5b4d0c17785ab7a9e64a2646970667358221220c64a92eedd9b2d1d07e501e4d674a285b8c329f2761ed3c86d6a0acb170cf30b64736f6c63430006040033"

// DeployBSCValidatorSet deploys a new Ethereum contract, binding an instance of BSCValidatorSet to it.
func DeployBSCValidatorSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BSCValidatorSet, error) {
	parsed, err := abi.JSON(strings.NewReader(BSCValidatorSetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BSCValidatorSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BSCValidatorSet{BSCValidatorSetCaller: BSCValidatorSetCaller{contract: contract}, BSCValidatorSetTransactor: BSCValidatorSetTransactor{contract: contract}, BSCValidatorSetFilterer: BSCValidatorSetFilterer{contract: contract}}, nil
}

// BSCValidatorSet is an auto generated Go binding around an Ethereum contract.
type BSCValidatorSet struct {
	BSCValidatorSetCaller     // Read-only binding to the contract
	BSCValidatorSetTransactor // Write-only binding to the contract
	BSCValidatorSetFilterer   // Log filterer for contract events
}

// BSCValidatorSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type BSCValidatorSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCValidatorSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BSCValidatorSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCValidatorSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BSCValidatorSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCValidatorSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BSCValidatorSetSession struct {
	Contract     *BSCValidatorSet  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BSCValidatorSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BSCValidatorSetCallerSession struct {
	Contract *BSCValidatorSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BSCValidatorSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BSCValidatorSetTransactorSession struct {
	Contract     *BSCValidatorSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BSCValidatorSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type BSCValidatorSetRaw struct {
	Contract *BSCValidatorSet // Generic contract binding to access the raw methods on
}

// BSCValidatorSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BSCValidatorSetCallerRaw struct {
	Contract *BSCValidatorSetCaller // Generic read-only contract binding to access the raw methods on
}

// BSCValidatorSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BSCValidatorSetTransactorRaw struct {
	Contract *BSCValidatorSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBSCValidatorSet creates a new instance of BSCValidatorSet, bound to a specific deployed contract.
func NewBSCValidatorSet(address common.Address, backend bind.ContractBackend) (*BSCValidatorSet, error) {
	contract, err := bindBSCValidatorSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSet{BSCValidatorSetCaller: BSCValidatorSetCaller{contract: contract}, BSCValidatorSetTransactor: BSCValidatorSetTransactor{contract: contract}, BSCValidatorSetFilterer: BSCValidatorSetFilterer{contract: contract}}, nil
}

// NewBSCValidatorSetCaller creates a new read-only instance of BSCValidatorSet, bound to a specific deployed contract.
func NewBSCValidatorSetCaller(address common.Address, caller bind.ContractCaller) (*BSCValidatorSetCaller, error) {
	contract, err := bindBSCValidatorSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetCaller{contract: contract}, nil
}

// NewBSCValidatorSetTransactor creates a new write-only instance of BSCValidatorSet, bound to a specific deployed contract.
func NewBSCValidatorSetTransactor(address common.Address, transactor bind.ContractTransactor) (*BSCValidatorSetTransactor, error) {
	contract, err := bindBSCValidatorSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetTransactor{contract: contract}, nil
}

// NewBSCValidatorSetFilterer creates a new log filterer instance of BSCValidatorSet, bound to a specific deployed contract.
func NewBSCValidatorSetFilterer(address common.Address, filterer bind.ContractFilterer) (*BSCValidatorSetFilterer, error) {
	contract, err := bindBSCValidatorSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetFilterer{contract: contract}, nil
}

// bindBSCValidatorSet binds a generic wrapper to an already deployed contract.
func bindBSCValidatorSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BSCValidatorSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BSCValidatorSet *BSCValidatorSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BSCValidatorSet.Contract.BSCValidatorSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BSCValidatorSet *BSCValidatorSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.BSCValidatorSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BSCValidatorSet *BSCValidatorSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.BSCValidatorSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BSCValidatorSet *BSCValidatorSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BSCValidatorSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BSCValidatorSet *BSCValidatorSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BSCValidatorSet *BSCValidatorSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.contract.Transact(opts, method, params...)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCaller) BINDCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "BIND_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetSession) BINDCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.BINDCHANNELID(&_BSCValidatorSet.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) BINDCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.BINDCHANNELID(&_BSCValidatorSet.CallOpts)
}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) BURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) BURNADDRESS() (common.Address, error) {
	return _BSCValidatorSet.Contract.BURNADDRESS(&_BSCValidatorSet.CallOpts)
}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) BURNADDRESS() (common.Address, error) {
	return _BSCValidatorSet.Contract.BURNADDRESS(&_BSCValidatorSet.CallOpts)
}

// BURNRATIOSCALE is a free data retrieval call binding the contract method 0x3de0f0d8.
//
// Solidity: function BURN_RATIO_SCALE() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) BURNRATIOSCALE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "BURN_RATIO_SCALE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BURNRATIOSCALE is a free data retrieval call binding the contract method 0x3de0f0d8.
//
// Solidity: function BURN_RATIO_SCALE() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) BURNRATIOSCALE() (*big.Int, error) {
	return _BSCValidatorSet.Contract.BURNRATIOSCALE(&_BSCValidatorSet.CallOpts)
}

// BURNRATIOSCALE is a free data retrieval call binding the contract method 0x3de0f0d8.
//
// Solidity: function BURN_RATIO_SCALE() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) BURNRATIOSCALE() (*big.Int, error) {
	return _BSCValidatorSet.Contract.BURNRATIOSCALE(&_BSCValidatorSet.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "CODE_OK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetSession) CODEOK() (uint32, error) {
	return _BSCValidatorSet.Contract.CODEOK(&_BSCValidatorSet.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) CODEOK() (uint32, error) {
	return _BSCValidatorSet.Contract.CODEOK(&_BSCValidatorSet.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "CROSS_CHAIN_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.CROSSCHAINCONTRACTADDR(&_BSCValidatorSet.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.CROSSCHAINCONTRACTADDR(&_BSCValidatorSet.CallOpts)
}

// DUSTYINCOMING is a free data retrieval call binding the contract method 0xd86222d5.
//
// Solidity: function DUSTY_INCOMING() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) DUSTYINCOMING(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "DUSTY_INCOMING")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DUSTYINCOMING is a free data retrieval call binding the contract method 0xd86222d5.
//
// Solidity: function DUSTY_INCOMING() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) DUSTYINCOMING() (*big.Int, error) {
	return _BSCValidatorSet.Contract.DUSTYINCOMING(&_BSCValidatorSet.CallOpts)
}

// DUSTYINCOMING is a free data retrieval call binding the contract method 0xd86222d5.
//
// Solidity: function DUSTY_INCOMING() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) DUSTYINCOMING() (*big.Int, error) {
	return _BSCValidatorSet.Contract.DUSTYINCOMING(&_BSCValidatorSet.CallOpts)
}

// ERRORFAILCHECKVALIDATORS is a free data retrieval call binding the contract method 0x81650b62.
//
// Solidity: function ERROR_FAIL_CHECK_VALIDATORS() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCaller) ERRORFAILCHECKVALIDATORS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "ERROR_FAIL_CHECK_VALIDATORS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORFAILCHECKVALIDATORS is a free data retrieval call binding the contract method 0x81650b62.
//
// Solidity: function ERROR_FAIL_CHECK_VALIDATORS() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetSession) ERRORFAILCHECKVALIDATORS() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORFAILCHECKVALIDATORS(&_BSCValidatorSet.CallOpts)
}

// ERRORFAILCHECKVALIDATORS is a free data retrieval call binding the contract method 0x81650b62.
//
// Solidity: function ERROR_FAIL_CHECK_VALIDATORS() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) ERRORFAILCHECKVALIDATORS() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORFAILCHECKVALIDATORS(&_BSCValidatorSet.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "ERROR_FAIL_DECODE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetSession) ERRORFAILDECODE() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORFAILDECODE(&_BSCValidatorSet.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORFAILDECODE(&_BSCValidatorSet.CallOpts)
}

// ERRORLENOFVALMISMATCH is a free data retrieval call binding the contract method 0x5d77156c.
//
// Solidity: function ERROR_LEN_OF_VAL_MISMATCH() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCaller) ERRORLENOFVALMISMATCH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "ERROR_LEN_OF_VAL_MISMATCH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORLENOFVALMISMATCH is a free data retrieval call binding the contract method 0x5d77156c.
//
// Solidity: function ERROR_LEN_OF_VAL_MISMATCH() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetSession) ERRORLENOFVALMISMATCH() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORLENOFVALMISMATCH(&_BSCValidatorSet.CallOpts)
}

// ERRORLENOFVALMISMATCH is a free data retrieval call binding the contract method 0x5d77156c.
//
// Solidity: function ERROR_LEN_OF_VAL_MISMATCH() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) ERRORLENOFVALMISMATCH() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORLENOFVALMISMATCH(&_BSCValidatorSet.CallOpts)
}

// ERRORRELAYFEETOOLARGE is a free data retrieval call binding the contract method 0x219f22d5.
//
// Solidity: function ERROR_RELAYFEE_TOO_LARGE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCaller) ERRORRELAYFEETOOLARGE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "ERROR_RELAYFEE_TOO_LARGE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORRELAYFEETOOLARGE is a free data retrieval call binding the contract method 0x219f22d5.
//
// Solidity: function ERROR_RELAYFEE_TOO_LARGE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetSession) ERRORRELAYFEETOOLARGE() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORRELAYFEETOOLARGE(&_BSCValidatorSet.CallOpts)
}

// ERRORRELAYFEETOOLARGE is a free data retrieval call binding the contract method 0x219f22d5.
//
// Solidity: function ERROR_RELAYFEE_TOO_LARGE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) ERRORRELAYFEETOOLARGE() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORRELAYFEETOOLARGE(&_BSCValidatorSet.CallOpts)
}

// ERRORUNKNOWNPACKAGETYPE is a free data retrieval call binding the contract method 0xeda5868c.
//
// Solidity: function ERROR_UNKNOWN_PACKAGE_TYPE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCaller) ERRORUNKNOWNPACKAGETYPE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "ERROR_UNKNOWN_PACKAGE_TYPE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORUNKNOWNPACKAGETYPE is a free data retrieval call binding the contract method 0xeda5868c.
//
// Solidity: function ERROR_UNKNOWN_PACKAGE_TYPE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetSession) ERRORUNKNOWNPACKAGETYPE() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORUNKNOWNPACKAGETYPE(&_BSCValidatorSet.CallOpts)
}

// ERRORUNKNOWNPACKAGETYPE is a free data retrieval call binding the contract method 0xeda5868c.
//
// Solidity: function ERROR_UNKNOWN_PACKAGE_TYPE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) ERRORUNKNOWNPACKAGETYPE() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORUNKNOWNPACKAGETYPE(&_BSCValidatorSet.CallOpts)
}

// EXPIRETIMESECONDGAP is a free data retrieval call binding the contract method 0x853230aa.
//
// Solidity: function EXPIRE_TIME_SECOND_GAP() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) EXPIRETIMESECONDGAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "EXPIRE_TIME_SECOND_GAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXPIRETIMESECONDGAP is a free data retrieval call binding the contract method 0x853230aa.
//
// Solidity: function EXPIRE_TIME_SECOND_GAP() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) EXPIRETIMESECONDGAP() (*big.Int, error) {
	return _BSCValidatorSet.Contract.EXPIRETIMESECONDGAP(&_BSCValidatorSet.CallOpts)
}

// EXPIRETIMESECONDGAP is a free data retrieval call binding the contract method 0x853230aa.
//
// Solidity: function EXPIRE_TIME_SECOND_GAP() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) EXPIRETIMESECONDGAP() (*big.Int, error) {
	return _BSCValidatorSet.Contract.EXPIRETIMESECONDGAP(&_BSCValidatorSet.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "GOV_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetSession) GOVCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.GOVCHANNELID(&_BSCValidatorSet.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) GOVCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.GOVCHANNELID(&_BSCValidatorSet.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) GOVHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "GOV_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) GOVHUBADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.GOVHUBADDR(&_BSCValidatorSet.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) GOVHUBADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.GOVHUBADDR(&_BSCValidatorSet.CallOpts)
}

// GetAllStakeInfo is a free data retrieval call binding the contract method 0x54b1a671.
//
// Solidity: function GetAllStakeInfo() view returns((address,uint256)[])
func (_BSCValidatorSet *BSCValidatorSetCaller) GetAllStakeInfo(opts *bind.CallOpts) ([]BSCValidatorSetStakeInfo, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "GetAllStakeInfo")

	if err != nil {
		return *new([]BSCValidatorSetStakeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]BSCValidatorSetStakeInfo)).(*[]BSCValidatorSetStakeInfo)

	return out0, err

}

// GetAllStakeInfo is a free data retrieval call binding the contract method 0x54b1a671.
//
// Solidity: function GetAllStakeInfo() view returns((address,uint256)[])
func (_BSCValidatorSet *BSCValidatorSetSession) GetAllStakeInfo() ([]BSCValidatorSetStakeInfo, error) {
	return _BSCValidatorSet.Contract.GetAllStakeInfo(&_BSCValidatorSet.CallOpts)
}

// GetAllStakeInfo is a free data retrieval call binding the contract method 0x54b1a671.
//
// Solidity: function GetAllStakeInfo() view returns((address,uint256)[])
func (_BSCValidatorSet *BSCValidatorSetCallerSession) GetAllStakeInfo() ([]BSCValidatorSetStakeInfo, error) {
	return _BSCValidatorSet.Contract.GetAllStakeInfo(&_BSCValidatorSet.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) INCENTIVIZEADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "INCENTIVIZE_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) INCENTIVIZEADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.INCENTIVIZEADDR(&_BSCValidatorSet.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) INCENTIVIZEADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.INCENTIVIZEADDR(&_BSCValidatorSet.CallOpts)
}

// INITBURNRATIO is a free data retrieval call binding the contract method 0x78dfed4a.
//
// Solidity: function INIT_BURN_RATIO() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) INITBURNRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "INIT_BURN_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITBURNRATIO is a free data retrieval call binding the contract method 0x78dfed4a.
//
// Solidity: function INIT_BURN_RATIO() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) INITBURNRATIO() (*big.Int, error) {
	return _BSCValidatorSet.Contract.INITBURNRATIO(&_BSCValidatorSet.CallOpts)
}

// INITBURNRATIO is a free data retrieval call binding the contract method 0x78dfed4a.
//
// Solidity: function INIT_BURN_RATIO() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) INITBURNRATIO() (*big.Int, error) {
	return _BSCValidatorSet.Contract.INITBURNRATIO(&_BSCValidatorSet.CallOpts)
}

// INITVALIDATORSETBYTES is a free data retrieval call binding the contract method 0xa5422d5c.
//
// Solidity: function INIT_VALIDATORSET_BYTES() view returns(bytes)
func (_BSCValidatorSet *BSCValidatorSetCaller) INITVALIDATORSETBYTES(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "INIT_VALIDATORSET_BYTES")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITVALIDATORSETBYTES is a free data retrieval call binding the contract method 0xa5422d5c.
//
// Solidity: function INIT_VALIDATORSET_BYTES() view returns(bytes)
func (_BSCValidatorSet *BSCValidatorSetSession) INITVALIDATORSETBYTES() ([]byte, error) {
	return _BSCValidatorSet.Contract.INITVALIDATORSETBYTES(&_BSCValidatorSet.CallOpts)
}

// INITVALIDATORSETBYTES is a free data retrieval call binding the contract method 0xa5422d5c.
//
// Solidity: function INIT_VALIDATORSET_BYTES() view returns(bytes)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) INITVALIDATORSETBYTES() ([]byte, error) {
	return _BSCValidatorSet.Contract.INITVALIDATORSETBYTES(&_BSCValidatorSet.CallOpts)
}

// JAILMESSAGETYPE is a free data retrieval call binding the contract method 0xbf9f4995.
//
// Solidity: function JAIL_MESSAGE_TYPE() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCaller) JAILMESSAGETYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "JAIL_MESSAGE_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// JAILMESSAGETYPE is a free data retrieval call binding the contract method 0xbf9f4995.
//
// Solidity: function JAIL_MESSAGE_TYPE() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetSession) JAILMESSAGETYPE() (uint8, error) {
	return _BSCValidatorSet.Contract.JAILMESSAGETYPE(&_BSCValidatorSet.CallOpts)
}

// JAILMESSAGETYPE is a free data retrieval call binding the contract method 0xbf9f4995.
//
// Solidity: function JAIL_MESSAGE_TYPE() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) JAILMESSAGETYPE() (uint8, error) {
	return _BSCValidatorSet.Contract.JAILMESSAGETYPE(&_BSCValidatorSet.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) LIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "LIGHT_CLIENT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.LIGHTCLIENTADDR(&_BSCValidatorSet.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.LIGHTCLIENTADDR(&_BSCValidatorSet.CallOpts)
}

// MAXNUMOFVALIDATORS is a free data retrieval call binding the contract method 0xe086c7b1.
//
// Solidity: function MAX_NUM_OF_VALIDATORS() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) MAXNUMOFVALIDATORS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "MAX_NUM_OF_VALIDATORS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXNUMOFVALIDATORS is a free data retrieval call binding the contract method 0xe086c7b1.
//
// Solidity: function MAX_NUM_OF_VALIDATORS() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) MAXNUMOFVALIDATORS() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MAXNUMOFVALIDATORS(&_BSCValidatorSet.CallOpts)
}

// MAXNUMOFVALIDATORS is a free data retrieval call binding the contract method 0xe086c7b1.
//
// Solidity: function MAX_NUM_OF_VALIDATORS() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) MAXNUMOFVALIDATORS() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MAXNUMOFVALIDATORS(&_BSCValidatorSet.CallOpts)
}

// MAXSTAKEFFFCOUNT is a free data retrieval call binding the contract method 0x8a10febe.
//
// Solidity: function MAX_STAKE_FFF_COUNT() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) MAXSTAKEFFFCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "MAX_STAKE_FFF_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSTAKEFFFCOUNT is a free data retrieval call binding the contract method 0x8a10febe.
//
// Solidity: function MAX_STAKE_FFF_COUNT() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) MAXSTAKEFFFCOUNT() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MAXSTAKEFFFCOUNT(&_BSCValidatorSet.CallOpts)
}

// MAXSTAKEFFFCOUNT is a free data retrieval call binding the contract method 0x8a10febe.
//
// Solidity: function MAX_STAKE_FFF_COUNT() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) MAXSTAKEFFFCOUNT() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MAXSTAKEFFFCOUNT(&_BSCValidatorSet.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) PRECISION() (*big.Int, error) {
	return _BSCValidatorSet.Contract.PRECISION(&_BSCValidatorSet.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) PRECISION() (*big.Int, error) {
	return _BSCValidatorSet.Contract.PRECISION(&_BSCValidatorSet.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "RELAYERHUB_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.RELAYERHUBCONTRACTADDR(&_BSCValidatorSet.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.RELAYERHUBCONTRACTADDR(&_BSCValidatorSet.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCaller) SLASHCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "SLASH_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetSession) SLASHCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.SLASHCHANNELID(&_BSCValidatorSet.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) SLASHCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.SLASHCHANNELID(&_BSCValidatorSet.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) SLASHCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "SLASH_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.SLASHCONTRACTADDR(&_BSCValidatorSet.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.SLASHCONTRACTADDR(&_BSCValidatorSet.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "STAKING_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetSession) STAKINGCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.STAKINGCHANNELID(&_BSCValidatorSet.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.STAKINGCHANNELID(&_BSCValidatorSet.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) SYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "SYSTEM_REWARD_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.SYSTEMREWARDADDR(&_BSCValidatorSet.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.SYSTEMREWARDADDR(&_BSCValidatorSet.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "TOKEN_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) TOKENHUBADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.TOKENHUBADDR(&_BSCValidatorSet.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.TOKENHUBADDR(&_BSCValidatorSet.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) TOKENMANAGERADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "TOKEN_MANAGER_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) TOKENMANAGERADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.TOKENMANAGERADDR(&_BSCValidatorSet.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) TOKENMANAGERADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.TOKENMANAGERADDR(&_BSCValidatorSet.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "TRANSFER_IN_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetSession) TRANSFERINCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.TRANSFERINCHANNELID(&_BSCValidatorSet.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.TRANSFERINCHANNELID(&_BSCValidatorSet.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "TRANSFER_OUT_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.TRANSFEROUTCHANNELID(&_BSCValidatorSet.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.TRANSFEROUTCHANNELID(&_BSCValidatorSet.CallOpts)
}

// VALIDATORSUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x5667515a.
//
// Solidity: function VALIDATORS_UPDATE_MESSAGE_TYPE() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCaller) VALIDATORSUPDATEMESSAGETYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "VALIDATORS_UPDATE_MESSAGE_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VALIDATORSUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x5667515a.
//
// Solidity: function VALIDATORS_UPDATE_MESSAGE_TYPE() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetSession) VALIDATORSUPDATEMESSAGETYPE() (uint8, error) {
	return _BSCValidatorSet.Contract.VALIDATORSUPDATEMESSAGETYPE(&_BSCValidatorSet.CallOpts)
}

// VALIDATORSUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x5667515a.
//
// Solidity: function VALIDATORS_UPDATE_MESSAGE_TYPE() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) VALIDATORSUPDATEMESSAGETYPE() (uint8, error) {
	return _BSCValidatorSet.Contract.VALIDATORSUPDATEMESSAGETYPE(&_BSCValidatorSet.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) VALIDATORCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "VALIDATOR_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.VALIDATORCONTRACTADDR(&_BSCValidatorSet.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.VALIDATORCONTRACTADDR(&_BSCValidatorSet.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_BSCValidatorSet *BSCValidatorSetCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "alreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_BSCValidatorSet *BSCValidatorSetSession) AlreadyInit() (bool, error) {
	return _BSCValidatorSet.Contract.AlreadyInit(&_BSCValidatorSet.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) AlreadyInit() (bool, error) {
	return _BSCValidatorSet.Contract.AlreadyInit(&_BSCValidatorSet.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_BSCValidatorSet *BSCValidatorSetCaller) BscChainID(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "bscChainID")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_BSCValidatorSet *BSCValidatorSetSession) BscChainID() (uint16, error) {
	return _BSCValidatorSet.Contract.BscChainID(&_BSCValidatorSet.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) BscChainID() (uint16, error) {
	return _BSCValidatorSet.Contract.BscChainID(&_BSCValidatorSet.CallOpts)
}

// BurnRatio is a free data retrieval call binding the contract method 0x5192c82c.
//
// Solidity: function burnRatio() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) BurnRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "burnRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnRatio is a free data retrieval call binding the contract method 0x5192c82c.
//
// Solidity: function burnRatio() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) BurnRatio() (*big.Int, error) {
	return _BSCValidatorSet.Contract.BurnRatio(&_BSCValidatorSet.CallOpts)
}

// BurnRatio is a free data retrieval call binding the contract method 0x5192c82c.
//
// Solidity: function burnRatio() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) BurnRatio() (*big.Int, error) {
	return _BSCValidatorSet.Contract.BurnRatio(&_BSCValidatorSet.CallOpts)
}

// BurnRatioInitialized is a free data retrieval call binding the contract method 0x152ad3b8.
//
// Solidity: function burnRatioInitialized() view returns(bool)
func (_BSCValidatorSet *BSCValidatorSetCaller) BurnRatioInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "burnRatioInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnRatioInitialized is a free data retrieval call binding the contract method 0x152ad3b8.
//
// Solidity: function burnRatioInitialized() view returns(bool)
func (_BSCValidatorSet *BSCValidatorSetSession) BurnRatioInitialized() (bool, error) {
	return _BSCValidatorSet.Contract.BurnRatioInitialized(&_BSCValidatorSet.CallOpts)
}

// BurnRatioInitialized is a free data retrieval call binding the contract method 0x152ad3b8.
//
// Solidity: function burnRatioInitialized() view returns(bool)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) BurnRatioInitialized() (bool, error) {
	return _BSCValidatorSet.Contract.BurnRatioInitialized(&_BSCValidatorSet.CallOpts)
}

// CurrStakeFFF is a free data retrieval call binding the contract method 0x101a875f.
//
// Solidity: function currStakeFFF() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) CurrStakeFFF(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "currStakeFFF")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrStakeFFF is a free data retrieval call binding the contract method 0x101a875f.
//
// Solidity: function currStakeFFF() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) CurrStakeFFF() (*big.Int, error) {
	return _BSCValidatorSet.Contract.CurrStakeFFF(&_BSCValidatorSet.CallOpts)
}

// CurrStakeFFF is a free data retrieval call binding the contract method 0x101a875f.
//
// Solidity: function currStakeFFF() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) CurrStakeFFF() (*big.Int, error) {
	return _BSCValidatorSet.Contract.CurrStakeFFF(&_BSCValidatorSet.CallOpts)
}

// CurrStakePeople is a free data retrieval call binding the contract method 0x93a101f3.
//
// Solidity: function currStakePeople() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) CurrStakePeople(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "currStakePeople")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrStakePeople is a free data retrieval call binding the contract method 0x93a101f3.
//
// Solidity: function currStakePeople() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) CurrStakePeople() (*big.Int, error) {
	return _BSCValidatorSet.Contract.CurrStakePeople(&_BSCValidatorSet.CallOpts)
}

// CurrStakePeople is a free data retrieval call binding the contract method 0x93a101f3.
//
// Solidity: function currStakePeople() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) CurrStakePeople() (*big.Int, error) {
	return _BSCValidatorSet.Contract.CurrStakePeople(&_BSCValidatorSet.CallOpts)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address consensusAddress, address feeAddress, address BBCFeeAddress, uint64 votingPower, bool jailed, uint256 incoming)
func (_BSCValidatorSet *BSCValidatorSetCaller) CurrentValidatorSet(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	BBCFeeAddress    common.Address
	VotingPower      uint64
	Jailed           bool
	Incoming         *big.Int
}, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "currentValidatorSet", arg0)

	outstruct := new(struct {
		ConsensusAddress common.Address
		FeeAddress       common.Address
		BBCFeeAddress    common.Address
		VotingPower      uint64
		Jailed           bool
		Incoming         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConsensusAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.FeeAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BBCFeeAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.VotingPower = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.Jailed = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Incoming = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address consensusAddress, address feeAddress, address BBCFeeAddress, uint64 votingPower, bool jailed, uint256 incoming)
func (_BSCValidatorSet *BSCValidatorSetSession) CurrentValidatorSet(arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	BBCFeeAddress    common.Address
	VotingPower      uint64
	Jailed           bool
	Incoming         *big.Int
}, error) {
	return _BSCValidatorSet.Contract.CurrentValidatorSet(&_BSCValidatorSet.CallOpts, arg0)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address consensusAddress, address feeAddress, address BBCFeeAddress, uint64 votingPower, bool jailed, uint256 incoming)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) CurrentValidatorSet(arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	BBCFeeAddress    common.Address
	VotingPower      uint64
	Jailed           bool
	Incoming         *big.Int
}, error) {
	return _BSCValidatorSet.Contract.CurrentValidatorSet(&_BSCValidatorSet.CallOpts, arg0)
}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) CurrentValidatorSetMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "currentValidatorSetMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) CurrentValidatorSetMap(arg0 common.Address) (*big.Int, error) {
	return _BSCValidatorSet.Contract.CurrentValidatorSetMap(&_BSCValidatorSet.CallOpts, arg0)
}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) CurrentValidatorSetMap(arg0 common.Address) (*big.Int, error) {
	return _BSCValidatorSet.Contract.CurrentValidatorSetMap(&_BSCValidatorSet.CallOpts, arg0)
}

// ExpireTimeSecondGap is a free data retrieval call binding the contract method 0x86249882.
//
// Solidity: function expireTimeSecondGap() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) ExpireTimeSecondGap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "expireTimeSecondGap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpireTimeSecondGap is a free data retrieval call binding the contract method 0x86249882.
//
// Solidity: function expireTimeSecondGap() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) ExpireTimeSecondGap() (*big.Int, error) {
	return _BSCValidatorSet.Contract.ExpireTimeSecondGap(&_BSCValidatorSet.CallOpts)
}

// ExpireTimeSecondGap is a free data retrieval call binding the contract method 0x86249882.
//
// Solidity: function expireTimeSecondGap() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) ExpireTimeSecondGap() (*big.Int, error) {
	return _BSCValidatorSet.Contract.ExpireTimeSecondGap(&_BSCValidatorSet.CallOpts)
}

// GetIncoming is a free data retrieval call binding the contract method 0x565c56b3.
//
// Solidity: function getIncoming(address validator) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) GetIncoming(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "getIncoming", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIncoming is a free data retrieval call binding the contract method 0x565c56b3.
//
// Solidity: function getIncoming(address validator) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) GetIncoming(validator common.Address) (*big.Int, error) {
	return _BSCValidatorSet.Contract.GetIncoming(&_BSCValidatorSet.CallOpts, validator)
}

// GetIncoming is a free data retrieval call binding the contract method 0x565c56b3.
//
// Solidity: function getIncoming(address validator) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) GetIncoming(validator common.Address) (*big.Int, error) {
	return _BSCValidatorSet.Contract.GetIncoming(&_BSCValidatorSet.CallOpts, validator)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_BSCValidatorSet *BSCValidatorSetCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_BSCValidatorSet *BSCValidatorSetSession) GetValidators() ([]common.Address, error) {
	return _BSCValidatorSet.Contract.GetValidators(&_BSCValidatorSet.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_BSCValidatorSet *BSCValidatorSetCallerSession) GetValidators() ([]common.Address, error) {
	return _BSCValidatorSet.Contract.GetValidators(&_BSCValidatorSet.CallOpts)
}

// MaxStakeFFFCount is a free data retrieval call binding the contract method 0xb2ccf3a4.
//
// Solidity: function maxStakeFFFCount() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) MaxStakeFFFCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "maxStakeFFFCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxStakeFFFCount is a free data retrieval call binding the contract method 0xb2ccf3a4.
//
// Solidity: function maxStakeFFFCount() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) MaxStakeFFFCount() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MaxStakeFFFCount(&_BSCValidatorSet.CallOpts)
}

// MaxStakeFFFCount is a free data retrieval call binding the contract method 0xb2ccf3a4.
//
// Solidity: function maxStakeFFFCount() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) MaxStakeFFFCount() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MaxStakeFFFCount(&_BSCValidatorSet.CallOpts)
}

// MaxStakePeopleCount is a free data retrieval call binding the contract method 0xf78ffcbc.
//
// Solidity: function maxStakePeopleCount() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) MaxStakePeopleCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "maxStakePeopleCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxStakePeopleCount is a free data retrieval call binding the contract method 0xf78ffcbc.
//
// Solidity: function maxStakePeopleCount() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) MaxStakePeopleCount() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MaxStakePeopleCount(&_BSCValidatorSet.CallOpts)
}

// MaxStakePeopleCount is a free data retrieval call binding the contract method 0xf78ffcbc.
//
// Solidity: function maxStakePeopleCount() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) MaxStakePeopleCount() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MaxStakePeopleCount(&_BSCValidatorSet.CallOpts)
}

// MiniStakeFFFCount is a free data retrieval call binding the contract method 0x1faea20b.
//
// Solidity: function miniStakeFFFCount() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) MiniStakeFFFCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "miniStakeFFFCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiniStakeFFFCount is a free data retrieval call binding the contract method 0x1faea20b.
//
// Solidity: function miniStakeFFFCount() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) MiniStakeFFFCount() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MiniStakeFFFCount(&_BSCValidatorSet.CallOpts)
}

// MiniStakeFFFCount is a free data retrieval call binding the contract method 0x1faea20b.
//
// Solidity: function miniStakeFFFCount() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) MiniStakeFFFCount() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MiniStakeFFFCount(&_BSCValidatorSet.CallOpts)
}

// MyBalance is a free data retrieval call binding the contract method 0xc9116b69.
//
// Solidity: function myBalance() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) MyBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "myBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MyBalance is a free data retrieval call binding the contract method 0xc9116b69.
//
// Solidity: function myBalance() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) MyBalance() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MyBalance(&_BSCValidatorSet.CallOpts)
}

// MyBalance is a free data retrieval call binding the contract method 0xc9116b69.
//
// Solidity: function myBalance() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) MyBalance() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MyBalance(&_BSCValidatorSet.CallOpts)
}

// NumOfJailed is a free data retrieval call binding the contract method 0xdaacdb66.
//
// Solidity: function numOfJailed() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) NumOfJailed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "numOfJailed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfJailed is a free data retrieval call binding the contract method 0xdaacdb66.
//
// Solidity: function numOfJailed() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) NumOfJailed() (*big.Int, error) {
	return _BSCValidatorSet.Contract.NumOfJailed(&_BSCValidatorSet.CallOpts)
}

// NumOfJailed is a free data retrieval call binding the contract method 0xdaacdb66.
//
// Solidity: function numOfJailed() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) NumOfJailed() (*big.Int, error) {
	return _BSCValidatorSet.Contract.NumOfJailed(&_BSCValidatorSet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) Owner() (common.Address, error) {
	return _BSCValidatorSet.Contract.Owner(&_BSCValidatorSet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) Owner() (common.Address, error) {
	return _BSCValidatorSet.Contract.Owner(&_BSCValidatorSet.CallOpts)
}

// StakeInfoIndexMap is a free data retrieval call binding the contract method 0x7dedf367.
//
// Solidity: function stakeInfoIndexMap(uint256 ) view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) StakeInfoIndexMap(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "stakeInfoIndexMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeInfoIndexMap is a free data retrieval call binding the contract method 0x7dedf367.
//
// Solidity: function stakeInfoIndexMap(uint256 ) view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) StakeInfoIndexMap(arg0 *big.Int) (common.Address, error) {
	return _BSCValidatorSet.Contract.StakeInfoIndexMap(&_BSCValidatorSet.CallOpts, arg0)
}

// StakeInfoIndexMap is a free data retrieval call binding the contract method 0x7dedf367.
//
// Solidity: function stakeInfoIndexMap(uint256 ) view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) StakeInfoIndexMap(arg0 *big.Int) (common.Address, error) {
	return _BSCValidatorSet.Contract.StakeInfoIndexMap(&_BSCValidatorSet.CallOpts, arg0)
}

// StakeInfoMap is a free data retrieval call binding the contract method 0xc5f30db5.
//
// Solidity: function stakeInfoMap(address ) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) StakeInfoMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "stakeInfoMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeInfoMap is a free data retrieval call binding the contract method 0xc5f30db5.
//
// Solidity: function stakeInfoMap(address ) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) StakeInfoMap(arg0 common.Address) (*big.Int, error) {
	return _BSCValidatorSet.Contract.StakeInfoMap(&_BSCValidatorSet.CallOpts, arg0)
}

// StakeInfoMap is a free data retrieval call binding the contract method 0xc5f30db5.
//
// Solidity: function stakeInfoMap(address ) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) StakeInfoMap(arg0 common.Address) (*big.Int, error) {
	return _BSCValidatorSet.Contract.StakeInfoMap(&_BSCValidatorSet.CallOpts, arg0)
}

// TotalInComing is a free data retrieval call binding the contract method 0x1ff18069.
//
// Solidity: function totalInComing() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) TotalInComing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "totalInComing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalInComing is a free data retrieval call binding the contract method 0x1ff18069.
//
// Solidity: function totalInComing() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) TotalInComing() (*big.Int, error) {
	return _BSCValidatorSet.Contract.TotalInComing(&_BSCValidatorSet.CallOpts)
}

// TotalInComing is a free data retrieval call binding the contract method 0x1ff18069.
//
// Solidity: function totalInComing() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) TotalInComing() (*big.Int, error) {
	return _BSCValidatorSet.Contract.TotalInComing(&_BSCValidatorSet.CallOpts)
}

// CreateAddress is a paid mutator transaction binding the contract method 0x2463cd09.
//
// Solidity: function CreateAddress(uint256 count) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) CreateAddress(opts *bind.TransactOpts, count *big.Int) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "CreateAddress", count)
}

// CreateAddress is a paid mutator transaction binding the contract method 0x2463cd09.
//
// Solidity: function CreateAddress(uint256 count) returns()
func (_BSCValidatorSet *BSCValidatorSetSession) CreateAddress(count *big.Int) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.CreateAddress(&_BSCValidatorSet.TransactOpts, count)
}

// CreateAddress is a paid mutator transaction binding the contract method 0x2463cd09.
//
// Solidity: function CreateAddress(uint256 count) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) CreateAddress(count *big.Int) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.CreateAddress(&_BSCValidatorSet.TransactOpts, count)
}

// DistributeBlockReward is a paid mutator transaction binding the contract method 0x50403540.
//
// Solidity: function DistributeBlockReward(uint256 reward) payable returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) DistributeBlockReward(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "DistributeBlockReward", reward)
}

// DistributeBlockReward is a paid mutator transaction binding the contract method 0x50403540.
//
// Solidity: function DistributeBlockReward(uint256 reward) payable returns()
func (_BSCValidatorSet *BSCValidatorSetSession) DistributeBlockReward(reward *big.Int) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.DistributeBlockReward(&_BSCValidatorSet.TransactOpts, reward)
}

// DistributeBlockReward is a paid mutator transaction binding the contract method 0x50403540.
//
// Solidity: function DistributeBlockReward(uint256 reward) payable returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) DistributeBlockReward(reward *big.Int) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.DistributeBlockReward(&_BSCValidatorSet.TransactOpts, reward)
}

// StakeFFF is a paid mutator transaction binding the contract method 0x0db6cccb.
//
// Solidity: function StakeFFF() payable returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) StakeFFF(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "StakeFFF")
}

// StakeFFF is a paid mutator transaction binding the contract method 0x0db6cccb.
//
// Solidity: function StakeFFF() payable returns()
func (_BSCValidatorSet *BSCValidatorSetSession) StakeFFF() (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.StakeFFF(&_BSCValidatorSet.TransactOpts)
}

// StakeFFF is a paid mutator transaction binding the contract method 0x0db6cccb.
//
// Solidity: function StakeFFF() payable returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) StakeFFF() (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.StakeFFF(&_BSCValidatorSet.TransactOpts)
}

// UnStakeFFF is a paid mutator transaction binding the contract method 0x30c62a13.
//
// Solidity: function UnStakeFFF() payable returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) UnStakeFFF(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "UnStakeFFF")
}

// UnStakeFFF is a paid mutator transaction binding the contract method 0x30c62a13.
//
// Solidity: function UnStakeFFF() payable returns()
func (_BSCValidatorSet *BSCValidatorSetSession) UnStakeFFF() (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.UnStakeFFF(&_BSCValidatorSet.TransactOpts)
}

// UnStakeFFF is a paid mutator transaction binding the contract method 0x30c62a13.
//
// Solidity: function UnStakeFFF() payable returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) UnStakeFFF() (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.UnStakeFFF(&_BSCValidatorSet.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address valAddr) payable returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) Deposit(opts *bind.TransactOpts, valAddr common.Address) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "deposit", valAddr)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address valAddr) payable returns()
func (_BSCValidatorSet *BSCValidatorSetSession) Deposit(valAddr common.Address) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.Deposit(&_BSCValidatorSet.TransactOpts, valAddr)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address valAddr) payable returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) Deposit(valAddr common.Address) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.Deposit(&_BSCValidatorSet.TransactOpts, valAddr)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) Felony(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "felony", validator)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_BSCValidatorSet *BSCValidatorSetSession) Felony(validator common.Address) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.Felony(&_BSCValidatorSet.TransactOpts, validator)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) Felony(validator common.Address) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.Felony(&_BSCValidatorSet.TransactOpts, validator)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) HandleAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "handleAckPackage", channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.HandleAckPackage(&_BSCValidatorSet.TransactOpts, channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.HandleAckPackage(&_BSCValidatorSet.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) HandleFailAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "handleFailAckPackage", channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.HandleFailAckPackage(&_BSCValidatorSet.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.HandleFailAckPackage(&_BSCValidatorSet.TransactOpts, channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_BSCValidatorSet *BSCValidatorSetTransactor) HandleSynPackage(opts *bind.TransactOpts, arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "handleSynPackage", arg0, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_BSCValidatorSet *BSCValidatorSetSession) HandleSynPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.HandleSynPackage(&_BSCValidatorSet.TransactOpts, arg0, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) HandleSynPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.HandleSynPackage(&_BSCValidatorSet.TransactOpts, arg0, msgBytes)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_BSCValidatorSet *BSCValidatorSetSession) Init() (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.Init(&_BSCValidatorSet.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) Init() (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.Init(&_BSCValidatorSet.TransactOpts)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) Misdemeanor(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "misdemeanor", validator)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_BSCValidatorSet *BSCValidatorSetSession) Misdemeanor(validator common.Address) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.Misdemeanor(&_BSCValidatorSet.TransactOpts, validator)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) Misdemeanor(validator common.Address) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.Misdemeanor(&_BSCValidatorSet.TransactOpts, validator)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_BSCValidatorSet *BSCValidatorSetSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.UpdateParam(&_BSCValidatorSet.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.UpdateParam(&_BSCValidatorSet.TransactOpts, key, value)
}

// BSCValidatorSetDistributeBlockRewardFuncIterator is returned from FilterDistributeBlockRewardFunc and is used to iterate over the raw logs and unpacked data for DistributeBlockRewardFunc events raised by the BSCValidatorSet contract.
type BSCValidatorSetDistributeBlockRewardFuncIterator struct {
	Event *BSCValidatorSetDistributeBlockRewardFunc // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetDistributeBlockRewardFuncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetDistributeBlockRewardFunc)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetDistributeBlockRewardFunc)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetDistributeBlockRewardFuncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetDistributeBlockRewardFuncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetDistributeBlockRewardFunc represents a DistributeBlockRewardFunc event raised by the BSCValidatorSet contract.
type BSCValidatorSetDistributeBlockRewardFunc struct {
	UserCount *big.Int
	Count     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDistributeBlockRewardFunc is a free log retrieval operation binding the contract event 0x77affbec109ff6a3609c802050db936c55bc48fb23db5c88f4a4818922e65239.
//
// Solidity: event DistributeBlockRewardFunc(uint256 userCount, uint256 count)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterDistributeBlockRewardFunc(opts *bind.FilterOpts) (*BSCValidatorSetDistributeBlockRewardFuncIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "DistributeBlockRewardFunc")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetDistributeBlockRewardFuncIterator{contract: _BSCValidatorSet.contract, event: "DistributeBlockRewardFunc", logs: logs, sub: sub}, nil
}

// WatchDistributeBlockRewardFunc is a free log subscription operation binding the contract event 0x77affbec109ff6a3609c802050db936c55bc48fb23db5c88f4a4818922e65239.
//
// Solidity: event DistributeBlockRewardFunc(uint256 userCount, uint256 count)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchDistributeBlockRewardFunc(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetDistributeBlockRewardFunc) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "DistributeBlockRewardFunc")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetDistributeBlockRewardFunc)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "DistributeBlockRewardFunc", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributeBlockRewardFunc is a log parse operation binding the contract event 0x77affbec109ff6a3609c802050db936c55bc48fb23db5c88f4a4818922e65239.
//
// Solidity: event DistributeBlockRewardFunc(uint256 userCount, uint256 count)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseDistributeBlockRewardFunc(log types.Log) (*BSCValidatorSetDistributeBlockRewardFunc, error) {
	event := new(BSCValidatorSetDistributeBlockRewardFunc)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "DistributeBlockRewardFunc", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetUserStakeFFFIterator is returned from FilterUserStakeFFF and is used to iterate over the raw logs and unpacked data for UserStakeFFF events raised by the BSCValidatorSet contract.
type BSCValidatorSetUserStakeFFFIterator struct {
	Event *BSCValidatorSetUserStakeFFF // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetUserStakeFFFIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetUserStakeFFF)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetUserStakeFFF)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetUserStakeFFFIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetUserStakeFFFIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetUserStakeFFF represents a UserStakeFFF event raised by the BSCValidatorSet contract.
type BSCValidatorSetUserStakeFFF struct {
	User  common.Address
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUserStakeFFF is a free log retrieval operation binding the contract event 0x26a5c0f8e27c8ffa9a8414a94d6c0137a2ec1446b6925835dc0cf5ba74d2967f.
//
// Solidity: event UserStakeFFF(address user, uint256 count)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterUserStakeFFF(opts *bind.FilterOpts) (*BSCValidatorSetUserStakeFFFIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "UserStakeFFF")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetUserStakeFFFIterator{contract: _BSCValidatorSet.contract, event: "UserStakeFFF", logs: logs, sub: sub}, nil
}

// WatchUserStakeFFF is a free log subscription operation binding the contract event 0x26a5c0f8e27c8ffa9a8414a94d6c0137a2ec1446b6925835dc0cf5ba74d2967f.
//
// Solidity: event UserStakeFFF(address user, uint256 count)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchUserStakeFFF(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetUserStakeFFF) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "UserStakeFFF")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetUserStakeFFF)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "UserStakeFFF", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserStakeFFF is a log parse operation binding the contract event 0x26a5c0f8e27c8ffa9a8414a94d6c0137a2ec1446b6925835dc0cf5ba74d2967f.
//
// Solidity: event UserStakeFFF(address user, uint256 count)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseUserStakeFFF(log types.Log) (*BSCValidatorSetUserStakeFFF, error) {
	event := new(BSCValidatorSetUserStakeFFF)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "UserStakeFFF", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetUserUnStakeFFFIterator is returned from FilterUserUnStakeFFF and is used to iterate over the raw logs and unpacked data for UserUnStakeFFF events raised by the BSCValidatorSet contract.
type BSCValidatorSetUserUnStakeFFFIterator struct {
	Event *BSCValidatorSetUserUnStakeFFF // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetUserUnStakeFFFIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetUserUnStakeFFF)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetUserUnStakeFFF)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetUserUnStakeFFFIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetUserUnStakeFFFIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetUserUnStakeFFF represents a UserUnStakeFFF event raised by the BSCValidatorSet contract.
type BSCValidatorSetUserUnStakeFFF struct {
	User  common.Address
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUserUnStakeFFF is a free log retrieval operation binding the contract event 0x433979128f828c623a175588cd9821124640d4353dd48d812d2b7273f1b76c3f.
//
// Solidity: event UserUnStakeFFF(address user, uint256 count)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterUserUnStakeFFF(opts *bind.FilterOpts) (*BSCValidatorSetUserUnStakeFFFIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "UserUnStakeFFF")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetUserUnStakeFFFIterator{contract: _BSCValidatorSet.contract, event: "UserUnStakeFFF", logs: logs, sub: sub}, nil
}

// WatchUserUnStakeFFF is a free log subscription operation binding the contract event 0x433979128f828c623a175588cd9821124640d4353dd48d812d2b7273f1b76c3f.
//
// Solidity: event UserUnStakeFFF(address user, uint256 count)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchUserUnStakeFFF(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetUserUnStakeFFF) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "UserUnStakeFFF")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetUserUnStakeFFF)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "UserUnStakeFFF", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserUnStakeFFF is a log parse operation binding the contract event 0x433979128f828c623a175588cd9821124640d4353dd48d812d2b7273f1b76c3f.
//
// Solidity: event UserUnStakeFFF(address user, uint256 count)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseUserUnStakeFFF(log types.Log) (*BSCValidatorSetUserUnStakeFFF, error) {
	event := new(BSCValidatorSetUserUnStakeFFF)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "UserUnStakeFFF", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetBatchTransferIterator is returned from FilterBatchTransfer and is used to iterate over the raw logs and unpacked data for BatchTransfer events raised by the BSCValidatorSet contract.
type BSCValidatorSetBatchTransferIterator struct {
	Event *BSCValidatorSetBatchTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetBatchTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetBatchTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetBatchTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetBatchTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetBatchTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetBatchTransfer represents a BatchTransfer event raised by the BSCValidatorSet contract.
type BSCValidatorSetBatchTransfer struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBatchTransfer is a free log retrieval operation binding the contract event 0xa217d08e65f80c73121cd9db834d81652d544bfbf452f6d04922b16c90a37b70.
//
// Solidity: event batchTransfer(uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterBatchTransfer(opts *bind.FilterOpts) (*BSCValidatorSetBatchTransferIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "batchTransfer")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetBatchTransferIterator{contract: _BSCValidatorSet.contract, event: "batchTransfer", logs: logs, sub: sub}, nil
}

// WatchBatchTransfer is a free log subscription operation binding the contract event 0xa217d08e65f80c73121cd9db834d81652d544bfbf452f6d04922b16c90a37b70.
//
// Solidity: event batchTransfer(uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchBatchTransfer(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetBatchTransfer) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "batchTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetBatchTransfer)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "batchTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchTransfer is a log parse operation binding the contract event 0xa217d08e65f80c73121cd9db834d81652d544bfbf452f6d04922b16c90a37b70.
//
// Solidity: event batchTransfer(uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseBatchTransfer(log types.Log) (*BSCValidatorSetBatchTransfer, error) {
	event := new(BSCValidatorSetBatchTransfer)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "batchTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetBatchTransferFailedIterator is returned from FilterBatchTransferFailed and is used to iterate over the raw logs and unpacked data for BatchTransferFailed events raised by the BSCValidatorSet contract.
type BSCValidatorSetBatchTransferFailedIterator struct {
	Event *BSCValidatorSetBatchTransferFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetBatchTransferFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetBatchTransferFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetBatchTransferFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetBatchTransferFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetBatchTransferFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetBatchTransferFailed represents a BatchTransferFailed event raised by the BSCValidatorSet contract.
type BSCValidatorSetBatchTransferFailed struct {
	Amount *big.Int
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBatchTransferFailed is a free log retrieval operation binding the contract event 0xa7cdeed7d0db45e3219a6e5d60838824c16f1d39991fcfe3f963029c844bf280.
//
// Solidity: event batchTransferFailed(uint256 indexed amount, string reason)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterBatchTransferFailed(opts *bind.FilterOpts, amount []*big.Int) (*BSCValidatorSetBatchTransferFailedIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "batchTransferFailed", amountRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetBatchTransferFailedIterator{contract: _BSCValidatorSet.contract, event: "batchTransferFailed", logs: logs, sub: sub}, nil
}

// WatchBatchTransferFailed is a free log subscription operation binding the contract event 0xa7cdeed7d0db45e3219a6e5d60838824c16f1d39991fcfe3f963029c844bf280.
//
// Solidity: event batchTransferFailed(uint256 indexed amount, string reason)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchBatchTransferFailed(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetBatchTransferFailed, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "batchTransferFailed", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetBatchTransferFailed)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "batchTransferFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchTransferFailed is a log parse operation binding the contract event 0xa7cdeed7d0db45e3219a6e5d60838824c16f1d39991fcfe3f963029c844bf280.
//
// Solidity: event batchTransferFailed(uint256 indexed amount, string reason)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseBatchTransferFailed(log types.Log) (*BSCValidatorSetBatchTransferFailed, error) {
	event := new(BSCValidatorSetBatchTransferFailed)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "batchTransferFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetBatchTransferLowerFailedIterator is returned from FilterBatchTransferLowerFailed and is used to iterate over the raw logs and unpacked data for BatchTransferLowerFailed events raised by the BSCValidatorSet contract.
type BSCValidatorSetBatchTransferLowerFailedIterator struct {
	Event *BSCValidatorSetBatchTransferLowerFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetBatchTransferLowerFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetBatchTransferLowerFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetBatchTransferLowerFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetBatchTransferLowerFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetBatchTransferLowerFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetBatchTransferLowerFailed represents a BatchTransferLowerFailed event raised by the BSCValidatorSet contract.
type BSCValidatorSetBatchTransferLowerFailed struct {
	Amount *big.Int
	Reason []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBatchTransferLowerFailed is a free log retrieval operation binding the contract event 0xbfa884552dd8921b6ce90bfe906952ae5b3b29be0cc1a951d4f62697635a3a45.
//
// Solidity: event batchTransferLowerFailed(uint256 indexed amount, bytes reason)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterBatchTransferLowerFailed(opts *bind.FilterOpts, amount []*big.Int) (*BSCValidatorSetBatchTransferLowerFailedIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "batchTransferLowerFailed", amountRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetBatchTransferLowerFailedIterator{contract: _BSCValidatorSet.contract, event: "batchTransferLowerFailed", logs: logs, sub: sub}, nil
}

// WatchBatchTransferLowerFailed is a free log subscription operation binding the contract event 0xbfa884552dd8921b6ce90bfe906952ae5b3b29be0cc1a951d4f62697635a3a45.
//
// Solidity: event batchTransferLowerFailed(uint256 indexed amount, bytes reason)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchBatchTransferLowerFailed(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetBatchTransferLowerFailed, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "batchTransferLowerFailed", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetBatchTransferLowerFailed)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "batchTransferLowerFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchTransferLowerFailed is a log parse operation binding the contract event 0xbfa884552dd8921b6ce90bfe906952ae5b3b29be0cc1a951d4f62697635a3a45.
//
// Solidity: event batchTransferLowerFailed(uint256 indexed amount, bytes reason)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseBatchTransferLowerFailed(log types.Log) (*BSCValidatorSetBatchTransferLowerFailed, error) {
	event := new(BSCValidatorSetBatchTransferLowerFailed)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "batchTransferLowerFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetDeprecatedDepositIterator is returned from FilterDeprecatedDeposit and is used to iterate over the raw logs and unpacked data for DeprecatedDeposit events raised by the BSCValidatorSet contract.
type BSCValidatorSetDeprecatedDepositIterator struct {
	Event *BSCValidatorSetDeprecatedDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetDeprecatedDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetDeprecatedDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetDeprecatedDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetDeprecatedDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetDeprecatedDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetDeprecatedDeposit represents a DeprecatedDeposit event raised by the BSCValidatorSet contract.
type BSCValidatorSetDeprecatedDeposit struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeprecatedDeposit is a free log retrieval operation binding the contract event 0xf177e5d6c5764d79c32883ed824111d9b13f5668cf6ab1cc12dd36791dd955b4.
//
// Solidity: event deprecatedDeposit(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterDeprecatedDeposit(opts *bind.FilterOpts, validator []common.Address) (*BSCValidatorSetDeprecatedDepositIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "deprecatedDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetDeprecatedDepositIterator{contract: _BSCValidatorSet.contract, event: "deprecatedDeposit", logs: logs, sub: sub}, nil
}

// WatchDeprecatedDeposit is a free log subscription operation binding the contract event 0xf177e5d6c5764d79c32883ed824111d9b13f5668cf6ab1cc12dd36791dd955b4.
//
// Solidity: event deprecatedDeposit(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchDeprecatedDeposit(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetDeprecatedDeposit, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "deprecatedDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetDeprecatedDeposit)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "deprecatedDeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeprecatedDeposit is a log parse operation binding the contract event 0xf177e5d6c5764d79c32883ed824111d9b13f5668cf6ab1cc12dd36791dd955b4.
//
// Solidity: event deprecatedDeposit(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseDeprecatedDeposit(log types.Log) (*BSCValidatorSetDeprecatedDeposit, error) {
	event := new(BSCValidatorSetDeprecatedDeposit)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "deprecatedDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetDirectTransferIterator is returned from FilterDirectTransfer and is used to iterate over the raw logs and unpacked data for DirectTransfer events raised by the BSCValidatorSet contract.
type BSCValidatorSetDirectTransferIterator struct {
	Event *BSCValidatorSetDirectTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetDirectTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetDirectTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetDirectTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetDirectTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetDirectTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetDirectTransfer represents a DirectTransfer event raised by the BSCValidatorSet contract.
type BSCValidatorSetDirectTransfer struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDirectTransfer is a free log retrieval operation binding the contract event 0x6c61d60f69a7beb3e1c80db7f39f37b208537cbb19da3174511b477812b2fc7d.
//
// Solidity: event directTransfer(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterDirectTransfer(opts *bind.FilterOpts, validator []common.Address) (*BSCValidatorSetDirectTransferIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "directTransfer", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetDirectTransferIterator{contract: _BSCValidatorSet.contract, event: "directTransfer", logs: logs, sub: sub}, nil
}

// WatchDirectTransfer is a free log subscription operation binding the contract event 0x6c61d60f69a7beb3e1c80db7f39f37b208537cbb19da3174511b477812b2fc7d.
//
// Solidity: event directTransfer(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchDirectTransfer(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetDirectTransfer, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "directTransfer", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetDirectTransfer)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "directTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDirectTransfer is a log parse operation binding the contract event 0x6c61d60f69a7beb3e1c80db7f39f37b208537cbb19da3174511b477812b2fc7d.
//
// Solidity: event directTransfer(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseDirectTransfer(log types.Log) (*BSCValidatorSetDirectTransfer, error) {
	event := new(BSCValidatorSetDirectTransfer)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "directTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetDirectTransferFailIterator is returned from FilterDirectTransferFail and is used to iterate over the raw logs and unpacked data for DirectTransferFail events raised by the BSCValidatorSet contract.
type BSCValidatorSetDirectTransferFailIterator struct {
	Event *BSCValidatorSetDirectTransferFail // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetDirectTransferFailIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetDirectTransferFail)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetDirectTransferFail)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetDirectTransferFailIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetDirectTransferFailIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetDirectTransferFail represents a DirectTransferFail event raised by the BSCValidatorSet contract.
type BSCValidatorSetDirectTransferFail struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDirectTransferFail is a free log retrieval operation binding the contract event 0x25d0ce7d2f0cec669a8c17efe49d195c13455bb8872b65fa610ac7f53fe4ca7d.
//
// Solidity: event directTransferFail(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterDirectTransferFail(opts *bind.FilterOpts, validator []common.Address) (*BSCValidatorSetDirectTransferFailIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "directTransferFail", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetDirectTransferFailIterator{contract: _BSCValidatorSet.contract, event: "directTransferFail", logs: logs, sub: sub}, nil
}

// WatchDirectTransferFail is a free log subscription operation binding the contract event 0x25d0ce7d2f0cec669a8c17efe49d195c13455bb8872b65fa610ac7f53fe4ca7d.
//
// Solidity: event directTransferFail(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchDirectTransferFail(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetDirectTransferFail, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "directTransferFail", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetDirectTransferFail)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "directTransferFail", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDirectTransferFail is a log parse operation binding the contract event 0x25d0ce7d2f0cec669a8c17efe49d195c13455bb8872b65fa610ac7f53fe4ca7d.
//
// Solidity: event directTransferFail(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseDirectTransferFail(log types.Log) (*BSCValidatorSetDirectTransferFail, error) {
	event := new(BSCValidatorSetDirectTransferFail)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "directTransferFail", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetFailReasonWithStrIterator is returned from FilterFailReasonWithStr and is used to iterate over the raw logs and unpacked data for FailReasonWithStr events raised by the BSCValidatorSet contract.
type BSCValidatorSetFailReasonWithStrIterator struct {
	Event *BSCValidatorSetFailReasonWithStr // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetFailReasonWithStrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetFailReasonWithStr)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetFailReasonWithStr)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetFailReasonWithStrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetFailReasonWithStrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetFailReasonWithStr represents a FailReasonWithStr event raised by the BSCValidatorSet contract.
type BSCValidatorSetFailReasonWithStr struct {
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFailReasonWithStr is a free log retrieval operation binding the contract event 0x70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2.
//
// Solidity: event failReasonWithStr(string message)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterFailReasonWithStr(opts *bind.FilterOpts) (*BSCValidatorSetFailReasonWithStrIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "failReasonWithStr")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetFailReasonWithStrIterator{contract: _BSCValidatorSet.contract, event: "failReasonWithStr", logs: logs, sub: sub}, nil
}

// WatchFailReasonWithStr is a free log subscription operation binding the contract event 0x70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2.
//
// Solidity: event failReasonWithStr(string message)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchFailReasonWithStr(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetFailReasonWithStr) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "failReasonWithStr")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetFailReasonWithStr)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "failReasonWithStr", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFailReasonWithStr is a log parse operation binding the contract event 0x70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2.
//
// Solidity: event failReasonWithStr(string message)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseFailReasonWithStr(log types.Log) (*BSCValidatorSetFailReasonWithStr, error) {
	event := new(BSCValidatorSetFailReasonWithStr)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "failReasonWithStr", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetFeeBurnedIterator is returned from FilterFeeBurned and is used to iterate over the raw logs and unpacked data for FeeBurned events raised by the BSCValidatorSet contract.
type BSCValidatorSetFeeBurnedIterator struct {
	Event *BSCValidatorSetFeeBurned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetFeeBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetFeeBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetFeeBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetFeeBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetFeeBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetFeeBurned represents a FeeBurned event raised by the BSCValidatorSet contract.
type BSCValidatorSetFeeBurned struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeBurned is a free log retrieval operation binding the contract event 0x627059660ea01c4733a328effb2294d2f86905bf806da763a89cee254de8bee5.
//
// Solidity: event feeBurned(uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterFeeBurned(opts *bind.FilterOpts) (*BSCValidatorSetFeeBurnedIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "feeBurned")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetFeeBurnedIterator{contract: _BSCValidatorSet.contract, event: "feeBurned", logs: logs, sub: sub}, nil
}

// WatchFeeBurned is a free log subscription operation binding the contract event 0x627059660ea01c4733a328effb2294d2f86905bf806da763a89cee254de8bee5.
//
// Solidity: event feeBurned(uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchFeeBurned(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetFeeBurned) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "feeBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetFeeBurned)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "feeBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeBurned is a log parse operation binding the contract event 0x627059660ea01c4733a328effb2294d2f86905bf806da763a89cee254de8bee5.
//
// Solidity: event feeBurned(uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseFeeBurned(log types.Log) (*BSCValidatorSetFeeBurned, error) {
	event := new(BSCValidatorSetFeeBurned)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "feeBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the BSCValidatorSet contract.
type BSCValidatorSetParamChangeIterator struct {
	Event *BSCValidatorSetParamChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetParamChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetParamChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetParamChange represents a ParamChange event raised by the BSCValidatorSet contract.
type BSCValidatorSetParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterParamChange(opts *bind.FilterOpts) (*BSCValidatorSetParamChangeIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetParamChangeIterator{contract: _BSCValidatorSet.contract, event: "paramChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetParamChange) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetParamChange)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "paramChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParamChange is a log parse operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseParamChange(log types.Log) (*BSCValidatorSetParamChange, error) {
	event := new(BSCValidatorSetParamChange)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "paramChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetSystemTransferIterator is returned from FilterSystemTransfer and is used to iterate over the raw logs and unpacked data for SystemTransfer events raised by the BSCValidatorSet contract.
type BSCValidatorSetSystemTransferIterator struct {
	Event *BSCValidatorSetSystemTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetSystemTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetSystemTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetSystemTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetSystemTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetSystemTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetSystemTransfer represents a SystemTransfer event raised by the BSCValidatorSet contract.
type BSCValidatorSetSystemTransfer struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSystemTransfer is a free log retrieval operation binding the contract event 0x6ecc855f9440a9282c90913bbc91619fd44f5ec0b462af28d127b116f130aa4d.
//
// Solidity: event systemTransfer(uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterSystemTransfer(opts *bind.FilterOpts) (*BSCValidatorSetSystemTransferIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "systemTransfer")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetSystemTransferIterator{contract: _BSCValidatorSet.contract, event: "systemTransfer", logs: logs, sub: sub}, nil
}

// WatchSystemTransfer is a free log subscription operation binding the contract event 0x6ecc855f9440a9282c90913bbc91619fd44f5ec0b462af28d127b116f130aa4d.
//
// Solidity: event systemTransfer(uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchSystemTransfer(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetSystemTransfer) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "systemTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetSystemTransfer)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "systemTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSystemTransfer is a log parse operation binding the contract event 0x6ecc855f9440a9282c90913bbc91619fd44f5ec0b462af28d127b116f130aa4d.
//
// Solidity: event systemTransfer(uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseSystemTransfer(log types.Log) (*BSCValidatorSetSystemTransfer, error) {
	event := new(BSCValidatorSetSystemTransfer)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "systemTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetUnexpectedPackageIterator is returned from FilterUnexpectedPackage and is used to iterate over the raw logs and unpacked data for UnexpectedPackage events raised by the BSCValidatorSet contract.
type BSCValidatorSetUnexpectedPackageIterator struct {
	Event *BSCValidatorSetUnexpectedPackage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetUnexpectedPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetUnexpectedPackage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetUnexpectedPackage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetUnexpectedPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetUnexpectedPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetUnexpectedPackage represents a UnexpectedPackage event raised by the BSCValidatorSet contract.
type BSCValidatorSetUnexpectedPackage struct {
	ChannelId uint8
	MsgBytes  []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnexpectedPackage is a free log retrieval operation binding the contract event 0x41ce201247b6ceb957dcdb217d0b8acb50b9ea0e12af9af4f5e7f38902101605.
//
// Solidity: event unexpectedPackage(uint8 channelId, bytes msgBytes)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterUnexpectedPackage(opts *bind.FilterOpts) (*BSCValidatorSetUnexpectedPackageIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "unexpectedPackage")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetUnexpectedPackageIterator{contract: _BSCValidatorSet.contract, event: "unexpectedPackage", logs: logs, sub: sub}, nil
}

// WatchUnexpectedPackage is a free log subscription operation binding the contract event 0x41ce201247b6ceb957dcdb217d0b8acb50b9ea0e12af9af4f5e7f38902101605.
//
// Solidity: event unexpectedPackage(uint8 channelId, bytes msgBytes)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchUnexpectedPackage(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetUnexpectedPackage) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "unexpectedPackage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetUnexpectedPackage)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "unexpectedPackage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnexpectedPackage is a log parse operation binding the contract event 0x41ce201247b6ceb957dcdb217d0b8acb50b9ea0e12af9af4f5e7f38902101605.
//
// Solidity: event unexpectedPackage(uint8 channelId, bytes msgBytes)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseUnexpectedPackage(log types.Log) (*BSCValidatorSetUnexpectedPackage, error) {
	event := new(BSCValidatorSetUnexpectedPackage)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "unexpectedPackage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetValidatorDepositIterator is returned from FilterValidatorDeposit and is used to iterate over the raw logs and unpacked data for ValidatorDeposit events raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorDepositIterator struct {
	Event *BSCValidatorSetValidatorDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetValidatorDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetValidatorDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetValidatorDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetValidatorDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetValidatorDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetValidatorDeposit represents a ValidatorDeposit event raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorDeposit struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorDeposit is a free log retrieval operation binding the contract event 0x93a090ecc682c002995fad3c85b30c5651d7fd29b0be5da9d784a3302aedc055.
//
// Solidity: event validatorDeposit(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterValidatorDeposit(opts *bind.FilterOpts, validator []common.Address) (*BSCValidatorSetValidatorDepositIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "validatorDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetValidatorDepositIterator{contract: _BSCValidatorSet.contract, event: "validatorDeposit", logs: logs, sub: sub}, nil
}

// WatchValidatorDeposit is a free log subscription operation binding the contract event 0x93a090ecc682c002995fad3c85b30c5651d7fd29b0be5da9d784a3302aedc055.
//
// Solidity: event validatorDeposit(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchValidatorDeposit(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetValidatorDeposit, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "validatorDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetValidatorDeposit)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorDeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorDeposit is a log parse operation binding the contract event 0x93a090ecc682c002995fad3c85b30c5651d7fd29b0be5da9d784a3302aedc055.
//
// Solidity: event validatorDeposit(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseValidatorDeposit(log types.Log) (*BSCValidatorSetValidatorDeposit, error) {
	event := new(BSCValidatorSetValidatorDeposit)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetValidatorEmptyJailedIterator is returned from FilterValidatorEmptyJailed and is used to iterate over the raw logs and unpacked data for ValidatorEmptyJailed events raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorEmptyJailedIterator struct {
	Event *BSCValidatorSetValidatorEmptyJailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetValidatorEmptyJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetValidatorEmptyJailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetValidatorEmptyJailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetValidatorEmptyJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetValidatorEmptyJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetValidatorEmptyJailed represents a ValidatorEmptyJailed event raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorEmptyJailed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorEmptyJailed is a free log retrieval operation binding the contract event 0xe209c46bebf57cf265d5d9009a00870e256d9150f3ed5281ab9d9eb3cec6e4be.
//
// Solidity: event validatorEmptyJailed(address indexed validator)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterValidatorEmptyJailed(opts *bind.FilterOpts, validator []common.Address) (*BSCValidatorSetValidatorEmptyJailedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "validatorEmptyJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetValidatorEmptyJailedIterator{contract: _BSCValidatorSet.contract, event: "validatorEmptyJailed", logs: logs, sub: sub}, nil
}

// WatchValidatorEmptyJailed is a free log subscription operation binding the contract event 0xe209c46bebf57cf265d5d9009a00870e256d9150f3ed5281ab9d9eb3cec6e4be.
//
// Solidity: event validatorEmptyJailed(address indexed validator)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchValidatorEmptyJailed(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetValidatorEmptyJailed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "validatorEmptyJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetValidatorEmptyJailed)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorEmptyJailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorEmptyJailed is a log parse operation binding the contract event 0xe209c46bebf57cf265d5d9009a00870e256d9150f3ed5281ab9d9eb3cec6e4be.
//
// Solidity: event validatorEmptyJailed(address indexed validator)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseValidatorEmptyJailed(log types.Log) (*BSCValidatorSetValidatorEmptyJailed, error) {
	event := new(BSCValidatorSetValidatorEmptyJailed)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorEmptyJailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetValidatorFelonyIterator is returned from FilterValidatorFelony and is used to iterate over the raw logs and unpacked data for ValidatorFelony events raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorFelonyIterator struct {
	Event *BSCValidatorSetValidatorFelony // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetValidatorFelonyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetValidatorFelony)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetValidatorFelony)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetValidatorFelonyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetValidatorFelonyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetValidatorFelony represents a ValidatorFelony event raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorFelony struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorFelony is a free log retrieval operation binding the contract event 0x3b6f9ef90462b512a1293ecec018670bf7b7f1876fb727590a8a6d7643130a70.
//
// Solidity: event validatorFelony(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterValidatorFelony(opts *bind.FilterOpts, validator []common.Address) (*BSCValidatorSetValidatorFelonyIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "validatorFelony", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetValidatorFelonyIterator{contract: _BSCValidatorSet.contract, event: "validatorFelony", logs: logs, sub: sub}, nil
}

// WatchValidatorFelony is a free log subscription operation binding the contract event 0x3b6f9ef90462b512a1293ecec018670bf7b7f1876fb727590a8a6d7643130a70.
//
// Solidity: event validatorFelony(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchValidatorFelony(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetValidatorFelony, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "validatorFelony", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetValidatorFelony)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorFelony", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorFelony is a log parse operation binding the contract event 0x3b6f9ef90462b512a1293ecec018670bf7b7f1876fb727590a8a6d7643130a70.
//
// Solidity: event validatorFelony(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseValidatorFelony(log types.Log) (*BSCValidatorSetValidatorFelony, error) {
	event := new(BSCValidatorSetValidatorFelony)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorFelony", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetValidatorJailedIterator is returned from FilterValidatorJailed and is used to iterate over the raw logs and unpacked data for ValidatorJailed events raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorJailedIterator struct {
	Event *BSCValidatorSetValidatorJailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetValidatorJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetValidatorJailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetValidatorJailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetValidatorJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetValidatorJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetValidatorJailed represents a ValidatorJailed event raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorJailed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorJailed is a free log retrieval operation binding the contract event 0xf226e7d8f547ff903d9d419cf5f54e0d7d07efa9584135a53a057c5f1f27f49a.
//
// Solidity: event validatorJailed(address indexed validator)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterValidatorJailed(opts *bind.FilterOpts, validator []common.Address) (*BSCValidatorSetValidatorJailedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "validatorJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetValidatorJailedIterator{contract: _BSCValidatorSet.contract, event: "validatorJailed", logs: logs, sub: sub}, nil
}

// WatchValidatorJailed is a free log subscription operation binding the contract event 0xf226e7d8f547ff903d9d419cf5f54e0d7d07efa9584135a53a057c5f1f27f49a.
//
// Solidity: event validatorJailed(address indexed validator)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchValidatorJailed(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetValidatorJailed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "validatorJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetValidatorJailed)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorJailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorJailed is a log parse operation binding the contract event 0xf226e7d8f547ff903d9d419cf5f54e0d7d07efa9584135a53a057c5f1f27f49a.
//
// Solidity: event validatorJailed(address indexed validator)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseValidatorJailed(log types.Log) (*BSCValidatorSetValidatorJailed, error) {
	event := new(BSCValidatorSetValidatorJailed)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorJailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetValidatorMisdemeanorIterator is returned from FilterValidatorMisdemeanor and is used to iterate over the raw logs and unpacked data for ValidatorMisdemeanor events raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorMisdemeanorIterator struct {
	Event *BSCValidatorSetValidatorMisdemeanor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetValidatorMisdemeanorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetValidatorMisdemeanor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetValidatorMisdemeanor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetValidatorMisdemeanorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetValidatorMisdemeanorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetValidatorMisdemeanor represents a ValidatorMisdemeanor event raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorMisdemeanor struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorMisdemeanor is a free log retrieval operation binding the contract event 0x8cd4e147d8af98a9e3b6724021b8bf6aed2e5dac71c38f2dce8161b82585b25d.
//
// Solidity: event validatorMisdemeanor(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterValidatorMisdemeanor(opts *bind.FilterOpts, validator []common.Address) (*BSCValidatorSetValidatorMisdemeanorIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "validatorMisdemeanor", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetValidatorMisdemeanorIterator{contract: _BSCValidatorSet.contract, event: "validatorMisdemeanor", logs: logs, sub: sub}, nil
}

// WatchValidatorMisdemeanor is a free log subscription operation binding the contract event 0x8cd4e147d8af98a9e3b6724021b8bf6aed2e5dac71c38f2dce8161b82585b25d.
//
// Solidity: event validatorMisdemeanor(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchValidatorMisdemeanor(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetValidatorMisdemeanor, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "validatorMisdemeanor", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetValidatorMisdemeanor)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorMisdemeanor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorMisdemeanor is a log parse operation binding the contract event 0x8cd4e147d8af98a9e3b6724021b8bf6aed2e5dac71c38f2dce8161b82585b25d.
//
// Solidity: event validatorMisdemeanor(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseValidatorMisdemeanor(log types.Log) (*BSCValidatorSetValidatorMisdemeanor, error) {
	event := new(BSCValidatorSetValidatorMisdemeanor)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorMisdemeanor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetValidatorSetUpdatedIterator is returned from FilterValidatorSetUpdated and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdated events raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorSetUpdatedIterator struct {
	Event *BSCValidatorSetValidatorSetUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetValidatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetValidatorSetUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetValidatorSetUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetValidatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetValidatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetValidatorSetUpdated represents a ValidatorSetUpdated event raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorSetUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdated is a free log retrieval operation binding the contract event 0xedd8d7296956dd970ab4de3f2fc03be2b0ffc615d20cd4c72c6e44f928630ebf.
//
// Solidity: event validatorSetUpdated()
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterValidatorSetUpdated(opts *bind.FilterOpts) (*BSCValidatorSetValidatorSetUpdatedIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "validatorSetUpdated")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetValidatorSetUpdatedIterator{contract: _BSCValidatorSet.contract, event: "validatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdated is a free log subscription operation binding the contract event 0xedd8d7296956dd970ab4de3f2fc03be2b0ffc615d20cd4c72c6e44f928630ebf.
//
// Solidity: event validatorSetUpdated()
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchValidatorSetUpdated(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetValidatorSetUpdated) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "validatorSetUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetValidatorSetUpdated)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorSetUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorSetUpdated is a log parse operation binding the contract event 0xedd8d7296956dd970ab4de3f2fc03be2b0ffc615d20cd4c72c6e44f928630ebf.
//
// Solidity: event validatorSetUpdated()
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseValidatorSetUpdated(log types.Log) (*BSCValidatorSetValidatorSetUpdated, error) {
	event := new(BSCValidatorSetValidatorSetUpdated)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BytesToTypesABI is the input ABI used to generate the binding from.
const BytesToTypesABI = "[]"

// BytesToTypesBin is the compiled bytecode used for deploying new contracts.
var BytesToTypesBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a33734ce4b05e2a4f24fb6a43c76c2a990deec01482ee930645ff1694633109f64736f6c63430006040033"

// DeployBytesToTypes deploys a new Ethereum contract, binding an instance of BytesToTypes to it.
func DeployBytesToTypes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesToTypes, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesToTypesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BytesToTypesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesToTypes{BytesToTypesCaller: BytesToTypesCaller{contract: contract}, BytesToTypesTransactor: BytesToTypesTransactor{contract: contract}, BytesToTypesFilterer: BytesToTypesFilterer{contract: contract}}, nil
}

// BytesToTypes is an auto generated Go binding around an Ethereum contract.
type BytesToTypes struct {
	BytesToTypesCaller     // Read-only binding to the contract
	BytesToTypesTransactor // Write-only binding to the contract
	BytesToTypesFilterer   // Log filterer for contract events
}

// BytesToTypesCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesToTypesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesToTypesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesToTypesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesToTypesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesToTypesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesToTypesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesToTypesSession struct {
	Contract     *BytesToTypes     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesToTypesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesToTypesCallerSession struct {
	Contract *BytesToTypesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BytesToTypesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesToTypesTransactorSession struct {
	Contract     *BytesToTypesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BytesToTypesRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesToTypesRaw struct {
	Contract *BytesToTypes // Generic contract binding to access the raw methods on
}

// BytesToTypesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesToTypesCallerRaw struct {
	Contract *BytesToTypesCaller // Generic read-only contract binding to access the raw methods on
}

// BytesToTypesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesToTypesTransactorRaw struct {
	Contract *BytesToTypesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesToTypes creates a new instance of BytesToTypes, bound to a specific deployed contract.
func NewBytesToTypes(address common.Address, backend bind.ContractBackend) (*BytesToTypes, error) {
	contract, err := bindBytesToTypes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesToTypes{BytesToTypesCaller: BytesToTypesCaller{contract: contract}, BytesToTypesTransactor: BytesToTypesTransactor{contract: contract}, BytesToTypesFilterer: BytesToTypesFilterer{contract: contract}}, nil
}

// NewBytesToTypesCaller creates a new read-only instance of BytesToTypes, bound to a specific deployed contract.
func NewBytesToTypesCaller(address common.Address, caller bind.ContractCaller) (*BytesToTypesCaller, error) {
	contract, err := bindBytesToTypes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesToTypesCaller{contract: contract}, nil
}

// NewBytesToTypesTransactor creates a new write-only instance of BytesToTypes, bound to a specific deployed contract.
func NewBytesToTypesTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesToTypesTransactor, error) {
	contract, err := bindBytesToTypes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesToTypesTransactor{contract: contract}, nil
}

// NewBytesToTypesFilterer creates a new log filterer instance of BytesToTypes, bound to a specific deployed contract.
func NewBytesToTypesFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesToTypesFilterer, error) {
	contract, err := bindBytesToTypes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesToTypesFilterer{contract: contract}, nil
}

// bindBytesToTypes binds a generic wrapper to an already deployed contract.
func bindBytesToTypes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesToTypesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesToTypes *BytesToTypesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesToTypes.Contract.BytesToTypesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesToTypes *BytesToTypesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesToTypes.Contract.BytesToTypesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesToTypes *BytesToTypesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesToTypes.Contract.BytesToTypesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesToTypes *BytesToTypesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesToTypes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesToTypes *BytesToTypesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesToTypes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesToTypes *BytesToTypesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesToTypes.Contract.contract.Transact(opts, method, params...)
}

// CmnPkgABI is the input ABI used to generate the binding from.
const CmnPkgABI = "[]"

// CmnPkgBin is the compiled bytecode used for deploying new contracts.
var CmnPkgBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122013cba6ca7d42b9596ab731c4e5624b92170ac09ff98e457a52610505ed13982a64736f6c63430006040033"

// DeployCmnPkg deploys a new Ethereum contract, binding an instance of CmnPkg to it.
func DeployCmnPkg(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CmnPkg, error) {
	parsed, err := abi.JSON(strings.NewReader(CmnPkgABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CmnPkgBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CmnPkg{CmnPkgCaller: CmnPkgCaller{contract: contract}, CmnPkgTransactor: CmnPkgTransactor{contract: contract}, CmnPkgFilterer: CmnPkgFilterer{contract: contract}}, nil
}

// CmnPkg is an auto generated Go binding around an Ethereum contract.
type CmnPkg struct {
	CmnPkgCaller     // Read-only binding to the contract
	CmnPkgTransactor // Write-only binding to the contract
	CmnPkgFilterer   // Log filterer for contract events
}

// CmnPkgCaller is an auto generated read-only Go binding around an Ethereum contract.
type CmnPkgCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CmnPkgTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CmnPkgTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CmnPkgFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CmnPkgFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CmnPkgSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CmnPkgSession struct {
	Contract     *CmnPkg           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CmnPkgCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CmnPkgCallerSession struct {
	Contract *CmnPkgCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CmnPkgTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CmnPkgTransactorSession struct {
	Contract     *CmnPkgTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CmnPkgRaw is an auto generated low-level Go binding around an Ethereum contract.
type CmnPkgRaw struct {
	Contract *CmnPkg // Generic contract binding to access the raw methods on
}

// CmnPkgCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CmnPkgCallerRaw struct {
	Contract *CmnPkgCaller // Generic read-only contract binding to access the raw methods on
}

// CmnPkgTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CmnPkgTransactorRaw struct {
	Contract *CmnPkgTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCmnPkg creates a new instance of CmnPkg, bound to a specific deployed contract.
func NewCmnPkg(address common.Address, backend bind.ContractBackend) (*CmnPkg, error) {
	contract, err := bindCmnPkg(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CmnPkg{CmnPkgCaller: CmnPkgCaller{contract: contract}, CmnPkgTransactor: CmnPkgTransactor{contract: contract}, CmnPkgFilterer: CmnPkgFilterer{contract: contract}}, nil
}

// NewCmnPkgCaller creates a new read-only instance of CmnPkg, bound to a specific deployed contract.
func NewCmnPkgCaller(address common.Address, caller bind.ContractCaller) (*CmnPkgCaller, error) {
	contract, err := bindCmnPkg(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CmnPkgCaller{contract: contract}, nil
}

// NewCmnPkgTransactor creates a new write-only instance of CmnPkg, bound to a specific deployed contract.
func NewCmnPkgTransactor(address common.Address, transactor bind.ContractTransactor) (*CmnPkgTransactor, error) {
	contract, err := bindCmnPkg(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CmnPkgTransactor{contract: contract}, nil
}

// NewCmnPkgFilterer creates a new log filterer instance of CmnPkg, bound to a specific deployed contract.
func NewCmnPkgFilterer(address common.Address, filterer bind.ContractFilterer) (*CmnPkgFilterer, error) {
	contract, err := bindCmnPkg(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CmnPkgFilterer{contract: contract}, nil
}

// bindCmnPkg binds a generic wrapper to an already deployed contract.
func bindCmnPkg(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CmnPkgABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CmnPkg *CmnPkgRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CmnPkg.Contract.CmnPkgCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CmnPkg *CmnPkgRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CmnPkg.Contract.CmnPkgTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CmnPkg *CmnPkgRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CmnPkg.Contract.CmnPkgTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CmnPkg *CmnPkgCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CmnPkg.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CmnPkg *CmnPkgTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CmnPkg.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CmnPkg *CmnPkgTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CmnPkg.Contract.contract.Transact(opts, method, params...)
}

// IApplicationABI is the input ABI used to generate the binding from.
const IApplicationABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleAckPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleFailAckPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleSynPackage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"responsePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IApplicationFuncSigs maps the 4-byte function signature to its string representation.
var IApplicationFuncSigs = map[string]string{
	"831d65d1": "handleAckPackage(uint8,bytes)",
	"c8509d81": "handleFailAckPackage(uint8,bytes)",
	"1182b875": "handleSynPackage(uint8,bytes)",
}

// IApplication is an auto generated Go binding around an Ethereum contract.
type IApplication struct {
	IApplicationCaller     // Read-only binding to the contract
	IApplicationTransactor // Write-only binding to the contract
	IApplicationFilterer   // Log filterer for contract events
}

// IApplicationCaller is an auto generated read-only Go binding around an Ethereum contract.
type IApplicationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IApplicationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IApplicationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IApplicationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IApplicationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IApplicationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IApplicationSession struct {
	Contract     *IApplication     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IApplicationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IApplicationCallerSession struct {
	Contract *IApplicationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IApplicationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IApplicationTransactorSession struct {
	Contract     *IApplicationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IApplicationRaw is an auto generated low-level Go binding around an Ethereum contract.
type IApplicationRaw struct {
	Contract *IApplication // Generic contract binding to access the raw methods on
}

// IApplicationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IApplicationCallerRaw struct {
	Contract *IApplicationCaller // Generic read-only contract binding to access the raw methods on
}

// IApplicationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IApplicationTransactorRaw struct {
	Contract *IApplicationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIApplication creates a new instance of IApplication, bound to a specific deployed contract.
func NewIApplication(address common.Address, backend bind.ContractBackend) (*IApplication, error) {
	contract, err := bindIApplication(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IApplication{IApplicationCaller: IApplicationCaller{contract: contract}, IApplicationTransactor: IApplicationTransactor{contract: contract}, IApplicationFilterer: IApplicationFilterer{contract: contract}}, nil
}

// NewIApplicationCaller creates a new read-only instance of IApplication, bound to a specific deployed contract.
func NewIApplicationCaller(address common.Address, caller bind.ContractCaller) (*IApplicationCaller, error) {
	contract, err := bindIApplication(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IApplicationCaller{contract: contract}, nil
}

// NewIApplicationTransactor creates a new write-only instance of IApplication, bound to a specific deployed contract.
func NewIApplicationTransactor(address common.Address, transactor bind.ContractTransactor) (*IApplicationTransactor, error) {
	contract, err := bindIApplication(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IApplicationTransactor{contract: contract}, nil
}

// NewIApplicationFilterer creates a new log filterer instance of IApplication, bound to a specific deployed contract.
func NewIApplicationFilterer(address common.Address, filterer bind.ContractFilterer) (*IApplicationFilterer, error) {
	contract, err := bindIApplication(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IApplicationFilterer{contract: contract}, nil
}

// bindIApplication binds a generic wrapper to an already deployed contract.
func bindIApplication(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IApplicationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IApplication *IApplicationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IApplication.Contract.IApplicationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IApplication *IApplicationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IApplication.Contract.IApplicationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IApplication *IApplicationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IApplication.Contract.IApplicationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IApplication *IApplicationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IApplication.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IApplication *IApplicationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IApplication.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IApplication *IApplicationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IApplication.Contract.contract.Transact(opts, method, params...)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_IApplication *IApplicationTransactor) HandleAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.contract.Transact(opts, "handleAckPackage", channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_IApplication *IApplicationSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.Contract.HandleAckPackage(&_IApplication.TransactOpts, channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_IApplication *IApplicationTransactorSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.Contract.HandleAckPackage(&_IApplication.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_IApplication *IApplicationTransactor) HandleFailAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.contract.Transact(opts, "handleFailAckPackage", channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_IApplication *IApplicationSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.Contract.HandleFailAckPackage(&_IApplication.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_IApplication *IApplicationTransactorSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.Contract.HandleFailAckPackage(&_IApplication.TransactOpts, channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 channelId, bytes msgBytes) returns(bytes responsePayload)
func (_IApplication *IApplicationTransactor) HandleSynPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.contract.Transact(opts, "handleSynPackage", channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 channelId, bytes msgBytes) returns(bytes responsePayload)
func (_IApplication *IApplicationSession) HandleSynPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.Contract.HandleSynPackage(&_IApplication.TransactOpts, channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 channelId, bytes msgBytes) returns(bytes responsePayload)
func (_IApplication *IApplicationTransactorSession) HandleSynPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.Contract.HandleSynPackage(&_IApplication.TransactOpts, channelId, msgBytes)
}

// IBSCValidatorSetABI is the input ABI used to generate the binding from.
const IBSCValidatorSetABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"felony\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"misdemeanor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IBSCValidatorSetFuncSigs maps the 4-byte function signature to its string representation.
var IBSCValidatorSetFuncSigs = map[string]string{
	"35409f7f": "felony(address)",
	"eb57e202": "misdemeanor(address)",
}

// IBSCValidatorSet is an auto generated Go binding around an Ethereum contract.
type IBSCValidatorSet struct {
	IBSCValidatorSetCaller     // Read-only binding to the contract
	IBSCValidatorSetTransactor // Write-only binding to the contract
	IBSCValidatorSetFilterer   // Log filterer for contract events
}

// IBSCValidatorSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBSCValidatorSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBSCValidatorSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBSCValidatorSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBSCValidatorSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBSCValidatorSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBSCValidatorSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBSCValidatorSetSession struct {
	Contract     *IBSCValidatorSet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBSCValidatorSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBSCValidatorSetCallerSession struct {
	Contract *IBSCValidatorSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IBSCValidatorSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBSCValidatorSetTransactorSession struct {
	Contract     *IBSCValidatorSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IBSCValidatorSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBSCValidatorSetRaw struct {
	Contract *IBSCValidatorSet // Generic contract binding to access the raw methods on
}

// IBSCValidatorSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBSCValidatorSetCallerRaw struct {
	Contract *IBSCValidatorSetCaller // Generic read-only contract binding to access the raw methods on
}

// IBSCValidatorSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBSCValidatorSetTransactorRaw struct {
	Contract *IBSCValidatorSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBSCValidatorSet creates a new instance of IBSCValidatorSet, bound to a specific deployed contract.
func NewIBSCValidatorSet(address common.Address, backend bind.ContractBackend) (*IBSCValidatorSet, error) {
	contract, err := bindIBSCValidatorSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBSCValidatorSet{IBSCValidatorSetCaller: IBSCValidatorSetCaller{contract: contract}, IBSCValidatorSetTransactor: IBSCValidatorSetTransactor{contract: contract}, IBSCValidatorSetFilterer: IBSCValidatorSetFilterer{contract: contract}}, nil
}

// NewIBSCValidatorSetCaller creates a new read-only instance of IBSCValidatorSet, bound to a specific deployed contract.
func NewIBSCValidatorSetCaller(address common.Address, caller bind.ContractCaller) (*IBSCValidatorSetCaller, error) {
	contract, err := bindIBSCValidatorSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBSCValidatorSetCaller{contract: contract}, nil
}

// NewIBSCValidatorSetTransactor creates a new write-only instance of IBSCValidatorSet, bound to a specific deployed contract.
func NewIBSCValidatorSetTransactor(address common.Address, transactor bind.ContractTransactor) (*IBSCValidatorSetTransactor, error) {
	contract, err := bindIBSCValidatorSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBSCValidatorSetTransactor{contract: contract}, nil
}

// NewIBSCValidatorSetFilterer creates a new log filterer instance of IBSCValidatorSet, bound to a specific deployed contract.
func NewIBSCValidatorSetFilterer(address common.Address, filterer bind.ContractFilterer) (*IBSCValidatorSetFilterer, error) {
	contract, err := bindIBSCValidatorSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBSCValidatorSetFilterer{contract: contract}, nil
}

// bindIBSCValidatorSet binds a generic wrapper to an already deployed contract.
func bindIBSCValidatorSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBSCValidatorSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBSCValidatorSet *IBSCValidatorSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBSCValidatorSet.Contract.IBSCValidatorSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBSCValidatorSet *IBSCValidatorSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.IBSCValidatorSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBSCValidatorSet *IBSCValidatorSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.IBSCValidatorSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBSCValidatorSet *IBSCValidatorSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBSCValidatorSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBSCValidatorSet *IBSCValidatorSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBSCValidatorSet *IBSCValidatorSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.contract.Transact(opts, method, params...)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_IBSCValidatorSet *IBSCValidatorSetTransactor) Felony(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _IBSCValidatorSet.contract.Transact(opts, "felony", validator)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_IBSCValidatorSet *IBSCValidatorSetSession) Felony(validator common.Address) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.Felony(&_IBSCValidatorSet.TransactOpts, validator)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_IBSCValidatorSet *IBSCValidatorSetTransactorSession) Felony(validator common.Address) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.Felony(&_IBSCValidatorSet.TransactOpts, validator)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_IBSCValidatorSet *IBSCValidatorSetTransactor) Misdemeanor(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _IBSCValidatorSet.contract.Transact(opts, "misdemeanor", validator)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_IBSCValidatorSet *IBSCValidatorSetSession) Misdemeanor(validator common.Address) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.Misdemeanor(&_IBSCValidatorSet.TransactOpts, validator)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_IBSCValidatorSet *IBSCValidatorSetTransactorSession) Misdemeanor(validator common.Address) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.Misdemeanor(&_IBSCValidatorSet.TransactOpts, validator)
}

// ILightClientABI is the input ABI used to generate the binding from.
const ILightClientABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getAppHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getSubmitter\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"isHeaderSynced\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ILightClientFuncSigs maps the 4-byte function signature to its string representation.
var ILightClientFuncSigs = map[string]string{
	"cba510a9": "getAppHash(uint64)",
	"dda83148": "getSubmitter(uint64)",
	"df5fe704": "isHeaderSynced(uint64)",
}

// ILightClient is an auto generated Go binding around an Ethereum contract.
type ILightClient struct {
	ILightClientCaller     // Read-only binding to the contract
	ILightClientTransactor // Write-only binding to the contract
	ILightClientFilterer   // Log filterer for contract events
}

// ILightClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILightClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILightClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILightClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILightClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILightClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILightClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILightClientSession struct {
	Contract     *ILightClient     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILightClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILightClientCallerSession struct {
	Contract *ILightClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ILightClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILightClientTransactorSession struct {
	Contract     *ILightClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ILightClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILightClientRaw struct {
	Contract *ILightClient // Generic contract binding to access the raw methods on
}

// ILightClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILightClientCallerRaw struct {
	Contract *ILightClientCaller // Generic read-only contract binding to access the raw methods on
}

// ILightClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILightClientTransactorRaw struct {
	Contract *ILightClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILightClient creates a new instance of ILightClient, bound to a specific deployed contract.
func NewILightClient(address common.Address, backend bind.ContractBackend) (*ILightClient, error) {
	contract, err := bindILightClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILightClient{ILightClientCaller: ILightClientCaller{contract: contract}, ILightClientTransactor: ILightClientTransactor{contract: contract}, ILightClientFilterer: ILightClientFilterer{contract: contract}}, nil
}

// NewILightClientCaller creates a new read-only instance of ILightClient, bound to a specific deployed contract.
func NewILightClientCaller(address common.Address, caller bind.ContractCaller) (*ILightClientCaller, error) {
	contract, err := bindILightClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILightClientCaller{contract: contract}, nil
}

// NewILightClientTransactor creates a new write-only instance of ILightClient, bound to a specific deployed contract.
func NewILightClientTransactor(address common.Address, transactor bind.ContractTransactor) (*ILightClientTransactor, error) {
	contract, err := bindILightClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILightClientTransactor{contract: contract}, nil
}

// NewILightClientFilterer creates a new log filterer instance of ILightClient, bound to a specific deployed contract.
func NewILightClientFilterer(address common.Address, filterer bind.ContractFilterer) (*ILightClientFilterer, error) {
	contract, err := bindILightClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILightClientFilterer{contract: contract}, nil
}

// bindILightClient binds a generic wrapper to an already deployed contract.
func bindILightClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ILightClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILightClient *ILightClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILightClient.Contract.ILightClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILightClient *ILightClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILightClient.Contract.ILightClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILightClient *ILightClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILightClient.Contract.ILightClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILightClient *ILightClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILightClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILightClient *ILightClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILightClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILightClient *ILightClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILightClient.Contract.contract.Transact(opts, method, params...)
}

// GetAppHash is a free data retrieval call binding the contract method 0xcba510a9.
//
// Solidity: function getAppHash(uint64 height) view returns(bytes32)
func (_ILightClient *ILightClientCaller) GetAppHash(opts *bind.CallOpts, height uint64) ([32]byte, error) {
	var out []interface{}
	err := _ILightClient.contract.Call(opts, &out, "getAppHash", height)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAppHash is a free data retrieval call binding the contract method 0xcba510a9.
//
// Solidity: function getAppHash(uint64 height) view returns(bytes32)
func (_ILightClient *ILightClientSession) GetAppHash(height uint64) ([32]byte, error) {
	return _ILightClient.Contract.GetAppHash(&_ILightClient.CallOpts, height)
}

// GetAppHash is a free data retrieval call binding the contract method 0xcba510a9.
//
// Solidity: function getAppHash(uint64 height) view returns(bytes32)
func (_ILightClient *ILightClientCallerSession) GetAppHash(height uint64) ([32]byte, error) {
	return _ILightClient.Contract.GetAppHash(&_ILightClient.CallOpts, height)
}

// GetSubmitter is a free data retrieval call binding the contract method 0xdda83148.
//
// Solidity: function getSubmitter(uint64 height) view returns(address)
func (_ILightClient *ILightClientCaller) GetSubmitter(opts *bind.CallOpts, height uint64) (common.Address, error) {
	var out []interface{}
	err := _ILightClient.contract.Call(opts, &out, "getSubmitter", height)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSubmitter is a free data retrieval call binding the contract method 0xdda83148.
//
// Solidity: function getSubmitter(uint64 height) view returns(address)
func (_ILightClient *ILightClientSession) GetSubmitter(height uint64) (common.Address, error) {
	return _ILightClient.Contract.GetSubmitter(&_ILightClient.CallOpts, height)
}

// GetSubmitter is a free data retrieval call binding the contract method 0xdda83148.
//
// Solidity: function getSubmitter(uint64 height) view returns(address)
func (_ILightClient *ILightClientCallerSession) GetSubmitter(height uint64) (common.Address, error) {
	return _ILightClient.Contract.GetSubmitter(&_ILightClient.CallOpts, height)
}

// IsHeaderSynced is a free data retrieval call binding the contract method 0xdf5fe704.
//
// Solidity: function isHeaderSynced(uint64 height) view returns(bool)
func (_ILightClient *ILightClientCaller) IsHeaderSynced(opts *bind.CallOpts, height uint64) (bool, error) {
	var out []interface{}
	err := _ILightClient.contract.Call(opts, &out, "isHeaderSynced", height)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHeaderSynced is a free data retrieval call binding the contract method 0xdf5fe704.
//
// Solidity: function isHeaderSynced(uint64 height) view returns(bool)
func (_ILightClient *ILightClientSession) IsHeaderSynced(height uint64) (bool, error) {
	return _ILightClient.Contract.IsHeaderSynced(&_ILightClient.CallOpts, height)
}

// IsHeaderSynced is a free data retrieval call binding the contract method 0xdf5fe704.
//
// Solidity: function isHeaderSynced(uint64 height) view returns(bool)
func (_ILightClient *ILightClientCallerSession) IsHeaderSynced(height uint64) (bool, error) {
	return _ILightClient.Contract.IsHeaderSynced(&_ILightClient.CallOpts, height)
}

// IParamSubscriberABI is the input ABI used to generate the binding from.
const IParamSubscriberABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"updateParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IParamSubscriberFuncSigs maps the 4-byte function signature to its string representation.
var IParamSubscriberFuncSigs = map[string]string{
	"ac431751": "updateParam(string,bytes)",
}

// IParamSubscriber is an auto generated Go binding around an Ethereum contract.
type IParamSubscriber struct {
	IParamSubscriberCaller     // Read-only binding to the contract
	IParamSubscriberTransactor // Write-only binding to the contract
	IParamSubscriberFilterer   // Log filterer for contract events
}

// IParamSubscriberCaller is an auto generated read-only Go binding around an Ethereum contract.
type IParamSubscriberCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IParamSubscriberTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IParamSubscriberTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IParamSubscriberFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IParamSubscriberFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IParamSubscriberSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IParamSubscriberSession struct {
	Contract     *IParamSubscriber // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IParamSubscriberCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IParamSubscriberCallerSession struct {
	Contract *IParamSubscriberCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IParamSubscriberTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IParamSubscriberTransactorSession struct {
	Contract     *IParamSubscriberTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IParamSubscriberRaw is an auto generated low-level Go binding around an Ethereum contract.
type IParamSubscriberRaw struct {
	Contract *IParamSubscriber // Generic contract binding to access the raw methods on
}

// IParamSubscriberCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IParamSubscriberCallerRaw struct {
	Contract *IParamSubscriberCaller // Generic read-only contract binding to access the raw methods on
}

// IParamSubscriberTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IParamSubscriberTransactorRaw struct {
	Contract *IParamSubscriberTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIParamSubscriber creates a new instance of IParamSubscriber, bound to a specific deployed contract.
func NewIParamSubscriber(address common.Address, backend bind.ContractBackend) (*IParamSubscriber, error) {
	contract, err := bindIParamSubscriber(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IParamSubscriber{IParamSubscriberCaller: IParamSubscriberCaller{contract: contract}, IParamSubscriberTransactor: IParamSubscriberTransactor{contract: contract}, IParamSubscriberFilterer: IParamSubscriberFilterer{contract: contract}}, nil
}

// NewIParamSubscriberCaller creates a new read-only instance of IParamSubscriber, bound to a specific deployed contract.
func NewIParamSubscriberCaller(address common.Address, caller bind.ContractCaller) (*IParamSubscriberCaller, error) {
	contract, err := bindIParamSubscriber(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IParamSubscriberCaller{contract: contract}, nil
}

// NewIParamSubscriberTransactor creates a new write-only instance of IParamSubscriber, bound to a specific deployed contract.
func NewIParamSubscriberTransactor(address common.Address, transactor bind.ContractTransactor) (*IParamSubscriberTransactor, error) {
	contract, err := bindIParamSubscriber(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IParamSubscriberTransactor{contract: contract}, nil
}

// NewIParamSubscriberFilterer creates a new log filterer instance of IParamSubscriber, bound to a specific deployed contract.
func NewIParamSubscriberFilterer(address common.Address, filterer bind.ContractFilterer) (*IParamSubscriberFilterer, error) {
	contract, err := bindIParamSubscriber(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IParamSubscriberFilterer{contract: contract}, nil
}

// bindIParamSubscriber binds a generic wrapper to an already deployed contract.
func bindIParamSubscriber(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IParamSubscriberABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IParamSubscriber *IParamSubscriberRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IParamSubscriber.Contract.IParamSubscriberCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IParamSubscriber *IParamSubscriberRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IParamSubscriber.Contract.IParamSubscriberTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IParamSubscriber *IParamSubscriberRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IParamSubscriber.Contract.IParamSubscriberTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IParamSubscriber *IParamSubscriberCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IParamSubscriber.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IParamSubscriber *IParamSubscriberTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IParamSubscriber.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IParamSubscriber *IParamSubscriberTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IParamSubscriber.Contract.contract.Transact(opts, method, params...)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_IParamSubscriber *IParamSubscriberTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _IParamSubscriber.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_IParamSubscriber *IParamSubscriberSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _IParamSubscriber.Contract.UpdateParam(&_IParamSubscriber.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_IParamSubscriber *IParamSubscriberTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _IParamSubscriber.Contract.UpdateParam(&_IParamSubscriber.TransactOpts, key, value)
}

// IRelayerHubABI is the input ABI used to generate the binding from.
const IRelayerHubABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IRelayerHubFuncSigs maps the 4-byte function signature to its string representation.
var IRelayerHubFuncSigs = map[string]string{
	"541d5548": "isRelayer(address)",
}

// IRelayerHub is an auto generated Go binding around an Ethereum contract.
type IRelayerHub struct {
	IRelayerHubCaller     // Read-only binding to the contract
	IRelayerHubTransactor // Write-only binding to the contract
	IRelayerHubFilterer   // Log filterer for contract events
}

// IRelayerHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRelayerHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayerHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRelayerHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayerHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRelayerHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayerHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRelayerHubSession struct {
	Contract     *IRelayerHub      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRelayerHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRelayerHubCallerSession struct {
	Contract *IRelayerHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IRelayerHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRelayerHubTransactorSession struct {
	Contract     *IRelayerHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IRelayerHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRelayerHubRaw struct {
	Contract *IRelayerHub // Generic contract binding to access the raw methods on
}

// IRelayerHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRelayerHubCallerRaw struct {
	Contract *IRelayerHubCaller // Generic read-only contract binding to access the raw methods on
}

// IRelayerHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRelayerHubTransactorRaw struct {
	Contract *IRelayerHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRelayerHub creates a new instance of IRelayerHub, bound to a specific deployed contract.
func NewIRelayerHub(address common.Address, backend bind.ContractBackend) (*IRelayerHub, error) {
	contract, err := bindIRelayerHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRelayerHub{IRelayerHubCaller: IRelayerHubCaller{contract: contract}, IRelayerHubTransactor: IRelayerHubTransactor{contract: contract}, IRelayerHubFilterer: IRelayerHubFilterer{contract: contract}}, nil
}

// NewIRelayerHubCaller creates a new read-only instance of IRelayerHub, bound to a specific deployed contract.
func NewIRelayerHubCaller(address common.Address, caller bind.ContractCaller) (*IRelayerHubCaller, error) {
	contract, err := bindIRelayerHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRelayerHubCaller{contract: contract}, nil
}

// NewIRelayerHubTransactor creates a new write-only instance of IRelayerHub, bound to a specific deployed contract.
func NewIRelayerHubTransactor(address common.Address, transactor bind.ContractTransactor) (*IRelayerHubTransactor, error) {
	contract, err := bindIRelayerHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRelayerHubTransactor{contract: contract}, nil
}

// NewIRelayerHubFilterer creates a new log filterer instance of IRelayerHub, bound to a specific deployed contract.
func NewIRelayerHubFilterer(address common.Address, filterer bind.ContractFilterer) (*IRelayerHubFilterer, error) {
	contract, err := bindIRelayerHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRelayerHubFilterer{contract: contract}, nil
}

// bindIRelayerHub binds a generic wrapper to an already deployed contract.
func bindIRelayerHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRelayerHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRelayerHub *IRelayerHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRelayerHub.Contract.IRelayerHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRelayerHub *IRelayerHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRelayerHub.Contract.IRelayerHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRelayerHub *IRelayerHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRelayerHub.Contract.IRelayerHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRelayerHub *IRelayerHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRelayerHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRelayerHub *IRelayerHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRelayerHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRelayerHub *IRelayerHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRelayerHub.Contract.contract.Transact(opts, method, params...)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) view returns(bool)
func (_IRelayerHub *IRelayerHubCaller) IsRelayer(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _IRelayerHub.contract.Call(opts, &out, "isRelayer", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) view returns(bool)
func (_IRelayerHub *IRelayerHubSession) IsRelayer(sender common.Address) (bool, error) {
	return _IRelayerHub.Contract.IsRelayer(&_IRelayerHub.CallOpts, sender)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) view returns(bool)
func (_IRelayerHub *IRelayerHubCallerSession) IsRelayer(sender common.Address) (bool, error) {
	return _IRelayerHub.Contract.IsRelayer(&_IRelayerHub.CallOpts, sender)
}

// ISlashIndicatorABI is the input ABI used to generate the binding from.
const ISlashIndicatorABI = "[{\"inputs\":[],\"name\":\"clean\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ISlashIndicatorFuncSigs maps the 4-byte function signature to its string representation.
var ISlashIndicatorFuncSigs = map[string]string{
	"fc4333cd": "clean()",
}

// ISlashIndicator is an auto generated Go binding around an Ethereum contract.
type ISlashIndicator struct {
	ISlashIndicatorCaller     // Read-only binding to the contract
	ISlashIndicatorTransactor // Write-only binding to the contract
	ISlashIndicatorFilterer   // Log filterer for contract events
}

// ISlashIndicatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISlashIndicatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlashIndicatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISlashIndicatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlashIndicatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISlashIndicatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlashIndicatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISlashIndicatorSession struct {
	Contract     *ISlashIndicator  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISlashIndicatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISlashIndicatorCallerSession struct {
	Contract *ISlashIndicatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ISlashIndicatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISlashIndicatorTransactorSession struct {
	Contract     *ISlashIndicatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ISlashIndicatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISlashIndicatorRaw struct {
	Contract *ISlashIndicator // Generic contract binding to access the raw methods on
}

// ISlashIndicatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISlashIndicatorCallerRaw struct {
	Contract *ISlashIndicatorCaller // Generic read-only contract binding to access the raw methods on
}

// ISlashIndicatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISlashIndicatorTransactorRaw struct {
	Contract *ISlashIndicatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISlashIndicator creates a new instance of ISlashIndicator, bound to a specific deployed contract.
func NewISlashIndicator(address common.Address, backend bind.ContractBackend) (*ISlashIndicator, error) {
	contract, err := bindISlashIndicator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISlashIndicator{ISlashIndicatorCaller: ISlashIndicatorCaller{contract: contract}, ISlashIndicatorTransactor: ISlashIndicatorTransactor{contract: contract}, ISlashIndicatorFilterer: ISlashIndicatorFilterer{contract: contract}}, nil
}

// NewISlashIndicatorCaller creates a new read-only instance of ISlashIndicator, bound to a specific deployed contract.
func NewISlashIndicatorCaller(address common.Address, caller bind.ContractCaller) (*ISlashIndicatorCaller, error) {
	contract, err := bindISlashIndicator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISlashIndicatorCaller{contract: contract}, nil
}

// NewISlashIndicatorTransactor creates a new write-only instance of ISlashIndicator, bound to a specific deployed contract.
func NewISlashIndicatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ISlashIndicatorTransactor, error) {
	contract, err := bindISlashIndicator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISlashIndicatorTransactor{contract: contract}, nil
}

// NewISlashIndicatorFilterer creates a new log filterer instance of ISlashIndicator, bound to a specific deployed contract.
func NewISlashIndicatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ISlashIndicatorFilterer, error) {
	contract, err := bindISlashIndicator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISlashIndicatorFilterer{contract: contract}, nil
}

// bindISlashIndicator binds a generic wrapper to an already deployed contract.
func bindISlashIndicator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISlashIndicatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISlashIndicator *ISlashIndicatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISlashIndicator.Contract.ISlashIndicatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISlashIndicator *ISlashIndicatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlashIndicator.Contract.ISlashIndicatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISlashIndicator *ISlashIndicatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISlashIndicator.Contract.ISlashIndicatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISlashIndicator *ISlashIndicatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISlashIndicator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISlashIndicator *ISlashIndicatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlashIndicator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISlashIndicator *ISlashIndicatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISlashIndicator.Contract.contract.Transact(opts, method, params...)
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_ISlashIndicator *ISlashIndicatorTransactor) Clean(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlashIndicator.contract.Transact(opts, "clean")
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_ISlashIndicator *ISlashIndicatorSession) Clean() (*types.Transaction, error) {
	return _ISlashIndicator.Contract.Clean(&_ISlashIndicator.TransactOpts)
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_ISlashIndicator *ISlashIndicatorTransactorSession) Clean() (*types.Transaction, error) {
	return _ISlashIndicator.Contract.Clean(&_ISlashIndicator.TransactOpts)
}

// ISystemRewardABI is the input ABI used to generate the binding from.
const ISystemRewardABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ISystemRewardFuncSigs maps the 4-byte function signature to its string representation.
var ISystemRewardFuncSigs = map[string]string{
	"9a99b4f0": "claimRewards(address,uint256)",
}

// ISystemReward is an auto generated Go binding around an Ethereum contract.
type ISystemReward struct {
	ISystemRewardCaller     // Read-only binding to the contract
	ISystemRewardTransactor // Write-only binding to the contract
	ISystemRewardFilterer   // Log filterer for contract events
}

// ISystemRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISystemRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISystemRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISystemRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISystemRewardSession struct {
	Contract     *ISystemReward    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISystemRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISystemRewardCallerSession struct {
	Contract *ISystemRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ISystemRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISystemRewardTransactorSession struct {
	Contract     *ISystemRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ISystemRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISystemRewardRaw struct {
	Contract *ISystemReward // Generic contract binding to access the raw methods on
}

// ISystemRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISystemRewardCallerRaw struct {
	Contract *ISystemRewardCaller // Generic read-only contract binding to access the raw methods on
}

// ISystemRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISystemRewardTransactorRaw struct {
	Contract *ISystemRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISystemReward creates a new instance of ISystemReward, bound to a specific deployed contract.
func NewISystemReward(address common.Address, backend bind.ContractBackend) (*ISystemReward, error) {
	contract, err := bindISystemReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISystemReward{ISystemRewardCaller: ISystemRewardCaller{contract: contract}, ISystemRewardTransactor: ISystemRewardTransactor{contract: contract}, ISystemRewardFilterer: ISystemRewardFilterer{contract: contract}}, nil
}

// NewISystemRewardCaller creates a new read-only instance of ISystemReward, bound to a specific deployed contract.
func NewISystemRewardCaller(address common.Address, caller bind.ContractCaller) (*ISystemRewardCaller, error) {
	contract, err := bindISystemReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISystemRewardCaller{contract: contract}, nil
}

// NewISystemRewardTransactor creates a new write-only instance of ISystemReward, bound to a specific deployed contract.
func NewISystemRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*ISystemRewardTransactor, error) {
	contract, err := bindISystemReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISystemRewardTransactor{contract: contract}, nil
}

// NewISystemRewardFilterer creates a new log filterer instance of ISystemReward, bound to a specific deployed contract.
func NewISystemRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*ISystemRewardFilterer, error) {
	contract, err := bindISystemReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISystemRewardFilterer{contract: contract}, nil
}

// bindISystemReward binds a generic wrapper to an already deployed contract.
func bindISystemReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISystemRewardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISystemReward *ISystemRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISystemReward.Contract.ISystemRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISystemReward *ISystemRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISystemReward.Contract.ISystemRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISystemReward *ISystemRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISystemReward.Contract.ISystemRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISystemReward *ISystemRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISystemReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISystemReward *ISystemRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISystemReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISystemReward *ISystemRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISystemReward.Contract.contract.Transact(opts, method, params...)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address to, uint256 amount) returns(uint256 actualAmount)
func (_ISystemReward *ISystemRewardTransactor) ClaimRewards(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISystemReward.contract.Transact(opts, "claimRewards", to, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address to, uint256 amount) returns(uint256 actualAmount)
func (_ISystemReward *ISystemRewardSession) ClaimRewards(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISystemReward.Contract.ClaimRewards(&_ISystemReward.TransactOpts, to, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address to, uint256 amount) returns(uint256 actualAmount)
func (_ISystemReward *ISystemRewardTransactorSession) ClaimRewards(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISystemReward.Contract.ClaimRewards(&_ISystemReward.TransactOpts, to, amount)
}

// ITokenHubABI is the input ABI used to generate the binding from.
const ITokenHubABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipientAddrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"refundAddrs\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"expireTime\",\"type\":\"uint64\"}],\"name\":\"batchTransferOutBNB\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bep2Symbol\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"bindToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"getBep2SymbolByContractAddr\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bep2Symbol\",\"type\":\"bytes32\"}],\"name\":\"getContractAddrByBEP2Symbol\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMiniRelayFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expireTime\",\"type\":\"uint64\"}],\"name\":\"transferOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bep2Symbol\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"unbindToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ITokenHubFuncSigs maps the 4-byte function signature to its string representation.
var ITokenHubFuncSigs = map[string]string{
	"6e056520": "batchTransferOutBNB(address[],uint256[],address[],uint64)",
	"8eff336c": "bindToken(bytes32,address,uint256)",
	"bd466461": "getBep2SymbolByContractAddr(address)",
	"59b92789": "getContractAddrByBEP2Symbol(bytes32)",
	"149d14d9": "getMiniRelayFee()",
	"aa7415f5": "transferOut(address,address,uint256,uint64)",
	"b99328c5": "unbindToken(bytes32,address)",
}

// ITokenHub is an auto generated Go binding around an Ethereum contract.
type ITokenHub struct {
	ITokenHubCaller     // Read-only binding to the contract
	ITokenHubTransactor // Write-only binding to the contract
	ITokenHubFilterer   // Log filterer for contract events
}

// ITokenHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenHubSession struct {
	Contract     *ITokenHub        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenHubCallerSession struct {
	Contract *ITokenHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ITokenHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenHubTransactorSession struct {
	Contract     *ITokenHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ITokenHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenHubRaw struct {
	Contract *ITokenHub // Generic contract binding to access the raw methods on
}

// ITokenHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenHubCallerRaw struct {
	Contract *ITokenHubCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenHubTransactorRaw struct {
	Contract *ITokenHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITokenHub creates a new instance of ITokenHub, bound to a specific deployed contract.
func NewITokenHub(address common.Address, backend bind.ContractBackend) (*ITokenHub, error) {
	contract, err := bindITokenHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITokenHub{ITokenHubCaller: ITokenHubCaller{contract: contract}, ITokenHubTransactor: ITokenHubTransactor{contract: contract}, ITokenHubFilterer: ITokenHubFilterer{contract: contract}}, nil
}

// NewITokenHubCaller creates a new read-only instance of ITokenHub, bound to a specific deployed contract.
func NewITokenHubCaller(address common.Address, caller bind.ContractCaller) (*ITokenHubCaller, error) {
	contract, err := bindITokenHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenHubCaller{contract: contract}, nil
}

// NewITokenHubTransactor creates a new write-only instance of ITokenHub, bound to a specific deployed contract.
func NewITokenHubTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenHubTransactor, error) {
	contract, err := bindITokenHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenHubTransactor{contract: contract}, nil
}

// NewITokenHubFilterer creates a new log filterer instance of ITokenHub, bound to a specific deployed contract.
func NewITokenHubFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenHubFilterer, error) {
	contract, err := bindITokenHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenHubFilterer{contract: contract}, nil
}

// bindITokenHub binds a generic wrapper to an already deployed contract.
func bindITokenHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenHub *ITokenHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenHub.Contract.ITokenHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenHub *ITokenHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenHub.Contract.ITokenHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenHub *ITokenHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenHub.Contract.ITokenHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenHub *ITokenHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenHub *ITokenHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenHub *ITokenHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenHub.Contract.contract.Transact(opts, method, params...)
}

// GetBep2SymbolByContractAddr is a free data retrieval call binding the contract method 0xbd466461.
//
// Solidity: function getBep2SymbolByContractAddr(address contractAddr) view returns(bytes32)
func (_ITokenHub *ITokenHubCaller) GetBep2SymbolByContractAddr(opts *bind.CallOpts, contractAddr common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ITokenHub.contract.Call(opts, &out, "getBep2SymbolByContractAddr", contractAddr)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBep2SymbolByContractAddr is a free data retrieval call binding the contract method 0xbd466461.
//
// Solidity: function getBep2SymbolByContractAddr(address contractAddr) view returns(bytes32)
func (_ITokenHub *ITokenHubSession) GetBep2SymbolByContractAddr(contractAddr common.Address) ([32]byte, error) {
	return _ITokenHub.Contract.GetBep2SymbolByContractAddr(&_ITokenHub.CallOpts, contractAddr)
}

// GetBep2SymbolByContractAddr is a free data retrieval call binding the contract method 0xbd466461.
//
// Solidity: function getBep2SymbolByContractAddr(address contractAddr) view returns(bytes32)
func (_ITokenHub *ITokenHubCallerSession) GetBep2SymbolByContractAddr(contractAddr common.Address) ([32]byte, error) {
	return _ITokenHub.Contract.GetBep2SymbolByContractAddr(&_ITokenHub.CallOpts, contractAddr)
}

// GetContractAddrByBEP2Symbol is a free data retrieval call binding the contract method 0x59b92789.
//
// Solidity: function getContractAddrByBEP2Symbol(bytes32 bep2Symbol) view returns(address)
func (_ITokenHub *ITokenHubCaller) GetContractAddrByBEP2Symbol(opts *bind.CallOpts, bep2Symbol [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ITokenHub.contract.Call(opts, &out, "getContractAddrByBEP2Symbol", bep2Symbol)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContractAddrByBEP2Symbol is a free data retrieval call binding the contract method 0x59b92789.
//
// Solidity: function getContractAddrByBEP2Symbol(bytes32 bep2Symbol) view returns(address)
func (_ITokenHub *ITokenHubSession) GetContractAddrByBEP2Symbol(bep2Symbol [32]byte) (common.Address, error) {
	return _ITokenHub.Contract.GetContractAddrByBEP2Symbol(&_ITokenHub.CallOpts, bep2Symbol)
}

// GetContractAddrByBEP2Symbol is a free data retrieval call binding the contract method 0x59b92789.
//
// Solidity: function getContractAddrByBEP2Symbol(bytes32 bep2Symbol) view returns(address)
func (_ITokenHub *ITokenHubCallerSession) GetContractAddrByBEP2Symbol(bep2Symbol [32]byte) (common.Address, error) {
	return _ITokenHub.Contract.GetContractAddrByBEP2Symbol(&_ITokenHub.CallOpts, bep2Symbol)
}

// GetMiniRelayFee is a free data retrieval call binding the contract method 0x149d14d9.
//
// Solidity: function getMiniRelayFee() view returns(uint256)
func (_ITokenHub *ITokenHubCaller) GetMiniRelayFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ITokenHub.contract.Call(opts, &out, "getMiniRelayFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMiniRelayFee is a free data retrieval call binding the contract method 0x149d14d9.
//
// Solidity: function getMiniRelayFee() view returns(uint256)
func (_ITokenHub *ITokenHubSession) GetMiniRelayFee() (*big.Int, error) {
	return _ITokenHub.Contract.GetMiniRelayFee(&_ITokenHub.CallOpts)
}

// GetMiniRelayFee is a free data retrieval call binding the contract method 0x149d14d9.
//
// Solidity: function getMiniRelayFee() view returns(uint256)
func (_ITokenHub *ITokenHubCallerSession) GetMiniRelayFee() (*big.Int, error) {
	return _ITokenHub.Contract.GetMiniRelayFee(&_ITokenHub.CallOpts)
}

// BatchTransferOutBNB is a paid mutator transaction binding the contract method 0x6e056520.
//
// Solidity: function batchTransferOutBNB(address[] recipientAddrs, uint256[] amounts, address[] refundAddrs, uint64 expireTime) payable returns(bool)
func (_ITokenHub *ITokenHubTransactor) BatchTransferOutBNB(opts *bind.TransactOpts, recipientAddrs []common.Address, amounts []*big.Int, refundAddrs []common.Address, expireTime uint64) (*types.Transaction, error) {
	return _ITokenHub.contract.Transact(opts, "batchTransferOutBNB", recipientAddrs, amounts, refundAddrs, expireTime)
}

// BatchTransferOutBNB is a paid mutator transaction binding the contract method 0x6e056520.
//
// Solidity: function batchTransferOutBNB(address[] recipientAddrs, uint256[] amounts, address[] refundAddrs, uint64 expireTime) payable returns(bool)
func (_ITokenHub *ITokenHubSession) BatchTransferOutBNB(recipientAddrs []common.Address, amounts []*big.Int, refundAddrs []common.Address, expireTime uint64) (*types.Transaction, error) {
	return _ITokenHub.Contract.BatchTransferOutBNB(&_ITokenHub.TransactOpts, recipientAddrs, amounts, refundAddrs, expireTime)
}

// BatchTransferOutBNB is a paid mutator transaction binding the contract method 0x6e056520.
//
// Solidity: function batchTransferOutBNB(address[] recipientAddrs, uint256[] amounts, address[] refundAddrs, uint64 expireTime) payable returns(bool)
func (_ITokenHub *ITokenHubTransactorSession) BatchTransferOutBNB(recipientAddrs []common.Address, amounts []*big.Int, refundAddrs []common.Address, expireTime uint64) (*types.Transaction, error) {
	return _ITokenHub.Contract.BatchTransferOutBNB(&_ITokenHub.TransactOpts, recipientAddrs, amounts, refundAddrs, expireTime)
}

// BindToken is a paid mutator transaction binding the contract method 0x8eff336c.
//
// Solidity: function bindToken(bytes32 bep2Symbol, address contractAddr, uint256 decimals) returns()
func (_ITokenHub *ITokenHubTransactor) BindToken(opts *bind.TransactOpts, bep2Symbol [32]byte, contractAddr common.Address, decimals *big.Int) (*types.Transaction, error) {
	return _ITokenHub.contract.Transact(opts, "bindToken", bep2Symbol, contractAddr, decimals)
}

// BindToken is a paid mutator transaction binding the contract method 0x8eff336c.
//
// Solidity: function bindToken(bytes32 bep2Symbol, address contractAddr, uint256 decimals) returns()
func (_ITokenHub *ITokenHubSession) BindToken(bep2Symbol [32]byte, contractAddr common.Address, decimals *big.Int) (*types.Transaction, error) {
	return _ITokenHub.Contract.BindToken(&_ITokenHub.TransactOpts, bep2Symbol, contractAddr, decimals)
}

// BindToken is a paid mutator transaction binding the contract method 0x8eff336c.
//
// Solidity: function bindToken(bytes32 bep2Symbol, address contractAddr, uint256 decimals) returns()
func (_ITokenHub *ITokenHubTransactorSession) BindToken(bep2Symbol [32]byte, contractAddr common.Address, decimals *big.Int) (*types.Transaction, error) {
	return _ITokenHub.Contract.BindToken(&_ITokenHub.TransactOpts, bep2Symbol, contractAddr, decimals)
}

// TransferOut is a paid mutator transaction binding the contract method 0xaa7415f5.
//
// Solidity: function transferOut(address contractAddr, address recipient, uint256 amount, uint64 expireTime) payable returns(bool)
func (_ITokenHub *ITokenHubTransactor) TransferOut(opts *bind.TransactOpts, contractAddr common.Address, recipient common.Address, amount *big.Int, expireTime uint64) (*types.Transaction, error) {
	return _ITokenHub.contract.Transact(opts, "transferOut", contractAddr, recipient, amount, expireTime)
}

// TransferOut is a paid mutator transaction binding the contract method 0xaa7415f5.
//
// Solidity: function transferOut(address contractAddr, address recipient, uint256 amount, uint64 expireTime) payable returns(bool)
func (_ITokenHub *ITokenHubSession) TransferOut(contractAddr common.Address, recipient common.Address, amount *big.Int, expireTime uint64) (*types.Transaction, error) {
	return _ITokenHub.Contract.TransferOut(&_ITokenHub.TransactOpts, contractAddr, recipient, amount, expireTime)
}

// TransferOut is a paid mutator transaction binding the contract method 0xaa7415f5.
//
// Solidity: function transferOut(address contractAddr, address recipient, uint256 amount, uint64 expireTime) payable returns(bool)
func (_ITokenHub *ITokenHubTransactorSession) TransferOut(contractAddr common.Address, recipient common.Address, amount *big.Int, expireTime uint64) (*types.Transaction, error) {
	return _ITokenHub.Contract.TransferOut(&_ITokenHub.TransactOpts, contractAddr, recipient, amount, expireTime)
}

// UnbindToken is a paid mutator transaction binding the contract method 0xb99328c5.
//
// Solidity: function unbindToken(bytes32 bep2Symbol, address contractAddr) returns()
func (_ITokenHub *ITokenHubTransactor) UnbindToken(opts *bind.TransactOpts, bep2Symbol [32]byte, contractAddr common.Address) (*types.Transaction, error) {
	return _ITokenHub.contract.Transact(opts, "unbindToken", bep2Symbol, contractAddr)
}

// UnbindToken is a paid mutator transaction binding the contract method 0xb99328c5.
//
// Solidity: function unbindToken(bytes32 bep2Symbol, address contractAddr) returns()
func (_ITokenHub *ITokenHubSession) UnbindToken(bep2Symbol [32]byte, contractAddr common.Address) (*types.Transaction, error) {
	return _ITokenHub.Contract.UnbindToken(&_ITokenHub.TransactOpts, bep2Symbol, contractAddr)
}

// UnbindToken is a paid mutator transaction binding the contract method 0xb99328c5.
//
// Solidity: function unbindToken(bytes32 bep2Symbol, address contractAddr) returns()
func (_ITokenHub *ITokenHubTransactorSession) UnbindToken(bep2Symbol [32]byte, contractAddr common.Address) (*types.Transaction, error) {
	return _ITokenHub.Contract.UnbindToken(&_ITokenHub.TransactOpts, bep2Symbol, contractAddr)
}

// MemoryABI is the input ABI used to generate the binding from.
const MemoryABI = "[]"

// MemoryBin is the compiled bytecode used for deploying new contracts.
var MemoryBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209323a48b104f52ad6049da30c049f2d948fdd51593f63649bed67841938c956664736f6c63430006040033"

// DeployMemory deploys a new Ethereum contract, binding an instance of Memory to it.
func DeployMemory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Memory, error) {
	parsed, err := abi.JSON(strings.NewReader(MemoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MemoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Memory{MemoryCaller: MemoryCaller{contract: contract}, MemoryTransactor: MemoryTransactor{contract: contract}, MemoryFilterer: MemoryFilterer{contract: contract}}, nil
}

// Memory is an auto generated Go binding around an Ethereum contract.
type Memory struct {
	MemoryCaller     // Read-only binding to the contract
	MemoryTransactor // Write-only binding to the contract
	MemoryFilterer   // Log filterer for contract events
}

// MemoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MemoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MemoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MemoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MemorySession struct {
	Contract     *Memory           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MemoryCallerSession struct {
	Contract *MemoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MemoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MemoryTransactorSession struct {
	Contract     *MemoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MemoryRaw struct {
	Contract *Memory // Generic contract binding to access the raw methods on
}

// MemoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MemoryCallerRaw struct {
	Contract *MemoryCaller // Generic read-only contract binding to access the raw methods on
}

// MemoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MemoryTransactorRaw struct {
	Contract *MemoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMemory creates a new instance of Memory, bound to a specific deployed contract.
func NewMemory(address common.Address, backend bind.ContractBackend) (*Memory, error) {
	contract, err := bindMemory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Memory{MemoryCaller: MemoryCaller{contract: contract}, MemoryTransactor: MemoryTransactor{contract: contract}, MemoryFilterer: MemoryFilterer{contract: contract}}, nil
}

// NewMemoryCaller creates a new read-only instance of Memory, bound to a specific deployed contract.
func NewMemoryCaller(address common.Address, caller bind.ContractCaller) (*MemoryCaller, error) {
	contract, err := bindMemory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MemoryCaller{contract: contract}, nil
}

// NewMemoryTransactor creates a new write-only instance of Memory, bound to a specific deployed contract.
func NewMemoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MemoryTransactor, error) {
	contract, err := bindMemory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MemoryTransactor{contract: contract}, nil
}

// NewMemoryFilterer creates a new log filterer instance of Memory, bound to a specific deployed contract.
func NewMemoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MemoryFilterer, error) {
	contract, err := bindMemory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MemoryFilterer{contract: contract}, nil
}

// bindMemory binds a generic wrapper to an already deployed contract.
func bindMemory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MemoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Memory *MemoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Memory.Contract.MemoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Memory *MemoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Memory.Contract.MemoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Memory *MemoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Memory.Contract.MemoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Memory *MemoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Memory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Memory *MemoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Memory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Memory *MemoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Memory.Contract.contract.Transact(opts, method, params...)
}

// RLPDecodeABI is the input ABI used to generate the binding from.
const RLPDecodeABI = "[]"

// RLPDecodeBin is the compiled bytecode used for deploying new contracts.
var RLPDecodeBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220181baf58d2f79353375cbba9ae73c1d7d112f8842f92bc48cb128f13700b978964736f6c63430006040033"

// DeployRLPDecode deploys a new Ethereum contract, binding an instance of RLPDecode to it.
func DeployRLPDecode(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RLPDecode, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPDecodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RLPDecodeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RLPDecode{RLPDecodeCaller: RLPDecodeCaller{contract: contract}, RLPDecodeTransactor: RLPDecodeTransactor{contract: contract}, RLPDecodeFilterer: RLPDecodeFilterer{contract: contract}}, nil
}

// RLPDecode is an auto generated Go binding around an Ethereum contract.
type RLPDecode struct {
	RLPDecodeCaller     // Read-only binding to the contract
	RLPDecodeTransactor // Write-only binding to the contract
	RLPDecodeFilterer   // Log filterer for contract events
}

// RLPDecodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPDecodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPDecodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPDecodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPDecodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPDecodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPDecodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPDecodeSession struct {
	Contract     *RLPDecode        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPDecodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPDecodeCallerSession struct {
	Contract *RLPDecodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RLPDecodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPDecodeTransactorSession struct {
	Contract     *RLPDecodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RLPDecodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPDecodeRaw struct {
	Contract *RLPDecode // Generic contract binding to access the raw methods on
}

// RLPDecodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPDecodeCallerRaw struct {
	Contract *RLPDecodeCaller // Generic read-only contract binding to access the raw methods on
}

// RLPDecodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPDecodeTransactorRaw struct {
	Contract *RLPDecodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLPDecode creates a new instance of RLPDecode, bound to a specific deployed contract.
func NewRLPDecode(address common.Address, backend bind.ContractBackend) (*RLPDecode, error) {
	contract, err := bindRLPDecode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLPDecode{RLPDecodeCaller: RLPDecodeCaller{contract: contract}, RLPDecodeTransactor: RLPDecodeTransactor{contract: contract}, RLPDecodeFilterer: RLPDecodeFilterer{contract: contract}}, nil
}

// NewRLPDecodeCaller creates a new read-only instance of RLPDecode, bound to a specific deployed contract.
func NewRLPDecodeCaller(address common.Address, caller bind.ContractCaller) (*RLPDecodeCaller, error) {
	contract, err := bindRLPDecode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPDecodeCaller{contract: contract}, nil
}

// NewRLPDecodeTransactor creates a new write-only instance of RLPDecode, bound to a specific deployed contract.
func NewRLPDecodeTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPDecodeTransactor, error) {
	contract, err := bindRLPDecode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPDecodeTransactor{contract: contract}, nil
}

// NewRLPDecodeFilterer creates a new log filterer instance of RLPDecode, bound to a specific deployed contract.
func NewRLPDecodeFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPDecodeFilterer, error) {
	contract, err := bindRLPDecode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPDecodeFilterer{contract: contract}, nil
}

// bindRLPDecode binds a generic wrapper to an already deployed contract.
func bindRLPDecode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPDecodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPDecode *RLPDecodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPDecode.Contract.RLPDecodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPDecode *RLPDecodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPDecode.Contract.RLPDecodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPDecode *RLPDecodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPDecode.Contract.RLPDecodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPDecode *RLPDecodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPDecode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPDecode *RLPDecodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPDecode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPDecode *RLPDecodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPDecode.Contract.contract.Transact(opts, method, params...)
}

// RLPEncodeABI is the input ABI used to generate the binding from.
const RLPEncodeABI = "[]"

// RLPEncodeBin is the compiled bytecode used for deploying new contracts.
var RLPEncodeBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a7f8a9dea63ff28a573c4d26717079867be48faf1c6611bb94144eba7330cd6664736f6c63430006040033"

// DeployRLPEncode deploys a new Ethereum contract, binding an instance of RLPEncode to it.
func DeployRLPEncode(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RLPEncode, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPEncodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RLPEncodeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RLPEncode{RLPEncodeCaller: RLPEncodeCaller{contract: contract}, RLPEncodeTransactor: RLPEncodeTransactor{contract: contract}, RLPEncodeFilterer: RLPEncodeFilterer{contract: contract}}, nil
}

// RLPEncode is an auto generated Go binding around an Ethereum contract.
type RLPEncode struct {
	RLPEncodeCaller     // Read-only binding to the contract
	RLPEncodeTransactor // Write-only binding to the contract
	RLPEncodeFilterer   // Log filterer for contract events
}

// RLPEncodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPEncodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPEncodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPEncodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPEncodeSession struct {
	Contract     *RLPEncode        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPEncodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPEncodeCallerSession struct {
	Contract *RLPEncodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RLPEncodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPEncodeTransactorSession struct {
	Contract     *RLPEncodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RLPEncodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPEncodeRaw struct {
	Contract *RLPEncode // Generic contract binding to access the raw methods on
}

// RLPEncodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPEncodeCallerRaw struct {
	Contract *RLPEncodeCaller // Generic read-only contract binding to access the raw methods on
}

// RLPEncodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPEncodeTransactorRaw struct {
	Contract *RLPEncodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLPEncode creates a new instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncode(address common.Address, backend bind.ContractBackend) (*RLPEncode, error) {
	contract, err := bindRLPEncode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLPEncode{RLPEncodeCaller: RLPEncodeCaller{contract: contract}, RLPEncodeTransactor: RLPEncodeTransactor{contract: contract}, RLPEncodeFilterer: RLPEncodeFilterer{contract: contract}}, nil
}

// NewRLPEncodeCaller creates a new read-only instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeCaller(address common.Address, caller bind.ContractCaller) (*RLPEncodeCaller, error) {
	contract, err := bindRLPEncode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeCaller{contract: contract}, nil
}

// NewRLPEncodeTransactor creates a new write-only instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPEncodeTransactor, error) {
	contract, err := bindRLPEncode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeTransactor{contract: contract}, nil
}

// NewRLPEncodeFilterer creates a new log filterer instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPEncodeFilterer, error) {
	contract, err := bindRLPEncode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeFilterer{contract: contract}, nil
}

// bindRLPEncode binds a generic wrapper to an already deployed contract.
func bindRLPEncode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPEncodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPEncode *RLPEncodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPEncode.Contract.RLPEncodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPEncode *RLPEncodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPEncode.Contract.RLPEncodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPEncode *RLPEncodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPEncode.Contract.RLPEncodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPEncode *RLPEncodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPEncode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPEncode *RLPEncodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPEncode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPEncode *RLPEncodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPEncode.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220330623cf29bc2b24ef9154ec313ab44fb4077fbaf305c00404422c7fd32ab7bf64736f6c63430006040033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SystemABI is the input ABI used to generate the binding from.
const SystemABI = "[{\"inputs\":[],\"name\":\"BIND_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CODE_OK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_FAIL_DECODE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INCENTIVIZE_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM_REWARD_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_MANAGER_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bscChainID\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SystemFuncSigs maps the 4-byte function signature to its string representation.
var SystemFuncSigs = map[string]string{
	"3dffc387": "BIND_CHANNELID()",
	"ab51bb96": "CODE_OK()",
	"51e80672": "CROSS_CHAIN_CONTRACT_ADDR()",
	"0bee7a67": "ERROR_FAIL_DECODE()",
	"96713da9": "GOV_CHANNELID()",
	"9dc09262": "GOV_HUB_ADDR()",
	"6e47b482": "INCENTIVIZE_ADDR()",
	"dc927faf": "LIGHT_CLIENT_ADDR()",
	"a1a11bf5": "RELAYERHUB_CONTRACT_ADDR()",
	"7942fd05": "SLASH_CHANNELID()",
	"43756e5c": "SLASH_CONTRACT_ADDR()",
	"4bf6c882": "STAKING_CHANNELID()",
	"c81b1662": "SYSTEM_REWARD_ADDR()",
	"fd6a6879": "TOKEN_HUB_ADDR()",
	"75d47a0a": "TOKEN_MANAGER_ADDR()",
	"70fd5bad": "TRANSFER_IN_CHANNELID()",
	"fc3e5908": "TRANSFER_OUT_CHANNELID()",
	"f9a2bbc7": "VALIDATOR_CONTRACT_ADDR()",
	"a78abc16": "alreadyInit()",
	"493279b1": "bscChainID()",
}

// SystemBin is the compiled bytecode used for deploying new contracts.
var SystemBin = "0x608060405234801561001057600080fd5b506102ef806100206000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c806396713da9116100ad578063c81b166211610071578063c81b16621461021f578063dc927faf14610227578063f9a2bbc71461022f578063fc3e590814610237578063fd6a68791461023f5761012c565b806396713da9146101e35780639dc09262146101eb578063a1a11bf5146101f3578063a78abc16146101fb578063ab51bb96146102175761012c565b806351e80672116100f457806351e80672146101bb5780636e47b482146101c357806370fd5bad146101cb57806375d47a0a146101d35780637942fd05146101db5761012c565b80630bee7a67146101315780633dffc3871461015257806343756e5c14610170578063493279b1146101945780634bf6c882146101b3575b600080fd5b610139610247565b6040805163ffffffff9092168252519081900360200190f35b61015a61024c565b6040805160ff9092168252519081900360200190f35b610178610251565b604080516001600160a01b039092168252519081900360200190f35b61019c610257565b6040805161ffff9092168252519081900360200190f35b61015a61025c565b610178610261565b610178610267565b61015a61026d565b610178610272565b61015a610278565b61015a61027d565b610178610282565b610178610288565b61020361028e565b604080519115158252519081900360200190f35b610139610297565b61017861029c565b6101786102a2565b6101786102a8565b61015a6102ae565b6101786102b3565b606481565b600181565b61100181565b606081565b600881565b61200081565b61100581565b600281565b61100881565b600b81565b600981565b61100781565b61100681565b60005460ff1681565b600081565b61100281565b61100381565b61100081565b600381565b6110048156fea2646970667358221220f6bd4ecf625b4bcfbb66bc7843198370f23b5b7c20b542107de9b5c9ba2acf9464736f6c63430006040033"

// DeploySystem deploys a new Ethereum contract, binding an instance of System to it.
func DeploySystem(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *System, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SystemBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &System{SystemCaller: SystemCaller{contract: contract}, SystemTransactor: SystemTransactor{contract: contract}, SystemFilterer: SystemFilterer{contract: contract}}, nil
}

// System is an auto generated Go binding around an Ethereum contract.
type System struct {
	SystemCaller     // Read-only binding to the contract
	SystemTransactor // Write-only binding to the contract
	SystemFilterer   // Log filterer for contract events
}

// SystemCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemSession struct {
	Contract     *System           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemCallerSession struct {
	Contract *SystemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SystemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemTransactorSession struct {
	Contract     *SystemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemRaw struct {
	Contract *System // Generic contract binding to access the raw methods on
}

// SystemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemCallerRaw struct {
	Contract *SystemCaller // Generic read-only contract binding to access the raw methods on
}

// SystemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemTransactorRaw struct {
	Contract *SystemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystem creates a new instance of System, bound to a specific deployed contract.
func NewSystem(address common.Address, backend bind.ContractBackend) (*System, error) {
	contract, err := bindSystem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &System{SystemCaller: SystemCaller{contract: contract}, SystemTransactor: SystemTransactor{contract: contract}, SystemFilterer: SystemFilterer{contract: contract}}, nil
}

// NewSystemCaller creates a new read-only instance of System, bound to a specific deployed contract.
func NewSystemCaller(address common.Address, caller bind.ContractCaller) (*SystemCaller, error) {
	contract, err := bindSystem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemCaller{contract: contract}, nil
}

// NewSystemTransactor creates a new write-only instance of System, bound to a specific deployed contract.
func NewSystemTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemTransactor, error) {
	contract, err := bindSystem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemTransactor{contract: contract}, nil
}

// NewSystemFilterer creates a new log filterer instance of System, bound to a specific deployed contract.
func NewSystemFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemFilterer, error) {
	contract, err := bindSystem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemFilterer{contract: contract}, nil
}

// bindSystem binds a generic wrapper to an already deployed contract.
func bindSystem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_System *SystemRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _System.Contract.SystemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_System *SystemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _System.Contract.SystemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_System *SystemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _System.Contract.SystemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_System *SystemCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _System.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_System *SystemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _System.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_System *SystemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _System.Contract.contract.Transact(opts, method, params...)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_System *SystemCaller) BINDCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "BIND_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_System *SystemSession) BINDCHANNELID() (uint8, error) {
	return _System.Contract.BINDCHANNELID(&_System.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_System *SystemCallerSession) BINDCHANNELID() (uint8, error) {
	return _System.Contract.BINDCHANNELID(&_System.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_System *SystemCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "CODE_OK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_System *SystemSession) CODEOK() (uint32, error) {
	return _System.Contract.CODEOK(&_System.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_System *SystemCallerSession) CODEOK() (uint32, error) {
	return _System.Contract.CODEOK(&_System.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_System *SystemCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "CROSS_CHAIN_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_System *SystemSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _System.Contract.CROSSCHAINCONTRACTADDR(&_System.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_System *SystemCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _System.Contract.CROSSCHAINCONTRACTADDR(&_System.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_System *SystemCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "ERROR_FAIL_DECODE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_System *SystemSession) ERRORFAILDECODE() (uint32, error) {
	return _System.Contract.ERRORFAILDECODE(&_System.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_System *SystemCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _System.Contract.ERRORFAILDECODE(&_System.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_System *SystemCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "GOV_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_System *SystemSession) GOVCHANNELID() (uint8, error) {
	return _System.Contract.GOVCHANNELID(&_System.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_System *SystemCallerSession) GOVCHANNELID() (uint8, error) {
	return _System.Contract.GOVCHANNELID(&_System.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_System *SystemCaller) GOVHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "GOV_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_System *SystemSession) GOVHUBADDR() (common.Address, error) {
	return _System.Contract.GOVHUBADDR(&_System.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_System *SystemCallerSession) GOVHUBADDR() (common.Address, error) {
	return _System.Contract.GOVHUBADDR(&_System.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_System *SystemCaller) INCENTIVIZEADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "INCENTIVIZE_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_System *SystemSession) INCENTIVIZEADDR() (common.Address, error) {
	return _System.Contract.INCENTIVIZEADDR(&_System.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_System *SystemCallerSession) INCENTIVIZEADDR() (common.Address, error) {
	return _System.Contract.INCENTIVIZEADDR(&_System.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_System *SystemCaller) LIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "LIGHT_CLIENT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_System *SystemSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _System.Contract.LIGHTCLIENTADDR(&_System.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_System *SystemCallerSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _System.Contract.LIGHTCLIENTADDR(&_System.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_System *SystemCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "RELAYERHUB_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_System *SystemSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _System.Contract.RELAYERHUBCONTRACTADDR(&_System.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_System *SystemCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _System.Contract.RELAYERHUBCONTRACTADDR(&_System.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_System *SystemCaller) SLASHCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "SLASH_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_System *SystemSession) SLASHCHANNELID() (uint8, error) {
	return _System.Contract.SLASHCHANNELID(&_System.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_System *SystemCallerSession) SLASHCHANNELID() (uint8, error) {
	return _System.Contract.SLASHCHANNELID(&_System.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_System *SystemCaller) SLASHCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "SLASH_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_System *SystemSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _System.Contract.SLASHCONTRACTADDR(&_System.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_System *SystemCallerSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _System.Contract.SLASHCONTRACTADDR(&_System.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_System *SystemCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "STAKING_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_System *SystemSession) STAKINGCHANNELID() (uint8, error) {
	return _System.Contract.STAKINGCHANNELID(&_System.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_System *SystemCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _System.Contract.STAKINGCHANNELID(&_System.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_System *SystemCaller) SYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "SYSTEM_REWARD_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_System *SystemSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _System.Contract.SYSTEMREWARDADDR(&_System.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_System *SystemCallerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _System.Contract.SYSTEMREWARDADDR(&_System.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_System *SystemCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "TOKEN_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_System *SystemSession) TOKENHUBADDR() (common.Address, error) {
	return _System.Contract.TOKENHUBADDR(&_System.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_System *SystemCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _System.Contract.TOKENHUBADDR(&_System.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_System *SystemCaller) TOKENMANAGERADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "TOKEN_MANAGER_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_System *SystemSession) TOKENMANAGERADDR() (common.Address, error) {
	return _System.Contract.TOKENMANAGERADDR(&_System.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_System *SystemCallerSession) TOKENMANAGERADDR() (common.Address, error) {
	return _System.Contract.TOKENMANAGERADDR(&_System.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_System *SystemCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "TRANSFER_IN_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_System *SystemSession) TRANSFERINCHANNELID() (uint8, error) {
	return _System.Contract.TRANSFERINCHANNELID(&_System.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_System *SystemCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _System.Contract.TRANSFERINCHANNELID(&_System.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_System *SystemCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "TRANSFER_OUT_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_System *SystemSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _System.Contract.TRANSFEROUTCHANNELID(&_System.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_System *SystemCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _System.Contract.TRANSFEROUTCHANNELID(&_System.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_System *SystemCaller) VALIDATORCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "VALIDATOR_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_System *SystemSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _System.Contract.VALIDATORCONTRACTADDR(&_System.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_System *SystemCallerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _System.Contract.VALIDATORCONTRACTADDR(&_System.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_System *SystemCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "alreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_System *SystemSession) AlreadyInit() (bool, error) {
	return _System.Contract.AlreadyInit(&_System.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_System *SystemCallerSession) AlreadyInit() (bool, error) {
	return _System.Contract.AlreadyInit(&_System.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_System *SystemCaller) BscChainID(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "bscChainID")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_System *SystemSession) BscChainID() (uint16, error) {
	return _System.Contract.BscChainID(&_System.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_System *SystemCallerSession) BscChainID() (uint16, error) {
	return _System.Contract.BscChainID(&_System.CallOpts)
}

// TransferHelperABI is the input ABI used to generate the binding from.
const TransferHelperABI = "[]"

// TransferHelperBin is the compiled bytecode used for deploying new contracts.
var TransferHelperBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e414503167cbe4bcb728ac62f35649e58b31915260cc7e505b9771c72fffd60064736f6c63430006040033"

// DeployTransferHelper deploys a new Ethereum contract, binding an instance of TransferHelper to it.
func DeployTransferHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TransferHelper, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferHelperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TransferHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// TransferHelper is an auto generated Go binding around an Ethereum contract.
type TransferHelper struct {
	TransferHelperCaller     // Read-only binding to the contract
	TransferHelperTransactor // Write-only binding to the contract
	TransferHelperFilterer   // Log filterer for contract events
}

// TransferHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferHelperSession struct {
	Contract     *TransferHelper   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferHelperCallerSession struct {
	Contract *TransferHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TransferHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferHelperTransactorSession struct {
	Contract     *TransferHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TransferHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferHelperRaw struct {
	Contract *TransferHelper // Generic contract binding to access the raw methods on
}

// TransferHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferHelperCallerRaw struct {
	Contract *TransferHelperCaller // Generic read-only contract binding to access the raw methods on
}

// TransferHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferHelperTransactorRaw struct {
	Contract *TransferHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferHelper creates a new instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelper(address common.Address, backend bind.ContractBackend) (*TransferHelper, error) {
	contract, err := bindTransferHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// NewTransferHelperCaller creates a new read-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperCaller(address common.Address, caller bind.ContractCaller) (*TransferHelperCaller, error) {
	contract, err := bindTransferHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperCaller{contract: contract}, nil
}

// NewTransferHelperTransactor creates a new write-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferHelperTransactor, error) {
	contract, err := bindTransferHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperTransactor{contract: contract}, nil
}

// NewTransferHelperFilterer creates a new log filterer instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferHelperFilterer, error) {
	contract, err := bindTransferHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferHelperFilterer{contract: contract}, nil
}

// bindTransferHelper binds a generic wrapper to an already deployed contract.
func bindTransferHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.TransferHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transact(opts, method, params...)
}
