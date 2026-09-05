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

// PoolKeyVerifierMetaData contains all meta data concerning the PoolKeyVerifier contract.
var PoolKeyVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"compressProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"compressed\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provingKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyCompressedProof\",\"inputs\":[{\"name\":\"compressedProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"ProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicInputNotInField\",\"inputs\":[]}]",
	Bin: "0x60808060405234601557611615908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80630ea14a39146109f0578063233ace11146109b6578063449ccd1e146105e657806348306671146101955763b8e72af614610051575f80fd5b34610181576040366003190112610181576004356001600160401b03811161018157610081903690600401610ab8565b6024356001600160401b038111610181576100a0903690600401610ab8565b8101610100828203126101815780601f8301121561018157610100604051926100c98285610ae5565b8391810192831161018157905b82821061018557505050303b156101815760405163224e668f60e11b8152610120600482015261012481018390529283919083906101448401375f6101448484010152602482015f905b6008821061016757505050610144815f93601f80199101168101030181305afa801561015c5761014e575080f35b61015a91505f90610ae5565b005b6040513d5f823e3d90fd5b829350602080916001939451815201930191018492610120565b5f80fd5b81358152602091820191016100d6565b34610181576101803660031901126101815736608411610181573661018411610181576103006040516101c88282610ae5565b813682376101d7600435610de8565b6101e8602493929335604435610e53565b919392906101f7606435610de8565b9390926040519660408801965f5160206114205f395f51905f5289528860208101985f5160206111c05f395f51905f528a525f5160206114805f395f51905f5281525f5160206113a05f395f51905f52604060608401925f5160206113205f395f51905f5284525f5160206113605f395f51905f526084359583608082019780895286828660608160075afa911016818360808160065afa165f5160206114005f395f51905f5285525f5160206111405f395f51905f528852600160a43591828b5288848860608160075afa9310161616818360808160065afa165f5160206114605f395f51905f5285525f5160206115405f395f51905f52885260c43590818a5287838760608160075afa92101616818360808160065afa165f5160206113005f395f51905f5285525f5160206114a05f395f51905f52885260e43590818a5287838760608160075afa92101616818360808160065afa165f5160206112605f395f51905f5285525f5160206115605f395f51905f5288526101043590818a5287838760608160075afa92101616818360808160065afa165f5160206111e05f395f51905f5285525f5160206113805f395f51905f5288526101243590818a5287838760608160075afa92101616818360808160065afa165f5160206111605f395f51905f5285525f5160206115805f395f51905f5288526101443590818a5287838760608160075afa921016169160808160065afa16945f5160206111805f395f51905f528352526101643580955260608160075afa9210161660408a60808160065afa169851975198156105d75760209a8a528a8a015260408901526060880152608087015260a086015260c085015260e08401525f5160206113e05f395f51905f526101008401525f5160206112405f395f51905f526101208401525f5160206111a05f395f51905f526101408401525f5160206115c05f395f51905f526101608401525f5160206112805f395f51905f526101808401525f5160206114e05f395f51905f526101a08401525f5160206115205f395f51905f526101c08401525f5160206112005f395f51905f526101e08401525f5160206114405f395f51905f526102008401525f5160206115005f395f51905f526102208401526102408301526102608201525f5160206112205f395f51905f526102808201525f5160206112e05f395f51905f526102a08201525f5160206112c05f395f51905f526102c08201525f5160206113c05f395f51905f526102e08201526040519283916105a38484610ae5565b8336843760085afa159081156105ca575b506105bb57005b631ff3747d60e21b5f5260045ffd5b60019150511415816105b4565b63a54f8e2760e01b5f5260045ffd5b3461018157610120366003190112610181576004356001600160401b03811161018157610617903690600401610ab8565b36610124116101815761010061062d9114610b1c565b604051604081015f5160206114205f395f51905f52825260208201905f5160206111c05f395f51905f5282525f5160206114805f395f51905f528152606083015f5160206113205f395f51905f5281525f5160206113a05f395f51905f526040602435935f5160206113605f395f51905f52608088019580875284848460608160075afa911016838960808160065afa165f5160206114005f395f51905f5283525f5160206111405f395f51905f52865260016044359182895286868660608160075afa9310161616838960808160065afa165f5160206114605f395f51905f5283525f5160206115405f395f51905f5286526064359081885285858560608160075afa92101616838960808160065afa165f5160206113005f395f51905f5283525f5160206114a05f395f51905f5286526084359081885285858560608160075afa92101616838960808160065afa165f5160206112605f395f51905f5283525f5160206115605f395f51905f52865260a4359081885285858560608160075afa92101616838960808160065afa165f5160206111e05f395f51905f5283525f5160206113805f395f51905f52865260c4359081885285858560608160075afa92101616838960808160065afa165f5160206111605f395f51905f5283525f5160206115805f395f51905f52865260e4359081885285858560608160075afa92101616838960808160065afa16945f5160206111805f395f51905f528352526101043580955260608160075afa9210161660408360808160065afa169151905191156105d7576101006040519384375f5160206113e05f395f51905f526101008401525f5160206112405f395f51905f526101208401525f5160206111a05f395f51905f526101408401525f5160206115c05f395f51905f526101608401525f5160206112805f395f51905f526101808401525f5160206114e05f395f51905f526101a08401525f5160206115205f395f51905f526101c08401525f5160206112005f395f51905f526101e08401525f5160206114405f395f51905f526102008401525f5160206115005f395f51905f526102208401526102408301526102608201525f5160206112205f395f51905f526102808201525f5160206112e05f395f51905f526102a08201525f5160206112c05f395f51905f526102c08201525f5160206113c05f395f51905f526102e08201526020816103008160085afa905116156105bb57005b34610181575f3660031901126101815760206040517f190f9007ad878e9b8d8ef68770c984d3ada5d6ad32e4945393993317cb20b8d58152f35b34610181576020366003190112610181576004356001600160401b03811161018157610a20903690600401610ab8565b610a89608092610a4461010060405194610a3a8787610ae5565b8636873714610b1c565b610a5360208201358235610b5f565b8352610a708482013560a083013560408401356060850135610c00565b6020850152604084015260c060e0820135910135610b5f565b6060820152604051905f825b60048210610aa257505050f35b6020806001928551815201930191019091610a95565b9181601f84011215610181578235916001600160401b038311610181576020838186019501011161018157565b601f909101601f19168101906001600160401b03821190821017610b0857604052565b634e487b7160e01b5f52604160045260245ffd5b15610b2357565b60405162461bcd60e51b81526020600482015260146024820152730d2dcecc2d8d2c840e0e4dedecc40d8cadccee8d60631b6044820152606490fd5b905f5160206112a05f395f51905f528210801590610be9575b6105bb57811580610be1575b610bdb57610ba85f5160206112a05f395f51905f5260038185818180090908610f53565b818103610bb757505060011b90565b5f5160206112a05f395f51905f52809106810306145f146105bb57600190811b1790565b50505f90565b508015610b84565b505f5160206112a05f395f51905f52811015610b78565b919093925f5160206112a05f395f51905f528310801590610dd1575b8015610dba575b8015610da3575b6105bb578082868517171715610d9857908291610cfb5f5160206112a05f395f51905f5280808080888180808f9d5f5160206114c05f395f51905f528f839290839109099d8e0981848181800909085f5160206115a05f395f51905f52089a09818c8181800909085f5160206111205f395f51905f520806810306945f5160206112a05f395f51905f525f5160206113405f395f51905f5281610cd581808b80098187800908610f53565b8408095f5160206112a05f395f51905f52610cef826110b7565b80091415958691610f76565b929080821480610d8f575b15610d2d5750505050905f14610d255760ff60025b169060021b179190565b60ff5f610d1b565b5f5160206112a05f395f51905f52809106810306149182610d70575b5050156105bb5760019115610d685760ff60025b169060021b17179190565b60ff5f610d5d565b5f5160206112a05f395f51905f52919250819006810306145f80610d49565b50838314610d06565b50505090505f905f90565b505f5160206112a05f395f51905f52811015610c2a565b505f5160206112a05f395f51905f52821015610c23565b505f5160206112a05f395f51905f52851015610c1c565b8015610e4c578060011c915f5160206112a05f395f51905f528310156105bb57600180610e2b5f5160206112a05f395f51905f5260038188818180090908610f53565b931614610e3457565b905f5160206112a05f395f51905f5280910681030690565b505f905f90565b801580610f4b575b610f3f578060021c92825f5160206112a05f395f51905f528510801590610f28575b6105bb5784815f5160206112a05f395f51905f5280808080808080805f5160206114c05f395f51905f5281610ef29d8d0909998a0981898181800909085f5160206111205f395f51905f520806810306936002808a16149509818a8181800909085f5160206115a05f395f51905f5208610f76565b80929160018082961614610f04575050565b5f5160206112a05f395f51905f528093945080929550809106810306930681030690565b505f5160206112a05f395f51905f52811015610e7d565b50505f905f905f905f90565b508115610e5b565b90610f5d826110b7565b915f5160206112a05f395f51905f52838009036105bb57565b915f5160206112a05f395f51905f525f5160206113405f395f51905f5281610fbb93969496610fad82808a8009818a800908610f53565b906110ab575b860809610f53565b925f5160206112a05f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f5160206112a05f395f51905f5260a083015260208260c08160055afa915191156105bb575f5160206112a05f395f51905f528260019209036105bb575f5160206112a05f395f51905f52908209925f5160206112a05f395f51905f52808080878009068103068187800908149081159161108c575b506105bb57565b90505f5160206112a05f395f51905f528084860960020914155f611085565b81809106810306610fb3565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f5160206112a05f395f51905f5260a083015260208260c08160055afa915191156105bb5756fe2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750675054c18d2841bd71468019b842b1b23aa02438a33705c0efa57768f57510e05a3a35b2328817cd074d99e16c1d3677ecebda0eda88393038bd81ebad2801c24159eff8b1ca9f1aa31c2a81dc100eb84d0cd658f795a176e1a2307c3af79b51ec706a02698789f1502bd27ac0b79227818368ff08e4b41ac5f3544b21d33e72d248d67bd26f1e0d138ca1256f0254c71bcef05b4325a54d08afc388ee6c0aa0f439c5782f49e2c2888834d3d90b75b3eb626855ee9610f95aff9db622e443305e6f763994c89caa27ae4bb57efc1ec10e4d29de150da148a3549230e9e3f8d00c32bd2faccf4e2e754aa6b750b8a0190be4174b34e2f77faa2e2ed3173876001550f1d5c0992e63aa2a20d17fa64af18b3bb011db086fc08cb6b7d820728aa22c8328d6da45bc81cb4da134a57b834e70ea7e0fb45d7cae1b17fdc36181e7a2bfa19b544ad773d9161ae41016aaac6529bb4757f8fcd93c58e42969292ff9930644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd472dfa2ea4174b4ee7688f82d3a6f1c2854b5d7f9cd7994ccef4647b38880f0bb2101b34c555f2bafde3b88f60dac67a46d7021c0a9bd55778d421a744c3f97e6724b3307648e1fb030f94c1941d27d8e6feed6385612a5d360df61a8379df5326258eec22bb57deee1f99a825ac42391c3bc7921e25f699c1842ae89ebaaf4f88183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea42d7aaf7ccb79423e43f652c31d550f5d9925beb6316ff86c59fc7a154bb05720006f29dc618ea4371fd7c110e3adf624ac2aef62707b7febf042b3bf06d5d0b530644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000011bafba6fc31d2298bf916e7c771cd4fed2f83d2de4cde8c71986d98d0969732d1608cf61cc3b4ce3167ad77f52c25423b1a993866767407d6ef63a0d53d84e1e22c630611da24925ff051743c3d897fcecba733d09f4423a650039b109691fd705d175c744c5741c4c625115329abd8afd4961c6f0d0ba4e864578a7ac1306b6091356ea3fcfd7a0a483823cb07a95d0d186b3b82aad8efb39c10bbadcf2bc041a2c647c0deec318b0139d07557bfca36264b712b3571f86a92209ceae1d3d2029711ec6ea19e42b315e182d17e0316b0ca811043a29e7f6c0b756cf78304c220aba68f972b53cb223f1c2ff122a1ccc6f36cd2a6e26b38fce9b4a24c20a39c830644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd441f2be75b3008da80391cc1d8c0004f53e1af44d93acb9ff027fecc2c2fef69fa1864b48711a5c2dceb3f7090a40be643e669dd6b64e5abc8299502753cc241ce1afdffb6bda022ca6eeb4e78dc3a9253cab2ca21f5cc6ee1d2ffecfac0414fcf2358d9f7d6221402b2266a3accd2ad49d4aaccbef61ba45b9b776e7ae8b9b3081c703a014c756260e026d768a752e606f9b4d5083107779675cb6860526306e70a9fd137b69184e0b7e014ac27aaccdc47aa08300db40454e2520a894009d4652b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e52579148f0911000fb623ebedff4d8cef48e33148d0545718dd0848106bd936dfa264697066735822122016e58085b16762e8eaee46276d741e3d00f7625d4364cbb6d4941d30a1b05c9c64736f6c634300081c0033",
}

// PoolKeyVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolKeyVerifierMetaData.ABI instead.
var PoolKeyVerifierABI = PoolKeyVerifierMetaData.ABI

// PoolKeyVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoolKeyVerifierMetaData.Bin instead.
var PoolKeyVerifierBin = PoolKeyVerifierMetaData.Bin

// DeployPoolKeyVerifier deploys a new Ethereum contract, binding an instance of PoolKeyVerifier to it.
func DeployPoolKeyVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PoolKeyVerifier, error) {
	parsed, err := PoolKeyVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoolKeyVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PoolKeyVerifier{PoolKeyVerifierCaller: PoolKeyVerifierCaller{contract: contract}, PoolKeyVerifierTransactor: PoolKeyVerifierTransactor{contract: contract}, PoolKeyVerifierFilterer: PoolKeyVerifierFilterer{contract: contract}}, nil
}

// PoolKeyVerifier is an auto generated Go binding around an Ethereum contract.
type PoolKeyVerifier struct {
	PoolKeyVerifierCaller     // Read-only binding to the contract
	PoolKeyVerifierTransactor // Write-only binding to the contract
	PoolKeyVerifierFilterer   // Log filterer for contract events
}

// PoolKeyVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolKeyVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolKeyVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolKeyVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolKeyVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolKeyVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolKeyVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolKeyVerifierSession struct {
	Contract     *PoolKeyVerifier  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolKeyVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolKeyVerifierCallerSession struct {
	Contract *PoolKeyVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PoolKeyVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolKeyVerifierTransactorSession struct {
	Contract     *PoolKeyVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PoolKeyVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolKeyVerifierRaw struct {
	Contract *PoolKeyVerifier // Generic contract binding to access the raw methods on
}

// PoolKeyVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolKeyVerifierCallerRaw struct {
	Contract *PoolKeyVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// PoolKeyVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolKeyVerifierTransactorRaw struct {
	Contract *PoolKeyVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolKeyVerifier creates a new instance of PoolKeyVerifier, bound to a specific deployed contract.
func NewPoolKeyVerifier(address common.Address, backend bind.ContractBackend) (*PoolKeyVerifier, error) {
	contract, err := bindPoolKeyVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolKeyVerifier{PoolKeyVerifierCaller: PoolKeyVerifierCaller{contract: contract}, PoolKeyVerifierTransactor: PoolKeyVerifierTransactor{contract: contract}, PoolKeyVerifierFilterer: PoolKeyVerifierFilterer{contract: contract}}, nil
}

// NewPoolKeyVerifierCaller creates a new read-only instance of PoolKeyVerifier, bound to a specific deployed contract.
func NewPoolKeyVerifierCaller(address common.Address, caller bind.ContractCaller) (*PoolKeyVerifierCaller, error) {
	contract, err := bindPoolKeyVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolKeyVerifierCaller{contract: contract}, nil
}

// NewPoolKeyVerifierTransactor creates a new write-only instance of PoolKeyVerifier, bound to a specific deployed contract.
func NewPoolKeyVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolKeyVerifierTransactor, error) {
	contract, err := bindPoolKeyVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolKeyVerifierTransactor{contract: contract}, nil
}

// NewPoolKeyVerifierFilterer creates a new log filterer instance of PoolKeyVerifier, bound to a specific deployed contract.
func NewPoolKeyVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolKeyVerifierFilterer, error) {
	contract, err := bindPoolKeyVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolKeyVerifierFilterer{contract: contract}, nil
}

