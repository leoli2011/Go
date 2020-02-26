package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]interface{}, 4)
	m["company"] = "wg"
	m["subjects"] = []string{"go", "docker", "fabric", "Test"}
	m["isok"] = true
	m["price"] = 99

	result, _ := json.Marshal(m)
	fmt.Printf("result=%s\n", string(result))
}
