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
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[16]\",\"internalType\":\"uint256[16]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"input\",\"type\":\"uint256[16]\",\"internalType\":\"uint256[16]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidProofEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x60808060405234601557611c14908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace1114610ed65750806344f6369214610e3d5780634605cb8914610868578063b8e72af6146106ad5763da3496ab14610055575f80fd5b346106aa576102803660031901126106aa57366084116106aa5736610284116106aa57604051906103006100898184610f3b565b803684376100986004356111e7565b6100a9602495929535604435611252565b919392906100b86064356111e7565b9390926040519660408801965f51602061185f5f395f51905f5289528860208101985f5160206115bf5f395f51905f528a525f5160206119ff5f395f51905f5281525f5160206118df5f395f51905f52604060608401925f516020611a5f5f395f51905f5284525f516020611b9f5f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f51602061163f5f395f51905f5285525f51602061179f5f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f5160206119df5f395f51905f5285525f516020611bbf5f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f5160206118ff5f395f51905f5285525f51602061169f5f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f5160206116bf5f395f51905f5285525f516020611a9f5f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f516020611adf5f395f51905f5285525f51602061177f5f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f516020611b1f5f395f51905f5285525f51602061181f5f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f516020611a7f5f395f51905f5285525f516020611a3f5f395f51905f5288526101643590818a5287838760608160075afa92101616818360808160065afa165f51602061161f5f395f51905f5285525f51602061171f5f395f51905f5288526101843590818a5287838760608160075afa92101616818360808160065afa165f51602061197f5f395f51905f5285525f51602061173f5f395f51905f5288526101a43590818a5287838760608160075afa92101616818360808160065afa165f51602061157f5f395f51905f5285525f5160206115df5f395f51905f5288526101c43590818a5287838760608160075afa92101616818360808160065afa165f51602061195f5f395f51905f5285525f516020611abf5f395f51905f5288526101e43590818a5287838760608160075afa92101616818360808160065afa165f5160206115ff5f395f51905f5285525f51602061153f5f395f51905f5288526102043590818a5287838760608160075afa92101616818360808160065afa165f516020611b3f5f395f51905f5285525f51602061175f5f395f51905f5288526102243590818a5287838760608160075afa92101616818360808160065afa165f51602061187f5f395f51905f5285525f51602061191f5f395f51905f5288526102443590818a5287838760608160075afa921016169160808160065afa16945f5160206117bf5f395f51905f528352526102643580955260608160075afa9210161660408a60808160065afa1698519751981561069b5760209a9b9c8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f516020611b5f5f395f51905f526101008401525f51602061165f5f395f51905f526101208401525f516020611aff5f395f51905f526101408401525f5160206116df5f395f51905f526101608401525f51602061189f5f395f51905f526101808401525f5160206119bf5f395f51905f526101a08401525f5160206118bf5f395f51905f526101c08401525f51602061183f5f395f51905f526101e08401525f51602061199f5f395f51905f526102008401525f51602061159f5f395f51905f526102208401526102408301526102608201525f51602061155f5f395f51905f526102808201525f51602061167f5f395f51905f526102a08201525f51602061193f5f395f51905f526102c08201525f5160206117df5f395f51905f526102e08201526040519283916106668484610f3b565b8336843760085afa1590811561068e575b5061067f5780f35b631ff3747d60e21b8152600490fd5b600191505114155f610677565b63a54f8e2760e01b8c5260048cfd5b80fd5b5034610821576040366003190112610821576004356001600160401b038111610821576106de903690600401610f0e565b6024356001600160401b038111610821576106fd903690600401610f0e565b90916101008103610859578301610100848203126108215780601f85011215610821576040519361073061010086610f3b565b8490610100810192831161082157905b828210610849575050508101610200828203126108215780601f83011215610821576040519161077261020084610f3b565b8290610200810192831161082157905b82821061082557505050303b1561082157604051634605cb8960e01b8152915f600484015b6008821061080b5750505061010482015f905b601082106107f5575050505f8161030481305afa80156107ea576107dc575080f35b6107e891505f90610f3b565b005b6040513d5f823e3d90fd5b60208060019285518152019301910190916107ba565b60208060019285518152019301910190916107a7565b5f80fd5b8135815260209182019101610782565b634e487b7160e01b5f52604160045260245ffd5b8135815260209182019101610740565b63236bd13760e01b5f5260045ffd5b34610821576103003660031901126108215736610104116108215736610304116108215760405160408101905f51602061185f5f395f51905f52815260208101915f5160206115bf5f395f51905f5283525f5160206119ff5f395f51905f528152606082015f516020611a5f5f395f51905f5281525f5160206118df5f395f51905f52604061010435935f516020611b9f5f395f51905f52608087019580875284848460608160075afa911016838860808160065afa165f51602061163f5f395f51905f5283525f51602061179f5f395f51905f5286526001610124359182895286868660608160075afa9310161616838860808160065afa165f5160206119df5f395f51905f5283525f516020611bbf5f395f51905f528652610144359081885285858560608160075afa92101616838860808160065afa165f5160206118ff5f395f51905f5283525f51602061169f5f395f51905f528652610164359081885285858560608160075afa92101616838860808160065afa165f5160206116bf5f395f51905f5283525f516020611a9f5f395f51905f528652610184359081885285858560608160075afa92101616838860808160065afa165f516020611adf5f395f51905f5283525f51602061177f5f395f51905f5286526101a4359081885285858560608160075afa92101616838860808160065afa165f516020611b1f5f395f51905f5283525f51602061181f5f395f51905f5286526101c4359081885285858560608160075afa92101616838860808160065afa165f516020611a7f5f395f51905f5283525f516020611a3f5f395f51905f5286526101e4359081885285858560608160075afa92101616838860808160065afa165f51602061161f5f395f51905f5283525f51602061171f5f395f51905f528652610204359081885285858560608160075afa92101616838860808160065afa165f51602061197f5f395f51905f5283525f51602061173f5f395f51905f528652610224359081885285858560608160075afa92101616838860808160065afa165f51602061157f5f395f51905f5283525f5160206115df5f395f51905f528652610244359081885285858560608160075afa92101616838860808160065afa165f51602061195f5f395f51905f5283525f516020611abf5f395f51905f528652610264359081885285858560608160075afa92101616838860808160065afa165f5160206115ff5f395f51905f5283525f51602061153f5f395f51905f528652610284359081885285858560608160075afa92101616838860808160065afa165f516020611b3f5f395f51905f5283525f51602061175f5f395f51905f5286526102a4359081885285858560608160075afa92101616838860808160065afa165f51602061187f5f395f51905f5283525f51602061191f5f395f51905f5286526102c4359081885285858560608160075afa92101616838860808160065afa16945f5160206117bf5f395f51905f528352526102e43580955260608160075afa9210161660408260808160065afa16905191519015610e2e5760405191610100600484375f516020611b5f5f395f51905f526101008401525f51602061165f5f395f51905f526101208401525f516020611aff5f395f51905f526101408401525f5160206116df5f395f51905f526101608401525f51602061189f5f395f51905f526101808401525f5160206119bf5f395f51905f526101a08401525f5160206118bf5f395f51905f526101c08401525f51602061183f5f395f51905f526101e08401525f51602061199f5f395f51905f526102008401525f51602061159f5f395f51905f526102208401526102408301526102608201525f51602061155f5f395f51905f526102808201525f51602061167f5f395f51905f526102a08201525f51602061193f5f395f51905f526102c08201525f5160206117df5f395f51905f526102e08201526020816103008160085afa90511615610e1f57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b346108215761010036600319011261082157366101041161082157604051610e66608082610f3b565b6080368237610e79602435600435610f5e565b8152610e8f60843560a435604435606435610fff565b60208301526040820152610ea760e43560c435610f5e565b6060820152604051905f825b60048210610ec057608084f35b6020806001928551815201930191019091610eb3565b34610821575f36600319011261082157807f74e759dfc2477a128e4d830e7161f32d02f0cc8582b825d29a115c58201b7c0e60209252f35b9181601f84011215610821578235916001600160401b038311610821576020838186019501011161082157565b601f909101601f19168101906001600160401b0382119082101761083557604052565b905f5160206116ff5f395f51905f528210801590610fe8575b610e1f57811580610fe0575b610fda57610fa75f5160206116ff5f395f51905f5260038185818180090908611352565b818103610fb657505060011b90565b5f5160206116ff5f395f51905f52809106810306145f14610e1f57600190811b1790565b50505f90565b508015610f83565b505f5160206116ff5f395f51905f52811015610f77565b919093925f5160206116ff5f395f51905f5283108015906111d0575b80156111b9575b80156111a2575b610e1f578082868517171715611197579082916110fa5f5160206116ff5f395f51905f5280808080888180808f9d5f516020611a1f5f395f51905f528f839290839109099d8e0981848181800909085f516020611b7f5f395f51905f52089a09818c8181800909085f51602061151f5f395f51905f520806810306945f5160206116ff5f395f51905f525f5160206117ff5f395f51905f52816110d481808b80098187800908611352565b8408095f5160206116ff5f395f51905f526110ee826114b6565b80091415958691611375565b92908082148061118e575b1561112c5750505050905f146111245760ff60025b169060021b179190565b60ff5f61111a565b5f5160206116ff5f395f51905f5280910681030614918261116f575b505015610e1f57600191156111675760ff60025b169060021b17179190565b60ff5f61115c565b5f5160206116ff5f395f51905f52919250819006810306145f80611148565b50838314611105565b50505090505f905f90565b505f5160206116ff5f395f51905f52811015611029565b505f5160206116ff5f395f51905f52821015611022565b505f5160206116ff5f395f51905f5285101561101b565b801561124b578060011c915f5160206116ff5f395f51905f52831015610e1f5760018061122a5f5160206116ff5f395f51905f5260038188818180090908611352565b93161461123357565b905f5160206116ff5f395f51905f5280910681030690565b505f905f90565b80158061134a575b61133e578060021c92825f5160206116ff5f395f51905f528510801590611327575b610e1f5784815f5160206116ff5f395f51905f5280808080808080805f516020611a1f5f395f51905f52816112f19d8d0909998a0981898181800909085f51602061151f5f395f51905f520806810306936002808a16149509818a8181800909085f516020611b7f5f395f51905f5208611375565b80929160018082961614611303575050565b5f5160206116ff5f395f51905f528093945080929550809106810306930681030690565b505f5160206116ff5f395f51905f5281101561127c565b50505f905f905f905f90565b50811561125a565b9061135c826114b6565b915f5160206116ff5f395f51905f5283800903610e1f57565b915f5160206116ff5f395f51905f525f5160206117ff5f395f51905f52816113ba939694966113ac82808a8009818a800908611352565b906114aa575b860809611352565b925f5160206116ff5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206116ff5f395f51905f5260a083015260208260c08160055afa91519115610e1f575f5160206116ff5f395f51905f52826001920903610e1f575f5160206116ff5f395f51905f52908209925f5160206116ff5f395f51905f52808080878009068103068187800908149081159161148b575b50610e1f57565b90505f5160206116ff5f395f51905f528084860960020914155f611484565b818091068103066113b2565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206116ff5f395f51905f5260a083015260208260c08160055afa91519115610e1f5756fe2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77515958e9af29d92aeb5bc6d3371d00547e45b12dd4df8a96e44c34607ab348e70082c8dbfbf300944bd650da6eb8cd3eb5c72b9a7380b5455a2ffeb2e12cd66a1089a39537845d2fb9e3a1ac6e092c75d128b26a209f47716a4079f33be2d2a290edfab50053420d7a1d331ddb218d24311482f614c83b19f21e9cd92b5e7283b215ff54be8a978781a22bb1686c3b2e52344a55d37311ad2e8f30981375c7e4d0008f9c70469d3949e9f8a473b084cc4ee23cd6edab0313a417fe10af33d040a2b319c6d0b0d9ef6375a19c5cc0b1e39734f1c67ef3cf9d0083ef08cecab25a90cc0f4eb020523c9b74a8291332b080bac18976406cd649595eaf186c2ef33dd207f03dba15604a1a68a60674fd5a6db75c65e2381e1b46e17683d5e5f7f5187110cee34ea35479dddccf9f0572c5a15a013078b4ca5317fbfc78d7f45f27cf01bda1b59cc781b345b3ec5c9fb35180e599ccab1bce72e7ce7fb879168bcdfa7294bb0a1f5d94bf7817f4dbb7bebb91d71031b56ceb292680dd2e5012ffba46f0d1f2f54f65ae02f4e166a4cf177e5380d9f950c4bcc76bbe981dd803e6962ed13850560514e8deee38fd5de59c92a0be7536810aab197ebef27a737e6d382de30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd472b43b3cf6c295f76653ef868f921aeb9d918151a9ce2471cfbeebc68d888a39426a35ee4ee61b8799b149126a0311449767f3305d3fdb33bfe5a770efc5182ca2380743fd9b4da1881335fadfa97b87cca25f9ca19469cb73e179c0351a22ed513bb70dfe337483f1ab3cb35b4ec4375be88fb376062be64852e16b2d8f1343811102d14caee0b6137d2efa79659a8325f231792e2574ad274b9273b745d2cff2eae54517385d9e55279d7c2a905fce76cd0f9b75fe59782e36d704b4221601e2fabcc6fd5b3f5cf7b59b0484c78daf788dead458e049ff0796e8837913fe426183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea428522074d8e6edc62c75f70b4ee89565c21624e5d0fef377d482cf34da5961750ec77f90b88472421ebecae6164e6dff3bfae0158bac7bff374af62f270792f906adbb7e5d3f151bd2a212ac55fad6900d37af38f833f15d165dc2df222f340c1865b2edf24f2b642eed40b6a830d6dad214ff553d675874616ac070f7adbbd110a86718c50fd5bfd9d7cf4d374a21e4898816f5b40b394cd5c0e1892104680003610a25502eb8ba77758dfcaa95ba14d5999daf45e3d388986f31bc350767f130644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001051cb2d6e88560e52fcc19a56fe7c613ecee5934e34ed07eb58da45d57b3aa3619888d052a8c8b18a96c68f532787138b8053eee400a1026094a8c81de43d1031bd7079c355bb7b9293ae9bf11d0a8060628f6fe9c18c710f4e60fb51e12ce8b14db6fd72442643b78ff97593bd4ef37548d6e8a2537204eda2cfc56aa0ec05a1bd1bcd7c3b9b7fa2add4dca1828f5fc48eec558f92593cce0aebc95b90860b40c70140b00fdaaf43e75818cd186d0d9f12eb734f7ee81f047e0365e365b3be51af118a7ab1b302a40f38137fee418a751ff21dab4ba2895cceb7f842a703cd51387191d24e862d1620da8764b00ae86908293c4d74f9642eca4e4b9e0d3e42f151592f740f47bacf0dad9d3f259b06582cd907e7b68d5d96eb70357913a1e4930644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd440dd1afc0fc6bfe71bb7127b40db90b38c91e316e67f58df252f36cd800bab6bd24c9d1feb0b837f39faf73a5072e78b59d99a6257e74d991e43b71edf0cad27d1497141d59d2109c5eee387edc9ab4c8c96f49bfee76a3189358c6a3b3425c4d28d714d3477d2c5d75c8e27bd2500eb0393536f52823d014744bc1057735f678294fce4b27871b0ac548fd51d0f9848f402bdac134dae1cb7e1fbdb9262c73c31ebf883cd98a60f7d7cfd251f3d03e973ff172677e07df1405234377d2090976060cded89f86051fdaf4990496fc7c502888093fe506705d50eaf858d380fe8a1f36872a8b24f43a835aea7b501e0b9fc62e6302e15b181ccc0ed67f527c152910c5fbb3d306497a406c55832f7a8bf1bc2ef8ff20a87935fd05400039f4f5852a9a4223068e454090934020092e27270ef593fda81245c406aa0f8e47eb67b22b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e528be516dc1897b53c5b8e264827f14675e2fbae4e1154278378069750c4069d00de1620317faae9dc538c97a0db357048fdd85db2d9a786ad2469e7de45dbf4ca2646970667358221220a9532e62d6fbe51baac7eda9beb67e7656e612fe8fdd7b31fab50d6ef8839dab64736f6c634300081c0033",
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

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_PartialDecryptVerifier *PartialDecryptVerifierCaller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int) ([4]*big.Int, error) {
	var out []interface{}
	err := _PartialDecryptVerifier.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_PartialDecryptVerifier *PartialDecryptVerifierSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _PartialDecryptVerifier.Contract.CompressProof(&_PartialDecryptVerifier.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_PartialDecryptVerifier *PartialDecryptVerifierCallerSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
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

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xda3496ab.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[16] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [16]*big.Int) error {
	var out []interface{}
	err := _PartialDecryptVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xda3496ab.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[16] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [16]*big.Int) error {
	return _PartialDecryptVerifier.Contract.VerifyCompressedProof(&_PartialDecryptVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xda3496ab.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[16] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [16]*big.Int) error {
	return _PartialDecryptVerifier.Contract.VerifyCompressedProof(&_PartialDecryptVerifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x4605cb89.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[16] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, input [16]*big.Int) error {
	var out []interface{}
	err := _PartialDecryptVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x4605cb89.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[16] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierSession) VerifyProof(proof [8]*big.Int, input [16]*big.Int) error {
	return _PartialDecryptVerifier.Contract.VerifyProof(&_PartialDecryptVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x4605cb89.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[16] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCallerSession) VerifyProof(proof [8]*big.Int, input [16]*big.Int) error {
	return _PartialDecryptVerifier.Contract.VerifyProof(&_PartialDecryptVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCaller) VerifyProof0(opts *bind.CallOpts, proof []byte, input []byte) error {
	var out []interface{}
	err := _PartialDecryptVerifier.contract.Call(opts, &out, "verifyProof0", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierSession) VerifyProof0(proof []byte, input []byte) error {
	return _PartialDecryptVerifier.Contract.VerifyProof0(&_PartialDecryptVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCallerSession) VerifyProof0(proof []byte, input []byte) error {
	return _PartialDecryptVerifier.Contract.VerifyProof0(&_PartialDecryptVerifier.CallOpts, proof, input)
}
