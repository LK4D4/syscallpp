// generated file; DO NOT EDIT - use go generate in directory with source

// +build amd64,linux

package syscallpp

func GetName(num int) string {
	switch num {
	case 0:
		return "read"
	case 1:
		return "write"
	case 2:
		return "open"
	case 3:
		return "close"
	case 4:
		return "stat"
	case 5:
		return "fstat"
	case 6:
		return "lstat"
	case 7:
		return "poll"
	case 8:
		return "lseek"
	case 9:
		return "mmap"
	case 10:
		return "mprotect"
	case 11:
		return "munmap"
	case 12:
		return "brk"
	case 13:
		return "rt_sigaction"
	case 14:
		return "rt_sigprocmask"
	case 15:
		return "rt_sigreturn"
	case 16:
		return "ioctl"
	case 17:
		return "pread64"
	case 18:
		return "pwrite64"
	case 19:
		return "readv"
	case 20:
		return "writev"
	case 21:
		return "access"
	case 22:
		return "pipe"
	case 23:
		return "select"
	case 24:
		return "sched_yield"
	case 25:
		return "mremap"
	case 26:
		return "msync"
	case 27:
		return "mincore"
	case 28:
		return "madvise"
	case 29:
		return "shmget"
	case 30:
		return "shmat"
	case 31:
		return "shmctl"
	case 32:
		return "dup"
	case 33:
		return "dup2"
	case 34:
		return "pause"
	case 35:
		return "nanosleep"
	case 36:
		return "getitimer"
	case 37:
		return "alarm"
	case 38:
		return "setitimer"
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
	case 56:
		return "clone"
	case 57:
		return "fork"
	case 58:
		return "vfork"
	case 59:
		return "execve"
	case 60:
		return "exit"
	case 61:
		return "wait4"
	case 62:
		return "kill"
	case 63:
		return "uname"
	case 64:
		return "semget"
	case 65:
		return "semop"
	case 66:
		return "semctl"
	case 67:
		return "shmdt"
	case 68:
		return "msgget"
	case 69:
		return "msgsnd"
	case 70:
		return "msgrcv"
	case 71:
		return "msgctl"
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
	case 82:
		return "rename"
	case 83:
		return "mkdir"
	case 84:
		return "rmdir"
	case 85:
		return "creat"
	case 86:
		return "link"
	case 87:
		return "unlink"
	case 88:
		return "symlink"
	case 89:
		return "readlink"
	case 90:
		return "chmod"
	case 91:
		return "fchmod"
	case 92:
		return "chown"
	case 93:
		return "fchown"
	case 94:
		return "lchown"
	case 95:
		return "umask"
	case 96:
		return "gettimeofday"
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
	case 103:
		return "syslog"
	case 104:
		return "getgid"
	case 105:
		return "setuid"
	case 106:
		return "setgid"
	case 107:
		return "geteuid"
	case 108:
		return "getegid"
	case 109:
		return "setpgid"
	case 110:
		return "getppid"
	case 111:
		return "getpgrp"
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
	case 118:
		return "getresuid"
	case 119:
		return "setresgid"
	case 120:
		return "getresgid"
	case 121:
		return "getpgid"
	case 122:
		return "setfsuid"
	case 123:
		return "setfsgid"
	case 124:
		return "getsid"
	case 125:
		return "capget"
	case 126:
		return "capset"
	case 127:
		return "rt_sigpending"
	case 128:
		return "rt_sigtimedwait"
	case 129:
		return "rt_sigqueueinfo"
	case 130:
		return "rt_sigsuspend"
	case 131:
		return "sigaltstack"
	case 132:
		return "utime"
	case 133:
		return "mknod"
	case 134:
		return "uselib"
	case 135:
		return "personality"
	case 136:
		return "ustat"
	case 137:
		return "statfs"
	case 138:
		return "fstatfs"
	case 139:
		return "sysfs"
	case 140:
		return "getpriority"
	case 141:
		return "setpriority"
	case 142:
		return "sched_setparam"
	case 143:
		return "sched_getparam"
	case 144:
		return "sched_setscheduler"
	case 145:
		return "sched_getscheduler"
	case 146:
		return "sched_get_priority_max"
	case 147:
		return "sched_get_priority_min"
	case 148:
		return "sched_rr_get_interval"
	case 149:
		return "mlock"
	case 150:
		return "munlock"
	case 151:
		return "mlockall"
	case 152:
		return "munlockall"
	case 153:
		return "vhangup"
	case 154:
		return "modify_ldt"
	case 155:
		return "pivot_root"
	case 156:
		return "_sysctl"
	case 157:
		return "prctl"
	case 158:
		return "arch_prctl"
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
	case 166:
		return "umount2"
	case 167:
		return "swapon"
	case 168:
		return "swapoff"
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
	case 174:
		return "create_module"
	case 175:
		return "init_module"
	case 176:
		return "delete_module"
	case 177:
		return "get_kernel_syms"
	case 178:
		return "query_module"
	case 179:
		return "quotactl"
	case 180:
		return "nfsservctl"
	case 181:
		return "getpmsg"
	case 182:
		return "putpmsg"
	case 183:
		return "afs_syscall"
	case 184:
		return "tuxcall"
	case 185:
		return "security"
	case 186:
		return "gettid"
	case 187:
		return "readahead"
	case 188:
		return "setxattr"
	case 189:
		return "lsetxattr"
	case 190:
		return "fsetxattr"
	case 191:
		return "getxattr"
	case 192:
		return "lgetxattr"
	case 193:
		return "fgetxattr"
	case 194:
		return "listxattr"
	case 195:
		return "llistxattr"
	case 196:
		return "flistxattr"
	case 197:
		return "removexattr"
	case 198:
		return "lremovexattr"
	case 199:
		return "fremovexattr"
	case 200:
		return "tkill"
	case 201:
		return "time"
	case 202:
		return "futex"
	case 203:
		return "sched_setaffinity"
	case 204:
		return "sched_getaffinity"
	case 205:
		return "set_thread_area"
	case 206:
		return "io_setup"
	case 207:
		return "io_destroy"
	case 208:
		return "io_getevents"
	case 209:
		return "io_submit"
	case 210:
		return "io_cancel"
	case 211:
		return "get_thread_area"
	case 212:
		return "lookup_dcookie"
	case 213:
		return "epoll_create"
	case 214:
		return "epoll_ctl_old"
	case 215:
		return "epoll_wait_old"
	case 216:
		return "remap_file_pages"
	case 217:
		return "getdents64"
	case 218:
		return "set_tid_address"
	case 219:
		return "restart_syscall"
	case 220:
		return "semtimedop"
	case 221:
		return "fadvise64"
	case 222:
		return "timer_create"
	case 223:
		return "timer_settime"
	case 224:
		return "timer_gettime"
	case 225:
		return "timer_getoverrun"
	case 226:
		return "timer_delete"
	case 227:
		return "clock_settime"
	case 228:
		return "clock_gettime"
	case 229:
		return "clock_getres"
	case 230:
		return "clock_nanosleep"
	case 231:
		return "exit_group"
	case 232:
		return "epoll_wait"
	case 233:
		return "epoll_ctl"
	case 234:
		return "tgkill"
	case 235:
		return "utimes"
	case 236:
		return "vserver"
	case 237:
		return "mbind"
	case 238:
		return "set_mempolicy"
	case 239:
		return "get_mempolicy"
	case 240:
		return "mq_open"
	case 241:
		return "mq_unlink"
	case 242:
		return "mq_timedsend"
	case 243:
		return "mq_timedreceive"
	case 244:
		return "mq_notify"
	case 245:
		return "mq_getsetattr"
	case 246:
		return "kexec_load"
	case 247:
		return "waitid"
	case 248:
		return "add_key"
	case 249:
		return "request_key"
	case 250:
		return "keyctl"
	case 251:
		return "ioprio_set"
	case 252:
		return "ioprio_get"
	case 253:
		return "inotify_init"
	case 254:
		return "inotify_add_watch"
	case 255:
		return "inotify_rm_watch"
	case 256:
		return "migrate_pages"
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
	case 262:
		return "newfstatat"
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
	case 270:
		return "pselect6"
	case 271:
		return "ppoll"
	case 272:
		return "unshare"
	case 273:
		return "set_robust_list"
	case 274:
		return "get_robust_list"
	case 275:
		return "splice"
	case 276:
		return "tee"
	case 277:
		return "sync_file_range"
	case 278:
		return "vmsplice"
	case 279:
		return "move_pages"
	case 280:
		return "utimensat"
	case 281:
		return "epoll_pwait"
	case 282:
		return "signalfd"
	case 283:
		return "timerfd_create"
	case 284:
		return "eventfd"
	case 285:
		return "fallocate"
	case 286:
		return "timerfd_settime"
	case 287:
		return "timerfd_gettime"
	case 288:
		return "accept4"
	case 289:
		return "signalfd4"
	case 290:
		return "eventfd2"
	case 291:
		return "epoll_create1"
	case 292:
		return "dup3"
	case 293:
		return "pipe2"
	case 294:
		return "inotify_init1"
	case 295:
		return "preadv"
	case 296:
		return "pwritev"
	case 297:
		return "rt_tgsigqueueinfo"
	case 298:
		return "perf_event_open"
	case 299:
		return "recvmmsg"
	case 300:
		return "fanotify_init"
	case 301:
		return "fanotify_mark"
	case 302:
		return "prlimit64"
	case 303:
		return "name_to_handle_at"
	case 304:
		return "open_by_handle_at"
	case 305:
		return "clock_adjtime"
	case 306:
		return "syncfs"
	case 307:
		return "sendmmsg"
	case 308:
		return "setns"
	case 309:
		return "getcpu"
	case 310:
		return "process_vm_readv"
	case 311:
		return "process_vm_writev"
	case 312:
		return "kcmp"
	case 313:
		return "finit_module"
	case 314:
		return "sched_setattr"
	case 315:
		return "sched_getattr"
	case 316:
		return "renameat2"
	case 317:
		return "seccomp"
	case 318:
		return "getrandom"
	case 319:
		return "memfd_create"
	case 320:
		return "kexec_file_load"
	case 321:
		return "bpf"
	case 322:
		return "execveat"
	case 323:
		return "userfaultfd"
	case 324:
		return "membarrier"
	}
	return "unknown"
}

