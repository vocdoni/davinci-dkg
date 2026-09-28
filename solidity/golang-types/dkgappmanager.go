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
	Mode             uint8
	OpenSubmission   bool
	Submitters       []common.Address
	MaxCiphertexts   uint16
	NotBeforeBlock   uint64
	NotAfterBlock    uint64
	DecryptNotBefore uint64
	DecryptNotAfter  uint64
}

// DKGTypesApplication is an auto generated low-level Go binding around an user-defined struct.
type DKGTypesApplication struct {
	Creator         common.Address
	OrganizerPK     DKGTypesPoint
	OrganizerSecret *big.Int
	PoolIndex       uint8
	Policy          DKGTypesAppPolicy
	CreatedAtBlock  uint64
	Exists          bool
}

// DKGTypesPoint is an auto generated low-level Go binding around an user-defined struct.
type DKGTypesPoint struct {
	X *big.Int
	Y *big.Int
}

// DKGAppManagerMetaData contains all meta data concerning the DKGAppManager contract.
var DKGAppManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MANAGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApplication\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Application\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"organizerPK\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"organizerSecret\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"poolIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.AppPolicy\",\"components\":[{\"name\":\"mode\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.AppMode\"},{\"name\":\"openSubmission\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"submitters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"maxCiphertexts\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"decryptNotBefore\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"decryptNotAfter\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"createdAtBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApplicationKey\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrganizerPK\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredAids\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerApplication\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.AppPolicy\",\"components\":[{\"name\":\"mode\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.AppMode\"},{\"name\":\"openSubmission\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"submitters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"maxCiphertexts\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"decryptNotBefore\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"decryptNotAfter\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"pkOrgX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pkOrgY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrAx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrAy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrZ\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registrarAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireCanSubmitCiphertext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireDecryptionOpen\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revealOrganizerSecret\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"organizerSecret\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRegistrar\",\"inputs\":[{\"name\":\"r\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ApplicationRegistered\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"organizerPKx\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"organizerPKy\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"mode\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumDKGTypes.AppMode\"},{\"name\":\"poolIndex\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrganizerSecretRevealed\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"organizerSecret\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistrarSet\",\"inputs\":[{\"name\":\"registrar\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyRevealed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ApplicationAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionClosed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionLimitReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionNotOpen\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionNotYetAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidApplication\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEpoch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOrganizerSecret\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRegistrar\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSchnorrProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IsIdentity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotCanonical\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOnCurve\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistrar\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrganizerSecretNotRevealed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PointNotInSubgroup\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolExhausted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x60c0346100a157601f61261238819003918201601f19168301916001600160401b038311848410176100a5578084926020946040528339810103126100a157516001600160a01b0381168082036100a15715610092576080523360a05260405161255890816100ba8239608051818181610190015281816102c90152611233015260a0518181816102850152610a9b0152f35b63e6c4247b60e01b5f5260045ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe60806040526004361015610011575f80fd5b5f3560e01c80630fa5744c146100d4578063112fc40d146100cf5780631b2df850146100ca5780632b20e397146100c55780632fed2529146100c05780636ed93d64146100bb57806374a99aba146100b657806381fe92fb146100b1578063a59b7a4d146100ac578063ed78d71e146100a7578063f6d651b4146100a25763faab9d391461009d575f80fd5b610a7c565b6109a6565b610927565b6107dd565b610788565b6105e3565b610561565b6104b1565b6102f8565b6102b4565b610270565b6100f4565b600435906001600160a01b0319821682036100f057565b5f80fd5b346100f05760403660031901126100f05761010d6100d9565b61012960243561011c83610b35565b905f5260205260405f2090565b9061014361013f600984015460ff9060401c1690565b1590565b61026157604061018c9161015b600485015460ff1690565b82516356cbb5f360e01b81526001600160a01b0319909216600483015260ff16602482015291829081906044820190565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561025c575f905f9261022a575b50819281926101dd600583015460ff1690565b6101e681610334565b15610208575b50506040805192835260208301939093525090819081015b0390f35b90919250610221935060026001830154920154926116b1565b905f80806101ec565b905061024e915060403d604011610255575b6102468183610bd5565b810190610bf8565b905f6101ca565b503d61023c565b610c0e565b6378e9323b60e11b5f5260045ffd5b346100f0575f3660031901126100f0576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346100f0575f3660031901126100f0576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346100f0575f3660031901126100f0576002546040516001600160a01b039091168152602090f35b634e487b7160e01b5f52602160045260245ffd5b6002111561033e57565b610320565b90600282101561033e5752565b6001600160401b031690565b6001600160401b03169052565b9061010081019161037b828251610343565b60208101511515602083015260408101519261010060408401528351809152602061012084019401905f5b818110610410575050509060e080836103cd606061040d960151606086019061ffff169052565b6103df6080820151608086019061035c565b6103f160a082015160a086019061035c565b61040360c082015160c086019061035c565b015191019061035c565b90565b82516001600160a01b03168652602095860195909201916001016103a6565b61040d906020815260018060a01b038351166020820152602080840151805160408401520151606082015260408301516080820152610478606084015160a083019060ff169052565b61010060c061049560808601518383860152610120850190610369565b946104a860a082015160e086019061035c565b01511515910152565b346100f05760403660031901126100f0576102046105556105506104d36100d9565b61011c602435915f60c06040516104e981610b63565b8281526104f4610c49565b602082015282604082015282606082015260405161051181610b83565b838152836020820152606060408201528360608201528360808201528360a082015283838201528360e082015260808201528260a08201520152610b35565b610db6565b6040519182918261042f565b346100f0576101003660031901126100f05761057b6100d9565b6044356024356001600160401b0382116100f05761010060031983360301126100f0576105c4926064356084359060a4359260c435946105ba60e43590565b96600401916111e5565b005b61ffff8116036100f057565b6001600160a01b038116036100f057565b346100f05760803660031901126100f0576105fc6100d9565b60443561062360243561060e836105c6565b61011c6064359461061e866105d2565b610b35565b9161063961013f600985015460ff9060401c1690565b6102615760058301546106509060081c60ff161590565b9081610773575b506107645760078201546001600160401b03601082901c168015159081610749575b5061073a57610699605082901c6001600160401b0316610350565b610350565b8015159081610727575b506107185761ffff16801515919082610706575b50506106f75761069460086106cd920154610350565b80151590816106ed575b506106de57005b632300418b60e01b5f5260045ffd5b905042115f6106d7565b63464e67af60e01b5f5260045ffd5b61ffff90811691161190505f806106b7565b630410ff2960e31b5f5260045ffd5b436001600160401b03161190505f6106a3565b633deac39560e01b5f5260045ffd5b6107539150610350565b436001600160401b0316105f610679565b6330cd747160e01b5f5260045ffd5b610782915061013f9084611ae6565b5f610657565b346100f05760403660031901126100f0576107a16100d9565b6001600160a01b0319165f9081526020818152604080832060243584528252918290206001810154600290910154835191825291810191909152f35b346100f05760603660031901126100f0576107f66100d9565b602435906044359061080b8361011c83610b35565b600981015461081e9060401c60ff161590565b61026157600581015460ff1661083381610334565b6108df57600381019081546108df57831580156108c8575b6108aa576108588461216e565b90600183015414918215926108b9575b50506108aa578290556040519182526001600160a01b031916907fa6d794f0981501d1b02889a9815c10a2bd90f6486541351fd1258d2ed331867790602090a3005b634102642560e11b5f5260045ffd5b60020154141590505f80610868565b505f5160206124e35f395f51905f5284101561084b565b63a89ac15160e01b5f5260045ffd5b60206040818301928281528451809452019201905f5b8181106109115750505090565b8251845260209384019390920191600101610904565b346100f05760203660031901126100f0576001600160a01b03196109496100d9565b165f52600160205260405f206040519081602082549182815201915f5260205f20905f5b818110610990576102048561098481870382610bd5565b604051918291826108ee565b825484526020909301926001928301920161096d565b346100f05760403660031901126100f0576109ce6109c26100d9565b61011c60243591610b35565b60098101546109e19060401c60ff161590565b61026157600581015460ff166109f681610334565b1580610a70575b610a6157600781015460901c6001600160401b03168015159081610a57575b50610a4857600801546001600160401b0390610a3790610350565b1680151590816106ed57506106de57005b6314badd7360e31b5f5260045ffd5b905042105f610a1c565b630a4bc8e760e41b5f5260045ffd5b506003810154156109fd565b346100f05760203660031901126100f057600435610a99816105d2565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303610b27576001600160a01b03168015610b1857600280546001600160a01b031916821790556040519081527f6263309d5d4d1cfececd45a387cda7f14dccde21cf7a1bee1be6561075e6101490602090a1005b6372491c3360e11b5f5260045ffd5b6282b42960e81b5f5260045ffd5b6001600160a01b0319165f90815260208190526040902090565b634e487b7160e01b5f52604160045260245ffd5b60e081019081106001600160401b03821117610b7e57604052565b610b4f565b61010081019081106001600160401b03821117610b7e57604052565b604081019081106001600160401b03821117610b7e57604052565b608081019081106001600160401b03821117610b7e57604052565b601f909101601f19168101906001600160401b03821190821017610b7e57604052565b91908260409103126100f0576020825192015190565b6040513d5f823e3d90fd5b60405190610c2961018083610bd5565b565b60405190610c29604083610bd5565b60405190610c29608083610bd5565b60405190610c5682610b9f565b5f6020838281520152565b90604051610c6e81610b9f565b602060018294805484520154910152565b90604051918281549182825260208201905f5260205f20925f5b818110610cae575050610c2992500383610bd5565b84546001600160a01b0316835260019485019487945060209093019201610c99565b6001600160401b039091169052565b9060405191610ced83610b83565b8281549160ff831692600284101561033e57600360e092610d27610d1e610daf94610c2998885260ff9060081c1690565b15156020870152565b610d3360018201610c7f565b6040860152610da8610d9f6002830154610d5c610d518261ffff1690565b61ffff1660608a0152565b610d76601082901c6001600160401b031660808a01610cd0565b610d90605082901c6001600160401b031660a08a01610cd0565b60901c6001600160401b031690565b60c08701610cd0565b0154610350565b9101610cd0565b90610c29604051610dc681610b63565b83546001600160a01b0316815292839060c090610e4790600990610dec60018201610c61565b602086015260038101546040860152610e16610e0c600483015460ff1690565b60ff166060870152565b610e2260058201610cdf565b60808601520154610e3e610e3582610350565b60a08601610cd0565b60401c60ff1690565b1515910152565b5190610c29826105d2565b5190610c29826105c6565b6001600160401b038116036100f057565b5190610c2982610e64565b91908260e09103126100f057604051610e9881610b63565b60c0610f0a8183958051610eab816105c6565b85526020810151610ebb816105c6565b6020860152610ecc60408201610e59565b6040860152610edd60608201610e59565b6060860152610eee60808201610e75565b6080860152610eff60a08201610e75565b60a086015201610e75565b910152565b519060068210156100f057565b610240818303126100f057610220610ff091610f4e610f39610c19565b94610f4383610e4e565b865260208301610e80565b6020850152610f606101008201610f0f565b6040850152610f726101208201610e75565b6060850152610f846101408201610e75565b6080850152610f966101608201610e75565b60a085015261018081015160c08501526101a081015160e0850152610fbe6101c08201610e59565b610100850152610fd16101e08201610e59565b610120850152610fe46102008201610e59565b61014085015201610e59565b61016082015290565b6006111561033e57565b3560028110156100f05790565b908160209103126100f0575160ff811681036100f05790565b90600281101561033e5760ff80198354169116179055565b3580151581036100f05790565b903590601e19813603018212156100f057018035906001600160401b0382116100f057602001918160051b360383136100f057565b634e487b7160e01b5f52601160045260245ffd5b3561040d816105d2565b906001600160401b038311610b7e57600160401b8311610b7e578154838355808410611101575b506110d790915f5260205f2090565b5f5b8381106110e65750505050565b60019060206110f485611097565b94019381840155016110d9565b825f528360205f2091820191015b81811061111c57506110c8565b5f815560010161110f565b3561040d816105c6565b3561040d81610e64565b80546001600160401b0319166001600160401b03909216919091179055565b634e487b7160e01b5f52603260045260245ffd5b8054821015611183575f5260205f2001905f90565b61115a565b8054600160401b811015610b7e576111a59160018201815561116e565b819291549060031b91821b915f19901b1916179055565b92936111df60ff9296956060946080870198875260208701526040860190610343565b16910152565b6002549097919694959491939291906001600160a01b0316801515908161169d575b5061168e57604051635f2cdc7560e11b81526001600160a01b0319891660048201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016939061024081602481885afa90811561025c575f9161165f575b50805161128a906001600160a01b03165b6001600160a01b031690565b15611650576040600391015161129f81610ff9565b6112a881610ff9565b03611641576112b688611725565b6112c38861011c8b610b35565b9460098601926112d8845460ff9060401c1690565b611632576112e58561175e565b5f9760016112f287611003565b6112fb81610334565b036115b757505050506113489495506020876001945b60405163421adfbb60e01b81526001600160a01b03198c1660048201526024810192909252909687919082905f9082906044820190565b03925af193841561025c577f763f6464c291e9532f5ac2b75d76aa5283fc376d894a5e0b1e750ea0532cc7b1955f95611572575b5080546001600160a01b0319163317815561156d926115509290916115289190611502906113c46113ab610c2b565b8c8152602001899052600182018c905560028201899055565b60048101805460ff191660ff8b16179055611410600582016113ee6113e888611003565b82611029565b6113fa60208801611041565b815461ff00191690151560081b61ff0016179055565b61142a611420604087018761104e565b90600684016110a1565b6114ed61143960608701611127565b6114536007840191829061ffff1661ffff19825416179055565b61148661146260808901611131565b825462010000600160501b03191660109190911b62010000600160501b0316178255565b6114bb61149560a08901611131565b8254600160501b600160901b03191660509190911b600160501b600160901b0316178255565b6114c760c08801611131565b8154600160901b600160d01b03191660909190911b600160901b600160d01b0316179055565b60086114fb60e08701611131565b910161113b565b611515436001600160401b03168261113b565b805460ff60401b1916600160401b179055565b6001600160a01b031989165f90815260016020526040902061154b908990611188565b611003565b60405133986001600160a01b0319169690948594909291856111bc565b0390a4565b61156d93919550611550926115026115a46115289360203d6020116115b0575b61159c8183610bd5565b810190611010565b9793955050925061137c565b503d611592565b90919297506115c689896118f6565b6115d361013f8a8a61197b565b61162357878282611602956115fc8f968f968f87819961013f9b6115f786866118f6565b611eb8565b90611f60565b61161457602087611348969794611311565b6327f7eb4d60e11b5f5260045ffd5b63b28e789160e01b5f5260045ffd5b630b792c8f60e01b5f5260045ffd5b63268dbf6760e21b5f5260045ffd5b63d5b25b6360e01b5f5260045ffd5b61168191506102403d8111611687575b6116798183610bd5565b810190610f1c565b5f61126d565b503d61166f565b632ad9535160e01b5f5260045ffd5b6001600160a01b031690503314155f611207565b90939281158061171b575b6117145782158061170a575b61170357936116f36116fa926116ff95966116ed6116e4611b55565b80958194611bc7565b95611bc7565b8380611c12565b611cb2565b9091565b5090509190565b50600181146116c8565b9350509190565b50600185146116bc565b8015908115611736575b5061026157565b5f5160206124c35f395f51905f52915010155f61172f565b91908110156111835760051b0190565b6040810161176c818361104e565b9190506020821180156118da575b6117fa575f5b8281106118ae5750505060e081016001600160401b0361179f82611131565b16151580611891575b6117fa5760c08201906117bd61069483611131565b1515918261187b575b82611850575b50506117fa5760808101906117e361069483611131565b15159182611837575b82611809575b50506117fa57565b63d06b96b160e01b5f5260045ffd5b61182491925060a061181d61069492611131565b9301611131565b6001600160401b03909116115f806117f2565b915061184861069460a08301611131565b1515916117ec565b61186891925061186261069491611131565b92611131565b6001600160401b03909116115f806117cc565b915061188961069482611131565b1515916117c6565b5061189b81611131565b426001600160401b0390911611156117a8565b6118cd61127e6118c8836118c2868961104e565b9061174e565b611097565b156117fa57600101611780565b506118e760208401611041565b801561177a575081151561177a565b905f5160206124c35f395f51905f52821080611965575b156119565781158061194c575b61193d5761192791611d90565b1561192e57565b6361586bdd60e01b5f5260045ffd5b632b39517d60e21b5f5260045ffd5b506001811461191a565b63d7c7beeb60e01b5f5260045ffd5b505f5160206124c35f395f51905f52811061190d565b801580611adc575b611ad5576119918282611d90565b15611acf5761199e611ba3565b506119e96119aa611b55565b80926119c6826119b8611e60565b96602088015b938451612393565b60408501906119d98383518351906123bb565b606086015b519151905191611c12565b6119f1610c3a565b5f81526020810192600184526040820192600184525f60608401525f91610100805b611a42575050505051159182611a36575b5081611a2e575090565b905051151590565b5181511491505f611a24565b60011901805f5160206124e35f395f51905f52811c6003168515611aa057611a6b8588806123bb565b611a768588806123bb565b8481611a84575b5050611a13565b611a91611a999286611e94565b518880611c12565b5f84611a7d565b809150611aaf575b5080611a13565b611ac6919450611abf9083611e94565b518561241c565b82600193611aa8565b50505f90565b5050600190565b5060018214611983565b90600682018054928315611b3d57505f5b838110611b0657505050505f90565b611b10818361116e565b905460039190911b1c6001600160a01b0390811690841614611b3457600101611af7565b50505050600190565b546001600160a01b0392831692169190911492915050565b611b5d610c49565b50604051611b6a81610b9f565b5f5160206124c35f395f51905f5281527f0578d36fdd1172a8c3909ff8b278cb9adf026a6b5db6203e5d099f85f9afd71b602082015290565b60405190611bb082610bba565b5f6060838281528260208201528260408201520152565b9291610c2991611bd5611ba3565b5060405194611be386610bba565b5f86525f60208701525f60408701525f606087015285612393565b634e487b7160e01b5f52601260045260245ffd5b92908151926020820192835183518603918615611cad5786808086818080999881808d81809d9c816020819f01968188518c51820390089208099f519051900891518551900890099a818181038d089b089687958160608c0151606085015190099060200151900998604001519060400151900980089581818103880896089582868209895209606087015283096020850152099060400152565b611bfe565b90604082015160405190602082526020808301526020604083015260608201527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593efffffff60808201525f5160206124c35f395f51905f5260a082015260208160c08160055afa156100f0575f5160206124c35f395f51905f529051602082828651099401510990565b5f5160206124c35f395f51905f52119081611d51575090565b5f5160206124c35f395f51905f5291501090565b5f5160206124c35f395f51905f5203905f5160206124c35f395f51905f528211611d8b57565b611083565b5f5160206124c35f395f51905f528110801590611e26575b611acf575f5160206124c35f395f51905f528181920991800990805f5160206124c35f395f51905f5203915f5160206124c35f395f51905f528311611d8b575f5160206124c35f395f51905f528080838195097f1aee90f15f2189693df072d799fd11fc039b2959ebb7c867d075ca8cf4d7eb8e0960010892081490565b505f5160206124c35f395f51905f52821015611da8565b5f5b828110611e4b57505050565b602090611e56611ba3565b8184015201611e3f565b60405190611e6f608083610bd5565b610c29608083611e3d565b60405190610c29610200611e8e8185610bd5565b83611e3d565b9060048110156111835760051b0190565b5f5160206124e35f395f51905f52900690565b9390925f5160206124e35f395f51905f5295926040519460208601967f41ea6f3fa95eccd1f3b1ce8e05efa11027280aa0c6b4167fd6695db659c30b28885260018060a01b0319166040870152604c860152606c850152608c84015260ac83015260cc82015260cc8152611f2d60ec82610bd5565b5190200690565b9060108110156111835760051b0190565b9060048201809211611d8b57565b91908201809211611d8b57565b959491929395611f7361013f8686611d38565b801561215b575b61215157611f8a611f9091611ea5565b91611ea5565b958061214357505f5b61201a611fa4611b55565b8092611ffe82611fb2611e7a565b97600160208a510152600160408a510152611fd782611fd18b60800190565b51612281565b611fe6826101008b01516122c3565b611ff5826101808b015161232b565b602089016119be565b60408601906120118383518351906123bb565b606087016119de565b60045b600c8111156120ec575061202f610c3a565b925f845260016020850152600160408501525f60608501525f9260fc805b61206157505050505061040d939450612440565b600119019889600c83821c60021b1682821c6003161786156120c0576120888689806123bb565b6120938689806123bb565b85816120a4575b50505b909961204d565b6120b16120b99287611f34565b518980611c12565b5f8561209a565b806120cc575b5061209d565b6120e39196506120dc9085611f34565b518761241c565b6001945f6120c6565b60015b60048110612106575061210190611f45565b61201d565b8061213d8461212061211a60019587611f53565b89611f34565b5161212b868a611f34565b51612136858b611f34565b5191611c12565b016120ef565b61214c90611d65565b611f99565b505f955050505050565b5061216961013f8489611d38565b611f7a565b5f5160206124e35f395f51905f5290069081156122795761218d611ba3565b50612196611b55565b61219e611e60565b906121ce81602084016119c6825f5160206124a35f395f51905f525f5160206125035f395f51905f528451612393565b6121d6610c3a565b915f835260016020840152600160408401525f60608401525f94610100805b61220857505050506116ff919250611cb2565b600119018082811c6003168815612251576122248688806123bb565b61222f8688806123bb565b858161223d575b50506121f5565b611a9161224a9287611e94565b5f85612236565b809150612260575b50806121f5565b612270919750611abf9084611e94565b85600196612259565b5f9150600190565b90518015611cad57806060915f5160206125035f395f51905f52068352805f5160206124a35f395f51905f520680602085015260016040850152835109910152565b90518015611cad57806060917f0daaa7e6b25c28e6dc8dd1d48e9cc61cd07015c1d7c1b8d4590eb6f51d5346dc068352807f01666cafbf0a30da8b9ebeaf848a1da067a892296f1043188e1705402b6d68530680602085015260016040850152835109910152565b90518015611cad57806060917f136d609c4c856f5d277fab08c730cbdd1a776ce4728c6a2eb20ff22bccf26894068352807f21d66f0e2295ae954494f25889f9319cc1b4df71eff3f46ba9e4631b43fd7c950680602085015260016040850152835109910152565b9251908115611cad576060928280920685520680602085015260016040850152835109910152565b9151815191602081019081518315611cad578380808093604098088180808a818080808c5180099c518009818d810382089c08810380988782980908980151800980088103870894828682098a520960608801528309602086015209910152565b90606080918051845260208101516020850152604081015160408501520151910152565b9091604082018051938415159485612484575b505083612461575b50505090565b519192506020915f5160206124c35f395f51905f529109910151145f808061245b565b84519295505f5160206124c35f395f51905f52910914925f8061245356fe25797203f7a0b24925572e1cd16bf9edfce0051fb9e133774b3c257a872d7d8b30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001060c89ce5c263405370a08b6d0302b0bab3eedb83920ee0a677297dc392126f11561ff836ce19d358a4eb7a4c199e94c377c749ae6f2a277f1f9195afe553f9fa2646970667358221220d923e67031a7cb3185aef932f47990abcdea3b2a539791a781bdd396b6bd9af464736f6c634300081c0033",
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
// Solidity: function getApplication(bytes12 epochId, bytes32 aid) view returns((address,(uint256,uint256),uint256,uint8,(uint8,bool,address[],uint16,uint64,uint64,uint64,uint64),uint64,bool))
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
// Solidity: function getApplication(bytes12 epochId, bytes32 aid) view returns((address,(uint256,uint256),uint256,uint8,(uint8,bool,address[],uint16,uint64,uint64,uint64,uint64),uint64,bool))
func (_DKGAppManager *DKGAppManagerSession) GetApplication(epochId [12]byte, aid [32]byte) (DKGTypesApplication, error) {
	return _DKGAppManager.Contract.GetApplication(&_DKGAppManager.CallOpts, epochId, aid)
}

