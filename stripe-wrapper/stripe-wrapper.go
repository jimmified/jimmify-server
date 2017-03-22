package stripe

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"jimmify-server/db"
	"log"
	"os"
	"strconv"
)

func init() {
	stripe.Key = os.Getenv("JSTRIPEKEY")
}

//PrioritizeQuestion prioritizes ?
func PrioritizeQuestion(token string, qkey string) error {
	// Charge the user's card:
	params := &stripe.ChargeParams{
		Amount:   100,
		Currency: "usd",
	}

	params.SetSource(token)

	charge, err := charge.New(params)
	if err != nil {
		return err
	}
	log.Println(charge)

	i, err := strconv.ParseInt(qkey, 10, 64)
	if err != nil {
		return err
	}

	err = db.MoveToFront(i)
	return err
}
