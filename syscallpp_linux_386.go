// generated file; DO NOT EDIT - use go generate in directory with source

// +build 386,linux

package syscallpp

func GetName(num int) string {
	switch num {
	case 1:
		return "exit"
	case 3:
		return "read"
	case 4:
		return "write"
	case 6:
		return "close"
	case 12:
		return "chdir"
	case 13:
		return "time"
	case 16:
		return "lchown"
	case 20:
		return "getpid"
	case 21:
		return "mount"
	case 24:
		return "getuid"
	case 26:
		return "ptrace"
	case 29:
		return "pause"
	case 30:
		return "utime"
	case 36:
		return "sync"
	case 37:
		return "kill"
	case 41:
		return "dup"
	case 42:
		return "pipe"
	case 43:
		return "times"
	case 47:
		return "getgid"
	case 49:
		return "geteuid"
	case 50:
		return "getegid"
	case 51:
		return "acct"
	case 55:
		return "fcntl"
	case 57:
		return "setpgid"
	case 60:
		return "umask"
	case 61:
		return "chroot"
	case 62:
		return "ustat"
	case 63:
		return "dup2"
	case 64:
		return "getppid"
	case 66:
		return "setsid"
	case 70:
		return "setreuid"
	case 71:
		return "setregid"
	case 74:
		return "sethostname"
	case 75:
		return "setrlimit"
	case 76:
		return "getrlimit"
	case 77:
		return "getrusage"
	case 78:
		return "gettimeofday"
	case 79:
		return "settimeofday"
	case 80:
		return "getgroups"
	case 81:
		return "setgroups"
	case 82:
		return "select"
	case 88:
		return "reboot"
	case 91:
		return "munmap"
	case 92:
		return "truncate"
	case 93:
		return "ftruncate"
	case 94:
		return "fchmod"
	case 95:
		return "fchown"
	case 96:
		return "getpriority"
	case 97:
		return "setpriority"
	case 101:
		return "ioperm"
	case 106:
		return "stat"
	case 107:
		return "lstat"
	case 108:
		return "fstat"
	case 110:
		return "iopl"
	case 114:
		return "wait4"
	case 116:
		return "sysinfo"
	case 118:
		return "fsync"
	case 121:
		return "setdomainname"
	case 122:
		return "uname"
	case 124:
		return "adjtimex"
	case 125:
		return "mprotect"
	case 132:
		return "getpgid"
	case 133:
		return "fchdir"
	case 138:
		return "setfsuid"
	case 139:
		return "setfsgid"
	case 141:
		return "getdents"
	case 143:
		return "flock"
	case 148:
		return "fdatasync"
	case 150:
		return "mlock"
	case 151:
		return "munlock"
	case 152:
		return "mlockall"
	case 153:
		return "munlockall"
	case 162:
		return "nanosleep"
	case 164:
		return "setresuid"
	case 170:
		return "setresgid"
	case 183:
		return "getcwd"
	case 187:
		return "sendfile"
	case 192:
		return "mmap2"
	case 219:
		return "madvise"
	case 223:
		return "gettid"
	case 225:
		return "setxattr"
	case 228:
		return "getxattr"
	case 231:
		return "listxattr"
	case 234:
		return "removexattr"
	case 268:
		return "tgkill"
	case 269:
		return "utimes"
	case 292:
		return "openat"
	case 293:
		return "mkdirat"
	case 294:
		return "mknodat"
	case 295:
		return "fchownat"
	case 296:
		return "futimesat"
	case 298:
		return "unlinkat"
	case 299:
		return "renameat"
	case 300:
		return "linkat"
	case 301:
		return "symlinkat"
	case 302:
		return "readlinkat"
	case 303:
		return "fchmodat"
	case 304:
		return "faccessat"
	case 307:
		return "unshare"
	case 310:
		return "splice"
	case 312:
		return "tee"
	case 317:
		return "utimensat"
	case 321:
		return "fallocate"
	case 327:
		return "dup3"
	case 328:
		return "pipe2"
	}
	return "unknown"
}

