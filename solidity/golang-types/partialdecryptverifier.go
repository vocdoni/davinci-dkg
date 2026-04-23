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
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[13]\",\"internalType\":\"uint256[13]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"input\",\"type\":\"uint256[13]\",\"internalType\":\"uint256[13]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidInputEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x60808060405234601557611c0c908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c80631e11b6cc146109b0578063233ace111461097657806344f63692146108f8578063b8e72af6146108995763f7f338eb14610050575f80fd5b34610895576102a036600319011261089557366101041161089557366102a4116108955760405160408101907f1a6cb99620bd0176a7d8c9ab538ac16a27645c5f1ed6705fb8af864fc033788a815260208101917f0f0d1b7924967253c7b7f588c5967a0650f9fccc1eccabcced178f402474f32883527f29589fd532dc8fdf7a6cba62923c5d4d79cf52bb3343ddd7fa0152116b38031c8152606082017f2428844bf3517874cfb6efb71df2aa7d7935a7256e340e4dbb27396c42cffe6c81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f2c0d48b1ff1d653c7b4c91676a237f5886757bf1e1cb1739fb8afa298055f865608087019580875284848460608160075afa911016838860808160065afa167f2e8255201fc416d314f4c5c504a7abb2d93005dc97d0687e6f7773a2b8bc443f83527f057780c6047f3456e96890b6fa36b6409aef761ff9be62685a9afcf4e64b1fdf86526001610124359182895286868660608160075afa9310161616838860808160065afa167f2bf6606a36da0e7b9887c334d989ea94b14c6a87957f5312ac7ca5399226f1cb83527f0b206c2e50b5e484c45f65dfd367b54d2e33f88b9fb4d28c369fe63c0586b03a8652610144359081885285858560608160075afa92101616838860808160065afa167f024981fc71e6642a4a12ed94d6ef645c1b8503454a45e37a87c7ec1b59273e5883527f29c3460c639f4396c5b3b98c7a82e39fc53e8cf9e581e862c6f85a6c092776de8652610164359081885285858560608160075afa92101616838860808160065afa167f05f32dfd738ac4a586b67ef880fcfd5a478a18b7cab550805b3d286a93c27ef183527f1fce0e89d2abfed96fb62a1c5612634512c1b33a339633542eef3641e533352e8652610184359081885285858560608160075afa92101616838860808160065afa167f191844f1b738f0d703db2721671de5b29f357b5c784030de547b3571eaf689ee83527f1d7305cda3ece7aba8ea1ebd562a4f44ecc5776bbd61da984d0d48375fa2a93386526101a4359081885285858560608160075afa92101616838860808160065afa167f067054571ad86d0ff12df260528595ce10373ed178afd7139159f6dba4d0706a83527f1ee85c94fc03dd099854df180d8b0e621fd876bd361d8131b9b5ca57d38c85b386526101c4359081885285858560608160075afa92101616838860808160065afa167f2f6e9d40b7cd92ebbae23b5374b965ab005ef72b65a39317f243d5cfc515b7f883527f115a4f1f7a4d0d19a1aedce8a7dd675c7397773624dc448380809e5bf602597b86526101e4359081885285858560608160075afa92101616838860808160065afa167f1eec0153d200ed025faa9f757d38a8a2f181bdcd0e1a26b9e5d0bc344656309f83527f139ed23109dc19f72761b416258f0284c4415d4977bea6394347ebe43bc35f008652610204359081885285858560608160075afa92101616838860808160065afa167f047f1f1bcb7d0df95d5b7c037d5b3ffaf49294b57ac437eda0d8c19d260aedd183527f2c07b4de88722af23e84c9cd23db059ecc890281abcf5e2af69102884de9c7138652610224359081885285858560608160075afa92101616838860808160065afa167f20cfaba3d4dead1b93b00c1df2198c02e383b140ef79f64b68934de7b2030ee283527f1ab0a72d290807d415330ef5fc912a9c578a8c06921031ce42eca2faa295f5d88652610244359081885285858560608160075afa92101616838860808160065afa167f1e3ecfe0689b002b81906424ef693589b18672913ea34014548fa6859f34fa7b83527f029788b76137429f47b2e6232c75196945d6750b2465a89e2138669a8f0cbd708652610264359081885285858560608160075afa92101616838860808160065afa16947f2acb635d220a41ba389953d6c3865f10b7df5c8ab101de35c69765fd37eb4ac28352526102843580955260608160075afa9210161660408260808160065afa169051915190156108865760405191610100600484377f0213caf548a265a34f39111b9ddaa5686fd16d48294bca00c01441ca890575906101008401527f235b09ab07037402fc609c555a307c78078620ebf931f42e61d67347aea3d72d6101208401527f23268021218e0641b9e688e6e4706cc3b8dec80406633d212be8ac45754ddf2b6101408401527ea56730af981388c770baad19aaab899f4d6a8b0515daba2d0039501c0974b26101608401527f2abf0d1ca78fbaebf86f08596bbc1041d8e12476d5bd148bfc9f0d6f4e85a1fb6101808401527f29c9e6f46a272db8f7fe0c84bb47a7408a553cd22b2c04286f5daa93eee79b6f6101a08401527f0c210a68685fd530bb6687200ca2c370f3154269a4cdefb6fb563f9630e2c9f06101c08401527f0fcb090dc1f3db9374390e858e51ad97d7a5fe55584fcb9bcc30dcbf5832a8736101e08401527f022997a459a902af7621bd7fb7a911a140fccaaebf3ccf893ff36fa7d867caa66102008401527f0c5be8ce822c586688a9500c4f254e1897133601cc808ce8729ffdb94e818d336102208401526102408301526102608201527f2feba215b71deabe2485435e4a862c43d02ff9f0ce4e90c41646e45cd8e847186102808201527f2181941078d3b5a116a4eade43b600804ecab9ec919729fdd8b3f422746d002a6102a08201527f305ba2e96b238f9b5d6bad5eaa961ccfdaf8a27133af6afd9fc4e51e2da9cd1d6102c08201527f205e5e9407038e82c99088e0f579cd451576fe80c978e55770d8f18ce8b9cf256102e08201526020816103008160085afa9051161561087757005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b5f80fd5b346108955760403660031901126108955760043567ffffffffffffffff8111610895576108ca90369060040161127a565b6024359167ffffffffffffffff8311610895576108ee6108f693369060040161127a565b929091611359565b005b346108955761010036600319011261089557366101041161089557608060405161092282826112a8565b813682376109346024356004356116a5565b815261094a60843560a435604435606435611746565b6020830152604082015261096260e43560c4356116a5565b60608201526109746040518092611253565bf35b34610895575f3660031901126108955760206040517f3871aae077eee32665f8ee95220a1ef4ff279fc8c2387bcffedd1869b66466e88152f35b34610895576102203660031901126108955736608411610895573661022411610895576103006040516109e382826112a8565b813682376109f2600435611501565b610a0360249392933560443561156c565b91939290610a12606435611501565b9390926040519660408801967f1a6cb99620bd0176a7d8c9ab538ac16a27645c5f1ed6705fb8af864fc033788a89528860208101987f0f0d1b7924967253c7b7f588c5967a0650f9fccc1eccabcced178f402474f3288a527f29589fd532dc8fdf7a6cba62923c5d4d79cf52bb3343ddd7fa0152116b38031c81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f2428844bf3517874cfb6efb71df2aa7d7935a7256e340e4dbb27396c42cffe6c84527f2c0d48b1ff1d653c7b4c91676a237f5886757bf1e1cb1739fb8afa298055f8656084359583608082019780895286828660608160075afa911016818360808160065afa167f2e8255201fc416d314f4c5c504a7abb2d93005dc97d0687e6f7773a2b8bc443f85527f057780c6047f3456e96890b6fa36b6409aef761ff9be62685a9afcf4e64b1fdf8852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f2bf6606a36da0e7b9887c334d989ea94b14c6a87957f5312ac7ca5399226f1cb85527f0b206c2e50b5e484c45f65dfd367b54d2e33f88b9fb4d28c369fe63c0586b03a885260c43590818a5287838760608160075afa92101616818360808160065afa167f024981fc71e6642a4a12ed94d6ef645c1b8503454a45e37a87c7ec1b59273e5885527f29c3460c639f4396c5b3b98c7a82e39fc53e8cf9e581e862c6f85a6c092776de885260e43590818a5287838760608160075afa92101616818360808160065afa167f05f32dfd738ac4a586b67ef880fcfd5a478a18b7cab550805b3d286a93c27ef185527f1fce0e89d2abfed96fb62a1c5612634512c1b33a339633542eef3641e533352e88526101043590818a5287838760608160075afa92101616818360808160065afa167f191844f1b738f0d703db2721671de5b29f357b5c784030de547b3571eaf689ee85527f1d7305cda3ece7aba8ea1ebd562a4f44ecc5776bbd61da984d0d48375fa2a93388526101243590818a5287838760608160075afa92101616818360808160065afa167f067054571ad86d0ff12df260528595ce10373ed178afd7139159f6dba4d0706a85527f1ee85c94fc03dd099854df180d8b0e621fd876bd361d8131b9b5ca57d38c85b388526101443590818a5287838760608160075afa92101616818360808160065afa167f2f6e9d40b7cd92ebbae23b5374b965ab005ef72b65a39317f243d5cfc515b7f885527f115a4f1f7a4d0d19a1aedce8a7dd675c7397773624dc448380809e5bf602597b88526101643590818a5287838760608160075afa92101616818360808160065afa167f1eec0153d200ed025faa9f757d38a8a2f181bdcd0e1a26b9e5d0bc344656309f85527f139ed23109dc19f72761b416258f0284c4415d4977bea6394347ebe43bc35f0088526101843590818a5287838760608160075afa92101616818360808160065afa167f047f1f1bcb7d0df95d5b7c037d5b3ffaf49294b57ac437eda0d8c19d260aedd185527f2c07b4de88722af23e84c9cd23db059ecc890281abcf5e2af69102884de9c71388526101a43590818a5287838760608160075afa92101616818360808160065afa167f20cfaba3d4dead1b93b00c1df2198c02e383b140ef79f64b68934de7b2030ee285527f1ab0a72d290807d415330ef5fc912a9c578a8c06921031ce42eca2faa295f5d888526101c43590818a5287838760608160075afa92101616818360808160065afa167f1e3ecfe0689b002b81906424ef693589b18672913ea34014548fa6859f34fa7b85527f029788b76137429f47b2e6232c75196945d6750b2465a89e2138669a8f0cbd7088526101e43590818a5287838760608160075afa921016169160808160065afa16947f2acb635d220a41ba389953d6c3865f10b7df5c8ab101de35c69765fd37eb4ac28352526102043580955260608160075afa9210161660408a60808160065afa169851975198156108865760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f0213caf548a265a34f39111b9ddaa5686fd16d48294bca00c01441ca890575906101008401527f235b09ab07037402fc609c555a307c78078620ebf931f42e61d67347aea3d72d6101208401527f23268021218e0641b9e688e6e4706cc3b8dec80406633d212be8ac45754ddf2b6101408401527ea56730af981388c770baad19aaab899f4d6a8b0515daba2d0039501c0974b26101608401527f2abf0d1ca78fbaebf86f08596bbc1041d8e12476d5bd148bfc9f0d6f4e85a1fb6101808401527f29c9e6f46a272db8f7fe0c84bb47a7408a553cd22b2c04286f5daa93eee79b6f6101a08401527f0c210a68685fd530bb6687200ca2c370f3154269a4cdefb6fb563f9630e2c9f06101c08401527f0fcb090dc1f3db9374390e858e51ad97d7a5fe55584fcb9bcc30dcbf5832a8736101e08401527f022997a459a902af7621bd7fb7a911a140fccaaebf3ccf893ff36fa7d867caa66102008401527f0c5be8ce822c586688a9500c4f254e1897133601cc808ce8729ffdb94e818d336102208401526102408301526102608201527f2feba215b71deabe2485435e4a862c43d02ff9f0ce4e90c41646e45cd8e847186102808201527f2181941078d3b5a116a4eade43b600804ecab9ec919729fdd8b3f422746d002a6102a08201527f305ba2e96b238f9b5d6bad5eaa961ccfdaf8a27133af6afd9fc4e51e2da9cd1d6102c08201527f205e5e9407038e82c99088e0f579cd451576fe80c978e55770d8f18ce8b9cf256102e082015260405192839161122e84846112a8565b8336843760085afa15908115611246575b5061087757005b600191505114158161123f565b905f905b6004821061126457505050565b6020806001928551815201930191019091611257565b9181601f840112156108955782359167ffffffffffffffff8311610895576020838186019501011161089557565b90601f8019910116810190811067ffffffffffffffff8211176112ca57604052565b634e487b7160e01b5f52604160045260245ffd5b906101a0828203126108955780601f8301121561089557604051916113056101a0846112a8565b82906101a0810192831161089557905b8282106113225750505090565b8135815260209182019101611315565b905f905b600d821061134357505050565b6020806001928551815201930191019091611336565b93929190610100811461143b576080811461137d5763236bd13760e01b5f5260045ffd5b6101a0830361142c578401916080858403126108955782601f8601121561089557604051926113ad6080856112a8565b83956080810191821161089557955b81871061141c57505061141a939450816113db916114049301906112de565b6040516307846db360e21b6020820152926113fa906024850190611253565b60a4830190611332565b6102248152611415610244826112a8565b61197a565b565b86358152602096870196016113bc565b630c0b7e3560e11b5f5260045ffd5b6101a0830361142c57840191610100858403126108955782601f86011215610895576040519261146d610100856112a8565b8395610100810191821161089557955b8187106114f1575050611495929394508101906112de565b60405163f7f338eb60e01b6020820152915f602484015b600882106114db57505050906114ca61141a92610124830190611332565b6102a481526114156102c4826112a8565b60208060019285518152019301910190916114ac565b863581526020968701960161147d565b8015611565578060011c915f516020611bb75f395f51905f52831015610877576001806115445f516020611bb75f395f51905f52600381888181800909086119d7565b93161461154d57565b905f516020611bb75f395f51905f5280910681030690565b505f905f90565b80158061169d575b611691578060021c92825f516020611bb75f395f51905f52851080159061167a575b6108775784815f516020611bb75f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816116449d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5086119fa565b80929160018082961614611656575050565b5f516020611bb75f395f51905f528093945080929550809106810306930681030690565b505f516020611bb75f395f51905f52811015611596565b50505f905f905f905f90565b508115611574565b905f516020611bb75f395f51905f52821080159061172f575b61087757811580611727575b611721576116ee5f516020611bb75f395f51905f52600381858181800909086119d7565b8181036116fd57505060011b90565b5f516020611bb75f395f51905f52809106810306145f1461087757600190811b1790565b50505f90565b5080156116ca565b505f516020611bb75f395f51905f528110156116be565b919093925f516020611bb75f395f51905f528310801590611963575b801561194c575b8015611935575b61087757808286851717171561192a5790829161188d5f516020611bb75f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611bb75f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161186781808b800981878009086119d7565b8408095f516020611bb75f395f51905f5261188182611b4e565b800914159586916119fa565b929080821480611921575b156118bf5750505050905f146118b75760ff60025b169060021b179190565b60ff5f6118ad565b5f516020611bb75f395f51905f52809106810306149182611902575b50501561087757600191156118fa5760ff60025b169060021b17179190565b60ff5f6118ef565b5f516020611bb75f395f51905f52919250819006810306145f806118db565b50838314611898565b50505090505f905f90565b505f516020611bb75f395f51905f52811015611770565b505f516020611bb75f395f51905f52821015611769565b505f516020611bb75f395f51905f52851015611762565b5f8091602081519101305afa3d156119cf573d9067ffffffffffffffff82116112ca57604051916119b5601f8201601f1916602001846112a8565b82523d5f602084013e5b156119c75750565b602081519101fd5b6060906119bf565b906119e182611b4e565b915f516020611bb75f395f51905f528380090361087757565b915f516020611bb75f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611a5293969496611a4482808a8009818a8009086119d7565b90611b42575b8608096119d7565b925f516020611bb75f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611bb75f395f51905f5260a083015260208260c08160055afa91519115610877575f516020611bb75f395f51905f52826001920903610877575f516020611bb75f395f51905f52908209925f516020611bb75f395f51905f528080808780090681030681878009081490811591611b23575b5061087757565b90505f516020611bb75f395f51905f528084860960020914155f611b1c565b81809106810306611a4a565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611bb75f395f51905f5260a083015260208260c08160055afa915191156108775756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a26469706673582212207e36e7502b071e72cb0bcaec25e87500e9a7667a51f6ca267c0d4c97a264aee964736f6c634300081c0033",
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

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x1e11b6cc.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[13] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [13]*big.Int) error {
	var out []interface{}
	err := _PartialDecryptVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x1e11b6cc.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[13] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [13]*big.Int) error {
	return _PartialDecryptVerifier.Contract.VerifyCompressedProof(&_PartialDecryptVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x1e11b6cc.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[13] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [13]*big.Int) error {
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

// VerifyProof0 is a free data retrieval call binding the contract method 0xf7f338eb.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[13] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCaller) VerifyProof0(opts *bind.CallOpts, proof [8]*big.Int, input [13]*big.Int) error {
	var out []interface{}
	err := _PartialDecryptVerifier.contract.Call(opts, &out, "verifyProof0", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xf7f338eb.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[13] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierSession) VerifyProof0(proof [8]*big.Int, input [13]*big.Int) error {
	return _PartialDecryptVerifier.Contract.VerifyProof0(&_PartialDecryptVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xf7f338eb.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[13] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCallerSession) VerifyProof0(proof [8]*big.Int, input [13]*big.Int) error {
	return _PartialDecryptVerifier.Contract.VerifyProof0(&_PartialDecryptVerifier.CallOpts, proof, input)
}
