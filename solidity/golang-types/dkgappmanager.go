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

// DKGTypesAppPolicy is an auto generated low-level Go binding around an user-defined struct.
type DKGTypesAppPolicy struct {
	AuthorizedSubmitter common.Address
	MaxCiphertexts      uint16
	NotBeforeBlock      uint64
	NotAfterBlock       uint64
}

// DKGTypesApplication is an auto generated low-level Go binding around an user-defined struct.
type DKGTypesApplication struct {
	Creator        common.Address
	Mode           uint8
	DerivationS    *big.Int
	OrganizerPK    DKGTypesPoint
	Policy         DKGTypesAppPolicy
	CreatedAtBlock uint64
	Exists         bool
}

// DKGAppManagerMetaData contains all meta data concerning the DKGAppManager contract.
var DKGAppManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_partialDecryptVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MANAGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARTIAL_DECRYPT_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApplication\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"app\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Application\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mode\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.AppMode\"},{\"name\":\"derivationS\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"organizerPK\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.AppPolicy\",\"components\":[{\"name\":\"authorizedSubmitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxCiphertexts\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"createdAtBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCombineCorrection\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"mode\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"derivationS\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deltaOrgX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deltaOrgY\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredAids\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerApplication\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.AppPolicy\",\"components\":[{\"name\":\"authorizedSubmitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxCiphertexts\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerApplicationCoDec\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.AppPolicy\",\"components\":[{\"name\":\"authorizedSubmitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxCiphertexts\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"pkOrgX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pkOrgY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrAx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrAy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrZ\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requireCanSubmitCiphertext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitOrganizerShare\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deltaOrgX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deltaOrgY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dleqProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"dleqInput\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ApplicationRegistered\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"mode\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"derivationS\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"organizerPKx\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"organizerPKy\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrganizerShareSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"deltaOrgX\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"deltaOrgY\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyPartiallyDecrypted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ApplicationAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextNotSubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionLimitReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionNotYetAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPartialDecryptions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidApplication\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCiphertext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEpoch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSchnorrProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifier\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IsIdentity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotCanonical\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOnCurve\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]}]",
	Bin: "0x60c0346100d257601f61211138819003918201601f19168301916001600160401b038311848410176100d65780849260409485528339810103126100d257610052602061004b836100ea565b92016100ea565b906001600160a01b038116156100c3576001600160a01b038216156100b45760805260a05260405161201290816100ff82396080518181816101030152818161037c015281816106a80152611142015260a05181818161091d01526112ea0152f35b63baa3de5f60e01b5f5260045ffd5b63e6c4247b60e01b5f5260045ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036100d25756fe60806040526004361015610011575f80fd5b5f3560e01c806317476f00146100a45780631b2df8501461009f5780632fed25291461009a5780634ba849e71461009557806374a99aba14610090578063852507001461008b578063be5b346314610086578063bf192209146100815763ed78d71e1461007c575f80fd5b610985565b610908565b6108ba565b610674565b610637565b610596565b610489565b610367565b3461033157610160366003190112610331576100be610335565b6024356100ca3661034c565b60c4359060e435906101043591610124359161014435635f2cdc7560e11b608090815261030090607f196100ff8b6084610dff565b01817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561032c575f906102ff575b80516001600160a01b0316156102f0576060600391015161015a81610e1e565b61016381610e1e565b036102e15786156102d2576101888761017b8a610e28565b905f5260205260405f2090565b93600685019561019d875460ff9060401c1690565b6102c3578782826101db956101d16101d7968f968f878d99918a926101c284846117eb565b6101cc86866117eb565b611927565b90611a2c565b1590565b6102b4575f516020611fbd5f395f51905f529361024661026c9260046102af966102053382610e76565b805460ff60a01b1916600160a01b1781555f6001820155610240610227610ad2565b8b8152602001889052600282018b905560038201889055565b01610eb4565b610259436001600160401b031682610e95565b805460ff60401b1916600160401b179055565b61027e8561027988610e42565b610f4e565b60408051600181525f602082015290810194909452606084015233946001600160a01b031916929081906080820190565b0390a4005b6327f7eb4d60e11b5f5260045ffd5b630b792c8f60e01b5f5260045ffd5b6378e9323b60e11b5f5260045ffd5b63268dbf6760e21b5f5260045ffd5b63d5b25b6360e01b5f5260045ffd5b506103003d8111610325575b8061031861032092610a18565b608001610c3b565b61013a565b503d61030b565b610e13565b5f80fd5b600435906001600160a01b03198216820361033157565b608090604319011261033157604490565b5f91031261033157565b34610331575f366003190112610331576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b634e487b7160e01b5f52602160045260245ffd5b600211156103c957565b6103ab565b81516001600160a01b0316815260208201516101608201939290919060028310156103c9576020828101939093526040808201518184015260608083015180518286015285015160808086019190915283015180516001600160a01b031660a08601529485015161ffff1660c085810191909152918501516001600160401b0390811660e0860152940151909316610100830152610487926101409160a08101516001600160401b031661012085015201511515910152565b565b346103315760403660031901126103315761053f61052261051d6104ab610335565b61017b602435915f60c06040516104c181610a43565b8281528260208201528260408201526040516104dc81610a5e565b83815283602082015260608201526040516104f681610a79565b83815283602082015283604082015283606082015260808201528260a08201520152610e28565b610fe4565b6060810160208151015115610543575b50604051918291826103ce565b0390f35b61054b610ad2565b905f825260016020830152525f610532565b61ffff81160361033157565b9181601f84011215610331578235916001600160401b038311610331576020838186019501011161033157565b3461033157610160366003190112610331576105b0610335565b602435604435916105c08361055d565b6084356064356101043560e43560c43560a435610124356001600160401b038111610331576105f3903690600401610569565b610144359a90979196906001600160401b038c116103315761061c6106249c3690600401610569565b9b909a61112b565b005b6001600160a01b0381160361033157565b3461033157608036600319011261033157610624610653610335565b6044356024356106628261055d565b6064359261066f84610626565b6115bf565b346103315760c03660031901126103315761068d610335565b6024356106993661034c565b604051635f2cdc7560e11b81527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169061030081806106e38860048301610dff565b0381855afa90811561032c575f9161088b575b5080516001600160a01b0316156102f0576060600391015161071781610e1e565b61072081610e1e565b036102e15782156102d2576107388361017b86610e28565b91600683019061074d825460ff9060401c1690565b6102c357604080516319a9f63760e11b815293849081806107718b60048301610dff565b03915afa92831561032c575f9361085a575b5060208301908151156102f057925190516040516001600160a01b0319881660208201908152602c820193909352604c810191909152606c8082018790528152610809936102469290916004916107e7916107df608c82610aaf565b519020611730565b956107f23382610e76565b805460ff60a01b1916815586600182015501610eb4565b6108168261027985610e42565b604080515f8082526020820193909352908101919091526001606082015233926001600160a01b031916905f516020611fbd5f395f51905f529080608081016102af565b61087d91935060403d604011610884575b6108758183610aaf565b810190611708565b915f610783565b503d61086b565b6108ad91506103003d81116108b3575b6108a58183610aaf565b810190610d17565b5f6106f6565b503d61089b565b346103315760603660031901126103315760806108ed6108d8610335565b602435604435916108e88361055d565b611743565b9160ff60405194168452602084015260408301526060820152f35b34610331575f366003190112610331576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b60206040818301928281528451809452019201905f5b81811061096f5750505090565b8251845260209384019390920191600101610962565b34610331576020366003190112610331576001600160a01b03196109a7610335565b165f52600160205260405f206040519081602082549182815201915f5260205f20905f5b8181106109ee5761053f856109e281870382610aaf565b6040519182918261094c565b82548452602090930192600192830192016109cb565b634e487b7160e01b5f52604160045260245ffd5b6080601f91909101601f19168101906001600160401b03821190821017610a3e57604052565b610a04565b60e081019081106001600160401b03821117610a3e57604052565b604081019081106001600160401b03821117610a3e57604052565b608081019081106001600160401b03821117610a3e57604052565b60c081019081106001600160401b03821117610a3e57604052565b601f909101601f19168101906001600160401b03821190821017610a3e57604052565b60405190610487604083610aaf565b604051906104876101a083610aaf565b519061048782610626565b51906104878261055d565b6001600160401b0381160361033157565b519061048782610b07565b91908260e091031261033157604051610b3b81610a43565b60c0610bad8183958051610b4e8161055d565b85526020810151610b5e8161055d565b6020860152610b6f60408201610afc565b6040860152610b8060608201610afc565b6060860152610b9160808201610b18565b6080860152610ba260a08201610b18565b60a086015201610b18565b910152565b91908260c091031261033157604051610bca81610a94565b809280519081151582036103315760a0610bad91819385526020810151610bf08161055d565b6020860152610c0160408201610b18565b6040860152610c1260608201610b18565b6060860152610c2360808201610b18565b608086015201610b18565b5190600682101561033157565b610300607f1982011261033157610c78610c53610ae1565b91610c5e6080610af1565b8352610c6b8160a0610b23565b6020840152610180610bb2565b6040820152610c88610240610c2e565b6060820152610c98610260610b18565b6080820152610ca8610280610b18565b60a0820152610cb86102a0610b18565b60c08201526102c05160e08201526102e051610100820152610cdb610300610afc565b610120820152610cec610320610afc565b610140820152610cfd610340610afc565b610160820152610d0e610360610afc565b61018082015290565b61030081830312610331576102e0610d0e91610d5c610d34610ae1565b94610d3e83610af1565b8652610d4d8160208501610b23565b60208701526101008301610bb2565b6040850152610d6e6101c08201610c2e565b6060850152610d806101e08201610b18565b6080850152610d926102008201610b18565b60a0850152610da46102208201610b18565b60c085015261024081015160e0850152610260810151610100850152610dcd6102808201610afc565b610120850152610de06102a08201610afc565b610140850152610df36102c08201610afc565b61016085015201610afc565b6001600160a01b0319909116815260200190565b6040513d5f823e3d90fd5b600611156103c957565b6001600160a01b0319165f90815260208190526040902090565b6001600160a01b0319165f90815260016020526040902090565b6001600160a01b0319165f90815260026020526040902090565b80546001600160a01b0319166001600160a01b03909216919091179055565b80546001600160401b0319166001600160401b03909216919091179055565b6001606061048793610ed08135610eca81610626565b85610e76565b6020810135610ede8161055d565b845461ffff60a01b19811660a09290921b61ffff60a01b169182178655906040830135610f0a81610b07565b8560b01b8660f01b039060b01b16918560a01b8660f01b03191617178455013591610f3483610b07565b01610e95565b634e487b7160e01b5f52603260045260245ffd5b805490600160401b821015610a3e5760018201808255821015610f75575f5260205f200155565b610f3a565b90604051610f8781610a5e565b602060018294805484520154910152565b90604051610fa581610a79565b82546001600160a01b038116825260a081901c61ffff16602083015260b01c6001600160401b0390811660408301526001909301549092166060830152565b9060405191610ff283610a43565b80546001600160a01b0381168452839060a01c60ff1660028110156103c95761106860066104879460c09360208601526001810154604086015261103860028201610f7a565b606086015261104960048201610f98565b608086015201546001600160401b03811660a085015260401c60ff1690565b1515910152565b90816020910312610331575190565b908060209392818452848401375f828201840152601f01601f1916010190565b92906110b7906110c5959360408652604086019161107e565b92602081850391015261107e565b90565b90610200828203126103315780601f8301121561033157610200604051926110f08285610aaf565b8391810192831161033157905b82821061110a5750505090565b81358152602091820191016110fd565b906010811015610f755760051b0190565b96999b949b9a909a98919895929560018060a01b037f000000000000000000000000000000000000000000000000000000000000000016604051635f2cdc7560e11b815261030081806111818d60048301610dff565b0381855afa90811561032c575f916115a0575b5080516001600160a01b0316156102f057606060039101516111b581610e1e565b6111be81610e1e565b036102e15761ffff8b169d8e158f8115611594575b50611585576111e58e61017b8c610e28565b956111fb6101d7600689015460ff9060401c1690565b6102d257600160ff611212895460ff9060a01c1690565b61121b816103bf565b16036102d25760208f611268948f918e60405197889485938493632de546d560e01b85526004850191604091949361ffff91606085019660018060a01b0319168552602085015216910152565b03915afa92831561032c575f93611554575b508215611545576112a591888b91608093916040519384526020840152604083015260608201522090565b0361146a576112da60026112d28c8f6112c19061017b8e610e5c565b9061ffff165f5260205260405f2090565b015460ff1690565b611536576112e8898c6117eb565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031692833b156103315786935f9284926113406040519a8b9586948594635c73957b60e11b86526004860161109e565b03915afa91821561032c57889561135f9361151c575b508101906110c8565b91825161137b61136f8860a01c90565b6001600160601b031690565b149485159561150d575b85156114fe575b85156114ee575b85156114e0575b85156114d1575b5084156114c2575b5083156114af575b8315611499575b508215611489575b8215611479575b505061146a5761144d7f8b6045276e66f28a1293f2044b947b82818f03c318251187680b22778c8f0948936114266113fd610ad2565b9187835285602084015261140f610ad2565b928352600160208401526112c18961017b87610e5c565b6002602091828451805183550151600182015501910151151560ff80198354169116179055565b6040805194855260208501929092526001600160a01b03191692a4565b63d1fed5fd60e01b5f5260045ffd5b610140015114159050825f6113c7565b61012081015187141592506113c0565b610100820151600390910154141592505f6113b8565b60e08201516002820154141593506113b1565b60c0830151141593505f6113a9565b60a0840151141594505f6113a1565b60808401511515955061139a565b6060840151600214159550611393565b60408401518c1415955061138c565b60208401518b14159550611385565b8061152a5f61153093610aaf565b8061035d565b5f611356565b633466526160e01b5f5260045ffd5b6346f551f560e01b5f5260045ffd5b61157791935060203d60201161157e575b61156f8183610aaf565b81019061106f565b915f61127a565b503d611565565b634c4d29cd60e11b5f5260045ffd5b6101009150118f6111d3565b6115b991506103003d81116108b3576108a58183610aaf565b5f611194565b9092919280156117025761017b6115d592610e28565b60068101546115e89060401c60ff161590565b6102d25760046115f89101610f98565b805190916001600160a01b039091168015159190826116ea575b50506116db5760408101516001600160401b031680151590816116bf575b506116b05760608101516001600160401b0316801515908161169d575b5061168e576020015161ffff1680151591908261167c575b505061166d57565b63464e67af60e01b5f5260045ffd5b61ffff90811691161190505f80611665565b630410ff2960e31b5f5260045ffd5b436001600160401b03161190505f61164d565b633deac39560e01b5f5260045ffd5b6001600160401b03169050436001600160401b0316105f611630565b6330cd747160e01b5f5260045ffd5b6001600160a01b039182169116141590505f80611612565b50505050565b908160409103126103315760206040519161172283610a5e565b805183520151602082015290565b5f516020611f9d5f395f51905f52900690565b929181156117dd576117588261017b86610e28565b9161176e6101d7600685015460ff9060401c1690565b6102d257825460a01c60ff1694611784866103bf565b60ff861661179b5750505060010154905f90600190565b6117ad93509061017b6112c192610e5c565b6117be6101d7600283015460ff1690565b6117ce575f916001825492015490565b63032cddf960e11b5f5260045ffd5b505f92508291829150600190565b905f516020611f7d5f395f51905f5282108061185a575b1561184b57811580611841575b6118325761181c91611870565b1561182357565b6361586bdd60e01b5f5260045ffd5b632b39517d60e21b5f5260045ffd5b506001811461180f565b63d7c7beeb60e01b5f5260045ffd5b505f516020611f7d5f395f51905f528110611802565b5f516020611f7d5f395f51905f528110801590611910575b61190a575f516020611f7d5f395f51905f52818192099180095f516020611f7d5f395f51905f5282065f516020611f7d5f395f51905f5203905f516020611f7d5f395f51905f528211611905575f516020611f7d5f395f51905f528080838180965f0896095f516020611f5d5f395f51905f520960010892081490565b6119a3565b50505f90565b505f516020611f7d5f395f51905f52821015611888565b9390925f516020611f9d5f395f51905f5295926040519460208601967f41ea6f3fa95eccd1f3b1ce8e05efa11027280aa0c6b4167fd6695db659c30b28885260018060a01b0319166040870152604c860152606c850152608c84015260ac83015260cc82015260cc815261199c60ec82610aaf565b5190200690565b634e487b7160e01b5f52601160045260245ffd5b5f516020611f7d5f395f51905f5203905f516020611f7d5f395f51905f52821161190557565b604051906102006119ee8184610aaf565b368337565b908160021b918083046004149015171561190557565b908160011b918083046002149015171561190557565b9190820180921161190557565b611a3e611a4491969293949596611730565b92611730565b9480611cdc57505f5b611b7a611a586119dd565b92611a616119dd565b60018152927f1561ff836ce19d358a4eb7a4c199e94c377c749ae6f2a277f1f9195afe553f9f60808601527f25797203f7a0b24925572e1cd16bf9edfce0051fb9e133774b3c257a872d7d8b60808501527f0daaa7e6b25c28e6dc8dd1d48e9cc61cd07015c1d7c1b8d4590eb6f51d5346dc6101008601527f01666cafbf0a30da8b9ebeaf848a1da067a892296f1043188e1705402b6d68536101008501527f136d609c4c856f5d277fab08c730cbdd1a776ce4728c6a2eb20ff22bccf268946101808601527f21d66f0e2295ae954494f25889f9319cc1b4df71eff3f46ba9e4631b43fd7c95610180850152806020860152816020850152611b6682828181611d79565b604086810191825287018290525190611d79565b6060830152606083015260015b60048110611c4457505f92600196607e805b611bb4575050505050149182611bae57505090565b14919050565b5f1901611bc081611a09565b600383600c86841c60021b16921c16178615801590611c39575b611c18575b80611bec575b5080611b99565b9586829a611c09611c01611c10959a8a61111a565b51928861111a565b5192611d79565b989095611be5565b958981611c31939b611c2993611d79565b908181611d79565b989095611bdf565b5060018a1415611bda565b611c56611c50826119f3565b8461111a565b51611c63611c50836119f3565b519060015b60048110611c7b57505050600101611b87565b80611cd5611ca2611c8e6001948a61111a565b51611c99848a61111a565b51908787611d79565b611cb784611cb28a9594956119f3565b611a1f565b90611cce611cc886611cb28c6119f3565b8b61111a565b528961111a565b5201611c68565b611ce5906119b7565b611a4d565b5f516020611f7d5f395f51905f5290065f516020611f7d5f395f51905f52035f516020611f7d5f395f51905f528111611905575f516020611f7d5f395f51905f529060010890565b905f516020611f7d5f395f51905f5290065f516020611f7d5f395f51905f52035f516020611f7d5f395f51905f528111611905575f516020611f7d5f395f51905f52910890565b9193929093821580611eec575b611ee557801580611edb575b611ed5575f516020611f7d5f395f51905f52818409915f516020611f7d5f395f51905f528187095f516020611f7d5f395f51905f528185095f516020611f7d5f395f51905f52905f516020611f5d5f395f51905f5209935f516020611f7d5f395f51905f52907f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000009611e2391611d32565b935f516020611f7d5f395f51905f528460010893611e4090611cea565b965f516020611f7d5f395f51905f52888609611e5b90611ef6565b97885f516020611f7d5f395f51905f529109935f516020611f7d5f395f51905f529109915f516020611f7d5f395f51905f529109905f516020611f7d5f395f51905f529108905f516020611f7d5f395f51905f529109935f516020611f7d5f395f51905f5291095f516020611f7d5f395f51905f52910990565b50509190565b5060018214611d92565b9350919050565b5060018514611d86565b60405190602082526020808301526020604083015260608201527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593efffffff60808201525f516020611f7d5f395f51905f5260a082015260208160c08160055afa1561033157519056fe1aee90f15f2189693df072d799fd11fc039b2959ebb7c867d075ca8cf4d7eb8e30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001060c89ce5c263405370a08b6d0302b0bab3eedb83920ee0a677297dc392126f15c1bc55eb261d6ac466922a422fe62e9de8433120dc04979463fd16a85723792a2646970667358221220927682fbc21d004b79a7bb8de509289f7356fce3b84173d4e6ebccd72ad66c0364736f6c634300081c0033",
}

// DKGAppManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DKGAppManagerMetaData.ABI instead.
var DKGAppManagerABI = DKGAppManagerMetaData.ABI

// DKGAppManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DKGAppManagerMetaData.Bin instead.
var DKGAppManagerBin = DKGAppManagerMetaData.Bin

// DeployDKGAppManager deploys a new Ethereum contract, binding an instance of DKGAppManager to it.
func DeployDKGAppManager(auth *bind.TransactOpts, backend bind.ContractBackend, _manager common.Address, _partialDecryptVerifier common.Address) (common.Address, *types.Transaction, *DKGAppManager, error) {
	parsed, err := DKGAppManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DKGAppManagerBin), backend, _manager, _partialDecryptVerifier)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DKGAppManager{DKGAppManagerCaller: DKGAppManagerCaller{contract: contract}, DKGAppManagerTransactor: DKGAppManagerTransactor{contract: contract}, DKGAppManagerFilterer: DKGAppManagerFilterer{contract: contract}}, nil
}

// DKGAppManager is an auto generated Go binding around an Ethereum contract.
type DKGAppManager struct {
	DKGAppManagerCaller     // Read-only binding to the contract
	DKGAppManagerTransactor // Write-only binding to the contract
	DKGAppManagerFilterer   // Log filterer for contract events
}

// DKGAppManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DKGAppManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGAppManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DKGAppManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGAppManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DKGAppManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGAppManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DKGAppManagerSession struct {
	Contract     *DKGAppManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DKGAppManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DKGAppManagerCallerSession struct {
	Contract *DKGAppManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DKGAppManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DKGAppManagerTransactorSession struct {
	Contract     *DKGAppManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DKGAppManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DKGAppManagerRaw struct {
	Contract *DKGAppManager // Generic contract binding to access the raw methods on
}

// DKGAppManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DKGAppManagerCallerRaw struct {
	Contract *DKGAppManagerCaller // Generic read-only contract binding to access the raw methods on
}

// DKGAppManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DKGAppManagerTransactorRaw struct {
	Contract *DKGAppManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDKGAppManager creates a new instance of DKGAppManager, bound to a specific deployed contract.
func NewDKGAppManager(address common.Address, backend bind.ContractBackend) (*DKGAppManager, error) {
	contract, err := bindDKGAppManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DKGAppManager{DKGAppManagerCaller: DKGAppManagerCaller{contract: contract}, DKGAppManagerTransactor: DKGAppManagerTransactor{contract: contract}, DKGAppManagerFilterer: DKGAppManagerFilterer{contract: contract}}, nil
}

// NewDKGAppManagerCaller creates a new read-only instance of DKGAppManager, bound to a specific deployed contract.
func NewDKGAppManagerCaller(address common.Address, caller bind.ContractCaller) (*DKGAppManagerCaller, error) {
	contract, err := bindDKGAppManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DKGAppManagerCaller{contract: contract}, nil
}

// NewDKGAppManagerTransactor creates a new write-only instance of DKGAppManager, bound to a specific deployed contract.
func NewDKGAppManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*DKGAppManagerTransactor, error) {
	contract, err := bindDKGAppManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DKGAppManagerTransactor{contract: contract}, nil
}

