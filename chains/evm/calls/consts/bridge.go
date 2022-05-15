package consts

const BridgeABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"domainID\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"destinationDomainID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"handlerResponse\",\"type\":\"bytes\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EndKeygen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lowLevelData\",\"type\":\"bytes\"}],\"name\":\"FailedHandlerExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"KeyRefresh\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"originDomainID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StartKeygen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_MPCAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"_depositCounts\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_domainID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_expiry\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_fee\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToHandlerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getRoleMemberIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isValidForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usedNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminPauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminUnpauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"depositFunctionSig\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"depositFunctionDepositerOffset\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"executeFunctionSig\",\"type\":\"bytes4\"}],\"name\":\"adminSetGenericResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"domainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"adminSetDepositNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"adminSetForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"adminChangeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"adminWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationDomainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"domainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"revertOnFail\",\"type\":\"bool\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"transferFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startKeygen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"MPCAddress\",\"type\":\"address\"}],\"name\":\"endKeygen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refreshKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"domainID\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"}],\"name\":\"isProposalExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
const BridgeBin = "0x60806040523480156200001157600080fd5b5060405162002b0538038062002b05833981016040819052620000349162000394565b6000805460ff199081169091556002805490911660ff85161790556200006682620000fc602090811b6200159f17901c565b600260016101000a8154816001600160801b0302191690836001600160801b03160217905550620000a2816200015b60201b620015f81760201c565b6002805464ffffffffff92909216600160881b0264ffffffffff60881b19909216919091179055620000df6000620000d9620001b4565b620001f7565b620000f3620000ed620001b4565b62000207565b505050620003d3565b6000600160801b8210620001575760405162461bcd60e51b815260206004820152601e60248201527f76616c756520646f6573206e6f742066697420696e203132382062697473000060448201526064015b60405180910390fd5b5090565b6000650100000000008210620001575760405162461bcd60e51b815260206004820152601d60248201527f76616c756520646f6573206e6f742066697420696e203430206269747300000060448201526064016200014e565b60003360143610801590620001e157506001600160a01b03811660009081526006602052604090205460ff165b15620001f2575060131936013560601c5b919050565b6200020382826200025d565b5050565b62000211620002d8565b6000805460ff191660011790556040516001600160a01b03821681527f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2589060200160405180910390a150565b6000828152600160209081526040909120620002849183906200164f62000322821b17901c565b15620002035762000294620001b4565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60005460ff1615620003205760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016200014e565b565b600062000339836001600160a01b03841662000342565b90505b92915050565b60008181526001830160205260408120546200038b575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556200033c565b5060006200033c565b600080600060608486031215620003aa57600080fd5b835160ff81168114620003bc57600080fd5b602085015160409095015190969495509392505050565b61272280620003e36000396000f3fe6080604052600436106102045760003560e01c80639010d07c11610118578063ca15c873116100a0578063d547741f1161006f578063d547741f1461067c578063edc20c3c1461069c578063f5f63b39146106bc578063f8c39e44146106d1578063ffaac0eb1461070157600080fd5b8063ca15c873146105fc578063cb10f2151461061c578063d15ef64e1461063c578063d2e5fae91461065c57600080fd5b80639dd694f4116100e75780639dd694f414610523578063a217fddf1461054f578063bd2a182014610564578063c5b37c2214610584578063c5ec8970146105c157600080fd5b80639010d07c146104a357806391c404ac146104c357806391d14854146104e35780639ae0bf451461050357600080fd5b80634b0b919d1161019b5780635e1fab0f1161016a5780635e1fab0f146104035780636ba6db6b1461042357806380ae1c281461043857806384db809f1461044d5780638c0c26311461048357600080fd5b80634b0b919d146103515780634e0df3f61461039f5780635a1ad87c146103bf5780635c975abb146103df57600080fd5b8063248a9ca3116101d7578063248a9ca3146102c15780632f2ff15d146102f157806336568abe146103115780634603ae381461033157600080fd5b8063059972d21461020957806305e2ca171461024657806308a641041461025b578063202f5b8c146102a1575b600080fd5b34801561021557600080fd5b50600354610229906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b610259610254366004611e82565b610716565b005b34801561026757600080fd5b50610293610276366004611edb565b600760209081526000928352604080842090915290825290205481565b60405190815260200161023d565b3480156102ad57600080fd5b506102596102bc366004611f2c565b610910565b3480156102cd57600080fd5b506102936102dc366004611fd6565b60009081526001602052604090206002015490565b3480156102fd57600080fd5b5061025961030c366004612004565b610c9d565b34801561031d57600080fd5b5061025961032c366004612004565b610d2d565b34801561033d57600080fd5b5061025961034c366004612078565b610db7565b34801561035d57600080fd5b5061038761036c3660046120d7565b6004602052600090815260409020546001600160401b031681565b6040516001600160401b03909116815260200161023d565b3480156103ab57600080fd5b506102936103ba366004612004565b610e5b565b3480156103cb57600080fd5b506102596103da36600461210a565b610e87565b3480156103eb57600080fd5b5060005460ff165b604051901515815260200161023d565b34801561040f57600080fd5b5061025961041e366004612174565b610f3d565b34801561042f57600080fd5b50610259610fc9565b34801561044457600080fd5b50610259611055565b34801561045957600080fd5b50610229610468366004611fd6565b6005602052600090815260409020546001600160a01b031681565b34801561048f57600080fd5b5061025961049e366004612191565b61106f565b3480156104af57600080fd5b506102296104be3660046121bf565b6110db565b3480156104cf57600080fd5b506102596104de366004611fd6565b6110fa565b3480156104ef57600080fd5b506103f36104fe366004612004565b611194565b34801561050f57600080fd5b506103f361051e366004611edb565b6111ac565b34801561052f57600080fd5b5060025461053d9060ff1681565b60405160ff909116815260200161023d565b34801561055b57600080fd5b50610293600081565b34801561057057600080fd5b5061025961057f36600461224e565b6111fb565b34801561059057600080fd5b506002546105a99061010090046001600160801b031681565b6040516001600160801b03909116815260200161023d565b3480156105cd57600080fd5b506002546105e690600160881b900464ffffffffff1681565b60405164ffffffffff909116815260200161023d565b34801561060857600080fd5b50610293610617366004611fd6565b611231565b34801561062857600080fd5b506102596106373660046122e0565b611248565b34801561064857600080fd5b50610259610657366004612322565b6112d2565b34801561066857600080fd5b50610259610677366004612174565b611305565b34801561068857600080fd5b50610259610697366004612004565b61141d565b3480156106a857600080fd5b506102596106b7366004612357565b6114a0565b3480156106c857600080fd5b50610259611559565b3480156106dd57600080fd5b506103f36106ec366004612174565b60066020526000908152604090205460ff1681565b34801561070d57600080fd5b5061025961158c565b61071e611664565b60025461010090046001600160801b0316341461077b5760405162461bcd60e51b8152602060048201526016602482015275125b98dbdc9c9958dd08199959481cdd5c1c1b1a595960521b60448201526064015b60405180910390fd5b6000838152600560205260409020546001600160a01b0316806107e05760405162461bcd60e51b815260206004820181905260248201527f7265736f757263654944206e6f74206d617070656420746f2068616e646c65726044820152606401610772565b60ff8516600090815260046020526040812080548290610808906001600160401b0316612397565b91906101000a8154816001600160401b0302191690836001600160401b031602179055905060006108376116aa565b60405163b07e54bb60e01b815290915083906000906001600160a01b0383169063b07e54bb90610871908b9087908c908c906004016123e7565b6000604051808303816000875af1158015610890573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526108b8919081019061244c565b9050826001600160a01b03167f17bc3181e17a9620a479c24e6c606e474ba84fc036877b768926872e8cd0e11f8a8a878b8b876040516108fd969594939291906124ee565b60405180910390a2505050505050505050565b610918611664565b61092b88886001600160401b03166111ac565b1515600114156109925760405162461bcd60e51b815260206004820152602c60248201527f4465706f73697420776974682070726f7669646564206e6f6e636520616c726560448201526b18591e48195e1958dd5d195960a21b6064820152608401610772565b6000610a0289898989896040516020016109b095949392919061253f565b6040516020818303038152906040528051906020012085858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506116eb92505050565b6003549091506001600160a01b03808316911614610a5b5760405162461bcd60e51b815260206004820152601660248201527524b73b30b634b21036b2b9b9b0b3b29039b4b3b732b960511b6044820152606401610772565b60008581526005602090815260408083205490516001600160a01b039091169291610a8c9184918c918c910161257d565b60408051601f198184030181529190528051602090910120905081610ab36101008c6125bf565b60ff8d16600090815260076020526040812060016001600160401b03939093169290921b9190610ae56101008f6125e5565b6001600160401b031681526020810191909152604001600020805490911790558415610b725760405163712467f960e11b81526001600160a01b0382169063e248cff290610b3b908b908e908e9060040161260b565b600060405180830381600087803b158015610b5557600080fd5b505af1158015610b69573d6000803e3d6000fd5b50505050610c42565b60405163712467f960e11b81526001600160a01b0382169063e248cff290610ba2908b908e908e9060040161260b565b600060405180830381600087803b158015610bbc57600080fd5b505af1925050508015610bcd575060015b610c42573d808015610bfb576040519150601f19603f3d011682016040523d82523d6000602084013e610c00565b606091505b507fbd37c1f0d53bb2f33fe4c2104de272fcdeb4d2fef3acdbf1e4ddc3d6833ca37681604051610c309190612625565b60405180910390a15050505050610c93565b6040805160ff8e1681526001600160401b038d1660208201529081018390527f6018c584b8d99bafeda249b2429f5907d830e792222070c1b3a94aa76ee716779060600160405180910390a1505050505b5050505050505050565b600082815260016020526040902060020154610cbb906104fe6116aa565b610d1f5760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60448201526e0818591b5a5b881d1bc819dc985b9d608a1b6064820152608401610772565b610d29828261170f565b5050565b610d356116aa565b6001600160a01b0316816001600160a01b031614610dad5760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b6064820152608401610772565b610d298282611778565b610dbf6117e1565b60005b83811015610e5457848482818110610ddc57610ddc612638565b9050602002016020810190610df19190612174565b6001600160a01b03166108fc848484818110610e0f57610e0f612638565b905060200201359081150290604051600060405180830381858888f19350505050158015610e41573d6000803e3d6000fd5b5080610e4c8161264e565b915050610dc2565b5050505050565b60008281526001602081815260408084206001600160a01b038616855290920190529020545b92915050565b610e8f6117e1565b6000858152600560205260409081902080546001600160a01b0319166001600160a01b03898116918217909255915163de319d9960e01b81526004810188905290861660248201526001600160e01b03198086166044830152606482018590528316608482015287919063de319d999060a401600060405180830381600087803b158015610f1c57600080fd5b505af1158015610f30573d6000803e3d6000fd5b5050505050505050505050565b610f456117e1565b6000610f4f6116aa565b9050816001600160a01b0316816001600160a01b03161415610fb35760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f742072656e6f756e6365206f6e6573656c660000000000000000006044820152606401610772565b610fbe600083610c9d565b610d29600082610d2d565b610fd16117e1565b6003546001600160a01b03161561102a5760405162461bcd60e51b815260206004820152601a60248201527f4d5043206164647265737320697320616c7265616479207365740000000000006044820152606401610772565b6040517f24e723a5c27b62883404028b8dee9965934de6a46828cda2ff63bf9a5e65ce4390600090a1565b61105d6117e1565b61106d6110686116aa565b61183a565b565b6110776117e1565b6040516307b7ed9960e01b81526001600160a01b0382811660048301528391908216906307b7ed99906024015b600060405180830381600087803b1580156110be57600080fd5b505af11580156110d2573d6000803e3d6000fd5b50505050505050565b60008281526001602052604081206110f3908361188f565b9392505050565b6111026117e1565b60025461010090046001600160801b03168114156111625760405162461bcd60e51b815260206004820152601f60248201527f43757272656e742066656520697320657175616c20746f206e657720666565006044820152606401610772565b61116b8161159f565b600260016101000a8154816001600160801b0302191690836001600160801b0316021790555050565b60008281526001602052604081206110f3908361189b565b60006111ba61010083612669565b60ff84166000908152600760205260408120600190921b91906111df6101008661267d565b8152602001908152602001600020541660001415905092915050565b6112036117e1565b60405163025a3c9960e21b815282906001600160a01b03821690630968f264906110a4908590600401612625565b6000818152600160205260408120610e81906118bd565b6112506117e1565b6000828152600560205260409081902080546001600160a01b0319166001600160a01b038681169182179092559151635c7d1b9b60e11b815260048101859052908316602482015284919063b8fa373690604401600060405180830381600087803b1580156112be57600080fd5b505af1158015610c93573d6000803e3d6000fd5b6112da6117e1565b6001600160a01b03919091166000908152600660205260409020805460ff1916911515919091179055565b61130d6117e1565b6001600160a01b03811661136d5760405162461bcd60e51b815260206004820152602160248201527f4d504320616464726573732063616e2774206265207a65726f2d6164647265736044820152607360f81b6064820152608401610772565b6003546001600160a01b0316156113c65760405162461bcd60e51b815260206004820152601c60248201527f4d504320616464726573732063616e27742062652075706461746564000000006044820152606401610772565b600380546001600160a01b0319166001600160a01b0383161790556113f16113ec6116aa565b6118c7565b6040517f4187686ceef7b541a1f224d48d4cded8f2c535e0e58ac0f0514071b1de3dad5790600090a150565b60008281526001602052604090206002015461143b906104fe6116aa565b610dad5760405162461bcd60e51b815260206004820152603060248201527f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60448201526f2061646d696e20746f207265766f6b6560801b6064820152608401610772565b6114a86117e1565b60ff82166000908152600460205260409020546001600160401b03908116908216116115255760405162461bcd60e51b815260206004820152602660248201527f446f6573206e6f7420616c6c6f772064656372656d656e7473206f6620746865604482015265206e6f6e636560d01b6064820152608401610772565b60ff919091166000908152600460205260409020805467ffffffffffffffff19166001600160401b03909216919091179055565b6115616117e1565b6040517f034a03238f3c0f2cad22894b0fa8810b6ffcf678a2560e54a7e41b4e9cebd02e90600090a1565b6115946117e1565b61106d6113ec6116aa565b6000600160801b82106115f45760405162461bcd60e51b815260206004820152601e60248201527f76616c756520646f6573206e6f742066697420696e20313238206269747300006044820152606401610772565b5090565b60006501000000000082106115f45760405162461bcd60e51b815260206004820152601d60248201527f76616c756520646f6573206e6f742066697420696e20343020626974730000006044820152606401610772565b60006110f3836001600160a01b038416611912565b60005460ff161561106d5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610772565b600033601436108015906116d657506001600160a01b03811660009081526006602052604090205460ff165b156116e6575060131936013560601c5b919050565b60008060006116fa8585611961565b91509150611707816119d1565b509392505050565b6000828152600160205260409020611727908261164f565b15610d29576117346116aa565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60008281526001602052604090206117909082611b8f565b15610d295761179d6116aa565b6001600160a01b0316816001600160a01b0316837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45050565b6117ee60006104fe6116aa565b61106d5760405162461bcd60e51b815260206004820152601e60248201527f73656e64657220646f65736e277420686176652061646d696e20726f6c6500006044820152606401610772565b611842611664565b6000805460ff191660011790556040516001600160a01b03821681527f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258906020015b60405180910390a150565b60006110f38383611ba4565b6001600160a01b038116600090815260018301602052604081205415156110f3565b6000610e81825490565b6118cf611bce565b6000805460ff191690556040516001600160a01b03821681527f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa90602001611884565b600081815260018301602052604081205461195957508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610e81565b506000610e81565b6000808251604114156119985760208301516040840151606085015160001a61198c87828585611c17565b945094505050506119ca565b8251604014156119c257602083015160408401516119b7868383611d04565b9350935050506119ca565b506000905060025b9250929050565b60008160048111156119e5576119e5612691565b14156119ee5750565b6001816004811115611a0257611a02612691565b1415611a505760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610772565b6002816004811115611a6457611a64612691565b1415611ab25760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610772565b6003816004811115611ac657611ac6612691565b1415611b1f5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610772565b6004816004811115611b3357611b33612691565b1415611b8c5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610772565b50565b60006110f3836001600160a01b038416611d3d565b6000826000018281548110611bbb57611bbb612638565b9060005260206000200154905092915050565b60005460ff1661106d5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610772565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115611c4e5750600090506003611cfb565b8460ff16601b14158015611c6657508460ff16601c14155b15611c775750600090506004611cfb565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611ccb573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611cf457600060019250925050611cfb565b9150600090505b94509492505050565b6000806001600160ff1b03831681611d2160ff86901c601b6126a7565b9050611d2f87828885611c17565b935093505050935093915050565b60008181526001830160205260408120548015611e26576000611d616001836126bf565b8554909150600090611d75906001906126bf565b9050818114611dda576000866000018281548110611d9557611d95612638565b9060005260206000200154905080876000018481548110611db857611db8612638565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080611deb57611deb6126d6565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610e81565b6000915050610e81565b803560ff811681146116e657600080fd5b60008083601f840112611e5357600080fd5b5081356001600160401b03811115611e6a57600080fd5b6020830191508360208285010111156119ca57600080fd5b60008060008060608587031215611e9857600080fd5b611ea185611e30565b93506020850135925060408501356001600160401b03811115611ec357600080fd5b611ecf87828801611e41565b95989497509550505050565b60008060408385031215611eee57600080fd5b611ef783611e30565b946020939093013593505050565b80356001600160401b03811681146116e657600080fd5b803580151581146116e657600080fd5b60008060008060008060008060c0898b031215611f4857600080fd5b611f5189611e30565b9750611f5f60208a01611f05565b965060408901356001600160401b0380821115611f7b57600080fd5b611f878c838d01611e41565b909850965060608b0135955060808b0135915080821115611fa757600080fd5b50611fb48b828c01611e41565b9094509250611fc7905060a08a01611f1c565b90509295985092959890939650565b600060208284031215611fe857600080fd5b5035919050565b6001600160a01b0381168114611b8c57600080fd5b6000806040838503121561201757600080fd5b82359150602083013561202981611fef565b809150509250929050565b60008083601f84011261204657600080fd5b5081356001600160401b0381111561205d57600080fd5b6020830191508360208260051b85010111156119ca57600080fd5b6000806000806040858703121561208e57600080fd5b84356001600160401b03808211156120a557600080fd5b6120b188838901612034565b909650945060208701359150808211156120ca57600080fd5b50611ecf87828801612034565b6000602082840312156120e957600080fd5b6110f382611e30565b80356001600160e01b0319811681146116e657600080fd5b60008060008060008060c0878903121561212357600080fd5b863561212e81611fef565b955060208701359450604087013561214581611fef565b9350612153606088016120f2565b92506080870135915061216860a088016120f2565b90509295509295509295565b60006020828403121561218657600080fd5b81356110f381611fef565b600080604083850312156121a457600080fd5b82356121af81611fef565b9150602083013561202981611fef565b600080604083850312156121d257600080fd5b50508035926020909101359150565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b038111828210171561221f5761221f6121e1565b604052919050565b60006001600160401b03821115612240576122406121e1565b50601f01601f191660200190565b6000806040838503121561226157600080fd5b823561226c81611fef565b915060208301356001600160401b0381111561228757600080fd5b8301601f8101851361229857600080fd5b80356122ab6122a682612227565b6121f7565b8181528660208385010111156122c057600080fd5b816020840160208301376000602083830101528093505050509250929050565b6000806000606084860312156122f557600080fd5b833561230081611fef565b925060208401359150604084013561231781611fef565b809150509250925092565b6000806040838503121561233557600080fd5b823561234081611fef565b915061234e60208401611f1c565b90509250929050565b6000806040838503121561236a57600080fd5b61237383611e30565b915061234e60208401611f05565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b03808316818114156123b4576123b4612381565b6001019392505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8481526001600160a01b038416602082015260606040820181905260009061241290830184866123be565b9695505050505050565b60005b8381101561243757818101518382015260200161241f565b83811115612446576000848401525b50505050565b60006020828403121561245e57600080fd5b81516001600160401b0381111561247457600080fd5b8201601f8101841361248557600080fd5b80516124936122a682612227565b8181528560208385010111156124a857600080fd5b6124b982602083016020860161241c565b95945050505050565b600081518084526124da81602086016020860161241c565b601f01601f19169290920160200192915050565b60ff871681528560208201526001600160401b038516604082015260a06060820152600061252060a0830185876123be565b828103608084015261253281856124c2565b9998505050505050505050565b60ff861681526001600160401b038516602082015260806040820152600061256b6080830185876123be565b90508260608301529695505050505050565b6bffffffffffffffffffffffff198460601b168152818360148301376000910160140190815292915050565b634e487b7160e01b600052601260045260246000fd5b60006001600160401b03808416806125d9576125d96125a9565b92169190910692915050565b60006001600160401b03808416806125ff576125ff6125a9565b92169190910492915050565b8381526040602082015260006124b96040830184866123be565b6020815260006110f360208301846124c2565b634e487b7160e01b600052603260045260246000fd5b600060001982141561266257612662612381565b5060010190565b600082612678576126786125a9565b500690565b60008261268c5761268c6125a9565b500490565b634e487b7160e01b600052602160045260246000fd5b600082198211156126ba576126ba612381565b500190565b6000828210156126d1576126d1612381565b500390565b634e487b7160e01b600052603160045260246000fdfea2646970667358221220a4a8362757d7f55b77ad27c235fcdcb6cefecf362c6957968ebbedc49575abbd64736f6c634300080b0033"
