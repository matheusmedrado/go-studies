package module01
import "strings"

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//

//This is a more efficient way of concatenating strings
func Reverse(word string) string {

	var rev strings.Builder
	for i := len(word)-1; i>=0; i--{
		rev.WriteByte(word[i])
	}
	return rev.String() 
}

//The "regular" way would be like this:
// func Reverse(word string) string {

// 	var rev string
// 	for i := len(word)-1; i>=0; i--{
// 		rev += string(word[i])
// 	}
// 	return rev 
// }


