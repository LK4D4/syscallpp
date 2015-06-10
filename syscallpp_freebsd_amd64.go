// generated file; DO NOT EDIT - use go generate in directory with source

// +build amd64,freebsd

package syscallpp

func GetName(num int) string {
	switch num {
	case 0:
		return "exit"
	case 1:
		return "fork"
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
	case 14:
		return "obreak"
	case 15:
		return "getpid"
	case 16:
		return "mount"
	case 17:
		return "unmount"
	case 18:
		return "setuid"
	case 19:
		return "getuid"
	case 20:
		return "geteuid"
	case 21:
		return "ptrace"
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
	case 37:
		return "profil"
	case 38:
		return "ktrace"
	case 39:
		return "getgid"
	case 40:
		return "getlogin"
	case 41:
		return "setlogin"
	case 42:
		return "acct"
	case 43:
		return "sigaltstack"
	case 44:
		return "ioctl"
	case 45:
		return "reboot"
	case 46:
		return "revoke"
	case 47:
		return "symlink"
	case 48:
		return "readlink"
	case 49:
		return "execve"
	case 50:
		return "umask"
	case 51:
		return "chroot"
	case 52:
		return "msync"
	case 53:
		return "vfork"
	case 54:
		return "sbrk"
	case 55:
		return "sstk"
	case 56:
		return "ovadvise"
	case 57:
		return "munmap"
	case 58:
		return "mprotect"
	case 59:
		return "madvise"
	case 60:
		return "mincore"
	case 61:
		return "getgroups"
	case 62:
		return "setgroups"
	case 63:
		return "getpgrp"
	case 64:
		return "setpgid"
	case 65:
		return "setitimer"
	case 66:
		return "swapon"
	case 67:
		return "getitimer"
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
	case 83:
		return "readv"
	case 84:
		return "writev"
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
	case 101:
		return "quotactl"
	case 102:
		return "lgetfh"
	case 103:
		return "getfh"
	case 104:
		return "sysarch"
	case 105:
		return "rtprio"
	case 106:
		return "freebsd6_pread"
	case 107:
		return "freebsd6_pwrite"
	case 108:
		return "setfib"
	case 109:
		return "ntp_adjtime"
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
	case 121:
		return "freebsd6_mmap"
	case 122:
		return "freebsd6_lseek"
	case 123:
		return "freebsd6_truncate"
	case 124:
		return "freebsd6_ftruncate"
	case 125:
		return "__sysctl"
	case 126:
		return "mlock"
	case 127:
		return "munlock"
	case 128:
		return "undelete"
	case 129:
		return "futimes"
	case 130:
		return "getpgid"
	case 131:
		return "poll"
	case 132:
		return "clock_gettime"
	case 133:
		return "clock_settime"
	case 134:
		return "clock_getres"
	case 135:
		return "ktimer_create"
	case 136:
		return "ktimer_delete"
	case 137:
		return "ktimer_settime"
	case 138:
		return "ktimer_gettime"
	case 139:
		return "ktimer_getoverrun"
	case 140:
		return "nanosleep"
	case 141:
		return "ffclock_getcounter"
	case 142:
		return "ffclock_setestimate"
	case 143:
		return "ffclock_getestimate"
	case 144:
		return "clock_getcpuclockid2"
	case 145:
		return "ntp_gettime"
	case 146:
		return "minherit"
	case 147:
		return "rfork"
	case 148:
		return "openbsd_poll"
	case 149:
		return "issetugid"
	case 150:
		return "lchown"
	case 151:
		return "getdents"
	case 152:
		return "lchmod"
	case 153:
		return "lutimes"
	case 154:
		return "nstat"
	case 155:
		return "nfstat"
	case 156:
		return "nlstat"
	case 157:
		return "preadv"
	case 158:
		return "pwritev"
	case 159:
		return "fhopen"
	case 160:
		return "fhstat"
	case 161:
		return "modnext"
	case 162:
		return "modstat"
	case 163:
		return "modfnext"
	case 164:
		return "modfind"
	case 165:
		return "kldload"
	case 166:
		return "kldunload"
	case 167:
		return "kldfind"
	case 168:
		return "kldnext"
	case 169:
		return "kldstat"
	case 170:
		return "kldfirstmod"
	case 171:
		return "getsid"
	case 172:
		return "setresuid"
	case 173:
		return "setresgid"
	case 174:
		return "yield"
	case 175:
		return "mlockall"
	case 176:
		return "munlockall"
	case 177:
		return "__getcwd"
	case 178:
		return "sched_setparam"
	case 179:
		return "sched_getparam"
	case 180:
		return "sched_setscheduler"
	case 181:
		return "sched_getscheduler"
	case 182:
		return "sched_yield"
	case 183:
		return "sched_get_priority_max"
	case 184:
		return "sched_get_priority_min"
	case 185:
		return "sched_rr_get_interval"
	case 186:
		return "utrace"
	case 187:
		return "kldsym"
	case 188:
		return "jail"
	case 189:
		return "sigprocmask"
	case 190:
		return "sigsuspend"
	case 191:
		return "sigpending"
	case 192:
		return "sigtimedwait"
	case 193:
		return "sigwaitinfo"
	case 194:
		return "__acl_get_file"
	case 195:
		return "__acl_set_file"
	case 196:
		return "__acl_get_fd"
	case 197:
		return "__acl_set_fd"
	case 198:
		return "__acl_delete_file"
	case 199:
		return "__acl_delete_fd"
	case 200:
		return "__acl_aclcheck_file"
	case 201:
		return "__acl_aclcheck_fd"
	case 202:
		return "extattrctl"
	case 203:
		return "extattr_set_file"
	case 204:
		return "extattr_get_file"
	case 205:
		return "extattr_delete_file"
	case 206:
		return "getresuid"
	case 207:
		return "getresgid"
	case 208:
		return "kqueue"
	case 209:
		return "kevent"
	case 210:
		return "extattr_set_fd"
	case 211:
		return "extattr_get_fd"
	case 212:
		return "extattr_delete_fd"
	case 213:
		return "__setugid"
	case 214:
		return "eaccess"
	case 215:
		return "nmount"
	case 216:
		return "__mac_get_proc"
	case 217:
		return "__mac_set_proc"
	case 218:
		return "__mac_get_fd"
	case 219:
		return "__mac_get_file"
	case 220:
		return "__mac_set_fd"
	case 221:
		return "__mac_set_file"
	case 222:
		return "kenv"
	case 223:
		return "lchflags"
	case 224:
		return "uuidgen"
	case 225:
		return "sendfile"
	case 226:
		return "mac_syscall"
	case 227:
		return "getfsstat"
	case 228:
		return "statfs"
	case 229:
		return "fstatfs"
	case 230:
		return "fhstatfs"
	case 231:
		return "__mac_get_pid"
	case 232:
		return "__mac_get_link"
	case 233:
		return "__mac_set_link"
	case 234:
		return "extattr_set_link"
	case 235:
		return "extattr_get_link"
	case 236:
		return "extattr_delete_link"
	case 237:
		return "__mac_execve"
	case 238:
		return "sigaction"
	case 239:
		return "sigreturn"
	case 240:
		return "getcontext"
	case 241:
		return "setcontext"
	case 242:
		return "swapcontext"
	case 243:
		return "swapoff"
	case 244:
		return "__acl_get_link"
	case 245:
		return "__acl_set_link"
	case 246:
		return "__acl_delete_link"
	case 247:
		return "__acl_aclcheck_link"
	case 248:
		return "sigwait"
	case 249:
		return "thr_create"
	case 250:
		return "thr_exit"
	case 251:
		return "thr_self"
	case 252:
		return "thr_kill"
	case 253:
		return "_umtx_lock"
	case 254:
		return "_umtx_unlock"
	case 255:
		return "jail_attach"
	case 256:
		return "extattr_list_fd"
	case 257:
		return "extattr_list_file"
	case 258:
		return "extattr_list_link"
	case 259:
		return "thr_suspend"
	case 260:
		return "thr_wake"
	case 261:
		return "kldunloadf"
	case 262:
		return "audit"
	case 263:
		return "auditon"
	case 264:
		return "getauid"
	case 265:
		return "setauid"
	case 266:
		return "getaudit"
	case 267:
		return "setaudit"
	case 268:
		return "getaudit_addr"
	case 269:
		return "setaudit_addr"
	case 270:
		return "auditctl"
	case 271:
		return "_umtx_op"
	case 272:
		return "thr_new"
	case 273:
		return "sigqueue"
	case 274:
		return "abort2"
	case 275:
		return "thr_set_name"
	case 276:
		return "rtprio_thread"
	case 277:
		return "sctp_peeloff"
	case 278:
		return "sctp_generic_sendmsg"
	case 279:
		return "sctp_generic_sendmsg_iov"
	case 280:
		return "sctp_generic_recvmsg"
	case 281:
		return "pread"
	case 282:
		return "pwrite"
	case 283:
		return "mmap"
	case 284:
		return "lseek"
	case 285:
		return "truncate"
	case 286:
		return "ftruncate"
	case 287:
		return "thr_kill2"
	case 288:
		return "shm_open"
	case 289:
		return "shm_unlink"
	case 290:
		return "cpuset"
	case 291:
		return "cpuset_setid"
	case 292:
		return "cpuset_getid"
	case 293:
		return "cpuset_getaffinity"
	case 294:
		return "cpuset_setaffinity"
	case 295:
		return "faccessat"
	case 296:
		return "fchmodat"
	case 297:
		return "fchownat"
	case 298:
		return "fexecve"
	case 299:
		return "fstatat"
	case 300:
		return "futimesat"
	case 301:
		return "linkat"
	case 302:
		return "mkdirat"
	case 303:
		return "mkfifoat"
	case 304:
		return "mknodat"
	case 305:
		return "openat"
	case 306:
		return "readlinkat"
	case 307:
		return "renameat"
	case 308:
		return "symlinkat"
	case 309:
		return "unlinkat"
	case 310:
		return "posix_openpt"
	case 311:
		return "jail_get"
	case 312:
		return "jail_set"
	case 313:
		return "jail_remove"
	case 314:
		return "closefrom"
	case 315:
		return "lpathconf"
	case 316:
		return "cap_new"
	case 317:
		return "cap_getrights"
	case 318:
		return "cap_enter"
	case 319:
		return "cap_getmode"
	case 320:
		return "pdfork"
	case 321:
		return "pdkill"
	case 322:
		return "pdgetpid"
	case 323:
		return "pselect"
	case 324:
		return "getloginclass"
	case 325:
		return "setloginclass"
	case 326:
		return "rctl_get_racct"
	case 327:
		return "rctl_get_rules"
	case 328:
		return "rctl_get_limits"
	case 329:
		return "rctl_add_rule"
	case 330:
		return "rctl_remove_rule"
	case 331:
		return "posix_fallocate"
	case 332:
		return "posix_fadvise"
	case 333:
		return "wait6"
	case 334:
		return "bindat"
	case 335:
		return "connectat"
	case 336:
		return "chflagsat"
	case 337:
		return "accept4"
	case 338:
		return "pipe2"
	case 339:
		return "procctl"
	}
	return "unknown"
}

