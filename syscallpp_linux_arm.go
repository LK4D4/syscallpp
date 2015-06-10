// generated file; DO NOT EDIT - use go generate in directory with source

// +build arm,linux

package syscallpp

func GetName(num int) string {
	switch num {
	case 3:
		return "exit"
	case 5:
		return "read"
	case 6:
		return "write"
	case 8:
		return "close"
	case 13:
		return "chdir"
	case 14:
		return "time"
	case 17:
		return "lchown"
	case 19:
		return "getpid"
	case 20:
		return "mount"
	case 23:
		return "getuid"
	case 25:
		return "ptrace"
	case 27:
		return "pause"
	case 28:
		return "utime"
	case 31:
		return "sync"
	case 32:
		return "kill"
	case 36:
		return "dup"
	case 38:
		return "times"
	case 41:
		return "getgid"
	case 42:
		return "geteuid"
	case 43:
		return "getegid"
	case 44:
		return "acct"
	case 47:
		return "fcntl"
	case 48:
		return "setpgid"
	case 49:
		return "umask"
	case 50:
		return "chroot"
	case 51:
		return "ustat"
	case 52:
		return "dup2"
	case 53:
		return "getppid"
	case 55:
		return "setsid"
	case 57:
		return "setreuid"
	case 58:
		return "setregid"
	case 61:
		return "sethostname"
	case 62:
		return "setrlimit"
	case 63:
		return "getrlimit"
	case 64:
		return "getrusage"
	case 65:
		return "gettimeofday"
	case 66:
		return "settimeofday"
	case 67:
		return "getgroups"
	case 68:
		return "setgroups"
	case 69:
		return "select"
	case 74:
		return "reboot"
	case 77:
		return "munmap"
	case 78:
		return "truncate"
	case 79:
		return "ftruncate"
	case 80:
		return "fchmod"
	case 81:
		return "fchown"
	case 82:
		return "getpriority"
	case 83:
		return "setpriority"
	case 90:
		return "stat"
	case 91:
		return "lstat"
	case 92:
		return "fstat"
	case 95:
		return "wait4"
	case 97:
		return "sysinfo"
	case 99:
		return "fsync"
	case 102:
		return "setdomainname"
	case 103:
		return "uname"
	case 104:
		return "adjtimex"
	case 105:
		return "mprotect"
	case 110:
		return "getpgid"
	case 111:
		return "fchdir"
	case 115:
		return "setfsuid"
	case 116:
		return "setfsgid"
	case 118:
		return "getdents"
	case 120:
		return "flock"
	case 125:
		return "fdatasync"
	case 127:
		return "mlock"
	case 128:
		return "munlock"
	case 129:
		return "mlockall"
	case 130:
		return "munlockall"
	case 139:
		return "nanosleep"
	case 141:
		return "setresuid"
	case 145:
		return "setresgid"
	case 158:
		return "getcwd"
	case 162:
		return "sendfile"
	case 165:
		return "mmap2"
	case 193:
		return "madvise"
	case 195:
		return "gettid"
	case 197:
		return "setxattr"
	case 200:
		return "getxattr"
	case 203:
		return "listxattr"
	case 206:
		return "removexattr"
	case 237:
		return "tgkill"
	case 238:
		return "utimes"
	case 250:
		return "socket"
	case 251:
		return "bind"
	case 252:
		return "connect"
	case 253:
		return "listen"
	case 254:
		return "accept"
	case 255:
		return "getsockname"
	case 256:
		return "getpeername"
	case 257:
		return "socketpair"
	case 259:
		return "sendto"
	case 261:
		return "recvfrom"
	case 262:
		return "shutdown"
	case 263:
		return "setsockopt"
	case 264:
		return "getsockopt"
	case 265:
		return "sendmsg"
	case 266:
		return "recvmsg"
	case 291:
		return "openat"
	case 292:
		return "mkdirat"
	case 293:
		return "mknodat"
	case 294:
		return "fchownat"
	case 295:
		return "futimesat"
	case 297:
		return "unlinkat"
	case 298:
		return "renameat"
	case 299:
		return "linkat"
	case 300:
		return "symlinkat"
	case 301:
		return "readlinkat"
	case 302:
		return "fchmodat"
	case 303:
		return "faccessat"
	case 306:
		return "unshare"
	case 309:
		return "splice"
	case 311:
		return "tee"
	case 317:
		return "utimensat"
	case 321:
		return "fallocate"
	case 327:
		return "dup3"
	case 328:
		return "pipe2"
	case 335:
		return "accept4"
	}
	return "unknown"
}

