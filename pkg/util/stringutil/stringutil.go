// Copyright 2025 Robin Liu <robinliu27@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/robinlg/onexblog. The professional
// version of this repository is https://github.com/robinlg/onexblog.

package stringutil

import "unicode/utf8"

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func FindString(array []string, str string) int {
	for index, s := range array {
		if str == s {
			return index
		}
	}
	return -1
}

func StringIn(str string, array []string) bool {
	return FindString(array, str) > -1
}