func GetNum(name string) int {
	switch name {
	case "read":
		return 0
	case "write":
		return 1
	case "open":
		return 2
	case "close":
		return 3
	case "stat":
		return 4
	case "fstat":
		return 5
	case "lstat":
		return 6
	case "poll":
		return 7
	case "lseek":
		return 8
	case "mmap":
		return 9
	case "mprotect":
		return 10
	case "munmap":
		return 11
	case "brk":
		return 12
	case "rt_sigaction":
		return 13
	case "rt_sigprocmask":
		return 14
	case "rt_sigreturn":
		return 15
	case "ioctl":
		return 16
	case "pread64":
		return 17
	case "pwrite64":
		return 18
	case "readv":
		return 19
	case "writev":
		return 20
	case "access":
		return 21
	case "pipe":
		return 22
	case "select":
		return 23
	case "sched_yield":
		return 24
	case "mremap":
		return 25
	case "msync":
		return 26
	case "mincore":
		return 27
	case "madvise":
		return 28
	case "shmget":
		return 29
	case "shmat":
		return 30
	case "shmctl":
		return 31
	case "dup":
		return 32
	case "dup2":
		return 33
	case "pause":
		return 34
	case "nanosleep":
		return 35
	case "getitimer":
		return 36
	case "alarm":
		return 37
	case "setitimer":
		return 38
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
	case "clone":
		return 56
	case "fork":
		return 57
	case "vfork":
		return 58
	case "execve":
		return 59
	case "exit":
		return 60
	case "wait4":
		return 61
	case "kill":
		return 62
	case "uname":
		return 63
	case "semget":
		return 64
	case "semop":
		return 65
	case "semctl":
		return 66
	case "shmdt":
		return 67
	case "msgget":
		return 68
	case "msgsnd":
		return 69
	case "msgrcv":
		return 70
	case "msgctl":
		return 71
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
	case "rename":
		return 82
	case "mkdir":
		return 83
	case "rmdir":
		return 84
	case "creat":
		return 85
	case "link":
		return 86
	case "unlink":
		return 87
	case "symlink":
		return 88
	case "readlink":
		return 89
	case "chmod":
		return 90
	case "fchmod":
		return 91
	case "chown":
		return 92
	case "fchown":
		return 93
	case "lchown":
		return 94
	case "umask":
		return 95
	case "gettimeofday":
		return 96
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
	case "syslog":
		return 103
	case "getgid":
		return 104
	case "setuid":
		return 105
	case "setgid":
		return 106
	case "geteuid":
		return 107
	case "getegid":
		return 108
	case "setpgid":
		return 109
	case "getppid":
		return 110
	case "getpgrp":
		return 111
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
	case "getresuid":
		return 118
	case "setresgid":
		return 119
	case "getresgid":
		return 120
	case "getpgid":
		return 121
	case "setfsuid":
		return 122
	case "setfsgid":
		return 123
	case "getsid":
		return 124
	case "capget":
		return 125
	case "capset":
		return 126
	case "rt_sigpending":
		return 127
	case "rt_sigtimedwait":
		return 128
	case "rt_sigqueueinfo":
		return 129
	case "rt_sigsuspend":
		return 130
	case "sigaltstack":
		return 131
	case "utime":
		return 132
	case "mknod":
		return 133
	case "uselib":
		return 134
	case "personality":
		return 135
	case "ustat":
		return 136
	case "statfs":
		return 137
	case "fstatfs":
		return 138
	case "sysfs":
		return 139
	case "getpriority":
		return 140
	case "setpriority":
		return 141
	case "sched_setparam":
		return 142
	case "sched_getparam":
		return 143
	case "sched_setscheduler":
		return 144
	case "sched_getscheduler":
		return 145
	case "sched_get_priority_max":
		return 146
	case "sched_get_priority_min":
		return 147
	case "sched_rr_get_interval":
		return 148
	case "mlock":
		return 149
	case "munlock":
		return 150
	case "mlockall":
		return 151
	case "munlockall":
		return 152
	case "vhangup":
		return 153
	case "modify_ldt":
		return 154
	case "pivot_root":
		return 155
	case "_sysctl":
		return 156
	case "prctl":
		return 157
	case "arch_prctl":
		return 158
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
	case "umount2":
		return 166
	case "swapon":
		return 167
	case "swapoff":
		return 168
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
	case "create_module":
		return 174
	case "init_module":
		return 175
	case "delete_module":
		return 176
	case "get_kernel_syms":
		return 177
	case "query_module":
		return 178
	case "quotactl":
		return 179
	case "nfsservctl":
		return 180
	case "getpmsg":
		return 181
	case "putpmsg":
		return 182
	case "afs_syscall":
		return 183
	case "tuxcall":
		return 184
	case "security":
		return 185
	case "gettid":
		return 186
	case "readahead":
		return 187
	case "setxattr":
		return 188
	case "lsetxattr":
		return 189
	case "fsetxattr":
		return 190
	case "getxattr":
		return 191
	case "lgetxattr":
		return 192
	case "fgetxattr":
		return 193
	case "listxattr":
		return 194
	case "llistxattr":
		return 195
	case "flistxattr":
		return 196
	case "removexattr":
		return 197
	case "lremovexattr":
		return 198
	case "fremovexattr":
		return 199
	case "tkill":
		return 200
	case "time":
		return 201
	case "futex":
		return 202
	case "sched_setaffinity":
		return 203
	case "sched_getaffinity":
		return 204
	case "set_thread_area":
		return 205
	case "io_setup":
		return 206
	case "io_destroy":
		return 207
	case "io_getevents":
		return 208
	case "io_submit":
		return 209
	case "io_cancel":
		return 210
	case "get_thread_area":
		return 211
	case "lookup_dcookie":
		return 212
	case "epoll_create":
		return 213
	case "epoll_ctl_old":
		return 214
	case "epoll_wait_old":
		return 215
	case "remap_file_pages":
		return 216
	case "getdents64":
		return 217
	case "set_tid_address":
		return 218
	case "restart_syscall":
		return 219
	case "semtimedop":
		return 220
	case "fadvise64":
		return 221
	case "timer_create":
		return 222
	case "timer_settime":
		return 223
	case "timer_gettime":
		return 224
	case "timer_getoverrun":
		return 225
	case "timer_delete":
		return 226
	case "clock_settime":
		return 227
	case "clock_gettime":
		return 228
	case "clock_getres":
		return 229
	case "clock_nanosleep":
		return 230
	case "exit_group":
		return 231
	case "epoll_wait":
		return 232
	case "epoll_ctl":
		return 233
	case "tgkill":
		return 234
	case "utimes":
		return 235
	case "vserver":
		return 236
	case "mbind":
		return 237
	case "set_mempolicy":
		return 238
	case "get_mempolicy":
		return 239
	case "mq_open":
		return 240
	case "mq_unlink":
		return 241
	case "mq_timedsend":
		return 242
	case "mq_timedreceive":
		return 243
	case "mq_notify":
		return 244
	case "mq_getsetattr":
		return 245
	case "kexec_load":
		return 246
	case "waitid":
		return 247
	case "add_key":
		return 248
	case "request_key":
		return 249
	case "keyctl":
		return 250
	case "ioprio_set":
		return 251
	case "ioprio_get":
		return 252
	case "inotify_init":
		return 253
	case "inotify_add_watch":
		return 254
	case "inotify_rm_watch":
		return 255
	case "migrate_pages":
		return 256
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
	case "newfstatat":
		return 262
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
	case "pselect6":
		return 270
	case "ppoll":
		return 271
	case "unshare":
		return 272
	case "set_robust_list":
		return 273
	case "get_robust_list":
		return 274
	case "splice":
		return 275
	case "tee":
		return 276
	case "sync_file_range":
		return 277
	case "vmsplice":
		return 278
	case "move_pages":
		return 279
	case "utimensat":
		return 280
	case "epoll_pwait":
		return 281
	case "signalfd":
		return 282
	case "timerfd_create":
		return 283
	case "eventfd":
		return 284
	case "fallocate":
		return 285
	case "timerfd_settime":
		return 286
	case "timerfd_gettime":
		return 287
	case "accept4":
		return 288
	case "signalfd4":
		return 289
	case "eventfd2":
		return 290
	case "epoll_create1":
		return 291
	case "dup3":
		return 292
	case "pipe2":
		return 293
	case "inotify_init1":
		return 294
	case "preadv":
		return 295
	case "pwritev":
		return 296
	case "rt_tgsigqueueinfo":
		return 297
	case "perf_event_open":
		return 298
	case "recvmmsg":
		return 299
	case "fanotify_init":
		return 300
	case "fanotify_mark":
		return 301
	case "prlimit64":
		return 302
	case "name_to_handle_at":
		return 303
	case "open_by_handle_at":
		return 304
	case "clock_adjtime":
		return 305
	case "syncfs":
		return 306
	case "sendmmsg":
		return 307
	case "setns":
		return 308
	case "getcpu":
		return 309
	case "process_vm_readv":
		return 310
	case "process_vm_writev":
		return 311
	case "kcmp":
		return 312
	case "finit_module":
		return 313
	case "sched_setattr":
		return 314
	case "sched_getattr":
		return 315
	case "renameat2":
		return 316
	case "seccomp":
		return 317
	case "getrandom":
		return 318
	case "memfd_create":
		return 319
	case "kexec_file_load":
		return 320
	case "bpf":
		return 321
	case "execveat":
		return 322
	case "userfaultfd":
		return 323
	case "membarrier":
		return 324
	}
	return -1
}

