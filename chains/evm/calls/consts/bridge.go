package consts

const BridgeABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"domainID\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"accessControl\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAccessControl\",\"type\":\"address\"}],\"name\":\"AccessControlChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"destinationDomainID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"handlerResponse\",\"type\":\"bytes\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EndKeygen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lowLevelData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"originDomainID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"}],\"name\":\"FailedHandlerExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeHandler\",\"type\":\"address\"}],\"name\":\"FeeHandlerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"KeyRefresh\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pero\",\"type\":\"string\"}],\"name\":\"Pero\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"originDomainID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"}],\"name\":\"Retry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StartKeygen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_MPCAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_accessControl\",\"outputs\":[{\"internalType\":\"contractIAccessControlSegregator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"_depositCounts\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_domainID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_feeHandler\",\"outputs\":[{\"internalType\":\"contractIFeeHandler\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToHandlerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isValidForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usedNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminPauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminUnpauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"depositFunctionSig\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"depositFunctionDepositerOffset\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"executeFunctionSig\",\"type\":\"bytes4\"}],\"name\":\"adminSetGenericResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"domainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"adminSetDepositNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"adminSetForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAccessControl\",\"type\":\"address\"}],\"name\":\"adminChangeAccessControl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeHandler\",\"type\":\"address\"}],\"name\":\"adminChangeFeeHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"adminWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationDomainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"depositData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"originDomainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structBridge.Proposal[]\",\"name\":\"proposal\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"originDomainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structBridge.Proposal[]\",\"name\":\"proposals\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"executeProposals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startKeygen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"MPCAddress\",\"type\":\"address\"}],\"name\":\"endKeygen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"refreshKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"}],\"name\":\"retry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"domainID\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"}],\"name\":\"isProposalExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"originDomainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structBridge.Proposal[]\",\"name\":\"proposals\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
const BridgeBin = "0x6101606040523480156200001257600080fd5b50604051620033f4380380620033f4833981016040819052620000359162000253565b604080518082018252600681526542726964676560d01b6020808301918252835180850190945260058452640332e312e360dc1b908401526000805460ff191690558151902060e08190527f0e23d0b508e2034d01a5c31f12e9d9bbb31708c5518057dde31201ab93b17cef6101008190524660a0529192917f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6200011f8184846040805160208101859052908101839052606081018290524660808201523060a082015260009060c0016040516020818303038152906040528051906020012090509392505050565b6080523060c052610120525050505060ff821661014052600280546001600160a01b0319166001600160a01b038316179055620001656200015f6200016d565b620001b0565b5050620002a2565b600033601436108015906200019a57506001600160a01b03811660009081526005602052604090205460ff165b15620001ab575060131936013560601c5b919050565b620001ba62000206565b6000805460ff191660011790556040516001600160a01b03821681527f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2589060200160405180910390a150565b60005460ff1615620002515760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640160405180910390fd5b565b600080604083850312156200026757600080fd5b825160ff811681146200027957600080fd5b60208401519092506001600160a01b03811681146200029757600080fd5b809150509250929050565b60805160a05160c05160e0516101005161012051610140516130e2620003126000396000818161045d0152818161064701528181610d140152610e24015260006122ee0152600061233d01526000612318015260006122710152600061229b015260006122c501526130e26000f3fe6080604052600436106101b75760003560e01c80638c0c2631116100ec578063d15ef64e1161008a578063edc20c3c11610064578063edc20c3c14610551578063f8c39e4414610571578063fe4648f4146105a1578063ffaac0eb146105c157600080fd5b8063d15ef64e146104f1578063d2e5fae914610511578063d82367441461053157600080fd5b80639dd694f4116100c65780639dd694f41461044b578063a546e8a114610491578063bd2a1820146104b1578063cb10f215146104d157600080fd5b80638c0c2631146103eb5780639ae0bf451461040b5780639d33b6d41461042b57600080fd5b80635c975abb1161015957806376c5d76d1161013357806376c5d76d1461036057806380ae1c281461038057806384db809f146103955780638b63aebf146103cb57600080fd5b80635c975abb146103145780636ba6db6b1461033857806373c45c981461034d57600080fd5b8063366b488511610195578063366b48851461026657806344e8e430146102865780634b0b919d146102a65780635a1ad87c146102f457600080fd5b8063059972d2146101bc57806308a64104146101fe5780631f5c64c114610244575b600080fd5b3480156101c857600080fd5b506000546101e19061010090046001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561020a57600080fd5b5061023661021936600461250b565b600660209081526000928352604080842090915290825290205481565b6040519081526020016101f5565b34801561025057600080fd5b5061026461025f366004612646565b6105d6565b005b34801561027257600080fd5b5061026461028136600461278b565b610b64565b34801561029257600080fd5b506002546101e1906001600160a01b031681565b3480156102b257600080fd5b506102dc6102c13660046127db565b6003602052600090815260409020546001600160401b031681565b6040516001600160401b0390911681526020016101f5565b34801561030057600080fd5b5061026461030f366004612825565b610b9e565b34801561032057600080fd5b5060005460ff165b60405190151581526020016101f5565b34801561034457600080fd5b50610264610c69565b61026461035b3660046128cc565b610d0a565b34801561036c57600080fd5b5061026461037b366004612955565b611017565b34801561038c57600080fd5b5061026461156c565b3480156103a157600080fd5b506101e16103b03660046129ee565b6004602052600090815260409020546001600160a01b031681565b3480156103d757600080fd5b506102646103e6366004612a07565b611596565b3480156103f757600080fd5b50610264610406366004612a22565b6115fc565b34801561041757600080fd5b5061032861042636600461250b565b611678565b34801561043757600080fd5b50610264610446366004612a07565b6116c7565b34801561045757600080fd5b5061047f7f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff90911681526020016101f5565b34801561049d57600080fd5b506103286104ac366004612955565b61172d565b3480156104bd57600080fd5b506102646104cc366004612a55565b6119dc565b3480156104dd57600080fd5b506102646104ec366004612a98565b611a22565b3480156104fd57600080fd5b5061026461050c366004612ae2565b611ac6565b34801561051d57600080fd5b5061026461052c366004612a07565b611b09565b34801561053d57600080fd5b5061026461054c36600461278b565b611c3b565b34801561055d57600080fd5b5061026461056c366004612b19565b611c82565b34801561057d57600080fd5b5061032861058c366004612a07565b60056020526000908152604090205460ff1681565b3480156105ad57600080fd5b506001546101e1906001600160a01b031681565b3480156105cd57600080fd5b50610264611d4b565b6105de611dc1565b600082511161063e5760405162461bcd60e51b815260206004820152602160248201527f50726f706f73616c732063616e277420626520616e20656d70747920617272616044820152607960f81b60648201526084015b60405180910390fd5b600061069b82847f0000000000000000000000000000000000000000000000000000000000000000604051602001610677929190612b9b565b60405160208183030381529060405280519060200120611e0790919063ffffffff16565b6000549091506001600160a01b0380831661010090920416146106f95760405162461bcd60e51b815260206004820152601660248201527524b73b30b634b21036b2b9b9b0b3b29039b4b3b732b960511b6044820152606401610635565b60005b8351811015610b5e5761075284828151811061071a5761071a612c42565b60200260200101516000015185838151811061073857610738612c42565b6020026020010151602001516001600160401b0316611678565b1561075c57610b4c565b60006004600086848151811061077457610774612c42565b602002602001015160400151815260200190815260200160002060009054906101000a90046001600160a01b031690506000818684815181106107b9576107b9612c42565b6020026020010151606001516040516020016107d6929190612c58565b604051602081830303815290604052805190602001209050600082905061010087858151811061080857610808612c42565b60200260200101516020015161081e9190612ca6565b6001600160401b03166001901b6006600089878151811061084157610841612c42565b60200260200101516000015160ff1660ff16815260200190815260200160002060006101008a888151811061087857610878612c42565b60200260200101516020015161088e9190612ce2565b6001600160401b0316815260200190815260200160002060008282541792505081905550806001600160a01b031663e248cff28886815181106108d3576108d3612c42565b6020026020010151604001518987815181106108f1576108f1612c42565b6020026020010151606001516040518363ffffffff1660e01b815260040161091a929190612d08565b600060405180830381600087803b15801561093457600080fd5b505af1925050508015610945575060015b610ab4573d808015610973576040519150601f19603f3d011682016040523d82523d6000602084013e610978565b606091505b507f19f774a63ee465292252a9981ae52051acc13da671c698ac4b5bf25b38c5b6fc818987815181106109ad576109ad612c42565b6020026020010151600001518a88815181106109cb576109cb612c42565b6020026020010151602001516040516109e693929190612d21565b60405180910390a1610100888681518110610a0357610a03612c42565b602002602001015160200151610a199190612ca6565b6001600160401b03166001901b19600660008a8881518110610a3d57610a3d612c42565b60200260200101516000015160ff1660ff16815260200190815260200160002060006101008b8981518110610a7457610a74612c42565b602002602001015160200151610a8a9190612ce2565b6001600160401b0316815260208101919091526040016000208054909116905550610b4c92505050565b7f6018c584b8d99bafeda249b2429f5907d830e792222070c1b3a94aa76ee71677878581518110610ae757610ae7612c42565b602002602001015160000151888681518110610b0557610b05612c42565b60200260200101516020015184604051610b409392919060ff9390931683526001600160401b03919091166020830152604082015260600190565b60405180910390a15050505b80610b5681612d56565b9150506106fc565b50505050565b7f9069464c059b9a90135a3fdf2c47855263346b912894ad7562d989532c3fad4c81604051610b939190612d71565b60405180910390a150565b610bbb6000356001600160e01b031916610bb6611e2b565b611e6c565b60008581526004602081905260409182902080546001600160a01b0319166001600160a01b038a8116918217909255925163de319d9960e01b8152918201889052861660248201526001600160e01b03198086166044830152606482018590528316608482015287919063de319d999060a401600060405180830381600087803b158015610c4857600080fd5b505af1158015610c5c573d6000803e3d6000fd5b5050505050505050505050565b610c816000356001600160e01b031916610bb6611e2b565b60005461010090046001600160a01b031615610cdf5760405162461bcd60e51b815260206004820152601a60248201527f4d5043206164647265737320697320616c7265616479207365740000000000006044820152606401610635565b6040517f24e723a5c27b62883404028b8dee9965934de6a46828cda2ff63bf9a5e65ce4390600090a1565b610d12611dc1565b7f000000000000000000000000000000000000000000000000000000000000000060ff168660ff161415610d885760405162461bcd60e51b815260206004820152601f60248201527f43616e2774206465706f73697420746f2063757272656e7420646f6d61696e006044820152606401610635565b6000610d92611e2b565b6001549091506001600160a01b0316610df8573415610df35760405162461bcd60e51b815260206004820152601d60248201527f6e6f2046656548616e646c65722c206d73672e76616c756520213d20300000006044820152606401610635565b610e8b565b600154604051632530706560e01b81526001600160a01b03909116906325307065903490610e589085907f0000000000000000000000000000000000000000000000000000000000000000908d908d908d908d908d908d90600401612dad565b6000604051808303818588803b158015610e7157600080fd5b505af1158015610e85573d6000803e3d6000fd5b50505050505b6000868152600460205260409020546001600160a01b031680610ef05760405162461bcd60e51b815260206004820181905260248201527f7265736f757263654944206e6f74206d617070656420746f2068616e646c65726044820152606401610635565b60ff8816600090815260036020526040812080548290610f18906001600160401b0316612e09565b91906101000a8154816001600160401b0302191690836001600160401b031602179055905060008290506000816001600160a01b031663b07e54bb8b878c8c6040518563ffffffff1660e01b8152600401610f769493929190612e30565b6000604051808303816000875af1158015610f95573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610fbd9190810190612e65565b9050846001600160a01b03167f17bc3181e17a9620a479c24e6c606e474ba84fc036877b768926872e8cd0e11f8c8c868d8d8760405161100296959493929190612edb565b60405180910390a25050505050505050505050565b61101f611dc1565b6001831461106f5760405162461bcd60e51b815260206004820152601960248201527f50726f706f73616c206c656e677468206d7573742062652031000000000000006044820152606401610635565b6110e98484600081811061108557611085612c42565b90506020028101906110979190612f2c565b6110a59060208101906127db565b858560008181106110b8576110b8612c42565b90506020028101906110ca9190612f2c565b6110db906040810190602001612f4c565b6001600160401b0316611678565b1515600114156111505760405162461bcd60e51b815260206004820152602c60248201527f4465706f73697420776974682070726f7669646564206e6f6e636520616c726560448201526b18591e48195e1958dd5d195960a21b6064820152608401610635565b7f6c1c3e083c3e1456662cc3e333f73180a70281c9b1cc48e77461bd614d406f36604051611197906020808252600590820152647072696a6560d81b604082015260600190565b60405180910390a16111ab8484848461172d565b6111f75760405162461bcd60e51b815260206004820152601760248201527f496e76616c69642070726f706f73616c207369676e65720000000000000000006044820152606401610635565b7f6c1c3e083c3e1456662cc3e333f73180a70281c9b1cc48e77461bd614d406f3660405161123e90602080825260059082015264726164696d60d81b604082015260600190565b60405180910390a16000600460008686600081811061125f5761125f612c42565b90506020028101906112719190612f2c565b60409081013582526020820192909252016000908120546001600160a01b0316915081868683816112a4576112a4612c42565b90506020028101906112b69190612f2c565b6112c4906060810190612f67565b6040516020016112d693929190612fad565b60408051601f198184030181529190528051602090910120905081610100878760008161130557611305612c42565b90506020028101906113179190612f2c565b611328906040810190602001612f4c565b6113329190612ca6565b6001600160401b03166001901b600660008989600081811061135657611356612c42565b90506020028101906113689190612f2c565b6113769060208101906127db565b60ff1660ff16815260200190815260200160002060006101008a8a60008181106113a2576113a2612c42565b90506020028101906113b49190612f2c565b6113c5906040810190602001612f4c565b6113cf9190612ce2565b6001600160401b0316815260200190815260200160002060008282541792505081905550806001600160a01b031663e248cff28888600081811061141557611415612c42565b90506020028101906114279190612f2c565b604001358989600081811061143e5761143e612c42565b90506020028101906114509190612f2c565b61145e906060810190612f67565b6040518463ffffffff1660e01b815260040161147c93929190612fd9565b600060405180830381600087803b15801561149657600080fd5b505af11580156114aa573d6000803e3d6000fd5b505050507f6018c584b8d99bafeda249b2429f5907d830e792222070c1b3a94aa76ee71677878760008181106114e2576114e2612c42565b90506020028101906114f49190612f2c565b6115029060208101906127db565b8888600081811061151557611515612c42565b90506020028101906115279190612f2c565b611538906040810190602001612f4c565b6040805160ff90931683526001600160401b039091166020830152810184905260600160405180910390a150505050505050565b6115846000356001600160e01b031916610bb6611e2b565b61159461158f611e2b565b611f49565b565b6115ae6000356001600160e01b031916610bb6611e2b565b600180546001600160a01b0319166001600160a01b0383169081179091556040519081527f729170bd142e4965055b26a285faeedf03baf2b915bfc5a7c75d24b45815ff2c90602001610b93565b6116146000356001600160e01b031916610bb6611e2b565b6040516307b7ed9960e01b81526001600160a01b0382811660048301528391908216906307b7ed99906024015b600060405180830381600087803b15801561165b57600080fd5b505af115801561166f573d6000803e3d6000fd5b50505050505050565b600061168661010083612ff3565b60ff84166000908152600660205260408120600190921b91906116ab61010086613007565b8152602001908152602001600020541660001415905092915050565b6116df6000356001600160e01b031916610bb6611e2b565b600280546001600160a01b0319166001600160a01b0383169081179091556040519081527f497acaa34ac19c2a2a579ad43eca493b4fea820459e254e9383e7dda216b0f0490602001610b93565b600080846001600160401b0381111561174857611748612535565b604051908082528060200260200182016040528015611771578160200160208202803683370190505b50905060005b858110156118f7577fcc13634e956dd3d4ec8d808ee8bf294e1cd05a38f63fe7f234b079a0a4c36a708787838181106117b2576117b2612c42565b90506020028101906117c49190612f2c565b6117d29060208101906127db565b8888848181106117e4576117e4612c42565b90506020028101906117f69190612f2c565b611807906040810190602001612f4c565b89898581811061181957611819612c42565b905060200281019061182b9190612f2c565b604001358a8a8681811061184157611841612c42565b90506020028101906118539190612f2c565b611861906060810190612f67565b60405161186f92919061301b565b6040519081900381206118b2959493929160200194855260ff9390931660208501526001600160401b039190911660408401526060830152608082015260a00190565b604051602081830303815290604052805190602001208282815181106118da576118da612c42565b6020908102919091010152806118ef81612d56565b915050611777565b5060006119bb85858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040516119b592507f989d14110ba109ccad392cc18511d1f6ae3a85165c5960e49d72c2c67682fde5915061196a90879060200161302b565b6040516020818303038152906040528051906020012060405160200161199a929190918252602082015260400190565b60405160208183030381529060405280519060200120611f97565b90611e07565b60005461010090046001600160a01b03908116911614979650505050505050565b6119f46000356001600160e01b031916610bb6611e2b565b60405163025a3c9960e21b815282906001600160a01b03821690630968f26490611641908590600401612d71565b611a3a6000356001600160e01b031916610bb6611e2b565b60008281526004602081905260409182902080546001600160a01b0319166001600160a01b038781169182179092559251635c7d1b9b60e11b81529182018590528316602482015284919063b8fa373690604401600060405180830381600087803b158015611aa857600080fd5b505af1158015611abc573d6000803e3d6000fd5b5050505050505050565b611ade6000356001600160e01b031916610bb6611e2b565b6001600160a01b03919091166000908152600560205260409020805460ff1916911515919091179055565b611b216000356001600160e01b031916610bb6611e2b565b6001600160a01b038116611b815760405162461bcd60e51b815260206004820152602160248201527f4d504320616464726573732063616e2774206265206e756c6c2d6164647265736044820152607360f81b6064820152608401610635565b60005461010090046001600160a01b031615611bdf5760405162461bcd60e51b815260206004820152601c60248201527f4d504320616464726573732063616e27742062652075706461746564000000006044820152606401610635565b60008054610100600160a81b0319166101006001600160a01b03841602179055611c0f611c0a611e2b565b611feb565b6040517f4187686ceef7b541a1f224d48d4cded8f2c535e0e58ac0f0514071b1de3dad5790600090a150565b611c536000356001600160e01b031916610bb6611e2b565b7fe78d813a9260522f81d6c761e311727b2e19008daadd2b9e174be86bc4f06a4b81604051610b939190612d71565b611c9a6000356001600160e01b031916610bb6611e2b565b60ff82166000908152600360205260409020546001600160401b0390811690821611611d175760405162461bcd60e51b815260206004820152602660248201527f446f6573206e6f7420616c6c6f772064656372656d656e7473206f6620746865604482015265206e6f6e636560d01b6064820152608401610635565b60ff919091166000908152600360205260409020805467ffffffffffffffff19166001600160401b03909216919091179055565b611d636000356001600160e01b031916610bb6611e2b565b60005461010090046001600160a01b0316611db65760405162461bcd60e51b8152602060048201526013602482015272135410c81859191c995cdcc81b9bdd081cd95d606a1b6044820152606401610635565b611594611c0a611e2b565b60005460ff16156115945760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610635565b6000806000611e168585612036565b91509150611e23816120a6565b509392505050565b60003360143610801590611e5757506001600160a01b03811660009081526005602052604090205460ff165b15611e67575060131936013560601c5b919050565b600254604051631c72548760e21b81526001600160e01b0319841660048201526001600160a01b038381166024830152909116906371c9521c90604401602060405180830381865afa158015611ec6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611eea9190613061565b611f455760405162461bcd60e51b815260206004820152602660248201527f73656e64657220646f65736e277420686176652061636365737320746f2066756044820152653731ba34b7b760d11b6064820152608401610635565b5050565b611f51611dc1565b6000805460ff191660011790556040516001600160a01b03821681527f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25890602001610b93565b6000611fe5611fa4612264565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b92915050565b611ff361238b565b6000805460ff191690556040516001600160a01b03821681527f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa90602001610b93565b60008082516041141561206d5760208301516040840151606085015160001a612061878285856123d4565b9450945050505061209f565b825160401415612097576020830151604084015161208c8683836124c1565b93509350505061209f565b506000905060025b9250929050565b60008160048111156120ba576120ba61307e565b14156120c35750565b60018160048111156120d7576120d761307e565b14156121255760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610635565b60028160048111156121395761213961307e565b14156121875760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610635565b600381600481111561219b5761219b61307e565b14156121f45760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610635565b60048160048111156122085761220861307e565b14156122615760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610635565b50565b6000306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480156122bd57507f000000000000000000000000000000000000000000000000000000000000000046145b156122e757507f000000000000000000000000000000000000000000000000000000000000000090565b50604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091527f0000000000000000000000000000000000000000000000000000000000000000828401527f000000000000000000000000000000000000000000000000000000000000000060608301524660808301523060a0808401919091528351808403909101815260c0909201909252805191012090565b60005460ff166115945760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610635565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561240b57506000905060036124b8565b8460ff16601b1415801561242357508460ff16601c14155b1561243457506000905060046124b8565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015612488573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166124b1576000600192509250506124b8565b9150600090505b94509492505050565b6000806001600160ff1b038316816124de60ff86901c601b613094565b90506124ec878288856123d4565b935093505050935093915050565b803560ff81168114611e6757600080fd5b6000806040838503121561251e57600080fd5b612527836124fa565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b038111828210171561256d5761256d612535565b60405290565b604051601f8201601f191681016001600160401b038111828210171561259b5761259b612535565b604052919050565b80356001600160401b0381168114611e6757600080fd5b60006001600160401b038211156125d3576125d3612535565b50601f01601f191660200190565b60006125f46125ef846125ba565b612573565b905082815283838301111561260857600080fd5b828260208301376000602084830101529392505050565b600082601f83011261263057600080fd5b61263f838335602085016125e1565b9392505050565b6000806040838503121561265957600080fd5b82356001600160401b038082111561267057600080fd5b818501915085601f83011261268457600080fd5b813560208282111561269857612698612535565b8160051b6126a7828201612573565b928352848101820192828101908a8511156126c157600080fd5b83870192505b8483101561275d578235868111156126df5760008081fd5b87016080818d03601f19018113156126f75760008081fd5b6126ff61254b565b61270a8784016124fa565b8152612718604084016125a3565b81880152606083810135604083015291830135918983111561273a5760008081fd5b6127488f898587010161261f565b908201528452505091830191908301906126c7565b975050508601359250508082111561277457600080fd5b506127818582860161261f565b9150509250929050565b60006020828403121561279d57600080fd5b81356001600160401b038111156127b357600080fd5b8201601f810184136127c457600080fd5b6127d3848235602084016125e1565b949350505050565b6000602082840312156127ed57600080fd5b61263f826124fa565b80356001600160a01b0381168114611e6757600080fd5b80356001600160e01b031981168114611e6757600080fd5b60008060008060008060c0878903121561283e57600080fd5b612847876127f6565b95506020870135945061285c604088016127f6565b935061286a6060880161280d565b92506080870135915061287f60a0880161280d565b90509295509295509295565b60008083601f84011261289d57600080fd5b5081356001600160401b038111156128b457600080fd5b60208301915083602082850101111561209f57600080fd5b600080600080600080608087890312156128e557600080fd5b6128ee876124fa565b95506020870135945060408701356001600160401b038082111561291157600080fd5b61291d8a838b0161288b565b9096509450606089013591508082111561293657600080fd5b5061294389828a0161288b565b979a9699509497509295939492505050565b6000806000806040858703121561296b57600080fd5b84356001600160401b038082111561298257600080fd5b818701915087601f83011261299657600080fd5b8135818111156129a557600080fd5b8860208260051b85010111156129ba57600080fd5b6020928301965094509086013590808211156129d557600080fd5b506129e28782880161288b565b95989497509550505050565b600060208284031215612a0057600080fd5b5035919050565b600060208284031215612a1957600080fd5b61263f826127f6565b60008060408385031215612a3557600080fd5b612a3e836127f6565b9150612a4c602084016127f6565b90509250929050565b60008060408385031215612a6857600080fd5b612a71836127f6565b915060208301356001600160401b03811115612a8c57600080fd5b6127818582860161261f565b600080600060608486031215612aad57600080fd5b612ab6846127f6565b925060208401359150612acb604085016127f6565b90509250925092565b801515811461226157600080fd5b60008060408385031215612af557600080fd5b612afe836127f6565b91506020830135612b0e81612ad4565b809150509250929050565b60008060408385031215612b2c57600080fd5b612b35836124fa565b9150612a4c602084016125a3565b60005b83811015612b5e578181015183820152602001612b46565b83811115610b5e5750506000910152565b60008151808452612b87816020860160208601612b43565b601f01601f19169290920160200192915050565b60006040808301818452808651808352606092508286019150828160051b8701016020808a0160005b84811015612c2257898403605f190186528151805160ff168552838101516001600160401b03168486015288810151898601528701516080888601819052612c0e81870183612b6f565b978501979550505090820190600101612bc4565b5050819650612c358189018a60ff169052565b5050505050509392505050565b634e487b7160e01b600052603260045260246000fd5b6bffffffffffffffffffffffff198360601b16815260008251612c82816014850160208701612b43565b919091016014019392505050565b634e487b7160e01b600052601260045260246000fd5b60006001600160401b0380841680612cc057612cc0612c90565b92169190910692915050565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b0380841680612cfc57612cfc612c90565b92169190910492915050565b8281526040602082015260006127d36040830184612b6f565b606081526000612d346060830186612b6f565b905060ff841660208301526001600160401b0383166040830152949350505050565b6000600019821415612d6a57612d6a612ccc565b5060010190565b60208152600061263f6020830184612b6f565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60018060a01b038916815260ff8816602082015260ff8716604082015285606082015260c060808201526000612de760c083018688612d84565b82810360a0840152612dfa818587612d84565b9b9a5050505050505050505050565b60006001600160401b0380831681811415612e2657612e26612ccc565b6001019392505050565b8481526001600160a01b0384166020820152606060408201819052600090612e5b9083018486612d84565b9695505050505050565b600060208284031215612e7757600080fd5b81516001600160401b03811115612e8d57600080fd5b8201601f81018413612e9e57600080fd5b8051612eac6125ef826125ba565b818152856020838501011115612ec157600080fd5b612ed2826020830160208601612b43565b95945050505050565b60ff871681528560208201526001600160401b038516604082015260a060608201526000612f0d60a083018587612d84565b8281036080840152612f1f8185612b6f565b9998505050505050505050565b60008235607e19833603018112612f4257600080fd5b9190910192915050565b600060208284031215612f5e57600080fd5b61263f826125a3565b6000808335601e19843603018112612f7e57600080fd5b8301803591506001600160401b03821115612f9857600080fd5b60200191503681900382131561209f57600080fd5b6bffffffffffffffffffffffff198460601b168152818360148301376000910160140190815292915050565b838152604060208201526000612ed2604083018486612d84565b60008261300257613002612c90565b500690565b60008261301657613016612c90565b500490565b8183823760009101908152919050565b815160009082906020808601845b8381101561305557815185529382019390820190600101613039565b50929695505050505050565b60006020828403121561307357600080fd5b815161263f81612ad4565b634e487b7160e01b600052602160045260246000fd5b600082198211156130a7576130a7612ccc565b50019056fea264697066735822122063202b5939495642556f05eefeaf15cd6680e27fdf52bf07048030504e9b500364736f6c634300080b0033"