func GetNum(name string) int {
	switch name {
	case "exit":
		return 1
	case "read":
		return 3
	case "write":
		return 4
	case "close":
		return 6
	case "chdir":
		return 12
	case "time":
		return 13
	case "lchown":
		return 16
	case "getpid":
		return 20
	case "mount":
		return 21
	case "getuid":
		return 24
	case "ptrace":
		return 26
	case "pause":
		return 29
	case "utime":
		return 30
	case "sync":
		return 36
	case "kill":
		return 37
	case "dup":
		return 41
	case "pipe":
		return 42
	case "times":
		return 43
	case "getgid":
		return 47
	case "geteuid":
		return 49
	case "getegid":
		return 50
	case "acct":
		return 51
	case "fcntl":
		return 55
	case "setpgid":
		return 57
	case "umask":
		return 60
	case "chroot":
		return 61
	case "ustat":
		return 62
	case "dup2":
		return 63
	case "getppid":
		return 64
	case "setsid":
		return 66
	case "setreuid":
		return 70
	case "setregid":
		return 71
	case "sethostname":
		return 74
	case "setrlimit":
		return 75
	case "getrlimit":
		return 76
	case "getrusage":
		return 77
	case "gettimeofday":
		return 78
	case "settimeofday":
		return 79
	case "getgroups":
		return 80
	case "setgroups":
		return 81
	case "select":
		return 82
	case "reboot":
		return 88
	case "munmap":
		return 91
	case "truncate":
		return 92
	case "ftruncate":
		return 93
	case "fchmod":
		return 94
	case "fchown":
		return 95
	case "getpriority":
		return 96
	case "setpriority":
		return 97
	case "ioperm":
		return 101
	case "stat":
		return 106
	case "lstat":
		return 107
	case "fstat":
		return 108
	case "iopl":
		return 110
	case "wait4":
		return 114
	case "sysinfo":
		return 116
	case "fsync":
		return 118
	case "setdomainname":
		return 121
	case "uname":
		return 122
	case "adjtimex":
		return 124
	case "mprotect":
		return 125
	case "getpgid":
		return 132
	case "fchdir":
		return 133
	case "setfsuid":
		return 138
	case "setfsgid":
		return 139
	case "getdents":
		return 141
	case "flock":
		return 143
	case "fdatasync":
		return 148
	case "mlock":
		return 150
	case "munlock":
		return 151
	case "mlockall":
		return 152
	case "munlockall":
		return 153
	case "nanosleep":
		return 162
	case "setresuid":
		return 164
	case "setresgid":
		return 170
	case "getcwd":
		return 183
	case "sendfile":
		return 187
	case "mmap2":
		return 192
	case "madvise":
		return 219
	case "gettid":
		return 223
	case "setxattr":
		return 225
	case "getxattr":
		return 228
	case "listxattr":
		return 231
	case "removexattr":
		return 234
	case "tgkill":
		return 268
	case "utimes":
		return 269
	case "openat":
		return 292
	case "mkdirat":
		return 293
	case "mknodat":
		return 294
	case "fchownat":
		return 295
	case "futimesat":
		return 296
	case "unlinkat":
		return 298
	case "renameat":
		return 299
	case "linkat":
		return 300
	case "symlinkat":
		return 301
	case "readlinkat":
		return 302
	case "fchmodat":
		return 303
	case "faccessat":
		return 304
	case "unshare":
		return 307
	case "splice":
		return 310
	case "tee":
		return 312
	case "utimensat":
		return 317
	case "fallocate":
		return 321
	case "dup3":
		return 327
	case "pipe2":
		return 328
	}
	return -1
}