func GetNum(name string) int {
	switch name {
	case "exit":
		return 0
	case "fork":
		return 1
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
	case "obreak":
		return 14
	case "getpid":
		return 15
	case "mount":
		return 16
	case "unmount":
		return 17
	case "setuid":
		return 18
	case "getuid":
		return 19
	case "geteuid":
		return 20
	case "ptrace":
		return 21
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
	case "profil":
		return 37
	case "ktrace":
		return 38
	case "getgid":
		return 39
	case "getlogin":
		return 40
	case "setlogin":
		return 41
	case "acct":
		return 42
	case "sigaltstack":
		return 43
	case "ioctl":
		return 44
	case "reboot":
		return 45
	case "revoke":
		return 46
	case "symlink":
		return 47
	case "readlink":
		return 48
	case "execve":
		return 49
	case "umask":
		return 50
	case "chroot":
		return 51
	case "msync":
		return 52
	case "vfork":
		return 53
	case "sbrk":
		return 54
	case "sstk":
		return 55
	case "ovadvise":
		return 56
	case "munmap":
		return 57
	case "mprotect":
		return 58
	case "madvise":
		return 59
	case "mincore":
		return 60
	case "getgroups":
		return 61
	case "setgroups":
		return 62
	case "getpgrp":
		return 63
	case "setpgid":
		return 64
	case "setitimer":
		return 65
	case "swapon":
		return 66
	case "getitimer":
		return 67
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
	case "readv":
		return 83
	case "writev":
		return 84
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
	case "quotactl":
		return 101
	case "lgetfh":
		return 102
	case "getfh":
		return 103
	case "sysarch":
		return 104
	case "rtprio":
		return 105
	case "freebsd6_pread":
		return 106
	case "freebsd6_pwrite":
		return 107
	case "setfib":
		return 108
	case "ntp_adjtime":
		return 109
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
	case "freebsd6_mmap":
		return 121
	case "freebsd6_lseek":
		return 122
	case "freebsd6_truncate":
		return 123
	case "freebsd6_ftruncate":
		return 124
	case "__sysctl":
		return 125
	case "mlock":
		return 126
	case "munlock":
		return 127
	case "undelete":
		return 128
	case "futimes":
		return 129
	case "getpgid":
		return 130
	case "poll":
		return 131
	case "clock_gettime":
		return 132
	case "clock_settime":
		return 133
	case "clock_getres":
		return 134
	case "ktimer_create":
		return 135
	case "ktimer_delete":
		return 136
	case "ktimer_settime":
		return 137
	case "ktimer_gettime":
		return 138
	case "ktimer_getoverrun":
		return 139
	case "nanosleep":
		return 140
	case "ffclock_getcounter":
		return 141
	case "ffclock_setestimate":
		return 142
	case "ffclock_getestimate":
		return 143
	case "clock_getcpuclockid2":
		return 144
	case "ntp_gettime":
		return 145
	case "minherit":
		return 146
	case "rfork":
		return 147
	case "openbsd_poll":
		return 148
	case "issetugid":
		return 149
	case "lchown":
		return 150
	case "getdents":
		return 151
	case "lchmod":
		return 152
	case "lutimes":
		return 153
	case "nstat":
		return 154
	case "nfstat":
		return 155
	case "nlstat":
		return 156
	case "preadv":
		return 157
	case "pwritev":
		return 158
	case "fhopen":
		return 159
	case "fhstat":
		return 160
	case "modnext":
		return 161
	case "modstat":
		return 162
	case "modfnext":
		return 163
	case "modfind":
		return 164
	case "kldload":
		return 165
	case "kldunload":
		return 166
	case "kldfind":
		return 167
	case "kldnext":
		return 168
	case "kldstat":
		return 169
	case "kldfirstmod":
		return 170
	case "getsid":
		return 171
	case "setresuid":
		return 172
	case "setresgid":
		return 173
	case "yield":
		return 174
	case "mlockall":
		return 175
	case "munlockall":
		return 176
	case "__getcwd":
		return 177
	case "sched_setparam":
		return 178
	case "sched_getparam":
		return 179
	case "sched_setscheduler":
		return 180
	case "sched_getscheduler":
		return 181
	case "sched_yield":
		return 182
	case "sched_get_priority_max":
		return 183
	case "sched_get_priority_min":
		return 184
	case "sched_rr_get_interval":
		return 185
	case "utrace":
		return 186
	case "kldsym":
		return 187
	case "jail":
		return 188
	case "sigprocmask":
		return 189
	case "sigsuspend":
		return 190
	case "sigpending":
		return 191
	case "sigtimedwait":
		return 192
	case "sigwaitinfo":
		return 193
	case "__acl_get_file":
		return 194
	case "__acl_set_file":
		return 195
	case "__acl_get_fd":
		return 196
	case "__acl_set_fd":
		return 197
	case "__acl_delete_file":
		return 198
	case "__acl_delete_fd":
		return 199
	case "__acl_aclcheck_file":
		return 200
	case "__acl_aclcheck_fd":
		return 201
	case "extattrctl":
		return 202
	case "extattr_set_file":
		return 203
	case "extattr_get_file":
		return 204
	case "extattr_delete_file":
		return 205
	case "getresuid":
		return 206
	case "getresgid":
		return 207
	case "kqueue":
		return 208
	case "kevent":
		return 209
	case "extattr_set_fd":
		return 210
	case "extattr_get_fd":
		return 211
	case "extattr_delete_fd":
		return 212
	case "__setugid":
		return 213
	case "eaccess":
		return 214
	case "nmount":
		return 215
	case "__mac_get_proc":
		return 216
	case "__mac_set_proc":
		return 217
	case "__mac_get_fd":
		return 218
	case "__mac_get_file":
		return 219
	case "__mac_set_fd":
		return 220
	case "__mac_set_file":
		return 221
	case "kenv":
		return 222
	case "lchflags":
		return 223
	case "uuidgen":
		return 224
	case "sendfile":
		return 225
	case "mac_syscall":
		return 226
	case "getfsstat":
		return 227
	case "statfs":
		return 228
	case "fstatfs":
		return 229
	case "fhstatfs":
		return 230
	case "__mac_get_pid":
		return 231
	case "__mac_get_link":
		return 232
	case "__mac_set_link":
		return 233
	case "extattr_set_link":
		return 234
	case "extattr_get_link":
		return 235
	case "extattr_delete_link":
		return 236
	case "__mac_execve":
		return 237
	case "sigaction":
		return 238
	case "sigreturn":
		return 239
	case "getcontext":
		return 240
	case "setcontext":
		return 241
	case "swapcontext":
		return 242
	case "swapoff":
		return 243
	case "__acl_get_link":
		return 244
	case "__acl_set_link":
		return 245
	case "__acl_delete_link":
		return 246
	case "__acl_aclcheck_link":
		return 247
	case "sigwait":
		return 248
	case "thr_create":
		return 249
	case "thr_exit":
		return 250
	case "thr_self":
		return 251
	case "thr_kill":
		return 252
	case "_umtx_lock":
		return 253
	case "_umtx_unlock":
		return 254
	case "jail_attach":
		return 255
	case "extattr_list_fd":
		return 256
	case "extattr_list_file":
		return 257
	case "extattr_list_link":
		return 258
	case "thr_suspend":
		return 259
	case "thr_wake":
		return 260
	case "kldunloadf":
		return 261
	case "audit":
		return 262
	case "auditon":
		return 263
	case "getauid":
		return 264
	case "setauid":
		return 265
	case "getaudit":
		return 266
	case "setaudit":
		return 267
	case "getaudit_addr":
		return 268
	case "setaudit_addr":
		return 269
	case "auditctl":
		return 270
	case "_umtx_op":
		return 271
	case "thr_new":
		return 272
	case "sigqueue":
		return 273
	case "abort2":
		return 274
	case "thr_set_name":
		return 275
	case "rtprio_thread":
		return 276
	case "sctp_peeloff":
		return 277
	case "sctp_generic_sendmsg":
		return 278
	case "sctp_generic_sendmsg_iov":
		return 279
	case "sctp_generic_recvmsg":
		return 280
	case "pread":
		return 281
	case "pwrite":
		return 282
	case "mmap":
		return 283
	case "lseek":
		return 284
	case "truncate":
		return 285
	case "ftruncate":
		return 286
	case "thr_kill2":
		return 287
	case "shm_open":
		return 288
	case "shm_unlink":
		return 289
	case "cpuset":
		return 290
	case "cpuset_setid":
		return 291
	case "cpuset_getid":
		return 292
	case "cpuset_getaffinity":
		return 293
	case "cpuset_setaffinity":
		return 294
	case "faccessat":
		return 295
	case "fchmodat":
		return 296
	case "fchownat":
		return 297
	case "fexecve":
		return 298
	case "fstatat":
		return 299
	case "futimesat":
		return 300
	case "linkat":
		return 301
	case "mkdirat":
		return 302
	case "mkfifoat":
		return 303
	case "mknodat":
		return 304
	case "openat":
		return 305
	case "readlinkat":
		return 306
	case "renameat":
		return 307
	case "symlinkat":
		return 308
	case "unlinkat":
		return 309
	case "posix_openpt":
		return 310
	case "jail_get":
		return 311
	case "jail_set":
		return 312
	case "jail_remove":
		return 313
	case "closefrom":
		return 314
	case "lpathconf":
		return 315
	case "cap_new":
		return 316
	case "cap_getrights":
		return 317
	case "cap_enter":
		return 318
	case "cap_getmode":
		return 319
	case "pdfork":
		return 320
	case "pdkill":
		return 321
	case "pdgetpid":
		return 322
	case "pselect":
		return 323
	case "getloginclass":
		return 324
	case "setloginclass":
		return 325
	case "rctl_get_racct":
		return 326
	case "rctl_get_rules":
		return 327
	case "rctl_get_limits":
		return 328
	case "rctl_add_rule":
		return 329
	case "rctl_remove_rule":
		return 330
	case "posix_fallocate":
		return 331
	case "posix_fadvise":
		return 332
	case "wait6":
		return 333
	case "bindat":
		return 334
	case "connectat":
		return 335
	case "chflagsat":
		return 336
	case "accept4":
		return 337
	case "pipe2":
		return 338
	case "procctl":
		return 339
	}
	return -1
}

