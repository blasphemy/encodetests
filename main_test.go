package jobinator

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"testing"

	"github.com/vmihailenco/msgpack"
)

type testData struct {
	AString string
	AInt    int
}

type testData2 struct {
	TD           testData
	LongerString string
	BigNum       int64
	BigFloat     float64
}

var (
	td = testData{
		AString: "test",
		AInt:    9,
	}
	td2 = testData2{
		TD:           td,
		LongerString: "this is a very long string. this is a very long string. this is a very long string.",
		BigNum:       9923492394,
		BigFloat:     1.234343234234,
	}
)

func BenchmarkJsonSimple(b *testing.B) {
	for n := 0; n < b.N; n++ {
		by, _ := json.Marshal(td)
		out := &testData{}
		json.Unmarshal(by, out)
	}
	return
}

func BenchmarkMsgPackSimple(b *testing.B) {
	for n := 0; n < b.N; n++ {
		by, _ := msgpack.Marshal(td)
		out := &testData{}
		msgpack.Unmarshal(by, out)
	}
	return
}

func BenchmarkGobSimple(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b := &bytes.Buffer{}
		gobencoder := gob.NewEncoder(b)
		gobdecoder := gob.NewDecoder(b)
		gobencoder.Encode(td)
		t2 := &testData{}
		gobdecoder.Decode(t2)
	}
}

func BenchmarkJsonComplex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		by, err := json.Marshal(td2)
		if err != nil {
			b.Fail()
		}
		out := &testData{}
		err = json.Unmarshal(by, out)
		if err != nil {
			b.Fail()
		}
	}
	return
}

func BenchmarkMsgPackComplex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		by, err := msgpack.Marshal(td2)
		if err != nil {
			b.Fail()
		}
		out := &testData{}
		err = msgpack.Unmarshal(by, out)
		if err != nil {
			b.Fail()
		}
	}
	return
}

func BenchmarkGobComplex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b := &bytes.Buffer{}
		gobencoder := gob.NewEncoder(b)
		gobdecoder := gob.NewDecoder(b)
		gobencoder.Encode(td2)
		t2 := &testData2{}
		gobdecoder.Decode(t2)
	}
}
