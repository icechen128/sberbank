package mobile

import (
	"context"
	"encoding/json"
	"fmt"
	acquiring "github.com/icechen128/sberbank"
	"github.com/icechen128/sberbank/endpoints"
	"github.com/icechen128/sberbank/schema"
	"net/http"
)

type Client struct {
	API acquiring.API
}

// ApplePaymentRequest is used for building PayWithApplePay request
type ApplePaymentRequest struct {
	OrderNumber          string `json:"orderNumber"`
	Merchant             string `json:"merchant"`
	PaymentToken         string `json:"paymentToken"`
	Description          string `json:"description,omitempty"`
	PreAuth              bool   `json:"preAuth,omitempty"`
	AdditionalParameters map[string]string
}

// PayWithApplePay request
// see https://securepayments.sberbank.ru/wiki/doku.php/integration:api:rest:requests:payment_applepay
func PayWithApplePay(ctx context.Context, applePaymentRequest ApplePaymentRequest) (*schema.ApplePaymentResponse, *http.Response, error) {
	return getClient().PayWithApplePay(ctx, applePaymentRequest)
}

// PayWithApplePay request
// see https://securepayments.sberbank.ru/wiki/doku.php/integration:api:rest:requests:payment_applepay
func (c Client) PayWithApplePay(ctx context.Context, applePaymentRequest ApplePaymentRequest) (*schema.ApplePaymentResponse, *http.Response, error) {
	path := endpoints.ApplePay

	if err := validateApplePaymentRequest(applePaymentRequest); err != nil {
		return nil, nil, err
	}

	var response schema.ApplePaymentResponse

	req, err := c.API.NewRequest(ctx, "GET", path, applePaymentRequest)

	if err != nil {
		return nil, nil, err
	}
	result, err := c.API.Do(req, &response)
	if err != nil {
		return nil, result, err
	}
	_ = json.NewDecoder(result.Body).Decode(&response)

	return &response, result, err
}

func validateApplePaymentRequest(request ApplePaymentRequest) error {
	if request.OrderNumber == "" || request.Merchant == "" || request.PaymentToken == "" {
		return fmt.Errorf("orderNumber, merchant and PaymentToken are required")
	}
	return nil
}

// GooglePaymentRequest request
// "OrderNumber" _required_ used to pass orderId to api
// "Merchant" _required_ merchant is API
// "PaymentToken" _required_ token encoded in base64
// "Language" in response
// "Description" check API Docs
// "ClientId" client ID
// "IP" clients ip to store
// "PreAuth" preAuth flag
// "Amount" sum in penny
// "AdditionalParameters" used in some endpoints of API
type GooglePaymentRequest struct {
	OrderNumber          string            `json:"orderNumber"`
	Merchant             string            `json:"merchant"`
	PaymentToken         string            `json:"paymentToken"`
	Language             string            `json:"language"`
	Description          string            `json:"description,omitempty"`
	ClientId             string            `json:"clientId,omitempty"`
	IP                   string            `json:"ip,omitempty"`
	Amount               int               `json:"amount,omitempty"`
	CurrencyCode         int               `json:"currencyCode,omitempty"`
	Email                string            `json:"email,omitempty"`
	Phone                string            `json:"phone,omitempty"`
	ReturnUrl            string            `json:"returnUrl"`
	FailUrl              string            `json:"failUrl,omitempty"`
	PreAuth              bool              `json:"preAuth,omitempty"`
	AdditionalParameters map[string]string `json:"additionalParameters"`
}

// PayWithGooglePay request
// see https://securepayments.sberbank.ru/wiki/doku.php/integration:api:rest:requests:payment_googlepay
func PayWithGooglePay(ctx context.Context, googlePaymentRequest GooglePaymentRequest) (*schema.GooglePaymentResponse, *http.Response, error) {
	return getClient().PayWithGooglePay(ctx, googlePaymentRequest)
}

// PayWithGooglePay request
// see https://securepayments.sberbank.ru/wiki/doku.php/integration:api:rest:requests:payment_googlepay
func (c Client) PayWithGooglePay(ctx context.Context, googlePaymentRequest GooglePaymentRequest) (*schema.GooglePaymentResponse, *http.Response, error) {
	path := endpoints.GooglePay

	if err := validateGooglePayRequest(googlePaymentRequest); err != nil {
		return nil, nil, err
	}

	var response schema.GooglePaymentResponse

	req, err := c.API.NewRequest(ctx, "GET", path, googlePaymentRequest)

	if err != nil {
		return nil, nil, err
	}
	result, err := c.API.Do(req, &response)
	if err != nil {
		return nil, result, err
	}
	_ = json.NewDecoder(result.Body).Decode(&response)

	return &response, result, err
}

func validateGooglePayRequest(request GooglePaymentRequest) error {
	if request.OrderNumber == "" || request.Merchant == "" || request.PaymentToken == "" {
		return fmt.Errorf("orderNumber, merchant and PaymentToken are required")
	}
	return nil
}

// SamsungPaymentRequest is used to send PayWithSamsungPay request
// "OrderNumber" _required_ used to pass orderId to api
// "Merchant" _required_ merchant is API
// "PaymentToken" _required_ token encoded in base64
// "Language" in response
// "Description" check API Docs
// "ClientId" client ID
// "IP" clients ip to store
// "PreAuth" preauth flag
// "AdditionalParameters" used in some endpoints of API
type SamsungPaymentRequest struct {
	OrderNumber          string            `json:"orderNumber"`
	Merchant             string            `json:"merchant"`
	PaymentToken         string            `json:"paymentToken"`
	Language             string            `json:"language"`
	Description          string            `json:"description,omitempty"`
	ClientId             string            `json:"clientId,omitempty"`
	IP                   string            `json:"ip,omitempty"`
	PreAuth              bool              `json:"preAuth,omitempty"`
	AdditionalParameters map[string]string `json:"additionalParameters"`
}

// PayWithSamsungPay is used to send PayWithSamsungPay request
func PayWithSamsungPay(ctx context.Context, samsungPaymentRequest SamsungPaymentRequest) (*schema.SamsungPaymentResponse, *http.Response, error) {
	return getClient().PayWithSamsungPay(ctx, samsungPaymentRequest)
}

// PayWithSamsungPay is used to send PayWithSamsungPay request
func (c Client) PayWithSamsungPay(ctx context.Context, samsungPaymentRequest SamsungPaymentRequest) (*schema.SamsungPaymentResponse, *http.Response, error) {
	path := endpoints.SamsungPay

	if err := validateSamsungPaymentRequest(samsungPaymentRequest); err != nil {
		return nil, nil, err
	}

	var response schema.SamsungPaymentResponse
	req, err := c.API.NewRequest(ctx, "GET", path, samsungPaymentRequest)

	if err != nil {
		return nil, nil, err
	}
	result, err := c.API.Do(req, &response)
	if err != nil {
		return nil, result, err
	}
	_ = json.NewDecoder(result.Body).Decode(&response)

	return &response, result, err
}

func validateSamsungPaymentRequest(request SamsungPaymentRequest) error {
	if request.OrderNumber == "" || request.Merchant == "" || request.PaymentToken == "" {
		return fmt.Errorf("orderNumber, merchant and PaymentToken are required")
	}
	return nil
}

func getClient() Client {
	return Client{acquiring.GetAPI()}
}
