package main

import (
	"flag"
	"path/filepath"
	"net"
	"sync"
	"log"
	"time"
	"net/http"
    "image"
	"image/gif"
	"image/color"
	"io/ioutil"
    "bufio"
    "math"
    "math/rand"
	// "strings"
    "fmt"
    "os"
    "io"
)


var palette = []color.Color{color.White,color.Black}
const(
    whiteIndex = 0//first color in palette
    blackIndex = 1// next color in palette
)


var mu sync.Mutex
var count int

type Integer  int
//oop 面向对象
func (a Integer) Less(b Integer) bool{
    return a<b
}
func (a *Integer) Add(b Integer) {
    *a += b
}

//opp 面向过程
func Integer_Less(a Integer,b Integer) bool{
    return a<b
}


type Rect struct {
    x,y float64
    width,height float64
}
func (r *Rect)Area() float64{
    return r.width*r.height
}
func NewRect(x,y,width,height float64) *Rect{
    return &Rect{x,y,width,height}
}


func main() {

    // //demo2
    // var s,sep string
    // for i:=1;i<len(os.Args);i++{
    //     s+=sep+os.Args[i]
    //     sep=" "
    // }
    // fmt.Println(s)

    // //demo3
    // s=""
    // sep=" "
    // for _,arg := range os.Args[1:]{
    //     s+=sep+arg
    //     sep = " "
    // }
    // fmt.Println(s)

    // //demo4
    // fmt.Println(strings.Join(os.Args[1:]," "))
    // // fmt.Println(os.Args[1:])

    // //demo5
    // counts:=make(map[string]int)
    // files := os.Args[1:]
    // if len(files)==0{
    //     countLines(os.Stdin,counts)
    // }else{
    //     for _,arg:=range files{
    //         f,err:=os.Open(arg)
    //         if err!=nil{
    //             fmt.Fprintf(os.Stderr,"dup:%v\n",err)
    //             continue
    //         }
    //         countLines(f,counts)
    //         f.Close()
    //     }
    // }
   
    // for line,n:=range counts{
    //     if n>1{
    //         fmt.Printf("%d\t%s\n",n,line)
    //     }
    // }


    // //demo6
    // clearMap(counts)
    // fmt.Println("==============zhang zhi wei==============")
    // for _,filename :=range os.Args[1:]{
    //     data,err := ioutil.ReadFile(filename)
    //     if err!=nil{
    //         fmt.Fprintf(os.Stderr,"dup:%v\n",err)
    //         continue
    //     }
    //     for _,line:=range strings.Split(string(data),"\n"){
    //         counts[line]++
    //     }
    // }
    // for line,n:=range counts{
    //     if n>1{
    //         fmt.Printf("%d\t%s\n",n,line)
    //     }
    // }

// //demo7
// lissajous(os.Stdout)


// //demo8
// for _,url:=range os.Args[1:]{
//     resp,err:=http.Get(url)
//     if err!=nil{
//         fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
//         os.Exit(1)
//     }
//     b,err:=ioutil.ReadAll(resp.Body)
//     resp.Body.Close()
//     if err!=nil{
//         fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
//         os.Exit(1)        
//     }
//     fmt.Printf("%s",b)
// }

// //demo9 concurrency
// start:=time.Now()
// ch:=make(chan string)
// for _,url:=range os.Args[1:]{
//     go fetch(url,ch) //start a go routine
// }
// for range os.Args[1:]{
//     fmt.Println(<-ch) //retrieve from channel ch
// }

// fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds)


// //demo10
// http.HandleFunc("/",handler)
// http.HandleFunc("/count",counter)
// http.HandleFunc("/gif", func(w http.ResponseWriter, r *http.Request) {
//     lissajous(w)
//     })
// log.Fatal(http.ListenAndServe("localhost:8000",nil))



// //demo11
// var a Integer =1
// if a.Less(2){
//     fmt.Println("oop func: ",a,"Less 2")
// }
// a.Add(2)
// fmt.Println("a =", a)

// if Integer_Less(a,10){
//     fmt.Println("opp func: ",a,"Less 2")
// }


// //demo12
// go spinner(100*time.Millisecond)
// const n = 45
// fibN:=fib(n)
// fmt.Printf("\rFibonacci(%d) = %d \n",n,fibN)


// //demo13
// //demo13 client
// go tcpClient()
// //demo13 server
// listener,err:=net.Listen("tcp","localhost:8000")
// if err!=nil{
//     log.Fatal(err)
// }
// for{
//     conn,err:=listener.Accept()
//     if err !=nil{
//         log.Fatal(err)
//         continue
//     }
//     go handleConn(conn)
// }



// //demo14
// naturals := make(chan int)
// squares := make(chan int)

// //counter
// go func() {
//     for x := 0; x < 100; x++ {
//     naturals <- x
//     }
//     close(naturals)
//     }()
    
// go func() {
//     for{
//         x,ok:= <-naturals
//         if !ok{
//             break
//         }
//         squares<- x*x
//     }
//     close(squares)
// }()
// // for{
// //     x,ok:= <-squares
// //     if !ok{
// //         break
// //     }
// //     fmt.Println(x)
// // }
// // Printer (in main goroutine)
// for x := range squares {
//     fmt.Println(x)
//     }
    

// //demo15
// naturals :=make(chan int)
// squares := make(chan int)
// go counter2(naturals)
// go squarer(squares,naturals)
// printer(squares)


// //demo16
// strResult := mirroredQuery()
// fmt.Println("Message Size：",len(strResult),";\nMessage Content:",strResult)


// //demo17

// abort := make(chan struct{})
// go func(){
//     os.Stdin.Read(make([]byte,1))
//     abort<-struct{}{}
// }()

// fmt.Println("Commencing countdown. Press return to abort.")
// ticker := time.NewTicker(1 * time.Second)

// for countdown:=10;countdown>0;countdown--{
// fmt.Println(countdown)
// select {
// case <-ticker.C: // receive from the ticker's channel
//     //do nothing.
// case <-abort:
//     fmt.Println("Launch aborted!")
//     ticker.Stop()
//     return
// }
// }
// //to do sth 
// ticker.Stop()
// fmt.Println("Launch Successed!")




}

