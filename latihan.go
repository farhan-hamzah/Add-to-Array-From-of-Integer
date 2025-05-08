package main
import "fmt"
const NMAX int = 100
func main(){
	var n, i, k, jum, x int
	var A[NMAX]int
	jum = 0
	fmt.Scan(&n, &k)
	for i = 0; i <n; i++{
		fmt.Scan(&A[i])
	}
	for i = 0; i < n; i++{
		x = A[i]
		jum = (jum*10)+x
	}
	jum = jum+k
	fmt.Print(jum)
}