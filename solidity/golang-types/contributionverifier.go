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

// ContributionVerifierMetaData contains all meta data concerning the ContributionVerifier contract.
var ContributionVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"input\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidInputEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofEncoding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x608080604052346015576117b0908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c908163233ace1114610e1e5750806344f6369214610da057806348306671146106fc578063a6047e6c146100b65763b8e72af614610053575f80fd5b346100b25760403660031901126100b25760043567ffffffffffffffff81116100b257610084903690600401610e7d565b6024359167ffffffffffffffff83116100b2576100a86100b0933690600401610e7d565b929091610f5c565b005b5f80fd5b346100b2576102003660031901126100b25736610104116100b25736610204116100b25760405160408101907f1deb9598c797efafa344dac01b2ef8227bcd006b73e66b05af2b60ad6eb36f83815260208101917f089c2f6c8f99e3d79aabb885de3b27e68a9474857f0b5316c05bb533cb7fd0f183527f0fe6649deeeba6dc2c20a887a36e6164705a835590703355adf7c3b450410f418152606082017f11adc6797dada83c66d5cf4eed51e76afeb479e78ed0109746a0780c98c207db81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f2cee34fd28c2fc8bad55c72dd2b7f23c1aa058adcf6ac95ff0b0f636b2fc4e1e608087019580875284848460608160075afa911016838860808160065afa167f2a6c2cef26260b8de1af38acd80cc8af221b27872fae2545c99776da6475b48783527f0985182b2fef350ffce9d00a43a23490688e93be3918d4c2bf0ae89d5c12735d86526001610124359182895286868660608160075afa9310161616838860808160065afa167f222a7066dca9422605590f195e02fcb787f541d0d2ae47840cc660a018569f9983527f04c6be73271f6d2181d208c567dd125e1af655adaa0ab8b1f3d6cef6467ead2f8652610144359081885285858560608160075afa92101616838860808160065afa167f133c91503fd9604a9e9196b13ff4b1e993c56b07afa75543f4f536c367fc54cf83527f0e2dabab0753b3c19269570909e081b13e1a1a387144e37e15a36e76b1063f208652610164359081885285858560608160075afa92101616838860808160065afa167f085a1778db479b71dc18ba08272631c70177ee996f5be8c18f823cb09e516b0783527e5427028e6ac0f8af2ac209eab6368cbaf08fc80203b75187e440940210bd318652610184359081885285858560608160075afa92101616838860808160065afa167f1415224de12a88b02dfb614660975d24305482e6b76731284dc57564d849424883527f0e71a1332ab59c54621351055c8f824b1a31df7898d80dfbd038d9602432162386526101a4359081885285858560608160075afa92101616838860808160065afa167f2b8d74f2d5a94d5938121968b9e75030c28e1c82039ffe76d4078c61194f284683527e81fdddb45160f6268a677d40278ccf1a4a1cd7052da87b4f2cbd1a5ab1be8386526101c4359081885285858560608160075afa92101616838860808160065afa16947f1d2d8c39e19cdc4f27a85ea15d89a72ff4dca48254780a7a607c18f9776e8b0a8352526101e43580955260608160075afa9210161660408260808160065afa169051915190156106ed5760405191610100600484377f22a59c3832538e978f28a8900a7cf07f436a679a70161223b3dc052222068b2b6101008401527f060cde84e3b45f2d53f881bf56c499aea5e0b6e4927adaf239c1d1ac1a74fc7b6101208401527f0352867b924a81d2cf6267064969324f03daf574f0611443a63e18f77491eae46101408401527f2747f0e3c40a120a5a7189a2c6dfb520f2464b70197aa39db5359950f49741be6101608401527f1647d993f8d7ef6d3927a22a60a223d8450490628f8d42d511175563a9c98fcc6101808401527f2f8040d09af248bebc5c095b17c7f30cbdebe5933fd82a86d42f0ccec31526ff6101a08401527f2ac207b335c17f9b0e7d8dad1707b6bf9f210e501a9b4432a12166702a6bfd886101c08401527f24dc18aca2eebe1210cbb7ded3fbf65e6db9f832925dd3409825a8c6483f6a386101e08401527f1d936514f337eb76d2defed1a14ce6b71b0d94a6df8a59b2cb842cbc8800343d6102008401527f2c8b4bbfab90c8dfc25bea09d87d80e1195b677de22d597657406de855b0d6a86102208401526102408301526102608201527f0fb87ac94995bb395cfd24002a570a8f4adec1ad91d02bb9d27fcda9b9a2a92e6102808201527f2cddeb137593d4e531729e4852e31a95e830e23f5e518818f8f6ead37888dbe16102a08201527f04e0f203e9a4d0d4b1c7107c7fe6d7efb1f33aba71262c18b53bb7fd7fe9bbd46102c08201527f046e38c0e47d9c8f1790abd52057f093d193e12b725a2c0b7ea920363290bf836102e08201526020816103008160085afa905116156106de57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b346100b2576101803660031901126100b257366084116100b25736610184116100b25761030060405161072f8282610eab565b8136823761073e60043561137a565b61074f6024939293356044356113e5565b9193929061075e60643561137a565b9390926040519660408801967f1deb9598c797efafa344dac01b2ef8227bcd006b73e66b05af2b60ad6eb36f8389528860208101987f089c2f6c8f99e3d79aabb885de3b27e68a9474857f0b5316c05bb533cb7fd0f18a527f0fe6649deeeba6dc2c20a887a36e6164705a835590703355adf7c3b450410f4181527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f11adc6797dada83c66d5cf4eed51e76afeb479e78ed0109746a0780c98c207db84527f2cee34fd28c2fc8bad55c72dd2b7f23c1aa058adcf6ac95ff0b0f636b2fc4e1e6084359583608082019780895286828660608160075afa911016818360808160065afa167f2a6c2cef26260b8de1af38acd80cc8af221b27872fae2545c99776da6475b48785527f0985182b2fef350ffce9d00a43a23490688e93be3918d4c2bf0ae89d5c12735d8852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f222a7066dca9422605590f195e02fcb787f541d0d2ae47840cc660a018569f9985527f04c6be73271f6d2181d208c567dd125e1af655adaa0ab8b1f3d6cef6467ead2f885260c43590818a5287838760608160075afa92101616818360808160065afa167f133c91503fd9604a9e9196b13ff4b1e993c56b07afa75543f4f536c367fc54cf85527f0e2dabab0753b3c19269570909e081b13e1a1a387144e37e15a36e76b1063f20885260e43590818a5287838760608160075afa92101616818360808160065afa167f085a1778db479b71dc18ba08272631c70177ee996f5be8c18f823cb09e516b0785527e5427028e6ac0f8af2ac209eab6368cbaf08fc80203b75187e440940210bd3188526101043590818a5287838760608160075afa92101616818360808160065afa167f1415224de12a88b02dfb614660975d24305482e6b76731284dc57564d849424885527f0e71a1332ab59c54621351055c8f824b1a31df7898d80dfbd038d9602432162388526101243590818a5287838760608160075afa92101616818360808160065afa167f2b8d74f2d5a94d5938121968b9e75030c28e1c82039ffe76d4078c61194f284685527e81fdddb45160f6268a677d40278ccf1a4a1cd7052da87b4f2cbd1a5ab1be8388526101443590818a5287838760608160075afa921016169160808160065afa16947f1d2d8c39e19cdc4f27a85ea15d89a72ff4dca48254780a7a607c18f9776e8b0a8352526101643580955260608160075afa9210161660408a60808160065afa169851975198156106ed5760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f22a59c3832538e978f28a8900a7cf07f436a679a70161223b3dc052222068b2b6101008401527f060cde84e3b45f2d53f881bf56c499aea5e0b6e4927adaf239c1d1ac1a74fc7b6101208401527f0352867b924a81d2cf6267064969324f03daf574f0611443a63e18f77491eae46101408401527f2747f0e3c40a120a5a7189a2c6dfb520f2464b70197aa39db5359950f49741be6101608401527f1647d993f8d7ef6d3927a22a60a223d8450490628f8d42d511175563a9c98fcc6101808401527f2f8040d09af248bebc5c095b17c7f30cbdebe5933fd82a86d42f0ccec31526ff6101a08401527f2ac207b335c17f9b0e7d8dad1707b6bf9f210e501a9b4432a12166702a6bfd886101c08401527f24dc18aca2eebe1210cbb7ded3fbf65e6db9f832925dd3409825a8c6483f6a386101e08401527f1d936514f337eb76d2defed1a14ce6b71b0d94a6df8a59b2cb842cbc8800343d6102008401527f2c8b4bbfab90c8dfc25bea09d87d80e1195b677de22d597657406de855b0d6a86102208401526102408301526102608201527f0fb87ac94995bb395cfd24002a570a8f4adec1ad91d02bb9d27fcda9b9a2a92e6102808201527f2cddeb137593d4e531729e4852e31a95e830e23f5e518818f8f6ead37888dbe16102a08201527f04e0f203e9a4d0d4b1c7107c7fe6d7efb1f33aba71262c18b53bb7fd7fe9bbd46102c08201527f046e38c0e47d9c8f1790abd52057f093d193e12b725a2c0b7ea920363290bf836102e0820152604051928391610d7b8484610eab565b8336843760085afa15908115610d93575b506106de57005b6001915051141581610d8c565b346100b2576101003660031901126100b25736610104116100b2576080604051610dca8282610eab565b81368237610ddc6024356004356110a5565b8152610df260843560a435604435606435611146565b60208301526040820152610e0a60e43560c4356110a5565b6060820152610e1c6040518092610e56565bf35b346100b2575f3660031901126100b257807fe662b9f5b48e5c825efe2667a1b8e144cd262f125c06fcc794f6f198dea9f15360209252f35b905f905b60048210610e6757505050565b6020806001928551815201930191019091610e5a565b9181601f840112156100b25782359167ffffffffffffffff83116100b257602083818601950101116100b257565b90601f8019910116810190811067ffffffffffffffff821117610ecd57604052565b634e487b7160e01b5f52604160045260245ffd5b90610100828203126100b25780601f830112156100b25760405191610f0861010084610eab565b829061010081019283116100b257905b828210610f255750505090565b8135815260209182019101610f18565b905f905b60088210610f4657505050565b6020806001928551815201930191019091610f39565b93929190610100811461103e5760808114610f805763236bd13760e01b5f5260045ffd5b610100830361102f578401916080858403126100b25782601f860112156100b25760405192610fb0608085610eab565b8395608081019182116100b257955b81871061101f57505061101d93945081610fde91611007930190610ee1565b604051634830667160e01b602082015292610ffd906024850190610e56565b60a4830190610f35565b61018481526110186101a482610eab565b61151e565b565b8635815260209687019601610fbf565b630c0b7e3560e11b5f5260045ffd5b909293610100830361102f5761101d93611061826110949461106a940190610ee1565b93810190610ee1565b6040516329811f9b60e21b602082015292611089906024850190610f35565b610124830190610f35565b610204815261101861022482610eab565b905f51602061175b5f395f51905f52821080159061112f575b6106de57811580611127575b611121576110ee5f51602061175b5f395f51905f526003818581818009090861157b565b8181036110fd57505060011b90565b5f51602061175b5f395f51905f52809106810306145f146106de57600190811b1790565b50505f90565b5080156110ca565b505f51602061175b5f395f51905f528110156110be565b919093925f51602061175b5f395f51905f528310801590611363575b801561134c575b8015611335575b6106de57808286851717171561132a5790829161128d5f51602061175b5f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f51602061175b5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161126781808b8009818780090861157b565b8408095f51602061175b5f395f51905f52611281826116f2565b8009141595869161159e565b929080821480611321575b156112bf5750505050905f146112b75760ff60025b169060021b179190565b60ff5f6112ad565b5f51602061175b5f395f51905f52809106810306149182611302575b5050156106de57600191156112fa5760ff60025b169060021b17179190565b60ff5f6112ef565b5f51602061175b5f395f51905f52919250819006810306145f806112db565b50838314611298565b50505090505f905f90565b505f51602061175b5f395f51905f52811015611170565b505f51602061175b5f395f51905f52821015611169565b505f51602061175b5f395f51905f52851015611162565b80156113de578060011c915f51602061175b5f395f51905f528310156106de576001806113bd5f51602061175b5f395f51905f526003818881818009090861157b565b9316146113c657565b905f51602061175b5f395f51905f5280910681030690565b505f905f90565b801580611516575b61150a578060021c92825f51602061175b5f395f51905f5285108015906114f3575b6106de5784815f51602061175b5f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816114bd9d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50861159e565b809291600180829616146114cf575050565b5f51602061175b5f395f51905f528093945080929550809106810306930681030690565b505f51602061175b5f395f51905f5281101561140f565b50505f905f905f905f90565b5081156113ed565b5f8091602081519101305afa3d15611573573d9067ffffffffffffffff8211610ecd5760405191611559601f8201601f191660200184610eab565b82523d5f602084013e5b1561156b5750565b602081519101fd5b606090611563565b90611585826116f2565b915f51602061175b5f395f51905f52838009036106de57565b915f51602061175b5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816115f6939694966115e882808a8009818a80090861157b565b906116e6575b86080961157b565b925f51602061175b5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f51602061175b5f395f51905f5260a083015260208260c08160055afa915191156106de575f51602061175b5f395f51905f528260019209036106de575f51602061175b5f395f51905f52908209925f51602061175b5f395f51905f5280808087800906810306818780090814908115916116c7575b506106de57565b90505f51602061175b5f395f51905f528084860960020914155f6116c0565b818091068103066115ee565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f51602061175b5f395f51905f5260a083015260208260c08160055afa915191156106de5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220f3f06f2f563982afcee2607183b4dcaae74a532860363aaac98ea3fd9d989bbb64736f6c634300081c0033",
}

// ContributionVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use ContributionVerifierMetaData.ABI instead.
var ContributionVerifierABI = ContributionVerifierMetaData.ABI

// ContributionVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContributionVerifierMetaData.Bin instead.
var ContributionVerifierBin = ContributionVerifierMetaData.Bin

// DeployContributionVerifier deploys a new Ethereum contract, binding an instance of ContributionVerifier to it.
func DeployContributionVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContributionVerifier, error) {
	parsed, err := ContributionVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContributionVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContributionVerifier{ContributionVerifierCaller: ContributionVerifierCaller{contract: contract}, ContributionVerifierTransactor: ContributionVerifierTransactor{contract: contract}, ContributionVerifierFilterer: ContributionVerifierFilterer{contract: contract}}, nil
}

// ContributionVerifier is an auto generated Go binding around an Ethereum contract.
type ContributionVerifier struct {
	ContributionVerifierCaller     // Read-only binding to the contract
	ContributionVerifierTransactor // Write-only binding to the contract
	ContributionVerifierFilterer   // Log filterer for contract events
}

// ContributionVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContributionVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContributionVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContributionVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContributionVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContributionVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContributionVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContributionVerifierSession struct {
	Contract     *ContributionVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ContributionVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContributionVerifierCallerSession struct {
	Contract *ContributionVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ContributionVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContributionVerifierTransactorSession struct {
	Contract     *ContributionVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ContributionVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContributionVerifierRaw struct {
	Contract *ContributionVerifier // Generic contract binding to access the raw methods on
}

// ContributionVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContributionVerifierCallerRaw struct {
	Contract *ContributionVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// ContributionVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContributionVerifierTransactorRaw struct {
	Contract *ContributionVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContributionVerifier creates a new instance of ContributionVerifier, bound to a specific deployed contract.
func NewContributionVerifier(address common.Address, backend bind.ContractBackend) (*ContributionVerifier, error) {
	contract, err := bindContributionVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContributionVerifier{ContributionVerifierCaller: ContributionVerifierCaller{contract: contract}, ContributionVerifierTransactor: ContributionVerifierTransactor{contract: contract}, ContributionVerifierFilterer: ContributionVerifierFilterer{contract: contract}}, nil
}

// NewContributionVerifierCaller creates a new read-only instance of ContributionVerifier, bound to a specific deployed contract.
func NewContributionVerifierCaller(address common.Address, caller bind.ContractCaller) (*ContributionVerifierCaller, error) {
	contract, err := bindContributionVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContributionVerifierCaller{contract: contract}, nil
}

// NewContributionVerifierTransactor creates a new write-only instance of ContributionVerifier, bound to a specific deployed contract.
func NewContributionVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*ContributionVerifierTransactor, error) {
	contract, err := bindContributionVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContributionVerifierTransactor{contract: contract}, nil
}

// NewContributionVerifierFilterer creates a new log filterer instance of ContributionVerifier, bound to a specific deployed contract.
func NewContributionVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*ContributionVerifierFilterer, error) {
	contract, err := bindContributionVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContributionVerifierFilterer{contract: contract}, nil
}

