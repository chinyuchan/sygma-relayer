// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package feeHandler

import (
	"strings"

	"github.com/ChainSafe/chainbridge-core/chains/evm/calls"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/contracts"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/transactor"
	"github.com/ChainSafe/chainbridge-core/types"
	"github.com/ChainSafe/sygma-relayer/chains/evm/calls/consts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type FeeRouter struct {
	contracts.Contract
}

func NewFeeRouter(
	client calls.ContractCallerDispatcher,
	feeRouterAddress common.Address,
	transactor transactor.Transactor,
) *FeeRouter {
	a, _ := abi.JSON(strings.NewReader(consts.FeeRouterABI))
	b := common.FromHex(consts.FeeRouterBin)
	return &FeeRouter{contracts.NewContract(feeRouterAddress, a, b, client, transactor)}
}

// AdminSetResourceHandler sets handler for provided domainID and resourceID. https://github.com/ChainSafe/sygma-relayer-solidity/blob/master/contracts/handlers/FeeHandlerRouter.sol#L54
func (c *FeeRouter) AdminSetResourceHandler(destDomainID uint8, resourceID types.ResourceID, feeHandlerAddress common.Address, opts transactor.TransactOptions) (*common.Hash, error) {
	return c.ExecuteTransaction(
		"adminSetResourceHandler",
		opts,
		destDomainID,
		resourceID,
		feeHandlerAddress,
	)
}

func (f *FeeRouter) CollectFee(
	sender common.Address,
	fromDomainID uint8,
	destDomainID uint8,
	resourceID types.ResourceID,
	data []byte,
	feeData []byte,
	opts transactor.TransactOptions,
) (*common.Hash, error) {
	return f.ExecuteTransaction(
		"collectFee",
		opts,
		sender,
		fromDomainID,
		destDomainID,
		resourceID,
		data,
		feeData,
	)
}

func (f *FeeRouter) CalculateFee(
	sender common.Address,
	fromDomainID uint8,
	destDomainID uint8,
	resourceID types.ResourceID,
	data []byte,
	feeData []byte,
	opts transactor.TransactOptions,
) (*common.Hash, error) {
	return f.ExecuteTransaction(
		"calculateFee",
		opts,
		sender,
		fromDomainID,
		destDomainID,
		resourceID,
		data,
		feeData,
	)
}
