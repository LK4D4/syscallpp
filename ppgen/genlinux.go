package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const numPrefix = "#define __NR_"

var archNumFiles = map[string]string{
	"386":   "arch/x86/include/generated/uapi/asm/unistd_32.h",
	"amd64": "arch/x86/include/generated/uapi/asm/unistd_64.h",
	"arm":   "arch/arm/include/uapi/asm/unistd.h",
	"arm64": "arch/arm64/include/uapi/asm/unistd.h",
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

func parseLinuxNumbers(arch string) ([]sc, error) {
	fName, ok := archNumFiles[arch]
	if !ok {
		return nil, fmt.Errorf("Unsupported arch: %s", arch)
	}
	fName = filepath.Join(srcPath, fName)
	f, err := os.Open(fName)
	if err != nil {
		return nil, err
	}
	var scs []sc
	s := bufio.NewScanner(f)
	for s.Scan() {
		ln := s.Text()
		if !strings.HasPrefix(ln, numPrefix) {
			continue
		}
		switch arch {
		case "386", "amd64", "ppc64":
			parsedSc, err := parseSimpleNumber(ln[len(numPrefix):])
			if err != nil {
				return nil, err
			}
			scs = append(scs, parsedSc)
		}
	}
	if s.Err() != nil {
		return nil, s.Err()
	}
	return scs, nil
}
