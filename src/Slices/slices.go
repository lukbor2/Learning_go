package main 

import (
	"fmt"
	"sort"
)

func InsertStringSliceCopy(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}

func main() {
	s := []string{"A", "B", "C", "D", "E", "F", "G"}
	t := s[:5]
	// [A B C D E]
	u := s[3 : len(s)-1] // [D E F]
	fmt.Println("Before Changing the value: ")
	fmt.Println(s, t, u)
	u[1] = "x" //The value changes in all 3 slices because all slices refer to the same underlying array
	fmt.Println("After changing the value: ")
	fmt.Println( s, t, u)
	
	//Example of append()
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println("Example of append()")
	fmt.Println(slice1, slice2)
	
	//Example of copy()
	slice3 := []int{1,2,3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println("Example of copy()")
	fmt.Println(slice3, slice4)
	
	//Indexing and Slicing Slices
	s1 := []string{"A", "B", "C", "D", "E", "F", "G"}
	t1 := s1[2:6]
	fmt.Println("Indexing and slicing")
	fmt.Println(t1, s1, "=", s1[:4], "+", s1[4:])
	s1[3] = "x"
	t1[len(t1)-1] = "y"
	fmt.Println(t1, s1, "=", s1[:4], "+", s1[4:])
	
	//Iterating slices without modifying values
	amounts := []float64{237.81, 261.87, 273.93, 279.99, 281.07, 303.17, 231.47, 227.33, 209.23, 197.09}
	sum := 0.0
	for _, amount := range amounts {
		sum += amount
	}
	fmt.Println("Iterating no changes")
	fmt.Printf("Sum %.1f --> %.1f\n", amounts, sum)
	
	//Iterating slices modifying values
	for i := range amounts {
		amounts[i] *= 1.05
		sum += amounts[i]
	}
	fmt.Println("Iterating making changes")
	fmt.Printf("Sum %.1f --> %.1f\n", amounts, sum)
	
	//Using append() function
	s2 := []string{"A", "B", "C", "D", "E", "F", "G"}
	t2 := []string{"K", "L", "M", "N"}
	u2 := []string{"m", "n", "o","p", "q", "r"}
	s2 = append(s2, "h", "i", "j") //Append individual values
	s2 = append(s2, t2...)			//Append all of a slice values
	s2 = append(s2, u2[2:5]...)		//Append a sub-slice
	b := []byte{'U', 'V'}
	letters := "wxy"
	b = append(b, letters...)	//Append a string's bytes to a byte slice
	fmt.Println("Using the append() function")
	fmt.Printf("%v\n%s\n", s2, b)
	
	//Using the InsertStringSliceCopy() function
	s3 := []string{"M", "N", "O", "P", "Q", "R"}
	x := InsertStringSliceCopy(s3, []string{"a", "b", "c"}, 0)	// At the front
	y := InsertStringSliceCopy(s3, []string{"x", "y"}, 3) 		// In the middle
	z := InsertStringSliceCopy(s3, []string{"z"}, len(s3))		// At the end
	fmt.Println("Using the InsertStringSliceCopy function")
	fmt.Printf("%v\n%v\n%v\n%v\n", s3, x, y, z)
	
	//Sorting slices with standard functions
	files := []string{"Test.conf", "util.go", "Makefile", "misc.go", "main.go"}
	fmt.Printf("Unsorted:%q\n", files)
	sort.Strings(files) // Standard library sort function
	fmt.Printf("Underlying bytes: %q\n", files)
	
	//Searching a slice using standard search function
	target := "Makefile"
	sort.Strings(files)
	fmt.Printf("%q\n", files)
	i := sort.Search(len(files), func(i int) bool { return files[i] >= target })
	if i < len(files) && files[i] == target {
		fmt.Println("Using standard sort.Search")
		fmt.Printf("found \"%s\" at files[%d]\n", files[i], i)
	}

}

