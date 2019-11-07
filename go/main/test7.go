package main

import ("fmt"

)
func main(){
a := addr(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
		fmt.Println(s)

	}

}
type iaddr func(int) (int,iaddr)
func addr(base int) iaddr{
	return func(v int)(int,iaddr){
		return base+v,addr(base+v)
	}
}