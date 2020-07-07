package url

import (
	"strings"
)

func ParsePath(path string,pathList [][]string) (map[string]string,bool) {
	var param map[string]string
	pathSplit := PathToSlice(path)

	found := false

	for _,obj := range pathList {
		if param,found = checkPathConfig(pathSplit,obj) ; found {
			break
		}
	}

	return param,found
}

func checkPathConfig(path []string,pattern []string) (map[string]string,bool){
	if len(path) != len(pattern){
		return nil,false
	}

	data := map[string]string{}

	same := true

	for idx,obj := range path {
		patternIdx := pattern[idx]

		if patternIdx == obj {
			continue
		}

		if string(patternIdx[0]) == ":" && len(obj) != 0{
			patternIdx = strings.Replace(patternIdx,":","",1)
			data[patternIdx] = obj
			continue
		}

		same = false
		break

	}

	return data,same
}

func PathToSlice(path string) []string{
	path = strings.Replace(path,"/","",1)
	path = strings.ReplaceAll(path,"/"," ")
	pathSplit := strings.Split(path," ")
	lengthPath := len(pathSplit)
	if pathSplit[lengthPath - 1] == "" {
		pathSplit = pathSplit[:lengthPath - 1]
	}

	return pathSplit
}