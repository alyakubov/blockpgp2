pragma solidity ^0.4.22;
import "./lfingers.sol";
import "./lrevoc.sol";
import "./lksrv.sol";
import "./cbase.sol";

// https://medium.com/daox/avoiding-out-of-gas-error-in-large-ethereum-smart-contracts-18961b1fc0c6


/*interface Interf_iden {
    function getOrchestrAddr() external view returns(address);

    function getIdenOwnCert(bytes _finger) external view returns(bytes);
    function getIdenLoadDate(bytes _finger) external view returns(uint);
    function getIdenEmail(bytes _finger) external view returns(string);
    function getIdenName(bytes _finger) external view returns(string);
    function getIdenCertUser(bytes _finger) external view returns(address);

    function getSigntLen(bytes _finger) external view returns(uint len, int err);
    function newCertificate(string _email, bytes _finger, bytes _ownCert, string _name, address _certUser) external;
    function revokeSignt( bytes _finger, bytes _signedFinger, uint _revocationDate) external;
    function getSignedFinger(bytes _finger, uint _ind ) external view
                returns(bytes _signedFinger, bytes _proposedCert, uint _revocationDate, int _err);
    function acceptProposedCert(bytes _introducingFinger, bytes _finger) external;
    function newSignt(bytes _finger, bytes _signedFinger, bytes _proposedCert) external;
}*/


