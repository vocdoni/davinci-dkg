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
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"input\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidInputEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x608080604052346015576115c0908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace11146109f55750806344f636921461095c5780634830667114610529578063a6047e6c146101545763b8e72af614610055575f80fd5b34610132576040366003190112610132576004356001600160401b03811161013257610085903690600401610a2d565b906024356001600160401b038111610132576100a5903690600401610a2d565b92610100810361014557610100840361013657826100c8916100d1940190610a91565b92810190610a91565b303b1561013257610102906100f7604051936329811f9b60e21b85526004850190610ae3565b610104830190610ae3565b5f8161020481305afa801561012757610119575080f35b61012591505f90610a5a565b005b6040513d5f823e3d90fd5b5f80fd5b630c0b7e3560e11b5f5260045ffd5b63236bd13760e01b5f5260045ffd5b34610132576102003660031901126101325736610104116101325736610204116101325760405160408101905f51602061154b5f395f51905f52815260208101915f5160206113ab5f395f51905f5283525f5160206111ab5f395f51905f528152606082015f51602061124b5f395f51905f5281525f51602061138b5f395f51905f52604061010435935f5160206113cb5f395f51905f52608087019580875284848460608160075afa911016838860808160065afa165f51602061128b5f395f51905f5283525f51602061140b5f395f51905f5286526001610124359182895286868660608160075afa9310161616838860808160065afa165f5160206114eb5f395f51905f5283525f5160206111cb5f395f51905f528652610144359081885285858560608160075afa92101616838860808160065afa165f51602061152b5f395f51905f5283525f5160206112cb5f395f51905f528652610164359081885285858560608160075afa92101616838860808160065afa165f5160206114cb5f395f51905f5283525f51602061136b5f395f51905f528652610184359081885285858560608160075afa92101616838860808160065afa165f51602061146b5f395f51905f5283525f51602061114b5f395f51905f5286526101a4359081885285858560608160075afa92101616838860808160065afa165f51602061142b5f395f51905f5283525f51602061134b5f395f51905f5286526101c4359081885285858560608160075afa92101616838860808160065afa16945f51602061116b5f395f51905f528352526101e43580955260608160075afa9210161660408260808160065afa1690519151901561051a5760405191610100600484375f51602061132b5f395f51905f526101008401525f51602061122b5f395f51905f526101208401525f5160206110eb5f395f51905f526101408401525f5160206114ab5f395f51905f526101608401525f51602061118b5f395f51905f526101808401525f51602061148b5f395f51905f526101a08401525f51602061110b5f395f51905f526101c08401525f51602061126b5f395f51905f526101e08401525f51602061150b5f395f51905f526102008401525f5160206113eb5f395f51905f526102208401526102408301526102608201525f5160206112ab5f395f51905f526102808201525f51602061112b5f395f51905f526102a08201525f51602061120b5f395f51905f526102c08201525f51602061130b5f395f51905f526102e08201526020816103008160085afa9051161561050b57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b346101325761018036600319011261013257366084116101325736610184116101325761030060405161055c8282610a5a565b8136823761056b600435610d93565b61057c602493929335604435610dfe565b9193929061058b606435610d93565b9390926040519660408801965f51602061154b5f395f51905f5289528860208101985f5160206113ab5f395f51905f528a525f5160206111ab5f395f51905f5281525f51602061138b5f395f51905f52604060608401925f51602061124b5f395f51905f5284525f5160206113cb5f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f51602061128b5f395f51905f5285525f51602061140b5f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f5160206114eb5f395f51905f5285525f5160206111cb5f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f51602061152b5f395f51905f5285525f5160206112cb5f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f5160206114cb5f395f51905f5285525f51602061136b5f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f51602061146b5f395f51905f5285525f51602061114b5f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f51602061142b5f395f51905f5285525f51602061134b5f395f51905f5288526101443590818a5287838760608160075afa921016169160808160065afa16945f51602061116b5f395f51905f528352526101643580955260608160075afa9210161660408a60808160065afa1698519751981561051a5760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f51602061132b5f395f51905f526101008401525f51602061122b5f395f51905f526101208401525f5160206110eb5f395f51905f526101408401525f5160206114ab5f395f51905f526101608401525f51602061118b5f395f51905f526101808401525f51602061148b5f395f51905f526101a08401525f51602061110b5f395f51905f526101c08401525f51602061126b5f395f51905f526101e08401525f51602061150b5f395f51905f526102008401525f5160206113eb5f395f51905f526102208401526102408301526102608201525f5160206112ab5f395f51905f526102808201525f51602061112b5f395f51905f526102a08201525f51602061120b5f395f51905f526102c08201525f51602061130b5f395f51905f526102e08201526040519283916109378484610a5a565b8336843760085afa1590811561094f575b5061050b57005b6001915051141581610948565b346101325761010036600319011261013257366101041161013257604051610985608082610a5a565b6080368237610998602435600435610b0a565b81526109ae60843560a435604435606435610bab565b602083015260408201526109c660e43560c435610b0a565b6060820152604051905f825b600482106109df57608084f35b60208060019285518152019301910190916109d2565b34610132575f36600319011261013257807fd747a935a6680c8b6446389ab994a7d939aa12b0def66c921f85fc984b1e69d960209252f35b9181601f84011215610132578235916001600160401b038311610132576020838186019501011161013257565b601f909101601f19168101906001600160401b03821190821017610a7d57604052565b634e487b7160e01b5f52604160045260245ffd5b90610100828203126101325780601f830112156101325761010060405192610ab98285610a5a565b8391810192831161013257905b828210610ad35750505090565b8135815260209182019101610ac6565b905f905b60088210610af457505050565b6020806001928551815201930191019091610ae7565b905f5160206111eb5f395f51905f528210801590610b94575b61050b57811580610b8c575b610b8657610b535f5160206111eb5f395f51905f5260038185818180090908610efe565b818103610b6257505060011b90565b5f5160206111eb5f395f51905f52809106810306145f1461050b57600190811b1790565b50505f90565b508015610b2f565b505f5160206111eb5f395f51905f52811015610b23565b919093925f5160206111eb5f395f51905f528310801590610d7c575b8015610d65575b8015610d4e575b61050b578082868517171715610d4357908291610ca65f5160206111eb5f395f51905f5280808080888180808f9d5f51602061144b5f395f51905f528f839290839109099d8e0981848181800909085f51602061156b5f395f51905f52089a09818c8181800909085f5160206110cb5f395f51905f520806810306945f5160206111eb5f395f51905f525f5160206112eb5f395f51905f5281610c8081808b80098187800908610efe565b8408095f5160206111eb5f395f51905f52610c9a82611062565b80091415958691610f21565b929080821480610d3a575b15610cd85750505050905f14610cd05760ff60025b169060021b179190565b60ff5f610cc6565b5f5160206111eb5f395f51905f52809106810306149182610d1b575b50501561050b5760019115610d135760ff60025b169060021b17179190565b60ff5f610d08565b5f5160206111eb5f395f51905f52919250819006810306145f80610cf4565b50838314610cb1565b50505090505f905f90565b505f5160206111eb5f395f51905f52811015610bd5565b505f5160206111eb5f395f51905f52821015610bce565b505f5160206111eb5f395f51905f52851015610bc7565b8015610df7578060011c915f5160206111eb5f395f51905f5283101561050b57600180610dd65f5160206111eb5f395f51905f5260038188818180090908610efe565b931614610ddf57565b905f5160206111eb5f395f51905f5280910681030690565b505f905f90565b801580610ef6575b610eea578060021c92825f5160206111eb5f395f51905f528510801590610ed3575b61050b5784815f5160206111eb5f395f51905f5280808080808080805f51602061144b5f395f51905f5281610e9d9d8d0909998a0981898181800909085f5160206110cb5f395f51905f520806810306936002808a16149509818a8181800909085f51602061156b5f395f51905f5208610f21565b80929160018082961614610eaf575050565b5f5160206111eb5f395f51905f528093945080929550809106810306930681030690565b505f5160206111eb5f395f51905f52811015610e28565b50505f905f905f905f90565b508115610e06565b90610f0882611062565b915f5160206111eb5f395f51905f528380090361050b57565b915f5160206111eb5f395f51905f525f5160206112eb5f395f51905f5281610f6693969496610f5882808a8009818a800908610efe565b90611056575b860809610efe565b925f5160206111eb5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206111eb5f395f51905f5260a083015260208260c08160055afa9151911561050b575f5160206111eb5f395f51905f5282600192090361050b575f5160206111eb5f395f51905f52908209925f5160206111eb5f395f51905f528080808780090681030681878009081490811591611037575b5061050b57565b90505f5160206111eb5f395f51905f528084860960020914155f611030565b81809106810306610f5e565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206111eb5f395f51905f5260a083015260208260c08160055afa9151911561050b5756fe2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775026f9a032930c11258400468f829c469e18540a6c3d74eb90cf45a751518e14f0d372e87fb0eba9197cf967a63322a6d69710eba34536d7f218172b830d445b41d5967d8c0730431736737274e8a94bc7ebce6ccdd7f11683253c4e950cec27b0fa8d22519d4a1a38f4859c181c1e7c58f2e76795ffa710bd73bb26b4c841c6a0b0b93e809491216a84909052399d676dd16d3e19ac4b84dcb5704b814e0724b2130e7f91727b18dd53616d663c7bc6bdbfd9a21801b931704b9fce6386f41412739bf3679e1bcb6307329617d1e1cf11290270b5977a464cd11b0e932a8f0bd16dce1c5aba4cbc918d755a4849790c4ce3cbd8afd305033498251856045820f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd472bcaa2bdc1174fcca0002cbbeb5036c0b610d1a577725a68b4bb7abb50fdb1cb1e6017fd86e71811a30c11afaef40f9021d51e726a77a44d173b660b4257c53c024400d7ac761ae5ac97374e44fdae6d21b3b3be85bdcc043055edadecf23fdb12b58e82036d63fb3f7369a1ad4cca51f3595353b352bdb330d16e2e9398e0572375cdd2501e1d523486a577e9e6bf292291855b09d2176c92c4c58add83693328f873305dc6fcbeea4e3c7aecd0fbd690981f724406e9ab0b89a0efb5d87bf500d7765ac65fbe01ea882d3bd61d1a7ce074c78baf86efc4eb6e3aa27b15627b183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea41a7cc5be87df529341ba85ca3abe1262e9b110e7f804a6e845f4e1d608990f3511347a3f2a0a3162e10d75294bfa4919c65f69493b6bec48d078a0001ff6d0c91a9fc9639fa2b987f48f0aa961b3544bc3decc76a7003935b4962435d918cbc800a39800043fb38d6c0b5df9b12f499bea7d564ba9bf69c27df51a674ac6e8e330644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000010f91734bb0ad69cf4f70913c877db9ea163b509b15d7b0324448a7a151299fb41970c3e50ae7ec5123e90210641fdf82c94c8befb405b45a104aa6607d29c7c62fed65097e48c8d78a8bb229e320da0c14692e5092eef287460400796f9717f908e3082133043e6cd541d429277f667d74f84ef0961fc08f4c3941b1411044b7147cd79ad580f275e4b48d71220833469f1ea34fccf58e0f5aa1d915a758c84130644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4416e6fedcb05880336b2172f6dc46f68a55d370ab20f384e1d379919ade3a993f1ecb00bc78dd88f87fec15af2df2c63cac8761b9110b743987b6219f7a07ecf725e2b605f8ba6f425d184089ac2a36d119b6dff7078337c27a5441d3f8c9e4071c4546671e7161161ff4746317683a9b40ee5cc255268d4cea1a6599507aaf29085c1719d6b4e6aea79f565e0c94e3d7588a9d000e89e0852dd5e77a98d235420c9c4fd9be09acb83913f05f60cbeb1b72cf2f589a0f98b1e8f62e0b1c35054528ebf299c47e33d26bb490416208a4a03a148dc2e161b3c470ed3d690505c594083fb2adf3d6de8c1db84f64ebcc48fdd43dc31e4793cb92b698346874e49aa72b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5a2646970667358221220bf99b135b2bbad1e4fb9745f7cdeeb467571342bcf4e2855d6c81cc6cc460ea964736f6c634300081c0033",
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

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x48306671.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [8]*big.Int) error {
	var out []interface{}
	err := _ContributionVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x48306671.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [8]*big.Int) error {
	return _ContributionVerifier.Contract.VerifyCompressedProof(&_ContributionVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x48306671.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [8]*big.Int) error {
	return _ContributionVerifier.Contract.VerifyCompressedProof(&_ContributionVerifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xa6047e6c.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, input [8]*big.Int) error {
	var out []interface{}
	err := _ContributionVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0xa6047e6c.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierSession) VerifyProof(proof [8]*big.Int, input [8]*big.Int) error {
	return _ContributionVerifier.Contract.VerifyProof(&_ContributionVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xa6047e6c.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierCallerSession) VerifyProof(proof [8]*big.Int, input [8]*big.Int) error {
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