// GetApplication is a free data retrieval call binding the contract method 0x2fed2529.
//
// Solidity: function getApplication(bytes12 epochId, bytes32 aid) view returns((address,(uint256,uint256),uint256,uint8,(uint8,bool,address[],uint16,uint64,uint64,uint64,uint64),uint64,bool))
func (_DKGAppManager *DKGAppManagerCallerSession) GetApplication(epochId [12]byte, aid [32]byte) (DKGTypesApplication, error) {
	return _DKGAppManager.Contract.GetApplication(&_DKGAppManager.CallOpts, epochId, aid)
}

// GetApplicationKey is a free data retrieval call binding the contract method 0x0fa5744c.
//
// Solidity: function getApplicationKey(bytes12 epochId, bytes32 aid) view returns(uint256 x, uint256 y)
func (_DKGAppManager *DKGAppManagerCaller) GetApplicationKey(opts *bind.CallOpts, epochId [12]byte, aid [32]byte) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _DKGAppManager.contract.Call(opts, &out, "getApplicationKey", epochId, aid)

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetApplicationKey is a free data retrieval call binding the contract method 0x0fa5744c.
//
// Solidity: function getApplicationKey(bytes12 epochId, bytes32 aid) view returns(uint256 x, uint256 y)
func (_DKGAppManager *DKGAppManagerSession) GetApplicationKey(epochId [12]byte, aid [32]byte) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _DKGAppManager.Contract.GetApplicationKey(&_DKGAppManager.CallOpts, epochId, aid)
}

