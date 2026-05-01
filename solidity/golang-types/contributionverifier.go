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
	Bin: "0x608080604052346015576119ab908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c908163233ace1114610fba5750806344f6369214610f3c57806386836a97146107ca57806394e4398a146100b65763b8e72af614610053575f80fd5b346100b25760403660031901126100b25760043567ffffffffffffffff81116100b257610084903690600401611019565b6024359167ffffffffffffffff83116100b2576100a86100b0933690600401611019565b9290916110f8565b005b5f80fd5b346100b2576102403660031901126100b25736610104116100b25736610244116100b25760405160408101907f2e6a7f620caa46da297b2dc8442ad836ab1e4e14aacf2ff3e024b7427a3d17d3815260208101917f2e4f821ba5be6852b2108882ea4da5b1a52cffff5ed619728dbd1b82b04eef7983527f18ff6d7aa5289f427a9007ad4e4027563fb0e84e1b541ff5b187ca7c61b4cb778152606082017f2cdfac3a84d951467dd822d5bdbdb6e290c0cc0a92195af211ee64a400013ba381527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f15e03a6351fa81acb913ac497a699abc2012403bf29c6eeb67e7796927ae98e8608087019580875284848460608160075afa911016838860808160065afa167f14c63f48b63251a2a89a1fae11c793040bd8290c5b5c53bd0141f3149a16bb9f83527f16c6ef47cd3655ea2b79aae20e86de737d0421de57ce13e3dcc5fb3728f13eba86526001610124359182895286868660608160075afa9310161616838860808160065afa167f1014b59ed009bee22bf798c9b449f9f872c21089bc1b08e397e29c1a0fcdea0b83527f21e812b865f82b14619b57aaf6dc57ea8a1595d1a3f7b5416328600db3f632758652610144359081885285858560608160075afa92101616838860808160065afa167f28dcb2449e9dcc89e22167c98ab738552f4b0d9dcc2211f617a9fab6284b0ca283527f2b91f2e6d71a49c7013132f6f1ed64f4ce0213c35c60fbe0fd7c4da33e78e1ce8652610164359081885285858560608160075afa92101616838860808160065afa167f2bc5259d6441b486bb8106863167f39515f872b04ee4eff2c0e210324d1fac2683527f0b165e353cc64696e3e55220c1ed8987b28f931aa4ad4677a486fe5d040406228652610184359081885285858560608160075afa92101616838860808160065afa167f09e29aeeba77d827047970f0b4d6f2b447b12dace2c4f17da6e7125b8bc4e54683527f1aa23863bb80f01c8cc30f4acf5df5ece4d5b2122d58f6c20caf08169645097286526101a4359081885285858560608160075afa92101616838860808160065afa167f2b73838cad7c0e52bc249b07e2414a0664dc63518efca6f0f67a4765615ad1f683527f2eefcd79209bb570aeb20f7dd6bec51b42ce7d40a3f4ea25ca7a91dc3fcdf14686526101c4359081885285858560608160075afa92101616838860808160065afa167f1a5a443cf3e51a5715d62b7f501be40d08c42ea241b6230ea4592d149e6f23d883527f1f847a19c86216f4d664f952302bb8e7d125587f255972dba79b990348ca64f786526101e4359081885285858560608160075afa92101616838860808160065afa167f290832d9609ccd62d520329bb766d07d7d614388b470789c5547846d74a7b93083527f193afb1629491caa6a86cceb343905a8cf1d2f53890629b6e2d61a189bf48e248652610204359081885285858560608160075afa92101616838860808160065afa16947f1a56ac9547ac6cb773e65733197189003c88f9b801cc3b1401df684e3111cb7c8352526102243580955260608160075afa9210161660408260808160065afa169051915190156107bb5760405191610100600484377f1d832ef5edec7a607e88320f9834502fbd866b9e8a9acd4bef12c1ec743f0a396101008401527f1d925bad7fc71ed2aef762d974bbda336ac5e2bccb1a041c50d350c6d23c71466101208401527f155fe27725c1b1a9bf3130891f7356ffe0ef84a6fdf77b9b887c60f30221d0ca6101408401527f0e84f0b89cb626dcc1ca1f21032b567d54f6abc7e8fa7384e1d460da6883ff2e6101608401527f10f1788c5726845b289d565ba4cc76a4fdbd4fe50ba38d39c2b93e569401cfaa6101808401527f04485a267ee87d29fe9db33cba40a03f63025725219b22b06d45682a399bb9826101a08401527f1f6c57aa0dcfbb2ff32ddee3e13b45717cbf8c555c300e83fb7a4644842e41376101c08401527f20e8dfc1b82f9f377464c7d914c2d6a222567270cbaba9d04b9e7ae2b57ddcb06101e08401527f0b60ec864ab77a986265995d7a8fab209dbbfaeb7a97d2ae3807e45e0725e4a06102008401527f23745c02485e2762a8e10f19f03ab8cccd2f744beeab7dbd702868109e56668c6102208401526102408301526102608201527f2583ca627d81e116d701c497b64e3d3ba2413adf307af88e65200776b48c917f6102808201527f22c43cd60f5b100a4cbf9ff999de9a25badd08dee74e6ede961026ffb173c43d6102a08201527f041ecf3ecd5cfd09b0f3226f08349654103b27614e778146346db8933696322a6102c08201527f2c072273d95bead57657e9921632bacbe3365c8408c021a82df1415bf40514386102e08201526020816103008160085afa905116156107ac57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b346100b2576101c03660031901126100b257366084116100b257366101c4116100b2576103006040516107fd8282611047565b8136823761080c600435611575565b61081d6024939293356044356115e0565b9193929061082c606435611575565b9390926040519660408801967f2e6a7f620caa46da297b2dc8442ad836ab1e4e14aacf2ff3e024b7427a3d17d389528860208101987f2e4f821ba5be6852b2108882ea4da5b1a52cffff5ed619728dbd1b82b04eef798a527f18ff6d7aa5289f427a9007ad4e4027563fb0e84e1b541ff5b187ca7c61b4cb7781527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f2cdfac3a84d951467dd822d5bdbdb6e290c0cc0a92195af211ee64a400013ba384527f15e03a6351fa81acb913ac497a699abc2012403bf29c6eeb67e7796927ae98e86084359583608082019780895286828660608160075afa911016818360808160065afa167f14c63f48b63251a2a89a1fae11c793040bd8290c5b5c53bd0141f3149a16bb9f85527f16c6ef47cd3655ea2b79aae20e86de737d0421de57ce13e3dcc5fb3728f13eba8852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f1014b59ed009bee22bf798c9b449f9f872c21089bc1b08e397e29c1a0fcdea0b85527f21e812b865f82b14619b57aaf6dc57ea8a1595d1a3f7b5416328600db3f63275885260c43590818a5287838760608160075afa92101616818360808160065afa167f28dcb2449e9dcc89e22167c98ab738552f4b0d9dcc2211f617a9fab6284b0ca285527f2b91f2e6d71a49c7013132f6f1ed64f4ce0213c35c60fbe0fd7c4da33e78e1ce885260e43590818a5287838760608160075afa92101616818360808160065afa167f2bc5259d6441b486bb8106863167f39515f872b04ee4eff2c0e210324d1fac2685527f0b165e353cc64696e3e55220c1ed8987b28f931aa4ad4677a486fe5d0404062288526101043590818a5287838760608160075afa92101616818360808160065afa167f09e29aeeba77d827047970f0b4d6f2b447b12dace2c4f17da6e7125b8bc4e54685527f1aa23863bb80f01c8cc30f4acf5df5ece4d5b2122d58f6c20caf08169645097288526101243590818a5287838760608160075afa92101616818360808160065afa167f2b73838cad7c0e52bc249b07e2414a0664dc63518efca6f0f67a4765615ad1f685527f2eefcd79209bb570aeb20f7dd6bec51b42ce7d40a3f4ea25ca7a91dc3fcdf14688526101443590818a5287838760608160075afa92101616818360808160065afa167f1a5a443cf3e51a5715d62b7f501be40d08c42ea241b6230ea4592d149e6f23d885527f1f847a19c86216f4d664f952302bb8e7d125587f255972dba79b990348ca64f788526101643590818a5287838760608160075afa92101616818360808160065afa167f290832d9609ccd62d520329bb766d07d7d614388b470789c5547846d74a7b93085527f193afb1629491caa6a86cceb343905a8cf1d2f53890629b6e2d61a189bf48e2488526101843590818a5287838760608160075afa921016169160808160065afa16947f1a56ac9547ac6cb773e65733197189003c88f9b801cc3b1401df684e3111cb7c8352526101a43580955260608160075afa9210161660408a60808160065afa169851975198156107bb5760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f1d832ef5edec7a607e88320f9834502fbd866b9e8a9acd4bef12c1ec743f0a396101008401527f1d925bad7fc71ed2aef762d974bbda336ac5e2bccb1a041c50d350c6d23c71466101208401527f155fe27725c1b1a9bf3130891f7356ffe0ef84a6fdf77b9b887c60f30221d0ca6101408401527f0e84f0b89cb626dcc1ca1f21032b567d54f6abc7e8fa7384e1d460da6883ff2e6101608401527f10f1788c5726845b289d565ba4cc76a4fdbd4fe50ba38d39c2b93e569401cfaa6101808401527f04485a267ee87d29fe9db33cba40a03f63025725219b22b06d45682a399bb9826101a08401527f1f6c57aa0dcfbb2ff32ddee3e13b45717cbf8c555c300e83fb7a4644842e41376101c08401527f20e8dfc1b82f9f377464c7d914c2d6a222567270cbaba9d04b9e7ae2b57ddcb06101e08401527f0b60ec864ab77a986265995d7a8fab209dbbfaeb7a97d2ae3807e45e0725e4a06102008401527f23745c02485e2762a8e10f19f03ab8cccd2f744beeab7dbd702868109e56668c6102208401526102408301526102608201527f2583ca627d81e116d701c497b64e3d3ba2413adf307af88e65200776b48c917f6102808201527f22c43cd60f5b100a4cbf9ff999de9a25badd08dee74e6ede961026ffb173c43d6102a08201527f041ecf3ecd5cfd09b0f3226f08349654103b27614e778146346db8933696322a6102c08201527f2c072273d95bead57657e9921632bacbe3365c8408c021a82df1415bf40514386102e0820152604051928391610f178484611047565b8336843760085afa15908115610f2f575b506107ac57005b6001915051141581610f28565b346100b2576101003660031901126100b25736610104116100b2576080604051610f668282611047565b81368237610f786024356004356112a0565b8152610f8e60843560a435604435606435611341565b60208301526040820152610fa660e43560c4356112a0565b6060820152610fb86040518092610ff2565bf35b346100b2575f3660031901126100b257807f5c77f8f911a3b7933667cc03856a40146de22b1461006e333f6e623b350bb48c60209252f35b905f905b6004821061100357505050565b6020806001928551815201930191019091610ff6565b9181601f840112156100b25782359167ffffffffffffffff83116100b257602083818601950101116100b257565b90601f8019910116810190811067ffffffffffffffff82111761106957604052565b634e487b7160e01b5f52604160045260245ffd5b90610140828203126100b25780601f830112156100b257604051916110a461014084611047565b829061014081019283116100b257905b8282106110c15750505090565b81358152602091820191016110b4565b905f905b600a82106110e257505050565b60208060019285518152019301910190916110d5565b9392919061010081146111da576080811461111c5763236bd13760e01b5f5260045ffd5b61014083036111cb578401916080858403126100b25782601f860112156100b2576040519261114c608085611047565b8395608081019182116100b257955b8187106111bb5750506111b99394508161117a916111a393019061107d565b6040516386836a9760e01b602082015292611199906024850190610ff2565b60a48301906110d1565b6101c481526111b46101e482611047565b611719565b565b863581526020968701960161115b565b630c0b7e3560e11b5f5260045ffd5b61014083036111cb57840191610100858403126100b25782601f860112156100b2576040519261120c61010085611047565b839561010081019182116100b257955b8187106112905750506112349293945081019061107d565b604051634a721cc560e11b6020820152915f602484015b6008821061127a57505050906112696111b9926101248301906110d1565b61024481526111b461026482611047565b602080600192855181520193019101909161124b565b863581526020968701960161121c565b905f5160206119565f395f51905f52821080159061132a575b6107ac57811580611322575b61131c576112e95f5160206119565f395f51905f5260038185818180090908611776565b8181036112f857505060011b90565b5f5160206119565f395f51905f52809106810306145f146107ac57600190811b1790565b50505f90565b5080156112c5565b505f5160206119565f395f51905f528110156112b9565b919093925f5160206119565f395f51905f52831080159061155e575b8015611547575b8015611530575b6107ac578082868517171715611525579082916114885f5160206119565f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f5160206119565f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161146281808b80098187800908611776565b8408095f5160206119565f395f51905f5261147c826118ed565b80091415958691611799565b92908082148061151c575b156114ba5750505050905f146114b25760ff60025b169060021b179190565b60ff5f6114a8565b5f5160206119565f395f51905f528091068103061491826114fd575b5050156107ac57600191156114f55760ff60025b169060021b17179190565b60ff5f6114ea565b5f5160206119565f395f51905f52919250819006810306145f806114d6565b50838314611493565b50505090505f905f90565b505f5160206119565f395f51905f5281101561136b565b505f5160206119565f395f51905f52821015611364565b505f5160206119565f395f51905f5285101561135d565b80156115d9578060011c915f5160206119565f395f51905f528310156107ac576001806115b85f5160206119565f395f51905f5260038188818180090908611776565b9316146115c157565b905f5160206119565f395f51905f5280910681030690565b505f905f90565b801580611711575b611705578060021c92825f5160206119565f395f51905f5285108015906116ee575b6107ac5784815f5160206119565f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816116b89d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611799565b809291600180829616146116ca575050565b5f5160206119565f395f51905f528093945080929550809106810306930681030690565b505f5160206119565f395f51905f5281101561160a565b50505f905f905f905f90565b5081156115e8565b5f8091602081519101305afa3d1561176e573d9067ffffffffffffffff82116110695760405191611754601f8201601f191660200184611047565b82523d5f602084013e5b156117665750565b602081519101fd5b60609061175e565b90611780826118ed565b915f5160206119565f395f51905f52838009036107ac57565b915f5160206119565f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816117f1939694966117e382808a8009818a800908611776565b906118e1575b860809611776565b925f5160206119565f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206119565f395f51905f5260a083015260208260c08160055afa915191156107ac575f5160206119565f395f51905f528260019209036107ac575f5160206119565f395f51905f52908209925f5160206119565f395f51905f5280808087800906810306818780090814908115916118c2575b506107ac57565b90505f5160206119565f395f51905f528084860960020914155f6118bb565b818091068103066117e9565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206119565f395f51905f5260a083015260208260c08160055afa915191156107ac5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a264697066735822122082e6710d39f2af5080cdd77b864057fe762732d19230520abba3f4f32f1ee0ef64736f6c634300081c0033",
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