func GetNum(name string) int {
	switch name {
	case "exit":
		return 3
	case "read":
		return 5
	case "write":
		return 6
	case "close":
		return 8
	case "chdir":
		return 13
	case "time":
		return 14
	case "lchown":
		return 17
	case "getpid":
		return 19
	case "mount":
		return 20
	case "getuid":
		return 23
	case "ptrace":
		return 25
	case "pause":
		return 27
	case "utime":
		return 28
	case "sync":
		return 31
	case "kill":
		return 32
	case "dup":
		return 36
	case "times":
		return 38
	case "getgid":
		return 41
	case "geteuid":
		return 42
	case "getegid":
		return 43
	case "acct":
		return 44
	case "fcntl":
		return 47
	case "setpgid":
		return 48
	case "umask":
		return 49
	case "chroot":
		return 50
	case "ustat":
		return 51
	case "dup2":
		return 52
	case "getppid":
		return 53
	case "setsid":
		return 55
	case "setreuid":
		return 57
	case "setregid":
		return 58
	case "sethostname":
		return 61
	case "setrlimit":
		return 62
	case "getrlimit":
		return 63
	case "getrusage":
		return 64
	case "gettimeofday":
		return 65
	case "settimeofday":
		return 66
	case "getgroups":
		return 67
	case "setgroups":
		return 68
	case "select":
		return 69
	case "reboot":
		return 74
	case "munmap":
		return 77
	case "truncate":
		return 78
	case "ftruncate":
		return 79
	case "fchmod":
		return 80
	case "fchown":
		return 81
	case "getpriority":
		return 82
	case "setpriority":
		return 83
	case "stat":
		return 90
	case "lstat":
		return 91
	case "fstat":
		return 92
	case "wait4":
		return 95
	case "sysinfo":
		return 97
	case "fsync":
		return 99
	case "setdomainname":
		return 102
	case "uname":
		return 103
	case "adjtimex":
		return 104
	case "mprotect":
		return 105
	case "getpgid":
		return 110
	case "fchdir":
		return 111
	case "setfsuid":
		return 115
	case "setfsgid":
		return 116
	case "getdents":
		return 118
	case "flock":
		return 120
	case "fdatasync":
		return 125
	case "mlock":
		return 127
	case "munlock":
		return 128
	case "mlockall":
		return 129
	case "munlockall":
		return 130
	case "nanosleep":
		return 139
	case "setresuid":
		return 141
	case "setresgid":
		return 145
	case "getcwd":
		return 158
	case "sendfile":
		return 162
	case "mmap2":
		return 165
	case "madvise":
		return 193
	case "gettid":
		return 195
	case "setxattr":
		return 197
	case "getxattr":
		return 200
	case "listxattr":
		return 203
	case "removexattr":
		return 206
	case "tgkill":
		return 237
	case "utimes":
		return 238
	case "socket":
		return 250
	case "bind":
		return 251
	case "connect":
		return 252
	case "listen":
		return 253
	case "accept":
		return 254
	case "getsockname":
		return 255
	case "getpeername":
		return 256
	case "socketpair":
		return 257
	case "sendto":
		return 259
	case "recvfrom":
		return 261
	case "shutdown":
		return 262
	case "setsockopt":
		return 263
	case "getsockopt":
		return 264
	case "sendmsg":
		return 265
	case "recvmsg":
		return 266
	case "openat":
		return 291
	case "mkdirat":
		return 292
	case "mknodat":
		return 293
	case "fchownat":
		return 294
	case "futimesat":
		return 295
	case "unlinkat":
		return 297
	case "renameat":
		return 298
	case "linkat":
		return 299
	case "symlinkat":
		return 300
	case "readlinkat":
		return 301
	case "fchmodat":
		return 302
	case "faccessat":
		return 303
	case "unshare":
		return 306
	case "splice":
		return 309
	case "tee":
		return 311
	case "utimensat":
		return 317
	case "fallocate":
		return 321
	case "dup3":
		return 327
	case "pipe2":
		return 328
	case "accept4":
		return 335
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
	case "stat":
		return []ArgType{ARG_STR, ARG_PTR}
	case "lstat":
		return []ArgType{ARG_STR, ARG_PTR}
	case "fstat":
		return []ArgType{ARG_INT, ARG_PTR}
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
	case "accept4":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR, ARG_INT}
	}
	return nil
}
