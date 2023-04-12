package main

import "fmt"

func main(){

  m := map[string]int{}

  m["test"] = 5
  fmt.Println(m["test"])
  //定義されていない場合は0が返ってくる 
  v, ok := m["a"]
  fmt.Println(v,ok)
  // v = 0 ok = false
  //okの部分は存在すればtrue,存在しなければfalse

  m["b"] = 0
  v, ok = m["b"]
  //v =0 ok = true
  fmt.Println(v,ok)
}
