//
// Copyright 2021, Offchain Labs, Inc. All rights reserved.
//

package precompiles

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"math/big"
)

type ArbOwner struct{}

func (con ArbOwner) AddAllowedSender(caller common.Address, st *state.StateDB, addr common.Address) error {
	return errors.New("unimplemented")
}

func (con ArbOwner) AddAllowedSenderGasCost(addr common.Address) *big.Int {
	return nil
}

func (con ArbOwner) AddChainOwner(caller common.Address, st *state.StateDB, newOwner common.Address) error {
	return errors.New("unimplemented")
}

func (con ArbOwner) AddChainOwnerGasCost(newOwner common.Address) *big.Int {
	return nil
}

func (con ArbOwner) AllowAllSenders(caller common.Address, st *state.StateDB) error {
	return errors.New("unimplemented")
}

func (con ArbOwner) AllowAllSendersGasCost() *big.Int {
	return nil
}

func (con ArbOwner) AddMappingException(caller common.Address, st *state.StateDB, from *big.Int, to *big.Int) error {
	return errors.New("unimplemented")
}

func (con ArbOwner) AddMappingExceptionGasCost(from *big.Int, to *big.Int) *big.Int {
	return nil
}

func (con ArbOwner) AllowOnlyOwnerToSend(caller common.Address, st *state.StateDB) error {
	return errors.New("unimplemented")
}

func (con ArbOwner) AllowOnlyOwnerToSendGasCost() *big.Int {
	return nil
}

func (con ArbOwner) AddToReserveFunds(caller common.Address, st *state.StateDB, value *big.Int) error {
	return errors.New("unimplemented")
}

func (con ArbOwner) AddToReserveFundsGasCost() *big.Int {
	return nil
}

func (con ArbOwner) ContinueCodeUpload(caller common.Address, st *state.StateDB, marshalledCode []byte) error {
	return errors.New("unimplemented")
}

func (con ArbOwner) ContinueCodeUploadGasCost(marshalledCode []byte) *big.Int {
	return nil
}

func (con ArbOwner) CreateChainParameter(
	caller common.Address,
	st *state.StateDB,
	which [32]byte,
	value *big.Int,
) error {
	return errors.New("unimplemented")
}

func (con ArbOwner) CreateChainParameterGasCost(which [32]byte, value *big.Int) *big.Int {
	return nil
}

func (con ArbOwner) DeployContract(
	caller common.Address,
	st *state.StateDB,
	value *big.Int,
	constructorData []byte,
	deemedSender common.Address,
	deemedNonce *big.Int,
) (common.Address, error) {
	return common.Address{}, errors.New("unimplemented")
}

func (con ArbOwner) DeployContractGasCost(
	constructorData []byte,
	deemedSender common.Address,
	deemedNonce *big.Int,
) *big.Int {
	return nil
}

func (con ArbOwner) FinishCodeUploadAsArbosUpgrade(
	caller common.Address,
	st *state.StateDB,
	newCodeHash [32]byte,
	oldCodeHash [32]byte,
) error {
	return errors.New("unimplemented")
}

func (con ArbOwner) FinishCodeUploadAsArbosUpgradeGasCost(newCodeHash [32]byte, oldCodeHash [32]byte) *big.Int {
	return nil
}

func (con ArbOwner) GetAllAllowedSenders(caller common.Address, st *state.StateDB) ([]byte, error) {
	return nil, errors.New("unimplemented")
}

func (con ArbOwner) GetAllAllowedSendersGasCost() *big.Int {
	return nil
}

func (con ArbOwner) GetAllChainOwners(caller common.Address, st *state.StateDB) ([]byte, error) {
	return nil, errors.New("unimplemented")
}

func (con ArbOwner) GetAllChainOwnersGasCost() *big.Int {
	return nil
}

func (con ArbOwner) GetAllFairGasPriceSenders(caller common.Address, st *state.StateDB) ([]byte, error) {
	return nil, errors.New("unimplemented")
}

func (con ArbOwner) GetAllFairGasPriceSendersGasCost() *big.Int {
	return nil
}

func (con ArbOwner) GetAllMappingExceptions(caller common.Address, st *state.StateDB) ([]byte, error) {
	return nil, errors.New("unimplemented")
}

func (con ArbOwner) GetAllMappingExceptionsGasCost() *big.Int {
	return nil
}

func (con ArbOwner) GetChainParameter(caller common.Address, st *state.StateDB, which [32]byte) (*big.Int, error) {
	return nil, errors.New("unimplemented")
}

func (con ArbOwner) GetChainParameterGasCost(which [32]byte) *big.Int {
	return nil
}

