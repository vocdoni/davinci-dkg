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

// RevealSubmitVerifierMetaData contains all meta data concerning the RevealSubmitVerifier contract.
var RevealSubmitVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"input\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x6080806040523460155761134d908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace1114610c17575080632a07d99a146107a057806344f6369214610707578063a6708604146101f65763b8e72af614610055575f80fd5b346101be5760403660031901126101be5760043567ffffffffffffffff81116101be57610086903690600401610c4f565b60243567ffffffffffffffff81116101be576100a6903690600401610c4f565b90918301610100848203126101be5780601f850112156101be57604051936100d061010086610c7d565b849061010081019283116101be57905b8282106101e657505050810160a0828203126101be5780601f830112156101be576040519161011060a084610c7d565b829060a081019283116101be57905b8282106101c257505050303b156101be57604051631503eccd60e11b8152915f600484015b600882106101a85750505061010482015f905b60058210610192575050505f816101a481305afa801561018757610179575080f35b61018591505f90610c7d565b005b6040513d5f823e3d90fd5b6020806001928551815201930191019091610157565b6020806001928551815201930191019091610144565b5f80fd5b813581526020918201910161011f565b634e487b7160e01b5f52604160045260245ffd5b81358152602091820191016100e0565b346101be576101203660031901126101be57366084116101be5736610124116101be576103006040516102298282610c7d565b81368237610238600435610f74565b610249602493929335604435610fdf565b91939290610258606435610f74565b9390926040519660408801967e59e0867e4f7640daa9c99ec6ce5c79781a41c3e6f2c29b10444add1c2c449289528860208101987f232d9ceec10f9a1fb72a5485a3d104269d3db6444c7d571f442dab86e69a9b208a525f81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401925f84527f216cd25f7339674d9b2263327897e2c30fa7fdcfe1657d2610f0c240d5145c9d6084359583608082019780895286828660608160075afa911016818360808160065afa165f85525f8852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f07bd6188f3e06f35b2ffc3afd3d9bc7eeb89e3ddfe9b8a7462a48c087056035d85527f0441337153b256144d2116dbcdac3d66db8b22c58a63491d8b534c5ecc7262eb885260c43590818a5287838760608160075afa92101616818360808160065afa167f252b4b44f1ff28415b674836a1965a09215bd06fe8c9ff92b14ff6bd41bae07b85527f0e8534691e6787bae3f65336c6950165f8f55a4b0bd76b4442b1535c34185810885260e43590818a5287838760608160075afa921016169160808160065afa16947f204b3b6108faed2383700d77832a7dcdc39b0626a083bb4754a0a44211668e898352526101043580955260608160075afa9210161660408a60808160065afa169851975198156106f85760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f2a2d3173efb6376bccb67c049a285129b7b1baf3366a23e467abea48f7ee8ae46101008401527f13588c3e22e775c754371bb1330d408d363c691864996f7d04163dd7d9188e946101208401527f06e4ae1006150950cda0dc0a00fd1755809b5da2c0ba034fdefd05d901ccb6586101408401527f1cb91807286cf767c50a1301e662b60b1dc57d586fe5ecc0061b81f0971de0cd6101608401527f07b0c39e9a071e71aa92de63b066addc79bb440b483f4138e22c13fcbd493fda6101808401527f10d594dbe423530c8e8b639151eae8acc26e402d7ae616a3ece83414a6eb168d6101a08401527f1cf6b23605a12a3fab920d9fd8d10b33b6e6fe4e9c10d931a89cbc9fcf6d2d306101c08401527f05a6b4b0a8f35d8d7bf0a9c7a98f7e53058321bc2654761a3262156f215d5a966101e08401527f1368b3ba41632a59d72da741fedd5e6747815225e52a572e5f053cad3b76fa266102008401527f053df329321a7672ec459579e3dbd17a7657840b2a5550bdd9e0bb150b2f38e66102208401526102408301526102608201527f213870d101427a83487a0ba0bc6699b3cca36c483a4048aebb7fca56f838020c6102808201527f23c22621a35996595e5d44313aa59b73f564e0ad0ca4dc12dd091a68de5f67c36102a08201527f1cd31cdfdba27caea015a1a4f94d2b7a30e4056bbe7781dea61e5027a557ec276102c08201527f25f6ff6e7e89ab89e18902ae13e8a52947cac89c95d818b937d3fe989b6c26e46102e08201526040519283916106c48484610c7d565b8336843760085afa159081156106eb575b506106dc57005b631ff3747d60e21b5f5260045ffd5b60019150511415816106d5565b63a54f8e2760e01b5f5260045ffd5b346101be576101003660031901126101be5736610104116101be57604051610730608082610c7d565b6080368237610743602435600435610c9f565b815261075960843560a435604435606435610d40565b6020830152604082015261077160e43560c435610c9f565b6060820152604051905f825b6004821061078a57608084f35b602080600192855181520193019101909161077d565b346101be576101a03660031901126101be5736610104116101be57366101a4116101be5760405160408101907e59e0867e4f7640daa9c99ec6ce5c79781a41c3e6f2c29b10444add1c2c4492815260208101917f232d9ceec10f9a1fb72a5485a3d104269d3db6444c7d571f442dab86e69a9b2083525f8152606082015f81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f216cd25f7339674d9b2263327897e2c30fa7fdcfe1657d2610f0c240d5145c9d608087019580875284848460608160075afa911016838860808160065afa165f83525f86526001610124359182895286868660608160075afa9310161616838860808160065afa167f07bd6188f3e06f35b2ffc3afd3d9bc7eeb89e3ddfe9b8a7462a48c087056035d83527f0441337153b256144d2116dbcdac3d66db8b22c58a63491d8b534c5ecc7262eb8652610144359081885285858560608160075afa92101616838860808160065afa167f252b4b44f1ff28415b674836a1965a09215bd06fe8c9ff92b14ff6bd41bae07b83527f0e8534691e6787bae3f65336c6950165f8f55a4b0bd76b4442b1535c341858108652610164359081885285858560608160075afa92101616838860808160065afa16947f204b3b6108faed2383700d77832a7dcdc39b0626a083bb4754a0a44211668e898352526101843580955260608160075afa9210161660408260808160065afa169051915190156106f85760405191610100600484377f2a2d3173efb6376bccb67c049a285129b7b1baf3366a23e467abea48f7ee8ae46101008401527f13588c3e22e775c754371bb1330d408d363c691864996f7d04163dd7d9188e946101208401527f06e4ae1006150950cda0dc0a00fd1755809b5da2c0ba034fdefd05d901ccb6586101408401527f1cb91807286cf767c50a1301e662b60b1dc57d586fe5ecc0061b81f0971de0cd6101608401527f07b0c39e9a071e71aa92de63b066addc79bb440b483f4138e22c13fcbd493fda6101808401527f10d594dbe423530c8e8b639151eae8acc26e402d7ae616a3ece83414a6eb168d6101a08401527f1cf6b23605a12a3fab920d9fd8d10b33b6e6fe4e9c10d931a89cbc9fcf6d2d306101c08401527f05a6b4b0a8f35d8d7bf0a9c7a98f7e53058321bc2654761a3262156f215d5a966101e08401527f1368b3ba41632a59d72da741fedd5e6747815225e52a572e5f053cad3b76fa266102008401527f053df329321a7672ec459579e3dbd17a7657840b2a5550bdd9e0bb150b2f38e66102208401526102408301526102608201527f213870d101427a83487a0ba0bc6699b3cca36c483a4048aebb7fca56f838020c6102808201527f23c22621a35996595e5d44313aa59b73f564e0ad0ca4dc12dd091a68de5f67c36102a08201527f1cd31cdfdba27caea015a1a4f94d2b7a30e4056bbe7781dea61e5027a557ec276102c08201527f25f6ff6e7e89ab89e18902ae13e8a52947cac89c95d818b937d3fe989b6c26e46102e08201526020816103008160085afa905116156106dc57005b346101be575f3660031901126101be57807f413a4850abadf1894fc17fc886857c9f6df7395e37db93c75310b7cdc386018660209252f35b9181601f840112156101be5782359167ffffffffffffffff83116101be57602083818601950101116101be57565b90601f8019910116810190811067ffffffffffffffff8211176101d257604052565b905f5160206112f85f395f51905f528210801590610d29575b6106dc57811580610d21575b610d1b57610ce85f5160206112f85f395f51905f5260038185818180090908611118565b818103610cf757505060011b90565b5f5160206112f85f395f51905f52809106810306145f146106dc57600190811b1790565b50505f90565b508015610cc4565b505f5160206112f85f395f51905f52811015610cb8565b919093925f5160206112f85f395f51905f528310801590610f5d575b8015610f46575b8015610f2f575b6106dc578082868517171715610f2457908291610e875f5160206112f85f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f5160206112f85f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481610e6181808b80098187800908611118565b8408095f5160206112f85f395f51905f52610e7b8261128f565b8009141595869161113b565b929080821480610f1b575b15610eb95750505050905f14610eb15760ff60025b169060021b179190565b60ff5f610ea7565b5f5160206112f85f395f51905f52809106810306149182610efc575b5050156106dc5760019115610ef45760ff60025b169060021b17179190565b60ff5f610ee9565b5f5160206112f85f395f51905f52919250819006810306145f80610ed5565b50838314610e92565b50505090505f905f90565b505f5160206112f85f395f51905f52811015610d6a565b505f5160206112f85f395f51905f52821015610d63565b505f5160206112f85f395f51905f52851015610d5c565b8015610fd8578060011c915f5160206112f85f395f51905f528310156106dc57600180610fb75f5160206112f85f395f51905f5260038188818180090908611118565b931614610fc057565b905f5160206112f85f395f51905f5280910681030690565b505f905f90565b801580611110575b611104578060021c92825f5160206112f85f395f51905f5285108015906110ed575b6106dc5784815f5160206112f85f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816110b79d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50861113b565b809291600180829616146110c9575050565b5f5160206112f85f395f51905f528093945080929550809106810306930681030690565b505f5160206112f85f395f51905f52811015611009565b50505f905f905f905f90565b508115610fe7565b906111228261128f565b915f5160206112f85f395f51905f52838009036106dc57565b915f5160206112f85f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816111939396949661118582808a8009818a800908611118565b90611283575b860809611118565b925f5160206112f85f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206112f85f395f51905f5260a083015260208260c08160055afa915191156106dc575f5160206112f85f395f51905f528260019209036106dc575f5160206112f85f395f51905f52908209925f5160206112f85f395f51905f528080808780090681030681878009081490811591611264575b506106dc57565b90505f5160206112f85f395f51905f528084860960020914155f61125d565b8180910681030661118b565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206112f85f395f51905f5260a083015260208260c08160055afa915191156106dc5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220c9d9f078fc6ec483caef5543e6ad9217c012411a7ab1ed146152b9f2d519e17f64736f6c634300081c0033",
}

// RevealSubmitVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use RevealSubmitVerifierMetaData.ABI instead.
var RevealSubmitVerifierABI = RevealSubmitVerifierMetaData.ABI

// RevealSubmitVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RevealSubmitVerifierMetaData.Bin instead.
var RevealSubmitVerifierBin = RevealSubmitVerifierMetaData.Bin

// DeployRevealSubmitVerifier deploys a new Ethereum contract, binding an instance of RevealSubmitVerifier to it.
func DeployRevealSubmitVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RevealSubmitVerifier, error) {
	parsed, err := RevealSubmitVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RevealSubmitVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RevealSubmitVerifier{RevealSubmitVerifierCaller: RevealSubmitVerifierCaller{contract: contract}, RevealSubmitVerifierTransactor: RevealSubmitVerifierTransactor{contract: contract}, RevealSubmitVerifierFilterer: RevealSubmitVerifierFilterer{contract: contract}}, nil
}

// RevealSubmitVerifier is an auto generated Go binding around an Ethereum contract.
type RevealSubmitVerifier struct {
	RevealSubmitVerifierCaller     // Read-only binding to the contract
	RevealSubmitVerifierTransactor // Write-only binding to the contract
	RevealSubmitVerifierFilterer   // Log filterer for contract events
}

// RevealSubmitVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type RevealSubmitVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevealSubmitVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RevealSubmitVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevealSubmitVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RevealSubmitVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevealSubmitVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RevealSubmitVerifierSession struct {
	Contract     *RevealSubmitVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RevealSubmitVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RevealSubmitVerifierCallerSession struct {
	Contract *RevealSubmitVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// RevealSubmitVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RevealSubmitVerifierTransactorSession struct {
	Contract     *RevealSubmitVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// RevealSubmitVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type RevealSubmitVerifierRaw struct {
	Contract *RevealSubmitVerifier // Generic contract binding to access the raw methods on
}

// RevealSubmitVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RevealSubmitVerifierCallerRaw struct {
	Contract *RevealSubmitVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// RevealSubmitVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RevealSubmitVerifierTransactorRaw struct {
	Contract *RevealSubmitVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRevealSubmitVerifier creates a new instance of RevealSubmitVerifier, bound to a specific deployed contract.
func NewRevealSubmitVerifier(address common.Address, backend bind.ContractBackend) (*RevealSubmitVerifier, error) {
	contract, err := bindRevealSubmitVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RevealSubmitVerifier{RevealSubmitVerifierCaller: RevealSubmitVerifierCaller{contract: contract}, RevealSubmitVerifierTransactor: RevealSubmitVerifierTransactor{contract: contract}, RevealSubmitVerifierFilterer: RevealSubmitVerifierFilterer{contract: contract}}, nil
}

// NewRevealSubmitVerifierCaller creates a new read-only instance of RevealSubmitVerifier, bound to a specific deployed contract.
func NewRevealSubmitVerifierCaller(address common.Address, caller bind.ContractCaller) (*RevealSubmitVerifierCaller, error) {
	contract, err := bindRevealSubmitVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RevealSubmitVerifierCaller{contract: contract}, nil
}

// NewRevealSubmitVerifierTransactor creates a new write-only instance of RevealSubmitVerifier, bound to a specific deployed contract.
func NewRevealSubmitVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*RevealSubmitVerifierTransactor, error) {
	contract, err := bindRevealSubmitVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RevealSubmitVerifierTransactor{contract: contract}, nil
}

