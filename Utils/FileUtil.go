package Utils

import (
	"fmt"
	"io/ioutil"
)

func GetPostPictures(key string) []string {

	files, err := ioutil.ReadDir("./PostsPicture/" + key)
	ans := make([]string, cap(files))
	if err != nil {
		fmt.Println(err)
	}
	for k, f := range files {
		ans[k] = f.Name()
	}

	return ans

}

func GetStrinFile(filepath string) (string, error) {

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
		return "", err
	}
	return string(data), nil

}
