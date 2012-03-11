package go_koans

import (
  "testing"
  "os"
  "fmt"
  "runtime"
  "io/ioutil"
  "path"
  "strings"
)

func TestKoans(t *testing.T) {
  testBasics()
  testStrings()
  testArrays()

  fmt.Printf("\n%c[32;1mYou won life. Good job.\n\n", 27)
}

func assert(o bool) {
  if !o {
    fmt.Printf("\n%c[35m%s\n\n", 27, __getRecentLine())
    os.Exit(1)
  }
}

func __getRecentLine() string {
  _, file, line, _ := runtime.Caller(2)
  buf, _ := ioutil.ReadFile(file)
  code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
  return fmt.Sprintf("%v:%d\n%s", path.Base(file), line, code)
}
