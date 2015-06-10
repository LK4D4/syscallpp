// generated file; DO NOT EDIT - use go generate in directory with source

// +build arm64,linux

package syscallpp

func GetName(num int) string {
	switch num {
	case 5:
		return "setxattr"
	case 8:
		return "getxattr"
	case 11:
		return "listxattr"
	case 14:
		return "removexattr"
	case 17:
		return "getcwd"
	case 23:
		return "dup"
	case 24:
		return "dup3"
	case 25:
		return "fcntl"
	case 32:
		return "flock"
	case 33:
		return "mknodat"
	case 34:
		return "mkdirat"
	case 35:
		return "unlinkat"
	case 36:
		return "symlinkat"
	case 37:
		return "linkat"
	case 38:
		return "renameat"
	case 40:
		return "mount"
	case 43:
		return "statfs"
	case 44:
		return "fstatfs"
	case 45:
		return "truncate"
	case 46:
		return "ftruncate"
	case 47:
		return "fallocate"
	case 48:
		return "faccessat"
	case 49:
		return "chdir"
	case 50:
		return "fchdir"
	case 51:
		return "chroot"
	case 52:
		return "fchmod"
	case 53:
		return "fchmodat"
	case 54:
		return "fchownat"
	case 55:
		return "fchown"
	case 56:
		return "openat"
	case 57:
		return "close"
	case 59:
		return "pipe2"
	case 63:
		return "read"
	case 64:
		return "write"
	case 71:
		return "sendfile"
	case 76:
		return "splice"
	case 77:
		return "tee"
	case 78:
		return "readlinkat"
	case 79:
		return "fstatat"
	case 80:
		return "fstat"
	case 81:
		return "sync"
	case 82:
		return "fsync"
	case 83:
		return "fdatasync"
	case 89:
		return "utimensat"
	case 90:
		return "acct"
	case 94:
		return "exit"
	case 98:
		return "unshare"
	case 102:
		return "nanosleep"
	case 118:
		return "ptrace"
	case 130:
		return "kill"
	case 132:
		return "tgkill"
	case 141:
		return "setpriority"
	case 142:
		return "getpriority"
	case 143:
		return "reboot"
	case 144:
		return "setregid"
	case 146:
		return "setreuid"
	case 148:
		return "setresuid"
	case 150:
		return "setresgid"
	case 152:
		return "setfsuid"
	case 153:
		return "setfsgid"
	case 154:
		return "times"
	case 155:
		return "setpgid"
	case 156:
		return "getpgid"
	case 158:
		return "setsid"
	case 159:
		return "getgroups"
	case 160:
		return "setgroups"
	case 161:
		return "uname"
	case 162:
		return "sethostname"
	case 163:
		return "setdomainname"
	case 164:
		return "getrlimit"
	case 165:
		return "setrlimit"
	case 166:
		return "getrusage"
	case 167:
		return "umask"
	case 170:
		return "gettimeofday"
	case 171:
		return "settimeofday"
	case 172:
		return "adjtimex"
	case 173:
		return "getpid"
	case 174:
		return "getppid"
	case 175:
		return "getuid"
	case 176:
		return "geteuid"
	case 177:
		return "getgid"
	case 178:
		return "getegid"
	case 179:
		return "gettid"
	case 180:
		return "sysinfo"
	case 199:
		return "socket"
	case 200:
		return "socketpair"
	case 201:
		return "bind"
	case 202:
		return "listen"
	case 203:
		return "accept"
	case 204:
		return "connect"
	case 205:
		return "getsockname"
	case 206:
		return "getpeername"
	case 207:
		return "sendto"
	case 208:
		return "recvfrom"
	case 209:
		return "setsockopt"
	case 210:
		return "getsockopt"
	case 211:
		return "shutdown"
	case 212:
		return "sendmsg"
	case 213:
		return "recvmsg"
	case 216:
		return "munmap"
	case 223:
		return "mmap"
	case 227:
		return "mprotect"
	case 229:
		return "mlock"
	case 230:
		return "munlock"
	case 231:
		return "mlockall"
	case 232:
		return "munlockall"
	case 234:
		return "madvise"
	case 243:
		return "accept4"
	case 246:
		return "wait4"
	}
	return "unknown"
}

