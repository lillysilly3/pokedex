package pokecache

import (
    "testing"
    "time"
)

func TestAddGet(t *testing.T) {
    cases := []struct {
        key string
        val []byte
    }{
        {key: "https://pokeapi.co/api/v2/location-area", val: []byte("testdata")},
        {key: "https://pokeapi.co/api/v2/location-area?offset=20", val: []byte("testdata2")},
    }

    for _, c := range cases {
        cache := NewCache(5 * time.Second)
        cache.Add(c.key, c.val)
        got, ok := cache.Get(c.key)
        if !ok {
            t.Errorf("expected to find key %s", c.key)
            continue
        }
        if string(got) != string(c.val) {
            t.Errorf("expected %s, got %s", c.val, got)
        }
    }
}

func TestGetMissing(t *testing.T) {
    cache := NewCache(5 * time.Second)
    _, ok := cache.Get("https://pokeapi.co/api/v2/location-area")
    if ok {
        t.Error("expected key to not be found")
    }
}

func TestReapLoop(t *testing.T) {
    interval := 10 * time.Millisecond
    cache := NewCache(interval)

    cache.Add("https://pokeapi.co/api/v2/location-area", []byte("testdata"))

    _, ok := cache.Get("https://pokeapi.co/api/v2/location-area")
    if !ok {
        t.Error("expected to find key before reap")
    }

    time.Sleep(interval * 2)

    _, ok = cache.Get("https://pokeapi.co/api/v2/location-area")
    if ok {
        t.Error("expected key to be reaped")
    }
}

func TestOverwrite(t *testing.T) {
    cache := NewCache(5 * time.Second)
    cache.Add("key", []byte("first"))
    cache.Add("key", []byte("second"))
    got, ok := cache.Get("key")
    if !ok {
        t.Error("expected to find key")
    }
    if string(got) != "second" {
        t.Errorf("expected second, got %s", got)
    }
}

func TestReapOnlyExpired(t *testing.T) {
    interval := 10 * time.Millisecond
    cache := NewCache(interval)

    cache.Add("old", []byte("olddata"))
    time.Sleep(interval * 2)
    cache.Add("new", []byte("newdata")) // added after sleep, should survive

    time.Sleep(interval / 2) // wait a bit but not enough to reap "new"

    _, ok := cache.Get("old")
    if ok {
        t.Error("expected old key to be reaped")
    }
    _, ok = cache.Get("new")
    if !ok {
        t.Error("expected new key to still exist")
    }
}

func TestEmptyKeyAndVal(t *testing.T) {
    cache := NewCache(5 * time.Second)
    cache.Add("", []byte{})
    got, ok := cache.Get("")
    if !ok {
        t.Error("expected to find empty key")
    }
    if len(got) != 0 {
        t.Errorf("expected empty val, got %s", got)
    }
}

func TestConcurrent(t *testing.T) {
    cache := NewCache(5 * time.Second)
    for i := 0; i < 100; i++ {
        go cache.Add("key", []byte("val"))
        go cache.Get("key")
    }
}