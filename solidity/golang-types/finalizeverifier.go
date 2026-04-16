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
	Bin: "0x6080806040523460155761177c908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace11146110465750806344f6369214610fad5780635f89feef146108a45780638a3ae438146101f95763b8e72af614610055575f80fd5b346101c15760403660031901126101c15760043567ffffffffffffffff81116101c15761008690369060040161107e565b60243567ffffffffffffffff81116101c1576100a690369060040161107e565b90918301610100848203126101c15780601f850112156101c157604051936100d0610100866110ac565b849061010081019283116101c157905b8282106101e9575050508101610120828203126101c15780601f830112156101c15760405191610112610120846110ac565b829061012081019283116101c157905b8282106101c557505050303b156101c1576040516311475c8760e31b8152915f600484015b600882106101ab5750505061010482015f905b60098210610195575050505f8161022481305afa801561018a5761017c575080f35b61018891505f906110ac565b005b6040513d5f823e3d90fd5b602080600192855181520193019101909161015a565b6020806001928551815201930191019091610147565b5f80fd5b8135815260209182019101610122565b634e487b7160e01b5f52604160045260245ffd5b81358152602091820191016100e0565b346101c1576102203660031901126101c15736610104116101c15736610224116101c15760405160408101907f030a68dabf230ac33bd3211332ff270edb49b5c71317121823e5338e2e9b5458815260208101917f1198272025eb6e97d465d5ee11182cbd4d28fb931eca2f5a5f9c89dbd97bccc783527f2f69e36b72321b89f9247d89dcec83aa2731dfd6c65dbab1f043d5aafd280d218152606082017f28c2bf34cce8f6477e5708ee11d6ef876f2b1f41c69ae558d9d214d330f6e68e81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f254152df584a61eab3732c13e53c085c8c3e7b2093031b3408eb9627ab7e91f3608087019580875284848460608160075afa911016838860808160065afa167f0d572a533d326450de44bf89f9768cb8ee07dfcc4556e975e43c73ef4ecfc19883527f2fee2b3788ddf1b833181d8f93b191ee254cc1f940e746b5f7b28a04f1f4c12186526001610124359182895286868660608160075afa9310161616838860808160065afa167f14f96d8dd8509f287446c81c69f6afeb99cd3b479784ca125bc3fa8bca398fb583527f206997b4d6fd2ae6436825d917bb8695f52ccd47ab016311edb9d951e17dc5e18652610144359081885285858560608160075afa92101616838860808160065afa167f2c5a88ea34dc7585ef7522f6c5935dcde4bcffc8ce681afb94b92a4f94d4331e83527f1cb9796c9edaa17bc3ca074f72ef04314a4ff4f8e440da6e6dbb95a069db09d78652610164359081885285858560608160075afa92101616838860808160065afa167f09d8bb9d2a122476d09307b5dae6d56eed672c57393753f72fc89316bc8204a983527f1701db9fc0d1f4020996b0a314873d867ebd4bf9e79dedf7d82835f6cbe70b638652610184359081885285858560608160075afa92101616838860808160065afa167f22d059f40b179be7d58d1860b30760743855501d4001352b5b8454af4b6aff9783527f08c2d6126d902148f3b6de090ddef1992d78b7346f3d51d9eab621afe535935f86526101a4359081885285858560608160075afa92101616838860808160065afa167ee6a9fcee8f452d62d7309eb57297cc40b78fdaffd22fdd45fe8b694b47e56d83527f27fb399018ed5a23b76c09a9085a06dd76e14bd44a39a6e9916099cc39cea8b786526101c4359081885285858560608160075afa92101616838860808160065afa167f04a8f3c5144f66d34ca0f8886d53ac9831b4b11587a1f5e50f4cea597785689f83527f131ce5ad39d2a5e42487fb021f3b08e89b891e061e1d313542b04a00d380300c86526101e4359081885285858560608160075afa92101616838860808160065afa16947f10206bf3a50ea3a65aad52a60a077ca9bfc4e6c05a2870696767d4c0bc6b2db78352526102043580955260608160075afa9210161660408260808160065afa169051915190156108955760405191610100600484377eeca5ae89a65aac73662524ae6c6abaa7bb068c93dda1fd4fd0c9f71d2006df6101008401527f214c623c7c765be1286963b0b804bed5eb5af215b700daefdcf6263cdf78a2d16101208401527f2f1c2daa5b2f1fb0710f04e074beb4845c800f9932f2ab31ca03f5e4dda20f346101408401527e200754b5e3db6fd954d35ec868d68371fdb21bebb3ff294a1a4493edb07a446101608401527f1b0aacd47a9fd9a97a1be755e3a4a227edcf9a93f4d07112578e4cb9833781f36101808401527f1ea50dbebc1189458d2ffbb95d7c29d58bcce215cb66b782acc23ebd58a8b5706101a08401527f04558e7076f57aa2a3d279fef345dc8d1ae9c2bff29c4f940584d9cd500b72e36101c08401527f0d428c2018e96d994a9df0927288919890c3bad00562bad355fd5983d2f30b366101e08401527f1533a1f9688d6093aed51b1367b6fe0a8d6cad3e4df351f35968e8c1342ea1626102008401527f07197ee3d6d0adc8c8f9a28621ab81e981936085cae2a0378653961e0e0ef0f96102208401526102408301526102608201527f2d4c7d78663f8750b960ab969058043c98d97a6309ef1e57e278928930e794696102808201527f24abaf759860dc8db481358077cedd238ddcf0b4181782fc6766bb6ad46dca106102a08201527f2e99668a5d893e05d6be4efcec6823fb585b3f9b9d84fd82ce54eb68a1f476d56102c08201527f1d22677cc81e3428387a8ce9c3bbca03b2148b3494e9e47543414488b16ff89d6102e08201526020816103008160085afa9051161561088657005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b346101c1576101a03660031901126101c157366084116101c157366101a4116101c1576103006040516108d782826110ac565b813682376108e66004356113a3565b6108f760249392933560443561140e565b919392906109066064356113a3565b9390926040519660408801967f030a68dabf230ac33bd3211332ff270edb49b5c71317121823e5338e2e9b545889528860208101987f1198272025eb6e97d465d5ee11182cbd4d28fb931eca2f5a5f9c89dbd97bccc78a527f2f69e36b72321b89f9247d89dcec83aa2731dfd6c65dbab1f043d5aafd280d2181527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f28c2bf34cce8f6477e5708ee11d6ef876f2b1f41c69ae558d9d214d330f6e68e84527f254152df584a61eab3732c13e53c085c8c3e7b2093031b3408eb9627ab7e91f36084359583608082019780895286828660608160075afa911016818360808160065afa167f0d572a533d326450de44bf89f9768cb8ee07dfcc4556e975e43c73ef4ecfc19885527f2fee2b3788ddf1b833181d8f93b191ee254cc1f940e746b5f7b28a04f1f4c1218852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f14f96d8dd8509f287446c81c69f6afeb99cd3b479784ca125bc3fa8bca398fb585527f206997b4d6fd2ae6436825d917bb8695f52ccd47ab016311edb9d951e17dc5e1885260c43590818a5287838760608160075afa92101616818360808160065afa167f2c5a88ea34dc7585ef7522f6c5935dcde4bcffc8ce681afb94b92a4f94d4331e85527f1cb9796c9edaa17bc3ca074f72ef04314a4ff4f8e440da6e6dbb95a069db09d7885260e43590818a5287838760608160075afa92101616818360808160065afa167f09d8bb9d2a122476d09307b5dae6d56eed672c57393753f72fc89316bc8204a985527f1701db9fc0d1f4020996b0a314873d867ebd4bf9e79dedf7d82835f6cbe70b6388526101043590818a5287838760608160075afa92101616818360808160065afa167f22d059f40b179be7d58d1860b30760743855501d4001352b5b8454af4b6aff9785527f08c2d6126d902148f3b6de090ddef1992d78b7346f3d51d9eab621afe535935f88526101243590818a5287838760608160075afa92101616818360808160065afa167ee6a9fcee8f452d62d7309eb57297cc40b78fdaffd22fdd45fe8b694b47e56d85527f27fb399018ed5a23b76c09a9085a06dd76e14bd44a39a6e9916099cc39cea8b788526101443590818a5287838760608160075afa92101616818360808160065afa167f04a8f3c5144f66d34ca0f8886d53ac9831b4b11587a1f5e50f4cea597785689f85527f131ce5ad39d2a5e42487fb021f3b08e89b891e061e1d313542b04a00d380300c88526101643590818a5287838760608160075afa921016169160808160065afa16947f10206bf3a50ea3a65aad52a60a077ca9bfc4e6c05a2870696767d4c0bc6b2db78352526101843580955260608160075afa9210161660408a60808160065afa169851975198156108955760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527eeca5ae89a65aac73662524ae6c6abaa7bb068c93dda1fd4fd0c9f71d2006df6101008401527f214c623c7c765be1286963b0b804bed5eb5af215b700daefdcf6263cdf78a2d16101208401527f2f1c2daa5b2f1fb0710f04e074beb4845c800f9932f2ab31ca03f5e4dda20f346101408401527e200754b5e3db6fd954d35ec868d68371fdb21bebb3ff294a1a4493edb07a446101608401527f1b0aacd47a9fd9a97a1be755e3a4a227edcf9a93f4d07112578e4cb9833781f36101808401527f1ea50dbebc1189458d2ffbb95d7c29d58bcce215cb66b782acc23ebd58a8b5706101a08401527f04558e7076f57aa2a3d279fef345dc8d1ae9c2bff29c4f940584d9cd500b72e36101c08401527f0d428c2018e96d994a9df0927288919890c3bad00562bad355fd5983d2f30b366101e08401527f1533a1f9688d6093aed51b1367b6fe0a8d6cad3e4df351f35968e8c1342ea1626102008401527f07197ee3d6d0adc8c8f9a28621ab81e981936085cae2a0378653961e0e0ef0f96102208401526102408301526102608201527f2d4c7d78663f8750b960ab969058043c98d97a6309ef1e57e278928930e794696102808201527f24abaf759860dc8db481358077cedd238ddcf0b4181782fc6766bb6ad46dca106102a08201527f2e99668a5d893e05d6be4efcec6823fb585b3f9b9d84fd82ce54eb68a1f476d56102c08201527f1d22677cc81e3428387a8ce9c3bbca03b2148b3494e9e47543414488b16ff89d6102e0820152604051928391610f8884846110ac565b8336843760085afa15908115610fa0575b5061088657005b6001915051141581610f99565b346101c1576101003660031901126101c15736610104116101c157604051610fd66080826110ac565b6080368237610fe96024356004356110ce565b8152610fff60843560a43560443560643561116f565b6020830152604082015261101760e43560c4356110ce565b6060820152604051905f825b6004821061103057608084f35b6020806001928551815201930191019091611023565b346101c1575f3660031901126101c157807fa5a0381635973c2599e2e0bd559739beb8912aa8607280773b4f2f24f9e7f03560209252f35b9181601f840112156101c15782359167ffffffffffffffff83116101c157602083818601950101116101c157565b90601f8019910116810190811067ffffffffffffffff8211176101d557604052565b905f5160206117275f395f51905f528210801590611158575b61088657811580611150575b61114a576111175f5160206117275f395f51905f5260038185818180090908611547565b81810361112657505060011b90565b5f5160206117275f395f51905f52809106810306145f1461088657600190811b1790565b50505f90565b5080156110f3565b505f5160206117275f395f51905f528110156110e7565b919093925f5160206117275f395f51905f52831080159061138c575b8015611375575b801561135e575b610886578082868517171715611353579082916112b65f5160206117275f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f5160206117275f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161129081808b80098187800908611547565b8408095f5160206117275f395f51905f526112aa826116be565b8009141595869161156a565b92908082148061134a575b156112e85750505050905f146112e05760ff60025b169060021b179190565b60ff5f6112d6565b5f5160206117275f395f51905f5280910681030614918261132b575b50501561088657600191156113235760ff60025b169060021b17179190565b60ff5f611318565b5f5160206117275f395f51905f52919250819006810306145f80611304565b508383146112c1565b50505090505f905f90565b505f5160206117275f395f51905f52811015611199565b505f5160206117275f395f51905f52821015611192565b505f5160206117275f395f51905f5285101561118b565b8015611407578060011c915f5160206117275f395f51905f52831015610886576001806113e65f5160206117275f395f51905f5260038188818180090908611547565b9316146113ef57565b905f5160206117275f395f51905f5280910681030690565b505f905f90565b80158061153f575b611533578060021c92825f5160206117275f395f51905f52851080159061151c575b6108865784815f5160206117275f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816114e69d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50861156a565b809291600180829616146114f8575050565b5f5160206117275f395f51905f528093945080929550809106810306930681030690565b505f5160206117275f395f51905f52811015611438565b50505f905f905f905f90565b508115611416565b90611551826116be565b915f5160206117275f395f51905f528380090361088657565b915f5160206117275f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816115c2939694966115b482808a8009818a800908611547565b906116b2575b860809611547565b925f5160206117275f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206117275f395f51905f5260a083015260208260c08160055afa91519115610886575f5160206117275f395f51905f52826001920903610886575f5160206117275f395f51905f52908209925f5160206117275f395f51905f528080808780090681030681878009081490811591611693575b5061088657565b90505f5160206117275f395f51905f528084860960020914155f61168c565b818091068103066115ba565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206117275f395f51905f5260a083015260208260c08160055afa915191156108865756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220096c265a02aa3ce8ec39f7681a806df6334063ae3ff13c1525cf0d57953ec1ea64736f6c634300081c0033",
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
