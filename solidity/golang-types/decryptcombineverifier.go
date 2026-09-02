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
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[11]\",\"internalType\":\"uint256[11]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"input\",\"type\":\"uint256[11]\",\"internalType\":\"uint256[11]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidProofEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x6080806040523460155761182e908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace1114610c305750806344f6369214610b975780638261a6531461072057806398ea1c081461020f5763b8e72af614610055575f80fd5b346101c85760403660031901126101c8576004356001600160401b0381116101c857610085903690600401610c68565b6024356001600160401b0381116101c8576100a4903690600401610c68565b90916101008103610200578301610100848203126101c85780601f850112156101c857604051936100d761010086610c95565b849061010081019283116101c857905b8282106101f0575050508101610160828203126101c85780601f830112156101c8576040519161011961016084610c95565b829061016081019283116101c857905b8282106101cc57505050303b156101c857604051638261a65360e01b8152915f600484015b600882106101b25750505061010482015f905b600b821061019c575050505f8161026481305afa801561019157610183575080f35b61018f91505f90610c95565b005b6040513d5f823e3d90fd5b6020806001928551815201930191019091610161565b602080600192855181520193019101909161014e565b5f80fd5b8135815260209182019101610129565b634e487b7160e01b5f52604160045260245ffd5b81358152602091820191016100e7565b63236bd13760e01b5f5260045ffd5b346101c8576101e03660031901126101c857366084116101c857366101e4116101c8576103006040516102428282610c95565b81368237610251600435610f41565b610262602493929335604435610fac565b91939290610271606435610f41565b9390926040519660408801965f5160206115995f395f51905f5289528860208101985f5160206114195f395f51905f528a525f5160206117795f395f51905f5281525f5160206115b95f395f51905f52604060608401925f5160206112995f395f51905f5284525f5160206115595f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f5160206116595f395f51905f5285525f5160206113395f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f5160206113595f395f51905f5285525f5160206114995f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f5160206115f95f395f51905f5285525f5160206116d95f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f5160206115795f395f51905f5285525f5160206117395f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f5160206117595f395f51905f5285525f5160206117195f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f5160206116995f395f51905f5285525f5160206115195f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f5160206112f95f395f51905f5285525f5160206114f95f395f51905f5288526101643590818a5287838760608160075afa92101616818360808160065afa165f5160206112b95f395f51905f5285525f5160206114395f395f51905f5288526101843590818a5287838760608160075afa92101616818360808160065afa165f5160206116395f395f51905f5285525f5160206117d95f395f51905f5288526101a43590818a5287838760608160075afa921016169160808160065afa16945f5160206112d95f395f51905f528352526101c43580955260608160075afa9210161660408a60808160065afa169851975198156107115760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f5160206113f95f395f51905f526101008401525f5160206116795f395f51905f526101208401525f5160206117b95f395f51905f526101408401525f5160206113d95f395f51905f526101608401525f5160206116b95f395f51905f526101808401525f5160206114595f395f51905f526101a08401525f5160206115d95f395f51905f526101c08401525f5160206113995f395f51905f526101e08401525f5160206114795f395f51905f526102008401525f5160206114b95f395f51905f526102208401526102408301526102608201525f5160206116f95f395f51905f526102808201525f5160206113195f395f51905f526102a08201525f5160206113b95f395f51905f526102c08201525f5160206115395f395f51905f526102e08201526040519283916106dd8484610c95565b8336843760085afa15908115610704575b506106f557005b631ff3747d60e21b5f5260045ffd5b60019150511415816106ee565b63a54f8e2760e01b5f5260045ffd5b346101c8576102603660031901126101c85736610104116101c85736610264116101c85760405160408101905f5160206115995f395f51905f52815260208101915f5160206114195f395f51905f5283525f5160206117795f395f51905f528152606082015f5160206112995f395f51905f5281525f5160206115b95f395f51905f52604061010435935f5160206115595f395f51905f52608087019580875284848460608160075afa911016838860808160065afa165f5160206116595f395f51905f5283525f5160206113395f395f51905f5286526001610124359182895286868660608160075afa9310161616838860808160065afa165f5160206113595f395f51905f5283525f5160206114995f395f51905f528652610144359081885285858560608160075afa92101616838860808160065afa165f5160206115f95f395f51905f5283525f5160206116d95f395f51905f528652610164359081885285858560608160075afa92101616838860808160065afa165f5160206115795f395f51905f5283525f5160206117395f395f51905f528652610184359081885285858560608160075afa92101616838860808160065afa165f5160206117595f395f51905f5283525f5160206117195f395f51905f5286526101a4359081885285858560608160075afa92101616838860808160065afa165f5160206116995f395f51905f5283525f5160206115195f395f51905f5286526101c4359081885285858560608160075afa92101616838860808160065afa165f5160206112f95f395f51905f5283525f5160206114f95f395f51905f5286526101e4359081885285858560608160075afa92101616838860808160065afa165f5160206112b95f395f51905f5283525f5160206114395f395f51905f528652610204359081885285858560608160075afa92101616838860808160065afa165f5160206116395f395f51905f5283525f5160206117d95f395f51905f528652610224359081885285858560608160075afa92101616838860808160065afa16945f5160206112d95f395f51905f528352526102443580955260608160075afa9210161660408260808160065afa169051915190156107115760405191610100600484375f5160206113f95f395f51905f526101008401525f5160206116795f395f51905f526101208401525f5160206117b95f395f51905f526101408401525f5160206113d95f395f51905f526101608401525f5160206116b95f395f51905f526101808401525f5160206114595f395f51905f526101a08401525f5160206115d95f395f51905f526101c08401525f5160206113995f395f51905f526101e08401525f5160206114795f395f51905f526102008401525f5160206114b95f395f51905f526102208401526102408301526102608201525f5160206116f95f395f51905f526102808201525f5160206113195f395f51905f526102a08201525f5160206113b95f395f51905f526102c08201525f5160206115395f395f51905f526102e08201526020816103008160085afa905116156106f557005b346101c8576101003660031901126101c85736610104116101c857604051610bc0608082610c95565b6080368237610bd3602435600435610cb8565b8152610be960843560a435604435606435610d59565b60208301526040820152610c0160e43560c435610cb8565b6060820152604051905f825b60048210610c1a57608084f35b6020806001928551815201930191019091610c0d565b346101c8575f3660031901126101c857807f23b50690255b6580e58c4f76addf9359f378856e884fec3bd3cc1e2c2960ecb360209252f35b9181601f840112156101c8578235916001600160401b0383116101c857602083818601950101116101c857565b601f909101601f19168101906001600160401b038211908210176101dc57604052565b905f5160206113795f395f51905f528210801590610d42575b6106f557811580610d3a575b610d3457610d015f5160206113795f395f51905f52600381858181800909086110ac565b818103610d1057505060011b90565b5f5160206113795f395f51905f52809106810306145f146106f557600190811b1790565b50505f90565b508015610cdd565b505f5160206113795f395f51905f52811015610cd1565b919093925f5160206113795f395f51905f528310801590610f2a575b8015610f13575b8015610efc575b6106f5578082868517171715610ef157908291610e545f5160206113795f395f51905f5280808080888180808f9d5f5160206116195f395f51905f528f839290839109099d8e0981848181800909085f5160206117995f395f51905f52089a09818c8181800909085f5160206112795f395f51905f520806810306945f5160206113795f395f51905f525f5160206114d95f395f51905f5281610e2e81808b800981878009086110ac565b8408095f5160206113795f395f51905f52610e4882611210565b800914159586916110cf565b929080821480610ee8575b15610e865750505050905f14610e7e5760ff60025b169060021b179190565b60ff5f610e74565b5f5160206113795f395f51905f52809106810306149182610ec9575b5050156106f55760019115610ec15760ff60025b169060021b17179190565b60ff5f610eb6565b5f5160206113795f395f51905f52919250819006810306145f80610ea2565b50838314610e5f565b50505090505f905f90565b505f5160206113795f395f51905f52811015610d83565b505f5160206113795f395f51905f52821015610d7c565b505f5160206113795f395f51905f52851015610d75565b8015610fa5578060011c915f5160206113795f395f51905f528310156106f557600180610f845f5160206113795f395f51905f52600381888181800909086110ac565b931614610f8d57565b905f5160206113795f395f51905f5280910681030690565b505f905f90565b8015806110a4575b611098578060021c92825f5160206113795f395f51905f528510801590611081575b6106f55784815f5160206113795f395f51905f5280808080808080805f5160206116195f395f51905f528161104b9d8d0909998a0981898181800909085f5160206112795f395f51905f520806810306936002808a16149509818a8181800909085f5160206117995f395f51905f52086110cf565b8092916001808296161461105d575050565b5f5160206113795f395f51905f528093945080929550809106810306930681030690565b505f5160206113795f395f51905f52811015610fd6565b50505f905f905f905f90565b508115610fb4565b906110b682611210565b915f5160206113795f395f51905f52838009036106f557565b915f5160206113795f395f51905f525f5160206114d95f395f51905f52816111149396949661110682808a8009818a8009086110ac565b90611204575b8608096110ac565b925f5160206113795f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206113795f395f51905f5260a083015260208260c08160055afa915191156106f5575f5160206113795f395f51905f528260019209036106f5575f5160206113795f395f51905f52908209925f5160206113795f395f51905f5280808087800906810306818780090814908115916111e5575b506106f557565b90505f5160206113795f395f51905f528084860960020914155f6111de565b8180910681030661110c565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206113795f395f51905f5260a083015260208260c08160055afa915191156106f55756fe2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775188f2571c23b709d0542877cedaaedbc09af3f1095b5cb48ac5662c30ed25faa1d857d91b85b0c8c6d9716db11cd192be5836fd9cbcfe063837f9e164fb78b4d20a0cad618979e31184b38f23dc0030268417f404b5f9b0fb7d592787e3c65ca01c5857c70b08da5fff9f73e52059d185e30af4b3341ae6162ccee6fc1ee56d62e21de06192fcd246480c70323b5e903097a93619ce8d8d265f49f0a3b5dd96a2e76a8a93dcf480ef880e8313677a302ec257a80f66459d868f405ef3e0ac03e0aaf4752b77ff1ea3bd00b63c34afd182592906efb458e6b275dfb10c331025930644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4706b4456f99199ef7dfd74ae98c901cb3f7d1524fabd359b53eab0694648f5c03145a089a87aa6bf0ed2feed32e0d8d11bd53bc95875eac664afc1e1405a93ffc295658bcd7efce611c8cca6eacea3434e670c901da592482b41c18591a71362b2a4fe04d736c9ae981ee3d15f4459323d52dd3e8c423bb2faf1c6b75ab2ec87f089b32a2b7b7ab40d15d170d4de21bd61c68351ec0df8383b83d652a9cfe92d9044fb50a9fce8547c1e51290bcd2e6485291f69d2a72bd9e83ce5ff3e86b338817dbb2ebf3d72af87c1b8845b089dd961878c51d41bc3067ff1bc0e174dbc88827680e18742bef6409428ce0911b0e4fe56e73c5508f9172147dbcf9f45670da084a617ed9f393f7546d3c4b8fc0d8e02780c70bae16ffc5ab794ec5fb34b2df251dff080a9b2afc5b42138a9a3ece7c4cff557f6830352c029f70ddb5ec6ed9183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea40ecc5143101d8fdb53c86baf2587a3a94a9eb12059a7ec8dd35674246407d1ce0f4a3264609bffb220277bd5c5748585383422d35ec84c7cf9db7b9a32fa81d702ba9a11443394343993b3be2baa5a15b206805afb61abd0385f7c9cf92dc5b414497f896c3ee9b661e26dc3ad45c87ad592cacdff44d2dff242307dc1bec71d1b8ea39027a31cf941c9a6e4ff37be54e44d080ac8fb2c8af1f35379f0a20ace0a7579b6b3cecb48d3204e8fb9014422588a249c580ec4f6aa60132c021e744b30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000012b97264996d88a555a18bf32e76688c90cb3748a54c0daa175dcf7f3c46ed0491587afd4e8645136111c6a2adc7168e9469bd656d0b716be20bab5a0ffdfab7030644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd441979761ec74aa626db838c9811d2ff5a68597649aa783527126424783869690d2fd0f595ed0cbfad279323dbe9bcb95e1ff2ef7bff194b77e0ace36d87ee97161e07ebc0816524fe15f7b5aaea95fc2366d9544b57cec839aca4661c3a2657e621b33b16714f741451f9eb40dd94f8b6f51ba123545eeec5f23a140efcda236f259cd4caeced3cc5d180b55a3bf0d8a0d6544c0c4f18f80c6ae08091e1e29c5220b25c007b530398b7db2db36562c7207ccf4ec5a02db6a6c2707d874ee9319b04e4e33da5e4464294ad72c16622c7a911b01d25e2030a3bb2c977d228c780c91e4ae8f747b50d21d120dee3c9d3b4acff3623795fc02581ab8a76910ffb83571496298096bc2d6b69e5f50725c14de14ac8fba820be2209b80383dd386fc07c1e914b363af7265b7e7fb6a8da8c3a06b5d2c9fc55e89957cdde00492e902acd0b88eb4f6d9e1b9b8b9a731017193091f5191eacaef8585f358428b993db15062b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50a6b29c65bca840df86710531a3062c2457f9346c87bc5024031ee176011a4cd14922534b8456360a25c78d538c7a22396d09461021070fdd30abe8316b9db6ea2646970667358221220aafc848e3787c5cf427204bcc96b82b8f6ba825e52ec15a92f93b52c4f8cb7fb64736f6c634300081c0033",
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

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x98ea1c08.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[11] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [11]*big.Int) error {
	var out []interface{}
	err := _DecryptCombineVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x98ea1c08.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[11] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [11]*big.Int) error {
	return _DecryptCombineVerifier.Contract.VerifyCompressedProof(&_DecryptCombineVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x98ea1c08.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[11] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [11]*big.Int) error {
	return _DecryptCombineVerifier.Contract.VerifyCompressedProof(&_DecryptCombineVerifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x8261a653.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[11] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, input [11]*big.Int) error {
	var out []interface{}
	err := _DecryptCombineVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x8261a653.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[11] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierSession) VerifyProof(proof [8]*big.Int, input [11]*big.Int) error {
	return _DecryptCombineVerifier.Contract.VerifyProof(&_DecryptCombineVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x8261a653.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[11] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerSession) VerifyProof(proof [8]*big.Int, input [11]*big.Int) error {
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
