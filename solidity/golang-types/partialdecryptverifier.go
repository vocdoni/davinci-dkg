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
	Bin: "0x60808060405234601557611c14908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace1114610ed65750806344f6369214610e3d5780634605cb8914610868578063b8e72af6146106ad5763da3496ab14610055575f80fd5b346106aa576102803660031901126106aa57366084116106aa5736610284116106aa57604051906103006100898184610f3b565b803684376100986004356111e7565b6100a9602495929535604435611252565b919392906100b86064356111e7565b9390926040519660408801965f51602061155f5f395f51905f5289528860208101985f516020611a7f5f395f51905f528a525f5160206115bf5f395f51905f5281525f51602061171f5f395f51905f52604060608401925f51602061183f5f395f51905f5284525f51602061179f5f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f516020611abf5f395f51905f5285525f51602061153f5f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f51602061163f5f395f51905f5285525f5160206119ff5f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f51602061199f5f395f51905f5285525f5160206116ff5f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f516020611bbf5f395f51905f5285525f51602061189f5f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f5160206117ff5f395f51905f5285525f5160206119df5f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f51602061167f5f395f51905f5285525f516020611adf5f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f51602061191f5f395f51905f5285525f516020611a9f5f395f51905f5288526101643590818a5287838760608160075afa92101616818360808160065afa165f51602061175f5f395f51905f5285525f5160206119bf5f395f51905f5288526101843590818a5287838760608160075afa92101616818360808160065afa165f516020611a3f5f395f51905f5285525f5160206115ff5f395f51905f5288526101a43590818a5287838760608160075afa92101616818360808160065afa165f5160206118df5f395f51905f5285525f516020611aff5f395f51905f5288526101c43590818a5287838760608160075afa92101616818360808160065afa165f516020611b9f5f395f51905f5285525f5160206115df5f395f51905f5288526101e43590818a5287838760608160075afa92101616818360808160065afa165f5160206116df5f395f51905f5285525f5160206117df5f395f51905f5288526102043590818a5287838760608160075afa92101616818360808160065afa165f51602061169f5f395f51905f5285525f516020611a5f5f395f51905f5288526102243590818a5287838760608160075afa92101616818360808160065afa165f51602061173f5f395f51905f5285525f51602061177f5f395f51905f5288526102443590818a5287838760608160075afa921016169160808160065afa16945f51602061181f5f395f51905f528352526102643580955260608160075afa9210161660408a60808160065afa1698519751981561069b5760209a9b9c8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f5160206118bf5f395f51905f526101008401525f516020611b5f5f395f51905f526101208401525f5160206117bf5f395f51905f526101408401525f5160206118ff5f395f51905f526101608401525f51602061185f5f395f51905f526101808401525f516020611b1f5f395f51905f526101a08401525f51602061197f5f395f51905f526101c08401525f51602061157f5f395f51905f526101e08401525f51602061193f5f395f51905f526102008401525f516020611a1f5f395f51905f526102208401526102408301526102608201525f516020611b3f5f395f51905f526102808201525f51602061187f5f395f51905f526102a08201525f51602061161f5f395f51905f526102c08201525f5160206116bf5f395f51905f526102e08201526040519283916106668484610f3b565b8336843760085afa1590811561068e575b5061067f5780f35b631ff3747d60e21b8152600490fd5b600191505114155f610677565b63a54f8e2760e01b8c5260048cfd5b80fd5b5034610821576040366003190112610821576004356001600160401b038111610821576106de903690600401610f0e565b6024356001600160401b038111610821576106fd903690600401610f0e565b90916101008103610859578301610100848203126108215780601f85011215610821576040519361073061010086610f3b565b8490610100810192831161082157905b828210610849575050508101610200828203126108215780601f83011215610821576040519161077261020084610f3b565b8290610200810192831161082157905b82821061082557505050303b1561082157604051634605cb8960e01b8152915f600484015b6008821061080b5750505061010482015f905b601082106107f5575050505f8161030481305afa80156107ea576107dc575080f35b6107e891505f90610f3b565b005b6040513d5f823e3d90fd5b60208060019285518152019301910190916107ba565b60208060019285518152019301910190916107a7565b5f80fd5b8135815260209182019101610782565b634e487b7160e01b5f52604160045260245ffd5b8135815260209182019101610740565b63236bd13760e01b5f5260045ffd5b34610821576103003660031901126108215736610104116108215736610304116108215760405160408101905f51602061155f5f395f51905f52815260208101915f516020611a7f5f395f51905f5283525f5160206115bf5f395f51905f528152606082015f51602061183f5f395f51905f5281525f51602061171f5f395f51905f52604061010435935f51602061179f5f395f51905f52608087019580875284848460608160075afa911016838860808160065afa165f516020611abf5f395f51905f5283525f51602061153f5f395f51905f5286526001610124359182895286868660608160075afa9310161616838860808160065afa165f51602061163f5f395f51905f5283525f5160206119ff5f395f51905f528652610144359081885285858560608160075afa92101616838860808160065afa165f51602061199f5f395f51905f5283525f5160206116ff5f395f51905f528652610164359081885285858560608160075afa92101616838860808160065afa165f516020611bbf5f395f51905f5283525f51602061189f5f395f51905f528652610184359081885285858560608160075afa92101616838860808160065afa165f5160206117ff5f395f51905f5283525f5160206119df5f395f51905f5286526101a4359081885285858560608160075afa92101616838860808160065afa165f51602061167f5f395f51905f5283525f516020611adf5f395f51905f5286526101c4359081885285858560608160075afa92101616838860808160065afa165f51602061191f5f395f51905f5283525f516020611a9f5f395f51905f5286526101e4359081885285858560608160075afa92101616838860808160065afa165f51602061175f5f395f51905f5283525f5160206119bf5f395f51905f528652610204359081885285858560608160075afa92101616838860808160065afa165f516020611a3f5f395f51905f5283525f5160206115ff5f395f51905f528652610224359081885285858560608160075afa92101616838860808160065afa165f5160206118df5f395f51905f5283525f516020611aff5f395f51905f528652610244359081885285858560608160075afa92101616838860808160065afa165f516020611b9f5f395f51905f5283525f5160206115df5f395f51905f528652610264359081885285858560608160075afa92101616838860808160065afa165f5160206116df5f395f51905f5283525f5160206117df5f395f51905f528652610284359081885285858560608160075afa92101616838860808160065afa165f51602061169f5f395f51905f5283525f516020611a5f5f395f51905f5286526102a4359081885285858560608160075afa92101616838860808160065afa165f51602061173f5f395f51905f5283525f51602061177f5f395f51905f5286526102c4359081885285858560608160075afa92101616838860808160065afa16945f51602061181f5f395f51905f528352526102e43580955260608160075afa9210161660408260808160065afa16905191519015610e2e5760405191610100600484375f5160206118bf5f395f51905f526101008401525f516020611b5f5f395f51905f526101208401525f5160206117bf5f395f51905f526101408401525f5160206118ff5f395f51905f526101608401525f51602061185f5f395f51905f526101808401525f516020611b1f5f395f51905f526101a08401525f51602061197f5f395f51905f526101c08401525f51602061157f5f395f51905f526101e08401525f51602061193f5f395f51905f526102008401525f516020611a1f5f395f51905f526102208401526102408301526102608201525f516020611b3f5f395f51905f526102808201525f51602061187f5f395f51905f526102a08201525f51602061161f5f395f51905f526102c08201525f5160206116bf5f395f51905f526102e08201526020816103008160085afa90511615610e1f57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b346108215761010036600319011261082157366101041161082157604051610e66608082610f3b565b6080368237610e79602435600435610f5e565b8152610e8f60843560a435604435606435610fff565b60208301526040820152610ea760e43560c435610f5e565b6060820152604051905f825b60048210610ec057608084f35b6020806001928551815201930191019091610eb3565b34610821575f36600319011261082157807ffabc38153fad944fbc64c51867e8df0f0c3e8a73721f8ce449a0c99134129d6460209252f35b9181601f84011215610821578235916001600160401b038311610821576020838186019501011161082157565b601f909101601f19168101906001600160401b0382119082101761083557604052565b905f51602061159f5f395f51905f528210801590610fe8575b610e1f57811580610fe0575b610fda57610fa75f51602061159f5f395f51905f5260038185818180090908611352565b818103610fb657505060011b90565b5f51602061159f5f395f51905f52809106810306145f14610e1f57600190811b1790565b50505f90565b508015610f83565b505f51602061159f5f395f51905f52811015610f77565b919093925f51602061159f5f395f51905f5283108015906111d0575b80156111b9575b80156111a2575b610e1f578082868517171715611197579082916110fa5f51602061159f5f395f51905f5280808080888180808f9d5f51602061195f5f395f51905f528f839290839109099d8e0981848181800909085f516020611b7f5f395f51905f52089a09818c8181800909085f51602061151f5f395f51905f520806810306945f51602061159f5f395f51905f525f51602061165f5f395f51905f52816110d481808b80098187800908611352565b8408095f51602061159f5f395f51905f526110ee826114b6565b80091415958691611375565b92908082148061118e575b1561112c5750505050905f146111245760ff60025b169060021b179190565b60ff5f61111a565b5f51602061159f5f395f51905f5280910681030614918261116f575b505015610e1f57600191156111675760ff60025b169060021b17179190565b60ff5f61115c565b5f51602061159f5f395f51905f52919250819006810306145f80611148565b50838314611105565b50505090505f905f90565b505f51602061159f5f395f51905f52811015611029565b505f51602061159f5f395f51905f52821015611022565b505f51602061159f5f395f51905f5285101561101b565b801561124b578060011c915f51602061159f5f395f51905f52831015610e1f5760018061122a5f51602061159f5f395f51905f5260038188818180090908611352565b93161461123357565b905f51602061159f5f395f51905f5280910681030690565b505f905f90565b80158061134a575b61133e578060021c92825f51602061159f5f395f51905f528510801590611327575b610e1f5784815f51602061159f5f395f51905f5280808080808080805f51602061195f5f395f51905f52816112f19d8d0909998a0981898181800909085f51602061151f5f395f51905f520806810306936002808a16149509818a8181800909085f516020611b7f5f395f51905f5208611375565b80929160018082961614611303575050565b5f51602061159f5f395f51905f528093945080929550809106810306930681030690565b505f51602061159f5f395f51905f5281101561127c565b50505f905f905f905f90565b50811561125a565b9061135c826114b6565b915f51602061159f5f395f51905f5283800903610e1f57565b915f51602061159f5f395f51905f525f51602061165f5f395f51905f52816113ba939694966113ac82808a8009818a800908611352565b906114aa575b860809611352565b925f51602061159f5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f51602061159f5f395f51905f5260a083015260208260c08160055afa91519115610e1f575f51602061159f5f395f51905f52826001920903610e1f575f51602061159f5f395f51905f52908209925f51602061159f5f395f51905f52808080878009068103068187800908149081159161148b575b50610e1f57565b90505f51602061159f5f395f51905f528084860960020914155f611484565b818091068103066113b2565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f51602061159f5f395f51905f5260a083015260208260c08160055afa91519115610e1f5756fe2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7752f49e9948be8f0e7b3ff7d4a2d9ec96de63f26b957431bde7f56d00d66b6a1b20539befe24d64abcbc5e55c58653ef165985873df9d4e06f00b9baa13741a0191b3e85e96ffffe8d51483203ae677c52deede4380771ce8b360feb22ae47c80e30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd470915e779650a6224e1264c61b8bf7056b6aa4005231f74653f1df6687832d7212de1141e5427a1463f486042d266b1c1acc83e4f4a4d61a1f0d3ab7e22e159070414771496c909b21c8e261957bf280504d268e7937dbe64f5eb1efc4d4e539c1ccef5a264910c8d63af9d860cbae43ab1f19a84a26a072734e66652e446086103d045f16c276605c979a8270cecc26c54e31b11fcf71a2373560016704632be183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea408b8fa8351ba34c9787b5060b1f7027a795be4c0ccf006c34f6e3305beb9fb5e13128242c2de72a4bf747a98654eb737e6c51fd12efd1e8d630bb36e9d39e735056e97923712f3fc4a47fb907b9c66c0ae88924102fda1175b36b07f05bd5e732df29a56e548b2e98c93a20953ae2463e30f462f3d029c8380b5974c0b7dcbfa2b9221e2b8cc994161c21eb2149a1bbd7b3deebdd5a184d5e11a9ebf5f3d1be430644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000012641d2c39179218631960d17d07b4bcf236dfb1bb00b812a8ab744b7740d1352145a9bca4a2251e85815046b7ebc50c7fa39780638f3cd08a5e5ae5e29c5a8a617322ee1ebba092a670b9f3e6b64d03a2e915b6be25ac93e525484e7746aa3a008c6da0db64ff640032f997de9ddc2fd87454e8f1593c986a0c4b2f80cebf53b2e181fb63ef315855ea1dba0801cfb4d3c1e1b7a14fb138aebe537baebd6c3372838dbc7e01e2eb0f684c326e2c1d0add4aef5ab323874d0d7a7c31e4973f23210f93e325cbf50f8fe82d5337dba9ce8ba48faff83d72bbd16ccbddf0a219d3321a4b06dfc8f11ff8066ec233f03cfe797a24bbdff1717bc8e4028746edeb8d8263739095dd57cffc01068c7a6ed9ada78c81083670003fca91837e770d2a1ed224924d7e236d0c805435221e1243cbe149235ee2bbde2b102020dd65749822a0bbe4a28a2f0edcdfdc83ef495738002f9d570f7736aeb287a82d572f214aaed146fb9d82122281cf8ece67516835511faa39c451f34cdde0135a186251f7680220e369cce2d908740d9b05c3942b11c4defcd0e1ffc458fbb33152bfac76e4600525e4457ad1ce168dd4feeb6589bc0f49307b18ff7916d80d65ffaab8f1164108a05e9c61e794190f8ae69cbe88a6af502c2c4a19866bb39239077905094ce195c261995c094e91c90089bbf021d3734d14d5c00230575f75afc0eaa1044e12d7c06e5a5937bde9ea2d9e87663896d054f40c2ef3ec6d8a62448b8f0e735ba30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4422be953f0faef2ef9b8255f9e2a4fe660a81dd94d22ce40368e628d5ba80cb81267547e49df55980f0da6930fcc8fa22733d6308374d11191727017a20b5471b2532ea3ebe4a8e56ff5018417a70e86084e82728fb794362a74e0a45dc964a852da93252b024b65d61d2a6f6c85ebbcfab1aab7ee2a003dc2aeaeb96172c8d310d9d15a2ec5488f6c40cb3426433c27d9a869ac562030002992e1e3bb66f873820d5856ca5cf2754516dceb5ec0b39e4df3acfd3c3127784245228ecd806a7791e8cad970d1305e7f44842339ae3d4618462b2c303d0460ae875f1508efc8039082beb77f112f8a400198d53992a34a1c0739b721867856f4bef45e65d6957da113f177c21d83973027bf7e2779d63225dedea5814daac8adf51fba5da4bbec70d1e55c86976ef2c52f64e57e68afd51bd8819377c0b10b9eadd0fda8d98a0932b544b47c6d6e9348b2bb298588c1976896f772b747cc315912433bb6267a4671c078f410842b106db9d4cfada61f3685bb112242d52291afe0d267caa82eab92556d28190adc23826c64264ac2d883bb6803862a31028bb4a241a1e8e6d24f42449b3620b6978826db145232ade00d0899ef737db872027b706235fe4ace14a009ec2bd323fc97a1748adc91237a80c3b8e4e15bea3b486ae49bf5343e953801830d5c83872b13e90d0963fc5f4dd3c53693525dc4583aefd617086b306d0572b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e52990e4c09b665606afa6411bc065804a1383b8f0796fa8038e1a88f33719678f2c281f0564c98efeaa2835b43be49659dc91e2858d212a534db051a2e0ed77a8a264697066735822122024290dbaac4cbb799b3f4cfa06f29a87312f676c0f8f24a1ad76047be8d04d5864736f6c634300081c0033",
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
