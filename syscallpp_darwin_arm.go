// generated file; DO NOT EDIT - use go generate in directory with source

// +build arm,darwin

package syscallpp

func GetName(num int) string {
	switch num {
	case 1:
		return "exit"
	case 3:
		return "read"
	case 4:
		return "write"
	case 5:
		return "open"
	case 6:
		return "close"
	case 7:
		return "wait4"
	case 8:
		return "link"
	case 9:
		return "unlink"
	case 10:
		return "chdir"
	case 11:
		return "fchdir"
	case 12:
		return "mknod"
	case 13:
		return "chmod"
	case 14:
		return "chown"
	case 18:
		return "getpid"
	case 19:
		return "setuid"
	case 20:
		return "getuid"
	case 21:
		return "geteuid"
	case 22:
		return "ptrace"
	case 23:
		return "recvmsg"
	case 24:
		return "sendmsg"
	case 25:
		return "recvfrom"
	case 26:
		return "accept"
	case 27:
		return "getpeername"
	case 28:
		return "getsockname"
	case 29:
		return "access"
	case 30:
		return "chflags"
	case 31:
		return "fchflags"
	case 32:
		return "sync"
	case 33:
		return "kill"
	case 34:
		return "getppid"
	case 35:
		return "dup"
	case 36:
		return "pipe"
	case 37:
		return "getegid"
	case 40:
		return "getgid"
	case 43:
		return "setlogin"
	case 49:
		return "revoke"
	case 50:
		return "symlink"
	case 51:
		return "readlink"
	case 53:
		return "umask"
	case 54:
		return "chroot"
	case 60:
		return "munmap"
	case 61:
		return "mprotect"
	case 64:
		return "getgroups"
	case 65:
		return "setgroups"
	case 66:
		return "getpgrp"
	case 67:
		return "setpgid"
	case 71:
		return "getdtablesize"
	case 72:
		return "dup2"
	case 73:
		return "fcntl"
	case 74:
		return "select"
	case 75:
		return "fsync"
	case 76:
		return "setpriority"
	case 77:
		return "socket"
	case 78:
		return "connect"
	case 79:
		return "getpriority"
	case 80:
		return "bind"
	case 81:
		return "setsockopt"
	case 82:
		return "listen"
	case 84:
		return "gettimeofday"
	case 85:
		return "getrusage"
	case 86:
		return "getsockopt"
	case 89:
		return "settimeofday"
	case 90:
		return "fchown"
	case 91:
		return "fchmod"
	case 92:
		return "setreuid"
	case 93:
		return "setregid"
	case 94:
		return "rename"
	case 95:
		return "flock"
	case 96:
		return "mkfifo"
	case 97:
		return "sendto"
	case 98:
		return "shutdown"
	case 99:
		return "socketpair"
	case 100:
		return "mkdir"
	case 101:
		return "rmdir"
	case 102:
		return "utimes"
	case 103:
		return "futimes"
	case 104:
		return "adjtime"
	case 106:
		return "setsid"
	case 107:
		return "getpgid"
	case 108:
		return "setprivexec"
	case 109:
		return "pread"
	case 110:
		return "pwrite"
	case 112:
		return "statfs"
	case 113:
		return "fstatfs"
	case 114:
		return "unmount"
	case 123:
		return "setgid"
	case 124:
		return "setegid"
	case 125:
		return "seteuid"
	case 128:
		return "stat"
	case 129:
		return "fstat"
	case 130:
		return "lstat"
	case 131:
		return "pathconf"
	case 132:
		return "fpathconf"
	case 133:
		return "getrlimit"
	case 134:
		return "setrlimit"
	case 135:
		return "getdirentries"
	case 136:
		return "mmap"
	case 138:
		return "truncate"
	case 139:
		return "ftruncate"
	case 141:
		return "mlock"
	case 142:
		return "munlock"
	case 143:
		return "undelete"
	case 160:
		return "exchangedata"
	case 235:
		return "getsid"
	case 248:
		return "mlockall"
	case 249:
		return "munlockall"
	case 250:
		return "issetugid"
	case 283:
		return "kqueue"
	case 284:
		return "kevent"
	case 285:
		return "lchown"
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
	case "open":
		return 5
	case "close":
		return 6
	case "wait4":
		return 7
	case "link":
		return 8
	case "unlink":
		return 9
	case "chdir":
		return 10
	case "fchdir":
		return 11
	case "mknod":
		return 12
	case "chmod":
		return 13
	case "chown":
		return 14
	case "getpid":
		return 18
	case "setuid":
		return 19
	case "getuid":
		return 20
	case "geteuid":
		return 21
	case "ptrace":
		return 22
	case "recvmsg":
		return 23
	case "sendmsg":
		return 24
	case "recvfrom":
		return 25
	case "accept":
		return 26
	case "getpeername":
		return 27
	case "getsockname":
		return 28
	case "access":
		return 29
	case "chflags":
		return 30
	case "fchflags":
		return 31
	case "sync":
		return 32
	case "kill":
		return 33
	case "getppid":
		return 34
	case "dup":
		return 35
	case "pipe":
		return 36
	case "getegid":
		return 37
	case "getgid":
		return 40
	case "setlogin":
		return 43
	case "revoke":
		return 49
	case "symlink":
		return 50
	case "readlink":
		return 51
	case "umask":
		return 53
	case "chroot":
		return 54
	case "munmap":
		return 60
	case "mprotect":
		return 61
	case "getgroups":
		return 64
	case "setgroups":
		return 65
	case "getpgrp":
		return 66
	case "setpgid":
		return 67
	case "getdtablesize":
		return 71
	case "dup2":
		return 72
	case "fcntl":
		return 73
	case "select":
		return 74
	case "fsync":
		return 75
	case "setpriority":
		return 76
	case "socket":
		return 77
	case "connect":
		return 78
	case "getpriority":
		return 79
	case "bind":
		return 80
	case "setsockopt":
		return 81
	case "listen":
		return 82
	case "gettimeofday":
		return 84
	case "getrusage":
		return 85
	case "getsockopt":
		return 86
	case "settimeofday":
		return 89
	case "fchown":
		return 90
	case "fchmod":
		return 91
	case "setreuid":
		return 92
	case "setregid":
		return 93
	case "rename":
		return 94
	case "flock":
		return 95
	case "mkfifo":
		return 96
	case "sendto":
		return 97
	case "shutdown":
		return 98
	case "socketpair":
		return 99
	case "mkdir":
		return 100
	case "rmdir":
		return 101
	case "utimes":
		return 102
	case "futimes":
		return 103
	case "adjtime":
		return 104
	case "setsid":
		return 106
	case "getpgid":
		return 107
	case "setprivexec":
		return 108
	case "pread":
		return 109
	case "pwrite":
		return 110
	case "statfs":
		return 112
	case "fstatfs":
		return 113
	case "unmount":
		return 114
	case "setgid":
		return 123
	case "setegid":
		return 124
	case "seteuid":
		return 125
	case "stat":
		return 128
	case "fstat":
		return 129
	case "lstat":
		return 130
	case "pathconf":
		return 131
	case "fpathconf":
		return 132
	case "getrlimit":
		return 133
	case "setrlimit":
		return 134
	case "getdirentries":
		return 135
	case "mmap":
		return 136
	case "truncate":
		return 138
	case "ftruncate":
		return 139
	case "mlock":
		return 141
	case "munlock":
		return 142
	case "undelete":
		return 143
	case "exchangedata":
		return 160
	case "getsid":
		return 235
	case "mlockall":
		return 248
	case "munlockall":
		return 249
	case "issetugid":
		return 250
	case "kqueue":
		return 283
	case "kevent":
		return 284
	case "lchown":
		return 285
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
	case "open":
		return []ArgType{ARG_STR, ARG_INT, ARG_INT}
	case "close":
		return []ArgType{ARG_INT}
	case "wait4":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR}
	case "link":
		return []ArgType{ARG_STR, ARG_STR}
	case "unlink":
		return []ArgType{ARG_STR}
	case "chdir":
		return []ArgType{ARG_STR}
	case "fchdir":
		return []ArgType{ARG_INT}
	case "mknod":
		return []ArgType{ARG_STR, ARG_INT, ARG_INT}
	case "chmod":
		return []ArgType{ARG_STR, ARG_INT}
	case "chown":
		return []ArgType{ARG_STR, ARG_INT, ARG_INT}
	case "getpid":
		return []ArgType(nil)
	case "setuid":
		return []ArgType{ARG_INT}
	case "getuid":
		return []ArgType(nil)
	case "geteuid":
		return []ArgType(nil)
	case "ptrace":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT}
	case "recvmsg":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT}
	case "sendmsg":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT}
	case "recvfrom":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_PTR}
	case "accept":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "getpeername":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "getsockname":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "access":
		return []ArgType{ARG_STR, ARG_INT}
	case "chflags":
		return []ArgType{ARG_STR, ARG_INT}
	case "fchflags":
		return []ArgType{ARG_INT, ARG_INT}
	case "sync":
		return []ArgType(nil)
	case "kill":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "getppid":
		return []ArgType(nil)
	case "dup":
		return []ArgType{ARG_INT}
	case "pipe":
		return []ArgType(nil)
	case "getegid":
		return []ArgType(nil)
	case "getgid":
		return []ArgType(nil)
	case "setlogin":
		return []ArgType{ARG_STR}
	case "revoke":
		return []ArgType{ARG_STR}
	case "symlink":
		return []ArgType{ARG_STR, ARG_STR}
	case "readlink":
		return []ArgType{ARG_STR, ARG_PTR}
	case "umask":
		return []ArgType{ARG_INT}
	case "chroot":
		return []ArgType{ARG_STR}
	case "munmap":
		return []ArgType{ARG_INT, ARG_INT}
	case "mprotect":
		return []ArgType{ARG_PTR, ARG_INT}
	case "getgroups":
		return []ArgType{ARG_INT, ARG_PTR}
	case "setgroups":
		return []ArgType{ARG_INT, ARG_PTR}
	case "getpgrp":
		return []ArgType(nil)
	case "setpgid":
		return []ArgType{ARG_INT, ARG_INT}
	case "getdtablesize":
		return []ArgType(nil)
	case "dup2":
		return []ArgType{ARG_INT, ARG_INT}
	case "fcntl":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "select":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}
	case "fsync":
		return []ArgType{ARG_INT}
	case "setpriority":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "socket":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "connect":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "getpriority":
		return []ArgType{ARG_INT, ARG_INT}
	case "bind":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "setsockopt":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_PTR, ARG_INT}
	case "listen":
		return []ArgType{ARG_INT, ARG_INT}
	case "gettimeofday":
		return []ArgType{ARG_PTR}
	case "getrusage":
		return []ArgType{ARG_INT, ARG_PTR}
	case "getsockopt":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_PTR, ARG_PTR}
	case "settimeofday":
		return []ArgType{ARG_PTR}
	case "fchown":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "fchmod":
		return []ArgType{ARG_INT, ARG_INT}
	case "setreuid":
		return []ArgType{ARG_INT, ARG_INT}
	case "setregid":
		return []ArgType{ARG_INT, ARG_INT}
	case "rename":
		return []ArgType{ARG_STR, ARG_STR}
	case "flock":
		return []ArgType{ARG_INT, ARG_INT}
	case "mkfifo":
		return []ArgType{ARG_STR, ARG_INT}
	case "sendto":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_PTR}
	case "shutdown":
		return []ArgType{ARG_INT, ARG_INT}
	case "socketpair":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_PTR}
	case "mkdir":
		return []ArgType{ARG_STR, ARG_INT}
	case "rmdir":
		return []ArgType{ARG_STR}
	case "utimes":
		return []ArgType{ARG_STR, ARG_PTR}
	case "futimes":
		return []ArgType{ARG_INT, ARG_PTR}
	case "adjtime":
		return []ArgType{ARG_PTR, ARG_PTR}
	case "setsid":
		return []ArgType(nil)
	case "getpgid":
		return []ArgType{ARG_INT}
	case "setprivexec":
		return []ArgType{ARG_INT}
	case "pread":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "pwrite":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "statfs":
		return []ArgType{ARG_STR, ARG_PTR}
	case "fstatfs":
		return []ArgType{ARG_INT, ARG_PTR}
	case "unmount":
		return []ArgType{ARG_STR, ARG_INT}
	case "setgid":
		return []ArgType{ARG_INT}
	case "setegid":
		return []ArgType{ARG_INT}
	case "seteuid":
		return []ArgType{ARG_INT}
	case "stat":
		return []ArgType{ARG_STR, ARG_PTR}
	case "fstat":
		return []ArgType{ARG_INT, ARG_PTR}
	case "lstat":
		return []ArgType{ARG_STR, ARG_PTR}
	case "pathconf":
		return []ArgType{ARG_STR, ARG_INT}
	case "fpathconf":
		return []ArgType{ARG_INT, ARG_INT}
	case "getrlimit":
		return []ArgType{ARG_INT, ARG_PTR}
	case "setrlimit":
		return []ArgType{ARG_INT, ARG_PTR}
	case "getdirentries":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "mmap":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_PTR}
	case "truncate":
		return []ArgType{ARG_STR, ARG_PTR}
	case "ftruncate":
		return []ArgType{ARG_INT, ARG_PTR}
	case "mlock":
		return []ArgType{ARG_PTR}
	case "munlock":
		return []ArgType{ARG_PTR}
	case "undelete":
		return []ArgType{ARG_STR}
	case "exchangedata":
		return []ArgType{ARG_STR, ARG_STR, ARG_INT}
	case "getsid":
		return []ArgType{ARG_INT}
	case "mlockall":
		return []ArgType{ARG_INT}
	case "munlockall":
		return []ArgType(nil)
	case "issetugid":
		return []ArgType(nil)
	case "kqueue":
		return []ArgType(nil)
	case "kevent":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_INT, ARG_PTR}
	case "lchown":
		return []ArgType{ARG_STR, ARG_INT, ARG_INT}
	}
	return nil
}
