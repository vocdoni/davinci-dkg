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
	Bin: "0x608080604052346015576116ed908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80630ea14a3914610a88578063233ace1114610a4e578063454c28a31461063e5780635f89feef146101ad5763b8e72af614610051575f80fd5b3461018a57604036600319011261018a576004356001600160401b03811161018a57610081903690600401610b50565b6024356001600160401b03811161018a576100a0903690600401610b50565b610100830361019e5781016101208282031261018a5780601f8301121561018a57610120604051926100d28285610b7d565b8391810192831161018a57905b82821061018e57505050303b1561018a5760405163454c28a360e01b8152610140600482015261014481018390529283919083906101648401375f6101648484010152602482015f905b6009821061017057505050610164815f93601f80199101168101030181305afa801561016557610157575080f35b61016391505f90610b7d565b005b6040513d5f823e3d90fd5b829350602080916001939451815201930191018492610129565b5f80fd5b81358152602091820191016100df565b63236bd13760e01b5f5260045ffd5b3461018a576101a036600319011261018a573660841161018a57366101a41161018a576103006040516101e08282610b7d565b813682376101ef600435610e80565b610200602493929335604435610eeb565b9193929061020f606435610e80565b9390926040519660408801965f5160206115d85f395f51905f5289528860208101985f5160206115785f395f51905f528a525f5160206113785f395f51905f5281525f5160206114585f395f51905f52604060608401925f5160206112585f395f51905f5284525f5160206114385f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f5160206113985f395f51905f5285525f5160206115985f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f5160206112185f395f51905f5285525f5160206115f85f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f5160206114b85f395f51905f5285525f5160206115185f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f5160206111b85f395f51905f5285525f5160206113585f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f5160206112785f395f51905f5285525f5160206114985f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f5160206115b85f395f51905f5285525f5160206114185f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f5160206112b85f395f51905f5285525f5160206112985f395f51905f5288526101643590818a5287838760608160075afa921016169160808160065afa16945f5160206114785f395f51905f528352526101843580955260608160075afa9210161660408a60808160065afa1698519751981561062f5760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f5160206112f85f395f51905f526101008401525f5160206116785f395f51905f526101208401525f5160206116385f395f51905f526101408401525f5160206116985f395f51905f526101608401525f5160206113d85f395f51905f526101808401525f5160206113f85f395f51905f526101a08401525f5160206116185f395f51905f526101c08401525f5160206111f85f395f51905f526101e08401525f5160206114f85f395f51905f526102008401525f5160206112385f395f51905f526102208401526102408301526102608201525f5160206114d85f395f51905f526102808201525f5160206115585f395f51905f526102a08201525f5160206113385f395f51905f526102c08201525f5160206112d85f395f51905f526102e08201526040519283916105fb8484610b7d565b8336843760085afa15908115610622575b5061061357005b631ff3747d60e21b5f5260045ffd5b600191505114158161060c565b63a54f8e2760e01b5f5260045ffd5b3461018a5761014036600319011261018a576004356001600160401b03811161018a5761066f903690600401610b50565b366101441161018a576101006106859114610bb4565b604051604081015f5160206115d85f395f51905f52825260208201905f5160206115785f395f51905f5282525f5160206113785f395f51905f528152606083015f5160206112585f395f51905f5281525f5160206114585f395f51905f526040602435935f5160206114385f395f51905f52608088019580875284848460608160075afa911016838960808160065afa165f5160206113985f395f51905f5283525f5160206115985f395f51905f52865260016044359182895286868660608160075afa9310161616838960808160065afa165f5160206112185f395f51905f5283525f5160206115f85f395f51905f5286526064359081885285858560608160075afa92101616838960808160065afa165f5160206114b85f395f51905f5283525f5160206115185f395f51905f5286526084359081885285858560608160075afa92101616838960808160065afa165f5160206111b85f395f51905f5283525f5160206113585f395f51905f52865260a4359081885285858560608160075afa92101616838960808160065afa165f5160206112785f395f51905f5283525f5160206114985f395f51905f52865260c4359081885285858560608160075afa92101616838960808160065afa165f5160206115b85f395f51905f5283525f5160206114185f395f51905f52865260e4359081885285858560608160075afa92101616838960808160065afa165f5160206112b85f395f51905f5283525f5160206112985f395f51905f528652610104359081885285858560608160075afa92101616838960808160065afa16945f5160206114785f395f51905f528352526101243580955260608160075afa9210161660408360808160065afa1691519051911561062f576101006040519384375f5160206112f85f395f51905f526101008401525f5160206116785f395f51905f526101208401525f5160206116385f395f51905f526101408401525f5160206116985f395f51905f526101608401525f5160206113d85f395f51905f526101808401525f5160206113f85f395f51905f526101a08401525f5160206116185f395f51905f526101c08401525f5160206111f85f395f51905f526101e08401525f5160206114f85f395f51905f526102008401525f5160206112385f395f51905f526102208401526102408301526102608201525f5160206114d85f395f51905f526102808201525f5160206115585f395f51905f526102a08201525f5160206113385f395f51905f526102c08201525f5160206112d85f395f51905f526102e08201526020816103008160085afa9051161561061357005b3461018a575f36600319011261018a5760206040517f580a117c037fe7ca3465d9e87421e09eec03f7f9dc5d4965369bedd4ec6fad108152f35b3461018a57602036600319011261018a576004356001600160401b03811161018a57610ab8903690600401610b50565b610b21608092610adc61010060405194610ad28787610b7d565b8636873714610bb4565b610aeb60208201358235610bf7565b8352610b088482013560a083013560408401356060850135610c98565b6020850152604084015260c060e0820135910135610bf7565b6060820152604051905f825b60048210610b3a57505050f35b6020806001928551815201930191019091610b2d565b9181601f8401121561018a578235916001600160401b03831161018a576020838186019501011161018a57565b601f909101601f19168101906001600160401b03821190821017610ba057604052565b634e487b7160e01b5f52604160045260245ffd5b15610bbb57565b60405162461bcd60e51b81526020600482015260146024820152730d2dcecc2d8d2c840e0e4dedecc40d8cadccee8d60631b6044820152606490fd5b905f5160206113185f395f51905f528210801590610c81575b61061357811580610c79575b610c7357610c405f5160206113185f395f51905f5260038185818180090908610feb565b818103610c4f57505060011b90565b5f5160206113185f395f51905f52809106810306145f1461061357600190811b1790565b50505f90565b508015610c1c565b505f5160206113185f395f51905f52811015610c10565b919093925f5160206113185f395f51905f528310801590610e69575b8015610e52575b8015610e3b575b610613578082868517171715610e3057908291610d935f5160206113185f395f51905f5280808080888180808f9d5f5160206115385f395f51905f528f839290839109099d8e0981848181800909085f5160206116585f395f51905f52089a09818c8181800909085f5160206111d85f395f51905f520806810306945f5160206113185f395f51905f525f5160206113b85f395f51905f5281610d6d81808b80098187800908610feb565b8408095f5160206113185f395f51905f52610d878261114f565b8009141595869161100e565b929080821480610e27575b15610dc55750505050905f14610dbd5760ff60025b169060021b179190565b60ff5f610db3565b5f5160206113185f395f51905f52809106810306149182610e08575b5050156106135760019115610e005760ff60025b169060021b17179190565b60ff5f610df5565b5f5160206113185f395f51905f52919250819006810306145f80610de1565b50838314610d9e565b50505090505f905f90565b505f5160206113185f395f51905f52811015610cc2565b505f5160206113185f395f51905f52821015610cbb565b505f5160206113185f395f51905f52851015610cb4565b8015610ee4578060011c915f5160206113185f395f51905f5283101561061357600180610ec35f5160206113185f395f51905f5260038188818180090908610feb565b931614610ecc57565b905f5160206113185f395f51905f5280910681030690565b505f905f90565b801580610fe3575b610fd7578060021c92825f5160206113185f395f51905f528510801590610fc0575b6106135784815f5160206113185f395f51905f5280808080808080805f5160206115385f395f51905f5281610f8a9d8d0909998a0981898181800909085f5160206111d85f395f51905f520806810306936002808a16149509818a8181800909085f5160206116585f395f51905f520861100e565b80929160018082961614610f9c575050565b5f5160206113185f395f51905f528093945080929550809106810306930681030690565b505f5160206113185f395f51905f52811015610f15565b50505f905f905f905f90565b508115610ef3565b90610ff58261114f565b915f5160206113185f395f51905f528380090361061357565b915f5160206113185f395f51905f525f5160206113b85f395f51905f52816110539396949661104582808a8009818a800908610feb565b90611143575b860809610feb565b925f5160206113185f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206113185f395f51905f5260a083015260208260c08160055afa91519115610613575f5160206113185f395f51905f52826001920903610613575f5160206113185f395f51905f52908209925f5160206113185f395f51905f528080808780090681030681878009081490811591611124575b5061061357565b90505f5160206113185f395f51905f528084860960020914155f61111d565b8180910681030661104b565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206113185f395f51905f5260a083015260208260c08160055afa915191156106135756fe2c8dbc8c703fe6a549505cb135b6b6700f15d741c710600151861b6a7550ac442fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77519440437276d9f6a1d690a6c14156f5ffa930516e577d24d139151093cb637d11b78760521864fd637e7b77e0df1d6e5d70750224dc974545e598d3bc8d1a9940a20753c5c33d971188fd86b79c4eb683b4fa6c68800f4eb9eecc5943e57ddd808d24293b132a5a5a1e7107aaa6b52d5322aef320a77eae409ea84263a4b42f706916a1e86898d0f21249505a7a591ec3017d74519a7c92378b17ac9aab8609909c4fa0d54f09b2cc9e45b13486681db52eb6c815a195eee8f810f084a9eb2e1146770e07cc054a154e875a52f2cd7f6b4f28954d576c686f18ff0b28c3a163b05b9de381c22217d1e34e32b3e7340a77defdf7a3ce6e36b57821fae055dfb1c0bb46d9c1fd49a19cc0d77db1122922339ba0d015a65eafb9702269f48181c2730644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd471365056e8df437180e8003418433daf0ad526edf0135e96dde80a62a9cb372d81616ae6a774fcb75d2803bc2b06cc6ff30de86302c84fbc3c35f3ef6d68fd9c52f2f7234f1f38e2583e1250371f0c7f2302a353060a78d685576cd8a56b2527c0c9e19cae04e330a15a5f940babba38ce790d46fb0f8286c20b5b8ff21d6a300183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea40716f697e353c5d9d39df6468ffaecd4e9935db724779bc24767163361ba5dae075f1c98ed8da7ad260ddd6840dca6c62c7d9943e9199a5628abd45b30eda6d921f0bb6d308da591937003724c5ae2a206e2a6288b7380470ee05cd2fca4cbec2e956e92a605a2844a1297e0c91c367b05bae5a95054bc349944ff0e9861fafb30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000012192677b650ee0a0201be908c5e9800cb1bdf263f76e1abd7ffb78ab4062a50e17dbaa64c96e362f2afe774da3c89eb21fb808a368ea615b3f61a4ae598427c22a5292e2f506cd53a2d373d109b9fbb31e9cc989167debf07fa80801a5a1700c2de65512b72bc222b016138d8d4898237e002a228b2db518f67d0c97772446d71207f129043423701287fb38707f6a5bc2f712adf5b40857dd68e4c52c284354141a46f2d4a42edd4a851201625a10287bd6e139f30232965fc1cf7137c997ac30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4403599f6d26ce610c0e29c79ca084bdc38bc099124bf7f08e834a651a11833f2e06844f24a4a6311d8c3eba216d5b46803076c216db851283847010cc8bd8476e2139e52f679015503f90eadd2fe4a030a009ca7830081e56e0003baa0471415905fb653a2ca0e69afc77978028ddf096275247b01994d39cf01f8310fbbe838827f97e12e83679fddba97e8e9ac2fb2c2cd1440901ea637d4e03b2d28ce9e9b60841b69eb6af4aa7ef8e8513c452fe326b860270a871ea295875537f5114989c0055943c5b405b48dd854de7f0653ede2a9f06822a2868e699b5f4a4156b5e4a201205087c26b43f9e7cc00813b7e50705979c1f31fa542c13ee2ff04960a1d22b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e51a136430c6cc777a9229ad9703cbc21f9592268528abbe5b470fb3c5a4e66a1011efcb758016cfed1c81f6b68010e5970b94ea696130136a13f5961e1efac7dfa2646970667358221220c30e9241797223b777513bc10532495c754e63b5eba1d4d87ffe402d42f8637564736f6c634300081c0033",
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
