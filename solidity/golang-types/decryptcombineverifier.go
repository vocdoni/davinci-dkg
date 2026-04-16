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

// DecryptCombineVerifierMetaData contains all meta data concerning the DecryptCombineVerifier contract.
var DecryptCombineVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[7]\",\"internalType\":\"uint256[7]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"input\",\"type\":\"uint256[7]\",\"internalType\":\"uint256[7]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x60808060405234601557611609908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80631042664e1461092a578063233ace11146108f057806344f6369214610857578063b8e72af6146106b55763f6d3293a14610051575f80fd5b346106b2576101603660031901126106b257366084116106b25736610164116106b257604051906103006100858184610f39565b80368437610094600435611230565b6100a560249592953560443561129b565b919392906100b4606435611230565b9390926040519660408801967f1583e83195e85805dbd435a872cfe0edd16807c3cd03d016f18213c9a87afdf889528860208101987f115493ce0d22f6ff84c8d974b3a15c65593697b658948f5228412011f8be22da8a527f041a78e6f4b12e0aab1104d324f5b4c17369394989ab2bccb7dfa3348483e18781527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f0eb32d2807d8c8689a82f6e72e2d0d69aeb9e07a645bd47aa47a0cbf2e3fb71584527f27e62c3f043050f47e9847f3601e478a10a921de9974bfbcfc934b11f65239ae6084359583608082019780895286828660608160075afa911016818360808160065afa167f167cddd130fdf83402360cb2439ff31dffb80c464b4047440dad6885cba1447885527f182f2557fb069a23eef70a4d45eba6861b4e3fdfe9d7331d1bcd14f96c074d0e8852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f2f8285c6ec61a83544d8022eb79010541fb814922050373dda3d9b4c328e284985527f2a440c01b593474525f4e6d9a1fd9ca3bde0f8bc226ee822165d6790f68b8564885260c43590818a5287838760608160075afa92101616818360808160065afa167f07983431235a991980b118faffdbe1cd1585825c342271ea01d3f682680fed0185527f152b81de5ccda0d19165af99a629f1d746760245a356434537b4ec269d830f0b885260e43590818a5287838760608160075afa92101616818360808160065afa167f19fea2eef249c7e2db398142c2b8f2ffb1c85e90721a63d3d73195a1443a232085527f2d8408bb04f61810312d3257c68fd7e1a8b9fb4b21396aee69ea757975f2c25088526101043590818a5287838760608160075afa92101616818360808160065afa167f270d82fe7f89c0303ea1e5657e6537a1250697a0dfb34522bcf50458fc05907385527f034622472227d6c3bca801e6bc5e88126b0904769ef8e6c20b1b7620042cb41288526101243590818a5287838760608160075afa921016169160808160065afa16947f13ce2ebab1c5949bcb7225fbd098b9e1974a2e806f5f3aa0a6d5dac8b98438488352526101443580955260608160075afa9210161660408a60808160065afa169851975198156106a35760209a9b9c8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f2f2b722d4d657a0b71ef03a3f7f5d57fbba5e1bef1a2087ac43ceb62ee846e216101008401527f07ce0d78242326aabcd81700e847de16e29179f5030fbd6cb54ce13625dc25b46101208401527f272b0d5a90bb9aad9ea451e5d4e0aaa22b08ee99a9bd0379b2dad311ee79510c6101408401527f0ec2825218ef54e1f2747b9132611c8e04eaa2fed318744c0d6b8682d9630d8a6101608401527f27623dc6297fe1e20111683c8b6b0568d6362331df0a4c603f84358ffc76ce2a6101808401527f2e554113150e1466e142c3d97e69ed113086719ea21af9ee1a8ab996de5e46de6101a08401527f20860ee0791ee69aaa4c543857f053fc408fbedc770e374e28d762a862d753326101c08401527f0fb19f3950661231f7f52d6fc2a483ded424e7db9c8ba29bd0e9f4c25485b26c6101e08401527f25d3c02e0a16e98310581a1cd8b67b72083b4f2636bd242b89dbe6101337d5456102008401527f2a298c50b674f4f8ef1a1b24ce2f08cf07c00f3d26c9f8a4fbd7c248b8ae50ed6102208401526102408301526102608201527f028cf8e5559c6d6797e8f2e43a22997827d8bfcbb678693d8057a5075f30a4a46102808201527f25eaa168a0eea89a699216797e61140308ccfc571447bf8e2a8a26a5c7db07ca6102a08201527f1976aa5afaa38bd490535d33da2760188dc233759abaf0d4333fae065901ed156102c08201527eec028f346d44ea0388816814996a1f54f22ea233b94809c4ee9a2c77b1336b6102e082015260405192839161066e8484610f39565b8336843760085afa15908115610696575b506106875780f35b631ff3747d60e21b8152600490fd5b600191505114155f61067f565b63a54f8e2760e01b8c5260048cfd5b80fd5b503461081f57604036600319011261081f5760043567ffffffffffffffff811161081f576106e7903690600401610f0b565b60243567ffffffffffffffff811161081f57610707903690600401610f0b565b909183016101008482031261081f5780601f8501121561081f576040519361073161010086610f39565b8490610100810192831161081f57905b82821061084757505050810160e08282031261081f5780601f8301121561081f576040519161077160e084610f39565b829060e0810192831161081f57905b82821061082357505050303b1561081f57604051630821332760e11b8152915f600484015b600882106108095750505061010482015f905b600782106107f3575050505f816101e481305afa80156107e8576107da575080f35b6107e691505f90610f39565b005b6040513d5f823e3d90fd5b60208060019285518152019301910190916107b8565b60208060019285518152019301910190916107a5565b5f80fd5b8135815260209182019101610780565b634e487b7160e01b5f52604160045260245ffd5b8135815260209182019101610741565b3461081f5761010036600319011261081f57366101041161081f57604051610880608082610f39565b6080368237610893602435600435610f5b565b81526108a960843560a435604435606435610ffc565b602083015260408201526108c160e43560c435610f5b565b6060820152604051905f825b600482106108da57608084f35b60208060019285518152019301910190916108cd565b3461081f575f36600319011261081f5760206040517f43d858ca8532f6b67880bc3a61cd9e3a39b4db1d4d59828d3e84452b7dc76a4f8152f35b3461081f576101e036600319011261081f57366101041161081f57366101e41161081f5760405160408101907f1583e83195e85805dbd435a872cfe0edd16807c3cd03d016f18213c9a87afdf8815260208101917f115493ce0d22f6ff84c8d974b3a15c65593697b658948f5228412011f8be22da83527f041a78e6f4b12e0aab1104d324f5b4c17369394989ab2bccb7dfa3348483e1878152606082017f0eb32d2807d8c8689a82f6e72e2d0d69aeb9e07a645bd47aa47a0cbf2e3fb71581527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f27e62c3f043050f47e9847f3601e478a10a921de9974bfbcfc934b11f65239ae608087019580875284848460608160075afa911016838860808160065afa167f167cddd130fdf83402360cb2439ff31dffb80c464b4047440dad6885cba1447883527f182f2557fb069a23eef70a4d45eba6861b4e3fdfe9d7331d1bcd14f96c074d0e86526001610124359182895286868660608160075afa9310161616838860808160065afa167f2f8285c6ec61a83544d8022eb79010541fb814922050373dda3d9b4c328e284983527f2a440c01b593474525f4e6d9a1fd9ca3bde0f8bc226ee822165d6790f68b85648652610144359081885285858560608160075afa92101616838860808160065afa167f07983431235a991980b118faffdbe1cd1585825c342271ea01d3f682680fed0183527f152b81de5ccda0d19165af99a629f1d746760245a356434537b4ec269d830f0b8652610164359081885285858560608160075afa92101616838860808160065afa167f19fea2eef249c7e2db398142c2b8f2ffb1c85e90721a63d3d73195a1443a232083527f2d8408bb04f61810312d3257c68fd7e1a8b9fb4b21396aee69ea757975f2c2508652610184359081885285858560608160075afa92101616838860808160065afa167f270d82fe7f89c0303ea1e5657e6537a1250697a0dfb34522bcf50458fc05907383527f034622472227d6c3bca801e6bc5e88126b0904769ef8e6c20b1b7620042cb41286526101a4359081885285858560608160075afa92101616838860808160065afa16947f13ce2ebab1c5949bcb7225fbd098b9e1974a2e806f5f3aa0a6d5dac8b98438488352526101c43580955260608160075afa9210161660408260808160065afa16905191519015610efc5760405191610100600484377f2f2b722d4d657a0b71ef03a3f7f5d57fbba5e1bef1a2087ac43ceb62ee846e216101008401527f07ce0d78242326aabcd81700e847de16e29179f5030fbd6cb54ce13625dc25b46101208401527f272b0d5a90bb9aad9ea451e5d4e0aaa22b08ee99a9bd0379b2dad311ee79510c6101408401527f0ec2825218ef54e1f2747b9132611c8e04eaa2fed318744c0d6b8682d9630d8a6101608401527f27623dc6297fe1e20111683c8b6b0568d6362331df0a4c603f84358ffc76ce2a6101808401527f2e554113150e1466e142c3d97e69ed113086719ea21af9ee1a8ab996de5e46de6101a08401527f20860ee0791ee69aaa4c543857f053fc408fbedc770e374e28d762a862d753326101c08401527f0fb19f3950661231f7f52d6fc2a483ded424e7db9c8ba29bd0e9f4c25485b26c6101e08401527f25d3c02e0a16e98310581a1cd8b67b72083b4f2636bd242b89dbe6101337d5456102008401527f2a298c50b674f4f8ef1a1b24ce2f08cf07c00f3d26c9f8a4fbd7c248b8ae50ed6102208401526102408301526102608201527f028cf8e5559c6d6797e8f2e43a22997827d8bfcbb678693d8057a5075f30a4a46102808201527f25eaa168a0eea89a699216797e61140308ccfc571447bf8e2a8a26a5c7db07ca6102a08201527f1976aa5afaa38bd490535d33da2760188dc233759abaf0d4333fae065901ed156102c08201527eec028f346d44ea0388816814996a1f54f22ea233b94809c4ee9a2c77b1336b6102e08201526020816103008160085afa90511615610eed57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b9181601f8401121561081f5782359167ffffffffffffffff831161081f576020838186019501011161081f57565b90601f8019910116810190811067ffffffffffffffff82111761083357604052565b905f5160206115b45f395f51905f528210801590610fe5575b610eed57811580610fdd575b610fd757610fa45f5160206115b45f395f51905f52600381858181800909086113d4565b818103610fb357505060011b90565b5f5160206115b45f395f51905f52809106810306145f14610eed57600190811b1790565b50505f90565b508015610f80565b505f5160206115b45f395f51905f52811015610f74565b919093925f5160206115b45f395f51905f528310801590611219575b8015611202575b80156111eb575b610eed5780828685171717156111e0579082916111435f5160206115b45f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f5160206115b45f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161111d81808b800981878009086113d4565b8408095f5160206115b45f395f51905f526111378261154b565b800914159586916113f7565b9290808214806111d7575b156111755750505050905f1461116d5760ff60025b169060021b179190565b60ff5f611163565b5f5160206115b45f395f51905f528091068103061491826111b8575b505015610eed57600191156111b05760ff60025b169060021b17179190565b60ff5f6111a5565b5f5160206115b45f395f51905f52919250819006810306145f80611191565b5083831461114e565b50505090505f905f90565b505f5160206115b45f395f51905f52811015611026565b505f5160206115b45f395f51905f5282101561101f565b505f5160206115b45f395f51905f52851015611018565b8015611294578060011c915f5160206115b45f395f51905f52831015610eed576001806112735f5160206115b45f395f51905f52600381888181800909086113d4565b93161461127c57565b905f5160206115b45f395f51905f5280910681030690565b505f905f90565b8015806113cc575b6113c0578060021c92825f5160206115b45f395f51905f5285108015906113a9575b610eed5784815f5160206115b45f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816113739d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5086113f7565b80929160018082961614611385575050565b5f5160206115b45f395f51905f528093945080929550809106810306930681030690565b505f5160206115b45f395f51905f528110156112c5565b50505f905f905f905f90565b5081156112a3565b906113de8261154b565b915f5160206115b45f395f51905f5283800903610eed57565b915f5160206115b45f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161144f9396949661144182808a8009818a8009086113d4565b9061153f575b8608096113d4565b925f5160206115b45f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206115b45f395f51905f5260a083015260208260c08160055afa91519115610eed575f5160206115b45f395f51905f52826001920903610eed575f5160206115b45f395f51905f52908209925f5160206115b45f395f51905f528080808780090681030681878009081490811591611520575b50610eed57565b90505f5160206115b45f395f51905f528084860960020914155f611519565b81809106810306611447565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206115b45f395f51905f5260a083015260208260c08160055afa91519115610eed5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a26469706673582212202ace5e98ba4c2e5b9088355abcb10fc86d88513d1d56dc701ddbe9c670b835ec64736f6c634300081c0033",
}

// DecryptCombineVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use DecryptCombineVerifierMetaData.ABI instead.
var DecryptCombineVerifierABI = DecryptCombineVerifierMetaData.ABI

// DecryptCombineVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DecryptCombineVerifierMetaData.Bin instead.
var DecryptCombineVerifierBin = DecryptCombineVerifierMetaData.Bin

// DeployDecryptCombineVerifier deploys a new Ethereum contract, binding an instance of DecryptCombineVerifier to it.
func DeployDecryptCombineVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DecryptCombineVerifier, error) {
	parsed, err := DecryptCombineVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DecryptCombineVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DecryptCombineVerifier{DecryptCombineVerifierCaller: DecryptCombineVerifierCaller{contract: contract}, DecryptCombineVerifierTransactor: DecryptCombineVerifierTransactor{contract: contract}, DecryptCombineVerifierFilterer: DecryptCombineVerifierFilterer{contract: contract}}, nil
}

// DecryptCombineVerifier is an auto generated Go binding around an Ethereum contract.
type DecryptCombineVerifier struct {
	DecryptCombineVerifierCaller     // Read-only binding to the contract
	DecryptCombineVerifierTransactor // Write-only binding to the contract
	DecryptCombineVerifierFilterer   // Log filterer for contract events
}

// DecryptCombineVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type DecryptCombineVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecryptCombineVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DecryptCombineVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecryptCombineVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DecryptCombineVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecryptCombineVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DecryptCombineVerifierSession struct {
	Contract     *DecryptCombineVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DecryptCombineVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DecryptCombineVerifierCallerSession struct {
	Contract *DecryptCombineVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// DecryptCombineVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DecryptCombineVerifierTransactorSession struct {
	Contract     *DecryptCombineVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// DecryptCombineVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type DecryptCombineVerifierRaw struct {
	Contract *DecryptCombineVerifier // Generic contract binding to access the raw methods on
}

// DecryptCombineVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DecryptCombineVerifierCallerRaw struct {
	Contract *DecryptCombineVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// DecryptCombineVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DecryptCombineVerifierTransactorRaw struct {
	Contract *DecryptCombineVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDecryptCombineVerifier creates a new instance of DecryptCombineVerifier, bound to a specific deployed contract.
func NewDecryptCombineVerifier(address common.Address, backend bind.ContractBackend) (*DecryptCombineVerifier, error) {
	contract, err := bindDecryptCombineVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DecryptCombineVerifier{DecryptCombineVerifierCaller: DecryptCombineVerifierCaller{contract: contract}, DecryptCombineVerifierTransactor: DecryptCombineVerifierTransactor{contract: contract}, DecryptCombineVerifierFilterer: DecryptCombineVerifierFilterer{contract: contract}}, nil
}

// NewDecryptCombineVerifierCaller creates a new read-only instance of DecryptCombineVerifier, bound to a specific deployed contract.
func NewDecryptCombineVerifierCaller(address common.Address, caller bind.ContractCaller) (*DecryptCombineVerifierCaller, error) {
	contract, err := bindDecryptCombineVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DecryptCombineVerifierCaller{contract: contract}, nil
}

// NewDecryptCombineVerifierTransactor creates a new write-only instance of DecryptCombineVerifier, bound to a specific deployed contract.
func NewDecryptCombineVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*DecryptCombineVerifierTransactor, error) {
	contract, err := bindDecryptCombineVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DecryptCombineVerifierTransactor{contract: contract}, nil
}