// GetApplicationKey is a free data retrieval call binding the contract method 0x0fa5744c.
//
// Solidity: function getApplicationKey(bytes12 epochId, bytes32 aid) view returns(uint256 x, uint256 y)
func (_DKGAppManager *DKGAppManagerCallerSession) GetApplicationKey(epochId [12]byte, aid [32]byte) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _DKGAppManager.Contract.GetApplicationKey(&_DKGAppManager.CallOpts, epochId, aid)
}

// GetOrganizerPK is a free data retrieval call binding the contract method 0x81fe92fb.
//
// Solidity: function getOrganizerPK(bytes12 epochId, bytes32 aid) view returns(uint256, uint256)
func (_DKGAppManager *DKGAppManagerCaller) GetOrganizerPK(opts *bind.CallOpts, epochId [12]byte, aid [32]byte) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _DKGAppManager.contract.Call(opts, &out, "getOrganizerPK", epochId, aid)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetOrganizerPK is a free data retrieval call binding the contract method 0x81fe92fb.
//
// Solidity: function getOrganizerPK(bytes12 epochId, bytes32 aid) view returns(uint256, uint256)
func (_DKGAppManager *DKGAppManagerSession) GetOrganizerPK(epochId [12]byte, aid [32]byte) (*big.Int, *big.Int, error) {
	return _DKGAppManager.Contract.GetOrganizerPK(&_DKGAppManager.CallOpts, epochId, aid)
}

