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
	Bin: "0x6080806040523460155761160b908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80631042664e1461092b578063233ace11146108f157806344f6369214610858578063b8e72af6146106b65763f6d3293a14610051575f80fd5b346106b3576101603660031901126106b357366084116106b35736610164116106b357604051906103006100858184610f3b565b80368437610094600435611232565b6100a560249592953560443561129d565b919392906100b4606435611232565b9390926040519660408801967f24ff501db207556ddb380bc6d992efe6e2ca8230d07ade3619f8ed43ef42263689528860208101987f1c8bcc960d75397071d60556f8193db9abaff6598be4052387be46ee6be31c618a527f1769df031c511b8592fa1b00a3b8954622fdd856612f3d127dd55d52145015e081527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f0cb2c71f7667c9f97c15fea5d080c7d36ef5dfa468fcfd3e42e3605243bf832884527f204f94caeae196fbd190ebc9bb8644f2f990e90c92f14b9e084ae3f794be2e0f6084359583608082019780895286828660608160075afa911016818360808160065afa167f02ea109ddaa8ac9ab1c0af64bd47ab26ec6be8da830b190cd1cb6eb16e5ab79185527f037711161bc3e32a8e0b81315c99a61f7cacd851f4cd1c885f392d1740243e5f8852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f23c29e914522ac7804ef5178d37675ff9aeef6270e469cf712115a778bd97abb85527f127d3da9ac231282e1f2ee969bb0fc55235999ccda6508d898de34e5e0080f91885260c43590818a5287838760608160075afa92101616818360808160065afa167f21659e448daf8b1e13d654a99b13a7551992ee0adce3097d5de6e456c1cf78b085527f1577d33a5e2be6fcb2681ad1d6263b73329022fbe7e26649921e06d4facf76f4885260e43590818a5287838760608160075afa92101616818360808160065afa167f03928412d2f839c886dc1cffc4f6bb4d1cdea524a22e7c19e2053ece09bcca3a85527f06d037f674e658c93cb787cecfa10118a02459f8e668efbf3d38158a23b560e188526101043590818a5287838760608160075afa92101616818360808160065afa167f0d30b210233dce062ae116aa6d7c3a6c3b6827afd5f3c4a5f7bf04356a22070385527f0dca20eaf8a2fd6419a73de9c0da63a4c5694f9aa7cb032dab993ad670770e7c88526101243590818a5287838760608160075afa921016169160808160065afa16947f06e63f0e5eb23e8340e0a4033461235bc30a13ad79e6a6ee156e87e276faa5af8352526101443580955260608160075afa9210161660408a60808160065afa169851975198156106a45760209a9b9c8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f1b957fb6f2e0fdf1b7c0b0802df9df99219fe090a7c973436000c5782b5e37a86101008401527f1d9b8592b01afbbaefbe443a024ccb87ec7abaf98e9ce3655c646b0936d498596101208401527f11f2aa12b04adfc2c93b89dc05cf4cffbbfd8b0bbfc00c86fa6d536cb6e5c9e36101408401527f0de3a0e09bc1c585acac5eba06fe37f7987fea4f31982736186632b775c1a5ce6101608401527f259013a40df170674fa0fefdb851aa9619d5f394b81f859a66c98f498e2f3ca46101808401527f23f03a145e74e81d1d1b7fd44b8ef98474cc8e98e46bee468f4f6ae67d4817406101a08401527f2a5f5936051aa2987b38f2c814b7399f2b50fe2e4d4f00b564ab0eee5b355de16101c08401527f0c3607ebb312b3292183be91f611c9fa41da20d9b8c26028be2b0ba16d8c2c836101e08401527f13f88fc4517c7fef4fa01daa9d56ba95da21a368270356eec401fe0b65bbaeba6102008401527f2ac8cc9e491532f6d70b6e89eee1c79ab825f80f577359901c6f34d88ee5fee66102208401526102408301526102608201527f20910a76f4aee5f0412948d472ba4a9396998e6cbbda72908585d37503a269336102808201527f158a7dce4fbc834dde5705059b7167e554566fdd51864c7508ebeafc68eaa6ed6102a08201527f0bd059852f29b878691531bbc4c5b4f3d6a0842645825fa0cbdc411e639dec446102c08201527f2f74c3ae7c9185ad5e7674108b5b452e34f3f6efe32b6c9f3c8e0cf8386bacaf6102e082015260405192839161066f8484610f3b565b8336843760085afa15908115610697575b506106885780f35b631ff3747d60e21b8152600490fd5b600191505114155f610680565b63a54f8e2760e01b8c5260048cfd5b80fd5b50346108205760403660031901126108205760043567ffffffffffffffff8111610820576106e8903690600401610f0d565b60243567ffffffffffffffff811161082057610708903690600401610f0d565b90918301610100848203126108205780601f85011215610820576040519361073261010086610f3b565b8490610100810192831161082057905b82821061084857505050810160e0828203126108205780601f83011215610820576040519161077260e084610f3b565b829060e0810192831161082057905b82821061082457505050303b1561082057604051630821332760e11b8152915f600484015b6008821061080a5750505061010482015f905b600782106107f4575050505f816101e481305afa80156107e9576107db575080f35b6107e791505f90610f3b565b005b6040513d5f823e3d90fd5b60208060019285518152019301910190916107b9565b60208060019285518152019301910190916107a6565b5f80fd5b8135815260209182019101610781565b634e487b7160e01b5f52604160045260245ffd5b8135815260209182019101610742565b346108205761010036600319011261082057366101041161082057604051610881608082610f3b565b6080368237610894602435600435610f5d565b81526108aa60843560a435604435606435610ffe565b602083015260408201526108c260e43560c435610f5d565b6060820152604051905f825b600482106108db57608084f35b60208060019285518152019301910190916108ce565b34610820575f3660031901126108205760206040517f4e1ae67d1bb1e750c48fa01e0f40783926027cd66180d9e72f8d9eed6d99cd368152f35b34610820576101e036600319011261082057366101041161082057366101e4116108205760405160408101907f24ff501db207556ddb380bc6d992efe6e2ca8230d07ade3619f8ed43ef422636815260208101917f1c8bcc960d75397071d60556f8193db9abaff6598be4052387be46ee6be31c6183527f1769df031c511b8592fa1b00a3b8954622fdd856612f3d127dd55d52145015e08152606082017f0cb2c71f7667c9f97c15fea5d080c7d36ef5dfa468fcfd3e42e3605243bf832881527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f204f94caeae196fbd190ebc9bb8644f2f990e90c92f14b9e084ae3f794be2e0f608087019580875284848460608160075afa911016838860808160065afa167f02ea109ddaa8ac9ab1c0af64bd47ab26ec6be8da830b190cd1cb6eb16e5ab79183527f037711161bc3e32a8e0b81315c99a61f7cacd851f4cd1c885f392d1740243e5f86526001610124359182895286868660608160075afa9310161616838860808160065afa167f23c29e914522ac7804ef5178d37675ff9aeef6270e469cf712115a778bd97abb83527f127d3da9ac231282e1f2ee969bb0fc55235999ccda6508d898de34e5e0080f918652610144359081885285858560608160075afa92101616838860808160065afa167f21659e448daf8b1e13d654a99b13a7551992ee0adce3097d5de6e456c1cf78b083527f1577d33a5e2be6fcb2681ad1d6263b73329022fbe7e26649921e06d4facf76f48652610164359081885285858560608160075afa92101616838860808160065afa167f03928412d2f839c886dc1cffc4f6bb4d1cdea524a22e7c19e2053ece09bcca3a83527f06d037f674e658c93cb787cecfa10118a02459f8e668efbf3d38158a23b560e18652610184359081885285858560608160075afa92101616838860808160065afa167f0d30b210233dce062ae116aa6d7c3a6c3b6827afd5f3c4a5f7bf04356a22070383527f0dca20eaf8a2fd6419a73de9c0da63a4c5694f9aa7cb032dab993ad670770e7c86526101a4359081885285858560608160075afa92101616838860808160065afa16947f06e63f0e5eb23e8340e0a4033461235bc30a13ad79e6a6ee156e87e276faa5af8352526101c43580955260608160075afa9210161660408260808160065afa16905191519015610efe5760405191610100600484377f1b957fb6f2e0fdf1b7c0b0802df9df99219fe090a7c973436000c5782b5e37a86101008401527f1d9b8592b01afbbaefbe443a024ccb87ec7abaf98e9ce3655c646b0936d498596101208401527f11f2aa12b04adfc2c93b89dc05cf4cffbbfd8b0bbfc00c86fa6d536cb6e5c9e36101408401527f0de3a0e09bc1c585acac5eba06fe37f7987fea4f31982736186632b775c1a5ce6101608401527f259013a40df170674fa0fefdb851aa9619d5f394b81f859a66c98f498e2f3ca46101808401527f23f03a145e74e81d1d1b7fd44b8ef98474cc8e98e46bee468f4f6ae67d4817406101a08401527f2a5f5936051aa2987b38f2c814b7399f2b50fe2e4d4f00b564ab0eee5b355de16101c08401527f0c3607ebb312b3292183be91f611c9fa41da20d9b8c26028be2b0ba16d8c2c836101e08401527f13f88fc4517c7fef4fa01daa9d56ba95da21a368270356eec401fe0b65bbaeba6102008401527f2ac8cc9e491532f6d70b6e89eee1c79ab825f80f577359901c6f34d88ee5fee66102208401526102408301526102608201527f20910a76f4aee5f0412948d472ba4a9396998e6cbbda72908585d37503a269336102808201527f158a7dce4fbc834dde5705059b7167e554566fdd51864c7508ebeafc68eaa6ed6102a08201527f0bd059852f29b878691531bbc4c5b4f3d6a0842645825fa0cbdc411e639dec446102c08201527f2f74c3ae7c9185ad5e7674108b5b452e34f3f6efe32b6c9f3c8e0cf8386bacaf6102e08201526020816103008160085afa90511615610eef57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b9181601f840112156108205782359167ffffffffffffffff8311610820576020838186019501011161082057565b90601f8019910116810190811067ffffffffffffffff82111761083457604052565b905f5160206115b65f395f51905f528210801590610fe7575b610eef57811580610fdf575b610fd957610fa65f5160206115b65f395f51905f52600381858181800909086113d6565b818103610fb557505060011b90565b5f5160206115b65f395f51905f52809106810306145f14610eef57600190811b1790565b50505f90565b508015610f82565b505f5160206115b65f395f51905f52811015610f76565b919093925f5160206115b65f395f51905f52831080159061121b575b8015611204575b80156111ed575b610eef5780828685171717156111e2579082916111455f5160206115b65f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f5160206115b65f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161111f81808b800981878009086113d6565b8408095f5160206115b65f395f51905f526111398261154d565b800914159586916113f9565b9290808214806111d9575b156111775750505050905f1461116f5760ff60025b169060021b179190565b60ff5f611165565b5f5160206115b65f395f51905f528091068103061491826111ba575b505015610eef57600191156111b25760ff60025b169060021b17179190565b60ff5f6111a7565b5f5160206115b65f395f51905f52919250819006810306145f80611193565b50838314611150565b50505090505f905f90565b505f5160206115b65f395f51905f52811015611028565b505f5160206115b65f395f51905f52821015611021565b505f5160206115b65f395f51905f5285101561101a565b8015611296578060011c915f5160206115b65f395f51905f52831015610eef576001806112755f5160206115b65f395f51905f52600381888181800909086113d6565b93161461127e57565b905f5160206115b65f395f51905f5280910681030690565b505f905f90565b8015806113ce575b6113c2578060021c92825f5160206115b65f395f51905f5285108015906113ab575b610eef5784815f5160206115b65f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816113759d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5086113f9565b80929160018082961614611387575050565b5f5160206115b65f395f51905f528093945080929550809106810306930681030690565b505f5160206115b65f395f51905f528110156112c7565b50505f905f905f905f90565b5081156112a5565b906113e08261154d565b915f5160206115b65f395f51905f5283800903610eef57565b915f5160206115b65f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816114519396949661144382808a8009818a8009086113d6565b90611541575b8608096113d6565b925f5160206115b65f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206115b65f395f51905f5260a083015260208260c08160055afa91519115610eef575f5160206115b65f395f51905f52826001920903610eef575f5160206115b65f395f51905f52908209925f5160206115b65f395f51905f528080808780090681030681878009081490811591611522575b50610eef57565b90505f5160206115b65f395f51905f528084860960020914155f61151b565b81809106810306611449565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206115b65f395f51905f5260a083015260208260c08160055afa91519115610eef5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a264697066735822122027b7f31a60cb804beb702dde7ab580eca3fa271695a78d109c179410d98441e864736f6c634300081c0033",
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
