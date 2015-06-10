// generated file; DO NOT EDIT - use go generate in directory with source

// +build amd64,linux

package syscallpp

func GetName(num int) string {
	switch num {
	case 0:
		return "read"
	case 1:
		return "write"
	case 3:
		return "close"
	case 4:
		return "stat"
	case 5:
		return "fstat"
	case 6:
		return "lstat"
	case 9:
		return "mmap"
	case 10:
		return "mprotect"
	case 11:
		return "munmap"
	case 22:
		return "pipe"
	case 23:
		return "select"
	case 28:
		return "madvise"
	case 32:
		return "dup"
	case 33:
		return "dup2"
	case 34:
		return "pause"
	case 35:
		return "nanosleep"
	case 39:
		return "getpid"
	case 40:
		return "sendfile"
	case 41:
		return "socket"
	case 42:
		return "connect"
	case 43:
		return "accept"
	case 44:
		return "sendto"
	case 45:
		return "recvfrom"
	case 46:
		return "sendmsg"
	case 47:
		return "recvmsg"
	case 48:
		return "shutdown"
	case 49:
		return "bind"
	case 50:
		return "listen"
	case 51:
		return "getsockname"
	case 52:
		return "getpeername"
	case 53:
		return "socketpair"
	case 54:
		return "setsockopt"
	case 55:
		return "getsockopt"
	case 60:
		return "exit"
	case 61:
		return "wait4"
	case 62:
		return "kill"
	case 63:
		return "uname"
	case 72:
		return "fcntl"
	case 73:
		return "flock"
	case 74:
		return "fsync"
	case 75:
		return "fdatasync"
	case 76:
		return "truncate"
	case 77:
		return "ftruncate"
	case 78:
		return "getdents"
	case 79:
		return "getcwd"
	case 80:
		return "chdir"
	case 81:
		return "fchdir"
	case 91:
		return "fchmod"
	case 93:
		return "fchown"
	case 94:
		return "lchown"
	case 95:
		return "umask"
	case 97:
		return "getrlimit"
	case 98:
		return "getrusage"
	case 99:
		return "sysinfo"
	case 100:
		return "times"
	case 101:
		return "ptrace"
	case 102:
		return "getuid"
	case 104:
		return "getgid"
	case 107:
		return "geteuid"
	case 108:
		return "getegid"
	case 109:
		return "setpgid"
	case 110:
		return "getppid"
	case 112:
		return "setsid"
	case 113:
		return "setreuid"
	case 114:
		return "setregid"
	case 115:
		return "getgroups"
	case 116:
		return "setgroups"
	case 117:
		return "setresuid"
	case 119:
		return "setresgid"
	case 121:
		return "getpgid"
	case 122:
		return "setfsuid"
	case 123:
		return "setfsgid"
	case 132:
		return "utime"
	case 136:
		return "ustat"
	case 137:
		return "statfs"
	case 138:
		return "fstatfs"
	case 140:
		return "getpriority"
	case 141:
		return "setpriority"
	case 149:
		return "mlock"
	case 150:
		return "munlock"
	case 151:
		return "mlockall"
	case 152:
		return "munlockall"
	case 159:
		return "adjtimex"
	case 160:
		return "setrlimit"
	case 161:
		return "chroot"
	case 162:
		return "sync"
	case 163:
		return "acct"
	case 164:
		return "settimeofday"
	case 165:
		return "mount"
	case 169:
		return "reboot"
	case 170:
		return "sethostname"
	case 171:
		return "setdomainname"
	case 172:
		return "iopl"
	case 173:
		return "ioperm"
	case 186:
		return "gettid"
	case 188:
		return "setxattr"
	case 191:
		return "getxattr"
	case 194:
		return "listxattr"
	case 197:
		return "removexattr"
	case 234:
		return "tgkill"
	case 235:
		return "utimes"
	case 257:
		return "openat"
	case 258:
		return "mkdirat"
	case 259:
		return "mknodat"
	case 260:
		return "fchownat"
	case 261:
		return "futimesat"
	case 263:
		return "unlinkat"
	case 264:
		return "renameat"
	case 265:
		return "linkat"
	case 266:
		return "symlinkat"
	case 267:
		return "readlinkat"
	case 268:
		return "fchmodat"
	case 269:
		return "faccessat"
	case 272:
		return "unshare"
	case 275:
		return "splice"
	case 276:
		return "tee"
	case 280:
		return "utimensat"
	case 285:
		return "fallocate"
	case 288:
		return "accept4"
	case 292:
		return "dup3"
	case 293:
		return "pipe2"
	}
	return "unknown"
}

