/*
@Time : 2020/4/2 3:11 PM
*/
package log

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestSetLevel(t *testing.T) {
	SetLevel(ErrorLevel)
	if infoLog.Writer() != ioutil.Discard || errorLog.Writer() != os.Stdout {
		t.Fatal("failed to set log level")
	}
	SetLevel(Disabled)
	if infoLog.Writer() != ioutil.Discard || errorLog.Writer() != ioutil.Discard {
		t.Fatal("failed to set log level")
	}
}
