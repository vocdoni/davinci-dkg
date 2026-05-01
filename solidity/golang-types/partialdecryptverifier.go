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
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[16]\",\"internalType\":\"uint256[16]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"input\",\"type\":\"uint256[16]\",\"internalType\":\"uint256[16]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidInputEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x60808060405234601557611e71908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c908163233ace11146114805750806344f63692146114025780634605cb8914610aa9578063b8e72af614610a4a5763da3496ab14610053575f80fd5b34610a4657610280366003190112610a465736608411610a46573661028411610a4657610300604051610086828261150d565b81368237610095600435611a98565b6100a6602493929335604435611b03565b919392906100b5606435611a98565b9390926040519660408801967f28058ec429811f31ccc717c2f35037da897c8d7c49a39afdaf5463088e5371c489528860208101987f0c36d8b95a28c88752a778b5b8dcdcf5305a698044fd0c3f3c6b37731545a9a08a527f249cbb6260c1fcecc3414dc396b0ae269ff2db4f9771e3d19a2ffd4e7572549281527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f14b1df9eeeeff14c150f0a42d11e2ce3be92f8601b9c52c61a8441e0a661f4f784527f25abb187af90ffda14d20ab1f876eff14dfd6b76360322d3597f7a44fe19052f6084359583608082019780895286828660608160075afa911016818360808160065afa167f0edb424a57930415143d8060dca2d7d261d2bd5659e1e7818064f9f3476dc9fc85527f27d25099732a1b6acb2af808a63d47508cd0f21694bc4025a16b36f2d9f0c68a8852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f0ec2d26710fa59099d3b48dbc1661b1592349fbc31fd1fee631541ada71a040185527f0ac47166f383b3738fcf44be314e99b1db4bde71c52326036c879f2ef1f7f4c4885260c43590818a5287838760608160075afa92101616818360808160065afa167f09c3532d5a447ad8fd9fe75b101d26f274845625a9f528d704e209520fa504af85527f2b7acd2744f72bb023c74dcede6242c22d9554e390405cc0a96a72c6c6e50b1e885260e43590818a5287838760608160075afa92101616818360808160065afa167f2eb3b1e5bb087b568f08842fabe5ea1cab02f47e20ce294114730031fc50c04785527f0eb386c18d7b6c6c7d64abbdee5822d2887fb26cf47d26b1c9de20a3ef11a9e688526101043590818a5287838760608160075afa92101616818360808160065afa167f2ee78847ada9f79a6bd1940739763bc983a804bb44642c68091d2a1fd8dd30f685527f2d2064dbed9599de21efb1a6d17ce8857cd336ddf1b7740333f72b46d6f645ac88526101243590818a5287838760608160075afa92101616818360808160065afa167f274588c13e7f0a2be656e371b04bd6d61172ebfed84e8ba962219a8dca39e12a85527f0ae2cf180dcc657e2ce2a95e25b59fbac75ad35d34473c2d8efe6f48071a881288526101443590818a5287838760608160075afa92101616818360808160065afa167f125ea33d2b31142e2b2ef1a10b61bce20f02cce74527789089d8e25690cf236a85527f22d3321831392f2bdc0d10629966fe1df9381227728cc0a1d28205ac2057b32388526101643590818a5287838760608160075afa92101616818360808160065afa167f2f5210eb1554cf119e144b3568c4aa8e4ec43f958b418620f39776be4db901b885527f22636fb93f60ede88581c2e90a8747b7115a2dcba2a93a08765d0be16a2fb0d788526101843590818a5287838760608160075afa92101616818360808160065afa167e349907fddea37eb730025481032b01758790d726e5d560327ed39beb82b2f585527f0e5d1f1cc4f262516345f6b47f0b0cec0a0148d83e7d14455ec6975cc6a940f288526101a43590818a5287838760608160075afa92101616818360808160065afa167f0100ec879b925a45d2fbcbf3dda68e9d4b545c58dd70bdb4236e66a9b938a31185527f1c98168999316562be8daeaab75c81684520eba9a1e2c2b714b002eb054d0f7b88526101c43590818a5287838760608160075afa92101616818360808160065afa167f0485d8912b44fc0de139c3bd5f3530b8af0d6a971ceed2c7013b1d25bd2a7d3985527f285b4c40d6f92beb1e0ef13d66879b4381eac403e99bf1387972e1f5286a5e3588526101e43590818a5287838760608160075afa92101616818360808160065afa167f1c1cd4c0740066154f521e66a4128bdfe713829d2968cc29cbd5e2a66257a05185527f1f0eca679b677e8ce739777991b53455fd8eb2eec6ad1429384595ad10449a4088526102043590818a5287838760608160075afa92101616818360808160065afa167f205782b035b9689e32bbe4e20190677d766499e8faaca267565f08e708978a3d85527f2253359dd7b9d1685b319e02ec75e6f0639abee0808a6670fe6c816ae7a033e588526102243590818a5287838760608160075afa92101616818360808160065afa167f0422c199e6b8a024a53858381daba4651afd7e96b3bfd2abdb232d5cf59afbf585527f27435bcf62618c0d39b88d7498e825ba2efb5781809707d9a67e6b5d3f4991f888526102443590818a5287838760608160075afa921016169160808160065afa16947f036bf43bc528fb2b2b6022c81165f2addb3248237537f1bbb15f1df776818f068352526102643580955260608160075afa9210161660408a60808160065afa16985197519815610a375760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f0cc1e2fecbea7c434f596caf004fe64b21e3bd3143a44801c3005e2b35ef8aa06101008401527f08aaa064c6bd16ae2494e3507fc643cdd7f6e76ad91249aac9047abba7c1eed06101208401527f1b3a7e08dae8548fa87109b4f8d7b72f2273ca011ab4bd610ebc96cc760dbfed6101408401527f29a5f24a357044671fc60d5eaaa00cb4db7918bd6df9e6561d212b7f4c7feb036101608401527f20d330044366153caa7733ed55a72bc5c9c370efe3a01b40ea73d4bec44a25536101808401527f2f799743ddd22acd8a5636979093084794a9c0ce0c85b85cb516eff2841498256101a08401527f2b78861890caa23adcfc9d4a285a913900bc186e73af2805f46b0c85a5963d096101c08401527f1a79cb0b18d17d0741f046b8853fa9baa97b42ec10efd42c949b64f37dbfa0396101e08401527f0ea4c004927fa47a3d5954ccf4b8c5d7e2631ad9ecdf713b7844c612b5a8117d6102008401527f07af2bff9ec850670033d71194fdcd8de136b5767274d16fd6ba624afda7eeb16102208401526102408301526102608201527f0917eb184be5aae54dfbf8600c92aa00afefc9fbcbb04ef9f4fc2850c8ee4ac56102808201527f186b05cad5e6ecc017a3b3489dbaefeef51743307aaf1bb02e6720c2416350936102a08201527f06b7ea5ea2b4a7d689644000e9d5a0434d971407d6e7bd1a82c5f4f81c7b57d46102c08201527f03194239964487332a496113b29dbcc1150c2c92a36168e8d4c7c62b8f2a76e36102e0820152604051928391610a03848461150d565b8336843760085afa15908115610a2a575b50610a1b57005b631ff3747d60e21b5f5260045ffd5b600191505114155f610a14565b63a54f8e2760e01b5f5260045ffd5b5f80fd5b34610a46576040366003190112610a465760043567ffffffffffffffff8111610a4657610a7b9036906004016114df565b6024359167ffffffffffffffff8311610a4657610a9f610aa79336906004016114df565b9290916115be565b005b34610a4657610300366003190112610a46573661010411610a46573661030411610a465760405160408101907f28058ec429811f31ccc717c2f35037da897c8d7c49a39afdaf5463088e5371c4815260208101917f0c36d8b95a28c88752a778b5b8dcdcf5305a698044fd0c3f3c6b37731545a9a083527f249cbb6260c1fcecc3414dc396b0ae269ff2db4f9771e3d19a2ffd4e757254928152606082017f14b1df9eeeeff14c150f0a42d11e2ce3be92f8601b9c52c61a8441e0a661f4f781527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f25abb187af90ffda14d20ab1f876eff14dfd6b76360322d3597f7a44fe19052f608087019580875284848460608160075afa911016838860808160065afa167f0edb424a57930415143d8060dca2d7d261d2bd5659e1e7818064f9f3476dc9fc83527f27d25099732a1b6acb2af808a63d47508cd0f21694bc4025a16b36f2d9f0c68a86526001610124359182895286868660608160075afa9310161616838860808160065afa167f0ec2d26710fa59099d3b48dbc1661b1592349fbc31fd1fee631541ada71a040183527f0ac47166f383b3738fcf44be314e99b1db4bde71c52326036c879f2ef1f7f4c48652610144359081885285858560608160075afa92101616838860808160065afa167f09c3532d5a447ad8fd9fe75b101d26f274845625a9f528d704e209520fa504af83527f2b7acd2744f72bb023c74dcede6242c22d9554e390405cc0a96a72c6c6e50b1e8652610164359081885285858560608160075afa92101616838860808160065afa167f2eb3b1e5bb087b568f08842fabe5ea1cab02f47e20ce294114730031fc50c04783527f0eb386c18d7b6c6c7d64abbdee5822d2887fb26cf47d26b1c9de20a3ef11a9e68652610184359081885285858560608160075afa92101616838860808160065afa167f2ee78847ada9f79a6bd1940739763bc983a804bb44642c68091d2a1fd8dd30f683527f2d2064dbed9599de21efb1a6d17ce8857cd336ddf1b7740333f72b46d6f645ac86526101a4359081885285858560608160075afa92101616838860808160065afa167f274588c13e7f0a2be656e371b04bd6d61172ebfed84e8ba962219a8dca39e12a83527f0ae2cf180dcc657e2ce2a95e25b59fbac75ad35d34473c2d8efe6f48071a881286526101c4359081885285858560608160075afa92101616838860808160065afa167f125ea33d2b31142e2b2ef1a10b61bce20f02cce74527789089d8e25690cf236a83527f22d3321831392f2bdc0d10629966fe1df9381227728cc0a1d28205ac2057b32386526101e4359081885285858560608160075afa92101616838860808160065afa167f2f5210eb1554cf119e144b3568c4aa8e4ec43f958b418620f39776be4db901b883527f22636fb93f60ede88581c2e90a8747b7115a2dcba2a93a08765d0be16a2fb0d78652610204359081885285858560608160075afa92101616838860808160065afa167e349907fddea37eb730025481032b01758790d726e5d560327ed39beb82b2f583527f0e5d1f1cc4f262516345f6b47f0b0cec0a0148d83e7d14455ec6975cc6a940f28652610224359081885285858560608160075afa92101616838860808160065afa167f0100ec879b925a45d2fbcbf3dda68e9d4b545c58dd70bdb4236e66a9b938a31183527f1c98168999316562be8daeaab75c81684520eba9a1e2c2b714b002eb054d0f7b8652610244359081885285858560608160075afa92101616838860808160065afa167f0485d8912b44fc0de139c3bd5f3530b8af0d6a971ceed2c7013b1d25bd2a7d3983527f285b4c40d6f92beb1e0ef13d66879b4381eac403e99bf1387972e1f5286a5e358652610264359081885285858560608160075afa92101616838860808160065afa167f1c1cd4c0740066154f521e66a4128bdfe713829d2968cc29cbd5e2a66257a05183527f1f0eca679b677e8ce739777991b53455fd8eb2eec6ad1429384595ad10449a408652610284359081885285858560608160075afa92101616838860808160065afa167f205782b035b9689e32bbe4e20190677d766499e8faaca267565f08e708978a3d83527f2253359dd7b9d1685b319e02ec75e6f0639abee0808a6670fe6c816ae7a033e586526102a4359081885285858560608160075afa92101616838860808160065afa167f0422c199e6b8a024a53858381daba4651afd7e96b3bfd2abdb232d5cf59afbf583527f27435bcf62618c0d39b88d7498e825ba2efb5781809707d9a67e6b5d3f4991f886526102c4359081885285858560608160075afa92101616838860808160065afa16947f036bf43bc528fb2b2b6022c81165f2addb3248237537f1bbb15f1df776818f068352526102e43580955260608160075afa9210161660408260808160065afa16905191519015610a375760405191610100600484377f0cc1e2fecbea7c434f596caf004fe64b21e3bd3143a44801c3005e2b35ef8aa06101008401527f08aaa064c6bd16ae2494e3507fc643cdd7f6e76ad91249aac9047abba7c1eed06101208401527f1b3a7e08dae8548fa87109b4f8d7b72f2273ca011ab4bd610ebc96cc760dbfed6101408401527f29a5f24a357044671fc60d5eaaa00cb4db7918bd6df9e6561d212b7f4c7feb036101608401527f20d330044366153caa7733ed55a72bc5c9c370efe3a01b40ea73d4bec44a25536101808401527f2f799743ddd22acd8a5636979093084794a9c0ce0c85b85cb516eff2841498256101a08401527f2b78861890caa23adcfc9d4a285a913900bc186e73af2805f46b0c85a5963d096101c08401527f1a79cb0b18d17d0741f046b8853fa9baa97b42ec10efd42c949b64f37dbfa0396101e08401527f0ea4c004927fa47a3d5954ccf4b8c5d7e2631ad9ecdf713b7844c612b5a8117d6102008401527f07af2bff9ec850670033d71194fdcd8de136b5767274d16fd6ba624afda7eeb16102208401526102408301526102608201527f0917eb184be5aae54dfbf8600c92aa00afefc9fbcbb04ef9f4fc2850c8ee4ac56102808201527f186b05cad5e6ecc017a3b3489dbaefeef51743307aaf1bb02e6720c2416350936102a08201527f06b7ea5ea2b4a7d689644000e9d5a0434d971407d6e7bd1a82c5f4f81c7b57d46102c08201527f03194239964487332a496113b29dbcc1150c2c92a36168e8d4c7c62b8f2a76e36102e08201526020816103008160085afa90511615610a1b57005b34610a4657610100366003190112610a46573661010411610a4657608060405161142c828261150d565b8136823761143e602435600435611766565b815261145460843560a435604435606435611807565b6020830152604082015261146c60e43560c435611766565b606082015261147e60405180926114b8565bf35b34610a46575f366003190112610a4657807fe36767c950be07ed840c6f59c9a3c3cd54f08594a256d18eceb07997a9ff524560209252f35b905f905b600482106114c957505050565b60208060019285518152019301910190916114bc565b9181601f84011215610a465782359167ffffffffffffffff8311610a465760208381860195010111610a4657565b90601f8019910116810190811067ffffffffffffffff82111761152f57604052565b634e487b7160e01b5f52604160045260245ffd5b9061020082820312610a465780601f83011215610a46576040519161156a6102008461150d565b82906102008101928311610a4657905b8282106115875750505090565b813581526020918201910161157a565b905f905b601082106115a857505050565b602080600192855181520193019101909161159b565b9392919061010081146116a057608081146115e25763236bd13760e01b5f5260045ffd5b610200830361169157840191608085840312610a465782601f86011215610a46576040519261161260808561150d565b839560808101918211610a4657955b81871061168157505061167f9394508161164091611669930190611543565b60405163da3496ab60e01b60208201529261165f9060248501906114b8565b60a4830190611597565b610284815261167a6102a48261150d565b611a3b565b565b8635815260209687019601611621565b630c0b7e3560e11b5f5260045ffd5b61020083036116915784019161010085840312610a465782601f86011215610a4657604051926116d26101008561150d565b83956101008101918211610a4657955b8187106117565750506116fa92939450810190611543565b604051634605cb8960e01b6020820152915f602484015b60088210611740575050509061172f61167f92610124830190611597565b610304815261167a6103248261150d565b6020806001928551815201930191019091611711565b86358152602096870196016116e2565b905f516020611e1c5f395f51905f5282108015906117f0575b610a1b578115806117e8575b6117e2576117af5f516020611e1c5f395f51905f5260038185818180090908611c3c565b8181036117be57505060011b90565b5f516020611e1c5f395f51905f52809106810306145f14610a1b57600190811b1790565b50505f90565b50801561178b565b505f516020611e1c5f395f51905f5281101561177f565b919093925f516020611e1c5f395f51905f528310801590611a24575b8015611a0d575b80156119f6575b610a1b5780828685171717156119eb5790829161194e5f516020611e1c5f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611e1c5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161192881808b80098187800908611c3c565b8408095f516020611e1c5f395f51905f5261194282611db3565b80091415958691611c5f565b9290808214806119e2575b156119805750505050905f146119785760ff60025b169060021b179190565b60ff5f61196e565b5f516020611e1c5f395f51905f528091068103061491826119c3575b505015610a1b57600191156119bb5760ff60025b169060021b17179190565b60ff5f6119b0565b5f516020611e1c5f395f51905f52919250819006810306145f8061199c565b50838314611959565b50505090505f905f90565b505f516020611e1c5f395f51905f52811015611831565b505f516020611e1c5f395f51905f5282101561182a565b505f516020611e1c5f395f51905f52851015611823565b5f8091602081519101305afa3d15611a90573d9067ffffffffffffffff821161152f5760405191611a76601f8201601f19166020018461150d565b82523d5f602084013e5b15611a885750565b602081519101fd5b606090611a80565b8015611afc578060011c915f516020611e1c5f395f51905f52831015610a1b57600180611adb5f516020611e1c5f395f51905f5260038188818180090908611c3c565b931614611ae457565b905f516020611e1c5f395f51905f5280910681030690565b505f905f90565b801580611c34575b611c28578060021c92825f516020611e1c5f395f51905f528510801590611c11575b610a1b5784815f516020611e1c5f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4481611bdb9d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611c5f565b80929160018082961614611bed575050565b5f516020611e1c5f395f51905f528093945080929550809106810306930681030690565b505f516020611e1c5f395f51905f52811015611b2d565b50505f905f905f905f90565b508115611b0b565b90611c4682611db3565b915f516020611e1c5f395f51905f5283800903610a1b57565b915f516020611e1c5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611cb793969496611ca982808a8009818a800908611c3c565b90611da7575b860809611c3c565b925f516020611e1c5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611e1c5f395f51905f5260a083015260208260c08160055afa91519115610a1b575f516020611e1c5f395f51905f52826001920903610a1b575f516020611e1c5f395f51905f52908209925f516020611e1c5f395f51905f528080808780090681030681878009081490811591611d88575b50610a1b57565b90505f516020611e1c5f395f51905f528084860960020914155f611d81565b81809106810306611caf565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611e1c5f395f51905f5260a083015260208260c08160055afa91519115610a1b5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220f0466c84a9089d66a68f891f369c4b71c72626298cfc1a75689b725d5d4b1cb364736f6c634300081c0033",
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
