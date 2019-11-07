package main
import(
	"fmt"
	"github.com/garyburd/redigo/redis"
)
func main(){
c,err := redis.Dial("tcp","127.0.0.1:6379")
if err != nil{
	fmt.Println(err)
	return
}
defer c.Close()
v,err := c.Do("SET","name","red")
	if err != nil{
		fmt.Println(err)
		return
	}
	v, err = redis.String(c.Do("GET", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
}