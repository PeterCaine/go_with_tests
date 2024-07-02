package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(s, l string) string {
  if len(s) == 0{
    s = "world"
  }
  return greetingPrefix(l) + s +"!" 
}

func greetingPrefix(l string) (prefix string) {
  switch l {
  case "Spanish":
    prefix = spanishHelloPrefix
  case "French":
    prefix = frenchHelloPrefix
  default:
    prefix = englishHelloPrefix
  }
  return 
}

func main(){
  fmt.Println(Hello("Chris", "English"))
}
