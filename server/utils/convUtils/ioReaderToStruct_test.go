//
// EPITECH PROJECT, 2025
// AREA
// File description:
// ioReaderToStruct_test
//

package conv_utils

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"
)

func TestIoReaderToStruct(t *testing.T) {
	t.Run("Default Test", func(t *testing.T) {
		person := &personTest{
			Name: "Miaou",
			Age:  17,
		}
		b, _ := json.Marshal(person)
		r := bytes.NewReader(b)
		c := io.NopCloser(r)
		m, err := IoReaderToStruct[personTest](&c)
		if err != nil {
			t.Errorf(err.Error())
			return
		}
		if *m != *person {
			t.Errorf("The conversion did not work")
			return
		}
	})
	t.Run("Empty test", func(t *testing.T) {
		_, err := IoReaderToStruct[personTest](nil)
		if err == nil {
			t.Errorf("Error was not catched")
			return
		}
	})
}