func GetNum(name string) int {
	switch name {
	case "setxattr":
		return 5
	case "getxattr":
		return 8
	case "listxattr":
		return 11
	case "removexattr":
		return 14
	case "getcwd":
		return 17
	case "dup":
		return 23
	case "dup3":
		return 24
	case "fcntl":
		return 25
	case "flock":
		return 32
	case "mknodat":
		return 33
	case "mkdirat":
		return 34
	case "unlinkat":
		return 35
	case "symlinkat":
		return 36
	case "linkat":
		return 37
	case "renameat":
		return 38
	case "mount":
		return 40
	case "statfs":
		return 43
	case "fstatfs":
		return 44
	case "truncate":
		return 45
	case "ftruncate":
		return 46
	case "fallocate":
		return 47
	case "faccessat":
		return 48
	case "chdir":
		return 49
	case "fchdir":
		return 50
	case "chroot":
		return 51
	case "fchmod":
		return 52
	case "fchmodat":
		return 53
	case "fchownat":
		return 54
	case "fchown":
		return 55
	case "openat":
		return 56
	case "close":
		return 57
	case "pipe2":
		return 59
	case "read":
		return 63
	case "write":
		return 64
	case "sendfile":
		return 71
	case "splice":
		return 76
	case "tee":
		return 77
	case "readlinkat":
		return 78
	case "fstatat":
		return 79
	case "fstat":
		return 80
	case "sync":
		return 81
	case "fsync":
		return 82
	case "fdatasync":
		return 83
	case "utimensat":
		return 89
	case "acct":
		return 90
	case "exit":
		return 94
	case "unshare":
		return 98
	case "nanosleep":
		return 102
	case "ptrace":
		return 118
	case "kill":
		return 130
	case "tgkill":
		return 132
	case "setpriority":
		return 141
	case "getpriority":
		return 142
	case "reboot":
		return 143
	case "setregid":
		return 144
	case "setreuid":
		return 146
	case "setresuid":
		return 148
	case "setresgid":
		return 150
	case "setfsuid":
		return 152
	case "setfsgid":
		return 153
	case "times":
		return 154
	case "setpgid":
		return 155
	case "getpgid":
		return 156
	case "setsid":
		return 158
	case "getgroups":
		return 159
	case "setgroups":
		return 160
	case "uname":
		return 161
	case "sethostname":
		return 162
	case "setdomainname":
		return 163
	case "getrlimit":
		return 164
	case "setrlimit":
		return 165
	case "getrusage":
		return 166
	case "umask":
		return 167
	case "gettimeofday":
		return 170
	case "settimeofday":
		return 171
	case "adjtimex":
		return 172
	case "getpid":
		return 173
	case "getppid":
		return 174
	case "getuid":
		return 175
	case "geteuid":
		return 176
	case "getgid":
		return 177
	case "getegid":
		return 178
	case "gettid":
		return 179
	case "sysinfo":
		return 180
	case "socket":
		return 199
	case "socketpair":
		return 200
	case "bind":
		return 201
	case "listen":
		return 202
	case "accept":
		return 203
	case "connect":
		return 204
	case "getsockname":
		return 205
	case "getpeername":
		return 206
	case "sendto":
		return 207
	case "recvfrom":
		return 208
	case "setsockopt":
		return 209
	case "getsockopt":
		return 210
	case "shutdown":
		return 211
	case "sendmsg":
		return 212
	case "recvmsg":
		return 213
	case "munmap":
		return 216
	case "mmap":
		return 223
	case "mprotect":
		return 227
	case "mlock":
		return 229
	case "munlock":
		return 230
	case "mlockall":
		return 231
	case "munlockall":
		return 232
	case "madvise":
		return 234
	case "accept4":
		return 243
	case "wait4":
		return 246
	}
	return -1
}

