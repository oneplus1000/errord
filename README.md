# Error Extension

### Example

```GO
package main

import (
  "github.com/oneplus1000/errorx"
  "errors"
  "log"
)

var ErrXxx = errors.New("xxx fail")

func main() {
  err := makeError01()
  if err != nil {
    panic(errorx.StackString(err))
  }
}

func makeError01() error {
  err := makeError02()
  if err != nil {
    return errorx.Errorf("makeError02 fail : %w",err)
  }
  return nil
}

func makeError02() error {
  //do somthing and fail function return ErrXxx
  return errorx.Errorf("do somthing fail : %w",ErrXxx)
}


```
