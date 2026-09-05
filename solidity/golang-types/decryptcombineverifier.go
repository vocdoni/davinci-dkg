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
	Bin: "0x608080604052346015576116ed908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80630ea14a3914610a88578063233ace1114610a4e578063454c28a31461063e5780635f89feef146101ad5763b8e72af614610051575f80fd5b3461018a57604036600319011261018a576004356001600160401b03811161018a57610081903690600401610b50565b6024356001600160401b03811161018a576100a0903690600401610b50565b610100830361019e5781016101208282031261018a5780601f8301121561018a57610120604051926100d28285610b7d565b8391810192831161018a57905b82821061018e57505050303b1561018a5760405163454c28a360e01b8152610140600482015261014481018390529283919083906101648401375f6101648484010152602482015f905b6009821061017057505050610164815f93601f80199101168101030181305afa801561016557610157575080f35b61016391505f90610b7d565b005b6040513d5f823e3d90fd5b829350602080916001939451815201930191018492610129565b5f80fd5b81358152602091820191016100df565b63236bd13760e01b5f5260045ffd5b3461018a576101a036600319011261018a573660841161018a57366101a41161018a576103006040516101e08282610b7d565b813682376101ef600435610e80565b610200602493929335604435610eeb565b9193929061020f606435610e80565b9390926040519660408801965f5160206112185f395f51905f5289528860208101985f5160206114385f395f51905f528a525f5160206115185f395f51905f5281525f5160206114985f395f51905f52604060608401925f5160206115b85f395f51905f5284525f5160206113385f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f5160206115385f395f51905f5285525f5160206112785f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f5160206113b85f395f51905f5285525f5160206116585f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f5160206112985f395f51905f5285525f5160206115785f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f5160206113785f395f51905f5285525f5160206113585f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f5160206114185f395f51905f5285525f5160206113f85f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f5160206114585f395f51905f5285525f5160206112585f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f5160206114b85f395f51905f5285525f5160206114f85f395f51905f5288526101643590818a5287838760608160075afa921016169160808160065afa16945f5160206111f85f395f51905f528352526101843580955260608160075afa9210161660408a60808160065afa1698519751981561062f5760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f5160206113185f395f51905f526101008401525f5160206114d85f395f51905f526101208401525f5160206112385f395f51905f526101408401525f5160206112b85f395f51905f526101608401525f5160206116185f395f51905f526101808401525f5160206114785f395f51905f526101a08401525f5160206116385f395f51905f526101c08401525f5160206116785f395f51905f526101e08401525f5160206113985f395f51905f526102008401525f5160206111d85f395f51905f526102208401526102408301526102608201525f5160206115d85f395f51905f526102808201525f5160206115585f395f51905f526102a08201525f5160206112d85f395f51905f526102c08201525f5160206115985f395f51905f526102e08201526040519283916105fb8484610b7d565b8336843760085afa15908115610622575b5061061357005b631ff3747d60e21b5f5260045ffd5b600191505114158161060c565b63a54f8e2760e01b5f5260045ffd5b3461018a5761014036600319011261018a576004356001600160401b03811161018a5761066f903690600401610b50565b366101441161018a576101006106859114610bb4565b604051604081015f5160206112185f395f51905f52825260208201905f5160206114385f395f51905f5282525f5160206115185f395f51905f528152606083015f5160206115b85f395f51905f5281525f5160206114985f395f51905f526040602435935f5160206113385f395f51905f52608088019580875284848460608160075afa911016838960808160065afa165f5160206115385f395f51905f5283525f5160206112785f395f51905f52865260016044359182895286868660608160075afa9310161616838960808160065afa165f5160206113b85f395f51905f5283525f5160206116585f395f51905f5286526064359081885285858560608160075afa92101616838960808160065afa165f5160206112985f395f51905f5283525f5160206115785f395f51905f5286526084359081885285858560608160075afa92101616838960808160065afa165f5160206113785f395f51905f5283525f5160206113585f395f51905f52865260a4359081885285858560608160075afa92101616838960808160065afa165f5160206114185f395f51905f5283525f5160206113f85f395f51905f52865260c4359081885285858560608160075afa92101616838960808160065afa165f5160206114585f395f51905f5283525f5160206112585f395f51905f52865260e4359081885285858560608160075afa92101616838960808160065afa165f5160206114b85f395f51905f5283525f5160206114f85f395f51905f528652610104359081885285858560608160075afa92101616838960808160065afa16945f5160206111f85f395f51905f528352526101243580955260608160075afa9210161660408360808160065afa1691519051911561062f576101006040519384375f5160206113185f395f51905f526101008401525f5160206114d85f395f51905f526101208401525f5160206112385f395f51905f526101408401525f5160206112b85f395f51905f526101608401525f5160206116185f395f51905f526101808401525f5160206114785f395f51905f526101a08401525f5160206116385f395f51905f526101c08401525f5160206116785f395f51905f526101e08401525f5160206113985f395f51905f526102008401525f5160206111d85f395f51905f526102208401526102408301526102608201525f5160206115d85f395f51905f526102808201525f5160206115585f395f51905f526102a08201525f5160206112d85f395f51905f526102c08201525f5160206115985f395f51905f526102e08201526020816103008160085afa9051161561061357005b3461018a575f36600319011261018a5760206040517fd70de162a56ac4801077f857f9491015916bb238db7fb9b0cfe5061eceae53058152f35b3461018a57602036600319011261018a576004356001600160401b03811161018a57610ab8903690600401610b50565b610b21608092610adc61010060405194610ad28787610b7d565b8636873714610bb4565b610aeb60208201358235610bf7565b8352610b088482013560a083013560408401356060850135610c98565b6020850152604084015260c060e0820135910135610bf7565b6060820152604051905f825b60048210610b3a57505050f35b6020806001928551815201930191019091610b2d565b9181601f8401121561018a578235916001600160401b03831161018a576020838186019501011161018a57565b601f909101601f19168101906001600160401b03821190821017610ba057604052565b634e487b7160e01b5f52604160045260245ffd5b15610bbb57565b60405162461bcd60e51b81526020600482015260146024820152730d2dcecc2d8d2c840e0e4dedecc40d8cadccee8d60631b6044820152606490fd5b905f5160206112f85f395f51905f528210801590610c81575b61061357811580610c79575b610c7357610c405f5160206112f85f395f51905f5260038185818180090908610feb565b818103610c4f57505060011b90565b5f5160206112f85f395f51905f52809106810306145f1461061357600190811b1790565b50505f90565b508015610c1c565b505f5160206112f85f395f51905f52811015610c10565b919093925f5160206112f85f395f51905f528310801590610e69575b8015610e52575b8015610e3b575b610613578082868517171715610e3057908291610d935f5160206112f85f395f51905f5280808080888180808f9d5f5160206115f85f395f51905f528f839290839109099d8e0981848181800909085f5160206116985f395f51905f52089a09818c8181800909085f5160206111b85f395f51905f520806810306945f5160206112f85f395f51905f525f5160206113d85f395f51905f5281610d6d81808b80098187800908610feb565b8408095f5160206112f85f395f51905f52610d878261114f565b8009141595869161100e565b929080821480610e27575b15610dc55750505050905f14610dbd5760ff60025b169060021b179190565b60ff5f610db3565b5f5160206112f85f395f51905f52809106810306149182610e08575b5050156106135760019115610e005760ff60025b169060021b17179190565b60ff5f610df5565b5f5160206112f85f395f51905f52919250819006810306145f80610de1565b50838314610d9e565b50505090505f905f90565b505f5160206112f85f395f51905f52811015610cc2565b505f5160206112f85f395f51905f52821015610cbb565b505f5160206112f85f395f51905f52851015610cb4565b8015610ee4578060011c915f5160206112f85f395f51905f5283101561061357600180610ec35f5160206112f85f395f51905f5260038188818180090908610feb565b931614610ecc57565b905f5160206112f85f395f51905f5280910681030690565b505f905f90565b801580610fe3575b610fd7578060021c92825f5160206112f85f395f51905f528510801590610fc0575b6106135784815f5160206112f85f395f51905f5280808080808080805f5160206115f85f395f51905f5281610f8a9d8d0909998a0981898181800909085f5160206111b85f395f51905f520806810306936002808a16149509818a8181800909085f5160206116985f395f51905f520861100e565b80929160018082961614610f9c575050565b5f5160206112f85f395f51905f528093945080929550809106810306930681030690565b505f5160206112f85f395f51905f52811015610f15565b50505f905f905f905f90565b508115610ef3565b90610ff58261114f565b915f5160206112f85f395f51905f528380090361061357565b915f5160206112f85f395f51905f525f5160206113d85f395f51905f52816110539396949661104582808a8009818a800908610feb565b90611143575b860809610feb565b925f5160206112f85f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206112f85f395f51905f5260a083015260208260c08160055afa91519115610613575f5160206112f85f395f51905f52826001920903610613575f5160206112f85f395f51905f52908209925f5160206112f85f395f51905f528080808780090681030681878009081490811591611124575b5061061357565b90505f5160206112f85f395f51905f528084860960020914155f61111d565b8180910681030661104b565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206112f85f395f51905f5260a083015260208260c08160055afa915191156106135756fe2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7752a8fe459e7d0d352816f2e00500422ab178c28b9164fa5979fc72d7fa9903ed312962289249d9403c29cf8934de999ef702d1ce45af90ed8256ca80add4eb2bd1cc7b9e4bf2c1cc55608cfc9732d1cccc5809ab41fefda50fae538b0df696c8f20bf36558b83e7bf2d5c5528d89f4a3252d1120ccfbeebab5e71536da168020710f8ddfb921ee8b86f7c4cd5d6ddbbd3544f8d4e2ba8219a3dbea6058fd02522274035024b7413a2a52aba6c721c951e7708be7c2b32a6cb26017581dbbf33b929ae1300c076bcfcd4d728dd0a3a2f0b17335481d90d1ff77f5d4d16997112a208bb74164e00751655ea67d19c0711787a1616663403427274c8469d9ca8893900feb5c65ee1722f334cc1d17cafdd88b1437cb922d4fa31ede96f8753d26e6030644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4711efc3e2a655adf57ae706a84505b91591a7809dd049417c8c636c66d3c4878a2619d2c3e26206b94173551b54e1b2821bb3b448f3d993cfc3dc6eb179fd6a8614dcf554521bbfeeb57e85c88e4e00e77e058bb91332c6fb0c5239660c1730d90f540fc38bcedef99c8bc9b175bbe3ce3df7debbfdca1e73b6eea40a70166bb0269c66aca3eb31eaadea05948d8193634316ac9039be92539a862059215b550825b3aac7e8d2fa3dc06215580c0238948ec738b7717b027bb6d994b37b56846a183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea42168a722a7c970c0a6659ffc52901699cc520c1e64b03cdb98df10f62bceab85266f56bfb23367baab34a6b8db947e4321b6fbc52b5a61d15a856c43eb0cf7b02c0f4107a1298e403da21157056a2ff668bb0e5ae2568ece5bba9e5138b6e512270f1a35278a9a6a28f085d5de068e16a29eec067ea81a87d759d0c70dbd92d205bcc1aaa1e517ae89b4a59c1693bc87fd9a9bf60bdaa9e35e85b7aacf08e9bf30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000124dc28b18212aa7f370d847f47276a000f5e3247ef7b912b63e0d882d94c2d582a3b2f3d3dd52b965ab0292bf2324c8507937633789fc705ac553c5b9a390f7b12870b709c34582cea7e01b1c2669bc4a28053a92031bb3738a0d5b8cb42f9601347c6142850d6d57dac8238e68cea714d126b7a926fedc5c0f7dd88553293da2796ffca33905f26572cf6efee83fdbe26d363224f37b7497bad38b4f3841abf050a7adc618f41747017ced88448e9c0eb55f7f7bd6bedfc142a36fb12ad272b223b1a2ebe0c816a41f0feda8ee423cc4353a6aec1a4913487a037f57a4ecd9c27f47eb8df8e87e8893647b6a66cfb0a77378b7301df925bd3e5610ca16a91e507ba5b72dd6a974e3f2f8df8a43db53cb86ef47c16ec8c780fd69b4f7d08bb5511798df22b204e8bf0899cfe25c53c1c65ff0bbc32c557677960c2dfb5fa064530644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4429f8a5036212e7f38dc3258e1f8bd0fecb7cf7f8c254a6ce013a3d77e7df79b2082e4f222f1ada33f2754d14d87c8124d5d35a1618dbb31ca42c42938eb3a2730f1805380b9795c39748dc857b59f23a3f2c82c8013fd582758386d73a7aee1b2a2c2dbbb6ae9ec1b16ac1a00231e95a61bc624f6c2e0e41794cd7fe7dd056ce2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5a264697066735822122055a588054d1e31342eba8d263f5b6b03ecee5120dfa84dbfcd4faf172da8bb6c64736f6c634300081c0033",
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