// bindContributionVerifier binds a generic wrapper to an already deployed contract.
func bindContributionVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContributionVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContributionVerifier *ContributionVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContributionVerifier.Contract.ContributionVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContributionVerifier *ContributionVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContributionVerifier.Contract.ContributionVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContributionVerifier *ContributionVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContributionVerifier.Contract.ContributionVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContributionVerifier *ContributionVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContributionVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContributionVerifier *ContributionVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContributionVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContributionVerifier *ContributionVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContributionVerifier.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_ContributionVerifier *ContributionVerifierCaller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int) ([4]*big.Int, error) {
	var out []interface{}
	err := _ContributionVerifier.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_ContributionVerifier *ContributionVerifierSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _ContributionVerifier.Contract.CompressProof(&_ContributionVerifier.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_ContributionVerifier *ContributionVerifierCallerSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _ContributionVerifier.Contract.CompressProof(&_ContributionVerifier.CallOpts, proof)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_ContributionVerifier *ContributionVerifierCaller) ProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContributionVerifier.contract.Call(opts, &out, "provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_ContributionVerifier *ContributionVerifierSession) ProvingKeyHash() ([32]byte, error) {
	return _ContributionVerifier.Contract.ProvingKeyHash(&_ContributionVerifier.CallOpts)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_ContributionVerifier *ContributionVerifierCallerSession) ProvingKeyHash() ([32]byte, error) {
	return _ContributionVerifier.Contract.ProvingKeyHash(&_ContributionVerifier.CallOpts)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x48306671.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [8]*big.Int) error {
	var out []interface{}
	err := _ContributionVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x48306671.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [8]*big.Int) error {
	return _ContributionVerifier.Contract.VerifyCompressedProof(&_ContributionVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x48306671.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [8]*big.Int) error {
	return _ContributionVerifier.Contract.VerifyCompressedProof(&_ContributionVerifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xa6047e6c.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, input [8]*big.Int) error {
	var out []interface{}
	err := _ContributionVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0xa6047e6c.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierSession) VerifyProof(proof [8]*big.Int, input [8]*big.Int) error {
	return _ContributionVerifier.Contract.VerifyProof(&_ContributionVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xa6047e6c.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[8] input) view returns()
func (_ContributionVerifier *ContributionVerifierCallerSession) VerifyProof(proof [8]*big.Int, input [8]*big.Int) error {
	return _ContributionVerifier.Contract.VerifyProof(&_ContributionVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_ContributionVerifier *ContributionVerifierCaller) VerifyProof0(opts *bind.CallOpts, proof []byte, input []byte) error {
	var out []interface{}
	err := _ContributionVerifier.contract.Call(opts, &out, "verifyProof0", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_ContributionVerifier *ContributionVerifierSession) VerifyProof0(proof []byte, input []byte) error {
	return _ContributionVerifier.Contract.VerifyProof0(&_ContributionVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_ContributionVerifier *ContributionVerifierCallerSession) VerifyProof0(proof []byte, input []byte) error {
	return _ContributionVerifier.Contract.VerifyProof0(&_ContributionVerifier.CallOpts, proof, input)
}
