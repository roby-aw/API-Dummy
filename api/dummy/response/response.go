package response

import "api-redeem-point/business/dummy"

type ResponseGetProduct struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  []dummy.StockProduct
}
