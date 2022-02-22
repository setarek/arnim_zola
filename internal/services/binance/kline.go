package binance

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance"
)

func GetKlines(apiKey, secretKey, symbol string) ([]*binance.Kline, error) {
	client := binance.NewClient(apiKey, secretKey)
	klines, err := client.NewKlinesService().Symbol(symbol).
		Interval("1d").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return klines, nil
}
