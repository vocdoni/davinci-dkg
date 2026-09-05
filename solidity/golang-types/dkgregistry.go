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

// IDKGRegistryNodeKey is an auto generated low-level Go binding around an user-defined struct.
type IDKGRegistryNodeKey struct {
	Operator          common.Address
	PubX              *big.Int
	PubY              *big.Int
	Status            uint8
	LastActiveBlock   uint64
	RegisteredAtBlock uint64
}

// DKGRegistryMetaData contains all meta data concerning the DKGRegistry contract.
var DKGRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"inactivityWindow\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"INACTIVITY_WINDOW\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activeCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNode\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDKGRegistry.NodeKey\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pubX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pubY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIDKGRegistry.NodeStatus\"},{\"name\":\"lastActiveBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"registeredAtBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"heartbeat\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isActive\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"manager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"markActive\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nodeCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reactivate\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reap\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerKey\",\"inputs\":[{\"name\":\"pubX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pubY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrAx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrAy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrZ\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setManager\",\"inputs\":[{\"name\":\"m\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateKey\",\"inputs\":[{\"name\":\"pubX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pubY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrAx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrAy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrZ\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ManagerSet\",\"inputs\":[{\"name\":\"manager\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeMarkedActive\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"atBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeReactivated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeReaped\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"lastActiveBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubX\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"pubY\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubX\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"pubY\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSchnorrProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ManagerAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ManagerNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInactive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PointIsIdentity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PointNotCanonical\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PointNotInSubgroup\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PointNotOnCurve\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StillActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x60c03461009357601f6115e838819003918201601f19168301916001600160401b038311848410176100975780849260209460405283398101031261009357516001600160401b0381168082036100935715610084576080523360a05260405161153c90816100ac82396080518181816105c101526108d7015260a051816101e60152f35b630eda9c3d60e31b5f5260045ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe6080806040526004361015610012575f80fd5b5f3560e01c90816303ac08f6146108c5575080633defb962146108615780634331ed1f14610837578063481c6a751461080f5780636da49b83146107e757806381c7362d1461068b5780638af9f493146105695780639d209048146104465780639f8a13d7146103f8578063b82856e11461028f578063d0ebdbe7146101be578063d18611d6146100de5763f06f37bc146100ab575f80fd5b346100da5760203660031901126100da576004356001600160a01b03811681036100da576100d890610a1f565b005b5f80fd5b346100da575f3660031901126100da57335f525f602052600360405f200160ff81541660038110156101aa5760020361019b57805460ff19166001178155436001600160401b03169061013c9082906101378282610928565b61094d565b60018054600160401b600160801b03198116604091821c6001600160401b03168301821b600160401b600160801b0316179091555190335f5160206114e75f395f51905f525f80a281525f5160206114875f395f51905f5260203392a2005b63442d617b60e11b5f5260045ffd5b634e487b7160e01b5f52602160045260245ffd5b346100da5760203660031901126100da576004356001600160a01b038116908190036100da577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633036102815760025460ff8160a01c16610272578115610263576001600160a81b0319168117600160a01b176002557f60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa695f80a2005b63e6c4247b60e01b5f5260045ffd5b634294267360e11b5f5260045ffd5b6282b42960e81b5f5260045ffd5b346100da5761029d36610906565b939291906102ab8385610ac1565b6102b58183610aeb565b335f525f60205260405f2091600383019560ff87541660038110156101aa576103e957828287926102ed6102f3968996878733610ddc565b90610e9f565b156103da5780546001600160a01b031916331781556001808201849055600291909101829055835460ff1916178355436001600160401b03169261033d9084906101378282610928565b600180546001600160801b031981166001600160401b03808316840181169182176001600160401b0319909316909117604090811c9091168301811b600160401b600160801b0316919091179091558051928352602083019190915233917f99140a41575033d78b1016979e49f1b8a4943ef274d75edf0dac1bc3ccbce5f79190a26040519081525f5160206114875f395f51905f5260203392a2005b6327f7eb4d60e11b5f5260045ffd5b630ea075bf60e21b5f5260045ffd5b346100da5760203660031901126100da576004356001600160a01b038116908190036100da575f525f60205260ff600360405f2001541660038110156101aa57602090600160405191148152f35b346100da5760203660031901126100da576004356001600160a01b038116908190036100da575f60a060405161047b81610997565b82815282602082015282604082015282606082015282608082015201525f525f60205260405f206040516104ae81610997565b81546001600160a01b03168152600182015460208201908152600283015460408301908152600393840154919360608401919060ff8416908110156101aa5782526001600160401b03600884901c81166080860190815260489490941c1660a085019081526040805195516001600160a01b031686529551602086015290519484019490945251929060038410156101aa57606083019390935291516001600160401b039081166080830152915190911660a082015260c090f35b346100da5760203660031901126100da576004356001600160a01b038116908190036100da57805f525f602052600360405f2001805460ff811660038110156101aa5760010361067c576105ec906001600160401b037f000000000000000000000000000000000000000000000000000000000000000081169160081c16610976565b43111561066d57805460ff1916600217815560018054600160401b600160801b03198116604091821c6001600160401b039081165f1901831b600160401b600160801b0316919091179092559154915160089290921c1681527f17b35aacc7270dcc7c9993688488c4a6267c1ca2e4ab73b83c6411855a54376f90602090a2005b63785bbc6d60e11b5f5260045ffd5b634065aaf160e11b5f5260045ffd5b346100da5761069936610906565b939291906106a78385610ac1565b6106b18183610aeb565b335f525f60205260405f2091600383019560ff87541660038110156101aa57156107d857828287926102ed6106ea968996878733610ddc565b156103da5760018101839055600201819055436001600160401b0316926107118482610928565b60ff81541660038110156101aa576002859114610777575b505060405191825260208201527f1e2215a8512058e371c99f86c2731c45755267c9d5fb9eb3c911230fa9b55cfc60403392a26040519081525f5160206114875f395f51905f5260203392a2005b815460ff1916600117825561078b9161094d565b60018054600160401b600160801b03198116604091821c6001600160401b0316830190911b600160401b600160801b0316179055335f5160206114e75f395f51905f525f80a28383610729565b63aba4733960e01b5f5260045ffd5b346100da575f3660031901126100da576001546040516001600160401b039091168152602090f35b346100da575f3660031901126100da576002546040516001600160a01b039091168152602090f35b346100da575f3660031901126100da576001546040805191901c6001600160401b03168152602090f35b346100da575f3660031901126100da57335f525f602052600360405f200160ff81541660038110156101aa5760010361067c57436001600160401b0316906108aa908290610928565b6040519081525f5160206114875f395f51905f5260203392a2005b346100da575f3660031901126100da577f00000000000000000000000000000000000000000000000000000000000000006001600160401b03168152602090f35b60a09060031901126100da576004359060243590604435906064359060843590565b8054610100600160481b03191660089290921b610100600160481b0316919091179055565b8054600160481b600160881b03191660489290921b600160481b600160881b0316919091179055565b9190820180921161098357565b634e487b7160e01b5f52601160045260245ffd5b60c081019081106001600160401b038211176109b257604052565b634e487b7160e01b5f52604160045260245ffd5b608081019081106001600160401b038211176109b257604052565b604081019081106001600160401b038211176109b257604052565b601f909101601f19168101906001600160401b038211908210176109b257604052565b6002546001600160a01b03168015610ab2573303610aa35760018060a01b0316805f525f602052600360405f200180549060ff821660038110156101aa57600103610a9e576001600160401b034381169260081c168214610a9e5781610a955f5160206114875f395f51905f5293602093610928565b604051908152a2565b505050565b63607e454560e11b5f5260045ffd5b6321f7ab5360e01b5f5260045ffd5b90610ad591610ad08282610aeb565b610b72565b15610adc57565b63b28e789160e01b5f5260045ffd5b905f5160206114a75f395f51905f528210801590610b5b575b610b4c57811580610b42575b610b3357610b1d91610d11565b15610b2457565b63a3d28e1360e01b5f5260045ffd5b6332d0802760e11b5f5260045ffd5b5060018114610b10565b630593753d60e51b5f5260045ffd5b505f5160206114a75f395f51905f52811015610b04565b801580610d07575b610d0057610b888282610d11565b15610cfa57610b95610e56565b50610b9e6112d2565b908160809360405194610bb26080876109fc565b5f5b818110610cde57505091610bd282610bf4946020880193845161132b565b6040850190610be5838351835190611366565b606086015191519051916113c7565b604051610c00816109c6565b5f81526020810192600184526040820192600184525f60608401525f91610100805b610c51575050505051159182610c45575b5081610c3d575090565b905051151590565b5181511491505f610c33565b60011901805f5160206114c75f395f51905f52811c6003168515610caf57610c7a858880611366565b610c85858880611366565b8481610c93575b5050610c22565b610ca0610ca892866112c1565b5188806113c7565b5f84610c8c565b809150610cbe575b5080610c22565b610cd5919450610cce90836112c1565b5185611462565b82600193610cb7565b602091929350610cec610e56565b818801520190849291610bb4565b50505f90565b5050600190565b5060018214610b7a565b5f5160206114a75f395f51905f528110801590610da7575b610cfa575f5160206114a75f395f51905f528181920991800990805f5160206114a75f395f51905f5203915f5160206114a75f395f51905f528311610983575f5160206114a75f395f51905f528080838195097f1aee90f15f2189693df072d799fd11fc039b2959ebb7c867d075ca8cf4d7eb8e0960010892081490565b505f5160206114a75f395f51905f52821015610d29565b8115610dc8570690565b634e487b7160e01b5f52601260045260245ffd5b925f5160206114c75f395f51905f529491926040519360208501957f4599aabb337c91d65fe440ef7e20c6dcc72c2459fd0901c45add50b08b3bb34d875260018060601b03199060601b16604086015260548501526074840152609483015260b482015260b48152610e4f60d4826109fc565b5190200690565b60405190610e63826109c6565b5f6060838281528260208201528260408201520152565b906010811015610e8b5760051b0190565b634e487b7160e01b5f52603260045260245ffd5b94939190945f5160206114a75f395f51905f528210806112ab575b158015611278575b61126e575f5160206114c75f395f51905f5280910695069180155f1461124557505f945b610eee6112d2565b956102009560405196610f0181896109fc565b5f5b81811061123057505060016020885101526001604088510152608087015188518015610dc857806060917f1561ff836ce19d358a4eb7a4c199e94c377c749ae6f2a277f1f9195afe553f9f068352807f25797203f7a0b24925572e1cd16bf9edfce0051fb9e133774b3c257a872d7d8b068060208501526001604085015283510991015261010087015188518015610dc857806060917f0daaa7e6b25c28e6dc8dd1d48e9cc61cd07015c1d7c1b8d4590eb6f51d5346dc068352807f01666cafbf0a30da8b9ebeaf848a1da067a892296f1043188e1705402b6d68530680602085015260016040850152835109910152610180870151918851928315610dc85789809361107792606087611099987f136d609c4c856f5d277fab08c730cbdd1a776ce4728c6a2eb20ff22bccf26894068352807f21d66f0e2295ae954494f25889f9319cc1b4df71eff3f46ba9e4631b43fd7c95068060208501526001604085015283510991015260208b0193845161132b565b604088019061108a838351835190611366565b606089015191519051916113c7565b60045b600c8111156111c75750604051956110b3876109c6565b5f87526020870195600187526040880194600186525f60608a01525f9360fc805b6111365750505050505081518015159586611118575b505050836110f9575b50505090565b5f5160206114a75f395f51905f529293505190099051145f80806110f3565b519295505f5160206114a75f395f51905f52910914925f80806110ea565b600119018082811c60021b600c1684821c60031617868d8915611196579081611163828261116995611366565b80611366565b868d82611179575b5050506110d4565b61118661118e9389610e7a565b5190806113c7565b5f868d611171565b50508091506111a7575b50806110d4565b6111be9196506111b79085610e7a565b518b611462565b846001956111a0565b60015b600481106111f35750600481018091111561109c57634e487b7160e01b5f52601160045260245ffd5b8061122a8961120d61120760019587610976565b8b610e7a565b51611218868c610e7a565b51611223858d610e7a565b51916113c7565b016111ca565b60209061123b610e56565b818b015201610f03565b5f5160206114a75f395f51905f52035f5160206114a75f395f51905f5281116109835794610ee6565b5050505050505f90565b505f5160206114a75f395f51905f52831080611295575b15610ec2565b505f5160206114a75f395f51905f52851061128f565b505f5160206114a75f395f51905f528410610eba565b906004811015610e8b5760051b0190565b5f60206040516112e1816109e1565b82815201526040516112f2816109e1565b5f5160206114a75f395f51905f5281527f0578d36fdd1172a8c3909ff8b278cb9adf026a6b5db6203e5d099f85f9afd71b602082015290565b929061133f92611346925193848093610dbe565b8552610dbe565b90816020840152600160408401528251918115610dc85760609209910152565b9151815191602081019081518315610dc8578380808093604098088180808a818080808c5180099c518009818d810382089c08810380988782980908980151800980088103870894828682098a520960608801528309602086015209910152565b92908151926020820192835183518603918615610dc85786808086818080999881808d81809d9c816020819f01968188518c51820390089208099f519051900891518551900890099a818181038d089b089687958160608c0151606085015190099060200151900998604001519060400151900980089581818103880896089582868209895209606087015283096020850152099060400152565b9060608091805184526020810151602085015260408101516040850152015191015256fe02c36b03f66c867a89d996a43b2ea1f9c0e5740578642d17ef1b3d259073e72c30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001060c89ce5c263405370a08b6d0302b0bab3eedb83920ee0a677297dc392126f1f979d653049f5a10edd541959ecb5c2ced8fd1b0adaefc8fd66744fe70742052a264697066735822122002b68f0ef8fe437ebae4c283e822b94deebc4a80d8dd25e37e813a48e0ebb26864736f6c634300081c0033",
}

// DKGRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use DKGRegistryMetaData.ABI instead.
var DKGRegistryABI = DKGRegistryMetaData.ABI

// DKGRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DKGRegistryMetaData.Bin instead.
var DKGRegistryBin = DKGRegistryMetaData.Bin

// DeployDKGRegistry deploys a new Ethereum contract, binding an instance of DKGRegistry to it.
func DeployDKGRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, inactivityWindow uint64) (common.Address, *types.Transaction, *DKGRegistry, error) {
	parsed, err := DKGRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DKGRegistryBin), backend, inactivityWindow)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DKGRegistry{DKGRegistryCaller: DKGRegistryCaller{contract: contract}, DKGRegistryTransactor: DKGRegistryTransactor{contract: contract}, DKGRegistryFilterer: DKGRegistryFilterer{contract: contract}}, nil
}

// DKGRegistry is an auto generated Go binding around an Ethereum contract.
type DKGRegistry struct {
	DKGRegistryCaller     // Read-only binding to the contract
	DKGRegistryTransactor // Write-only binding to the contract
	DKGRegistryFilterer   // Log filterer for contract events
}

// DKGRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DKGRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DKGRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DKGRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DKGRegistrySession struct {
	Contract     *DKGRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DKGRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DKGRegistryCallerSession struct {
	Contract *DKGRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DKGRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DKGRegistryTransactorSession struct {
	Contract     *DKGRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DKGRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DKGRegistryRaw struct {
	Contract *DKGRegistry // Generic contract binding to access the raw methods on
}

// DKGRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DKGRegistryCallerRaw struct {
	Contract *DKGRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// DKGRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DKGRegistryTransactorRaw struct {
	Contract *DKGRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDKGRegistry creates a new instance of DKGRegistry, bound to a specific deployed contract.
func NewDKGRegistry(address common.Address, backend bind.ContractBackend) (*DKGRegistry, error) {
	contract, err := bindDKGRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DKGRegistry{DKGRegistryCaller: DKGRegistryCaller{contract: contract}, DKGRegistryTransactor: DKGRegistryTransactor{contract: contract}, DKGRegistryFilterer: DKGRegistryFilterer{contract: contract}}, nil
}

// NewDKGRegistryCaller creates a new read-only instance of DKGRegistry, bound to a specific deployed contract.
func NewDKGRegistryCaller(address common.Address, caller bind.ContractCaller) (*DKGRegistryCaller, error) {
	contract, err := bindDKGRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DKGRegistryCaller{contract: contract}, nil
}

// NewDKGRegistryTransactor creates a new write-only instance of DKGRegistry, bound to a specific deployed contract.
func NewDKGRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*DKGRegistryTransactor, error) {
	contract, err := bindDKGRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DKGRegistryTransactor{contract: contract}, nil
}

// NewDKGRegistryFilterer creates a new log filterer instance of DKGRegistry, bound to a specific deployed contract.
func NewDKGRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*DKGRegistryFilterer, error) {
	contract, err := bindDKGRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DKGRegistryFilterer{contract: contract}, nil
}