func GetArgsTypes(name string) []ArgType {
	switch name {
	case "exit":
		return []ArgType{ARG_INT}
	case "fork":
		return []ArgType(nil)
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
	case "obreak":
		return []ArgType(nil)
	case "getpid":
		return []ArgType(nil)
	case "mount":
		return []ArgType(nil)
	case "unmount":
		return []ArgType{ARG_STR, ARG_INT}
	case "setuid":
		return []ArgType{ARG_INT}
	case "getuid":
		return []ArgType(nil)
	case "geteuid":
		return []ArgType(nil)
	case "ptrace":
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
	case "profil":
		return []ArgType(nil)
	case "ktrace":
		return []ArgType(nil)
	case "getgid":
		return []ArgType(nil)
	case "getlogin":
		return []ArgType(nil)
	case "setlogin":
		return []ArgType{ARG_STR}
	case "acct":
		return []ArgType(nil)
	case "sigaltstack":
		return []ArgType(nil)
	case "ioctl":
		return []ArgType(nil)
	case "reboot":
		return []ArgType(nil)
	case "revoke":
		return []ArgType{ARG_STR}
	case "symlink":
		return []ArgType{ARG_STR, ARG_STR}
	case "readlink":
		return []ArgType{ARG_STR, ARG_PTR}
	case "execve":
		return []ArgType(nil)
	case "umask":
		return []ArgType{ARG_INT}
	case "chroot":
		return []ArgType{ARG_STR}
	case "msync":
		return []ArgType(nil)
	case "vfork":
		return []ArgType(nil)
	case "sbrk":
		return []ArgType(nil)
	case "sstk":
		return []ArgType(nil)
	case "ovadvise":
		return []ArgType(nil)
	case "munmap":
		return []ArgType{ARG_INT, ARG_INT}
	case "mprotect":
		return []ArgType(nil)
	case "madvise":
		return []ArgType(nil)
	case "mincore":
		return []ArgType(nil)
	case "getgroups":
		return []ArgType{ARG_INT, ARG_PTR}
	case "setgroups":
		return []ArgType{ARG_INT, ARG_PTR}
	case "getpgrp":
		return []ArgType(nil)
	case "setpgid":
		return []ArgType{ARG_INT, ARG_INT}
	case "setitimer":
		return []ArgType(nil)
	case "swapon":
		return []ArgType(nil)
	case "getitimer":
		return []ArgType(nil)
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
	case "readv":
		return []ArgType(nil)
	case "writev":
		return []ArgType(nil)
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
	case "quotactl":
		return []ArgType(nil)
	case "lgetfh":
		return []ArgType(nil)
	case "getfh":
		return []ArgType(nil)
	case "sysarch":
		return []ArgType(nil)
	case "rtprio":
		return []ArgType(nil)
	case "freebsd6_pread":
		return []ArgType(nil)
	case "freebsd6_pwrite":
		return []ArgType(nil)
	case "setfib":
		return []ArgType(nil)
	case "ntp_adjtime":
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
	case "freebsd6_mmap":
		return []ArgType(nil)
	case "freebsd6_lseek":
		return []ArgType(nil)
	case "freebsd6_truncate":
		return []ArgType(nil)
	case "freebsd6_ftruncate":
		return []ArgType(nil)
	case "__sysctl":
		return []ArgType(nil)
	case "mlock":
		return []ArgType(nil)
	case "munlock":
		return []ArgType(nil)
	case "undelete":
		return []ArgType{ARG_STR}
	case "futimes":
		return []ArgType{ARG_INT, ARG_PTR}
	case "getpgid":
		return []ArgType{ARG_INT}
	case "poll":
		return []ArgType(nil)
	case "clock_gettime":
		return []ArgType(nil)
	case "clock_settime":
		return []ArgType(nil)
	case "clock_getres":
		return []ArgType(nil)
	case "ktimer_create":
		return []ArgType(nil)
	case "ktimer_delete":
		return []ArgType(nil)
	case "ktimer_settime":
		return []ArgType(nil)
	case "ktimer_gettime":
		return []ArgType(nil)
	case "ktimer_getoverrun":
		return []ArgType(nil)
	case "nanosleep":
		return []ArgType{ARG_PTR, ARG_PTR}
	case "ffclock_getcounter":
		return []ArgType(nil)
	case "ffclock_setestimate":
		return []ArgType(nil)
	case "ffclock_getestimate":
		return []ArgType(nil)
	case "clock_getcpuclockid2":
		return []ArgType(nil)
	case "ntp_gettime":
		return []ArgType(nil)
	case "minherit":
		return []ArgType(nil)
	case "rfork":
		return []ArgType(nil)
	case "openbsd_poll":
		return []ArgType(nil)
	case "issetugid":
		return []ArgType(nil)
	case "lchown":
		return []ArgType{ARG_STR, ARG_INT, ARG_INT}
	case "getdents":
		return []ArgType(nil)
	case "lchmod":
		return []ArgType(nil)
	case "lutimes":
		return []ArgType(nil)
	case "nstat":
		return []ArgType(nil)
	case "nfstat":
		return []ArgType(nil)
	case "nlstat":
		return []ArgType(nil)
	case "preadv":
		return []ArgType(nil)
	case "pwritev":
		return []ArgType(nil)
	case "fhopen":
		return []ArgType(nil)
	case "fhstat":
		return []ArgType(nil)
	case "modnext":
		return []ArgType(nil)
	case "modstat":
		return []ArgType(nil)
	case "modfnext":
		return []ArgType(nil)
	case "modfind":
		return []ArgType(nil)
	case "kldload":
		return []ArgType(nil)
	case "kldunload":
		return []ArgType(nil)
	case "kldfind":
		return []ArgType(nil)
	case "kldnext":
		return []ArgType(nil)
	case "kldstat":
		return []ArgType(nil)
	case "kldfirstmod":
		return []ArgType(nil)
	case "getsid":
		return []ArgType{ARG_INT}
	case "setresuid":
		return []ArgType(nil)
	case "setresgid":
		return []ArgType(nil)
	case "yield":
		return []ArgType(nil)
	case "mlockall":
		return []ArgType(nil)
	case "munlockall":
		return []ArgType(nil)
	case "__getcwd":
		return []ArgType(nil)
	case "sched_setparam":
		return []ArgType(nil)
	case "sched_getparam":
		return []ArgType(nil)
	case "sched_setscheduler":
		return []ArgType(nil)
	case "sched_getscheduler":
		return []ArgType(nil)
	case "sched_yield":
		return []ArgType(nil)
	case "sched_get_priority_max":
		return []ArgType(nil)
	case "sched_get_priority_min":
		return []ArgType(nil)
	case "sched_rr_get_interval":
		return []ArgType(nil)
	case "utrace":
		return []ArgType(nil)
	case "kldsym":
		return []ArgType(nil)
	case "jail":
		return []ArgType(nil)
	case "sigprocmask":
		return []ArgType(nil)
	case "sigsuspend":
		return []ArgType(nil)
	case "sigpending":
		return []ArgType(nil)
	case "sigtimedwait":
		return []ArgType(nil)
	case "sigwaitinfo":
		return []ArgType(nil)
	case "__acl_get_file":
		return []ArgType(nil)
	case "__acl_set_file":
		return []ArgType(nil)
	case "__acl_get_fd":
		return []ArgType(nil)
	case "__acl_set_fd":
		return []ArgType(nil)
	case "__acl_delete_file":
		return []ArgType(nil)
	case "__acl_delete_fd":
		return []ArgType(nil)
	case "__acl_aclcheck_file":
		return []ArgType(nil)
	case "__acl_aclcheck_fd":
		return []ArgType(nil)
	case "extattrctl":
		return []ArgType(nil)
	case "extattr_set_file":
		return []ArgType(nil)
	case "extattr_get_file":
		return []ArgType(nil)
	case "extattr_delete_file":
		return []ArgType(nil)
	case "getresuid":
		return []ArgType(nil)
	case "getresgid":
		return []ArgType(nil)
	case "kqueue":
		return []ArgType(nil)
	case "kevent":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_INT, ARG_PTR}
	case "extattr_set_fd":
		return []ArgType(nil)
	case "extattr_get_fd":
		return []ArgType(nil)
	case "extattr_delete_fd":
		return []ArgType(nil)
	case "__setugid":
		return []ArgType(nil)
	case "eaccess":
		return []ArgType(nil)
	case "nmount":
		return []ArgType(nil)
	case "__mac_get_proc":
		return []ArgType(nil)
	case "__mac_set_proc":
		return []ArgType(nil)
	case "__mac_get_fd":
		return []ArgType(nil)
	case "__mac_get_file":
		return []ArgType(nil)
	case "__mac_set_fd":
		return []ArgType(nil)
	case "__mac_set_file":
		return []ArgType(nil)
	case "kenv":
		return []ArgType(nil)
	case "lchflags":
		return []ArgType(nil)
	case "uuidgen":
		return []ArgType(nil)
	case "sendfile":
		return []ArgType(nil)
	case "mac_syscall":
		return []ArgType(nil)
	case "getfsstat":
		return []ArgType{ARG_PTR, ARG_INT}
	case "statfs":
		return []ArgType{ARG_STR, ARG_PTR}
	case "fstatfs":
		return []ArgType{ARG_INT, ARG_PTR}
	case "fhstatfs":
		return []ArgType(nil)
	case "__mac_get_pid":
		return []ArgType(nil)
	case "__mac_get_link":
		return []ArgType(nil)
	case "__mac_set_link":
		return []ArgType(nil)
	case "extattr_set_link":
		return []ArgType(nil)
	case "extattr_get_link":
		return []ArgType(nil)
	case "extattr_delete_link":
		return []ArgType(nil)
	case "__mac_execve":
		return []ArgType(nil)
	case "sigaction":
		return []ArgType(nil)
	case "sigreturn":
		return []ArgType(nil)
	case "getcontext":
		return []ArgType(nil)
	case "setcontext":
		return []ArgType(nil)
	case "swapcontext":
		return []ArgType(nil)
	case "swapoff":
		return []ArgType(nil)
	case "__acl_get_link":
		return []ArgType(nil)
	case "__acl_set_link":
		return []ArgType(nil)
	case "__acl_delete_link":
		return []ArgType(nil)
	case "__acl_aclcheck_link":
		return []ArgType(nil)
	case "sigwait":
		return []ArgType(nil)
	case "thr_create":
		return []ArgType(nil)
	case "thr_exit":
		return []ArgType(nil)
	case "thr_self":
		return []ArgType(nil)
	case "thr_kill":
		return []ArgType(nil)
	case "_umtx_lock":
		return []ArgType(nil)
	case "_umtx_unlock":
		return []ArgType(nil)
	case "jail_attach":
		return []ArgType(nil)
	case "extattr_list_fd":
		return []ArgType(nil)
	case "extattr_list_file":
		return []ArgType(nil)
	case "extattr_list_link":
		return []ArgType(nil)
	case "thr_suspend":
		return []ArgType(nil)
	case "thr_wake":
		return []ArgType(nil)
	case "kldunloadf":
		return []ArgType(nil)
	case "audit":
		return []ArgType(nil)
	case "auditon":
		return []ArgType(nil)
	case "getauid":
		return []ArgType(nil)
	case "setauid":
		return []ArgType(nil)
	case "getaudit":
		return []ArgType(nil)
	case "setaudit":
		return []ArgType(nil)
	case "getaudit_addr":
		return []ArgType(nil)
	case "setaudit_addr":
		return []ArgType(nil)
	case "auditctl":
		return []ArgType(nil)
	case "_umtx_op":
		return []ArgType(nil)
	case "thr_new":
		return []ArgType(nil)
	case "sigqueue":
		return []ArgType(nil)
	case "abort2":
		return []ArgType(nil)
	case "thr_set_name":
		return []ArgType(nil)
	case "rtprio_thread":
		return []ArgType(nil)
	case "sctp_peeloff":
		return []ArgType(nil)
	case "sctp_generic_sendmsg":
		return []ArgType(nil)
	case "sctp_generic_sendmsg_iov":
		return []ArgType(nil)
	case "sctp_generic_recvmsg":
		return []ArgType(nil)
	case "pread":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "pwrite":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "mmap":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_PTR}
	case "lseek":
		return []ArgType(nil)
	case "truncate":
		return []ArgType{ARG_STR, ARG_PTR}
	case "ftruncate":
		return []ArgType{ARG_INT, ARG_PTR}
	case "thr_kill2":
		return []ArgType(nil)
	case "shm_open":
		return []ArgType(nil)
	case "shm_unlink":
		return []ArgType(nil)
	case "cpuset":
		return []ArgType(nil)
	case "cpuset_setid":
		return []ArgType(nil)
	case "cpuset_getid":
		return []ArgType(nil)
	case "cpuset_getaffinity":
		return []ArgType(nil)
	case "cpuset_setaffinity":
		return []ArgType(nil)
	case "faccessat":
		return []ArgType(nil)
	case "fchmodat":
		return []ArgType(nil)
	case "fchownat":
		return []ArgType(nil)
	case "fexecve":
		return []ArgType(nil)
	case "fstatat":
		return []ArgType(nil)
	case "futimesat":
		return []ArgType(nil)
	case "linkat":
		return []ArgType(nil)
	case "mkdirat":
		return []ArgType(nil)
	case "mkfifoat":
		return []ArgType(nil)
	case "mknodat":
		return []ArgType(nil)
	case "openat":
		return []ArgType(nil)
	case "readlinkat":
		return []ArgType(nil)
	case "renameat":
		return []ArgType(nil)
	case "symlinkat":
		return []ArgType(nil)
	case "unlinkat":
		return []ArgType(nil)
	case "posix_openpt":
		return []ArgType(nil)
	case "jail_get":
		return []ArgType(nil)
	case "jail_set":
		return []ArgType(nil)
	case "jail_remove":
		return []ArgType(nil)
	case "closefrom":
		return []ArgType(nil)
	case "lpathconf":
		return []ArgType(nil)
	case "cap_new":
		return []ArgType(nil)
	case "cap_getrights":
		return []ArgType(nil)
	case "cap_enter":
		return []ArgType(nil)
	case "cap_getmode":
		return []ArgType(nil)
	case "pdfork":
		return []ArgType(nil)
	case "pdkill":
		return []ArgType(nil)
	case "pdgetpid":
		return []ArgType(nil)
	case "pselect":
		return []ArgType(nil)
	case "getloginclass":
		return []ArgType(nil)
	case "setloginclass":
		return []ArgType(nil)
	case "rctl_get_racct":
		return []ArgType(nil)
	case "rctl_get_rules":
		return []ArgType(nil)
	case "rctl_get_limits":
		return []ArgType(nil)
	case "rctl_add_rule":
		return []ArgType(nil)
	case "rctl_remove_rule":
		return []ArgType(nil)
	case "posix_fallocate":
		return []ArgType(nil)
	case "posix_fadvise":
		return []ArgType(nil)
	case "wait6":
		return []ArgType(nil)
	case "bindat":
		return []ArgType(nil)
	case "connectat":
		return []ArgType(nil)
	case "chflagsat":
		return []ArgType(nil)
	case "accept4":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR, ARG_INT}
	case "pipe2":
		return []ArgType(nil)
	case "procctl":
		return []ArgType(nil)
	}
	return nil
}
