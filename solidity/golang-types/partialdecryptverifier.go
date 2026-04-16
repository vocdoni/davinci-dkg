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

// PartialDecryptVerifierMetaData contains all meta data concerning the PartialDecryptVerifier contract.
var PartialDecryptVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[13]\",\"internalType\":\"uint256[13]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"input\",\"type\":\"uint256[13]\",\"internalType\":\"uint256[13]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidInputEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x60808060405234601557611c0c908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c80631e11b6cc146109b0578063233ace111461097657806344f63692146108f8578063b8e72af6146108995763f7f338eb14610050575f80fd5b34610895576102a036600319011261089557366101041161089557366102a4116108955760405160408101907f020d271c5ac881f20db4b43b127a3550338c1d76026d7a67aa053ec3751a1df4815260208101917f3047670fca9a27c6d8a1a4cc5e07623b37b896c6fbd30d5ddbb793e1bf86952083527f2240c5df788da0a3b99c44cd62593f9c81e9f870c7436a2d2c58a101f0638c648152606082017f01450713e986b79eccdeb941c3508f14476040b27ed0bf34a95f7f47c1b1f1ac81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f071355fc88b0ca164bcacc6fbf11349c86fb8c3c8ecad17dce72e050fc8e37ad608087019580875284848460608160075afa911016838860808160065afa167f15c690c30aac0d9bf17964d90c08c22c1ac88692e887f36b3a88bb1595512fc783527f12fbca7eaefbacb70e11493659cebbc8d8d05b27ebf9e1e82fc7772e3a68a5e386526001610124359182895286868660608160075afa9310161616838860808160065afa167f29b279db6f6806379a7c55b118ef92f41bbd98b11e9ee1c79dcf1b6f122b337383527f132b4d3b5f5ca3bcc0ec5bf81e4afdd83bf9e38341f190a290e366ed2fd35e728652610144359081885285858560608160075afa92101616838860808160065afa167e484b98a0dfecb91ab9c867dbd7b3857f8d877289a6463e492773e3ba26f7bd83527f05535834f94ae90686bec01717ceb24556ff705bd9e115e62b6455ccb5df17398652610164359081885285858560608160075afa92101616838860808160065afa167f0fd9f1147db33a8f37543caf732c0733730b9611ad8e9f4162aa42e88f63872783527f0d9dbf4a462cedde163dc9b27863c38158b8bde98b1f5b83a2273392a641d99e8652610184359081885285858560608160075afa92101616838860808160065afa167f07325b357fc456c7b30042e2219ab3b38fae9d99f4cd7ab3b33130750b2dece083527f1fde9f44bd29d0df6f75d5537f878a47b64f3dd68557b688b9f166b3397c5d7586526101a4359081885285858560608160075afa92101616838860808160065afa167f02bb06f4874c5e0d6cfad7f58d7047d220662f27a77e313250da37693218ff8183527f1c07a28de964b34a67142a92467c8c45842b92bc07a1ce2e2fb5527491bd90d386526101c4359081885285858560608160075afa92101616838860808160065afa167f11f14f61df17a3bbf35d24ae91cc880c5aa49ca62d3d279b67d9604e3ba2548083527f148ac317c76458d197cc0167a434262b9dc839c705efda6c34e7a0f037efd73b86526101e4359081885285858560608160075afa92101616838860808160065afa167f15d509d34119e69f7f1f38aa453fff9ab6ed496d24966e763caf91a9819c61d583527f1491f0e1667b3fd62487ef5de71292489631d99dd162792fc3392f78d9c719fd8652610204359081885285858560608160075afa92101616838860808160065afa167f284907989aa68ff9f787f22c7e95b79f86a655ddc95201ad4e2828b609747b0483527f2415b930f05b8cd877099fa468beced30ce03289d27afb521f0e215a0e95bcc18652610224359081885285858560608160075afa92101616838860808160065afa167f197c6b9b9d568b6facba361a87f3a83fa3afae97f74c0f72f1e00372d6fd928483527f1fa323d88acc37176d3c758bbdefac2ce405e235fd414b16dbfe05b7a308fca18652610244359081885285858560608160075afa92101616838860808160065afa167f1d954fb300e7deb5fa5744888fe36f224ccc4fd99f5b3184fb6609ab6c797a2583527f19a6cf04d713c862182c99360070ed5334434a836dbf6f418ba340e6a2a622008652610264359081885285858560608160075afa92101616838860808160065afa16947f07e5d04daf7e99efed449bdf7c78b9e6d146c81293d85d71c43eaa4463c6651c8352526102843580955260608160075afa9210161660408260808160065afa169051915190156108865760405191610100600484377f1c395128f063daff897eaed7bbf9dc1a9cf0cc875fbc619b764cf5634b4221976101008401527f28cb0ef4e32f21b28886b083ddcd523ab7d1ce471644c8895d7f75dd72c4534f6101208401527f085e3338d6cb0ae205bdf37415e379a651cecba40f529702213cb58a1d0f93906101408401527f11fbef20e9f353335b558059094a644146c11fa99a697772eb0bab3b02e364f66101608401527f2818ce49497351b0dae316e307fd4c392fd6c6f79adb6fafb33e9221a82c23486101808401527f11b54150b3389bad7f1193b98f9f2a6356ada25ca49714c1cf0a65a03ee01b7e6101a08401527f152a8b54e3162c4c54b8dc490001de7fcd8c506a191c87251424da198eab4f936101c08401527f067071e1ccf862e7d55bf47f982e188035883d8ea7aea0b387749bb80d6dc17c6101e08401527f11ab83d2def166250af6a1a4d17590f96cd5ef6e38d48fe7bb66771a9d2107f06102008401527f203a2a5687fd2719ab8c469490fc3d153a5664fc430e67c41dc841e6227578d56102208401526102408301526102608201527f1af40b0007fcb504c33fc1dd84f6d070c80c78325787625304e28ec9aeb0b3de6102808201527f1c74a05dde7562b4e4a0f95eb67c54053e740d158248c14431e070d4d546124f6102a08201527f20669647e4ecda3b278c34e592adb47bb493493efe6c332456cf7035203449556102c08201527f2d0501277e0eaa39eb8a565649450f5fa061815bdcd4dd5ce19e63bd7a61f25d6102e08201526020816103008160085afa9051161561087757005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b5f80fd5b346108955760403660031901126108955760043567ffffffffffffffff8111610895576108ca90369060040161127a565b6024359167ffffffffffffffff8311610895576108ee6108f693369060040161127a565b929091611359565b005b346108955761010036600319011261089557366101041161089557608060405161092282826112a8565b813682376109346024356004356116a5565b815261094a60843560a435604435606435611746565b6020830152604082015261096260e43560c4356116a5565b60608201526109746040518092611253565bf35b34610895575f3660031901126108955760206040517f3871aae077eee32665f8ee95220a1ef4ff279fc8c2387bcffedd1869b66466e88152f35b34610895576102203660031901126108955736608411610895573661022411610895576103006040516109e382826112a8565b813682376109f2600435611501565b610a0360249392933560443561156c565b91939290610a12606435611501565b9390926040519660408801967f020d271c5ac881f20db4b43b127a3550338c1d76026d7a67aa053ec3751a1df489528860208101987f3047670fca9a27c6d8a1a4cc5e07623b37b896c6fbd30d5ddbb793e1bf8695208a527f2240c5df788da0a3b99c44cd62593f9c81e9f870c7436a2d2c58a101f0638c6481527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f01450713e986b79eccdeb941c3508f14476040b27ed0bf34a95f7f47c1b1f1ac84527f071355fc88b0ca164bcacc6fbf11349c86fb8c3c8ecad17dce72e050fc8e37ad6084359583608082019780895286828660608160075afa911016818360808160065afa167f15c690c30aac0d9bf17964d90c08c22c1ac88692e887f36b3a88bb1595512fc785527f12fbca7eaefbacb70e11493659cebbc8d8d05b27ebf9e1e82fc7772e3a68a5e38852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f29b279db6f6806379a7c55b118ef92f41bbd98b11e9ee1c79dcf1b6f122b337385527f132b4d3b5f5ca3bcc0ec5bf81e4afdd83bf9e38341f190a290e366ed2fd35e72885260c43590818a5287838760608160075afa92101616818360808160065afa167e484b98a0dfecb91ab9c867dbd7b3857f8d877289a6463e492773e3ba26f7bd85527f05535834f94ae90686bec01717ceb24556ff705bd9e115e62b6455ccb5df1739885260e43590818a5287838760608160075afa92101616818360808160065afa167f0fd9f1147db33a8f37543caf732c0733730b9611ad8e9f4162aa42e88f63872785527f0d9dbf4a462cedde163dc9b27863c38158b8bde98b1f5b83a2273392a641d99e88526101043590818a5287838760608160075afa92101616818360808160065afa167f07325b357fc456c7b30042e2219ab3b38fae9d99f4cd7ab3b33130750b2dece085527f1fde9f44bd29d0df6f75d5537f878a47b64f3dd68557b688b9f166b3397c5d7588526101243590818a5287838760608160075afa92101616818360808160065afa167f02bb06f4874c5e0d6cfad7f58d7047d220662f27a77e313250da37693218ff8185527f1c07a28de964b34a67142a92467c8c45842b92bc07a1ce2e2fb5527491bd90d388526101443590818a5287838760608160075afa92101616818360808160065afa167f11f14f61df17a3bbf35d24ae91cc880c5aa49ca62d3d279b67d9604e3ba2548085527f148ac317c76458d197cc0167a434262b9dc839c705efda6c34e7a0f037efd73b88526101643590818a5287838760608160075afa92101616818360808160065afa167f15d509d34119e69f7f1f38aa453fff9ab6ed496d24966e763caf91a9819c61d585527f1491f0e1667b3fd62487ef5de71292489631d99dd162792fc3392f78d9c719fd88526101843590818a5287838760608160075afa92101616818360808160065afa167f284907989aa68ff9f787f22c7e95b79f86a655ddc95201ad4e2828b609747b0485527f2415b930f05b8cd877099fa468beced30ce03289d27afb521f0e215a0e95bcc188526101a43590818a5287838760608160075afa92101616818360808160065afa167f197c6b9b9d568b6facba361a87f3a83fa3afae97f74c0f72f1e00372d6fd928485527f1fa323d88acc37176d3c758bbdefac2ce405e235fd414b16dbfe05b7a308fca188526101c43590818a5287838760608160075afa92101616818360808160065afa167f1d954fb300e7deb5fa5744888fe36f224ccc4fd99f5b3184fb6609ab6c797a2585527f19a6cf04d713c862182c99360070ed5334434a836dbf6f418ba340e6a2a6220088526101e43590818a5287838760608160075afa921016169160808160065afa16947f07e5d04daf7e99efed449bdf7c78b9e6d146c81293d85d71c43eaa4463c6651c8352526102043580955260608160075afa9210161660408a60808160065afa169851975198156108865760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f1c395128f063daff897eaed7bbf9dc1a9cf0cc875fbc619b764cf5634b4221976101008401527f28cb0ef4e32f21b28886b083ddcd523ab7d1ce471644c8895d7f75dd72c4534f6101208401527f085e3338d6cb0ae205bdf37415e379a651cecba40f529702213cb58a1d0f93906101408401527f11fbef20e9f353335b558059094a644146c11fa99a697772eb0bab3b02e364f66101608401527f2818ce49497351b0dae316e307fd4c392fd6c6f79adb6fafb33e9221a82c23486101808401527f11b54150b3389bad7f1193b98f9f2a6356ada25ca49714c1cf0a65a03ee01b7e6101a08401527f152a8b54e3162c4c54b8dc490001de7fcd8c506a191c87251424da198eab4f936101c08401527f067071e1ccf862e7d55bf47f982e188035883d8ea7aea0b387749bb80d6dc17c6101e08401527f11ab83d2def166250af6a1a4d17590f96cd5ef6e38d48fe7bb66771a9d2107f06102008401527f203a2a5687fd2719ab8c469490fc3d153a5664fc430e67c41dc841e6227578d56102208401526102408301526102608201527f1af40b0007fcb504c33fc1dd84f6d070c80c78325787625304e28ec9aeb0b3de6102808201527f1c74a05dde7562b4e4a0f95eb67c54053e740d158248c14431e070d4d546124f6102a08201527f20669647e4ecda3b278c34e592adb47bb493493efe6c332456cf7035203449556102c08201527f2d0501277e0eaa39eb8a565649450f5fa061815bdcd4dd5ce19e63bd7a61f25d6102e082015260405192839161122e84846112a8565b8336843760085afa15908115611246575b5061087757005b600191505114158161123f565b905f905b6004821061126457505050565b6020806001928551815201930191019091611257565b9181601f840112156108955782359167ffffffffffffffff8311610895576020838186019501011161089557565b90601f8019910116810190811067ffffffffffffffff8211176112ca57604052565b634e487b7160e01b5f52604160045260245ffd5b906101a0828203126108955780601f8301121561089557604051916113056101a0846112a8565b82906101a0810192831161089557905b8282106113225750505090565b8135815260209182019101611315565b905f905b600d821061134357505050565b6020806001928551815201930191019091611336565b93929190610100811461143b576080811461137d5763236bd13760e01b5f5260045ffd5b6101a0830361142c578401916080858403126108955782601f8601121561089557604051926113ad6080856112a8565b83956080810191821161089557955b81871061141c57505061141a939450816113db916114049301906112de565b6040516307846db360e21b6020820152926113fa906024850190611253565b60a4830190611332565b6102248152611415610244826112a8565b61197a565b565b86358152602096870196016113bc565b630c0b7e3560e11b5f5260045ffd5b6101a0830361142c57840191610100858403126108955782601f86011215610895576040519261146d610100856112a8565b8395610100810191821161089557955b8187106114f1575050611495929394508101906112de565b60405163f7f338eb60e01b6020820152915f602484015b600882106114db57505050906114ca61141a92610124830190611332565b6102a481526114156102c4826112a8565b60208060019285518152019301910190916114ac565b863581526020968701960161147d565b8015611565578060011c915f516020611bb75f395f51905f52831015610877576001806115445f516020611bb75f395f51905f52600381888181800909086119d7565b93161461154d57565b905f516020611bb75f395f51905f5280910681030690565b505f905f90565b80158061169d575b611691578060021c92825f516020611bb75f395f51905f52851080159061167a575b6108775784815f516020611bb75f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816116449d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5086119fa565b80929160018082961614611656575050565b5f516020611bb75f395f51905f528093945080929550809106810306930681030690565b505f516020611bb75f395f51905f52811015611596565b50505f905f905f905f90565b508115611574565b905f516020611bb75f395f51905f52821080159061172f575b61087757811580611727575b611721576116ee5f516020611bb75f395f51905f52600381858181800909086119d7565b8181036116fd57505060011b90565b5f516020611bb75f395f51905f52809106810306145f1461087757600190811b1790565b50505f90565b5080156116ca565b505f516020611bb75f395f51905f528110156116be565b919093925f516020611bb75f395f51905f528310801590611963575b801561194c575b8015611935575b61087757808286851717171561192a5790829161188d5f516020611bb75f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611bb75f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161186781808b800981878009086119d7565b8408095f516020611bb75f395f51905f5261188182611b4e565b800914159586916119fa565b929080821480611921575b156118bf5750505050905f146118b75760ff60025b169060021b179190565b60ff5f6118ad565b5f516020611bb75f395f51905f52809106810306149182611902575b50501561087757600191156118fa5760ff60025b169060021b17179190565b60ff5f6118ef565b5f516020611bb75f395f51905f52919250819006810306145f806118db565b50838314611898565b50505090505f905f90565b505f516020611bb75f395f51905f52811015611770565b505f516020611bb75f395f51905f52821015611769565b505f516020611bb75f395f51905f52851015611762565b5f8091602081519101305afa3d156119cf573d9067ffffffffffffffff82116112ca57604051916119b5601f8201601f1916602001846112a8565b82523d5f602084013e5b156119c75750565b602081519101fd5b6060906119bf565b906119e182611b4e565b915f516020611bb75f395f51905f528380090361087757565b915f516020611bb75f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611a5293969496611a4482808a8009818a8009086119d7565b90611b42575b8608096119d7565b925f516020611bb75f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611bb75f395f51905f5260a083015260208260c08160055afa91519115610877575f516020611bb75f395f51905f52826001920903610877575f516020611bb75f395f51905f52908209925f516020611bb75f395f51905f528080808780090681030681878009081490811591611b23575b5061087757565b90505f516020611bb75f395f51905f528084860960020914155f611b1c565b81809106810306611a4a565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611bb75f395f51905f5260a083015260208260c08160055afa915191156108775756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220c2a7d2f0578146723240d093846526db88efdfb0a5d51c9e5de851338b5496ae64736f6c634300081c0033",
}

// PartialDecryptVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use PartialDecryptVerifierMetaData.ABI instead.
var PartialDecryptVerifierABI = PartialDecryptVerifierMetaData.ABI

// PartialDecryptVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PartialDecryptVerifierMetaData.Bin instead.
var PartialDecryptVerifierBin = PartialDecryptVerifierMetaData.Bin

// DeployPartialDecryptVerifier deploys a new Ethereum contract, binding an instance of PartialDecryptVerifier to it.
func DeployPartialDecryptVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PartialDecryptVerifier, error) {
	parsed, err := PartialDecryptVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PartialDecryptVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PartialDecryptVerifier{PartialDecryptVerifierCaller: PartialDecryptVerifierCaller{contract: contract}, PartialDecryptVerifierTransactor: PartialDecryptVerifierTransactor{contract: contract}, PartialDecryptVerifierFilterer: PartialDecryptVerifierFilterer{contract: contract}}, nil
}

// PartialDecryptVerifier is an auto generated Go binding around an Ethereum contract.
type PartialDecryptVerifier struct {
	PartialDecryptVerifierCaller     // Read-only binding to the contract
	PartialDecryptVerifierTransactor // Write-only binding to the contract
	PartialDecryptVerifierFilterer   // Log filterer for contract events
}

// PartialDecryptVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type PartialDecryptVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartialDecryptVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PartialDecryptVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartialDecryptVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PartialDecryptVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartialDecryptVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PartialDecryptVerifierSession struct {
	Contract     *PartialDecryptVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PartialDecryptVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PartialDecryptVerifierCallerSession struct {
	Contract *PartialDecryptVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// PartialDecryptVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PartialDecryptVerifierTransactorSession struct {
	Contract     *PartialDecryptVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// PartialDecryptVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type PartialDecryptVerifierRaw struct {
	Contract *PartialDecryptVerifier // Generic contract binding to access the raw methods on
}

// PartialDecryptVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PartialDecryptVerifierCallerRaw struct {
	Contract *PartialDecryptVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// PartialDecryptVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PartialDecryptVerifierTransactorRaw struct {
	Contract *PartialDecryptVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPartialDecryptVerifier creates a new instance of PartialDecryptVerifier, bound to a specific deployed contract.
func NewPartialDecryptVerifier(address common.Address, backend bind.ContractBackend) (*PartialDecryptVerifier, error) {
	contract, err := bindPartialDecryptVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PartialDecryptVerifier{PartialDecryptVerifierCaller: PartialDecryptVerifierCaller{contract: contract}, PartialDecryptVerifierTransactor: PartialDecryptVerifierTransactor{contract: contract}, PartialDecryptVerifierFilterer: PartialDecryptVerifierFilterer{contract: contract}}, nil
}

// NewPartialDecryptVerifierCaller creates a new read-only instance of PartialDecryptVerifier, bound to a specific deployed contract.
func NewPartialDecryptVerifierCaller(address common.Address, caller bind.ContractCaller) (*PartialDecryptVerifierCaller, error) {
	contract, err := bindPartialDecryptVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PartialDecryptVerifierCaller{contract: contract}, nil
}

// NewPartialDecryptVerifierTransactor creates a new write-only instance of PartialDecryptVerifier, bound to a specific deployed contract.
func NewPartialDecryptVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*PartialDecryptVerifierTransactor, error) {
	contract, err := bindPartialDecryptVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PartialDecryptVerifierTransactor{contract: contract}, nil
}

