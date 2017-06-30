package main
import(
  
	//bulit-in package
	"time"

  //extends
  _ "SEP.DataListener/disp"
  . "bakerstreet-club/logs"
)

func main(){

  	//log success info
  	Info(" => system named SEP.DataListener init success .. ")
  	//prevent run out of main func
  	for{time.Sleep(1<<63 - 1)}
}