func GetNum(name string) int {
	switch name {
	case "read":
		return 0
	case "write":
		return 1
	case "close":
		return 3
	case "stat":
		return 4
	case "fstat":
		return 5
	case "lstat":
		return 6
	case "mmap":
		return 9
	case "mprotect":
		return 10
	case "munmap":
		return 11
	case "pipe":
		return 22
	case "select":
		return 23
	case "madvise":
		return 28
	case "dup":
		return 32
	case "dup2":
		return 33
	case "pause":
		return 34
	case "nanosleep":
		return 35
	case "getpid":
		return 39
	case "sendfile":
		return 40
	case "socket":
		return 41
	case "connect":
		return 42
	case "accept":
		return 43
	case "sendto":
		return 44
	case "recvfrom":
		return 45
	case "sendmsg":
		return 46
	case "recvmsg":
		return 47
	case "shutdown":
		return 48
	case "bind":
		return 49
	case "listen":
		return 50
	case "getsockname":
		return 51
	case "getpeername":
		return 52
	case "socketpair":
		return 53
	case "setsockopt":
		return 54
	case "getsockopt":
		return 55
	case "exit":
		return 60
	case "wait4":
		return 61
	case "kill":
		return 62
	case "uname":
		return 63
	case "fcntl":
		return 72
	case "flock":
		return 73
	case "fsync":
		return 74
	case "fdatasync":
		return 75
	case "truncate":
		return 76
	case "ftruncate":
		return 77
	case "getdents":
		return 78
	case "getcwd":
		return 79
	case "chdir":
		return 80
	case "fchdir":
		return 81
	case "fchmod":
		return 91
	case "fchown":
		return 93
	case "lchown":
		return 94
	case "umask":
		return 95
	case "getrlimit":
		return 97
	case "getrusage":
		return 98
	case "sysinfo":
		return 99
	case "times":
		return 100
	case "ptrace":
		return 101
	case "getuid":
		return 102
	case "getgid":
		return 104
	case "geteuid":
		return 107
	case "getegid":
		return 108
	case "setpgid":
		return 109
	case "getppid":
		return 110
	case "setsid":
		return 112
	case "setreuid":
		return 113
	case "setregid":
		return 114
	case "getgroups":
		return 115
	case "setgroups":
		return 116
	case "setresuid":
		return 117
	case "setresgid":
		return 119
	case "getpgid":
		return 121
	case "setfsuid":
		return 122
	case "setfsgid":
		return 123
	case "utime":
		return 132
	case "ustat":
		return 136
	case "statfs":
		return 137
	case "fstatfs":
		return 138
	case "getpriority":
		return 140
	case "setpriority":
		return 141
	case "mlock":
		return 149
	case "munlock":
		return 150
	case "mlockall":
		return 151
	case "munlockall":
		return 152
	case "adjtimex":
		return 159
	case "setrlimit":
		return 160
	case "chroot":
		return 161
	case "sync":
		return 162
	case "acct":
		return 163
	case "settimeofday":
		return 164
	case "mount":
		return 165
	case "reboot":
		return 169
	case "sethostname":
		return 170
	case "setdomainname":
		return 171
	case "iopl":
		return 172
	case "ioperm":
		return 173
	case "gettid":
		return 186
	case "setxattr":
		return 188
	case "getxattr":
		return 191
	case "listxattr":
		return 194
	case "removexattr":
		return 197
	case "tgkill":
		return 234
	case "utimes":
		return 235
	case "openat":
		return 257
	case "mkdirat":
		return 258
	case "mknodat":
		return 259
	case "fchownat":
		return 260
	case "futimesat":
		return 261
	case "unlinkat":
		return 263
	case "renameat":
		return 264
	case "linkat":
		return 265
	case "symlinkat":
		return 266
	case "readlinkat":
		return 267
	case "fchmodat":
		return 268
	case "faccessat":
		return 269
	case "unshare":
		return 272
	case "splice":
		return 275
	case "tee":
		return 276
	case "utimensat":
		return 280
	case "fallocate":
		return 285
	case "accept4":
		return 288
	case "dup3":
		return 292
	case "pipe2":
		return 293
	}
	return -1
}