// bindDKGRegistry binds a generic wrapper to an already deployed contract.
func bindDKGRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DKGRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DKGRegistry *DKGRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKGRegistry.Contract.DKGRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DKGRegistry *DKGRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKGRegistry.Contract.DKGRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DKGRegistry *DKGRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKGRegistry.Contract.DKGRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DKGRegistry *DKGRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKGRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DKGRegistry *DKGRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKGRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DKGRegistry *DKGRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKGRegistry.Contract.contract.Transact(opts, method, params...)
}

// INACTIVITYWINDOW is a free data retrieval call binding the contract method 0x03ac08f6.
//
// Solidity: function INACTIVITY_WINDOW() view returns(uint64)
func (_DKGRegistry *DKGRegistryCaller) INACTIVITYWINDOW(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DKGRegistry.contract.Call(opts, &out, "INACTIVITY_WINDOW")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// INACTIVITYWINDOW is a free data retrieval call binding the contract method 0x03ac08f6.
//
// Solidity: function INACTIVITY_WINDOW() view returns(uint64)
func (_DKGRegistry *DKGRegistrySession) INACTIVITYWINDOW() (uint64, error) {
	return _DKGRegistry.Contract.INACTIVITYWINDOW(&_DKGRegistry.CallOpts)
}

// INACTIVITYWINDOW is a free data retrieval call binding the contract method 0x03ac08f6.
//
// Solidity: function INACTIVITY_WINDOW() view returns(uint64)
func (_DKGRegistry *DKGRegistryCallerSession) INACTIVITYWINDOW() (uint64, error) {
	return _DKGRegistry.Contract.INACTIVITYWINDOW(&_DKGRegistry.CallOpts)
}

// ActiveCount is a free data retrieval call binding the contract method 0x4331ed1f.
//
// Solidity: function activeCount() view returns(uint64)
func (_DKGRegistry *DKGRegistryCaller) ActiveCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DKGRegistry.contract.Call(opts, &out, "activeCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ActiveCount is a free data retrieval call binding the contract method 0x4331ed1f.
//
// Solidity: function activeCount() view returns(uint64)
func (_DKGRegistry *DKGRegistrySession) ActiveCount() (uint64, error) {
	return _DKGRegistry.Contract.ActiveCount(&_DKGRegistry.CallOpts)
}

// ActiveCount is a free data retrieval call binding the contract method 0x4331ed1f.
//
// Solidity: function activeCount() view returns(uint64)
func (_DKGRegistry *DKGRegistryCallerSession) ActiveCount() (uint64, error) {
	return _DKGRegistry.Contract.ActiveCount(&_DKGRegistry.CallOpts)
}

// GetNode is a free data retrieval call binding the contract method 0x9d209048.
//
// Solidity: function getNode(address operator) view returns((address,uint256,uint256,uint8,uint64,uint64))
func (_DKGRegistry *DKGRegistryCaller) GetNode(opts *bind.CallOpts, operator common.Address) (IDKGRegistryNodeKey, error) {
	var out []interface{}
	err := _DKGRegistry.contract.Call(opts, &out, "getNode", operator)

	if err != nil {
		return *new(IDKGRegistryNodeKey), err
	}

	out0 := *abi.ConvertType(out[0], new(IDKGRegistryNodeKey)).(*IDKGRegistryNodeKey)

	return out0, err

}

// GetNode is a free data retrieval call binding the contract method 0x9d209048.
//
// Solidity: function getNode(address operator) view returns((address,uint256,uint256,uint8,uint64,uint64))
func (_DKGRegistry *DKGRegistrySession) GetNode(operator common.Address) (IDKGRegistryNodeKey, error) {
	return _DKGRegistry.Contract.GetNode(&_DKGRegistry.CallOpts, operator)
}

// GetNode is a free data retrieval call binding the contract method 0x9d209048.
//
// Solidity: function getNode(address operator) view returns((address,uint256,uint256,uint8,uint64,uint64))
func (_DKGRegistry *DKGRegistryCallerSession) GetNode(operator common.Address) (IDKGRegistryNodeKey, error) {
	return _DKGRegistry.Contract.GetNode(&_DKGRegistry.CallOpts, operator)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address operator) view returns(bool)
func (_DKGRegistry *DKGRegistryCaller) IsActive(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _DKGRegistry.contract.Call(opts, &out, "isActive", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address operator) view returns(bool)
func (_DKGRegistry *DKGRegistrySession) IsActive(operator common.Address) (bool, error) {
	return _DKGRegistry.Contract.IsActive(&_DKGRegistry.CallOpts, operator)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address operator) view returns(bool)
func (_DKGRegistry *DKGRegistryCallerSession) IsActive(operator common.Address) (bool, error) {
	return _DKGRegistry.Contract.IsActive(&_DKGRegistry.CallOpts, operator)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_DKGRegistry *DKGRegistryCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKGRegistry.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_DKGRegistry *DKGRegistrySession) Manager() (common.Address, error) {
	return _DKGRegistry.Contract.Manager(&_DKGRegistry.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_DKGRegistry *DKGRegistryCallerSession) Manager() (common.Address, error) {
	return _DKGRegistry.Contract.Manager(&_DKGRegistry.CallOpts)
}

// NodeCount is a free data retrieval call binding the contract method 0x6da49b83.
//
// Solidity: function nodeCount() view returns(uint64)
func (_DKGRegistry *DKGRegistryCaller) NodeCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DKGRegistry.contract.Call(opts, &out, "nodeCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NodeCount is a free data retrieval call binding the contract method 0x6da49b83.
//
// Solidity: function nodeCount() view returns(uint64)
func (_DKGRegistry *DKGRegistrySession) NodeCount() (uint64, error) {
	return _DKGRegistry.Contract.NodeCount(&_DKGRegistry.CallOpts)
}

// NodeCount is a free data retrieval call binding the contract method 0x6da49b83.
//
// Solidity: function nodeCount() view returns(uint64)
func (_DKGRegistry *DKGRegistryCallerSession) NodeCount() (uint64, error) {
	return _DKGRegistry.Contract.NodeCount(&_DKGRegistry.CallOpts)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x3defb962.
//
// Solidity: function heartbeat() returns()
func (_DKGRegistry *DKGRegistryTransactor) Heartbeat(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKGRegistry.contract.Transact(opts, "heartbeat")
}

// Heartbeat is a paid mutator transaction binding the contract method 0x3defb962.
//
// Solidity: function heartbeat() returns()
func (_DKGRegistry *DKGRegistrySession) Heartbeat() (*types.Transaction, error) {
	return _DKGRegistry.Contract.Heartbeat(&_DKGRegistry.TransactOpts)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x3defb962.
//
// Solidity: function heartbeat() returns()
func (_DKGRegistry *DKGRegistryTransactorSession) Heartbeat() (*types.Transaction, error) {
	return _DKGRegistry.Contract.Heartbeat(&_DKGRegistry.TransactOpts)
}

// MarkActive is a paid mutator transaction binding the contract method 0xf06f37bc.
//
// Solidity: function markActive(address operator) returns()
func (_DKGRegistry *DKGRegistryTransactor) MarkActive(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _DKGRegistry.contract.Transact(opts, "markActive", operator)
}

// MarkActive is a paid mutator transaction binding the contract method 0xf06f37bc.
//
// Solidity: function markActive(address operator) returns()
func (_DKGRegistry *DKGRegistrySession) MarkActive(operator common.Address) (*types.Transaction, error) {
	return _DKGRegistry.Contract.MarkActive(&_DKGRegistry.TransactOpts, operator)
}

// MarkActive is a paid mutator transaction binding the contract method 0xf06f37bc.
//
// Solidity: function markActive(address operator) returns()
func (_DKGRegistry *DKGRegistryTransactorSession) MarkActive(operator common.Address) (*types.Transaction, error) {
	return _DKGRegistry.Contract.MarkActive(&_DKGRegistry.TransactOpts, operator)
}

// Reactivate is a paid mutator transaction binding the contract method 0xd18611d6.
//
// Solidity: function reactivate() returns()
func (_DKGRegistry *DKGRegistryTransactor) Reactivate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKGRegistry.contract.Transact(opts, "reactivate")
}

// Reactivate is a paid mutator transaction binding the contract method 0xd18611d6.
//
// Solidity: function reactivate() returns()
func (_DKGRegistry *DKGRegistrySession) Reactivate() (*types.Transaction, error) {
	return _DKGRegistry.Contract.Reactivate(&_DKGRegistry.TransactOpts)
}

// Reactivate is a paid mutator transaction binding the contract method 0xd18611d6.
//
// Solidity: function reactivate() returns()
func (_DKGRegistry *DKGRegistryTransactorSession) Reactivate() (*types.Transaction, error) {
	return _DKGRegistry.Contract.Reactivate(&_DKGRegistry.TransactOpts)
}

// Reap is a paid mutator transaction binding the contract method 0x8af9f493.
//
// Solidity: function reap(address operator) returns()
func (_DKGRegistry *DKGRegistryTransactor) Reap(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _DKGRegistry.contract.Transact(opts, "reap", operator)
}

// Reap is a paid mutator transaction binding the contract method 0x8af9f493.
//
// Solidity: function reap(address operator) returns()
func (_DKGRegistry *DKGRegistrySession) Reap(operator common.Address) (*types.Transaction, error) {
	return _DKGRegistry.Contract.Reap(&_DKGRegistry.TransactOpts, operator)
}

// Reap is a paid mutator transaction binding the contract method 0x8af9f493.
//
// Solidity: function reap(address operator) returns()
func (_DKGRegistry *DKGRegistryTransactorSession) Reap(operator common.Address) (*types.Transaction, error) {
	return _DKGRegistry.Contract.Reap(&_DKGRegistry.TransactOpts, operator)
}

// RegisterKey is a paid mutator transaction binding the contract method 0xb82856e1.
//
// Solidity: function registerKey(uint256 pubX, uint256 pubY, uint256 schnorrAx, uint256 schnorrAy, uint256 schnorrZ) returns()
func (_DKGRegistry *DKGRegistryTransactor) RegisterKey(opts *bind.TransactOpts, pubX *big.Int, pubY *big.Int, schnorrAx *big.Int, schnorrAy *big.Int, schnorrZ *big.Int) (*types.Transaction, error) {
	return _DKGRegistry.contract.Transact(opts, "registerKey", pubX, pubY, schnorrAx, schnorrAy, schnorrZ)
}

// RegisterKey is a paid mutator transaction binding the contract method 0xb82856e1.
//
// Solidity: function registerKey(uint256 pubX, uint256 pubY, uint256 schnorrAx, uint256 schnorrAy, uint256 schnorrZ) returns()
func (_DKGRegistry *DKGRegistrySession) RegisterKey(pubX *big.Int, pubY *big.Int, schnorrAx *big.Int, schnorrAy *big.Int, schnorrZ *big.Int) (*types.Transaction, error) {
	return _DKGRegistry.Contract.RegisterKey(&_DKGRegistry.TransactOpts, pubX, pubY, schnorrAx, schnorrAy, schnorrZ)
}

// RegisterKey is a paid mutator transaction binding the contract method 0xb82856e1.
//
// Solidity: function registerKey(uint256 pubX, uint256 pubY, uint256 schnorrAx, uint256 schnorrAy, uint256 schnorrZ) returns()
func (_DKGRegistry *DKGRegistryTransactorSession) RegisterKey(pubX *big.Int, pubY *big.Int, schnorrAx *big.Int, schnorrAy *big.Int, schnorrZ *big.Int) (*types.Transaction, error) {
	return _DKGRegistry.Contract.RegisterKey(&_DKGRegistry.TransactOpts, pubX, pubY, schnorrAx, schnorrAy, schnorrZ)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address m) returns()
func (_DKGRegistry *DKGRegistryTransactor) SetManager(opts *bind.TransactOpts, m common.Address) (*types.Transaction, error) {
	return _DKGRegistry.contract.Transact(opts, "setManager", m)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address m) returns()
func (_DKGRegistry *DKGRegistrySession) SetManager(m common.Address) (*types.Transaction, error) {
	return _DKGRegistry.Contract.SetManager(&_DKGRegistry.TransactOpts, m)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address m) returns()
func (_DKGRegistry *DKGRegistryTransactorSession) SetManager(m common.Address) (*types.Transaction, error) {
	return _DKGRegistry.Contract.SetManager(&_DKGRegistry.TransactOpts, m)
}

// UpdateKey is a paid mutator transaction binding the contract method 0x81c7362d.
//
// Solidity: function updateKey(uint256 pubX, uint256 pubY, uint256 schnorrAx, uint256 schnorrAy, uint256 schnorrZ) returns()
func (_DKGRegistry *DKGRegistryTransactor) UpdateKey(opts *bind.TransactOpts, pubX *big.Int, pubY *big.Int, schnorrAx *big.Int, schnorrAy *big.Int, schnorrZ *big.Int) (*types.Transaction, error) {
	return _DKGRegistry.contract.Transact(opts, "updateKey", pubX, pubY, schnorrAx, schnorrAy, schnorrZ)
}

// UpdateKey is a paid mutator transaction binding the contract method 0x81c7362d.
//
// Solidity: function updateKey(uint256 pubX, uint256 pubY, uint256 schnorrAx, uint256 schnorrAy, uint256 schnorrZ) returns()
func (_DKGRegistry *DKGRegistrySession) UpdateKey(pubX *big.Int, pubY *big.Int, schnorrAx *big.Int, schnorrAy *big.Int, schnorrZ *big.Int) (*types.Transaction, error) {
	return _DKGRegistry.Contract.UpdateKey(&_DKGRegistry.TransactOpts, pubX, pubY, schnorrAx, schnorrAy, schnorrZ)
}

// UpdateKey is a paid mutator transaction binding the contract method 0x81c7362d.
//
// Solidity: function updateKey(uint256 pubX, uint256 pubY, uint256 schnorrAx, uint256 schnorrAy, uint256 schnorrZ) returns()
func (_DKGRegistry *DKGRegistryTransactorSession) UpdateKey(pubX *big.Int, pubY *big.Int, schnorrAx *big.Int, schnorrAy *big.Int, schnorrZ *big.Int) (*types.Transaction, error) {
	return _DKGRegistry.Contract.UpdateKey(&_DKGRegistry.TransactOpts, pubX, pubY, schnorrAx, schnorrAy, schnorrZ)
}

// DKGRegistryManagerSetIterator is returned from FilterManagerSet and is used to iterate over the raw logs and unpacked data for ManagerSet events raised by the DKGRegistry contract.
type DKGRegistryManagerSetIterator struct {
	Event *DKGRegistryManagerSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DKGRegistryManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRegistryManagerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DKGRegistryManagerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DKGRegistryManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRegistryManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRegistryManagerSet represents a ManagerSet event raised by the DKGRegistry contract.
type DKGRegistryManagerSet struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerSet is a free log retrieval operation binding the contract event 0x60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa69.
//
// Solidity: event ManagerSet(address indexed manager)
func (_DKGRegistry *DKGRegistryFilterer) FilterManagerSet(opts *bind.FilterOpts, manager []common.Address) (*DKGRegistryManagerSetIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _DKGRegistry.contract.FilterLogs(opts, "ManagerSet", managerRule)
	if err != nil {
		return nil, err
	}
	return &DKGRegistryManagerSetIterator{contract: _DKGRegistry.contract, event: "ManagerSet", logs: logs, sub: sub}, nil
}

// WatchManagerSet is a free log subscription operation binding the contract event 0x60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa69.
//
// Solidity: event ManagerSet(address indexed manager)
func (_DKGRegistry *DKGRegistryFilterer) WatchManagerSet(opts *bind.WatchOpts, sink chan<- *DKGRegistryManagerSet, manager []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _DKGRegistry.contract.WatchLogs(opts, "ManagerSet", managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRegistryManagerSet)
				if err := _DKGRegistry.contract.UnpackLog(event, "ManagerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseManagerSet is a log parse operation binding the contract event 0x60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa69.
//
// Solidity: event ManagerSet(address indexed manager)
func (_DKGRegistry *DKGRegistryFilterer) ParseManagerSet(log types.Log) (*DKGRegistryManagerSet, error) {
	event := new(DKGRegistryManagerSet)
	if err := _DKGRegistry.contract.UnpackLog(event, "ManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRegistryNodeMarkedActiveIterator is returned from FilterNodeMarkedActive and is used to iterate over the raw logs and unpacked data for NodeMarkedActive events raised by the DKGRegistry contract.
type DKGRegistryNodeMarkedActiveIterator struct {
	Event *DKGRegistryNodeMarkedActive // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DKGRegistryNodeMarkedActiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRegistryNodeMarkedActive)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DKGRegistryNodeMarkedActive)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DKGRegistryNodeMarkedActiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRegistryNodeMarkedActiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRegistryNodeMarkedActive represents a NodeMarkedActive event raised by the DKGRegistry contract.
type DKGRegistryNodeMarkedActive struct {
	Operator common.Address
	AtBlock  uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNodeMarkedActive is a free log retrieval operation binding the contract event 0x02c36b03f66c867a89d996a43b2ea1f9c0e5740578642d17ef1b3d259073e72c.
//
// Solidity: event NodeMarkedActive(address indexed operator, uint64 atBlock)
func (_DKGRegistry *DKGRegistryFilterer) FilterNodeMarkedActive(opts *bind.FilterOpts, operator []common.Address) (*DKGRegistryNodeMarkedActiveIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.FilterLogs(opts, "NodeMarkedActive", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DKGRegistryNodeMarkedActiveIterator{contract: _DKGRegistry.contract, event: "NodeMarkedActive", logs: logs, sub: sub}, nil
}

// WatchNodeMarkedActive is a free log subscription operation binding the contract event 0x02c36b03f66c867a89d996a43b2ea1f9c0e5740578642d17ef1b3d259073e72c.
//
// Solidity: event NodeMarkedActive(address indexed operator, uint64 atBlock)
func (_DKGRegistry *DKGRegistryFilterer) WatchNodeMarkedActive(opts *bind.WatchOpts, sink chan<- *DKGRegistryNodeMarkedActive, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.WatchLogs(opts, "NodeMarkedActive", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRegistryNodeMarkedActive)
				if err := _DKGRegistry.contract.UnpackLog(event, "NodeMarkedActive", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeMarkedActive is a log parse operation binding the contract event 0x02c36b03f66c867a89d996a43b2ea1f9c0e5740578642d17ef1b3d259073e72c.
//
// Solidity: event NodeMarkedActive(address indexed operator, uint64 atBlock)
func (_DKGRegistry *DKGRegistryFilterer) ParseNodeMarkedActive(log types.Log) (*DKGRegistryNodeMarkedActive, error) {
	event := new(DKGRegistryNodeMarkedActive)
	if err := _DKGRegistry.contract.UnpackLog(event, "NodeMarkedActive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRegistryNodeReactivatedIterator is returned from FilterNodeReactivated and is used to iterate over the raw logs and unpacked data for NodeReactivated events raised by the DKGRegistry contract.
type DKGRegistryNodeReactivatedIterator struct {
	Event *DKGRegistryNodeReactivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DKGRegistryNodeReactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRegistryNodeReactivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DKGRegistryNodeReactivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DKGRegistryNodeReactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRegistryNodeReactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRegistryNodeReactivated represents a NodeReactivated event raised by the DKGRegistry contract.
type DKGRegistryNodeReactivated struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNodeReactivated is a free log retrieval operation binding the contract event 0xf979d653049f5a10edd541959ecb5c2ced8fd1b0adaefc8fd66744fe70742052.
//
// Solidity: event NodeReactivated(address indexed operator)
func (_DKGRegistry *DKGRegistryFilterer) FilterNodeReactivated(opts *bind.FilterOpts, operator []common.Address) (*DKGRegistryNodeReactivatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.FilterLogs(opts, "NodeReactivated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DKGRegistryNodeReactivatedIterator{contract: _DKGRegistry.contract, event: "NodeReactivated", logs: logs, sub: sub}, nil
}

// WatchNodeReactivated is a free log subscription operation binding the contract event 0xf979d653049f5a10edd541959ecb5c2ced8fd1b0adaefc8fd66744fe70742052.
//
// Solidity: event NodeReactivated(address indexed operator)
func (_DKGRegistry *DKGRegistryFilterer) WatchNodeReactivated(opts *bind.WatchOpts, sink chan<- *DKGRegistryNodeReactivated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.WatchLogs(opts, "NodeReactivated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRegistryNodeReactivated)
				if err := _DKGRegistry.contract.UnpackLog(event, "NodeReactivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeReactivated is a log parse operation binding the contract event 0xf979d653049f5a10edd541959ecb5c2ced8fd1b0adaefc8fd66744fe70742052.
//
// Solidity: event NodeReactivated(address indexed operator)
func (_DKGRegistry *DKGRegistryFilterer) ParseNodeReactivated(log types.Log) (*DKGRegistryNodeReactivated, error) {
	event := new(DKGRegistryNodeReactivated)
	if err := _DKGRegistry.contract.UnpackLog(event, "NodeReactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRegistryNodeReapedIterator is returned from FilterNodeReaped and is used to iterate over the raw logs and unpacked data for NodeReaped events raised by the DKGRegistry contract.
type DKGRegistryNodeReapedIterator struct {
	Event *DKGRegistryNodeReaped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DKGRegistryNodeReapedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRegistryNodeReaped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DKGRegistryNodeReaped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DKGRegistryNodeReapedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRegistryNodeReapedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRegistryNodeReaped represents a NodeReaped event raised by the DKGRegistry contract.
type DKGRegistryNodeReaped struct {
	Operator        common.Address
	LastActiveBlock uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodeReaped is a free log retrieval operation binding the contract event 0x17b35aacc7270dcc7c9993688488c4a6267c1ca2e4ab73b83c6411855a54376f.
//
// Solidity: event NodeReaped(address indexed operator, uint64 lastActiveBlock)
func (_DKGRegistry *DKGRegistryFilterer) FilterNodeReaped(opts *bind.FilterOpts, operator []common.Address) (*DKGRegistryNodeReapedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.FilterLogs(opts, "NodeReaped", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DKGRegistryNodeReapedIterator{contract: _DKGRegistry.contract, event: "NodeReaped", logs: logs, sub: sub}, nil
}

// WatchNodeReaped is a free log subscription operation binding the contract event 0x17b35aacc7270dcc7c9993688488c4a6267c1ca2e4ab73b83c6411855a54376f.
//
// Solidity: event NodeReaped(address indexed operator, uint64 lastActiveBlock)
func (_DKGRegistry *DKGRegistryFilterer) WatchNodeReaped(opts *bind.WatchOpts, sink chan<- *DKGRegistryNodeReaped, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.WatchLogs(opts, "NodeReaped", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRegistryNodeReaped)
				if err := _DKGRegistry.contract.UnpackLog(event, "NodeReaped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeReaped is a log parse operation binding the contract event 0x17b35aacc7270dcc7c9993688488c4a6267c1ca2e4ab73b83c6411855a54376f.
//
// Solidity: event NodeReaped(address indexed operator, uint64 lastActiveBlock)
func (_DKGRegistry *DKGRegistryFilterer) ParseNodeReaped(log types.Log) (*DKGRegistryNodeReaped, error) {
	event := new(DKGRegistryNodeReaped)
	if err := _DKGRegistry.contract.UnpackLog(event, "NodeReaped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRegistryNodeRegisteredIterator is returned from FilterNodeRegistered and is used to iterate over the raw logs and unpacked data for NodeRegistered events raised by the DKGRegistry contract.
type DKGRegistryNodeRegisteredIterator struct {
	Event *DKGRegistryNodeRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DKGRegistryNodeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRegistryNodeRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DKGRegistryNodeRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DKGRegistryNodeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRegistryNodeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRegistryNodeRegistered represents a NodeRegistered event raised by the DKGRegistry contract.
type DKGRegistryNodeRegistered struct {
	Operator common.Address
	PubX     *big.Int
	PubY     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNodeRegistered is a free log retrieval operation binding the contract event 0x99140a41575033d78b1016979e49f1b8a4943ef274d75edf0dac1bc3ccbce5f7.
//
// Solidity: event NodeRegistered(address indexed operator, uint256 pubX, uint256 pubY)
func (_DKGRegistry *DKGRegistryFilterer) FilterNodeRegistered(opts *bind.FilterOpts, operator []common.Address) (*DKGRegistryNodeRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.FilterLogs(opts, "NodeRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DKGRegistryNodeRegisteredIterator{contract: _DKGRegistry.contract, event: "NodeRegistered", logs: logs, sub: sub}, nil
}

// WatchNodeRegistered is a free log subscription operation binding the contract event 0x99140a41575033d78b1016979e49f1b8a4943ef274d75edf0dac1bc3ccbce5f7.
//
// Solidity: event NodeRegistered(address indexed operator, uint256 pubX, uint256 pubY)
func (_DKGRegistry *DKGRegistryFilterer) WatchNodeRegistered(opts *bind.WatchOpts, sink chan<- *DKGRegistryNodeRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.WatchLogs(opts, "NodeRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRegistryNodeRegistered)
				if err := _DKGRegistry.contract.UnpackLog(event, "NodeRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeRegistered is a log parse operation binding the contract event 0x99140a41575033d78b1016979e49f1b8a4943ef274d75edf0dac1bc3ccbce5f7.
//
// Solidity: event NodeRegistered(address indexed operator, uint256 pubX, uint256 pubY)
func (_DKGRegistry *DKGRegistryFilterer) ParseNodeRegistered(log types.Log) (*DKGRegistryNodeRegistered, error) {
	event := new(DKGRegistryNodeRegistered)
	if err := _DKGRegistry.contract.UnpackLog(event, "NodeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRegistryNodeUpdatedIterator is returned from FilterNodeUpdated and is used to iterate over the raw logs and unpacked data for NodeUpdated events raised by the DKGRegistry contract.
type DKGRegistryNodeUpdatedIterator struct {
	Event *DKGRegistryNodeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DKGRegistryNodeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRegistryNodeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DKGRegistryNodeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DKGRegistryNodeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRegistryNodeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRegistryNodeUpdated represents a NodeUpdated event raised by the DKGRegistry contract.
type DKGRegistryNodeUpdated struct {
	Operator common.Address
	PubX     *big.Int
	PubY     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNodeUpdated is a free log retrieval operation binding the contract event 0x1e2215a8512058e371c99f86c2731c45755267c9d5fb9eb3c911230fa9b55cfc.
//
// Solidity: event NodeUpdated(address indexed operator, uint256 pubX, uint256 pubY)
func (_DKGRegistry *DKGRegistryFilterer) FilterNodeUpdated(opts *bind.FilterOpts, operator []common.Address) (*DKGRegistryNodeUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.FilterLogs(opts, "NodeUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DKGRegistryNodeUpdatedIterator{contract: _DKGRegistry.contract, event: "NodeUpdated", logs: logs, sub: sub}, nil
}

// WatchNodeUpdated is a free log subscription operation binding the contract event 0x1e2215a8512058e371c99f86c2731c45755267c9d5fb9eb3c911230fa9b55cfc.
//
// Solidity: event NodeUpdated(address indexed operator, uint256 pubX, uint256 pubY)
func (_DKGRegistry *DKGRegistryFilterer) WatchNodeUpdated(opts *bind.WatchOpts, sink chan<- *DKGRegistryNodeUpdated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.WatchLogs(opts, "NodeUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRegistryNodeUpdated)
				if err := _DKGRegistry.contract.UnpackLog(event, "NodeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeUpdated is a log parse operation binding the contract event 0x1e2215a8512058e371c99f86c2731c45755267c9d5fb9eb3c911230fa9b55cfc.
//
// Solidity: event NodeUpdated(address indexed operator, uint256 pubX, uint256 pubY)
func (_DKGRegistry *DKGRegistryFilterer) ParseNodeUpdated(log types.Log) (*DKGRegistryNodeUpdated, error) {
	event := new(DKGRegistryNodeUpdated)
	if err := _DKGRegistry.contract.UnpackLog(event, "NodeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
