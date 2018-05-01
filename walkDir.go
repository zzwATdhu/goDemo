package main

import (
	"flag"
	"path/filepath"
	"sync"
	"time"
	"io/ioutil"
    "fmt"
    "os"
)


var verbose = flag.Bool("v", false, "show verbose progress messages")

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

//cancellation
var cancelFlag = make(chan struct{})
func cancelled() bool {
	select {
		case <-cancelFlag:
			return true
		default:
			return false
	}
}

func main() {

	flag.Parse()
	roots := flag.Args()
	if len(roots) ==0{
    	roots = []string{"."}
	}


	// Cancel traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(cancelFlag)
	}()


	//traverse the file tree.
	fileSizes :=make(chan int64)
	var n sync.WaitGroup
	for _,root:=range roots{
    	n.Add(1)
    	go walkDir(root,&n,fileSizes)
	}

	go func(){
    	n.Wait()
    	close(fileSizes)
	}()


	var tick <-chan time.Time
	if *verbose{
    	tick = time.Tick(500*time.Millisecond)
	}
	var nfiles,nbytes int64

	loop:
    	for{
        	select{
				case <-cancelFlag:
					// Drain fileSizes to allow existing goroutines to finish.
					for range fileSizes {
						// Do nothing.
					}
        		case size,ok:=<-fileSizes:
            		if !ok{
                		break loop//fileSizes was closed
            		}
            		nfiles++
            		nbytes +=size
        		case <-tick:
            		printDiskUsage(nfiles, nbytes)
        	}
    	}
    printDiskUsage(nfiles, nbytes) // final totals
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSize chan<-int64){
	defer n.Done()
	if cancelled() {
		return
	}
    for _,entry:=range dirents(dir){
        if entry.IsDir(){
            n.Add(1)
            subdir:=filepath.Join(dir,entry.Name())
            go walkDir(subdir,n,fileSize)
        }else{
            fileSize<-entry.Size()
        }
    }
}



// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo{
	select {
		case sema <- struct{}{}: // acquire token
		case <-cancelFlag:
			return nil // cancelled
		}

    defer func() { <-sema }() // release token

    entries,err :=ioutil.ReadDir(dir)
    if err!=nil{
        fmt.Fprintf(os.Stderr,"walkDir:%v\n",err)
        return nil
    }
    return entries
}


func printDiskUsage(nfiles, nbytes int64) {
    fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}