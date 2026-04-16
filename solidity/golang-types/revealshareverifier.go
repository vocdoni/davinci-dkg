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
	Bin: "0x6080806040523460155761160b908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80631042664e1461092b578063233ace11146108f157806344f6369214610858578063b8e72af6146106b65763f6d3293a14610051575f80fd5b346106b3576101603660031901126106b357366084116106b35736610164116106b357604051906103006100858184610f3b565b80368437610094600435611232565b6100a560249592953560443561129d565b919392906100b4606435611232565b9390926040519660408801967f189af10acccde3b9e62eda5d151a464ce2fa5626216f5fa10bcc0d2cb1a5d3ba89528860208101987f08e5c84e0dbd24150caae7dc6b4acc546794566f75fcc07b9d4e8a6107020be88a527f16cb8714d05702f33c0a41bd6194baa7dbcee8c92f43b838471fd3540cae320a81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f2e2d59fbe5b172b17104019c118b0580dff76f9b6352207369fa4689117d4e6d84527f05b327c42392634a781f0e7be0a47e7b2d91c8546546b5cc82f0348359c4740a6084359583608082019780895286828660608160075afa911016818360808160065afa167f1556c92014d8212833cbbf01dbbb47b1b8010cb9a90164f2b10840770a5b3bf985527f12e1da809443e45c4d320f6f54d4de7fe16168e74ac9847310b6b7d581f695798852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f1c40cb3facdbaa571b68dda167d0e19ec123587e786d0896597ce20cafd1187285527f1ae493f043da079c69f0a6b84a5ad39d80f687bc444b2e64c3de0f07524f7c95885260c43590818a5287838760608160075afa92101616818360808160065afa167f02c8cbfbced0f3e8e617a7571b53ab2c9fcda970182e9941b4364cea521bca5485527f2f1647a1a5cd30774a68fe6bcf740d3de6a666945bfda645d6503f1780ae0878885260e43590818a5287838760608160075afa92101616818360808160065afa167f0cdce2350179eea6a3bd03cc624eae4b00729a801ca150e4209a7e0a34174f3e85527f078ba5f602f5a85885a57ea2503f5fc8437608ea6bfa5c5688e446f90931879088526101043590818a5287838760608160075afa92101616818360808160065afa167f2f245495350b8a1d135d1bddeb11fc02d8b85fb2e5a67b2c67b7329d231dafe685527f20bc7adbf908a2e0bea84d2dff1d7aa3def249fcf4187aed48a9fafab8c2ed3d88526101243590818a5287838760608160075afa921016169160808160065afa16947f1c829d895a999d663fe41d651e7fcc5c382913c6abc2f16935a213dabdc3e4848352526101443580955260608160075afa9210161660408a60808160065afa169851975198156106a45760209a9b9c8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f2838fed8ad0f50f5c5b45ca3bf1deeb3920f0a4bbea03a835841d5395daca3446101008401527f018be3423accfefca029a82b59eb1984c76c77c972b396524ce66037a2cc8d586101208401527f26b88e614dace2eb0a8343cff7a10bb405e62e04fa1d7a77987de9977e49bc406101408401527f0684f4cfed99a316eb087139f80a62e7c127b99365dc074f657d89548107ee526101608401527f109fe36ec2e6a93486323002b660fbf890cbd012c8896d8c1c71636f60527ca06101808401527f2355bb079c069459f0e9021ed63433f34c44163b41927b0c0e1a145fe46a8c0f6101a08401527f0d641f6e50b211747d74acebeee7a7c80f05bd50769595ac6b1547e09e9617a66101c08401527f2bf1649225ea1670a0c7b74d17dce2e95b7fc22fdf78ef00b8ac687c8977861c6101e08401527f28caf3ddc3b25e441fee0ee9b75c7ac378e8ca89efa1ac6c79e12ce7950b9f026102008401527f14c311f7d2da0219f1fdf28cf0dab6908da003157726fe2122f402e11f72db836102208401526102408301526102608201527f2ee0e9bdec4d650a038834a2909da4e3a3e7209ae38ddb790da1f55e0d0ac6066102808201527f19f11855cce896796f61b3eaa7103efdfa2ff958f334965cc03ad60c7cb4f3f86102a08201527f130c36359a90a805e9dabb2bde54f1d454a27e77db4038c92779661c30d857806102c08201527f29c9cae2f8dd5384ce3d836378daf14017aede28f7fb0075a293582a0f5b96e86102e082015260405192839161066f8484610f3b565b8336843760085afa15908115610697575b506106885780f35b631ff3747d60e21b8152600490fd5b600191505114155f610680565b63a54f8e2760e01b8c5260048cfd5b80fd5b50346108205760403660031901126108205760043567ffffffffffffffff8111610820576106e8903690600401610f0d565b60243567ffffffffffffffff811161082057610708903690600401610f0d565b90918301610100848203126108205780601f85011215610820576040519361073261010086610f3b565b8490610100810192831161082057905b82821061084857505050810160e0828203126108205780601f83011215610820576040519161077260e084610f3b565b829060e0810192831161082057905b82821061082457505050303b1561082057604051630821332760e11b8152915f600484015b6008821061080a5750505061010482015f905b600782106107f4575050505f816101e481305afa80156107e9576107db575080f35b6107e791505f90610f3b565b005b6040513d5f823e3d90fd5b60208060019285518152019301910190916107b9565b60208060019285518152019301910190916107a6565b5f80fd5b8135815260209182019101610781565b634e487b7160e01b5f52604160045260245ffd5b8135815260209182019101610742565b346108205761010036600319011261082057366101041161082057604051610881608082610f3b565b6080368237610894602435600435610f5d565b81526108aa60843560a435604435606435610ffe565b602083015260408201526108c260e43560c435610f5d565b6060820152604051905f825b600482106108db57608084f35b60208060019285518152019301910190916108ce565b34610820575f3660031901126108205760206040517f4e1ae67d1bb1e750c48fa01e0f40783926027cd66180d9e72f8d9eed6d99cd368152f35b34610820576101e036600319011261082057366101041161082057366101e4116108205760405160408101907f189af10acccde3b9e62eda5d151a464ce2fa5626216f5fa10bcc0d2cb1a5d3ba815260208101917f08e5c84e0dbd24150caae7dc6b4acc546794566f75fcc07b9d4e8a6107020be883527f16cb8714d05702f33c0a41bd6194baa7dbcee8c92f43b838471fd3540cae320a8152606082017f2e2d59fbe5b172b17104019c118b0580dff76f9b6352207369fa4689117d4e6d81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f05b327c42392634a781f0e7be0a47e7b2d91c8546546b5cc82f0348359c4740a608087019580875284848460608160075afa911016838860808160065afa167f1556c92014d8212833cbbf01dbbb47b1b8010cb9a90164f2b10840770a5b3bf983527f12e1da809443e45c4d320f6f54d4de7fe16168e74ac9847310b6b7d581f6957986526001610124359182895286868660608160075afa9310161616838860808160065afa167f1c40cb3facdbaa571b68dda167d0e19ec123587e786d0896597ce20cafd1187283527f1ae493f043da079c69f0a6b84a5ad39d80f687bc444b2e64c3de0f07524f7c958652610144359081885285858560608160075afa92101616838860808160065afa167f02c8cbfbced0f3e8e617a7571b53ab2c9fcda970182e9941b4364cea521bca5483527f2f1647a1a5cd30774a68fe6bcf740d3de6a666945bfda645d6503f1780ae08788652610164359081885285858560608160075afa92101616838860808160065afa167f0cdce2350179eea6a3bd03cc624eae4b00729a801ca150e4209a7e0a34174f3e83527f078ba5f602f5a85885a57ea2503f5fc8437608ea6bfa5c5688e446f9093187908652610184359081885285858560608160075afa92101616838860808160065afa167f2f245495350b8a1d135d1bddeb11fc02d8b85fb2e5a67b2c67b7329d231dafe683527f20bc7adbf908a2e0bea84d2dff1d7aa3def249fcf4187aed48a9fafab8c2ed3d86526101a4359081885285858560608160075afa92101616838860808160065afa16947f1c829d895a999d663fe41d651e7fcc5c382913c6abc2f16935a213dabdc3e4848352526101c43580955260608160075afa9210161660408260808160065afa16905191519015610efe5760405191610100600484377f2838fed8ad0f50f5c5b45ca3bf1deeb3920f0a4bbea03a835841d5395daca3446101008401527f018be3423accfefca029a82b59eb1984c76c77c972b396524ce66037a2cc8d586101208401527f26b88e614dace2eb0a8343cff7a10bb405e62e04fa1d7a77987de9977e49bc406101408401527f0684f4cfed99a316eb087139f80a62e7c127b99365dc074f657d89548107ee526101608401527f109fe36ec2e6a93486323002b660fbf890cbd012c8896d8c1c71636f60527ca06101808401527f2355bb079c069459f0e9021ed63433f34c44163b41927b0c0e1a145fe46a8c0f6101a08401527f0d641f6e50b211747d74acebeee7a7c80f05bd50769595ac6b1547e09e9617a66101c08401527f2bf1649225ea1670a0c7b74d17dce2e95b7fc22fdf78ef00b8ac687c8977861c6101e08401527f28caf3ddc3b25e441fee0ee9b75c7ac378e8ca89efa1ac6c79e12ce7950b9f026102008401527f14c311f7d2da0219f1fdf28cf0dab6908da003157726fe2122f402e11f72db836102208401526102408301526102608201527f2ee0e9bdec4d650a038834a2909da4e3a3e7209ae38ddb790da1f55e0d0ac6066102808201527f19f11855cce896796f61b3eaa7103efdfa2ff958f334965cc03ad60c7cb4f3f86102a08201527f130c36359a90a805e9dabb2bde54f1d454a27e77db4038c92779661c30d857806102c08201527f29c9cae2f8dd5384ce3d836378daf14017aede28f7fb0075a293582a0f5b96e86102e08201526020816103008160085afa90511615610eef57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b9181601f840112156108205782359167ffffffffffffffff8311610820576020838186019501011161082057565b90601f8019910116810190811067ffffffffffffffff82111761083457604052565b905f5160206115b65f395f51905f528210801590610fe7575b610eef57811580610fdf575b610fd957610fa65f5160206115b65f395f51905f52600381858181800909086113d6565b818103610fb557505060011b90565b5f5160206115b65f395f51905f52809106810306145f14610eef57600190811b1790565b50505f90565b508015610f82565b505f5160206115b65f395f51905f52811015610f76565b919093925f5160206115b65f395f51905f52831080159061121b575b8015611204575b80156111ed575b610eef5780828685171717156111e2579082916111455f5160206115b65f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f5160206115b65f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161111f81808b800981878009086113d6565b8408095f5160206115b65f395f51905f526111398261154d565b800914159586916113f9565b9290808214806111d9575b156111775750505050905f1461116f5760ff60025b169060021b179190565b60ff5f611165565b5f5160206115b65f395f51905f528091068103061491826111ba575b505015610eef57600191156111b25760ff60025b169060021b17179190565b60ff5f6111a7565b5f5160206115b65f395f51905f52919250819006810306145f80611193565b50838314611150565b50505090505f905f90565b505f5160206115b65f395f51905f52811015611028565b505f5160206115b65f395f51905f52821015611021565b505f5160206115b65f395f51905f5285101561101a565b8015611296578060011c915f5160206115b65f395f51905f52831015610eef576001806112755f5160206115b65f395f51905f52600381888181800909086113d6565b93161461127e57565b905f5160206115b65f395f51905f5280910681030690565b505f905f90565b8015806113ce575b6113c2578060021c92825f5160206115b65f395f51905f5285108015906113ab575b610eef5784815f5160206115b65f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816113759d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5086113f9565b80929160018082961614611387575050565b5f5160206115b65f395f51905f528093945080929550809106810306930681030690565b505f5160206115b65f395f51905f528110156112c7565b50505f905f905f905f90565b5081156112a5565b906113e08261154d565b915f5160206115b65f395f51905f5283800903610eef57565b915f5160206115b65f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816114519396949661144382808a8009818a8009086113d6565b90611541575b8608096113d6565b925f5160206115b65f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206115b65f395f51905f5260a083015260208260c08160055afa91519115610eef575f5160206115b65f395f51905f52826001920903610eef575f5160206115b65f395f51905f52908209925f5160206115b65f395f51905f528080808780090681030681878009081490811591611522575b50610eef57565b90505f5160206115b65f395f51905f528084860960020914155f61151b565b81809106810306611449565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206115b65f395f51905f5260a083015260208260c08160055afa91519115610eef5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220ed8d1ee32cd49ce353593a86848678c432fd8619a376a6ee082773f22f9e7c1064736f6c634300081c0033",
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
