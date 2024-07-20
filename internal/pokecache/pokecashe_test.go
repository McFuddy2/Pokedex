package pokecache

import (
	"testing"
	"time"
	
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond *10)
	if cache.cache == nil {
		t.Error("Cache is Nil")
	}

}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond *10)

	cases := []struct {
		inputKey	string
		inputVal	[]byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "bananas",
			inputVal: []byte("and apples"),
		},
		{
			inputKey: "",
			inputVal: []byte("im with blank"),
		},
	}
	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("correct key not found! Missing key: %v", cas.inputKey)
			continue
		} 
		if string(actual) != string(cas.inputVal) {
			t.Errorf("correct key not found! expected: %v actual: %v", 
				string(cas.inputVal), 
				string(actual))
		}
	}
}



func TestReap(t *testing.T) {
	interval := time.Millisecond *10
	cache :=NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(keyOne)
	if ok{
		t.Error("data has not been reaped!")
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond *10
	cache :=NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval/2)

	_, ok := cache.Get(keyOne)
	if !ok{
		t.Error("data has been prematurely reaped!")
	}
}