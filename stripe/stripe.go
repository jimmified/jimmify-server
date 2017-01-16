package stripe

import (
	"fmt"
	"github.com/stripe/stripe-go"
	"jimmify-server/db"
	"os"
)

func init() {
	stripe.Key = os.Getenv("JSTRIPEKEY")
}

func prioritizeQuestion(chargeID string, qkey int64) {
	//c, err := charge.Get(id, nil)
	err = nil
	if err == nil {
		// Charge Exists
		// Prioritize
		db.MoveToFront(qkey)
	} else {
		fmt.Println(err)
	}

}
