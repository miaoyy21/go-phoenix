package base

import (
	"encoding/json"
	"testing"
)


type Allen struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Description string `json:"description"`
}

//func BenchmarkEncodingJson(b *testing.B) {
//	b.Logf("BenchmarkEncodingJson Size :: %9.2f M \n", float64(bytes.NewBufferString(data).Len())/float64(1<<20))
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		var users []Allen
//
//		if err := json.Unmarshal([]byte(data), &users); err != nil {
//			b.Fatal(err)
//		}
//	}
//}
//
//func BenchmarkSonic1(b *testing.B) {
//	b.Logf("BenchmarkSonic1 Size :: %9.2f M \n", float64(bytes.NewBufferString(data).Len())/float64(1<<20))
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		var users []Allen
//
//		if err := sonic.Unmarshal([]byte(data), &users); err != nil {
//			b.Fatal(err)
//		}
//	}
//}
//
//func BenchmarkSonic2(b *testing.B) {
//	b.Logf("BenchmarkSonic2 Size :: %9.2f M \n", float64(bytes.NewBufferString(data).Len())/float64(1<<20))
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		var users []Allen
//
//		if err := sonic.UnmarshalString(data, &users); err != nil {
//			b.Fatal(err)
//		}
//	}
//}

type Bob struct {
	Id   int    `json:"id"`
	Name string `json:"aaa,bbb"`
}

func TestCompareMap(t *testing.T) {
	var bob = Bob{
		Id:   123,
		Name: "BOB",
	}

	bs, err := json.Marshal(bob)
	if err != nil {
		t.Fatalf("%s \n", err.Error())
	}

	t.Logf("JSON %s \n", string(bs))
}