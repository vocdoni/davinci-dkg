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

// RevealShareVerifierMetaData contains all meta data concerning the RevealShareVerifier contract.
var RevealShareVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[7]\",\"internalType\":\"uint256[7]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"input\",\"type\":\"uint256[7]\",\"internalType\":\"uint256[7]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x6080806040523460155761160b908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80631042664e1461092b578063233ace11146108f157806344f6369214610858578063b8e72af6146106b65763f6d3293a14610051575f80fd5b346106b3576101603660031901126106b357366084116106b35736610164116106b357604051906103006100858184610f3b565b80368437610094600435611232565b6100a560249592953560443561129d565b919392906100b4606435611232565b9390926040519660408801967f0dfb26ff37b679c1c9f64380ee92d8056017a1daeceac11f25578ee59bbbb2cb89528860208101987f1f402de6af9d5597d99032e3322a62a744e6c75b75e48a75fc6be86c1bb4abc88a527f01ba57caf2dd0ef97400ff7847161ce4dddacb3cf6df86e975a3092ef4ec4cbd81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f101dfe7b067da65c9cf7bf3ffe6b4da4052196a13a03d98cdd83bb502149dd4084527f285e235e9aa51daf6644c1d8db1bcb7155fbbbe346d8d4954544281444a3e51a6084359583608082019780895286828660608160075afa911016818360808160065afa167f1d10def94709e1da4643a80863150fb716c567064a44e6643e693f5b309ab46a85527f1a4463dc9bc21f16ec0e74f6bb7c0072fff9068e586aa659b7b076d1b2d9b09c8852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f1c8c32efdcd4dfc7fe88794bb241042946ad6944306e0aa2a04b97354d70150885527f10c00683bf679e37dfb6dbf161a29d3bb08589095af373b2bf9726cf67539e9e885260c43590818a5287838760608160075afa92101616818360808160065afa167f1aedfa0cb2fab6fc6e352dc0d1c0661fe23798ad4a249db556c8df6e68c6c42985527f2183614418c533233ac0f8054c19e4ce9de9d69856d3f83f17460f40a1ab47de885260e43590818a5287838760608160075afa92101616818360808160065afa167f1bb7fa5dae2eb590e5d44bec6863582c701098671e9c6a223252aeb0776af67685527f293e29d7dd18d1427304c85230692667fcc4c366273f953c51990c28383e5ac188526101043590818a5287838760608160075afa92101616818360808160065afa167f2c5feb1355cffb9df96d15312b52d60c6cd02a79a2604f8cc157ec8fadb537ee85527f2d42cf984a0ebe1a82ddd33fece9cf15c95510ea90546e905e9a85d662affb5588526101243590818a5287838760608160075afa921016169160808160065afa16947f199e7da818353ecc131688006b67af228d1f4c79dab5015975d3da71a8547d4f8352526101443580955260608160075afa9210161660408a60808160065afa169851975198156106a45760209a9b9c8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f03799f712eb0dbd68cf72393c65a47ac92ef43e1e84ff7d97495a72b9316f4a56101008401527f03c720838a19cb63a9447552c8d6e5553cfc22c58f3ce6634893bf58bebfea7c6101208401527f10d4839fbf942fbc277fc01bb04db9686a13169c569034dc6b9a9801b3b99b9f6101408401527f0bd1622089ed6ab4a26a3e93a93401aa95eaf142a51020d935b31a2b9714d48e6101608401527f0a6ced41020f9fc692f95d1e001e90bd616998e88209d178a5cd89bbc86c0e7b6101808401527f2edd4c02ad8154e40af01b76944e56950cdf1f53123ed6d94877e101fa09142d6101a08401527f18736cad4dee17ff032780bf31d62f073fd7dba456f7c89e1680e83b50aef40e6101c08401527f2498a73464d4d591d93503d794c2b6cb380f1511c9cd4b10ef87132027a1e10e6101e08401527f2b0589c30078a901e83c090cae1029d5c4247c53354360cd65015b4c34cd84746102008401527f19cc16dae81a96834123d5a2dd8d05882f0d209822b15bdca92814e7cdc378a76102208401526102408301526102608201527f2ab924f200306e4956687fce0ebbf9aa1f9c80f759682200fda36b8918e910596102808201527f1e9d1585762a2d6ac4e0edf659afbadd8e17697dba80904ee1d3a841aeb1f4336102a08201527f1f68193313b0fa03f6c0abfd2d96557cf45ea0c1c17893b400c3d7f0196140966102c08201527f2ed2409d093a4939782232832e918422a7b729a2c9f709ccaa9bdf59b9be056c6102e082015260405192839161066f8484610f3b565b8336843760085afa15908115610697575b506106885780f35b631ff3747d60e21b8152600490fd5b600191505114155f610680565b63a54f8e2760e01b8c5260048cfd5b80fd5b50346108205760403660031901126108205760043567ffffffffffffffff8111610820576106e8903690600401610f0d565b60243567ffffffffffffffff811161082057610708903690600401610f0d565b90918301610100848203126108205780601f85011215610820576040519361073261010086610f3b565b8490610100810192831161082057905b82821061084857505050810160e0828203126108205780601f83011215610820576040519161077260e084610f3b565b829060e0810192831161082057905b82821061082457505050303b1561082057604051630821332760e11b8152915f600484015b6008821061080a5750505061010482015f905b600782106107f4575050505f816101e481305afa80156107e9576107db575080f35b6107e791505f90610f3b565b005b6040513d5f823e3d90fd5b60208060019285518152019301910190916107b9565b60208060019285518152019301910190916107a6565b5f80fd5b8135815260209182019101610781565b634e487b7160e01b5f52604160045260245ffd5b8135815260209182019101610742565b346108205761010036600319011261082057366101041161082057604051610881608082610f3b565b6080368237610894602435600435610f5d565b81526108aa60843560a435604435606435610ffe565b602083015260408201526108c260e43560c435610f5d565b6060820152604051905f825b600482106108db57608084f35b60208060019285518152019301910190916108ce565b34610820575f3660031901126108205760206040517f4e1ae67d1bb1e750c48fa01e0f40783926027cd66180d9e72f8d9eed6d99cd368152f35b34610820576101e036600319011261082057366101041161082057366101e4116108205760405160408101907f0dfb26ff37b679c1c9f64380ee92d8056017a1daeceac11f25578ee59bbbb2cb815260208101917f1f402de6af9d5597d99032e3322a62a744e6c75b75e48a75fc6be86c1bb4abc883527f01ba57caf2dd0ef97400ff7847161ce4dddacb3cf6df86e975a3092ef4ec4cbd8152606082017f101dfe7b067da65c9cf7bf3ffe6b4da4052196a13a03d98cdd83bb502149dd4081527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f285e235e9aa51daf6644c1d8db1bcb7155fbbbe346d8d4954544281444a3e51a608087019580875284848460608160075afa911016838860808160065afa167f1d10def94709e1da4643a80863150fb716c567064a44e6643e693f5b309ab46a83527f1a4463dc9bc21f16ec0e74f6bb7c0072fff9068e586aa659b7b076d1b2d9b09c86526001610124359182895286868660608160075afa9310161616838860808160065afa167f1c8c32efdcd4dfc7fe88794bb241042946ad6944306e0aa2a04b97354d70150883527f10c00683bf679e37dfb6dbf161a29d3bb08589095af373b2bf9726cf67539e9e8652610144359081885285858560608160075afa92101616838860808160065afa167f1aedfa0cb2fab6fc6e352dc0d1c0661fe23798ad4a249db556c8df6e68c6c42983527f2183614418c533233ac0f8054c19e4ce9de9d69856d3f83f17460f40a1ab47de8652610164359081885285858560608160075afa92101616838860808160065afa167f1bb7fa5dae2eb590e5d44bec6863582c701098671e9c6a223252aeb0776af67683527f293e29d7dd18d1427304c85230692667fcc4c366273f953c51990c28383e5ac18652610184359081885285858560608160075afa92101616838860808160065afa167f2c5feb1355cffb9df96d15312b52d60c6cd02a79a2604f8cc157ec8fadb537ee83527f2d42cf984a0ebe1a82ddd33fece9cf15c95510ea90546e905e9a85d662affb5586526101a4359081885285858560608160075afa92101616838860808160065afa16947f199e7da818353ecc131688006b67af228d1f4c79dab5015975d3da71a8547d4f8352526101c43580955260608160075afa9210161660408260808160065afa16905191519015610efe5760405191610100600484377f03799f712eb0dbd68cf72393c65a47ac92ef43e1e84ff7d97495a72b9316f4a56101008401527f03c720838a19cb63a9447552c8d6e5553cfc22c58f3ce6634893bf58bebfea7c6101208401527f10d4839fbf942fbc277fc01bb04db9686a13169c569034dc6b9a9801b3b99b9f6101408401527f0bd1622089ed6ab4a26a3e93a93401aa95eaf142a51020d935b31a2b9714d48e6101608401527f0a6ced41020f9fc692f95d1e001e90bd616998e88209d178a5cd89bbc86c0e7b6101808401527f2edd4c02ad8154e40af01b76944e56950cdf1f53123ed6d94877e101fa09142d6101a08401527f18736cad4dee17ff032780bf31d62f073fd7dba456f7c89e1680e83b50aef40e6101c08401527f2498a73464d4d591d93503d794c2b6cb380f1511c9cd4b10ef87132027a1e10e6101e08401527f2b0589c30078a901e83c090cae1029d5c4247c53354360cd65015b4c34cd84746102008401527f19cc16dae81a96834123d5a2dd8d05882f0d209822b15bdca92814e7cdc378a76102208401526102408301526102608201527f2ab924f200306e4956687fce0ebbf9aa1f9c80f759682200fda36b8918e910596102808201527f1e9d1585762a2d6ac4e0edf659afbadd8e17697dba80904ee1d3a841aeb1f4336102a08201527f1f68193313b0fa03f6c0abfd2d96557cf45ea0c1c17893b400c3d7f0196140966102c08201527f2ed2409d093a4939782232832e918422a7b729a2c9f709ccaa9bdf59b9be056c6102e08201526020816103008160085afa90511615610eef57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b9181601f840112156108205782359167ffffffffffffffff8311610820576020838186019501011161082057565b90601f8019910116810190811067ffffffffffffffff82111761083457604052565b905f5160206115b65f395f51905f528210801590610fe7575b610eef57811580610fdf575b610fd957610fa65f5160206115b65f395f51905f52600381858181800909086113d6565b818103610fb557505060011b90565b5f5160206115b65f395f51905f52809106810306145f14610eef57600190811b1790565b50505f90565b508015610f82565b505f5160206115b65f395f51905f52811015610f76565b919093925f5160206115b65f395f51905f52831080159061121b575b8015611204575b80156111ed575b610eef5780828685171717156111e2579082916111455f5160206115b65f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f5160206115b65f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161111f81808b800981878009086113d6565b8408095f5160206115b65f395f51905f526111398261154d565b800914159586916113f9565b9290808214806111d9575b156111775750505050905f1461116f5760ff60025b169060021b179190565b60ff5f611165565b5f5160206115b65f395f51905f528091068103061491826111ba575b505015610eef57600191156111b25760ff60025b169060021b17179190565b60ff5f6111a7565b5f5160206115b65f395f51905f52919250819006810306145f80611193565b50838314611150565b50505090505f905f90565b505f5160206115b65f395f51905f52811015611028565b505f5160206115b65f395f51905f52821015611021565b505f5160206115b65f395f51905f5285101561101a565b8015611296578060011c915f5160206115b65f395f51905f52831015610eef576001806112755f5160206115b65f395f51905f52600381888181800909086113d6565b93161461127e57565b905f5160206115b65f395f51905f5280910681030690565b505f905f90565b8015806113ce575b6113c2578060021c92825f5160206115b65f395f51905f5285108015906113ab575b610eef5784815f5160206115b65f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816113759d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5086113f9565b80929160018082961614611387575050565b5f5160206115b65f395f51905f528093945080929550809106810306930681030690565b505f5160206115b65f395f51905f528110156112c7565b50505f905f905f905f90565b5081156112a5565b906113e08261154d565b915f5160206115b65f395f51905f5283800903610eef57565b915f5160206115b65f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816114519396949661144382808a8009818a8009086113d6565b90611541575b8608096113d6565b925f5160206115b65f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206115b65f395f51905f5260a083015260208260c08160055afa91519115610eef575f5160206115b65f395f51905f52826001920903610eef575f5160206115b65f395f51905f52908209925f5160206115b65f395f51905f528080808780090681030681878009081490811591611522575b50610eef57565b90505f5160206115b65f395f51905f528084860960020914155f61151b565b81809106810306611449565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206115b65f395f51905f5260a083015260208260c08160055afa91519115610eef5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220272fd4804e3164750cceb4f8e55663d28b333f315df88fb05a27e4715966d4d764736f6c634300081c0033",
}

// RevealShareVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use RevealShareVerifierMetaData.ABI instead.
var RevealShareVerifierABI = RevealShareVerifierMetaData.ABI

// RevealShareVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RevealShareVerifierMetaData.Bin instead.
var RevealShareVerifierBin = RevealShareVerifierMetaData.Bin

// DeployRevealShareVerifier deploys a new Ethereum contract, binding an instance of RevealShareVerifier to it.
func DeployRevealShareVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RevealShareVerifier, error) {
	parsed, err := RevealShareVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RevealShareVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RevealShareVerifier{RevealShareVerifierCaller: RevealShareVerifierCaller{contract: contract}, RevealShareVerifierTransactor: RevealShareVerifierTransactor{contract: contract}, RevealShareVerifierFilterer: RevealShareVerifierFilterer{contract: contract}}, nil
}

// RevealShareVerifier is an auto generated Go binding around an Ethereum contract.
type RevealShareVerifier struct {
	RevealShareVerifierCaller     // Read-only binding to the contract
	RevealShareVerifierTransactor // Write-only binding to the contract
	RevealShareVerifierFilterer   // Log filterer for contract events
}

// RevealShareVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type RevealShareVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevealShareVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RevealShareVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevealShareVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RevealShareVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevealShareVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RevealShareVerifierSession struct {
	Contract     *RevealShareVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RevealShareVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RevealShareVerifierCallerSession struct {
	Contract *RevealShareVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// RevealShareVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RevealShareVerifierTransactorSession struct {
	Contract     *RevealShareVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// RevealShareVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type RevealShareVerifierRaw struct {
	Contract *RevealShareVerifier // Generic contract binding to access the raw methods on
}

// RevealShareVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RevealShareVerifierCallerRaw struct {
	Contract *RevealShareVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// RevealShareVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RevealShareVerifierTransactorRaw struct {
	Contract *RevealShareVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRevealShareVerifier creates a new instance of RevealShareVerifier, bound to a specific deployed contract.
func NewRevealShareVerifier(address common.Address, backend bind.ContractBackend) (*RevealShareVerifier, error) {
	contract, err := bindRevealShareVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RevealShareVerifier{RevealShareVerifierCaller: RevealShareVerifierCaller{contract: contract}, RevealShareVerifierTransactor: RevealShareVerifierTransactor{contract: contract}, RevealShareVerifierFilterer: RevealShareVerifierFilterer{contract: contract}}, nil
}

// NewRevealShareVerifierCaller creates a new read-only instance of RevealShareVerifier, bound to a specific deployed contract.
func NewRevealShareVerifierCaller(address common.Address, caller bind.ContractCaller) (*RevealShareVerifierCaller, error) {
	contract, err := bindRevealShareVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RevealShareVerifierCaller{contract: contract}, nil
}

// NewRevealShareVerifierTransactor creates a new write-only instance of RevealShareVerifier, bound to a specific deployed contract.
func NewRevealShareVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*RevealShareVerifierTransactor, error) {
	contract, err := bindRevealShareVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RevealShareVerifierTransactor{contract: contract}, nil
}

