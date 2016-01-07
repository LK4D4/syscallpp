package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var archNumFiles = map[string]string{
	"386":   "arch/x86/include/generated/uapi/asm/unistd_32.h",
	"amd64": "arch/x86/include/generated/uapi/asm/unistd_64.h",
	"arm":   "arch/arm/include/uapi/asm/unistd.h",
	"arm64": "include/uapi/asm-generic/unistd.h",
	"ppc64": "arch/powerpc/include/uapi/asm/unistd.h",
}
var srcPath string

func init() {
	flag.StringVar(&srcPath, "linux-src", "/usr/src/linux", "path to linux sources")
}

func parseSimpleNumber(ln string) (sc, error) {
	parts := strings.Fields(ln)
	num, err := strconv.Atoi(parts[1])
	if err != nil {
		return sc{}, err
	}
	return sc{
		name:   parts[0],
		number: num,
	}, nil
}

func parseSimpleArchNumbers(s *bufio.Scanner) ([]sc, error) {
	const numPrefix = "#define __NR_"

	var scs []sc
	for s.Scan() {
		ln := s.Text()
		if !strings.HasPrefix(ln, numPrefix) {
			continue
		}
		parsedSc, err := parseSimpleNumber(ln[len(numPrefix):])
		if err != nil {
			return nil, err
		}
		scs = append(scs, parsedSc)
	}
	if s.Err() != nil {
		return nil, s.Err()
	}
	return scs, nil
}

func parseArm64Numbers(s *bufio.Scanner) ([]sc, error) {
	const numPrefix = "#define __NR_"
	const sc3264Prefix = "__NR3264_"
	const num3264Prefix = "#define " + sc3264Prefix
	cache3264 := map[string]int{}
	mapScTo3264 := map[string]string{}
	var scs []sc
	for s.Scan() {
		ln := s.Text()
		if strings.HasPrefix(ln, numPrefix+"syscalls") {
			continue
		}
		if strings.HasPrefix(ln, numPrefix) {
			ln = ln[len(numPrefix):]
			parts := strings.Fields(ln)
			if strings.HasPrefix(parts[1], sc3264Prefix) {
				mapScTo3264[parts[0]] = parts[1][len(sc3264Prefix):]
				continue
			}
			num, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, err
			}
			scs = append(scs, sc{name: parts[0], number: num})
		}
		if strings.HasPrefix(ln, num3264Prefix) {
			ln = ln[len(num3264Prefix):]
			parts := strings.Fields(ln)
			num, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, err
			}
			cache3264[parts[0]] = num
		}
	}
	for real, sc3264 := range mapScTo3264 {
		scs = append(scs, sc{name: real, number: cache3264[sc3264]})
	}
	if s.Err() != nil {
		return nil, s.Err()
	}
	return scs, nil
}

func parseArm32Numbers(s *bufio.Scanner) ([]sc, error) {
	const numPrefix = "#define __NR_"

	var scs []sc
	for s.Scan() {
		ln := s.Text()
		if !strings.HasPrefix(ln, numPrefix) {
			continue
		}
		ln = ln[len(numPrefix):]
		parts := strings.SplitN(ln, " ", 2)
		if strings.HasSuffix(parts[0], "SYSCALL_BASE") {
			continue
		}
		if parts[0] == "sync_file_range2" {
			for _, psc := range scs {
				if psc.name == "arm_sync_file_range" {
					scs = append(scs, sc{name: "sync_file_range2", number: psc.number})
				}
			}
			continue
		}
		strNum := parts[1][len("(__NR_SYSCALL_BASE+"):]
		strNum = strings.Trim(strNum, " )")
		num, err := strconv.Atoi(strNum)
		if err != nil {
			return nil, err
		}
		scs = append(scs, sc{name: parts[0], number: num})
	}
	if s.Err() != nil {
		return nil, s.Err()
	}
	return scs, nil
}

func parseLinuxNumbers(arch string) ([]sc, error) {
	fName, ok := archNumFiles[arch]
	if !ok {
		return nil, fmt.Errorf("Unsupported arch: %s", arch)
	}
	fName = filepath.Join(srcPath, fName)
	cmd := exec.Command("gcc", "-E", "-dD", fName)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	s := bufio.NewScanner(&stdout)
	switch arch {
	case "386", "amd64", "ppc64":
		return parseSimpleArchNumbers(s)
	case "arm64":
		return parseArm64Numbers(s)
	case "arm":
		return parseArm32Numbers(s)
	default:
		return nil, fmt.Errorf("Unsupported arch: %s", arch)
	}
}
