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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MANAGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApplication\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Application\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"organizerPK\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"organizerSecret\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"poolIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.AppPolicy\",\"components\":[{\"name\":\"mode\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.AppMode\"},{\"name\":\"openSubmission\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"submitters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"maxCiphertexts\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"decryptNotBefore\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"decryptNotAfter\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"createdAtBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApplicationKey\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrganizerPK\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredAids\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerApplication\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.AppPolicy\",\"components\":[{\"name\":\"mode\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.AppMode\"},{\"name\":\"openSubmission\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"submitters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"maxCiphertexts\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"decryptNotBefore\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"decryptNotAfter\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"pkOrgX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pkOrgY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrAx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrAy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrZ\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requireCanSubmitCiphertext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireDecryptionOpen\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revealOrganizerSecret\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"organizerSecret\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ApplicationRegistered\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"organizerPKx\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"organizerPKy\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"mode\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumDKGTypes.AppMode\"},{\"name\":\"poolIndex\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrganizerSecretRevealed\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"organizerSecret\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyRevealed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ApplicationAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionClosed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionLimitReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionNotOpen\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionNotYetAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidApplication\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEpoch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOrganizerSecret\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSchnorrProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IsIdentity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotCanonical\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOnCurve\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrganizerSecretNotRevealed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PointNotInSubgroup\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolExhausted\",\"inputs\":[]}]",
	Bin: "0x60a03461008d57601f61245f38819003918201601f19168301916001600160401b038311848410176100915780849260209460405283398101031261008d57516001600160a01b03811680820361008d571561007e576080526040516123b990816100a682396080518181816101400152818161025001526110bb0152f35b63e6c4247b60e01b5f5260045ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe60806040526004361015610011575f80fd5b5f3560e01c80630fa5744c146100a45780631b2df8501461009f5780632fed25291461009a5780636ed93d641461009557806374a99aba1461009057806381fe92fb1461008b578063a59b7a4d14610086578063ed78d71e146100815763f6d651b41461007c575f80fd5b610905565b610886565b61073c565b6106e7565b610542565b6104c0565b610410565b61023b565b34610220576040366003190112610220576100bd610224565b6100d96024356100cc836109db565b905f5260205260405f2090565b906100f36100ef600984015460ff9060401c1690565b1590565b61021157604061013c9161010b600485015460ff1690565b82516356cbb5f360e01b81526001600160a01b0319909216600483015260ff16602482015291829081906044820190565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561020c575f905f926101da575b508192819261018d600583015460ff1690565b61019681610293565b156101b8575b50506040805192835260208301939093525090819081015b0390f35b909192506101d193506002600183015492015492611506565b9082808061019c565b90506101fe915060403d604011610205575b6101f68183610a7b565b810190610a9e565b908361017a565b503d6101ec565b610ab4565b6378e9323b60e11b5f5260045ffd5b5f80fd5b600435906001600160a01b03198216820361022057565b34610220575f366003190112610220576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b634e487b7160e01b5f52602160045260245ffd5b6002111561029d57565b61027f565b90600282101561029d5752565b6001600160401b031690565b6001600160401b03169052565b906101008101916102da8282516102a2565b60208101511515602083015260408101519261010060408401528351809152602061012084019401905f5b81811061036f575050509060e0808361032c606061036c960151606086019061ffff169052565b61033e608082015160808601906102bb565b61035060a082015160a08601906102bb565b61036260c082015160c08601906102bb565b01519101906102bb565b90565b82516001600160a01b0316865260209586019590920191600101610305565b61036c906020815260018060a01b0383511660208201526020808401518051604084015201516060820152604083015160808201526103d7606084015160a083019060ff169052565b61010060c06103f4608086015183838601526101208501906102c8565b9461040760a082015160e08601906102bb565b01511515910152565b34610220576040366003190112610220576101b46104b46104af610432610224565b6100cc602435915f60c060405161044881610a09565b828152610453610aef565b602082015282604082015282606082015260405161047081610a29565b838152836020820152606060408201528360608201528360808201528360a082015283838201528360e082015260808201528260a082015201526109db565b610c5c565b6040519182918261038e565b3461022057610100366003190112610220576104da610224565b6044356024356001600160401b03821161022057610100600319833603011261022057610523926064356084359060a4359260c4359461051960e43590565b966004019161108b565b005b61ffff81160361022057565b6001600160a01b0381160361022057565b346102205760803660031901126102205761055b610224565b60443561058260243561056d83610525565b6100cc6064359461057d86610531565b6109db565b916105986100ef600985015460ff9060401c1690565b6102115760058301546105af9060081c60ff161590565b90816106d2575b506106c35760078201546001600160401b03601082901c1680151590816106a8575b50610699576105f8605082901c6001600160401b03166102af565b6102af565b8015159081610686575b506106775761ffff16801515919082610665575b5050610656576105f3600861062c9201546102af565b801515908161064c575b5061063d57005b632300418b60e01b5f5260045ffd5b905042115f610636565b63464e67af60e01b5f5260045ffd5b61ffff90811691161190505f80610616565b630410ff2960e31b5f5260045ffd5b436001600160401b03161190505f610602565b633deac39560e01b5f5260045ffd5b6106b291506102af565b436001600160401b0316105f6105d8565b6330cd747160e01b5f5260045ffd5b6106e191506100ef9084611947565b5f6105b6565b3461022057604036600319011261022057610700610224565b6001600160a01b0319165f9081526020818152604080832060243584528252918290206001810154600290910154835191825291810191909152f35b3461022057606036600319011261022057610755610224565b602435906044359061076a836100cc836109db565b600981015461077d9060401c60ff161590565b61021157600581015460ff1661079281610293565b61083e576003810190815461083e5783158015610827575b610809576107b784611fcf565b9060018301541491821592610818575b5050610809578290556040519182526001600160a01b031916907fa6d794f0981501d1b02889a9815c10a2bd90f6486541351fd1258d2ed331867790602090a3005b634102642560e11b5f5260045ffd5b60020154141590505f806107c7565b505f5160206123445f395f51905f528410156107aa565b63a89ac15160e01b5f5260045ffd5b60206040818301928281528451809452019201905f5b8181106108705750505090565b8251845260209384019390920191600101610863565b34610220576020366003190112610220576001600160a01b03196108a8610224565b165f52600160205260405f206040519081602082549182815201915f5260205f20905f5b8181106108ef576101b4856108e381870382610a7b565b6040519182918261084d565b82548452602090930192600192830192016108cc565b346102205760403660031901126102205761092d610921610224565b6100cc602435916109db565b60098101546109409060401c60ff161590565b61021157600581015460ff1661095581610293565b15806109cf575b6109c057600781015460901c6001600160401b031680151590816109b6575b506109a757600801546001600160401b0390610996906102af565b16801515908161064c575061063d57005b6314badd7360e31b5f5260045ffd5b905042105f61097b565b630a4bc8e760e41b5f5260045ffd5b5060038101541561095c565b6001600160a01b0319165f90815260208190526040902090565b634e487b7160e01b5f52604160045260245ffd5b60e081019081106001600160401b03821117610a2457604052565b6109f5565b61010081019081106001600160401b03821117610a2457604052565b604081019081106001600160401b03821117610a2457604052565b608081019081106001600160401b03821117610a2457604052565b601f909101601f19168101906001600160401b03821190821017610a2457604052565b9190826040910312610220576020825192015190565b6040513d5f823e3d90fd5b60405190610acf61018083610a7b565b565b60405190610acf604083610a7b565b60405190610acf608083610a7b565b60405190610afc82610a45565b5f6020838281520152565b90604051610b1481610a45565b602060018294805484520154910152565b90604051918281549182825260208201905f5260205f20925f5b818110610b54575050610acf92500383610a7b565b84546001600160a01b0316835260019485019487945060209093019201610b3f565b6001600160401b039091169052565b9060405191610b9383610a29565b8281549160ff831692600284101561029d57600360e092610bcd610bc4610c5594610acf98885260ff9060081c1690565b15156020870152565b610bd960018201610b25565b6040860152610c4e610c456002830154610c02610bf78261ffff1690565b61ffff1660608a0152565b610c1c601082901c6001600160401b031660808a01610b76565b610c36605082901c6001600160401b031660a08a01610b76565b60901c6001600160401b031690565b60c08701610b76565b01546102af565b9101610b76565b90610acf604051610c6c81610a09565b83546001600160a01b0316815292839060c090610ced90600990610c9260018201610b07565b602086015260038101546040860152610cbc610cb2600483015460ff1690565b60ff166060870152565b610cc860058201610b85565b60808601520154610ce4610cdb826102af565b60a08601610b76565b60401c60ff1690565b1515910152565b5190610acf82610531565b5190610acf82610525565b6001600160401b0381160361022057565b5190610acf82610d0a565b91908260e091031261022057604051610d3e81610a09565b60c0610db08183958051610d5181610525565b85526020810151610d6181610525565b6020860152610d7260408201610cff565b6040860152610d8360608201610cff565b6060860152610d9460808201610d1b565b6080860152610da560a08201610d1b565b60a086015201610d1b565b910152565b5190600682101561022057565b6102408183031261022057610220610e9691610df4610ddf610abf565b94610de983610cf4565b865260208301610d26565b6020850152610e066101008201610db5565b6040850152610e186101208201610d1b565b6060850152610e2a6101408201610d1b565b6080850152610e3c6101608201610d1b565b60a085015261018081015160c08501526101a081015160e0850152610e646101c08201610cff565b610100850152610e776101e08201610cff565b610120850152610e8a6102008201610cff565b61014085015201610cff565b61016082015290565b6006111561029d57565b3560028110156102205790565b90816020910312610220575160ff811681036102205790565b90600281101561029d5760ff80198354169116179055565b3580151581036102205790565b903590601e198136030182121561022057018035906001600160401b03821161022057602001918160051b3603831361022057565b634e487b7160e01b5f52601160045260245ffd5b3561036c81610531565b906001600160401b038311610a2457600160401b8311610a24578154838355808410610fa7575b50610f7d90915f5260205f2090565b5f5b838110610f8c5750505050565b6001906020610f9a85610f3d565b9401938184015501610f7f565b825f528360205f2091820191015b818110610fc25750610f6e565b5f8155600101610fb5565b3561036c81610525565b3561036c81610d0a565b80546001600160401b0319166001600160401b03909216919091179055565b634e487b7160e01b5f52603260045260245ffd5b8054821015611029575f5260205f2001905f90565b611000565b8054600160401b811015610a245761104b91600182018155611014565b819291549060031b91821b915f19901b1916179055565b929361108560ff92969560609460808701988752602087015260408601906102a2565b16910152565b604051635f2cdc7560e11b81526001600160a01b0319821660048201529097919694959491936001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001693929161024081602481885afa90811561020c575f916114d7575b5080516001600160a01b0316156114c8576040600391015161111781610e9f565b61112081610e9f565b036114b95761112e8861157a565b61113b886100cc8b6109db565b946009860192611150845460ff9060401c1690565b6114aa5761115d856115b3565b5f97600161116a87610ea9565b61117381610293565b0361142f57505050506111c09495506020876001945b60405163421adfbb60e01b81526001600160a01b03198c1660048201526024810192909252909687919082905f9082906044820190565b03925af193841561020c577f763f6464c291e9532f5ac2b75d76aa5283fc376d894a5e0b1e750ea0532cc7b1955f956113ea575b5080546001600160a01b031916331781556113e5926113c89290916113a0919061137a9061123c611223610ad1565b8c8152602001899052600182018c905560028201899055565b60048101805460ff191660ff8b161790556112886005820161126661126088610ea9565b82610ecf565b61127260208801610ee7565b815461ff00191690151560081b61ff0016179055565b6112a26112986040870187610ef4565b9060068401610f47565b6113656112b160608701610fcd565b6112cb6007840191829061ffff1661ffff19825416179055565b6112fe6112da60808901610fd7565b825462010000600160501b03191660109190911b62010000600160501b0316178255565b61133361130d60a08901610fd7565b8254600160501b600160901b03191660509190911b600160501b600160901b0316178255565b61133f60c08801610fd7565b8154600160901b600160d01b03191660909190911b600160901b600160d01b0316179055565b600861137360e08701610fd7565b9101610fe1565b61138d436001600160401b031682610fe1565b805460ff60401b1916600160401b179055565b6001600160a01b031989165f9081526001602052604090206113c390899061102e565b610ea9565b60405133986001600160a01b031916969094859490929185611062565b0390a4565b6113e5939195506113c89261137a61141c6113a09360203d602011611428575b6114148183610a7b565b810190610eb6565b979395505092506111f4565b503d61140a565b909192975061143e8989611757565b61144b6100ef8a8a6117dc565b61149b5787828261147a956114748f968f968f8781996100ef9b61146f8686611757565b611d19565b90611dc1565b61148c576020876111c0969794611189565b6327f7eb4d60e11b5f5260045ffd5b63b28e789160e01b5f5260045ffd5b630b792c8f60e01b5f5260045ffd5b63268dbf6760e21b5f5260045ffd5b63d5b25b6360e01b5f5260045ffd5b6114f991506102403d81116114ff575b6114f18183610a7b565b810190610dc2565b5f6110f6565b503d6114e7565b909392811580611570575b6115695782158061155f575b611558579361154861154f9261155495966115426115396119b6565b80958194611a28565b95611a28565b8380611a73565b611b13565b9091565b5090509190565b506001811461151d565b9350509190565b5060018514611511565b801590811561158b575b5061021157565b5f5160206123245f395f51905f52915010155f611584565b91908110156110295760051b0190565b604081016115c18183610ef4565b91905060208211801561173b575b61164f575f5b8281106117035750505060e081016001600160401b036115f482610fd7565b161515806116e6575b61164f5760c08201906116126105f383610fd7565b151591826116d0575b826116a5575b505061164f5760808101906116386105f383610fd7565b1515918261168c575b8261165e575b505061164f57565b63d06b96b160e01b5f5260045ffd5b61167991925060a06116726105f392610fd7565b9301610fd7565b6001600160401b03909116115f80611647565b915061169d6105f360a08301610fd7565b151591611641565b6116bd9192506116b76105f391610fd7565b92610fd7565b6001600160401b03909116115f80611621565b91506116de6105f382610fd7565b15159161161b565b506116f081610fd7565b426001600160401b0390911611156115fd565b61172e61172261171d836117178689610ef4565b906115a3565b610f3d565b6001600160a01b031690565b1561164f576001016115d5565b5061174860208401610ee7565b80156115cf57508115156115cf565b905f5160206123245f395f51905f528210806117c6575b156117b7578115806117ad575b61179e5761178891611bf1565b1561178f57565b6361586bdd60e01b5f5260045ffd5b632b39517d60e21b5f5260045ffd5b506001811461177b565b63d7c7beeb60e01b5f5260045ffd5b505f5160206123245f395f51905f52811061176e565b80158061193d575b611936576117f28282611bf1565b15611930576117ff611a04565b5061184a61180b6119b6565b809261182782611819611cc1565b96602088015b9384516121f4565b604085019061183a83835183519061221c565b606086015b519151905191611a73565b611852610ae0565b5f81526020810192600184526040820192600184525f60608401525f91610100805b6118a3575050505051159182611897575b508161188f575090565b905051151590565b5181511491505f611885565b60011901805f5160206123445f395f51905f52811c6003168515611901576118cc85888061221c565b6118d785888061221c565b84816118e5575b5050611874565b6118f26118fa9286611cf5565b518880611a73565b5f846118de565b809150611910575b5080611874565b6119279194506119209083611cf5565b518561227d565b82600193611909565b50505f90565b5050600190565b50600182146117e4565b9060068201805492831561199e57505f5b83811061196757505050505f90565b6119718183611014565b905460039190911b1c6001600160a01b039081169084161461199557600101611958565b50505050600190565b546001600160a01b0392831692169190911492915050565b6119be610aef565b506040516119cb81610a45565b5f5160206123245f395f51905f5281527f0578d36fdd1172a8c3909ff8b278cb9adf026a6b5db6203e5d099f85f9afd71b602082015290565b60405190611a1182610a60565b5f6060838281528260208201528260408201520152565b9291610acf91611a36611a04565b5060405194611a4486610a60565b5f86525f60208701525f60408701525f6060870152856121f4565b634e487b7160e01b5f52601260045260245ffd5b92908151926020820192835183518603918615611b0e5786808086818080999881808d81809d9c816020819f01968188518c51820390089208099f519051900891518551900890099a818181038d089b089687958160608c0151606085015190099060200151900998604001519060400151900980089581818103880896089582868209895209606087015283096020850152099060400152565b611a5f565b90604082015160405190602082526020808301526020604083015260608201527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593efffffff60808201525f5160206123245f395f51905f5260a082015260208160c08160055afa15610220575f5160206123245f395f51905f529051602082828651099401510990565b5f5160206123245f395f51905f52119081611bb2575090565b5f5160206123245f395f51905f5291501090565b5f5160206123245f395f51905f5203905f5160206123245f395f51905f528211611bec57565b610f29565b5f5160206123245f395f51905f528110801590611c87575b611930575f5160206123245f395f51905f528181920991800990805f5160206123245f395f51905f5203915f5160206123245f395f51905f528311611bec575f5160206123245f395f51905f528080838195097f1aee90f15f2189693df072d799fd11fc039b2959ebb7c867d075ca8cf4d7eb8e0960010892081490565b505f5160206123245f395f51905f52821015611c09565b5f5b828110611cac57505050565b602090611cb7611a04565b8184015201611ca0565b60405190611cd0608083610a7b565b610acf608083611c9e565b60405190610acf610200611cef8185610a7b565b83611c9e565b9060048110156110295760051b0190565b5f5160206123445f395f51905f52900690565b9390925f5160206123445f395f51905f5295926040519460208601967f41ea6f3fa95eccd1f3b1ce8e05efa11027280aa0c6b4167fd6695db659c30b28885260018060a01b0319166040870152604c860152606c850152608c84015260ac83015260cc82015260cc8152611d8e60ec82610a7b565b5190200690565b9060108110156110295760051b0190565b9060048201809211611bec57565b91908201809211611bec57565b959491929395611dd46100ef8686611b99565b8015611fbc575b611fb257611deb611df191611d06565b91611d06565b9580611fa457505f5b611e7b611e056119b6565b8092611e5f82611e13611cdb565b97600160208a510152600160408a510152611e3882611e328b60800190565b516120e2565b611e47826101008b0151612124565b611e56826101808b015161218c565b6020890161181f565b6040860190611e7283835183519061221c565b6060870161183f565b60045b600c811115611f4d5750611e90610ae0565b925f845260016020850152600160408501525f60608501525f9260fc805b611ec257505050505061036c9394506122a1565b600119019889600c83821c60021b1682821c600316178615611f2157611ee986898061221c565b611ef486898061221c565b8581611f05575b50505b9099611eae565b611f12611f1a9287611d95565b518980611a73565b5f85611efb565b80611f2d575b50611efe565b611f44919650611f3d9085611d95565b518761227d565b6001945f611f27565b60015b60048110611f675750611f6290611da6565b611e7e565b80611f9e84611f81611f7b60019587611db4565b89611d95565b51611f8c868a611d95565b51611f97858b611d95565b5191611a73565b01611f50565b611fad90611bc6565b611dfa565b505f955050505050565b50611fca6100ef8489611b99565b611ddb565b5f5160206123445f395f51905f5290069081156120da57611fee611a04565b50611ff76119b6565b611fff611cc1565b9061202f8160208401611827825f5160206123045f395f51905f525f5160206123645f395f51905f5284516121f4565b612037610ae0565b915f835260016020840152600160408401525f60608401525f94610100805b6120695750505050611554919250611b13565b600119018082811c60031688156120b25761208586888061221c565b61209086888061221c565b858161209e575b5050612056565b6118f26120ab9287611cf5565b5f85612097565b8091506120c1575b5080612056565b6120d19197506119209084611cf5565b856001966120ba565b5f9150600190565b90518015611b0e57806060915f5160206123645f395f51905f52068352805f5160206123045f395f51905f520680602085015260016040850152835109910152565b90518015611b0e57806060917f0daaa7e6b25c28e6dc8dd1d48e9cc61cd07015c1d7c1b8d4590eb6f51d5346dc068352807f01666cafbf0a30da8b9ebeaf848a1da067a892296f1043188e1705402b6d68530680602085015260016040850152835109910152565b90518015611b0e57806060917f136d609c4c856f5d277fab08c730cbdd1a776ce4728c6a2eb20ff22bccf26894068352807f21d66f0e2295ae954494f25889f9319cc1b4df71eff3f46ba9e4631b43fd7c950680602085015260016040850152835109910152565b9251908115611b0e576060928280920685520680602085015260016040850152835109910152565b9151815191602081019081518315611b0e578380808093604098088180808a818080808c5180099c518009818d810382089c08810380988782980908980151800980088103870894828682098a520960608801528309602086015209910152565b90606080918051845260208101516020850152604081015160408501520151910152565b90916040820180519384151594856122e5575b5050836122c2575b50505090565b519192506020915f5160206123245f395f51905f529109910151145f80806122bc565b84519295505f5160206123245f395f51905f52910914925f806122b456fe25797203f7a0b24925572e1cd16bf9edfce0051fb9e133774b3c257a872d7d8b30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001060c89ce5c263405370a08b6d0302b0bab3eedb83920ee0a677297dc392126f11561ff836ce19d358a4eb7a4c199e94c377c749ae6f2a277f1f9195afe553f9fa26469706673582212205a09bda7bd3fcbb43be2b37d5b5cc9ab6ef50082a98e10a2d42ea24e128977d964736f6c634300081c0033",
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
