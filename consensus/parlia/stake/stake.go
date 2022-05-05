// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stake

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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"div\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"mul\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SafeMathFuncSigs maps the 4-byte function signature to its string representation.
var SafeMathFuncSigs = map[string]string{
	"771602f7": "add(uint256,uint256)",
	"a391c15b": "div(uint256,uint256)",
	"c8a4ac9c": "mul(uint256,uint256)",
	"b67d77c5": "sub(uint256,uint256)",
}

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x61018c610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c8063771602f71461005b578063a391c15b14610090578063b67d77c5146100b3578063c8a4ac9c146100d6575b600080fd5b61007e6004803603604081101561007157600080fd5b50803590602001356100f9565b60408051918252519081900360200190f35b61007e600480360360408110156100a657600080fd5b508035906020013561010f565b61007e600480360360408110156100c957600080fd5b5080359060200135610124565b61007e600480360360408110156100ec57600080fd5b5080359060200135610136565b60008282018381101561010857fe5b9392505050565b60008082848161011b57fe5b04949350505050565b60008282111561013057fe5b50900390565b600082820283158061015057508284828161014d57fe5b04145b61010857fefea264697066735822122087673e8010980c6cbbcce9f20836d7a8369e8891f72c43a8398ad5401bbcd49364736f6c63430006040033"

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

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 a, uint256 b) pure returns(uint256)
func (_SafeMath *SafeMathCaller) Add(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SafeMath.contract.Call(opts, &out, "add", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 a, uint256 b) pure returns(uint256)
func (_SafeMath *SafeMathSession) Add(a *big.Int, b *big.Int) (*big.Int, error) {
	return _SafeMath.Contract.Add(&_SafeMath.CallOpts, a, b)
}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 a, uint256 b) pure returns(uint256)
func (_SafeMath *SafeMathCallerSession) Add(a *big.Int, b *big.Int) (*big.Int, error) {
	return _SafeMath.Contract.Add(&_SafeMath.CallOpts, a, b)
}