//return the quickest response via channel
func mirroredQuery()string{
    responses:=make(chan string,3)
    go fetch("https://www.baidu.com/",responses)
    go fetch("https://www.csdn.net/",responses) 
    go fetch("http://www.runoob.com/go/go-for-loop.html",responses)
    return <-responses //return the quickest response
}


func counter2(out chan<-int){
    for x:=0;x<100;x++{
        out<-x
    }
    close(out)
}

func squarer(out chan<-int,in <-chan int){
    for v:=range in{
        out<-v*v
    }
    close(out)
}

func printer(in <-chan int){
    for v:=range in{
        fmt.Println(v)
    }
}


func tcpClient(){

    time.Sleep(1*time.Second)

    conn,err := net.Dial("tcp","localhost:8000")
    if err!=nil{
        log.Fatal(err)
    }
    defer conn.Close()
    mustCopy(os.Stdout,conn)
}

func mustCopy(dst io.Writer,src io.Reader){
    if _,err := io.Copy(dst,src);err!=nil{
        log.Fatal(err)
    }
}


func handleConn(c net.Conn){
    defer c.Close()
    for{
        _,err := io.WriteString(c,time.Now().Format("15:04:05\n"))
        if err!=nil{
            return
        }
        time.Sleep(1*time.Second)
    }
}


func spinner(delay time.Duration){
    for{
        for _, r := range `-\|/` {
            fmt.Printf("\r%c",r)
            time.Sleep(delay)
        }
    }
}

func fib(x int) int{
    if x<2{
        return x
    }
    return fib(x-1)+fib(x-2)
}

func countLines(f *os.File,counts map[string]int){
    input:=bufio.NewScanner(f)
    for input.Scan(){
        counts[input.Text()]++
    }
}

func clearMap(dataMap map[string]int){
    for k:=range dataMap{
        delete(dataMap,k)
    }
}


func lissajous(out io.Writer){
const (
    cycle = 5
    res =0.001
    size = 100
    nframes = 64
    delay = 8
)

freq:=rand.Float64()*3.0
anim:=gif.GIF{LoopCount:nframes}
phase:=0.0
for i:=0;i<nframes;i++{
    rect:=image.Rect(0,0,2*size+1,2*size+1)
    img:=image.NewPaletted(rect,palette)
    for t:=0.0;t<cycle*2*math.Pi;t+=res{
        x:=math.Sin(t)
        y:=math.Sin(t*freq+phase)
        img.SetColorIndex(size+int(x*size+0.5),size+int(y*size+0.5),blackIndex)
    }
    phase+=0.1
    anim.Delay = append(anim.Delay,delay)
    anim.Image = append(anim.Image,img)
}
gif.EncodeAll(out,&anim)
}

func fetch(url string,out chan<-string){
    start:=time.Now()
    resp,err:=http.Get(url)
    if err!=nil{
        out<-fmt.Sprint(err) //send to channel ch
        return
    }
nbytes,err:=io.Copy(ioutil.Discard,resp.Body)
resp.Body.Close()
if err!=nil{
    out<-fmt.Sprint("while reading %s:%v",url,err)
    return
}
secs:=time.Since(start).Seconds()
out<-fmt.Sprintf("%.2fs %7d %s",secs,nbytes,url)
}


func handler(w http.ResponseWriter,r *http.Request){
    mu.Lock()
    count++
    mu.Unlock()

    fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
    for k, v := range r.Header {
        fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
    }
    fmt.Fprintf(w, "Host = %q\n", r.Host)
    fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
    if err := r.ParseForm(); err != nil {
        log.Print(err)
    }
    for k, v := range r.Form {
        fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
    }

}

func counter(w http.ResponseWriter,r *http.Request){
    mu.Lock()
    fmt.Fprintf(w,"Count %d\n",count)
    mu.Unlock()
}