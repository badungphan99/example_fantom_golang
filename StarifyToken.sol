pragma solidity >=0.7.0 <0.9.0;

library SafeMath {
  function sub(uint256 a, uint256 b) internal pure returns (uint256) {
    assert(b <= a);
    return a - b;
  }

  function add(uint256 a, uint256 b) internal pure returns (uint256) {
    uint256 c = a + b;
    assert(c >= a);
    return c;
  }
}

contract StarifyToken {
  string public constant name = "StarifyToken";
  string public constant symbol = "SCash";
  uint8 public constant decimals = 18;

  event Transfer(address indexed _from, address indexed _to, uint _tokens);

  using SafeMath for uint256;

  mapping(address => uint256) balances;

  address owner;

  uint256 _totalSupply;

  constructor(uint256 total) public {
    owner = msg.sender;
    _totalSupply = total;
    balances[owner] = _totalSupply;
  }

  function mint(uint256 amount) public returns (bool){
    require(msg.sender == owner);
    balances[owner] = balances[owner].add(amount);
    _totalSupply = _totalSupply.add(amount);
    return true;
  }

  function totalSupply() public view returns(uint256) {
    return _totalSupply;
  }

  function balanceOF(address tokenOwner) public view returns (uint256){
    return balances[tokenOwner];
  }

  function transfer(address receiver, uint256 amount) public returns (bool) {
    require(amount <= balances[msg.sender]);
    balances[msg.sender] = balances[msg.sender].sub(amount);
    balances[receiver] = balances[receiver].add(amount);
    emit Transfer(msg.sender, receiver, amount);
    return true;
  }
}