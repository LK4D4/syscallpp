// generated file; DO NOT EDIT - use go generate in directory with source

// +build 386,darwin

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
	case 16:
		return "getpid"
	case 17:
		return "setuid"
	case 18:
		return "getuid"
	case 19:
		return "geteuid"
	case 20:
		return "ptrace"
	case 21:
		return "recvmsg"
	case 22:
		return "sendmsg"
	case 23:
		return "recvfrom"
	case 24:
		return "accept"
	case 25:
		return "getpeername"
	case 26:
		return "getsockname"
	case 27:
		return "access"
	case 28:
		return "chflags"
	case 29:
		return "fchflags"
	case 30:
		return "sync"
	case 31:
		return "kill"
	case 32:
		return "getppid"
	case 33:
		return "dup"
	case 34:
		return "pipe"
	case 35:
		return "getegid"
	case 38:
		return "getgid"
	case 41:
		return "setlogin"
	case 47:
		return "revoke"
	case 48:
		return "symlink"
	case 49:
		return "readlink"
	case 51:
		return "umask"
	case 52:
		return "chroot"
	case 55:
		return "munmap"
	case 56:
		return "mprotect"
	case 59:
		return "getgroups"
	case 60:
		return "setgroups"
	case 61:
		return "getpgrp"
	case 62:
		return "setpgid"
	case 66:
		return "getdtablesize"
	case 67:
		return "dup2"
	case 68:
		return "fcntl"
	case 69:
		return "select"
	case 70:
		return "fsync"
	case 71:
		return "setpriority"
	case 72:
		return "socket"
	case 73:
		return "connect"
	case 74:
		return "getpriority"
	case 75:
		return "bind"
	case 76:
		return "setsockopt"
	case 77:
		return "listen"
	case 79:
		return "gettimeofday"
	case 80:
		return "getrusage"
	case 81:
		return "getsockopt"
	case 84:
		return "settimeofday"
	case 85:
		return "fchown"
	case 86:
		return "fchmod"
	case 87:
		return "setreuid"
	case 88:
		return "setregid"
	case 89:
		return "rename"
	case 90:
		return "flock"
	case 91:
		return "mkfifo"
	case 92:
		return "sendto"
	case 93:
		return "shutdown"
	case 94:
		return "socketpair"
	case 95:
		return "mkdir"
	case 96:
		return "rmdir"
	case 97:
		return "utimes"
	case 98:
		return "futimes"
	case 99:
		return "adjtime"
	case 101:
		return "setsid"
	case 102:
		return "getpgid"
	case 103:
		return "setprivexec"
	case 104:
		return "pread"
	case 105:
		return "pwrite"
	case 107:
		return "statfs"
	case 108:
		return "fstatfs"
	case 109:
		return "unmount"
	case 117:
		return "setgid"
	case 118:
		return "setegid"
	case 119:
		return "seteuid"
	case 123:
		return "stat"
	case 124:
		return "fstat"
	case 125:
		return "lstat"
	case 126:
		return "pathconf"
	case 127:
		return "fpathconf"
	case 128:
		return "getrlimit"
	case 129:
		return "setrlimit"
	case 130:
		return "getdirentries"
	case 131:
		return "mmap"
	case 133:
		return "truncate"
	case 134:
		return "ftruncate"
	case 136:
		return "mlock"
	case 137:
		return "munlock"
	case 138:
		return "undelete"
	case 153:
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
		return 16
	case "setuid":
		return 17
	case "getuid":
		return 18
	case "geteuid":
		return 19
	case "ptrace":
		return 20
	case "recvmsg":
		return 21
	case "sendmsg":
		return 22
	case "recvfrom":
		return 23
	case "accept":
		return 24
	case "getpeername":
		return 25
	case "getsockname":
		return 26
	case "access":
		return 27
	case "chflags":
		return 28
	case "fchflags":
		return 29
	case "sync":
		return 30
	case "kill":
		return 31
	case "getppid":
		return 32
	case "dup":
		return 33
	case "pipe":
		return 34
	case "getegid":
		return 35
	case "getgid":
		return 38
	case "setlogin":
		return 41
	case "revoke":
		return 47
	case "symlink":
		return 48
	case "readlink":
		return 49
	case "umask":
		return 51
	case "chroot":
		return 52
	case "munmap":
		return 55
	case "mprotect":
		return 56
	case "getgroups":
		return 59
	case "setgroups":
		return 60
	case "getpgrp":
		return 61
	case "setpgid":
		return 62
	case "getdtablesize":
		return 66
	case "dup2":
		return 67
	case "fcntl":
		return 68
	case "select":
		return 69
	case "fsync":
		return 70
	case "setpriority":
		return 71
	case "socket":
		return 72
	case "connect":
		return 73
	case "getpriority":
		return 74
	case "bind":
		return 75
	case "setsockopt":
		return 76
	case "listen":
		return 77
	case "gettimeofday":
		return 79
	case "getrusage":
		return 80
	case "getsockopt":
		return 81
	case "settimeofday":
		return 84
	case "fchown":
		return 85
	case "fchmod":
		return 86
	case "setreuid":
		return 87
	case "setregid":
		return 88
	case "rename":
		return 89
	case "flock":
		return 90
	case "mkfifo":
		return 91
	case "sendto":
		return 92
	case "shutdown":
		return 93
	case "socketpair":
		return 94
	case "mkdir":
		return 95
	case "rmdir":
		return 96
	case "utimes":
		return 97
	case "futimes":
		return 98
	case "adjtime":
		return 99
	case "setsid":
		return 101
	case "getpgid":
		return 102
	case "setprivexec":
		return 103
	case "pread":
		return 104
	case "pwrite":
		return 105
	case "statfs":
		return 107
	case "fstatfs":
		return 108
	case "unmount":
		return 109
	case "setgid":
		return 117
	case "setegid":
		return 118
	case "seteuid":
		return 119
	case "stat":
		return 123
	case "fstat":
		return 124
	case "lstat":
		return 125
	case "pathconf":
		return 126
	case "fpathconf":
		return 127
	case "getrlimit":
		return 128
	case "setrlimit":
		return 129
	case "getdirentries":
		return 130
	case "mmap":
		return 131
	case "truncate":
		return 133
	case "ftruncate":
		return 134
	case "mlock":
		return 136
	case "munlock":
		return 137
	case "undelete":
		return 138
	case "exchangedata":
		return 153
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
