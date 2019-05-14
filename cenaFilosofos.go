package main

import (
     "fmt"
     "time"
)

func comida(i int, sinc chan int) {
     time.Sleep(10 * time.Duration((i%5)+i) * time.Millisecond)
     sinc <- i
}

func main() {
     var sincronizador1 chan int = make(chan int)
     var sincronizador2 chan int = make(chan int)
     var sincronizador3 chan int = make(chan int)
     var sincronizador4 chan int = make(chan int)
     var sincronizador5 chan int = make(chan int)
     go comida(1, sincronizador1)
     go comida(2, sincronizador2)
     go comida(3, sincronizador3)
     go comida(4, sincronizador4)
     go comida(5, sincronizador5)
     for i := 0; i < 5; i += 1 {
          select {
          case i1 := <-sincronizador1:
               fmt.Println("primer filosofo.",
                    "Tu número es el", i1)
          case i2 := <-sincronizador2:
               fmt.Println("segundo filosofo.",
                    "Tu número es el", i2)
          case i3 := <-sincronizador3:
               fmt.Println("tercer filosofo.",
                    "Tu número es el", i3)
          case i4 := <-sincronizador4:
               fmt.Println("cuarto filosofo.",
                    "Tu número es el", i4)
          case i5 := <-sincronizador5:
               fmt.Println("quinto filosofo.",
                    "Tu número es el", i5)
          }
          fmt.Println("filosofo atendido")
     }
     fmt.Println("Fin")
}