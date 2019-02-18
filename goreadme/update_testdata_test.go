// +build testdata

package goreadme_test

import "fmt"

func init() {
	err := updateTestdata()
	if err != nil {
		panic(fmt.Sprintf("error updating test data: %v", err))
	}
}
