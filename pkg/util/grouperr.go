package util

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"runtime"
	"time"
)

func main() {
	if err := Call(); err != nil {
		return
	}
	Save()
}

func Save() {
	fmt.Println("Save ing")
}

func Call2(ctx context.Context) error {
	ctx2, _ := context.WithCancel(ctx)
	group, errCtx := errgroup.WithContext(ctx2)

	// 参数处理
	fs := []func(){
		func() { panic("Call2 1") },
		func() { fmt.Println("Call2 123") },
		func() { fmt.Println("Call2 123") },
	}

	for i, v := range fs {
		f := v
		group.Go(func() (err error) {
			defer func() { // 捕获异常
				if recErr := recover(); recErr != nil {
					err = errors.New(GetCurFuncStack())
				}
			}()

			time.Sleep(time.Duration(i) * 5 * time.Second)
			err = CheckGoroutineErr(errCtx) // 绝大部分情况下，不会触发这个错误(协程太少，速度太快)
			if err != nil {
				return
			}
			f()
			return
		})

	}

	// 捕获err
	err := group.Wait()
	if err != nil {
		errStr := fmt.Sprintf("预授信参数处理错误，err:%v", err.Error())
		fmt.Println(errStr)
		return errors.New(errStr)
	}

	return nil
}

func Call() error {
	ctx, cannel := context.WithCancel(context.Background())
	group, errCtx := errgroup.WithContext(ctx)

	// 参数处理
	fs := []func(){
		//func() { panic("1")},

		func() { Call2(errCtx) },
		func() { fmt.Println("123") },
		func() { fmt.Println("123") },
	}

	for i, v := range fs {
		f := v
		group.Go(func() (err error) {
			defer func() { // 捕获异常
				if recErr := recover(); recErr != nil {
					err = errors.New(GetCurFuncStack())
				}
			}()
			if i == 0 {
				cannel()
			}

			time.Sleep(time.Duration(i) * 5 * time.Second)
			err = CheckGoroutineErr(errCtx) // 绝大部分情况下，不会触发这个错误(协程太少，速度太快)
			if err != nil {
				return
			}
			f()
			return
		})

	}

	// 捕获err
	err := group.Wait()
	if err != nil {
		errStr := fmt.Sprintf("预授信参数处理错误，err:%v", err.Error())
		fmt.Println(errStr)
		return errors.New(errStr)
	}

	return nil
}

//校验是否有协程已发生错误
func CheckGoroutineErr(errContext context.Context) error {

	select {
	case <-errContext.Done():
		fmt.Println("协程发生错误")
		return errContext.Err()
	default:
		return nil
	}
}

func GetCurFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func GetCurFuncStack() string {
	const size = 64 << 10
	buf := make([]byte, size)
	buf = buf[:runtime.Stack(buf, false)]
	return string(buf)
}