// NewRevealShareVerifierFilterer creates a new log filterer instance of RevealShareVerifier, bound to a specific deployed contract.
func NewRevealShareVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*RevealShareVerifierFilterer, error) {
	contract, err := bindRevealShareVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RevealShareVerifierFilterer{contract: contract}, nil
}

// bindRevealShareVerifier binds a generic wrapper to an already deployed contract.
func bindRevealShareVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RevealShareVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevealShareVerifier *RevealShareVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevealShareVerifier.Contract.RevealShareVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevealShareVerifier *RevealShareVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevealShareVerifier.Contract.RevealShareVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevealShareVerifier *RevealShareVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevealShareVerifier.Contract.RevealShareVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevealShareVerifier *RevealShareVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevealShareVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevealShareVerifier *RevealShareVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevealShareVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevealShareVerifier *RevealShareVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevealShareVerifier.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_RevealShareVerifier *RevealShareVerifierCaller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int) ([4]*big.Int, error) {
	var out []interface{}
	err := _RevealShareVerifier.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_RevealShareVerifier *RevealShareVerifierSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _RevealShareVerifier.Contract.CompressProof(&_RevealShareVerifier.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_RevealShareVerifier *RevealShareVerifierCallerSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _RevealShareVerifier.Contract.CompressProof(&_RevealShareVerifier.CallOpts, proof)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_RevealShareVerifier *RevealShareVerifierCaller) ProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RevealShareVerifier.contract.Call(opts, &out, "provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_RevealShareVerifier *RevealShareVerifierSession) ProvingKeyHash() ([32]byte, error) {
	return _RevealShareVerifier.Contract.ProvingKeyHash(&_RevealShareVerifier.CallOpts)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_RevealShareVerifier *RevealShareVerifierCallerSession) ProvingKeyHash() ([32]byte, error) {
	return _RevealShareVerifier.Contract.ProvingKeyHash(&_RevealShareVerifier.CallOpts)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xf6d3293a.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[7] input) view returns()
func (_RevealShareVerifier *RevealShareVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [7]*big.Int) error {
	var out []interface{}
	err := _RevealShareVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xf6d3293a.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[7] input) view returns()
func (_RevealShareVerifier *RevealShareVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [7]*big.Int) error {
	return _RevealShareVerifier.Contract.VerifyCompressedProof(&_RevealShareVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xf6d3293a.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[7] input) view returns()
func (_RevealShareVerifier *RevealShareVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [7]*big.Int) error {
	return _RevealShareVerifier.Contract.VerifyCompressedProof(&_RevealShareVerifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x1042664e.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[7] input) view returns()
func (_RevealShareVerifier *RevealShareVerifierCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, input [7]*big.Int) error {
	var out []interface{}
	err := _RevealShareVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x1042664e.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[7] input) view returns()
func (_RevealShareVerifier *RevealShareVerifierSession) VerifyProof(proof [8]*big.Int, input [7]*big.Int) error {
	return _RevealShareVerifier.Contract.VerifyProof(&_RevealShareVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x1042664e.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[7] input) view returns()
func (_RevealShareVerifier *RevealShareVerifierCallerSession) VerifyProof(proof [8]*big.Int, input [7]*big.Int) error {
	return _RevealShareVerifier.Contract.VerifyProof(&_RevealShareVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_RevealShareVerifier *RevealShareVerifierCaller) VerifyProof0(opts *bind.CallOpts, proof []byte, input []byte) error {
	var out []interface{}
	err := _RevealShareVerifier.contract.Call(opts, &out, "verifyProof0", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_RevealShareVerifier *RevealShareVerifierSession) VerifyProof0(proof []byte, input []byte) error {
	return _RevealShareVerifier.Contract.VerifyProof0(&_RevealShareVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_RevealShareVerifier *RevealShareVerifierCallerSession) VerifyProof0(proof []byte, input []byte) error {
	return _RevealShareVerifier.Contract.VerifyProof0(&_RevealShareVerifier.CallOpts, proof, input)
}
