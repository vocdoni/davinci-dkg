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

// ContributionVerifierMetaData contains all meta data concerning the ContributionVerifier contract.
var ContributionVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[10]\",\"internalType\":\"uint256[10]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"input\",\"type\":\"uint256[10]\",\"internalType\":\"uint256[10]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidInputEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x608080604052346015576119ab908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c908163233ace1114610fba5750806344f6369214610f3c57806386836a97146107ca57806394e4398a146100b65763b8e72af614610053575f80fd5b346100b25760403660031901126100b25760043567ffffffffffffffff81116100b257610084903690600401611019565b6024359167ffffffffffffffff83116100b2576100a86100b0933690600401611019565b9290916110f8565b005b5f80fd5b346100b2576102403660031901126100b25736610104116100b25736610244116100b25760405160408101907f1f06539765aefa237260d6cc23988fdb9e2feeb1f709761a9cdd45f80154876b815260208101917f0843bb6979d96290cbda398296129ea86f751877db6449ceaf4ccb137a45065883527f15888dedaa3eb08afff2440e44a1128732315c24ebd329daa4ead3e59a96a56b8152606082017f2c0c134b820c863bbc19a2f0e2d0e8efc16f1073afcb230944d0a6156f1029ed81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f0cfb51f6c9bd3f7971c4e8307d0633784eb005702dbf8cef5076267000030822608087019580875284848460608160075afa911016838860808160065afa167f1f238a17994ae5f5a9355d3aede58d1ca6d6fc79aa4231f8f5fe2dafca667b3b83527f26a4db90591596fa26c0e6d81425f642afa1dc803498ebe62d50707143806edd86526001610124359182895286868660608160075afa9310161616838860808160065afa167f2b72f1c6810df9cc1ba76ce747b7969a599a0a9812aa928c7ed59015a33d062083527f07dcf5a7dfee2eeb87e7be527cd105051226b9845ef9d96f5fd7a7e7797c285e8652610144359081885285858560608160075afa92101616838860808160065afa167f2e0d3c6e85cf3d413df489a48d73114442b65dfbbe33b02ebf841998ac6ebe5a83527f30560883d97a7b339189b525cc0c2647e9e230825b24cd8f67730c231910cec38652610164359081885285858560608160075afa92101616838860808160065afa167f084c64171eb82a860dd85e8a062c1602aec9a3d466ef4b97712e30ecd3b61e4283527f0cb232124ac739697b839c636f404df64a5348c43113b40584671d1a07b12ac98652610184359081885285858560608160075afa92101616838860808160065afa167f037500750a5a9227ebf3a7a7db7a02a4c542b7d0706bdd3b68c6283c605a07b083527f0ec4d142f4be1940e0a4b4ceb0de525b8c2757ddc7225f052f0edfa67493cac486526101a4359081885285858560608160075afa92101616838860808160065afa167f1e33762dfaa15f8e925600ef120b2284f7f94ecace83c450faea4f821d2d1ae083527f302179fb1bd6ba24aa5d7112382f7a731464a21593e68a9690fe33eeafcd742286526101c4359081885285858560608160075afa92101616838860808160065afa167f26c0c3a886c1ffaa915405addc2e38413cbffcac983ace3ee37c17d75877f13583527f05d3a4b680c9da68e881f02e57f35aa91eef82451eea079dea68f56bddd5feac86526101e4359081885285858560608160075afa92101616838860808160065afa167f19628d6c5c56f42e5fa8491a664b0d6ccd18a9a5bcdccb72ce68f8dd69f84f3a83527f045cb1f61ed3624b4c985fec728b60506ade2a0e4006523df02f30ae1ddaabc68652610204359081885285858560608160075afa92101616838860808160065afa16947f0f497af6708dc96c2993781c222d6a80d31793ae53d89238b6430f3447cab47d8352526102243580955260608160075afa9210161660408260808160065afa169051915190156107bb5760405191610100600484377f2edfe6d90bba3bf6070a218397249cb7ccb46117bdbc25d04a74bbd838d1bfa46101008401527f2812d255e1d15b8f3e33cd19b5927c07b356339779709bdd42b0b15d6a6a76e76101208401527f145084bfca2c6dc3a68546212604549114e6fb5fb3166580c97471c0043e2acf6101408401527f0a6e26681f24ac14d9792dc25cef2fb38bd105314ab593c2f70c3bc4d30b4edb6101608401527f0e9ab95bd296973779b76f7dbec5fbf5b03532df36695896f29820bd8fe47d706101808401527f05dc7e6aa29795670f284bbcc43cf35a39d0b8400583db7896fa8cd39de6ede46101a08401527f2da7a5f3c263196b720eea4bb47928a133acd95bcdb2a0123e3ea5e2d55d4bfc6101c08401527f14f6d91d850cb00656d06aa758a2718c4544e36722cddfb968192ffb22b3bbfa6101e08401527f0cd41ff0856235457705012f8871f3a6a34304c0b997714604ed243b4cd820426102008401527f0fa75f376c64e0d6c87448ed2c46f1788d1b827132edf14aefceaf5b7831b4ce6102208401526102408301526102608201527f02bb056046f42456d52cbca0621059c2a3a8962615c0873715c7c74ee8fdd2166102808201527f1fe20aa4fa923a0bca8cbbf52703c818c2033ed65cce08348412f8cac201b0ec6102a08201527f1625784d7fc5cd8a3446e0f4ced19c6f55c75ac8b1e7742b520636eff3df57a86102c08201527f1c474fe6248d139598186a2f7feb074d34ad16343c69e964b38b2bf191c0c9416102e08201526020816103008160085afa905116156107ac57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b346100b2576101c03660031901126100b257366084116100b257366101c4116100b2576103006040516107fd8282611047565b8136823761080c600435611575565b61081d6024939293356044356115e0565b9193929061082c606435611575565b9390926040519660408801967f1f06539765aefa237260d6cc23988fdb9e2feeb1f709761a9cdd45f80154876b89528860208101987f0843bb6979d96290cbda398296129ea86f751877db6449ceaf4ccb137a4506588a527f15888dedaa3eb08afff2440e44a1128732315c24ebd329daa4ead3e59a96a56b81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f2c0c134b820c863bbc19a2f0e2d0e8efc16f1073afcb230944d0a6156f1029ed84527f0cfb51f6c9bd3f7971c4e8307d0633784eb005702dbf8cef50762670000308226084359583608082019780895286828660608160075afa911016818360808160065afa167f1f238a17994ae5f5a9355d3aede58d1ca6d6fc79aa4231f8f5fe2dafca667b3b85527f26a4db90591596fa26c0e6d81425f642afa1dc803498ebe62d50707143806edd8852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f2b72f1c6810df9cc1ba76ce747b7969a599a0a9812aa928c7ed59015a33d062085527f07dcf5a7dfee2eeb87e7be527cd105051226b9845ef9d96f5fd7a7e7797c285e885260c43590818a5287838760608160075afa92101616818360808160065afa167f2e0d3c6e85cf3d413df489a48d73114442b65dfbbe33b02ebf841998ac6ebe5a85527f30560883d97a7b339189b525cc0c2647e9e230825b24cd8f67730c231910cec3885260e43590818a5287838760608160075afa92101616818360808160065afa167f084c64171eb82a860dd85e8a062c1602aec9a3d466ef4b97712e30ecd3b61e4285527f0cb232124ac739697b839c636f404df64a5348c43113b40584671d1a07b12ac988526101043590818a5287838760608160075afa92101616818360808160065afa167f037500750a5a9227ebf3a7a7db7a02a4c542b7d0706bdd3b68c6283c605a07b085527f0ec4d142f4be1940e0a4b4ceb0de525b8c2757ddc7225f052f0edfa67493cac488526101243590818a5287838760608160075afa92101616818360808160065afa167f1e33762dfaa15f8e925600ef120b2284f7f94ecace83c450faea4f821d2d1ae085527f302179fb1bd6ba24aa5d7112382f7a731464a21593e68a9690fe33eeafcd742288526101443590818a5287838760608160075afa92101616818360808160065afa167f26c0c3a886c1ffaa915405addc2e38413cbffcac983ace3ee37c17d75877f13585527f05d3a4b680c9da68e881f02e57f35aa91eef82451eea079dea68f56bddd5feac88526101643590818a5287838760608160075afa92101616818360808160065afa167f19628d6c5c56f42e5fa8491a664b0d6ccd18a9a5bcdccb72ce68f8dd69f84f3a85527f045cb1f61ed3624b4c985fec728b60506ade2a0e4006523df02f30ae1ddaabc688526101843590818a5287838760608160075afa921016169160808160065afa16947f0f497af6708dc96c2993781c222d6a80d31793ae53d89238b6430f3447cab47d8352526101a43580955260608160075afa9210161660408a60808160065afa169851975198156107bb5760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f2edfe6d90bba3bf6070a218397249cb7ccb46117bdbc25d04a74bbd838d1bfa46101008401527f2812d255e1d15b8f3e33cd19b5927c07b356339779709bdd42b0b15d6a6a76e76101208401527f145084bfca2c6dc3a68546212604549114e6fb5fb3166580c97471c0043e2acf6101408401527f0a6e26681f24ac14d9792dc25cef2fb38bd105314ab593c2f70c3bc4d30b4edb6101608401527f0e9ab95bd296973779b76f7dbec5fbf5b03532df36695896f29820bd8fe47d706101808401527f05dc7e6aa29795670f284bbcc43cf35a39d0b8400583db7896fa8cd39de6ede46101a08401527f2da7a5f3c263196b720eea4bb47928a133acd95bcdb2a0123e3ea5e2d55d4bfc6101c08401527f14f6d91d850cb00656d06aa758a2718c4544e36722cddfb968192ffb22b3bbfa6101e08401527f0cd41ff0856235457705012f8871f3a6a34304c0b997714604ed243b4cd820426102008401527f0fa75f376c64e0d6c87448ed2c46f1788d1b827132edf14aefceaf5b7831b4ce6102208401526102408301526102608201527f02bb056046f42456d52cbca0621059c2a3a8962615c0873715c7c74ee8fdd2166102808201527f1fe20aa4fa923a0bca8cbbf52703c818c2033ed65cce08348412f8cac201b0ec6102a08201527f1625784d7fc5cd8a3446e0f4ced19c6f55c75ac8b1e7742b520636eff3df57a86102c08201527f1c474fe6248d139598186a2f7feb074d34ad16343c69e964b38b2bf191c0c9416102e0820152604051928391610f178484611047565b8336843760085afa15908115610f2f575b506107ac57005b6001915051141581610f28565b346100b2576101003660031901126100b25736610104116100b2576080604051610f668282611047565b81368237610f786024356004356112a0565b8152610f8e60843560a435604435606435611341565b60208301526040820152610fa660e43560c4356112a0565b6060820152610fb86040518092610ff2565bf35b346100b2575f3660031901126100b257807fe662b9f5b48e5c825efe2667a1b8e144cd262f125c06fcc794f6f198dea9f15360209252f35b905f905b6004821061100357505050565b6020806001928551815201930191019091610ff6565b9181601f840112156100b25782359167ffffffffffffffff83116100b257602083818601950101116100b257565b90601f8019910116810190811067ffffffffffffffff82111761106957604052565b634e487b7160e01b5f52604160045260245ffd5b90610140828203126100b25780601f830112156100b257604051916110a461014084611047565b829061014081019283116100b257905b8282106110c15750505090565b81358152602091820191016110b4565b905f905b600a82106110e257505050565b60208060019285518152019301910190916110d5565b9392919061010081146111da576080811461111c5763236bd13760e01b5f5260045ffd5b61014083036111cb578401916080858403126100b25782601f860112156100b2576040519261114c608085611047565b8395608081019182116100b257955b8187106111bb5750506111b99394508161117a916111a393019061107d565b6040516386836a9760e01b602082015292611199906024850190610ff2565b60a48301906110d1565b6101c481526111b46101e482611047565b611719565b565b863581526020968701960161115b565b630c0b7e3560e11b5f5260045ffd5b61014083036111cb57840191610100858403126100b25782601f860112156100b2576040519261120c61010085611047565b839561010081019182116100b257955b8187106112905750506112349293945081019061107d565b604051634a721cc560e11b6020820152915f602484015b6008821061127a57505050906112696111b9926101248301906110d1565b61024481526111b461026482611047565b602080600192855181520193019101909161124b565b863581526020968701960161121c565b905f5160206119565f395f51905f52821080159061132a575b6107ac57811580611322575b61131c576112e95f5160206119565f395f51905f5260038185818180090908611776565b8181036112f857505060011b90565b5f5160206119565f395f51905f52809106810306145f146107ac57600190811b1790565b50505f90565b5080156112c5565b505f5160206119565f395f51905f528110156112b9565b919093925f5160206119565f395f51905f52831080159061155e575b8015611547575b8015611530575b6107ac578082868517171715611525579082916114885f5160206119565f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f5160206119565f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161146281808b80098187800908611776565b8408095f5160206119565f395f51905f5261147c826118ed565b80091415958691611799565b92908082148061151c575b156114ba5750505050905f146114b25760ff60025b169060021b179190565b60ff5f6114a8565b5f5160206119565f395f51905f528091068103061491826114fd575b5050156107ac57600191156114f55760ff60025b169060021b17179190565b60ff5f6114ea565b5f5160206119565f395f51905f52919250819006810306145f806114d6565b50838314611493565b50505090505f905f90565b505f5160206119565f395f51905f5281101561136b565b505f5160206119565f395f51905f52821015611364565b505f5160206119565f395f51905f5285101561135d565b80156115d9578060011c915f5160206119565f395f51905f528310156107ac576001806115b85f5160206119565f395f51905f5260038188818180090908611776565b9316146115c157565b905f5160206119565f395f51905f5280910681030690565b505f905f90565b801580611711575b611705578060021c92825f5160206119565f395f51905f5285108015906116ee575b6107ac5784815f5160206119565f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816116b89d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611799565b809291600180829616146116ca575050565b5f5160206119565f395f51905f528093945080929550809106810306930681030690565b505f5160206119565f395f51905f5281101561160a565b50505f905f905f905f90565b5081156115e8565b5f8091602081519101305afa3d1561176e573d9067ffffffffffffffff82116110695760405191611754601f8201601f191660200184611047565b82523d5f602084013e5b156117665750565b602081519101fd5b60609061175e565b90611780826118ed565b915f5160206119565f395f51905f52838009036107ac57565b915f5160206119565f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816117f1939694966117e382808a8009818a800908611776565b906118e1575b860809611776565b925f5160206119565f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206119565f395f51905f5260a083015260208260c08160055afa915191156107ac575f5160206119565f395f51905f528260019209036107ac575f5160206119565f395f51905f52908209925f5160206119565f395f51905f5280808087800906810306818780090814908115916118c2575b506107ac57565b90505f5160206119565f395f51905f528084860960020914155f6118bb565b818091068103066117e9565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206119565f395f51905f5260a083015260208260c08160055afa915191156107ac5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220e7b5d28e836d9c4dfb0f8406918b6c28433e2bfdacd1eb9c6a5c9e67e663ee3464736f6c634300081c0033",
}

// ContributionVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use ContributionVerifierMetaData.ABI instead.
var ContributionVerifierABI = ContributionVerifierMetaData.ABI

// ContributionVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContributionVerifierMetaData.Bin instead.
var ContributionVerifierBin = ContributionVerifierMetaData.Bin

// DeployContributionVerifier deploys a new Ethereum contract, binding an instance of ContributionVerifier to it.
func DeployContributionVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContributionVerifier, error) {
	parsed, err := ContributionVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContributionVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContributionVerifier{ContributionVerifierCaller: ContributionVerifierCaller{contract: contract}, ContributionVerifierTransactor: ContributionVerifierTransactor{contract: contract}, ContributionVerifierFilterer: ContributionVerifierFilterer{contract: contract}}, nil
}

// ContributionVerifier is an auto generated Go binding around an Ethereum contract.
type ContributionVerifier struct {
	ContributionVerifierCaller     // Read-only binding to the contract
	ContributionVerifierTransactor // Write-only binding to the contract
	ContributionVerifierFilterer   // Log filterer for contract events
}

// ContributionVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContributionVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContributionVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContributionVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContributionVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContributionVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContributionVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContributionVerifierSession struct {
	Contract     *ContributionVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ContributionVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContributionVerifierCallerSession struct {
	Contract *ContributionVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ContributionVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContributionVerifierTransactorSession struct {
	Contract     *ContributionVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ContributionVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContributionVerifierRaw struct {
	Contract *ContributionVerifier // Generic contract binding to access the raw methods on
}

// ContributionVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContributionVerifierCallerRaw struct {
	Contract *ContributionVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// ContributionVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContributionVerifierTransactorRaw struct {
	Contract *ContributionVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContributionVerifier creates a new instance of ContributionVerifier, bound to a specific deployed contract.
func NewContributionVerifier(address common.Address, backend bind.ContractBackend) (*ContributionVerifier, error) {
	contract, err := bindContributionVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContributionVerifier{ContributionVerifierCaller: ContributionVerifierCaller{contract: contract}, ContributionVerifierTransactor: ContributionVerifierTransactor{contract: contract}, ContributionVerifierFilterer: ContributionVerifierFilterer{contract: contract}}, nil
}

// NewContributionVerifierCaller creates a new read-only instance of ContributionVerifier, bound to a specific deployed contract.
func NewContributionVerifierCaller(address common.Address, caller bind.ContractCaller) (*ContributionVerifierCaller, error) {
	contract, err := bindContributionVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContributionVerifierCaller{contract: contract}, nil
}

// NewContributionVerifierTransactor creates a new write-only instance of ContributionVerifier, bound to a specific deployed contract.
func NewContributionVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*ContributionVerifierTransactor, error) {
	contract, err := bindContributionVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContributionVerifierTransactor{contract: contract}, nil
}

// NewContributionVerifierFilterer creates a new log filterer instance of ContributionVerifier, bound to a specific deployed contract.
func NewContributionVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*ContributionVerifierFilterer, error) {
	contract, err := bindContributionVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContributionVerifierFilterer{contract: contract}, nil
}

