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
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidInputEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x60808060405234601557611645908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80630ea14a3914610a20578063233ace11146109e6578063449ccd1e1461061657806348306671146101c55763b8e72af614610051575f80fd5b34610193576040366003190112610193576004356001600160401b03811161019357610081903690600401610ae8565b6024356001600160401b038111610193576100a0903690600401610ae8565b61010083036101b65761010081036101a7578101610100828203126101935780601f8301121561019357610100604051926100db8285610b15565b8391810192831161019357905b82821061019757505050303b156101935760405163224e668f60e11b8152610120600482015261012481018390529283919083906101448401375f6101448484010152602482015f905b6008821061017957505050610144815f93601f80199101168101030181305afa801561016e57610160575080f35b61016c91505f90610b15565b005b6040513d5f823e3d90fd5b829350602080916001939451815201930191018492610132565b5f80fd5b81358152602091820191016100e8565b630c0b7e3560e11b5f5260045ffd5b63236bd13760e01b5f5260045ffd5b34610193576101803660031901126101935736608411610193573661018411610193576103006040516101f88282610b15565b81368237610207600435610e18565b610218602493929335604435610e83565b91939290610227606435610e18565b9390926040519660408801965f5160206112705f395f51905f5289528860208101985f5160206111b05f395f51905f528a525f5160206113505f395f51905f5281525f5160206113d05f395f51905f52604060608401925f5160206115b05f395f51905f5284525f5160206113f05f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f5160206112305f395f51905f5285525f5160206114105f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f5160206111705f395f51905f5285525f5160206115d05f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f5160206112905f395f51905f5285525f5160206115905f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f5160206113905f395f51905f5285525f5160206113105f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f5160206111f05f395f51905f5285525f5160206114905f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f5160206114b05f395f51905f5285525f5160206114705f395f51905f5288526101443590818a5287838760608160075afa921016169160808160065afa16945f5160206111905f395f51905f528352526101643580955260608160075afa9210161660408a60808160065afa169851975198156106075760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f5160206114505f395f51905f526101008401525f5160206113705f395f51905f526101208401525f5160206115105f395f51905f526101408401525f5160206112f05f395f51905f526101608401525f5160206115505f395f51905f526101808401525f5160206111d05f395f51905f526101a08401525f5160206114d05f395f51905f526101c08401525f5160206114305f395f51905f526101e08401525f5160206112105f395f51905f526102008401525f5160206113b05f395f51905f526102208401526102408301526102608201525f5160206114f05f395f51905f526102808201525f5160206112d05f395f51905f526102a08201525f5160206112b05f395f51905f526102c08201525f5160206115705f395f51905f526102e08201526040519283916105d38484610b15565b8336843760085afa159081156105fa575b506105eb57005b631ff3747d60e21b5f5260045ffd5b60019150511415816105e4565b63a54f8e2760e01b5f5260045ffd5b3461019357610120366003190112610193576004356001600160401b03811161019357610647903690600401610ae8565b36610124116101935761010061065d9114610b4c565b604051604081015f5160206112705f395f51905f52825260208201905f5160206111b05f395f51905f5282525f5160206113505f395f51905f528152606083015f5160206115b05f395f51905f5281525f5160206113d05f395f51905f526040602435935f5160206113f05f395f51905f52608088019580875284848460608160075afa911016838960808160065afa165f5160206112305f395f51905f5283525f5160206114105f395f51905f52865260016044359182895286868660608160075afa9310161616838960808160065afa165f5160206111705f395f51905f5283525f5160206115d05f395f51905f5286526064359081885285858560608160075afa92101616838960808160065afa165f5160206112905f395f51905f5283525f5160206115905f395f51905f5286526084359081885285858560608160075afa92101616838960808160065afa165f5160206113905f395f51905f5283525f5160206113105f395f51905f52865260a4359081885285858560608160075afa92101616838960808160065afa165f5160206111f05f395f51905f5283525f5160206114905f395f51905f52865260c4359081885285858560608160075afa92101616838960808160065afa165f5160206114b05f395f51905f5283525f5160206114705f395f51905f52865260e4359081885285858560608160075afa92101616838960808160065afa16945f5160206111905f395f51905f528352526101043580955260608160075afa9210161660408360808160065afa16915190519115610607576101006040519384375f5160206114505f395f51905f526101008401525f5160206113705f395f51905f526101208401525f5160206115105f395f51905f526101408401525f5160206112f05f395f51905f526101608401525f5160206115505f395f51905f526101808401525f5160206111d05f395f51905f526101a08401525f5160206114d05f395f51905f526101c08401525f5160206114305f395f51905f526101e08401525f5160206112105f395f51905f526102008401525f5160206113b05f395f51905f526102208401526102408301526102608201525f5160206114f05f395f51905f526102808201525f5160206112d05f395f51905f526102a08201525f5160206112b05f395f51905f526102c08201525f5160206115705f395f51905f526102e08201526020816103008160085afa905116156105eb57005b34610193575f3660031901126101935760206040517fbff0af27903c29e2e759e5b9653430e743cd8d56d24a2a4556fad311ec650a6a8152f35b34610193576020366003190112610193576004356001600160401b03811161019357610a50903690600401610ae8565b610ab9608092610a7461010060405194610a6a8787610b15565b8636873714610b4c565b610a8360208201358235610b8f565b8352610aa08482013560a083013560408401356060850135610c30565b6020850152604084015260c060e0820135910135610b8f565b6060820152604051905f825b60048210610ad257505050f35b6020806001928551815201930191019091610ac5565b9181601f84011215610193578235916001600160401b038311610193576020838186019501011161019357565b601f909101601f19168101906001600160401b03821190821017610b3857604052565b634e487b7160e01b5f52604160045260245ffd5b15610b5357565b60405162461bcd60e51b81526020600482015260146024820152730d2dcecc2d8d2c840e0e4dedecc40d8cadccee8d60631b6044820152606490fd5b905f5160206112505f395f51905f528210801590610c19575b6105eb57811580610c11575b610c0b57610bd85f5160206112505f395f51905f5260038185818180090908610f83565b818103610be757505060011b90565b5f5160206112505f395f51905f52809106810306145f146105eb57600190811b1790565b50505f90565b508015610bb4565b505f5160206112505f395f51905f52811015610ba8565b919093925f5160206112505f395f51905f528310801590610e01575b8015610dea575b8015610dd3575b6105eb578082868517171715610dc857908291610d2b5f5160206112505f395f51905f5280808080888180808f9d5f5160206115305f395f51905f528f839290839109099d8e0981848181800909085f5160206115f05f395f51905f52089a09818c8181800909085f5160206111505f395f51905f520806810306945f5160206112505f395f51905f525f5160206113305f395f51905f5281610d0581808b80098187800908610f83565b8408095f5160206112505f395f51905f52610d1f826110e7565b80091415958691610fa6565b929080821480610dbf575b15610d5d5750505050905f14610d555760ff60025b169060021b179190565b60ff5f610d4b565b5f5160206112505f395f51905f52809106810306149182610da0575b5050156105eb5760019115610d985760ff60025b169060021b17179190565b60ff5f610d8d565b5f5160206112505f395f51905f52919250819006810306145f80610d79565b50838314610d36565b50505090505f905f90565b505f5160206112505f395f51905f52811015610c5a565b505f5160206112505f395f51905f52821015610c53565b505f5160206112505f395f51905f52851015610c4c565b8015610e7c578060011c915f5160206112505f395f51905f528310156105eb57600180610e5b5f5160206112505f395f51905f5260038188818180090908610f83565b931614610e6457565b905f5160206112505f395f51905f5280910681030690565b505f905f90565b801580610f7b575b610f6f578060021c92825f5160206112505f395f51905f528510801590610f58575b6105eb5784815f5160206112505f395f51905f5280808080808080805f5160206115305f395f51905f5281610f229d8d0909998a0981898181800909085f5160206111505f395f51905f520806810306936002808a16149509818a8181800909085f5160206115f05f395f51905f5208610fa6565b80929160018082961614610f34575050565b5f5160206112505f395f51905f528093945080929550809106810306930681030690565b505f5160206112505f395f51905f52811015610ead565b50505f905f905f905f90565b508115610e8b565b90610f8d826110e7565b915f5160206112505f395f51905f52838009036105eb57565b915f5160206112505f395f51905f525f5160206113305f395f51905f5281610feb93969496610fdd82808a8009818a800908610f83565b906110db575b860809610f83565b925f5160206112505f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206112505f395f51905f5260a083015260208260c08160055afa915191156105eb575f5160206112505f395f51905f528260019209036105eb575f5160206112505f395f51905f52908209925f5160206112505f395f51905f5280808087800906810306818780090814908115916110bc575b506105eb57565b90505f5160206112505f395f51905f528084860960020914155f6110b5565b81809106810306610fe3565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206112505f395f51905f5260a083015260208260c08160055afa915191156105eb5756fe2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7751990f0c237a2c2d4a993cf9c8360a57f019fb3dc7579755c1fe813e406f721bf07f1978f5ea453b7fe60ff545fb7825409c135e6d311de1ff18c6b3193cf241421ec27911dc8f4c02590c6db844b275be5ef84b9d453747e70e7183e82a38b3a1b8797d436a8c13dd48d5581ade0fdb8d7daf888ab92b9035cf4f2df2641b046181972ec10c609f370a3f2f57d281c1573c556f87a6ce049833faf6aa2ab9f6f1ff70dcb6210c23ca796e5418e082fb294c0317518d1932022c927fdbd0a601513cac61690ac7de15d09739b7fa45ca83a74305787d22d31309d0a3c819dad2930644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4720cb32efe7251c7a3d6cd56f913d2aeb9500f2c93e21ab4646a9b64560c7223f22e53f60f0ee6bd6eabdb6bfec4cbdcfe30d79fd5817cdb9dbfa7a2b9cba4caf1e916c65f9abdc0af2bfde301fa0d22a539a1b71d9288a3abcf3680c1a80da2b2f936be819bd214baf216e309ae57c93f2e74eac810067811582aee098ec6f4a124448df6b3bb0bfc15e79a19c70c647806a243a8ad8cf7a180ad085a284e36d0cc79be8601bc375698ff1a756015e0b43214524173a556b19d004c7ace2ea03183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea423734184db1b65df226d95e3c6de3a821778c189e2fdb200a4eadf058ac70ea81c0ffc9942d93204ff2e4e0b1c9b2ebdd31cb57baea688253823ff3f9ddeabeb1a2e92d4059c2174535b99bf7d5ef67c1789fafbf53191df3d944e856dedf38d057ee6a16f6186403cddfe474c985ce8abb7339f7330ae0689b220ee941e25fb30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000010a3f32e84c80d9224c41d01040d634fc23a335c3e7f6f1f000e91d1aea42604405fc9fb814b4e6e19f91b7623de94b118c063ceb8717f88d7f675c2cfc78153025d6e5dfa37fab8aea074a1b8aec887e42ab9756a99d5249f411f0b26dbe021a0b61955e8c41d899556737574d20e8788562764222ac35f3b6db3b1b61cb8601056dbf991c12ecea2d0791187b9958002e5366ca0473f0bd13bfecb73f6639c50058cf4dbd3e55fefd44957d1db27bd6d776945c312400d7d3b517a24e866cb10015412eb1d24b50fef8123f42e53d4fbef2478a65e52965fad8517e98417acc2bdeeeab45e3f93d5528d125884dfe643b664f9d7cfa892943ffea47ff67f9aa1fecb3c84051319495bb64280539b734b650ff5bbc89b1b2ad94aca284d2071026d3550d9fe24bc88015706c0ccef3113d10ebfaa63a7511172bedc9df86a1f130644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd442ee49555fab7b176c0052bf27a8067716ca0e8bc03cabf9490c66f2825a51ac101e6fac558decbfc635b7120890078f9ae287b14fb302503dd38d785f6be51002b749fb6bbf56def728e54cee43ed68ebc6e938f72dd97999d311a931485869d10d22ec007e1fc0b204dd5b44c0b4eb48d8d034c91a28eabb7c412bd2e1e37891481132c863ffc4a50389b1d2b376ede5ca3443aca363ebf7c40f46a4021d49f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5a2646970667358221220c0821718a29310aa066c9be7c7be651528defba8fd365b1492dc8a0b3851a6e464736f6c634300081c0033",
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

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed)
func (_ContributionVerifier *ContributionVerifierCaller) CompressProof(opts *bind.CallOpts, proof []byte) ([4]*big.Int, error) {
	var out []interface{}
	err := _ContributionVerifier.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed)
func (_ContributionVerifier *ContributionVerifierSession) CompressProof(proof []byte) ([4]*big.Int, error) {
	return _ContributionVerifier.Contract.CompressProof(&_ContributionVerifier.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed)
func (_ContributionVerifier *ContributionVerifierCallerSession) CompressProof(proof []byte) ([4]*big.Int, error) {
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

// VerifyProof is a free data retrieval call binding the contract method 0x449ccd1e.
//
// Solidity: function verifyProof(bytes proof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierCaller) VerifyProof(opts *bind.CallOpts, proof []byte, input [8]*big.Int) error {
	var out []interface{}
	err := _ContributionVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x449ccd1e.
//
// Solidity: function verifyProof(bytes proof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierSession) VerifyProof(proof []byte, input [8]*big.Int) error {
	return _ContributionVerifier.Contract.VerifyProof(&_ContributionVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x449ccd1e.
//
// Solidity: function verifyProof(bytes proof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierCallerSession) VerifyProof(proof []byte, input [8]*big.Int) error {
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
