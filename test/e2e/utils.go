package e2e

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/otiai10/copy"
	"github.com/thoas/go-funk"
	"kusionstack.io/kpm/pkg/reporter"
)

// LoadFirstFileWithExt read the first file with extention 'ext' in 'dir' and return the content.
func LoadFirstFileWithExt(dir string, ext string) string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		reporter.ExitWithReport("kpm_e2e: failed to load file, the dir not exists.")
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ext {
			data, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				reporter.ExitWithReport("kpm_e2e: the file exists, but failed to read file.")
			}
			return string(data)
		}
	}

	reporter.ExitWithReport("kpm_e2e: failed to load file, the file not exists.")
	return ""
}

// Copy will copy file from 'srcPath' to 'dstPath'.
func Copy(srcPath, dstPath string) {
	src, err := os.Open(srcPath)
	if err != nil {
		reporter.ExitWithReport("kpm_e2e: failed to copy file from src.")
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		reporter.ExitWithReport("kpm_e2e: failed to copy file to dst.")
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		reporter.ExitWithReport("kpm_e2e: failed to copy file.")
	}
}

// CopyDir will copy dir from 'srcDir' to 'dstDir'.
func CopyDir(srcDir, dstDir string) {
	err := copy.Copy(srcDir, dstDir)
	if err != nil {
		reporter.ExitWithReport("kpm_e2e: failed to copy dir.")
	}
}

var KEYS = []string{"<workspace>", "<ignore>"}

// IsIgnore will reture whether the expected result in 'expectedStr' should be ignored.
func IsIgnore(expectedStr string) bool {
	return strings.Contains(expectedStr, "<ignore>")
}

// ReplaceAllKeyByValue will replace all 'key's by 'value' in 'originStr'.
func ReplaceAllKeyByValue(originStr, key, value string) string {
	if !funk.Contains(KEYS, key) {
		reporter.ExitWithReport("kpm_e2e: unknown key.", key)
	} else {
		return strings.ReplaceAll(originStr, key, value)
	}

	return originStr
}