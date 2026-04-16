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
	Bin: "0x60808060405234601557611609908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80631042664e1461092a578063233ace11146108f057806344f6369214610857578063b8e72af6146106b55763f6d3293a14610051575f80fd5b346106b2576101603660031901126106b257366084116106b25736610164116106b257604051906103006100858184610f39565b80368437610094600435611230565b6100a560249592953560443561129b565b919392906100b4606435611230565b9390926040519660408801967f22616cb87d911a26c3849e8bf9d7d737e22c5740079210544ea637f6b7a69a7f89528860208101987f293c52ad9eeba39035640ff607c2ef6651296aca06210d090c0f1874758b664c8a527f2048a82eb085f9c3256416d422de5e2b9adf11ed6da165ce134c846b63affd4981527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f21dc619f20cf03ed31f07ffe588d9d3e7a44d3f88cda556c3725f7f98b1eecac84527f07dc7ad79a9808a927a9a40619cb2c0bdba69b5d3014764ec4cf2904eeae05276084359583608082019780895286828660608160075afa911016818360808160065afa167f0e837b3a0cd200f1a6bd87ca9acb5882f620a947f7be0b3c8087c96a706ffc0585527f0cba852403c9cc46927532811d70bf51856b1e66596ff1f1915c3e7a3446c7f18852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f01362ef886dee3933a1e08e484dae4106d03d328347ac55b61de8f04f586641185527f1c5d342dc2098598caac3f797d25946a85f18e1e8600658477da568cb9daa40d885260c43590818a5287838760608160075afa92101616818360808160065afa167f29bffa9ff66845f961f1bf3a8d3c4240b4d18967cbfc53d667acb2e65fc159d585527f3061fa9b45eb1330feac45f16e79576ccfa5f13aea16d1b463176df8bdd0385b885260e43590818a5287838760608160075afa92101616818360808160065afa167f28a39e89384c3c2a5bf8caae458c0ddfa79a88856b648740c85df5e4ff101d8d85527f032e520f93e514c2b5dd296b58dda2483179d18392aa91230a4898a6a7a67fef88526101043590818a5287838760608160075afa92101616818360808160065afa167f2c84f1024b74af6e3f9365ead3a49fb30d63e1ae7eba0e49f397562ffff19cac85527f2fc47059e7b8929c7193d1b74adc7de90329e3c462bd4894a0c18a8c839ee3a188526101243590818a5287838760608160075afa921016169160808160065afa16947f24f5014d7c18eac60f6bc59c0c4307ecd286757dadd00b1477b3faee7de62c658352526101443580955260608160075afa9210161660408a60808160065afa169851975198156106a35760209a9b9c8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f2beeb11ae49ec94b82be8791950ff00f8cf64c6b64cb9f97cbfd67ac812557156101008401527f243ec0b54c4bf1708d44273a3cb5cf6bc4f56cbc9bb1da523585d0e94d184aeb6101208401527f2050ff5ce48ad21ad8cdb6a6f8c2bf6701a69ab92b6bf09e18c743bfd1d499666101408401527f2eb0f5e41eb8a0e77c6f52a685f4f067352473826c3fb43dec2e13660a58fed86101608401527f26cfbcebc50eec2452b53845a8dd427aeec670cb3f91b398c6515a68550db2ce6101808401527f0422da97f7a9051d3c5e2da00de6c824cd5b5fb4c865dce6cc92ce2ce2914c726101a08401527f15fa15bb2ce01c708d8c9433489f5fb785d03d322b59682a9ba4ff33a42799f86101c08401527f035e59d89bd93de2a0facc1d25f7844b83ca6882eaf2925fafa7c4a7079f00aa6101e08401527e43d646f51b8a54e280bde95dfe30316573d70e16a495fb4f660dc746a1ecff6102008401527f0ee2ec4d55c37351e61a320da4e70938900589aabbcfe06440039361a19342686102208401526102408301526102608201527f0e8aeed7d0d859fe6e4f04b6a825feb0dc9dbc157c1507c549019ee94b3242eb6102808201527f1647bc2998853e647cc3cd33bb882015a68f88664fc710fc4a1435f508feb49a6102a08201527f0ecd6e6cc10a3df47a9bd346f371590bd595aa30c5aef352e302452c038d6bf86102c08201527f024eb1f6b485c4045dc9e5a9acc33f47911a70dd460a8df0286eb11382abfb266102e082015260405192839161066e8484610f39565b8336843760085afa15908115610696575b506106875780f35b631ff3747d60e21b8152600490fd5b600191505114155f61067f565b63a54f8e2760e01b8c5260048cfd5b80fd5b503461081f57604036600319011261081f5760043567ffffffffffffffff811161081f576106e7903690600401610f0b565b60243567ffffffffffffffff811161081f57610707903690600401610f0b565b909183016101008482031261081f5780601f8501121561081f576040519361073161010086610f39565b8490610100810192831161081f57905b82821061084757505050810160e08282031261081f5780601f8301121561081f576040519161077160e084610f39565b829060e0810192831161081f57905b82821061082357505050303b1561081f57604051630821332760e11b8152915f600484015b600882106108095750505061010482015f905b600782106107f3575050505f816101e481305afa80156107e8576107da575080f35b6107e691505f90610f39565b005b6040513d5f823e3d90fd5b60208060019285518152019301910190916107b8565b60208060019285518152019301910190916107a5565b5f80fd5b8135815260209182019101610780565b634e487b7160e01b5f52604160045260245ffd5b8135815260209182019101610741565b3461081f5761010036600319011261081f57366101041161081f57604051610880608082610f39565b6080368237610893602435600435610f5b565b81526108a960843560a435604435606435610ffc565b602083015260408201526108c160e43560c435610f5b565b6060820152604051905f825b600482106108da57608084f35b60208060019285518152019301910190916108cd565b3461081f575f36600319011261081f5760206040517f43d858ca8532f6b67880bc3a61cd9e3a39b4db1d4d59828d3e84452b7dc76a4f8152f35b3461081f576101e036600319011261081f57366101041161081f57366101e41161081f5760405160408101907f22616cb87d911a26c3849e8bf9d7d737e22c5740079210544ea637f6b7a69a7f815260208101917f293c52ad9eeba39035640ff607c2ef6651296aca06210d090c0f1874758b664c83527f2048a82eb085f9c3256416d422de5e2b9adf11ed6da165ce134c846b63affd498152606082017f21dc619f20cf03ed31f07ffe588d9d3e7a44d3f88cda556c3725f7f98b1eecac81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f07dc7ad79a9808a927a9a40619cb2c0bdba69b5d3014764ec4cf2904eeae0527608087019580875284848460608160075afa911016838860808160065afa167f0e837b3a0cd200f1a6bd87ca9acb5882f620a947f7be0b3c8087c96a706ffc0583527f0cba852403c9cc46927532811d70bf51856b1e66596ff1f1915c3e7a3446c7f186526001610124359182895286868660608160075afa9310161616838860808160065afa167f01362ef886dee3933a1e08e484dae4106d03d328347ac55b61de8f04f586641183527f1c5d342dc2098598caac3f797d25946a85f18e1e8600658477da568cb9daa40d8652610144359081885285858560608160075afa92101616838860808160065afa167f29bffa9ff66845f961f1bf3a8d3c4240b4d18967cbfc53d667acb2e65fc159d583527f3061fa9b45eb1330feac45f16e79576ccfa5f13aea16d1b463176df8bdd0385b8652610164359081885285858560608160075afa92101616838860808160065afa167f28a39e89384c3c2a5bf8caae458c0ddfa79a88856b648740c85df5e4ff101d8d83527f032e520f93e514c2b5dd296b58dda2483179d18392aa91230a4898a6a7a67fef8652610184359081885285858560608160075afa92101616838860808160065afa167f2c84f1024b74af6e3f9365ead3a49fb30d63e1ae7eba0e49f397562ffff19cac83527f2fc47059e7b8929c7193d1b74adc7de90329e3c462bd4894a0c18a8c839ee3a186526101a4359081885285858560608160075afa92101616838860808160065afa16947f24f5014d7c18eac60f6bc59c0c4307ecd286757dadd00b1477b3faee7de62c658352526101c43580955260608160075afa9210161660408260808160065afa16905191519015610efc5760405191610100600484377f2beeb11ae49ec94b82be8791950ff00f8cf64c6b64cb9f97cbfd67ac812557156101008401527f243ec0b54c4bf1708d44273a3cb5cf6bc4f56cbc9bb1da523585d0e94d184aeb6101208401527f2050ff5ce48ad21ad8cdb6a6f8c2bf6701a69ab92b6bf09e18c743bfd1d499666101408401527f2eb0f5e41eb8a0e77c6f52a685f4f067352473826c3fb43dec2e13660a58fed86101608401527f26cfbcebc50eec2452b53845a8dd427aeec670cb3f91b398c6515a68550db2ce6101808401527f0422da97f7a9051d3c5e2da00de6c824cd5b5fb4c865dce6cc92ce2ce2914c726101a08401527f15fa15bb2ce01c708d8c9433489f5fb785d03d322b59682a9ba4ff33a42799f86101c08401527f035e59d89bd93de2a0facc1d25f7844b83ca6882eaf2925fafa7c4a7079f00aa6101e08401527e43d646f51b8a54e280bde95dfe30316573d70e16a495fb4f660dc746a1ecff6102008401527f0ee2ec4d55c37351e61a320da4e70938900589aabbcfe06440039361a19342686102208401526102408301526102608201527f0e8aeed7d0d859fe6e4f04b6a825feb0dc9dbc157c1507c549019ee94b3242eb6102808201527f1647bc2998853e647cc3cd33bb882015a68f88664fc710fc4a1435f508feb49a6102a08201527f0ecd6e6cc10a3df47a9bd346f371590bd595aa30c5aef352e302452c038d6bf86102c08201527f024eb1f6b485c4045dc9e5a9acc33f47911a70dd460a8df0286eb11382abfb266102e08201526020816103008160085afa90511615610eed57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b9181601f8401121561081f5782359167ffffffffffffffff831161081f576020838186019501011161081f57565b90601f8019910116810190811067ffffffffffffffff82111761083357604052565b905f5160206115b45f395f51905f528210801590610fe5575b610eed57811580610fdd575b610fd757610fa45f5160206115b45f395f51905f52600381858181800909086113d4565b818103610fb357505060011b90565b5f5160206115b45f395f51905f52809106810306145f14610eed57600190811b1790565b50505f90565b508015610f80565b505f5160206115b45f395f51905f52811015610f74565b919093925f5160206115b45f395f51905f528310801590611219575b8015611202575b80156111eb575b610eed5780828685171717156111e0579082916111435f5160206115b45f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f5160206115b45f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161111d81808b800981878009086113d4565b8408095f5160206115b45f395f51905f526111378261154b565b800914159586916113f7565b9290808214806111d7575b156111755750505050905f1461116d5760ff60025b169060021b179190565b60ff5f611163565b5f5160206115b45f395f51905f528091068103061491826111b8575b505015610eed57600191156111b05760ff60025b169060021b17179190565b60ff5f6111a5565b5f5160206115b45f395f51905f52919250819006810306145f80611191565b5083831461114e565b50505090505f905f90565b505f5160206115b45f395f51905f52811015611026565b505f5160206115b45f395f51905f5282101561101f565b505f5160206115b45f395f51905f52851015611018565b8015611294578060011c915f5160206115b45f395f51905f52831015610eed576001806112735f5160206115b45f395f51905f52600381888181800909086113d4565b93161461127c57565b905f5160206115b45f395f51905f5280910681030690565b505f905f90565b8015806113cc575b6113c0578060021c92825f5160206115b45f395f51905f5285108015906113a9575b610eed5784815f5160206115b45f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816113739d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5086113f7565b80929160018082961614611385575050565b5f5160206115b45f395f51905f528093945080929550809106810306930681030690565b505f5160206115b45f395f51905f528110156112c5565b50505f905f905f905f90565b5081156112a3565b906113de8261154b565b915f5160206115b45f395f51905f5283800903610eed57565b915f5160206115b45f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161144f9396949661144182808a8009818a8009086113d4565b9061153f575b8608096113d4565b925f5160206115b45f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206115b45f395f51905f5260a083015260208260c08160055afa91519115610eed575f5160206115b45f395f51905f52826001920903610eed575f5160206115b45f395f51905f52908209925f5160206115b45f395f51905f528080808780090681030681878009081490811591611520575b50610eed57565b90505f5160206115b45f395f51905f528084860960020914155f611519565b81809106810306611447565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206115b45f395f51905f5260a083015260208260c08160055afa91519115610eed5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220e965ca8ad8755a2237222d1a885b7b1fd95a263c2fbdb879b881de0884e2e83664736f6c634300081c0033",
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
