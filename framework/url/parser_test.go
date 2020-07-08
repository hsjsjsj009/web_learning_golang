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
		{"dfg"}, // Path => /dfg or /dfg/
		{"dfg","asd",":dfg"}, // Path => /dfg/asd/:dfg
		{"asd",":asd","dfg",":dfg"}, // Path => /asd/:asd/dfg/:dfg
	}
	url := "/asd/asd"

	param,found,idx := ParsePath(url,list)

	output := map[string]string{
		"asd":"asd",
	}

	assertions.Equal(0,idx)
	assertions.Equal(true,found)
	assertions.Equal(fmt.Sprint(output),fmt.Sprint(param))

	url = "/dfg"

	param,found,idx = ParsePath(url,list)

	output = map[string]string{}

	assertions.Equal(1,idx)
	assertions.Equal(true,found)
	assertions.Equal(fmt.Sprint(output),fmt.Sprint(param))

	url = "/dfg/"

	param,found,idx = ParsePath(url,list)

	output = map[string]string{}

	assertions.Equal(1,idx)
	assertions.Equal(true,found)
	assertions.Equal(fmt.Sprint(output),fmt.Sprint(param))

	url = "/dfg/asd/dfg"

	param,found,idx = ParsePath(url,list)

	output = map[string]string{
		"dfg":"dfg",
	}

	assertions.Equal(2,idx)
	assertions.Equal(true,found)
	assertions.Equal(fmt.Sprint(output),fmt.Sprint(param))

	url = "/asd/asd/dfg/dfg"

	param,found,idx = ParsePath(url,list)

	output = map[string]string{
		"asd":"asd",
		"dfg":"dfg",
	}

	assertions.Equal(3,idx)
	assertions.Equal(true,found)
	assertions.Equal(fmt.Sprint(output),fmt.Sprint(param))

}

func TestParsePathNotFound(t *testing.T){
	assertions := assert.New(t)

	list := [][]string{{"asd",":asd"},{"dfg"},{"asd"}}

	url := "/dfg/asd/"

	_,found,idx := ParsePath(url,list)

	assertions.Equal(-1,idx)
	assertions.Equal(false,found)

}