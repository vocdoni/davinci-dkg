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
	Bin: "0x608080604052346015576119cf908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80631e11b6cc146107f8578063233ace11146107be57806344f6369214610725578063b8e72af61461056a5763f7f338eb14610051575f80fd5b34610567576102a036600319011261056757366101041161056757366102a4116105675760405160408101905f51602061141a5f395f51905f52815260208101915f5160206118fa5f395f51905f5283525f51602061183a5f395f51905f528152606082015f51602061165a5f395f51905f5281525f5160206116ba5f395f51905f52604061010435935f5160206115da5f395f51905f52608087019580875284848460608160075afa911016838860808160065afa165f5160206117fa5f395f51905f5283525f51602061157a5f395f51905f5286526001610124359182895286868660608160075afa9310161616838860808160065afa165f51602061171a5f395f51905f5283525f51602061151a5f395f51905f528652610144359081885285858560608160075afa92101616838860808160065afa165f51602061185a5f395f51905f5283525f51602061161a5f395f51905f528652610164359081885285858560608160075afa92101616838860808160065afa165f51602061195a5f395f51905f5283525f5160206114ba5f395f51905f528652610184359081885285858560608160075afa92101616838860808160065afa165f51602061189a5f395f51905f5283525f5160206118da5f395f51905f5286526101a4359081885285858560608160075afa92101616838860808160065afa165f51602061181a5f395f51905f5283525f51602061159a5f395f51905f5286526101c4359081885285858560608160075afa92101616838860808160065afa165f51602061153a5f395f51905f5283525f51602061163a5f395f51905f5286526101e4359081885285858560608160075afa92101616838860808160065afa165f51602061169a5f395f51905f5283525f51602061173a5f395f51905f528652610204359081885285858560608160075afa92101616838860808160065afa165f5160206114fa5f395f51905f5283525f51602061147a5f395f51905f528652610224359081885285858560608160075afa92101616838860808160065afa165f5160206113fa5f395f51905f5283525f5160206114da5f395f51905f528652610244359081885285858560608160075afa92101616838860808160065afa165f51602061155a5f395f51905f5283525f5160206117da5f395f51905f528652610264359081885285858560608160075afa92101616838860808160065afa16945f5160206115fa5f395f51905f528352526102843580955260608160075afa9210161660408260808160065afa169051915190156105585760405191610100600484375f5160206113da5f395f51905f526101008401525f51602061149a5f395f51905f526101208401525f5160206113ba5f395f51905f526101408401525f51602061167a5f395f51905f526101608401525f51602061143a5f395f51905f526101808401525f51602061191a5f395f51905f526101a08401525f5160206118ba5f395f51905f526101c08401525f51602061193a5f395f51905f526101e08401525f5160206116da5f395f51905f526102008401525f51602061177a5f395f51905f526102208401526102408301526102608201525f51602061187a5f395f51905f526102808201525f51602061179a5f395f51905f526102a08201525f51602061175a5f395f51905f526102c08201525f5160206116fa5f395f51905f526102e08201526020816103008160085afa905116156105495780f35b631ff3747d60e21b8152600490fd5b63a54f8e2760e01b8352600483fd5b80fd5b50346106de5760403660031901126106de576004356001600160401b0381116106de5761059b903690600401610d89565b6024356001600160401b0381116106de576105ba903690600401610d89565b90916101008103610716578301610100848203126106de5780601f850112156106de57604051936105ed61010086610db6565b849061010081019283116106de57905b8282106107065750505081016101a0828203126106de5780601f830112156106de576040519161062f6101a084610db6565b82906101a081019283116106de57905b8282106106e257505050303b156106de5760405163f7f338eb60e01b8152915f600484015b600882106106c85750505061010482015f905b600d82106106b2575050505f816102a481305afa80156106a757610699575080f35b6106a591505f90610db6565b005b6040513d5f823e3d90fd5b6020806001928551815201930191019091610677565b6020806001928551815201930191019091610664565b5f80fd5b813581526020918201910161063f565b634e487b7160e01b5f52604160045260245ffd5b81358152602091820191016105fd565b63236bd13760e01b5f5260045ffd5b346106de576101003660031901126106de5736610104116106de5760405161074e608082610db6565b6080368237610761602435600435610f44565b815261077760843560a435604435606435610fe5565b6020830152604082015261078f60e43560c435610f44565b6060820152604051905f825b600482106107a857608084f35b602080600192855181520193019101909161079b565b346106de575f3660031901126106de5760206040517f45db382e20ec94816f3fd8bb0fd1c54a5a550eb6a842c26790fac5d5583a442f8152f35b346106de576102203660031901126106de57366084116106de5736610224116106de5761030060405161082b8282610db6565b8136823761083a600435610dd9565b61084b602493929335604435610e44565b9193929061085a606435610dd9565b9390926040519660408801965f51602061141a5f395f51905f5289528860208101985f5160206118fa5f395f51905f528a525f51602061183a5f395f51905f5281525f5160206116ba5f395f51905f52604060608401925f51602061165a5f395f51905f5284525f5160206115da5f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f5160206117fa5f395f51905f5285525f51602061157a5f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f51602061171a5f395f51905f5285525f51602061151a5f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f51602061185a5f395f51905f5285525f51602061161a5f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f51602061195a5f395f51905f5285525f5160206114ba5f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f51602061189a5f395f51905f5285525f5160206118da5f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f51602061181a5f395f51905f5285525f51602061159a5f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f51602061153a5f395f51905f5285525f51602061163a5f395f51905f5288526101643590818a5287838760608160075afa92101616818360808160065afa165f51602061169a5f395f51905f5285525f51602061173a5f395f51905f5288526101843590818a5287838760608160075afa92101616818360808160065afa165f5160206114fa5f395f51905f5285525f51602061147a5f395f51905f5288526101a43590818a5287838760608160075afa92101616818360808160065afa165f5160206113fa5f395f51905f5285525f5160206114da5f395f51905f5288526101c43590818a5287838760608160075afa92101616818360808160065afa165f51602061155a5f395f51905f5285525f5160206117da5f395f51905f5288526101e43590818a5287838760608160075afa921016169160808160065afa16945f5160206115fa5f395f51905f528352526102043580955260608160075afa9210161660408a60808160065afa16985197519815610d7a5760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f5160206113da5f395f51905f526101008401525f51602061149a5f395f51905f526101208401525f5160206113ba5f395f51905f526101408401525f51602061167a5f395f51905f526101608401525f51602061143a5f395f51905f526101808401525f51602061191a5f395f51905f526101a08401525f5160206118ba5f395f51905f526101c08401525f51602061193a5f395f51905f526101e08401525f5160206116da5f395f51905f526102008401525f51602061177a5f395f51905f526102208401526102408301526102608201525f51602061187a5f395f51905f526102808201525f51602061179a5f395f51905f526102a08201525f51602061175a5f395f51905f526102c08201525f5160206116fa5f395f51905f526102e0820152604051928391610d468484610db6565b8336843760085afa15908115610d6d575b50610d5e57005b631ff3747d60e21b5f5260045ffd5b6001915051141581610d57565b63a54f8e2760e01b5f5260045ffd5b9181601f840112156106de578235916001600160401b0383116106de57602083818601950101116106de57565b601f909101601f19168101906001600160401b038211908210176106f257604052565b8015610e3d578060011c915f51602061145a5f395f51905f52831015610d5e57600180610e1c5f51602061145a5f395f51905f52600381888181800909086111cd565b931614610e2557565b905f51602061145a5f395f51905f5280910681030690565b505f905f90565b801580610f3c575b610f30578060021c92825f51602061145a5f395f51905f528510801590610f19575b610d5e5784815f51602061145a5f395f51905f5280808080808080805f5160206117ba5f395f51905f5281610ee39d8d0909998a0981898181800909085f51602061139a5f395f51905f520806810306936002808a16149509818a8181800909085f51602061197a5f395f51905f52086111f0565b80929160018082961614610ef5575050565b5f51602061145a5f395f51905f528093945080929550809106810306930681030690565b505f51602061145a5f395f51905f52811015610e6e565b50505f905f905f905f90565b508115610e4c565b905f51602061145a5f395f51905f528210801590610fce575b610d5e57811580610fc6575b610fc057610f8d5f51602061145a5f395f51905f52600381858181800909086111cd565b818103610f9c57505060011b90565b5f51602061145a5f395f51905f52809106810306145f14610d5e57600190811b1790565b50505f90565b508015610f69565b505f51602061145a5f395f51905f52811015610f5d565b919093925f51602061145a5f395f51905f5283108015906111b6575b801561119f575b8015611188575b610d5e57808286851717171561117d579082916110e05f51602061145a5f395f51905f5280808080888180808f9d5f5160206117ba5f395f51905f528f839290839109099d8e0981848181800909085f51602061197a5f395f51905f52089a09818c8181800909085f51602061139a5f395f51905f520806810306945f51602061145a5f395f51905f525f5160206115ba5f395f51905f52816110ba81808b800981878009086111cd565b8408095f51602061145a5f395f51905f526110d482611331565b800914159586916111f0565b929080821480611174575b156111125750505050905f1461110a5760ff60025b169060021b179190565b60ff5f611100565b5f51602061145a5f395f51905f52809106810306149182611155575b505015610d5e576001911561114d5760ff60025b169060021b17179190565b60ff5f611142565b5f51602061145a5f395f51905f52919250819006810306145f8061112e565b508383146110eb565b50505090505f905f90565b505f51602061145a5f395f51905f5281101561100f565b505f51602061145a5f395f51905f52821015611008565b505f51602061145a5f395f51905f52851015611001565b906111d782611331565b915f51602061145a5f395f51905f5283800903610d5e57565b915f51602061145a5f395f51905f525f5160206115ba5f395f51905f52816112359396949661122782808a8009818a8009086111cd565b90611325575b8608096111cd565b925f51602061145a5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f51602061145a5f395f51905f5260a083015260208260c08160055afa91519115610d5e575f51602061145a5f395f51905f52826001920903610d5e575f51602061145a5f395f51905f52908209925f51602061145a5f395f51905f528080808780090681030681878009081490811591611306575b50610d5e57565b90505f51602061145a5f395f51905f528084860960020914155f6112ff565b8180910681030661122d565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f51602061145a5f395f51905f5260a083015260208260c08160055afa91519115610d5e5756fe2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775249f7f346953737519ea5d7cdadeb473ef2ab4e1ca63502e704e13991ccb558a2d33a4fbbb9676aa1ca6cc1213c90dcead5d19a0aea01577d581b81651cafad7223ee7aea78251ca6532a58252383795892038bbf9ae82c051065fe5d0266c12301158b216db295597bf6b1726dcd8293a14997ab9ad77cfaccf18a23e8a576113f5e57e9fa01b3ca8450ec77a171b283fc261f486acd5ae5ca02f11412d0a0230644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd470d64f751f620427e73b86d4d5ea213101dee4160555b19f4bf5a3a3c6535162d19fefd9dff39fc7bb2616bf6a3376b18220f2d42d226481d1938561f5358d80d126560a3e90b0536cf94a1df1672f67c38c856e34f7fecfa930ee39fcb9c4dd22124ec9cfffefe6f731f0c72723fc682db20ca313ac5648d4b994e19c665634e195c5ed918f57a9b2d1127e1e2ee33e4bfef00966e4999ada6c4fc0c000305cd203f6390cce3bd5c7132aff0a38b377538a33f5a215ea059c87136dadeb26dbd0444b0c28afe295b3d653d27d7864f2ba9d047fa95be21c88aa38b272a9135cd2230ca37eb15a6f3a047941cc83576f34610ff4d20f2edfc6b84abd581171aae0f6af107362ddbcae3d6300bfe6ffda96987b40439039fb33f7ed2f6088097ea2f5eb47a4102d1ba69c3af94f5202273e60e196655e66ac715f00141d991e2cf183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4004ce7ffcaaa944aec2a11b8a836500579dd85e233d83efa83c1e3c177479dad0d3747a9cbf590224d779d3e34629d52162cc0155036ff4c74899f837faec2262b8693b1d4f7e9dae37fafbf7721c7f3b1213e55e951ee5aaee599bf20ab668013b68fb383cfb0eec1243a94c51d552caae8f46d2cb03fe8887c47d5d08ca62511f2ce0f00902e1e872e50140fe831fe4c82458f26a54313f68d46f2e2edadd2164c27074e1865bc9e8bb54316ec9cf48980605bf020d8892bf7cfb2da89c0222dad76d3075b4e646f65b0309b264fdf7dc7a8c3013d1b86d6c4697a0db87b8430644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000115971583ce5a49b81491f184f52e3b8daf86c22c46db4fef3f36301a663dbb2416344b2a7242876f3f10738fe82d5ae2d98ccaee197759f0ed0de1339d0c789828b68646c1bf9f87be84e4450325da25de8037825b728b990ef0014a0a0c048f0ea27b1e4c4a3c7b42395816dab840ae6a36d058c65ae909430d7116d70e4e441b2ebf372bd2e1aa36d91086d05feff1802b161583b8194937450cb89be84b0324c90759820602a5cd4c664d3686d4ba41acc9db3d913b93c4f8be1547300e3b149c235233a840402abfc94a8986156e9ac10943ac7177e1a3519926470452c630644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd440210f3dc8e97d8a61b8e2b655f1662eea9760daf550d090c9bd33012829263550601864b7ed105c2a51d84f8f7d3749d9650337c6093d841eba9589bdaf06a4400dca34695e4b0f3b187f17ac64089037a8df1017b1729390b2c8006a77ecf6e0d9a7029306a7c10dc90f2a5e91518bcf9399d465a7004aee42a8e8b24df7d892655fc79d7db586aa32381c716e996ab516b54be031722150ee74b129383a7931d5499a818688d8eeadf5fecf372d6ac6d6a74179f4d43dab63a026504252f022118d644dde5cc5090d52b5c759964cb98655dae495eb373946006d86c0df9620cf35bf44f2bc74dffcbd34ae3f55d920f11d9c9fd7b64b50bbcfcd9c3fa34b71a433a61c795b5a0d5a364c3988327c3089267ce3a3f78267762e00e1ebe846602ea88ad2e0563ddf12c23f37313c3672539f61701b09ac001f2e3ba564e66e908fc21124fb01d02e9aa69d5c77ecef063073ef8471b1b3964fd7edb6481871b16515bd00819fc053f7ce3690befb5efb7fe2cb88f63eca459158354b2fb856202587406a9dc003c3be067ee35d5421bb7223b02b429e1b2f268ee2e8d0206cc2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5a26469706673582212204d9bdfdf771599123a42d61e48310a95b8fc64925199863ce23da9ab762706f464736f6c634300081c0033",
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
