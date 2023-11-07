package cache

import (
	cache "github.com/zekroTJA/timedmap"
	"reforce.pattern/config"
	"time"
)

var ch = Init()
var cfg = config.CFG.Cache

func Init() *cache.TimedMap {
	return cache.New(1 * time.Minute)
}

func Set(key interface{}, val Value, dur time.Duration) {
	ch.Set(key, val, dur)
}

func Get(key interface{}) (Value, bool) {
	token, ok := ch.GetValue(key).(Value)
	return token, ok
}

func Remove(key string) {
	ch.Remove(key)
}
func Flush() {
	ch.Flush()
}