// NewPartialDecryptVerifierFilterer creates a new log filterer instance of PartialDecryptVerifier, bound to a specific deployed contract.
func NewPartialDecryptVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*PartialDecryptVerifierFilterer, error) {
	contract, err := bindPartialDecryptVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PartialDecryptVerifierFilterer{contract: contract}, nil
}

// bindPartialDecryptVerifier binds a generic wrapper to an already deployed contract.
func bindPartialDecryptVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PartialDecryptVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PartialDecryptVerifier *PartialDecryptVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PartialDecryptVerifier.Contract.PartialDecryptVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PartialDecryptVerifier *PartialDecryptVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartialDecryptVerifier.Contract.PartialDecryptVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PartialDecryptVerifier *PartialDecryptVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PartialDecryptVerifier.Contract.PartialDecryptVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PartialDecryptVerifier *PartialDecryptVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PartialDecryptVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PartialDecryptVerifier *PartialDecryptVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartialDecryptVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PartialDecryptVerifier *PartialDecryptVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PartialDecryptVerifier.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_PartialDecryptVerifier *PartialDecryptVerifierCaller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int) ([4]*big.Int, error) {
	var out []interface{}
	err := _PartialDecryptVerifier.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_PartialDecryptVerifier *PartialDecryptVerifierSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _PartialDecryptVerifier.Contract.CompressProof(&_PartialDecryptVerifier.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_PartialDecryptVerifier *PartialDecryptVerifierCallerSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _PartialDecryptVerifier.Contract.CompressProof(&_PartialDecryptVerifier.CallOpts, proof)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_PartialDecryptVerifier *PartialDecryptVerifierCaller) ProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PartialDecryptVerifier.contract.Call(opts, &out, "provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_PartialDecryptVerifier *PartialDecryptVerifierSession) ProvingKeyHash() ([32]byte, error) {
	return _PartialDecryptVerifier.Contract.ProvingKeyHash(&_PartialDecryptVerifier.CallOpts)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_PartialDecryptVerifier *PartialDecryptVerifierCallerSession) ProvingKeyHash() ([32]byte, error) {
	return _PartialDecryptVerifier.Contract.ProvingKeyHash(&_PartialDecryptVerifier.CallOpts)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x1e11b6cc.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[13] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [13]*big.Int) error {
	var out []interface{}
	err := _PartialDecryptVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x1e11b6cc.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[13] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [13]*big.Int) error {
	return _PartialDecryptVerifier.Contract.VerifyCompressedProof(&_PartialDecryptVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x1e11b6cc.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[13] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [13]*big.Int) error {
	return _PartialDecryptVerifier.Contract.VerifyCompressedProof(&_PartialDecryptVerifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCaller) VerifyProof(opts *bind.CallOpts, proof []byte, input []byte) error {
	var out []interface{}
	err := _PartialDecryptVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierSession) VerifyProof(proof []byte, input []byte) error {
	return _PartialDecryptVerifier.Contract.VerifyProof(&_PartialDecryptVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCallerSession) VerifyProof(proof []byte, input []byte) error {
	return _PartialDecryptVerifier.Contract.VerifyProof(&_PartialDecryptVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xf7f338eb.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[13] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCaller) VerifyProof0(opts *bind.CallOpts, proof [8]*big.Int, input [13]*big.Int) error {
	var out []interface{}
	err := _PartialDecryptVerifier.contract.Call(opts, &out, "verifyProof0", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xf7f338eb.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[13] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierSession) VerifyProof0(proof [8]*big.Int, input [13]*big.Int) error {
	return _PartialDecryptVerifier.Contract.VerifyProof0(&_PartialDecryptVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xf7f338eb.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[13] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCallerSession) VerifyProof0(proof [8]*big.Int, input [13]*big.Int) error {
	return _PartialDecryptVerifier.Contract.VerifyProof0(&_PartialDecryptVerifier.CallOpts, proof, input)
}