func GetArgsTypes(name string) []ArgType {
	switch name {
	case "read":
		return []ArgType{ARG_INT, ARG_PTR}
	case "write":
		return []ArgType{ARG_INT, ARG_PTR}
	case "open":
		return []ArgType{ARG_STR, ARG_INT, ARG_INT}
	case "close":
		return []ArgType{ARG_INT}
	case "stat":
		return []ArgType{ARG_STR, ARG_PTR}
	case "fstat":
		return []ArgType{ARG_INT, ARG_PTR}
	case "lstat":
		return []ArgType{ARG_STR, ARG_PTR}
	case "poll":
		return []ArgType(nil)
	case "lseek":
		return []ArgType(nil)
	case "mmap":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_PTR}
	case "mprotect":
		return []ArgType{ARG_PTR, ARG_INT}
	case "munmap":
		return []ArgType{ARG_INT, ARG_INT}
	case "brk":
		return []ArgType(nil)
	case "rt_sigaction":
		return []ArgType(nil)
	case "rt_sigprocmask":
		return []ArgType(nil)
	case "rt_sigreturn":
		return []ArgType(nil)
	case "ioctl":
		return []ArgType(nil)
	case "pread64":
		return []ArgType(nil)
	case "pwrite64":
		return []ArgType(nil)
	case "readv":
		return []ArgType(nil)
	case "writev":
		return []ArgType(nil)
	case "access":
		return []ArgType{ARG_STR, ARG_INT}
	case "pipe":
		return []ArgType{ARG_PTR}
	case "select":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}
	case "sched_yield":
		return []ArgType(nil)
	case "mremap":
		return []ArgType(nil)
	case "msync":
		return []ArgType(nil)
	case "mincore":
		return []ArgType(nil)
	case "madvise":
		return []ArgType{ARG_PTR, ARG_INT}
	case "shmget":
		return []ArgType(nil)
	case "shmat":
		return []ArgType(nil)
	case "shmctl":
		return []ArgType(nil)
	case "dup":
		return []ArgType{ARG_INT}
	case "dup2":
		return []ArgType{ARG_INT, ARG_INT}
	case "pause":
		return []ArgType(nil)
	case "nanosleep":
		return []ArgType{ARG_PTR, ARG_PTR}
	case "getitimer":
		return []ArgType(nil)
	case "alarm":
		return []ArgType(nil)
	case "setitimer":
		return []ArgType(nil)
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
	case "clone":
		return []ArgType(nil)
	case "fork":
		return []ArgType(nil)
	case "vfork":
		return []ArgType(nil)
	case "execve":
		return []ArgType(nil)
	case "exit":
		return []ArgType{ARG_INT}
	case "wait4":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR}
	case "kill":
		return []ArgType{ARG_INT, ARG_PTR}
	case "uname":
		return []ArgType{ARG_PTR}
	case "semget":
		return []ArgType(nil)
	case "semop":
		return []ArgType(nil)
	case "semctl":
		return []ArgType(nil)
	case "shmdt":
		return []ArgType(nil)
	case "msgget":
		return []ArgType(nil)
	case "msgsnd":
		return []ArgType(nil)
	case "msgrcv":
		return []ArgType(nil)
	case "msgctl":
		return []ArgType(nil)
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
	case "rename":
		return []ArgType{ARG_STR, ARG_STR}
	case "mkdir":
		return []ArgType{ARG_STR, ARG_INT}
	case "rmdir":
		return []ArgType{ARG_STR}
	case "creat":
		return []ArgType{ARG_STR, ARG_INT}
	case "link":
		return []ArgType{ARG_STR, ARG_STR}
	case "unlink":
		return []ArgType{ARG_STR}
	case "symlink":
		return []ArgType{ARG_STR, ARG_STR}
	case "readlink":
		return []ArgType{ARG_STR, ARG_PTR}
	case "chmod":
		return []ArgType{ARG_STR, ARG_INT}
	case "fchmod":
		return []ArgType{ARG_INT, ARG_INT}
	case "chown":
		return []ArgType{ARG_STR, ARG_INT, ARG_INT}
	case "fchown":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "lchown":
		return []ArgType{ARG_STR, ARG_INT, ARG_INT}
	case "umask":
		return []ArgType{ARG_INT}
	case "gettimeofday":
		return []ArgType(nil)
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
	case "syslog":
		return []ArgType(nil)
	case "getgid":
		return []ArgType(nil)
	case "setuid":
		return []ArgType{ARG_INT}
	case "setgid":
		return []ArgType{ARG_INT}
	case "geteuid":
		return []ArgType(nil)
	case "getegid":
		return []ArgType(nil)
	case "setpgid":
		return []ArgType{ARG_INT, ARG_INT}
	case "getppid":
		return []ArgType(nil)
	case "getpgrp":
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
	case "getresuid":
		return []ArgType(nil)
	case "setresgid":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "getresgid":
		return []ArgType(nil)
	case "getpgid":
		return []ArgType{ARG_INT}
	case "setfsuid":
		return []ArgType{ARG_INT}
	case "setfsgid":
		return []ArgType{ARG_INT}
	case "getsid":
		return []ArgType(nil)
	case "capget":
		return []ArgType(nil)
	case "capset":
		return []ArgType(nil)
	case "rt_sigpending":
		return []ArgType(nil)
	case "rt_sigtimedwait":
		return []ArgType(nil)
	case "rt_sigqueueinfo":
		return []ArgType(nil)
	case "rt_sigsuspend":
		return []ArgType(nil)
	case "sigaltstack":
		return []ArgType(nil)
	case "utime":
		return []ArgType{ARG_STR, ARG_PTR}
	case "mknod":
		return []ArgType{ARG_STR, ARG_INT, ARG_INT}
	case "uselib":
		return []ArgType(nil)
	case "personality":
		return []ArgType(nil)
	case "ustat":
		return []ArgType{ARG_INT, ARG_PTR}
	case "statfs":
		return []ArgType{ARG_STR, ARG_PTR}
	case "fstatfs":
		return []ArgType{ARG_INT, ARG_PTR}
	case "sysfs":
		return []ArgType(nil)
	case "getpriority":
		return []ArgType{ARG_INT, ARG_INT}
	case "setpriority":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "sched_setparam":
		return []ArgType(nil)
	case "sched_getparam":
		return []ArgType(nil)
	case "sched_setscheduler":
		return []ArgType(nil)
	case "sched_getscheduler":
		return []ArgType(nil)
	case "sched_get_priority_max":
		return []ArgType(nil)
	case "sched_get_priority_min":
		return []ArgType(nil)
	case "sched_rr_get_interval":
		return []ArgType(nil)
	case "mlock":
		return []ArgType{ARG_PTR}
	case "munlock":
		return []ArgType{ARG_PTR}
	case "mlockall":
		return []ArgType{ARG_INT}
	case "munlockall":
		return []ArgType(nil)
	case "vhangup":
		return []ArgType(nil)
	case "modify_ldt":
		return []ArgType(nil)
	case "pivot_root":
		return []ArgType(nil)
	case "_sysctl":
		return []ArgType(nil)
	case "prctl":
		return []ArgType(nil)
	case "arch_prctl":
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
	case "umount2":
		return []ArgType(nil)
	case "swapon":
		return []ArgType(nil)
	case "swapoff":
		return []ArgType(nil)
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
	case "create_module":
		return []ArgType(nil)
	case "init_module":
		return []ArgType(nil)
	case "delete_module":
		return []ArgType(nil)
	case "get_kernel_syms":
		return []ArgType(nil)
	case "query_module":
		return []ArgType(nil)
	case "quotactl":
		return []ArgType(nil)
	case "nfsservctl":
		return []ArgType(nil)
	case "getpmsg":
		return []ArgType(nil)
	case "putpmsg":
		return []ArgType(nil)
	case "afs_syscall":
		return []ArgType(nil)
	case "tuxcall":
		return []ArgType(nil)
	case "security":
		return []ArgType(nil)
	case "gettid":
		return []ArgType(nil)
	case "readahead":
		return []ArgType(nil)
	case "setxattr":
		return []ArgType{ARG_STR, ARG_STR, ARG_PTR, ARG_INT}
	case "lsetxattr":
		return []ArgType(nil)
	case "fsetxattr":
		return []ArgType(nil)
	case "getxattr":
		return []ArgType{ARG_STR, ARG_STR, ARG_PTR}
	case "lgetxattr":
		return []ArgType(nil)
	case "fgetxattr":
		return []ArgType(nil)
	case "listxattr":
		return []ArgType{ARG_STR, ARG_PTR}
	case "llistxattr":
		return []ArgType(nil)
	case "flistxattr":
		return []ArgType(nil)
	case "removexattr":
		return []ArgType{ARG_STR, ARG_STR}
	case "lremovexattr":
		return []ArgType(nil)
	case "fremovexattr":
		return []ArgType(nil)
	case "tkill":
		return []ArgType(nil)
	case "time":
		return []ArgType(nil)
	case "futex":
		return []ArgType(nil)
	case "sched_setaffinity":
		return []ArgType(nil)
	case "sched_getaffinity":
		return []ArgType(nil)
	case "set_thread_area":
		return []ArgType(nil)
	case "io_setup":
		return []ArgType(nil)
	case "io_destroy":
		return []ArgType(nil)
	case "io_getevents":
		return []ArgType(nil)
	case "io_submit":
		return []ArgType(nil)
	case "io_cancel":
		return []ArgType(nil)
	case "get_thread_area":
		return []ArgType(nil)
	case "lookup_dcookie":
		return []ArgType(nil)
	case "epoll_create":
		return []ArgType(nil)
	case "epoll_ctl_old":
		return []ArgType(nil)
	case "epoll_wait_old":
		return []ArgType(nil)
	case "remap_file_pages":
		return []ArgType(nil)
	case "getdents64":
		return []ArgType(nil)
	case "set_tid_address":
		return []ArgType(nil)
	case "restart_syscall":
		return []ArgType(nil)
	case "semtimedop":
		return []ArgType(nil)
	case "fadvise64":
		return []ArgType(nil)
	case "timer_create":
		return []ArgType(nil)
	case "timer_settime":
		return []ArgType(nil)
	case "timer_gettime":
		return []ArgType(nil)
	case "timer_getoverrun":
		return []ArgType(nil)
	case "timer_delete":
		return []ArgType(nil)
	case "clock_settime":
		return []ArgType(nil)
	case "clock_gettime":
		return []ArgType(nil)
	case "clock_getres":
		return []ArgType(nil)
	case "clock_nanosleep":
		return []ArgType(nil)
	case "exit_group":
		return []ArgType(nil)
	case "epoll_wait":
		return []ArgType(nil)
	case "epoll_ctl":
		return []ArgType(nil)
	case "tgkill":
		return []ArgType{ARG_INT, ARG_INT, ARG_PTR}
	case "utimes":
		return []ArgType{ARG_STR, ARG_PTR}
	case "vserver":
		return []ArgType(nil)
	case "mbind":
		return []ArgType(nil)
	case "set_mempolicy":
		return []ArgType(nil)
	case "get_mempolicy":
		return []ArgType(nil)
	case "mq_open":
		return []ArgType(nil)
	case "mq_unlink":
		return []ArgType(nil)
	case "mq_timedsend":
		return []ArgType(nil)
	case "mq_timedreceive":
		return []ArgType(nil)
	case "mq_notify":
		return []ArgType(nil)
	case "mq_getsetattr":
		return []ArgType(nil)
	case "kexec_load":
		return []ArgType(nil)
	case "waitid":
		return []ArgType(nil)
	case "add_key":
		return []ArgType(nil)
	case "request_key":
		return []ArgType(nil)
	case "keyctl":
		return []ArgType(nil)
	case "ioprio_set":
		return []ArgType(nil)
	case "ioprio_get":
		return []ArgType(nil)
	case "inotify_init":
		return []ArgType(nil)
	case "inotify_add_watch":
		return []ArgType(nil)
	case "inotify_rm_watch":
		return []ArgType(nil)
	case "migrate_pages":
		return []ArgType(nil)
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
	case "newfstatat":
		return []ArgType(nil)
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
	case "pselect6":
		return []ArgType(nil)
	case "ppoll":
		return []ArgType(nil)
	case "unshare":
		return []ArgType{ARG_INT}
	case "set_robust_list":
		return []ArgType(nil)
	case "get_robust_list":
		return []ArgType(nil)
	case "splice":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_INT, ARG_INT}
	case "tee":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT}
	case "sync_file_range":
		return []ArgType(nil)
	case "vmsplice":
		return []ArgType(nil)
	case "move_pages":
		return []ArgType(nil)
	case "utimensat":
		return []ArgType{ARG_INT, ARG_STR, ARG_PTR}
	case "epoll_pwait":
		return []ArgType(nil)
	case "signalfd":
		return []ArgType(nil)
	case "timerfd_create":
		return []ArgType(nil)
	case "eventfd":
		return []ArgType(nil)
	case "fallocate":
		return []ArgType{ARG_INT, ARG_INT, ARG_PTR, ARG_PTR}
	case "timerfd_settime":
		return []ArgType(nil)
	case "timerfd_gettime":
		return []ArgType(nil)
	case "accept4":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR, ARG_INT}
	case "signalfd4":
		return []ArgType(nil)
	case "eventfd2":
		return []ArgType(nil)
	case "epoll_create1":
		return []ArgType(nil)
	case "dup3":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "pipe2":
		return []ArgType{ARG_PTR, ARG_INT}
	case "inotify_init1":
		return []ArgType(nil)
	case "preadv":
		return []ArgType(nil)
	case "pwritev":
		return []ArgType(nil)
	case "rt_tgsigqueueinfo":
		return []ArgType(nil)
	case "perf_event_open":
		return []ArgType(nil)
	case "recvmmsg":
		return []ArgType(nil)
	case "fanotify_init":
		return []ArgType(nil)
	case "fanotify_mark":
		return []ArgType(nil)
	case "prlimit64":
		return []ArgType(nil)
	case "name_to_handle_at":
		return []ArgType(nil)
	case "open_by_handle_at":
		return []ArgType(nil)
	case "clock_adjtime":
		return []ArgType(nil)
	case "syncfs":
		return []ArgType(nil)
	case "sendmmsg":
		return []ArgType(nil)
	case "setns":
		return []ArgType(nil)
	case "getcpu":
		return []ArgType(nil)
	case "process_vm_readv":
		return []ArgType(nil)
	case "process_vm_writev":
		return []ArgType(nil)
	case "kcmp":
		return []ArgType(nil)
	case "finit_module":
		return []ArgType(nil)
	case "sched_setattr":
		return []ArgType(nil)
	case "sched_getattr":
		return []ArgType(nil)
	case "renameat2":
		return []ArgType(nil)
	case "seccomp":
		return []ArgType(nil)
	case "getrandom":
		return []ArgType(nil)
	case "memfd_create":
		return []ArgType(nil)
	case "kexec_file_load":
		return []ArgType(nil)
	case "bpf":
		return []ArgType(nil)
	case "execveat":
		return []ArgType(nil)
	case "userfaultfd":
		return []ArgType(nil)
	case "membarrier":
		return []ArgType(nil)
	}
	return nil
}
