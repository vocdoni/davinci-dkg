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
	Bin: "0x60808060405234601557611c14908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace1114610ed65750806344f6369214610e3d5780634605cb8914610868578063b8e72af6146106ad5763da3496ab14610055575f80fd5b346106aa576102803660031901126106aa57366084116106aa5736610284116106aa57604051906103006100898184610f3b565b803684376100986004356111e7565b6100a9602495929535604435611252565b919392906100b86064356111e7565b9390926040519660408801965f51602061193f5f395f51905f5289528860208101985f516020611b3f5f395f51905f528a525f51602061157f5f395f51905f5281525f5160206118df5f395f51905f52604060608401925f51602061197f5f395f51905f5284525f5160206117ff5f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f51602061179f5f395f51905f5285525f5160206116ff5f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f51602061171f5f395f51905f5285525f51602061177f5f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f516020611b5f5f395f51905f5285525f516020611a3f5f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f51602061195f5f395f51905f5285525f5160206116bf5f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f51602061161f5f395f51905f5285525f516020611abf5f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f516020611a1f5f395f51905f5285525f5160206119ff5f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f516020611b7f5f395f51905f5285525f51602061189f5f395f51905f5288526101643590818a5287838760608160075afa92101616818360808160065afa165f5160206115bf5f395f51905f5285525f51602061185f5f395f51905f5288526101843590818a5287838760608160075afa92101616818360808160065afa165f51602061191f5f395f51905f5285525f516020611a9f5f395f51905f5288526101a43590818a5287838760608160075afa92101616818360808160065afa165f5160206118bf5f395f51905f5285525f5160206115ff5f395f51905f5288526101c43590818a5287838760608160075afa92101616818360808160065afa165f516020611b1f5f395f51905f5285525f51602061183f5f395f51905f5288526101e43590818a5287838760608160075afa92101616818360808160065afa165f5160206117df5f395f51905f5285525f51602061163f5f395f51905f5288526102043590818a5287838760608160075afa92101616818360808160065afa165f51602061155f5f395f51905f5285525f51602061153f5f395f51905f5288526102243590818a5287838760608160075afa92101616818360808160065afa165f51602061187f5f395f51905f5285525f51602061159f5f395f51905f5288526102443590818a5287838760608160075afa921016169160808160065afa16945f51602061169f5f395f51905f528352526102643580955260608160075afa9210161660408a60808160065afa1698519751981561069b5760209a9b9c8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f5160206119df5f395f51905f526101008401525f516020611adf5f395f51905f526101208401525f5160206118ff5f395f51905f526101408401525f51602061199f5f395f51905f526101608401525f51602061181f5f395f51905f526101808401525f51602061167f5f395f51905f526101a08401525f5160206117bf5f395f51905f526101c08401525f516020611b9f5f395f51905f526101e08401525f5160206115df5f395f51905f526102008401525f5160206119bf5f395f51905f526102208401526102408301526102608201525f516020611a7f5f395f51905f526102808201525f516020611aff5f395f51905f526102a08201525f5160206116df5f395f51905f526102c08201525f51602061173f5f395f51905f526102e08201526040519283916106668484610f3b565b8336843760085afa1590811561068e575b5061067f5780f35b631ff3747d60e21b8152600490fd5b600191505114155f610677565b63a54f8e2760e01b8c5260048cfd5b80fd5b5034610821576040366003190112610821576004356001600160401b038111610821576106de903690600401610f0e565b6024356001600160401b038111610821576106fd903690600401610f0e565b90916101008103610859578301610100848203126108215780601f85011215610821576040519361073061010086610f3b565b8490610100810192831161082157905b828210610849575050508101610200828203126108215780601f83011215610821576040519161077261020084610f3b565b8290610200810192831161082157905b82821061082557505050303b1561082157604051634605cb8960e01b8152915f600484015b6008821061080b5750505061010482015f905b601082106107f5575050505f8161030481305afa80156107ea576107dc575080f35b6107e891505f90610f3b565b005b6040513d5f823e3d90fd5b60208060019285518152019301910190916107ba565b60208060019285518152019301910190916107a7565b5f80fd5b8135815260209182019101610782565b634e487b7160e01b5f52604160045260245ffd5b8135815260209182019101610740565b63236bd13760e01b5f5260045ffd5b34610821576103003660031901126108215736610104116108215736610304116108215760405160408101905f51602061193f5f395f51905f52815260208101915f516020611b3f5f395f51905f5283525f51602061157f5f395f51905f528152606082015f51602061197f5f395f51905f5281525f5160206118df5f395f51905f52604061010435935f5160206117ff5f395f51905f52608087019580875284848460608160075afa911016838860808160065afa165f51602061179f5f395f51905f5283525f5160206116ff5f395f51905f5286526001610124359182895286868660608160075afa9310161616838860808160065afa165f51602061171f5f395f51905f5283525f51602061177f5f395f51905f528652610144359081885285858560608160075afa92101616838860808160065afa165f516020611b5f5f395f51905f5283525f516020611a3f5f395f51905f528652610164359081885285858560608160075afa92101616838860808160065afa165f51602061195f5f395f51905f5283525f5160206116bf5f395f51905f528652610184359081885285858560608160075afa92101616838860808160065afa165f51602061161f5f395f51905f5283525f516020611abf5f395f51905f5286526101a4359081885285858560608160075afa92101616838860808160065afa165f516020611a1f5f395f51905f5283525f5160206119ff5f395f51905f5286526101c4359081885285858560608160075afa92101616838860808160065afa165f516020611b7f5f395f51905f5283525f51602061189f5f395f51905f5286526101e4359081885285858560608160075afa92101616838860808160065afa165f5160206115bf5f395f51905f5283525f51602061185f5f395f51905f528652610204359081885285858560608160075afa92101616838860808160065afa165f51602061191f5f395f51905f5283525f516020611a9f5f395f51905f528652610224359081885285858560608160075afa92101616838860808160065afa165f5160206118bf5f395f51905f5283525f5160206115ff5f395f51905f528652610244359081885285858560608160075afa92101616838860808160065afa165f516020611b1f5f395f51905f5283525f51602061183f5f395f51905f528652610264359081885285858560608160075afa92101616838860808160065afa165f5160206117df5f395f51905f5283525f51602061163f5f395f51905f528652610284359081885285858560608160075afa92101616838860808160065afa165f51602061155f5f395f51905f5283525f51602061153f5f395f51905f5286526102a4359081885285858560608160075afa92101616838860808160065afa165f51602061187f5f395f51905f5283525f51602061159f5f395f51905f5286526102c4359081885285858560608160075afa92101616838860808160065afa16945f51602061169f5f395f51905f528352526102e43580955260608160075afa9210161660408260808160065afa16905191519015610e2e5760405191610100600484375f5160206119df5f395f51905f526101008401525f516020611adf5f395f51905f526101208401525f5160206118ff5f395f51905f526101408401525f51602061199f5f395f51905f526101608401525f51602061181f5f395f51905f526101808401525f51602061167f5f395f51905f526101a08401525f5160206117bf5f395f51905f526101c08401525f516020611b9f5f395f51905f526101e08401525f5160206115df5f395f51905f526102008401525f5160206119bf5f395f51905f526102208401526102408301526102608201525f516020611a7f5f395f51905f526102808201525f516020611aff5f395f51905f526102a08201525f5160206116df5f395f51905f526102c08201525f51602061173f5f395f51905f526102e08201526020816103008160085afa90511615610e1f57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b346108215761010036600319011261082157366101041161082157604051610e66608082610f3b565b6080368237610e79602435600435610f5e565b8152610e8f60843560a435604435606435610fff565b60208301526040820152610ea760e43560c435610f5e565b6060820152604051905f825b60048210610ec057608084f35b6020806001928551815201930191019091610eb3565b34610821575f36600319011261082157807f8d5cb7d575354215d07b8b2721e1a5465424a85f78ca2d0beb56041efd2b082060209252f35b9181601f84011215610821578235916001600160401b038311610821576020838186019501011161082157565b601f909101601f19168101906001600160401b0382119082101761083557604052565b905f51602061165f5f395f51905f528210801590610fe8575b610e1f57811580610fe0575b610fda57610fa75f51602061165f5f395f51905f5260038185818180090908611352565b818103610fb657505060011b90565b5f51602061165f5f395f51905f52809106810306145f14610e1f57600190811b1790565b50505f90565b508015610f83565b505f51602061165f5f395f51905f52811015610f77565b919093925f51602061165f5f395f51905f5283108015906111d0575b80156111b9575b80156111a2575b610e1f578082868517171715611197579082916110fa5f51602061165f5f395f51905f5280808080888180808f9d5f516020611a5f5f395f51905f528f839290839109099d8e0981848181800909085f516020611bbf5f395f51905f52089a09818c8181800909085f51602061151f5f395f51905f520806810306945f51602061165f5f395f51905f525f51602061175f5f395f51905f52816110d481808b80098187800908611352565b8408095f51602061165f5f395f51905f526110ee826114b6565b80091415958691611375565b92908082148061118e575b1561112c5750505050905f146111245760ff60025b169060021b179190565b60ff5f61111a565b5f51602061165f5f395f51905f5280910681030614918261116f575b505015610e1f57600191156111675760ff60025b169060021b17179190565b60ff5f61115c565b5f51602061165f5f395f51905f52919250819006810306145f80611148565b50838314611105565b50505090505f905f90565b505f51602061165f5f395f51905f52811015611029565b505f51602061165f5f395f51905f52821015611022565b505f51602061165f5f395f51905f5285101561101b565b801561124b578060011c915f51602061165f5f395f51905f52831015610e1f5760018061122a5f51602061165f5f395f51905f5260038188818180090908611352565b93161461123357565b905f51602061165f5f395f51905f5280910681030690565b505f905f90565b80158061134a575b61133e578060021c92825f51602061165f5f395f51905f528510801590611327575b610e1f5784815f51602061165f5f395f51905f5280808080808080805f516020611a5f5f395f51905f52816112f19d8d0909998a0981898181800909085f51602061151f5f395f51905f520806810306936002808a16149509818a8181800909085f516020611bbf5f395f51905f5208611375565b80929160018082961614611303575050565b5f51602061165f5f395f51905f528093945080929550809106810306930681030690565b505f51602061165f5f395f51905f5281101561127c565b50505f905f905f905f90565b50811561125a565b9061135c826114b6565b915f51602061165f5f395f51905f5283800903610e1f57565b915f51602061165f5f395f51905f525f51602061175f5f395f51905f52816113ba939694966113ac82808a8009818a800908611352565b906114aa575b860809611352565b925f51602061165f5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f51602061165f5f395f51905f5260a083015260208260c08160055afa91519115610e1f575f51602061165f5f395f51905f52826001920903610e1f575f51602061165f5f395f51905f52908209925f51602061165f5f395f51905f52808080878009068103068187800908149081159161148b575b50610e1f57565b90505f51602061165f5f395f51905f528084860960020914155f611484565b818091068103066113b2565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f51602061165f5f395f51905f5260a083015260208260c08160055afa91519115610e1f5756fe2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775209b8d24b8585347768b30941f3c6c5dfba4726eaec1df192fd4ae60dc3128be0f7f659f4fa846222ac85d417bccf2cfbaca1d8ab96af978a0fe4c33277b18cc1637f62fd1e4efd97c58c6caf5a7ca83122ba8938488867363b5325be4e45d7319a812a0e88a2c22760c201a4be30ed4bf104ec4f141d3ec2bed233b4fe8b35c23e86e349502a15382e3a78c0870b3f5856c93f985823c2104b3dbeb9d5959f817028e81dba491144dc9543e164c9c36088e6562a4e0994780e121056b40cd78271442516230bd2e5d377fc0a42e85cba78e6a486ebf857b5db97d0ab904b3e11e403614a757f4c318b8298899e0ec536d29854e7b0c1912e43a5a16235d43c72ebb8857ff06beb9921d84f635dab1caa69204511c59861f7aa63c1af9dbfac830644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4722f563682fc87ab208149692587ae71db70e5f4282ad5e114859c290c9fdc0e114480a81e681e7d36aa2f5cdf94e7534192600e7745acede7f36062044fde1ca1f332b079f6bb85c7312efa93cc6e755c1c55c60428c80be97abeae7faf313730aa86e7078f5854591f2abbcaf45c6fa3e4b2906e1b040fe541741e9c0e507580654efaed8855afdfa44e69955ede360e46c958722a84eddd3e38a89db82f9fc26205db28e981c1e99843ef34c1c2c016128154de6a927df9dcfb6c467543d3f23a5bdb8073ae5933f515c5f6b76230f0968ac425ac92c488a29ce606fe9675d183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea417de3591e541b8c30d939914eb5a9af22d50e2fe2ba84b9974db422915cb94ff177f0af792a15b3012a4c4843627fd354e2d9ec5e4b6b7de8b0c6fde59cb0fe00a38fd3448dba055259ef9872ec2b74c4e35284c87cce0a62a4f360ea58bb67c12901620458746f50d1e38678d60d708ee8f1e5f4c421599567e6756e0651d0519f0ef645e8c19d82639c9a8be4a242ac3be1f8db7cbfca76f424aba9d50b6302ba854d266a0011bce0f755fbb5c73fd5c3fde9d843db2ba006b96c645add07b138cb6bd4e7938b19130d71dcfd2b088dc572de92d14c22f4267bc811f07db171da7d6512dfdca8e6f0fb47c654355c68cd008ecbf15c825eb062f7e9b14862f0a8700d7ea1704cd789141cc962fd8f5f76008974b8ade6b8fd97db42ec5ce4a10c8c1b69e8cc4401e01ea59a02cc90d1db99a2233ea8e4758ee409dbda543d30368e72a84304d67a2397c11538c82cf1bc474be5b0615f690dbc415fb8a668030644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000110bec214abe0d8065cb6c3eb1298fc9d63c4241503932fdf5b148d62e389ad1922f16cab8f3587ea7e4eecbf63b55808ceb6a0a131e8e8241203f371d1a0592a051a75d17be4c1d5e4d881c45c9f98496451c7c2571850d94180ef445a742ee32599954aff1b60a8e45317ec1083a9c7c6600d250c7c6cea7645daa5416263a62a768e2b53d16c5126db47c9a4a18211b429c9ba1196b25cb6a4c671365c110501a136e9a06f4c5fc9d91d79fda7ecd537ea4cad1db502928262d0e519def1871d9d831f6aba6402903293278abb53b0c3d24feb10286620e6e3434d2a418e2a21d788b60920023302095384dd3d4b973cdc6fdaa4810c0ce183830d111216d12ef460e26360b4b6e1e9de585e424fcdef3a0a08d8aef27568e8ae14ffad31741772cbf5bd69d9ff01bae7bb2afd1ba8e65d424f7f9923634585199a10e0346c014428a9c286a86410e949e1e3873b27d854fd4638ba685aa4ad1b1b915e1c7930644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4413bb804e5c59b111241965c40136eba871dc5e6005f0da8e9b8a8276e60469400961d77742deb77fb5017ca043e871c4f657c293c03bf1077e2308f61e7c866d2422f719a57f1a00a0da8eea3f2cdc1171d264e33862e9fbe91a74ae750019330a37e4c909d1c7575db44c86b9c5610b1e66b4f5081fdafef08bd7427c0f8ad40cca7cb2ec32ce562a0700f3796b1c7f5d9783ac1de6e5b9a9170533ad15a51103f5442bb759929b2e826f34181d3e78815e4535bb86ca2ca4b1b7a8f045014e26257341d8310a59f366f1585349240685ed855c7409c07f6ccbc6fa050f307a0fae037c4ef22ee2d43e171f721b8c0ee836240c190204958f32cf4a0ec9dc07164df06a822e213008bd06fa860d9b53c1f5a087fe87b840fc52459ce9dffe4f0c790a2851d70cc408ed33e53396bf7537ba4421f5134952cec63f9b5f8c315c2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5a26469706673582212206726a673591384f809a67f30f001ff9fe1639b033d97961dc795bc217be3f65064736f6c634300081c0033",
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
