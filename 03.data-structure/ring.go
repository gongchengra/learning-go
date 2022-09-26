package main

// from https://pkg.go.dev/container/ring@go1.19#example-Ring.Do
import (
	"container/ring"
	"fmt"
)

func main() {
	{
		// Create a new ring of size 5
		r := ring.New(5)
		// Get the length of the ring
		n := r.Len()
		// Initialize the ring with some integer values
		for i := 0; i < n; i++ {
			r.Value = i
			r = r.Next()
		}
		// Iterate through the ring and print its contents
		r.Do(func(p any) {
			fmt.Println(p.(int))
		})
	}
	{
		// Create a new ring of size 4
		r := ring.New(4)
		// Print out its length
		fmt.Println(r.Len())
	}
	{
		// Create two rings, r and s, of size 2
		r, s := ring.New(2), ring.New(2)
		// Get the length of the ring
		lr, ls := r.Len(), s.Len()
		// Initialize r with 0s
		for i := 0; i < lr; i++ {
			r.Value = 0
			r = r.Next()
		}
		// Initialize s with 1s
		for j := 0; j < ls; j++ {
			s.Value = 1
			s = s.Next()
		}
		// Link ring r and ring s
		rs := r.Link(s)
		// Iterate through the combined ring and print its contents
		rs.Do(func(p any) {
			fmt.Println(p.(int))
		})
	}
	{
		// Create a new ring of size 5
		r := ring.New(5)
		// Get the length of the ring
		n := r.Len()
		// Initialize the ring with some integer values
		for i := 0; i < n; i++ {
			r.Value = i
			r = r.Next()
		}
		// Move the pointer forward by three steps
		r = r.Move(3)
		// Iterate through the ring and print its contents
		r.Do(func(p any) {
			fmt.Println(p.(int))
		})
	}
	{
		// Create a new ring of size 5
		r := ring.New(5)
		// Get the length of the ring
		n := r.Len()
		// Initialize the ring with some integer values
		for i := 0; i < n; i++ {
			r.Value = i
			r = r.Next()
		}
		// Iterate through the ring and print its contents
		for j := 0; j < n; j++ {
			fmt.Println(r.Value)
			r = r.Next()
		}
	}
	{
		// Create a new ring of size 5
		r := ring.New(5)
		// Get the length of the ring
		n := r.Len()
		// Initialize the ring with some integer values
		for i := 0; i < n; i++ {
			r.Value = i
			r = r.Next()
		}
		// Iterate through the ring backwards and print its contents
		for j := 0; j < n; j++ {
			r = r.Prev()
			fmt.Println(r.Value)
		}
	}
	{
		// Create a new ring of size 6
		r := ring.New(6)
		// Get the length of the ring
		n := r.Len()
		// Initialize the ring with some integer values
		for i := 0; i < n; i++ {
			r.Value = i
			r = r.Next()
		}
		// Unlink three elements from r, starting from r.Next()
		r.Unlink(3)
		// Iterate through the remaining ring and print its contents
		r.Do(func(p any) {
			fmt.Println(p.(int))
		})
	}
}
