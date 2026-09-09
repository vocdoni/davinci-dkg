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
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[9]\",\"internalType\":\"uint256[9]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"uint256[9]\",\"internalType\":\"uint256[9]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidProofEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x608080604052346015576116ed908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80630ea14a3914610a88578063233ace1114610a4e578063454c28a31461063e5780635f89feef146101ad5763b8e72af614610051575f80fd5b3461018a57604036600319011261018a576004356001600160401b03811161018a57610081903690600401610b50565b6024356001600160401b03811161018a576100a0903690600401610b50565b610100830361019e5781016101208282031261018a5780601f8301121561018a57610120604051926100d28285610b7d565b8391810192831161018a57905b82821061018e57505050303b1561018a5760405163454c28a360e01b8152610140600482015261014481018390529283919083906101648401375f6101648484010152602482015f905b6009821061017057505050610164815f93601f80199101168101030181305afa801561016557610157575080f35b61016391505f90610b7d565b005b6040513d5f823e3d90fd5b829350602080916001939451815201930191018492610129565b5f80fd5b81358152602091820191016100df565b63236bd13760e01b5f5260045ffd5b3461018a576101a036600319011261018a573660841161018a57366101a41161018a576103006040516101e08282610b7d565b813682376101ef600435610e80565b610200602493929335604435610eeb565b9193929061020f606435610e80565b9390926040519660408801965f5160206115f85f395f51905f5289528860208101985f5160206114d85f395f51905f528a525f5160206114f85f395f51905f5281525f5160206113d85f395f51905f52604060608401925f5160206113b85f395f51905f5284525f5160206116785f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f5160206112185f395f51905f5285525f5160206113985f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f5160206115185f395f51905f5285525f5160206115b85f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f5160206116385f395f51905f5285525f5160206116985f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f5160206112d85f395f51905f5285525f5160206112385f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f5160206112f85f395f51905f5285525f5160206114585f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f5160206114985f395f51905f5285525f5160206115985f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f5160206113585f395f51905f5285525f5160206111d85f395f51905f5288526101643590818a5287838760608160075afa921016169160808160065afa16945f5160206112585f395f51905f528352526101843580955260608160075afa9210161660408a60808160065afa1698519751981561062f5760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f5160206115585f395f51905f526101008401525f5160206114b85f395f51905f526101208401525f5160206115d85f395f51905f526101408401525f5160206112b85f395f51905f526101608401525f5160206111f85f395f51905f526101808401525f5160206113385f395f51905f526101a08401525f5160206114185f395f51905f526101c08401525f5160206114385f395f51905f526101e08401525f5160206115385f395f51905f526102008401525f5160206116185f395f51905f526102208401526102408301526102608201525f5160206114785f395f51905f526102808201525f5160206113185f395f51905f526102a08201525f5160206113f85f395f51905f526102c08201525f5160206112785f395f51905f526102e08201526040519283916105fb8484610b7d565b8336843760085afa15908115610622575b5061061357005b631ff3747d60e21b5f5260045ffd5b600191505114158161060c565b63a54f8e2760e01b5f5260045ffd5b3461018a5761014036600319011261018a576004356001600160401b03811161018a5761066f903690600401610b50565b366101441161018a576101006106859114610bb4565b604051604081015f5160206115f85f395f51905f52825260208201905f5160206114d85f395f51905f5282525f5160206114f85f395f51905f528152606083015f5160206113b85f395f51905f5281525f5160206113d85f395f51905f526040602435935f5160206116785f395f51905f52608088019580875284848460608160075afa911016838960808160065afa165f5160206112185f395f51905f5283525f5160206113985f395f51905f52865260016044359182895286868660608160075afa9310161616838960808160065afa165f5160206115185f395f51905f5283525f5160206115b85f395f51905f5286526064359081885285858560608160075afa92101616838960808160065afa165f5160206116385f395f51905f5283525f5160206116985f395f51905f5286526084359081885285858560608160075afa92101616838960808160065afa165f5160206112d85f395f51905f5283525f5160206112385f395f51905f52865260a4359081885285858560608160075afa92101616838960808160065afa165f5160206112f85f395f51905f5283525f5160206114585f395f51905f52865260c4359081885285858560608160075afa92101616838960808160065afa165f5160206114985f395f51905f5283525f5160206115985f395f51905f52865260e4359081885285858560608160075afa92101616838960808160065afa165f5160206113585f395f51905f5283525f5160206111d85f395f51905f528652610104359081885285858560608160075afa92101616838960808160065afa16945f5160206112585f395f51905f528352526101243580955260608160075afa9210161660408360808160065afa1691519051911561062f576101006040519384375f5160206115585f395f51905f526101008401525f5160206114b85f395f51905f526101208401525f5160206115d85f395f51905f526101408401525f5160206112b85f395f51905f526101608401525f5160206111f85f395f51905f526101808401525f5160206113385f395f51905f526101a08401525f5160206114185f395f51905f526101c08401525f5160206114385f395f51905f526101e08401525f5160206115385f395f51905f526102008401525f5160206116185f395f51905f526102208401526102408301526102608201525f5160206114785f395f51905f526102808201525f5160206113185f395f51905f526102a08201525f5160206113f85f395f51905f526102c08201525f5160206112785f395f51905f526102e08201526020816103008160085afa9051161561061357005b3461018a575f36600319011261018a5760206040517f0a82f57b9a605ea4fb7f72d305887e7b441e405ef42893eb16c0dfa462d937858152f35b3461018a57602036600319011261018a576004356001600160401b03811161018a57610ab8903690600401610b50565b610b21608092610adc61010060405194610ad28787610b7d565b8636873714610bb4565b610aeb60208201358235610bf7565b8352610b088482013560a083013560408401356060850135610c98565b6020850152604084015260c060e0820135910135610bf7565b6060820152604051905f825b60048210610b3a57505050f35b6020806001928551815201930191019091610b2d565b9181601f8401121561018a578235916001600160401b03831161018a576020838186019501011161018a57565b601f909101601f19168101906001600160401b03821190821017610ba057604052565b634e487b7160e01b5f52604160045260245ffd5b15610bbb57565b60405162461bcd60e51b81526020600482015260146024820152730d2dcecc2d8d2c840e0e4dedecc40d8cadccee8d60631b6044820152606490fd5b905f5160206112985f395f51905f528210801590610c81575b61061357811580610c79575b610c7357610c405f5160206112985f395f51905f5260038185818180090908610feb565b818103610c4f57505060011b90565b5f5160206112985f395f51905f52809106810306145f1461061357600190811b1790565b50505f90565b508015610c1c565b505f5160206112985f395f51905f52811015610c10565b919093925f5160206112985f395f51905f528310801590610e69575b8015610e52575b8015610e3b575b610613578082868517171715610e3057908291610d935f5160206112985f395f51905f5280808080888180808f9d5f5160206115785f395f51905f528f839290839109099d8e0981848181800909085f5160206116585f395f51905f52089a09818c8181800909085f5160206111b85f395f51905f520806810306945f5160206112985f395f51905f525f5160206113785f395f51905f5281610d6d81808b80098187800908610feb565b8408095f5160206112985f395f51905f52610d878261114f565b8009141595869161100e565b929080821480610e27575b15610dc55750505050905f14610dbd5760ff60025b169060021b179190565b60ff5f610db3565b5f5160206112985f395f51905f52809106810306149182610e08575b5050156106135760019115610e005760ff60025b169060021b17179190565b60ff5f610df5565b5f5160206112985f395f51905f52919250819006810306145f80610de1565b50838314610d9e565b50505090505f905f90565b505f5160206112985f395f51905f52811015610cc2565b505f5160206112985f395f51905f52821015610cbb565b505f5160206112985f395f51905f52851015610cb4565b8015610ee4578060011c915f5160206112985f395f51905f5283101561061357600180610ec35f5160206112985f395f51905f5260038188818180090908610feb565b931614610ecc57565b905f5160206112985f395f51905f5280910681030690565b505f905f90565b801580610fe3575b610fd7578060021c92825f5160206112985f395f51905f528510801590610fc0575b6106135784815f5160206112985f395f51905f5280808080808080805f5160206115785f395f51905f5281610f8a9d8d0909998a0981898181800909085f5160206111b85f395f51905f520806810306936002808a16149509818a8181800909085f5160206116585f395f51905f520861100e565b80929160018082961614610f9c575050565b5f5160206112985f395f51905f528093945080929550809106810306930681030690565b505f5160206112985f395f51905f52811015610f15565b50505f905f905f905f90565b508115610ef3565b90610ff58261114f565b915f5160206112985f395f51905f528380090361061357565b915f5160206112985f395f51905f525f5160206113785f395f51905f52816110539396949661104582808a8009818a800908610feb565b90611143575b860809610feb565b925f5160206112985f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206112985f395f51905f5260a083015260208260c08160055afa91519115610613575f5160206112985f395f51905f52826001920903610613575f5160206112985f395f51905f52908209925f5160206112985f395f51905f528080808780090681030681878009081490811591611124575b5061061357565b90505f5160206112985f395f51905f528084860960020914155f61111d565b8180910681030661104b565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206112985f395f51905f5260a083015260208260c08160055afa915191156106135756fe2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775169976666b32ce7c404eea24cb025dda64f2f292703a860c5c5dd9b86228af350123bdc1d3e27e087d80bcbbc34a8f014cb093a2e9bc2e597df0d15a9fc3354624b8da4819859137b3036cbefa57add2ddcfcdeb462d5429cd569cb89e7b489a10ec7f39e2d3b1ac2b805fb398deae41aaf5922e4dfe559d8ced2ba8b746ca36224fd1fb589549e338b73b2ebde788c22f347468e954f9dd5d33f31d217666e71f26279a2d69b9be0ab3372ae922b9ad2ffcc413737bc8e9cf768143d27041e130644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47132751202b0b801e6a341177eed4a16612aa63d62966e0a0d648c96db90bec6d27231a355a0de15227e95979e9fe5bb26ab2e4140f33f0ac5c75a0da0038534f28d7a2296c4bdec44a108af587a98a264ad89fbd8fd8b981c3cbc3f72f8c07672b8f2036a314c45f2da3524cceb5a247dd3c461642f9d0a7da03fba45f1998ad2b223a0806b59222d06a5c2457b326295d0001ec4a8ff391a41fd3d6143356e42a03c99058d6fec6ae1f87ac15a758ed136280cac636577d0fa8cd8109a88640183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea41efa3fee7446a8534ca417c48ed63761aace4bbf6185f0d6e73873e07f68ed532bcbbd3ab6864bbad78e7ee091460b69a1b8374a33dcdeb25804799e265972e430644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000010ad88a49bbf03d23d294efbdf5ba81f43dea8265e35638b2b3f0f4c27a9143c809ec4dab9eaa75b69a5dc9e2965fcede1fb51c07fdb6b1654ec465e71b3f8ab806be84d8ffed058907ccdbb8b1dcb6a37f1bd4f6eefdb6c95be4c533821c1f3d21ae11419c7c0c9c01da850b61380bd4f2d5d9c5034fbf26d7304db01b4789bd11becd48184c62ed1a3ad5715faeaef0c77e00eaac56688999f8ed727ccab79c231495e843a03a2cf6929da1fc737b1bdb421630db6486a17fd907642ccc2cbd1975254e2303b182f4fc3d927d46d38169ebcdef2791fec12d3e7f042a2d68bc26746f31da18294d70fe5b1b302570c67c6d0eca44756d5d8fdadd28ee02bbad0ee62957195dbc0e0962210017866ba6825325d0e7697d01d605245326140fcc181db40c3dd9dbea9cc9b50536007ec49236bf94e82cc6916be262fe677f028b2192b0bff726c95cb32505ed28fe204aab6edf2a935973edfeb1fd95509716a618f94dbcec35365b14f4b5716a71bcd7f08d694e28f9dc19645236be1184a9c330644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd441985ed4fb04198aa989f308388582cc356ba51432b844e361fd0635247d89b502d5031e08cdd2cd7a5cb38201503119d3494e4936cf9481e8ba96ac2c73b92791cbe340464c9dc037b693fec4b50234471d0ca619934f1d89a55d6125ecbc3071c8f6ace87541d8d559b3fa87e045be387f2ea1a4931ab1e9ebed6aa9c6e5c082f5df6294ad27bd5717a2faedd7f0b98f0376e74ca43779596e17cf2ec306e460807afdcbe80f6db3a01a925ece8dbdfef4a49ca4d45f455a639846ee2b0adc82b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50c634439eddd5e192164b431428abdb7b42417263067a76eb4463394551bd7d427130dfbf5ead5f2c9fc380561eecd9d519806e528d00c17a414cbeb3f67abcba2646970667358221220f2a718465b6b8a2c4a05f503e2ad9cb3285935a9b81fcd911ff62573eca75ef664736f6c634300081c0033",
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

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed)
func (_DecryptCombineVerifier *DecryptCombineVerifierCaller) CompressProof(opts *bind.CallOpts, proof []byte) ([4]*big.Int, error) {
	var out []interface{}
	err := _DecryptCombineVerifier.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed)