func (con ArbOwner) GetTotalOfEthBalances(caller common.Address, st *state.StateDB) (*big.Int, error) {
	return nil, errors.New("unimplemented")
}

func (con ArbOwner) GetTotalOfEthBalancesGasCost() *big.Int {
	return nil
}

func (con ArbOwner) GetLastUpgradeHash(caller common.Address, st *state.StateDB) ([32]byte, error) {
	return [32]byte{}, errors.New("unimplemented")
}

func (con ArbOwner) GetLastUpgradeHashGasCost() *big.Int {
	return nil
}

func (con ArbOwner) GetUploadedCodeHash(caller common.Address, st *state.StateDB) ([32]byte, error) {
	return [32]byte{}, errors.New("unimplemented")
}

func (con ArbOwner) GetUploadedCodeHashGasCost() *big.Int {
	return nil
}

func (con ArbOwner) IsAllowedSender(caller common.Address, st *state.StateDB, addr common.Address) (bool, error) {
	return false, errors.New("unimplemented")
}

func (con ArbOwner) IsAllowedSenderGasCost(addr common.Address) *big.Int {
	return nil
}

func (con ArbOwner) IsChainOwner(caller common.Address, st *state.StateDB, addr common.Address) (bool, error) {
	return false, errors.New("unimplemented")
}

func (con ArbOwner) IsChainOwnerGasCost(addr common.Address) *big.Int {
	return nil
}

func (con ArbOwner) IsFairGasPriceSender(caller common.Address, st *state.StateDB, addr common.Address) (bool, error) {
	return false, errors.New("unimplemented")
}

func (con ArbOwner) IsFairGasPriceSenderGasCost(addr common.Address) *big.Int {
	return nil
}

func (con ArbOwner) IsMappingException(
	caller common.Address,
	st *state.StateDB,
	from *big.Int,
	to *big.Int,
) (bool, error) {
	return false, errors.New("unimplemented")
}

func (con ArbOwner) IsMappingExceptionGasCost(from *big.Int, to *big.Int) *big.Int {
	return nil
}

func (con ArbOwner) RemoveAllowedSender(caller common.Address, st *state.StateDB, addr common.Address) error {
	return errors.New("unimplemented")
}

func (con ArbOwner) RemoveAllowedSenderGasCost(addr common.Address) *big.Int {
	return nil
}

func (con ArbOwner) RemoveChainOwner(caller common.Address, st *state.StateDB, addr common.Address) error {
	return errors.New("unimplemented")
}

func (con ArbOwner) RemoveChainOwnerGasCost(addr common.Address) *big.Int {
	return nil
}

func (con ArbOwner) RemoveMappingException(
	caller common.Address,
	st *state.StateDB,
	from *big.Int,
	to *big.Int,
) error {
	return errors.New("unimplemented")
}

func (con ArbOwner) RemoveMappingExceptionGasCost(from *big.Int, to *big.Int) *big.Int {
	return nil
}

func (con ArbOwner) SerializeAllParameters(caller common.Address, st *state.StateDB) ([]byte, error) {
	return nil, errors.New("unimplemented")
}

func (con ArbOwner) SerializeAllParametersGasCost() *big.Int {
	return nil
}

func (con ArbOwner) SetChainParameter(caller common.Address, st *state.StateDB, which [32]byte, value *big.Int) error {
	return errors.New("unimplemented")
}

func (con ArbOwner) SetChainParameterGasCost(which [32]byte, value *big.Int) *big.Int {
	return nil
}

func (con ArbOwner) SetFairGasPriceSender(
	caller common.Address,
	st *state.StateDB,
	addr common.Address,
	isFairGasPriceSender bool,
) error {
	return errors.New("unimplemented")
}

func (con ArbOwner) SetFairGasPriceSenderGasCost(addr common.Address, isFairGasPriceSender bool) *big.Int {
	return nil
}

func (con ArbOwner) SetL1GasPriceEstimate(caller common.Address, st *state.StateDB, priceInGwei *big.Int) error {
	return errors.New("unimplemented")
}

func (con ArbOwner) SetL1GasPriceEstimateGasCost(priceInGwei *big.Int) *big.Int {
	return nil
}

func (con ArbOwner) StartCodeUpload(caller common.Address, st *state.StateDB) error {
	return errors.New("unimplemented")
}

func (con ArbOwner) StartCodeUploadGasCost() *big.Int {
	return nil
}

func (con ArbOwner) StartCodeUploadWithCheck(caller common.Address, st *state.StateDB, oldCodeHash [32]byte) error {
	return errors.New("unimplemented")
}

func (con ArbOwner) StartCodeUploadWithCheckGasCost(oldCodeHash [32]byte) *big.Int {
	return nil
}
