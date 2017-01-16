package stripe

import (
	"github.com/stripe/stripe-go"
	"jimmify-server/db"
	"os"
)

func init() {
	stripe.Key = os.Getenv("JSTRIPEKEY")
}

//PrioritizeQuestion prioritizes ?
func PrioritizeQuestion(chargeID string, qkey int64) error {
	//_, err := charge.Get(id, nil)

	//if err != nil {
	//	return err
	//}

	err := db.AddCharge(chargeID)

	if err != nil {
		return err
	}

	err = db.MoveToFront(qkey)
	return err
}
