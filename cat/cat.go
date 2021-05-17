package main

import (
	"log"
	"os"
)

func getFile(fName string) (*os.File, func(), error) {
	file, err := os.Open(fName)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		file.Close()
	}, err
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Command does not have file argument")
	}

	fName := os.Args[1]

	file, closer, err := getFile(fName)
	defer closer()

	if err != nil {
		log.Fatalf("File %s could not be opened\n", fName)
	}

	data := make([]byte, 2048)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalf("File %s could not be read\n", fName)
	}

	count, err = os.Stdout.Write(data[:count])
	if err != nil {
		log.Fatalf("Data of file %s could not be written to stdout\n", fName)
	}

}
