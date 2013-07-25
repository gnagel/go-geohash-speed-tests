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
var format_tag = "[%-30s][%-12s]\t NS Per/Call %10.d (ns)\t Total for %d Calls = %5f (ms)\n"

func main() {
	var start time.Time
	var end  time.Time
	var ms_total float64
	var ns_per_call int64

	//
	// "github.com/broady/gogeohash"
	//
	start = time.Now()
	for i := 0; i < num_loops; i++ {
		geohash.Encode(latitude, longitude)
	}
	end = time.Now()
	ms_total = float64(end.Sub(start).Nanoseconds() / int64(1000000))
	ns_per_call = end.Sub(start).Nanoseconds() / int64(num_loops)
	fmt.Printf(format_tag, "encode", "gogeohash", ns_per_call, num_loops, ms_total)

	start = time.Now()
	for i := 0; i < num_loops; i++ {
		geohash.Decode(geostr)
	}
	end = time.Now()
	ms_total = float64(end.Sub(start).Nanoseconds() / int64(1000000))
	ns_per_call = end.Sub(start).Nanoseconds() / int64(num_loops)
	fmt.Printf(format_tag, "decode", "gogeohash", ns_per_call, num_loops, ms_total)

	//
	// "github.com/gnagel/go-geohash"
	//
	start = time.Now()
	for i := 0; i < num_loops; i++ {
		go_geohash.Encode(latitude, longitude, precision)
	}
	end = time.Now()
	ms_total = float64(end.Sub(start).Nanoseconds() / int64(1000000))
	ns_per_call = end.Sub(start).Nanoseconds() / int64(num_loops)
	fmt.Printf(format_tag, "encode", "go-geohash", ns_per_call, num_loops, ms_total)

	start = time.Now()
	for i := 0; i < num_loops; i++ {
		go_geohash.Decode(geostr)
	}
	end = time.Now()
	ms_total = float64(end.Sub(start).Nanoseconds() / int64(1000000))
	ns_per_call = end.Sub(start).Nanoseconds() / int64(num_loops)
	fmt.Printf(format_tag, "decode", "go-geohash", ns_per_call, num_loops, ms_total)

	start = time.Now()
	for i := 0; i < num_loops; i++ {
		go_geohash.DecodeBoundBox(geostr)
	}
	end = time.Now()
	ms_total = float64(end.Sub(start).Nanoseconds() / int64(1000000))
	ns_per_call = end.Sub(start).Nanoseconds() / int64(num_loops)
	fmt.Printf(format_tag, "decode_bbox", "go-geohash", ns_per_call, num_loops, ms_total)

	start = time.Now()
	for i := 0; i < num_loops; i++ {
		go_geohash.Neighbor(neighbor, directions)
	}
	end = time.Now()
	ms_total = float64(end.Sub(start).Nanoseconds() / int64(1000000))
	ns_per_call = end.Sub(start).Nanoseconds() / int64(num_loops)
	fmt.Printf(format_tag, "neighbor", "go-geohash", ns_per_call, num_loops, ms_total)
}
