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
	Bin: "0x60808060405234601557611c14908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace1114610ed65750806344f6369214610e3d5780634605cb8914610868578063b8e72af6146106ad5763da3496ab14610055575f80fd5b346106aa576102803660031901126106aa57366084116106aa5736610284116106aa57604051906103006100898184610f3b565b803684376100986004356111e7565b6100a9602495929535604435611252565b919392906100b86064356111e7565b9390926040519660408801965f5160206115df5f395f51905f5289528860208101985f51602061189f5f395f51905f528a525f516020611aff5f395f51905f5281525f51602061185f5f395f51905f52604060608401925f51602061175f5f395f51905f5284525f51602061167f5f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f51602061193f5f395f51905f5285525f51602061191f5f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f5160206118df5f395f51905f5285525f51602061179f5f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f516020611adf5f395f51905f5285525f5160206115ff5f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f516020611b1f5f395f51905f5285525f516020611a7f5f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f516020611bbf5f395f51905f5285525f5160206119ff5f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f51602061159f5f395f51905f5285525f516020611abf5f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f51602061155f5f395f51905f5285525f51602061187f5f395f51905f5288526101643590818a5287838760608160075afa92101616818360808160065afa165f5160206119df5f395f51905f5285525f51602061153f5f395f51905f5288526101843590818a5287838760608160075afa92101616818360808160065afa165f51602061165f5f395f51905f5285525f5160206116bf5f395f51905f5288526101a43590818a5287838760608160075afa92101616818360808160065afa165f51602061171f5f395f51905f5285525f51602061183f5f395f51905f5288526101c43590818a5287838760608160075afa92101616818360808160065afa165f51602061197f5f395f51905f5285525f51602061177f5f395f51905f5288526101e43590818a5287838760608160075afa92101616818360808160065afa165f516020611b7f5f395f51905f5285525f51602061199f5f395f51905f5288526102043590818a5287838760608160075afa92101616818360808160065afa165f5160206117df5f395f51905f5285525f5160206118ff5f395f51905f5288526102243590818a5287838760608160075afa92101616818360808160065afa165f51602061195f5f395f51905f5285525f51602061163f5f395f51905f5288526102443590818a5287838760608160075afa921016169160808160065afa16945f516020611a1f5f395f51905f528352526102643580955260608160075afa9210161660408a60808160065afa1698519751981561069b5760209a9b9c8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f5160206117ff5f395f51905f526101008401525f51602061173f5f395f51905f526101208401525f5160206116ff5f395f51905f526101408401525f5160206116df5f395f51905f526101608401525f5160206119bf5f395f51905f526101808401525f516020611b9f5f395f51905f526101a08401525f5160206118bf5f395f51905f526101c08401525f51602061169f5f395f51905f526101e08401525f51602061157f5f395f51905f526102008401525f5160206115bf5f395f51905f526102208401526102408301526102608201525f51602061181f5f395f51905f526102808201525f516020611a3f5f395f51905f526102a08201525f516020611b3f5f395f51905f526102c08201525f516020611a9f5f395f51905f526102e08201526040519283916106668484610f3b565b8336843760085afa1590811561068e575b5061067f5780f35b631ff3747d60e21b8152600490fd5b600191505114155f610677565b63a54f8e2760e01b8c5260048cfd5b80fd5b5034610821576040366003190112610821576004356001600160401b038111610821576106de903690600401610f0e565b6024356001600160401b038111610821576106fd903690600401610f0e565b90916101008103610859578301610100848203126108215780601f85011215610821576040519361073061010086610f3b565b8490610100810192831161082157905b828210610849575050508101610200828203126108215780601f83011215610821576040519161077261020084610f3b565b8290610200810192831161082157905b82821061082557505050303b1561082157604051634605cb8960e01b8152915f600484015b6008821061080b5750505061010482015f905b601082106107f5575050505f8161030481305afa80156107ea576107dc575080f35b6107e891505f90610f3b565b005b6040513d5f823e3d90fd5b60208060019285518152019301910190916107ba565b60208060019285518152019301910190916107a7565b5f80fd5b8135815260209182019101610782565b634e487b7160e01b5f52604160045260245ffd5b8135815260209182019101610740565b63236bd13760e01b5f5260045ffd5b34610821576103003660031901126108215736610104116108215736610304116108215760405160408101905f5160206115df5f395f51905f52815260208101915f51602061189f5f395f51905f5283525f516020611aff5f395f51905f528152606082015f51602061175f5f395f51905f5281525f51602061185f5f395f51905f52604061010435935f51602061167f5f395f51905f52608087019580875284848460608160075afa911016838860808160065afa165f51602061193f5f395f51905f5283525f51602061191f5f395f51905f5286526001610124359182895286868660608160075afa9310161616838860808160065afa165f5160206118df5f395f51905f5283525f51602061179f5f395f51905f528652610144359081885285858560608160075afa92101616838860808160065afa165f516020611adf5f395f51905f5283525f5160206115ff5f395f51905f528652610164359081885285858560608160075afa92101616838860808160065afa165f516020611b1f5f395f51905f5283525f516020611a7f5f395f51905f528652610184359081885285858560608160075afa92101616838860808160065afa165f516020611bbf5f395f51905f5283525f5160206119ff5f395f51905f5286526101a4359081885285858560608160075afa92101616838860808160065afa165f51602061159f5f395f51905f5283525f516020611abf5f395f51905f5286526101c4359081885285858560608160075afa92101616838860808160065afa165f51602061155f5f395f51905f5283525f51602061187f5f395f51905f5286526101e4359081885285858560608160075afa92101616838860808160065afa165f5160206119df5f395f51905f5283525f51602061153f5f395f51905f528652610204359081885285858560608160075afa92101616838860808160065afa165f51602061165f5f395f51905f5283525f5160206116bf5f395f51905f528652610224359081885285858560608160075afa92101616838860808160065afa165f51602061171f5f395f51905f5283525f51602061183f5f395f51905f528652610244359081885285858560608160075afa92101616838860808160065afa165f51602061197f5f395f51905f5283525f51602061177f5f395f51905f528652610264359081885285858560608160075afa92101616838860808160065afa165f516020611b7f5f395f51905f5283525f51602061199f5f395f51905f528652610284359081885285858560608160075afa92101616838860808160065afa165f5160206117df5f395f51905f5283525f5160206118ff5f395f51905f5286526102a4359081885285858560608160075afa92101616838860808160065afa165f51602061195f5f395f51905f5283525f51602061163f5f395f51905f5286526102c4359081885285858560608160075afa92101616838860808160065afa16945f516020611a1f5f395f51905f528352526102e43580955260608160075afa9210161660408260808160065afa16905191519015610e2e5760405191610100600484375f5160206117ff5f395f51905f526101008401525f51602061173f5f395f51905f526101208401525f5160206116ff5f395f51905f526101408401525f5160206116df5f395f51905f526101608401525f5160206119bf5f395f51905f526101808401525f516020611b9f5f395f51905f526101a08401525f5160206118bf5f395f51905f526101c08401525f51602061169f5f395f51905f526101e08401525f51602061157f5f395f51905f526102008401525f5160206115bf5f395f51905f526102208401526102408301526102608201525f51602061181f5f395f51905f526102808201525f516020611a3f5f395f51905f526102a08201525f516020611b3f5f395f51905f526102c08201525f516020611a9f5f395f51905f526102e08201526020816103008160085afa90511615610e1f57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b346108215761010036600319011261082157366101041161082157604051610e66608082610f3b565b6080368237610e79602435600435610f5e565b8152610e8f60843560a435604435606435610fff565b60208301526040820152610ea760e43560c435610f5e565b6060820152604051905f825b60048210610ec057608084f35b6020806001928551815201930191019091610eb3565b34610821575f36600319011261082157807f132b2e301c3d3b4c8914546c41dd20329d6fe6bab58a96faeffa5d6e48fec5d060209252f35b9181601f84011215610821578235916001600160401b038311610821576020838186019501011161082157565b601f909101601f19168101906001600160401b0382119082101761083557604052565b905f51602061161f5f395f51905f528210801590610fe8575b610e1f57811580610fe0575b610fda57610fa75f51602061161f5f395f51905f5260038185818180090908611352565b818103610fb657505060011b90565b5f51602061161f5f395f51905f52809106810306145f14610e1f57600190811b1790565b50505f90565b508015610f83565b505f51602061161f5f395f51905f52811015610f77565b919093925f51602061161f5f395f51905f5283108015906111d0575b80156111b9575b80156111a2575b610e1f578082868517171715611197579082916110fa5f51602061161f5f395f51905f5280808080888180808f9d5f516020611a5f5f395f51905f528f839290839109099d8e0981848181800909085f516020611b5f5f395f51905f52089a09818c8181800909085f51602061151f5f395f51905f520806810306945f51602061161f5f395f51905f525f5160206117bf5f395f51905f52816110d481808b80098187800908611352565b8408095f51602061161f5f395f51905f526110ee826114b6565b80091415958691611375565b92908082148061118e575b1561112c5750505050905f146111245760ff60025b169060021b179190565b60ff5f61111a565b5f51602061161f5f395f51905f5280910681030614918261116f575b505015610e1f57600191156111675760ff60025b169060021b17179190565b60ff5f61115c565b5f51602061161f5f395f51905f52919250819006810306145f80611148565b50838314611105565b50505090505f905f90565b505f51602061161f5f395f51905f52811015611029565b505f51602061161f5f395f51905f52821015611022565b505f51602061161f5f395f51905f5285101561101b565b801561124b578060011c915f51602061161f5f395f51905f52831015610e1f5760018061122a5f51602061161f5f395f51905f5260038188818180090908611352565b93161461123357565b905f51602061161f5f395f51905f5280910681030690565b505f905f90565b80158061134a575b61133e578060021c92825f51602061161f5f395f51905f528510801590611327575b610e1f5784815f51602061161f5f395f51905f5280808080808080805f516020611a5f5f395f51905f52816112f19d8d0909998a0981898181800909085f51602061151f5f395f51905f520806810306936002808a16149509818a8181800909085f516020611b5f5f395f51905f5208611375565b80929160018082961614611303575050565b5f51602061161f5f395f51905f528093945080929550809106810306930681030690565b505f51602061161f5f395f51905f5281101561127c565b50505f905f905f905f90565b50811561125a565b9061135c826114b6565b915f51602061161f5f395f51905f5283800903610e1f57565b915f51602061161f5f395f51905f525f5160206117bf5f395f51905f52816113ba939694966113ac82808a8009818a800908611352565b906114aa575b860809611352565b925f51602061161f5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f51602061161f5f395f51905f5260a083015260208260c08160055afa91519115610e1f575f51602061161f5f395f51905f52826001920903610e1f575f51602061161f5f395f51905f52908209925f51602061161f5f395f51905f52808080878009068103068187800908149081159161148b575b50610e1f57565b90505f51602061161f5f395f51905f528084860960020914155f611484565b818091068103066113b2565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f51602061161f5f395f51905f5260a083015260208260c08160055afa91519115610e1f5756fe2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775023427a61bdbd3e08c8e42f68c74f06531dd55d3c263c3da0e68f082d91510610bf23b14ef5939c72eeced381e7d9e5c7766acedea1c3365b7103c141471c4930b1247d06087a876ac038110160c2053c690f4c39f9fee8d4cf823976aedbced134e7598bad41c70d9c22b1735a1eb4cb597ed82ccfe4ac68c0c384889af39772fd29140fcbb7346062bbbaed766654f75d43af066e69e359519296534cf18491a33e5a8f8c464fc82ba1524b3477f6775ea0c38892d450f014356106ef07a0429a887da234a73576cc21f8b987053018b39dabbfcc20641d2e6d2e8f49f3b2130644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd471add9653bf6c1fce4a2f349bf7bb0b956e3891085838c6e3993545a2d833f21814ae747f99b10834ad649e6f5660a05b74b0cdef47c9e439ecb4b67ab066d65f0440d8ec4ae28f97edeae3ce1899b1cc2995de8ad6d9a64643dbd1c102ed1e7206c6673d3b7ddedb3a24402fce864f92ef7d0c3e88ffd26ce7d29a4b0cc1e6ec2792d5a18c0cfdc3e946f3b18836a541f16e6af17a5a55167dfc4c4171ea0c2921b90c21b0fb5c55cb1d967525d79eb8dfb54e1ac018cecfadbd500621df95a50010a1ffc564e8b381c04f4fdf7c909005f0ae6544c538395060cc85d8bbd132034a63967a949774f8fb0c245d79c212136cf17358d4fa839a53f1374af4947e09cb97ef1785343115cbb9121690d42c72d84c8b7c05407a8a992fb3db15be832d8a45c6afe5eef874d6cf990c5988f5a63ed21d45c7729c214b3c45585fb4621084e37a108d0c9af35d3def0242a365e089a565a3aaea7e22febf32b162a5ff23425c0cf8ad06e06e7f3fb014ded74afb690454d715a434f2da7f303ea6fc66183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea40bd6095c2aa165b00d3dc3bc36588e2b29c326952ad89eb0428d620a583f8f770baa6b2844a3ba8fff420ec1dcb22eab84b97fa8699ae44e82bcbca00433844f26f0a8344563e83592a11b017b658b1aa01871416ae300c0381cccc82c2cba5425a24d5b7d47641d22bd1c147035eb26b3e9569e0d497315fb5014d925d14bac30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001169dbd6b0d19cf6356b51b2c9e1d32bf8ae067748f99ba9b879d50380388b9040a8be9043113a88ba015ef47690bbc6ca09f3c753ff067c6db3243ba0e6c727429178149d5d19799e6d59e5b789d4027252d5d3ec25abd1085a5cc098375e6e7081c4a9a03ee2b0a5d7779b0a574d7274c1149ccbfaf771e4defd563b3354695272fd18787da08e5727f5bdb3aa20f2cdd48771a5b913e80ee88e29d676da7600111672cec39dcc041c84ebc8a6c6eeac2894b78f536e5214ae4274f83b2cf811d19227146623fd2f807193b06d2b552295ff4de8c35b34d61e5ac1fc743277213a181f7ff3ef06d94bdc265fb35182dbfa595ea017b72b4baccba9f37e95dbc16631a46838c94de78baa9056fd47088c61840cca77521a0fd700f0589d7861312d2afae0a5a5eadc82d6971884eff50974645d43a28b86b10a9e613daa53ce227de06d5fe21ef6c823491349fbebf386577151615360053214faa765ddc4f0a24e586302ef0680aacd5d98944dc79c99e545d492631f6ee49dff75081ec4b6b1eadd92122eb8718104316e6e1ff643d658a769ccbb990994522f8627a17d16f252d79d6744e55d9204c66c915225d0152bfb511850f5f3037109ab69e0d32372edc6c0c8fb02bc6c379c0a4d1f4759d861fb901b25297fcbf242bb882ec7fe830644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44264eaf30066f8ee3a29a85dee8f1b9d915e8b8b33b3e0dac5df994ec110715ff07f6f493b70095522626980c2b61e683a6ce1cc2938a55079d5c7c22cee2a2cf2bcf256e799dff4f67a569d0a7dacdb245152005fe91475908d052f800f3cab92e3df3c0e4a486405281193c349d6d03c3903d333007b622b3e51cd36aa216f12f52256e7eabf00c797a091a8ebd756104a7ab4c58bbe41fd96c460adfd2a4e91e83dd00b6f83b16ce34827df3b26bf5642c7dcb55aa291676e1efd6a89c76872b2582b44738838d95ae917392832b13303514c6f5966d59bc9daf695c065a1f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5201b44ade7e4473915b9da4d287a1bb0ade2af47ac66179e0a5e3722b601a26d14b15c5a71c72195eea144d0a1e19bceb29723da40297cb910543e48c946e847251fc8bc1e15bbe3325c4b1074a709ea33771c4bb50073ef8fba22613005ff71a2646970667358221220bce213b86f24151415aa56c2dc17e9d599893589c2b769b91656183cb21b16d364736f6c634300081c0033",
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