func GetArgsTypes(name string) []ArgType {
	switch name {
	case "exit":
		return []ArgType{ARG_INT}
	case "read":
		return []ArgType{ARG_INT, ARG_PTR}
	case "write":
		return []ArgType{ARG_INT, ARG_PTR}
	case "close":
		return []ArgType{ARG_INT}
	case "chdir":
		return []ArgType{ARG_STR}
	case "time":
		return []ArgType{ARG_PTR}
	case "lchown":
		return []ArgType{ARG_STR, ARG_INT, ARG_INT}
	case "getpid":
		return []ArgType(nil)
	case "mount":
		return []ArgType{ARG_STR, ARG_STR, ARG_STR, ARG_INT, ARG_PTR}
	case "getuid":
		return []ArgType(nil)
	case "ptrace":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT}
	case "pause":
		return []ArgType(nil)
	case "utime":
		return []ArgType{ARG_STR, ARG_PTR}
	case "sync":
		return []ArgType(nil)
	case "kill":
		return []ArgType{ARG_INT, ARG_PTR}
	case "dup":
		return []ArgType{ARG_INT}
	case "pipe":
		return []ArgType{ARG_PTR}
	case "times":
		return []ArgType{ARG_PTR}
	case "getgid":
		return []ArgType(nil)
	case "geteuid":
		return []ArgType(nil)
	case "getegid":
		return []ArgType(nil)
	case "acct":
		return []ArgType{ARG_STR}
	case "fcntl":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "setpgid":
		return []ArgType{ARG_INT, ARG_INT}
	case "umask":
		return []ArgType{ARG_INT}
	case "chroot":
		return []ArgType{ARG_STR}
	case "ustat":
		return []ArgType{ARG_INT, ARG_PTR}
	case "dup2":
		return []ArgType{ARG_INT, ARG_INT}
	case "getppid":
		return []ArgType(nil)
	case "setsid":
		return []ArgType(nil)
	case "setreuid":
		return []ArgType{ARG_INT, ARG_INT}
	case "setregid":
		return []ArgType{ARG_INT, ARG_INT}
	case "sethostname":
		return []ArgType{ARG_PTR}
	case "setrlimit":
		return []ArgType{ARG_INT, ARG_PTR}
	case "getrlimit":
		return []ArgType{ARG_INT, ARG_PTR}
	case "getrusage":
		return []ArgType{ARG_INT, ARG_PTR}
	case "gettimeofday":
		return []ArgType{ARG_PTR}
	case "settimeofday":
		return []ArgType{ARG_PTR}
	case "getgroups":
		return []ArgType{ARG_INT, ARG_PTR}
	case "setgroups":
		return []ArgType{ARG_INT, ARG_PTR}
	case "select":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}
	case "reboot":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_STR}
	case "munmap":
		return []ArgType{ARG_INT, ARG_INT}
	case "truncate":
		return []ArgType{ARG_STR, ARG_PTR}
	case "ftruncate":
		return []ArgType{ARG_INT, ARG_PTR}
	case "fchmod":
		return []ArgType{ARG_INT, ARG_INT}
	case "fchown":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "getpriority":
		return []ArgType{ARG_INT, ARG_INT}
	case "setpriority":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "ioperm":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "stat":
		return []ArgType{ARG_STR, ARG_PTR}
	case "lstat":
		return []ArgType{ARG_STR, ARG_PTR}
	case "fstat":
		return []ArgType{ARG_INT, ARG_PTR}
	case "iopl":
		return []ArgType{ARG_INT}
	case "wait4":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR}
	case "sysinfo":
		return []ArgType{ARG_PTR}
	case "fsync":
		return []ArgType{ARG_INT}
	case "setdomainname":
		return []ArgType{ARG_PTR}
	case "uname":
		return []ArgType{ARG_PTR}
	case "adjtimex":
		return []ArgType{ARG_PTR}
	case "mprotect":
		return []ArgType{ARG_PTR, ARG_INT}
	case "getpgid":
		return []ArgType{ARG_INT}
	case "fchdir":
		return []ArgType{ARG_INT}
	case "setfsuid":
		return []ArgType{ARG_INT}
	case "setfsgid":
		return []ArgType{ARG_INT}
	case "getdents":
		return []ArgType{ARG_INT, ARG_PTR}
	case "flock":
		return []ArgType{ARG_INT, ARG_INT}
	case "fdatasync":
		return []ArgType{ARG_INT}
	case "mlock":
		return []ArgType{ARG_PTR}
	case "munlock":
		return []ArgType{ARG_PTR}
	case "mlockall":
		return []ArgType{ARG_INT}
	case "munlockall":
		return []ArgType(nil)
	case "nanosleep":
		return []ArgType{ARG_PTR, ARG_PTR}
	case "setresuid":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "setresgid":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "getcwd":
		return []ArgType{ARG_PTR}
	case "sendfile":
		return []ArgType{ARG_INT, ARG_INT, ARG_PTR, ARG_INT}
	case "mmap2":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_INT}
	case "madvise":
		return []ArgType{ARG_PTR, ARG_INT}
	case "gettid":
		return []ArgType(nil)
	case "setxattr":
		return []ArgType{ARG_STR, ARG_STR, ARG_PTR, ARG_INT}
	case "getxattr":
		return []ArgType{ARG_STR, ARG_STR, ARG_PTR}
	case "listxattr":
		return []ArgType{ARG_STR, ARG_PTR}
	case "removexattr":
		return []ArgType{ARG_STR, ARG_STR}
	case "tgkill":
		return []ArgType{ARG_INT, ARG_INT, ARG_PTR}
	case "utimes":
		return []ArgType{ARG_STR, ARG_PTR}
	case "openat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_INT}
	case "mkdirat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT}
	case "mknodat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_INT}
	case "fchownat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_INT, ARG_INT}
	case "futimesat":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "unlinkat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT}
	case "renameat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_STR}
	case "linkat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_STR, ARG_INT}
	case "symlinkat":
		return []ArgType{ARG_STR, ARG_INT, ARG_STR}
	case "readlinkat":
		return []ArgType{ARG_INT, ARG_STR, ARG_PTR}
	case "fchmodat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_INT}
	case "faccessat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_INT}
	case "unshare":
		return []ArgType{ARG_INT}
	case "splice":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_INT, ARG_INT}
	case "tee":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT}
	case "utimensat":
		return []ArgType{ARG_INT, ARG_STR, ARG_PTR}
	case "fallocate":
		return []ArgType{ARG_INT, ARG_INT, ARG_PTR, ARG_PTR}
	case "dup3":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "pipe2":
		return []ArgType{ARG_PTR, ARG_INT}
	}
	return nil
}
