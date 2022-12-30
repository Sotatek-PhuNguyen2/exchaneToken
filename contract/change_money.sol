// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract change_money {
    address private owner;

    constructor (){
        owner = msg.sender;
    }

    event NewTransactionExchange (address, address, uint256, uint256);
    event NewTransactionReceived(address);

    mapping (bytes => bool) public isReceived;
    uint256 private nonce;

    function getMessageHash(
        address _to,
        uint256 _amount,
        string memory network,
        uint256 _nonce
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_to, _amount, network, _nonce));
    }

    function getEthSignedMessageHash(bytes32 _messageHash)
        public
        pure
        returns (bytes32)
    {
        return
            keccak256(
                abi.encodePacked("\x19Ethereum Signed Message:\n32", _messageHash)
            );
    }

    /* 4. Verify signature */
    function verify(
        address _to,
        bytes memory signature,
        uint256 _value,
        string memory network,
        uint256 _nonce
    ) public view returns (bool) {
        bytes32 messageHash = getMessageHash(_to, _value, network, _nonce);
        bytes32 ethSignedMessageHash = getEthSignedMessageHash(messageHash);
        return recoverSigner(ethSignedMessageHash, signature) == owner;
    }

    function recoverSigner(bytes32 _ethSignedMessageHash, bytes memory _signature)
        private
        pure
        returns (address)
    {
        (bytes32 r, bytes32 s, uint8 v) = splitSignature(_signature);

        return ecrecover(_ethSignedMessageHash, v, r, s);
    }

    function splitSignature(bytes memory sig)
        private
        pure
        returns (
            bytes32 r,
            bytes32 s,
            uint8 v
        )
    {
        require(sig.length == 65, "invalid signature length");

        assembly {
            /*
            First 32 bytes stores the length of the signature

            add(sig, 32) = pointer of sig + 32
            effectively, skips first 32 bytes of signature

            mload(p) loads next 32 bytes starting at the memory address p into memory
            */

            // first 32 bytes, after the length prefix
            r := mload(add(sig, 32))
            // second 32 bytes
            s := mload(add(sig, 64))
            // final byte (first byte of the next 32 bytes)
            v := byte(0, mload(add(sig, 96)))
        }

        // implicitly return (r, s, v)
    }

    function sendTokenToContract(address _address) public payable {
        require(msg.value >= 0.00001 ether, "Failed to send, require value >= 0.00001");
        nonce++;
        emit NewTransactionExchange(msg.sender,_address, msg.value, nonce);
    }

    function receiveToken(bytes memory _signature,address _address,string memory network, uint256 _value, uint256 _nonce) public payable{
        require (verify( _address , _signature, _value, network, _nonce) == true, "Authentication failed");
        require (isReceived[_signature]==false,"Address is received");
        (bool sent,) = _address.call{value: _value}("");
        require(sent, "Failed to send");
        isReceived[_signature] = true;
        emit NewTransactionReceived(_address);
    }
} 