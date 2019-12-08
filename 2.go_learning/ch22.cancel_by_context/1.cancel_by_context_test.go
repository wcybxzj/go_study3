package ch22_cancel_by_context

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	select {
		case <-ctx.Done():
			return true
		default:
			return false
	}
}

//比较简单的context cancel使用
/*
goroutine i: 2 Cancelled
goroutine i: 0 Cancelled
goroutine i: 1 Cancelled
goroutine i: 3 Cancelled
goroutine i: 4 Cancelled
 */
func TestV1(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Microsecond)
			}
			fmt.Println("goroutine i:", i, "Cancelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}

func worker(i int, ctx context.Context, wg *sync.WaitGroup) {
	for {
		if isCancelled(ctx) {
			break
		}
		time.Sleep(time.Microsecond)
	}
	fmt.Println("goroutine i:", i, "Cancelled")
	wg.Done()
}

func CreateWorker(start int, end int, ctx context.Context, wg *sync.WaitGroup) {
	var wgWorker sync.WaitGroup
	myCtx, _:= context.WithCancel(ctx)

	for i := start; i < end; i++ {
		wgWorker.Add(1)
		go worker(i, myCtx, &wgWorker)
	}
	//cancel()
	wgWorker.Wait()
	wg.Done()
}

/*
输出:
goroutine i: 9 Cancelled
goroutine i: 4 Cancelled
goroutine i: 0 Cancelled
goroutine i: 1 Cancelled
goroutine i: 2 Cancelled
goroutine i: 3 Cancelled
goroutine i: 8 Cancelled
goroutine i: 5 Cancelled
goroutine i: 6 Cancelled
goroutine i: 7 Cancelled
*/
func TestV2(t *testing.T) {
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go CreateWorker(0,5, ctx, &wg)

	wg.Add(1)
	go CreateWorker(5,10, ctx, &wg)

	cancel()
	wg.Wait()
}