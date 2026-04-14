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
	Bin: "0x60808060405234601557611782908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace111461104c5750806344f6369214610fb35780635f89feef146108a75780638a3ae438146101f95763b8e72af614610055575f80fd5b346101c15760403660031901126101c15760043567ffffffffffffffff81116101c157610086903690600401611084565b60243567ffffffffffffffff81116101c1576100a6903690600401611084565b90918301610100848203126101c15780601f850112156101c157604051936100d0610100866110b2565b849061010081019283116101c157905b8282106101e9575050508101610120828203126101c15780601f830112156101c15760405191610112610120846110b2565b829061012081019283116101c157905b8282106101c557505050303b156101c1576040516311475c8760e31b8152915f600484015b600882106101ab5750505061010482015f905b60098210610195575050505f8161022481305afa801561018a5761017c575080f35b61018891505f906110b2565b005b6040513d5f823e3d90fd5b602080600192855181520193019101909161015a565b6020806001928551815201930191019091610147565b5f80fd5b8135815260209182019101610122565b634e487b7160e01b5f52604160045260245ffd5b81358152602091820191016100e0565b346101c1576102203660031901126101c15736610104116101c15736610224116101c15760405160408101907f269e9b4e5aa9333bd05b91764b44b59fd5a4295d397c05b95f993615b7b59f76815260208101917f2b0d92d115b9d3c87e9d491d63db3723a093922f4617669a1d11e53830140eef83527f2e65c50f3bf067bdd1f9640aa7dc9e34dc8b5eb034fe75a1b0029ccf71b563af8152606082017f070fb3b8bf1ddbaea3dc583cd450faaf951127e1ad786fa3ce85ebd27f3d4e1581527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f042cc62a2a5dd2a45775d86f0ef84f1f77c71ff887f884e3fc1eb1cd2f4d138b608087019580875284848460608160075afa911016838860808160065afa167f29830d2963711662d45b0f372af869d45d13744350ea138cc38b4f00443ff66e83527f29b54e5a656521a69a695a54f4dcc9457d4567103510bfa684e5f8d690c27cee86526001610124359182895286868660608160075afa9310161616838860808160065afa167f033ad8f3e62b74b17a735be0ae39fe20ad921122d54151e49210c38b5fc62e0b83527f11f8f3ebf5c2c82cb77ded63a285500540bec141fe7c695d436c87d48f9742a88652610144359081885285858560608160075afa92101616838860808160065afa167f18d0ea125cb1c043d5bfbcabf94aedb1e5790110febd6180191673d19903f85483527f0998837541d9cfbb74a2e80681ed4cde3b9be0d863fb6b4cb275c27c75826edd8652610164359081885285858560608160075afa92101616838860808160065afa167f05b2d781402f6c497b9ba537da9431808a9bea3b99d6530e7e7555c16929e3a283527f277b49a5544e5a07e2302231fd829ac639cff20937bd6615c87e0edd3707e4248652610184359081885285858560608160075afa92101616838860808160065afa167f2c681053525b39dcbce6528819cf0ee526e6367cac8d611cb9da9120a6acf15283527f21f198ddcabc87cd0c62a365135b83f37eb4d43bcb9d34bc854a862c52843dbd86526101a4359081885285858560608160075afa92101616838860808160065afa167f026f739bdb57decfe80a76bc2a46b8d4b75e6476c611e59ce10871286e4f4bcb83527f3061639ef760511c3b559eb98baeea681b0f9719ab3538968048c33ad0cf96b686526101c4359081885285858560608160075afa92101616838860808160065afa167f0685ea14bd7c113301145722b3297eed0edc09b0327b2fecb22b6ce50b240ce183527f253092d1df10842663df2b8a52e78124b4bb65e7266b6fce63f933a6111bafdb86526101e4359081885285858560608160075afa92101616838860808160065afa16947f11f94524b380cd342fd82cfee54b9aad03052332a5c82d46822ea982eab95f028352526102043580955260608160075afa9210161660408260808160065afa169051915190156108985760405191610100600484377f268f819b9bad433b22b26258064ee2b2c52ec468afa1b773395dd51afd7d5ba06101008401527f05a3612e176d0f8a10fdd4aa69ee0d6a71ab0876aa04c4230d7be825ac80cbd66101208401527f26432b67a131c7d9131540ba7f2e06650776d84d95d1000621311c73a901dc826101408401527f2b8e24812dd0ac58a67f9f470fbbd27f533d8dbb489325dc87a42e48d7f87b566101608401527f10bedb923124267bfa76ecf8ed463b2799a0f8484a975d94af680fd915aea9df6101808401527f038293f90d003291dcc24fbc72ce82c06946db678b6d268bed95fefad59c5c5a6101a08401527f1b320bcac8345edee7def0fffafc5bb24589d4eecd4ebf8e27cdfaf4044a91116101c08401527f027b70cdd82e6fb0d02bfafe6770f9a554166a49a0f701c242751c5d858d1b536101e08401527f1703ef3aca90619360428d645f8b30beeafe5e6c7d3c44ccc6b4f4bb7b89013b6102008401527f2011b4f712fd17f6b3086494533c63e498d3e62a30c174d59248715f2b6c90a76102208401526102408301526102608201527f20c030d5e5d852e9769a9ef81a5258d1bf0d02007604f6406007f769286390466102808201527f0c454d5deeb97770e0b1f4367b4420cac96247829e6b9e1c28b82dc7e0015f9d6102a08201527f1dd122831b310da54da5debdcfc727054711cd22983e928f9173abc8a5124d566102c08201527f2337202b89750339b3b0686bbf51ab3f988e4ac6bb0ee7cc876d57051be191ad6102e08201526020816103008160085afa9051161561088957005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b346101c1576101a03660031901126101c157366084116101c157366101a4116101c1576103006040516108da82826110b2565b813682376108e96004356113a9565b6108fa602493929335604435611414565b919392906109096064356113a9565b9390926040519660408801967f269e9b4e5aa9333bd05b91764b44b59fd5a4295d397c05b95f993615b7b59f7689528860208101987f2b0d92d115b9d3c87e9d491d63db3723a093922f4617669a1d11e53830140eef8a527f2e65c50f3bf067bdd1f9640aa7dc9e34dc8b5eb034fe75a1b0029ccf71b563af81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f070fb3b8bf1ddbaea3dc583cd450faaf951127e1ad786fa3ce85ebd27f3d4e1584527f042cc62a2a5dd2a45775d86f0ef84f1f77c71ff887f884e3fc1eb1cd2f4d138b6084359583608082019780895286828660608160075afa911016818360808160065afa167f29830d2963711662d45b0f372af869d45d13744350ea138cc38b4f00443ff66e85527f29b54e5a656521a69a695a54f4dcc9457d4567103510bfa684e5f8d690c27cee8852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f033ad8f3e62b74b17a735be0ae39fe20ad921122d54151e49210c38b5fc62e0b85527f11f8f3ebf5c2c82cb77ded63a285500540bec141fe7c695d436c87d48f9742a8885260c43590818a5287838760608160075afa92101616818360808160065afa167f18d0ea125cb1c043d5bfbcabf94aedb1e5790110febd6180191673d19903f85485527f0998837541d9cfbb74a2e80681ed4cde3b9be0d863fb6b4cb275c27c75826edd885260e43590818a5287838760608160075afa92101616818360808160065afa167f05b2d781402f6c497b9ba537da9431808a9bea3b99d6530e7e7555c16929e3a285527f277b49a5544e5a07e2302231fd829ac639cff20937bd6615c87e0edd3707e42488526101043590818a5287838760608160075afa92101616818360808160065afa167f2c681053525b39dcbce6528819cf0ee526e6367cac8d611cb9da9120a6acf15285527f21f198ddcabc87cd0c62a365135b83f37eb4d43bcb9d34bc854a862c52843dbd88526101243590818a5287838760608160075afa92101616818360808160065afa167f026f739bdb57decfe80a76bc2a46b8d4b75e6476c611e59ce10871286e4f4bcb85527f3061639ef760511c3b559eb98baeea681b0f9719ab3538968048c33ad0cf96b688526101443590818a5287838760608160075afa92101616818360808160065afa167f0685ea14bd7c113301145722b3297eed0edc09b0327b2fecb22b6ce50b240ce185527f253092d1df10842663df2b8a52e78124b4bb65e7266b6fce63f933a6111bafdb88526101643590818a5287838760608160075afa921016169160808160065afa16947f11f94524b380cd342fd82cfee54b9aad03052332a5c82d46822ea982eab95f028352526101843580955260608160075afa9210161660408a60808160065afa169851975198156108985760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f268f819b9bad433b22b26258064ee2b2c52ec468afa1b773395dd51afd7d5ba06101008401527f05a3612e176d0f8a10fdd4aa69ee0d6a71ab0876aa04c4230d7be825ac80cbd66101208401527f26432b67a131c7d9131540ba7f2e06650776d84d95d1000621311c73a901dc826101408401527f2b8e24812dd0ac58a67f9f470fbbd27f533d8dbb489325dc87a42e48d7f87b566101608401527f10bedb923124267bfa76ecf8ed463b2799a0f8484a975d94af680fd915aea9df6101808401527f038293f90d003291dcc24fbc72ce82c06946db678b6d268bed95fefad59c5c5a6101a08401527f1b320bcac8345edee7def0fffafc5bb24589d4eecd4ebf8e27cdfaf4044a91116101c08401527f027b70cdd82e6fb0d02bfafe6770f9a554166a49a0f701c242751c5d858d1b536101e08401527f1703ef3aca90619360428d645f8b30beeafe5e6c7d3c44ccc6b4f4bb7b89013b6102008401527f2011b4f712fd17f6b3086494533c63e498d3e62a30c174d59248715f2b6c90a76102208401526102408301526102608201527f20c030d5e5d852e9769a9ef81a5258d1bf0d02007604f6406007f769286390466102808201527f0c454d5deeb97770e0b1f4367b4420cac96247829e6b9e1c28b82dc7e0015f9d6102a08201527f1dd122831b310da54da5debdcfc727054711cd22983e928f9173abc8a5124d566102c08201527f2337202b89750339b3b0686bbf51ab3f988e4ac6bb0ee7cc876d57051be191ad6102e0820152604051928391610f8e84846110b2565b8336843760085afa15908115610fa6575b5061088957005b6001915051141581610f9f565b346101c1576101003660031901126101c15736610104116101c157604051610fdc6080826110b2565b6080368237610fef6024356004356110d4565b815261100560843560a435604435606435611175565b6020830152604082015261101d60e43560c4356110d4565b6060820152604051905f825b6004821061103657608084f35b6020806001928551815201930191019091611029565b346101c1575f3660031901126101c157807fa5a0381635973c2599e2e0bd559739beb8912aa8607280773b4f2f24f9e7f03560209252f35b9181601f840112156101c15782359167ffffffffffffffff83116101c157602083818601950101116101c157565b90601f8019910116810190811067ffffffffffffffff8211176101d557604052565b905f51602061172d5f395f51905f52821080159061115e575b61088957811580611156575b6111505761111d5f51602061172d5f395f51905f526003818581818009090861154d565b81810361112c57505060011b90565b5f51602061172d5f395f51905f52809106810306145f1461088957600190811b1790565b50505f90565b5080156110f9565b505f51602061172d5f395f51905f528110156110ed565b919093925f51602061172d5f395f51905f528310801590611392575b801561137b575b8015611364575b610889578082868517171715611359579082916112bc5f51602061172d5f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f51602061172d5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161129681808b8009818780090861154d565b8408095f51602061172d5f395f51905f526112b0826116c4565b80091415958691611570565b929080821480611350575b156112ee5750505050905f146112e65760ff60025b169060021b179190565b60ff5f6112dc565b5f51602061172d5f395f51905f52809106810306149182611331575b50501561088957600191156113295760ff60025b169060021b17179190565b60ff5f61131e565b5f51602061172d5f395f51905f52919250819006810306145f8061130a565b508383146112c7565b50505090505f905f90565b505f51602061172d5f395f51905f5281101561119f565b505f51602061172d5f395f51905f52821015611198565b505f51602061172d5f395f51905f52851015611191565b801561140d578060011c915f51602061172d5f395f51905f52831015610889576001806113ec5f51602061172d5f395f51905f526003818881818009090861154d565b9316146113f557565b905f51602061172d5f395f51905f5280910681030690565b505f905f90565b801580611545575b611539578060021c92825f51602061172d5f395f51905f528510801590611522575b6108895784815f51602061172d5f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816114ec9d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611570565b809291600180829616146114fe575050565b5f51602061172d5f395f51905f528093945080929550809106810306930681030690565b505f51602061172d5f395f51905f5281101561143e565b50505f905f905f905f90565b50811561141c565b90611557826116c4565b915f51602061172d5f395f51905f528380090361088957565b915f51602061172d5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816115c8939694966115ba82808a8009818a80090861154d565b906116b8575b86080961154d565b925f51602061172d5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f51602061172d5f395f51905f5260a083015260208260c08160055afa91519115610889575f51602061172d5f395f51905f52826001920903610889575f51602061172d5f395f51905f52908209925f51602061172d5f395f51905f528080808780090681030681878009081490811591611699575b5061088957565b90505f51602061172d5f395f51905f528084860960020914155f611692565b818091068103066115c0565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f51602061172d5f395f51905f5260a083015260208260c08160055afa915191156108895756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a264697066735822122044892a15e180c7e6b11742e14676eee7fe2d24117e8427e11b81d4c94eabf3c964736f6c634300081c0033",
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
