package stripe

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"jimmify-server/db"
	"jimmify-server/firebase"
	"log"
	"os"
)

func init() {
	stripe.Key = os.Getenv("JSTRIPEKEY")
}

//PrioritizeQuestion prioritizes ?
func PrioritizeQuestion(token string, qkey int64) error {
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

	err = db.MoveToFront(qkey)
	if err != nil {
		return err
	}
	firebase.Push("Jimmy Payment", "Dolla Dolla Bill Ya'll")
	return err
}
