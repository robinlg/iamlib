// Copyright 2025 Robin Liu <robinliu27@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/robinlg/onexblog. The professional
// version of this repository is https://github.com/robinlg/onexblog.

package fileutil

import "os"

// EnsureDirAll will create a directory at the given path along with any necessary parents if they don't already exist.
func EnsureDirAll(path string) error {
	return os.MkdirAll(path, 0755)
}
