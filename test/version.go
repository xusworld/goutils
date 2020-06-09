package test

import (
	"strconv"
	"strings"
)

// CompareVersion 比较字符串版本号
func CompareVersion(lhs, rhs string) int {
	leftVersions := strings.Split(strings.Trim(lhs, " "), ".")
	rightVersions := strings.Split(strings.Trim(rhs, " "), ".")

	var numsToCompare int

	leftStrLen := len(leftVersions)
	rightStrLen := len(rightVersions)

	if leftStrLen > rightStrLen {
		numsToCompare = rightStrLen
	} else {
		numsToCompare = leftStrLen
	}

	for idx := 0; idx < numsToCompare; idx++ {
		left, _ := strconv.Atoi(leftVersions[idx])
		right, _ := strconv.Atoi(rightVersions[idx])

		if left > right {
			return 1 // bigger than
		} else if left < right {
			return -1 // less than
		}
	}

	if leftStrLen > numsToCompare {
		return 1 // bigger than
	}
	if rightStrLen > numsToCompare {
		return -1 // less than
	}
	return 0 // equal to
}

// CompareVersionGT 判断 LeftVersion 是否 大于 RightVersion, GT = greater than
func CompareVersionGT(lhs, rhs string) bool {
	return CompareVersion(lhs, rhs) > 0
}

// CompareVersionGE 判断 LeftVersion 是否 大于 RightVersion, GE = greater equal
func CompareVersionGE(lhs, rhs string) bool {
	return CompareVersion(lhs, rhs) >= 0
}

// CompareVersionLT 判断 LeftVersion 是否 小于 RightVersion, LT = less than
func CompareVersionLT(lhs, rhs string) bool {
	return CompareVersion(lhs, rhs) < 0
}

// CompareVersionLE 判断 LeftVersion 是否 小于等于 RightVersion, LE = less equal
func CompareVersionLE(lhs, rhs string) bool {
	return CompareVersion(lhs, rhs) <= 0
}
