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
	Bin: "0x60808060405234601557611b0c908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c80631e11b6cc14610930578063233ace11146108f657806344f6369214610878578063b8e72af6146108195763f7f338eb14610050575f80fd5b34610815576102a036600319011261081557366101041161081557366102a4116108155760405160408101907f09f303f08b70f5b2dd57c051020fc277280324d7d6d766b6009cb35039188068815260208101917f0f6bd4f9794af59bc6d82964d1201d338b6040c7173f8eef782fb5d56fdde02a83525f8152606082015f81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f15b03463eecbe6c9d19e3e6277a32c5f63e451aa178ef9ae16a9e14c52d30015608087019580875284848460608160075afa911016838860808160065afa165f83525f86526001610124359182895286868660608160075afa9310161616838860808160065afa167f1c9f66855b9c6cf02a4186f33f478dc660d7ee9613394f5f5c0e9bf2eaf1178183527f22b7f63e8d7524c166d958c2c88e732b239faf8274392ae1cbb051cd2aa8c09b8652610144359081885285858560608160075afa92101616838860808160065afa167f2441b4ab66f79ab9604a972ebd29e9c63946d40b6ce80db7cefd50152a1f285f83527f0b4c30e914a1e89e99e3b78fa81990ffba181e3fb9ab1511f5614999a4203d728652610164359081885285858560608160075afa92101616838860808160065afa167f115c8f53a84ab942c511799b9a6aecbc9b477397805f5ec4b6f5e6e8f6cb3ea683527f02756fed000629a4c9987fb9f4dc2428955e8d9b94305660c99111671b8705748652610184359081885285858560608160075afa92101616838860808160065afa167f0e7217b67fc95126bc5ca05ce6619a5d3509f37261a6ef41f45b9078826fe5a183527f0e805cbe24707ee674068f0eb4401cfe5d015deae46e648fb501c79a6294c78d86526101a4359081885285858560608160075afa92101616838860808160065afa167f1e4236e27fff0dee996ca2bf453e3c585ffc60f3f72d88a104c9e88b45457fa483527e268bca6ef630b7184ff79ca3002c0c62164a3d4b78ac234cc23fb4a65ab6dd86526101c4359081885285858560608160075afa92101616838860808160065afa167f14f657c0d7ffae0775b08a7f3bd2031c4cda80a12ea25fc7107c27b596d6fa8183527f03c7c69117a97f87a830c6fb5a14fd5f148ee98a2799b9c68fae9f487301c99886526101e4359081885285858560608160075afa92101616838860808160065afa167f1d428e51316a19d0021cc47219c82dff841343fb7abde094a40921013ed9707383527f21ff6fc90258712cb529e45fdcd1e0c625cf0e56be607426d40c5bce562bd28b8652610204359081885285858560608160075afa92101616838860808160065afa167f1b86d879768757b388fb3d89c8fbf067facfa303496f83b704b3a50ef837000c83527f2c16f64fe1047d727d6d5615701c893d6f9873de1f9ee3e5d3177df2399041468652610224359081885285858560608160075afa92101616838860808160065afa167f075823e5219cb1cec740b4b54b7e8698f2aacccd16075da5b7095e18fb9e5ad083527f0ac9ebfd0ae4f5be8ab622ff13e9dd789ba5df59a2125830ef129c5da96f2dd98652610244359081885285858560608160075afa92101616838860808160065afa167f02d41f7b36711505602833f7955443e2db70db44662fdceb61b4ff39d2a38fe283527f180c764f071761459b4a06bdd13ab7420908a3a1ab45c60af19ad5b2e1858dd28652610264359081885285858560608160075afa92101616838860808160065afa16947f0370ec4aa86c8b9f93cec22174e2df01a382085e147f197df1e0545b5e60241c8352526102843580955260608160075afa9210161660408260808160065afa169051915190156108065760405191610100600484377f0e7fe6144d14b43bc8d9ea669f4de057044ba03ea4b188da11935d303ddf5e086101008401527f193dec86a2e6d307a28789cefae262b46098d4eaf35c950bdae39eea77d2630f6101208401527f1dac3b63858c423c49c7d687e6403789a0e05903c9cebaa75a5656ed8d8e2f236101408401527f01edec6340b037050143b81ae55c6af4343b23175e3036f8e6314222832e68286101608401527f03171eb6c18658f2cf1a531b3eeb13d5a0b780667223bb0d90a78dda30bd3d4b6101808401527f2e86aa3686de6749df5993e02c2d6e09409119046bce29b5946d287f810136c16101a08401527f1c60c78c7cbb833c8213c12c129b0ca36320fb859ce37d579a61a4c2fab5d9116101c08401527f2b978724b2e8559c7116f1fdc5065af6950396dae8b6d980769378723f2afaae6101e08401527f29b84b4438e5db567745a98658bed74c548e288fb305de714a360457bcb66a1a6102008401527f2d12dba5badf2198be044e8437dec2901c838236386db897360933ee695cf9af6102208401526102408301526102608201527f1d2660618f474c7308718d50b9ccb4a4992cb7215cdd83fb2f9ad1d0b92e49ae6102808201527f1800750c731be72f941bb91de5b77178a5bdeda86df24e16e1b2424dc52eeb2f6102a08201527f2d417f101de8bede8f7d7d212e5b0b358c552c3e7d5867412ed42ab17745e1576102c08201527f2b180deb60c40d3bcdf3eac189b3a6cf8566ab7b424c2cc991cdd0c2b508ab216102e08201526020816103008160085afa905116156107f757005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b5f80fd5b346108155760403660031901126108155760043567ffffffffffffffff81116108155761084a90369060040161117a565b6024359167ffffffffffffffff83116108155761086e61087693369060040161117a565b929091611259565b005b34610815576101003660031901126108155736610104116108155760806040516108a282826111a8565b813682376108b46024356004356115a5565b81526108ca60843560a435604435606435611646565b602083015260408201526108e260e43560c4356115a5565b60608201526108f46040518092611153565bf35b34610815575f3660031901126108155760206040517f3871aae077eee32665f8ee95220a1ef4ff279fc8c2387bcffedd1869b66466e88152f35b346108155761022036600319011261081557366084116108155736610224116108155761030060405161096382826111a8565b81368237610972600435611401565b61098360249392933560443561146c565b91939290610992606435611401565b9390926040519660408801967f09f303f08b70f5b2dd57c051020fc277280324d7d6d766b6009cb3503918806889528860208101987f0f6bd4f9794af59bc6d82964d1201d338b6040c7173f8eef782fb5d56fdde02a8a525f81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401925f84527f15b03463eecbe6c9d19e3e6277a32c5f63e451aa178ef9ae16a9e14c52d300156084359583608082019780895286828660608160075afa911016818360808160065afa165f85525f8852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f1c9f66855b9c6cf02a4186f33f478dc660d7ee9613394f5f5c0e9bf2eaf1178185527f22b7f63e8d7524c166d958c2c88e732b239faf8274392ae1cbb051cd2aa8c09b885260c43590818a5287838760608160075afa92101616818360808160065afa167f2441b4ab66f79ab9604a972ebd29e9c63946d40b6ce80db7cefd50152a1f285f85527f0b4c30e914a1e89e99e3b78fa81990ffba181e3fb9ab1511f5614999a4203d72885260e43590818a5287838760608160075afa92101616818360808160065afa167f115c8f53a84ab942c511799b9a6aecbc9b477397805f5ec4b6f5e6e8f6cb3ea685527f02756fed000629a4c9987fb9f4dc2428955e8d9b94305660c99111671b87057488526101043590818a5287838760608160075afa92101616818360808160065afa167f0e7217b67fc95126bc5ca05ce6619a5d3509f37261a6ef41f45b9078826fe5a185527f0e805cbe24707ee674068f0eb4401cfe5d015deae46e648fb501c79a6294c78d88526101243590818a5287838760608160075afa92101616818360808160065afa167f1e4236e27fff0dee996ca2bf453e3c585ffc60f3f72d88a104c9e88b45457fa485527e268bca6ef630b7184ff79ca3002c0c62164a3d4b78ac234cc23fb4a65ab6dd88526101443590818a5287838760608160075afa92101616818360808160065afa167f14f657c0d7ffae0775b08a7f3bd2031c4cda80a12ea25fc7107c27b596d6fa8185527f03c7c69117a97f87a830c6fb5a14fd5f148ee98a2799b9c68fae9f487301c99888526101643590818a5287838760608160075afa92101616818360808160065afa167f1d428e51316a19d0021cc47219c82dff841343fb7abde094a40921013ed9707385527f21ff6fc90258712cb529e45fdcd1e0c625cf0e56be607426d40c5bce562bd28b88526101843590818a5287838760608160075afa92101616818360808160065afa167f1b86d879768757b388fb3d89c8fbf067facfa303496f83b704b3a50ef837000c85527f2c16f64fe1047d727d6d5615701c893d6f9873de1f9ee3e5d3177df23990414688526101a43590818a5287838760608160075afa92101616818360808160065afa167f075823e5219cb1cec740b4b54b7e8698f2aacccd16075da5b7095e18fb9e5ad085527f0ac9ebfd0ae4f5be8ab622ff13e9dd789ba5df59a2125830ef129c5da96f2dd988526101c43590818a5287838760608160075afa92101616818360808160065afa167f02d41f7b36711505602833f7955443e2db70db44662fdceb61b4ff39d2a38fe285527f180c764f071761459b4a06bdd13ab7420908a3a1ab45c60af19ad5b2e1858dd288526101e43590818a5287838760608160075afa921016169160808160065afa16947f0370ec4aa86c8b9f93cec22174e2df01a382085e147f197df1e0545b5e60241c8352526102043580955260608160075afa9210161660408a60808160065afa169851975198156108065760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f0e7fe6144d14b43bc8d9ea669f4de057044ba03ea4b188da11935d303ddf5e086101008401527f193dec86a2e6d307a28789cefae262b46098d4eaf35c950bdae39eea77d2630f6101208401527f1dac3b63858c423c49c7d687e6403789a0e05903c9cebaa75a5656ed8d8e2f236101408401527f01edec6340b037050143b81ae55c6af4343b23175e3036f8e6314222832e68286101608401527f03171eb6c18658f2cf1a531b3eeb13d5a0b780667223bb0d90a78dda30bd3d4b6101808401527f2e86aa3686de6749df5993e02c2d6e09409119046bce29b5946d287f810136c16101a08401527f1c60c78c7cbb833c8213c12c129b0ca36320fb859ce37d579a61a4c2fab5d9116101c08401527f2b978724b2e8559c7116f1fdc5065af6950396dae8b6d980769378723f2afaae6101e08401527f29b84b4438e5db567745a98658bed74c548e288fb305de714a360457bcb66a1a6102008401527f2d12dba5badf2198be044e8437dec2901c838236386db897360933ee695cf9af6102208401526102408301526102608201527f1d2660618f474c7308718d50b9ccb4a4992cb7215cdd83fb2f9ad1d0b92e49ae6102808201527f1800750c731be72f941bb91de5b77178a5bdeda86df24e16e1b2424dc52eeb2f6102a08201527f2d417f101de8bede8f7d7d212e5b0b358c552c3e7d5867412ed42ab17745e1576102c08201527f2b180deb60c40d3bcdf3eac189b3a6cf8566ab7b424c2cc991cdd0c2b508ab216102e082015260405192839161112e84846111a8565b8336843760085afa15908115611146575b506107f757005b600191505114158161113f565b905f905b6004821061116457505050565b6020806001928551815201930191019091611157565b9181601f840112156108155782359167ffffffffffffffff8311610815576020838186019501011161081557565b90601f8019910116810190811067ffffffffffffffff8211176111ca57604052565b634e487b7160e01b5f52604160045260245ffd5b906101a0828203126108155780601f8301121561081557604051916112056101a0846111a8565b82906101a0810192831161081557905b8282106112225750505090565b8135815260209182019101611215565b905f905b600d821061124357505050565b6020806001928551815201930191019091611236565b93929190610100811461133b576080811461127d5763236bd13760e01b5f5260045ffd5b6101a0830361132c578401916080858403126108155782601f8601121561081557604051926112ad6080856111a8565b83956080810191821161081557955b81871061131c57505061131a939450816112db916113049301906111de565b6040516307846db360e21b6020820152926112fa906024850190611153565b60a4830190611232565b6102248152611315610244826111a8565b61187a565b565b86358152602096870196016112bc565b630c0b7e3560e11b5f5260045ffd5b6101a0830361132c57840191610100858403126108155782601f86011215610815576040519261136d610100856111a8565b8395610100810191821161081557955b8187106113f1575050611395929394508101906111de565b60405163f7f338eb60e01b6020820152915f602484015b600882106113db57505050906113ca61131a92610124830190611232565b6102a481526113156102c4826111a8565b60208060019285518152019301910190916113ac565b863581526020968701960161137d565b8015611465578060011c915f516020611ab75f395f51905f528310156107f7576001806114445f516020611ab75f395f51905f52600381888181800909086118d7565b93161461144d57565b905f516020611ab75f395f51905f5280910681030690565b505f905f90565b80158061159d575b611591578060021c92825f516020611ab75f395f51905f52851080159061157a575b6107f75784815f516020611ab75f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816115449d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5086118fa565b80929160018082961614611556575050565b5f516020611ab75f395f51905f528093945080929550809106810306930681030690565b505f516020611ab75f395f51905f52811015611496565b50505f905f905f905f90565b508115611474565b905f516020611ab75f395f51905f52821080159061162f575b6107f757811580611627575b611621576115ee5f516020611ab75f395f51905f52600381858181800909086118d7565b8181036115fd57505060011b90565b5f516020611ab75f395f51905f52809106810306145f146107f757600190811b1790565b50505f90565b5080156115ca565b505f516020611ab75f395f51905f528110156115be565b919093925f516020611ab75f395f51905f528310801590611863575b801561184c575b8015611835575b6107f757808286851717171561182a5790829161178d5f516020611ab75f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611ab75f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161176781808b800981878009086118d7565b8408095f516020611ab75f395f51905f5261178182611a4e565b800914159586916118fa565b929080821480611821575b156117bf5750505050905f146117b75760ff60025b169060021b179190565b60ff5f6117ad565b5f516020611ab75f395f51905f52809106810306149182611802575b5050156107f757600191156117fa5760ff60025b169060021b17179190565b60ff5f6117ef565b5f516020611ab75f395f51905f52919250819006810306145f806117db565b50838314611798565b50505090505f905f90565b505f516020611ab75f395f51905f52811015611670565b505f516020611ab75f395f51905f52821015611669565b505f516020611ab75f395f51905f52851015611662565b5f8091602081519101305afa3d156118cf573d9067ffffffffffffffff82116111ca57604051916118b5601f8201601f1916602001846111a8565b82523d5f602084013e5b156118c75750565b602081519101fd5b6060906118bf565b906118e182611a4e565b915f516020611ab75f395f51905f52838009036107f757565b915f516020611ab75f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816119529396949661194482808a8009818a8009086118d7565b90611a42575b8608096118d7565b925f516020611ab75f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611ab75f395f51905f5260a083015260208260c08160055afa915191156107f7575f516020611ab75f395f51905f528260019209036107f7575f516020611ab75f395f51905f52908209925f516020611ab75f395f51905f528080808780090681030681878009081490811591611a23575b506107f757565b90505f516020611ab75f395f51905f528084860960020914155f611a1c565b8180910681030661194a565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611ab75f395f51905f5260a083015260208260c08160055afa915191156107f75756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220a2487ca0114bc2c724b6492b0c76f8b05c046d2cddd778841b5cdbb441e40c2a64736f6c634300081c0033",
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
