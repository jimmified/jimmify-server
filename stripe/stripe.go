package stripe

import (
	"fmt"
	"github.com/stripe/stripe-go"
	"os"
)

func init() {
	stripe.Key = os.Getenv("JSTRIPEKEY")
}

func prioritizeQuestion(chargeID string, qkey int64) {
	c, err := charge.Get(id, nil)

	if err == nil {
		// Charge Exists
		// Prioritize
		fmt.Println(c)
	} else {
		fmt.Println(err)
	}

}
