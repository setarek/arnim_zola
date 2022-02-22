package spot

import (
	"github.com/gin-gonic/gin"
	"github.com/setarek/arnim_zola/config"
	"github.com/setarek/arnim_zola/internal/services/binance"
	"github.com/setarek/arnim_zola/pkg/trade_tools"
	"github.com/setarek/arnim_zola/pkg/utils"
	"net/http"
)

type SpotHandler struct {}

type CalculateRequest struct {
	Symbol   string    `json:"symbol"`
}

type CalculateResponse struct {
	Symbol       string              `json:"symbol"`
	Suggestions  []SuggestionPoint   `json:"suggestions"`
}

type ResponseError struct {
	ErrorMessage    string    `json:"error_message"`
}

type SuggestionPoint struct {
	Buy    float64   `json:"buy"`
	Sell   float64   `json:"sell"`
}

func (h *SpotHandler) CalculateOptPoint(ctx *gin.Context)  {
	var request CalculateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil || request.Symbol == "" {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			ErrorMessage: "invalid parameter",
		})
		return
	}

	config := config.GetConfig()
	klines, err := binance.GetKlines(config.GetString("api_key"), config.GetString("api_secret"), request.Symbol)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			ErrorMessage: "error while get klines from binance",
		})
		return
	}

	var suggestionList []SuggestionPoint
	for _, k := range klines {
		var suggest SuggestionPoint
		high := utils.ParseFloat64(k.High)
		low := utils.ParseFloat64(k.Low)
		closingPrice := utils.ParseFloat64(k.Close)
		pp := trade_tools.CalculatePivotPoint(high, low, closingPrice)
		firstSupport := trade_tools.CalculateFirstSupport(pp, high)
		secondSupport := trade_tools.CalculateSecondSupport(pp, high, low)
		suggest.Buy = firstSupport
		if secondSupport < firstSupport {
			suggest.Buy = secondSupport
		}

		firstResistance := trade_tools.CalculateFirstResistance(pp, low)
		secondResistance := trade_tools.CalculateSecondResistance(pp, high, low)
		suggest.Sell = firstResistance
		if secondResistance > firstResistance {
			suggest.Sell = secondResistance
		}
		suggestionList = append(suggestionList, suggest)
	}

	ctx.JSON(http.StatusOK, CalculateResponse{
		Symbol: request.Symbol,
		Suggestions: suggestionList,
	})
	return

}
