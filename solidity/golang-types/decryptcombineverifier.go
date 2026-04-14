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
	Bin: "0x6080806040523460155761160b908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80631042664e1461092b578063233ace11146108f157806344f6369214610858578063b8e72af6146106b65763f6d3293a14610051575f80fd5b346106b3576101603660031901126106b357366084116106b35736610164116106b357604051906103006100858184610f3b565b80368437610094600435611232565b6100a560249592953560443561129d565b919392906100b4606435611232565b9390926040519660408801967f264723aed10461372902b8cb59989a6db06363250cb194873f512a35810cd6ed89528860208101987f0b12cf75020b4246f4577d710cba877f879dea57cc19fbb0e67f948713deec288a527f26f3fbf28531fe55e68ebff96b4bfc35133770324c3ef5b240c67e4bfc0e1a3281527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f0a5e2420c2189cd7e1090f03599a3c4f19b4821b47b37f584e57da071c35807b84527f1f07706085dc0d2f0bcadbf4afca1ad5488ada07a99d53662fabd8fe4169caac6084359583608082019780895286828660608160075afa911016818360808160065afa167f2ac6ec179c6ae56536e4d1e680e2d483984492f9e0117c64655dde539f75b1da85527f14a434c4f617a04ae0290f688fa750a19f5eea6e0597aaf09d3f4db45b8260f68852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f20ebdbbd22060979e897f19e6cb284ff5eba2171dff28a7a33b34b139feb217d85527f26589095799acfb568a1891efb6dff4009725fc0a76c568e27c937762b52786e885260c43590818a5287838760608160075afa92101616818360808160065afa167f171eb4412f2f06bf8f412e70ed4aa824d35bf5934126153c2436adf150301bd285527f2c1f25118570dbbf9b5e9d078ac5a335fd7d4a47063dba4629c371540251e0bd885260e43590818a5287838760608160075afa92101616818360808160065afa167f1a6533270219b48afbc55f101501b49a85a5b7728a11b0eb2d07b390aab6722d85527f1caf582a2514437fbb3ef9d8a601299f3e9b65d063c6c5798a76722c827e60f088526101043590818a5287838760608160075afa92101616818360808160065afa167f2370893ba19861e3cc82002b345d52fe1d87cae2b45d0a0b26a2ddd748a1843685527f0aa921633f0389fb1c056a147cb664b1bb66c1932f081a4e8dae5c18c3b910dc88526101243590818a5287838760608160075afa921016169160808160065afa16947f07adafe13dcfa4543254d40a8be4f0f93f6218243080499036de530281fb6b238352526101443580955260608160075afa9210161660408a60808160065afa169851975198156106a45760209a9b9c8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f1dd5ffeff4605e89257aaf6435ec388474989f3fb1d9ba9d82b2daac035c97486101008401527f256a0a6fefaa97ace83877259d7fa49df504d723fb4535be9af0e4b7e594e8246101208401527f0e4948c49bbf2f00ad2622e2bc14e685e89da270905d84963a300fffcb0aff5f6101408401527f1c3715c405f5ad7e074061910aae03a28366898141bacefeb3f2cea2074532b76101608401527f1c32bb063a0d81f64144cfa9ff7c42cdd3850a93b2bf1137f220feddbe7f43d06101808401527f1fc29003abdcd9cff08b408d8e34f37d77f0098876e1a43c4778e52c1f892bc46101a08401527f1e021dac9841eca4f7c72b46c84d7a7979c4753a215baa01d86af165f60dbe766101c08401527f0a042b4acfbe5ce82bc110764ec88162605b68c963f1480ecda4655455804c4a6101e08401527f0d566ade2a9d3ebcdb922c2c363aae02013739953040815f5aa4a26005fcc8136102008401527f064366a0031bd2d2d24ebb00581d2d5d14023b766fe47242f0a6a472720b7be06102208401526102408301526102608201527f1685f51b6277fd5d9b567d7cf889305483a88c47160e6f95c0cae7ce4c2ad3046102808201527f01c822019e6abdb3ec4f37897d2bced2767e4f1d9d36a2c5ea0256b4fed0ed036102a08201527f04fed583f1925f04ac7ef095d1e02d93c0cfbc1454fa3880634ca4a26b36a8d06102c08201527f1027c41e36ca79686f4e378e4dbbf8c8fb97787fa91693ab712a5a4d1430934e6102e082015260405192839161066f8484610f3b565b8336843760085afa15908115610697575b506106885780f35b631ff3747d60e21b8152600490fd5b600191505114155f610680565b63a54f8e2760e01b8c5260048cfd5b80fd5b50346108205760403660031901126108205760043567ffffffffffffffff8111610820576106e8903690600401610f0d565b60243567ffffffffffffffff811161082057610708903690600401610f0d565b90918301610100848203126108205780601f85011215610820576040519361073261010086610f3b565b8490610100810192831161082057905b82821061084857505050810160e0828203126108205780601f83011215610820576040519161077260e084610f3b565b829060e0810192831161082057905b82821061082457505050303b1561082057604051630821332760e11b8152915f600484015b6008821061080a5750505061010482015f905b600782106107f4575050505f816101e481305afa80156107e9576107db575080f35b6107e791505f90610f3b565b005b6040513d5f823e3d90fd5b60208060019285518152019301910190916107b9565b60208060019285518152019301910190916107a6565b5f80fd5b8135815260209182019101610781565b634e487b7160e01b5f52604160045260245ffd5b8135815260209182019101610742565b346108205761010036600319011261082057366101041161082057604051610881608082610f3b565b6080368237610894602435600435610f5d565b81526108aa60843560a435604435606435610ffe565b602083015260408201526108c260e43560c435610f5d565b6060820152604051905f825b600482106108db57608084f35b60208060019285518152019301910190916108ce565b34610820575f3660031901126108205760206040517f43d858ca8532f6b67880bc3a61cd9e3a39b4db1d4d59828d3e84452b7dc76a4f8152f35b34610820576101e036600319011261082057366101041161082057366101e4116108205760405160408101907f264723aed10461372902b8cb59989a6db06363250cb194873f512a35810cd6ed815260208101917f0b12cf75020b4246f4577d710cba877f879dea57cc19fbb0e67f948713deec2883527f26f3fbf28531fe55e68ebff96b4bfc35133770324c3ef5b240c67e4bfc0e1a328152606082017f0a5e2420c2189cd7e1090f03599a3c4f19b4821b47b37f584e57da071c35807b81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f1f07706085dc0d2f0bcadbf4afca1ad5488ada07a99d53662fabd8fe4169caac608087019580875284848460608160075afa911016838860808160065afa167f2ac6ec179c6ae56536e4d1e680e2d483984492f9e0117c64655dde539f75b1da83527f14a434c4f617a04ae0290f688fa750a19f5eea6e0597aaf09d3f4db45b8260f686526001610124359182895286868660608160075afa9310161616838860808160065afa167f20ebdbbd22060979e897f19e6cb284ff5eba2171dff28a7a33b34b139feb217d83527f26589095799acfb568a1891efb6dff4009725fc0a76c568e27c937762b52786e8652610144359081885285858560608160075afa92101616838860808160065afa167f171eb4412f2f06bf8f412e70ed4aa824d35bf5934126153c2436adf150301bd283527f2c1f25118570dbbf9b5e9d078ac5a335fd7d4a47063dba4629c371540251e0bd8652610164359081885285858560608160075afa92101616838860808160065afa167f1a6533270219b48afbc55f101501b49a85a5b7728a11b0eb2d07b390aab6722d83527f1caf582a2514437fbb3ef9d8a601299f3e9b65d063c6c5798a76722c827e60f08652610184359081885285858560608160075afa92101616838860808160065afa167f2370893ba19861e3cc82002b345d52fe1d87cae2b45d0a0b26a2ddd748a1843683527f0aa921633f0389fb1c056a147cb664b1bb66c1932f081a4e8dae5c18c3b910dc86526101a4359081885285858560608160075afa92101616838860808160065afa16947f07adafe13dcfa4543254d40a8be4f0f93f6218243080499036de530281fb6b238352526101c43580955260608160075afa9210161660408260808160065afa16905191519015610efe5760405191610100600484377f1dd5ffeff4605e89257aaf6435ec388474989f3fb1d9ba9d82b2daac035c97486101008401527f256a0a6fefaa97ace83877259d7fa49df504d723fb4535be9af0e4b7e594e8246101208401527f0e4948c49bbf2f00ad2622e2bc14e685e89da270905d84963a300fffcb0aff5f6101408401527f1c3715c405f5ad7e074061910aae03a28366898141bacefeb3f2cea2074532b76101608401527f1c32bb063a0d81f64144cfa9ff7c42cdd3850a93b2bf1137f220feddbe7f43d06101808401527f1fc29003abdcd9cff08b408d8e34f37d77f0098876e1a43c4778e52c1f892bc46101a08401527f1e021dac9841eca4f7c72b46c84d7a7979c4753a215baa01d86af165f60dbe766101c08401527f0a042b4acfbe5ce82bc110764ec88162605b68c963f1480ecda4655455804c4a6101e08401527f0d566ade2a9d3ebcdb922c2c363aae02013739953040815f5aa4a26005fcc8136102008401527f064366a0031bd2d2d24ebb00581d2d5d14023b766fe47242f0a6a472720b7be06102208401526102408301526102608201527f1685f51b6277fd5d9b567d7cf889305483a88c47160e6f95c0cae7ce4c2ad3046102808201527f01c822019e6abdb3ec4f37897d2bced2767e4f1d9d36a2c5ea0256b4fed0ed036102a08201527f04fed583f1925f04ac7ef095d1e02d93c0cfbc1454fa3880634ca4a26b36a8d06102c08201527f1027c41e36ca79686f4e378e4dbbf8c8fb97787fa91693ab712a5a4d1430934e6102e08201526020816103008160085afa90511615610eef57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b9181601f840112156108205782359167ffffffffffffffff8311610820576020838186019501011161082057565b90601f8019910116810190811067ffffffffffffffff82111761083457604052565b905f5160206115b65f395f51905f528210801590610fe7575b610eef57811580610fdf575b610fd957610fa65f5160206115b65f395f51905f52600381858181800909086113d6565b818103610fb557505060011b90565b5f5160206115b65f395f51905f52809106810306145f14610eef57600190811b1790565b50505f90565b508015610f82565b505f5160206115b65f395f51905f52811015610f76565b919093925f5160206115b65f395f51905f52831080159061121b575b8015611204575b80156111ed575b610eef5780828685171717156111e2579082916111455f5160206115b65f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f5160206115b65f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161111f81808b800981878009086113d6565b8408095f5160206115b65f395f51905f526111398261154d565b800914159586916113f9565b9290808214806111d9575b156111775750505050905f1461116f5760ff60025b169060021b179190565b60ff5f611165565b5f5160206115b65f395f51905f528091068103061491826111ba575b505015610eef57600191156111b25760ff60025b169060021b17179190565b60ff5f6111a7565b5f5160206115b65f395f51905f52919250819006810306145f80611193565b50838314611150565b50505090505f905f90565b505f5160206115b65f395f51905f52811015611028565b505f5160206115b65f395f51905f52821015611021565b505f5160206115b65f395f51905f5285101561101a565b8015611296578060011c915f5160206115b65f395f51905f52831015610eef576001806112755f5160206115b65f395f51905f52600381888181800909086113d6565b93161461127e57565b905f5160206115b65f395f51905f5280910681030690565b505f905f90565b8015806113ce575b6113c2578060021c92825f5160206115b65f395f51905f5285108015906113ab575b610eef5784815f5160206115b65f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816113759d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5086113f9565b80929160018082961614611387575050565b5f5160206115b65f395f51905f528093945080929550809106810306930681030690565b505f5160206115b65f395f51905f528110156112c7565b50505f905f905f905f90565b5081156112a5565b906113e08261154d565b915f5160206115b65f395f51905f5283800903610eef57565b915f5160206115b65f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816114519396949661144382808a8009818a8009086113d6565b90611541575b8608096113d6565b925f5160206115b65f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206115b65f395f51905f5260a083015260208260c08160055afa91519115610eef575f5160206115b65f395f51905f52826001920903610eef575f5160206115b65f395f51905f52908209925f5160206115b65f395f51905f528080808780090681030681878009081490811591611522575b50610eef57565b90505f5160206115b65f395f51905f528084860960020914155f61151b565b81809106810306611449565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206115b65f395f51905f5260a083015260208260c08160055afa91519115610eef5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a264697066735822122080a643c21025145074512829077239cbc0f30756e8e0aeac3ba188dc2bbf16f464736f6c634300081c0033",
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
