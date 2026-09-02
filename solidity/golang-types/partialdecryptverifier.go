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
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[15]\",\"internalType\":\"uint256[15]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"input\",\"type\":\"uint256[15]\",\"internalType\":\"uint256[15]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidProofEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x60808060405234601557611b4f908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c8063216df43414610878578063233ace111461083e57806344f63692146107a5578063b8e72af6146105ea5763c6acb0f114610051575f80fd5b346105e7576102e03660031901126105e75736610104116105e757366102e4116105e75760405160408101905f516020611a7a5f395f51905f52815260208101915f516020611a5a5f395f51905f5283525f51602061157a5f395f51905f528152606082015f51602061185a5f395f51905f5281525f51602061187a5f395f51905f52604061010435935f51602061181a5f395f51905f52608087019580875284848460608160075afa911016838860808160065afa165f5160206116fa5f395f51905f5283525f51602061195a5f395f51905f5286526001610124359182895286868660608160075afa9310161616838860808160065afa165f51602061151a5f395f51905f5283525f516020611ada5f395f51905f528652610144359081885285858560608160075afa92101616838860808160065afa165f51602061175a5f395f51905f5283525f51602061177a5f395f51905f528652610164359081885285858560608160075afa92101616838860808160065afa165f51602061149a5f395f51905f5283525f5160206117ba5f395f51905f528652610184359081885285858560608160075afa92101616838860808160065afa165f51602061169a5f395f51905f5283525f51602061179a5f395f51905f5286526101a4359081885285858560608160075afa92101616838860808160065afa165f5160206118fa5f395f51905f5283525f51602061155a5f395f51905f5286526101c4359081885285858560608160075afa92101616838860808160065afa165f51602061191a5f395f51905f5283525f51602061161a5f395f51905f5286526101e4359081885285858560608160075afa92101616838860808160065afa165f51602061165a5f395f51905f5283525f5160206119fa5f395f51905f528652610204359081885285858560608160075afa92101616838860808160065afa165f516020611a3a5f395f51905f5283525f5160206119ba5f395f51905f528652610224359081885285858560608160075afa92101616838860808160065afa165f51602061199a5f395f51905f5283525f5160206116ba5f395f51905f528652610244359081885285858560608160075afa92101616838860808160065afa165f5160206115ba5f395f51905f5283525f51602061153a5f395f51905f528652610264359081885285858560608160075afa92101616838860808160065afa165f51602061173a5f395f51905f5283525f5160206118da5f395f51905f528652610284359081885285858560608160075afa92101616838860808160065afa165f5160206118ba5f395f51905f5283525f51602061163a5f395f51905f5286526102a4359081885285858560608160075afa92101616838860808160065afa16945f5160206114da5f395f51905f528352526102c43580955260608160075afa9210161660408260808160065afa169051915190156105d85760405191610100600484375f51602061189a5f395f51905f526101008401525f516020611a1a5f395f51905f526101208401525f51602061193a5f395f51905f526101408401525f5160206117da5f395f51905f526101608401525f5160206115fa5f395f51905f526101808401525f51602061167a5f395f51905f526101a08401525f516020611aba5f395f51905f526101c08401525f51602061183a5f395f51905f526101e08401525f516020611a9a5f395f51905f526102008401525f5160206117fa5f395f51905f526102208401526102408301526102608201525f5160206116da5f395f51905f526102808201525f5160206114fa5f395f51905f526102a08201525f51602061159a5f395f51905f526102c08201525f51602061197a5f395f51905f526102e08201526020816103008160085afa905116156105c95780f35b631ff3747d60e21b8152600490fd5b63a54f8e2760e01b8352600483fd5b80fd5b503461075e57604036600319011261075e576004356001600160401b03811161075e5761061b903690600401610e89565b6024356001600160401b03811161075e5761063a903690600401610e89565b909161010081036107965783016101008482031261075e5780601f8501121561075e576040519361066d61010086610eb6565b8490610100810192831161075e57905b8282106107865750505081016101e08282031261075e5780601f8301121561075e57604051916106af6101e084610eb6565b82906101e0810192831161075e57905b82821061076257505050303b1561075e5760405163c6acb0f160e01b8152915f600484015b600882106107485750505061010482015f905b600f8210610732575050505f816102e481305afa801561072757610719575080f35b61072591505f90610eb6565b005b6040513d5f823e3d90fd5b60208060019285518152019301910190916106f7565b60208060019285518152019301910190916106e4565b5f80fd5b81358152602091820191016106bf565b634e487b7160e01b5f52604160045260245ffd5b813581526020918201910161067d565b63236bd13760e01b5f5260045ffd5b3461075e5761010036600319011261075e57366101041161075e576040516107ce608082610eb6565b60803682376107e1602435600435611044565b81526107f760843560a4356044356064356110e5565b6020830152604082015261080f60e43560c435611044565b6060820152604051905f825b6004821061082857608084f35b602080600192855181520193019101909161081b565b3461075e575f36600319011261075e5760206040517f9c536a045045acc6d2bc25066b8af160f7b2cc9d6affb3f1a5fe6a8b4299b1a18152f35b3461075e5761026036600319011261075e573660841161075e57366102641161075e576103006040516108ab8282610eb6565b813682376108ba600435610ed9565b6108cb602493929335604435610f44565b919392906108da606435610ed9565b9390926040519660408801965f516020611a7a5f395f51905f5289528860208101985f516020611a5a5f395f51905f528a525f51602061157a5f395f51905f5281525f51602061187a5f395f51905f52604060608401925f51602061185a5f395f51905f5284525f51602061181a5f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f5160206116fa5f395f51905f5285525f51602061195a5f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f51602061151a5f395f51905f5285525f516020611ada5f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f51602061175a5f395f51905f5285525f51602061177a5f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f51602061149a5f395f51905f5285525f5160206117ba5f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f51602061169a5f395f51905f5285525f51602061179a5f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f5160206118fa5f395f51905f5285525f51602061155a5f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f51602061191a5f395f51905f5285525f51602061161a5f395f51905f5288526101643590818a5287838760608160075afa92101616818360808160065afa165f51602061165a5f395f51905f5285525f5160206119fa5f395f51905f5288526101843590818a5287838760608160075afa92101616818360808160065afa165f516020611a3a5f395f51905f5285525f5160206119ba5f395f51905f5288526101a43590818a5287838760608160075afa92101616818360808160065afa165f51602061199a5f395f51905f5285525f5160206116ba5f395f51905f5288526101c43590818a5287838760608160075afa92101616818360808160065afa165f5160206115ba5f395f51905f5285525f51602061153a5f395f51905f5288526101e43590818a5287838760608160075afa92101616818360808160065afa165f51602061173a5f395f51905f5285525f5160206118da5f395f51905f5288526102043590818a5287838760608160075afa92101616818360808160065afa165f5160206118ba5f395f51905f5285525f51602061163a5f395f51905f5288526102243590818a5287838760608160075afa921016169160808160065afa16945f5160206114da5f395f51905f528352526102443580955260608160075afa9210161660408a60808160065afa16985197519815610e7a5760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f51602061189a5f395f51905f526101008401525f516020611a1a5f395f51905f526101208401525f51602061193a5f395f51905f526101408401525f5160206117da5f395f51905f526101608401525f5160206115fa5f395f51905f526101808401525f51602061167a5f395f51905f526101a08401525f516020611aba5f395f51905f526101c08401525f51602061183a5f395f51905f526101e08401525f516020611a9a5f395f51905f526102008401525f5160206117fa5f395f51905f526102208401526102408301526102608201525f5160206116da5f395f51905f526102808201525f5160206114fa5f395f51905f526102a08201525f51602061159a5f395f51905f526102c08201525f51602061197a5f395f51905f526102e0820152604051928391610e468484610eb6565b8336843760085afa15908115610e6d575b50610e5e57005b631ff3747d60e21b5f5260045ffd5b6001915051141581610e57565b63a54f8e2760e01b5f5260045ffd5b9181601f8401121561075e578235916001600160401b03831161075e576020838186019501011161075e57565b601f909101601f19168101906001600160401b0382119082101761077257604052565b8015610f3d578060011c915f5160206115da5f395f51905f52831015610e5e57600180610f1c5f5160206115da5f395f51905f52600381888181800909086112cd565b931614610f2557565b905f5160206115da5f395f51905f5280910681030690565b505f905f90565b80158061103c575b611030578060021c92825f5160206115da5f395f51905f528510801590611019575b610e5e5784815f5160206115da5f395f51905f5280808080808080805f5160206119da5f395f51905f5281610fe39d8d0909998a0981898181800909085f5160206114ba5f395f51905f520806810306936002808a16149509818a8181800909085f516020611afa5f395f51905f52086112f0565b80929160018082961614610ff5575050565b5f5160206115da5f395f51905f528093945080929550809106810306930681030690565b505f5160206115da5f395f51905f52811015610f6e565b50505f905f905f905f90565b508115610f4c565b905f5160206115da5f395f51905f5282108015906110ce575b610e5e578115806110c6575b6110c05761108d5f5160206115da5f395f51905f52600381858181800909086112cd565b81810361109c57505060011b90565b5f5160206115da5f395f51905f52809106810306145f14610e5e57600190811b1790565b50505f90565b508015611069565b505f5160206115da5f395f51905f5281101561105d565b919093925f5160206115da5f395f51905f5283108015906112b6575b801561129f575b8015611288575b610e5e57808286851717171561127d579082916111e05f5160206115da5f395f51905f5280808080888180808f9d5f5160206119da5f395f51905f528f839290839109099d8e0981848181800909085f516020611afa5f395f51905f52089a09818c8181800909085f5160206114ba5f395f51905f520806810306945f5160206115da5f395f51905f525f51602061171a5f395f51905f52816111ba81808b800981878009086112cd565b8408095f5160206115da5f395f51905f526111d482611431565b800914159586916112f0565b929080821480611274575b156112125750505050905f1461120a5760ff60025b169060021b179190565b60ff5f611200565b5f5160206115da5f395f51905f52809106810306149182611255575b505015610e5e576001911561124d5760ff60025b169060021b17179190565b60ff5f611242565b5f5160206115da5f395f51905f52919250819006810306145f8061122e565b508383146111eb565b50505090505f905f90565b505f5160206115da5f395f51905f5281101561110f565b505f5160206115da5f395f51905f52821015611108565b505f5160206115da5f395f51905f52851015611101565b906112d782611431565b915f5160206115da5f395f51905f5283800903610e5e57565b915f5160206115da5f395f51905f525f51602061171a5f395f51905f52816113359396949661132782808a8009818a8009086112cd565b90611425575b8608096112cd565b925f5160206115da5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206115da5f395f51905f5260a083015260208260c08160055afa91519115610e5e575f5160206115da5f395f51905f52826001920903610e5e575f5160206115da5f395f51905f52908209925f5160206115da5f395f51905f528080808780090681030681878009081490811591611406575b50610e5e57565b90505f5160206115da5f395f51905f528084860960020914155f6113ff565b8180910681030661132d565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206115da5f395f51905f5260a083015260208260c08160055afa91519115610e5e5756fe048cc2aea7ba94b63103bb252e525dd9fb0f2d89f74651441b89a2b7265742fa2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775203b348a6be9032f9d5d8e7bb27cd3384eb992ec321a9bf6bcc3a220c93912bc0348c35b08a6718399c76afff0ab8aaec4531469f42b00c957dd85142d459ef525e530d90e9de6eebbcc4ab5b6daf4c7516a0d0651000555c6c829d13be8a8fd2f8015cf761ba00f506f42f5fe79d7617b5d31c2c55388e19d0cf4a723f1196917ba62be799d03512806b154ef610565f4740a452b5de2bbe7eacba33271462c25f1bef57fac29ecc9b71bd9295fdff008b2121110adedcaa65176f6a50dec6902fcfe5f2b58fe97cf3aeb551357a23de2fec130e344030cff2a4637cdff5e6b0dfd559c2dc180be59c4f72bdd0e0c6844dd5ba29999cf39e369d26388f9032330644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd471aa1a6fcbda7a76886b037f8b35c42a97488fa5447ca649dff26dcffa7072d9518b0ff1cf18de681e23a0917c71284c4f0c084e61d3fd85242daafb820a01d7907ec6b4be5898500b9fea31c070fb00474a1a8ab54b5e88758d2112f44ebfa5023745a878a21de544bc06cdd8d134f7750ec17d99a82c101b1af05347dafb0221704ebdfd76df9cb7fa1d9e9f2b37da324f4ec964735b72aafe318241f765b5d0850983fcd28966ef1d14944877aeed4e672cce5b97648b9e94e48c621b1aedb06c52a57b73e9ef8313ae073ddf0363b7aa9ab6172a201bffe16181978c85faf18e755f7506f1596862f49685133d6aa0da9d3e453367228f3ae707ebb5463b30f0f87c221b7f2ec99bdf5fc946ef96cb95be706564443a7226acfc2a55c37c0183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea42c284334fb827cbfc713d190d09d61882b9a00c58eb1da0e93eea76ff490e97c15acf830ac12c262e48571f47b5b86b1b8e4aeaee93fb51b99eb205ba7e458030995472ba6ced9edc2a5908a49f984c729a54c6810f864697972c0986312fbb907ab8446e2a2873daa4db5d431c81c279bae74d5174a7d82e901aa3e6d0666d506099c963f19069f43e925f95b061a951461c84e75546eea31c1a9dafe2cd31f06752fa5a47b224351c7b5ede602dc92e14ee97c47794a1d1a35aeb50116600b2d1ebbd4dc63369a54e80a7d1dd6cf802b2e23f9c992e313160625b56d9162ad243208a4731d876b4a0810569d845687f92b69f194b229e013b71198e325b09f14be2e95afbc7c50468fdc2f09cb482e7a85798b9f5b9ed619fbc0c0f8e3e76e195cdd179249b85728a7ea7503b1f32fddce0ad9e6a435b0a7f346c747ad6b9030644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000012b4c399d22d0f8ddca59a74249b3814b1f4fdb10819145b632d2c16fc0902aa81d4f5fa194e66eab4f8d50fc45619c2978060caef06dd5aeae4a31e1b489c9f9077343380e80a8c27a5f65b44d2bf145c09067065b176650f3b5993977238b4f078e400837fc2b85a7f25010fc1a81d6e78d1b275c3cbdf0ad5b5cf0d9c7308004d11b2438af726503af92df948ea79e345226e3a8c4fe2a18e10b4691e0a0721b4f610c86011e883949dcf288f3970e8119cb0627d91d633e13272f2a09b3112960161454342129be154f5b7a2073706ce37a61a66202d9f7542265e7661d832d03cab51edc5b349edfcc189834162a209928d68464cb4c31200b94ca7baf162f3b8803604e8012dc3f8c6661823c5a0831a7d1a5b33f50cdcd12d5a503744907f3182ae61e34dd3344e23b6f086c80cdbed1a4b01143b680f3d01d4fcf7b2130644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd442acb4773e47884de7e31b9d65e59c8eca9e2d989322922d3f881af0ba50e0df8020364fd54991a08a5b0af7a768fdf9d59142c867bda0222c32d613e8082e97701fa63811496ec6cabf107946a8b33de3d89924f02357ca853d3019e7785005a10f1d510ca2896cc7f9ff5cf1b627260cec1820e8c4cf28d9bc8b6de079b9e4e0c8f5bb9acd2a66f74041af93595151568970fb66567139ee41df7eed6dbf2f2008560d22d7173b99d80031e9e322c3bf7db812b1fbf9547442cbad1a21d63c12cc9ce5dafb1a959f49ea94eff85dc4458b7057f76770447b80c0f35e30f5b8a19e66bb7949cdcb7faf779c88e04267a7cad2fcb8a6e80264e7a2dc4327e0d442b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5a264697066735822122059e8357fe7f4abb06872e20b0f4b562ce9939cbf9167cf6933b872388aa208b864736f6c634300081c0033",
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

// VerifyProof0 is a free data retrieval call binding the contract method 0xc6acb0f1.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[15] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCaller) VerifyProof0(opts *bind.CallOpts, proof [8]*big.Int, input [15]*big.Int) error {
	var out []interface{}
	err := _PartialDecryptVerifier.contract.Call(opts, &out, "verifyProof0", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xc6acb0f1.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[15] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierSession) VerifyProof0(proof [8]*big.Int, input [15]*big.Int) error {
	return _PartialDecryptVerifier.Contract.VerifyProof0(&_PartialDecryptVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xc6acb0f1.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[15] input) view returns()
func (_PartialDecryptVerifier *PartialDecryptVerifierCallerSession) VerifyProof0(proof [8]*big.Int, input [15]*big.Int) error {
	return _PartialDecryptVerifier.Contract.VerifyProof0(&_PartialDecryptVerifier.CallOpts, proof, input)
}
