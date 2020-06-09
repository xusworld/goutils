package goutils

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestListDir(t *testing.T) {
	filenames := ListDir("./")
	fmt.Println(filenames)
}

// 测试 bufio.Scanner的坑
func TestReadLines(t *testing.T) {
	veryLongString := RandomStringFromCharset(70000) // MaxScanTokenSize = 64 * 1024 = 65536 < 70000

	scanner := bufio.NewScanner(strings.NewReader(veryLongString))
	for scanner.Scan() {
		t.Log(scanner.Text())
	}

	if scanner.Err() == bufio.ErrTooLong  {
		t.Errorf("bufio.Scan() : %s", bufio.ErrTooLong)
	}
	
	// solution
	newScanner := bufio.NewScanner(strings.NewReader(veryLongString))
	newScanner.Buffer([]byte{}, bufio.MaxScanTokenSize * 2)
	for newScanner.Scan() {
		t.Logf("bufio.Scan() success")
	}
}

func TestReadLinesV2(t *testing.T) {
	lines,_ , _:= ReadLinesV2("./tmp/name")
	t.Logf("%v", lines)
}

//