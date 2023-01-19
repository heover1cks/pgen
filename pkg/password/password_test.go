package password

import (
	"sync"
	"testing"
	"time"
)

func TestGeneratePass(t *testing.T) {
	wg := sync.WaitGroup{}

	opts := NewOptions()
	opts.WithPassLength(128)
	opts.WithUpperCase(true)
	opts.WithMaxSpecialCharacters(10)

	wg.Add(1)
	go func() {
		defer wg.Done()
		// GeneratePassword.
		start1 := time.Now()
		println(opts.GeneratePassword())
		println("function1) elapsed: ", time.Since(start1).String())
		println()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// GeneratePasswordV3.
		start3 := time.Now()
		println(opts.GeneratePasswordV3())
		println("function3) elapsed: ", time.Since(start3).String())
		println()
	}()
	wg.Wait()
}
