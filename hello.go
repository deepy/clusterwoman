package main

import "clusterwoman/lib"

func main() {
	// debian-fish 3c:97:0e:a3:22:40
	// debian-cow 3c:97:0e:9e:cf:93
	// debian-panda 3c:97:0e:c3:93:76
	println("Waking up debian-panda")
	lib.WakeHost("3c:97:0e:c3:93:76")
}
