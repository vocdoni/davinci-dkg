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
	Bin: "0x6080806040523460155761160b908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80631042664e1461092b578063233ace11146108f157806344f6369214610858578063b8e72af6146106b65763f6d3293a14610051575f80fd5b346106b3576101603660031901126106b357366084116106b35736610164116106b357604051906103006100858184610f3b565b80368437610094600435611232565b6100a560249592953560443561129d565b919392906100b4606435611232565b9390926040519660408801967f29091a7cbe7b2b261b5ef38972789a8ee371e6741b2f12b5ccf10e86a2204ad689528860208101987f2beadf5f4ce81aa2124c69e88be39ec4a6053d091969791dd9cb17458cf20d0e8a527f29fe65379c8906fbbb84f8afea88e9c8e9e11aa649d6a700eaeb2628179ac4ed81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f2f4e25bb967625e37c4c5b001d48273831bafda9779df3c0ef53eeeccf46bac784527f16de5411ae2d94320fc7edce61a5d876ef937a9c74891ea60771d085192b15056084359583608082019780895286828660608160075afa911016818360808160065afa167f1f245e93b16e3d131cbd7623d1961d27135369b36445f0c0923c98438670c47f85527f2dd8c95f1e479329051cb164a14d7cb32b2497b5fe0c4312554cd85e502cc2fb8852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f05601ad8beaf2cd5a1a9d0687de823acc236ca45f0a44849998f3d81efbe5ec685527f170ba708093e301b1597c394a5f65f8d376ad968415a7068c613abf50a563ee4885260c43590818a5287838760608160075afa92101616818360808160065afa167f16628b31301b1600a69d968d936bd7d47a69feac18c48c9bf065d6d423248d2f85527f12b435f79b86bfd7b857da5c98df2ece09303ec48cb6237dd2339a0c41d91efc885260e43590818a5287838760608160075afa92101616818360808160065afa167f1173cec5f60153e4451f7465e687d7714f96ba46aab0883e36a7c08eb6ae34ff85527f0913b2933797c4a56ecda4677439776f2c8e6d674402742d9600b8dfca1486c888526101043590818a5287838760608160075afa92101616818360808160065afa167f29b1af23ed8bb8fb1dad4101c35cffcd641a3aca3273942fb24b2c6427bb17bc85527f2ca0b0600bf8edc870565eae25721d4c5695484cca3b0d5c62f123a1f3dac99c88526101243590818a5287838760608160075afa921016169160808160065afa16947f1cf42e5e0f3d84f1c62ecfc0d1efedd5b21876067ca32fd69411395671fccb698352526101443580955260608160075afa9210161660408a60808160065afa169851975198156106a45760209a9b9c8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f053a12d5591a91f730a875c251daf4deab6593d8f504afe08db2d44d40814c5f6101008401527f199f3fe1a02d3f345d61fa3c037b831ca513bf20902f84c4becdeb97f3cf650b6101208401527f2ec7a3ac59be069c2d3e933faa7b9a45d7932a995c3a3dc7730db88fe79067036101408401527f0e177eb2c38a684ecc14af54be36952d2a8d482e22417ab26d4c3eebeead6ee06101608401527f111b5755850332e2a5a47d6d87379392c804e9f97bfc2119cf0f1c758c61f7806101808401527f0ff1096d05e226b41874788fd0957b05b9174bf0d2b227c0025d4dc4de32bac06101a08401527f184c2695e92c429c31dec4682aef9a5553851ba38abc1b41817f55d97cfd1ecd6101c08401527f0e2c668b3c87c8c254ff96f08ccfbd872cc754991cfad691e19b37206dfb2bef6101e08401527f23d3c03debcea6d78a90158bb769a6b019fab1247e07e788356cb9cbe027325d6102008401527f29da952b9301f0fd636381403e3ae6991815ad4b650463f10c49f5a340040cba6102208401526102408301526102608201527f249916fb3dcb13fe06c43b5a8541961b5189f1be2f0194116015cb1a2afb13b36102808201527f150d2a3658bca769187f6a6c7565fd1b9c83dd06fcfdefa4fccdc0f9c5c7924a6102a08201527f1bd96b8a912303db6fd710d6f2388eb00f0c3d22430e01ccdd0d03d0004e6c756102c08201527f2c00feaad35b9e1f0a56dcb225b9a28eac6172e226185daf0eff6b411b1e90636102e082015260405192839161066f8484610f3b565b8336843760085afa15908115610697575b506106885780f35b631ff3747d60e21b8152600490fd5b600191505114155f610680565b63a54f8e2760e01b8c5260048cfd5b80fd5b50346108205760403660031901126108205760043567ffffffffffffffff8111610820576106e8903690600401610f0d565b60243567ffffffffffffffff811161082057610708903690600401610f0d565b90918301610100848203126108205780601f85011215610820576040519361073261010086610f3b565b8490610100810192831161082057905b82821061084857505050810160e0828203126108205780601f83011215610820576040519161077260e084610f3b565b829060e0810192831161082057905b82821061082457505050303b1561082057604051630821332760e11b8152915f600484015b6008821061080a5750505061010482015f905b600782106107f4575050505f816101e481305afa80156107e9576107db575080f35b6107e791505f90610f3b565b005b6040513d5f823e3d90fd5b60208060019285518152019301910190916107b9565b60208060019285518152019301910190916107a6565b5f80fd5b8135815260209182019101610781565b634e487b7160e01b5f52604160045260245ffd5b8135815260209182019101610742565b346108205761010036600319011261082057366101041161082057604051610881608082610f3b565b6080368237610894602435600435610f5d565b81526108aa60843560a435604435606435610ffe565b602083015260408201526108c260e43560c435610f5d565b6060820152604051905f825b600482106108db57608084f35b60208060019285518152019301910190916108ce565b34610820575f3660031901126108205760206040517f4e1ae67d1bb1e750c48fa01e0f40783926027cd66180d9e72f8d9eed6d99cd368152f35b34610820576101e036600319011261082057366101041161082057366101e4116108205760405160408101907f29091a7cbe7b2b261b5ef38972789a8ee371e6741b2f12b5ccf10e86a2204ad6815260208101917f2beadf5f4ce81aa2124c69e88be39ec4a6053d091969791dd9cb17458cf20d0e83527f29fe65379c8906fbbb84f8afea88e9c8e9e11aa649d6a700eaeb2628179ac4ed8152606082017f2f4e25bb967625e37c4c5b001d48273831bafda9779df3c0ef53eeeccf46bac781527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f16de5411ae2d94320fc7edce61a5d876ef937a9c74891ea60771d085192b1505608087019580875284848460608160075afa911016838860808160065afa167f1f245e93b16e3d131cbd7623d1961d27135369b36445f0c0923c98438670c47f83527f2dd8c95f1e479329051cb164a14d7cb32b2497b5fe0c4312554cd85e502cc2fb86526001610124359182895286868660608160075afa9310161616838860808160065afa167f05601ad8beaf2cd5a1a9d0687de823acc236ca45f0a44849998f3d81efbe5ec683527f170ba708093e301b1597c394a5f65f8d376ad968415a7068c613abf50a563ee48652610144359081885285858560608160075afa92101616838860808160065afa167f16628b31301b1600a69d968d936bd7d47a69feac18c48c9bf065d6d423248d2f83527f12b435f79b86bfd7b857da5c98df2ece09303ec48cb6237dd2339a0c41d91efc8652610164359081885285858560608160075afa92101616838860808160065afa167f1173cec5f60153e4451f7465e687d7714f96ba46aab0883e36a7c08eb6ae34ff83527f0913b2933797c4a56ecda4677439776f2c8e6d674402742d9600b8dfca1486c88652610184359081885285858560608160075afa92101616838860808160065afa167f29b1af23ed8bb8fb1dad4101c35cffcd641a3aca3273942fb24b2c6427bb17bc83527f2ca0b0600bf8edc870565eae25721d4c5695484cca3b0d5c62f123a1f3dac99c86526101a4359081885285858560608160075afa92101616838860808160065afa16947f1cf42e5e0f3d84f1c62ecfc0d1efedd5b21876067ca32fd69411395671fccb698352526101c43580955260608160075afa9210161660408260808160065afa16905191519015610efe5760405191610100600484377f053a12d5591a91f730a875c251daf4deab6593d8f504afe08db2d44d40814c5f6101008401527f199f3fe1a02d3f345d61fa3c037b831ca513bf20902f84c4becdeb97f3cf650b6101208401527f2ec7a3ac59be069c2d3e933faa7b9a45d7932a995c3a3dc7730db88fe79067036101408401527f0e177eb2c38a684ecc14af54be36952d2a8d482e22417ab26d4c3eebeead6ee06101608401527f111b5755850332e2a5a47d6d87379392c804e9f97bfc2119cf0f1c758c61f7806101808401527f0ff1096d05e226b41874788fd0957b05b9174bf0d2b227c0025d4dc4de32bac06101a08401527f184c2695e92c429c31dec4682aef9a5553851ba38abc1b41817f55d97cfd1ecd6101c08401527f0e2c668b3c87c8c254ff96f08ccfbd872cc754991cfad691e19b37206dfb2bef6101e08401527f23d3c03debcea6d78a90158bb769a6b019fab1247e07e788356cb9cbe027325d6102008401527f29da952b9301f0fd636381403e3ae6991815ad4b650463f10c49f5a340040cba6102208401526102408301526102608201527f249916fb3dcb13fe06c43b5a8541961b5189f1be2f0194116015cb1a2afb13b36102808201527f150d2a3658bca769187f6a6c7565fd1b9c83dd06fcfdefa4fccdc0f9c5c7924a6102a08201527f1bd96b8a912303db6fd710d6f2388eb00f0c3d22430e01ccdd0d03d0004e6c756102c08201527f2c00feaad35b9e1f0a56dcb225b9a28eac6172e226185daf0eff6b411b1e90636102e08201526020816103008160085afa90511615610eef57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b9181601f840112156108205782359167ffffffffffffffff8311610820576020838186019501011161082057565b90601f8019910116810190811067ffffffffffffffff82111761083457604052565b905f5160206115b65f395f51905f528210801590610fe7575b610eef57811580610fdf575b610fd957610fa65f5160206115b65f395f51905f52600381858181800909086113d6565b818103610fb557505060011b90565b5f5160206115b65f395f51905f52809106810306145f14610eef57600190811b1790565b50505f90565b508015610f82565b505f5160206115b65f395f51905f52811015610f76565b919093925f5160206115b65f395f51905f52831080159061121b575b8015611204575b80156111ed575b610eef5780828685171717156111e2579082916111455f5160206115b65f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f5160206115b65f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161111f81808b800981878009086113d6565b8408095f5160206115b65f395f51905f526111398261154d565b800914159586916113f9565b9290808214806111d9575b156111775750505050905f1461116f5760ff60025b169060021b179190565b60ff5f611165565b5f5160206115b65f395f51905f528091068103061491826111ba575b505015610eef57600191156111b25760ff60025b169060021b17179190565b60ff5f6111a7565b5f5160206115b65f395f51905f52919250819006810306145f80611193565b50838314611150565b50505090505f905f90565b505f5160206115b65f395f51905f52811015611028565b505f5160206115b65f395f51905f52821015611021565b505f5160206115b65f395f51905f5285101561101a565b8015611296578060011c915f5160206115b65f395f51905f52831015610eef576001806112755f5160206115b65f395f51905f52600381888181800909086113d6565b93161461127e57565b905f5160206115b65f395f51905f5280910681030690565b505f905f90565b8015806113ce575b6113c2578060021c92825f5160206115b65f395f51905f5285108015906113ab575b610eef5784815f5160206115b65f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816113759d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5086113f9565b80929160018082961614611387575050565b5f5160206115b65f395f51905f528093945080929550809106810306930681030690565b505f5160206115b65f395f51905f528110156112c7565b50505f905f905f905f90565b5081156112a5565b906113e08261154d565b915f5160206115b65f395f51905f5283800903610eef57565b915f5160206115b65f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816114519396949661144382808a8009818a8009086113d6565b90611541575b8608096113d6565b925f5160206115b65f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206115b65f395f51905f5260a083015260208260c08160055afa91519115610eef575f5160206115b65f395f51905f52826001920903610eef575f5160206115b65f395f51905f52908209925f5160206115b65f395f51905f528080808780090681030681878009081490811591611522575b50610eef57565b90505f5160206115b65f395f51905f528084860960020914155f61151b565b81809106810306611449565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206115b65f395f51905f5260a083015260208260c08160055afa91519115610eef5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220aa9fe5ddc6669c6ec69ea3d5bb050ecf6f373c06a9e4a1193f67f622a36e6e7c64736f6c634300081c0033",
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
