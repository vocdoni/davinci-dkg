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
	Bin: "0x608080604052346015576119cf908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80631e11b6cc146107f8578063233ace11146107be57806344f6369214610725578063b8e72af61461056a5763f7f338eb14610051575f80fd5b34610567576102a036600319011261056757366101041161056757366102a4116105675760405160408101905f51602061183a5f395f51905f52815260208101915f5160206116ba5f395f51905f5283525f5160206115fa5f395f51905f528152606082015f51602061165a5f395f51905f5281525f51602061167a5f395f51905f52604061010435935f51602061191a5f395f51905f52608087019580875284848460608160075afa911016838860808160065afa165f5160206113fa5f395f51905f5283525f51602061149a5f395f51905f5286526001610124359182895286868660608160075afa9310161616838860808160065afa165f51602061173a5f395f51905f5283525f51602061155a5f395f51905f528652610144359081885285858560608160075afa92101616838860808160065afa165f51602061151a5f395f51905f5283525f51602061193a5f395f51905f528652610164359081885285858560608160075afa92101616838860808160065afa165f51602061169a5f395f51905f5283525f51602061159a5f395f51905f528652610184359081885285858560608160075afa92101616838860808160065afa165f51602061197a5f395f51905f5283525f51602061185a5f395f51905f5286526101a4359081885285858560608160075afa92101616838860808160065afa165f5160206114ba5f395f51905f5283525f5160206116da5f395f51905f5286526101c4359081885285858560608160075afa92101616838860808160065afa165f51602061175a5f395f51905f5283525f5160206118da5f395f51905f5286526101e4359081885285858560608160075afa92101616838860808160065afa165f51602061157a5f395f51905f5283525f51602061163a5f395f51905f528652610204359081885285858560608160075afa92101616838860808160065afa165f51602061145a5f395f51905f5283525f5160206118ba5f395f51905f528652610224359081885285858560608160075afa92101616838860808160065afa165f51602061143a5f395f51905f5283525f5160206116fa5f395f51905f528652610244359081885285858560608160075afa92101616838860808160065afa165f5160206113ba5f395f51905f5283525f51602061179a5f395f51905f528652610264359081885285858560608160075afa92101616838860808160065afa16945f51602061141a5f395f51905f528352526102843580955260608160075afa9210161660408260808160065afa169051915190156105585760405191610100600484375f51602061177a5f395f51905f526101008401525f5160206113da5f395f51905f526101208401525f51602061161a5f395f51905f526101408401525f51602061181a5f395f51905f526101608401525f5160206117fa5f395f51905f526101808401525f5160206115da5f395f51905f526101a08401525f5160206118fa5f395f51905f526101c08401525f51602061187a5f395f51905f526101e08401525f51602061153a5f395f51905f526102008401525f51602061171a5f395f51905f526102208401526102408301526102608201525f51602061189a5f395f51905f526102808201525f5160206117da5f395f51905f526102a08201525f51602061147a5f395f51905f526102c08201525f5160206114fa5f395f51905f526102e08201526020816103008160085afa905116156105495780f35b631ff3747d60e21b8152600490fd5b63a54f8e2760e01b8352600483fd5b80fd5b50346106de5760403660031901126106de576004356001600160401b0381116106de5761059b903690600401610d89565b6024356001600160401b0381116106de576105ba903690600401610d89565b90916101008103610716578301610100848203126106de5780601f850112156106de57604051936105ed61010086610db6565b849061010081019283116106de57905b8282106107065750505081016101a0828203126106de5780601f830112156106de576040519161062f6101a084610db6565b82906101a081019283116106de57905b8282106106e257505050303b156106de5760405163f7f338eb60e01b8152915f600484015b600882106106c85750505061010482015f905b600d82106106b2575050505f816102a481305afa80156106a757610699575080f35b6106a591505f90610db6565b005b6040513d5f823e3d90fd5b6020806001928551815201930191019091610677565b6020806001928551815201930191019091610664565b5f80fd5b813581526020918201910161063f565b634e487b7160e01b5f52604160045260245ffd5b81358152602091820191016105fd565b63236bd13760e01b5f5260045ffd5b346106de576101003660031901126106de5736610104116106de5760405161074e608082610db6565b6080368237610761602435600435610f44565b815261077760843560a435604435606435610fe5565b6020830152604082015261078f60e43560c435610f44565b6060820152604051905f825b600482106107a857608084f35b602080600192855181520193019101909161079b565b346106de575f3660031901126106de5760206040517fe16f45b822e85899f3d9f1972077f428531cb04e885c653d0716463d15433cf28152f35b346106de576102203660031901126106de57366084116106de5736610224116106de5761030060405161082b8282610db6565b8136823761083a600435610dd9565b61084b602493929335604435610e44565b9193929061085a606435610dd9565b9390926040519660408801965f51602061183a5f395f51905f5289528860208101985f5160206116ba5f395f51905f528a525f5160206115fa5f395f51905f5281525f51602061167a5f395f51905f52604060608401925f51602061165a5f395f51905f5284525f51602061191a5f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f5160206113fa5f395f51905f5285525f51602061149a5f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f51602061173a5f395f51905f5285525f51602061155a5f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f51602061151a5f395f51905f5285525f51602061193a5f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f51602061169a5f395f51905f5285525f51602061159a5f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f51602061197a5f395f51905f5285525f51602061185a5f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f5160206114ba5f395f51905f5285525f5160206116da5f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f51602061175a5f395f51905f5285525f5160206118da5f395f51905f5288526101643590818a5287838760608160075afa92101616818360808160065afa165f51602061157a5f395f51905f5285525f51602061163a5f395f51905f5288526101843590818a5287838760608160075afa92101616818360808160065afa165f51602061145a5f395f51905f5285525f5160206118ba5f395f51905f5288526101a43590818a5287838760608160075afa92101616818360808160065afa165f51602061143a5f395f51905f5285525f5160206116fa5f395f51905f5288526101c43590818a5287838760608160075afa92101616818360808160065afa165f5160206113ba5f395f51905f5285525f51602061179a5f395f51905f5288526101e43590818a5287838760608160075afa921016169160808160065afa16945f51602061141a5f395f51905f528352526102043580955260608160075afa9210161660408a60808160065afa16985197519815610d7a5760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f51602061177a5f395f51905f526101008401525f5160206113da5f395f51905f526101208401525f51602061161a5f395f51905f526101408401525f51602061181a5f395f51905f526101608401525f5160206117fa5f395f51905f526101808401525f5160206115da5f395f51905f526101a08401525f5160206118fa5f395f51905f526101c08401525f51602061187a5f395f51905f526101e08401525f51602061153a5f395f51905f526102008401525f51602061171a5f395f51905f526102208401526102408301526102608201525f51602061189a5f395f51905f526102808201525f5160206117da5f395f51905f526102a08201525f51602061147a5f395f51905f526102c08201525f5160206114fa5f395f51905f526102e0820152604051928391610d468484610db6565b8336843760085afa15908115610d6d575b50610d5e57005b631ff3747d60e21b5f5260045ffd5b6001915051141581610d57565b63a54f8e2760e01b5f5260045ffd5b9181601f840112156106de578235916001600160401b0383116106de57602083818601950101116106de57565b601f909101601f19168101906001600160401b038211908210176106f257604052565b8015610e3d578060011c915f5160206114da5f395f51905f52831015610d5e57600180610e1c5f5160206114da5f395f51905f52600381888181800909086111cd565b931614610e2557565b905f5160206114da5f395f51905f5280910681030690565b505f905f90565b801580610f3c575b610f30578060021c92825f5160206114da5f395f51905f528510801590610f19575b610d5e5784815f5160206114da5f395f51905f5280808080808080805f5160206117ba5f395f51905f5281610ee39d8d0909998a0981898181800909085f51602061139a5f395f51905f520806810306936002808a16149509818a8181800909085f51602061195a5f395f51905f52086111f0565b80929160018082961614610ef5575050565b5f5160206114da5f395f51905f528093945080929550809106810306930681030690565b505f5160206114da5f395f51905f52811015610e6e565b50505f905f905f905f90565b508115610e4c565b905f5160206114da5f395f51905f528210801590610fce575b610d5e57811580610fc6575b610fc057610f8d5f5160206114da5f395f51905f52600381858181800909086111cd565b818103610f9c57505060011b90565b5f5160206114da5f395f51905f52809106810306145f14610d5e57600190811b1790565b50505f90565b508015610f69565b505f5160206114da5f395f51905f52811015610f5d565b919093925f5160206114da5f395f51905f5283108015906111b6575b801561119f575b8015611188575b610d5e57808286851717171561117d579082916110e05f5160206114da5f395f51905f5280808080888180808f9d5f5160206117ba5f395f51905f528f839290839109099d8e0981848181800909085f51602061195a5f395f51905f52089a09818c8181800909085f51602061139a5f395f51905f520806810306945f5160206114da5f395f51905f525f5160206115ba5f395f51905f52816110ba81808b800981878009086111cd565b8408095f5160206114da5f395f51905f526110d482611331565b800914159586916111f0565b929080821480611174575b156111125750505050905f1461110a5760ff60025b169060021b179190565b60ff5f611100565b5f5160206114da5f395f51905f52809106810306149182611155575b505015610d5e576001911561114d5760ff60025b169060021b17179190565b60ff5f611142565b5f5160206114da5f395f51905f52919250819006810306145f8061112e565b508383146110eb565b50505090505f905f90565b505f5160206114da5f395f51905f5281101561100f565b505f5160206114da5f395f51905f52821015611008565b505f5160206114da5f395f51905f52851015611001565b906111d782611331565b915f5160206114da5f395f51905f5283800903610d5e57565b915f5160206114da5f395f51905f525f5160206115ba5f395f51905f52816112359396949661122782808a8009818a8009086111cd565b90611325575b8608096111cd565b925f5160206114da5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206114da5f395f51905f5260a083015260208260c08160055afa91519115610d5e575f5160206114da5f395f51905f52826001920903610d5e575f5160206114da5f395f51905f52908209925f5160206114da5f395f51905f528080808780090681030681878009081490811591611306575b50610d5e57565b90505f5160206114da5f395f51905f528084860960020914155f6112ff565b8180910681030661122d565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206114da5f395f51905f5260a083015260208260c08160055afa91519115610d5e5756fe2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7751ae0ecd4eb14313309467863a8bd4819e06952daec6947994a82050088daf20e2fca0d397146a3d8fd450258e7ad3317b416ab61b8148f899003928cea8f249e232eeb6b102330d96e93ffc5764b948562c2e01ae9a7f18114ac2e1f34a49cce1b1f18f238585b9cd701f3ecb2b619bf3ef0217a60438cd03cf73c2a3ca6640623e5989aaa82179a2c322f955f6e8bba0f6b23e41c2d0f89dffb9c989ac4f0841e790ada729892a03fc307a0cbf6d1372fa19a77e8f7409473039abf61ab807117373ad29ed13284c10aac91dcc2492dac78a0a59436df10e090ca3612b850530300376fcdb707810971d4f5b8b3f7565d98edcbd63d0242425690f3509ad8672324b372e1557a06302e8bbba2a8f7861c5a8e6530a73036c160a41d40ddaf3230644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4703f4383a19d685948f09a95ab675c6ed74d52347a5c9d15a2f76c435cd1de6021c239cba35d541c944a5e1834ceef72b7a76b44103ecbc178fc3c2e97f1e95f309bf34d02541d8fff7c6aee2acff8eef84985d1a391fe8b8c6dd0d61a45f08a120b16e0397aab6dfaec694e0151827f91c7755f55105ca797e902146942a963e27e512b088466f62629326e48ffa7ea01caeaf0435ee8a751552ca4d15095d970dc8c41730358aaa353d0faea45c04690d9cce25d42946928ab9ed28b3019318183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea424dac6f34fbee5b3047ab6d7a9ebfc95e5e0686a391a1da0c138af8d0840aa4b008a0ac816c540dab4e1697d3e7cb4bc8a9fabad08cb75715fe06502f851819b275586f220269c8309f4aee77d427d5b56ae0be42891b9a00dabab49976548fa0903a03a9e99157bd53b4ccdb30f37a3d52fa5569ae895f62067a1b67251cd22213421b8283cd8aa0d8ed60a0410845417973ba110cfc04f3ca843b525ed267430644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000105746596911f9310447c8cba16d7b71e3121dcccfdc17d2b633a50f22d4b1880160c75b7bbff8dc83b8a8281a5323ca447f75ed30fb65ba2dda98d7f1c84c3e801ec0ddc413d715149e1d6af9070054008e6844f8c61870479335ff2e11e35f626fcb30bd572e8d8602a9544abe163f0bbad7da048a2184540652f8ccc364284264a8d158bb5290c38daa4815125f35d76deadf995d930e721507dc918f4bfb524e11002ff1d7da7c127354db636b733cf0fa96455925659d6657a507fffd789260537be78cf0a364162c69fbde3d7ee901aeac798b8b5c3aa1b7992fcbc553a108986f8ebb7e50f3abc46c9da532a3186a8d962fbea450fe105cd994aa4891b257d8f6258779773ac872b054f46337e156640ef1961e4afca4cd6c1a7a2030030644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd440cd20623e0006bfab9dc5b3be557abac03e180b3b56a60591a8775effb96feeb2be37dac692769b185743044ad6d58aec309860dc1c395a673f1faefb4f0adfe10cc368bc3d5754a6b6dc08be5123659d0529c4be8f8ca75fa26551bafff9e1e083b0152ad9ed61fa3fd60b899c9f493517559189e5cec35067aa811588c89272fabb18e0527667d0f192ab65cdec6ec35dcd20448e52a1391388957a8d36095295c73e7e2abfbc56600f33eb19a4f5f966a05f11a75417a84038f1450c7806511b071067db5e2e8dde9f57467519b80629632a92747ca9be1306e030330bf9c208604b061b06138813824118d93c80053b647c9923a15659ee91bd7d36e8a021453470b49354ec84ec5825327c4432b6d275bb4e6192abe88c60641711d49e60898b13fb709e684c7e5cc7173b9e259004d8f175eb29cb6110c360b536c518b09a50c17458fa6bb0daf58e93b93f7f84c156813d140a19e83cf6ce900f5045d13071e8915cbed7659e70d485da557fc75fba51acf21785e716927a6666e37532b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50389d8c972b81e127b5325037bd2337c13f2c81b5274f94cb73fe87a7f6b1072a264697066735822122023f0b29b3ac7d40ce7ad538f606807289d09a9c20162dc9cfdbc98b38c91debe64736f6c634300081c0033",
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