// Div is a free data retrieval call binding the contract method 0xa391c15b.
//
// Solidity: function div(uint256 a, uint256 b) pure returns(uint256)
func (_SafeMath *SafeMathCaller) Div(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SafeMath.contract.Call(opts, &out, "div", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Div is a free data retrieval call binding the contract method 0xa391c15b.
//
// Solidity: function div(uint256 a, uint256 b) pure returns(uint256)
func (_SafeMath *SafeMathSession) Div(a *big.Int, b *big.Int) (*big.Int, error) {
	return _SafeMath.Contract.Div(&_SafeMath.CallOpts, a, b)
}

// Div is a free data retrieval call binding the contract method 0xa391c15b.
//
// Solidity: function div(uint256 a, uint256 b) pure returns(uint256)
func (_SafeMath *SafeMathCallerSession) Div(a *big.Int, b *big.Int) (*big.Int, error) {
	return _SafeMath.Contract.Div(&_SafeMath.CallOpts, a, b)
}

// Mul is a free data retrieval call binding the contract method 0xc8a4ac9c.
//
// Solidity: function mul(uint256 a, uint256 b) pure returns(uint256)
func (_SafeMath *SafeMathCaller) Mul(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SafeMath.contract.Call(opts, &out, "mul", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Mul is a free data retrieval call binding the contract method 0xc8a4ac9c.
//
// Solidity: function mul(uint256 a, uint256 b) pure returns(uint256)
func (_SafeMath *SafeMathSession) Mul(a *big.Int, b *big.Int) (*big.Int, error) {
	return _SafeMath.Contract.Mul(&_SafeMath.CallOpts, a, b)
}

// Mul is a free data retrieval call binding the contract method 0xc8a4ac9c.
//
// Solidity: function mul(uint256 a, uint256 b) pure returns(uint256)
func (_SafeMath *SafeMathCallerSession) Mul(a *big.Int, b *big.Int) (*big.Int, error) {
	return _SafeMath.Contract.Mul(&_SafeMath.CallOpts, a, b)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_SafeMath *SafeMathCaller) Sub(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SafeMath.contract.Call(opts, &out, "sub", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_SafeMath *SafeMathSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _SafeMath.Contract.Sub(&_SafeMath.CallOpts, a, b)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_SafeMath *SafeMathCallerSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _SafeMath.Contract.Sub(&_SafeMath.CallOpts, a, b)
}

// StakeABI is the input ABI used to generate the binding from.
const StakeABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"DistributeBlockRewardFunc\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"UserStakeFFF\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"UserUnStakeFFF\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"CreateAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"DistributeBlockReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"StakeFFF\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UnStakeFFF\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currStakeFFF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currStakePeople\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxStakeFFFCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxStakePeopleCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miniStakeFFFCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"myBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeInfoIndexMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakeInfoMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StakeFuncSigs maps the 4-byte function signature to its string representation.
var StakeFuncSigs = map[string]string{
	"2463cd09": "CreateAddress(uint256)",
	"50403540": "DistributeBlockReward(uint256)",
	"0db6cccb": "StakeFFF()",
	"30c62a13": "UnStakeFFF()",
	"101a875f": "currStakeFFF()",
	"93a101f3": "currStakePeople()",
	"b2ccf3a4": "maxStakeFFFCount()",
	"f78ffcbc": "maxStakePeopleCount()",
	"1faea20b": "miniStakeFFFCount()",
	"c9116b69": "myBalance()",
	"8da5cb5b": "owner()",
	"7dedf367": "stakeInfoIndexMap(uint256)",
	"c5f30db5": "stakeInfoMap(address)",
}

// StakeBin is the compiled bytecode used for deploying new contracts.
var StakeBin = "0x608060405269152d02c7e14af6800000600055670de0b6b3a7640000600155620186a060025534801561003157600080fd5b50600580546001600160a01b03191633179055610967806100536000396000f3fe6080604052600436106100c25760003560e01c80637dedf3671161007f578063b2ccf3a411610059578063b2ccf3a4146101cc578063c5f30db5146101e1578063c9116b6914610214578063f78ffcbc14610229576100c2565b80637dedf3671461015c5780638da5cb5b146101a257806393a101f3146101b7576100c2565b80630db6cccb146100c7578063101a875f146100d15780631faea20b146100f85780632463cd091461010d57806330c62a1314610137578063504035401461013f575b600080fd5b6100cf61023e565b005b3480156100dd57600080fd5b506100e661024a565b60408051918252519081900360200190f35b34801561010457600080fd5b506100e6610250565b34801561011957600080fd5b506100cf6004803603602081101561013057600080fd5b5035610256565b6100cf6102af565b6100cf6004803603602081101561015557600080fd5b50356103ac565b34801561016857600080fd5b506101866004803603602081101561017f57600080fd5b50356105d5565b604080516001600160a01b039092168252519081900360200190f35b3480156101ae57600080fd5b506101866105f0565b3480156101c357600080fd5b506100e66105ff565b3480156101d857600080fd5b506100e6610605565b3480156101ed57600080fd5b506100e66004803603602081101561020457600080fd5b50356001600160a01b031661060b565b34801561022057600080fd5b506100e661061d565b34801561023557600080fd5b506100e6610621565b6102483334610627565b565b60045481565b60015481565b60005b818110156102ab576040805160208082018490524282840152434060608084019190915283518084039091018152608090920190925280519101206001546102a2908290610627565b50600101610259565b5050565b336000818152600660205260409020546102fe576040805162461bcd60e51b815260206004820152600b60248201526a1b5d5cdd081cdd185ad95960aa1b604482015290519081900360640190fd5b3360009081526006602052604080822054905190916001600160a01b0384169183156108fc0291849190818181858888f19350505050158015610345573d6000803e3d6000fd5b5033600081815260066020908152604080832080546004805491909103905560038054600019019055929092558151928352820183905280517f433979128f828c623a175588cd9821124640d4353dd48d812d2b7273f1b76c3f9281900390910190a15050565b6005546001600160a01b03163314610401576040805162461bcd60e51b815260206004820152601360248201527221b0b63632b91034b9903737ba1037bbb732b960691b604482015290519081900360640190fd5b60005b600354811015610594576000818152600760209081526040808320546001600160a01b0316808452600683528184205482516332292b2760e21b81526004810182905260248101889052925191949093909273__$acdf508e5482d1bb24a67f58b1c293047c$__9263c8a4ac9c92604480840193919291829003018186803b15801561048f57600080fd5b505af41580156104a3573d6000803e3d6000fd5b505050506040513d60208110156104b957600080fd5b5051600480546040805163a391c15b60e01b8152928301939093526024820152905173__$acdf508e5482d1bb24a67f58b1c293047c$__9163a391c15b916044808301926020929190829003018186803b15801561051657600080fd5b505af415801561052a573d6000803e3d6000fd5b505050506040513d602081101561054057600080fd5b5051905061054e8382610816565b6040516001600160a01b0384169082156108fc029083906000818181858888f19350505050158015610584573d6000803e3d6000fd5b5050600190920191506104049050565b50600354604080519182526020820183905280517f77affbec109ff6a3609c802050db936c55bc48fb23db5c88f4a4818922e652399281900390910190a150565b6007602052600090815260409020546001600160a01b031681565b6005546001600160a01b031681565b60035481565b60005481565b60066020526000908152604090205481565b4790565b60025481565b600154819081108015906106565750600080546001600160a01b03851682526006602052604090912054820111155b6106a0576040805162461bcd60e51b81526020600482015260166024820152751cdd185ad94818dbdd5b9d081b5d5cdd081d985b1a5960521b604482015290519081900360640190fd5b6001600160a01b0383166000908152600660205260409020541580156106cd575060025460035460010111155b806106ef57506001600160a01b03831660009081526006602052604090205415155b610740576040805162461bcd60e51b815260206004820152601960248201527f7374616b652070656f706c65206f7574206f662072616e676500000000000000604482015290519081900360640190fd5b6001600160a01b0383166000908152600660205260409020546107a65760038054600090815260076020908152604080832080546001600160a01b0319166001600160a01b038916908117909155835260069091529020829055805460010190556107c5565b6001600160a01b03831660009081526006602052604090208054820190555b6004805482019055604080516001600160a01b03851681526020810183905281517f26a5c0f8e27c8ffa9a8414a94d6c0137a2ec1446b6925835dc0cf5ba74d2967f929181900390910190a1505050565b604080516000808252602082019092526001600160a01b0384169083906040518082805190602001908083835b602083106108625780518252601f199092019160209182019101610843565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d80600081146108c4576040519150601f19603f3d011682016040523d82523d6000602084013e6108c9565b606091505b50509050806109095760405162461bcd60e51b815260040180806020018281038252602381526020018061090f6023913960400191505060405180910390fd5b50505056fe5472616e7366657248656c7065723a205452585f5452414e534645525f4641494c4544a264697066735822122001cfc6e31865b773107d90a52f400e68c3b2c369421fac97b7c73e40fc6c4bf264736f6c63430006040033"

// DeployStake deploys a new Ethereum contract, binding an instance of Stake to it.
func DeployStake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Stake, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	safeMathAddr, _, _, _ := DeploySafeMath(auth, backend)
	StakeBin = strings.Replace(StakeBin, "__$acdf508e5482d1bb24a67f58b1c293047c$__", safeMathAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StakeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stake{StakeCaller: StakeCaller{contract: contract}, StakeTransactor: StakeTransactor{contract: contract}, StakeFilterer: StakeFilterer{contract: contract}}, nil
}

// Stake is an auto generated Go binding around an Ethereum contract.
type Stake struct {
	StakeCaller     // Read-only binding to the contract
	StakeTransactor // Write-only binding to the contract
	StakeFilterer   // Log filterer for contract events
}

// StakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeSession struct {
	Contract     *Stake            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeCallerSession struct {
	Contract *StakeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeTransactorSession struct {
	Contract     *StakeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeRaw struct {
	Contract *Stake // Generic contract binding to access the raw methods on
}

// StakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeCallerRaw struct {
	Contract *StakeCaller // Generic read-only contract binding to access the raw methods on
}

// StakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeTransactorRaw struct {
	Contract *StakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStake creates a new instance of Stake, bound to a specific deployed contract.
func NewStake(address common.Address, backend bind.ContractBackend) (*Stake, error) {
	contract, err := bindStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stake{StakeCaller: StakeCaller{contract: contract}, StakeTransactor: StakeTransactor{contract: contract}, StakeFilterer: StakeFilterer{contract: contract}}, nil
}

// NewStakeCaller creates a new read-only instance of Stake, bound to a specific deployed contract.
func NewStakeCaller(address common.Address, caller bind.ContractCaller) (*StakeCaller, error) {
	contract, err := bindStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeCaller{contract: contract}, nil
}

// NewStakeTransactor creates a new write-only instance of Stake, bound to a specific deployed contract.
func NewStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeTransactor, error) {
	contract, err := bindStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeTransactor{contract: contract}, nil
}

// NewStakeFilterer creates a new log filterer instance of Stake, bound to a specific deployed contract.
func NewStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeFilterer, error) {
	contract, err := bindStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeFilterer{contract: contract}, nil
}

// bindStake binds a generic wrapper to an already deployed contract.
func bindStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.StakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transact(opts, method, params...)
}

// CurrStakeFFF is a free data retrieval call binding the contract method 0x101a875f.
//
// Solidity: function currStakeFFF() view returns(uint256)
func (_Stake *StakeCaller) CurrStakeFFF(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "currStakeFFF")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrStakeFFF is a free data retrieval call binding the contract method 0x101a875f.
//
// Solidity: function currStakeFFF() view returns(uint256)
func (_Stake *StakeSession) CurrStakeFFF() (*big.Int, error) {
	return _Stake.Contract.CurrStakeFFF(&_Stake.CallOpts)
}

// CurrStakeFFF is a free data retrieval call binding the contract method 0x101a875f.
//
// Solidity: function currStakeFFF() view returns(uint256)
func (_Stake *StakeCallerSession) CurrStakeFFF() (*big.Int, error) {
	return _Stake.Contract.CurrStakeFFF(&_Stake.CallOpts)
}

// CurrStakePeople is a free data retrieval call binding the contract method 0x93a101f3.
//
// Solidity: function currStakePeople() view returns(uint256)
func (_Stake *StakeCaller) CurrStakePeople(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "currStakePeople")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrStakePeople is a free data retrieval call binding the contract method 0x93a101f3.
//
// Solidity: function currStakePeople() view returns(uint256)
func (_Stake *StakeSession) CurrStakePeople() (*big.Int, error) {
	return _Stake.Contract.CurrStakePeople(&_Stake.CallOpts)
}

// CurrStakePeople is a free data retrieval call binding the contract method 0x93a101f3.
//
// Solidity: function currStakePeople() view returns(uint256)
func (_Stake *StakeCallerSession) CurrStakePeople() (*big.Int, error) {
	return _Stake.Contract.CurrStakePeople(&_Stake.CallOpts)
}

// MaxStakeFFFCount is a free data retrieval call binding the contract method 0xb2ccf3a4.
//
// Solidity: function maxStakeFFFCount() view returns(uint256)
func (_Stake *StakeCaller) MaxStakeFFFCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "maxStakeFFFCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxStakeFFFCount is a free data retrieval call binding the contract method 0xb2ccf3a4.
//
// Solidity: function maxStakeFFFCount() view returns(uint256)
func (_Stake *StakeSession) MaxStakeFFFCount() (*big.Int, error) {
	return _Stake.Contract.MaxStakeFFFCount(&_Stake.CallOpts)
}

// MaxStakeFFFCount is a free data retrieval call binding the contract method 0xb2ccf3a4.
//
// Solidity: function maxStakeFFFCount() view returns(uint256)
func (_Stake *StakeCallerSession) MaxStakeFFFCount() (*big.Int, error) {
	return _Stake.Contract.MaxStakeFFFCount(&_Stake.CallOpts)
}

// MaxStakePeopleCount is a free data retrieval call binding the contract method 0xf78ffcbc.
//
// Solidity: function maxStakePeopleCount() view returns(uint256)
func (_Stake *StakeCaller) MaxStakePeopleCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "maxStakePeopleCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxStakePeopleCount is a free data retrieval call binding the contract method 0xf78ffcbc.
//
// Solidity: function maxStakePeopleCount() view returns(uint256)
func (_Stake *StakeSession) MaxStakePeopleCount() (*big.Int, error) {
	return _Stake.Contract.MaxStakePeopleCount(&_Stake.CallOpts)
}

// MaxStakePeopleCount is a free data retrieval call binding the contract method 0xf78ffcbc.
//
// Solidity: function maxStakePeopleCount() view returns(uint256)
func (_Stake *StakeCallerSession) MaxStakePeopleCount() (*big.Int, error) {
	return _Stake.Contract.MaxStakePeopleCount(&_Stake.CallOpts)
}

// MiniStakeFFFCount is a free data retrieval call binding the contract method 0x1faea20b.
//
// Solidity: function miniStakeFFFCount() view returns(uint256)
func (_Stake *StakeCaller) MiniStakeFFFCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "miniStakeFFFCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiniStakeFFFCount is a free data retrieval call binding the contract method 0x1faea20b.
//
// Solidity: function miniStakeFFFCount() view returns(uint256)
func (_Stake *StakeSession) MiniStakeFFFCount() (*big.Int, error) {
	return _Stake.Contract.MiniStakeFFFCount(&_Stake.CallOpts)
}

// MiniStakeFFFCount is a free data retrieval call binding the contract method 0x1faea20b.
//
// Solidity: function miniStakeFFFCount() view returns(uint256)
func (_Stake *StakeCallerSession) MiniStakeFFFCount() (*big.Int, error) {
	return _Stake.Contract.MiniStakeFFFCount(&_Stake.CallOpts)
}

// MyBalance is a free data retrieval call binding the contract method 0xc9116b69.
//
// Solidity: function myBalance() view returns(uint256)
func (_Stake *StakeCaller) MyBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "myBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MyBalance is a free data retrieval call binding the contract method 0xc9116b69.
//
// Solidity: function myBalance() view returns(uint256)
func (_Stake *StakeSession) MyBalance() (*big.Int, error) {
	return _Stake.Contract.MyBalance(&_Stake.CallOpts)
}

// MyBalance is a free data retrieval call binding the contract method 0xc9116b69.
//
// Solidity: function myBalance() view returns(uint256)
func (_Stake *StakeCallerSession) MyBalance() (*big.Int, error) {
	return _Stake.Contract.MyBalance(&_Stake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeSession) Owner() (common.Address, error) {
	return _Stake.Contract.Owner(&_Stake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeCallerSession) Owner() (common.Address, error) {
	return _Stake.Contract.Owner(&_Stake.CallOpts)
}

// StakeInfoIndexMap is a free data retrieval call binding the contract method 0x7dedf367.
//
// Solidity: function stakeInfoIndexMap(uint256 ) view returns(address)
func (_Stake *StakeCaller) StakeInfoIndexMap(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "stakeInfoIndexMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeInfoIndexMap is a free data retrieval call binding the contract method 0x7dedf367.
//
// Solidity: function stakeInfoIndexMap(uint256 ) view returns(address)
func (_Stake *StakeSession) StakeInfoIndexMap(arg0 *big.Int) (common.Address, error) {
	return _Stake.Contract.StakeInfoIndexMap(&_Stake.CallOpts, arg0)
}

// StakeInfoIndexMap is a free data retrieval call binding the contract method 0x7dedf367.
//
// Solidity: function stakeInfoIndexMap(uint256 ) view returns(address)
func (_Stake *StakeCallerSession) StakeInfoIndexMap(arg0 *big.Int) (common.Address, error) {
	return _Stake.Contract.StakeInfoIndexMap(&_Stake.CallOpts, arg0)
}

// StakeInfoMap is a free data retrieval call binding the contract method 0xc5f30db5.
//
// Solidity: function stakeInfoMap(address ) view returns(uint256)
func (_Stake *StakeCaller) StakeInfoMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "stakeInfoMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeInfoMap is a free data retrieval call binding the contract method 0xc5f30db5.
//
// Solidity: function stakeInfoMap(address ) view returns(uint256)
func (_Stake *StakeSession) StakeInfoMap(arg0 common.Address) (*big.Int, error) {
	return _Stake.Contract.StakeInfoMap(&_Stake.CallOpts, arg0)
}

// StakeInfoMap is a free data retrieval call binding the contract method 0xc5f30db5.
//
// Solidity: function stakeInfoMap(address ) view returns(uint256)
func (_Stake *StakeCallerSession) StakeInfoMap(arg0 common.Address) (*big.Int, error) {
	return _Stake.Contract.StakeInfoMap(&_Stake.CallOpts, arg0)
}

// CreateAddress is a paid mutator transaction binding the contract method 0x2463cd09.
//
// Solidity: function CreateAddress(uint256 count) returns()
func (_Stake *StakeTransactor) CreateAddress(opts *bind.TransactOpts, count *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "CreateAddress", count)
}

// CreateAddress is a paid mutator transaction binding the contract method 0x2463cd09.
//
// Solidity: function CreateAddress(uint256 count) returns()
func (_Stake *StakeSession) CreateAddress(count *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.CreateAddress(&_Stake.TransactOpts, count)
}

// CreateAddress is a paid mutator transaction binding the contract method 0x2463cd09.
//
// Solidity: function CreateAddress(uint256 count) returns()
func (_Stake *StakeTransactorSession) CreateAddress(count *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.CreateAddress(&_Stake.TransactOpts, count)
}

// DistributeBlockReward is a paid mutator transaction binding the contract method 0x50403540.
//
// Solidity: function DistributeBlockReward(uint256 reward) payable returns()
func (_Stake *StakeTransactor) DistributeBlockReward(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "DistributeBlockReward", reward)
}

// DistributeBlockReward is a paid mutator transaction binding the contract method 0x50403540.
//
// Solidity: function DistributeBlockReward(uint256 reward) payable returns()
func (_Stake *StakeSession) DistributeBlockReward(reward *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.DistributeBlockReward(&_Stake.TransactOpts, reward)
}

// DistributeBlockReward is a paid mutator transaction binding the contract method 0x50403540.
//
// Solidity: function DistributeBlockReward(uint256 reward) payable returns()
func (_Stake *StakeTransactorSession) DistributeBlockReward(reward *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.DistributeBlockReward(&_Stake.TransactOpts, reward)
}

// StakeFFF is a paid mutator transaction binding the contract method 0x0db6cccb.
//
// Solidity: function StakeFFF() payable returns()
func (_Stake *StakeTransactor) StakeFFF(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "StakeFFF")
}

// StakeFFF is a paid mutator transaction binding the contract method 0x0db6cccb.
//
// Solidity: function StakeFFF() payable returns()
func (_Stake *StakeSession) StakeFFF() (*types.Transaction, error) {
	return _Stake.Contract.StakeFFF(&_Stake.TransactOpts)
}

// StakeFFF is a paid mutator transaction binding the contract method 0x0db6cccb.
//
// Solidity: function StakeFFF() payable returns()
func (_Stake *StakeTransactorSession) StakeFFF() (*types.Transaction, error) {
	return _Stake.Contract.StakeFFF(&_Stake.TransactOpts)
}

// UnStakeFFF is a paid mutator transaction binding the contract method 0x30c62a13.
//
// Solidity: function UnStakeFFF() payable returns()
func (_Stake *StakeTransactor) UnStakeFFF(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "UnStakeFFF")
}

// UnStakeFFF is a paid mutator transaction binding the contract method 0x30c62a13.
//
// Solidity: function UnStakeFFF() payable returns()
func (_Stake *StakeSession) UnStakeFFF() (*types.Transaction, error) {
	return _Stake.Contract.UnStakeFFF(&_Stake.TransactOpts)
}

// UnStakeFFF is a paid mutator transaction binding the contract method 0x30c62a13.
//
// Solidity: function UnStakeFFF() payable returns()
func (_Stake *StakeTransactorSession) UnStakeFFF() (*types.Transaction, error) {
	return _Stake.Contract.UnStakeFFF(&_Stake.TransactOpts)
}

// StakeDistributeBlockRewardFuncIterator is returned from FilterDistributeBlockRewardFunc and is used to iterate over the raw logs and unpacked data for DistributeBlockRewardFunc events raised by the Stake contract.
type StakeDistributeBlockRewardFuncIterator struct {
	Event *StakeDistributeBlockRewardFunc // Event containing the contract specifics and raw log

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
func (it *StakeDistributeBlockRewardFuncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeDistributeBlockRewardFunc)
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
		it.Event = new(StakeDistributeBlockRewardFunc)
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
func (it *StakeDistributeBlockRewardFuncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeDistributeBlockRewardFuncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeDistributeBlockRewardFunc represents a DistributeBlockRewardFunc event raised by the Stake contract.
type StakeDistributeBlockRewardFunc struct {
	UserCount *big.Int
	Count     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDistributeBlockRewardFunc is a free log retrieval operation binding the contract event 0x77affbec109ff6a3609c802050db936c55bc48fb23db5c88f4a4818922e65239.
//
// Solidity: event DistributeBlockRewardFunc(uint256 userCount, uint256 count)
func (_Stake *StakeFilterer) FilterDistributeBlockRewardFunc(opts *bind.FilterOpts) (*StakeDistributeBlockRewardFuncIterator, error) {

	logs, sub, err := _Stake.contract.FilterLogs(opts, "DistributeBlockRewardFunc")
	if err != nil {
		return nil, err
	}
	return &StakeDistributeBlockRewardFuncIterator{contract: _Stake.contract, event: "DistributeBlockRewardFunc", logs: logs, sub: sub}, nil
}

// WatchDistributeBlockRewardFunc is a free log subscription operation binding the contract event 0x77affbec109ff6a3609c802050db936c55bc48fb23db5c88f4a4818922e65239.
//
// Solidity: event DistributeBlockRewardFunc(uint256 userCount, uint256 count)
func (_Stake *StakeFilterer) WatchDistributeBlockRewardFunc(opts *bind.WatchOpts, sink chan<- *StakeDistributeBlockRewardFunc) (event.Subscription, error) {

	logs, sub, err := _Stake.contract.WatchLogs(opts, "DistributeBlockRewardFunc")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeDistributeBlockRewardFunc)
				if err := _Stake.contract.UnpackLog(event, "DistributeBlockRewardFunc", log); err != nil {
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
func (_Stake *StakeFilterer) ParseDistributeBlockRewardFunc(log types.Log) (*StakeDistributeBlockRewardFunc, error) {
	event := new(StakeDistributeBlockRewardFunc)
	if err := _Stake.contract.UnpackLog(event, "DistributeBlockRewardFunc", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeUserStakeFFFIterator is returned from FilterUserStakeFFF and is used to iterate over the raw logs and unpacked data for UserStakeFFF events raised by the Stake contract.
type StakeUserStakeFFFIterator struct {
	Event *StakeUserStakeFFF // Event containing the contract specifics and raw log

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
func (it *StakeUserStakeFFFIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeUserStakeFFF)
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
		it.Event = new(StakeUserStakeFFF)
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
func (it *StakeUserStakeFFFIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeUserStakeFFFIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeUserStakeFFF represents a UserStakeFFF event raised by the Stake contract.
type StakeUserStakeFFF struct {
	User  common.Address
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUserStakeFFF is a free log retrieval operation binding the contract event 0x26a5c0f8e27c8ffa9a8414a94d6c0137a2ec1446b6925835dc0cf5ba74d2967f.
//
// Solidity: event UserStakeFFF(address user, uint256 count)
func (_Stake *StakeFilterer) FilterUserStakeFFF(opts *bind.FilterOpts) (*StakeUserStakeFFFIterator, error) {

	logs, sub, err := _Stake.contract.FilterLogs(opts, "UserStakeFFF")
	if err != nil {
		return nil, err
	}
	return &StakeUserStakeFFFIterator{contract: _Stake.contract, event: "UserStakeFFF", logs: logs, sub: sub}, nil
}

// WatchUserStakeFFF is a free log subscription operation binding the contract event 0x26a5c0f8e27c8ffa9a8414a94d6c0137a2ec1446b6925835dc0cf5ba74d2967f.
//
// Solidity: event UserStakeFFF(address user, uint256 count)
func (_Stake *StakeFilterer) WatchUserStakeFFF(opts *bind.WatchOpts, sink chan<- *StakeUserStakeFFF) (event.Subscription, error) {

	logs, sub, err := _Stake.contract.WatchLogs(opts, "UserStakeFFF")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeUserStakeFFF)
				if err := _Stake.contract.UnpackLog(event, "UserStakeFFF", log); err != nil {
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
func (_Stake *StakeFilterer) ParseUserStakeFFF(log types.Log) (*StakeUserStakeFFF, error) {
	event := new(StakeUserStakeFFF)
	if err := _Stake.contract.UnpackLog(event, "UserStakeFFF", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeUserUnStakeFFFIterator is returned from FilterUserUnStakeFFF and is used to iterate over the raw logs and unpacked data for UserUnStakeFFF events raised by the Stake contract.
type StakeUserUnStakeFFFIterator struct {
	Event *StakeUserUnStakeFFF // Event containing the contract specifics and raw log

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
func (it *StakeUserUnStakeFFFIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeUserUnStakeFFF)
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
		it.Event = new(StakeUserUnStakeFFF)
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
func (it *StakeUserUnStakeFFFIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeUserUnStakeFFFIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeUserUnStakeFFF represents a UserUnStakeFFF event raised by the Stake contract.
type StakeUserUnStakeFFF struct {
	User  common.Address
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUserUnStakeFFF is a free log retrieval operation binding the contract event 0x433979128f828c623a175588cd9821124640d4353dd48d812d2b7273f1b76c3f.
//
// Solidity: event UserUnStakeFFF(address user, uint256 count)
func (_Stake *StakeFilterer) FilterUserUnStakeFFF(opts *bind.FilterOpts) (*StakeUserUnStakeFFFIterator, error) {

	logs, sub, err := _Stake.contract.FilterLogs(opts, "UserUnStakeFFF")
	if err != nil {
		return nil, err
	}
	return &StakeUserUnStakeFFFIterator{contract: _Stake.contract, event: "UserUnStakeFFF", logs: logs, sub: sub}, nil
}

// WatchUserUnStakeFFF is a free log subscription operation binding the contract event 0x433979128f828c623a175588cd9821124640d4353dd48d812d2b7273f1b76c3f.
//
// Solidity: event UserUnStakeFFF(address user, uint256 count)
func (_Stake *StakeFilterer) WatchUserUnStakeFFF(opts *bind.WatchOpts, sink chan<- *StakeUserUnStakeFFF) (event.Subscription, error) {

	logs, sub, err := _Stake.contract.WatchLogs(opts, "UserUnStakeFFF")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeUserUnStakeFFF)
				if err := _Stake.contract.UnpackLog(event, "UserUnStakeFFF", log); err != nil {
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
func (_Stake *StakeFilterer) ParseUserUnStakeFFF(log types.Log) (*StakeUserUnStakeFFF, error) {
	event := new(StakeUserUnStakeFFF)
	if err := _Stake.contract.UnpackLog(event, "UserUnStakeFFF", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransferHelperABI is the input ABI used to generate the binding from.
const TransferHelperABI = "[]"

// TransferHelperBin is the compiled bytecode used for deploying new contracts.
var TransferHelperBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209f364ef21be806f75d59da71b3922406f64fbbee47408995b746ab15d6d5d4be64736f6c63430006040033"

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
