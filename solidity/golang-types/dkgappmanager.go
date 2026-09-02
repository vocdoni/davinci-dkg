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
	OrganizerPK    DKGTypesPoint
	Policy         DKGTypesAppPolicy
	CreatedAtBlock uint64
	Exists         bool
}

// DKGAppManagerMetaData contains all meta data concerning the DKGAppManager contract.
var DKGAppManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MANAGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApplication\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Application\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"organizerPK\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.AppPolicy\",\"components\":[{\"name\":\"authorizedSubmitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxCiphertexts\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"createdAtBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrganizerShareHash\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredAids\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerApplication\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.AppPolicy\",\"components\":[{\"name\":\"authorizedSubmitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxCiphertexts\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"pkOrgX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pkOrgY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrAx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrAy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrZ\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requireCanSubmitCiphertext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitOrganizerShare\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deltaX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deltaY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"a1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"a1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"a2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"a2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ApplicationRegistered\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"organizerPKx\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"organizerPKy\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrganizerShareSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"deltaX\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"deltaY\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"a1x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"a1y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"a2x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"a2y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyCombined\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ApplicationAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextNotSubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionLimitReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionNotYetAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidApplication\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCiphertext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEpoch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSchnorrProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IsIdentity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotCanonical\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOnCurve\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]}]",
	Bin: "0x60a03461008d57601f61175038819003918201601f19168301916001600160401b038311848410176100915780849260209460405283398101031261008d57516001600160a01b03811680820361008d571561007e576080526040516116aa90816100a68239608051818181610186015281816106dd0152610bed0152f35b63e6c4247b60e01b5f5260045ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe6080806040526004361015610012575f80fd5b5f3560e01c9081631b2df85014610bdb575080632c268ea114610b805780632fed252914610a215780634eef36be1461067357806374a99aba1461052c578063bf37878b146101225763ed78d71e14610069575f80fd5b3461011e57602036600319011261011e576001600160a01b031961008b610c1c565b165f52600160205260405f20604051806020835491828152019081935f5260205f20905f5b81811061010857505050816100c6910382610c95565b604051918291602083019060208452518091526040830191905f5b8181106100ef575050500390f35b82518452859450602093840193909201916001016100e1565b82548452602090930192600192830192016100b0565b5f80fd5b3461011e5761016036600319011261011e5761013c610c1c565b602435608036604319011261011e57604051635f2cdc7560e11b81526001600160a01b03199092166004830181905260c4359260e4359061012435906101043590610240816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610521575f916104f2575b5080516001600160a01b0316156104e3576040015160068110156104cf57600219016104c057841580156104a9575b61049a57835f525f60205260405f20855f5260205260405f2090600582019260ff845460401c1661048b578785916102ac936102298484610f1b565b6102338282610f1b565b5f5160206116555f395f51905f5260405160208101907f41ea6f3fa95eccd1f3b1ce8e05efa11027280aa0c6b4167fd6695db659c30b2882528b60408201528c604c82015285606c82015286608c8201528360ac8201528460cc82015260cc815261029f60ec82610c95565b5190200661014435611096565b1561047c576040516102bd81610c7a565b6044356001600160a01b038116919082810361011e5781526064359061ffff8216820361011e57602081019182526084356001600160401b038116810361011e576040820190815260a435926001600160401b038416840361011e576004946060840194855215610474575b85546001600160a01b03191633178655604051889060209061034a81610c5f565b8d8152015260018681018c9055600287018990559251600387018054925193516001600160f01b03199093166001600160a01b03929092169190911760a09390931b61ffff60a01b169290921760b09190911b600160b01b600160f01b031617905590519190920180546001600160401b0319166001600160401b0392831617905582546001600160481b0319164390911617600160401b9081179092555f8481526020919091526040902080549091811015610460576001810180835581101561044c5784915f5260205f20015560405193845260208401527f0f12fb8a3aa491e558d2d037d94c69b56e8d02f8ff80590c47fb41a5eaaec86b60403394a4005b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b338352610329565b6327f7eb4d60e11b5f5260045ffd5b630b792c8f60e01b5f5260045ffd5b6378e9323b60e11b5f5260045ffd5b505f5160206116355f395f51905f528510156101ed565b63268dbf6760e21b5f5260045ffd5b634e487b7160e01b5f52602160045260245ffd5b63d5b25b6360e01b5f5260045ffd5b61051491506102403d811161051a575b61050c8183610c95565b810190610d63565b876101be565b503d610502565b6040513d5f823e3d90fd5b3461011e57608036600319011261011e57610545610c1c565b61054d610c33565b906064356001600160a01b038116919082900361011e5760018060a01b0319165f525f60205260405f206024355f5260205260405f2060ff600582015460401c161561049a57600361059f9101610cf4565b9060018060a01b03825116036106645760408101516001600160401b03168015159081610651575b506106425760608101516001600160401b0316801515908161062f575b50610620576020015161ffff168015159182610612575b505061060357005b63464e67af60e01b5f5260045ffd5b61ffff1611905081806105fb565b630410ff2960e31b5f5260045ffd5b436001600160401b0316119050836105e4565b633deac39560e01b5f5260045ffd5b436001600160401b0316109050836105c7565b6330cd747160e01b5f5260045ffd5b3461011e576101c036600319011261011e5761068d610c1c565b602435610698610c33565b604051635f2cdc7560e11b81526001600160a01b031984166004820181905291936101243591610104359160e435916101a435916101843591610164359161014435917f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316919061024081602481865afa908115610521575f91610a02575b5080516001600160a01b0316156104e3576040015160068110156104cf57600219016104c05761ffff8c169b8c1580156109f7575b6109e8578a5f525f60205260405f208c5f5260205260ff600560405f20015460401c161561049a5761079f6020828e6040519384928392632de546d560e01b84528860048501610ef4565b0381875afa908115610521575f916109b6575b5080156109a75760806040516064358152608435602082015260a435604082015260c4356060820152200361090357610806928c60609360405195869485938493639bbada6760e01b855260048501610ef4565b03915afa8015610521575f90610932575b6020915001516109235761082b8686610f1b565b6108358188610fa0565b158015610912575b610903575f5160206116555f395f51905f52841015610903577f0f4be0dda65b3c849f2a8f567a3ec58c0c9b322e10877432aebee6629ef2356a9660e09660405160208101908882528260408201528360608201528460808201528560a08201528660c0820152878a8201528981526108b861010082610c95565b5190208a5f52600260205260405f208c5f5260205260405f208d5f5260205260405f2055604051968752602087015260408601526060850152608084015260a083015260c0820152a4005b63d1fed5fd60e01b5f5260045ffd5b5061091d8383610fa0565b1561083d565b63955c0c4960e01b5f5260045ffd5b506060813d60601161099f575b8161094c60609383610c95565b8101031261011e57604051606081016001600160401b038111828210176104605760405261097982610d40565b81526020820151801515810361011e576020926040918484015201516040820152610817565b3d915061093f565b6346f551f560e01b5f5260045ffd5b90506020813d6020116109e0575b816109d160209383610c95565b8101031261011e57518e6107b2565b3d91506109c4565b634c4d29cd60e11b5f5260045ffd5b506101008d11610754565b610a1b91506102403d811161051a5761050c8183610c95565b8d61071f565b3461011e57604036600319011261011e57610a3a610c1c565b5f6080604051610a4981610c44565b828152604051610a5881610c5f565b8381528360208201526020820152610a6e610cd0565b6040820152826060820152015260018060a01b0319165f525f60205260405f206024355f5260205261012060405f2060405190610aaa82610c44565b80546001600160a01b03168252604051610ac381610c5f565b6001820154815260028201546020820152602083019081526005610ae960038401610cf4565b604085810191825291909301546001600160401b03808216606080880191825292841c60ff1615156080808901918252855198516001600160a01b039081168a529651805160208b810191909152908101518a88015297518051909716898601529686015161ffff169688019690965292840151811660a0870152920151821660c0850152511660e0830152511515610100820152f35b3461011e57606036600319011261011e57610b99610c1c565b610ba1610c33565b9060018060a01b0319165f52600260205260405f206024355f5260205261ffff60405f2091165f52602052602060405f2054604051908152f35b3461011e575f36600319011261011e577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b600435906001600160a01b03198216820361011e57565b6044359061ffff8216820361011e57565b60a081019081106001600160401b0382111761046057604052565b604081019081106001600160401b0382111761046057604052565b608081019081106001600160401b0382111761046057604052565b601f909101601f19168101906001600160401b0382119082101761046057604052565b60405190610cc582610c5f565b5f6020838281520152565b60405190610cdd82610c7a565b5f6060838281528260208201528260408201520152565b90604051610d0181610c7a565b82546001600160a01b038116825260a081901c61ffff16602083015260b01c6001600160401b0390811660408301526001909301549092166060830152565b519061ffff8216820361011e57565b51906001600160401b038216820361011e57565b80910390610240821261011e576040519161018083016001600160401b03811184821017610460576040528151906001600160a01b038216820361011e5760e0918452601f19011261011e5760405160e081016001600160401b0381118282101761046057604052610dd760208301610d40565b8152610de560408301610d40565b6020820152610df660608301610d40565b6040820152610e0760808301610d40565b6060820152610e1860a08301610d4f565b6080820152610e2960c08301610d4f565b60a0820152610e3a60e08301610d4f565b60c08201526020830152610100810151600681101561011e57610eeb91610220916040850152610e6d6101208201610d4f565b6060850152610e7f6101408201610d4f565b6080850152610e916101608201610d4f565b60a085015261018081015160c08501526101a081015160e0850152610eb96101c08201610d40565b610100850152610ecc6101e08201610d40565b610120850152610edf6102008201610d40565b61014085015201610d40565b61016082015290565b6001600160a01b03199091168152602081019190915261ffff909116604082015260600190565b905f5160206116355f395f51905f52821080610f8a575b15610f7b57811580610f71575b610f6257610f4c91610fa0565b15610f5357565b6361586bdd60e01b5f5260045ffd5b632b39517d60e21b5f5260045ffd5b5060018114610f3f565b63d7c7beeb60e01b5f5260045ffd5b505f5160206116355f395f51905f528110610f32565b5f5160206116355f395f51905f528110801590611050575b61104a575f5160206116355f395f51905f528181920991800990805f5160206116355f395f51905f5203915f5160206116355f395f51905f528311611036575f5160206116355f395f51905f528080838195097f1aee90f15f2189693df072d799fd11fc039b2959ebb7c867d075ca8cf4d7eb8e0960010892081490565b634e487b7160e01b5f52601160045260245ffd5b50505f90565b505f5160206116355f395f51905f52821015610fb8565b8115611071570690565b634e487b7160e01b5f52601260045260245ffd5b90601081101561044c5760051b0190565b94939190945f5160206116355f395f51905f52821080611522575b1580156114ef575b6114e5575f5160206116555f395f51905f5280910695069180155f146114bc57505f945b6110e5610cb8565b50604051956110f387610c5f565b5f5160206116355f395f51905f5287527f0578d36fdd1172a8c3909ff8b278cb9adf026a6b5db6203e5d099f85f9afd71b60208801526040519561020061113a8189610c95565b5f5b8181106114a75750506001602088510152600160408851015260808701518851801561107157806060917f1561ff836ce19d358a4eb7a4c199e94c377c749ae6f2a277f1f9195afe553f9f068352807f25797203f7a0b24925572e1cd16bf9edfce0051fb9e133774b3c257a872d7d8b06806020850152600160408501528351099101526101008701518851801561107157806060917f0daaa7e6b25c28e6dc8dd1d48e9cc61cd07015c1d7c1b8d4590eb6f51d5346dc068352807f01666cafbf0a30da8b9ebeaf848a1da067a892296f1043188e1705402b6d68530680602085015260016040850152835109910152610180870151918851928315611071576060846112b3957f136d609c4c856f5d277fab08c730cbdd1a776ce4728c6a2eb20ff22bccf26894068352807f21d66f0e2295ae954494f25889f9319cc1b4df71eff3f46ba9e4631b43fd7c9506806020850152600160408501528351099101526112ba60208901918251938b5195868093611067565b8552611067565b918260208201526001604082015280518415611071576112ff948b946060920991015260408801906112f0838351835190611538565b60608901519151905191611599565b60045b600c81111561143d57506040519561131987610c7a565b5f875260208701956001875260408801946001865260608901905f82525f9460fc805b61139f575050505050505081518015159586611381575b50505083611362575b50505090565b5f5160206116355f395f51905f529293505190099051145f808061135c565b519295505f5160206116355f395f51905f52910914925f8080611353565b600119018082811c60021b600c1684821c60031617878e8a156113fd576113c7828280611538565b6113d2828280611538565b826113e0575b50505061133c565b6113ed6113f5938a611085565b519080611599565b5f878e6113d8565b505080915061140e575b508061133c565b606091975061141d9086611085565b5180518d5260208101518c5260408101518a520151835285600196611407565b60015b600481106114695750600481018091111561130257634e487b7160e01b5f52601160045260245ffd5b80820190818311611036576114a1896114846001948b611085565b5161148f868c611085565b5161149a858d611085565b5191611599565b01611440565b6020906114b2610cd0565b818b01520161113c565b5f5160206116355f395f51905f52035f5160206116355f395f51905f52811161103657946110dd565b5050505050505f90565b505f5160206116355f395f51905f5283108061150c575b156110b9565b505f5160206116355f395f51905f528510611506565b505f5160206116355f395f51905f5284106110b1565b9151815191602081019081518315611071578380808093604098088180808a818080808c5180099c518009818d810382089c08810380988782980908980151800980088103870894828682098a520960608801528309602086015209910152565b929081519260208201928351835186039186156110715786808086818080999881808d81809d9c816020819f01968188518c51820390089208099f519051900891518551900890099a818181038d089b089687958160608c015160608501519009906020015190099860400151906040015190098008958181810388089608958286820989520960608701528309602085015209906040015256fe30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001060c89ce5c263405370a08b6d0302b0bab3eedb83920ee0a677297dc392126f1a264697066735822122024d2e8e47b6727c19ee3bdb4413f0a2ac8694d66902be367db0dd20b0df1d30f64736f6c634300081c0033",
}

// DKGAppManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DKGAppManagerMetaData.ABI instead.
var DKGAppManagerABI = DKGAppManagerMetaData.ABI

// DKGAppManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DKGAppManagerMetaData.Bin instead.
var DKGAppManagerBin = DKGAppManagerMetaData.Bin

// DeployDKGAppManager deploys a new Ethereum contract, binding an instance of DKGAppManager to it.
func DeployDKGAppManager(auth *bind.TransactOpts, backend bind.ContractBackend, _manager common.Address) (common.Address, *types.Transaction, *DKGAppManager, error) {
	parsed, err := DKGAppManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DKGAppManagerBin), backend, _manager)
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

// GetApplication is a free data retrieval call binding the contract method 0x2fed2529.
//
// Solidity: function getApplication(bytes12 epochId, bytes32 aid) view returns((address,(uint256,uint256),(address,uint16,uint64,uint64),uint64,bool))
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
// Solidity: function getApplication(bytes12 epochId, bytes32 aid) view returns((address,(uint256,uint256),(address,uint16,uint64,uint64),uint64,bool))
func (_DKGAppManager *DKGAppManagerSession) GetApplication(epochId [12]byte, aid [32]byte) (DKGTypesApplication, error) {
	return _DKGAppManager.Contract.GetApplication(&_DKGAppManager.CallOpts, epochId, aid)
}

