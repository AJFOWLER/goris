package main
// run go mod init goris before this.
import (
	"fmt"
	"os"
	"log"
	"bufio"
)

type record struct {
	title string
	abstract string
}

func main(){
	// open a file
	f, err := os.Open("file.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer f.Close()

	//read the file scanning
	scan := bufio.NewScanner(f)

	// .ris files = between TY - AND ER - is the record.
	// Abstract is N2 or AB
	// Title is T1 or TT or TI
	//var holder map[string]record
	index := 0
	for scan.Scan(){
		// a record starts with TY - and ends with ER - 
		// first split into records
		fmt.Println(_)
		fmt.Println(v)
		t := scan.Text()
		key := t[:2]
		value := t[2:]
		fmt.Printf("Key is %s and value is %s \n", key, value)
		index++
	}
}

