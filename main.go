package main

import (
	"fmt"

	"github.com/sahajavidya/learngows/stringutils"
	"github.com/sahajavidya/mathutils"
	log "github.com/sirupsen/logrus"
)

func main() {
	s := "sahajavidya"

	fmt.Println(stringutils.Upper(s))
	fmt.Println("Addition of two Number:", mathutils.Add(1, 4))
	log.Info("Using Third Party package")
	fmt.Println("Multiply of two Number:", mathutils.Multiply(2, 4))

}
