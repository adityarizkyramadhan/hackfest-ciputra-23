package payment

import (
	"fmt"
	"os"

	"github.com/adityarizkyramadhan/hackfest-ciputra-23/model"
	"github.com/gofrs/uuid"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/client"
	"github.com/xendit/xendit-go/invoice"
)

var xenCli = client.New(os.Getenv("XENDIT_KEY"))

type PaymentRequest struct {
	Amount      float64 `json:"amount" binding:"required"`
	Description string  `json:"description" binding:"required"`
}

// Create
func CreatePayment(arg *PaymentRequest, user *model.User) (*xendit.Invoice, error) {
	resp, err := xenCli.Invoice.Create(&invoice.CreateParams{
		ForUserID:   "",
		ExternalID:  uuid.Must(uuid.NewV6()).String(),
		Amount:      arg.Amount,
		Description: arg.Description,
		Customer: xendit.InvoiceCustomer{
			GivenNames: user.Name,
		},
		InvoiceDuration:              100,
		SuccessRedirectURL:           fmt.Sprintf("%s/api/v1/payment/success", os.Getenv("BASE_URL")),
		FailureRedirectURL:           fmt.Sprintf("%s/api/v1/payment/failure", os.Getenv("BASE_URL")),
		PaymentMethods:               []string{"BNI", "BSI", "BRI", "BSS", "MANDIRI", "PERMATA", "BJB", "OVO", "DANA", "QRIS"},
		Currency:                     "IDR",
		MidLabel:                     "INV",
		ReminderTimeUnit:             "days",
		ReminderTime:                 7,
		Locale:                       "en",
		ShouldAuthenticateCreditCard: true,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