// bindPoolKeyVerifier binds a generic wrapper to an already deployed contract.
func bindPoolKeyVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolKeyVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolKeyVerifier *PoolKeyVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolKeyVerifier.Contract.PoolKeyVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolKeyVerifier *PoolKeyVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolKeyVerifier.Contract.PoolKeyVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolKeyVerifier *PoolKeyVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolKeyVerifier.Contract.PoolKeyVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolKeyVerifier *PoolKeyVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolKeyVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolKeyVerifier *PoolKeyVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolKeyVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolKeyVerifier *PoolKeyVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolKeyVerifier.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed)
func (_PoolKeyVerifier *PoolKeyVerifierCaller) CompressProof(opts *bind.CallOpts, proof []byte) ([4]*big.Int, error) {
	var out []interface{}
	err := _PoolKeyVerifier.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed)
func (_PoolKeyVerifier *PoolKeyVerifierSession) CompressProof(proof []byte) ([4]*big.Int, error) {
	return _PoolKeyVerifier.Contract.CompressProof(&_PoolKeyVerifier.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed)
func (_PoolKeyVerifier *PoolKeyVerifierCallerSession) CompressProof(proof []byte) ([4]*big.Int, error) {
	return _PoolKeyVerifier.Contract.CompressProof(&_PoolKeyVerifier.CallOpts, proof)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_PoolKeyVerifier *PoolKeyVerifierCaller) ProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolKeyVerifier.contract.Call(opts, &out, "provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_PoolKeyVerifier *PoolKeyVerifierSession) ProvingKeyHash() ([32]byte, error) {
	return _PoolKeyVerifier.Contract.ProvingKeyHash(&_PoolKeyVerifier.CallOpts)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_PoolKeyVerifier *PoolKeyVerifierCallerSession) ProvingKeyHash() ([32]byte, error) {
	return _PoolKeyVerifier.Contract.ProvingKeyHash(&_PoolKeyVerifier.CallOpts)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x48306671.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[8] input) view returns()
func (_PoolKeyVerifier *PoolKeyVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [8]*big.Int) error {
	var out []interface{}
	err := _PoolKeyVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x48306671.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[8] input) view returns()
func (_PoolKeyVerifier *PoolKeyVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [8]*big.Int) error {
	return _PoolKeyVerifier.Contract.VerifyCompressedProof(&_PoolKeyVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x48306671.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[8] input) view returns()
func (_PoolKeyVerifier *PoolKeyVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [8]*big.Int) error {
	return _PoolKeyVerifier.Contract.VerifyCompressedProof(&_PoolKeyVerifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x449ccd1e.
//
// Solidity: function verifyProof(bytes proof, uint256[8] input) view returns()
func (_PoolKeyVerifier *PoolKeyVerifierCaller) VerifyProof(opts *bind.CallOpts, proof []byte, input [8]*big.Int) error {
	var out []interface{}
	err := _PoolKeyVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x449ccd1e.
//
// Solidity: function verifyProof(bytes proof, uint256[8] input) view returns()
func (_PoolKeyVerifier *PoolKeyVerifierSession) VerifyProof(proof []byte, input [8]*big.Int) error {
	return _PoolKeyVerifier.Contract.VerifyProof(&_PoolKeyVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x449ccd1e.
//
// Solidity: function verifyProof(bytes proof, uint256[8] input) view returns()
func (_PoolKeyVerifier *PoolKeyVerifierCallerSession) VerifyProof(proof []byte, input [8]*big.Int) error {
	return _PoolKeyVerifier.Contract.VerifyProof(&_PoolKeyVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_PoolKeyVerifier *PoolKeyVerifierCaller) VerifyProof0(opts *bind.CallOpts, proof []byte, input []byte) error {
	var out []interface{}
	err := _PoolKeyVerifier.contract.Call(opts, &out, "verifyProof0", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_PoolKeyVerifier *PoolKeyVerifierSession) VerifyProof0(proof []byte, input []byte) error {
	return _PoolKeyVerifier.Contract.VerifyProof0(&_PoolKeyVerifier.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes proof, bytes input) view returns()
func (_PoolKeyVerifier *PoolKeyVerifierCallerSession) VerifyProof0(proof []byte, input []byte) error {
	return _PoolKeyVerifier.Contract.VerifyProof0(&_PoolKeyVerifier.CallOpts, proof, input)
}
