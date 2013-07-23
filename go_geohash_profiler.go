package main

import "time"
import "fmt"

// 2-Similar implementations of geohash-ing
import "github.com/gnagel/go-geohash"
import "github.com/broady/gogeohash"

var precision uint8 = 12
var longitude = 112.5584
var latitude = 37.8324
var geostr = "ww8p1r4t8"
var neighbor = "dqcjq"
var directions = [2]int{1, 0}
var num_loops int = 1000 * 1000
var format_tag = "[%-30s][%-12s]\t MS Per/Call %.4f (ms)\t Total for %d Calls = %5f (ms)\n"

func main() {
	var start time.Time
	var ms_total float64
	var ms_per_call float64

	//
	// "github.com/broady/gogeohash"
	//
	start = time.Now()
	for i := 0; i < num_loops; i++ {
		geohash.Encode(latitude, longitude)
	}
	ms_total = float64(time.Now().Sub(start).Nanoseconds() / int64(1000000))
	ms_per_call = ms_total / float64(num_loops)
	fmt.Printf(format_tag, "encode", "gogeohash", ms_per_call, num_loops, ms_total)

	start = time.Now()
	for i := 0; i < num_loops; i++ {
		geohash.Decode(geostr)
	}
	ms_total = float64(time.Now().Sub(start).Nanoseconds() / int64(1000000))
	ms_per_call = ms_total / float64(num_loops)
	fmt.Printf(format_tag, "decode", "gogeohash", ms_per_call, num_loops, ms_total)

	//
	// "github.com/gnagel/go-geohash"
	//
	start = time.Now()
	for i := 0; i < num_loops; i++ {
		go_geohash.Encode(latitude, longitude, precision)
	}
	ms_total = float64(time.Now().Sub(start).Nanoseconds() / int64(1000000))
	ms_per_call = ms_total / float64(num_loops)
	fmt.Printf(format_tag, "encode", "go-geohash", ms_per_call, num_loops, ms_total)

	start = time.Now()
	for i := 0; i < num_loops; i++ {
		go_geohash.Decode(geostr)
	}
	ms_total = float64(time.Now().Sub(start).Nanoseconds() / int64(1000000))
	ms_per_call = ms_total / float64(num_loops)
	fmt.Printf(format_tag, "decode", "go-geohash", ms_per_call, num_loops, ms_total)

	start = time.Now()
	for i := 0; i < num_loops; i++ {
		go_geohash.DecodeBoundBox(geostr)
	}
	ms_total = float64(time.Now().Sub(start).Nanoseconds() / int64(1000000))
	ms_per_call = ms_total / float64(num_loops)
	fmt.Printf(format_tag, "decode_bbox", "go-geohash", ms_per_call, num_loops, ms_total)

	start = time.Now()
	for i := 0; i < num_loops; i++ {
		go_geohash.Neighbor(neighbor, directions)
	}
	ms_total = float64(time.Now().Sub(start).Nanoseconds() / int64(1000000))
	ms_per_call = ms_total / float64(num_loops)
	fmt.Printf(format_tag, "neighbor", "go-geohash", ms_per_call, num_loops, ms_total)
}