// GetApplication is a free data retrieval call binding the contract method 0x2fed2529.
//
// Solidity: function getApplication(bytes12 epochId, bytes32 aid) view returns((address,(uint256,uint256),(address,uint16,uint64,uint64),uint64,bool))
func (_DKGAppManager *DKGAppManagerCallerSession) GetApplication(epochId [12]byte, aid [32]byte) (DKGTypesApplication, error) {
	return _DKGAppManager.Contract.GetApplication(&_DKGAppManager.CallOpts, epochId, aid)
}

// GetOrganizerShareHash is a free data retrieval call binding the contract method 0x2c268ea1.
//
// Solidity: function getOrganizerShareHash(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) view returns(bytes32)
func (_DKGAppManager *DKGAppManagerCaller) GetOrganizerShareHash(opts *bind.CallOpts, epochId [12]byte, aid [32]byte, ciphertextIndex uint16) ([32]byte, error) {
	var out []interface{}
	err := _DKGAppManager.contract.Call(opts, &out, "getOrganizerShareHash", epochId, aid, ciphertextIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOrganizerShareHash is a free data retrieval call binding the contract method 0x2c268ea1.
//
// Solidity: function getOrganizerShareHash(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) view returns(bytes32)
func (_DKGAppManager *DKGAppManagerSession) GetOrganizerShareHash(epochId [12]byte, aid [32]byte, ciphertextIndex uint16) ([32]byte, error) {
	return _DKGAppManager.Contract.GetOrganizerShareHash(&_DKGAppManager.CallOpts, epochId, aid, ciphertextIndex)
}

// GetOrganizerShareHash is a free data retrieval call binding the contract method 0x2c268ea1.
//
// Solidity: function getOrganizerShareHash(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) view returns(bytes32)
func (_DKGAppManager *DKGAppManagerCallerSession) GetOrganizerShareHash(epochId [12]byte, aid [32]byte, ciphertextIndex uint16) ([32]byte, error) {
	return _DKGAppManager.Contract.GetOrganizerShareHash(&_DKGAppManager.CallOpts, epochId, aid, ciphertextIndex)
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

// RegisterApplication is a paid mutator transaction binding the contract method 0xbf37878b.
//
// Solidity: function registerApplication(bytes12 epochId, bytes32 aid, (address,uint16,uint64,uint64) policy, uint256 pkOrgX, uint256 pkOrgY, uint256 schnorrAx, uint256 schnorrAy, uint256 schnorrZ) returns()
func (_DKGAppManager *DKGAppManagerTransactor) RegisterApplication(opts *bind.TransactOpts, epochId [12]byte, aid [32]byte, policy DKGTypesAppPolicy, pkOrgX *big.Int, pkOrgY *big.Int, schnorrAx *big.Int, schnorrAy *big.Int, schnorrZ *big.Int) (*types.Transaction, error) {
	return _DKGAppManager.contract.Transact(opts, "registerApplication", epochId, aid, policy, pkOrgX, pkOrgY, schnorrAx, schnorrAy, schnorrZ)
}

// RegisterApplication is a paid mutator transaction binding the contract method 0xbf37878b.
//
// Solidity: function registerApplication(bytes12 epochId, bytes32 aid, (address,uint16,uint64,uint64) policy, uint256 pkOrgX, uint256 pkOrgY, uint256 schnorrAx, uint256 schnorrAy, uint256 schnorrZ) returns()
func (_DKGAppManager *DKGAppManagerSession) RegisterApplication(epochId [12]byte, aid [32]byte, policy DKGTypesAppPolicy, pkOrgX *big.Int, pkOrgY *big.Int, schnorrAx *big.Int, schnorrAy *big.Int, schnorrZ *big.Int) (*types.Transaction, error) {
	return _DKGAppManager.Contract.RegisterApplication(&_DKGAppManager.TransactOpts, epochId, aid, policy, pkOrgX, pkOrgY, schnorrAx, schnorrAy, schnorrZ)
}

// RegisterApplication is a paid mutator transaction binding the contract method 0xbf37878b.
//
// Solidity: function registerApplication(bytes12 epochId, bytes32 aid, (address,uint16,uint64,uint64) policy, uint256 pkOrgX, uint256 pkOrgY, uint256 schnorrAx, uint256 schnorrAy, uint256 schnorrZ) returns()
func (_DKGAppManager *DKGAppManagerTransactorSession) RegisterApplication(epochId [12]byte, aid [32]byte, policy DKGTypesAppPolicy, pkOrgX *big.Int, pkOrgY *big.Int, schnorrAx *big.Int, schnorrAy *big.Int, schnorrZ *big.Int) (*types.Transaction, error) {
	return _DKGAppManager.Contract.RegisterApplication(&_DKGAppManager.TransactOpts, epochId, aid, policy, pkOrgX, pkOrgY, schnorrAx, schnorrAy, schnorrZ)
}

// SubmitOrganizerShare is a paid mutator transaction binding the contract method 0x4eef36be.
//
// Solidity: function submitOrganizerShare(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, uint256 deltaX, uint256 deltaY, uint256 a1x, uint256 a1y, uint256 a2x, uint256 a2y, uint256 z) returns()
func (_DKGAppManager *DKGAppManagerTransactor) SubmitOrganizerShare(opts *bind.TransactOpts, epochId [12]byte, aid [32]byte, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int, deltaX *big.Int, deltaY *big.Int, a1x *big.Int, a1y *big.Int, a2x *big.Int, a2y *big.Int, z *big.Int) (*types.Transaction, error) {
	return _DKGAppManager.contract.Transact(opts, "submitOrganizerShare", epochId, aid, ciphertextIndex, c1x, c1y, c2x, c2y, deltaX, deltaY, a1x, a1y, a2x, a2y, z)
}

// SubmitOrganizerShare is a paid mutator transaction binding the contract method 0x4eef36be.
//
// Solidity: function submitOrganizerShare(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, uint256 deltaX, uint256 deltaY, uint256 a1x, uint256 a1y, uint256 a2x, uint256 a2y, uint256 z) returns()
func (_DKGAppManager *DKGAppManagerSession) SubmitOrganizerShare(epochId [12]byte, aid [32]byte, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int, deltaX *big.Int, deltaY *big.Int, a1x *big.Int, a1y *big.Int, a2x *big.Int, a2y *big.Int, z *big.Int) (*types.Transaction, error) {
	return _DKGAppManager.Contract.SubmitOrganizerShare(&_DKGAppManager.TransactOpts, epochId, aid, ciphertextIndex, c1x, c1y, c2x, c2y, deltaX, deltaY, a1x, a1y, a2x, a2y, z)
}

// SubmitOrganizerShare is a paid mutator transaction binding the contract method 0x4eef36be.
//
// Solidity: function submitOrganizerShare(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, uint256 deltaX, uint256 deltaY, uint256 a1x, uint256 a1y, uint256 a2x, uint256 a2y, uint256 z) returns()
func (_DKGAppManager *DKGAppManagerTransactorSession) SubmitOrganizerShare(epochId [12]byte, aid [32]byte, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int, deltaX *big.Int, deltaY *big.Int, a1x *big.Int, a1y *big.Int, a2x *big.Int, a2y *big.Int, z *big.Int) (*types.Transaction, error) {
	return _DKGAppManager.Contract.SubmitOrganizerShare(&_DKGAppManager.TransactOpts, epochId, aid, ciphertextIndex, c1x, c1y, c2x, c2y, deltaX, deltaY, a1x, a1y, a2x, a2y, z)
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
	OrganizerPKx *big.Int
	OrganizerPKy *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterApplicationRegistered is a free log retrieval operation binding the contract event 0x0f12fb8a3aa491e558d2d037d94c69b56e8d02f8ff80590c47fb41a5eaaec86b.
//
// Solidity: event ApplicationRegistered(bytes12 indexed epochId, bytes32 indexed aid, address indexed creator, uint256 organizerPKx, uint256 organizerPKy)
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

// WatchApplicationRegistered is a free log subscription operation binding the contract event 0x0f12fb8a3aa491e558d2d037d94c69b56e8d02f8ff80590c47fb41a5eaaec86b.
//
// Solidity: event ApplicationRegistered(bytes12 indexed epochId, bytes32 indexed aid, address indexed creator, uint256 organizerPKx, uint256 organizerPKy)
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

// ParseApplicationRegistered is a log parse operation binding the contract event 0x0f12fb8a3aa491e558d2d037d94c69b56e8d02f8ff80590c47fb41a5eaaec86b.
//
// Solidity: event ApplicationRegistered(bytes12 indexed epochId, bytes32 indexed aid, address indexed creator, uint256 organizerPKx, uint256 organizerPKy)
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
	DeltaX          *big.Int
	DeltaY          *big.Int
	A1x             *big.Int
	A1y             *big.Int
	A2x             *big.Int
	A2y             *big.Int
	Z               *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOrganizerShareSubmitted is a free log retrieval operation binding the contract event 0x0f4be0dda65b3c849f2a8f567a3ec58c0c9b322e10877432aebee6629ef2356a.
//
// Solidity: event OrganizerShareSubmitted(bytes12 indexed epochId, bytes32 indexed aid, uint16 indexed ciphertextIndex, uint256 deltaX, uint256 deltaY, uint256 a1x, uint256 a1y, uint256 a2x, uint256 a2y, uint256 z)
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

// WatchOrganizerShareSubmitted is a free log subscription operation binding the contract event 0x0f4be0dda65b3c849f2a8f567a3ec58c0c9b322e10877432aebee6629ef2356a.
//
// Solidity: event OrganizerShareSubmitted(bytes12 indexed epochId, bytes32 indexed aid, uint16 indexed ciphertextIndex, uint256 deltaX, uint256 deltaY, uint256 a1x, uint256 a1y, uint256 a2x, uint256 a2y, uint256 z)
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

// ParseOrganizerShareSubmitted is a log parse operation binding the contract event 0x0f4be0dda65b3c849f2a8f567a3ec58c0c9b322e10877432aebee6629ef2356a.
//
// Solidity: event OrganizerShareSubmitted(bytes12 indexed epochId, bytes32 indexed aid, uint16 indexed ciphertextIndex, uint256 deltaX, uint256 deltaY, uint256 a1x, uint256 a1y, uint256 a2x, uint256 a2y, uint256 z)
func (_DKGAppManager *DKGAppManagerFilterer) ParseOrganizerShareSubmitted(log types.Log) (*DKGAppManagerOrganizerShareSubmitted, error) {
	event := new(DKGAppManagerOrganizerShareSubmitted)
	if err := _DKGAppManager.contract.UnpackLog(event, "OrganizerShareSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
