package prialgorithm

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)
const GUARD_FILE_PATH   = "/guard_of_sync_opec_data"
const 	filePath = "test_remove"
func TestReverseListNode(t *testing.T) {
	//a := &ListNode{Val: 1}
	//a.Next = &ListNode{Val: 2}
	//a.Next.Next = &ListNode{Val: 3}
	//a.Next.Next.Next = &ListNode{Val: 4}
	//a.Next.Next.Next.Next = &ListNode{Val: 5}
	//
	////a.Next.Next.Next.Next.Next = a.Next
	//
	//b := FindMiddleNode(a)
	//t.Log(os.Args[0])
	///home/dashi/data/gemini/gemini/file
	fmt.Println(splitFileIndex("/home/dashi/data/gemini/gemini_rc/file")+GUARD_FILE_PATH)
}

func TestCreateCornTab(t *testing.T) {

	//var fivesChan =make(<-chan time.Duration*5)
	//select {
	//case true:
	//
	//}
	ticker1 := time.Tick(15 * time.Second)
	ticker2 := time.Tick(10 * time.Second)
	for {
		select {
		case <-ticker1:
			fmt.Println("15s")
			go Remove1()
		case <-ticker2:
			fmt.Println("10s")
			go Remove2()
		}
	}

}

func Remove1() {

	var tmpFile *os.File

	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("111111111%s has already exists, exit. \n", filePath)
		return
	} else if os.IsNotExist(err) {
		var innerErr error
		tmpFile, innerErr = os.Create(filePath)
		if innerErr != nil {
			fmt.Printf("1111111os.Create() failed||err=%v \n", innerErr)
			return
		}
		fmt.Printf("1111111111%s \n","Create")

	} else {
		fmt.Printf("111111syncOpecData() failed||err=%v \n", err)
		return
	}
	if tmpFile != nil {
		defer func() {
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
			go func() {
				<- sigs
				tmpFile.Close()
				err := os.Remove(filePath)
				if err != nil {
					fmt.Errorf("--f os.Remove() failed||err=%v", err)
					return
				}
				os.Exit(0)
			}()
			tmpFile.Close()
			err := os.Remove(filePath)
			if err != nil {
				fmt.Printf("1111111os.Remove() failed||err=%v \n", err)
				return
			}
			fmt.Printf("1111111111111%s \n","delete")

		}()
	}
	time.Sleep(10*time.Second)
	fmt.Printf("1111111111111111111 %s \n", filePath)


}

func Remove2() {

	var tmpFile *os.File

	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf( "222222%s has already exists, exit. \n", filePath)
		return
	} else if os.IsNotExist(err) {
		var innerErr error
		tmpFile, innerErr = os.Create(filePath)
		if innerErr != nil {
			fmt.Printf("222222os.Create() failed||err=%v \n", innerErr)
			return
		}
		fmt.Printf("222222222222%s \n","Create")
	} else {
		fmt.Printf("22222222syncOpecData() failed||err=%v \n", err)
		return
	}
	if tmpFile != nil {
		defer func() {
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
			go func() {
				<- sigs
				tmpFile.Close()
				err := os.Remove(filePath)
				if err != nil {
					fmt.Errorf("--f os.Remove() failed||err=%v", err)
					return
				}
				os.Exit(0)
			}()
			tmpFile.Close()
			err := os.Remove(filePath)
			if err != nil {
				fmt.Printf("2222222os.Remove() failed||err=%v \n", err)
				return
			}
			fmt.Printf("222222222222%s \n","delete")

		}()
	}
	time.Sleep(10*time.Second)
	fmt.Printf("222222222222222222222 %s \n", filePath)

}

func Fibonacci()  {

}