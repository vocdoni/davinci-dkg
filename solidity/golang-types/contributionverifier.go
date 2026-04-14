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
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"input\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidInputEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x608080604052346015576117b2908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c908163233ace1114610e205750806344f6369214610da257806348306671146106fd578063a6047e6c146100b65763b8e72af614610053575f80fd5b346100b25760403660031901126100b25760043567ffffffffffffffff81116100b257610084903690600401610e7f565b6024359167ffffffffffffffff83116100b2576100a86100b0933690600401610e7f565b929091610f5e565b005b5f80fd5b346100b2576102003660031901126100b25736610104116100b25736610204116100b25760405160408101907f2853547a9392d5f4eeb0244598dbd1df91b5180dd481d76e1cae727f9e925675815260208101917f0da718f7b63e1bb1816f95daae813fc404e2994229574a78946565235243ef3583527f2ae58d3d57cac23cf9b1de916e7ab3302a507e53f00ce4a88b31f12542455aaa8152606082017f305139469a39efab092b653b605c449c5e0252c62927deb0cdca93163c7d572f81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f2f82b5788b780bf09174615d170a6cce6169653c758976148e099d54cabb3b5d608087019580875284848460608160075afa911016838860808160065afa167f2c95e3574b4bb87fc9b8c913e873b8456997ea65affca7839eaf067ac5e95d5883527f0ee13fb4314e706189f78621bba918f9da30941fc6e71c7198055faa689f43be86526001610124359182895286868660608160075afa9310161616838860808160065afa167f254c347d61132999ec9e4755c758c0c192b6dfa1bc047c7c516fd1aaee4547d183527f0543c74fc8545133183554a3e7d554d876bccce80d659019fed06fff98bc41c68652610144359081885285858560608160075afa92101616838860808160065afa167f01af09e56ee12d6222e24b4015238ed1004ef67a6de55b8a1b8cf6f584035f2683527f19a1faf11a67d51e7243ce65764180cd689f43ddbc4d31e2006f961e72c552c98652610164359081885285858560608160075afa92101616838860808160065afa167f163a7f0b920b48748043a7b0a2e54e42307f06594dab774fcc7fa8de083d2bec83527f25737cd0f1e21b9d35c5eca57bc9e1f7e98777d2a0958ff87e99213007cde3b78652610184359081885285858560608160075afa92101616838860808160065afa167f2ea50f09947661fd8eb5677ea625dbd7cf3e7fd9294139e48a50ef3e4bdb817f83527ea3699dd3b53180fa30124851cfc095ea3106cff9e2287930ae69a945ed532d86526101a4359081885285858560608160075afa92101616838860808160065afa167f2f29d72225cd3be92a06bafb498421fe05f7a77119bd397e9c3907741489e35283527f2f001e73172502003d2707ecc2d3232292a55b0b0ec00675100688745e8479d086526101c4359081885285858560608160075afa92101616838860808160065afa16947f205972844984651edf043e355347f3f3c17c6f45afd3facff896e03e914974fd8352526101e43580955260608160075afa9210161660408260808160065afa169051915190156106ee5760405191610100600484377f093097eb0e445d6a027d195b7156de58a8781655e9a58ed7577fbe2f6cd4cabb6101008401527f1b42cf4659a4409335900b759776f3d62c4c8881cba9fc7a6db8121c582952ac6101208401527f2b3c881c8f6087d1beafa6b55fc23cf5468e06c3516810d7923ee8c228df6b5f6101408401527f1898c0940b9274ae0101ae127cb2800b8e8e9fc9b4ee131a968f0adb9af379616101608401527f01415ebb93ddea26cd55b98f0ed1d9caf0e7587349cc1376cb9089b5b7e85c736101808401527f02729135a06fabf0fe31c44c022c71ba26e85afbfeebd945c068605169be117c6101a08401527f1e5d9628a4948a70df243191d032859da7c2b1c2cde4fafbae547db15d573c576101c08401527f1056677abb5ee09b4dc3abf829432afddec9f61076d31ef2406098dd502312386101e08401527f04dbba85b2bb47cbc6c6a32ff37bcbed352c78baff61616eaab77f0d78f26cfc6102008401527f06c418d207699c152893d997e1312a994f587214bbd69e63fdd28bb3c18935966102208401526102408301526102608201527f1f12c85d89347e8b231f564fa7dea658c94f16b2f0948dd34ab1f73c78b0a1d46102808201527f1cbf5e4688806074b2625bca28a30ad2ce1279817e91c94efaa5fa52cd1b2fca6102a08201527f2dd5ef3d375a16f7f98a63200c5bc9404e1a4032d058ff7f76a692735cf843c46102c08201527f1156e3a8669a5365541b50b1e6e2a3e4f79e776b0b347ff63ca65e11b9f183006102e08201526020816103008160085afa905116156106df57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b346100b2576101803660031901126100b257366084116100b25736610184116100b2576103006040516107308282610ead565b8136823761073f60043561137c565b6107506024939293356044356113e7565b9193929061075f60643561137c565b9390926040519660408801967f2853547a9392d5f4eeb0244598dbd1df91b5180dd481d76e1cae727f9e92567589528860208101987f0da718f7b63e1bb1816f95daae813fc404e2994229574a78946565235243ef358a527f2ae58d3d57cac23cf9b1de916e7ab3302a507e53f00ce4a88b31f12542455aaa81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f305139469a39efab092b653b605c449c5e0252c62927deb0cdca93163c7d572f84527f2f82b5788b780bf09174615d170a6cce6169653c758976148e099d54cabb3b5d6084359583608082019780895286828660608160075afa911016818360808160065afa167f2c95e3574b4bb87fc9b8c913e873b8456997ea65affca7839eaf067ac5e95d5885527f0ee13fb4314e706189f78621bba918f9da30941fc6e71c7198055faa689f43be8852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f254c347d61132999ec9e4755c758c0c192b6dfa1bc047c7c516fd1aaee4547d185527f0543c74fc8545133183554a3e7d554d876bccce80d659019fed06fff98bc41c6885260c43590818a5287838760608160075afa92101616818360808160065afa167f01af09e56ee12d6222e24b4015238ed1004ef67a6de55b8a1b8cf6f584035f2685527f19a1faf11a67d51e7243ce65764180cd689f43ddbc4d31e2006f961e72c552c9885260e43590818a5287838760608160075afa92101616818360808160065afa167f163a7f0b920b48748043a7b0a2e54e42307f06594dab774fcc7fa8de083d2bec85527f25737cd0f1e21b9d35c5eca57bc9e1f7e98777d2a0958ff87e99213007cde3b788526101043590818a5287838760608160075afa92101616818360808160065afa167f2ea50f09947661fd8eb5677ea625dbd7cf3e7fd9294139e48a50ef3e4bdb817f85527ea3699dd3b53180fa30124851cfc095ea3106cff9e2287930ae69a945ed532d88526101243590818a5287838760608160075afa92101616818360808160065afa167f2f29d72225cd3be92a06bafb498421fe05f7a77119bd397e9c3907741489e35285527f2f001e73172502003d2707ecc2d3232292a55b0b0ec00675100688745e8479d088526101443590818a5287838760608160075afa921016169160808160065afa16947f205972844984651edf043e355347f3f3c17c6f45afd3facff896e03e914974fd8352526101643580955260608160075afa9210161660408a60808160065afa169851975198156106ee5760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f093097eb0e445d6a027d195b7156de58a8781655e9a58ed7577fbe2f6cd4cabb6101008401527f1b42cf4659a4409335900b759776f3d62c4c8881cba9fc7a6db8121c582952ac6101208401527f2b3c881c8f6087d1beafa6b55fc23cf5468e06c3516810d7923ee8c228df6b5f6101408401527f1898c0940b9274ae0101ae127cb2800b8e8e9fc9b4ee131a968f0adb9af379616101608401527f01415ebb93ddea26cd55b98f0ed1d9caf0e7587349cc1376cb9089b5b7e85c736101808401527f02729135a06fabf0fe31c44c022c71ba26e85afbfeebd945c068605169be117c6101a08401527f1e5d9628a4948a70df243191d032859da7c2b1c2cde4fafbae547db15d573c576101c08401527f1056677abb5ee09b4dc3abf829432afddec9f61076d31ef2406098dd502312386101e08401527f04dbba85b2bb47cbc6c6a32ff37bcbed352c78baff61616eaab77f0d78f26cfc6102008401527f06c418d207699c152893d997e1312a994f587214bbd69e63fdd28bb3c18935966102208401526102408301526102608201527f1f12c85d89347e8b231f564fa7dea658c94f16b2f0948dd34ab1f73c78b0a1d46102808201527f1cbf5e4688806074b2625bca28a30ad2ce1279817e91c94efaa5fa52cd1b2fca6102a08201527f2dd5ef3d375a16f7f98a63200c5bc9404e1a4032d058ff7f76a692735cf843c46102c08201527f1156e3a8669a5365541b50b1e6e2a3e4f79e776b0b347ff63ca65e11b9f183006102e0820152604051928391610d7d8484610ead565b8336843760085afa15908115610d95575b506106df57005b6001915051141581610d8e565b346100b2576101003660031901126100b25736610104116100b2576080604051610dcc8282610ead565b81368237610dde6024356004356110a7565b8152610df460843560a435604435606435611148565b60208301526040820152610e0c60e43560c4356110a7565b6060820152610e1e6040518092610e58565bf35b346100b2575f3660031901126100b257807fe662b9f5b48e5c825efe2667a1b8e144cd262f125c06fcc794f6f198dea9f15360209252f35b905f905b60048210610e6957505050565b6020806001928551815201930191019091610e5c565b9181601f840112156100b25782359167ffffffffffffffff83116100b257602083818601950101116100b257565b90601f8019910116810190811067ffffffffffffffff821117610ecf57604052565b634e487b7160e01b5f52604160045260245ffd5b90610100828203126100b25780601f830112156100b25760405191610f0a61010084610ead565b829061010081019283116100b257905b828210610f275750505090565b8135815260209182019101610f1a565b905f905b60088210610f4857505050565b6020806001928551815201930191019091610f3b565b9392919061010081146110405760808114610f825763236bd13760e01b5f5260045ffd5b6101008303611031578401916080858403126100b25782601f860112156100b25760405192610fb2608085610ead565b8395608081019182116100b257955b81871061102157505061101f93945081610fe091611009930190610ee3565b604051634830667160e01b602082015292610fff906024850190610e58565b60a4830190610f37565b610184815261101a6101a482610ead565b611520565b565b8635815260209687019601610fc1565b630c0b7e3560e11b5f5260045ffd5b90929361010083036110315761101f93611063826110969461106c940190610ee3565b93810190610ee3565b6040516329811f9b60e21b60208201529261108b906024850190610f37565b610124830190610f37565b610204815261101a61022482610ead565b905f51602061175d5f395f51905f528210801590611131575b6106df57811580611129575b611123576110f05f51602061175d5f395f51905f526003818581818009090861157d565b8181036110ff57505060011b90565b5f51602061175d5f395f51905f52809106810306145f146106df57600190811b1790565b50505f90565b5080156110cc565b505f51602061175d5f395f51905f528110156110c0565b919093925f51602061175d5f395f51905f528310801590611365575b801561134e575b8015611337575b6106df57808286851717171561132c5790829161128f5f51602061175d5f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f51602061175d5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161126981808b8009818780090861157d565b8408095f51602061175d5f395f51905f52611283826116f4565b800914159586916115a0565b929080821480611323575b156112c15750505050905f146112b95760ff60025b169060021b179190565b60ff5f6112af565b5f51602061175d5f395f51905f52809106810306149182611304575b5050156106df57600191156112fc5760ff60025b169060021b17179190565b60ff5f6112f1565b5f51602061175d5f395f51905f52919250819006810306145f806112dd565b5083831461129a565b50505090505f905f90565b505f51602061175d5f395f51905f52811015611172565b505f51602061175d5f395f51905f5282101561116b565b505f51602061175d5f395f51905f52851015611164565b80156113e0578060011c915f51602061175d5f395f51905f528310156106df576001806113bf5f51602061175d5f395f51905f526003818881818009090861157d565b9316146113c857565b905f51602061175d5f395f51905f5280910681030690565b505f905f90565b801580611518575b61150c578060021c92825f51602061175d5f395f51905f5285108015906114f5575b6106df5784815f51602061175d5f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816114bf9d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5086115a0565b809291600180829616146114d1575050565b5f51602061175d5f395f51905f528093945080929550809106810306930681030690565b505f51602061175d5f395f51905f52811015611411565b50505f905f905f905f90565b5081156113ef565b5f8091602081519101305afa3d15611575573d9067ffffffffffffffff8211610ecf576040519161155b601f8201601f191660200184610ead565b82523d5f602084013e5b1561156d5750565b602081519101fd5b606090611565565b90611587826116f4565b915f51602061175d5f395f51905f52838009036106df57565b915f51602061175d5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816115f8939694966115ea82808a8009818a80090861157d565b906116e8575b86080961157d565b925f51602061175d5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f51602061175d5f395f51905f5260a083015260208260c08160055afa915191156106df575f51602061175d5f395f51905f528260019209036106df575f51602061175d5f395f51905f52908209925f51602061175d5f395f51905f5280808087800906810306818780090814908115916116c9575b506106df57565b90505f51602061175d5f395f51905f528084860960020914155f6116c2565b818091068103066115f0565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f51602061175d5f395f51905f5260a083015260208260c08160055afa915191156106df5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a26469706673582212206d5e667676b2f8f10f6d6388165db3bf849e263a6db0a4347a938742be54a46f64736f6c634300081c0033",
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

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x48306671.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [8]*big.Int) error {
	var out []interface{}
	err := _ContributionVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x48306671.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [8]*big.Int) error {
	return _ContributionVerifier.Contract.VerifyCompressedProof(&_ContributionVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x48306671.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [8]*big.Int) error {
	return _ContributionVerifier.Contract.VerifyCompressedProof(&_ContributionVerifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xa6047e6c.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, input [8]*big.Int) error {
	var out []interface{}
	err := _ContributionVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0xa6047e6c.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierSession) VerifyProof(proof [8]*big.Int, input [8]*big.Int) error {
	return _ContributionVerifier.Contract.VerifyProof(&_ContributionVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xa6047e6c.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierCallerSession) VerifyProof(proof [8]*big.Int, input [8]*big.Int) error {
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
