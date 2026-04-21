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

// FinalizeVerifierMetaData contains all meta data concerning the FinalizeVerifier contract.
var FinalizeVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[9]\",\"internalType\":\"uint256[9]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"input\",\"type\":\"uint256[9]\",\"internalType\":\"uint256[9]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x60808060405234601557611782908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace111461104c5750806344f6369214610fb35780635f89feef146108a75780638a3ae438146101f95763b8e72af614610055575f80fd5b346101c15760403660031901126101c15760043567ffffffffffffffff81116101c157610086903690600401611084565b60243567ffffffffffffffff81116101c1576100a6903690600401611084565b90918301610100848203126101c15780601f850112156101c157604051936100d0610100866110b2565b849061010081019283116101c157905b8282106101e9575050508101610120828203126101c15780601f830112156101c15760405191610112610120846110b2565b829061012081019283116101c157905b8282106101c557505050303b156101c1576040516311475c8760e31b8152915f600484015b600882106101ab5750505061010482015f905b60098210610195575050505f8161022481305afa801561018a5761017c575080f35b61018891505f906110b2565b005b6040513d5f823e3d90fd5b602080600192855181520193019101909161015a565b6020806001928551815201930191019091610147565b5f80fd5b8135815260209182019101610122565b634e487b7160e01b5f52604160045260245ffd5b81358152602091820191016100e0565b346101c1576102203660031901126101c15736610104116101c15736610224116101c15760405160408101907f0e7e181dd4bd26da44827f473c494e1872642ae63b91048e6c936401c698576b815260208101917f066b4fe133cba1a4aa490725516599f176a1187097856bfb416ef3b2177e1af583527f0204c41342681f9b72c9bd368e31c882fb42929ff511196563716eaf9b7681c78152606082017f066942433850eb7b155c095d541024bd2cea6e5339d6ce9fe16d0b02c5a2500f81527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604061010435937f0cbbd4a891af012549ef8d1354fefc995efc9c3714c49864f91120bb4c5416ef608087019580875284848460608160075afa911016838860808160065afa167f2a98e8540ce0fcc4fb2bfd089003b25d5b564b8e9873a6391995df849f531aa883527f0d3a8c6ab9aaf05a403d884ba434a7617e76ab726bb4d21e0448ee8dcc126f5386526001610124359182895286868660608160075afa9310161616838860808160065afa167f087c74c81eea3adc281c5700a7526e7e9970cab20688d3520454844f9a27ed2f83527f1457795b0a1621a96145208987f40ed32824f6b689bf749dda3e8dff5d0015448652610144359081885285858560608160075afa92101616838860808160065afa167f1c939aefe7f90b9a31ac4dfd5cdbfb7dd7bf3f7625062215384fa6340eac26e683527f290968ff3cacd7569a16ae16cbb711976c665a90f19099d22ae9f455dfb7d8088652610164359081885285858560608160075afa92101616838860808160065afa167f1f9abdb14f5da8d7adc52f66f78009047795c910216fd1a09afd02c4250f0fbb83527f0c7cab7eab79a9335d96d653c88d4d2254074147de25c0f46ce54e7ae60e2b618652610184359081885285858560608160075afa92101616838860808160065afa167f08bee54b2de51a6dd6c28a7905d0b1cdd6c712268b01a3dcd4b015949399221f83527f0134bef6d92b7d2e4a04f5a3673f39c0ff569a61de6174bf60f4bc7390f54b9886526101a4359081885285858560608160075afa92101616838860808160065afa167f07260487dd382cc8e81a3ed3308489216d6b433ed17e1a84b42f826357beb74b83527f11908b5e8d61447269bc49fba764cd8a667ea8a57f11ad6f8f0c003bdef6d40b86526101c4359081885285858560608160075afa92101616838860808160065afa167f1001fb11d740da9de01e058ae6b8ae84d9d2d6f4bc7209d27a223342753d32f383527f0b91ea8efa3373d86820cdf69d4cc5a7baeb3da03146a8fad55e7b4f49b9834d86526101e4359081885285858560608160075afa92101616838860808160065afa16947f0ab607b52809566d4f7f9f3ca34028035d87af130dffe98e82669a196f4292dd8352526102043580955260608160075afa9210161660408260808160065afa169051915190156108985760405191610100600484377f03889e9a1198f87013d39c4ba71b7be1d1d7ed1c504cb38530915245ef6f95386101008401527f1c847f19455205060178b6a7ff8797fe5882bf75c46a08363e059079802ba4cd6101208401527f078e56f7aa72dc3f99e6e2a7f3941c9df94b897311d057f2e47b3702bdaee5b96101408401527f1b8f3d16ac5234cdab49b7275e789d7c21173a1efaee9466eef73f9a98319af56101608401527f23bf348f95fbc1b80c083a2d52b71cdc6a84b1bd44460402caf2f5a00ef9f5666101808401527f11118a354000aa03b776dd75f23b7d479b23d253c4b7030d66ec124742eaf9fc6101a08401527f0c8b089bf52b4a84804e6121014ff2a05a7efdaed0d4d58da54d7e0151c2ec436101c08401527f0b6e873c847b95d3a30a81859ca07bed2f2b45f4034cd183cf458bd9a84844786101e08401527f22a138ecd8d8cc536e5ae5f75f0f743d8d39d178399b4be0f1400ab1fed499c26102008401527f2e0cce5213b01d231195ed57117e232c9080cadf9828a48416ec1a51608c54926102208401526102408301526102608201527f06de82bfc4f4b1cb6477d796951e9ca1dd9504fb96e90ed6d6c17bdd88fe67196102808201527f275d1980dcfc0ce573b4a78d5f73cce72de48b8ae5b9ee726170b4b7d206887b6102a08201527f0fd744f1e8c1fa3262ad6060ca4999b2cfea43a7201604054f7fef4f01d9cdd06102c08201527f11e82ba5505177c753bcbd173fb596fc10c52950b2e1ad5d8ac32f3720de6ebb6102e08201526020816103008160085afa9051161561088957005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b346101c1576101a03660031901126101c157366084116101c157366101a4116101c1576103006040516108da82826110b2565b813682376108e96004356113a9565b6108fa602493929335604435611414565b919392906109096064356113a9565b9390926040519660408801967f0e7e181dd4bd26da44827f473c494e1872642ae63b91048e6c936401c698576b89528860208101987f066b4fe133cba1a4aa490725516599f176a1187097856bfb416ef3b2177e1af58a527f0204c41342681f9b72c9bd368e31c882fb42929ff511196563716eaf9b7681c781527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604060608401927f066942433850eb7b155c095d541024bd2cea6e5339d6ce9fe16d0b02c5a2500f84527f0cbbd4a891af012549ef8d1354fefc995efc9c3714c49864f91120bb4c5416ef6084359583608082019780895286828660608160075afa911016818360808160065afa167f2a98e8540ce0fcc4fb2bfd089003b25d5b564b8e9873a6391995df849f531aa885527f0d3a8c6ab9aaf05a403d884ba434a7617e76ab726bb4d21e0448ee8dcc126f538852600160a43591828b5288848860608160075afa9310161616818360808160065afa167f087c74c81eea3adc281c5700a7526e7e9970cab20688d3520454844f9a27ed2f85527f1457795b0a1621a96145208987f40ed32824f6b689bf749dda3e8dff5d001544885260c43590818a5287838760608160075afa92101616818360808160065afa167f1c939aefe7f90b9a31ac4dfd5cdbfb7dd7bf3f7625062215384fa6340eac26e685527f290968ff3cacd7569a16ae16cbb711976c665a90f19099d22ae9f455dfb7d808885260e43590818a5287838760608160075afa92101616818360808160065afa167f1f9abdb14f5da8d7adc52f66f78009047795c910216fd1a09afd02c4250f0fbb85527f0c7cab7eab79a9335d96d653c88d4d2254074147de25c0f46ce54e7ae60e2b6188526101043590818a5287838760608160075afa92101616818360808160065afa167f08bee54b2de51a6dd6c28a7905d0b1cdd6c712268b01a3dcd4b015949399221f85527f0134bef6d92b7d2e4a04f5a3673f39c0ff569a61de6174bf60f4bc7390f54b9888526101243590818a5287838760608160075afa92101616818360808160065afa167f07260487dd382cc8e81a3ed3308489216d6b433ed17e1a84b42f826357beb74b85527f11908b5e8d61447269bc49fba764cd8a667ea8a57f11ad6f8f0c003bdef6d40b88526101443590818a5287838760608160075afa92101616818360808160065afa167f1001fb11d740da9de01e058ae6b8ae84d9d2d6f4bc7209d27a223342753d32f385527f0b91ea8efa3373d86820cdf69d4cc5a7baeb3da03146a8fad55e7b4f49b9834d88526101643590818a5287838760608160075afa921016169160808160065afa16947f0ab607b52809566d4f7f9f3ca34028035d87af130dffe98e82669a196f4292dd8352526101843580955260608160075afa9210161660408a60808160065afa169851975198156108985760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401527f03889e9a1198f87013d39c4ba71b7be1d1d7ed1c504cb38530915245ef6f95386101008401527f1c847f19455205060178b6a7ff8797fe5882bf75c46a08363e059079802ba4cd6101208401527f078e56f7aa72dc3f99e6e2a7f3941c9df94b897311d057f2e47b3702bdaee5b96101408401527f1b8f3d16ac5234cdab49b7275e789d7c21173a1efaee9466eef73f9a98319af56101608401527f23bf348f95fbc1b80c083a2d52b71cdc6a84b1bd44460402caf2f5a00ef9f5666101808401527f11118a354000aa03b776dd75f23b7d479b23d253c4b7030d66ec124742eaf9fc6101a08401527f0c8b089bf52b4a84804e6121014ff2a05a7efdaed0d4d58da54d7e0151c2ec436101c08401527f0b6e873c847b95d3a30a81859ca07bed2f2b45f4034cd183cf458bd9a84844786101e08401527f22a138ecd8d8cc536e5ae5f75f0f743d8d39d178399b4be0f1400ab1fed499c26102008401527f2e0cce5213b01d231195ed57117e232c9080cadf9828a48416ec1a51608c54926102208401526102408301526102608201527f06de82bfc4f4b1cb6477d796951e9ca1dd9504fb96e90ed6d6c17bdd88fe67196102808201527f275d1980dcfc0ce573b4a78d5f73cce72de48b8ae5b9ee726170b4b7d206887b6102a08201527f0fd744f1e8c1fa3262ad6060ca4999b2cfea43a7201604054f7fef4f01d9cdd06102c08201527f11e82ba5505177c753bcbd173fb596fc10c52950b2e1ad5d8ac32f3720de6ebb6102e0820152604051928391610f8e84846110b2565b8336843760085afa15908115610fa6575b5061088957005b6001915051141581610f9f565b346101c1576101003660031901126101c15736610104116101c157604051610fdc6080826110b2565b6080368237610fef6024356004356110d4565b815261100560843560a435604435606435611175565b6020830152604082015261101d60e43560c4356110d4565b6060820152604051905f825b6004821061103657608084f35b6020806001928551815201930191019091611029565b346101c1575f3660031901126101c157807fa5a0381635973c2599e2e0bd559739beb8912aa8607280773b4f2f24f9e7f03560209252f35b9181601f840112156101c15782359167ffffffffffffffff83116101c157602083818601950101116101c157565b90601f8019910116810190811067ffffffffffffffff8211176101d557604052565b905f51602061172d5f395f51905f52821080159061115e575b61088957811580611156575b6111505761111d5f51602061172d5f395f51905f526003818581818009090861154d565b81810361112c57505060011b90565b5f51602061172d5f395f51905f52809106810306145f1461088957600190811b1790565b50505f90565b5080156110f9565b505f51602061172d5f395f51905f528110156110ed565b919093925f51602061172d5f395f51905f528310801590611392575b801561137b575b8015611364575b610889578082868517171715611359579082916112bc5f51602061172d5f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f51602061172d5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161129681808b8009818780090861154d565b8408095f51602061172d5f395f51905f526112b0826116c4565b80091415958691611570565b929080821480611350575b156112ee5750505050905f146112e65760ff60025b169060021b179190565b60ff5f6112dc565b5f51602061172d5f395f51905f52809106810306149182611331575b50501561088957600191156113295760ff60025b169060021b17179190565b60ff5f61131e565b5f51602061172d5f395f51905f52919250819006810306145f8061130a565b508383146112c7565b50505090505f905f90565b505f51602061172d5f395f51905f5281101561119f565b505f51602061172d5f395f51905f52821015611198565b505f51602061172d5f395f51905f52851015611191565b801561140d578060011c915f51602061172d5f395f51905f52831015610889576001806113ec5f51602061172d5f395f51905f526003818881818009090861154d565b9316146113f557565b905f51602061172d5f395f51905f5280910681030690565b505f905f90565b801580611545575b611539578060021c92825f51602061172d5f395f51905f528510801590611522575b6108895784815f51602061172d5f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816114ec9d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611570565b809291600180829616146114fe575050565b5f51602061172d5f395f51905f528093945080929550809106810306930681030690565b505f51602061172d5f395f51905f5281101561143e565b50505f905f905f905f90565b50811561141c565b90611557826116c4565b915f51602061172d5f395f51905f528380090361088957565b915f51602061172d5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816115c8939694966115ba82808a8009818a80090861154d565b906116b8575b86080961154d565b925f51602061172d5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f51602061172d5f395f51905f5260a083015260208260c08160055afa91519115610889575f51602061172d5f395f51905f52826001920903610889575f51602061172d5f395f51905f52908209925f51602061172d5f395f51905f528080808780090681030681878009081490811591611699575b5061088957565b90505f51602061172d5f395f51905f528084860960020914155f611692565b818091068103066115c0565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f51602061172d5f395f51905f5260a083015260208260c08160055afa915191156108895756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a264697066735822122035a3439495a73896507ef6f6fe22ea03c8efa61e2c32a2b782e75699ca90b72764736f6c634300081c0033",
}

// FinalizeVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use FinalizeVerifierMetaData.ABI instead.
var FinalizeVerifierABI = FinalizeVerifierMetaData.ABI

// FinalizeVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FinalizeVerifierMetaData.Bin instead.
var FinalizeVerifierBin = FinalizeVerifierMetaData.Bin

// DeployFinalizeVerifier deploys a new Ethereum contract, binding an instance of FinalizeVerifier to it.
func DeployFinalizeVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FinalizeVerifier, error) {
	parsed, err := FinalizeVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FinalizeVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FinalizeVerifier{FinalizeVerifierCaller: FinalizeVerifierCaller{contract: contract}, FinalizeVerifierTransactor: FinalizeVerifierTransactor{contract: contract}, FinalizeVerifierFilterer: FinalizeVerifierFilterer{contract: contract}}, nil
}

// FinalizeVerifier is an auto generated Go binding around an Ethereum contract.
type FinalizeVerifier struct {
	FinalizeVerifierCaller     // Read-only binding to the contract
	FinalizeVerifierTransactor // Write-only binding to the contract
	FinalizeVerifierFilterer   // Log filterer for contract events
}

// FinalizeVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type FinalizeVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizeVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FinalizeVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizeVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FinalizeVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizeVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FinalizeVerifierSession struct {
	Contract     *FinalizeVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FinalizeVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FinalizeVerifierCallerSession struct {
	Contract *FinalizeVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// FinalizeVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FinalizeVerifierTransactorSession struct {
	Contract     *FinalizeVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// FinalizeVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type FinalizeVerifierRaw struct {
	Contract *FinalizeVerifier // Generic contract binding to access the raw methods on
}

// FinalizeVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FinalizeVerifierCallerRaw struct {
	Contract *FinalizeVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// FinalizeVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FinalizeVerifierTransactorRaw struct {
	Contract *FinalizeVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFinalizeVerifier creates a new instance of FinalizeVerifier, bound to a specific deployed contract.
func NewFinalizeVerifier(address common.Address, backend bind.ContractBackend) (*FinalizeVerifier, error) {
	contract, err := bindFinalizeVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FinalizeVerifier{FinalizeVerifierCaller: FinalizeVerifierCaller{contract: contract}, FinalizeVerifierTransactor: FinalizeVerifierTransactor{contract: contract}, FinalizeVerifierFilterer: FinalizeVerifierFilterer{contract: contract}}, nil
}

// NewFinalizeVerifierCaller creates a new read-only instance of FinalizeVerifier, bound to a specific deployed contract.
func NewFinalizeVerifierCaller(address common.Address, caller bind.ContractCaller) (*FinalizeVerifierCaller, error) {
	contract, err := bindFinalizeVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FinalizeVerifierCaller{contract: contract}, nil
}

// NewFinalizeVerifierTransactor creates a new write-only instance of FinalizeVerifier, bound to a specific deployed contract.
func NewFinalizeVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*FinalizeVerifierTransactor, error) {
	contract, err := bindFinalizeVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FinalizeVerifierTransactor{contract: contract}, nil
}

