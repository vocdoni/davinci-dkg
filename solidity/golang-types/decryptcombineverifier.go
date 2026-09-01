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

// DecryptCombineVerifierMetaData contains all meta data concerning the DecryptCombineVerifier contract.
var DecryptCombineVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[13]\",\"internalType\":\"uint256[13]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"input\",\"type\":\"uint256[13]\",\"internalType\":\"uint256[13]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidProofEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x608080604052346015576119cf908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80631e11b6cc146107f8578063233ace11146107be57806344f6369214610725578063b8e72af61461056a5763f7f338eb14610051575f80fd5b34610567576102a036600319011261056757366101041161056757366102a4116105675760405160408101905f51602061179a5f395f51905f52815260208101915f51602061155a5f395f51905f5283525f5160206113da5f395f51905f528152606082015f51602061181a5f395f51905f5281525f51602061169a5f395f51905f52604061010435935f5160206117ba5f395f51905f52608087019580875284848460608160075afa911016838860808160065afa165f51602061185a5f395f51905f5283525f51602061171a5f395f51905f5286526001610124359182895286868660608160075afa9310161616838860808160065afa165f51602061147a5f395f51905f5283525f51602061163a5f395f51905f528652610144359081885285858560608160075afa92101616838860808160065afa165f51602061167a5f395f51905f5283525f51602061191a5f395f51905f528652610164359081885285858560608160075afa92101616838860808160065afa165f5160206118fa5f395f51905f5283525f51602061161a5f395f51905f528652610184359081885285858560608160075afa92101616838860808160065afa165f5160206114fa5f395f51905f5283525f5160206115fa5f395f51905f5286526101a4359081885285858560608160075afa92101616838860808160065afa165f5160206118da5f395f51905f5283525f5160206113fa5f395f51905f5286526101c4359081885285858560608160075afa92101616838860808160065afa165f5160206113ba5f395f51905f5283525f51602061189a5f395f51905f5286526101e4359081885285858560608160075afa92101616838860808160065afa165f51602061193a5f395f51905f5283525f51602061141a5f395f51905f528652610204359081885285858560608160075afa92101616838860808160065afa165f5160206114ba5f395f51905f5283525f51602061149a5f395f51905f528652610224359081885285858560608160075afa92101616838860808160065afa165f5160206116da5f395f51905f5283525f51602061173a5f395f51905f528652610244359081885285858560608160075afa92101616838860808160065afa165f51602061153a5f395f51905f5283525f5160206118ba5f395f51905f528652610264359081885285858560608160075afa92101616838860808160065afa16945f51602061177a5f395f51905f528352526102843580955260608160075afa9210161660408260808160065afa169051915190156105585760405191610100600484375f51602061145a5f395f51905f526101008401525f51602061143a5f395f51905f526101208401525f5160206116ba5f395f51905f526101408401525f51602061187a5f395f51905f526101608401525f51602061197a5f395f51905f526101808401525f51602061159a5f395f51905f526101a08401525f51602061165a5f395f51905f526101c08401525f5160206117fa5f395f51905f526101e08401525f5160206116fa5f395f51905f526102008401525f5160206117da5f395f51905f526102208401526102408301526102608201525f5160206115ba5f395f51905f526102808201525f51602061151a5f395f51905f526102a08201525f51602061175a5f395f51905f526102c08201525f5160206115da5f395f51905f526102e08201526020816103008160085afa905116156105495780f35b631ff3747d60e21b8152600490fd5b63a54f8e2760e01b8352600483fd5b80fd5b50346106de5760403660031901126106de576004356001600160401b0381116106de5761059b903690600401610d89565b6024356001600160401b0381116106de576105ba903690600401610d89565b90916101008103610716578301610100848203126106de5780601f850112156106de57604051936105ed61010086610db6565b849061010081019283116106de57905b8282106107065750505081016101a0828203126106de5780601f830112156106de576040519161062f6101a084610db6565b82906101a081019283116106de57905b8282106106e257505050303b156106de5760405163f7f338eb60e01b8152915f600484015b600882106106c85750505061010482015f905b600d82106106b2575050505f816102a481305afa80156106a757610699575080f35b6106a591505f90610db6565b005b6040513d5f823e3d90fd5b6020806001928551815201930191019091610677565b6020806001928551815201930191019091610664565b5f80fd5b813581526020918201910161063f565b634e487b7160e01b5f52604160045260245ffd5b81358152602091820191016105fd565b63236bd13760e01b5f5260045ffd5b346106de576101003660031901126106de5736610104116106de5760405161074e608082610db6565b6080368237610761602435600435610f44565b815261077760843560a435604435606435610fe5565b6020830152604082015261078f60e43560c435610f44565b6060820152604051905f825b600482106107a857608084f35b602080600192855181520193019101909161079b565b346106de575f3660031901126106de5760206040517f192977b7b1dd4f4acfe86cb9c81cefa38856e9500b2f2168742e609ab3f9bff78152f35b346106de576102203660031901126106de57366084116106de5736610224116106de5761030060405161082b8282610db6565b8136823761083a600435610dd9565b61084b602493929335604435610e44565b9193929061085a606435610dd9565b9390926040519660408801965f51602061179a5f395f51905f5289528860208101985f51602061155a5f395f51905f528a525f5160206113da5f395f51905f5281525f51602061169a5f395f51905f52604060608401925f51602061181a5f395f51905f5284525f5160206117ba5f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f51602061185a5f395f51905f5285525f51602061171a5f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f51602061147a5f395f51905f5285525f51602061163a5f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f51602061167a5f395f51905f5285525f51602061191a5f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f5160206118fa5f395f51905f5285525f51602061161a5f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f5160206114fa5f395f51905f5285525f5160206115fa5f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f5160206118da5f395f51905f5285525f5160206113fa5f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f5160206113ba5f395f51905f5285525f51602061189a5f395f51905f5288526101643590818a5287838760608160075afa92101616818360808160065afa165f51602061193a5f395f51905f5285525f51602061141a5f395f51905f5288526101843590818a5287838760608160075afa92101616818360808160065afa165f5160206114ba5f395f51905f5285525f51602061149a5f395f51905f5288526101a43590818a5287838760608160075afa92101616818360808160065afa165f5160206116da5f395f51905f5285525f51602061173a5f395f51905f5288526101c43590818a5287838760608160075afa92101616818360808160065afa165f51602061153a5f395f51905f5285525f5160206118ba5f395f51905f5288526101e43590818a5287838760608160075afa921016169160808160065afa16945f51602061177a5f395f51905f528352526102043580955260608160075afa9210161660408a60808160065afa16985197519815610d7a5760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f51602061145a5f395f51905f526101008401525f51602061143a5f395f51905f526101208401525f5160206116ba5f395f51905f526101408401525f51602061187a5f395f51905f526101608401525f51602061197a5f395f51905f526101808401525f51602061159a5f395f51905f526101a08401525f51602061165a5f395f51905f526101c08401525f5160206117fa5f395f51905f526101e08401525f5160206116fa5f395f51905f526102008401525f5160206117da5f395f51905f526102208401526102408301526102608201525f5160206115ba5f395f51905f526102808201525f51602061151a5f395f51905f526102a08201525f51602061175a5f395f51905f526102c08201525f5160206115da5f395f51905f526102e0820152604051928391610d468484610db6565b8336843760085afa15908115610d6d575b50610d5e57005b631ff3747d60e21b5f5260045ffd5b6001915051141581610d57565b63a54f8e2760e01b5f5260045ffd5b9181601f840112156106de578235916001600160401b0383116106de57602083818601950101116106de57565b601f909101601f19168101906001600160401b038211908210176106f257604052565b8015610e3d578060011c915f5160206114da5f395f51905f52831015610d5e57600180610e1c5f5160206114da5f395f51905f52600381888181800909086111cd565b931614610e2557565b905f5160206114da5f395f51905f5280910681030690565b505f905f90565b801580610f3c575b610f30578060021c92825f5160206114da5f395f51905f528510801590610f19575b610d5e5784815f5160206114da5f395f51905f5280808080808080805f51602061183a5f395f51905f5281610ee39d8d0909998a0981898181800909085f51602061139a5f395f51905f520806810306936002808a16149509818a8181800909085f51602061195a5f395f51905f52086111f0565b80929160018082961614610ef5575050565b5f5160206114da5f395f51905f528093945080929550809106810306930681030690565b505f5160206114da5f395f51905f52811015610e6e565b50505f905f905f905f90565b508115610e4c565b905f5160206114da5f395f51905f528210801590610fce575b610d5e57811580610fc6575b610fc057610f8d5f5160206114da5f395f51905f52600381858181800909086111cd565b818103610f9c57505060011b90565b5f5160206114da5f395f51905f52809106810306145f14610d5e57600190811b1790565b50505f90565b508015610f69565b505f5160206114da5f395f51905f52811015610f5d565b919093925f5160206114da5f395f51905f5283108015906111b6575b801561119f575b8015611188575b610d5e57808286851717171561117d579082916110e05f5160206114da5f395f51905f5280808080888180808f9d5f51602061183a5f395f51905f528f839290839109099d8e0981848181800909085f51602061195a5f395f51905f52089a09818c8181800909085f51602061139a5f395f51905f520806810306945f5160206114da5f395f51905f525f51602061157a5f395f51905f52816110ba81808b800981878009086111cd565b8408095f5160206114da5f395f51905f526110d482611331565b800914159586916111f0565b929080821480611174575b156111125750505050905f1461110a5760ff60025b169060021b179190565b60ff5f611100565b5f5160206114da5f395f51905f52809106810306149182611155575b505015610d5e576001911561114d5760ff60025b169060021b17179190565b60ff5f611142565b5f5160206114da5f395f51905f52919250819006810306145f8061112e565b508383146110eb565b50505090505f905f90565b505f5160206114da5f395f51905f5281101561100f565b505f5160206114da5f395f51905f52821015611008565b505f5160206114da5f395f51905f52851015611001565b906111d782611331565b915f5160206114da5f395f51905f5283800903610d5e57565b915f5160206114da5f395f51905f525f51602061157a5f395f51905f52816112359396949661122782808a8009818a8009086111cd565b90611325575b8608096111cd565b925f5160206114da5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206114da5f395f51905f5260a083015260208260c08160055afa91519115610d5e575f5160206114da5f395f51905f52826001920903610d5e575f5160206114da5f395f51905f52908209925f5160206114da5f395f51905f528080808780090681030681878009081490811591611306575b50610d5e57565b90505f5160206114da5f395f51905f528084860960020914155f6112ff565b8180910681030661122d565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206114da5f395f51905f5260a083015260208260c08160055afa91519115610d5e5756fe2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750fc7c90afab484b23044be521084946c43ccaf3833604b6b3a399e1368fd742625caccc103522fce4f21608d8c97ee99676291284448b3c45a40a3d03d6dd79b03025e942529410a6851ca9be89e76a06469077deb472a0f3763662943630e1009079ae132b517fd3f8153e92a4259caab13419fd3cb4eb897e04297cb4bdb3a2143ecad846a2a5eec33089fdf94e0cd003994144b0ae1dcb98af0aff6eccdd91d048a960ce42c4145051ab52fc5840e73ade56670ef4be27450c4fec6e7953e004b5b2b53c8d90bd240f448431b8ff9f7890cd6a7db7359f540127f5e6105d22403b36df8ea8599c8c31c8cb56c326cacd63fcd5b4a74a6ef8c5f658f92e61c1417fc65145006ef5dbe8ef0aeea585a847647df2e0d243e73607bd0a688e7fc30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd470ce53cca7d2b3ac0dc5fa4a3dfa4d1d5fa0664d1af93bb3ec58f4af687fe89ed0eb291812015fde5501dbaf9d66b2aaab51ac85e53cb399a889d9c7b453ac24c0de53371faac18ec5086abc06a89bff9916bacbf04d198ab954fe5174f7c1717126f02f915b450b465673145a126833d50b4b50e2552b9276cecd9ef65f818b9183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea423a023c931808c9146e722457dbe7102c1d75606c6c24573d294b2e896e15cfd281b1b24de4be79cd5f4c8000de35e32b7173e06bea53faf8ed1ef3d9a61496f27e83a997782320162e4e8aa65f2655bd5b706bff5ec3a3d5c575c034c2883f22d06fe45dbb2207930da094de2b032e23ea9fbb257a7ae2261704d1cd367379604a9fe4cc941285ef483d240498cd79580f632243bd084966a76bdc38501772c2b2f65aa2873fa92d39320db22c9812d21af7c993418983b619f14238d2cefe51f8386ebd103c618410c5545dc01aa82599d0e71f79d88680c928275e2c58e361b92e5251586a1df1c80969e966d858bee2e32f4c3a0ef67c4e718037018687a30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001254101a761ecca2ce6422fd230ce36bd6425120ec26c39a3aad17d27447c32ed29fc5f9314b81526dd0d4e865af9319cc0cd40bcb7aeada011e2f35c47a16bef0d31f6d0942c33817a55c1bcd080ecbd7fc09b3e2fa29013db5c18e741374340056f0647eb623f9d2dc89f3d55a411124a174572d166bb3a3670eba4240fd1e50da3a6067fd43c2a99dea94e33a97d6132fe2eaed24a4a99cb84f0e6c071fc7b1dee8b0b8bcaf639f2e600f93bbdc5af03dbbf53fb84133514063f4d625f2a230a8311bb88e8faa1913e06f49cadef22ebd9f7373b3d37dfadaa982b35e157190ed2a81664ac5255fd03fd37cd5ed222ff03f91220b3f119eff5d4501824bbb11008f04243ae3565e22fddcbff1d7d9d5adc6c209bc7c10f067016ca7cee0a4d08654b0b68d3d298306064e7dd04ca8a637049a38a33eb8ba111a71abb3605ec033c86471dabfbd03a3359d7f2ae8ac9a41a06998bea87d297e0e49853658fdc0feb904a03e4b2fca5e160b74291e2321798b1c9bf52d3886d602e72a0486d3b30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd440ddd085ee77ea1455e79056efd13b38383af15d0464cbca9512d28228104f59a19b610d31ce17c1a4ecc990492304b58f134ab7c7f7c7657c21d0aa8af204a360a76bb50bf87570621cc2d07c2826407bade81b2271b222377194f9c3593ca171036f82a3277b1f4fd1f55ff674fec9b57db9a1254d6e0d4f57134784f56ade1260ea5aa541d17f5283f52cf9badd285d3861fc326c3f002b16b018283aca14c018bb0a0f1a1a735478ef31721c352aa35b9b8a84ec3c05d0d63448a7ee3f59c2f6e7f720e894a14e593ccb584f47ab386db969eb65cf2a5520d0eb927dedbd9043153dc5eb7f5d7a7277bfb1297485aeb03692e7b6557977cab08bebbff9b3e2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50b4bf19d650c1ca646033c3121a8f3dc3ad62416168e88bb802978c46422369aa2646970667358221220e6bc5276e01d33dea6ad9e5777aeac92f9a5392ab715a7a161ced4008f283f5764736f6c634300081c0033",
}

// DecryptCombineVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use DecryptCombineVerifierMetaData.ABI instead.
var DecryptCombineVerifierABI = DecryptCombineVerifierMetaData.ABI

// DecryptCombineVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DecryptCombineVerifierMetaData.Bin instead.
var DecryptCombineVerifierBin = DecryptCombineVerifierMetaData.Bin

// DeployDecryptCombineVerifier deploys a new Ethereum contract, binding an instance of DecryptCombineVerifier to it.
func DeployDecryptCombineVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DecryptCombineVerifier, error) {
	parsed, err := DecryptCombineVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DecryptCombineVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DecryptCombineVerifier{DecryptCombineVerifierCaller: DecryptCombineVerifierCaller{contract: contract}, DecryptCombineVerifierTransactor: DecryptCombineVerifierTransactor{contract: contract}, DecryptCombineVerifierFilterer: DecryptCombineVerifierFilterer{contract: contract}}, nil
}

// DecryptCombineVerifier is an auto generated Go binding around an Ethereum contract.
type DecryptCombineVerifier struct {
	DecryptCombineVerifierCaller     // Read-only binding to the contract
	DecryptCombineVerifierTransactor // Write-only binding to the contract
	DecryptCombineVerifierFilterer   // Log filterer for contract events
}

// DecryptCombineVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type DecryptCombineVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecryptCombineVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DecryptCombineVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecryptCombineVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DecryptCombineVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecryptCombineVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DecryptCombineVerifierSession struct {
	Contract     *DecryptCombineVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DecryptCombineVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DecryptCombineVerifierCallerSession struct {
	Contract *DecryptCombineVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// DecryptCombineVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DecryptCombineVerifierTransactorSession struct {
	Contract     *DecryptCombineVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// DecryptCombineVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type DecryptCombineVerifierRaw struct {
	Contract *DecryptCombineVerifier // Generic contract binding to access the raw methods on
}

// DecryptCombineVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DecryptCombineVerifierCallerRaw struct {
	Contract *DecryptCombineVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// DecryptCombineVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DecryptCombineVerifierTransactorRaw struct {
	Contract *DecryptCombineVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDecryptCombineVerifier creates a new instance of DecryptCombineVerifier, bound to a specific deployed contract.
func NewDecryptCombineVerifier(address common.Address, backend bind.ContractBackend) (*DecryptCombineVerifier, error) {
	contract, err := bindDecryptCombineVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DecryptCombineVerifier{DecryptCombineVerifierCaller: DecryptCombineVerifierCaller{contract: contract}, DecryptCombineVerifierTransactor: DecryptCombineVerifierTransactor{contract: contract}, DecryptCombineVerifierFilterer: DecryptCombineVerifierFilterer{contract: contract}}, nil
}

// NewDecryptCombineVerifierCaller creates a new read-only instance of DecryptCombineVerifier, bound to a specific deployed contract.
func NewDecryptCombineVerifierCaller(address common.Address, caller bind.ContractCaller) (*DecryptCombineVerifierCaller, error) {
	contract, err := bindDecryptCombineVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DecryptCombineVerifierCaller{contract: contract}, nil
}

// NewDecryptCombineVerifierTransactor creates a new write-only instance of DecryptCombineVerifier, bound to a specific deployed contract.
func NewDecryptCombineVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*DecryptCombineVerifierTransactor, error) {
	contract, err := bindDecryptCombineVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DecryptCombineVerifierTransactor{contract: contract}, nil
}

