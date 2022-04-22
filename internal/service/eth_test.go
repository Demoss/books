package service

import (
	"JMIND/internal/domain"
	"JMIND/internal/query"
	"testing"
)

func TestSuccessGetInfoAboutBlock(t *testing.T) {
	block := query.BlockByNumber{
		Result: query.Result{
			Number: "109789",
			Transactions: []query.Transactions{
				{BlockNumber: "109789",
					Value: "0x455f32e9884c2400"},
			},
		},
	}
	expected := domain.BlockInfo{
		Transactions: 1,
		Count:        4.998770090000001,
	}

	transactions, err := countTransactions(block)
	if err != nil {
		return
	}
	count, err := getCount(block)
	if err != nil {
		return
	}
	res := domain.BlockInfo{
		Transactions: transactions,
		Count:        count,
	}

	if expected != res {
		t.Errorf("expected: %v \n response: %v", expected, res)
	}
}

func TestEmptyBlock(t *testing.T) {
	block := query.BlockByNumber{
		Result: query.Result{
			Number:       "",
			Transactions: []query.Transactions{},
		},
	}
	expected := domain.BlockInfo{
		Transactions: 0,
		Count:        0,
	}

	transactions, err := countTransactions(block)
	if err != nil {
		return
	}
	count, err := getCount(block)
	if err != nil {
		return
	}
	res := domain.BlockInfo{
		Transactions: transactions,
		Count:        count,
	}

	if expected != res {
		t.Errorf("expected: %v \n response: %v", expected, res)
	}
}
