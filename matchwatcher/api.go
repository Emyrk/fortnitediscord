package matchwatcher

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"
)

var nodepath = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "Emyrk", "fortnitediscord", "matchwatcher", "gonode", "index.js")

var _ = fmt.Println

func GetStatisicsWithTimeout(name string, timeout int) (*PlayerStats, error) {
	c1 := make(chan *PlayerStats, 1)
	e1 := make(chan error, 1)
	go func() {
		res, err := GetStatisics(name)
		if err != nil {
			e1 <- err
			return
		}
		c1 <- res

	}()

	// Here's the `select` implementing a timeout.
	// `res := <-c1` awaits the result and `<-Time.After`
	// awaits a value to be sent after the timeout of
	// 1s. Since `select` proceeds with the first
	// receive that's ready, we'll take the timeout case
	// if the operation takes more than the allowed 1s.
	select {
	case res := <-c1:
		return res, nil
	case err := <-e1:
		return nil, err
	case <-time.After(time.Duration(timeout) * time.Second):
		fmt.Errorf("Timeout on node execute")
	}
	return nil, fmt.Errorf("impossible")
}

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
