// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package golangtypes

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// FinalizeVerifierMetaData contains all meta data concerning the FinalizeVerifier contract.
var FinalizeVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[9]\",\"internalType\":\"uint256[9]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"input\",\"type\":\"uint256[9]\",\"internalType\":\"uint256[9]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x60808060405234601557611696908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace1114610b185750806344f6369214610a7f5780635f89feef1461060c5780638a3ae438146101f75763b8e72af614610055575f80fd5b346101bf5760403660031901126101bf576004356001600160401b0381116101bf57610085903690600401610b50565b6024356001600160401b0381116101bf576100a4903690600401610b50565b90918301610100848203126101bf5780601f850112156101bf57604051936100ce61010086610b7d565b849061010081019283116101bf57905b8282106101e7575050508101610120828203126101bf5780601f830112156101bf576040519161011061012084610b7d565b829061012081019283116101bf57905b8282106101c357505050303b156101bf576040516311475c8760e31b8152915f600484015b600882106101a95750505061010482015f905b60098210610193575050505f8161022481305afa80156101885761017a575080f35b61018691505f90610b7d565b005b6040513d5f823e3d90fd5b6020806001928551815201930191019091610158565b6020806001928551815201930191019091610145565b5f80fd5b8135815260209182019101610120565b634e487b7160e01b5f52604160045260245ffd5b81358152602091820191016100de565b346101bf576102203660031901126101bf5736610104116101bf5736610224116101bf5760405160408101905f5160206113215f395f51905f52815260208101915f5160206114e15f395f51905f5283525f5160206115615f395f51905f528152606082015f5160206114615f395f51905f5281525f5160206113c15f395f51905f52604061010435935f5160206115a15f395f51905f52608087019580875284848460608160075afa911016838860808160065afa165f5160206114415f395f51905f5283525f5160206112c15f395f51905f5286526001610124359182895286868660608160075afa9310161616838860808160065afa165f5160206113e15f395f51905f5283525f5160206111615f395f51905f528652610144359081885285858560608160075afa92101616838860808160065afa165f5160206114015f395f51905f5283525f5160206111a15f395f51905f528652610164359081885285858560608160075afa92101616838860808160065afa165f5160206113015f395f51905f5283525f5160206116415f395f51905f528652610184359081885285858560608160075afa92101616838860808160065afa165f5160206112a15f395f51905f5283525f5160206111c15f395f51905f5286526101a4359081885285858560608160075afa92101616838860808160065afa165f5160206113415f395f51905f5283525f5160206115c15f395f51905f5286526101c4359081885285858560608160075afa92101616838860808160065afa165f5160206112015f395f51905f5283525f5160206112415f395f51905f5286526101e4359081885285858560608160075afa92101616838860808160065afa16945f5160206112615f395f51905f528352526102043580955260608160075afa9210161660408260808160065afa169051915190156105fd5760405191610100600484375f5160206116215f395f51905f526101008401525f5160206115215f395f51905f526101208401525f5160206115815f395f51905f526101408401525f5160206114815f395f51905f526101608401525f5160206113815f395f51905f526101808401525f5160206113a15f395f51905f526101a08401525f5160206114c15f395f51905f526101c08401525f5160206111e15f395f51905f526101e08401525f5160206113615f395f51905f526102008401525f5160206115e15f395f51905f526102208401526102408301526102608201525f5160206112215f395f51905f526102808201525f5160206115415f395f51905f526102a08201525f5160206114215f395f51905f526102c08201525f5160206114a15f395f51905f526102e08201526020816103008160085afa905116156105ee57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b346101bf576101a03660031901126101bf57366084116101bf57366101a4116101bf5761030060405161063f8282610b7d565b8136823761064e600435610e29565b61065f602493929335604435610e94565b9193929061066e606435610e29565b9390926040519660408801965f5160206113215f395f51905f5289528860208101985f5160206114e15f395f51905f528a525f5160206115615f395f51905f5281525f5160206113c15f395f51905f52604060608401925f5160206114615f395f51905f5284525f5160206115a15f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f5160206114415f395f51905f5285525f5160206112c15f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f5160206113e15f395f51905f5285525f5160206111615f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f5160206114015f395f51905f5285525f5160206111a15f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f5160206113015f395f51905f5285525f5160206116415f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f5160206112a15f395f51905f5285525f5160206111c15f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f5160206113415f395f51905f5285525f5160206115c15f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f5160206112015f395f51905f5285525f5160206112415f395f51905f5288526101643590818a5287838760608160075afa921016169160808160065afa16945f5160206112615f395f51905f528352526101843580955260608160075afa9210161660408a60808160065afa169851975198156105fd5760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f5160206116215f395f51905f526101008401525f5160206115215f395f51905f526101208401525f5160206115815f395f51905f526101408401525f5160206114815f395f51905f526101608401525f5160206113815f395f51905f526101808401525f5160206113a15f395f51905f526101a08401525f5160206114c15f395f51905f526101c08401525f5160206111e15f395f51905f526101e08401525f5160206113615f395f51905f526102008401525f5160206115e15f395f51905f526102208401526102408301526102608201525f5160206112215f395f51905f526102808201525f5160206115415f395f51905f526102a08201525f5160206114215f395f51905f526102c08201525f5160206114a15f395f51905f526102e0820152604051928391610a5a8484610b7d565b8336843760085afa15908115610a72575b506105ee57005b6001915051141581610a6b565b346101bf576101003660031901126101bf5736610104116101bf57604051610aa8608082610b7d565b6080368237610abb602435600435610ba0565b8152610ad160843560a435604435606435610c41565b60208301526040820152610ae960e43560c435610ba0565b6060820152604051905f825b60048210610b0257608084f35b6020806001928551815201930191019091610af5565b346101bf575f3660031901126101bf57807f0cf61bd16ef3d481b31b40d655f1e901429f88911b1abdc675cf23dadcf86a0960209252f35b9181601f840112156101bf578235916001600160401b0383116101bf57602083818601950101116101bf57565b601f909101601f19168101906001600160401b038211908210176101d357604052565b905f5160206112815f395f51905f528210801590610c2a575b6105ee57811580610c22575b610c1c57610be95f5160206112815f395f51905f5260038185818180090908610f94565b818103610bf857505060011b90565b5f5160206112815f395f51905f52809106810306145f146105ee57600190811b1790565b50505f90565b508015610bc5565b505f5160206112815f395f51905f52811015610bb9565b919093925f5160206112815f395f51905f528310801590610e12575b8015610dfb575b8015610de4575b6105ee578082868517171715610dd957908291610d3c5f5160206112815f395f51905f5280808080888180808f9d5f5160206115015f395f51905f528f839290839109099d8e0981848181800909085f5160206116015f395f51905f52089a09818c8181800909085f5160206111815f395f51905f520806810306945f5160206112815f395f51905f525f5160206112e15f395f51905f5281610d1681808b80098187800908610f94565b8408095f5160206112815f395f51905f52610d30826110f8565b80091415958691610fb7565b929080821480610dd0575b15610d6e5750505050905f14610d665760ff60025b169060021b179190565b60ff5f610d5c565b5f5160206112815f395f51905f52809106810306149182610db1575b5050156105ee5760019115610da95760ff60025b169060021b17179190565b60ff5f610d9e565b5f5160206112815f395f51905f52919250819006810306145f80610d8a565b50838314610d47565b50505090505f905f90565b505f5160206112815f395f51905f52811015610c6b565b505f5160206112815f395f51905f52821015610c64565b505f5160206112815f395f51905f52851015610c5d565b8015610e8d578060011c915f5160206112815f395f51905f528310156105ee57600180610e6c5f5160206112815f395f51905f5260038188818180090908610f94565b931614610e7557565b905f5160206112815f395f51905f5280910681030690565b505f905f90565b801580610f8c575b610f80578060021c92825f5160206112815f395f51905f528510801590610f69575b6105ee5784815f5160206112815f395f51905f5280808080808080805f5160206115015f395f51905f5281610f339d8d0909998a0981898181800909085f5160206111815f395f51905f520806810306936002808a16149509818a8181800909085f5160206116015f395f51905f5208610fb7565b80929160018082961614610f45575050565b5f5160206112815f395f51905f528093945080929550809106810306930681030690565b505f5160206112815f395f51905f52811015610ebe565b50505f905f905f905f90565b508115610e9c565b90610f9e826110f8565b915f5160206112815f395f51905f52838009036105ee57565b915f5160206112815f395f51905f525f5160206112e15f395f51905f5281610ffc93969496610fee82808a8009818a800908610f94565b906110ec575b860809610f94565b925f5160206112815f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206112815f395f51905f5260a083015260208260c08160055afa915191156105ee575f5160206112815f395f51905f528260019209036105ee575f5160206112815f395f51905f52908209925f5160206112815f395f51905f5280808087800906810306818780090814908115916110cd575b506105ee57565b90505f5160206112815f395f51905f528084860960020914155f6110c6565b81809106810306610ff4565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206112815f395f51905f5260a083015260208260c08160055afa915191156105ee5756fe04c4b760e57455941e99953f801a730759de6f33e06ea7d4e9864de2bf7f20622fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7751f105f44cde8aadc3ef2fabe93ba53de37d0727af60d073209c2721c1c97bd6a183425c5153d3186990b81332849ae36e76bf41b7ce6d1901811e5b39e816a401a7952434af1a3d085fa161b53e482e911e1b45437fabdb20266e4c279ff9e0b2c33e0f7f4c6e704939398e0c1d851c00b5a61009e638ad1f97cebcc4e040aa21bdd92f5a3d1ba460a898bf9174e440fd5edb105c76c608b7be33428aeaa30fe160a5abbea61f3c1c6b26152a0f803ab2438a843640806cf6a4ea5c481d50243074b28f177916387ad2fc19b3277449caa4d3a7ffd75508775c0f08df5a3fde730644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd472c4d66518021a50c85b8820e8c7c69024842e8658e04ba4a72d25e3a02780f5e1c2e97cfef693d81f1bb57f407e312f1d93d4be8a25869684a11e0b266777e9b183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea40af0aa60c1934f85e905e794d1acc4c473dff87d5ce9213395684e0624a4baaf1e50a2f08fef07f96a02b2c021f1de268bfeaffc2cd4b01c60f858a0d8ee86ce269fc008a5b2ea0e125c03296f22c6142286112e0f9fd60e99702ba01d82f3470457d22e9e575ba25f0bd07f1687173d4abc1a071ea2dc5c0fc7ed879657fd4812ec2eda8a58bc921ab189a991a5e99c732f2ba7dacfba879250f72624d54a7012e5c0bc75b916e8b7c8c125a200417b5a9379228067ac7ac9b08ab67930b8cd30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000103b450dd03d703259580040d3fce6becb44c5d838bab330352a62f8dd90c81b52407fcfa46a31f71481baa35b15493145c2199ab93f5a40da8947c0646d0bd7c1cb4ccdd5fd5d59a2e41bc50769f26a77046ea33a99a6564ae76568ae25ff8542cea599665544d3f178ca5e892cfc9708b6f221fd7414df58e2671efb2ac9eef19c39ed9c3b93dea2f2d19b3ef6f1334f7674257727c3abea892680a64af4aa20fa0fe2a975e823afaa7cc2d8e9abb866b73c7aca10354d274e74dd88fc1f80e1ccbbb709f838e543c010614108e417a329662ba35a9f106f8a18c5d22e26bd514617eabcc0bdced3b43afa4205c6d728bf662ea0c4f6b7a6d92ecb16b1853cb2ea4fd806110e671daaf40c28912094587bd4af0eca351573047edbae34a34cb30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd440001b58e847d8c0099b7dc3ea25095149b0e133e44fdb46607de475d5e96dbab0c0203f707972ad30bee5d820c6b70aea95952bd85a58f9d65063a9c207e0cd2161ea7d7e5e28dfb431e9ba16dd07d4808cb01aa628b88c057762b871bb2b7f90e7c7ddda95c1666267f90f097319c2e4a0636ba23b786127c43d000aceb508c05394c48c3bdd9cac6eacec742470dc5513cbcbbec5a4067aa5ee2fbcfbe63702032a2e156fe1aa27fb1c32a460e3957282a1bf57bad6af829dccbeaa89d5b341d10241b5451cac1deaebcc2dbb63c8f77f1734ecba34d7c210c0a02bd9f197e2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50588d4a8f4f70d65792d4c99d65f45d15a11bec38362d741cdbe556c5be4af2a2091c49a68b50a88f6d5db294d5651549c4510178cbd5b57ac9172011e218610a26469706673582212206a419b8a253075e88cc443dfd0f4a712ded61c125b83ab6c4445439b8759df1d64736f6c634300081c0033",
}

// FinalizeVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use FinalizeVerifierMetaData.ABI instead.
var FinalizeVerifierABI = FinalizeVerifierMetaData.ABI

// FinalizeVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FinalizeVerifierMetaData.Bin instead.
var FinalizeVerifierBin = FinalizeVerifierMetaData.Bin

// DeployFinalizeVerifier deploys a new Ethereum contract, binding an instance of FinalizeVerifier to it.
func DeployFinalizeVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FinalizeVerifier, error) {
	parsed, err := FinalizeVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FinalizeVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FinalizeVerifier{FinalizeVerifierCaller: FinalizeVerifierCaller{contract: contract}, FinalizeVerifierTransactor: FinalizeVerifierTransactor{contract: contract}, FinalizeVerifierFilterer: FinalizeVerifierFilterer{contract: contract}}, nil
}

// FinalizeVerifier is an auto generated Go binding around an Ethereum contract.
type FinalizeVerifier struct {
	FinalizeVerifierCaller     // Read-only binding to the contract
	FinalizeVerifierTransactor // Write-only binding to the contract
	FinalizeVerifierFilterer   // Log filterer for contract events
}

// FinalizeVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type FinalizeVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizeVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FinalizeVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizeVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FinalizeVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizeVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FinalizeVerifierSession struct {
	Contract     *FinalizeVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FinalizeVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FinalizeVerifierCallerSession struct {
	Contract *FinalizeVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// FinalizeVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FinalizeVerifierTransactorSession struct {
	Contract     *FinalizeVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// FinalizeVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type FinalizeVerifierRaw struct {
	Contract *FinalizeVerifier // Generic contract binding to access the raw methods on
}

// FinalizeVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FinalizeVerifierCallerRaw struct {
	Contract *FinalizeVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// FinalizeVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FinalizeVerifierTransactorRaw struct {
	Contract *FinalizeVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFinalizeVerifier creates a new instance of FinalizeVerifier, bound to a specific deployed contract.
func NewFinalizeVerifier(address common.Address, backend bind.ContractBackend) (*FinalizeVerifier, error) {
	contract, err := bindFinalizeVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FinalizeVerifier{FinalizeVerifierCaller: FinalizeVerifierCaller{contract: contract}, FinalizeVerifierTransactor: FinalizeVerifierTransactor{contract: contract}, FinalizeVerifierFilterer: FinalizeVerifierFilterer{contract: contract}}, nil
}

// NewFinalizeVerifierCaller creates a new read-only instance of FinalizeVerifier, bound to a specific deployed contract.
func NewFinalizeVerifierCaller(address common.Address, caller bind.ContractCaller) (*FinalizeVerifierCaller, error) {
	contract, err := bindFinalizeVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FinalizeVerifierCaller{contract: contract}, nil
}

// NewFinalizeVerifierTransactor creates a new write-only instance of FinalizeVerifier, bound to a specific deployed contract.
func NewFinalizeVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*FinalizeVerifierTransactor, error) {
	contract, err := bindFinalizeVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FinalizeVerifierTransactor{contract: contract}, nil
}