func GetArgsTypes(name string) []ArgType {
	switch name {
	case "read":
		return []ArgType{ARG_INT, ARG_PTR}
	case "write":
		return []ArgType{ARG_INT, ARG_PTR}
	case "close":
		return []ArgType{ARG_INT}
	case "stat":
		return []ArgType{ARG_STR, ARG_PTR}
	case "fstat":
		return []ArgType{ARG_INT, ARG_PTR}
	case "lstat":
		return []ArgType{ARG_STR, ARG_PTR}
	case "mmap":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_PTR}
	case "mprotect":
		return []ArgType{ARG_PTR, ARG_INT}
	case "munmap":
		return []ArgType{ARG_INT, ARG_INT}
	case "pipe":
		return []ArgType{ARG_PTR}
	case "select":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}
	case "madvise":
		return []ArgType{ARG_PTR, ARG_INT}
	case "dup":
		return []ArgType{ARG_INT}
	case "dup2":
		return []ArgType{ARG_INT, ARG_INT}
	case "pause":
		return []ArgType(nil)
	case "nanosleep":
		return []ArgType{ARG_PTR, ARG_PTR}
	case "getpid":
		return []ArgType(nil)
	case "sendfile":
		return []ArgType{ARG_INT, ARG_INT, ARG_PTR, ARG_INT}
	case "socket":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "connect":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "accept":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "sendto":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_PTR}
	case "recvfrom":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_PTR}
	case "sendmsg":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT}
	case "recvmsg":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT}
	case "shutdown":
		return []ArgType{ARG_INT, ARG_INT}
	case "bind":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "listen":
		return []ArgType{ARG_INT, ARG_INT}
	case "getsockname":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "getpeername":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "socketpair":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_PTR}
	case "setsockopt":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_PTR, ARG_INT}
	case "getsockopt":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_PTR, ARG_PTR}
	case "exit":
		return []ArgType{ARG_INT}
	case "wait4":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR}
	case "kill":
		return []ArgType{ARG_INT, ARG_PTR}
	case "uname":
		return []ArgType{ARG_PTR}
	case "fcntl":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "flock":
		return []ArgType{ARG_INT, ARG_INT}
	case "fsync":
		return []ArgType{ARG_INT}
	case "fdatasync":
		return []ArgType{ARG_INT}
	case "truncate":
		return []ArgType{ARG_STR, ARG_PTR}
	case "ftruncate":
		return []ArgType{ARG_INT, ARG_PTR}
	case "getdents":
		return []ArgType{ARG_INT, ARG_PTR}
	case "getcwd":
		return []ArgType{ARG_PTR}
	case "chdir":
		return []ArgType{ARG_STR}
	case "fchdir":
		return []ArgType{ARG_INT}
	case "fchmod":
		return []ArgType{ARG_INT, ARG_INT}
	case "fchown":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "lchown":
		return []ArgType{ARG_STR, ARG_INT, ARG_INT}
	case "umask":
		return []ArgType{ARG_INT}
	case "getrlimit":
		return []ArgType{ARG_INT, ARG_PTR}
	case "getrusage":
		return []ArgType{ARG_INT, ARG_PTR}
	case "sysinfo":
		return []ArgType{ARG_PTR}
	case "times":
		return []ArgType{ARG_PTR}
	case "ptrace":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT}
	case "getuid":
		return []ArgType(nil)
	case "getgid":
		return []ArgType(nil)
	case "geteuid":
		return []ArgType(nil)
	case "getegid":
		return []ArgType(nil)
	case "setpgid":
		return []ArgType{ARG_INT, ARG_INT}
	case "getppid":
		return []ArgType(nil)
	case "setsid":
		return []ArgType(nil)
	case "setreuid":
		return []ArgType{ARG_INT, ARG_INT}
	case "setregid":
		return []ArgType{ARG_INT, ARG_INT}
	case "getgroups":
		return []ArgType{ARG_INT, ARG_PTR}
	case "setgroups":
		return []ArgType{ARG_INT, ARG_PTR}
	case "setresuid":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "setresgid":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "getpgid":
		return []ArgType{ARG_INT}
	case "setfsuid":
		return []ArgType{ARG_INT}
	case "setfsgid":
		return []ArgType{ARG_INT}
	case "utime":
		return []ArgType{ARG_STR, ARG_PTR}
	case "ustat":
		return []ArgType{ARG_INT, ARG_PTR}
	case "statfs":
		return []ArgType{ARG_STR, ARG_PTR}
	case "fstatfs":
		return []ArgType{ARG_INT, ARG_PTR}
	case "getpriority":
		return []ArgType{ARG_INT, ARG_INT}
	case "setpriority":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "mlock":
		return []ArgType{ARG_PTR}
	case "munlock":
		return []ArgType{ARG_PTR}
	case "mlockall":
		return []ArgType{ARG_INT}
	case "munlockall":
		return []ArgType(nil)
	case "adjtimex":
		return []ArgType{ARG_PTR}
	case "setrlimit":
		return []ArgType{ARG_INT, ARG_PTR}
	case "chroot":
		return []ArgType{ARG_STR}
	case "sync":
		return []ArgType(nil)
	case "acct":
		return []ArgType{ARG_STR}
	case "settimeofday":
		return []ArgType{ARG_PTR}
	case "mount":
		return []ArgType{ARG_STR, ARG_STR, ARG_STR, ARG_INT, ARG_PTR}
	case "reboot":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_STR}
	case "sethostname":
		return []ArgType{ARG_PTR}
	case "setdomainname":
		return []ArgType{ARG_PTR}
	case "iopl":
		return []ArgType{ARG_INT}
	case "ioperm":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
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
	case "accept4":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR, ARG_INT}
	case "dup3":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "pipe2":
		return []ArgType{ARG_PTR, ARG_INT}
	}
	return nil
}
