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
	Bin: "0x60808060405234601557611b96908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80630ea14a3914610db1578063216df434146107a0578063233ace1114610766578063b8e72af6146106095763eb879eb814610051575f80fd5b3461060657610200366003190112610606576004356001600160401b03811161060457610082903690600401610e79565b3661020411610600576101006100989114610edd565b604051604081015f516020611ae15f395f51905f52825260208201905f5160206117215f395f51905f5282525f5160206117615f395f51905f528152606083015f5160206118c15f395f51905f5281525f5160206118615f395f51905f526040602435935f5160206117e15f395f51905f52608088019580875284848460608160075afa911016838960808160065afa165f5160206119015f395f51905f5283525f5160206117c15f395f51905f52865260016044359182895286868660608160075afa9310161616838960808160065afa165f5160206115815f395f51905f5283525f516020611b215f395f51905f5286526064359081885285858560608160075afa92101616838960808160065afa165f5160206119a15f395f51905f5283525f516020611ac15f395f51905f5286526084359081885285858560608160075afa92101616838960808160065afa165f5160206115c15f395f51905f5283525f5160206117815f395f51905f52865260a4359081885285858560608160075afa92101616838960808160065afa165f5160206118e15f395f51905f5283525f5160206118815f395f51905f52865260c4359081885285858560608160075afa92101616838960808160065afa165f516020611a415f395f51905f5283525f5160206115215f395f51905f52865260e4359081885285858560608160075afa92101616838960808160065afa165f5160206115015f395f51905f5283525f5160206118015f395f51905f528652610104359081885285858560608160075afa92101616838960808160065afa165f5160206119c15f395f51905f5283525f5160206116e15f395f51905f528652610124359081885285858560608160075afa92101616838960808160065afa165f5160206115615f395f51905f5283525f5160206116415f395f51905f528652610144359081885285858560608160075afa92101616838960808160065afa165f5160206118415f395f51905f5283525f5160206115a15f395f51905f528652610164359081885285858560608160075afa92101616838960808160065afa165f5160206116015f395f51905f5283525f516020611a015f395f51905f528652610184359081885285858560608160075afa92101616838960808160065afa165f5160206116215f395f51905f5283525f5160206118215f395f51905f5286526101a4359081885285858560608160075afa92101616838960808160065afa165f5160206119e15f395f51905f5283525f5160206116c15f395f51905f5286526101c4359081885285858560608160075afa92101616838960808160065afa16945f5160206116815f395f51905f528352526101e43580955260608160075afa9210161660408360808160065afa169151905191156105f1576101006040519384375f516020611aa15f395f51905f526101008401525f5160206119815f395f51905f526101208401525f516020611b015f395f51905f526101408401525f5160206115e15f395f51905f526101608401525f516020611a815f395f51905f526101808401525f5160206119615f395f51905f526101a08401525f5160206119415f395f51905f526101c08401525f5160206118a15f395f51905f526101e08401525f516020611a615f395f51905f526102008401525f5160206115415f395f51905f526102208401526102408301526102608201525f5160206117015f395f51905f526102808201525f5160206117415f395f51905f526102a08201525f5160206116615f395f51905f526102c08201525f5160206119215f395f51905f526102e08201526020816103008160085afa905116156105e25780f35b631ff3747d60e21b8152600490fd5b63a54f8e2760e01b8452600484fd5b8280fd5b505b80fd5b5034610743576040366003190112610743576004356001600160401b0381116107435761063a903690600401610e79565b6024356001600160401b03811161074357610659903690600401610e79565b61010083036107575781016101e0828203126107435780601f83011215610743576101e06040519261068b8285610ea6565b8391810192831161074357905b82821061074757505050303b1561074357604051631d70f3d760e31b8152610200600482015261020481018390529283919083906102248401375f6102248484010152602482015f905b600f821061072957505050610224815f93601f80199101168101030181305afa801561071e57610710575080f35b61071c91505f90610ea6565b005b6040513d5f823e3d90fd5b8293506020809160019394518152019301910184926106e2565b5f80fd5b8135815260209182019101610698565b63236bd13760e01b5f5260045ffd5b34610743575f3660031901126107435760206040517fdb96fe2f3ef7ecd40c137f819812bd892d6a29f4046dd7337524571dec8c1d9d8152f35b34610743576102603660031901126107435736608411610743573661026411610743576103006040516107d38282610ea6565b813682376107e26004356111a9565b6107f3602493929335604435611214565b919392906108026064356111a9565b9390926040519660408801965f516020611ae15f395f51905f5289528860208101985f5160206117215f395f51905f528a525f5160206117615f395f51905f5281525f5160206118615f395f51905f52604060608401925f5160206118c15f395f51905f5284525f5160206117e15f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f5160206119015f395f51905f5285525f5160206117c15f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f5160206115815f395f51905f5285525f516020611b215f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f5160206119a15f395f51905f5285525f516020611ac15f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f5160206115c15f395f51905f5285525f5160206117815f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f5160206118e15f395f51905f5285525f5160206118815f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f516020611a415f395f51905f5285525f5160206115215f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f5160206115015f395f51905f5285525f5160206118015f395f51905f5288526101643590818a5287838760608160075afa92101616818360808160065afa165f5160206119c15f395f51905f5285525f5160206116e15f395f51905f5288526101843590818a5287838760608160075afa92101616818360808160065afa165f5160206115615f395f51905f5285525f5160206116415f395f51905f5288526101a43590818a5287838760608160075afa92101616818360808160065afa165f5160206118415f395f51905f5285525f5160206115a15f395f51905f5288526101c43590818a5287838760608160075afa92101616818360808160065afa165f5160206116015f395f51905f5285525f516020611a015f395f51905f5288526101e43590818a5287838760608160075afa92101616818360808160065afa165f5160206116215f395f51905f5285525f5160206118215f395f51905f5288526102043590818a5287838760608160075afa92101616818360808160065afa165f5160206119e15f395f51905f5285525f5160206116c15f395f51905f5288526102243590818a5287838760608160075afa921016169160808160065afa16945f5160206116815f395f51905f528352526102443580955260608160075afa9210161660408a60808160065afa16985197519815610da25760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f516020611aa15f395f51905f526101008401525f5160206119815f395f51905f526101208401525f516020611b015f395f51905f526101408401525f5160206115e15f395f51905f526101608401525f516020611a815f395f51905f526101808401525f5160206119615f395f51905f526101a08401525f5160206119415f395f51905f526101c08401525f5160206118a15f395f51905f526101e08401525f516020611a615f395f51905f526102008401525f5160206115415f395f51905f526102208401526102408301526102608201525f5160206117015f395f51905f526102808201525f5160206117415f395f51905f526102a08201525f5160206116615f395f51905f526102c08201525f5160206119215f395f51905f526102e0820152604051928391610d6e8484610ea6565b8336843760085afa15908115610d95575b50610d8657005b631ff3747d60e21b5f5260045ffd5b6001915051141581610d7f565b63a54f8e2760e01b5f5260045ffd5b34610743576020366003190112610743576004356001600160401b03811161074357610de1903690600401610e79565b610e4a608092610e0561010060405194610dfb8787610ea6565b8636873714610edd565b610e1460208201358235610f20565b8352610e318482013560a083013560408401356060850135610fc1565b6020850152604084015260c060e0820135910135610f20565b6060820152604051905f825b60048210610e6357505050f35b6020806001928551815201930191019091610e56565b9181601f84011215610743578235916001600160401b038311610743576020838186019501011161074357565b601f909101601f19168101906001600160401b03821190821017610ec957604052565b634e487b7160e01b5f52604160045260245ffd5b15610ee457565b60405162461bcd60e51b81526020600482015260146024820152730d2dcecc2d8d2c840e0e4dedecc40d8cadccee8d60631b6044820152606490fd5b905f5160206116a15f395f51905f528210801590610faa575b610d8657811580610fa2575b610f9c57610f695f5160206116a15f395f51905f5260038185818180090908611314565b818103610f7857505060011b90565b5f5160206116a15f395f51905f52809106810306145f14610d8657600190811b1790565b50505f90565b508015610f45565b505f5160206116a15f395f51905f52811015610f39565b919093925f5160206116a15f395f51905f528310801590611192575b801561117b575b8015611164575b610d86578082868517171715611159579082916110bc5f5160206116a15f395f51905f5280808080888180808f9d5f516020611a215f395f51905f528f839290839109099d8e0981848181800909085f516020611b415f395f51905f52089a09818c8181800909085f5160206114e15f395f51905f520806810306945f5160206116a15f395f51905f525f5160206117a15f395f51905f528161109681808b80098187800908611314565b8408095f5160206116a15f395f51905f526110b082611478565b80091415958691611337565b929080821480611150575b156110ee5750505050905f146110e65760ff60025b169060021b179190565b60ff5f6110dc565b5f5160206116a15f395f51905f52809106810306149182611131575b505015610d8657600191156111295760ff60025b169060021b17179190565b60ff5f61111e565b5f5160206116a15f395f51905f52919250819006810306145f8061110a565b508383146110c7565b50505090505f905f90565b505f5160206116a15f395f51905f52811015610feb565b505f5160206116a15f395f51905f52821015610fe4565b505f5160206116a15f395f51905f52851015610fdd565b801561120d578060011c915f5160206116a15f395f51905f52831015610d86576001806111ec5f5160206116a15f395f51905f5260038188818180090908611314565b9316146111f557565b905f5160206116a15f395f51905f5280910681030690565b505f905f90565b80158061130c575b611300578060021c92825f5160206116a15f395f51905f5285108015906112e9575b610d865784815f5160206116a15f395f51905f5280808080808080805f516020611a215f395f51905f52816112b39d8d0909998a0981898181800909085f5160206114e15f395f51905f520806810306936002808a16149509818a8181800909085f516020611b415f395f51905f5208611337565b809291600180829616146112c5575050565b5f5160206116a15f395f51905f528093945080929550809106810306930681030690565b505f5160206116a15f395f51905f5281101561123e565b50505f905f905f905f90565b50811561121c565b9061131e82611478565b915f5160206116a15f395f51905f5283800903610d8657565b915f5160206116a15f395f51905f525f5160206117a15f395f51905f528161137c9396949661136e82808a8009818a800908611314565b9061146c575b860809611314565b925f5160206116a15f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206116a15f395f51905f5260a083015260208260c08160055afa91519115610d86575f5160206116a15f395f51905f52826001920903610d86575f5160206116a15f395f51905f52908209925f5160206116a15f395f51905f52808080878009068103068187800908149081159161144d575b50610d8657565b90505f5160206116a15f395f51905f528084860960020914155f611446565b81809106810306611374565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206116a15f395f51905f5260a083015260208260c08160055afa91519115610d865756fe2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775022a50c85f717445baaeb2f9e263a3281b8e7ecbafaf70af6edc717e712292af08d0ec72c688b1bc85cc54bb82b1275c247f11d4b7a2845bd2c50384980be97216871b0c36414fe8e4de7baa8e7845b60ec4ff049dc0ea0a662b19c8a40ff7082b7b33a011172560a34d40e0aabe7858dc73be62a7f614d977e3e57cf818976a1689557159cba091a47d14a01b02cb526f485f505333ed221104075dcf25df9c2efb9a6207037782ebaa83b66a371a8f0264b7f481c3bf332f6b65407992074112b00e25cdd01bca8296e1ca608b9c78cd1bf8c35fd00e59b54df5ea43cc49cf008a38cd9db40a705700996a4856dd1f05676641b562823bc5c907e05422303907d580bf0c58f140b20f8a2595342f6a9637aea0d2ee6634cc70519f6154029d160515b065b12506cf7909614a20d79dc10d07321eb33cf1a116585190645b5d2d89841dcc78f733fdc1002ddbdf41503b0e6cb04b723402ff64e2ccf0697870016375cb43e225ce3549af8878662300d30a03bf095b6a545e0cea403f210dd01a95b67ae9f1fef70866d9440128800f373b775f6da90e45608d84f7f7894ebe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47197a01f0c7927e85fe00c105fdfccc418c113231839eaf421ff73e07e6c158881f4d88dc01062fa4b21d39c1a47c5bfaa59b7c9258280e36362041664594d1ea0aabe70148fd2f845cf80a4596540a63c10022711c3023884e1ec6fdae1e13bc20f6babdf98f3656f9500a2ed919e12ca10bba91b8b16403a966c4639a57757e12c5091ecb5b8a5063bb3bf359ed40768174f36830ca692187cf9864fa92c28901d8dee32f27c40f0e397be809d723ab8aef03f4748312c42b4ec35438cd99a2134217e3df25f7b5bee0ccb2ba5f81d75e8a62a92feb05b8019c7c2f853bc826183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea41ba0e404b822117a699b4e8bcfdac65546c0fec1e7c624675f514952c272cea704f4525b367750d0655855cf2e0482027d4f460cd754d81de8e408d291d913e41c0c108537ede436ef9336a194df9e2ae6f9b64e227fbd5c565620eb0bbc027b09358a6ad01fb65615061276a6437866ef941cebe5b3f3fd6f3edea20e11d72715e9d9d9af81ef98881f39e85cf1905e0a64b7698f39a49ee5e9aedbb638c6c630644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000116a4287a01abc11fd0f7fe2e64cd19e1dc090f761150f4d63be0113d961c27d41f631a9ac3e2176184ad63a236a8a08836d24fda651f5300087f7f8edd213e03167a4075df392570a1bda150f338e3df22869657d159ad57da23821c861272ad1da7ce5a1fbbeb3fa75406ee8323aa0afa14df26194cf2c9fd99b3e41b4e125a075e179956e3e3cfaecbc597adcb48e9bfa586080524495c6b773664b5a232f527f39f98eb4b5311ce814abf8a735684e457e792148cfb10558dbdc84d9faf380f55aaee59fa5028003aa8aae0fdfa80525d302f869b28ca96f0e58d5421613e08e1f343c319ff142bc7d1b0ee08f4b085f982456fd9aba4664c0938b3a448d4021203cc48d9db85ecfd5bcba1ea09e40d2e4065ebf18366dc1a5a1e32fd306629cf3a694e79d562e6918978899befd8033bf3d1068fcb1ce6ba24e79fbc51922377157b9f9f22705b5c54b535adf34ee7f5998267f1a7e3388e8171f2dc7b541777aedb2d80a68beac05562267a85f3db214f7939a944ea27f4f80a03be7ba61fc7256cab73a49b4933ea6eebf43cdd315bb99f5969fc194da69bc57ac6b9bc30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd440abe0680afd9abe690cba963d09b880a4a9c8da6385fce83323d539a620e174019d0e3b692dc1f86134a10c2d9c06a473dda87650f3bafb07942717872099e1715a389f86750e8773241958e9eb9c98698c2736a1103425cf0f69e8b77478cea089fb3de7c149bc01aea29e0a16a92895f78db567a553982067deae8b7ad7cef234951c729677b310b997c2b43dc95d525e39e8b08693037fe4d8e36f760f60909e1a3ed1e7e4c634f33e68b8f1afbfee6f369c6cb0df8eef9bfbf257d86b849282ad3bfd46c14f3b41c596456268c684f8391494664a6a3385e38aca8dd14a02728f41786857bccf20abbfe76a303cd4d5e460719a415a7879bf3f49c6b775a2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5a2646970667358221220c551d8974de7baa6f7f4f65faa68e2b862add14338e878a1ccd0d821d24c469964736f6c634300081c0033",
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