// NewFinalizeVerifierFilterer creates a new log filterer instance of FinalizeVerifier, bound to a specific deployed contract.
func NewFinalizeVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*FinalizeVerifierFilterer, error) {
	contract, err := bindFinalizeVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FinalizeVerifierFilterer{contract: contract}, nil
}

// bindFinalizeVerifier binds a generic wrapper to an already deployed contract.
func bindFinalizeVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FinalizeVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalizeVerifier *FinalizeVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FinalizeVerifier.Contract.FinalizeVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalizeVerifier *FinalizeVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalizeVerifier.Contract.FinalizeVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalizeVerifier *FinalizeVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalizeVerifier.Contract.FinalizeVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalizeVerifier *FinalizeVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FinalizeVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalizeVerifier *FinalizeVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalizeVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalizeVerifier *FinalizeVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalizeVerifier.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_FinalizeVerifier *FinalizeVerifierCaller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int) ([4]*big.Int, error) {
	var out []interface{}
	err := _FinalizeVerifier.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_FinalizeVerifier *FinalizeVerifierSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _FinalizeVerifier.Contract.CompressProof(&_FinalizeVerifier.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_FinalizeVerifier *FinalizeVerifierCallerSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _FinalizeVerifier.Contract.CompressProof(&_FinalizeVerifier.CallOpts, proof)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_FinalizeVerifier *FinalizeVerifierCaller) ProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FinalizeVerifier.contract.Call(opts, &out, "provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_FinalizeVerifier *FinalizeVerifierSession) ProvingKeyHash() ([32]byte, error) {
	return _FinalizeVerifier.Contract.ProvingKeyHash(&_FinalizeVerifier.CallOpts)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_FinalizeVerifier *FinalizeVerifierCallerSession) ProvingKeyHash() ([32]byte, error) {
	return _FinalizeVerifier.Contract.ProvingKeyHash(&_FinalizeVerifier.CallOpts)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5f89feef.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[9] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [9]*big.Int) error {
	var out []interface{}
	err := _FinalizeVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5f89feef.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[9] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [9]*big.Int) error {
	return _FinalizeVerifier.Contract.VerifyCompressedProof(&_FinalizeVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5f89feef.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[9] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [9]*big.Int) error {
	return _FinalizeVerifier.Contract.VerifyCompressedProof(&_FinalizeVerifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x8a3ae438.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[9] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, input [9]*big.Int) error {
	var out []interface{}
	err := _FinalizeVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x8a3ae438.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[9] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierSession) VerifyProof(proof [8]*big.Int, input [9]*big.Int) error {
	return _FinalizeVerifier.Contract.VerifyProof(&_FinalizeVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x8a3ae438.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[9] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCallerSession) VerifyProof(proof [8]*big.Int, input [9]*big.Int) error {
	return _FinalizeVerifier.Contract.VerifyProof(&_FinalizeVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCaller) VerifyProof0(opts *bind.CallOpts, proof []byte, input []byte) error {
	var out []interface{}
	err := _FinalizeVerifier.contract.Call(opts, &out, "verifyProof0", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_FinalizeVerifier *FinalizeVerifierSession) VerifyProof0(proof []byte, input []byte) error {
	return _FinalizeVerifier.Contract.VerifyProof0(&_FinalizeVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCallerSession) VerifyProof0(proof []byte, input []byte) error {
	return _FinalizeVerifier.Contract.VerifyProof0(&_FinalizeVerifier.CallOpts, proof, input)
}
