// generated file; DO NOT EDIT - use go generate in directory with source

// +build arm,freebsd

package syscallpp

func GetName(num int) string {
	switch num {
	case 0:
		return "exit"
	case 2:
		return "read"
	case 3:
		return "write"
	case 4:
		return "open"
	case 5:
		return "close"
	case 6:
		return "wait4"
	case 7:
		return "link"
	case 8:
		return "unlink"
	case 9:
		return "chdir"
	case 10:
		return "fchdir"
	case 11:
		return "mknod"
	case 12:
		return "chmod"
	case 13:
		return "chown"
	case 15:
		return "getpid"
	case 17:
		return "unmount"
	case 18:
		return "setuid"
	case 19:
		return "getuid"
	case 20:
		return "geteuid"
	case 22:
		return "recvmsg"
	case 23:
		return "sendmsg"
	case 24:
		return "recvfrom"
	case 25:
		return "accept"
	case 26:
		return "getpeername"
	case 27:
		return "getsockname"
	case 28:
		return "access"
	case 29:
		return "chflags"
	case 30:
		return "fchflags"
	case 31:
		return "sync"
	case 32:
		return "kill"
	case 33:
		return "getppid"
	case 34:
		return "dup"
	case 35:
		return "pipe"
	case 36:
		return "getegid"
	case 39:
		return "getgid"
	case 41:
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
	case 57:
		return "munmap"
	case 61:
		return "getgroups"
	case 62:
		return "setgroups"
	case 63:
		return "getpgrp"
	case 64:
		return "setpgid"
	case 68:
		return "getdtablesize"
	case 69:
		return "dup2"
	case 70:
		return "fcntl"
	case 71:
		return "select"
	case 72:
		return "fsync"
	case 73:
		return "setpriority"
	case 74:
		return "socket"
	case 75:
		return "connect"
	case 76:
		return "getpriority"
	case 77:
		return "bind"
	case 78:
		return "setsockopt"
	case 79:
		return "listen"
	case 80:
		return "gettimeofday"
	case 81:
		return "getrusage"
	case 82:
		return "getsockopt"
	case 85:
		return "settimeofday"
	case 86:
		return "fchown"
	case 87:
		return "fchmod"
	case 88:
		return "setreuid"
	case 89:
		return "setregid"
	case 90:
		return "rename"
	case 91:
		return "flock"
	case 92:
		return "mkfifo"
	case 93:
		return "sendto"
	case 94:
		return "shutdown"
	case 95:
		return "socketpair"
	case 96:
		return "mkdir"
	case 97:
		return "rmdir"
	case 98:
		return "utimes"
	case 99:
		return "adjtime"
	case 100:
		return "setsid"
	case 110:
		return "setgid"
	case 111:
		return "setegid"
	case 112:
		return "seteuid"
	case 113:
		return "stat"
	case 114:
		return "fstat"
	case 115:
		return "lstat"
	case 116:
		return "pathconf"
	case 117:
		return "fpathconf"
	case 118:
		return "getrlimit"
	case 119:
		return "setrlimit"
	case 120:
		return "getdirentries"
	case 128:
		return "undelete"
	case 129:
		return "futimes"
	case 130:
		return "getpgid"
	case 140:
		return "nanosleep"
	case 149:
		return "issetugid"
	case 150:
		return "lchown"
	case 171:
		return "getsid"
	case 208:
		return "kqueue"
	case 209:
		return "kevent"
	case 228:
		return "statfs"
	case 229:
		return "fstatfs"
	case 281:
		return "pread"
	case 282:
		return "pwrite"
	case 283:
		return "mmap"
	case 285:
		return "truncate"
	case 286:
		return "ftruncate"
	case 337:
		return "accept4"
	}
	return "unknown"
}

