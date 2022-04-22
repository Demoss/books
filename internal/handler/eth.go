package handler

import (
	"JMIND/internal/query"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
)

func (h *Handler) GetInfoByBlock(c *gin.Context) {
	block, err := strconv.Atoi(c.Param("block"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	hexInt := decimalToHex(block)

	logrus.Infof("hex number for searcing block = %v", hexInt)

	url := fmt.Sprintf("https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=%v&boolean=true&apikey=%v", hexInt, os.Getenv("apiKey"))

	res, err := http.Get(url)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var transactionsBlock query.BlockByNumber

	err = json.NewDecoder(res.Body).Decode(&transactionsBlock)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	a, err := h.services.Eth.GetInfoAboutBlock(c, transactionsBlock)
	if err != nil {
		logrus.Errorf("data didn`t save due to error: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Infof("data successfully add")

	c.JSON(http.StatusOK, map[string]interface{}{
		"transactions": a.Transactions,
		"amount":       a.Count,
	})
}

func decimalToHex(block int) string {
	hexInt := strconv.FormatInt(int64(block), 16)
	hexInt = "0x" + hexInt
	return hexInt
}
