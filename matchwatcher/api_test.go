package matchwatcher_test

import (
	"testing"

	"fmt"
	. "github.com/Emyrk/fortnitediscord/matchwatcher"
)

func TestAPI(t *testing.T) {
	s, err := GetStatisticsJson("Emyrks")
	fmt.Println(err, string(s))
	p, err := GetStatisics("Emyrks")
	fmt.Println(err, p)
}
