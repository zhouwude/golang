package main

import (
	"fmt"
)

// Send the sequence 2, 3, 4, ... to returned channel
func generate1() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; i < 20; i++ {
			ch <- i
		}
	}()
	return ch
}

// Filter out input values divisible by 'prime', send rest to returned channel
func filter1(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate1()
		for {
			prime := <-ch
			ch = filter1(ch, prime)
			out <- prime
		}
	}()
	return out
}

func mainsieve2() {
	primes := sieve()
	for {
		fmt.Println(<-primes)
	}
}
