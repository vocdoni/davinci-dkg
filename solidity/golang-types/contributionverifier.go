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
	Bin: "0x608080604052346015576119a9908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c908163233ace1114610fb85750806344f6369214610f3a57806386836a97146107c957806394e4398a146100b65763b8e72af614610053575f80fd5b346100b25760403660031901126100b25760043567ffffffffffffffff81116100b257610084903690600401611017565b6024359167ffffffffffffffff83116100b2576100a86100b0933690600401611017565b9290916110f6565b005b5f80fd5b346100b2576102403660031901126100b25736610104116100b25736610244116100b25760405160408101907f1a4ce018958b648ea8e90255d05fa8a06f500c8f4bb5bdc2dbfb8c94384abc3f815260208101917f0124b07a81da24dcb3b44df47070b67487505f484a3764de6d46c527c20b4a8c83527f19f93ab8397298cba1be5ed3aaf760cf984a0ee3a3c909fbeaa4305c4bc616bd8152606082017f1455cd75657373e20c8a59e59efe209c29288f15cd4b9a327a76f867b413955d81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f2d7f799faf433b62041c82eda9732260550a38221e0f88a3e38bdbe51e4b9f6b608087019580875284848460608160075afa911016838860808160065afa167f02649c61640bb89d6ebe5307c9812ee0a98e07673f62279b66d2249f8799161183527f1394b02a375e8206a35e987064713748a0ca5eea650dc5f65e51610096abf2fd86526001610124359182895286868660608160075afa9310161616838860808160065afa167f1e1d44d64eb9acbabd8b2bdecb1aa0ce4916109d2fd2e3360b65c9a9d2e40b6f83527f2c45121fd613e2f56b6c1af83b8609dd69d31147ff9a7952989a27c937adbd528652610144359081885285858560608160075afa92101616838860808160065afa167f16481f880aace08ed61fc20f881debd9365e1100d08c3b6a1353f6b4359556c183527f133703ae79c040a9993dfbc150ae1e2d1143ffa1240b659ab0246b4441450e988652610164359081885285858560608160075afa92101616838860808160065afa167f017ded42a81158d5f32d8903c4660c1c767a40371f6f305a6d36b0e4dbd5282183527f20577c1a2c0b77dbe570445e6737d2b34b57c90baa180da3f3d7c7e4d51e01628652610184359081885285858560608160075afa92101616838860808160065afa167f09eca315e67305984c7c89152202ee4a4a2d2ffdaefd9b02188df0f1fa38ad4f83527f0b5c99414c132f5319fbc9a35922a886514d0034108a7575c32f6ef310f5613986526101a4359081885285858560608160075afa92101616838860808160065afa167f234a9919f30c2346b3636ddca6855470d9f8b1fc5a7c7499970f78fe9cb5202983527f24537a90ac7aeca65b30dc2af151195756c1aa1845ea839c202f5c924637123786526101c4359081885285858560608160075afa92101616838860808160065afa167f18e36d157b0248d13548dcbbade6c0e97aea488869e568f7e3227a7a422519e483527f06eb2a59dcf0554a3890814b21c7c68d7cdc7f0ea65deecd569299b5b2dcf71f86526101e4359081885285858560608160075afa92101616838860808160065afa167f20d485d8ecd782f4b25c1a87050b15bfb6ba860b354e7edea9299d488d77e99b83527f2f3eb5a11edad72ea77f0df0a8a44487b1f732c9f87fbc4ea541ac10856bd0df8652610204359081885285858560608160075afa92101616838860808160065afa16947f2bd822cb93e1dfa10cb8c9435c82bd80fae41e9ba72cc9762f67378bdcd08ea68352526102243580955260608160075afa9210161660408260808160065afa169051915190156107ba5760405191610100600484377f13701fb5f6b37a1246cbeb200a97c451833a9662013f867d409885aed1918d536101008401527f2da3e5cb10f2f84674fe164381d5792a1c7315656f1df046db03b72c6fab61e56101208401527f05a6482852cb3ca265e33d448e4c797a94168ad8626db0bc448c5b2850e362126101408401527f05508c0898ecc45a3f9e1b01380f67f2adfa9375f963f3e26c97eddc5a3337e26101608401527f2548a34bc187c01c66010a5a8dd8df2f3a93e61c4a2b10a10b0a8b9338b737bf6101808401527f275bf2fd1347f2b12e0f1e199eedbe8d5a743bb091663e8f3ee7696b74ce83cb6101a08401527f2f939b0f8f6de4385c014e6ee332df86ceed043b0e1337409385aac19f4132266101c08401527f0dbb3b7846f235da796b410d31c79b7fbc4046c2f0048700d97bc176d0e21c9e6101e08401527f253f29c8d9b6fdeb0a6942a8fedb5e285f57dfdf93d4e8a4e8dddf26308686c06102008401527f27484467f52de1090b22068ae757abe3f93e0c2e00d3305e6af1a758e6d6532d6102208401526102408301526102608201527f2ef45856706eb494beeaad777245b075cad4dc1d9d411deecfc6baa2c1505ab46102808201527f191138bc6f22d3cac6dec4a9ee288bc7cb1ae1719a25558eaee9a5e35f23281b6102a08201527f2fe58d0c2ba1b0805b4fc1c3016d8ade3d60fbb9b9b458746df7cb41c90ffd156102c08201527e499a905c745074f0a3e35f3e85ee6a3c9f46ab1cc9b868188339f3bd4205856102e08201526020816103008160085afa905116156107ab57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b346100b2576101c03660031901126100b257366084116100b257366101c4116100b2576103006040516107fc8282611045565b8136823761080b600435611573565b61081c6024939293356044356115de565b9193929061082b606435611573565b9390926040519660408801967f1a4ce018958b648ea8e90255d05fa8a06f500c8f4bb5bdc2dbfb8c94384abc3f89528860208101987f0124b07a81da24dcb3b44df47070b67487505f484a3764de6d46c527c20b4a8c8a527f19f93ab8397298cba1be5ed3aaf760cf984a0ee3a3c909fbeaa4305c4bc616bd81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f1455cd75657373e20c8a59e59efe209c29288f15cd4b9a327a76f867b413955d84527f2d7f799faf433b62041c82eda9732260550a38221e0f88a3e38bdbe51e4b9f6b6084359583608082019780895286828660608160075afa911016818360808160065afa167f02649c61640bb89d6ebe5307c9812ee0a98e07673f62279b66d2249f8799161185527f1394b02a375e8206a35e987064713748a0ca5eea650dc5f65e51610096abf2fd8852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f1e1d44d64eb9acbabd8b2bdecb1aa0ce4916109d2fd2e3360b65c9a9d2e40b6f85527f2c45121fd613e2f56b6c1af83b8609dd69d31147ff9a7952989a27c937adbd52885260c43590818a5287838760608160075afa92101616818360808160065afa167f16481f880aace08ed61fc20f881debd9365e1100d08c3b6a1353f6b4359556c185527f133703ae79c040a9993dfbc150ae1e2d1143ffa1240b659ab0246b4441450e98885260e43590818a5287838760608160075afa92101616818360808160065afa167f017ded42a81158d5f32d8903c4660c1c767a40371f6f305a6d36b0e4dbd5282185527f20577c1a2c0b77dbe570445e6737d2b34b57c90baa180da3f3d7c7e4d51e016288526101043590818a5287838760608160075afa92101616818360808160065afa167f09eca315e67305984c7c89152202ee4a4a2d2ffdaefd9b02188df0f1fa38ad4f85527f0b5c99414c132f5319fbc9a35922a886514d0034108a7575c32f6ef310f5613988526101243590818a5287838760608160075afa92101616818360808160065afa167f234a9919f30c2346b3636ddca6855470d9f8b1fc5a7c7499970f78fe9cb5202985527f24537a90ac7aeca65b30dc2af151195756c1aa1845ea839c202f5c924637123788526101443590818a5287838760608160075afa92101616818360808160065afa167f18e36d157b0248d13548dcbbade6c0e97aea488869e568f7e3227a7a422519e485527f06eb2a59dcf0554a3890814b21c7c68d7cdc7f0ea65deecd569299b5b2dcf71f88526101643590818a5287838760608160075afa92101616818360808160065afa167f20d485d8ecd782f4b25c1a87050b15bfb6ba860b354e7edea9299d488d77e99b85527f2f3eb5a11edad72ea77f0df0a8a44487b1f732c9f87fbc4ea541ac10856bd0df88526101843590818a5287838760608160075afa921016169160808160065afa16947f2bd822cb93e1dfa10cb8c9435c82bd80fae41e9ba72cc9762f67378bdcd08ea68352526101a43580955260608160075afa9210161660408a60808160065afa169851975198156107ba5760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f13701fb5f6b37a1246cbeb200a97c451833a9662013f867d409885aed1918d536101008401527f2da3e5cb10f2f84674fe164381d5792a1c7315656f1df046db03b72c6fab61e56101208401527f05a6482852cb3ca265e33d448e4c797a94168ad8626db0bc448c5b2850e362126101408401527f05508c0898ecc45a3f9e1b01380f67f2adfa9375f963f3e26c97eddc5a3337e26101608401527f2548a34bc187c01c66010a5a8dd8df2f3a93e61c4a2b10a10b0a8b9338b737bf6101808401527f275bf2fd1347f2b12e0f1e199eedbe8d5a743bb091663e8f3ee7696b74ce83cb6101a08401527f2f939b0f8f6de4385c014e6ee332df86ceed043b0e1337409385aac19f4132266101c08401527f0dbb3b7846f235da796b410d31c79b7fbc4046c2f0048700d97bc176d0e21c9e6101e08401527f253f29c8d9b6fdeb0a6942a8fedb5e285f57dfdf93d4e8a4e8dddf26308686c06102008401527f27484467f52de1090b22068ae757abe3f93e0c2e00d3305e6af1a758e6d6532d6102208401526102408301526102608201527f2ef45856706eb494beeaad777245b075cad4dc1d9d411deecfc6baa2c1505ab46102808201527f191138bc6f22d3cac6dec4a9ee288bc7cb1ae1719a25558eaee9a5e35f23281b6102a08201527f2fe58d0c2ba1b0805b4fc1c3016d8ade3d60fbb9b9b458746df7cb41c90ffd156102c08201527e499a905c745074f0a3e35f3e85ee6a3c9f46ab1cc9b868188339f3bd4205856102e0820152604051928391610f158484611045565b8336843760085afa15908115610f2d575b506107ab57005b6001915051141581610f26565b346100b2576101003660031901126100b25736610104116100b2576080604051610f648282611045565b81368237610f7660243560043561129e565b8152610f8c60843560a43560443560643561133f565b60208301526040820152610fa460e43560c43561129e565b6060820152610fb66040518092610ff0565bf35b346100b2575f3660031901126100b257807fe662b9f5b48e5c825efe2667a1b8e144cd262f125c06fcc794f6f198dea9f15360209252f35b905f905b6004821061100157505050565b6020806001928551815201930191019091610ff4565b9181601f840112156100b25782359167ffffffffffffffff83116100b257602083818601950101116100b257565b90601f8019910116810190811067ffffffffffffffff82111761106757604052565b634e487b7160e01b5f52604160045260245ffd5b90610140828203126100b25780601f830112156100b257604051916110a261014084611045565b829061014081019283116100b257905b8282106110bf5750505090565b81358152602091820191016110b2565b905f905b600a82106110e057505050565b60208060019285518152019301910190916110d3565b9392919061010081146111d8576080811461111a5763236bd13760e01b5f5260045ffd5b61014083036111c9578401916080858403126100b25782601f860112156100b2576040519261114a608085611045565b8395608081019182116100b257955b8187106111b95750506111b793945081611178916111a193019061107b565b6040516386836a9760e01b602082015292611197906024850190610ff0565b60a48301906110cf565b6101c481526111b26101e482611045565b611717565b565b8635815260209687019601611159565b630c0b7e3560e11b5f5260045ffd5b61014083036111c957840191610100858403126100b25782601f860112156100b2576040519261120a61010085611045565b839561010081019182116100b257955b81871061128e5750506112329293945081019061107b565b604051634a721cc560e11b6020820152915f602484015b6008821061127857505050906112676111b7926101248301906110cf565b61024481526111b261026482611045565b6020806001928551815201930191019091611249565b863581526020968701960161121a565b905f5160206119545f395f51905f528210801590611328575b6107ab57811580611320575b61131a576112e75f5160206119545f395f51905f5260038185818180090908611774565b8181036112f657505060011b90565b5f5160206119545f395f51905f52809106810306145f146107ab57600190811b1790565b50505f90565b5080156112c3565b505f5160206119545f395f51905f528110156112b7565b919093925f5160206119545f395f51905f52831080159061155c575b8015611545575b801561152e575b6107ab578082868517171715611523579082916114865f5160206119545f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f5160206119545f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161146081808b80098187800908611774565b8408095f5160206119545f395f51905f5261147a826118eb565b80091415958691611797565b92908082148061151a575b156114b85750505050905f146114b05760ff60025b169060021b179190565b60ff5f6114a6565b5f5160206119545f395f51905f528091068103061491826114fb575b5050156107ab57600191156114f35760ff60025b169060021b17179190565b60ff5f6114e8565b5f5160206119545f395f51905f52919250819006810306145f806114d4565b50838314611491565b50505090505f905f90565b505f5160206119545f395f51905f52811015611369565b505f5160206119545f395f51905f52821015611362565b505f5160206119545f395f51905f5285101561135b565b80156115d7578060011c915f5160206119545f395f51905f528310156107ab576001806115b65f5160206119545f395f51905f5260038188818180090908611774565b9316146115bf57565b905f5160206119545f395f51905f5280910681030690565b505f905f90565b80158061170f575b611703578060021c92825f5160206119545f395f51905f5285108015906116ec575b6107ab5784815f5160206119545f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816116b69d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611797565b809291600180829616146116c8575050565b5f5160206119545f395f51905f528093945080929550809106810306930681030690565b505f5160206119545f395f51905f52811015611608565b50505f905f905f905f90565b5081156115e6565b5f8091602081519101305afa3d1561176c573d9067ffffffffffffffff82116110675760405191611752601f8201601f191660200184611045565b82523d5f602084013e5b156117645750565b602081519101fd5b60609061175c565b9061177e826118eb565b915f5160206119545f395f51905f52838009036107ab57565b915f5160206119545f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816117ef939694966117e182808a8009818a800908611774565b906118df575b860809611774565b925f5160206119545f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206119545f395f51905f5260a083015260208260c08160055afa915191156107ab575f5160206119545f395f51905f528260019209036107ab575f5160206119545f395f51905f52908209925f5160206119545f395f51905f5280808087800906810306818780090814908115916118c0575b506107ab57565b90505f5160206119545f395f51905f528084860960020914155f6118b9565b818091068103066117e7565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206119545f395f51905f5260a083015260208260c08160055afa915191156107ab5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220216644fa5c76c5e628286cb6854a718c07ff149a903fda1f83ba4fbb4e66b76f64736f6c634300081c0033",
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
