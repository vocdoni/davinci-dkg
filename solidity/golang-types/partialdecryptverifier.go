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

// PartialDecryptVerifierMetaData contains all meta data concerning the PartialDecryptVerifier contract.
var PartialDecryptVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[15]\",\"internalType\":\"uint256[15]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"uint256[15]\",\"internalType\":\"uint256[15]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidProofEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x60808060405234601557611b96908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80630ea14a3914610db1578063216df434146107a0578063233ace1114610766578063b8e72af6146106095763eb879eb814610051575f80fd5b3461060657610200366003190112610606576004356001600160401b03811161060457610082903690600401610e79565b3661020411610600576101006100989114610edd565b604051604081015f5160206117a15f395f51905f52825260208201905f5160206115615f395f51905f5282525f5160206115e15f395f51905f528152606083015f5160206118e15f395f51905f5281525f5160206118415f395f51905f526040602435935f5160206118815f395f51905f52608088019580875284848460608160075afa911016838960808160065afa165f5160206119c15f395f51905f5283525f516020611b215f395f51905f52865260016044359182895286868660608160075afa9310161616838960808160065afa165f5160206116015f395f51905f5283525f516020611ac15f395f51905f5286526064359081885285858560608160075afa92101616838960808160065afa165f5160206119815f395f51905f5283525f5160206119015f395f51905f5286526084359081885285858560608160075afa92101616838960808160065afa165f5160206119a15f395f51905f5283525f5160206115415f395f51905f52865260a4359081885285858560608160075afa92101616838960808160065afa165f5160206118215f395f51905f5283525f516020611a815f395f51905f52865260c4359081885285858560608160075afa92101616838960808160065afa165f5160206115815f395f51905f5283525f516020611a015f395f51905f52865260e4359081885285858560608160075afa92101616838960808160065afa165f5160206117215f395f51905f5283525f5160206118015f395f51905f528652610104359081885285858560608160075afa92101616838960808160065afa165f5160206119615f395f51905f5283525f5160206115215f395f51905f528652610124359081885285858560608160075afa92101616838960808160065afa165f5160206117c15f395f51905f5283525f516020611a215f395f51905f528652610144359081885285858560608160075afa92101616838960808160065afa165f5160206116c15f395f51905f5283525f516020611aa15f395f51905f528652610164359081885285858560608160075afa92101616838960808160065afa165f516020611a415f395f51905f5283525f516020611ae15f395f51905f528652610184359081885285858560608160075afa92101616838960808160065afa165f5160206115015f395f51905f5283525f5160206116415f395f51905f5286526101a4359081885285858560608160075afa92101616838960808160065afa165f5160206116215f395f51905f5283525f5160206115c15f395f51905f5286526101c4359081885285858560608160075afa92101616838960808160065afa16945f5160206117615f395f51905f528352526101e43580955260608160075afa9210161660408360808160065afa169151905191156105f1576101006040519384375f5160206116815f395f51905f526101008401525f5160206116a15f395f51905f526101208401525f5160206116e15f395f51905f526101408401525f516020611a615f395f51905f526101608401525f5160206119215f395f51905f526101808401525f5160206117e15f395f51905f526101a08401525f5160206118c15f395f51905f526101c08401525f5160206115a15f395f51905f526101e08401525f516020611b415f395f51905f526102008401525f5160206117015f395f51905f526102208401526102408301526102608201525f5160206118615f395f51905f526102808201525f5160206117415f395f51905f526102a08201525f5160206118a15f395f51905f526102c08201525f5160206119415f395f51905f526102e08201526020816103008160085afa905116156105e25780f35b631ff3747d60e21b8152600490fd5b63a54f8e2760e01b8452600484fd5b8280fd5b505b80fd5b5034610743576040366003190112610743576004356001600160401b0381116107435761063a903690600401610e79565b6024356001600160401b03811161074357610659903690600401610e79565b61010083036107575781016101e0828203126107435780601f83011215610743576101e06040519261068b8285610ea6565b8391810192831161074357905b82821061074757505050303b1561074357604051631d70f3d760e31b8152610200600482015261020481018390529283919083906102248401375f6102248484010152602482015f905b600f821061072957505050610224815f93601f80199101168101030181305afa801561071e57610710575080f35b61071c91505f90610ea6565b005b6040513d5f823e3d90fd5b8293506020809160019394518152019301910184926106e2565b5f80fd5b8135815260209182019101610698565b63236bd13760e01b5f5260045ffd5b34610743575f3660031901126107435760206040517f0e827380dafca1282092a4afd188d67cbdaced28146a590e77a992415b05e94f8152f35b34610743576102603660031901126107435736608411610743573661026411610743576103006040516107d38282610ea6565b813682376107e26004356111a9565b6107f3602493929335604435611214565b919392906108026064356111a9565b9390926040519660408801965f5160206117a15f395f51905f5289528860208101985f5160206115615f395f51905f528a525f5160206115e15f395f51905f5281525f5160206118415f395f51905f52604060608401925f5160206118e15f395f51905f5284525f5160206118815f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f5160206119c15f395f51905f5285525f516020611b215f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f5160206116015f395f51905f5285525f516020611ac15f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f5160206119815f395f51905f5285525f5160206119015f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f5160206119a15f395f51905f5285525f5160206115415f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f5160206118215f395f51905f5285525f516020611a815f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f5160206115815f395f51905f5285525f516020611a015f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f5160206117215f395f51905f5285525f5160206118015f395f51905f5288526101643590818a5287838760608160075afa92101616818360808160065afa165f5160206119615f395f51905f5285525f5160206115215f395f51905f5288526101843590818a5287838760608160075afa92101616818360808160065afa165f5160206117c15f395f51905f5285525f516020611a215f395f51905f5288526101a43590818a5287838760608160075afa92101616818360808160065afa165f5160206116c15f395f51905f5285525f516020611aa15f395f51905f5288526101c43590818a5287838760608160075afa92101616818360808160065afa165f516020611a415f395f51905f5285525f516020611ae15f395f51905f5288526101e43590818a5287838760608160075afa92101616818360808160065afa165f5160206115015f395f51905f5285525f5160206116415f395f51905f5288526102043590818a5287838760608160075afa92101616818360808160065afa165f5160206116215f395f51905f5285525f5160206115c15f395f51905f5288526102243590818a5287838760608160075afa921016169160808160065afa16945f5160206117615f395f51905f528352526102443580955260608160075afa9210161660408a60808160065afa16985197519815610da25760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f5160206116815f395f51905f526101008401525f5160206116a15f395f51905f526101208401525f5160206116e15f395f51905f526101408401525f516020611a615f395f51905f526101608401525f5160206119215f395f51905f526101808401525f5160206117e15f395f51905f526101a08401525f5160206118c15f395f51905f526101c08401525f5160206115a15f395f51905f526101e08401525f516020611b415f395f51905f526102008401525f5160206117015f395f51905f526102208401526102408301526102608201525f5160206118615f395f51905f526102808201525f5160206117415f395f51905f526102a08201525f5160206118a15f395f51905f526102c08201525f5160206119415f395f51905f526102e0820152604051928391610d6e8484610ea6565b8336843760085afa15908115610d95575b50610d8657005b631ff3747d60e21b5f5260045ffd5b6001915051141581610d7f565b63a54f8e2760e01b5f5260045ffd5b34610743576020366003190112610743576004356001600160401b03811161074357610de1903690600401610e79565b610e4a608092610e0561010060405194610dfb8787610ea6565b8636873714610edd565b610e1460208201358235610f20565b8352610e318482013560a083013560408401356060850135610fc1565b6020850152604084015260c060e0820135910135610f20565b6060820152604051905f825b60048210610e6357505050f35b6020806001928551815201930191019091610e56565b9181601f84011215610743578235916001600160401b038311610743576020838186019501011161074357565b601f909101601f19168101906001600160401b03821190821017610ec957604052565b634e487b7160e01b5f52604160045260245ffd5b15610ee457565b60405162461bcd60e51b81526020600482015260146024820152730d2dcecc2d8d2c840e0e4dedecc40d8cadccee8d60631b6044820152606490fd5b905f5160206116615f395f51905f528210801590610faa575b610d8657811580610fa2575b610f9c57610f695f5160206116615f395f51905f5260038185818180090908611314565b818103610f7857505060011b90565b5f5160206116615f395f51905f52809106810306145f14610d8657600190811b1790565b50505f90565b508015610f45565b505f5160206116615f395f51905f52811015610f39565b919093925f5160206116615f395f51905f528310801590611192575b801561117b575b8015611164575b610d86578082868517171715611159579082916110bc5f5160206116615f395f51905f5280808080888180808f9d5f5160206119e15f395f51905f528f839290839109099d8e0981848181800909085f516020611b015f395f51905f52089a09818c8181800909085f5160206114e15f395f51905f520806810306945f5160206116615f395f51905f525f5160206117815f395f51905f528161109681808b80098187800908611314565b8408095f5160206116615f395f51905f526110b082611478565b80091415958691611337565b929080821480611150575b156110ee5750505050905f146110e65760ff60025b169060021b179190565b60ff5f6110dc565b5f5160206116615f395f51905f52809106810306149182611131575b505015610d8657600191156111295760ff60025b169060021b17179190565b60ff5f61111e565b5f5160206116615f395f51905f52919250819006810306145f8061110a565b508383146110c7565b50505090505f905f90565b505f5160206116615f395f51905f52811015610feb565b505f5160206116615f395f51905f52821015610fe4565b505f5160206116615f395f51905f52851015610fdd565b801561120d578060011c915f5160206116615f395f51905f52831015610d86576001806111ec5f5160206116615f395f51905f5260038188818180090908611314565b9316146111f557565b905f5160206116615f395f51905f5280910681030690565b505f905f90565b80158061130c575b611300578060021c92825f5160206116615f395f51905f5285108015906112e9575b610d865784815f5160206116615f395f51905f5280808080808080805f5160206119e15f395f51905f52816112b39d8d0909998a0981898181800909085f5160206114e15f395f51905f520806810306936002808a16149509818a8181800909085f516020611b015f395f51905f5208611337565b809291600180829616146112c5575050565b5f5160206116615f395f51905f528093945080929550809106810306930681030690565b505f5160206116615f395f51905f5281101561123e565b50505f905f905f905f90565b50811561121c565b9061131e82611478565b915f5160206116615f395f51905f5283800903610d8657565b915f5160206116615f395f51905f525f5160206117815f395f51905f528161137c9396949661136e82808a8009818a800908611314565b9061146c575b860809611314565b925f5160206116615f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206116615f395f51905f5260a083015260208260c08160055afa91519115610d86575f5160206116615f395f51905f52826001920903610d86575f5160206116615f395f51905f52908209925f5160206116615f395f51905f52808080878009068103068187800908149081159161144d575b50610d8657565b90505f5160206116615f395f51905f528084860960020914155f611446565b81809106810306611374565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206116615f395f51905f5260a083015260208260c08160055afa91519115610d865756fe2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77526cf294544dec821665a2ebcb7ab99a708a8f8f4a1db7a4369c77fa17a4d077a09ad4fdf1b709400f903a080822009c0fa864f67b079b4d8e08a6be9ca08633616b7cc0a506dbb804ffae48e97e3f68beb2fe3016434a1a3e9d36eabe3777a3421f440d43df703e9bc737d85da67c3336bd4ffe0028c03aa7afac3abf78e47e621e4f964d34720a025afe9bc0583708fd0110bd01308f3fe29daf641ccfd85e60bfaf3a0165f45c08b3ad2080bcfe01e416b56bd04c664ddc37975883e9a240203e79b4ef6a1a1c2956ba3ea9aa60cc77ed0a4cfb69555685b3c66e58a0fccab1a666ad477d5aa0b209a5a8dc431ad52991b7fe30c095eb8e09a7d6eef55cd922a3ccfed7b9b79d8ecb1dc070fc5cf6ba3680b28fd5f18296645090eba8af4b316f78337e339e2aa31dcccfb510608f84b3a7358a78b43192986c63e0450339d1014d8784a60b18dcf9804fe7b515e6a4b8c819f663a236186b332617b8f3d8b30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47220144d19daf2aeadcfe4f961dc6c8715a95376a752826eb847d155bdb35de1716fd085b1fff23d0ffa13bd1d93365e79cf3e42fe423fe0c3732ded6cbbc7a9522bbf01c13f4bb1ef9105439decedd0dafc956ec9f8425b64ffc29e03cbdf4631cb491e62f569f293e2938f290c21821b8d370753c04f4edb9b6238a4c47d2e71098d4cf579ba6c4e8d727bd059e0d7a3d00fa4922d091f6934ae64063080a8a2e001d0413e9a741ecdb1185e3d29ef0492a67f0740d616d350e63b5eb5681440253b4080d57b5b17c9ddbd053abafcb35e63752100d241701b8fd3c2e269c5a0f1f5183d7eeb656aaeff6d7149809d17098eab073b393a1d8cad77b53552242183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea411e2867613b1389593520f3c68be25bb368fe51c5f711e1390317c8d8451add4135b7b14e3e1918a5a7d2f8c72a7a5220011af2df602000fbf0ee29c9d885c4724e0fa24b0d0baf9c6398c376f76456c9a99d6ae3f7e4c257fc2f1e2cb236c7a00563415eb7639b690bc4f224fe77565b22ab87962144333dd04a5df4311bd342529d2d7caaf1bd61e150be4ba6da4cb8575785831b2d0f1a3574198ba435f5930644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000010e1f477a03be9546f22d41da204a099eaffc86c57380bf6a3144a13e371c395d01aaf4709ada6f019d77721fc380140e1d929a2d11d53a4eeffb020c04c7ab441e8c0e042b99288b6b337433ca8740f0d7da6462ebd982f589f263ee9fbcddee1a04a01b188ed32d81f6e92580adcea351bec22af5a803cd0370e0c818364c250d527200839ec1024adb20eef16e2c0ba66728a731e1d92d97c25b1436ac954006b00c2203491cfe6aed5745b2fac7e2628656f409f403b7c5446454a54d1fbd2c486c9f12e1714130a941bd599920ef00f86403277930422ae2f5182ae32e212dc0f8668cbd954d1f69480e82ac850110bfd8583db384a5962a9ef3fbeb5f0d0fe745e8c4f9394478bdb67fe1d9cc6a47090206d0093439b9ee41296034d2b0036897281257a972df325736bdd5d88cb7a3f7eadb3c0708f89bb2a9aeac28972a98b597f804669f81f21c773deca92d0bf08c67a19b2ae1427a163ba84f23092299d29c063ea0c225975041dfb812b8ea1293df0f3a7119fe7935713b450b2130644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44171aa078afd98a2c15f328b1e61d4c3a7e1ef5cc6e872ef61881106f901b550810d54f16a054e3952dfc6ada0c5cf282f0208fa7d162999331f354bf2eecbb071d0075b5fc7a473dd5c554db8807abbf7661f750f3d32c7649acc48bdd596bc90d1c8d7c9f65a01727d924880705b8c64982581129645fd401ccdceb062bd6810e318fbb5bf92f979365cceaf35b9eb6d412dc7e69cea0e77654dea6738c2afa02cc808f359e32d9f2ac77e58528689f30da12ecb5f4b14af9f65495eb0530bd2d6d94809f7b2cd4f5015749988945349f04fe0c5be736b0c3440efc3fe59f1615ed833989d1d3288da34b8536a6da57cf6d8b707250d532d1fc8e8547e1e56e2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e528986945f930fa7249baa5ba6185077706c52425a37d2331501dff8ffd8d68f42c897e89a758a787bd0dd39d731fdc8ce31e6df779f2dd32617128d13792fd1aa264697066735822122095c60d54d12242875f368586e91145c3ed62123cf2b90b86055177f5b7b13b1d64736f6c634300081c0033",
}

// PartialDecryptVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use PartialDecryptVerifierMetaData.ABI instead.
var PartialDecryptVerifierABI = PartialDecryptVerifierMetaData.ABI

// PartialDecryptVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PartialDecryptVerifierMetaData.Bin instead.
var PartialDecryptVerifierBin = PartialDecryptVerifierMetaData.Bin

// DeployPartialDecryptVerifier deploys a new Ethereum contract, binding an instance of PartialDecryptVerifier to it.
func DeployPartialDecryptVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PartialDecryptVerifier, error) {
	parsed, err := PartialDecryptVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PartialDecryptVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PartialDecryptVerifier{PartialDecryptVerifierCaller: PartialDecryptVerifierCaller{contract: contract}, PartialDecryptVerifierTransactor: PartialDecryptVerifierTransactor{contract: contract}, PartialDecryptVerifierFilterer: PartialDecryptVerifierFilterer{contract: contract}}, nil
}

// PartialDecryptVerifier is an auto generated Go binding around an Ethereum contract.
type PartialDecryptVerifier struct {
	PartialDecryptVerifierCaller     // Read-only binding to the contract
	PartialDecryptVerifierTransactor // Write-only binding to the contract
	PartialDecryptVerifierFilterer   // Log filterer for contract events
}

// PartialDecryptVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type PartialDecryptVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartialDecryptVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PartialDecryptVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartialDecryptVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PartialDecryptVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartialDecryptVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PartialDecryptVerifierSession struct {
	Contract     *PartialDecryptVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PartialDecryptVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PartialDecryptVerifierCallerSession struct {
	Contract *PartialDecryptVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// PartialDecryptVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PartialDecryptVerifierTransactorSession struct {
	Contract     *PartialDecryptVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// PartialDecryptVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type PartialDecryptVerifierRaw struct {
	Contract *PartialDecryptVerifier // Generic contract binding to access the raw methods on
}

// PartialDecryptVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PartialDecryptVerifierCallerRaw struct {
	Contract *PartialDecryptVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// PartialDecryptVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PartialDecryptVerifierTransactorRaw struct {
	Contract *PartialDecryptVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPartialDecryptVerifier creates a new instance of PartialDecryptVerifier, bound to a specific deployed contract.
func NewPartialDecryptVerifier(address common.Address, backend bind.ContractBackend) (*PartialDecryptVerifier, error) {
	contract, err := bindPartialDecryptVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PartialDecryptVerifier{PartialDecryptVerifierCaller: PartialDecryptVerifierCaller{contract: contract}, PartialDecryptVerifierTransactor: PartialDecryptVerifierTransactor{contract: contract}, PartialDecryptVerifierFilterer: PartialDecryptVerifierFilterer{contract: contract}}, nil
}

// NewPartialDecryptVerifierCaller creates a new read-only instance of PartialDecryptVerifier, bound to a specific deployed contract.
func NewPartialDecryptVerifierCaller(address common.Address, caller bind.ContractCaller) (*PartialDecryptVerifierCaller, error) {
	contract, err := bindPartialDecryptVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PartialDecryptVerifierCaller{contract: contract}, nil
}

// NewPartialDecryptVerifierTransactor creates a new write-only instance of PartialDecryptVerifier, bound to a specific deployed contract.
func NewPartialDecryptVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*PartialDecryptVerifierTransactor, error) {
	contract, err := bindPartialDecryptVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PartialDecryptVerifierTransactor{contract: contract}, nil
}

