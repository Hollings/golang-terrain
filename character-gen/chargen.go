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


func getOrCreateNames() string {
    file1, err1 := os.Open("generatedNames.txt")
    fullName := ""
    if err1 != nil {
        // If names file doesn't exist, generate one
        file, err := os.Create("generatedNames.txt")
        if err != nil {
            log.Fatal("Cannot create file", err)
        }
        defer file.Close()
        // createFile("generatedNames.txt")
        firstNames :=  generateNamesFromFile("names.txt",3)
        lastNames := generateNamesFromFile("last.txt",5)

        // for i := 0; i < 100; i++ {
        //     fmt.Print(firstNames[i] + " ")
        //     fmt.Println(lastNames[i])
        // }
        // createFile("generatedNames.txt")
        // file, err := os.Open("generatedNames.txt")
        // if err != nil {
           

        // defer file.Close()

        for i := 0; i < 10000; i++ {
            fmt.Print(firstNames[i] + " ")
            fmt.Println(lastNames[i])
            fmt.Fprintf(file,firstNames[i] + " " + lastNames[i] + "\n")
        }
        fullName = firstNames[10] + " " + lastNames[10] 
        
    }else{
        // If names file exists already, just pull a random line from it
        fullName = randomLine("generatedNames.txt")
        
    }
    defer file1.Close()
    return fullName
}

func generateNamesFromFile(fileName string, length int) [10000]string{
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
    var returnArray [10000]string 
    for i := 0; i < 10000; i++ {
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
func randomLine(fileName string) string{
	file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var lines []string
    for scanner.Scan() {
        lines = append(lines, string(scanner.Text()))
    }
    chosenLine := lines[rand.Intn(len(lines))]
    return chosenLine
}



func main() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	

    for i := 0; i < 5; i++ {
	    //fullName := firstNames[i] + " " + lastNames[i] 
	    

        fullName := getOrCreateNames()
        fmt.Println("---")
	   	fmt.Println(fullName)
	   	fmt.Println("the "+ randomLine("adverbs.txt") + " " + randomLine("adjectives.txt") +" "+[]string{0:"Male",1:"Female"}[rand.Intn(2)] + " " + randomLine("races.txt"))
	   // 	fmt.Println("likes " + randomLine("verbs.txt"))
	  	// fmt.Println("hates " + randomLine("verbs.txt"))

    }
    fmt.Println("---")	

    // fmt.Print(allSyllables[rand.Intn(len(allSyllables))])
    // fmt.Print(allSyllables[rand.Intn(len(allSyllables))])

    // if err := scanner.Err(); err != nil {
    //     log.Fatal(err)
    // }

}