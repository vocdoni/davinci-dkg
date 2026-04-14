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
	Bin: "0x6080806040523460155761134f908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace1114610c19575080632a07d99a146107a157806344f6369214610708578063a6708604146101f65763b8e72af614610055575f80fd5b346101be5760403660031901126101be5760043567ffffffffffffffff81116101be57610086903690600401610c51565b60243567ffffffffffffffff81116101be576100a6903690600401610c51565b90918301610100848203126101be5780601f850112156101be57604051936100d061010086610c7f565b849061010081019283116101be57905b8282106101e657505050810160a0828203126101be5780601f830112156101be576040519161011060a084610c7f565b829060a081019283116101be57905b8282106101c257505050303b156101be57604051631503eccd60e11b8152915f600484015b600882106101a85750505061010482015f905b60058210610192575050505f816101a481305afa801561018757610179575080f35b61018591505f90610c7f565b005b6040513d5f823e3d90fd5b6020806001928551815201930191019091610157565b6020806001928551815201930191019091610144565b5f80fd5b813581526020918201910161011f565b634e487b7160e01b5f52604160045260245ffd5b81358152602091820191016100e0565b346101be576101203660031901126101be57366084116101be5736610124116101be576103006040516102298282610c7f565b81368237610238600435610f76565b610249602493929335604435610fe1565b91939290610258606435610f76565b9390926040519660408801967f1bea450a92e9e831305c1a7e7a5e6d455d370a7624868a8f0c9752f13f54e5af89528860208101987f1788d2b3677ec6c158a366455a18ec1a45e6f9a7d5412982e62f9279296e60728a525f81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401925f84527f104522c4f4a7a61450d066811a77f79b3b9987155cd3d79f1da9e73f283dd9386084359583608082019780895286828660608160075afa911016818360808160065afa165f85525f8852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f1d0f32cc7d67b01e21323d27c4dd1114e914ef4724cb9d7e45e857f7e4d4a30985527f0c4bc3ef291a22fed0e1f53562e50ed9670f16f5b6c6abaca97aeaca92d9a701885260c43590818a5287838760608160075afa92101616818360808160065afa167f2c37eb4fe25c6552abfc49a6b5cc4b3ac8f780b1e763fefb8a0b0f19a367f60285527f13b4312d8b59a06dfe3b86ca2b800db766e284052ffcc04792fcc34cc8653a89885260e43590818a5287838760608160075afa921016169160808160065afa16947f0f2fec9a5d8b31e58544247e5724c3685078acdb19c615f2da950b69016f5de08352526101043580955260608160075afa9210161660408a60808160065afa169851975198156106f95760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f0aefd3992441537971500a76dde5df99e49c9d30cd68eeecb6a34156ca7b60266101008401527f07f5089d876bcbe36796e3780f6c6859473ddc226b8a337c72ec17522441e34d6101208401527f0ed9fcfb120d476efeb965e42eb2b356bbc3b825b758fa7703ab864b4af200046101408401527f273b7b0e0a96c6d4971dcc2f7de5e2b34bc6d0ae843a75fb147573856bafcdff6101608401527f15629ef1b4ab41d023e71c7c748e03a509845181816ee6399993820054a288256101808401527f10497953ac0c16c08f2d18eab58821398581528b6d2c67b5be1584fb740a02756101a08401527f246f4faaa0cc8908f402bb325f23b1b8c3bfa3792976ede6fbac9cfb528ac31e6101c08401527f119c92e65b64c4978e0f57b63f10096a0a8b798d7334486888971f8962882b416101e08401527f113ef95606e059738a21c0b33fad0a65c6d644deec57acad95d3267c7afbecda6102008401527f2586ba06795af5b5a4439feb5580e846ae4ef9024e329a2aff63ceabb82f3f4d6102208401526102408301526102608201527f1b901da5ee2fb64acb23d0496161423c56c44b4ad5a9de95fab0e351b156b1ef6102808201527f2e57f7765e1c11067bcf269771432b6f3bcf7dd95f5a6df0773ca29014325bfc6102a08201527f1c620eb18ba15cc786ee2cc9dc76a4c7c14b3f7378a969bcb0dd3a1d019b90ff6102c08201527f1056c792afa78ed4a68e0730235af9dfa43acd738286ffaddf4e64595f04f0146102e08201526040519283916106c58484610c7f565b8336843760085afa159081156106ec575b506106dd57005b631ff3747d60e21b5f5260045ffd5b60019150511415816106d6565b63a54f8e2760e01b5f5260045ffd5b346101be576101003660031901126101be5736610104116101be57604051610731608082610c7f565b6080368237610744602435600435610ca1565b815261075a60843560a435604435606435610d42565b6020830152604082015261077260e43560c435610ca1565b6060820152604051905f825b6004821061078b57608084f35b602080600192855181520193019101909161077e565b346101be576101a03660031901126101be5736610104116101be57366101a4116101be5760405160408101907f1bea450a92e9e831305c1a7e7a5e6d455d370a7624868a8f0c9752f13f54e5af815260208101917f1788d2b3677ec6c158a366455a18ec1a45e6f9a7d5412982e62f9279296e607283525f8152606082015f81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f104522c4f4a7a61450d066811a77f79b3b9987155cd3d79f1da9e73f283dd938608087019580875284848460608160075afa911016838860808160065afa165f83525f86526001610124359182895286868660608160075afa9310161616838860808160065afa167f1d0f32cc7d67b01e21323d27c4dd1114e914ef4724cb9d7e45e857f7e4d4a30983527f0c4bc3ef291a22fed0e1f53562e50ed9670f16f5b6c6abaca97aeaca92d9a7018652610144359081885285858560608160075afa92101616838860808160065afa167f2c37eb4fe25c6552abfc49a6b5cc4b3ac8f780b1e763fefb8a0b0f19a367f60283527f13b4312d8b59a06dfe3b86ca2b800db766e284052ffcc04792fcc34cc8653a898652610164359081885285858560608160075afa92101616838860808160065afa16947f0f2fec9a5d8b31e58544247e5724c3685078acdb19c615f2da950b69016f5de08352526101843580955260608160075afa9210161660408260808160065afa169051915190156106f95760405191610100600484377f0aefd3992441537971500a76dde5df99e49c9d30cd68eeecb6a34156ca7b60266101008401527f07f5089d876bcbe36796e3780f6c6859473ddc226b8a337c72ec17522441e34d6101208401527f0ed9fcfb120d476efeb965e42eb2b356bbc3b825b758fa7703ab864b4af200046101408401527f273b7b0e0a96c6d4971dcc2f7de5e2b34bc6d0ae843a75fb147573856bafcdff6101608401527f15629ef1b4ab41d023e71c7c748e03a509845181816ee6399993820054a288256101808401527f10497953ac0c16c08f2d18eab58821398581528b6d2c67b5be1584fb740a02756101a08401527f246f4faaa0cc8908f402bb325f23b1b8c3bfa3792976ede6fbac9cfb528ac31e6101c08401527f119c92e65b64c4978e0f57b63f10096a0a8b798d7334486888971f8962882b416101e08401527f113ef95606e059738a21c0b33fad0a65c6d644deec57acad95d3267c7afbecda6102008401527f2586ba06795af5b5a4439feb5580e846ae4ef9024e329a2aff63ceabb82f3f4d6102208401526102408301526102608201527f1b901da5ee2fb64acb23d0496161423c56c44b4ad5a9de95fab0e351b156b1ef6102808201527f2e57f7765e1c11067bcf269771432b6f3bcf7dd95f5a6df0773ca29014325bfc6102a08201527f1c620eb18ba15cc786ee2cc9dc76a4c7c14b3f7378a969bcb0dd3a1d019b90ff6102c08201527f1056c792afa78ed4a68e0730235af9dfa43acd738286ffaddf4e64595f04f0146102e08201526020816103008160085afa905116156106dd57005b346101be575f3660031901126101be57807f413a4850abadf1894fc17fc886857c9f6df7395e37db93c75310b7cdc386018660209252f35b9181601f840112156101be5782359167ffffffffffffffff83116101be57602083818601950101116101be57565b90601f8019910116810190811067ffffffffffffffff8211176101d257604052565b905f5160206112fa5f395f51905f528210801590610d2b575b6106dd57811580610d23575b610d1d57610cea5f5160206112fa5f395f51905f526003818581818009090861111a565b818103610cf957505060011b90565b5f5160206112fa5f395f51905f52809106810306145f146106dd57600190811b1790565b50505f90565b508015610cc6565b505f5160206112fa5f395f51905f52811015610cba565b919093925f5160206112fa5f395f51905f528310801590610f5f575b8015610f48575b8015610f31575b6106dd578082868517171715610f2657908291610e895f5160206112fa5f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f5160206112fa5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481610e6381808b8009818780090861111a565b8408095f5160206112fa5f395f51905f52610e7d82611291565b8009141595869161113d565b929080821480610f1d575b15610ebb5750505050905f14610eb35760ff60025b169060021b179190565b60ff5f610ea9565b5f5160206112fa5f395f51905f52809106810306149182610efe575b5050156106dd5760019115610ef65760ff60025b169060021b17179190565b60ff5f610eeb565b5f5160206112fa5f395f51905f52919250819006810306145f80610ed7565b50838314610e94565b50505090505f905f90565b505f5160206112fa5f395f51905f52811015610d6c565b505f5160206112fa5f395f51905f52821015610d65565b505f5160206112fa5f395f51905f52851015610d5e565b8015610fda578060011c915f5160206112fa5f395f51905f528310156106dd57600180610fb95f5160206112fa5f395f51905f526003818881818009090861111a565b931614610fc257565b905f5160206112fa5f395f51905f5280910681030690565b505f905f90565b801580611112575b611106578060021c92825f5160206112fa5f395f51905f5285108015906110ef575b6106dd5784815f5160206112fa5f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816110b99d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50861113d565b809291600180829616146110cb575050565b5f5160206112fa5f395f51905f528093945080929550809106810306930681030690565b505f5160206112fa5f395f51905f5281101561100b565b50505f905f905f905f90565b508115610fe9565b9061112482611291565b915f5160206112fa5f395f51905f52838009036106dd57565b915f5160206112fa5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816111959396949661118782808a8009818a80090861111a565b90611285575b86080961111a565b925f5160206112fa5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206112fa5f395f51905f5260a083015260208260c08160055afa915191156106dd575f5160206112fa5f395f51905f528260019209036106dd575f5160206112fa5f395f51905f52908209925f5160206112fa5f395f51905f528080808780090681030681878009081490811591611266575b506106dd57565b90505f5160206112fa5f395f51905f528084860960020914155f61125f565b8180910681030661118d565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206112fa5f395f51905f5260a083015260208260c08160055afa915191156106dd5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a26469706673582212201fc9c391199381f7ea47f1faae3b237fd4c54ddf5709e9d945532b0ab35ac6b064736f6c634300081c0033",
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