func GetArgsTypes(name string) []ArgType {
	switch name {
	case "setxattr":
		return []ArgType{ARG_STR, ARG_STR, ARG_PTR, ARG_INT}
	case "getxattr":
		return []ArgType{ARG_STR, ARG_STR, ARG_PTR}
	case "listxattr":
		return []ArgType{ARG_STR, ARG_PTR}
	case "removexattr":
		return []ArgType{ARG_STR, ARG_STR}
	case "getcwd":
		return []ArgType{ARG_PTR}
	case "dup":
		return []ArgType{ARG_INT}
	case "dup3":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "fcntl":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "flock":
		return []ArgType{ARG_INT, ARG_INT}
	case "mknodat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_INT}
	case "mkdirat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT}
	case "unlinkat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT}
	case "symlinkat":
		return []ArgType{ARG_STR, ARG_INT, ARG_STR}
	case "linkat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_STR, ARG_INT}
	case "renameat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_STR}
	case "mount":
		return []ArgType{ARG_STR, ARG_STR, ARG_STR, ARG_INT, ARG_PTR}
	case "statfs":
		return []ArgType{ARG_STR, ARG_PTR}
	case "fstatfs":
		return []ArgType{ARG_INT, ARG_PTR}
	case "truncate":
		return []ArgType{ARG_STR, ARG_PTR}
	case "ftruncate":
		return []ArgType{ARG_INT, ARG_PTR}
	case "fallocate":
		return []ArgType{ARG_INT, ARG_INT, ARG_PTR, ARG_PTR}
	case "faccessat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_INT}
	case "chdir":
		return []ArgType{ARG_STR}
	case "fchdir":
		return []ArgType{ARG_INT}
	case "chroot":
		return []ArgType{ARG_STR}
	case "fchmod":
		return []ArgType{ARG_INT, ARG_INT}
	case "fchmodat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_INT}
	case "fchownat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_INT, ARG_INT}
	case "fchown":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "openat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_INT}
	case "close":
		return []ArgType{ARG_INT}
	case "pipe2":
		return []ArgType{ARG_PTR, ARG_INT}
	case "read":
		return []ArgType{ARG_INT, ARG_PTR}
	case "write":
		return []ArgType{ARG_INT, ARG_PTR}
	case "sendfile":
		return []ArgType{ARG_INT, ARG_INT, ARG_PTR, ARG_INT}
	case "splice":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_INT, ARG_INT}
	case "tee":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT}
	case "readlinkat":
		return []ArgType{ARG_INT, ARG_STR, ARG_PTR}
	case "fstatat":
		return []ArgType{ARG_INT, ARG_STR, ARG_PTR, ARG_INT}
	case "fstat":
		return []ArgType{ARG_INT, ARG_PTR}
	case "sync":
		return []ArgType(nil)
	case "fsync":
		return []ArgType{ARG_INT}
	case "fdatasync":
		return []ArgType{ARG_INT}
	case "utimensat":
		return []ArgType{ARG_INT, ARG_STR, ARG_PTR}
	case "acct":
		return []ArgType{ARG_STR}
	case "exit":
		return []ArgType{ARG_INT}
	case "unshare":
		return []ArgType{ARG_INT}
	case "nanosleep":
		return []ArgType{ARG_PTR, ARG_PTR}
	case "ptrace":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT}
	case "kill":
		return []ArgType{ARG_INT, ARG_PTR}
	case "tgkill":
		return []ArgType{ARG_INT, ARG_INT, ARG_PTR}
	case "setpriority":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "getpriority":
		return []ArgType{ARG_INT, ARG_INT}
	case "reboot":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_STR}
	case "setregid":
		return []ArgType{ARG_INT, ARG_INT}
	case "setreuid":
		return []ArgType{ARG_INT, ARG_INT}
	case "setresuid":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "setresgid":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "setfsuid":
		return []ArgType{ARG_INT}
	case "setfsgid":
		return []ArgType{ARG_INT}
	case "times":
		return []ArgType{ARG_PTR}
	case "setpgid":
		return []ArgType{ARG_INT, ARG_INT}
	case "getpgid":
		return []ArgType{ARG_INT}
	case "setsid":
		return []ArgType(nil)
	case "getgroups":
		return []ArgType{ARG_INT, ARG_PTR}
	case "setgroups":
		return []ArgType{ARG_INT, ARG_PTR}
	case "uname":
		return []ArgType{ARG_PTR}
	case "sethostname":
		return []ArgType{ARG_PTR}
	case "setdomainname":
		return []ArgType{ARG_PTR}
	case "getrlimit":
		return []ArgType{ARG_INT, ARG_PTR}
	case "setrlimit":
		return []ArgType{ARG_INT, ARG_PTR}
	case "getrusage":
		return []ArgType{ARG_INT, ARG_PTR}
	case "umask":
		return []ArgType{ARG_INT}
	case "gettimeofday":
		return []ArgType{ARG_PTR}
	case "settimeofday":
		return []ArgType{ARG_PTR}
	case "adjtimex":
		return []ArgType{ARG_PTR}
	case "getpid":
		return []ArgType(nil)
	case "getppid":
		return []ArgType(nil)
	case "getuid":
		return []ArgType(nil)
	case "geteuid":
		return []ArgType(nil)
	case "getgid":
		return []ArgType(nil)
	case "getegid":
		return []ArgType(nil)
	case "gettid":
		return []ArgType(nil)
	case "sysinfo":
		return []ArgType{ARG_PTR}
	case "socket":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "socketpair":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_PTR}
	case "bind":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "listen":
		return []ArgType{ARG_INT, ARG_INT}
	case "accept":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "connect":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "getsockname":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "getpeername":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "sendto":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_PTR}
	case "recvfrom":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_PTR}
	case "setsockopt":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_PTR, ARG_INT}
	case "getsockopt":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_PTR, ARG_PTR}
	case "shutdown":
		return []ArgType{ARG_INT, ARG_INT}
	case "sendmsg":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT}
	case "recvmsg":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT}
	case "munmap":
		return []ArgType{ARG_INT, ARG_INT}
	case "mmap":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_PTR}
	case "mprotect":
		return []ArgType{ARG_PTR, ARG_INT}
	case "mlock":
		return []ArgType{ARG_PTR}
	case "munlock":
		return []ArgType{ARG_PTR}
	case "mlockall":
		return []ArgType{ARG_INT}
	case "munlockall":
		return []ArgType(nil)
	case "madvise":
		return []ArgType{ARG_PTR, ARG_INT}
	case "accept4":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR, ARG_INT}
	case "wait4":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR}
	}
	return nil
}
