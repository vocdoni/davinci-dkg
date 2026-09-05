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
	Bin: "0x60808060405234601557611b96908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80630ea14a3914610db1578063216df434146107a0578063233ace1114610766578063b8e72af6146106095763eb879eb814610051575f80fd5b3461060657610200366003190112610606576004356001600160401b03811161060457610082903690600401610e79565b3661020411610600576101006100989114610edd565b604051604081015f516020611a015f395f51905f52825260208201905f5160206119e15f395f51905f5282525f5160206118215f395f51905f528152606083015f5160206119a15f395f51905f5281525f5160206118c15f395f51905f526040602435935f5160206119415f395f51905f52608088019580875284848460608160075afa911016838960808160065afa165f5160206116e15f395f51905f5283525f5160206116615f395f51905f52865260016044359182895286868660608160075afa9310161616838960808160065afa165f516020611a215f395f51905f5283525f516020611b215f395f51905f5286526064359081885285858560608160075afa92101616838960808160065afa165f5160206118015f395f51905f5283525f5160206115215f395f51905f5286526084359081885285858560608160075afa92101616838960808160065afa165f5160206119615f395f51905f5283525f516020611b015f395f51905f52865260a4359081885285858560608160075afa92101616838960808160065afa165f5160206118e15f395f51905f5283525f5160206118a15f395f51905f52865260c4359081885285858560608160075afa92101616838960808160065afa165f516020611ac15f395f51905f5283525f516020611a815f395f51905f52865260e4359081885285858560608160075afa92101616838960808160065afa165f5160206117415f395f51905f5283525f5160206115815f395f51905f528652610104359081885285858560608160075afa92101616838960808160065afa165f516020611ae15f395f51905f5283525f5160206118815f395f51905f528652610124359081885285858560608160075afa92101616838960808160065afa165f5160206115e15f395f51905f5283525f5160206119815f395f51905f528652610144359081885285858560608160075afa92101616838960808160065afa165f5160206117c15f395f51905f5283525f5160206116415f395f51905f528652610164359081885285858560608160075afa92101616838960808160065afa165f5160206115415f395f51905f5283525f5160206117015f395f51905f528652610184359081885285858560608160075afa92101616838960808160065afa165f516020611a615f395f51905f5283525f5160206114e15f395f51905f5286526101a4359081885285858560608160075afa92101616838960808160065afa165f5160206117815f395f51905f5283525f5160206117e15f395f51905f5286526101c4359081885285858560608160075afa92101616838960808160065afa16945f5160206116c15f395f51905f528352526101e43580955260608160075afa9210161660408360808160065afa169151905191156105f1576101006040519384375f5160206115615f395f51905f526101008401525f516020611aa15f395f51905f526101208401525f5160206117215f395f51905f526101408401525f5160206116215f395f51905f526101608401525f5160206117615f395f51905f526101808401525f5160206119215f395f51905f526101a08401525f5160206116815f395f51905f526101c08401525f5160206118415f395f51905f526101e08401525f5160206115a15f395f51905f526102008401525f5160206118615f395f51905f526102208401526102408301526102608201525f5160206119c15f395f51905f526102808201525f5160206115c15f395f51905f526102a08201525f5160206119015f395f51905f526102c08201525f5160206116015f395f51905f526102e08201526020816103008160085afa905116156105e25780f35b631ff3747d60e21b8152600490fd5b63a54f8e2760e01b8452600484fd5b8280fd5b505b80fd5b5034610743576040366003190112610743576004356001600160401b0381116107435761063a903690600401610e79565b6024356001600160401b03811161074357610659903690600401610e79565b61010083036107575781016101e0828203126107435780601f83011215610743576101e06040519261068b8285610ea6565b8391810192831161074357905b82821061074757505050303b1561074357604051631d70f3d760e31b8152610200600482015261020481018390529283919083906102248401375f6102248484010152602482015f905b600f821061072957505050610224815f93601f80199101168101030181305afa801561071e57610710575080f35b61071c91505f90610ea6565b005b6040513d5f823e3d90fd5b8293506020809160019394518152019301910184926106e2565b5f80fd5b8135815260209182019101610698565b63236bd13760e01b5f5260045ffd5b34610743575f3660031901126107435760206040517f80205078ebcf34a4bb2e68928450e1afd388a0dfd8a0810e6fedc1ae5f855b578152f35b34610743576102603660031901126107435736608411610743573661026411610743576103006040516107d38282610ea6565b813682376107e26004356111a9565b6107f3602493929335604435611214565b919392906108026064356111a9565b9390926040519660408801965f516020611a015f395f51905f5289528860208101985f5160206119e15f395f51905f528a525f5160206118215f395f51905f5281525f5160206118c15f395f51905f52604060608401925f5160206119a15f395f51905f5284525f5160206119415f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f5160206116e15f395f51905f5285525f5160206116615f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f516020611a215f395f51905f5285525f516020611b215f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f5160206118015f395f51905f5285525f5160206115215f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f5160206119615f395f51905f5285525f516020611b015f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f5160206118e15f395f51905f5285525f5160206118a15f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f516020611ac15f395f51905f5285525f516020611a815f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f5160206117415f395f51905f5285525f5160206115815f395f51905f5288526101643590818a5287838760608160075afa92101616818360808160065afa165f516020611ae15f395f51905f5285525f5160206118815f395f51905f5288526101843590818a5287838760608160075afa92101616818360808160065afa165f5160206115e15f395f51905f5285525f5160206119815f395f51905f5288526101a43590818a5287838760608160075afa92101616818360808160065afa165f5160206117c15f395f51905f5285525f5160206116415f395f51905f5288526101c43590818a5287838760608160075afa92101616818360808160065afa165f5160206115415f395f51905f5285525f5160206117015f395f51905f5288526101e43590818a5287838760608160075afa92101616818360808160065afa165f516020611a615f395f51905f5285525f5160206114e15f395f51905f5288526102043590818a5287838760608160075afa92101616818360808160065afa165f5160206117815f395f51905f5285525f5160206117e15f395f51905f5288526102243590818a5287838760608160075afa921016169160808160065afa16945f5160206116c15f395f51905f528352526102443580955260608160075afa9210161660408a60808160065afa16985197519815610da25760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f5160206115615f395f51905f526101008401525f516020611aa15f395f51905f526101208401525f5160206117215f395f51905f526101408401525f5160206116215f395f51905f526101608401525f5160206117615f395f51905f526101808401525f5160206119215f395f51905f526101a08401525f5160206116815f395f51905f526101c08401525f5160206118415f395f51905f526101e08401525f5160206115a15f395f51905f526102008401525f5160206118615f395f51905f526102208401526102408301526102608201525f5160206119c15f395f51905f526102808201525f5160206115c15f395f51905f526102a08201525f5160206119015f395f51905f526102c08201525f5160206116015f395f51905f526102e0820152604051928391610d6e8484610ea6565b8336843760085afa15908115610d95575b50610d8657005b631ff3747d60e21b5f5260045ffd5b6001915051141581610d7f565b63a54f8e2760e01b5f5260045ffd5b34610743576020366003190112610743576004356001600160401b03811161074357610de1903690600401610e79565b610e4a608092610e0561010060405194610dfb8787610ea6565b8636873714610edd565b610e1460208201358235610f20565b8352610e318482013560a083013560408401356060850135610fc1565b6020850152604084015260c060e0820135910135610f20565b6060820152604051905f825b60048210610e6357505050f35b6020806001928551815201930191019091610e56565b9181601f84011215610743578235916001600160401b038311610743576020838186019501011161074357565b601f909101601f19168101906001600160401b03821190821017610ec957604052565b634e487b7160e01b5f52604160045260245ffd5b15610ee457565b60405162461bcd60e51b81526020600482015260146024820152730d2dcecc2d8d2c840e0e4dedecc40d8cadccee8d60631b6044820152606490fd5b905f5160206116a15f395f51905f528210801590610faa575b610d8657811580610fa2575b610f9c57610f695f5160206116a15f395f51905f5260038185818180090908611314565b818103610f7857505060011b90565b5f5160206116a15f395f51905f52809106810306145f14610d8657600190811b1790565b50505f90565b508015610f45565b505f5160206116a15f395f51905f52811015610f39565b919093925f5160206116a15f395f51905f528310801590611192575b801561117b575b8015611164575b610d86578082868517171715611159579082916110bc5f5160206116a15f395f51905f5280808080888180808f9d5f516020611a415f395f51905f528f839290839109099d8e0981848181800909085f516020611b415f395f51905f52089a09818c8181800909085f5160206115015f395f51905f520806810306945f5160206116a15f395f51905f525f5160206117a15f395f51905f528161109681808b80098187800908611314565b8408095f5160206116a15f395f51905f526110b082611478565b80091415958691611337565b929080821480611150575b156110ee5750505050905f146110e65760ff60025b169060021b179190565b60ff5f6110dc565b5f5160206116a15f395f51905f52809106810306149182611131575b505015610d8657600191156111295760ff60025b169060021b17179190565b60ff5f61111e565b5f5160206116a15f395f51905f52919250819006810306145f8061110a565b508383146110c7565b50505090505f905f90565b505f5160206116a15f395f51905f52811015610feb565b505f5160206116a15f395f51905f52821015610fe4565b505f5160206116a15f395f51905f52851015610fdd565b801561120d578060011c915f5160206116a15f395f51905f52831015610d86576001806111ec5f5160206116a15f395f51905f5260038188818180090908611314565b9316146111f557565b905f5160206116a15f395f51905f5280910681030690565b505f905f90565b80158061130c575b611300578060021c92825f5160206116a15f395f51905f5285108015906112e9575b610d865784815f5160206116a15f395f51905f5280808080808080805f516020611a415f395f51905f52816112b39d8d0909998a0981898181800909085f5160206115015f395f51905f520806810306936002808a16149509818a8181800909085f516020611b415f395f51905f5208611337565b809291600180829616146112c5575050565b5f5160206116a15f395f51905f528093945080929550809106810306930681030690565b505f5160206116a15f395f51905f5281101561123e565b50505f905f905f905f90565b50811561121c565b9061131e82611478565b915f5160206116a15f395f51905f5283800903610d8657565b915f5160206116a15f395f51905f525f5160206117a15f395f51905f528161137c9396949661136e82808a8009818a800908611314565b9061146c575b860809611314565b925f5160206116a15f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206116a15f395f51905f5260a083015260208260c08160055afa91519115610d86575f5160206116a15f395f51905f52826001920903610d86575f5160206116a15f395f51905f52908209925f5160206116a15f395f51905f52808080878009068103068187800908149081159161144d575b50610d8657565b90505f5160206116a15f395f51905f528084860960020914155f611446565b81809106810306611374565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206116a15f395f51905f5260a083015260208260c08160055afa91519115610d865756fe293a824b8bf093cc1167b080d2f135c5b9071805c2929d2495af1799a6bed8cb2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750e16af04a1005bc15f5310b7cfbcca5c693951ecec752ac95ac39d87be376a2906f3b4982f9e855f3159aeb33328ebb50133a0d9a295dbb5791617631febe9e719f55f2646e2cf0bbe615f02aadf78add018a45910d59ad0488934e94d5c249a09d96b38b3187438cd8a10851330f3c9bc5497ab9f5ce8c9112e98e8eb7e568b2f9875b8e3987b7a14a4861697fa330adb9c18f08b3080d58727ccb810471be003e86ea1c25e250db638cb350234ec530b0c2d069c57cb919cf2b7e1b5dc507d1a90273d4a68c6e90faa78a7b86265b4fb79e974e65e84837357d0980c1755651898b9d621e0f889cd125bb25874c1c0fe38915b403df821e8b1ceca7cd52ba70ffd18b1258e0714c7b996946a5f4753ad09ea63494d464cd9d213651dea583b180fbc7e6e962a0e2310f73ea2ae7d2eace003c974b2387835e22db60ab160db008add9c742285d20e32ddde3d657737baade833e6b2cf020d01bdaa04355257154bed76f5ea219826f7f29e675150ab62dc5acc19e9551d402a704cb553693330644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4709a2c850f5a758d4e2d67e5f552acd3fdeb4d2a359728ef38d4a20765fa0b5210316091e4d561de8c9f6c63a29d95082d4ebb2908f916a87801ae4b88f12edb52c2bd4146e3278d3e26e48f214e0d2652ae11e10e214d992d79eb5e71accc12100c3a515f796fc1da42393deeffc10f26978149d43a3371fbb1ee419a6a9cf2f29c0f8f9389912a42c60a4961dc99c57718bfd64eeb711b7762f4b9a174751ce0513b4aad7c9c98fea7cc7dba8fee4cb9270daa9e0b47e019058e837c1ca7f420bc1c5da37c089a72d452378ba05004202dd96c389fc125c8adb04a19e5fda89183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea425a5cec85c30e3eea695d02be565660e244bb48a065468adf30c27a2f851904c05dd3166413adc51bd95bfdac9605905e8f0ceee980be49b14521c1bc941200f2dbb12b7a55387dcd759fe70b9d7ff3aa3fc5696830efa276e179696ccc58c9210542138e0917ed6de96027057c5ec05760fc84fd8a4000ddb348efec0c9d31a0fd0fed94e724494f2814c39c1c8379e7d9282d6084dbf21f3738691176b5458051ca066490da14fa0445112ee2143d2626a6d11d340e313dcde98376e246ffc1dde1d0726a60c7eab6fece6b5ed7c58d51f62401c2d03b878f13a4bd6b9615914b5f9a7e4b2146071cece5a92213ca71ae23460a55221d49356fbdf3bdf7d3b30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000106231921fc17c4622d497db6712094e53c1d5a0f344901f1620855f9351613d52bc07576bff18ec7836df92e4c8b7042f8a809b34aef8699bca1175179deeefb2bdc4e4d5878651a1e8b95b391a4e029efaeeceba570afe6e70bb6c52e80569f2f8e11a3ad09f812997c312b3d9668f8f32f5d113e44ce59820db925f8e273b620147e8edf58e3e7a42b28b21d189a73299040da4d9ca8a717669f98ed021e5613a4a096a6d5b66d47cc34e7dbc83a3a8732697179cefe138f6aa64b1425330112c82895327837848ee927e4e5ba2622cf9d43858727fad429a11031e100564a2955dadfb142d0b2a2bc32f6171c159fb180d939a8d9717f0172dc6ffd438f4b119f468a0b10443032d2ff4d0de1b5734cc99fb875ad6a0cdacf57d2024b8f172f1d17083971edccb7e32656468ae99867e779a904479e99d0bf8f8abac1f05d1540e016119c0bd35d61d69c8b473b41181cceaa6fef4a463daadae690b2628d30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd441ec9be9c3306148695279c07837001d1e254da0068cf3f3c83f11143f140693421738d31d4ae47f92e070582a3d40443e4c979713e8c28aa774c595620edbb972c2900c4b15e3aa248962cdba31d62d503ec4f0393ced431d8dfa7d4d12cb38525296e1d47b8c97a63e7954d784d99af8c15d3de6bb024febdf117fdc6d70fe20642de04f0045a4146de95ca7b7865e82ff2013b6f8297441155b962ff6443b40db9e8ff2542b51c294c26a7b29f1351ed6c5c714a572f36ce4daa0dc033338c2eafdee6943729ade2147bbaf92fc9fc22db36b6b723980c10d36454d1c648442b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5a2646970667358221220e9591a44976f44394b31db4a101e7978dc80a755cfff8d10db8d8bb2a9f879c564736f6c634300081c0033",
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