// NewDecryptCombineVerifierFilterer creates a new log filterer instance of DecryptCombineVerifier, bound to a specific deployed contract.
func NewDecryptCombineVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*DecryptCombineVerifierFilterer, error) {
	contract, err := bindDecryptCombineVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DecryptCombineVerifierFilterer{contract: contract}, nil
}

// bindDecryptCombineVerifier binds a generic wrapper to an already deployed contract.
func bindDecryptCombineVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DecryptCombineVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DecryptCombineVerifier *DecryptCombineVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DecryptCombineVerifier.Contract.DecryptCombineVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DecryptCombineVerifier *DecryptCombineVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DecryptCombineVerifier.Contract.DecryptCombineVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DecryptCombineVerifier *DecryptCombineVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DecryptCombineVerifier.Contract.DecryptCombineVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DecryptCombineVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DecryptCombineVerifier *DecryptCombineVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DecryptCombineVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DecryptCombineVerifier *DecryptCombineVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DecryptCombineVerifier.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_DecryptCombineVerifier *DecryptCombineVerifierCaller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int) ([4]*big.Int, error) {
	var out []interface{}
	err := _DecryptCombineVerifier.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_DecryptCombineVerifier *DecryptCombineVerifierSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _DecryptCombineVerifier.Contract.CompressProof(&_DecryptCombineVerifier.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _DecryptCombineVerifier.Contract.CompressProof(&_DecryptCombineVerifier.CallOpts, proof)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_DecryptCombineVerifier *DecryptCombineVerifierCaller) ProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DecryptCombineVerifier.contract.Call(opts, &out, "provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_DecryptCombineVerifier *DecryptCombineVerifierSession) ProvingKeyHash() ([32]byte, error) {
	return _DecryptCombineVerifier.Contract.ProvingKeyHash(&_DecryptCombineVerifier.CallOpts)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerSession) ProvingKeyHash() ([32]byte, error) {
	return _DecryptCombineVerifier.Contract.ProvingKeyHash(&_DecryptCombineVerifier.CallOpts)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xf6d3293a.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[7] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [7]*big.Int) error {
	var out []interface{}
	err := _DecryptCombineVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xf6d3293a.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[7] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [7]*big.Int) error {
	return _DecryptCombineVerifier.Contract.VerifyCompressedProof(&_DecryptCombineVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xf6d3293a.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[7] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [7]*big.Int) error {
	return _DecryptCombineVerifier.Contract.VerifyCompressedProof(&_DecryptCombineVerifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x1042664e.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[7] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, input [7]*big.Int) error {
	var out []interface{}
	err := _DecryptCombineVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x1042664e.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[7] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierSession) VerifyProof(proof [8]*big.Int, input [7]*big.Int) error {
	return _DecryptCombineVerifier.Contract.VerifyProof(&_DecryptCombineVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x1042664e.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[7] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerSession) VerifyProof(proof [8]*big.Int, input [7]*big.Int) error {
	return _DecryptCombineVerifier.Contract.VerifyProof(&_DecryptCombineVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCaller) VerifyProof0(opts *bind.CallOpts, proof []byte, input []byte) error {
	var out []interface{}
	err := _DecryptCombineVerifier.contract.Call(opts, &out, "verifyProof0", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierSession) VerifyProof0(proof []byte, input []byte) error {
	return _DecryptCombineVerifier.Contract.VerifyProof0(&_DecryptCombineVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerSession) VerifyProof0(proof []byte, input []byte) error {
	return _DecryptCombineVerifier.Contract.VerifyProof0(&_DecryptCombineVerifier.CallOpts, proof, input)
}
