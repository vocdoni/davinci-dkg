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
	Bin: "0x60808060405234601557611645908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80630ea14a3914610a20578063233ace11146109e6578063449ccd1e1461061657806348306671146101c55763b8e72af614610051575f80fd5b34610193576040366003190112610193576004356001600160401b03811161019357610081903690600401610ae8565b6024356001600160401b038111610193576100a0903690600401610ae8565b61010083036101b65761010081036101a7578101610100828203126101935780601f8301121561019357610100604051926100db8285610b15565b8391810192831161019357905b82821061019757505050303b156101935760405163224e668f60e11b8152610120600482015261012481018390529283919083906101448401375f6101448484010152602482015f905b6008821061017957505050610144815f93601f80199101168101030181305afa801561016e57610160575080f35b61016c91505f90610b15565b005b6040513d5f823e3d90fd5b829350602080916001939451815201930191018492610132565b5f80fd5b81358152602091820191016100e8565b630c0b7e3560e11b5f5260045ffd5b63236bd13760e01b5f5260045ffd5b34610193576101803660031901126101935736608411610193573661018411610193576103006040516101f88282610b15565b81368237610207600435610e18565b610218602493929335604435610e83565b91939290610227606435610e18565b9390926040519660408801965f5160206115105f395f51905f5289528860208101985f5160206112105f395f51905f528a525f5160206114705f395f51905f5281525f5160206113d05f395f51905f52604060608401925f5160206113705f395f51905f5284525f5160206114b05f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f5160206113105f395f51905f5285525f5160206115d05f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f5160206112d05f395f51905f5285525f5160206112f05f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f5160206113f05f395f51905f5285525f5160206114d05f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f5160206114505f395f51905f5285525f5160206113b05f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f5160206114905f395f51905f5285525f5160206112b05f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f5160206111d05f395f51905f5285525f5160206115705f395f51905f5288526101443590818a5287838760608160075afa921016169160808160065afa16945f5160206115b05f395f51905f528352526101643580955260608160075afa9210161660408a60808160065afa169851975198156106075760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f5160206111b05f395f51905f526101008401525f5160206113905f395f51905f526101208401525f5160206112705f395f51905f526101408401525f5160206115305f395f51905f526101608401525f5160206114105f395f51905f526101808401525f5160206111f05f395f51905f526101a08401525f5160206112305f395f51905f526101c08401525f5160206113305f395f51905f526101e08401525f5160206112505f395f51905f526102008401525f5160206111905f395f51905f526102208401526102408301526102608201525f5160206114305f395f51905f526102808201525f5160206115505f395f51905f526102a08201525f5160206115905f395f51905f526102c08201525f5160206111705f395f51905f526102e08201526040519283916105d38484610b15565b8336843760085afa159081156105fa575b506105eb57005b631ff3747d60e21b5f5260045ffd5b60019150511415816105e4565b63a54f8e2760e01b5f5260045ffd5b3461019357610120366003190112610193576004356001600160401b03811161019357610647903690600401610ae8565b36610124116101935761010061065d9114610b4c565b604051604081015f5160206115105f395f51905f52825260208201905f5160206112105f395f51905f5282525f5160206114705f395f51905f528152606083015f5160206113705f395f51905f5281525f5160206113d05f395f51905f526040602435935f5160206114b05f395f51905f52608088019580875284848460608160075afa911016838960808160065afa165f5160206113105f395f51905f5283525f5160206115d05f395f51905f52865260016044359182895286868660608160075afa9310161616838960808160065afa165f5160206112d05f395f51905f5283525f5160206112f05f395f51905f5286526064359081885285858560608160075afa92101616838960808160065afa165f5160206113f05f395f51905f5283525f5160206114d05f395f51905f5286526084359081885285858560608160075afa92101616838960808160065afa165f5160206114505f395f51905f5283525f5160206113b05f395f51905f52865260a4359081885285858560608160075afa92101616838960808160065afa165f5160206114905f395f51905f5283525f5160206112b05f395f51905f52865260c4359081885285858560608160075afa92101616838960808160065afa165f5160206111d05f395f51905f5283525f5160206115705f395f51905f52865260e4359081885285858560608160075afa92101616838960808160065afa16945f5160206115b05f395f51905f528352526101043580955260608160075afa9210161660408360808160065afa16915190519115610607576101006040519384375f5160206111b05f395f51905f526101008401525f5160206113905f395f51905f526101208401525f5160206112705f395f51905f526101408401525f5160206115305f395f51905f526101608401525f5160206114105f395f51905f526101808401525f5160206111f05f395f51905f526101a08401525f5160206112305f395f51905f526101c08401525f5160206113305f395f51905f526101e08401525f5160206112505f395f51905f526102008401525f5160206111905f395f51905f526102208401526102408301526102608201525f5160206114305f395f51905f526102808201525f5160206115505f395f51905f526102a08201525f5160206115905f395f51905f526102c08201525f5160206111705f395f51905f526102e08201526020816103008160085afa905116156105eb57005b34610193575f3660031901126101935760206040517f3ff7007dbb761e27713637a497ef763e3a159564aeac58baaccbc3610ea0a6cf8152f35b34610193576020366003190112610193576004356001600160401b03811161019357610a50903690600401610ae8565b610ab9608092610a7461010060405194610a6a8787610b15565b8636873714610b4c565b610a8360208201358235610b8f565b8352610aa08482013560a083013560408401356060850135610c30565b6020850152604084015260c060e0820135910135610b8f565b6060820152604051905f825b60048210610ad257505050f35b6020806001928551815201930191019091610ac5565b9181601f84011215610193578235916001600160401b038311610193576020838186019501011161019357565b601f909101601f19168101906001600160401b03821190821017610b3857604052565b634e487b7160e01b5f52604160045260245ffd5b15610b5357565b60405162461bcd60e51b81526020600482015260146024820152730d2dcecc2d8d2c840e0e4dedecc40d8cadccee8d60631b6044820152606490fd5b905f5160206112905f395f51905f528210801590610c19575b6105eb57811580610c11575b610c0b57610bd85f5160206112905f395f51905f5260038185818180090908610f83565b818103610be757505060011b90565b5f5160206112905f395f51905f52809106810306145f146105eb57600190811b1790565b50505f90565b508015610bb4565b505f5160206112905f395f51905f52811015610ba8565b919093925f5160206112905f395f51905f528310801590610e01575b8015610dea575b8015610dd3575b6105eb578082868517171715610dc857908291610d2b5f5160206112905f395f51905f5280808080888180808f9d5f5160206114f05f395f51905f528f839290839109099d8e0981848181800909085f5160206115f05f395f51905f52089a09818c8181800909085f5160206111505f395f51905f520806810306945f5160206112905f395f51905f525f5160206113505f395f51905f5281610d0581808b80098187800908610f83565b8408095f5160206112905f395f51905f52610d1f826110e7565b80091415958691610fa6565b929080821480610dbf575b15610d5d5750505050905f14610d555760ff60025b169060021b179190565b60ff5f610d4b565b5f5160206112905f395f51905f52809106810306149182610da0575b5050156105eb5760019115610d985760ff60025b169060021b17179190565b60ff5f610d8d565b5f5160206112905f395f51905f52919250819006810306145f80610d79565b50838314610d36565b50505090505f905f90565b505f5160206112905f395f51905f52811015610c5a565b505f5160206112905f395f51905f52821015610c53565b505f5160206112905f395f51905f52851015610c4c565b8015610e7c578060011c915f5160206112905f395f51905f528310156105eb57600180610e5b5f5160206112905f395f51905f5260038188818180090908610f83565b931614610e6457565b905f5160206112905f395f51905f5280910681030690565b505f905f90565b801580610f7b575b610f6f578060021c92825f5160206112905f395f51905f528510801590610f58575b6105eb5784815f5160206112905f395f51905f5280808080808080805f5160206114f05f395f51905f5281610f229d8d0909998a0981898181800909085f5160206111505f395f51905f520806810306936002808a16149509818a8181800909085f5160206115f05f395f51905f5208610fa6565b80929160018082961614610f34575050565b5f5160206112905f395f51905f528093945080929550809106810306930681030690565b505f5160206112905f395f51905f52811015610ead565b50505f905f905f905f90565b508115610e8b565b90610f8d826110e7565b915f5160206112905f395f51905f52838009036105eb57565b915f5160206112905f395f51905f525f5160206113505f395f51905f5281610feb93969496610fdd82808a8009818a800908610f83565b906110db575b860809610f83565b925f5160206112905f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206112905f395f51905f5260a083015260208260c08160055afa915191156105eb575f5160206112905f395f51905f528260019209036105eb575f5160206112905f395f51905f52908209925f5160206112905f395f51905f5280808087800906810306818780090814908115916110bc575b506105eb57565b90505f5160206112905f395f51905f528084860960020914155f6110b5565b81809106810306610fe3565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206112905f395f51905f5260a083015260208260c08160055afa915191156105eb5756fe2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77519b5b51380ff837b289cbd0733edaa57ab07ce7d9d4a9357d99bd74dc02ddf9005ff84922ec3e02899e1dee571c07fa57b6294fec33e930843117cf22131b3512f0eae0a6d43240d5b66e49f0d5552426a1dda06ef1d71ab08aa4c6bae16912a2cf1fcd65ee147f720476d64d8594cf9eb5e1d20809299f701789bb58d73d24a2181d3c2ff67a3e14b8edf64a04e4a753445f0e91829284d307826775fda9fdb17a05d04438499b2ba23a47c23981c66dfb93002130022d27d98b232b7d349c82db17bf2eb219db5eb96d084b3be454cc16483c2bcbc11c6a31164cb4a7ad69712ad9e7397ad923d1350bf3487e086c9873c4edbfaa3d39272dbee735d39057619df63051ba199c4a5a6d12674a806ef9657ef93539ebe10c86b0176c0d7d4c230644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd470f863715db69a03d94481e561402ebbb72059ac5105e3a35bbfddaeb99342a6f135189d5cacc0ff531f48f1579de15f2fc7a0b5060d5ce20d5aa18ef058efc9d112c8641be58a444f252be5639ceef6eb140d794e49e6ceaeaec463fbe39ac2f125df4eda873f137493505c81db0f7abb375f79420c328166dddadda3a2de0df2b42c1dc21b6b47d7ec0c761f1d1f2c061ce97414bc9c1d5d21de5e111aa4948183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea41c37db09cbc64991a2d21cfdadd8a49b9563db11bb0b7b4e6e9a83caccb3ca3d2b8de2a5cfa1b4b635abbe6df77b63eca9229ae853343cfe87a5e7605d9f7eeb1ae5411b2a133b83a5b32c3b17c6747ed354fdb6c53ab5d9342fc0abd144c65330644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000117965e12bcab6bc670b717479f56716ff72f04dd568cc1ce164037b37947b5de15db9e0182f5ac7c2137e882896a6cd36be1c9c13289ef81e16e60b564186ba21896de5db042383accbd62a429a2b75710a863b9ee452d8dcdcedb495a89e4602a5249a355749ccfe0335aab56b68477ded338386838d36e2655535b4a243f671851a4ad8709115dbfdcf6d239630b5976ee0ea4fc337d8d85d9974b925a8bd01df18f00336597603dc5fca563817d3b6e9cf45c73536a1720920d272851f05711379f9d9b1470c13f6ec28730174c52076ace94272ec82ebec6b8e1eb4d9af21482a546dd39af0aebe6388e32cfef964a4ef2699da878b69b0ccbff0d6c4f9030644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd441a9a3f45b0d9e4c7e23afe89483d31320d4be06b731fc29c6021b1f6fd68bd7314c3213063314a0b68b6fa02449927384ac79597799c4d4dae600ca9947f09160832b1469b343d36c4bc7c3fa4d7c6f9f5614b3770823b599b8f3cd67aa9533814af2c3a54bc734bc1ca57ec1aa4ae537f62b90cdf9b293a0a4b567ed55dba1a162127fc85c1fd1240046567bbf220ed146b5a703fa79b0419506c0be3e1b2912e48c8c177cf52d0f2cfebc9f3f5602eeed3568532f841c98959d2fb6b2dbd852c63408c7cd4985d0e6b9796ff4c7829d1d323fd6d681102fb300bf710e3b0db2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5a26469706673582212207511c0073d5b800588e0556ec1b7598d306ea657477c4b3041b3104059767d4a64736f6c634300081c0033",
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