// NewPartialDecryptVerifierFilterer creates a new log filterer instance of PartialDecryptVerifier, bound to a specific deployed contract.
func NewPartialDecryptVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*PartialDecryptVerifierFilterer, error) {
	contract, err := bindPartialDecryptVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PartialDecryptVerifierFilterer{contract: contract}, nil
}

// bindPartialDecryptVerifier binds a generic wrapper to an already deployed contract.
func bindPartialDecryptVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PartialDecryptVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PartialDecryptVerifier *PartialDecryptVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PartialDecryptVerifier.Contract.PartialDecryptVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PartialDecryptVerifier *PartialDecryptVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartialDecryptVerifier.Contract.PartialDecryptVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PartialDecryptVerifier *PartialDecryptVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PartialDecryptVerifier.Contract.PartialDecryptVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PartialDecryptVerifier *PartialDecryptVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PartialDecryptVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PartialDecryptVerifier *PartialDecryptVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartialDecryptVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PartialDecryptVerifier *PartialDecryptVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PartialDecryptVerifier.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed)
func (_PartialDecryptVerifier *PartialDecryptVerifierCaller) CompressProof(opts *bind.CallOpts, proof []byte) ([4]*big.Int, error) {
	var out []interface{}
	err := _PartialDecryptVerifier.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed)
