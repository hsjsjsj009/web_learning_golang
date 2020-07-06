package url

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePathCorrect(t *testing.T){
	assertions := assert.New(t)

	list := [][]string{
		{"asd",":asd"}, // Path => /asd/:asd
		{"dfg"}, // Path => /dfg
		{"asd",""},  // Path => /asd/
		{"dfg","asd",":dfg"}, // Path => /dfg/asd/:dfg
		{"asd",":asd","dfg",":dfg"}, // Path => /asd/:asd/dfg/"dfg
	}
	url := "/asd/asd"

	param,found := ParsePath(url,list)

	output := map[string]string{
		"asd":"asd",
	}

	assertions.Equal(true,found)
	assertions.Equal(fmt.Sprint(output),fmt.Sprint(param))

	url = "/asd/"

	param,found = ParsePath(url,list)

	output = map[string]string{}

	assertions.Equal(true,found)
	assertions.Equal(fmt.Sprint(output),fmt.Sprint(param))

	url = "/dfg"

	param,found = ParsePath(url,list)

	output = map[string]string{}

	assertions.Equal(true,found)
	assertions.Equal(fmt.Sprint(output),fmt.Sprint(param))

	url = "/asd/"

	param,found = ParsePath(url,list)

	output = map[string]string{}

	assertions.Equal(true,found)
	assertions.Equal(fmt.Sprint(output),fmt.Sprint(param))


}

func TestParsePathNotFound(t *testing.T){
	assertions := assert.New(t)

	list := [][]string{{"asd",":asd"},{"dfg"},{"asd"}}
	url := "/asd/"

	_,found := ParsePath(url,list)

	assertions.Equal(false,found)

	url = "/dfg/asd/"

	_,found = ParsePath(url,list)

	assertions.Equal(false,found)

}