contract LU_PGP_orchestr is LU_PGP_base {

    //LU_PGP_fingers internal cFingers;
    //LU_PGP_revoc internal cRevoc;
    //LU_PGP_ksrv internal cKsrv;
    //LU_PGP_iden internal cIden;

    //address internal cFingersAddr;
    //address internal cRevocAddr;
    //address internal cKsrvAddr;
    address internal cIdenAddr;
    function getIdenAddr() external view returns(address) {
        return cIdenAddr;
    }

    /*function getFingersAddr() external view returns(address) {
        return cFingersAddr;
    }
    function getRevocAddr() external view returns(address) {
        return cRevocAddr;
    }
    function getKsrvAddr() external view returns(address) {
        return cKsrvAddr;
    }*/

    /*
    using LIB_iden for LIB_iden.DaIden;
    LIB_iden.DaIden internal cIden;
    */

    using LIB_fingers for LIB_fingers.DaFingers;
    LIB_fingers.DaFingers public cFingers;

    using LIB_revoc for LIB_revoc.DaRevoc;
    LIB_revoc.DaRevoc public cRevoc;

    using LIB_ksrv for LIB_ksrv.DaKsrv;
    LIB_ksrv.DaKsrv public cKsrv;

    function InitOrchestr(address _cOrchestrAddr, address _cOridenAddr) external onlyOwner //setOrchestrAddr(_cOrchestrAddr)
    {
        cOrchestrAddr = _cOrchestrAddr; //-- already in modifier
        cIdenAddr = _cOridenAddr;

        /*
        cIden.Init(_cOrchestrAddr);
        //cIdenAddr = _cIdenAddr;
        //cIden = Interf_iden(cIdenAddr);
        */

        cFingers.Init(_cOrchestrAddr, _cOridenAddr);
        //cFingersAddr = _cFingersAddr;
        //cFingers = LU_PGP_fingers(cFingersAddr);

        cRevoc.Init(_cOrchestrAddr, _cOridenAddr);
        //cRevocAddr = _cRevocAddr;
        //cRevoc = LU_PGP_revoc(cRevocAddr);

        cKsrv.Init(_cOrchestrAddr, _cOridenAddr);
        //cKsrvAddr = _cKsrvAddr;
        //cKsrv = LU_PGP_ksrv(cKsrvAddr);
    }

    modifier checkRef {

        require(cOrchestrAddr != 0);
        //require(cFingersAddr != 0);
        //require(cRevocAddr != 0);
        //require(cKsrvAddr != 0);
        //require(cIdenAddr != 0);
        /*
        require(cIden.getOrchestrAddr() == cOrchestrAddr);
        */
        require(cFingers.getOrchestrAddr() == cOrchestrAddr);
        require(cRevoc.getOrchestrAddr() == cOrchestrAddr);
        require(cKsrv.getOrchestrAddr() == cOrchestrAddr);
        _;
    }

    /*function Init(address _cPgpAddr, address _cFingersAddr, address _cRevocAddr, address _cKsrvAddr, 
            address _cIdenAddr, address _cUtilsAddr) 
            public onlyOwner setAddr(_cPgpAddr, _cFingersAddr, _cRevocAddr, _cKsrvAddr, _cIdenAddr, _cUtilsAddr)
    {
    }*/

    function getEvIndCertId(bytes _finger) external pure returns(uint) {
        return uint( keccak256( _finger ) );
    }
    function getEvIndSignId(bytes _finger, bytes _signedFinger) external pure returns(uint) {
        return uint( keccak256( LIB_fingers.concateBytes(_finger, _signedFinger) ) );
    }

    /* Shell functions  */
    function getCertRevocation(bytes _finger) external view checkRef returns(bool isRevoked, uint revDate) {
        return cRevoc.getCertRevocation(_finger);
    }
    function getFingersLen(string _email) external view checkRef returns(uint len, int err) {
        return cFingers.getFingersLen( _email );
    }
    function isFingersFinger(string _email, bytes _finger) external view checkRef returns(uint _ind, int _err) { 
        return cFingers.isFingersFinger( _email, _finger );
    }
    function newFinger(string _email, bytes _finger) external checkRef {
        cFingers.newFinger(_email,_finger);
    }
    function revokeCertificate(bytes _finger, uint _revocationDate) external checkRef {
        cRevoc.revokeCertificate(_finger, _revocationDate);
    }
    function revokeSignt( bytes _finger, bytes _signedFinger, uint _revocationDate) external checkRef {
        cRevoc.revokeSignt( _finger, _signedFinger, _revocationDate);
    }
    function getFingersItem(string _email, uint _ind) external view checkRef returns(bytes _res, int _err ) { 
        return cFingers.getFingersItem(_email, _ind);
    }
    function findSigntIndex( bytes _finger, bytes _signedFinger) external view checkRef
             returns(uint _ind, bytes _proposedCert, int _err, string _errMsg) {
        return cFingers.findSigntIndex( _finger, _signedFinger);
    }
    function getSigntRevocation(bytes _finger, uint _ind) external view checkRef returns(bool isRevoked, uint revDate) { 
        return cRevoc.getSigntRevocation(_finger, _ind);
    }
    function getSignRevokeInd(bytes _finger, uint _arrInd) external pure returns(uint) {
        return LIB_fingers.getSignRevokeInd(_finger, _arrInd);
    }

    function checkRights(bytes _finger, address _certUser) external view checkRef returns(int err, string errMsg) {
        return cFingers.checkRights(_finger, _certUser);
    }
    function isKsrvCertFound( bytes _finger, bytes _proposedCert) external view checkRef
             returns(uint isFound, uint _ind, int _err, string _errMsg){
        return cKsrv.isKsrvCertFound( _finger, _proposedCert);
    }

    /*function extractCertFromArmored(bytes armoured) external pure returns (bytes extract) {
        return LIB_fingers.extractCertFromArmored(armoured);
    }*/
    function calcCertHash(bytes _armouredCert) pure external returns(uint) {
        return LIB_fingers.calcCertHash(_armouredCert);
        /*bytes memory extract = LIB_fingers.extractCertFromArmored(_armouredCert);
        if (extract.length == 0) {
            return 0;
        }
        return uint(sha256(extract));*/
    }

    /*function newLog(bytes _str, bytes32 _uparam, bytes32 _iparam, address _addr) external {
        cFingers.newLog(_str, _uparam, _iparam, _addr);
    }
    function getLog(uint _ind) external view returns(bytes, bytes32, bytes32) {
        return cFingers.getLog(_ind);
    }*/


    /*
    function newCertificate(string _email, bytes _finger, bytes _ownCert, string _name, address _certUser)
            external checkRef {
        cIden.newCertificate(_email, _finger, _ownCert, _name, _certUser);
    }
    function newSignt(bytes _finger, bytes _signedFinger, bytes _proposedCert) external checkRef {
        cIden.newSignt( _finger, _signedFinger, _proposedCert);
    }
    function getSigntLen(bytes _finger) external view checkRef returns(uint len, int err) { 
        return cIden.getSigntLen(_finger);
    }
    function getSignedFinger(bytes _finger, uint _ind ) external view checkRef
            returns(bytes _signedFinger, bytes _proposedCert, uint _revocationDate, int _err) {
        return cIden.getSignedFinger(_finger,_ind);
    }
    function acceptProposedCert(bytes _introducingFinger, bytes _finger) external checkRef {
        cIden.acceptProposedCert(_introducingFinger, _finger);
    }
    function getIdenOwnCert(bytes _finger) external view checkRef returns(bytes) {
        return cIden.getIdenOwnCert(_finger);
    }
    function getIdenName(bytes _finger) external view checkRef returns(string) { 
        return cIden.getIdenName(_finger);
    }
    function getIdenLoadDate(bytes _finger) external view checkRef returns(uint) { 
        return cIden.getIdenLoadDate(_finger);
    }
    function getIdenEmail(bytes _finger) external view checkRef returns(string) { 
        return cIden.getIdenEmail(_finger);
    }
    function getIdenCertUser(bytes _finger) external view checkRef returns(address) { 
        return cIden.getIdenCertUser(_finger); 
    }
    */

}

