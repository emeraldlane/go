package collections

import (
  "bytes"
  "fmt"
)

type Element struct {
  fmt.Stringer
  Next *Element
  Value interface{}
}

type LinkedList struct {
  root *Element
}

func (l LinkedList) InsertFront(value interface{}) {

}

func (e Element) String() string {
  return fmt.Sprint(e.Value)
}

func (l LinkedList) String() string {
  var buffer bytes.Buffer
  first := l.root
  if first != nil {
    buffer.WriteString(fmt.Sprint(first))
    for next := first.Next; next != nil && next != first; next = next.Next {
      buffer.WriteString(", ")
      buffer.WriteString(fmt.Sprint(*next))
    }
  }
  return buffer.String()
}