func (_PartialDecryptVerifier *PartialDecryptVerifierSession) CompressProof(proof []byte) ([4]*big.Int, error) {
	return _PartialDecryptVerifier.Contract.CompressProof(&_PartialDecryptVerifier.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed)
func (_PartialDecryptVerifier *PartialDecryptVerifierCallerSession) CompressProof(proof []byte) ([4]*big.Int, error) {
	return _PartialDecryptVerifier.Contract.CompressProof(&_PartialDecryptVerifier.CallOpts, proof)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_PartialDecryptVerifier *PartialDecryptVerifierCaller) ProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PartialDecryptVerifier.contract.Call(opts, &out, "provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_PartialDecryptVerifier *PartialDecryptVerifierSession) ProvingKeyHash() ([32]byte, error) {
	return _PartialDecryptVerifier.Contract.ProvingKeyHash(&_PartialDecryptVerifier.CallOpts)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_PartialDecryptVerifier *PartialDecryptVerifierCallerSession) ProvingKeyHash() ([32]byte, error) {
	return _PartialDecryptVerifier.Contract.ProvingKeyHash(&_PartialDecryptVerifier.CallOpts)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x216df434.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[15] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [15]*big.Int) error {
	var out []interface{}
	err := _PartialDecryptVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x216df434.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[15] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [15]*big.Int) error {
	return _PartialDecryptVerifier.Contract.VerifyCompressedProof(&_PartialDecryptVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x216df434.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[15] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [15]*big.Int) error {
	return _PartialDecryptVerifier.Contract.VerifyCompressedProof(&_PartialDecryptVerifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCaller) VerifyProof(opts *bind.CallOpts, proof []byte, input []byte) error {
	var out []interface{}
	err := _PartialDecryptVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierSession) VerifyProof(proof []byte, input []byte) error {
	return _PartialDecryptVerifier.Contract.VerifyProof(&_PartialDecryptVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCallerSession) VerifyProof(proof []byte, input []byte) error {
	return _PartialDecryptVerifier.Contract.VerifyProof(&_PartialDecryptVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xeb879eb8.
//
// Solidity: function verifyProof(bytes proof, uint256[15] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCaller) VerifyProof0(opts *bind.CallOpts, proof []byte, input [15]*big.Int) error {
	var out []interface{}
	err := _PartialDecryptVerifier.contract.Call(opts, &out, "verifyProof0", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xeb879eb8.
//
// Solidity: function verifyProof(bytes proof, uint256[15] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierSession) VerifyProof0(proof []byte, input [15]*big.Int) error {
	return _PartialDecryptVerifier.Contract.VerifyProof0(&_PartialDecryptVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xeb879eb8.
//
// Solidity: function verifyProof(bytes proof, uint256[15] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCallerSession) VerifyProof0(proof []byte, input [15]*big.Int) error {
	return _PartialDecryptVerifier.Contract.VerifyProof0(&_PartialDecryptVerifier.CallOpts, proof, input)
}
