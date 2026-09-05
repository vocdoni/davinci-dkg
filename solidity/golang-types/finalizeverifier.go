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

// FinalizeVerifierMetaData contains all meta data concerning the FinalizeVerifier contract.
var FinalizeVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[7]\",\"internalType\":\"uint256[7]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"uint256[7]\",\"internalType\":\"uint256[7]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x6080806040523460155761157a908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80630ea14a3914610995578063233ace111461095b578063598da1d1146105ad578063b8e72af6146104695763f6d3293a14610051575f80fd5b3461046657610160366003190112610466573660841161046657366101641161046657604051906103006100858184610a8a565b80368437610094600435610d8d565b6100a5602495929535604435610df8565b919392906100b4606435610d8d565b9390926040519660408801965f5160206111255f395f51905f5289528860208101985f5160206111e55f395f51905f528a525f5160206114455f395f51905f5281525f5160206113655f395f51905f52604060608401925f5160206115055f395f51905f5284525f5160206111a55f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f5160206113255f395f51905f5285525f5160206111455f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f5160206114e55f395f51905f5285525f5160206114255f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f5160206112455f395f51905f5285525f5160206112855f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f5160206112a55f395f51905f5285525f5160206113055f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f5160206113e55f395f51905f5285525f5160206113a55f395f51905f5288526101243590818a5287838760608160075afa921016169160808160065afa16945f5160206114055f395f51905f528352526101443580955260608160075afa9210161660408a60808160065afa169851975198156104575760209a9b9c8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f5160206111655f395f51905f526101008401525f5160206112655f395f51905f526101208401525f5160206111855f395f51905f526101408401525f5160206113855f395f51905f526101608401525f5160206112c55f395f51905f526101808401525f5160206114855f395f51905f526101a08401525f5160206113c55f395f51905f526101c08401525f5160206114a55f395f51905f526101e08401525f5160206110e55f395f51905f526102008401525f5160206113455f395f51905f526102208401526102408301526102608201525f5160206111c55f395f51905f526102808201525f5160206114c55f395f51905f526102a08201525f5160206111055f395f51905f526102c08201525f5160206112255f395f51905f526102e08201526040519283916104228484610a8a565b8336843760085afa1590811561044a575b5061043b5780f35b631ff3747d60e21b8152600490fd5b600191505114155f610433565b63a54f8e2760e01b8c5260048cfd5b80fd5b5034610599576040366003190112610599576004356001600160401b0381116105995761049a903690600401610a5d565b6024356001600160401b038111610599576104b9903690600401610a5d565b810160e0828203126105995780601f8301121561059957604051916104df60e084610a8a565b829060e0810192831161059957905b82821061059d57505050303b156105995760405163598da1d160e01b8152610100600482015261010481018390529283919083906101248401375f6101248484010152602482015f905b6007821061057f57505050610124815f93601f80199101168101030181305afa801561057457610566575080f35b61057291505f90610a8a565b005b6040513d5f823e3d90fd5b829350602080916001939451815201930191018492610538565b5f80fd5b81358152602091820191016104ee565b3461059957610100366003190112610599576004356001600160401b038111610599576105de903690600401610a5d565b3661010411610599576101006105f49114610ac1565b604051604081015f5160206111255f395f51905f52825260208201905f5160206111e55f395f51905f5282525f5160206114455f395f51905f528152606083015f5160206115055f395f51905f5281525f5160206113655f395f51905f526040602435935f5160206111a55f395f51905f52608088019580875284848460608160075afa911016838960808160065afa165f5160206113255f395f51905f5283525f5160206111455f395f51905f52865260016044359182895286868660608160075afa9310161616838960808160065afa165f5160206114e55f395f51905f5283525f5160206114255f395f51905f5286526064359081885285858560608160075afa92101616838960808160065afa165f5160206112455f395f51905f5283525f5160206112855f395f51905f5286526084359081885285858560608160075afa92101616838960808160065afa165f5160206112a55f395f51905f5283525f5160206113055f395f51905f52865260a4359081885285858560608160075afa92101616838960808160065afa165f5160206113e55f395f51905f5283525f5160206113a55f395f51905f52865260c4359081885285858560608160075afa92101616838960808160065afa16945f5160206114055f395f51905f5283525260e43580955260608160075afa9210161660408360808160065afa1691519051911561094c576101006040519384375f5160206111655f395f51905f526101008401525f5160206112655f395f51905f526101208401525f5160206111855f395f51905f526101408401525f5160206113855f395f51905f526101608401525f5160206112c55f395f51905f526101808401525f5160206114855f395f51905f526101a08401525f5160206113c55f395f51905f526101c08401525f5160206114a55f395f51905f526101e08401525f5160206110e55f395f51905f526102008401525f5160206113455f395f51905f526102208401526102408301526102608201525f5160206111c55f395f51905f526102808201525f5160206114c55f395f51905f526102a08201525f5160206111055f395f51905f526102c08201525f5160206112255f395f51905f526102e08201526020816103008160085afa9051161561093d57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b34610599575f3660031901126105995760206040517f786fddb4f3190c1807ceb80865f0575c7f2f5ecf4ac1db52c890a37d280cea968152f35b34610599576020366003190112610599576004356001600160401b038111610599576109c5903690600401610a5d565b610a2e6080926109e9610100604051946109df8787610a8a565b8636873714610ac1565b6109f860208201358235610b04565b8352610a158482013560a083013560408401356060850135610ba5565b6020850152604084015260c060e0820135910135610b04565b6060820152604051905f825b60048210610a4757505050f35b6020806001928551815201930191019091610a3a565b9181601f84011215610599578235916001600160401b038311610599576020838186019501011161059957565b601f909101601f19168101906001600160401b03821190821017610aad57604052565b634e487b7160e01b5f52604160045260245ffd5b15610ac857565b60405162461bcd60e51b81526020600482015260146024820152730d2dcecc2d8d2c840e0e4dedecc40d8cadccee8d60631b6044820152606490fd5b905f5160206112055f395f51905f528210801590610b8e575b61093d57811580610b86575b610b8057610b4d5f5160206112055f395f51905f5260038185818180090908610ef8565b818103610b5c57505060011b90565b5f5160206112055f395f51905f52809106810306145f1461093d57600190811b1790565b50505f90565b508015610b29565b505f5160206112055f395f51905f52811015610b1d565b919093925f5160206112055f395f51905f528310801590610d76575b8015610d5f575b8015610d48575b61093d578082868517171715610d3d57908291610ca05f5160206112055f395f51905f5280808080888180808f9d5f5160206114655f395f51905f528f839290839109099d8e0981848181800909085f5160206115255f395f51905f52089a09818c8181800909085f5160206110c55f395f51905f520806810306945f5160206112055f395f51905f525f5160206112e55f395f51905f5281610c7a81808b80098187800908610ef8565b8408095f5160206112055f395f51905f52610c948261105c565b80091415958691610f1b565b929080821480610d34575b15610cd25750505050905f14610cca5760ff60025b169060021b179190565b60ff5f610cc0565b5f5160206112055f395f51905f52809106810306149182610d15575b50501561093d5760019115610d0d5760ff60025b169060021b17179190565b60ff5f610d02565b5f5160206112055f395f51905f52919250819006810306145f80610cee565b50838314610cab565b50505090505f905f90565b505f5160206112055f395f51905f52811015610bcf565b505f5160206112055f395f51905f52821015610bc8565b505f5160206112055f395f51905f52851015610bc1565b8015610df1578060011c915f5160206112055f395f51905f5283101561093d57600180610dd05f5160206112055f395f51905f5260038188818180090908610ef8565b931614610dd957565b905f5160206112055f395f51905f5280910681030690565b505f905f90565b801580610ef0575b610ee4578060021c92825f5160206112055f395f51905f528510801590610ecd575b61093d5784815f5160206112055f395f51905f5280808080808080805f5160206114655f395f51905f5281610e979d8d0909998a0981898181800909085f5160206110c55f395f51905f520806810306936002808a16149509818a8181800909085f5160206115255f395f51905f5208610f1b565b80929160018082961614610ea9575050565b5f5160206112055f395f51905f528093945080929550809106810306930681030690565b505f5160206112055f395f51905f52811015610e22565b50505f905f905f905f90565b508115610e00565b90610f028261105c565b915f5160206112055f395f51905f528380090361093d57565b915f5160206112055f395f51905f525f5160206112e55f395f51905f5281610f6093969496610f5282808a8009818a800908610ef8565b90611050575b860809610ef8565b925f5160206112055f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206112055f395f51905f5260a083015260208260c08160055afa9151911561093d575f5160206112055f395f51905f5282600192090361093d575f5160206112055f395f51905f52908209925f5160206112055f395f51905f528080808780090681030681878009081490811591611031575b5061093d57565b90505f5160206112055f395f51905f528084860960020914155f61102a565b81809106810306610f58565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206112055f395f51905f5260a083015260208260c08160055afa9151911561093d5756fe2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77513d6a48130b6f6eebe2f002cb232b7344ea6ab525fcd7062b6973c6aca6df44325bcd5b8ba3a938bc71194c61fb13914921a5e4a7de0b3077e0f5a1f6ec6f804200845e5e840c92816af94fb6009c9c457c62e81f68564155a5f1aa6519eba0910ff3f00a07cf3c9e94a797fd33fd4041afcf34e12c64271bedc27d8ed0189131a31cbd20dcb0526956f314a614b9fefae561dfafb7234cdf8122412c06a3b2c0f534cfa2d96a5b7ae5d71fe0f4654f035ba5fb5848d8eac62aaa3db3406b4861b2e8c1855be2058550decea687d0dcd4c566f5d451fe358bb62f7d9df89800f01c4dce70a453b3777894e6b9de9cd063bb0b19c198cdd96cbfef703711ed8b00971c4d99a72a6524f6de273fc54c14ef3840605cd0367af0f79c982adf77bf730644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd471d8943440dac0c63a647842e801cffb8d87971caabac7fa0ef66d4fe1bacffc21cea982c415a7f50c59ad5f0b451c31b7a3c63dd37515120f76406362b767c2d1c82e6b766a7a8c680e208a3e714e937d54bf0d429bb0d1bdcd0567c44a892a601d290caf8c4b279bf7757f920aed7b1dd33a55ec161fb80811e2ac7c62ebebe1277303d71b5c9b4fca552e908a5b5c22eed27704e104ee0e260754893a699471aad9bc06144a4f4e7452a374460a96f8890ece0f1f5efbb3d3ba24e8fbaa7f6183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea40f62265e46b2d276635f0c777de0dcb9c8775b9bf2fc89a3b706f59b7f11f679201ba11bae833dc84282f715f8b09a1e57421cea853ca6975167b2f8a91db39f07b99517fa771a4bfd81fc0a220a96ea06992647cae3e64215eab577e4f6ed3e30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001052da7dfe73df771731633a0a4ba358abeab8772b889505b4964b35024d83b860610449544ab66b9757ed14b91bb82a6a2b62d7be971245ff9bf4f7a83758b3d240c227c7037f111ea8f5ce29c02d6db67e058a67bd0a8a2293e12f056c0840221604f94e23ff576bc06f6442c236234e4584d6142b5f6b5e059974eea0bb5d42a85d48c3e67f52641bff5f67b20218c1c46bd9e0383405273ee436f057f5c0624aba406c1d306edf3c013b110586c1bdc1ff40d6feaecd7feef6243086942e207fc26a9edbf902f0ede00d2a4c25ac02f984daa53db07391cff5469943b524c30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44154a0a879d7776bb21266f7cc1f60d3fa11713c7bd634559afe8c8a4b2dd4bdc280f047a7243e873b3da1eabd9cbab4c3985263d4ccabd0929d308ce1213a7a628d69b703a73ee55845adf5d852a6e5b8701ddb452f72b25ae61a898ece379511f431b089fd5914557c59d80d0a1d74c9e01f4a0d3293eddc41f373d24fb10d0015924d864fa69eea46af77725e0a0d73344bda161a467a02c618a69c6d3f64f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5a26469706673582212208836193ade1742c4ff4ada5c85dedbdfbe3cba3931f687f66f070e48964de20a64736f6c634300081c0033",
}

// FinalizeVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use FinalizeVerifierMetaData.ABI instead.
var FinalizeVerifierABI = FinalizeVerifierMetaData.ABI

// FinalizeVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FinalizeVerifierMetaData.Bin instead.
var FinalizeVerifierBin = FinalizeVerifierMetaData.Bin

// DeployFinalizeVerifier deploys a new Ethereum contract, binding an instance of FinalizeVerifier to it.
func DeployFinalizeVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FinalizeVerifier, error) {
	parsed, err := FinalizeVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FinalizeVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FinalizeVerifier{FinalizeVerifierCaller: FinalizeVerifierCaller{contract: contract}, FinalizeVerifierTransactor: FinalizeVerifierTransactor{contract: contract}, FinalizeVerifierFilterer: FinalizeVerifierFilterer{contract: contract}}, nil
}

// FinalizeVerifier is an auto generated Go binding around an Ethereum contract.
type FinalizeVerifier struct {
	FinalizeVerifierCaller     // Read-only binding to the contract
	FinalizeVerifierTransactor // Write-only binding to the contract
	FinalizeVerifierFilterer   // Log filterer for contract events
}

// FinalizeVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type FinalizeVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizeVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FinalizeVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizeVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FinalizeVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizeVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FinalizeVerifierSession struct {
	Contract     *FinalizeVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FinalizeVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FinalizeVerifierCallerSession struct {
	Contract *FinalizeVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// FinalizeVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FinalizeVerifierTransactorSession struct {
	Contract     *FinalizeVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// FinalizeVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type FinalizeVerifierRaw struct {
	Contract *FinalizeVerifier // Generic contract binding to access the raw methods on
}

// FinalizeVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FinalizeVerifierCallerRaw struct {
	Contract *FinalizeVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// FinalizeVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FinalizeVerifierTransactorRaw struct {
	Contract *FinalizeVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFinalizeVerifier creates a new instance of FinalizeVerifier, bound to a specific deployed contract.
func NewFinalizeVerifier(address common.Address, backend bind.ContractBackend) (*FinalizeVerifier, error) {
	contract, err := bindFinalizeVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FinalizeVerifier{FinalizeVerifierCaller: FinalizeVerifierCaller{contract: contract}, FinalizeVerifierTransactor: FinalizeVerifierTransactor{contract: contract}, FinalizeVerifierFilterer: FinalizeVerifierFilterer{contract: contract}}, nil
}

// NewFinalizeVerifierCaller creates a new read-only instance of FinalizeVerifier, bound to a specific deployed contract.
func NewFinalizeVerifierCaller(address common.Address, caller bind.ContractCaller) (*FinalizeVerifierCaller, error) {
	contract, err := bindFinalizeVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FinalizeVerifierCaller{contract: contract}, nil
}

// NewFinalizeVerifierTransactor creates a new write-only instance of FinalizeVerifier, bound to a specific deployed contract.
func NewFinalizeVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*FinalizeVerifierTransactor, error) {
	contract, err := bindFinalizeVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FinalizeVerifierTransactor{contract: contract}, nil
}

