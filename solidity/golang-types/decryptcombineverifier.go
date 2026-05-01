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
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[13]\",\"internalType\":\"uint256[13]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"input\",\"type\":\"uint256[13]\",\"internalType\":\"uint256[13]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidInputEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x60808060405234601557611c0c908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c80631e11b6cc146109b0578063233ace111461097657806344f63692146108f8578063b8e72af6146108995763f7f338eb14610050575f80fd5b34610895576102a036600319011261089557366101041161089557366102a4116108955760405160408101907f212f8c9c30ef5086d167ac70a3ecc791344712937cbe9c81eace3670d6e360e7815260208101917f20975d217b2593dc2d06db9d9321637d4bd87e7a4a61df50274e1209e85c61bc83527f11b8de7782bdce83a93e0c6c82b8b272754c9c82b88c39cf41e4c4e0f255b7888152606082017f10075c2125e499fed523d45f169b6bd5e188e5e66047e6156a18bc49221ec25281527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f01f0763968c3953c2dbcd7ea35d7929756dab99fe380c4e1a5f1d4056d5a7baf608087019580875284848460608160075afa911016838860808160065afa167f0e9e6a2d38b98b57c84a1b640232bf488876d4adb099fc144616143166fd6c5f83527f0fac8283701ad712a895d2068f8bbdee9bf64306086dcb32c0362d54654fd6a186526001610124359182895286868660608160075afa9310161616838860808160065afa167f2c9206f6ce1cf15057bf9935d9261b97141f8b2cf65e6ac3c4829c5ebb69830883527f06f097cedc499659b999dff576f1fb734d9251692b0e5f13c7523b0d21ea109d8652610144359081885285858560608160075afa92101616838860808160065afa167f2210627388a6807e437cebf2337148e61f0460795428cba7650615140272985983527f085009bc331a387b4ff180326fba6b47553780388efc116869ec6f27353e7dd78652610164359081885285858560608160075afa92101616838860808160065afa167f23309a37e21dcc1d538a00d9f6f03aa8936d5023553a9d4226829bdc1d00183783527f04d5c862ed6de141828069b89cf8d38f57c0d859de822fe8f11388324786fd338652610184359081885285858560608160075afa92101616838860808160065afa167f24f5ce0636930fb138a43e685bde383ba9e8e8435a89e5a114490fbb3af7ce9d83527f12bfafd11848aa5cb0a92234c8d8df9e1e5ff438a3bb70ed1444c8790ad6e64c86526101a4359081885285858560608160075afa92101616838860808160065afa167f1a2f1a9b1317f91f6815ef2117e8b627f3a643814d001aff9fcf2a9569a062e383527f129af51e7ce61debf6f24a7ffd4bcb835eb6bebe9014cc67c610b9400b72935886526101c4359081885285858560608160075afa92101616838860808160065afa167f0c20d9e7d3e8d87fb101247992d9f47d3840454b8a7d2b72e55d1f7f3d8ab22b83527f1a3fa3028bc373df10ed4a636319233ceb73d1d752964cd22ac1dffcbe8c1b8e86526101e4359081885285858560608160075afa92101616838860808160065afa167f0efa6c70ee38e2280259dd4e116c152cf5687b11ab84e6828eee215cb7d012ef83527f1cf38529fdb7d92f1d05c9a0a446e9ccaa0db67ebc8acc1fef93a14cad8ede478652610204359081885285858560608160075afa92101616838860808160065afa167f2ba8e1e81c526c7e5316b3ab913b6a40c49f42615925443c816133fd3503be6983527f15f72290598781df6625146de17bf30287aa585ceb41c64228ac4dd7b653045d8652610224359081885285858560608160075afa92101616838860808160065afa167f0c74b6ef876b3cf7f4bed3439e558faa21beba0ebc7cec2abac53a4572e6f9ca83527f1f90f3643d710458d2f8286750e32e51867d650962d5056a812a24be9036f0628652610244359081885285858560608160075afa92101616838860808160065afa167f0922ec9f45ec9cb595d0d979f796fd67c0bbd94160dbcb510830a2531c95805f83527f24d79b0bc1ebc43c74eb36c3b69cd186b0e7dcb6164fad94b3dfd8965e51d46d8652610264359081885285858560608160075afa92101616838860808160065afa16947f27963007a92bc955699cb73d20b9db602fd7923f60f5158558e337fd11d73a6e8352526102843580955260608160075afa9210161660408260808160065afa169051915190156108865760405191610100600484377f1ac06d0b9a09ab713f1b41579a8d949a1fbcccacc3fac0df093c0c2f3b4d5a3a6101008401527f22642309db8a674fa7eeba054d243283cc19a3c56a0ee905d9c23961dcb2e28a6101208401527f10f53c0711c493f7c3648166a62550aeb7f5a256b628156a3f1ca84877c6f5236101408401527f045473b5c4b6f07ea986657c17d8407a72cab51459b048235799ef3f9ea4c47b6101608401527f0c5c0484ed474fb6499f3503536db0326a0984cd3a2e48a0ba7467b91a71a2fc6101808401527f013658d47eda956c6ab310da69fcc5ea9ea1bd6d631c0a94d96e7b9e6072362d6101a08401527f06c65e9a92419ea93f665280659f6ee9f9991d4900fcbc963bfbfb5f4e5567e06101c08401527f23042c2b56c20e2ad87d083e47831d44791b4b2e00486409edec4ad09162b5376101e08401527f18125d661a8ca05a5e6eb00e0612ea501ff034d472912138e0c1d77968e409416102008401527f27e30aee102e3617623e41824f4cec5cd2657a295305dcf29a42c9c1d05b41116102208401526102408301526102608201527f0f2c9d069849e24c63762ab3f1c3734d05ff194af6084460ddccc2fc754c98306102808201527f2d38807db56ab03f779d6b96af0a3e7415da8c00658725a09153175a653e8aab6102a08201527f2f570c72b206b40b9158a7ff62675041bcd2ea2713b02bc64373ead028800b676102c08201527ef3fa1b69fe7bbd3da387b68a7793f00ab1e5fb278a3f8d545d69bb950071496102e08201526020816103008160085afa9051161561087757005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b5f80fd5b346108955760403660031901126108955760043567ffffffffffffffff8111610895576108ca90369060040161127a565b6024359167ffffffffffffffff8311610895576108ee6108f693369060040161127a565b929091611359565b005b346108955761010036600319011261089557366101041161089557608060405161092282826112a8565b813682376109346024356004356116a5565b815261094a60843560a435604435606435611746565b6020830152604082015261096260e43560c4356116a5565b60608201526109746040518092611253565bf35b34610895575f3660031901126108955760206040517fd1cd294544102bb194baa6a3255de8e23899df4c1f489c4d45c23b84bb44af438152f35b34610895576102203660031901126108955736608411610895573661022411610895576103006040516109e382826112a8565b813682376109f2600435611501565b610a0360249392933560443561156c565b91939290610a12606435611501565b9390926040519660408801967f212f8c9c30ef5086d167ac70a3ecc791344712937cbe9c81eace3670d6e360e789528860208101987f20975d217b2593dc2d06db9d9321637d4bd87e7a4a61df50274e1209e85c61bc8a527f11b8de7782bdce83a93e0c6c82b8b272754c9c82b88c39cf41e4c4e0f255b78881527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f10075c2125e499fed523d45f169b6bd5e188e5e66047e6156a18bc49221ec25284527f01f0763968c3953c2dbcd7ea35d7929756dab99fe380c4e1a5f1d4056d5a7baf6084359583608082019780895286828660608160075afa911016818360808160065afa167f0e9e6a2d38b98b57c84a1b640232bf488876d4adb099fc144616143166fd6c5f85527f0fac8283701ad712a895d2068f8bbdee9bf64306086dcb32c0362d54654fd6a18852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f2c9206f6ce1cf15057bf9935d9261b97141f8b2cf65e6ac3c4829c5ebb69830885527f06f097cedc499659b999dff576f1fb734d9251692b0e5f13c7523b0d21ea109d885260c43590818a5287838760608160075afa92101616818360808160065afa167f2210627388a6807e437cebf2337148e61f0460795428cba7650615140272985985527f085009bc331a387b4ff180326fba6b47553780388efc116869ec6f27353e7dd7885260e43590818a5287838760608160075afa92101616818360808160065afa167f23309a37e21dcc1d538a00d9f6f03aa8936d5023553a9d4226829bdc1d00183785527f04d5c862ed6de141828069b89cf8d38f57c0d859de822fe8f11388324786fd3388526101043590818a5287838760608160075afa92101616818360808160065afa167f24f5ce0636930fb138a43e685bde383ba9e8e8435a89e5a114490fbb3af7ce9d85527f12bfafd11848aa5cb0a92234c8d8df9e1e5ff438a3bb70ed1444c8790ad6e64c88526101243590818a5287838760608160075afa92101616818360808160065afa167f1a2f1a9b1317f91f6815ef2117e8b627f3a643814d001aff9fcf2a9569a062e385527f129af51e7ce61debf6f24a7ffd4bcb835eb6bebe9014cc67c610b9400b72935888526101443590818a5287838760608160075afa92101616818360808160065afa167f0c20d9e7d3e8d87fb101247992d9f47d3840454b8a7d2b72e55d1f7f3d8ab22b85527f1a3fa3028bc373df10ed4a636319233ceb73d1d752964cd22ac1dffcbe8c1b8e88526101643590818a5287838760608160075afa92101616818360808160065afa167f0efa6c70ee38e2280259dd4e116c152cf5687b11ab84e6828eee215cb7d012ef85527f1cf38529fdb7d92f1d05c9a0a446e9ccaa0db67ebc8acc1fef93a14cad8ede4788526101843590818a5287838760608160075afa92101616818360808160065afa167f2ba8e1e81c526c7e5316b3ab913b6a40c49f42615925443c816133fd3503be6985527f15f72290598781df6625146de17bf30287aa585ceb41c64228ac4dd7b653045d88526101a43590818a5287838760608160075afa92101616818360808160065afa167f0c74b6ef876b3cf7f4bed3439e558faa21beba0ebc7cec2abac53a4572e6f9ca85527f1f90f3643d710458d2f8286750e32e51867d650962d5056a812a24be9036f06288526101c43590818a5287838760608160075afa92101616818360808160065afa167f0922ec9f45ec9cb595d0d979f796fd67c0bbd94160dbcb510830a2531c95805f85527f24d79b0bc1ebc43c74eb36c3b69cd186b0e7dcb6164fad94b3dfd8965e51d46d88526101e43590818a5287838760608160075afa921016169160808160065afa16947f27963007a92bc955699cb73d20b9db602fd7923f60f5158558e337fd11d73a6e8352526102043580955260608160075afa9210161660408a60808160065afa169851975198156108865760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f1ac06d0b9a09ab713f1b41579a8d949a1fbcccacc3fac0df093c0c2f3b4d5a3a6101008401527f22642309db8a674fa7eeba054d243283cc19a3c56a0ee905d9c23961dcb2e28a6101208401527f10f53c0711c493f7c3648166a62550aeb7f5a256b628156a3f1ca84877c6f5236101408401527f045473b5c4b6f07ea986657c17d8407a72cab51459b048235799ef3f9ea4c47b6101608401527f0c5c0484ed474fb6499f3503536db0326a0984cd3a2e48a0ba7467b91a71a2fc6101808401527f013658d47eda956c6ab310da69fcc5ea9ea1bd6d631c0a94d96e7b9e6072362d6101a08401527f06c65e9a92419ea93f665280659f6ee9f9991d4900fcbc963bfbfb5f4e5567e06101c08401527f23042c2b56c20e2ad87d083e47831d44791b4b2e00486409edec4ad09162b5376101e08401527f18125d661a8ca05a5e6eb00e0612ea501ff034d472912138e0c1d77968e409416102008401527f27e30aee102e3617623e41824f4cec5cd2657a295305dcf29a42c9c1d05b41116102208401526102408301526102608201527f0f2c9d069849e24c63762ab3f1c3734d05ff194af6084460ddccc2fc754c98306102808201527f2d38807db56ab03f779d6b96af0a3e7415da8c00658725a09153175a653e8aab6102a08201527f2f570c72b206b40b9158a7ff62675041bcd2ea2713b02bc64373ead028800b676102c08201527ef3fa1b69fe7bbd3da387b68a7793f00ab1e5fb278a3f8d545d69bb950071496102e082015260405192839161122e84846112a8565b8336843760085afa15908115611246575b5061087757005b600191505114158161123f565b905f905b6004821061126457505050565b6020806001928551815201930191019091611257565b9181601f840112156108955782359167ffffffffffffffff8311610895576020838186019501011161089557565b90601f8019910116810190811067ffffffffffffffff8211176112ca57604052565b634e487b7160e01b5f52604160045260245ffd5b906101a0828203126108955780601f8301121561089557604051916113056101a0846112a8565b82906101a0810192831161089557905b8282106113225750505090565b8135815260209182019101611315565b905f905b600d821061134357505050565b6020806001928551815201930191019091611336565b93929190610100811461143b576080811461137d5763236bd13760e01b5f5260045ffd5b6101a0830361142c578401916080858403126108955782601f8601121561089557604051926113ad6080856112a8565b83956080810191821161089557955b81871061141c57505061141a939450816113db916114049301906112de565b6040516307846db360e21b6020820152926113fa906024850190611253565b60a4830190611332565b6102248152611415610244826112a8565b61197a565b565b86358152602096870196016113bc565b630c0b7e3560e11b5f5260045ffd5b6101a0830361142c57840191610100858403126108955782601f86011215610895576040519261146d610100856112a8565b8395610100810191821161089557955b8187106114f1575050611495929394508101906112de565b60405163f7f338eb60e01b6020820152915f602484015b600882106114db57505050906114ca61141a92610124830190611332565b6102a481526114156102c4826112a8565b60208060019285518152019301910190916114ac565b863581526020968701960161147d565b8015611565578060011c915f516020611bb75f395f51905f52831015610877576001806115445f516020611bb75f395f51905f52600381888181800909086119d7565b93161461154d57565b905f516020611bb75f395f51905f5280910681030690565b505f905f90565b80158061169d575b611691578060021c92825f516020611bb75f395f51905f52851080159061167a575b6108775784815f516020611bb75f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816116449d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5086119fa565b80929160018082961614611656575050565b5f516020611bb75f395f51905f528093945080929550809106810306930681030690565b505f516020611bb75f395f51905f52811015611596565b50505f905f905f905f90565b508115611574565b905f516020611bb75f395f51905f52821080159061172f575b61087757811580611727575b611721576116ee5f516020611bb75f395f51905f52600381858181800909086119d7565b8181036116fd57505060011b90565b5f516020611bb75f395f51905f52809106810306145f1461087757600190811b1790565b50505f90565b5080156116ca565b505f516020611bb75f395f51905f528110156116be565b919093925f516020611bb75f395f51905f528310801590611963575b801561194c575b8015611935575b61087757808286851717171561192a5790829161188d5f516020611bb75f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611bb75f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161186781808b800981878009086119d7565b8408095f516020611bb75f395f51905f5261188182611b4e565b800914159586916119fa565b929080821480611921575b156118bf5750505050905f146118b75760ff60025b169060021b179190565b60ff5f6118ad565b5f516020611bb75f395f51905f52809106810306149182611902575b50501561087757600191156118fa5760ff60025b169060021b17179190565b60ff5f6118ef565b5f516020611bb75f395f51905f52919250819006810306145f806118db565b50838314611898565b50505090505f905f90565b505f516020611bb75f395f51905f52811015611770565b505f516020611bb75f395f51905f52821015611769565b505f516020611bb75f395f51905f52851015611762565b5f8091602081519101305afa3d156119cf573d9067ffffffffffffffff82116112ca57604051916119b5601f8201601f1916602001846112a8565b82523d5f602084013e5b156119c75750565b602081519101fd5b6060906119bf565b906119e182611b4e565b915f516020611bb75f395f51905f528380090361087757565b915f516020611bb75f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611a5293969496611a4482808a8009818a8009086119d7565b90611b42575b8608096119d7565b925f516020611bb75f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611bb75f395f51905f5260a083015260208260c08160055afa91519115610877575f516020611bb75f395f51905f52826001920903610877575f516020611bb75f395f51905f52908209925f516020611bb75f395f51905f528080808780090681030681878009081490811591611b23575b5061087757565b90505f516020611bb75f395f51905f528084860960020914155f611b1c565b81809106810306611a4a565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611bb75f395f51905f5260a083015260208260c08160055afa915191156108775756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a26469706673582212202779ad5d448ca5757883ec911a6908250c87294c3a9ce0b83ab9fa596071941e64736f6c634300081c0033",
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

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x1e11b6cc.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[13] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [13]*big.Int) error {
	var out []interface{}
	err := _DecryptCombineVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x1e11b6cc.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[13] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [13]*big.Int) error {
	return _DecryptCombineVerifier.Contract.VerifyCompressedProof(&_DecryptCombineVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x1e11b6cc.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[13] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [13]*big.Int) error {
	return _DecryptCombineVerifier.Contract.VerifyCompressedProof(&_DecryptCombineVerifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCaller) VerifyProof(opts *bind.CallOpts, proof []byte, input []byte) error {
	var out []interface{}
	err := _DecryptCombineVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierSession) VerifyProof(proof []byte, input []byte) error {
	return _DecryptCombineVerifier.Contract.VerifyProof(&_DecryptCombineVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerSession) VerifyProof(proof []byte, input []byte) error {
	return _DecryptCombineVerifier.Contract.VerifyProof(&_DecryptCombineVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xf7f338eb.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[13] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCaller) VerifyProof0(opts *bind.CallOpts, proof [8]*big.Int, input [13]*big.Int) error {
	var out []interface{}
	err := _DecryptCombineVerifier.contract.Call(opts, &out, "verifyProof0", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xf7f338eb.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[13] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierSession) VerifyProof0(proof [8]*big.Int, input [13]*big.Int) error {
	return _DecryptCombineVerifier.Contract.VerifyProof0(&_DecryptCombineVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xf7f338eb.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[13] input) view returns()
func (_DecryptCombineVerifier *DecryptCombineVerifierCallerSession) VerifyProof0(proof [8]*big.Int, input [13]*big.Int) error {
	return _DecryptCombineVerifier.Contract.VerifyProof0(&_DecryptCombineVerifier.CallOpts, proof, input)
}
