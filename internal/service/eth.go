package service

import (
	"JMIND"
	"JMIND/internal/domain"
	"JMIND/internal/query"
	"JMIND/internal/repository"
	"context"
	"math"
	"strconv"
)

type EthService struct {
	repo repository.Eth
}

func NewEthService(repo repository.Eth) *EthService {
	return &EthService{repo: repo}
}

func (s *EthService) GetInfoAboutBlock(ctx context.Context, number query.BlockByNumber) (domain.BlockInfo, error) {
	blockNumber := number.Result.Number

	block, err := s.repo.GetBlock(ctx, blockNumber)
	if err != nil {
		var info domain.BlockInfo
		info.BlockNumber = blockNumber

		transactions, err := countTransactions(number)
		if err != nil {
			return domain.BlockInfo{}, err
		}
		info.Transactions = transactions

		count, err := getCount(number)
		if err != nil {
			return domain.BlockInfo{}, err
		}
		info.Count = count

		err = s.repo.SaveBlock(ctx, info)
		if err != nil {
			return info, err
		}
		block = info
	}

	return block, nil
}
func countTransactions(number query.BlockByNumber) (int, error) {
	sum := 0
	for i := 0; i < len(number.Result.Transactions); i++ {
		sum++
	}
	return sum, nil
}

func getCount(number query.BlockByNumber) (float64, error) {
	amount := 0.0
	for i := 0; i < len(number.Result.Transactions); i++ {
		float, _ := strconv.ParseInt(JMIND.HexaNumberToInteger(number.Result.Transactions[i].Value), 16, 64)
		newF := float64(float)
		newF = newF * math.Pow(10, -18)
		amount += newF
	}
	return amount, nil
}
