package main

import (
	"path/filepath"
	"strings"
)

// https://leetcode.com/problems/simplify-path/description/
// Runtime 0 ms Beats 100% Memory 3.1 MB Beats 38.43%
func simplifyPath(path string) string {
	directories := []string{}
	for _, v := range strings.Split(path, "/") {
		if v == "" || v == "." {
			continue
		} else if v == ".." {
			if len(directories) > 0 {
				directories = directories[:len(directories)-1]
			}
		} else {
			directories = append(directories, v)
		}
	}

	return "/" + strings.Join(directories, "/")
}

func simplifyPath2(path string) string {
	path += "/"
	directories := []string{}
	segment := ""
	for _, v := range path {
		if v == 47 {
			if segment == ".." {
				if len(directories) > 0 {
					directories = directories[:len(directories)-1]
				}
			} else if segment != "." && segment != "" {
				directories = append(directories, segment)
			}
			segment = ""
		} else {
			segment += string(v)
		}
	}

	answer := ""
	for _, v := range directories {
		answer += "/" + v
	}

	if answer == "" {
		answer += "/"
	}
	return answer
}
func simplifyPath3(path string) string {
	return filepath.Clean(path)
}
