package main

import (
	"io"
	"log"
	"os"
)

type Voicer interface {
	voice()
}

type Duck struct {
	name string
	age  int
}

type Jaba struct {
	name  string
	color string
}

func (j *Jaba) voice() {
	log.Println("Kwa-kwa-kwa")
}

func (d *Duck) voice() {
	log.Println("Krya-krya-krya")
}

func trying(v Voicer, writer io.Writer) error {
	v.voice()
	return nil
}

func main() {
	d := Duck{
		name: "John",
		age:  15,
	}
	j := Jaba{
		name:  "Lena",
		color: "blue",
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_ = trying(&j, f)
	_ = trying(&d, f)

}
