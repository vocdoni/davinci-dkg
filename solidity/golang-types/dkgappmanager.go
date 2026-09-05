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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MANAGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApplication\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Application\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"organizerPK\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"organizerSecret\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"poolIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.AppPolicy\",\"components\":[{\"name\":\"mode\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.AppMode\"},{\"name\":\"openSubmission\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"submitters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"maxCiphertexts\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"decryptNotBefore\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"decryptNotAfter\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"createdAtBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrganizerPK\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredAids\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerApplication\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.AppPolicy\",\"components\":[{\"name\":\"mode\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.AppMode\"},{\"name\":\"openSubmission\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"submitters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"maxCiphertexts\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"decryptNotBefore\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"decryptNotAfter\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"pkOrgX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pkOrgY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrAx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrAy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrZ\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requireCanSubmitCiphertext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireDecryptionOpen\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revealOrganizerSecret\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"organizerSecret\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ApplicationRegistered\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"organizerPKx\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"organizerPKy\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"mode\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumDKGTypes.AppMode\"},{\"name\":\"poolIndex\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrganizerSecretRevealed\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"organizerSecret\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyRevealed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ApplicationAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionClosed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionLimitReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionNotOpen\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionNotYetAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidApplication\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEpoch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOrganizerSecret\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSchnorrProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IsIdentity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotCanonical\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOnCurve\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrganizerSecretNotRevealed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PointNotInSubgroup\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolExhausted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolKeyNotActive\",\"inputs\":[]}]",
	Bin: "0x60a034608057601f61224f38819003918201601f19168301916001600160401b03831184841017608457808492602094604052833981010312608057516001600160a01b038116808203608057156071576080526040516121b69081610099823960805181818160a60152610f1e0152f35b63e6c4247b60e01b5f5260045ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe60806040526004361015610011575f80fd5b5f3560e01c80631b2df850146100945780632fed25291461008f5780636ed93d641461008a57806374a99aba1461008557806381fe92fb14610080578063a59b7a4d1461007b578063ed78d71e146100765763f6d651b414610071575f80fd5b610799565b61071a565b6105d0565b61057f565b6103c7565b610345565b610284565b346100d8575f3660031901126100d8577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166080908152602090f35b5f80fd5b600435906001600160a01b0319821682036100d857565b634e487b7160e01b5f52602160045260245ffd5b6002111561011157565b6100f3565b9060028210156101115752565b6001600160401b031690565b6001600160401b03169052565b9061010081019161014e828251610116565b60208101511515602083015260408101519261010060408401528351809152602061012084019401905f5b8181106101e3575050509060e080836101a060606101e0960151606086019061ffff169052565b6101b26080820151608086019061012f565b6101c460a082015160a086019061012f565b6101d660c082015160c086019061012f565b015191019061012f565b90565b82516001600160a01b0316865260209586019590920191600101610179565b6101e0906020815260018060a01b03835116602082015260208084015180516040840152015160608201526040830151608082015261024b606084015160a083019060ff169052565b61010060c06102686080860151838386015261012085019061013c565b9461027b60a082015160e086019061012f565b01511515910152565b346100d85760403660031901126100d8576103416103356103306102a66100dc565b610323602435915f60c06040516102bc81610883565b8281526102c761092d565b60208201528260408201528260608201526040516102e4816108a3565b838152836020820152606060408201528360608201528360808201528360a082015283838201528360e082015260808201528260a08201520152610945565b905f5260205260405f2090565b610ab4565b60405191829182610202565b0390f35b346100d8576101003660031901126100d85761035f6100dc565b6044356024356001600160401b0382116100d85761010060031983360301126100d8576103a8926064356084359060a4359260c4359461039e60e43590565b9660040191610eee565b005b61ffff8116036100d857565b6001600160a01b038116036100d857565b346100d85760803660031901126100d8576103e06100dc565b6044356104076024356103f2836103aa565b61032360643594610402866103b6565b610945565b9161042161041d600985015460ff9060401c1690565b1590565b6105705760058301546104389060081c60ff161590565b908161055b575b5061054c5760078201546001600160401b03601082901c168015159081610531575b5061052257610481605082901c6001600160401b0316610123565b610123565b801515908161050f575b506105005761ffff168015159190826104ee575b50506104df5761047c60086104b5920154610123565b80151590816104d5575b506104c657005b632300418b60e01b5f5260045ffd5b905042115f6104bf565b63464e67af60e01b5f5260045ffd5b61ffff90811691161190505f8061049f565b630410ff2960e31b5f5260045ffd5b436001600160401b03161190505f61048b565b633deac39560e01b5f5260045ffd5b61053b9150610123565b436001600160401b0316105f610461565b6330cd747160e01b5f5260045ffd5b61056a915061041d908461173b565b5f61043f565b6378e9323b60e11b5f5260045ffd5b346100d85760403660031901126100d8576105986100dc565b6024359060018060a01b0319165f525f60205260405f20905f526020526040805f206002600182015491015482519182526020820152f35b346100d85760603660031901126100d8576105e96100dc565b60243590604435906105fe8361032383610945565b60098101546106119060401c60ff161590565b61057057600581015460ff1661062681610107565b6106d257600381019081546106d257831580156106bb575b61069d5761064b84611d28565b90600183015414918215926106ac575b505061069d578290556040519182526001600160a01b031916907fa6d794f0981501d1b02889a9815c10a2bd90f6486541351fd1258d2ed331867790602090a3005b634102642560e11b5f5260045ffd5b60020154141590505f8061065b565b505f5160206121415f395f51905f5284101561063e565b63a89ac15160e01b5f5260045ffd5b60206040818301928281528451809452019201905f5b8181106107045750505090565b82518452602093840193909201916001016106f7565b346100d85760203660031901126100d8576001600160a01b031961073c6100dc565b165f52600160205260405f206040519081602082549182815201915f5260205f20905f5b8181106107835761034185610777818703826108da565b604051918291826106e1565b8254845260209093019260019283019201610760565b346100d85760403660031901126100d8576107c16107b56100dc565b61032360243591610945565b60098101546107d49060401c60ff161590565b61057057600581015460ff166107e981610107565b1580610863575b61085457600781015460901c6001600160401b0316801515908161084a575b5061083b57600801546001600160401b039061082a90610123565b1680151590816104d557506104c657005b6314badd7360e31b5f5260045ffd5b905042105f61080f565b630a4bc8e760e41b5f5260045ffd5b506003810154156107f0565b634e487b7160e01b5f52604160045260245ffd5b60e081019081106001600160401b0382111761089e57604052565b61086f565b61010081019081106001600160401b0382111761089e57604052565b604081019081106001600160401b0382111761089e57604052565b601f909101601f19168101906001600160401b0382119082101761089e57604052565b6040519061090d610180836108da565b565b6040519061090d6040836108da565b6040519061090d6080836108da565b6040519061093a826108bf565b5f6020838281520152565b6001600160a01b0319165f90815260208190526040902090565b9060405161096c816108bf565b602060018294805484520154910152565b90604051918281549182825260208201905f5260205f20925f5b8181106109ac57505061090d925003836108da565b84546001600160a01b0316835260019485019487945060209093019201610997565b6001600160401b039091169052565b90604051916109eb836108a3565b8281549160ff831692600284101561011157600360e092610a25610a1c610aad9461090d98885260ff9060081c1690565b15156020870152565b610a316001820161097d565b6040860152610aa6610a9d6002830154610a5a610a4f8261ffff1690565b61ffff1660608a0152565b610a74601082901c6001600160401b031660808a016109ce565b610a8e605082901c6001600160401b031660a08a016109ce565b60901c6001600160401b031690565b60c087016109ce565b0154610123565b91016109ce565b9061090d604051610ac481610883565b83546001600160a01b0316815292839060c090610b4590600990610aea6001820161095f565b602086015260038101546040860152610b14610b0a600483015460ff1690565b60ff166060870152565b610b20600582016109dd565b60808601520154610b3c610b3382610123565b60a086016109ce565b60401c60ff1690565b1515910152565b519061090d826103b6565b519061090d826103aa565b6001600160401b038116036100d857565b519061090d82610b62565b91908260e09103126100d857604051610b9681610883565b60c0610c088183958051610ba9816103aa565b85526020810151610bb9816103aa565b6020860152610bca60408201610b57565b6040860152610bdb60608201610b57565b6060860152610bec60808201610b73565b6080860152610bfd60a08201610b73565b60a086015201610b73565b910152565b519060068210156100d857565b610240818303126100d857610220610cee91610c4c610c376108fd565b94610c4183610b4c565b865260208301610b7e565b6020850152610c5e6101008201610c0d565b6040850152610c706101208201610b73565b6060850152610c826101408201610b73565b6080850152610c946101608201610b73565b60a085015261018081015160c08501526101a081015160e0850152610cbc6101c08201610b57565b610100850152610ccf6101e08201610b57565b610120850152610ce26102008201610b57565b61014085015201610b57565b61016082015290565b6040513d5f823e3d90fd5b6006111561011157565b3560028110156100d85790565b908160209103126100d8575160ff811681036100d85790565b9060028110156101115760ff80198354169116179055565b3580151581036100d85790565b903590601e19813603018212156100d857018035906001600160401b0382116100d857602001918160051b360383136100d857565b634e487b7160e01b5f52601160045260245ffd5b356101e0816103b6565b906001600160401b03831161089e57600160401b831161089e578154838355808410610e0a575b50610de090915f5260205f2090565b5f5b838110610def5750505050565b6001906020610dfd85610da0565b9401938184015501610de2565b825f528360205f2091820191015b818110610e255750610dd1565b5f8155600101610e18565b356101e0816103aa565b356101e081610b62565b80546001600160401b0319166001600160401b03909216919091179055565b634e487b7160e01b5f52603260045260245ffd5b8054821015610e8c575f5260205f2001905f90565b610e63565b8054600160401b81101561089e57610eae91600182018155610e77565b819291549060031b91821b915f19901b1916179055565b9293610ee860ff9296956060946080870198875260208701526040860190610116565b16910152565b604051635f2cdc7560e11b81526001600160a01b0319821660048201529097919694959491936001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001693929161024081602481885afa908115611292575f9161133f575b5080516001600160a01b0316156113305760406003910151610f7a81610d02565b610f8381610d02565b0361132157610f918861136e565b610f9e886103238b610945565b946009860192610fb3845460ff9060401c1690565b61131257610fc0856113a7565b5f976001610fcd87610d0c565b610fd681610107565b0361129757505050506110239495506020876001945b60405163421adfbb60e01b81526001600160a01b03198c1660048201526024810192909252909687919082905f9082906044820190565b03925af1938415611292577f763f6464c291e9532f5ac2b75d76aa5283fc376d894a5e0b1e750ea0532cc7b1955f9561124d575b5080546001600160a01b031916331781556112489261122b92909161120391906111dd9061109f61108661090f565b8c8152602001899052600182018c905560028201899055565b60048101805460ff191660ff8b161790556110eb600582016110c96110c388610d0c565b82610d32565b6110d560208801610d4a565b815461ff00191690151560081b61ff0016179055565b6111056110fb6040870187610d57565b9060068401610daa565b6111c861111460608701610e30565b61112e6007840191829061ffff1661ffff19825416179055565b61116161113d60808901610e3a565b825462010000600160501b03191660109190911b62010000600160501b0316178255565b61119661117060a08901610e3a565b8254600160501b600160901b03191660509190911b600160501b600160901b0316178255565b6111a260c08801610e3a565b8154600160901b600160d01b03191660909190911b600160901b600160d01b0316179055565b60086111d660e08701610e3a565b9101610e44565b6111f0436001600160401b031682610e44565b805460ff60401b1916600160401b179055565b6001600160a01b031989165f908152600160205260409020611226908990610e91565b610d0c565b60405133986001600160a01b031916969094859490929185610ec5565b0390a4565b6112489391955061122b926111dd61127f6112039360203d60201161128b575b61127781836108da565b810190610d19565b97939550509250611057565b503d61126d565b610cf7565b90919297506112a6898961154b565b6112b361041d8a8a6115d0565b611303578782826112e2956112dc8f968f968f87819961041d9b6112d7868661154b565b611a72565b90611b1a565b6112f457602087611023969794610fec565b6327f7eb4d60e11b5f5260045ffd5b63b28e789160e01b5f5260045ffd5b630b792c8f60e01b5f5260045ffd5b63268dbf6760e21b5f5260045ffd5b63d5b25b6360e01b5f5260045ffd5b61136191506102403d8111611367575b61135981836108da565b810190610c1a565b5f610f59565b503d61134f565b801590811561137f575b5061057057565b5f5160206121215f395f51905f52915010155f611378565b9190811015610e8c5760051b0190565b604081016113b58183610d57565b91905060208211801561152f575b611443575f5b8281106114f75750505060e081016001600160401b036113e882610e3a565b161515806114da575b6114435760c082019061140661047c83610e3a565b151591826114c4575b82611499575b505061144357608081019061142c61047c83610e3a565b15159182611480575b82611452575b505061144357565b63d06b96b160e01b5f5260045ffd5b61146d91925060a061146661047c92610e3a565b9301610e3a565b6001600160401b03909116115f8061143b565b915061149161047c60a08301610e3a565b151591611435565b6114b19192506114ab61047c91610e3a565b92610e3a565b6001600160401b03909116115f80611415565b91506114d261047c82610e3a565b15159161140f565b506114e481610e3a565b426001600160401b0390911611156113f1565b6115226115166115118361150b8689610d57565b90611397565b610da0565b6001600160a01b031690565b15611443576001016113c9565b5061153c60208401610d4a565b80156113c357508115156113c3565b905f5160206121215f395f51905f528210806115ba575b156115ab578115806115a1575b6115925761157c91611816565b1561158357565b6361586bdd60e01b5f5260045ffd5b632b39517d60e21b5f5260045ffd5b506001811461156f565b63d7c7beeb60e01b5f5260045ffd5b505f5160206121215f395f51905f528110611562565b801580611731575b61172a576115e68282611816565b15611724576115f36118c3565b5061163e6115ff611df1565b809261161b8261160d61191a565b96602088015b938451611f56565b604085019061162e838351835190611f7e565b606086015b519151905191611fdf565b61164661091e565b5f81526020810192600184526040820192600184525f60608401525f91610100805b61169757505050505115918261168b575b5081611683575090565b905051151590565b5181511491505f611679565b60011901805f5160206121415f395f51905f52811c60031685156116f5576116c0858880611f7e565b6116cb858880611f7e565b84816116d9575b5050611668565b6116e66116ee928661194e565b518880611fdf565b5f846116d2565b809150611704575b5080611668565b61171b919450611714908361194e565b518561207a565b826001936116fd565b50505f90565b5050600190565b50600182146115d8565b9060068201805492831561179257505f5b83811061175b57505050505f90565b6117658183610e77565b905460039190911b1c6001600160a01b03908116908416146117895760010161174c565b50505050600190565b546001600160a01b0392831692169190911492915050565b5f5160206121215f395f51905f521190816117c3575090565b5f5160206121215f395f51905f5291501090565b634e487b7160e01b5f52601260045260245ffd5b5f5160206121215f395f51905f5203905f5160206121215f395f51905f52821161181157565b610d8c565b5f5160206121215f395f51905f5281108015906118ac575b611724575f5160206121215f395f51905f528181920991800990805f5160206121215f395f51905f5203915f5160206121215f395f51905f528311611811575f5160206121215f395f51905f528080838195097f1aee90f15f2189693df072d799fd11fc039b2959ebb7c867d075ca8cf4d7eb8e0960010892081490565b505f5160206121215f395f51905f5282101561182e565b60405190608082016001600160401b0381118382101761089e576040525f6060838281528260208201528260408201520152565b5f5b82811061190557505050565b6020906119106118c3565b81840152016118f9565b604051906119296080836108da565b61090d6080836118f7565b6040519061090d61020061194881856108da565b836118f7565b906004811015610e8c5760051b0190565b9092919261196b6118c3565b506119ab611977611df1565b809261198f8261198561191a565b9860208a01611613565b60408701906119a2838351835190611f7e565b60608801611633565b6119b361091e565b5f815260016020820152600160408201525f606082015280945f93610100805b6119df57505050505050565b600119018082811c6003168715611a30576119fb878780611f7e565b611a06878780611f7e565b8681611a14575b50506119d3565b611a21611a29928761194e565b518780611fdf565b5f86611a0d565b809150611a3f575b50806119d3565b611a56919650611a4f908461194e565b518461207a565b84600195611a38565b5f5160206121415f395f51905f52900690565b9390925f5160206121415f395f51905f5295926040519460208601967f41ea6f3fa95eccd1f3b1ce8e05efa11027280aa0c6b4167fd6695db659c30b28885260018060a01b0319166040870152604c860152606c850152608c84015260ac83015260cc82015260cc8152611ae760ec826108da565b5190200690565b906010811015610e8c5760051b0190565b906004820180921161181157565b9190820180921161181157565b959491929395611b2d61041d86866117aa565b8015611d15575b611d0b57611b44611b4a91611a5f565b91611a5f565b9580611cfd57505f5b611bd4611b5e611df1565b8092611bb882611b6c611934565b97600160208a510152600160408a510152611b9182611b8b8b60800190565b51611e3f565b611ba0826101008b0151611e86565b611baf826101808b0151611eee565b60208901611613565b6040860190611bcb838351835190611f7e565b60608701611633565b60045b600c811115611ca65750611be961091e565b925f845260016020850152600160408501525f60608501525f9260fc805b611c1b5750505050506101e093945061209e565b600119019889600c83821c60021b1682821c600316178615611c7a57611c42868980611f7e565b611c4d868980611f7e565b8581611c5e575b50505b9099611c07565b611c6b611c739287611aee565b518980611fdf565b5f85611c54565b80611c86575b50611c57565b611c9d919650611c969085611aee565b518761207a565b6001945f611c80565b60015b60048110611cc05750611cbb90611aff565b611bd7565b80611cf784611cda611cd460019587611b0d565b89611aee565b51611ce5868a611aee565b51611cf0858b611aee565b5191611fdf565b01611ca9565b611d06906117eb565b611b53565b505f955050505050565b50611d2361041d84896117aa565b611b34565b5f5160206121415f395f51905f5290068015611de9575f5160206121015f395f51905f525f5160206121615f395f51905f52611d639261195f565b90604082015160405190602082526020808301526020604083015260608201527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593efffffff60808201525f5160206121215f395f51905f5260a082015260208160c08160055afa156100d8575f5160206121215f395f51905f529051602082828651099401510990565b505f90600190565b611df961092d565b50604051611e06816108bf565b5f5160206121215f395f51905f5281527f0578d36fdd1172a8c3909ff8b278cb9adf026a6b5db6203e5d099f85f9afd71b602082015290565b90518015611e8157806060915f5160206121615f395f51905f52068352805f5160206121015f395f51905f520680602085015260016040850152835109910152565b6117d7565b90518015611e8157806060917f0daaa7e6b25c28e6dc8dd1d48e9cc61cd07015c1d7c1b8d4590eb6f51d5346dc068352807f01666cafbf0a30da8b9ebeaf848a1da067a892296f1043188e1705402b6d68530680602085015260016040850152835109910152565b90518015611e8157806060917f136d609c4c856f5d277fab08c730cbdd1a776ce4728c6a2eb20ff22bccf26894068352807f21d66f0e2295ae954494f25889f9319cc1b4df71eff3f46ba9e4631b43fd7c950680602085015260016040850152835109910152565b9251908115611e81576060928280920685520680602085015260016040850152835109910152565b9151815191602081019081518315611e81578380808093604098088180808a818080808c5180099c518009818d810382089c08810380988782980908980151800980088103870894828682098a520960608801528309602086015209910152565b92908151926020820192835183518603918615611e815786808086818080999881808d81809d9c816020819f01968188518c51820390089208099f519051900891518551900890099a818181038d089b089687958160608c0151606085015190099060200151900998604001519060400151900980089581818103880896089582868209895209606087015283096020850152099060400152565b90606080918051845260208101516020850152604081015160408501520151910152565b90916040820180519384151594856120e2575b5050836120bf575b50505090565b519192506020915f5160206121215f395f51905f529109910151145f80806120b9565b84519295505f5160206121215f395f51905f52910914925f806120b156fe25797203f7a0b24925572e1cd16bf9edfce0051fb9e133774b3c257a872d7d8b30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001060c89ce5c263405370a08b6d0302b0bab3eedb83920ee0a677297dc392126f11561ff836ce19d358a4eb7a4c199e94c377c749ae6f2a277f1f9195afe553f9fa26469706673582212205f75c46ba1db4bebb59b1727532b638dd40593432deead4da291700c6b25d7ef64736f6c634300081c0033",
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
