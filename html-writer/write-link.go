package htmlwriter

import (
	"fmt"
	"scheduler/model"
)

func CreateLink(l model.Link) {

	l.Name = "Test"
	fmt.Println(l)
}
