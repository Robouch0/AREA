//
// EPITECH PROJECT, 2025
// AREA
// File description:
// ioReaderToMap_test
//

package conv_utils

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"
)

func TestIoReaderToMap(t *testing.T) {
	t.Run("Default Test", func(t *testing.T) {

		b, _ := json.Marshal(&personTest{
			Name: "Miaou",
			Age:  17,
		})
		r := bytes.NewReader(b)
		c := io.NopCloser(r)
		m, err := IoReaderToMap(&c)
		if err != nil {
			t.Errorf(err.Error())
			return
		}
		if _, ok := m["name"]; !ok {
			t.Errorf("Conversion did not work")
		}
		if _, ok := m["age"]; !ok {
			t.Errorf("Conversion did not work")
		}
	})
	t.Run("Empty test", func(t *testing.T) {
		_, err := IoReaderToMap(nil)
		if err == nil {
			t.Errorf("Error was not catched")
			return
		}
	})
}
