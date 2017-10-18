package main

import "fmt"

func main() {
    // manually create
    var m1 = make(map[string]int)
    m1["A"] = 1

   // literally create
   var m2 = map[string]string {
       "007":"James Bond",
       "MissionImpossible":"Ethan", 
       "SpiderMan": "Peter Park" }
   
    fmt.Printf("m1=%v len=%d\nm2=%v len=%d\n", m1, len(m1), m2, len(m2)) 
    
    // iterate 
    fmt.Println("----Iteration -----")
    for k,v := range m2 {
        fmt.Printf("key=%v value=%v\n", k, v) 
    }

    name, exists := m2["Matrix"]
    fmt.Printf("Matrix exists in map? boolean=%v , value of key=%v\n", exists, name)
    
    name, exists = m2["007"]
    fmt.Printf("007 exists in map? boolean=%v , value of key=%v\n", exists, name)

    // ===============  good way to check map key =========================
    if val, ok := m2["SpiderMan"]; ok {
        fmt.Printf("SpiderMan also exists: actor=%v\n", val)
    }

    var m3 = map[string]map[string]int {
        "John":map[string]int {
            "Math":80,
            "Physics":60 },
        "Mary":map[string]int {
            "Math":90,
            "Physics":50 }, // note the comma here 
    }
    fmt.Printf("m3 is %v\n", m3)



}
