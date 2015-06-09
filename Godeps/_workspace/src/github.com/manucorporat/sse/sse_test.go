package sse

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeOnlyData(t *testing.T) {
	w := new(bytes.Buffer)
	err := Encode(w, Event{
		Data: "junk\n\njk\nid:fake",
	})
	assert.NoError(t, err)
	assert.Equal(t, w.String(), "data: junk\\n\\njk\\nid:fake\n\n")
}

func TestEncodeWithEvent(t *testing.T) {
	w := new(bytes.Buffer)
	err := Encode(w, Event{
		Event: "t\n:<>\r\test",
		Data:  "junk\n\njk\nid:fake",
	})
	assert.NoError(t, err)
	assert.Equal(t, w.String(), "event: t\\n:<>\\r\test\ndata: junk\\n\\njk\\nid:fake\n\n")
}

func TestEncodeWithId(t *testing.T) {
	w := new(bytes.Buffer)
	err := Encode(w, Event{
		Id:   "t\n:<>\r\test",
		Data: "junk\n\njk\nid:fa\rke",
	})
	assert.NoError(t, err)
	assert.Equal(t, w.String(), "id: t\\n:<>\\r\test\ndata: junk\\n\\njk\\nid:fa\\rke\n\n")
}

func TestEncodeWithRetry(t *testing.T) {
	w := new(bytes.Buffer)
	err := Encode(w, Event{
		Retry: 11,
		Data:  "junk\n\njk\nid:fake\n",
	})
	assert.NoError(t, err)
	assert.Equal(t, w.String(), "retry: 11\ndata: junk\\n\\njk\\nid:fake\\n\n\n")
}

func TestEncodeWithEverything(t *testing.T) {
	w := new(bytes.Buffer)
	err := Encode(w, Event{
		Event: "abc",
		Id:    "12345",
		Retry: 10,
		Data:  "some data",
	})
	assert.NoError(t, err)
	assert.Equal(t, w.String(), "id: 12345\nevent: abc\nretry: 10\ndata: some data\n\n")
}

func TestEncodeMap(t *testing.T) {
	w := new(bytes.Buffer)
	err := Encode(w, Event{
		Event: "a map",
		Data: map[string]interface{}{
			"foo": "b\n\rar",
			"bar": "id: 2",
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, w.String(), "event: a map\ndata: {\"bar\":\"id: 2\",\"foo\":\"b\\n\\rar\"}\n\n")
}

func TestEncodeSlice(t *testing.T) {
	w := new(bytes.Buffer)
	err := Encode(w, Event{
		Event: "a slice",
		Data:  []interface{}{1, "text", map[string]interface{}{"foo": "bar"}},
	})
	assert.NoError(t, err)
	assert.Equal(t, w.String(), "event: a slice\ndata: [1,\"text\",{\"foo\":\"bar\"}]\n\n")
}

func TestEncodeStruct(t *testing.T) {
	myStruct := struct {
		A int
		B string `json:"value"`
	}{1, "number"}

	w := new(bytes.Buffer)
	err := Encode(w, Event{
		Event: "a struct",
		Data:  myStruct,
	})
	assert.NoError(t, err)
	assert.Equal(t, w.String(), "event: a struct\ndata: {\"A\":1,\"value\":\"number\"}\n\n")

	w.Reset()
	err = Encode(w, Event{
		Event: "a struct",
		Data:  &myStruct,
	})
	assert.NoError(t, err)
	assert.Equal(t, w.String(), "event: a struct\ndata: {\"A\":1,\"value\":\"number\"}\n\n")
}

func TestEncodeInteger(t *testing.T) {
	w := new(bytes.Buffer)
	err := Encode(w, Event{
		Event: "an integer",
		Data:  1,
	})
	assert.NoError(t, err)
	assert.Equal(t, w.String(), "event: an integer\ndata: 1\n\n")
}

func TestEncodeFloat(t *testing.T) {
	w := new(bytes.Buffer)
	err := Encode(w, Event{
		Event: "Float",
		Data:  1.5,
	})
	assert.NoError(t, err)
	assert.Equal(t, w.String(), "event: Float\ndata: 1.5\n\n")
}

func TestEncodeStream(t *testing.T) {
	w := new(bytes.Buffer)

	Encode(w, Event{
		Event: "float",
		Data:  1.5,
	})

	Encode(w, Event{
		Id:   "123",
		Data: map[string]interface{}{"foo": "bar", "bar": "foo"},
	})

	Encode(w, Event{
		Id:    "124",
		Event: "chat",
		Data:  "hi! dude",
	})
	assert.Equal(t, w.String(), "event: float\ndata: 1.5\n\nid: 123\ndata: {\"bar\":\"foo\",\"foo\":\"bar\"}\n\nid: 124\nevent: chat\ndata: hi! dude\n\n")
}