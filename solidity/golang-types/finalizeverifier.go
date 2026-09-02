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
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[10]\",\"internalType\":\"uint256[10]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"input\",\"type\":\"uint256[10]\",\"internalType\":\"uint256[10]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x60808060405234601557611756908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace1114610b985750806344f6369214610aff57806386836a971461064c57806394e4398a146101f75763b8e72af614610055575f80fd5b346101bf5760403660031901126101bf576004356001600160401b0381116101bf57610085903690600401610bd0565b6024356001600160401b0381116101bf576100a4903690600401610bd0565b90918301610100848203126101bf5780601f850112156101bf57604051936100ce61010086610bfd565b849061010081019283116101bf57905b8282106101e7575050508101610140828203126101bf5780601f830112156101bf576040519161011061014084610bfd565b829061014081019283116101bf57905b8282106101c357505050303b156101bf57604051634a721cc560e11b8152915f600484015b600882106101a95750505061010482015f905b600a8210610193575050505f8161024481305afa80156101885761017a575080f35b61018691505f90610bfd565b005b6040513d5f823e3d90fd5b6020806001928551815201930191019091610158565b6020806001928551815201930191019091610145565b5f80fd5b8135815260209182019101610120565b634e487b7160e01b5f52604160045260245ffd5b81358152602091820191016100de565b346101bf576102403660031901126101bf5736610104116101bf5736610244116101bf5760405160408101905f5160206116015f395f51905f52815260208101915f5160206116415f395f51905f5283525f5160206114815f395f51905f528152606082015f5160206113615f395f51905f5281525f5160206114c15f395f51905f52604061010435935f5160206114615f395f51905f52608087019580875284848460608160075afa911016838860808160065afa165f5160206115015f395f51905f5283525f5160206115815f395f51905f5286526001610124359182895286868660608160075afa9310161616838860808160065afa165f5160206116a15f395f51905f5283525f5160206112215f395f51905f528652610144359081885285858560608160075afa92101616838860808160065afa165f5160206113015f395f51905f5283525f5160206115215f395f51905f528652610164359081885285858560608160075afa92101616838860808160065afa165f5160206116615f395f51905f5283525f5160206115c15f395f51905f528652610184359081885285858560608160075afa92101616838860808160065afa165f5160206113c15f395f51905f5283525f5160206113e15f395f51905f5286526101a4359081885285858560608160075afa92101616838860808160065afa165f5160206111e15f395f51905f5283525f5160206114e15f395f51905f5286526101c4359081885285858560608160075afa92101616838860808160065afa165f5160206116c15f395f51905f5283525f5160206112615f395f51905f5286526101e4359081885285858560608160075afa92101616838860808160065afa165f5160206116e15f395f51905f5283525f5160206115415f395f51905f528652610204359081885285858560608160075afa92101616838860808160065afa16945f5160206112c15f395f51905f528352526102243580955260608160075afa9210161660408260808160065afa1690519151901561063d5760405191610100600484375f5160206115a15f395f51905f526101008401525f5160206115615f395f51905f526101208401525f5160206112e15f395f51905f526101408401525f5160206113a15f395f51905f526101608401525f5160206114a15f395f51905f526101808401525f5160206113415f395f51905f526101a08401525f5160206113215f395f51905f526101c08401525f5160206116815f395f51905f526101e08401525f5160206112415f395f51905f526102008401525f5160206112a15f395f51905f526102208401526102408301526102608201525f5160206114215f395f51905f526102808201525f5160206114015f395f51905f526102a08201525f5160206116215f395f51905f526102c08201525f5160206112815f395f51905f526102e08201526020816103008160085afa9051161561062e57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b346101bf576101c03660031901126101bf57366084116101bf57366101c4116101bf5761030060405161067f8282610bfd565b8136823761068e600435610ea9565b61069f602493929335604435610f14565b919392906106ae606435610ea9565b9390926040519660408801965f5160206116015f395f51905f5289528860208101985f5160206116415f395f51905f528a525f5160206114815f395f51905f5281525f5160206114c15f395f51905f52604060608401925f5160206113615f395f51905f5284525f5160206114615f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f5160206115015f395f51905f5285525f5160206115815f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f5160206116a15f395f51905f5285525f5160206112215f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f5160206113015f395f51905f5285525f5160206115215f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f5160206116615f395f51905f5285525f5160206115c15f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f5160206113c15f395f51905f5285525f5160206113e15f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f5160206111e15f395f51905f5285525f5160206114e15f395f51905f5288526101443590818a5287838760608160075afa92101616818360808160065afa165f5160206116c15f395f51905f5285525f5160206112615f395f51905f5288526101643590818a5287838760608160075afa92101616818360808160065afa165f5160206116e15f395f51905f5285525f5160206115415f395f51905f5288526101843590818a5287838760608160075afa921016169160808160065afa16945f5160206112c15f395f51905f528352526101a43580955260608160075afa9210161660408a60808160065afa1698519751981561063d5760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f5160206115a15f395f51905f526101008401525f5160206115615f395f51905f526101208401525f5160206112e15f395f51905f526101408401525f5160206113a15f395f51905f526101608401525f5160206114a15f395f51905f526101808401525f5160206113415f395f51905f526101a08401525f5160206113215f395f51905f526101c08401525f5160206116815f395f51905f526101e08401525f5160206112415f395f51905f526102008401525f5160206112a15f395f51905f526102208401526102408301526102608201525f5160206114215f395f51905f526102808201525f5160206114015f395f51905f526102a08201525f5160206116215f395f51905f526102c08201525f5160206112815f395f51905f526102e0820152604051928391610ada8484610bfd565b8336843760085afa15908115610af2575b5061062e57005b6001915051141581610aeb565b346101bf576101003660031901126101bf5736610104116101bf57604051610b28608082610bfd565b6080368237610b3b602435600435610c20565b8152610b5160843560a435604435606435610cc1565b60208301526040820152610b6960e43560c435610c20565b6060820152604051905f825b60048210610b8257608084f35b6020806001928551815201930191019091610b75565b346101bf575f3660031901126101bf57807f252e7ff136ac4d02e9caac4aff0d8cc88d9b0d51aa53ee80cf7d857f52dfa8da60209252f35b9181601f840112156101bf578235916001600160401b0383116101bf57602083818601950101116101bf57565b601f909101601f19168101906001600160401b038211908210176101d357604052565b905f5160206113815f395f51905f528210801590610caa575b61062e57811580610ca2575b610c9c57610c695f5160206113815f395f51905f5260038185818180090908611014565b818103610c7857505060011b90565b5f5160206113815f395f51905f52809106810306145f1461062e57600190811b1790565b50505f90565b508015610c45565b505f5160206113815f395f51905f52811015610c39565b919093925f5160206113815f395f51905f528310801590610e92575b8015610e7b575b8015610e64575b61062e578082868517171715610e5957908291610dbc5f5160206113815f395f51905f5280808080888180808f9d5f5160206115e15f395f51905f528f839290839109099d8e0981848181800909085f5160206117015f395f51905f52089a09818c8181800909085f5160206112015f395f51905f520806810306945f5160206113815f395f51905f525f5160206114415f395f51905f5281610d9681808b80098187800908611014565b8408095f5160206113815f395f51905f52610db082611178565b80091415958691611037565b929080821480610e50575b15610dee5750505050905f14610de65760ff60025b169060021b179190565b60ff5f610ddc565b5f5160206113815f395f51905f52809106810306149182610e31575b50501561062e5760019115610e295760ff60025b169060021b17179190565b60ff5f610e1e565b5f5160206113815f395f51905f52919250819006810306145f80610e0a565b50838314610dc7565b50505090505f905f90565b505f5160206113815f395f51905f52811015610ceb565b505f5160206113815f395f51905f52821015610ce4565b505f5160206113815f395f51905f52851015610cdd565b8015610f0d578060011c915f5160206113815f395f51905f5283101561062e57600180610eec5f5160206113815f395f51905f5260038188818180090908611014565b931614610ef557565b905f5160206113815f395f51905f5280910681030690565b505f905f90565b80158061100c575b611000578060021c92825f5160206113815f395f51905f528510801590610fe9575b61062e5784815f5160206113815f395f51905f5280808080808080805f5160206115e15f395f51905f5281610fb39d8d0909998a0981898181800909085f5160206112015f395f51905f520806810306936002808a16149509818a8181800909085f5160206117015f395f51905f5208611037565b80929160018082961614610fc5575050565b5f5160206113815f395f51905f528093945080929550809106810306930681030690565b505f5160206113815f395f51905f52811015610f3e565b50505f905f905f905f90565b508115610f1c565b9061101e82611178565b915f5160206113815f395f51905f528380090361062e57565b915f5160206113815f395f51905f525f5160206114415f395f51905f528161107c9396949661106e82808a8009818a800908611014565b9061116c575b860809611014565b925f5160206113815f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206113815f395f51905f5260a083015260208260c08160055afa9151911561062e575f5160206113815f395f51905f5282600192090361062e575f5160206113815f395f51905f52908209925f5160206113815f395f51905f52808080878009068103068187800908149081159161114d575b5061062e57565b90505f5160206113815f395f51905f528084860960020914155f611146565b81809106810306611074565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206113815f395f51905f5260a083015260208260c08160055afa9151911561062e5756fe1f744aa434e4bbee29073b97341d57d556276dcf161b538725f8bfc2695edcf22fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750de134e2d06b9aaba70d1116eba6ad35959f49743ffe58eb731755dcfd7228f81f636bd19cbda582517f5d208a8b1b5370f698583dd8a91c655cab9fdc0abe9e10a44c4a43d89cd622cf60a5df1a7304fda8eaba2a530f14197847e3f25d4c25272122f42abc783db3621a0da2a2115c3b53afa5f0229ec444a8792ae91d4661212c55094cebc71965a68fcdc4f65aef1b467ecac27e10dee537709632f5aa241486051b71b870de3b2a6c62e4cbf675c8a49b6f1cb68371e2c48a192cf3976d295aad2eb714bb061aea84cdf658a5cff38467699546636ac91bb289eb6d11ed25dc302df30437928beb3a9b9860a2f2c97ecf6d637f5e94976df81a5838774b0c9f1076058b8d6a185d54354e90b198e090bbc355270374c877ba768ca0193725038351824b9205e0e1e635a243a14a3c8c6521b31c749541b87d5e4e37288c2d14c08496395d69eb586dfdb4580458a326fb3532a9817dab15a84752c1ea4630644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4719ebfbc8b042b63ccabc82266d121267e391a426e9464eeef1be80b24a0f0b461ab96d757326c565579bc84dd4955a427bc55787356eaacb3f2dd50330d38d830698d66b11f602d8c47e7e554850355b0ea054875ddd86605ee504a8b40a97d902d9976c4108adc8cdb74cce575231a0a8adbed98a1795fba605bedd699b9a230c43c0ec851c715ff8febf631dcba8bc9b637b4c5d1147c166fdcb3e1575a8b1183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea422be9d67b7cb98ae67d01491f596a373c7b930863e609fd157a7883f19dd8e8907bc5765dd669ddd7abb64c46b7ab62a92f5b7aaeb005d7a3862fb33373a9aab22d9ff6e6a469c6badcc11f6e14ae5b9bfe3deb7a5e843caf6eb13664885e28130644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001061c0ca9677e26b2b05faa28c03113af4832a8d2888aac1217bc1e55fc199781169473a7c45097fd64ff960bd6e7720cfb09e172f0f16799e7373cef5a37094214d17be8a6d2bc540f209f0bf50fb0c9da21ad833b6ce8d8683721062d131a19187ec35a35065057af1365ec6e431a4f029f16c83924c2e3a9aa168873bfdf4e0dc2652b9ed39aed811fdddea4df88c239d9d2e1ee7a5ab4bb81f08639ac9a001d015d6da613069c7bc9ad3dbdb71ad284d4230c0f5c6f375b67111b0fea4322160efbad4da82247d0f0b41a1b03e8fd08da5f717d1928d25aa6739ac0cce6ca06696d394b72fe49ae3597b33a106fd4289984c96a4537f92d1aa4858764bcb930644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd440a2da82eb9f3886e388dfa98c351f27d3226ebd33fcb6666f7e7d4f7db700b242bddadea86e168034f4b03d2f5c2bea40131056e8b8e853acec4c4639d8b2be81019e81524b62f7e0025d5079dae66077bc9e7e4a0fad8dcb8c15604dfcdc4481aa0a87d906560d004b118930ccfa21bb6d27e2ecb93d026f286512c4582e427230a5a812e42758fac998ff244e3490f82dc01790d21be1380eb7769f081add31b7bbb8229591dc759fc55f3f0aeb75942da614a4c2e694106ddd29afb1543a0072aade6ef768de8fe7ac2ccb830144166af46f28834ae01998c6745358c97b7038f4fc69a58e624aa1ab8c88bf0fa64bc98ffd444805f25ea17bcaba48436a12b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5a2646970667358221220cb02ce628dfe5c54bd935c74d51c703aa1999984236fe8df2cee9961ceb53ab564736f6c634300081c0033",
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

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x86836a97.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[10] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [10]*big.Int) error {
	var out []interface{}
	err := _FinalizeVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x86836a97.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[10] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [10]*big.Int) error {
	return _FinalizeVerifier.Contract.VerifyCompressedProof(&_FinalizeVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x86836a97.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[10] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [10]*big.Int) error {
	return _FinalizeVerifier.Contract.VerifyCompressedProof(&_FinalizeVerifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x94e4398a.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[10] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, input [10]*big.Int) error {
	var out []interface{}
	err := _FinalizeVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x94e4398a.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[10] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierSession) VerifyProof(proof [8]*big.Int, input [10]*big.Int) error {
	return _FinalizeVerifier.Contract.VerifyProof(&_FinalizeVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x94e4398a.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[10] input) view returns()
func (_FinalizeVerifier *FinalizeVerifierCallerSession) VerifyProof(proof [8]*big.Int, input [10]*big.Int) error {
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
