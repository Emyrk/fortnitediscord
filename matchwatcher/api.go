package matchwatcher

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

var nodepath = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "Emyrk", "fortnitediscord", "matchwatcher", "gonode", "index.js")

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
	data := exec.Command("node", nodepath, fmt.Sprintf(`"%s"`, name))
	// fmt.Println("node", nodepath, fmt.Sprintf(`"%s"`, name))

	d, err := data.Output()
	if err != nil {
		fmt.Println("Error output:", string(d))
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
	// fmt.Println(string(j))
	return j, nil
}
