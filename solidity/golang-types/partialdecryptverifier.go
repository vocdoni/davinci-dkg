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
	Bin: "0x60808060405234601557611b0c908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c80631e11b6cc14610930578063233ace11146108f657806344f6369214610878578063b8e72af6146108195763f7f338eb14610050575f80fd5b34610815576102a036600319011261081557366101041161081557366102a4116108155760405160408101907f046dd6e35a82da0ccdfde74df51fc4af904af86332d44b37bc7f38d6f8e71c19815260208101917f2d0bfde2eeead87ac3da47d3b8a01dfb780ac343c3a060313c5b36911eec5d1383525f8152606082015f81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f1d5de66b624d7cce5d7e9d82cfd8b26773fe4efa39959f9321acc99fb2b3c587608087019580875284848460608160075afa911016838860808160065afa165f83525f86526001610124359182895286868660608160075afa9310161616838860808160065afa167f0c9931e4178d925e9ff8af701bb5d57235fd273f1534402f9d57095efd26d04283527f01648ed928c630fa38200aa6c560b10a2ea701a18f022bc11c2886de779db5448652610144359081885285858560608160075afa92101616838860808160065afa167f06c53e6d4e4c4834bc8e5b5d22a8604ae61c62caeb6d45dc951aeb0ce617883683527f03f65f0c3e9e28cb452a3b8d222210199ed6b93259632c9185883e8a24f3a2178652610164359081885285858560608160075afa92101616838860808160065afa167f11724b93a605c8e4d435e404b0a24ae93f4f054a05c4c0c2c7f88fd4aac8c69683527f0a1a7fb89cd56d75b248ebcda043f996a988cb82df9640da90bc25792e39fe478652610184359081885285858560608160075afa92101616838860808160065afa167f2eb4d8fec9b9e287150cfab8b46b9e68944a06aed26a80ae8b79665646706bd883527f1a890aef81673ff8de73908da777356a1e6aa3d4f3a083fc7404192905c441aa86526101a4359081885285858560608160075afa92101616838860808160065afa167f25e7c8125b275a30461cdfd29dd6e340cab300d072bef820b41b983fbe78dc1583527f0579edb4b63a6dabf4e61eb8ff73b8e8f445b4261a24710d649bb7effdbd51a486526101c4359081885285858560608160075afa92101616838860808160065afa167f2ed44ccff0af3df603b883bc80cf3c43587e8958e0215b64944f06840046209f83527f2bc2f96de6accb6e4e0781761ec5d5e527fd7307ff42f49abff84c88bfc2ca7386526101e4359081885285858560608160075afa92101616838860808160065afa167f229a94e0ce01e05c36a5fe021208ea415b3f6cebd953fa715f636a51167357b483527f0179e196c3e567d5c7b1cbd162902877ccdb170c0b8d1119316249799a69243a8652610204359081885285858560608160075afa92101616838860808160065afa167f03df8b96f5c3d16e112b50c8a3789403a9e09742ba4c18d7242d648a8325438f83527ec20c3c3279445c2214da4d449b3ccf73eef7bcaf5986e66da55e454391c5598652610224359081885285858560608160075afa92101616838860808160065afa167f1ea8798a4eda9cebcb4288de4294ec04237028ff5a84ebce310769d3a04ee4ed83527f184754d55a88f5ec8d868d3545aff75e33f59ca226cafcad2f643bf144174eb08652610244359081885285858560608160075afa92101616838860808160065afa167f12f8e0e6a3af0d86ecf63866d7bd7fc2adb47169da0041e1e3d8acf158bb8cb883527f1c0ec98773e768d5074801002847839c36ebc51800d59c8765d0b52868e977898652610264359081885285858560608160075afa92101616838860808160065afa16947f2646e84aa751d857cea0b05cd2fea61f54b00075668c5e6e5b24579a579f029b8352526102843580955260608160075afa9210161660408260808160065afa169051915190156108065760405191610100600484377f1d10e265ac874ca8b2a3094ae5a17195ac3bfbce8117997a12a82f99e590edf06101008401527f13f635f9255671180a8dd7c0642e66dc96656fda40a7670ebcc617945122fb6d6101208401527f1e415ffb1f50123e74754d9473540fc0f543a1e6cee8026b7f861ab6451e5e506101408401527f2a05dec2e1efb9c734a99bbffd3a2cf33b18c7d84c6f625211d0f0edf39b6b646101608401527f18b4b861409eb2611b9162e49ded00061f3fadfd9be3d8e5c4580454cd8cea136101808401527f02a4b8d258c399339b7741379c1e4ba0ef1a2e7fb9617b4566c5ded59dca24af6101a08401527f265199206f0657b24a604b30ce0450303147180f4037c20c71f768c50f02d32d6101c08401527f1847f7b70933246f9dd5ee2414c83b36777fcb1cb8cee2126c05fad1b452ff7c6101e08401527f2151e4208a5d41999decefb5785e5ed2d468751424110bb2b0ac4e1f32bc0f766102008401527f09a43d9c30e17b0d6eaed25e6af8bcc9d3215559d98e4b549300b238620db0b66102208401526102408301526102608201527f1a5c6016485a6946b197abc73f9e1d9140848ef93e9d3c0f08b18bd8bcac25a06102808201527f303b274e15a17c343f156a7188f5419be27d04f77da45421e236aa1b4a3116686102a08201527f11f14d55690f8b72630acd6648a85efa607767fcab8f216d1d957459b950d8726102c08201527f1fd72d7bd0274767983a2e671bb8e0faf99d46b141b586fbb0e4baad701259086102e08201526020816103008160085afa905116156107f757005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b5f80fd5b346108155760403660031901126108155760043567ffffffffffffffff81116108155761084a90369060040161117a565b6024359167ffffffffffffffff83116108155761086e61087693369060040161117a565b929091611259565b005b34610815576101003660031901126108155736610104116108155760806040516108a282826111a8565b813682376108b46024356004356115a5565b81526108ca60843560a435604435606435611646565b602083015260408201526108e260e43560c4356115a5565b60608201526108f46040518092611153565bf35b34610815575f3660031901126108155760206040517f3871aae077eee32665f8ee95220a1ef4ff279fc8c2387bcffedd1869b66466e88152f35b346108155761022036600319011261081557366084116108155736610224116108155761030060405161096382826111a8565b81368237610972600435611401565b61098360249392933560443561146c565b91939290610992606435611401565b9390926040519660408801967f046dd6e35a82da0ccdfde74df51fc4af904af86332d44b37bc7f38d6f8e71c1989528860208101987f2d0bfde2eeead87ac3da47d3b8a01dfb780ac343c3a060313c5b36911eec5d138a525f81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401925f84527f1d5de66b624d7cce5d7e9d82cfd8b26773fe4efa39959f9321acc99fb2b3c5876084359583608082019780895286828660608160075afa911016818360808160065afa165f85525f8852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f0c9931e4178d925e9ff8af701bb5d57235fd273f1534402f9d57095efd26d04285527f01648ed928c630fa38200aa6c560b10a2ea701a18f022bc11c2886de779db544885260c43590818a5287838760608160075afa92101616818360808160065afa167f06c53e6d4e4c4834bc8e5b5d22a8604ae61c62caeb6d45dc951aeb0ce617883685527f03f65f0c3e9e28cb452a3b8d222210199ed6b93259632c9185883e8a24f3a217885260e43590818a5287838760608160075afa92101616818360808160065afa167f11724b93a605c8e4d435e404b0a24ae93f4f054a05c4c0c2c7f88fd4aac8c69685527f0a1a7fb89cd56d75b248ebcda043f996a988cb82df9640da90bc25792e39fe4788526101043590818a5287838760608160075afa92101616818360808160065afa167f2eb4d8fec9b9e287150cfab8b46b9e68944a06aed26a80ae8b79665646706bd885527f1a890aef81673ff8de73908da777356a1e6aa3d4f3a083fc7404192905c441aa88526101243590818a5287838760608160075afa92101616818360808160065afa167f25e7c8125b275a30461cdfd29dd6e340cab300d072bef820b41b983fbe78dc1585527f0579edb4b63a6dabf4e61eb8ff73b8e8f445b4261a24710d649bb7effdbd51a488526101443590818a5287838760608160075afa92101616818360808160065afa167f2ed44ccff0af3df603b883bc80cf3c43587e8958e0215b64944f06840046209f85527f2bc2f96de6accb6e4e0781761ec5d5e527fd7307ff42f49abff84c88bfc2ca7388526101643590818a5287838760608160075afa92101616818360808160065afa167f229a94e0ce01e05c36a5fe021208ea415b3f6cebd953fa715f636a51167357b485527f0179e196c3e567d5c7b1cbd162902877ccdb170c0b8d1119316249799a69243a88526101843590818a5287838760608160075afa92101616818360808160065afa167f03df8b96f5c3d16e112b50c8a3789403a9e09742ba4c18d7242d648a8325438f85527ec20c3c3279445c2214da4d449b3ccf73eef7bcaf5986e66da55e454391c55988526101a43590818a5287838760608160075afa92101616818360808160065afa167f1ea8798a4eda9cebcb4288de4294ec04237028ff5a84ebce310769d3a04ee4ed85527f184754d55a88f5ec8d868d3545aff75e33f59ca226cafcad2f643bf144174eb088526101c43590818a5287838760608160075afa92101616818360808160065afa167f12f8e0e6a3af0d86ecf63866d7bd7fc2adb47169da0041e1e3d8acf158bb8cb885527f1c0ec98773e768d5074801002847839c36ebc51800d59c8765d0b52868e9778988526101e43590818a5287838760608160075afa921016169160808160065afa16947f2646e84aa751d857cea0b05cd2fea61f54b00075668c5e6e5b24579a579f029b8352526102043580955260608160075afa9210161660408a60808160065afa169851975198156108065760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f1d10e265ac874ca8b2a3094ae5a17195ac3bfbce8117997a12a82f99e590edf06101008401527f13f635f9255671180a8dd7c0642e66dc96656fda40a7670ebcc617945122fb6d6101208401527f1e415ffb1f50123e74754d9473540fc0f543a1e6cee8026b7f861ab6451e5e506101408401527f2a05dec2e1efb9c734a99bbffd3a2cf33b18c7d84c6f625211d0f0edf39b6b646101608401527f18b4b861409eb2611b9162e49ded00061f3fadfd9be3d8e5c4580454cd8cea136101808401527f02a4b8d258c399339b7741379c1e4ba0ef1a2e7fb9617b4566c5ded59dca24af6101a08401527f265199206f0657b24a604b30ce0450303147180f4037c20c71f768c50f02d32d6101c08401527f1847f7b70933246f9dd5ee2414c83b36777fcb1cb8cee2126c05fad1b452ff7c6101e08401527f2151e4208a5d41999decefb5785e5ed2d468751424110bb2b0ac4e1f32bc0f766102008401527f09a43d9c30e17b0d6eaed25e6af8bcc9d3215559d98e4b549300b238620db0b66102208401526102408301526102608201527f1a5c6016485a6946b197abc73f9e1d9140848ef93e9d3c0f08b18bd8bcac25a06102808201527f303b274e15a17c343f156a7188f5419be27d04f77da45421e236aa1b4a3116686102a08201527f11f14d55690f8b72630acd6648a85efa607767fcab8f216d1d957459b950d8726102c08201527f1fd72d7bd0274767983a2e671bb8e0faf99d46b141b586fbb0e4baad701259086102e082015260405192839161112e84846111a8565b8336843760085afa15908115611146575b506107f757005b600191505114158161113f565b905f905b6004821061116457505050565b6020806001928551815201930191019091611157565b9181601f840112156108155782359167ffffffffffffffff8311610815576020838186019501011161081557565b90601f8019910116810190811067ffffffffffffffff8211176111ca57604052565b634e487b7160e01b5f52604160045260245ffd5b906101a0828203126108155780601f8301121561081557604051916112056101a0846111a8565b82906101a0810192831161081557905b8282106112225750505090565b8135815260209182019101611215565b905f905b600d821061124357505050565b6020806001928551815201930191019091611236565b93929190610100811461133b576080811461127d5763236bd13760e01b5f5260045ffd5b6101a0830361132c578401916080858403126108155782601f8601121561081557604051926112ad6080856111a8565b83956080810191821161081557955b81871061131c57505061131a939450816112db916113049301906111de565b6040516307846db360e21b6020820152926112fa906024850190611153565b60a4830190611232565b6102248152611315610244826111a8565b61187a565b565b86358152602096870196016112bc565b630c0b7e3560e11b5f5260045ffd5b6101a0830361132c57840191610100858403126108155782601f86011215610815576040519261136d610100856111a8565b8395610100810191821161081557955b8187106113f1575050611395929394508101906111de565b60405163f7f338eb60e01b6020820152915f602484015b600882106113db57505050906113ca61131a92610124830190611232565b6102a481526113156102c4826111a8565b60208060019285518152019301910190916113ac565b863581526020968701960161137d565b8015611465578060011c915f516020611ab75f395f51905f528310156107f7576001806114445f516020611ab75f395f51905f52600381888181800909086118d7565b93161461144d57565b905f516020611ab75f395f51905f5280910681030690565b505f905f90565b80158061159d575b611591578060021c92825f516020611ab75f395f51905f52851080159061157a575b6107f75784815f516020611ab75f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816115449d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5086118fa565b80929160018082961614611556575050565b5f516020611ab75f395f51905f528093945080929550809106810306930681030690565b505f516020611ab75f395f51905f52811015611496565b50505f905f905f905f90565b508115611474565b905f516020611ab75f395f51905f52821080159061162f575b6107f757811580611627575b611621576115ee5f516020611ab75f395f51905f52600381858181800909086118d7565b8181036115fd57505060011b90565b5f516020611ab75f395f51905f52809106810306145f146107f757600190811b1790565b50505f90565b5080156115ca565b505f516020611ab75f395f51905f528110156115be565b919093925f516020611ab75f395f51905f528310801590611863575b801561184c575b8015611835575b6107f757808286851717171561182a5790829161178d5f516020611ab75f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611ab75f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161176781808b800981878009086118d7565b8408095f516020611ab75f395f51905f5261178182611a4e565b800914159586916118fa565b929080821480611821575b156117bf5750505050905f146117b75760ff60025b169060021b179190565b60ff5f6117ad565b5f516020611ab75f395f51905f52809106810306149182611802575b5050156107f757600191156117fa5760ff60025b169060021b17179190565b60ff5f6117ef565b5f516020611ab75f395f51905f52919250819006810306145f806117db565b50838314611798565b50505090505f905f90565b505f516020611ab75f395f51905f52811015611670565b505f516020611ab75f395f51905f52821015611669565b505f516020611ab75f395f51905f52851015611662565b5f8091602081519101305afa3d156118cf573d9067ffffffffffffffff82116111ca57604051916118b5601f8201601f1916602001846111a8565b82523d5f602084013e5b156118c75750565b602081519101fd5b6060906118bf565b906118e182611a4e565b915f516020611ab75f395f51905f52838009036107f757565b915f516020611ab75f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816119529396949661194482808a8009818a8009086118d7565b90611a42575b8608096118d7565b925f516020611ab75f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611ab75f395f51905f5260a083015260208260c08160055afa915191156107f7575f516020611ab75f395f51905f528260019209036107f7575f516020611ab75f395f51905f52908209925f516020611ab75f395f51905f528080808780090681030681878009081490811591611a23575b506107f757565b90505f516020611ab75f395f51905f528084860960020914155f611a1c565b8180910681030661194a565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611ab75f395f51905f5260a083015260208260c08160055afa915191156107f75756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a26469706673582212208d23ffe251e6413008e06c2438f06183ec2da526b7e2bce3077829bca4b570c064736f6c634300081c0033",
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
