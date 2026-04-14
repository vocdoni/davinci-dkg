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
	Bin: "0x60808060405234601557611b0e908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c80631e11b6cc14610931578063233ace11146108f757806344f6369214610879578063b8e72af61461081a5763f7f338eb14610050575f80fd5b34610816576102a036600319011261081657366101041161081657366102a4116108165760405160408101907f065a94fc1427ca1fd0fd9c9938f07cd56044902f01ff1430d868b10d365f58a4815260208101917f28ee90356b55352644be315c7d4e7625d361fdc73065876cd18d35bfbec3ff9983525f8152606082015f81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f21d39847eafb40008bbad967ceec5b61d500a056a2d7bc98db5b072846e309ee608087019580875284848460608160075afa911016838860808160065afa165f83525f86526001610124359182895286868660608160075afa9310161616838860808160065afa167f1a5429cb79ef98c5660c7d7f55c0a1694741992d1a27cea1f5ca9195453f8afc83527f1ace023d1273d0dada0643b4c37e35eeb1e6222f9abc179a7193df59acee50c58652610144359081885285858560608160075afa92101616838860808160065afa167f0c53790b6729a058649775db6eda01b51c350719da71713b0df80bd059ecf69983527f08cc9850e6053363d95e5d79ff80f24afebaf832593762d459e9728c419606588652610164359081885285858560608160075afa92101616838860808160065afa167f2b33ce3ba77b6e69b3492e34fee6197f8b98546a44d917763594cb3908594ecf83527f206ae430aa9f585bf63f5014fa567105e43689eb18c52cbaa88ede0122b002d18652610184359081885285858560608160075afa92101616838860808160065afa167f1fc575c20d59e3b598e45ce1494e6f47c2e85416a9662098ad6acc432abadfbb83527f2bce8972359cf0f90279cc38067867734bcbd0c6ebd868223f0645a24008743786526101a4359081885285858560608160075afa92101616838860808160065afa167f2b28302c6f715116c64634823930d29d52f0dbb3ad90ff02fba0244a9751105583527f0e3c55ae2b2b130b0082107b19ddd6867f2e97cf935db6eecf159a7cb0db91b686526101c4359081885285858560608160075afa92101616838860808160065afa167f0ef5ca4987cb77d91813f39a7566a9d49d83e89666bc4bf66c4b9199bd0b2cd983527f2b64fb48ff14c073e84fa5ef0c8e14c7c937e501653097aba2034483bd01909086526101e4359081885285858560608160075afa92101616838860808160065afa167f06c58d9a92a5a4d4a822d055ebcd55cda32dc5dcdd59ecd0fbb419d90f277f9e83527f0e07dde44643b2251a328fb0d53231227195db00aa5fecfeacf7847a4adf52ef8652610204359081885285858560608160075afa92101616838860808160065afa167f0dcf69632e9961a4c4d87c8ebac3254f04ff32807f24150f7d6f8aa2023a75ef83527f29072b1fae3bc13146d9a3ca9121b3db4fb5097b8554c0a33a2339eaa14bda488652610224359081885285858560608160075afa92101616838860808160065afa167f2c9a278a6b046364e2499f5e6923bb16930f1f35cbb6655bbcc66f03c4b3f69b83527f225188ebc282192fdcfa2d9450fdc3762fa0ec2f3d0a8a6c1be2256e5c8d9b5f8652610244359081885285858560608160075afa92101616838860808160065afa167f070a788b3bfe8af6497d9eeb3d4cce32394b89df04798b07cc1ff48335e575ba83527f18023326033c95acbd6a9867cceb3061431194e3f50c185c77560e2efa5352c98652610264359081885285858560608160075afa92101616838860808160065afa16947f10f5a607876290ebbf0c407dabbeba887b5553729d3c0a5feb587ccdd15b66008352526102843580955260608160075afa9210161660408260808160065afa169051915190156108075760405191610100600484377f0ae19d4352eb399cc74046dcf7b8383482ccf4c66279d953d2c1efa3720c9b356101008401527f2c66518fb10650050160ee783f9b22549e42dcc994f58f10867bb7510308f1d86101208401527f2da95964eb990363b2af2e711df9f26e29bbb82104fa06979c813d24e67d6ec96101408401527f0c9f69dcaa3e768ae635ce5b5baffa2b3ed0c1a6897c98f0ad7103423110d29a6101608401527f073f3090ca26aad57d1fa4e8bc335ae5e2562c7a75e0c8809e50bc6ae0d8cc886101808401527f06d4be78cd995db24486cf5f0b680eb2c844f6c737e873de428d5b301fd100b96101a08401527f0244721531c55fab7d87d44a18377286d989653f2a499633e1d6a7774a3cac246101c08401527f2740e7a6515e44b675aaf3fdd672a992736a8cae51e8baebe95f5759ad59aa516101e08401527f0b5a922d7caf56de6918815dd7e5782a48f7f3b208f35412d95bc6dda7a451bf6102008401527f0ca4de36daa336a425749a2849c1683a9662335b09d9bdfedd79c3a2e0e13a576102208401526102408301526102608201527f072c36ba67c6b3221e33c495797b900a8f9fc14a609f9ce8c84b69e903977cbc6102808201527f16902914c61a626430be7e3bb9b93e674610a506e297119b20dd8146d120fe1d6102a08201527f203ec77b656a8c2ae4825bc54ca7d727b532c05425071982c5b67e89e28886b56102c08201527f16c3a3a844ebe967bc4939399b78ce6bd519cfbc0697491d1f1c46cfa3d704216102e08201526020816103008160085afa905116156107f857005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b5f80fd5b346108165760403660031901126108165760043567ffffffffffffffff81116108165761084b90369060040161117c565b6024359167ffffffffffffffff83116108165761086f61087793369060040161117c565b92909161125b565b005b34610816576101003660031901126108165736610104116108165760806040516108a382826111aa565b813682376108b56024356004356115a7565b81526108cb60843560a435604435606435611648565b602083015260408201526108e360e43560c4356115a7565b60608201526108f56040518092611155565bf35b34610816575f3660031901126108165760206040517f3871aae077eee32665f8ee95220a1ef4ff279fc8c2387bcffedd1869b66466e88152f35b346108165761022036600319011261081657366084116108165736610224116108165761030060405161096482826111aa565b81368237610973600435611403565b61098460249392933560443561146e565b91939290610993606435611403565b9390926040519660408801967f065a94fc1427ca1fd0fd9c9938f07cd56044902f01ff1430d868b10d365f58a489528860208101987f28ee90356b55352644be315c7d4e7625d361fdc73065876cd18d35bfbec3ff998a525f81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401925f84527f21d39847eafb40008bbad967ceec5b61d500a056a2d7bc98db5b072846e309ee6084359583608082019780895286828660608160075afa911016818360808160065afa165f85525f8852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f1a5429cb79ef98c5660c7d7f55c0a1694741992d1a27cea1f5ca9195453f8afc85527f1ace023d1273d0dada0643b4c37e35eeb1e6222f9abc179a7193df59acee50c5885260c43590818a5287838760608160075afa92101616818360808160065afa167f0c53790b6729a058649775db6eda01b51c350719da71713b0df80bd059ecf69985527f08cc9850e6053363d95e5d79ff80f24afebaf832593762d459e9728c41960658885260e43590818a5287838760608160075afa92101616818360808160065afa167f2b33ce3ba77b6e69b3492e34fee6197f8b98546a44d917763594cb3908594ecf85527f206ae430aa9f585bf63f5014fa567105e43689eb18c52cbaa88ede0122b002d188526101043590818a5287838760608160075afa92101616818360808160065afa167f1fc575c20d59e3b598e45ce1494e6f47c2e85416a9662098ad6acc432abadfbb85527f2bce8972359cf0f90279cc38067867734bcbd0c6ebd868223f0645a24008743788526101243590818a5287838760608160075afa92101616818360808160065afa167f2b28302c6f715116c64634823930d29d52f0dbb3ad90ff02fba0244a9751105585527f0e3c55ae2b2b130b0082107b19ddd6867f2e97cf935db6eecf159a7cb0db91b688526101443590818a5287838760608160075afa92101616818360808160065afa167f0ef5ca4987cb77d91813f39a7566a9d49d83e89666bc4bf66c4b9199bd0b2cd985527f2b64fb48ff14c073e84fa5ef0c8e14c7c937e501653097aba2034483bd01909088526101643590818a5287838760608160075afa92101616818360808160065afa167f06c58d9a92a5a4d4a822d055ebcd55cda32dc5dcdd59ecd0fbb419d90f277f9e85527f0e07dde44643b2251a328fb0d53231227195db00aa5fecfeacf7847a4adf52ef88526101843590818a5287838760608160075afa92101616818360808160065afa167f0dcf69632e9961a4c4d87c8ebac3254f04ff32807f24150f7d6f8aa2023a75ef85527f29072b1fae3bc13146d9a3ca9121b3db4fb5097b8554c0a33a2339eaa14bda4888526101a43590818a5287838760608160075afa92101616818360808160065afa167f2c9a278a6b046364e2499f5e6923bb16930f1f35cbb6655bbcc66f03c4b3f69b85527f225188ebc282192fdcfa2d9450fdc3762fa0ec2f3d0a8a6c1be2256e5c8d9b5f88526101c43590818a5287838760608160075afa92101616818360808160065afa167f070a788b3bfe8af6497d9eeb3d4cce32394b89df04798b07cc1ff48335e575ba85527f18023326033c95acbd6a9867cceb3061431194e3f50c185c77560e2efa5352c988526101e43590818a5287838760608160075afa921016169160808160065afa16947f10f5a607876290ebbf0c407dabbeba887b5553729d3c0a5feb587ccdd15b66008352526102043580955260608160075afa9210161660408a60808160065afa169851975198156108075760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f0ae19d4352eb399cc74046dcf7b8383482ccf4c66279d953d2c1efa3720c9b356101008401527f2c66518fb10650050160ee783f9b22549e42dcc994f58f10867bb7510308f1d86101208401527f2da95964eb990363b2af2e711df9f26e29bbb82104fa06979c813d24e67d6ec96101408401527f0c9f69dcaa3e768ae635ce5b5baffa2b3ed0c1a6897c98f0ad7103423110d29a6101608401527f073f3090ca26aad57d1fa4e8bc335ae5e2562c7a75e0c8809e50bc6ae0d8cc886101808401527f06d4be78cd995db24486cf5f0b680eb2c844f6c737e873de428d5b301fd100b96101a08401527f0244721531c55fab7d87d44a18377286d989653f2a499633e1d6a7774a3cac246101c08401527f2740e7a6515e44b675aaf3fdd672a992736a8cae51e8baebe95f5759ad59aa516101e08401527f0b5a922d7caf56de6918815dd7e5782a48f7f3b208f35412d95bc6dda7a451bf6102008401527f0ca4de36daa336a425749a2849c1683a9662335b09d9bdfedd79c3a2e0e13a576102208401526102408301526102608201527f072c36ba67c6b3221e33c495797b900a8f9fc14a609f9ce8c84b69e903977cbc6102808201527f16902914c61a626430be7e3bb9b93e674610a506e297119b20dd8146d120fe1d6102a08201527f203ec77b656a8c2ae4825bc54ca7d727b532c05425071982c5b67e89e28886b56102c08201527f16c3a3a844ebe967bc4939399b78ce6bd519cfbc0697491d1f1c46cfa3d704216102e082015260405192839161113084846111aa565b8336843760085afa15908115611148575b506107f857005b6001915051141581611141565b905f905b6004821061116657505050565b6020806001928551815201930191019091611159565b9181601f840112156108165782359167ffffffffffffffff8311610816576020838186019501011161081657565b90601f8019910116810190811067ffffffffffffffff8211176111cc57604052565b634e487b7160e01b5f52604160045260245ffd5b906101a0828203126108165780601f8301121561081657604051916112076101a0846111aa565b82906101a0810192831161081657905b8282106112245750505090565b8135815260209182019101611217565b905f905b600d821061124557505050565b6020806001928551815201930191019091611238565b93929190610100811461133d576080811461127f5763236bd13760e01b5f5260045ffd5b6101a0830361132e578401916080858403126108165782601f8601121561081657604051926112af6080856111aa565b83956080810191821161081657955b81871061131e57505061131c939450816112dd916113069301906111e0565b6040516307846db360e21b6020820152926112fc906024850190611155565b60a4830190611234565b6102248152611317610244826111aa565b61187c565b565b86358152602096870196016112be565b630c0b7e3560e11b5f5260045ffd5b6101a0830361132e57840191610100858403126108165782601f86011215610816576040519261136f610100856111aa565b8395610100810191821161081657955b8187106113f3575050611397929394508101906111e0565b60405163f7f338eb60e01b6020820152915f602484015b600882106113dd57505050906113cc61131c92610124830190611234565b6102a481526113176102c4826111aa565b60208060019285518152019301910190916113ae565b863581526020968701960161137f565b8015611467578060011c915f516020611ab95f395f51905f528310156107f8576001806114465f516020611ab95f395f51905f52600381888181800909086118d9565b93161461144f57565b905f516020611ab95f395f51905f5280910681030690565b505f905f90565b80158061159f575b611593578060021c92825f516020611ab95f395f51905f52851080159061157c575b6107f85784815f516020611ab95f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816115469d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5086118fc565b80929160018082961614611558575050565b5f516020611ab95f395f51905f528093945080929550809106810306930681030690565b505f516020611ab95f395f51905f52811015611498565b50505f905f905f905f90565b508115611476565b905f516020611ab95f395f51905f528210801590611631575b6107f857811580611629575b611623576115f05f516020611ab95f395f51905f52600381858181800909086118d9565b8181036115ff57505060011b90565b5f516020611ab95f395f51905f52809106810306145f146107f857600190811b1790565b50505f90565b5080156115cc565b505f516020611ab95f395f51905f528110156115c0565b919093925f516020611ab95f395f51905f528310801590611865575b801561184e575b8015611837575b6107f857808286851717171561182c5790829161178f5f516020611ab95f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611ab95f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161176981808b800981878009086118d9565b8408095f516020611ab95f395f51905f5261178382611a50565b800914159586916118fc565b929080821480611823575b156117c15750505050905f146117b95760ff60025b169060021b179190565b60ff5f6117af565b5f516020611ab95f395f51905f52809106810306149182611804575b5050156107f857600191156117fc5760ff60025b169060021b17179190565b60ff5f6117f1565b5f516020611ab95f395f51905f52919250819006810306145f806117dd565b5083831461179a565b50505090505f905f90565b505f516020611ab95f395f51905f52811015611672565b505f516020611ab95f395f51905f5282101561166b565b505f516020611ab95f395f51905f52851015611664565b5f8091602081519101305afa3d156118d1573d9067ffffffffffffffff82116111cc57604051916118b7601f8201601f1916602001846111aa565b82523d5f602084013e5b156118c95750565b602081519101fd5b6060906118c1565b906118e382611a50565b915f516020611ab95f395f51905f52838009036107f857565b915f516020611ab95f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816119549396949661194682808a8009818a8009086118d9565b90611a44575b8608096118d9565b925f516020611ab95f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611ab95f395f51905f5260a083015260208260c08160055afa915191156107f8575f516020611ab95f395f51905f528260019209036107f8575f516020611ab95f395f51905f52908209925f516020611ab95f395f51905f528080808780090681030681878009081490811591611a25575b506107f857565b90505f516020611ab95f395f51905f528084860960020914155f611a1e565b8180910681030661194c565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611ab95f395f51905f5260a083015260208260c08160055afa915191156107f85756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220b54a5fe924d6c48c8a99cf7c448b388ec132a5a9a03038b546dcbcd486bfd9f064736f6c634300081c0033",
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