// NewRevealSubmitVerifierFilterer creates a new log filterer instance of RevealSubmitVerifier, bound to a specific deployed contract.
func NewRevealSubmitVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*RevealSubmitVerifierFilterer, error) {
	contract, err := bindRevealSubmitVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RevealSubmitVerifierFilterer{contract: contract}, nil
}

// bindRevealSubmitVerifier binds a generic wrapper to an already deployed contract.
func bindRevealSubmitVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RevealSubmitVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevealSubmitVerifier *RevealSubmitVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevealSubmitVerifier.Contract.RevealSubmitVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevealSubmitVerifier *RevealSubmitVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevealSubmitVerifier.Contract.RevealSubmitVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevealSubmitVerifier *RevealSubmitVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevealSubmitVerifier.Contract.RevealSubmitVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevealSubmitVerifier *RevealSubmitVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevealSubmitVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevealSubmitVerifier *RevealSubmitVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevealSubmitVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevealSubmitVerifier *RevealSubmitVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevealSubmitVerifier.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_RevealSubmitVerifier *RevealSubmitVerifierCaller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int) ([4]*big.Int, error) {
	var out []interface{}
	err := _RevealSubmitVerifier.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_RevealSubmitVerifier *RevealSubmitVerifierSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _RevealSubmitVerifier.Contract.CompressProof(&_RevealSubmitVerifier.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_RevealSubmitVerifier *RevealSubmitVerifierCallerSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _RevealSubmitVerifier.Contract.CompressProof(&_RevealSubmitVerifier.CallOpts, proof)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_RevealSubmitVerifier *RevealSubmitVerifierCaller) ProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RevealSubmitVerifier.contract.Call(opts, &out, "provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_RevealSubmitVerifier *RevealSubmitVerifierSession) ProvingKeyHash() ([32]byte, error) {
	return _RevealSubmitVerifier.Contract.ProvingKeyHash(&_RevealSubmitVerifier.CallOpts)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_RevealSubmitVerifier *RevealSubmitVerifierCallerSession) ProvingKeyHash() ([32]byte, error) {
	return _RevealSubmitVerifier.Contract.ProvingKeyHash(&_RevealSubmitVerifier.CallOpts)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xa6708604.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[5] input) view returns()
func (_RevealSubmitVerifier *RevealSubmitVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [5]*big.Int) error {
	var out []interface{}
	err := _RevealSubmitVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xa6708604.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[5] input) view returns()
func (_RevealSubmitVerifier *RevealSubmitVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [5]*big.Int) error {
	return _RevealSubmitVerifier.Contract.VerifyCompressedProof(&_RevealSubmitVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xa6708604.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[5] input) view returns()
func (_RevealSubmitVerifier *RevealSubmitVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [5]*big.Int) error {
	return _RevealSubmitVerifier.Contract.VerifyCompressedProof(&_RevealSubmitVerifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x2a07d99a.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[5] input) view returns()
func (_RevealSubmitVerifier *RevealSubmitVerifierCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, input [5]*big.Int) error {
	var out []interface{}
	err := _RevealSubmitVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x2a07d99a.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[5] input) view returns()
func (_RevealSubmitVerifier *RevealSubmitVerifierSession) VerifyProof(proof [8]*big.Int, input [5]*big.Int) error {
	return _RevealSubmitVerifier.Contract.VerifyProof(&_RevealSubmitVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x2a07d99a.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[5] input) view returns()
func (_RevealSubmitVerifier *RevealSubmitVerifierCallerSession) VerifyProof(proof [8]*big.Int, input [5]*big.Int) error {
	return _RevealSubmitVerifier.Contract.VerifyProof(&_RevealSubmitVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_RevealSubmitVerifier *RevealSubmitVerifierCaller) VerifyProof0(opts *bind.CallOpts, proof []byte, input []byte) error {
	var out []interface{}
	err := _RevealSubmitVerifier.contract.Call(opts, &out, "verifyProof0", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_RevealSubmitVerifier *RevealSubmitVerifierSession) VerifyProof0(proof []byte, input []byte) error {
	return _RevealSubmitVerifier.Contract.VerifyProof0(&_RevealSubmitVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_RevealSubmitVerifier *RevealSubmitVerifierCallerSession) VerifyProof0(proof []byte, input []byte) error {
	return _RevealSubmitVerifier.Contract.VerifyProof0(&_RevealSubmitVerifier.CallOpts, proof, input)
}
