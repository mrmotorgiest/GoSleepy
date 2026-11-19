package errors

import "fmt"

func ReceiveError(err error, context string) {
	if err != nil {
		panic(fmt.Sprintf("Error in %s: %v", context, err))
	}
}
