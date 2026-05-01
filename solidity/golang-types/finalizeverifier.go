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
	Bin: "0x60808060405234601557611696908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace1114610b185750806344f6369214610a7f5780635f89feef1461060c5780638a3ae438146101f75763b8e72af614610055575f80fd5b346101bf5760403660031901126101bf576004356001600160401b0381116101bf57610085903690600401610b50565b6024356001600160401b0381116101bf576100a4903690600401610b50565b90918301610100848203126101bf5780601f850112156101bf57604051936100ce61010086610b7d565b849061010081019283116101bf57905b8282106101e7575050508101610120828203126101bf5780601f830112156101bf576040519161011061012084610b7d565b829061012081019283116101bf57905b8282106101c357505050303b156101bf576040516311475c8760e31b8152915f600484015b600882106101a95750505061010482015f905b60098210610193575050505f8161022481305afa80156101885761017a575080f35b61018691505f90610b7d565b005b6040513d5f823e3d90fd5b6020806001928551815201930191019091610158565b6020806001928551815201930191019091610145565b5f80fd5b8135815260209182019101610120565b634e487b7160e01b5f52604160045260245ffd5b81358152602091820191016100de565b346101bf576102203660031901126101bf5736610104116101bf5736610224116101bf5760405160408101905f5160206115a15f395f51905f52815260208101915f5160206114215f395f51905f5283525f5160206114415f395f51905f528152606082015f5160206113a15f395f51905f5281525f5160206113815f395f51905f52604061010435935f5160206114815f395f51905f52608087019580875284848460608160075afa911016838860808160065afa165f5160206113c15f395f51905f5283525f5160206111a15f395f51905f5286526001610124359182895286868660608160075afa9310161616838860808160065afa165f5160206114615f395f51905f5283525f5160206112c15f395f51905f528652610144359081885285858560608160075afa92101616838860808160065afa165f5160206112015f395f51905f5283525f5160206113615f395f51905f528652610164359081885285858560608160075afa92101616838860808160065afa165f5160206115415f395f51905f5283525f5160206116215f395f51905f528652610184359081885285858560608160075afa92101616838860808160065afa165f5160206115c15f395f51905f5283525f5160206112a15f395f51905f5286526101a4359081885285858560608160075afa92101616838860808160065afa165f5160206113415f395f51905f5283525f5160206112e15f395f51905f5286526101c4359081885285858560608160075afa92101616838860808160065afa165f5160206115e15f395f51905f5283525f5160206111e15f395f51905f5286526101e4359081885285858560608160075afa92101616838860808160065afa16945f5160206115015f395f51905f528352526102043580955260608160075afa9210161660408260808160065afa169051915190156105fd5760405191610100600484375f5160206112615f395f51905f526101008401525f5160206114015f395f51905f526101208401525f5160206112215f395f51905f526101408401525f5160206111815f395f51905f526101608401525f5160206112415f395f51905f526101808401525f5160206115615f395f51905f526101a08401525f5160206113215f395f51905f526101c08401525f5160206114a15f395f51905f526101e08401525f5160206113e15f395f51905f526102008401525f5160206116015f395f51905f526102208401526102408301526102608201525f5160206114e15f395f51905f526102808201525f5160206115215f395f51905f526102a08201525f5160206115815f395f51905f526102c08201525f5160206111c15f395f51905f526102e08201526020816103008160085afa905116156105ee57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b346101bf576101a03660031901126101bf57366084116101bf57366101a4116101bf5761030060405161063f8282610b7d565b8136823761064e600435610e29565b61065f602493929335604435610e94565b9193929061066e606435610e29565b9390926040519660408801965f5160206115a15f395f51905f5289528860208101985f5160206114215f395f51905f528a525f5160206114415f395f51905f5281525f5160206113815f395f51905f52604060608401925f5160206113a15f395f51905f5284525f5160206114815f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f5160206113c15f395f51905f5285525f5160206111a15f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f5160206114615f395f51905f5285525f5160206112c15f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f5160206112015f395f51905f5285525f5160206113615f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f5160206115415f395f51905f5285525f5160206116215f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f5160206115c15f395f51905f5285525f5160206112a15f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f5160206113415f395f51905f5285525f5160206112e15f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f5160206115e15f395f51905f5285525f5160206111e15f395f51905f5288526101643590818a5287838760608160075afa921016169160808160065afa16945f5160206115015f395f51905f528352526101843580955260608160075afa9210161660408a60808160065afa169851975198156105fd5760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f5160206112615f395f51905f526101008401525f5160206114015f395f51905f526101208401525f5160206112215f395f51905f526101408401525f5160206111815f395f51905f526101608401525f5160206112415f395f51905f526101808401525f5160206115615f395f51905f526101a08401525f5160206113215f395f51905f526101c08401525f5160206114a15f395f51905f526101e08401525f5160206113e15f395f51905f526102008401525f5160206116015f395f51905f526102208401526102408301526102608201525f5160206114e15f395f51905f526102808201525f5160206115215f395f51905f526102a08201525f5160206115815f395f51905f526102c08201525f5160206111c15f395f51905f526102e0820152604051928391610a5a8484610b7d565b8336843760085afa15908115610a72575b506105ee57005b6001915051141581610a6b565b346101bf576101003660031901126101bf5736610104116101bf57604051610aa8608082610b7d565b6080368237610abb602435600435610ba0565b8152610ad160843560a435604435606435610c41565b60208301526040820152610ae960e43560c435610ba0565b6060820152604051905f825b60048210610b0257608084f35b6020806001928551815201930191019091610af5565b346101bf575f3660031901126101bf57807fb7d50091bbf5d4629605846b345d0113903db148242238c55d039eb0c88e51d960209252f35b9181601f840112156101bf578235916001600160401b0383116101bf57602083818601950101116101bf57565b601f909101601f19168101906001600160401b038211908210176101d357604052565b905f5160206112815f395f51905f528210801590610c2a575b6105ee57811580610c22575b610c1c57610be95f5160206112815f395f51905f5260038185818180090908610f94565b818103610bf857505060011b90565b5f5160206112815f395f51905f52809106810306145f146105ee57600190811b1790565b50505f90565b508015610bc5565b505f5160206112815f395f51905f52811015610bb9565b919093925f5160206112815f395f51905f528310801590610e12575b8015610dfb575b8015610de4575b6105ee578082868517171715610dd957908291610d3c5f5160206112815f395f51905f5280808080888180808f9d5f5160206114c15f395f51905f528f839290839109099d8e0981848181800909085f5160206116415f395f51905f52089a09818c8181800909085f5160206111615f395f51905f520806810306945f5160206112815f395f51905f525f5160206113015f395f51905f5281610d1681808b80098187800908610f94565b8408095f5160206112815f395f51905f52610d30826110f8565b80091415958691610fb7565b929080821480610dd0575b15610d6e5750505050905f14610d665760ff60025b169060021b179190565b60ff5f610d5c565b5f5160206112815f395f51905f52809106810306149182610db1575b5050156105ee5760019115610da95760ff60025b169060021b17179190565b60ff5f610d9e565b5f5160206112815f395f51905f52919250819006810306145f80610d8a565b50838314610d47565b50505090505f905f90565b505f5160206112815f395f51905f52811015610c6b565b505f5160206112815f395f51905f52821015610c64565b505f5160206112815f395f51905f52851015610c5d565b8015610e8d578060011c915f5160206112815f395f51905f528310156105ee57600180610e6c5f5160206112815f395f51905f5260038188818180090908610f94565b931614610e7557565b905f5160206112815f395f51905f5280910681030690565b505f905f90565b801580610f8c575b610f80578060021c92825f5160206112815f395f51905f528510801590610f69575b6105ee5784815f5160206112815f395f51905f5280808080808080805f5160206114c15f395f51905f5281610f339d8d0909998a0981898181800909085f5160206111615f395f51905f520806810306936002808a16149509818a8181800909085f5160206116415f395f51905f5208610fb7565b80929160018082961614610f45575050565b5f5160206112815f395f51905f528093945080929550809106810306930681030690565b505f5160206112815f395f51905f52811015610ebe565b50505f905f905f905f90565b508115610e9c565b90610f9e826110f8565b915f5160206112815f395f51905f52838009036105ee57565b915f5160206112815f395f51905f525f5160206113015f395f51905f5281610ffc93969496610fee82808a8009818a800908610f94565b906110ec575b860809610f94565b925f5160206112815f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206112815f395f51905f5260a083015260208260c08160055afa915191156105ee575f5160206112815f395f51905f528260019209036105ee575f5160206112815f395f51905f52908209925f5160206112815f395f51905f5280808087800906810306818780090814908115916110cd575b506105ee57565b90505f5160206112815f395f51905f528084860960020914155f6110c6565b81809106810306610ff4565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206112815f395f51905f5260a083015260208260c08160055afa915191156105ee5756fe2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77503bffe1277ff37cb62e766c76c457cb2f6dca0b0f4e8eee2bb922aae05b6b92c061bddde25f87a59db159703a28d007ff36e444514b39dff51b58585f6891a350dafb26cea969f25ae4e31b413f61a90fd818d8ef6b8658694fb21a60991d803120e5f44d481f1c837e1b282a64c709ad190cc9396d3514666513347ca09c77f1f0373639716fc13b7343276a18471fc035b3f313ebbd637be88b0c6c9a4d73c2533f029c83dab87cbb7cee5af9acb06fd6fec22301147ae4f936c659cb7ba321e9cb26498cbe8c467c5171ca770acc954b1b5ab6a72b15b0a2b9c4d0aadcbd220f798e82e900df47b7fdb7611076f854798ecae4c27cffd7a1662a13c396dc330644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd472aa6e340d50144e60a4e018b040cdf05cca1815b428747c4ff2bb5f28b14606d1f6725d01a8c0433767278672f719068b1a180150402a78fa0f672ff8eca35fc235d6ada313c9282650b2ae29ad8e70f22242e5eadf7b7c4672d20dcb6f125ec183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4294ca3dac3a647cab572450072bd8baf4ef50a6f82df0d09894f76c3bc77cfb30c0eb3d073bd09e324e6996163144c5adfa89b937769fcef4c46f1b1da2577781b423456f9d7825db2bad7e1106587b426b8b58510ba6c1b322235c644b4c7f830644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000126f405b630b40940bbbc2cd31492f7fb92612de37704ed9de3fe49e343cbba581f30624c0ece90319acb451be8273fb53d857db0b0c9e26d96b8010e44d20b0005eebb1b32c08755958bece7acd8c538781f2eb9791c3506c7a8e8331810212d2baaa5b24e4fbc3e71127156e021552b2c04fe05067e2618cb96f8910b8de00016a049ffdeb9479af4be0e6a6fda5c6afc687b6a9cff3b244d4e93d79a93e8a003a678c0b4a2fd37af8af05a1708813af1eed14a797a4d9525ee28e545f6490113fd8905fbdc45583ed2b6a556785ea620ba12bd0295acf9ef615a00147f6ef9178b16e62e063591c43c2956c45dd1cbeb952ae88796cbcb8931b7b5472b93ca1b7f20113c9720f28ca87dbeb58a89410f54590ef74840a6de006e45683a7f2130644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44012a22b442edf98f1a0b2eb4dc642073a6724cd9508986b4c5274e6be538d50c09c0f5b76b2bec313e7554887e75a54f08b497c7a973e7a8154530d129e9998a201882e501e6c53820c15cebbd926934fbe82db67f667f657ecd91af7e60f167184aa5123f617f447476aee47e9a38901a4cac953bffef076c5e5b734b8506562636fc14db9b6122d613d3624a21e4eddf8af5ebca1a0c91eba36f8fefd6b5d622c2b30de542f8bca328e33a0a5df3c1d58d0374d8638405d41ff7cc94dd72320db5cfe1d4bd93960819c2512e06f60be9cc3229ef3051c451c4be15046be4f8169276a15907e5259303232795a045b01919d3d134eeb9b9051b9864f962dc542822fb62d198a8c5d9680118e75822e4be0192fd3fe06dee43cde5f9740d4d5111ce229f8f4cd70122dc8c9afe8f436df2185b5ed7141cd7690767aae44ef06a2c82750a2b3c4b2a8193ced51681da6588af5918039b540c8732b395136e3b3b2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5a2646970667358221220deea918348a55c427445c3c197c0cc9ba17be247b0e127b42ac64bee77387f6464736f6c634300081c0033",
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
