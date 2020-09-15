  // Copyright 2019 The Swarm Authors
  // This file is part of the Swarm library.
  //
  // The Swarm library is free software: you can redistribute it and/or modify
  // it under the terms of the GNU Lesser General Public License as published by
  // the Free Software Foundation, either version 3 of the License, or
  // (at your option) any later version.
  //
  // The Swarm library is distributed in the hope that it will be useful,
  // but WITHOUT ANY WARRANTY; without even the implied warranty of
  // MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
  // GNU Lesser General Public License for more details.
  //
  // You should have received a copy of the GNU Lesser General Public License
  // along with the Swarm library. If not, see <http://www.gnu.org/licenses/>.
  //
  // Code generated - DO NOT EDIT.
  // This file was autogenerated with 'npm run abigen' from ethersphere/swap-swear-and-swindle and any manual changes will be lost.
  package erc20simpleswap

  // ERC20SimpleSwapDeployedCode is the bytecode ERC20SimpleSwap will have after deployment
  const ERC20SimpleSwapDeployedCode = "0x608060405234801561001057600080fd5b50600436106101165760003560e01c8063b6343b0d116100a2578063b7ec1a3311610071578063b7ec1a331461049e578063c76a4d31146104a6578063d4c9a8e8146104cc578063e0bcf13a14610585578063fc0c546a1461058d57610116565b8063b6343b0d14610402578063b648b4171461044e578063b69ef8a81461046a578063b77703501461047257610116565b80631d143848116100e95780631d143848146103495780632e1a7d4d1461036d578063338f3fed1461038a57806381f03fcb146103b6578063946f46a2146103dc57610116565b80630d5f26591461011b57806312101021146101d65780631357e1dc146101f05780631633fb1d146101f8575b600080fd5b6101d46004803603606081101561013157600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561016057600080fd5b82018360208201111561017257600080fd5b803590602001918460018302840111600160201b8311171561019357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610595945050505050565b005b6101de6105a8565b60408051918252519081900360200190f35b6101de6105ae565b6101d4600480360360c081101561020e57600080fd5b6001600160a01b03823581169260208101359091169160408201359190810190608081016060820135600160201b81111561024857600080fd5b82018360208201111561025a57600080fd5b803590602001918460018302840111600160201b8311171561027b57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092958435959094909350604081019250602001359050600160201b8111156102d557600080fd5b8201836020820111156102e757600080fd5b803590602001918460018302840111600160201b8311171561030857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506105b4945050505050565b61035161062e565b604080516001600160a01b039092168252519081900360200190f35b6101d46004803603602081101561038357600080fd5b503561063d565b6101d4600480360360408110156103a057600080fd5b506001600160a01b03813516906020013561079e565b6101de600480360360208110156103cc57600080fd5b50356001600160a01b03166108ca565b6101d4600480360360208110156103f257600080fd5b50356001600160a01b03166108dc565b6104286004803603602081101561041857600080fd5b50356001600160a01b03166109b7565b604080519485526020850193909352838301919091526060830152519081900360800190f35b6104566109de565b604080519115158252519081900360200190f35b6101de6109ee565b6101d46004803603604081101561048857600080fd5b506001600160a01b038135169060200135610a6a565b6101de610b8c565b6101de600480360360208110156104bc57600080fd5b50356001600160a01b0316610ba7565b6101d4600480360360608110156104e257600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561051157600080fd5b82018360208201111561052357600080fd5b803590602001918460018302840111600160201b8311171561054457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610bd8945050505050565b6101de610ceb565b610351610cf1565b6105a3338484600085610d00565b505050565b60005481565b60035481565b6105ca6105c43033878987611105565b84611166565b6001600160a01b0316866001600160a01b0316146106195760405162461bcd60e51b81526004018080602001828103825260228152602001806115ae6022913960400191505060405180910390fd5b6106268686868585610d00565b505050505050565b6006546001600160a01b031681565b6006546001600160a01b03163314610695576040805162461bcd60e51b815260206004820152601660248201527529b4b6b83632a9bbb0b81d103737ba1034b9b9bab2b960511b604482015290519081900360640190fd5b61069d610b8c565b8111156106db5760405162461bcd60e51b81526004018080602001828103825260288152602001806116656028913960400191505060405180910390fd5b6001546006546040805163a9059cbb60e01b81526001600160a01b039283166004820152602481018590529051919092169163a9059cbb9160448083019260209291908290030181600087803b15801561073457600080fd5b505af1158015610748573d6000803e3d6000fd5b505050506040513d602081101561075e57600080fd5b505161079b5760405162461bcd60e51b81526004018080602001828103825260278152602001806116176027913960400191505060405180910390fd5b50565b6006546001600160a01b031633146107f6576040805162461bcd60e51b815260206004820152601660248201527529b4b6b83632a9bbb0b81d103737ba1034b9b9bab2b960511b604482015290519081900360640190fd5b6107fe6109ee565b60055461080b9083611181565b11156108485760405162461bcd60e51b81526004018080602001828103825260348152602001806115586034913960400191505060405180910390fd5b6001600160a01b0382166000908152600460205260409020805461086c9083611181565b815560055461087b9083611181565b60055560006003820155805460408051918252516001600160a01b038516917f2506c43272ded05d095b91dbba876e66e46888157d3e078db5691496e96c5fad919081900360200190a2505050565b60026020526000908152604090205481565b6001600160a01b03811660009081526004602052604090206003810154421080159061090b5750600381015415155b6109465760405162461bcd60e51b81526004018080602001828103825260258152602001806115f26025913960400191505060405180910390fd5b60018101548154610956916111db565b8155600060038201556001810154600554610970916111db565b600555805460408051918252516001600160a01b038416917f2506c43272ded05d095b91dbba876e66e46888157d3e078db5691496e96c5fad919081900360200190a25050565b60046020526000908152604090208054600182015460028301546003909301549192909184565b600654600160a01b900460ff1681565b600154604080516370a0823160e01b815230600482015290516000926001600160a01b0316916370a08231916024808301926020929190829003018186803b158015610a3957600080fd5b505afa158015610a4d573d6000803e3d6000fd5b505050506040513d6020811015610a6357600080fd5b5051905090565b6006546001600160a01b03163314610ac2576040805162461bcd60e51b815260206004820152601660248201527529b4b6b83632a9bbb0b81d103737ba1034b9b9bab2b960511b604482015290519081900360640190fd5b6001600160a01b03821660009081526004602052604090208054821115610b1a5760405162461bcd60e51b815260040180806020018281038252602781526020018061163e6027913960400191505060405180910390fd5b60008160020154600014610b32578160020154610b36565b6000545b4281016003840155600183018490556040805185815290519192506001600160a01b038616917fc8305077b495025ec4c1d977b176a762c350bb18cad4666ce1ee85c32b78698a9181900360200190a250505050565b6000610ba2600554610b9c6109ee565b906111db565b905090565b6001600160a01b038116600090815260046020526040812054610bd290610bcc610b8c565b90611181565b92915050565b6006546001600160a01b03163314610c30576040805162461bcd60e51b815260206004820152601660248201527529b4b6b83632a9bbb0b81d103737ba1034b9b9bab2b960511b604482015290519081900360640190fd5b610c44610c3e30858561121d565b82611166565b6001600160a01b0316836001600160a01b031614610c935760405162461bcd60e51b81526004018080602001828103825260228152602001806115ae6022913960400191505060405180910390fd5b6001600160a01b038316600081815260046020908152604091829020600201859055815185815291517f7b816003a769eb718bd9c66bdbd2dd5827da3f92bc6645276876bd7957b08cf09281900390910190a2505050565b60055481565b6001546001600160a01b031681565b6006546001600160a01b03163314610d8257610d20610c3e30878661121d565b6006546001600160a01b03908116911614610d82576040805162461bcd60e51b815260206004820152601d60248201527f53696d706c65537761703a20696e76616c696420697373756572536967000000604482015290519081900360640190fd5b6001600160a01b038516600090815260026020526040812054610da69085906111db565b90506000610dbc82610db789610ba7565b61126e565b6001600160a01b03881660009081526004602052604081205491925090610de490839061126e565b905084821015610e3b576040805162461bcd60e51b815260206004820152601d60248201527f53696d706c65537761703a2063616e6e6f74207061792063616c6c6572000000604482015290519081900360640190fd5b8015610e8e576001600160a01b038816600090815260046020526040902054610e6490826111db565b6001600160a01b038916600090815260046020526040902055600554610e8a90826111db565b6005555b6001600160a01b038816600090815260026020526040902054610eb19083611181565b6001600160a01b038916600090815260026020526040902055600354610ed79083611181565b6003556001546001600160a01b031663a9059cbb88610ef685896111db565b6040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015610f3c57600080fd5b505af1158015610f50573d6000803e3d6000fd5b505050506040513d6020811015610f6657600080fd5b5051610fa35760405162461bcd60e51b81526004018080602001828103825260278152602001806116176027913960400191505060405180910390fd5b8415611064576001546040805163a9059cbb60e01b81523360048201526024810188905290516001600160a01b039092169163a9059cbb916044808201926020929091908290030181600087803b158015610ffd57600080fd5b505af1158015611011573d6000803e3d6000fd5b505050506040513d602081101561102757600080fd5b50516110645760405162461bcd60e51b81526004018080602001828103825260278152602001806116176027913960400191505060405180910390fd5b6040805183815260208101889052808201879052905133916001600160a01b038a811692908c16917f950494fc3642fae5221b6c32e0e45765c95ebb382a04a71b160db0843e74c99f919081900360600190a48183146110fb576006805460ff60a01b1916600160a01b1790556040517f3f4449c047e11092ec54dc0751b6b4817a9162745de856c893a26e611d18ffc490600090a15b5050505050505050565b604080516bffffffffffffffffffffffff19606097881b811660208084019190915296881b8116603483015260488201959095529290951b9092166068820152607c8082019290925283518082039092018252609c01909252815191012090565b600061117a61117484611284565b836112d5565b9392505050565b60008282018381101561117a576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b600061117a83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506114c0565b604080516bffffffffffffffffffffffff19606095861b81166020808401919091529490951b9094166034850152604880850192909252805180850390920182526068909301909252815191012090565b600081831061127d578161117a565b5090919050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b6000815160411461132d576040805162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015290519081900360640190fd5b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561139e5760405162461bcd60e51b815260040180806020018281038252602281526020018061158c6022913960400191505060405180910390fd5b8060ff16601b141580156113b657508060ff16601c14155b156113f25760405162461bcd60e51b81526004018080602001828103825260228152602001806115d06022913960400191505060405180910390fd5b600060018783868660405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa15801561144e573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166114b6576040805162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015290519081900360640190fd5b9695505050505050565b6000818484111561154f5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156115145781810151838201526020016114fc565b50505050905090810190601f1680156115415780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50505090039056fe53696d706c65537761703a2068617264206465706f7369742063616e6e6f74206265206d6f7265207468616e2062616c616e636545434453413a20696e76616c6964207369676e6174757265202773272076616c756553696d706c65537761703a20696e76616c69642062656e656669636961727953696745434453413a20696e76616c6964207369676e6174757265202776272076616c756553696d706c65537761703a206465706f736974206e6f74207965742074696d6564206f757453696d706c65537761703a2053696d706c65537761703a207472616e73666572206661696c656453696d706c65537761703a2068617264206465706f736974206e6f742073756666696369656e7453696d706c65537761703a206c697175696442616c616e6365206e6f742073756666696369656e74a2646970667358221220c21fb187057a5c31dbbedbbf220875432d93aebab5beb8ed795e9e04676952f464736f6c634300060c0033"