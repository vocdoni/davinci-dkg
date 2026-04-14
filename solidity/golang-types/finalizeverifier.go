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
	Bin: "0x60808060405234601557611780908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace111461104a5750806344f6369214610fb15780635f89feef146108a65780638a3ae438146101f95763b8e72af614610055575f80fd5b346101c15760403660031901126101c15760043567ffffffffffffffff81116101c157610086903690600401611082565b60243567ffffffffffffffff81116101c1576100a6903690600401611082565b90918301610100848203126101c15780601f850112156101c157604051936100d0610100866110b0565b849061010081019283116101c157905b8282106101e9575050508101610120828203126101c15780601f830112156101c15760405191610112610120846110b0565b829061012081019283116101c157905b8282106101c557505050303b156101c1576040516311475c8760e31b8152915f600484015b600882106101ab5750505061010482015f905b60098210610195575050505f8161022481305afa801561018a5761017c575080f35b61018891505f906110b0565b005b6040513d5f823e3d90fd5b602080600192855181520193019101909161015a565b6020806001928551815201930191019091610147565b5f80fd5b8135815260209182019101610122565b634e487b7160e01b5f52604160045260245ffd5b81358152602091820191016100e0565b346101c1576102203660031901126101c15736610104116101c15736610224116101c15760405160408101907f0de41fe666e52ac8ead16cab4e8cce2838516521bec584def37e640361ae8ba7815260208101917f1dd6a75a8a9f295be544fc17183c5113a192548f25b7a9d9393597c03818c04d83527f2d7cd49ed5438f7a78ed22fbb82ad80132ab714280a97c213ddd962d3d654dce8152606082017f016f74ace6ad3077f63d46dfa48ac911249aaadeecb6c986e7ac7b132361edec81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f2045b97d3fe20d782702e03918a1a4fe2aea18dc4d6d0e1a7438bebb40822297608087019580875284848460608160075afa911016838860808160065afa167f27dff03d4d88616a4b626b57c7b1c4775beedbf8a22e618dc6eb9808a5f3504383527f2b4eaa1564379e88f342d6c047e34e636415fc82d25549563877d673d70fa92586526001610124359182895286868660608160075afa9310161616838860808160065afa167f109f4f019294a1a521abf5b25002450070aacde4611c820d8d3f00bd387f9e7483527f0642a2f44c92fbfac676a1ad25a8948342140817eda66eadf2db083b24ee31b48652610144359081885285858560608160075afa92101616838860808160065afa167f17d7dfc73c8d147ed76ddcb0cc3b1336dff703fbbf21a35649a031b83ceeea8883527f30438ff33fbe2a31f410e921816d69b082deb577a3adb63f7ca5b079fa3c38488652610164359081885285858560608160075afa92101616838860808160065afa167f0ef908f758b13e8814e1ea0c0ca3be94fc82e421a4a90430143aea23c691247e83527f135efbed41dd49fe7c420096bbb3676c4fac9c23a9d39b39b6550b7754a669928652610184359081885285858560608160075afa92101616838860808160065afa167f021477140f8f02f61d986d06cd04799f36c394a4b191756adef6cbabacc29a9183527f104225dacb680c434448922f3422f317a3f1502715ff4e8e6a20005dffb7fe5086526101a4359081885285858560608160075afa92101616838860808160065afa167f1c917d71789cf84c0f05fbf4e69298461675b1bec865e7b3c1b7fd8853701c1883527f1094a20ad33c26aa994667953547219034c661b636eed4c4e41a1e76abf2ce6f86526101c4359081885285858560608160075afa92101616838860808160065afa167ec5a5ae7591353d550dfb798cb5fa8669fe3bd3d70e6710e9c808beed36e2d083527f1e5b969d9c63d3c30d199e4d03fdb5d0d830a5f9c7079ebe8f1c3b1123c1a38b86526101e4359081885285858560608160075afa92101616838860808160065afa16947f15b1f4058a589b5cf9440090af0b0907a8427e99cdfafb4b0c9450b294dce8798352526102043580955260608160075afa9210161660408260808160065afa169051915190156108975760405191610100600484377f0963884d670da45101bfe996685cb5701efdf01b18470ddf7b2c61a26e2cf22d6101008401527f1b3de21ffee4004dbdcff08d09f07c57c0c55ac19520e7e29af4767a2dc14f226101208401527f2aac1e9ae0f81f590a0a33a04f400d65e07f31d6e9f89fbda73592ec3d7c0c716101408401527f10ffa358259a53551250274304807b5c5c2eb8f07b56ac630d9dbc3aa25002ca6101608401527f26cd44e557768bdfa98c7faac96f03803111276c938fca44696830c8257e3d076101808401527f29711dd6376909e4afa6ff5fdb186eea40839284d62f5044a87825f9a48f75146101a08401527f1987756debaf805b5ead11b0edb10843b705f9095fcc06877aa53664aaecafbe6101c08401527f152945ad06c41ee0c4aaa35ffeba2a86f418f254b2ab5338edeb4f50ae0c07fa6101e08401527f1d11f611311f89d77a629376d6fcd338908427fe939b0f6af4eb1db2d14daa106102008401527f2df0706a160ba1bc40c6190c2039e1d10e58a4d79d55d32bf3e0dded4dcbbc476102208401526102408301526102608201527f20d92ff22d68ca144cdad32cac845bbfb723c3e10cc300064f5b6d3fd21112c36102808201527f02e489adcceffca38bc603d9cb2377f9ab0e5f9aa8fab34be3017a8af1a6bcba6102a08201527f1b2b7c58ddbb1bda59318268bf8ed0142a49e95f70fea210b136a9f61726fa576102c08201527f1d9871c38d5049d1f15489905498701d5c3ac823da1148b36e80c33fd6a598656102e08201526020816103008160085afa9051161561088857005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b346101c1576101a03660031901126101c157366084116101c157366101a4116101c1576103006040516108d982826110b0565b813682376108e86004356113a7565b6108f9602493929335604435611412565b919392906109086064356113a7565b9390926040519660408801967f0de41fe666e52ac8ead16cab4e8cce2838516521bec584def37e640361ae8ba789528860208101987f1dd6a75a8a9f295be544fc17183c5113a192548f25b7a9d9393597c03818c04d8a527f2d7cd49ed5438f7a78ed22fbb82ad80132ab714280a97c213ddd962d3d654dce81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f016f74ace6ad3077f63d46dfa48ac911249aaadeecb6c986e7ac7b132361edec84527f2045b97d3fe20d782702e03918a1a4fe2aea18dc4d6d0e1a7438bebb408222976084359583608082019780895286828660608160075afa911016818360808160065afa167f27dff03d4d88616a4b626b57c7b1c4775beedbf8a22e618dc6eb9808a5f3504385527f2b4eaa1564379e88f342d6c047e34e636415fc82d25549563877d673d70fa9258852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f109f4f019294a1a521abf5b25002450070aacde4611c820d8d3f00bd387f9e7485527f0642a2f44c92fbfac676a1ad25a8948342140817eda66eadf2db083b24ee31b4885260c43590818a5287838760608160075afa92101616818360808160065afa167f17d7dfc73c8d147ed76ddcb0cc3b1336dff703fbbf21a35649a031b83ceeea8885527f30438ff33fbe2a31f410e921816d69b082deb577a3adb63f7ca5b079fa3c3848885260e43590818a5287838760608160075afa92101616818360808160065afa167f0ef908f758b13e8814e1ea0c0ca3be94fc82e421a4a90430143aea23c691247e85527f135efbed41dd49fe7c420096bbb3676c4fac9c23a9d39b39b6550b7754a6699288526101043590818a5287838760608160075afa92101616818360808160065afa167f021477140f8f02f61d986d06cd04799f36c394a4b191756adef6cbabacc29a9185527f104225dacb680c434448922f3422f317a3f1502715ff4e8e6a20005dffb7fe5088526101243590818a5287838760608160075afa92101616818360808160065afa167f1c917d71789cf84c0f05fbf4e69298461675b1bec865e7b3c1b7fd8853701c1885527f1094a20ad33c26aa994667953547219034c661b636eed4c4e41a1e76abf2ce6f88526101443590818a5287838760608160075afa92101616818360808160065afa167ec5a5ae7591353d550dfb798cb5fa8669fe3bd3d70e6710e9c808beed36e2d085527f1e5b969d9c63d3c30d199e4d03fdb5d0d830a5f9c7079ebe8f1c3b1123c1a38b88526101643590818a5287838760608160075afa921016169160808160065afa16947f15b1f4058a589b5cf9440090af0b0907a8427e99cdfafb4b0c9450b294dce8798352526101843580955260608160075afa9210161660408a60808160065afa169851975198156108975760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f0963884d670da45101bfe996685cb5701efdf01b18470ddf7b2c61a26e2cf22d6101008401527f1b3de21ffee4004dbdcff08d09f07c57c0c55ac19520e7e29af4767a2dc14f226101208401527f2aac1e9ae0f81f590a0a33a04f400d65e07f31d6e9f89fbda73592ec3d7c0c716101408401527f10ffa358259a53551250274304807b5c5c2eb8f07b56ac630d9dbc3aa25002ca6101608401527f26cd44e557768bdfa98c7faac96f03803111276c938fca44696830c8257e3d076101808401527f29711dd6376909e4afa6ff5fdb186eea40839284d62f5044a87825f9a48f75146101a08401527f1987756debaf805b5ead11b0edb10843b705f9095fcc06877aa53664aaecafbe6101c08401527f152945ad06c41ee0c4aaa35ffeba2a86f418f254b2ab5338edeb4f50ae0c07fa6101e08401527f1d11f611311f89d77a629376d6fcd338908427fe939b0f6af4eb1db2d14daa106102008401527f2df0706a160ba1bc40c6190c2039e1d10e58a4d79d55d32bf3e0dded4dcbbc476102208401526102408301526102608201527f20d92ff22d68ca144cdad32cac845bbfb723c3e10cc300064f5b6d3fd21112c36102808201527f02e489adcceffca38bc603d9cb2377f9ab0e5f9aa8fab34be3017a8af1a6bcba6102a08201527f1b2b7c58ddbb1bda59318268bf8ed0142a49e95f70fea210b136a9f61726fa576102c08201527f1d9871c38d5049d1f15489905498701d5c3ac823da1148b36e80c33fd6a598656102e0820152604051928391610f8c84846110b0565b8336843760085afa15908115610fa4575b5061088857005b6001915051141581610f9d565b346101c1576101003660031901126101c15736610104116101c157604051610fda6080826110b0565b6080368237610fed6024356004356110d2565b815261100360843560a435604435606435611173565b6020830152604082015261101b60e43560c4356110d2565b6060820152604051905f825b6004821061103457608084f35b6020806001928551815201930191019091611027565b346101c1575f3660031901126101c157807fa5a0381635973c2599e2e0bd559739beb8912aa8607280773b4f2f24f9e7f03560209252f35b9181601f840112156101c15782359167ffffffffffffffff83116101c157602083818601950101116101c157565b90601f8019910116810190811067ffffffffffffffff8211176101d557604052565b905f51602061172b5f395f51905f52821080159061115c575b61088857811580611154575b61114e5761111b5f51602061172b5f395f51905f526003818581818009090861154b565b81810361112a57505060011b90565b5f51602061172b5f395f51905f52809106810306145f1461088857600190811b1790565b50505f90565b5080156110f7565b505f51602061172b5f395f51905f528110156110eb565b919093925f51602061172b5f395f51905f528310801590611390575b8015611379575b8015611362575b610888578082868517171715611357579082916112ba5f51602061172b5f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f51602061172b5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161129481808b8009818780090861154b565b8408095f51602061172b5f395f51905f526112ae826116c2565b8009141595869161156e565b92908082148061134e575b156112ec5750505050905f146112e45760ff60025b169060021b179190565b60ff5f6112da565b5f51602061172b5f395f51905f5280910681030614918261132f575b50501561088857600191156113275760ff60025b169060021b17179190565b60ff5f61131c565b5f51602061172b5f395f51905f52919250819006810306145f80611308565b508383146112c5565b50505090505f905f90565b505f51602061172b5f395f51905f5281101561119d565b505f51602061172b5f395f51905f52821015611196565b505f51602061172b5f395f51905f5285101561118f565b801561140b578060011c915f51602061172b5f395f51905f52831015610888576001806113ea5f51602061172b5f395f51905f526003818881818009090861154b565b9316146113f357565b905f51602061172b5f395f51905f5280910681030690565b505f905f90565b801580611543575b611537578060021c92825f51602061172b5f395f51905f528510801590611520575b6108885784815f51602061172b5f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816114ea9d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50861156e565b809291600180829616146114fc575050565b5f51602061172b5f395f51905f528093945080929550809106810306930681030690565b505f51602061172b5f395f51905f5281101561143c565b50505f905f905f905f90565b50811561141a565b90611555826116c2565b915f51602061172b5f395f51905f528380090361088857565b915f51602061172b5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816115c6939694966115b882808a8009818a80090861154b565b906116b6575b86080961154b565b925f51602061172b5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f51602061172b5f395f51905f5260a083015260208260c08160055afa91519115610888575f51602061172b5f395f51905f52826001920903610888575f51602061172b5f395f51905f52908209925f51602061172b5f395f51905f528080808780090681030681878009081490811591611697575b5061088857565b90505f51602061172b5f395f51905f528084860960020914155f611690565b818091068103066115be565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f51602061172b5f395f51905f5260a083015260208260c08160055afa915191156108885756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a264697066735822122000e2c13b84b9a942cc0e4fd4c5747d00af80602ee2e5df9f018c3143c4655c5d64736f6c634300081c0033",
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