// NewDecryptCombineVerifierFilterer creates a new log filterer instance of DecryptCombineVerifier, bound to a specific deployed contract.
func NewDecryptCombineVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*DecryptCombineVerifierFilterer, error) {
	contract, err := bindDecryptCombineVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DecryptCombineVerifierFilterer{contract: contract}, nil
}

// bindDecryptCombineVerifier binds a generic wrapper to an already deployed contract.
func bindDecryptCombineVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DecryptCombineVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DecryptCombineVerifier *DecryptCombineVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DecryptCombineVerifier.Contract.DecryptCombineVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DecryptCombineVerifier *DecryptCombineVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DecryptCombineVerifier.Contract.DecryptCombineVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DecryptCombineVerifier *DecryptCombineVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DecryptCombineVerifier.Contract.DecryptCombineVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DecryptCombineVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DecryptCombineVerifier *DecryptCombineVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DecryptCombineVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DecryptCombineVerifier *DecryptCombineVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DecryptCombineVerifier.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_DecryptCombineVerifier *DecryptCombineVerifierCaller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int) ([4]*big.Int, error) {
	var out []interface{}
	err := _DecryptCombineVerifier.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_DecryptCombineVerifier *DecryptCombineVerifierSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _DecryptCombineVerifier.Contract.CompressProof(&_DecryptCombineVerifier.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _DecryptCombineVerifier.Contract.CompressProof(&_DecryptCombineVerifier.CallOpts, proof)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_DecryptCombineVerifier *DecryptCombineVerifierCaller) ProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DecryptCombineVerifier.contract.Call(opts, &out, "provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_DecryptCombineVerifier *DecryptCombineVerifierSession) ProvingKeyHash() ([32]byte, error) {
	return _DecryptCombineVerifier.Contract.ProvingKeyHash(&_DecryptCombineVerifier.CallOpts)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerSession) ProvingKeyHash() ([32]byte, error) {
	return _DecryptCombineVerifier.Contract.ProvingKeyHash(&_DecryptCombineVerifier.CallOpts)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x1e11b6cc.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[13] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [13]*big.Int) error {
	var out []interface{}
	err := _DecryptCombineVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x1e11b6cc.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[13] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [13]*big.Int) error {
	return _DecryptCombineVerifier.Contract.VerifyCompressedProof(&_DecryptCombineVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x1e11b6cc.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[13] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [13]*big.Int) error {
	return _DecryptCombineVerifier.Contract.VerifyCompressedProof(&_DecryptCombineVerifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCaller) VerifyProof(opts *bind.CallOpts, proof []byte, input []byte) error {
	var out []interface{}
	err := _DecryptCombineVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierSession) VerifyProof(proof []byte, input []byte) error {
	return _DecryptCombineVerifier.Contract.VerifyProof(&_DecryptCombineVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerSession) VerifyProof(proof []byte, input []byte) error {
	return _DecryptCombineVerifier.Contract.VerifyProof(&_DecryptCombineVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xf7f338eb.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[13] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCaller) VerifyProof0(opts *bind.CallOpts, proof [8]*big.Int, input [13]*big.Int) error {
	var out []interface{}
	err := _DecryptCombineVerifier.contract.Call(opts, &out, "verifyProof0", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xf7f338eb.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[13] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierSession) VerifyProof0(proof [8]*big.Int, input [13]*big.Int) error {
	return _DecryptCombineVerifier.Contract.VerifyProof0(&_DecryptCombineVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xf7f338eb.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[13] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerSession) VerifyProof0(proof [8]*big.Int, input [13]*big.Int) error {
	return _DecryptCombineVerifier.Contract.VerifyProof0(&_DecryptCombineVerifier.CallOpts, proof, input)
}
