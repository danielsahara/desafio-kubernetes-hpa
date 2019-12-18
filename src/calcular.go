package main

import(
	"fmt"
	"log"
	"net/http"
	"math"
)

func calcular() float64{
	var x = 0.0001;

	for i := 0; i <= 1000000; i++ {
		x += math.Sqrt(x)
	}
	return x
}

func home(w http.ResponseWriter, r *http.Request){
	calcular();

	fmt.Fprint(w, "Code.education Rocks Go!")
}

func handle(){
	http.HandleFunc("/", home)
}

func main(){
	handle()
	
	log.Fatal(http.ListenAndServe(":80", nil))
}