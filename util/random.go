package util

import (
	"math/rand"
	"strings"
	"time"
)
const alfabet = "abcdefghijklmnopqrestuvwxyz"
func init()  {

	rand.Seed(time.Now().UnixNano())
	
}
func RandomInt(min,max int64) int64 {
	return min + rand.Int63n(max-min + 1)//random int 
	
}

func RandomString(n int)  string{
	var sb strings.Builder
	k := len(alfabet)
	for i := 0; i<n; i++{
		c := alfabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney()  int64{
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EURO", "KSH", "USD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}