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
	Bin: "0x608080604052346015576119cf908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80631e11b6cc146107f8578063233ace11146107be57806344f6369214610725578063b8e72af61461056a5763f7f338eb14610051575f80fd5b34610567576102a036600319011261056757366101041161056757366102a4116105675760405160408101905f51602061173a5f395f51905f52815260208101915f51602061161a5f395f51905f5283525f51602061191a5f395f51905f528152606082015f51602061181a5f395f51905f5281525f51602061177a5f395f51905f52604061010435935f51602061139a5f395f51905f52608087019580875284848460608160075afa911016838860808160065afa165f51602061153a5f395f51905f5283525f51602061175a5f395f51905f5286526001610124359182895286868660608160075afa9310161616838860808160065afa165f5160206117fa5f395f51905f5283525f51602061171a5f395f51905f528652610144359081885285858560608160075afa92101616838860808160065afa165f51602061179a5f395f51905f5283525f5160206118fa5f395f51905f528652610164359081885285858560608160075afa92101616838860808160065afa165f51602061163a5f395f51905f5283525f51602061147a5f395f51905f528652610184359081885285858560608160075afa92101616838860808160065afa165f51602061193a5f395f51905f5283525f5160206118ba5f395f51905f5286526101a4359081885285858560608160075afa92101616838860808160065afa165f5160206117ba5f395f51905f5283525f5160206114da5f395f51905f5286526101c4359081885285858560608160075afa92101616838860808160065afa165f51602061183a5f395f51905f5283525f51602061165a5f395f51905f5286526101e4359081885285858560608160075afa92101616838860808160065afa165f5160206113ba5f395f51905f5283525f51602061141a5f395f51905f528652610204359081885285858560608160075afa92101616838860808160065afa165f51602061157a5f395f51905f5283525f51602061143a5f395f51905f528652610224359081885285858560608160075afa92101616838860808160065afa165f5160206115fa5f395f51905f5283525f5160206118da5f395f51905f528652610244359081885285858560608160075afa92101616838860808160065afa165f51602061189a5f395f51905f5283525f51602061169a5f395f51905f528652610264359081885285858560608160075afa92101616838860808160065afa16945f5160206117da5f395f51905f528352526102843580955260608160075afa9210161660408260808160065afa169051915190156105585760405191610100600484375f51602061145a5f395f51905f526101008401525f51602061155a5f395f51905f526101208401525f5160206114fa5f395f51905f526101408401525f5160206116ba5f395f51905f526101608401525f51602061149a5f395f51905f526101808401525f51602061187a5f395f51905f526101a08401525f51602061197a5f395f51905f526101c08401525f5160206115ba5f395f51905f526101e08401525f51602061151a5f395f51905f526102008401525f5160206116fa5f395f51905f526102208401526102408301526102608201525f5160206115da5f395f51905f526102808201525f5160206113fa5f395f51905f526102a08201525f51602061167a5f395f51905f526102c08201525f5160206114ba5f395f51905f526102e08201526020816103008160085afa905116156105495780f35b631ff3747d60e21b8152600490fd5b63a54f8e2760e01b8352600483fd5b80fd5b50346106de5760403660031901126106de576004356001600160401b0381116106de5761059b903690600401610d89565b6024356001600160401b0381116106de576105ba903690600401610d89565b90916101008103610716578301610100848203126106de5780601f850112156106de57604051936105ed61010086610db6565b849061010081019283116106de57905b8282106107065750505081016101a0828203126106de5780601f830112156106de576040519161062f6101a084610db6565b82906101a081019283116106de57905b8282106106e257505050303b156106de5760405163f7f338eb60e01b8152915f600484015b600882106106c85750505061010482015f905b600d82106106b2575050505f816102a481305afa80156106a757610699575080f35b6106a591505f90610db6565b005b6040513d5f823e3d90fd5b6020806001928551815201930191019091610677565b6020806001928551815201930191019091610664565b5f80fd5b813581526020918201910161063f565b634e487b7160e01b5f52604160045260245ffd5b81358152602091820191016105fd565b63236bd13760e01b5f5260045ffd5b346106de576101003660031901126106de5736610104116106de5760405161074e608082610db6565b6080368237610761602435600435610f44565b815261077760843560a435604435606435610fe5565b6020830152604082015261078f60e43560c435610f44565b6060820152604051905f825b600482106107a857608084f35b602080600192855181520193019101909161079b565b346106de575f3660031901126106de5760206040517fd8248e4c88328c26fd2135bae946e618485207eb738a69f8d2c56b07bd631a218152f35b346106de576102203660031901126106de57366084116106de5736610224116106de5761030060405161082b8282610db6565b8136823761083a600435610dd9565b61084b602493929335604435610e44565b9193929061085a606435610dd9565b9390926040519660408801965f51602061173a5f395f51905f5289528860208101985f51602061161a5f395f51905f528a525f51602061191a5f395f51905f5281525f51602061177a5f395f51905f52604060608401925f51602061181a5f395f51905f5284525f51602061139a5f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f51602061153a5f395f51905f5285525f51602061175a5f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f5160206117fa5f395f51905f5285525f51602061171a5f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f51602061179a5f395f51905f5285525f5160206118fa5f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f51602061163a5f395f51905f5285525f51602061147a5f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f51602061193a5f395f51905f5285525f5160206118ba5f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f5160206117ba5f395f51905f5285525f5160206114da5f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f51602061183a5f395f51905f5285525f51602061165a5f395f51905f5288526101643590818a5287838760608160075afa92101616818360808160065afa165f5160206113ba5f395f51905f5285525f51602061141a5f395f51905f5288526101843590818a5287838760608160075afa92101616818360808160065afa165f51602061157a5f395f51905f5285525f51602061143a5f395f51905f5288526101a43590818a5287838760608160075afa92101616818360808160065afa165f5160206115fa5f395f51905f5285525f5160206118da5f395f51905f5288526101c43590818a5287838760608160075afa92101616818360808160065afa165f51602061189a5f395f51905f5285525f51602061169a5f395f51905f5288526101e43590818a5287838760608160075afa921016169160808160065afa16945f5160206117da5f395f51905f528352526102043580955260608160075afa9210161660408a60808160065afa16985197519815610d7a5760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f51602061145a5f395f51905f526101008401525f51602061155a5f395f51905f526101208401525f5160206114fa5f395f51905f526101408401525f5160206116ba5f395f51905f526101608401525f51602061149a5f395f51905f526101808401525f51602061187a5f395f51905f526101a08401525f51602061197a5f395f51905f526101c08401525f5160206115ba5f395f51905f526101e08401525f51602061151a5f395f51905f526102008401525f5160206116fa5f395f51905f526102208401526102408301526102608201525f5160206115da5f395f51905f526102808201525f5160206113fa5f395f51905f526102a08201525f51602061167a5f395f51905f526102c08201525f5160206114ba5f395f51905f526102e0820152604051928391610d468484610db6565b8336843760085afa15908115610d6d575b50610d5e57005b631ff3747d60e21b5f5260045ffd5b6001915051141581610d57565b63a54f8e2760e01b5f5260045ffd5b9181601f840112156106de578235916001600160401b0383116106de57602083818601950101116106de57565b601f909101601f19168101906001600160401b038211908210176106f257604052565b8015610e3d578060011c915f51602061159a5f395f51905f52831015610d5e57600180610e1c5f51602061159a5f395f51905f52600381888181800909086111cd565b931614610e2557565b905f51602061159a5f395f51905f5280910681030690565b505f905f90565b801580610f3c575b610f30578060021c92825f51602061159a5f395f51905f528510801590610f19575b610d5e5784815f51602061159a5f395f51905f5280808080808080805f51602061185a5f395f51905f5281610ee39d8d0909998a0981898181800909085f5160206113da5f395f51905f520806810306936002808a16149509818a8181800909085f51602061195a5f395f51905f52086111f0565b80929160018082961614610ef5575050565b5f51602061159a5f395f51905f528093945080929550809106810306930681030690565b505f51602061159a5f395f51905f52811015610e6e565b50505f905f905f905f90565b508115610e4c565b905f51602061159a5f395f51905f528210801590610fce575b610d5e57811580610fc6575b610fc057610f8d5f51602061159a5f395f51905f52600381858181800909086111cd565b818103610f9c57505060011b90565b5f51602061159a5f395f51905f52809106810306145f14610d5e57600190811b1790565b50505f90565b508015610f69565b505f51602061159a5f395f51905f52811015610f5d565b919093925f51602061159a5f395f51905f5283108015906111b6575b801561119f575b8015611188575b610d5e57808286851717171561117d579082916110e05f51602061159a5f395f51905f5280808080888180808f9d5f51602061185a5f395f51905f528f839290839109099d8e0981848181800909085f51602061195a5f395f51905f52089a09818c8181800909085f5160206113da5f395f51905f520806810306945f51602061159a5f395f51905f525f5160206116da5f395f51905f52816110ba81808b800981878009086111cd565b8408095f51602061159a5f395f51905f526110d482611331565b800914159586916111f0565b929080821480611174575b156111125750505050905f1461110a5760ff60025b169060021b179190565b60ff5f611100565b5f51602061159a5f395f51905f52809106810306149182611155575b505015610d5e576001911561114d5760ff60025b169060021b17179190565b60ff5f611142565b5f51602061159a5f395f51905f52919250819006810306145f8061112e565b508383146110eb565b50505090505f905f90565b505f51602061159a5f395f51905f5281101561100f565b505f51602061159a5f395f51905f52821015611008565b505f51602061159a5f395f51905f52851015611001565b906111d782611331565b915f51602061159a5f395f51905f5283800903610d5e57565b915f51602061159a5f395f51905f525f5160206116da5f395f51905f52816112359396949661122782808a8009818a8009086111cd565b90611325575b8608096111cd565b925f51602061159a5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f51602061159a5f395f51905f5260a083015260208260c08160055afa91519115610d5e575f51602061159a5f395f51905f52826001920903610d5e575f51602061159a5f395f51905f52908209925f51602061159a5f395f51905f528080808780090681030681878009081490811591611306575b50610d5e57565b90505f51602061159a5f395f51905f528084860960020914155f6112ff565b8180910681030661122d565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f51602061159a5f395f51905f5260a083015260208260c08160055afa91519115610d5e5756fe23e46018e5423c5c58419ee02c3f76ca42ba9178c415e426cf2d04aa0b7b77a809772e7f413bc678085a6238b3849dd9785214619ec0d23cf3b3fe1dcd6282f02fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750dce4806925f96742a6fdfd19a2a4d6eab672bb48e1f0f1f2101276974b2f7f0205b0421e914e2c44f7828fffeca27e122162bf1bdf76f01ef3dd930099459d4236b677d6947bda2f2c10ac2f7f618792e9d1ba33f12d649af6e602417d59fc90378a3b71e801ff7f94aeba538cb7dccc013f0e1b0fb4d7b1abe4716449c5ec025894367bbcf222138c028a90f484f8d4699b97a775dccbc511138f2896dadf10d41b6ae879f568e36b5566e62426f10b3e72b6ea0f489733c74dd17466dc9702a087f8a501a431ad998cf02638f781628e3d3847523061038f34c45022a7d772c58952f6be49c89d25df42ce410114d4f8340ca1fc3edf8e4252840d8706f252f2ca56f7e3cc3d9c0676d678134207e61ace1571e3bce685788172f3021e2c50e8adfc0df5855a484589c8dbc0c06c1b8f9acf8fdcf8ce2eaafdacd88e04ec5048bbf17159d99a0caa49f8f8d5c4b937bea9decee4e307d8cd525b618daf5790027ba552aa40e568b3c307e66b8b4b5c2a2d75255af5fb34a3a4c23e81fca0d28f9af14a07d738a7cfcd2ce36f0faa2c50c3b97f8cf6dde5ab1395b1403e94430644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47014014c232e83facf4e95e5b36ff98b4a9c654c190894a513d7699c816c4650b043fadfb6d4a851b937cebd6e73d4e70de4a2709bc1e603c79db9593620926612e43dfd9d6dc490c91118f37b80ba5d476540f83aa3b446a1b42f4b695a0309903fa4dd53840340c2bb888d408ea1817f2d7c1dbc20193afd63578fa9a9813441843e145df1ee68530c7df85f7682adb5527e2fcf4686df82b5cf22824b0526a080114eef2b055bcfeba8abb6cac15190b2507a5873597fe18e8beccf2c7aef11480337ae3daf03824238028bcfed326a809a0c2e6d119e5cd809e0f5ce8f5f808455d4d3e774ed9a7fc513cc994283fb3f21d804bf3dfd7ab5c5e27ce16e6dc155dbaf26ffa76e6c8d255451c37acbf8f531f7a02e0dbb4850e627a46680481183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea41cb2f4728d9a10cef7543f9ebfc35cdc6bd0b70f87681505a096e642546eccee1a4fdb6fffffa3ddbb49ae58d09031cb8cb330a1d7145751037e6a36a03fdb0400586f386ed9aa84cb757094564b18803de745e6f1d2d14b338bf0c2f2a0f3af06700aca86924f455cced59bbdc171b748b995908ff7946b3180748f5cb5822730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000012a28e7b583d204f45e38687dfd5a8555a7d4fef7f823145401d40780f096d47b2a95e23481c8d7e04a88d8ff3d9f69e9feab4895c530c134605793a6162ffaa2217960010b07385d3f1e14c01f89623d48036ca21262b8ae6e285bce615fb6231bbc7df5f7c6d192c6a8892c9ad08256e21ac89ede1ddaaf5359c690ca38c65325065b38d892be93a1b61848617723efe9d7a89c077e2dd3546d98688e4aca7603dc43069e2e310153e9f316e234ca500060b9326f44d93b77587cf356d4bdf530644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd440b972a178f3ffcd1b810769bdf1023a3cce6a567fdb7fac9a0e9b982f822fdef07662025728a9b0001eecd4e540daf6389bbf97062dd077d781067fa92046349234e24a5cfc4a2197d92fe91b662f10191486c997bd0589c1b421516b3de07f51ebfcaaeac8b91cb9a255ae3fd2b9381fff87082af173de37538333f56ba56da0871819697838c61c4a21c4dc8ec3a356a33d86c7445dffe859ae63d1571fe5e10963498a3c2dec3310de41153d5192705b392d106550621ed6f736f9b439caf1bcbab5846ee6e6f410e370be30aa661c20c16b22f57e40321684fd533a4291d2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e52bc11c817bf7dbbeafcb0edffeefaf7ce8eece318fbdb8f2285b64e35e81afb5a264697066735822122059dc6574372978788349823a269922afd033d7d5a6fb8a4a2b6d37be9f149dfb64736f6c634300081c0033",
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
