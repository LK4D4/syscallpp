// generated file; DO NOT EDIT - use go generate in directory with source

// +build ppc64le,linux

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
	case 90:
		return "mmap"
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
	case 99:
		return "statfs"
	case 100:
		return "fstatfs"
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
	case 169:
		return "setresgid"
	case 182:
		return "getcwd"
	case 186:
		return "sendfile"
	case 198:
		return "madvise"
	case 200:
		return "gettid"
	case 202:
		return "setxattr"
	case 205:
		return "getxattr"
	case 208:
		return "listxattr"
	case 211:
		return "removexattr"
	case 241:
		return "tgkill"
	case 242:
		return "utimes"
	case 271:
		return "unshare"
	case 272:
		return "splice"
	case 273:
		return "tee"
	case 275:
		return "openat"
	case 276:
		return "mkdirat"
	case 277:
		return "mknodat"
	case 278:
		return "fchownat"
	case 279:
		return "futimesat"
	case 281:
		return "unlinkat"
	case 282:
		return "renameat"
	case 283:
		return "linkat"
	case 284:
		return "symlinkat"
	case 285:
		return "readlinkat"
	case 286:
		return "fchmodat"
	case 287:
		return "faccessat"
	case 293:
		return "utimensat"
	case 298:
		return "fallocate"
	case 305:
		return "dup3"
	case 306:
		return "pipe2"
	case 315:
		return "socket"
	case 316:
		return "bind"
	case 317:
		return "connect"
	case 318:
		return "listen"
	case 319:
		return "accept"
	case 320:
		return "getsockname"
	case 321:
		return "getpeername"
	case 322:
		return "socketpair"
	case 324:
		return "sendto"
	case 326:
		return "recvfrom"
	case 327:
		return "shutdown"
	case 328:
		return "setsockopt"
	case 329:
		return "getsockopt"
	case 330:
		return "sendmsg"
	case 331:
		return "recvmsg"
	case 333:
		return "accept4"
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
	case "mmap":
		return 90
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
	case "statfs":
		return 99
	case "fstatfs":
		return 100
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
		return 169
	case "getcwd":
		return 182
	case "sendfile":
		return 186
	case "madvise":
		return 198
	case "gettid":
		return 200
	case "setxattr":
		return 202
	case "getxattr":
		return 205
	case "listxattr":
		return 208
	case "removexattr":
		return 211
	case "tgkill":
		return 241
	case "utimes":
		return 242
	case "unshare":
		return 271
	case "splice":
		return 272
	case "tee":
		return 273
	case "openat":
		return 275
	case "mkdirat":
		return 276
	case "mknodat":
		return 277
	case "fchownat":
		return 278
	case "futimesat":
		return 279
	case "unlinkat":
		return 281
	case "renameat":
		return 282
	case "linkat":
		return 283
	case "symlinkat":
		return 284
	case "readlinkat":
		return 285
	case "fchmodat":
		return 286
	case "faccessat":
		return 287
	case "utimensat":
		return 293
	case "fallocate":
		return 298
	case "dup3":
		return 305
	case "pipe2":
		return 306
	case "socket":
		return 315
	case "bind":
		return 316
	case "connect":
		return 317
	case "listen":
		return 318
	case "accept":
		return 319
	case "getsockname":
		return 320
	case "getpeername":
		return 321
	case "socketpair":
		return 322
	case "sendto":
		return 324
	case "recvfrom":
		return 326
	case "shutdown":
		return 327
	case "setsockopt":
		return 328
	case "getsockopt":
		return 329
	case "sendmsg":
		return 330
	case "recvmsg":
		return 331
	case "accept4":
		return 333
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
	case "mmap":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_PTR}
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
	case "statfs":
		return []ArgType{ARG_STR, ARG_PTR}
	case "fstatfs":
		return []ArgType{ARG_INT, ARG_PTR}
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
	case "unshare":
		return []ArgType{ARG_INT}
	case "splice":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_INT, ARG_INT}
	case "tee":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT}
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
	case "utimensat":
		return []ArgType{ARG_INT, ARG_STR, ARG_PTR}
	case "fallocate":
		return []ArgType{ARG_INT, ARG_INT, ARG_PTR, ARG_PTR}
	case "dup3":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "pipe2":
		return []ArgType{ARG_PTR, ARG_INT}
	case "socket":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "bind":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "connect":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "listen":
		return []ArgType{ARG_INT, ARG_INT}
	case "accept":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "getsockname":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "getpeername":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "socketpair":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_PTR}
	case "sendto":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_PTR}
	case "recvfrom":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_PTR}
	case "shutdown":
		return []ArgType{ARG_INT, ARG_INT}
	case "setsockopt":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_PTR, ARG_INT}
	case "getsockopt":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_PTR, ARG_PTR}
	case "sendmsg":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT}
	case "recvmsg":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT}
	case "accept4":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR, ARG_INT}
	}
	return nil
}
