package main
import (
   "encoding/csv"
   "fmt"
   "os"
  "bufio"
	"log"
  //"reflect"
  "strconv"
  "math"
  
)
type CsvDataLines struct {
   Column1 string
   Column2 string
}
func main() {
  
    
	var words []string        // initialize array to store words from txt file
  var score []string
  
  var newscores []float64



// Program starts here
//---------------------------------------------------------------

  fmt.Println("Enter the name of the file you would like to use (with the .txt extension):")
  var filename string

  fmt.Scanln(&filename)    // input filename

 file, err := os.Open(filename)    // open txt file
  
	if err != nil {
		log.Fatal(err)
	}

	Scanner := bufio.NewScanner(file)  // scan the input text file
	Scanner.Split(bufio.ScanWords)    // scan word by word

	for Scanner.Scan() {
		words = append(words, Scanner.Text())
	}

  
  csvData, err := ReadCsvFile("socialsent.csv")  // read csv file
   if err != nil {
      panic(err)
   }

//---------------------------------------------------------------


  
fmt.Println()
fmt.Println()
fmt.Println("[word: current_score]")
for _, word := range words {    // loop through all the words in the text file

  for _, line := range csvData {
     if line[0] == word{      // if the word we are searching for if found in csv
       fmt.Println(line[0] + ":", line[1])
       score = append(score, line[1])  // we append the score of that word to another array, where we will do the calulations
   
// if its not found in csv, do nothing
      
     }
    }
  }



  
      for _, i := range score {    // convert all our scores from strings to floats
        j, err := strconv.ParseFloat(i, 8)
        if err != nil {
            panic(err)
        }
        newscores = append(newscores, j)
    }


// now add up all the floats for the final cumalative score
 var result float64 = 0.0  
 for _, v := range newscores {  
  result += v  
 }  
// round the float to 2 decmial places
  newresult := math.Round(result*100)/100

      fmt.Println()
      fmt.Println("Final Score:",newresult)


 
	if newresult < -5.0 {
		fmt.Println("Stars: 1");
	} else if newresult >= -5.0 && newresult < -1.0 {
		fmt.Println("Stars: 2");
	} else if newresult >= -1.0 && newresult < 1.0 {
		fmt.Println("Stars: 3")
	} else if newresult >= 1.0 && newresult < 5.0 {
		fmt.Println("Stars: 4")
	}else if newresult >= 5.0 {
		fmt.Println("Stars: 5")
	}else {
		fmt.Println("Out of range")
	}


  
  
}


//---------------------------------------------------------------


func ReadCsvFile(filename string) ([][]string, error) {  // function to read csv file
   // Open CSV file
   fileContent, err := os.Open(filename)
   if err != nil {
      return [][]string{}, err
   }
   defer fileContent.Close()

   // Read File into a Variable
   lines, err := csv.NewReader(fileContent).ReadAll()
   if err != nil {
      return [][]string{}, err
   }
   return lines, nil
}



