package matchwatcher_test

import (
	"testing"

	"fmt"
	. "github.com/Emyrk/fortnitediscord/matchwatcher"
)

func TestAPI(t *testing.T) {
	s, err := GetStatisticsJson("Emyrk")
	fmt.Println(err, s)
}
