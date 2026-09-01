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
	Bin: "0x608080604052346015576115c0908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace11146109f55750806344f636921461095c5780634830667114610529578063a6047e6c146101545763b8e72af614610055575f80fd5b34610132576040366003190112610132576004356001600160401b03811161013257610085903690600401610a2d565b906024356001600160401b038111610132576100a5903690600401610a2d565b92610100810361014557610100840361013657826100c8916100d1940190610a91565b92810190610a91565b303b1561013257610102906100f7604051936329811f9b60e21b85526004850190610ae3565b610104830190610ae3565b5f8161020481305afa801561012757610119575080f35b61012591505f90610a5a565b005b6040513d5f823e3d90fd5b5f80fd5b630c0b7e3560e11b5f5260045ffd5b63236bd13760e01b5f5260045ffd5b34610132576102003660031901126101325736610104116101325736610204116101325760405160408101905f51602061136b5f395f51905f52815260208101915f5160206112cb5f395f51905f5283525f5160206110eb5f395f51905f528152606082015f51602061144b5f395f51905f5281525f51602061130b5f395f51905f52604061010435935f51602061122b5f395f51905f52608087019580875284848460608160075afa911016838860808160065afa165f51602061120b5f395f51905f5283525f51602061142b5f395f51905f5286526001610124359182895286868660608160075afa9310161616838860808160065afa165f51602061118b5f395f51905f5283525f5160206114eb5f395f51905f528652610144359081885285858560608160075afa92101616838860808160065afa165f51602061154b5f395f51905f5283525f51602061128b5f395f51905f528652610164359081885285858560608160075afa92101616838860808160065afa165f5160206114ab5f395f51905f5283525f5160206112ab5f395f51905f528652610184359081885285858560608160075afa92101616838860808160065afa165f51602061148b5f395f51905f5283525f5160206113cb5f395f51905f5286526101a4359081885285858560608160075afa92101616838860808160065afa165f51602061112b5f395f51905f5283525f51602061138b5f395f51905f5286526101c4359081885285858560608160075afa92101616838860808160065afa16945f5160206112eb5f395f51905f528352526101e43580955260608160075afa9210161660408260808160065afa1690519151901561051a5760405191610100600484375f51602061146b5f395f51905f526101008401525f51602061114b5f395f51905f526101208401525f5160206113ab5f395f51905f526101408401525f51602061156b5f395f51905f526101608401525f51602061132b5f395f51905f526101808401525f51602061116b5f395f51905f526101a08401525f5160206111eb5f395f51905f526101c08401525f51602061134b5f395f51905f526101e08401525f5160206111ab5f395f51905f526102008401525f51602061140b5f395f51905f526102208401526102408301526102608201525f51602061110b5f395f51905f526102808201525f5160206114cb5f395f51905f526102a08201525f51602061150b5f395f51905f526102c08201525f51602061126b5f395f51905f526102e08201526020816103008160085afa9051161561050b57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b346101325761018036600319011261013257366084116101325736610184116101325761030060405161055c8282610a5a565b8136823761056b600435610d93565b61057c602493929335604435610dfe565b9193929061058b606435610d93565b9390926040519660408801965f51602061136b5f395f51905f5289528860208101985f5160206112cb5f395f51905f528a525f5160206110eb5f395f51905f5281525f51602061130b5f395f51905f52604060608401925f51602061144b5f395f51905f5284525f51602061122b5f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f51602061120b5f395f51905f5285525f51602061142b5f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f51602061118b5f395f51905f5285525f5160206114eb5f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f51602061154b5f395f51905f5285525f51602061128b5f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f5160206114ab5f395f51905f5285525f5160206112ab5f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f51602061148b5f395f51905f5285525f5160206113cb5f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f51602061112b5f395f51905f5285525f51602061138b5f395f51905f5288526101443590818a5287838760608160075afa921016169160808160065afa16945f5160206112eb5f395f51905f528352526101643580955260608160075afa9210161660408a60808160065afa1698519751981561051a5760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f51602061146b5f395f51905f526101008401525f51602061114b5f395f51905f526101208401525f5160206113ab5f395f51905f526101408401525f51602061156b5f395f51905f526101608401525f51602061132b5f395f51905f526101808401525f51602061116b5f395f51905f526101a08401525f5160206111eb5f395f51905f526101c08401525f51602061134b5f395f51905f526101e08401525f5160206111ab5f395f51905f526102008401525f51602061140b5f395f51905f526102208401526102408301526102608201525f51602061110b5f395f51905f526102808201525f5160206114cb5f395f51905f526102a08201525f51602061150b5f395f51905f526102c08201525f51602061126b5f395f51905f526102e08201526040519283916109378484610a5a565b8336843760085afa1590811561094f575b5061050b57005b6001915051141581610948565b346101325761010036600319011261013257366101041161013257604051610985608082610a5a565b6080368237610998602435600435610b0a565b81526109ae60843560a435604435606435610bab565b602083015260408201526109c660e43560c435610b0a565b6060820152604051905f825b600482106109df57608084f35b60208060019285518152019301910190916109d2565b34610132575f36600319011261013257807f3b7f6cd34cff4e0a52cb97665f76a4c735cae0719e10bf2e5795a563ed47235e60209252f35b9181601f84011215610132578235916001600160401b038311610132576020838186019501011161013257565b601f909101601f19168101906001600160401b03821190821017610a7d57604052565b634e487b7160e01b5f52604160045260245ffd5b90610100828203126101325780601f830112156101325761010060405192610ab98285610a5a565b8391810192831161013257905b828210610ad35750505090565b8135815260209182019101610ac6565b905f905b60088210610af457505050565b6020806001928551815201930191019091610ae7565b905f5160206111cb5f395f51905f528210801590610b94575b61050b57811580610b8c575b610b8657610b535f5160206111cb5f395f51905f5260038185818180090908610efe565b818103610b6257505060011b90565b5f5160206111cb5f395f51905f52809106810306145f1461050b57600190811b1790565b50505f90565b508015610b2f565b505f5160206111cb5f395f51905f52811015610b23565b919093925f5160206111cb5f395f51905f528310801590610d7c575b8015610d65575b8015610d4e575b61050b578082868517171715610d4357908291610ca65f5160206111cb5f395f51905f5280808080888180808f9d5f5160206113eb5f395f51905f528f839290839109099d8e0981848181800909085f51602061152b5f395f51905f52089a09818c8181800909085f5160206110cb5f395f51905f520806810306945f5160206111cb5f395f51905f525f51602061124b5f395f51905f5281610c8081808b80098187800908610efe565b8408095f5160206111cb5f395f51905f52610c9a82611062565b80091415958691610f21565b929080821480610d3a575b15610cd85750505050905f14610cd05760ff60025b169060021b179190565b60ff5f610cc6565b5f5160206111cb5f395f51905f52809106810306149182610d1b575b50501561050b5760019115610d135760ff60025b169060021b17179190565b60ff5f610d08565b5f5160206111cb5f395f51905f52919250819006810306145f80610cf4565b50838314610cb1565b50505090505f905f90565b505f5160206111cb5f395f51905f52811015610bd5565b505f5160206111cb5f395f51905f52821015610bce565b505f5160206111cb5f395f51905f52851015610bc7565b8015610df7578060011c915f5160206111cb5f395f51905f5283101561050b57600180610dd65f5160206111cb5f395f51905f5260038188818180090908610efe565b931614610ddf57565b905f5160206111cb5f395f51905f5280910681030690565b505f905f90565b801580610ef6575b610eea578060021c92825f5160206111cb5f395f51905f528510801590610ed3575b61050b5784815f5160206111cb5f395f51905f5280808080808080805f5160206113eb5f395f51905f5281610e9d9d8d0909998a0981898181800909085f5160206110cb5f395f51905f520806810306936002808a16149509818a8181800909085f51602061152b5f395f51905f5208610f21565b80929160018082961614610eaf575050565b5f5160206111cb5f395f51905f528093945080929550809106810306930681030690565b505f5160206111cb5f395f51905f52811015610e28565b50505f905f905f905f90565b508115610e06565b90610f0882611062565b915f5160206111cb5f395f51905f528380090361050b57565b915f5160206111cb5f395f51905f525f51602061124b5f395f51905f5281610f6693969496610f5882808a8009818a800908610efe565b90611056575b860809610efe565b925f5160206111cb5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206111cb5f395f51905f5260a083015260208260c08160055afa9151911561050b575f5160206111cb5f395f51905f5282600192090361050b575f5160206111cb5f395f51905f52908209925f5160206111cb5f395f51905f528080808780090681030681878009081490811591611037575b5061050b57565b90505f5160206111cb5f395f51905f528084860960020914155f611030565b81809106810306610f5e565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206111cb5f395f51905f5260a083015260208260c08160055afa9151911561050b5756fe2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775209d9ddb850e04ad2295d475e7299c2c0223bdde35b26580f5c7b94612974b8616df4af595e9734756c1706103dbbddd88fdb24aa0a3bdf78e9d1acc46d2be1b11882d8a76d80b2157e9794dc31b3801172df3b189b7be2a51a51720928946d127d466b56c2083d9d45dea7963df61a03b6b5b611000aeb33e0ca414cdb3a75a2bb572ed9c2f2572337bd34fd53a2d93d3118a6c2d7289b678338aa6eb9bb33608ad4e48f57cedc95110e0ff8ae1ff2dfe5e8e3db93336f050bb2e4b462a697e2b0d12d9219ece2357733c64d9db9e298016ba8015403e05627e7053d5dd8db730644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd470fc5d541ec7cb506255cab1e0eda1d5d74ad1c35040c20ea06d84b32a4a3d4011486cd9891d0636bd4784941201eb482dfec40d850609f235ffad76e3483796e25955f902f13ee2422ad3a5be0694762082de6a6a3daf1c4e3317da200e75400183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4275c2591cdee3d365a8e9d97ef28d639b894cd41803c8ffecf1cb927555f24dc07e902613c4255697b06d0274c48c107a04aec8226b5c339ea4daaadbe1d226d2ab899b6c7a84e5ac644ba68129fe2083595544d528aeb72b1e2b50361816a2a2eff4d28074b6f5b8a5491a271dbb28c3de2c2d55fa0b1c82df4903121512cd32cc158464135da1585215bcf31c57c5224fc1cc0d0d1d88c4879f46a7d471b1a30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000116f920fd6ecc50bae126a81c547900d7ee62493a008816efc3c3019ad60b621304167c5e6d72eef4f52451ec2700e3821426e7a8536f9f2a14e88b6a9bcaa1832ea0e3ca710220cdf045c0ff51924985daec2257e5afd64907951eb9d796cb5b08811cca65db2eebb39c952628da76a11c4206ee1b47f0d573170af83b8cedbd03c1430688113b10ac89663832f0fb51795f08190a73b68a68cc8715f9516a491f01dda6271c92d0d7a81526986ce544c0d14752f23e7ab90027ddc3f227387c30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd440a7807babea9b496871f78c245758992b87aa1924f02afc522c6b074283ef3a32d9327ae4adf0e79a06bffa56ad102574780457481bfe9e1a595a1eeb1faf8ac190c5080c2a2138343614bb54918b25ee49a166a3b765550c1ace17d60d8d76b23f7430c87bf3895764d3728de6a16f8bf4f3d3f41a94331b8d9c87495ef2c1204cd79ea57b5ff70d2b8a43b9d02d0601e6f4faf7c3bc970277c560e24ff3c2004d44c7bab55d1b888b63ee15bf48dda6ef781dcfed188c353bd9ea0b66220170904b96d2e510a2caec422d3118ce4df3db20da83975894ac76fa7f91f22298f1a5c03556daeb35576aa76a831a21c5e721e8ce32a7d51843d2d56ea4cee7e6c0628435d77dfccb7a22247fb32c8aa002a1d8daad388783514d603220bb3ac2e2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e52ae0091bc02d8c8dea626ed5e5e26fcf7903fb231489a8f5db7277cbaa8612241facde2fd9956218d4eee71e885437f2c3c42fe34d773178b881339e1ea1daa1a26469706673582212209f32d9426455e79145001b17c6eb6a6c942b8bdfcd71a5f3d799fe9c3b696c0364736f6c634300081c0033",
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
