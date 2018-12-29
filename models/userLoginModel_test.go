package models

import (
	"fmt"
	// "library/go-simplejson" // for json get
	// "system"
	"testing"
)

// func GetTypeMap(value interface{}) map[string]interface{} {

// 	switch inst := value.(type) {
// 	case map[string]interface{}:
// 		return inst
// 	default:
// 		return nil

// 	}

// }

func TestCheckUser(t *testing.T) {

	fmt.Println(GetUserId("gavin"))

}
