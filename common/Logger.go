package common

import "fmt"

func NewLogger()  {
	
}
 func d(msg string)  {

	fmt.Println(msg)
}
 func v(tag string,msg string)  {
	fmt.Printf("%s=%s",tag,msg)
}
