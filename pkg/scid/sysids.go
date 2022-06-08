package scid

var (
    SyscallId2Name = make(map[uint32]string)
    SyscallName2Id = make(map[string]uint32)
)

func init() {

    SyscallId2Name[0] = "read"
    SyscallId2Name[1] = "write"
    SyscallId2Name[2] = "open"
    SyscallId2Name[3] = "close"
    SyscallId2Name[4] = "stat"
    SyscallId2Name[5] = "fstat"
    SyscallId2Name[6] = "lstat"
    SyscallId2Name[7] = "poll"
    SyscallId2Name[8] = "lseek"
    SyscallId2Name[9] = "mmap"
    SyscallId2Name[10] = "mprotect"
    SyscallId2Name[11] = "munmap"
    SyscallId2Name[12] = "brk"
    SyscallId2Name[13] = "rt_sigaction"
    SyscallId2Name[14] = "rt_sigprocmask"
    SyscallId2Name[15] = "rt_sigreturn"
    SyscallId2Name[16] = "ioctl"
    SyscallId2Name[17] = "pread64"
    SyscallId2Name[18] = "pwrite64"
    SyscallId2Name[19] = "readv"
    SyscallId2Name[20] = "writev"
    SyscallId2Name[21] = "access"
    SyscallId2Name[22] = "pipe"
    SyscallId2Name[23] = "select"
    SyscallId2Name[24] = "sched_yield"
    SyscallId2Name[25] = "mremap"
    SyscallId2Name[26] = "msync"
    SyscallId2Name[27] = "mincore"
    SyscallId2Name[28] = "madvise"
    SyscallId2Name[29] = "shmget"
    SyscallId2Name[30] = "shmat"
    SyscallId2Name[31] = "shmctl"
    SyscallId2Name[32] = "dup"
    SyscallId2Name[33] = "dup2"
    SyscallId2Name[34] = "pause"
    SyscallId2Name[35] = "nanosleep"
    SyscallId2Name[36] = "getitimer"
    SyscallId2Name[37] = "alarm"
    SyscallId2Name[38] = "setitimer"
    SyscallId2Name[39] = "getpid"
    SyscallId2Name[40] = "sendfile"
    SyscallId2Name[41] = "socket"
    SyscallId2Name[42] = "connect"
    SyscallId2Name[43] = "accept"
    SyscallId2Name[44] = "sendto"
    SyscallId2Name[45] = "recvfrom"
    SyscallId2Name[46] = "sendmsg"
    SyscallId2Name[47] = "recvmsg"
    SyscallId2Name[48] = "shutdown"
    SyscallId2Name[49] = "bind"
    SyscallId2Name[50] = "listen"
    SyscallId2Name[51] = "getsockname"
    SyscallId2Name[52] = "getpeername"
    SyscallId2Name[53] = "socketpair"
    SyscallId2Name[54] = "setsockopt"
    SyscallId2Name[55] = "getsockopt"
    SyscallId2Name[56] = "clone"
    SyscallId2Name[57] = "fork"
    SyscallId2Name[58] = "vfork"
    SyscallId2Name[59] = "execve"
    SyscallId2Name[60] = "exit"
    SyscallId2Name[61] = "wait4"
    SyscallId2Name[62] = "kill"
    SyscallId2Name[63] = "uname"
    SyscallId2Name[64] = "semget"
    SyscallId2Name[65] = "semop"
    SyscallId2Name[66] = "semctl"
    SyscallId2Name[67] = "shmdt"
    SyscallId2Name[68] = "msgget"
    SyscallId2Name[69] = "msgsnd"
    SyscallId2Name[70] = "msgrcv"
    SyscallId2Name[71] = "msgctl"
    SyscallId2Name[72] = "fcntl"
    SyscallId2Name[73] = "flock"
    SyscallId2Name[74] = "fsync"
    SyscallId2Name[75] = "fdatasync"
    SyscallId2Name[76] = "truncate"
    SyscallId2Name[77] = "ftruncate"
    SyscallId2Name[78] = "getdents"
    SyscallId2Name[79] = "getcwd"
    SyscallId2Name[80] = "chdir"
    SyscallId2Name[81] = "fchdir"
    SyscallId2Name[82] = "rename"
    SyscallId2Name[83] = "mkdir"
    SyscallId2Name[84] = "rmdir"
    SyscallId2Name[85] = "creat"
    SyscallId2Name[86] = "link"
    SyscallId2Name[87] = "unlink"
    SyscallId2Name[88] = "symlink"
    SyscallId2Name[89] = "readlink"
    SyscallId2Name[90] = "chmod"
    SyscallId2Name[91] = "fchmod"
    SyscallId2Name[92] = "chown"
    SyscallId2Name[93] = "fchown"
    SyscallId2Name[94] = "lchown"
    SyscallId2Name[95] = "umask"
    SyscallId2Name[96] = "gettimeofday"
    SyscallId2Name[97] = "getrlimit"
    SyscallId2Name[98] = "getrusage"
    SyscallId2Name[99] = "sysinfo"
    SyscallId2Name[100] = "times"
    SyscallId2Name[101] = "ptrace"
    SyscallId2Name[102] = "getuid"
    SyscallId2Name[103] = "syslog"
    SyscallId2Name[104] = "getgid"
    SyscallId2Name[105] = "setuid"
    SyscallId2Name[106] = "setgid"
    SyscallId2Name[107] = "geteuid"
    SyscallId2Name[108] = "getegid"
    SyscallId2Name[109] = "setpgid"
    SyscallId2Name[110] = "getppid"
    SyscallId2Name[111] = "getpgrp"
    SyscallId2Name[112] = "setsid"
    SyscallId2Name[113] = "setreuid"
    SyscallId2Name[114] = "setregid"
    SyscallId2Name[115] = "getgroups"
    SyscallId2Name[116] = "setgroups"
    SyscallId2Name[117] = "setresuid"
    SyscallId2Name[118] = "getresuid"
    SyscallId2Name[119] = "setresgid"
    SyscallId2Name[120] = "getresgid"
    SyscallId2Name[121] = "getpgid"
    SyscallId2Name[122] = "setfsuid"
    SyscallId2Name[123] = "setfsgid"
    SyscallId2Name[124] = "getsid"
    SyscallId2Name[125] = "capget"
    SyscallId2Name[126] = "capset"
    SyscallId2Name[127] = "rt_sigpending"
    SyscallId2Name[128] = "rt_sigtimedwait"
    SyscallId2Name[129] = "rt_sigqueueinfo"
    SyscallId2Name[130] = "rt_sigsuspend"
    SyscallId2Name[131] = "sigaltstack"
    SyscallId2Name[132] = "utime"
    SyscallId2Name[133] = "mknod"
    SyscallId2Name[134] = "uselib"
    SyscallId2Name[135] = "personality"
    SyscallId2Name[136] = "ustat"
    SyscallId2Name[137] = "statfs"
    SyscallId2Name[138] = "fstatfs"
    SyscallId2Name[139] = "sysfs"
    SyscallId2Name[140] = "getpriority"
    SyscallId2Name[141] = "setpriority"
    SyscallId2Name[142] = "sched_setparam"
    SyscallId2Name[143] = "sched_getparam"
    SyscallId2Name[144] = "sched_setscheduler"
    SyscallId2Name[145] = "sched_getscheduler"
    SyscallId2Name[146] = "sched_get_priority_max"
    SyscallId2Name[147] = "sched_get_priority_min"
    SyscallId2Name[148] = "sched_rr_get_interval"
    SyscallId2Name[149] = "mlock"
    SyscallId2Name[150] = "munlock"
    SyscallId2Name[151] = "mlockall"
    SyscallId2Name[152] = "munlockall"
    SyscallId2Name[153] = "vhangup"
    SyscallId2Name[154] = "modify_ldt"
    SyscallId2Name[155] = "pivot_root"
    SyscallId2Name[156] = "_sysctl"
    SyscallId2Name[157] = "prctl"
    SyscallId2Name[158] = "arch_prctl"
    SyscallId2Name[159] = "adjtimex"
    SyscallId2Name[160] = "setrlimit"
    SyscallId2Name[161] = "chroot"
    SyscallId2Name[162] = "sync"
    SyscallId2Name[163] = "acct"
    SyscallId2Name[164] = "settimeofday"
    SyscallId2Name[165] = "mount"
    SyscallId2Name[166] = "umount2"
    SyscallId2Name[167] = "swapon"
    SyscallId2Name[168] = "swapoff"
    SyscallId2Name[169] = "reboot"
    SyscallId2Name[170] = "sethostname"
    SyscallId2Name[171] = "setdomainname"
    SyscallId2Name[172] = "iopl"
    SyscallId2Name[173] = "ioperm"
    SyscallId2Name[174] = "create_module"
    SyscallId2Name[175] = "init_module"
    SyscallId2Name[176] = "delete_module"
    SyscallId2Name[177] = "get_kernel_syms"
    SyscallId2Name[178] = "query_module"
    SyscallId2Name[179] = "quotactl"
    SyscallId2Name[180] = "nfsservctl"
    SyscallId2Name[181] = "getpmsg"
    SyscallId2Name[182] = "putpmsg"
    SyscallId2Name[183] = "afs_syscall"
    SyscallId2Name[184] = "tuxcall"
    SyscallId2Name[185] = "security"
    SyscallId2Name[186] = "gettid"
    SyscallId2Name[187] = "readahead"
    SyscallId2Name[188] = "setxattr"
    SyscallId2Name[189] = "lsetxattr"
    SyscallId2Name[190] = "fsetxattr"
    SyscallId2Name[191] = "getxattr"
    SyscallId2Name[192] = "lgetxattr"
    SyscallId2Name[193] = "fgetxattr"
    SyscallId2Name[194] = "listxattr"
    SyscallId2Name[195] = "llistxattr"
    SyscallId2Name[196] = "flistxattr"
    SyscallId2Name[197] = "removexattr"
    SyscallId2Name[198] = "lremovexattr"
    SyscallId2Name[199] = "fremovexattr"
    SyscallId2Name[200] = "tkill"
    SyscallId2Name[201] = "time"
    SyscallId2Name[202] = "futex"
    SyscallId2Name[203] = "sched_setaffinity"
    SyscallId2Name[204] = "sched_getaffinity"
    SyscallId2Name[205] = "set_thread_area"
    SyscallId2Name[206] = "io_setup"
    SyscallId2Name[207] = "io_destroy"
    SyscallId2Name[208] = "io_getevents"
    SyscallId2Name[209] = "io_submit"
    SyscallId2Name[210] = "io_cancel"
    SyscallId2Name[211] = "get_thread_area"
    SyscallId2Name[212] = "lookup_dcookie"
    SyscallId2Name[213] = "epoll_create"
    SyscallId2Name[214] = "epoll_ctl_old"
    SyscallId2Name[215] = "epoll_wait_old"
    SyscallId2Name[216] = "remap_file_pages"
    SyscallId2Name[217] = "getdents64"
    SyscallId2Name[218] = "set_tid_address"
    SyscallId2Name[219] = "restart_syscall"
    SyscallId2Name[220] = "semtimedop"
    SyscallId2Name[221] = "fadvise64"
    SyscallId2Name[222] = "timer_create"
    SyscallId2Name[223] = "timer_settime"
    SyscallId2Name[224] = "timer_gettime"
    SyscallId2Name[225] = "timer_getoverrun"
    SyscallId2Name[226] = "timer_delete"
    SyscallId2Name[227] = "clock_settime"
    SyscallId2Name[228] = "clock_gettime"
    SyscallId2Name[229] = "clock_getres"
    SyscallId2Name[230] = "clock_nanosleep"
    SyscallId2Name[231] = "exit_group"
    SyscallId2Name[232] = "epoll_wait"
    SyscallId2Name[233] = "epoll_ctl"
    SyscallId2Name[234] = "tgkill"
    SyscallId2Name[235] = "utimes"
    SyscallId2Name[236] = "vserver"
    SyscallId2Name[237] = "mbind"
    SyscallId2Name[238] = "set_mempolicy"
    SyscallId2Name[239] = "get_mempolicy"
    SyscallId2Name[240] = "mq_open"
    SyscallId2Name[241] = "mq_unlink"
    SyscallId2Name[242] = "mq_timedsend"
    SyscallId2Name[243] = "mq_timedreceive"
    SyscallId2Name[244] = "mq_notify"
    SyscallId2Name[245] = "mq_getsetattr"
    SyscallId2Name[246] = "kexec_load"
    SyscallId2Name[247] = "waitid"
    SyscallId2Name[248] = "add_key"
    SyscallId2Name[249] = "request_key"
    SyscallId2Name[250] = "keyctl"
    SyscallId2Name[251] = "ioprio_set"
    SyscallId2Name[252] = "ioprio_get"
    SyscallId2Name[253] = "inotify_init"
    SyscallId2Name[254] = "inotify_add_watch"
    SyscallId2Name[255] = "inotify_rm_watch"
    SyscallId2Name[256] = "migrate_pages"
    SyscallId2Name[257] = "openat"
    SyscallId2Name[258] = "mkdirat"
    SyscallId2Name[259] = "mknodat"
    SyscallId2Name[260] = "fchownat"
    SyscallId2Name[261] = "futimesat"
    SyscallId2Name[262] = "newfstatat"
    SyscallId2Name[263] = "unlinkat"
    SyscallId2Name[264] = "renameat"
    SyscallId2Name[265] = "linkat"
    SyscallId2Name[266] = "symlinkat"
    SyscallId2Name[267] = "readlinkat"
    SyscallId2Name[268] = "fchmodat"
    SyscallId2Name[269] = "faccessat"
    SyscallId2Name[270] = "pselect6"
    SyscallId2Name[271] = "ppoll"
    SyscallId2Name[272] = "unshare"
    SyscallId2Name[273] = "set_robust_list"
    SyscallId2Name[274] = "get_robust_list"
    SyscallId2Name[275] = "splice"
    SyscallId2Name[276] = "tee"
    SyscallId2Name[277] = "sync_file_range"
    SyscallId2Name[278] = "vmsplice"
    SyscallId2Name[279] = "move_pages"
    SyscallId2Name[280] = "utimensat"
    SyscallId2Name[281] = "epoll_pwait"
    SyscallId2Name[282] = "signalfd"
    SyscallId2Name[283] = "timerfd_create"
    SyscallId2Name[284] = "eventfd"
    SyscallId2Name[285] = "fallocate"
    SyscallId2Name[286] = "timerfd_settime"
    SyscallId2Name[287] = "timerfd_gettime"
    SyscallId2Name[288] = "accept4"
    SyscallId2Name[289] = "signalfd4"
    SyscallId2Name[290] = "eventfd2"
    SyscallId2Name[291] = "epoll_create1"
    SyscallId2Name[292] = "dup3"
    SyscallId2Name[293] = "pipe2"
    SyscallId2Name[294] = "inotify_init1"
    SyscallId2Name[295] = "preadv"
    SyscallId2Name[296] = "pwritev"
    SyscallId2Name[297] = "rt_tgsigqueueinfo"
    SyscallId2Name[298] = "perf_event_open"
    SyscallId2Name[299] = "recvmmsg"
    SyscallId2Name[300] = "fanotify_init"
    SyscallId2Name[301] = "fanotify_mark"
    SyscallId2Name[302] = "prlimit64"
    SyscallId2Name[303] = "name_to_handle_at"
    SyscallId2Name[304] = "open_by_handle_at"
    SyscallId2Name[305] = "clock_adjtime"
    SyscallId2Name[306] = "syncfs"
    SyscallId2Name[307] = "sendmmsg"
    SyscallId2Name[308] = "setns"
    SyscallId2Name[309] = "getcpu"
    SyscallId2Name[310] = "process_vm_readv"
    SyscallId2Name[311] = "process_vm_writev"
    SyscallId2Name[312] = "kcmp"
    SyscallId2Name[313] = "finit_module"
    SyscallId2Name[314] = "sched_setattr"
    SyscallId2Name[315] = "sched_getattr"
    SyscallId2Name[316] = "renameat2"
    SyscallId2Name[317] = "seccomp"
    SyscallId2Name[318] = "getrandom"
    SyscallId2Name[319] = "memfd_create"
    SyscallId2Name[320] = "kexec_file_load"
    SyscallId2Name[321] = "bpf"
    SyscallId2Name[322] = "execveat"
    SyscallId2Name[323] = "userfaultfd"
    SyscallId2Name[324] = "membarrier"
    SyscallId2Name[325] = "mlock2"
    SyscallId2Name[326] = "copy_file_range"
    SyscallId2Name[327] = "preadv2"
    SyscallId2Name[328] = "pwritev2"
    SyscallId2Name[329] = "pkey_mprotect"
    SyscallId2Name[330] = "pkey_alloc"
    SyscallId2Name[331] = "pkey_free"
    SyscallId2Name[332] = "statx"
    SyscallId2Name[333] = "io_pgetevents"
    SyscallId2Name[334] = "rseq"
    SyscallId2Name[424] = "pidfd_send_signal"
    SyscallId2Name[425] = "io_uring_setup"
    SyscallId2Name[426] = "io_uring_enter"
    SyscallId2Name[427] = "io_uring_register"
    SyscallId2Name[428] = "open_tree"
    SyscallId2Name[429] = "move_mount"
    SyscallId2Name[430] = "fsopen"
    SyscallId2Name[431] = "fsconfig"
    SyscallId2Name[432] = "fsmount"
    SyscallId2Name[433] = "fspick"
    SyscallId2Name[434] = "pidfd_open"
    SyscallId2Name[435] = "clone3"

    for id, name := range SyscallId2Name {
        SyscallName2Id[name] = id
    }
}
