package tools

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomCode() string {
	code := fmt.Sprintf("%04v",rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	return code
}
