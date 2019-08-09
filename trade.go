package alipay

import (
	"net/url"
)

// TradePagePay https://docs.open.alipay.com/api_1/alipay.trade.page.pay
func (this *Client) TradePagePay(param TradePagePay) (result *url.URL, err error) {
	p, err := this.URLValues(param)
	if err != nil {
		return nil, err
	}

	result, err = url.Parse(this.apiDomain + "?" + p.Encode())
	if err != nil {
		return nil, err
	}
	return result, err
}

// TradeAppPay https://docs.open.alipay.com/api_1/alipay.trade.app.pay
func (this *Client) TradeAppPay(param TradeAppPay) (result string, err error) {
	p, err := this.URLValues(param)
	if err != nil {
		return "", err
	}
	return p.Encode(), err
}

// TradeFastPayRefundQuery https://docs.open.alipay.com/api_1/alipay.trade.fastpay.refund.query
func (this *Client) TradeFastPayRefundQuery(param TradeFastPayRefundQuery) (result *TradeFastPayRefundQueryRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// TradeOrderSettle https://docs.open.alipay.com/api_1/alipay.trade.order.settle
func (this *Client) TradeOrderSettle(param TradeOrderSettle) (result *TradeOrderSettleRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// TradeClose https://docs.open.alipay.com/api_1/alipay.trade.close/
func (this *Client) TradeClose(param TradeClose) (result *TradeCloseRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// TradeCancel https://docs.open.alipay.com/api_1/alipay.trade.cancel/
func (this *Client) TradeCancel(param TradeCancel) (result *TradeCancelRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// TradeRefund https://docs.open.alipay.com/api_1/alipay.trade.refund/
func (this *Client) TradeRefund(param TradeRefund) (result *TradeRefundRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// TradePreCreate https://docs.open.alipay.com/api_1/alipay.trade.precreate/
func (this *Client) TradePreCreate(param TradePreCreate) (result *TradePreCreateRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// TradeQuery https://docs.open.alipay.com/api_1/alipay.trade.query/
func (this *Client) TradeQuery(param TradeQuery) (result *TradeQueryRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// TradeCreate https://docs.open.alipay.com/api_1/alipay.trade.create/
func (this *Client) TradeCreate(param TradeCreate) (result *TradeCreateRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// Trade https://docs.open.alipay.com/api_1/alipay.trade.pay/
func (this *Client) TradePay(param TradePay) (result *TradePayRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// TradeOrderInfoSync https://docs.open.alipay.com/api_1/alipay.trade.orderinfo.sync/
func (this *Client) TradeOrderInfoSync(param TradeOrderInfoSync) (result *TradeOrderInfoSyncRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}
