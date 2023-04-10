package main
import "fmt"

func main(){
  //0の初期化を明示的に行う
  var a int
  //宣言と初期化を同時
  var b int = 20
  //型推論
  var c = 30

  //同時に複数の宣言
  var d,e int = 1,2

  //0初期化も可能
  var f,g int

  //型が違う場合は推論でもOK
  var h,i = "hey",3

  //短縮系
  j := 4

  // 短縮系は（一つ以上の新しい変数を宣言する場合）に既存の変数の上書きも可能
  k,a,b := 5,33,55
  
}
 