// GetOrganizerPK is a free data retrieval call binding the contract method 0x81fe92fb.
//
// Solidity: function getOrganizerPK(bytes12 epochId, bytes32 aid) view returns(uint256, uint256)
func (_DKGAppManager *DKGAppManagerCallerSession) GetOrganizerPK(epochId [12]byte, aid [32]byte) (*big.Int, *big.Int, error) {
	return _DKGAppManager.Contract.GetOrganizerPK(&_DKGAppManager.CallOpts, epochId, aid)
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

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_DKGAppManager *DKGAppManagerCaller) Registrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKGAppManager.contract.Call(opts, &out, "registrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_DKGAppManager *DKGAppManagerSession) Registrar() (common.Address, error) {
	return _DKGAppManager.Contract.Registrar(&_DKGAppManager.CallOpts)
}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_DKGAppManager *DKGAppManagerCallerSession) Registrar() (common.Address, error) {
	return _DKGAppManager.Contract.Registrar(&_DKGAppManager.CallOpts)
}

// RegistrarAdmin is a free data retrieval call binding the contract method 0x112fc40d.
//
// Solidity: function registrarAdmin() view returns(address)
func (_DKGAppManager *DKGAppManagerCaller) RegistrarAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKGAppManager.contract.Call(opts, &out, "registrarAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistrarAdmin is a free data retrieval call binding the contract method 0x112fc40d.
//
// Solidity: function registrarAdmin() view returns(address)
func (_DKGAppManager *DKGAppManagerSession) RegistrarAdmin() (common.Address, error) {
	return _DKGAppManager.Contract.RegistrarAdmin(&_DKGAppManager.CallOpts)
}

// RegistrarAdmin is a free data retrieval call binding the contract method 0x112fc40d.
//
// Solidity: function registrarAdmin() view returns(address)
func (_DKGAppManager *DKGAppManagerCallerSession) RegistrarAdmin() (common.Address, error) {
	return _DKGAppManager.Contract.RegistrarAdmin(&_DKGAppManager.CallOpts)
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

// RequireDecryptionOpen is a free data retrieval call binding the contract method 0xf6d651b4.
//
// Solidity: function requireDecryptionOpen(bytes12 epochId, bytes32 aid) view returns()
func (_DKGAppManager *DKGAppManagerCaller) RequireDecryptionOpen(opts *bind.CallOpts, epochId [12]byte, aid [32]byte) error {
	var out []interface{}
	err := _DKGAppManager.contract.Call(opts, &out, "requireDecryptionOpen", epochId, aid)

	if err != nil {
		return err
	}

	return err

}

// RequireDecryptionOpen is a free data retrieval call binding the contract method 0xf6d651b4.
//
// Solidity: function requireDecryptionOpen(bytes12 epochId, bytes32 aid) view returns()
func (_DKGAppManager *DKGAppManagerSession) RequireDecryptionOpen(epochId [12]byte, aid [32]byte) error {
	return _DKGAppManager.Contract.RequireDecryptionOpen(&_DKGAppManager.CallOpts, epochId, aid)
}

// RequireDecryptionOpen is a free data retrieval call binding the contract method 0xf6d651b4.
//
// Solidity: function requireDecryptionOpen(bytes12 epochId, bytes32 aid) view returns()
func (_DKGAppManager *DKGAppManagerCallerSession) RequireDecryptionOpen(epochId [12]byte, aid [32]byte) error {
	return _DKGAppManager.Contract.RequireDecryptionOpen(&_DKGAppManager.CallOpts, epochId, aid)
}

// RegisterApplication is a paid mutator transaction binding the contract method 0x6ed93d64.
//
// Solidity: function registerApplication(bytes12 epochId, bytes32 aid, (uint8,bool,address[],uint16,uint64,uint64,uint64,uint64) policy, uint256 pkOrgX, uint256 pkOrgY, uint256 schnorrAx, uint256 schnorrAy, uint256 schnorrZ) returns()
func (_DKGAppManager *DKGAppManagerTransactor) RegisterApplication(opts *bind.TransactOpts, epochId [12]byte, aid [32]byte, policy DKGTypesAppPolicy, pkOrgX *big.Int, pkOrgY *big.Int, schnorrAx *big.Int, schnorrAy *big.Int, schnorrZ *big.Int) (*types.Transaction, error) {
	return _DKGAppManager.contract.Transact(opts, "registerApplication", epochId, aid, policy, pkOrgX, pkOrgY, schnorrAx, schnorrAy, schnorrZ)
}

// RegisterApplication is a paid mutator transaction binding the contract method 0x6ed93d64.
//
// Solidity: function registerApplication(bytes12 epochId, bytes32 aid, (uint8,bool,address[],uint16,uint64,uint64,uint64,uint64) policy, uint256 pkOrgX, uint256 pkOrgY, uint256 schnorrAx, uint256 schnorrAy, uint256 schnorrZ) returns()
func (_DKGAppManager *DKGAppManagerSession) RegisterApplication(epochId [12]byte, aid [32]byte, policy DKGTypesAppPolicy, pkOrgX *big.Int, pkOrgY *big.Int, schnorrAx *big.Int, schnorrAy *big.Int, schnorrZ *big.Int) (*types.Transaction, error) {
	return _DKGAppManager.Contract.RegisterApplication(&_DKGAppManager.TransactOpts, epochId, aid, policy, pkOrgX, pkOrgY, schnorrAx, schnorrAy, schnorrZ)
}

// RegisterApplication is a paid mutator transaction binding the contract method 0x6ed93d64.
//
// Solidity: function registerApplication(bytes12 epochId, bytes32 aid, (uint8,bool,address[],uint16,uint64,uint64,uint64,uint64) policy, uint256 pkOrgX, uint256 pkOrgY, uint256 schnorrAx, uint256 schnorrAy, uint256 schnorrZ) returns()
func (_DKGAppManager *DKGAppManagerTransactorSession) RegisterApplication(epochId [12]byte, aid [32]byte, policy DKGTypesAppPolicy, pkOrgX *big.Int, pkOrgY *big.Int, schnorrAx *big.Int, schnorrAy *big.Int, schnorrZ *big.Int) (*types.Transaction, error) {
	return _DKGAppManager.Contract.RegisterApplication(&_DKGAppManager.TransactOpts, epochId, aid, policy, pkOrgX, pkOrgY, schnorrAx, schnorrAy, schnorrZ)
}

// RevealOrganizerSecret is a paid mutator transaction binding the contract method 0xa59b7a4d.
//
// Solidity: function revealOrganizerSecret(bytes12 epochId, bytes32 aid, uint256 organizerSecret) returns()
func (_DKGAppManager *DKGAppManagerTransactor) RevealOrganizerSecret(opts *bind.TransactOpts, epochId [12]byte, aid [32]byte, organizerSecret *big.Int) (*types.Transaction, error) {
	return _DKGAppManager.contract.Transact(opts, "revealOrganizerSecret", epochId, aid, organizerSecret)
}

// RevealOrganizerSecret is a paid mutator transaction binding the contract method 0xa59b7a4d.
//
// Solidity: function revealOrganizerSecret(bytes12 epochId, bytes32 aid, uint256 organizerSecret) returns()
func (_DKGAppManager *DKGAppManagerSession) RevealOrganizerSecret(epochId [12]byte, aid [32]byte, organizerSecret *big.Int) (*types.Transaction, error) {
	return _DKGAppManager.Contract.RevealOrganizerSecret(&_DKGAppManager.TransactOpts, epochId, aid, organizerSecret)
}

// RevealOrganizerSecret is a paid mutator transaction binding the contract method 0xa59b7a4d.
//
// Solidity: function revealOrganizerSecret(bytes12 epochId, bytes32 aid, uint256 organizerSecret) returns()
func (_DKGAppManager *DKGAppManagerTransactorSession) RevealOrganizerSecret(epochId [12]byte, aid [32]byte, organizerSecret *big.Int) (*types.Transaction, error) {
	return _DKGAppManager.Contract.RevealOrganizerSecret(&_DKGAppManager.TransactOpts, epochId, aid, organizerSecret)
}

// SetRegistrar is a paid mutator transaction binding the contract method 0xfaab9d39.
//
// Solidity: function setRegistrar(address r) returns()
func (_DKGAppManager *DKGAppManagerTransactor) SetRegistrar(opts *bind.TransactOpts, r common.Address) (*types.Transaction, error) {
	return _DKGAppManager.contract.Transact(opts, "setRegistrar", r)
}

// SetRegistrar is a paid mutator transaction binding the contract method 0xfaab9d39.
//
// Solidity: function setRegistrar(address r) returns()
func (_DKGAppManager *DKGAppManagerSession) SetRegistrar(r common.Address) (*types.Transaction, error) {
	return _DKGAppManager.Contract.SetRegistrar(&_DKGAppManager.TransactOpts, r)
}

// SetRegistrar is a paid mutator transaction binding the contract method 0xfaab9d39.
//
// Solidity: function setRegistrar(address r) returns()
func (_DKGAppManager *DKGAppManagerTransactorSession) SetRegistrar(r common.Address) (*types.Transaction, error) {
	return _DKGAppManager.Contract.SetRegistrar(&_DKGAppManager.TransactOpts, r)
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
	Mode         uint8
	PoolIndex    uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterApplicationRegistered is a free log retrieval operation binding the contract event 0x763f6464c291e9532f5ac2b75d76aa5283fc376d894a5e0b1e750ea0532cc7b1.
//
// Solidity: event ApplicationRegistered(bytes12 indexed epochId, bytes32 indexed aid, address indexed creator, uint256 organizerPKx, uint256 organizerPKy, uint8 mode, uint8 poolIndex)
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

// WatchApplicationRegistered is a free log subscription operation binding the contract event 0x763f6464c291e9532f5ac2b75d76aa5283fc376d894a5e0b1e750ea0532cc7b1.
//
// Solidity: event ApplicationRegistered(bytes12 indexed epochId, bytes32 indexed aid, address indexed creator, uint256 organizerPKx, uint256 organizerPKy, uint8 mode, uint8 poolIndex)
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

// ParseApplicationRegistered is a log parse operation binding the contract event 0x763f6464c291e9532f5ac2b75d76aa5283fc376d894a5e0b1e750ea0532cc7b1.
//
// Solidity: event ApplicationRegistered(bytes12 indexed epochId, bytes32 indexed aid, address indexed creator, uint256 organizerPKx, uint256 organizerPKy, uint8 mode, uint8 poolIndex)
func (_DKGAppManager *DKGAppManagerFilterer) ParseApplicationRegistered(log types.Log) (*DKGAppManagerApplicationRegistered, error) {
	event := new(DKGAppManagerApplicationRegistered)
	if err := _DKGAppManager.contract.UnpackLog(event, "ApplicationRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGAppManagerOrganizerSecretRevealedIterator is returned from FilterOrganizerSecretRevealed and is used to iterate over the raw logs and unpacked data for OrganizerSecretRevealed events raised by the DKGAppManager contract.
type DKGAppManagerOrganizerSecretRevealedIterator struct {
	Event *DKGAppManagerOrganizerSecretRevealed // Event containing the contract specifics and raw log

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
func (it *DKGAppManagerOrganizerSecretRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGAppManagerOrganizerSecretRevealed)
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
		it.Event = new(DKGAppManagerOrganizerSecretRevealed)
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
func (it *DKGAppManagerOrganizerSecretRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGAppManagerOrganizerSecretRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGAppManagerOrganizerSecretRevealed represents a OrganizerSecretRevealed event raised by the DKGAppManager contract.
type DKGAppManagerOrganizerSecretRevealed struct {
	EpochId         [12]byte
	Aid             [32]byte
	OrganizerSecret *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOrganizerSecretRevealed is a free log retrieval operation binding the contract event 0xa6d794f0981501d1b02889a9815c10a2bd90f6486541351fd1258d2ed3318677.
//
// Solidity: event OrganizerSecretRevealed(bytes12 indexed epochId, bytes32 indexed aid, uint256 organizerSecret)
func (_DKGAppManager *DKGAppManagerFilterer) FilterOrganizerSecretRevealed(opts *bind.FilterOpts, epochId [][12]byte, aid [][32]byte) (*DKGAppManagerOrganizerSecretRevealedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var aidRule []interface{}
	for _, aidItem := range aid {
		aidRule = append(aidRule, aidItem)
	}

	logs, sub, err := _DKGAppManager.contract.FilterLogs(opts, "OrganizerSecretRevealed", epochIdRule, aidRule)
	if err != nil {
		return nil, err
	}
	return &DKGAppManagerOrganizerSecretRevealedIterator{contract: _DKGAppManager.contract, event: "OrganizerSecretRevealed", logs: logs, sub: sub}, nil
}

// WatchOrganizerSecretRevealed is a free log subscription operation binding the contract event 0xa6d794f0981501d1b02889a9815c10a2bd90f6486541351fd1258d2ed3318677.
//
// Solidity: event OrganizerSecretRevealed(bytes12 indexed epochId, bytes32 indexed aid, uint256 organizerSecret)
func (_DKGAppManager *DKGAppManagerFilterer) WatchOrganizerSecretRevealed(opts *bind.WatchOpts, sink chan<- *DKGAppManagerOrganizerSecretRevealed, epochId [][12]byte, aid [][32]byte) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var aidRule []interface{}
	for _, aidItem := range aid {
		aidRule = append(aidRule, aidItem)
	}

	logs, sub, err := _DKGAppManager.contract.WatchLogs(opts, "OrganizerSecretRevealed", epochIdRule, aidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGAppManagerOrganizerSecretRevealed)
				if err := _DKGAppManager.contract.UnpackLog(event, "OrganizerSecretRevealed", log); err != nil {
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

// ParseOrganizerSecretRevealed is a log parse operation binding the contract event 0xa6d794f0981501d1b02889a9815c10a2bd90f6486541351fd1258d2ed3318677.
//
// Solidity: event OrganizerSecretRevealed(bytes12 indexed epochId, bytes32 indexed aid, uint256 organizerSecret)
func (_DKGAppManager *DKGAppManagerFilterer) ParseOrganizerSecretRevealed(log types.Log) (*DKGAppManagerOrganizerSecretRevealed, error) {
	event := new(DKGAppManagerOrganizerSecretRevealed)
	if err := _DKGAppManager.contract.UnpackLog(event, "OrganizerSecretRevealed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGAppManagerRegistrarSetIterator is returned from FilterRegistrarSet and is used to iterate over the raw logs and unpacked data for RegistrarSet events raised by the DKGAppManager contract.
type DKGAppManagerRegistrarSetIterator struct {
	Event *DKGAppManagerRegistrarSet // Event containing the contract specifics and raw log

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
func (it *DKGAppManagerRegistrarSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGAppManagerRegistrarSet)
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
		it.Event = new(DKGAppManagerRegistrarSet)
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
func (it *DKGAppManagerRegistrarSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGAppManagerRegistrarSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGAppManagerRegistrarSet represents a RegistrarSet event raised by the DKGAppManager contract.
type DKGAppManagerRegistrarSet struct {
	Registrar common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRegistrarSet is a free log retrieval operation binding the contract event 0x6263309d5d4d1cfececd45a387cda7f14dccde21cf7a1bee1be6561075e61014.
//
// Solidity: event RegistrarSet(address registrar)
func (_DKGAppManager *DKGAppManagerFilterer) FilterRegistrarSet(opts *bind.FilterOpts) (*DKGAppManagerRegistrarSetIterator, error) {

	logs, sub, err := _DKGAppManager.contract.FilterLogs(opts, "RegistrarSet")
	if err != nil {
		return nil, err
	}
	return &DKGAppManagerRegistrarSetIterator{contract: _DKGAppManager.contract, event: "RegistrarSet", logs: logs, sub: sub}, nil
}

// WatchRegistrarSet is a free log subscription operation binding the contract event 0x6263309d5d4d1cfececd45a387cda7f14dccde21cf7a1bee1be6561075e61014.
//
// Solidity: event RegistrarSet(address registrar)
func (_DKGAppManager *DKGAppManagerFilterer) WatchRegistrarSet(opts *bind.WatchOpts, sink chan<- *DKGAppManagerRegistrarSet) (event.Subscription, error) {

	logs, sub, err := _DKGAppManager.contract.WatchLogs(opts, "RegistrarSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGAppManagerRegistrarSet)
				if err := _DKGAppManager.contract.UnpackLog(event, "RegistrarSet", log); err != nil {
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

// ParseRegistrarSet is a log parse operation binding the contract event 0x6263309d5d4d1cfececd45a387cda7f14dccde21cf7a1bee1be6561075e61014.
//
// Solidity: event RegistrarSet(address registrar)
func (_DKGAppManager *DKGAppManagerFilterer) ParseRegistrarSet(log types.Log) (*DKGAppManagerRegistrarSet, error) {
	event := new(DKGAppManagerRegistrarSet)
	if err := _DKGAppManager.contract.UnpackLog(event, "RegistrarSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
