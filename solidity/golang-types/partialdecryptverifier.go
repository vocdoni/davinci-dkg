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
	Bin: "0x60808060405234601557611b96908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80630ea14a3914610db1578063216df434146107a0578063233ace1114610766578063b8e72af6146106095763eb879eb814610051575f80fd5b3461060657610200366003190112610606576004356001600160401b03811161060457610082903690600401610e79565b3661020411610600576101006100989114610edd565b604051604081015f5160206117215f395f51905f52825260208201905f5160206117415f395f51905f5282525f5160206119215f395f51905f528152606083015f5160206118a15f395f51905f5281525f5160206118c15f395f51905f526040602435935f5160206117615f395f51905f52608088019580875284848460608160075afa911016838960808160065afa165f5160206116615f395f51905f5283525f516020611a015f395f51905f52865260016044359182895286868660608160075afa9310161616838960808160065afa165f516020611ae15f395f51905f5283525f516020611a615f395f51905f5286526064359081885285858560608160075afa92101616838960808160065afa165f5160206115415f395f51905f5283525f5160206116015f395f51905f5286526084359081885285858560608160075afa92101616838960808160065afa165f5160206117a15f395f51905f5283525f5160206118215f395f51905f52865260a4359081885285858560608160075afa92101616838960808160065afa165f5160206119a15f395f51905f5283525f5160206115c15f395f51905f52865260c4359081885285858560608160075afa92101616838960808160065afa165f5160206118e15f395f51905f5283525f516020611ac15f395f51905f52865260e4359081885285858560608160075afa92101616838960808160065afa165f516020611a415f395f51905f5283525f5160206114e15f395f51905f528652610104359081885285858560608160075afa92101616838960808160065afa165f5160206115215f395f51905f5283525f516020611a815f395f51905f528652610124359081885285858560608160075afa92101616838960808160065afa165f5160206119c15f395f51905f5283525f5160206119615f395f51905f528652610144359081885285858560608160075afa92101616838960808160065afa165f516020611aa15f395f51905f5283525f5160206115815f395f51905f528652610164359081885285858560608160075afa92101616838960808160065afa165f516020611b015f395f51905f5283525f5160206116a15f395f51905f528652610184359081885285858560608160075afa92101616838960808160065afa165f5160206117c15f395f51905f5283525f5160206119415f395f51905f5286526101a4359081885285858560608160075afa92101616838960808160065afa165f5160206116e15f395f51905f5283525f5160206118615f395f51905f5286526101c4359081885285858560608160075afa92101616838960808160065afa16945f5160206118815f395f51905f528352526101e43580955260608160075afa9210161660408360808160065afa169151905191156105f1576101006040519384375f5160206119e15f395f51905f526101008401525f5160206118415f395f51905f526101208401525f5160206117015f395f51905f526101408401525f5160206115a15f395f51905f526101608401525f5160206117815f395f51905f526101808401525f5160206116415f395f51905f526101a08401525f5160206119815f395f51905f526101c08401525f516020611b415f395f51905f526101e08401525f5160206116215f395f51905f526102008401525f5160206115615f395f51905f526102208401526102408301526102608201525f5160206115e15f395f51905f526102808201525f5160206117e15f395f51905f526102a08201525f5160206116c15f395f51905f526102c08201525f5160206119015f395f51905f526102e08201526020816103008160085afa905116156105e25780f35b631ff3747d60e21b8152600490fd5b63a54f8e2760e01b8452600484fd5b8280fd5b505b80fd5b5034610743576040366003190112610743576004356001600160401b0381116107435761063a903690600401610e79565b6024356001600160401b03811161074357610659903690600401610e79565b61010083036107575781016101e0828203126107435780601f83011215610743576101e06040519261068b8285610ea6565b8391810192831161074357905b82821061074757505050303b1561074357604051631d70f3d760e31b8152610200600482015261020481018390529283919083906102248401375f6102248484010152602482015f905b600f821061072957505050610224815f93601f80199101168101030181305afa801561071e57610710575080f35b61071c91505f90610ea6565b005b6040513d5f823e3d90fd5b8293506020809160019394518152019301910184926106e2565b5f80fd5b8135815260209182019101610698565b63236bd13760e01b5f5260045ffd5b34610743575f3660031901126107435760206040517fd80bfa3d4d43e86204180d8884a3b1bc5c60b5f5832974f3867d11eafb22f8658152f35b34610743576102603660031901126107435736608411610743573661026411610743576103006040516107d38282610ea6565b813682376107e26004356111a9565b6107f3602493929335604435611214565b919392906108026064356111a9565b9390926040519660408801965f5160206117215f395f51905f5289528860208101985f5160206117415f395f51905f528a525f5160206119215f395f51905f5281525f5160206118c15f395f51905f52604060608401925f5160206118a15f395f51905f5284525f5160206117615f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f5160206116615f395f51905f5285525f516020611a015f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f516020611ae15f395f51905f5285525f516020611a615f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f5160206115415f395f51905f5285525f5160206116015f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f5160206117a15f395f51905f5285525f5160206118215f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f5160206119a15f395f51905f5285525f5160206115c15f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f5160206118e15f395f51905f5285525f516020611ac15f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f516020611a415f395f51905f5285525f5160206114e15f395f51905f5288526101643590818a5287838760608160075afa92101616818360808160065afa165f5160206115215f395f51905f5285525f516020611a815f395f51905f5288526101843590818a5287838760608160075afa92101616818360808160065afa165f5160206119c15f395f51905f5285525f5160206119615f395f51905f5288526101a43590818a5287838760608160075afa92101616818360808160065afa165f516020611aa15f395f51905f5285525f5160206115815f395f51905f5288526101c43590818a5287838760608160075afa92101616818360808160065afa165f516020611b015f395f51905f5285525f5160206116a15f395f51905f5288526101e43590818a5287838760608160075afa92101616818360808160065afa165f5160206117c15f395f51905f5285525f5160206119415f395f51905f5288526102043590818a5287838760608160075afa92101616818360808160065afa165f5160206116e15f395f51905f5285525f5160206118615f395f51905f5288526102243590818a5287838760608160075afa921016169160808160065afa16945f5160206118815f395f51905f528352526102443580955260608160075afa9210161660408a60808160065afa16985197519815610da25760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f5160206119e15f395f51905f526101008401525f5160206118415f395f51905f526101208401525f5160206117015f395f51905f526101408401525f5160206115a15f395f51905f526101608401525f5160206117815f395f51905f526101808401525f5160206116415f395f51905f526101a08401525f5160206119815f395f51905f526101c08401525f516020611b415f395f51905f526101e08401525f5160206116215f395f51905f526102008401525f5160206115615f395f51905f526102208401526102408301526102608201525f5160206115e15f395f51905f526102808201525f5160206117e15f395f51905f526102a08201525f5160206116c15f395f51905f526102c08201525f5160206119015f395f51905f526102e0820152604051928391610d6e8484610ea6565b8336843760085afa15908115610d95575b50610d8657005b631ff3747d60e21b5f5260045ffd5b6001915051141581610d7f565b63a54f8e2760e01b5f5260045ffd5b34610743576020366003190112610743576004356001600160401b03811161074357610de1903690600401610e79565b610e4a608092610e0561010060405194610dfb8787610ea6565b8636873714610edd565b610e1460208201358235610f20565b8352610e318482013560a083013560408401356060850135610fc1565b6020850152604084015260c060e0820135910135610f20565b6060820152604051905f825b60048210610e6357505050f35b6020806001928551815201930191019091610e56565b9181601f84011215610743578235916001600160401b038311610743576020838186019501011161074357565b601f909101601f19168101906001600160401b03821190821017610ec957604052565b634e487b7160e01b5f52604160045260245ffd5b15610ee457565b60405162461bcd60e51b81526020600482015260146024820152730d2dcecc2d8d2c840e0e4dedecc40d8cadccee8d60631b6044820152606490fd5b905f5160206116815f395f51905f528210801590610faa575b610d8657811580610fa2575b610f9c57610f695f5160206116815f395f51905f5260038185818180090908611314565b818103610f7857505060011b90565b5f5160206116815f395f51905f52809106810306145f14610d8657600190811b1790565b50505f90565b508015610f45565b505f5160206116815f395f51905f52811015610f39565b919093925f5160206116815f395f51905f528310801590611192575b801561117b575b8015611164575b610d86578082868517171715611159579082916110bc5f5160206116815f395f51905f5280808080888180808f9d5f516020611a215f395f51905f528f839290839109099d8e0981848181800909085f516020611b215f395f51905f52089a09818c8181800909085f5160206115015f395f51905f520806810306945f5160206116815f395f51905f525f5160206118015f395f51905f528161109681808b80098187800908611314565b8408095f5160206116815f395f51905f526110b082611478565b80091415958691611337565b929080821480611150575b156110ee5750505050905f146110e65760ff60025b169060021b179190565b60ff5f6110dc565b5f5160206116815f395f51905f52809106810306149182611131575b505015610d8657600191156111295760ff60025b169060021b17179190565b60ff5f61111e565b5f5160206116815f395f51905f52919250819006810306145f8061110a565b508383146110c7565b50505090505f905f90565b505f5160206116815f395f51905f52811015610feb565b505f5160206116815f395f51905f52821015610fe4565b505f5160206116815f395f51905f52851015610fdd565b801561120d578060011c915f5160206116815f395f51905f52831015610d86576001806111ec5f5160206116815f395f51905f5260038188818180090908611314565b9316146111f557565b905f5160206116815f395f51905f5280910681030690565b505f905f90565b80158061130c575b611300578060021c92825f5160206116815f395f51905f5285108015906112e9575b610d865784815f5160206116815f395f51905f5280808080808080805f516020611a215f395f51905f52816112b39d8d0909998a0981898181800909085f5160206115015f395f51905f520806810306936002808a16149509818a8181800909085f516020611b215f395f51905f5208611337565b809291600180829616146112c5575050565b5f5160206116815f395f51905f528093945080929550809106810306930681030690565b505f5160206116815f395f51905f5281101561123e565b50505f905f905f905f90565b50811561121c565b9061131e82611478565b915f5160206116815f395f51905f5283800903610d8657565b915f5160206116815f395f51905f525f5160206118015f395f51905f528161137c9396949661136e82808a8009818a800908611314565b9061146c575b860809611314565b925f5160206116815f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206116815f395f51905f5260a083015260208260c08160055afa91519115610d86575f5160206116815f395f51905f52826001920903610d86575f5160206116815f395f51905f52908209925f5160206116815f395f51905f52808080878009068103068187800908149081159161144d575b50610d8657565b90505f5160206116815f395f51905f528084860960020914155f611446565b81809106810306611374565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206116815f395f51905f5260a083015260208260c08160055afa91519115610d865756fe14577d8420ff36fc8eee4b721753058bdfd4b5861ca62120ba1a3ea47d5773c52fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77513e54385b704698254b3b301dd0b45d64f86d6cf326bbf812587dfa7bbc79692112b0ce3fa1041092bc41f134e35f02dbe9d11dca0daabfa4cb675f25245beb11e47bf50c85eb5c4003ab6cf3bdc70018bb28feb4a83eac183824e48e11008330eddfefb9250a78399fcc051b4440397fc3419569d414dbd9ca718edde5f737c034460fa0338bec74be8cb35dc412696961f4024246a1dd386a24d59d4955c9e244f37f7ce71e61f1e24b34c65196c9d84251659011c7fabb1f26ff8e299265f2c38045a57e8265dbcbd98e695e34f6a803ebfc983100a870b8cea8124dfd1d908f63df9c8d64fb831d59d23334ac7995c44dc9c75418730866abe2817c75ac71a65e360ad3a06c6a4e0963b969f941acbda13ebef61f1b32cf276f78a2828d3000169eb0d85c4fc79a9384baaaa17aaf7957574b722b20f5b65dc98bca840a011e504ea6db7582bb0d56abc33dac2ff917074a44593a7a25c560988e64895f530644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd471805d1c53a44f07572aa7215c7fd92d7c0cd3208bfbd2d27a35e850127b0f0400286f644ed731b29967ac73d7daec66dbb9e81169859337f1cf1b17f4bb16ffb0b2152fb028e1615584a519deb650080fc0fe88d836ada3ac55038f23a63efa20ebe36953aa34d91a54108f8b91af7d99f4624122dc19acdefa837c6dbe41fe4123131df165e18481410e8d09a797ee6a412e7090843a69ddff00efadf260bf919699f9fc60893756af37cbc83259a37385276b93f1b2856879a1ec5840b759904ed0a89abf8e6e6268abe3986246ee925b4196aa14888b97aa607bae7fe55ca03d74b6dcefb4757ab9baece3e5a19164255d3464cd05a8f46dc62edb39ad55b0dc0488c63b5e63dc43f5b569eee1fa60b53a0eff6a53632aa0624975fdf61c714e2eb60f9db47ffc007c8f0aa5fda144cac7589a7e5b9610d84fc27de9978ae0d36f40e7a5504e48703e816026ef00a9482a71655fe86912d29cbdc7997debf183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea425bc8de223714287c946cc7cb33657e878b83c68de03a976c4182d4611a104701990964555e5f1dd3426d25ba10b7f15ba87c4d0e8b9c7cb2be6904cfd13f02f0069d018ac41e2c0242642af490f9f015763c2a30b5bacc294e01cf2d11d4031159e33cc7ba238665cb19b33328e3cc774e1061c7014de6cd86acb7ba90e6e0b2a38604a3cfa058b8ba603f66b5d8fd17180d009b7fa86c5cf999b775f92228e30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000111db54ea695c3553df0d3f8cf33d653cab158d1d9817a0a5103704dbc3aff00b2cea1b67aa24971ad5dac1e7ecc46cbc76a6d4a7f5db5f70d9c3cca47ea820f72a4c9eeeba60cddc00861f5b344ef4c5cd8fd601b3497ded74897035e77d56340df87a9c1a4a8742ef6d5ccdf4637911b87076aaba9a365c13e109b815fc7fbc08b2d8f4f3d9918acc9085038b839beb65340481d4acc419e7e66382cc88ac632297c6f5ce29304bdb7a24c78420dc76bc0c452c38bcb01f07616a044c840cf522cd9c41c84106852bb2770e5c57e616faf887e538956f455528c0ee73a394971f26f3d1780799740d48ca52abd4632f54fc8cdb70076d321f2388625b70f4da116b5a66d5a9857abc12b02a35332f2895069eff169cf4422a6ebd93598a2de21cb7371d65d5714f10d2da8563a642fb1973b66d357fb9275122f698b92d186530644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd441d028669e1fc054d7808db3baa899ec5cd8dc3ca5ab626ed676faba4033025402eed8b95cc332df638408182f1eca545b5c808e0d7e22024b9cd970d8f440d0722d05a89410a8f29373a250f65d32f1909e6da1b897bac2564dd5d53225ffc7d0c9459345ef245eb9077cabeb69461101a80645ae36233877e6cc03da36a31fe1cee09c509393e82a02ed38f409282902835b751be5bd60e38f55d3ca5301e3a064c2890d22a97582d21238809af20ba2b7a9a369a09b12cc8f052376209add9016bbf24e24a204e69ee97917856c4701bf0eb100806a50b9536ee9f076d89d22b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5266093e25b7f604f695b948f7fcdd058d564ba09d9e45fc56ae072b4c318304aa264697066735822122040d7bdf2579366af9cd6f7e132e1c7dea0cb099d94a98207031ba2ead238141464736f6c634300081c0033",
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