func GetNum(name string) int {
	switch name {
	case "exit":
		return 0
	case "read":
		return 2
	case "write":
		return 3
	case "open":
		return 4
	case "close":
		return 5
	case "wait4":
		return 6
	case "link":
		return 7
	case "unlink":
		return 8
	case "chdir":
		return 9
	case "fchdir":
		return 10
	case "mknod":
		return 11
	case "chmod":
		return 12
	case "chown":
		return 13
	case "getpid":
		return 15
	case "unmount":
		return 17
	case "setuid":
		return 18
	case "getuid":
		return 19
	case "geteuid":
		return 20
	case "recvmsg":
		return 22
	case "sendmsg":
		return 23
	case "recvfrom":
		return 24
	case "accept":
		return 25
	case "getpeername":
		return 26
	case "getsockname":
		return 27
	case "access":
		return 28
	case "chflags":
		return 29
	case "fchflags":
		return 30
	case "sync":
		return 31
	case "kill":
		return 32
	case "getppid":
		return 33
	case "dup":
		return 34
	case "pipe":
		return 35
	case "getegid":
		return 36
	case "getgid":
		return 39
	case "setlogin":
		return 41
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
		return 57
	case "getgroups":
		return 61
	case "setgroups":
		return 62
	case "getpgrp":
		return 63
	case "setpgid":
		return 64
	case "getdtablesize":
		return 68
	case "dup2":
		return 69
	case "fcntl":
		return 70
	case "select":
		return 71
	case "fsync":
		return 72
	case "setpriority":
		return 73
	case "socket":
		return 74
	case "connect":
		return 75
	case "getpriority":
		return 76
	case "bind":
		return 77
	case "setsockopt":
		return 78
	case "listen":
		return 79
	case "gettimeofday":
		return 80
	case "getrusage":
		return 81
	case "getsockopt":
		return 82
	case "settimeofday":
		return 85
	case "fchown":
		return 86
	case "fchmod":
		return 87
	case "setreuid":
		return 88
	case "setregid":
		return 89
	case "rename":
		return 90
	case "flock":
		return 91
	case "mkfifo":
		return 92
	case "sendto":
		return 93
	case "shutdown":
		return 94
	case "socketpair":
		return 95
	case "mkdir":
		return 96
	case "rmdir":
		return 97
	case "utimes":
		return 98
	case "adjtime":
		return 99
	case "setsid":
		return 100
	case "setgid":
		return 110
	case "setegid":
		return 111
	case "seteuid":
		return 112
	case "stat":
		return 113
	case "fstat":
		return 114
	case "lstat":
		return 115
	case "pathconf":
		return 116
	case "fpathconf":
		return 117
	case "getrlimit":
		return 118
	case "setrlimit":
		return 119
	case "getdirentries":
		return 120
	case "undelete":
		return 128
	case "futimes":
		return 129
	case "getpgid":
		return 130
	case "nanosleep":
		return 140
	case "issetugid":
		return 149
	case "lchown":
		return 150
	case "getsid":
		return 171
	case "kqueue":
		return 208
	case "kevent":
		return 209
	case "statfs":
		return 228
	case "fstatfs":
		return 229
	case "pread":
		return 281
	case "pwrite":
		return 282
	case "mmap":
		return 283
	case "truncate":
		return 285
	case "ftruncate":
		return 286
	case "accept4":
		return 337
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
	case "unmount":
		return []ArgType{ARG_STR, ARG_INT}
	case "setuid":
		return []ArgType{ARG_INT}
	case "getuid":
		return []ArgType(nil)
	case "geteuid":
		return []ArgType(nil)
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
		return []ArgType{ARG_INT, ARG_PTR}
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
	case "adjtime":
		return []ArgType{ARG_PTR, ARG_PTR}
	case "setsid":
		return []ArgType(nil)
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
	case "undelete":
		return []ArgType{ARG_STR}
	case "futimes":
		return []ArgType{ARG_INT, ARG_PTR}
	case "getpgid":
		return []ArgType{ARG_INT}
	case "nanosleep":
		return []ArgType{ARG_PTR, ARG_PTR}
	case "issetugid":
		return []ArgType(nil)
	case "lchown":
		return []ArgType{ARG_STR, ARG_INT, ARG_INT}
	case "getsid":
		return []ArgType{ARG_INT}
	case "kqueue":
		return []ArgType(nil)
	case "kevent":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_INT, ARG_PTR}
	case "statfs":
		return []ArgType{ARG_STR, ARG_PTR}
	case "fstatfs":
		return []ArgType{ARG_INT, ARG_PTR}
	case "pread":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "pwrite":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "mmap":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_PTR}
	case "truncate":
		return []ArgType{ARG_STR, ARG_PTR}
	case "ftruncate":
		return []ArgType{ARG_INT, ARG_PTR}
	case "accept4":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR, ARG_INT}
	}
	return nil
}
