package main
import "fmt"
func main() {
	superhero := map[string]map[string]string{

		    "superman" : map[string]string{
		    	"Realname" : "Clark Kent",
		    	"City" : "Metropolis",
	},
	       "Batman" : map[string]string{
		    	"Realname" : "Bruce wayan",
		    	"City" : "las vegas",
		   },
	}
  if temp, hero := superhero["superman"];hero{
  	fmt.Println(temp["Realname"],temp["City"])
  }
}
