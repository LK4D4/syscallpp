// generated file; DO NOT EDIT - use go generate in directory with source

// +build arm64,darwin

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
	case 37:
		return "getgid"
	case 40:
		return "setlogin"
	case 46:
		return "revoke"
	case 47:
		return "symlink"
	case 48:
		return "readlink"
	case 50:
		return "umask"
	case 51:
		return "chroot"
	case 54:
		return "munmap"
	case 55:
		return "mprotect"
	case 58:
		return "getgroups"
	case 59:
		return "setgroups"
	case 60:
		return "getpgrp"
	case 61:
		return "setpgid"
	case 65:
		return "getdtablesize"
	case 66:
		return "dup2"
	case 67:
		return "fcntl"
	case 68:
		return "select"
	case 69:
		return "fsync"
	case 70:
		return "setpriority"
	case 71:
		return "socket"
	case 72:
		return "connect"
	case 73:
		return "getpriority"
	case 74:
		return "bind"
	case 75:
		return "setsockopt"
	case 76:
		return "listen"
	case 78:
		return "gettimeofday"
	case 79:
		return "getrusage"
	case 80:
		return "getsockopt"
	case 83:
		return "settimeofday"
	case 84:
		return "fchown"
	case 85:
		return "fchmod"
	case 86:
		return "setreuid"
	case 87:
		return "setregid"
	case 88:
		return "rename"
	case 89:
		return "flock"
	case 90:
		return "mkfifo"
	case 91:
		return "sendto"
	case 92:
		return "shutdown"
	case 93:
		return "socketpair"
	case 94:
		return "mkdir"
	case 95:
		return "rmdir"
	case 96:
		return "utimes"
	case 97:
		return "futimes"
	case 98:
		return "adjtime"
	case 100:
		return "setsid"
	case 101:
		return "getpgid"
	case 102:
		return "setprivexec"
	case 103:
		return "pread"
	case 104:
		return "pwrite"
	case 106:
		return "statfs"
	case 107:
		return "fstatfs"
	case 108:
		return "unmount"
	case 116:
		return "setgid"
	case 117:
		return "setegid"
	case 118:
		return "seteuid"
	case 122:
		return "stat"
	case 123:
		return "fstat"
	case 124:
		return "lstat"
	case 125:
		return "pathconf"
	case 126:
		return "fpathconf"
	case 127:
		return "getrlimit"
	case 128:
		return "setrlimit"
	case 129:
		return "getdirentries"
	case 130:
		return "mmap"
	case 132:
		return "truncate"
	case 133:
		return "ftruncate"
	case 135:
		return "mlock"
	case 136:
		return "munlock"
	case 137:
		return "undelete"
	case 149:
		return "exchangedata"
	case 231:
		return "getsid"
	case 244:
		return "mlockall"
	case 245:
		return "munlockall"
	case 246:
		return "issetugid"
	case 277:
		return "kqueue"
	case 278:
		return "kevent"
	case 279:
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
		return 37
	case "setlogin":
		return 40
	case "revoke":
		return 46
	case "symlink":
		return 47
	case "readlink":
		return 48
	case "umask":
		return 50
	case "chroot":
		return 51
	case "munmap":
		return 54
	case "mprotect":
		return 55
	case "getgroups":
		return 58
	case "setgroups":
		return 59
	case "getpgrp":
		return 60
	case "setpgid":
		return 61
	case "getdtablesize":
		return 65
	case "dup2":
		return 66
	case "fcntl":
		return 67
	case "select":
		return 68
	case "fsync":
		return 69
	case "setpriority":
		return 70
	case "socket":
		return 71
	case "connect":
		return 72
	case "getpriority":
		return 73
	case "bind":
		return 74
	case "setsockopt":
		return 75
	case "listen":
		return 76
	case "gettimeofday":
		return 78
	case "getrusage":
		return 79
	case "getsockopt":
		return 80
	case "settimeofday":
		return 83
	case "fchown":
		return 84
	case "fchmod":
		return 85
	case "setreuid":
		return 86
	case "setregid":
		return 87
	case "rename":
		return 88
	case "flock":
		return 89
	case "mkfifo":
		return 90
	case "sendto":
		return 91
	case "shutdown":
		return 92
	case "socketpair":
		return 93
	case "mkdir":
		return 94
	case "rmdir":
		return 95
	case "utimes":
		return 96
	case "futimes":
		return 97
	case "adjtime":
		return 98
	case "setsid":
		return 100
	case "getpgid":
		return 101
	case "setprivexec":
		return 102
	case "pread":
		return 103
	case "pwrite":
		return 104
	case "statfs":
		return 106
	case "fstatfs":
		return 107
	case "unmount":
		return 108
	case "setgid":
		return 116
	case "setegid":
		return 117
	case "seteuid":
		return 118
	case "stat":
		return 122
	case "fstat":
		return 123
	case "lstat":
		return 124
	case "pathconf":
		return 125
	case "fpathconf":
		return 126
	case "getrlimit":
		return 127
	case "setrlimit":
		return 128
	case "getdirentries":
		return 129
	case "mmap":
		return 130
	case "truncate":
		return 132
	case "ftruncate":
		return 133
	case "mlock":
		return 135
	case "munlock":
		return 136
	case "undelete":
		return 137
	case "exchangedata":
		return 149
	case "getsid":
		return 231
	case "mlockall":
		return 244
	case "munlockall":
		return 245
	case "issetugid":
		return 246
	case "kqueue":
		return 277
	case "kevent":
		return 278
	case "lchown":
		return 279
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
