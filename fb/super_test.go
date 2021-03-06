package fb

import "testing"
import "fmt"

type SuperTest struct {
	t      *testing.T
	prefix string
}

func (st *SuperTest) Equal(expected interface{}, actual interface{}) {
	if expected != actual {
		st.t.Error(fmt.Sprintf("%s: Expected %v, got %v", st.prefix, expected, actual))
	}
}