// NewFinalizeVerifierFilterer creates a new log filterer instance of FinalizeVerifier, bound to a specific deployed contract.
func NewFinalizeVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*FinalizeVerifierFilterer, error) {
	contract, err := bindFinalizeVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FinalizeVerifierFilterer{contract: contract}, nil
}

// bindFinalizeVerifier binds a generic wrapper to an already deployed contract.
func bindFinalizeVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FinalizeVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalizeVerifier *FinalizeVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FinalizeVerifier.Contract.FinalizeVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalizeVerifier *FinalizeVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalizeVerifier.Contract.FinalizeVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalizeVerifier *FinalizeVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalizeVerifier.Contract.FinalizeVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalizeVerifier *FinalizeVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FinalizeVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalizeVerifier *FinalizeVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalizeVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalizeVerifier *FinalizeVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalizeVerifier.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_FinalizeVerifier *FinalizeVerifierCaller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int) ([4]*big.Int, error) {
	var out []interface{}
	err := _FinalizeVerifier.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_FinalizeVerifier *FinalizeVerifierSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _FinalizeVerifier.Contract.CompressProof(&_FinalizeVerifier.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_FinalizeVerifier *FinalizeVerifierCallerSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _FinalizeVerifier.Contract.CompressProof(&_FinalizeVerifier.CallOpts, proof)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_FinalizeVerifier *FinalizeVerifierCaller) ProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FinalizeVerifier.contract.Call(opts, &out, "provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_FinalizeVerifier *FinalizeVerifierSession) ProvingKeyHash() ([32]byte, error) {
	return _FinalizeVerifier.Contract.ProvingKeyHash(&_FinalizeVerifier.CallOpts)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_FinalizeVerifier *FinalizeVerifierCallerSession) ProvingKeyHash() ([32]byte, error) {
	return _FinalizeVerifier.Contract.ProvingKeyHash(&_FinalizeVerifier.CallOpts)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5f89feef.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[9] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [9]*big.Int) error {
	var out []interface{}
	err := _FinalizeVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5f89feef.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[9] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [9]*big.Int) error {
	return _FinalizeVerifier.Contract.VerifyCompressedProof(&_FinalizeVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5f89feef.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[9] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [9]*big.Int) error {
	return _FinalizeVerifier.Contract.VerifyCompressedProof(&_FinalizeVerifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x8a3ae438.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[9] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, input [9]*big.Int) error {
	var out []interface{}
	err := _FinalizeVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x8a3ae438.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[9] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierSession) VerifyProof(proof [8]*big.Int, input [9]*big.Int) error {
	return _FinalizeVerifier.Contract.VerifyProof(&_FinalizeVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x8a3ae438.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[9] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCallerSession) VerifyProof(proof [8]*big.Int, input [9]*big.Int) error {
	return _FinalizeVerifier.Contract.VerifyProof(&_FinalizeVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCaller) VerifyProof0(opts *bind.CallOpts, proof []byte, input []byte) error {
	var out []interface{}
	err := _FinalizeVerifier.contract.Call(opts, &out, "verifyProof0", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_FinalizeVerifier *FinalizeVerifierSession) VerifyProof0(proof []byte, input []byte) error {
	return _FinalizeVerifier.Contract.VerifyProof0(&_FinalizeVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCallerSession) VerifyProof0(proof []byte, input []byte) error {
	return _FinalizeVerifier.Contract.VerifyProof0(&_FinalizeVerifier.CallOpts, proof, input)
}