// NewFinalizeVerifierFilterer creates a new log filterer instance of FinalizeVerifier, bound to a specific deployed contract.
func NewFinalizeVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*FinalizeVerifierFilterer, error) {
	contract, err := bindFinalizeVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FinalizeVerifierFilterer{contract: contract}, nil
}

// bindFinalizeVerifier binds a generic wrapper to an already deployed contract.
func bindFinalizeVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FinalizeVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalizeVerifier *FinalizeVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FinalizeVerifier.Contract.FinalizeVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalizeVerifier *FinalizeVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalizeVerifier.Contract.FinalizeVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalizeVerifier *FinalizeVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalizeVerifier.Contract.FinalizeVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalizeVerifier *FinalizeVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FinalizeVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalizeVerifier *FinalizeVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalizeVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalizeVerifier *FinalizeVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalizeVerifier.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed)
func (_FinalizeVerifier *FinalizeVerifierCaller) CompressProof(opts *bind.CallOpts, proof []byte) ([4]*big.Int, error) {
	var out []interface{}
	err := _FinalizeVerifier.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed)
func (_FinalizeVerifier *FinalizeVerifierSession) CompressProof(proof []byte) ([4]*big.Int, error) {
	return _FinalizeVerifier.Contract.CompressProof(&_FinalizeVerifier.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed)
func (_FinalizeVerifier *FinalizeVerifierCallerSession) CompressProof(proof []byte) ([4]*big.Int, error) {
	return _FinalizeVerifier.Contract.CompressProof(&_FinalizeVerifier.CallOpts, proof)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_FinalizeVerifier *FinalizeVerifierCaller) ProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FinalizeVerifier.contract.Call(opts, &out, "provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_FinalizeVerifier *FinalizeVerifierSession) ProvingKeyHash() ([32]byte, error) {
	return _FinalizeVerifier.Contract.ProvingKeyHash(&_FinalizeVerifier.CallOpts)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_FinalizeVerifier *FinalizeVerifierCallerSession) ProvingKeyHash() ([32]byte, error) {
	return _FinalizeVerifier.Contract.ProvingKeyHash(&_FinalizeVerifier.CallOpts)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xf6d3293a.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[7] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [7]*big.Int) error {
	var out []interface{}
	err := _FinalizeVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xf6d3293a.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[7] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [7]*big.Int) error {
	return _FinalizeVerifier.Contract.VerifyCompressedProof(&_FinalizeVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xf6d3293a.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[7] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [7]*big.Int) error {
	return _FinalizeVerifier.Contract.VerifyCompressedProof(&_FinalizeVerifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x598da1d1.
//
// Solidity: function verifyProof(bytes proof, uint256[7] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCaller) VerifyProof(opts *bind.CallOpts, proof []byte, input [7]*big.Int) error {
	var out []interface{}
	err := _FinalizeVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x598da1d1.
//
// Solidity: function verifyProof(bytes proof, uint256[7] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierSession) VerifyProof(proof []byte, input [7]*big.Int) error {
	return _FinalizeVerifier.Contract.VerifyProof(&_FinalizeVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x598da1d1.
//
// Solidity: function verifyProof(bytes proof, uint256[7] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCallerSession) VerifyProof(proof []byte, input [7]*big.Int) error {
	return _FinalizeVerifier.Contract.VerifyProof(&_FinalizeVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCaller) VerifyProof0(opts *bind.CallOpts, proof []byte, input []byte) error {
	var out []interface{}
	err := _FinalizeVerifier.contract.Call(opts, &out, "verifyProof0", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_FinalizeVerifier *FinalizeVerifierSession) VerifyProof0(proof []byte, input []byte) error {
	return _FinalizeVerifier.Contract.VerifyProof0(&_FinalizeVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCallerSession) VerifyProof0(proof []byte, input []byte) error {
	return _FinalizeVerifier.Contract.VerifyProof0(&_FinalizeVerifier.CallOpts, proof, input)
}
