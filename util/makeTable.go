package util

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-yaml/yaml"
)

type Tables struct {
	B3 [][]string `yaml:"B3"`
	B4 [][]string `yaml:"B4"`
	M1 [][]string `yaml:"M1"`
}

var (
	text       string
	B3, B4, M1 string
	l          = []string{"B3", "B4", "M1"}
	row        = `%s & %s & %s \\ \hline
`
)

func MakeTable(filename string) string {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	t := Tables{}
	err = yaml.Unmarshal(buf, &t)
	if err != nil {
		log.Fatal(err)
	}

	_, err = yaml.Marshal(&t)
	if err != nil {
		log.Fatal(err)
	}

	for i, _ := range t.B3 {
		x := fmt.Sprintf(row, t.B3[i][0], t.B3[i][1], t.B3[i][2])
		B3 += x
	}

	for i, _ := range t.B4 {
		x := fmt.Sprintf(row, t.B4[i][0], t.B4[i][1], t.B4[i][2])
		B4 += x
	}

	for i, _ := range t.M1 {
		x := fmt.Sprintf(row, t.M1[i][0], t.M1[i][1], t.M1[i][2])
		M1 += x
	}

	for i, v := range []string{B3, B4, M1} {
		if len(v) > 0 {
			text += ReplaceParts(l[i], v)
			text += `
			`
		}
	}

	table := ReplaceTemplate(text)

	return table
}
