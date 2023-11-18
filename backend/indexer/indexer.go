package indexer

import (
	"context"
	"fmt"
	"strconv"
	"subman/chains"
	"subman/contract"
	"subman/converter"
	"subman/database"
	"subman/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.Background()

func Start() {
	ch := make(chan types.Chain, 1)
	for _, chain := range chains.Chains {
		go rabbitHole(chain, ch)
	}
}

func rabbitHole(chain types.Chain, ch chan types.Chain) {
	models, err := index(&chain)
	if err != nil {
		fmt.Println(err)
		ch <- chain
	}
	if len(models) > 0 {
		err = database.BulkWrite(models)
		if err != nil {
			fmt.Println(err)
			ch <- chain
		}
	}
	ch <- chain
}

func index(chain *types.Chain) ([]mongo.WriteModel, error) {
	var models []mongo.WriteModel
	client, err := ethclient.Dial(chain.RPC)
	if err != nil {
		return models, err
	}
	start, err := database.GetLastIndexedBlock(chain.ChainID)
	if err != nil {
		return models, err
	}
	endBlock, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return models, err
	}
	end := endBlock.Number.Uint64() - chain.Confirms
	if start+500 < end {
		end = start + 500
	}
	filter := bind.FilterOpts{
		Start:   start,
		End:     &end,
		Context: nil,
	}
	addr := common.HexToAddress(chain.Address)
	ctr, err := contract.NewContractFilterer(addr, client)
	if err != nil {
		return models, err
	}

	paymentProcessedModels, err := indexPaymentProcessed(ctr, filter)
	if err != nil {
		return models, err
	}
	models = append(models, paymentProcessedModels...)

	paymentCanceledModels, err := indexPaymentCanceled(ctr, filter)
	if err != nil {
		return models, err
	}
	models = append(models, paymentCanceledModels...)

	planActivatedModels, err := indexPlanActivated(ctr, filter)
	if err != nil {
		return models, err
	}
	models = append(models, planActivatedModels...)

	planDeactivatedModels, err := indexPlanDeactivated(ctr, filter)
	if err != nil {
		return models, err
	}
	models = append(models, planDeactivatedModels...)

	return models, nil
}

func indexPaymentProcessed(ctr *contract.ContractFilterer, filter bind.FilterOpts) ([]mongo.WriteModel, error) {
	var models []mongo.WriteModel
	iterator, err := ctr.FilterSubPaymentProcessed(&filter)
	if err != nil {
		return models, err
	}
	payments := []types.SubPayment{}
	for iterator.Next() {
		if iterator.Event.Raw.Removed {
			continue
		}
		payment, err := converter.ConvertToType(iterator.Event.SubPayment)
		if err != nil {
			return models, err
		}
		payments = append(payments, payment)
	}
	if len(payments) > 0 {
		for _, payment := range payments {
			model := database.NewDeadlineModel(payment, payment.EndTime)
			models = append(models, model)
		}
	}
	return models, nil
}

func indexPaymentCanceled(ctr *contract.ContractFilterer, filter bind.FilterOpts) ([]mongo.WriteModel, error) {
	var models []mongo.WriteModel
	iterator, err := ctr.FilterSubPaymentCanceled(&filter)
	if err != nil {
		return models, err
	}
	subscriptions := []types.SubPayment{}
	for iterator.Next() {
		if iterator.Event.Raw.Removed {
			continue
		}
		subscription, err := converter.ConvertToType(iterator.Event.SubPayment)
		if err != nil {
			return models, err
		}
		subscriptions = append(subscriptions, subscription)
	}
	if len(subscriptions) > 0 {
		for _, subscription := range subscriptions {
			model := database.NewCanceledModel(subscription)
			models = append(models, model)
		}
	}
	return models, nil
}

func indexPlanActivated(ctr *contract.ContractFilterer, filter bind.FilterOpts) ([]mongo.WriteModel, error) {
	var models []mongo.WriteModel
	iterator, err := ctr.FilterSubPlanActivated(&filter)
	if err != nil {
		return models, err
	}
	plans := []uint64{}
	for iterator.Next() {
		if iterator.Event.Raw.Removed {
			continue
		}
		planId, err := strconv.ParseUint(iterator.Event.SubPlanId.String(), 10, 64)
		if err != nil {
			return models, err
		}
		plans = append(plans, planId)
	}
	if len(plans) > 0 {
		for _, plan := range plans {
			model := database.NewActivatedModel(plan)
			models = append(models, model)
		}
	}
	return models, nil
}

func indexPlanDeactivated(ctr *contract.ContractFilterer, filter bind.FilterOpts) ([]mongo.WriteModel, error) {
	var models []mongo.WriteModel
	iterator, err := ctr.FilterSubPlanDeactivated(&filter)
	if err != nil {
		return models, err
	}
	plans := []uint64{}
	for iterator.Next() {
		if iterator.Event.Raw.Removed {
			continue
		}
		planId, err := strconv.ParseUint(iterator.Event.SubPlanId.String(), 10, 64)
		if err != nil {
			return models, err
		}
		plans = append(plans, planId)
	}
	if len(plans) > 0 {
		for _, plan := range plans {
			model := database.NewDeactivatedModel(plan)
			models = append(models, model)
		}
	}
	return models, nil
}