// NewDKGAppManagerFilterer creates a new log filterer instance of DKGAppManager, bound to a specific deployed contract.
func NewDKGAppManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*DKGAppManagerFilterer, error) {
	contract, err := bindDKGAppManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DKGAppManagerFilterer{contract: contract}, nil
}

// bindDKGAppManager binds a generic wrapper to an already deployed contract.
func bindDKGAppManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DKGAppManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DKGAppManager *DKGAppManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKGAppManager.Contract.DKGAppManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DKGAppManager *DKGAppManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKGAppManager.Contract.DKGAppManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DKGAppManager *DKGAppManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKGAppManager.Contract.DKGAppManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DKGAppManager *DKGAppManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKGAppManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DKGAppManager *DKGAppManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKGAppManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DKGAppManager *DKGAppManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKGAppManager.Contract.contract.Transact(opts, method, params...)
}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(address)
func (_DKGAppManager *DKGAppManagerCaller) MANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKGAppManager.contract.Call(opts, &out, "MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(address)
func (_DKGAppManager *DKGAppManagerSession) MANAGER() (common.Address, error) {
	return _DKGAppManager.Contract.MANAGER(&_DKGAppManager.CallOpts)
}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(address)
func (_DKGAppManager *DKGAppManagerCallerSession) MANAGER() (common.Address, error) {
	return _DKGAppManager.Contract.MANAGER(&_DKGAppManager.CallOpts)
}

// PARTIALDECRYPTVERIFIER is a free data retrieval call binding the contract method 0xbf192209.
//
// Solidity: function PARTIAL_DECRYPT_VERIFIER() view returns(address)
func (_DKGAppManager *DKGAppManagerCaller) PARTIALDECRYPTVERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKGAppManager.contract.Call(opts, &out, "PARTIAL_DECRYPT_VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PARTIALDECRYPTVERIFIER is a free data retrieval call binding the contract method 0xbf192209.
//
// Solidity: function PARTIAL_DECRYPT_VERIFIER() view returns(address)
func (_DKGAppManager *DKGAppManagerSession) PARTIALDECRYPTVERIFIER() (common.Address, error) {
	return _DKGAppManager.Contract.PARTIALDECRYPTVERIFIER(&_DKGAppManager.CallOpts)
}

// PARTIALDECRYPTVERIFIER is a free data retrieval call binding the contract method 0xbf192209.
//
// Solidity: function PARTIAL_DECRYPT_VERIFIER() view returns(address)
func (_DKGAppManager *DKGAppManagerCallerSession) PARTIALDECRYPTVERIFIER() (common.Address, error) {
	return _DKGAppManager.Contract.PARTIALDECRYPTVERIFIER(&_DKGAppManager.CallOpts)
}

// GetApplication is a free data retrieval call binding the contract method 0x2fed2529.
//
// Solidity: function getApplication(bytes12 epochId, bytes32 aid) view returns((address,uint8,uint256,(uint256,uint256),(address,uint16,uint64,uint64),uint64,bool) app)
func (_DKGAppManager *DKGAppManagerCaller) GetApplication(opts *bind.CallOpts, epochId [12]byte, aid [32]byte) (DKGTypesApplication, error) {
	var out []interface{}
	err := _DKGAppManager.contract.Call(opts, &out, "getApplication", epochId, aid)

	if err != nil {
		return *new(DKGTypesApplication), err
	}

	out0 := *abi.ConvertType(out[0], new(DKGTypesApplication)).(*DKGTypesApplication)

	return out0, err

}

// GetApplication is a free data retrieval call binding the contract method 0x2fed2529.
//
// Solidity: function getApplication(bytes12 epochId, bytes32 aid) view returns((address,uint8,uint256,(uint256,uint256),(address,uint16,uint64,uint64),uint64,bool) app)
func (_DKGAppManager *DKGAppManagerSession) GetApplication(epochId [12]byte, aid [32]byte) (DKGTypesApplication, error) {
	return _DKGAppManager.Contract.GetApplication(&_DKGAppManager.CallOpts, epochId, aid)
}

// GetApplication is a free data retrieval call binding the contract method 0x2fed2529.
//
// Solidity: function getApplication(bytes12 epochId, bytes32 aid) view returns((address,uint8,uint256,(uint256,uint256),(address,uint16,uint64,uint64),uint64,bool) app)
func (_DKGAppManager *DKGAppManagerCallerSession) GetApplication(epochId [12]byte, aid [32]byte) (DKGTypesApplication, error) {
	return _DKGAppManager.Contract.GetApplication(&_DKGAppManager.CallOpts, epochId, aid)
}

// GetCombineCorrection is a free data retrieval call binding the contract method 0xbe5b3463.
//
// Solidity: function getCombineCorrection(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) view returns(uint8 mode, uint256 derivationS, uint256 deltaOrgX, uint256 deltaOrgY)
func (_DKGAppManager *DKGAppManagerCaller) GetCombineCorrection(opts *bind.CallOpts, epochId [12]byte, aid [32]byte, ciphertextIndex uint16) (struct {
	Mode        uint8
	DerivationS *big.Int
	DeltaOrgX   *big.Int
	DeltaOrgY   *big.Int
}, error) {
	var out []interface{}
	err := _DKGAppManager.contract.Call(opts, &out, "getCombineCorrection", epochId, aid, ciphertextIndex)

	outstruct := new(struct {
		Mode        uint8
		DerivationS *big.Int
		DeltaOrgX   *big.Int
		DeltaOrgY   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mode = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.DerivationS = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DeltaOrgX = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DeltaOrgY = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCombineCorrection is a free data retrieval call binding the contract method 0xbe5b3463.
//
// Solidity: function getCombineCorrection(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) view returns(uint8 mode, uint256 derivationS, uint256 deltaOrgX, uint256 deltaOrgY)
func (_DKGAppManager *DKGAppManagerSession) GetCombineCorrection(epochId [12]byte, aid [32]byte, ciphertextIndex uint16) (struct {
	Mode        uint8
	DerivationS *big.Int
	DeltaOrgX   *big.Int
	DeltaOrgY   *big.Int
}, error) {
	return _DKGAppManager.Contract.GetCombineCorrection(&_DKGAppManager.CallOpts, epochId, aid, ciphertextIndex)
}

// GetCombineCorrection is a free data retrieval call binding the contract method 0xbe5b3463.
//
// Solidity: function getCombineCorrection(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) view returns(uint8 mode, uint256 derivationS, uint256 deltaOrgX, uint256 deltaOrgY)
func (_DKGAppManager *DKGAppManagerCallerSession) GetCombineCorrection(epochId [12]byte, aid [32]byte, ciphertextIndex uint16) (struct {
	Mode        uint8
	DerivationS *big.Int
	DeltaOrgX   *big.Int
	DeltaOrgY   *big.Int
}, error) {
	return _DKGAppManager.Contract.GetCombineCorrection(&_DKGAppManager.CallOpts, epochId, aid, ciphertextIndex)
}

// GetRegisteredAids is a free data retrieval call binding the contract method 0xed78d71e.
//
// Solidity: function getRegisteredAids(bytes12 epochId) view returns(bytes32[])
func (_DKGAppManager *DKGAppManagerCaller) GetRegisteredAids(opts *bind.CallOpts, epochId [12]byte) ([][32]byte, error) {
	var out []interface{}
	err := _DKGAppManager.contract.Call(opts, &out, "getRegisteredAids", epochId)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetRegisteredAids is a free data retrieval call binding the contract method 0xed78d71e.
//
// Solidity: function getRegisteredAids(bytes12 epochId) view returns(bytes32[])
func (_DKGAppManager *DKGAppManagerSession) GetRegisteredAids(epochId [12]byte) ([][32]byte, error) {
	return _DKGAppManager.Contract.GetRegisteredAids(&_DKGAppManager.CallOpts, epochId)
}

// GetRegisteredAids is a free data retrieval call binding the contract method 0xed78d71e.
//
// Solidity: function getRegisteredAids(bytes12 epochId) view returns(bytes32[])
func (_DKGAppManager *DKGAppManagerCallerSession) GetRegisteredAids(epochId [12]byte) ([][32]byte, error) {
	return _DKGAppManager.Contract.GetRegisteredAids(&_DKGAppManager.CallOpts, epochId)
}

// RequireCanSubmitCiphertext is a free data retrieval call binding the contract method 0x74a99aba.
//
// Solidity: function requireCanSubmitCiphertext(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, address sender) view returns()
func (_DKGAppManager *DKGAppManagerCaller) RequireCanSubmitCiphertext(opts *bind.CallOpts, epochId [12]byte, aid [32]byte, ciphertextIndex uint16, sender common.Address) error {
	var out []interface{}
	err := _DKGAppManager.contract.Call(opts, &out, "requireCanSubmitCiphertext", epochId, aid, ciphertextIndex, sender)

	if err != nil {
		return err
	}

	return err

}

// RequireCanSubmitCiphertext is a free data retrieval call binding the contract method 0x74a99aba.
//
// Solidity: function requireCanSubmitCiphertext(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, address sender) view returns()
func (_DKGAppManager *DKGAppManagerSession) RequireCanSubmitCiphertext(epochId [12]byte, aid [32]byte, ciphertextIndex uint16, sender common.Address) error {
	return _DKGAppManager.Contract.RequireCanSubmitCiphertext(&_DKGAppManager.CallOpts, epochId, aid, ciphertextIndex, sender)
}

// RequireCanSubmitCiphertext is a free data retrieval call binding the contract method 0x74a99aba.
//
// Solidity: function requireCanSubmitCiphertext(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, address sender) view returns()
func (_DKGAppManager *DKGAppManagerCallerSession) RequireCanSubmitCiphertext(epochId [12]byte, aid [32]byte, ciphertextIndex uint16, sender common.Address) error {
	return _DKGAppManager.Contract.RequireCanSubmitCiphertext(&_DKGAppManager.CallOpts, epochId, aid, ciphertextIndex, sender)
}

// RegisterApplication is a paid mutator transaction binding the contract method 0x85250700.
//
// Solidity: function registerApplication(bytes12 epochId, bytes32 aid, (address,uint16,uint64,uint64) policy) returns()
func (_DKGAppManager *DKGAppManagerTransactor) RegisterApplication(opts *bind.TransactOpts, epochId [12]byte, aid [32]byte, policy DKGTypesAppPolicy) (*types.Transaction, error) {
	return _DKGAppManager.contract.Transact(opts, "registerApplication", epochId, aid, policy)
}

// RegisterApplication is a paid mutator transaction binding the contract method 0x85250700.
//
// Solidity: function registerApplication(bytes12 epochId, bytes32 aid, (address,uint16,uint64,uint64) policy) returns()
func (_DKGAppManager *DKGAppManagerSession) RegisterApplication(epochId [12]byte, aid [32]byte, policy DKGTypesAppPolicy) (*types.Transaction, error) {
	return _DKGAppManager.Contract.RegisterApplication(&_DKGAppManager.TransactOpts, epochId, aid, policy)
}

// RegisterApplication is a paid mutator transaction binding the contract method 0x85250700.
//
// Solidity: function registerApplication(bytes12 epochId, bytes32 aid, (address,uint16,uint64,uint64) policy) returns()
func (_DKGAppManager *DKGAppManagerTransactorSession) RegisterApplication(epochId [12]byte, aid [32]byte, policy DKGTypesAppPolicy) (*types.Transaction, error) {
	return _DKGAppManager.Contract.RegisterApplication(&_DKGAppManager.TransactOpts, epochId, aid, policy)
}

// RegisterApplicationCoDec is a paid mutator transaction binding the contract method 0x17476f00.
//
// Solidity: function registerApplicationCoDec(bytes12 epochId, bytes32 aid, (address,uint16,uint64,uint64) policy, uint256 pkOrgX, uint256 pkOrgY, uint256 schnorrAx, uint256 schnorrAy, uint256 schnorrZ) returns()
func (_DKGAppManager *DKGAppManagerTransactor) RegisterApplicationCoDec(opts *bind.TransactOpts, epochId [12]byte, aid [32]byte, policy DKGTypesAppPolicy, pkOrgX *big.Int, pkOrgY *big.Int, schnorrAx *big.Int, schnorrAy *big.Int, schnorrZ *big.Int) (*types.Transaction, error) {
	return _DKGAppManager.contract.Transact(opts, "registerApplicationCoDec", epochId, aid, policy, pkOrgX, pkOrgY, schnorrAx, schnorrAy, schnorrZ)
}

// RegisterApplicationCoDec is a paid mutator transaction binding the contract method 0x17476f00.
//
// Solidity: function registerApplicationCoDec(bytes12 epochId, bytes32 aid, (address,uint16,uint64,uint64) policy, uint256 pkOrgX, uint256 pkOrgY, uint256 schnorrAx, uint256 schnorrAy, uint256 schnorrZ) returns()
func (_DKGAppManager *DKGAppManagerSession) RegisterApplicationCoDec(epochId [12]byte, aid [32]byte, policy DKGTypesAppPolicy, pkOrgX *big.Int, pkOrgY *big.Int, schnorrAx *big.Int, schnorrAy *big.Int, schnorrZ *big.Int) (*types.Transaction, error) {
	return _DKGAppManager.Contract.RegisterApplicationCoDec(&_DKGAppManager.TransactOpts, epochId, aid, policy, pkOrgX, pkOrgY, schnorrAx, schnorrAy, schnorrZ)
}

// RegisterApplicationCoDec is a paid mutator transaction binding the contract method 0x17476f00.
//
// Solidity: function registerApplicationCoDec(bytes12 epochId, bytes32 aid, (address,uint16,uint64,uint64) policy, uint256 pkOrgX, uint256 pkOrgY, uint256 schnorrAx, uint256 schnorrAy, uint256 schnorrZ) returns()
func (_DKGAppManager *DKGAppManagerTransactorSession) RegisterApplicationCoDec(epochId [12]byte, aid [32]byte, policy DKGTypesAppPolicy, pkOrgX *big.Int, pkOrgY *big.Int, schnorrAx *big.Int, schnorrAy *big.Int, schnorrZ *big.Int) (*types.Transaction, error) {
	return _DKGAppManager.Contract.RegisterApplicationCoDec(&_DKGAppManager.TransactOpts, epochId, aid, policy, pkOrgX, pkOrgY, schnorrAx, schnorrAy, schnorrZ)
}

// SubmitOrganizerShare is a paid mutator transaction binding the contract method 0x4ba849e7.
//
// Solidity: function submitOrganizerShare(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, uint256 deltaOrgX, uint256 deltaOrgY, bytes dleqProof, bytes dleqInput) returns()
func (_DKGAppManager *DKGAppManagerTransactor) SubmitOrganizerShare(opts *bind.TransactOpts, epochId [12]byte, aid [32]byte, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int, deltaOrgX *big.Int, deltaOrgY *big.Int, dleqProof []byte, dleqInput []byte) (*types.Transaction, error) {
	return _DKGAppManager.contract.Transact(opts, "submitOrganizerShare", epochId, aid, ciphertextIndex, c1x, c1y, c2x, c2y, deltaOrgX, deltaOrgY, dleqProof, dleqInput)
}

// SubmitOrganizerShare is a paid mutator transaction binding the contract method 0x4ba849e7.
//
// Solidity: function submitOrganizerShare(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, uint256 deltaOrgX, uint256 deltaOrgY, bytes dleqProof, bytes dleqInput) returns()
func (_DKGAppManager *DKGAppManagerSession) SubmitOrganizerShare(epochId [12]byte, aid [32]byte, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int, deltaOrgX *big.Int, deltaOrgY *big.Int, dleqProof []byte, dleqInput []byte) (*types.Transaction, error) {
	return _DKGAppManager.Contract.SubmitOrganizerShare(&_DKGAppManager.TransactOpts, epochId, aid, ciphertextIndex, c1x, c1y, c2x, c2y, deltaOrgX, deltaOrgY, dleqProof, dleqInput)
}

// SubmitOrganizerShare is a paid mutator transaction binding the contract method 0x4ba849e7.
//
// Solidity: function submitOrganizerShare(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, uint256 deltaOrgX, uint256 deltaOrgY, bytes dleqProof, bytes dleqInput) returns()
func (_DKGAppManager *DKGAppManagerTransactorSession) SubmitOrganizerShare(epochId [12]byte, aid [32]byte, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int, deltaOrgX *big.Int, deltaOrgY *big.Int, dleqProof []byte, dleqInput []byte) (*types.Transaction, error) {
	return _DKGAppManager.Contract.SubmitOrganizerShare(&_DKGAppManager.TransactOpts, epochId, aid, ciphertextIndex, c1x, c1y, c2x, c2y, deltaOrgX, deltaOrgY, dleqProof, dleqInput)
}

// DKGAppManagerApplicationRegisteredIterator is returned from FilterApplicationRegistered and is used to iterate over the raw logs and unpacked data for ApplicationRegistered events raised by the DKGAppManager contract.
type DKGAppManagerApplicationRegisteredIterator struct {
	Event *DKGAppManagerApplicationRegistered // Event containing the contract specifics and raw log

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
func (it *DKGAppManagerApplicationRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGAppManagerApplicationRegistered)
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
		it.Event = new(DKGAppManagerApplicationRegistered)
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
func (it *DKGAppManagerApplicationRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGAppManagerApplicationRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGAppManagerApplicationRegistered represents a ApplicationRegistered event raised by the DKGAppManager contract.
type DKGAppManagerApplicationRegistered struct {
	EpochId      [12]byte
	Aid          [32]byte
	Creator      common.Address
	Mode         uint8
	DerivationS  *big.Int
	OrganizerPKx *big.Int
	OrganizerPKy *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterApplicationRegistered is a free log retrieval operation binding the contract event 0x5c1bc55eb261d6ac466922a422fe62e9de8433120dc04979463fd16a85723792.
//
// Solidity: event ApplicationRegistered(bytes12 indexed epochId, bytes32 indexed aid, address indexed creator, uint8 mode, uint256 derivationS, uint256 organizerPKx, uint256 organizerPKy)
func (_DKGAppManager *DKGAppManagerFilterer) FilterApplicationRegistered(opts *bind.FilterOpts, epochId [][12]byte, aid [][32]byte, creator []common.Address) (*DKGAppManagerApplicationRegisteredIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var aidRule []interface{}
	for _, aidItem := range aid {
		aidRule = append(aidRule, aidItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _DKGAppManager.contract.FilterLogs(opts, "ApplicationRegistered", epochIdRule, aidRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &DKGAppManagerApplicationRegisteredIterator{contract: _DKGAppManager.contract, event: "ApplicationRegistered", logs: logs, sub: sub}, nil
}

// WatchApplicationRegistered is a free log subscription operation binding the contract event 0x5c1bc55eb261d6ac466922a422fe62e9de8433120dc04979463fd16a85723792.
//
// Solidity: event ApplicationRegistered(bytes12 indexed epochId, bytes32 indexed aid, address indexed creator, uint8 mode, uint256 derivationS, uint256 organizerPKx, uint256 organizerPKy)
func (_DKGAppManager *DKGAppManagerFilterer) WatchApplicationRegistered(opts *bind.WatchOpts, sink chan<- *DKGAppManagerApplicationRegistered, epochId [][12]byte, aid [][32]byte, creator []common.Address) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var aidRule []interface{}
	for _, aidItem := range aid {
		aidRule = append(aidRule, aidItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _DKGAppManager.contract.WatchLogs(opts, "ApplicationRegistered", epochIdRule, aidRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGAppManagerApplicationRegistered)
				if err := _DKGAppManager.contract.UnpackLog(event, "ApplicationRegistered", log); err != nil {
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

// ParseApplicationRegistered is a log parse operation binding the contract event 0x5c1bc55eb261d6ac466922a422fe62e9de8433120dc04979463fd16a85723792.
//
// Solidity: event ApplicationRegistered(bytes12 indexed epochId, bytes32 indexed aid, address indexed creator, uint8 mode, uint256 derivationS, uint256 organizerPKx, uint256 organizerPKy)
func (_DKGAppManager *DKGAppManagerFilterer) ParseApplicationRegistered(log types.Log) (*DKGAppManagerApplicationRegistered, error) {
	event := new(DKGAppManagerApplicationRegistered)
	if err := _DKGAppManager.contract.UnpackLog(event, "ApplicationRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGAppManagerOrganizerShareSubmittedIterator is returned from FilterOrganizerShareSubmitted and is used to iterate over the raw logs and unpacked data for OrganizerShareSubmitted events raised by the DKGAppManager contract.
type DKGAppManagerOrganizerShareSubmittedIterator struct {
	Event *DKGAppManagerOrganizerShareSubmitted // Event containing the contract specifics and raw log

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
func (it *DKGAppManagerOrganizerShareSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGAppManagerOrganizerShareSubmitted)
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
		it.Event = new(DKGAppManagerOrganizerShareSubmitted)
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
func (it *DKGAppManagerOrganizerShareSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGAppManagerOrganizerShareSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGAppManagerOrganizerShareSubmitted represents a OrganizerShareSubmitted event raised by the DKGAppManager contract.
type DKGAppManagerOrganizerShareSubmitted struct {
	EpochId         [12]byte
	Aid             [32]byte
	CiphertextIndex uint16
	DeltaOrgX       *big.Int
	DeltaOrgY       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOrganizerShareSubmitted is a free log retrieval operation binding the contract event 0x8b6045276e66f28a1293f2044b947b82818f03c318251187680b22778c8f0948.
//
// Solidity: event OrganizerShareSubmitted(bytes12 indexed epochId, bytes32 indexed aid, uint16 indexed ciphertextIndex, uint256 deltaOrgX, uint256 deltaOrgY)
func (_DKGAppManager *DKGAppManagerFilterer) FilterOrganizerShareSubmitted(opts *bind.FilterOpts, epochId [][12]byte, aid [][32]byte, ciphertextIndex []uint16) (*DKGAppManagerOrganizerShareSubmittedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var aidRule []interface{}
	for _, aidItem := range aid {
		aidRule = append(aidRule, aidItem)
	}
	var ciphertextIndexRule []interface{}
	for _, ciphertextIndexItem := range ciphertextIndex {
		ciphertextIndexRule = append(ciphertextIndexRule, ciphertextIndexItem)
	}

	logs, sub, err := _DKGAppManager.contract.FilterLogs(opts, "OrganizerShareSubmitted", epochIdRule, aidRule, ciphertextIndexRule)
	if err != nil {
		return nil, err
	}
	return &DKGAppManagerOrganizerShareSubmittedIterator{contract: _DKGAppManager.contract, event: "OrganizerShareSubmitted", logs: logs, sub: sub}, nil
}

// WatchOrganizerShareSubmitted is a free log subscription operation binding the contract event 0x8b6045276e66f28a1293f2044b947b82818f03c318251187680b22778c8f0948.
//
// Solidity: event OrganizerShareSubmitted(bytes12 indexed epochId, bytes32 indexed aid, uint16 indexed ciphertextIndex, uint256 deltaOrgX, uint256 deltaOrgY)
func (_DKGAppManager *DKGAppManagerFilterer) WatchOrganizerShareSubmitted(opts *bind.WatchOpts, sink chan<- *DKGAppManagerOrganizerShareSubmitted, epochId [][12]byte, aid [][32]byte, ciphertextIndex []uint16) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var aidRule []interface{}
	for _, aidItem := range aid {
		aidRule = append(aidRule, aidItem)
	}
	var ciphertextIndexRule []interface{}
	for _, ciphertextIndexItem := range ciphertextIndex {
		ciphertextIndexRule = append(ciphertextIndexRule, ciphertextIndexItem)
	}

	logs, sub, err := _DKGAppManager.contract.WatchLogs(opts, "OrganizerShareSubmitted", epochIdRule, aidRule, ciphertextIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGAppManagerOrganizerShareSubmitted)
				if err := _DKGAppManager.contract.UnpackLog(event, "OrganizerShareSubmitted", log); err != nil {
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

// ParseOrganizerShareSubmitted is a log parse operation binding the contract event 0x8b6045276e66f28a1293f2044b947b82818f03c318251187680b22778c8f0948.
//
// Solidity: event OrganizerShareSubmitted(bytes12 indexed epochId, bytes32 indexed aid, uint16 indexed ciphertextIndex, uint256 deltaOrgX, uint256 deltaOrgY)
func (_DKGAppManager *DKGAppManagerFilterer) ParseOrganizerShareSubmitted(log types.Log) (*DKGAppManagerOrganizerShareSubmitted, error) {
	event := new(DKGAppManagerOrganizerShareSubmitted)
	if err := _DKGAppManager.contract.UnpackLog(event, "OrganizerShareSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