// bindContributionVerifier binds a generic wrapper to an already deployed contract.
func bindContributionVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContributionVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContributionVerifier *ContributionVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContributionVerifier.Contract.ContributionVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContributionVerifier *ContributionVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContributionVerifier.Contract.ContributionVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContributionVerifier *ContributionVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContributionVerifier.Contract.ContributionVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContributionVerifier *ContributionVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContributionVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContributionVerifier *ContributionVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContributionVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContributionVerifier *ContributionVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContributionVerifier.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_ContributionVerifier *ContributionVerifierCaller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int) ([4]*big.Int, error) {
	var out []interface{}
	err := _ContributionVerifier.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_ContributionVerifier *ContributionVerifierSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _ContributionVerifier.Contract.CompressProof(&_ContributionVerifier.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_ContributionVerifier *ContributionVerifierCallerSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _ContributionVerifier.Contract.CompressProof(&_ContributionVerifier.CallOpts, proof)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_ContributionVerifier *ContributionVerifierCaller) ProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContributionVerifier.contract.Call(opts, &out, "provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_ContributionVerifier *ContributionVerifierSession) ProvingKeyHash() ([32]byte, error) {
	return _ContributionVerifier.Contract.ProvingKeyHash(&_ContributionVerifier.CallOpts)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_ContributionVerifier *ContributionVerifierCallerSession) ProvingKeyHash() ([32]byte, error) {
	return _ContributionVerifier.Contract.ProvingKeyHash(&_ContributionVerifier.CallOpts)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x86836a97.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[10] input) view returns()
func (_ContributionVerifier *ContributionVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [10]*big.Int) error {
	var out []interface{}
	err := _ContributionVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x86836a97.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[10] input) view returns()
func (_ContributionVerifier *ContributionVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [10]*big.Int) error {
	return _ContributionVerifier.Contract.VerifyCompressedProof(&_ContributionVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x86836a97.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[10] input) view returns()
func (_ContributionVerifier *ContributionVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [10]*big.Int) error {
	return _ContributionVerifier.Contract.VerifyCompressedProof(&_ContributionVerifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x94e4398a.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[10] input) view returns()
func (_ContributionVerifier *ContributionVerifierCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, input [10]*big.Int) error {
	var out []interface{}
	err := _ContributionVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x94e4398a.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[10] input) view returns()
func (_ContributionVerifier *ContributionVerifierSession) VerifyProof(proof [8]*big.Int, input [10]*big.Int) error {
	return _ContributionVerifier.Contract.VerifyProof(&_ContributionVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x94e4398a.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[10] input) view returns()
func (_ContributionVerifier *ContributionVerifierCallerSession) VerifyProof(proof [8]*big.Int, input [10]*big.Int) error {
	return _ContributionVerifier.Contract.VerifyProof(&_ContributionVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_ContributionVerifier *ContributionVerifierCaller) VerifyProof0(opts *bind.CallOpts, proof []byte, input []byte) error {
	var out []interface{}
	err := _ContributionVerifier.contract.Call(opts, &out, "verifyProof0", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_ContributionVerifier *ContributionVerifierSession) VerifyProof0(proof []byte, input []byte) error {
	return _ContributionVerifier.Contract.VerifyProof0(&_ContributionVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_ContributionVerifier *ContributionVerifierCallerSession) VerifyProof0(proof []byte, input []byte) error {
	return _ContributionVerifier.Contract.VerifyProof0(&_ContributionVerifier.CallOpts, proof, input)
}
