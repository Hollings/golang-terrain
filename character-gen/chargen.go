package main

import (
    //"reflect"
    "math/rand"
    "os"
    "fmt"
    "bufio"
    "log"
    "regexp"
    "unicode/utf8"
  	"time"
    "strings"

)
var (

)
func print(a string){
	fmt.Println(a)
}
type sylList struct {
	start []string
	middle []string
	end []string
}

func generateNameFromFile(fileName string, length int) [15]string{
	file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var sylList sylList
    for scanner.Scan() {
        syllables := splitSyl(string(scanner.Text()))
        for _, element := range syllables.start {
        	sylList.start = append(sylList.start, element)
		}
		for _, element := range syllables.middle {
        	sylList.middle = append(sylList.middle, element)
		}
		for _, element := range syllables.end {
        	sylList.end = append(sylList.end, element)
		}
    }

    var returnArray [15]string 
    for i := 0; i < 15; i++ {
    	returnArray[i] = ""
	    returnArray[i] += (strings.Title(strings.ToLower(sylList.start[rand.Intn(len(sylList.start))])))
	    for i := 0; i < rand.Intn(length); i++ {
	   	    returnArray[i] += (strings.ToLower(strings.ToLower(sylList.middle[rand.Intn(len(sylList.middle))])))
	    }
	    returnArray[i] += (strings.ToLower(strings.ToLower(sylList.end[rand.Intn(len(sylList.end))]))	)
    }

    return returnArray
}
func splitSyl(name string) sylList{ 
	// Splits a string into syllables
	var startSyllables []string
	var endSyllables []string
	var syllables []string

	var ex = regexp.MustCompile(`([^aeiouAEIOU]*[aeiouAEIOU]*)|[aeiouAEIOU]*[^aeiouAEIOU]*[aeiouAEIOU]*`)
	i:=0
	for(utf8.RuneCountInString(name) > 0){
	 	end := ex.FindStringIndex(name)[1]
	 	if i==0 {
			startSyllables = append(startSyllables, name[0:end])
	 	}else if utf8.RuneCountInString(name[end:]) == 0{
	 		endSyllables = append(endSyllables, name[0:end])
 		}else{
 			syllables = append(syllables, name[0:end])
 		}
 		i++
		syllables = append(syllables, name[0:end])
		name = name[end:]	
	}
	returns := sylList{start: startSyllables,middle: syllables,end:endSyllables}
	return returns
}



func main() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	
	firstNames :=  generateNameFromFile("names.txt",2)
    lastNames := generateNameFromFile("last.txt",3)

    for i := 0; i < 15; i++ {
    	fmt.Print(firstNames[i] + " ")
    	fmt.Println(lastNames[i])
    }


    fmt.Println("---")
    

    // fmt.Print(allSyllables[rand.Intn(len(allSyllables))])
    // fmt.Print(allSyllables[rand.Intn(len(allSyllables))])

    // if err := scanner.Err(); err != nil {
    //     log.Fatal(err)
    // }

}