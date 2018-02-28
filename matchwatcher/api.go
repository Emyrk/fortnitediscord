package matchwatcher

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
)

var _ = fmt.Println

func GetStatisics(name string) (*PlayerStats, error) {
	data, err := GetStatisticsJson(name)
	if err != nil {
		return nil, err
	}

	p := new(PlayerStats)
	err = json.Unmarshal(data, p)

	return p, err
}

func GetStatisticsJson(name string) ([]byte, error) {
	data := exec.Command("node", "gonode/index.js")
	var _ = data

	d, err := data.Output()
	if err != nil {
		return []byte{}, err
	}

	//fmt.Println(string(d))

	//reader := strings.NewReader(string(d))
	//b := bufio.NewReader(reader)
	//
	//var text []byte
	//for {
	//	l, _, err := b.ReadLine()
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Println(string(l))
	//	text = append(text, l...)
	//}

	r, _ := regexp.Compile("{.*")
	//j := r.Find([]byte(text))
	j := r.Find(d)
	fmt.Println(string(j))
	return j, nil
}