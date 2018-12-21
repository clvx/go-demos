package concepts

// _Defer_ is used to ensure that a function call is
// performed later in a program's execution, usually for
// purposes of cleanup. `defer` is often used where e.g.
// `ensure` and `finally` would be used in other languages.

import "fmt"
import "os"

// Suppose we wanted to create a file, write to it,
// and then close when we're done. Here's how we could
// do that with `defer`.
func main() {

    }

func CreateFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func WriteFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")

}

func CloseFile(f *os.File) {
    fmt.Println("closing")
    f.Close()
}