func (_DecryptCombineVerifier *DecryptCombineVerifierSession) CompressProof(proof []byte) ([4]*big.Int, error) {
	return _DecryptCombineVerifier.Contract.CompressProof(&_DecryptCombineVerifier.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed)
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerSession) CompressProof(proof []byte) ([4]*big.Int, error) {
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

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5f89feef.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[9] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [9]*big.Int) error {
	var out []interface{}
	err := _DecryptCombineVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5f89feef.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[9] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [9]*big.Int) error {
	return _DecryptCombineVerifier.Contract.VerifyCompressedProof(&_DecryptCombineVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5f89feef.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[9] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [9]*big.Int) error {
	return _DecryptCombineVerifier.Contract.VerifyCompressedProof(&_DecryptCombineVerifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x454c28a3.
//
// Solidity: function verifyProof(bytes proof, uint256[9] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCaller) VerifyProof(opts *bind.CallOpts, proof []byte, input [9]*big.Int) error {
	var out []interface{}
	err := _DecryptCombineVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x454c28a3.
//
// Solidity: function verifyProof(bytes proof, uint256[9] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierSession) VerifyProof(proof []byte, input [9]*big.Int) error {
	return _DecryptCombineVerifier.Contract.VerifyProof(&_DecryptCombineVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x454c28a3.
//
// Solidity: function verifyProof(bytes proof, uint256[9] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerSession) VerifyProof(proof []byte, input [9]*big.Int) error {
	return _DecryptCombineVerifier.Contract.VerifyProof(&_DecryptCombineVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCaller) VerifyProof0(opts *bind.CallOpts, proof []byte, input []byte) error {
	var out []interface{}
	err := _DecryptCombineVerifier.contract.Call(opts, &out, "verifyProof0", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierSession) VerifyProof0(proof []byte, input []byte) error {
	return _DecryptCombineVerifier.Contract.VerifyProof0(&_DecryptCombineVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerSession) VerifyProof0(proof []byte, input []byte) error {
	return _DecryptCombineVerifier.Contract.VerifyProof0(&_DecryptCombineVerifier.CallOpts, proof, input)
}
