package main
import "fmt"

func main(){
  
  //配列（あんまりつかわない）
  var x [3]int
  x[2] = 5

  fmt.Println(x[2])
  
  //初期化も可能[...]のように書く
  var y = [...]int{1,2,3,4,5}
  fmt.Println(y[1])
  //スライス
  var a = [] int {1,2,3}
  fmt.Println(a[0])
  //宣言のみ
  var b []int
  b = append(b,43)
  b = append(b,90)
  fmt.Println(b[1])
  }
