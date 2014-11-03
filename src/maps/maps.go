package main 

import (
	"fmt"
	"sort"
)

type Point struct{ x, y, z int }
        
func (point Point) String() string {
        return fmt.Sprintf("(%d,%d,%d)", point.x, point.y, point.z)
                
}

func main() {
        //Creating and printing a map
        massForPlanet := make(map[string]float64) // Same as: map[string]float64{}
        massForPlanet["Mercury"] = 0.06
        massForPlanet["Venus"] = 0.82
        massForPlanet["Earth"] = 1.00
        massForPlanet["Mars"] = 0.11
        fmt.Println(massForPlanet)
		
		//Using pointers as map index
		triangle := make(map[*Point]string, 3)
        triangle[&Point{89, 47, 27}] = "α"
        triangle[&Point{86, 65, 86}] = "β"
        triangle[&Point{7, 44, 45}] = "γ"
        fmt.Println("Priting Triangle....")
        fmt.Println(triangle)
        
        //Usisng for..loop with a map two variables
        populationForCity := map[string]int{"Istanbul": 12610000,"Karachi": 10620000, "Mumbai": 12690000, "Shanghai": 13680000}
        fmt.Println("Printing populationForCity")
        for city, population := range populationForCity {
                fmt.Printf("%-10s %8d\n", city, population)
        }
       
       //Map Lookups
        city := "Istanbul"
        if population, found := populationForCity[city]; found {fmt.Printf("%s's population is %d\n", city, population)} else {fmt.Printf("%s's population data is unavailable\n", city)}
        city = "Emerald City"
        _, present := populationForCity[city]
        fmt.Printf("%q is in the map == %t\n", city, present) 
        
        //Modifyig Maps
        fmt.Println("Modifying Maps")
        fmt.Println(len(populationForCity), populationForCity)
        delete(populationForCity, "Shanghai") //Delete
        fmt.Println(len(populationForCity), populationForCity)
        populationForCity["Karachi"] = 11620000 // Update
        fmt.Println(len(populationForCity), populationForCity)
        populationForCity["Beijing"] = 11290000 // Insert
        fmt.Println(len(populationForCity), populationForCity)

        //Changing Maps Keys
        oldKey, newKey := "Beijing", "Tokyo"
        value := populationForCity[oldKey]
        delete(populationForCity, oldKey)
        populationForCity[newKey] = value
        fmt.Println("Changing Maps Keys")
        fmt.Println(len(populationForCity), populationForCity)

        //Key-Ordered Map Iteration
        fmt.Println("Key-Ordered Maps")
        cities := make([]string, 0, len(populationForCity))
        for city := range populationForCity {
                cities = append(cities, city)
        }
        sort.Strings(cities)
        for _, city := range cities {
                fmt.Printf("%-10s %8d\n", city, populationForCity[city])
        }
        
        //Map Inversion
        fmt.Println("Map Inversion")
        cityForPopulation := make(map[int]string, len(populationForCity))
        for city, population := range populationForCity {
                cityForPopulation[population] = city
        }
        fmt.Println(cityForPopulation)
}

