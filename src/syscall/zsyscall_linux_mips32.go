
package syscall

import "unsafe"


func open(path string, mode int, perm uint32) (fd int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := Syscall(SYS_OPEN, uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(perm))
	use(unsafe.Pointer(_p0))
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func openat(dirfd int, path string, flags int, mode uint32) (fd int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := Syscall6(SYS_OPENAT, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(flags), uintptr(mode), 0, 0)
	use(unsafe.Pointer(_p0))
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func pipe(p *[2]_C_int) (err error) {
	_, _, e1 := RawSyscall(SYS_PIPE, uintptr(unsafe.Pointer(p)), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func pipe2(p *[2]_C_int, flags int) (err error) {
	_, _, e1 := RawSyscall(SYS_PIPE2, uintptr(unsafe.Pointer(p)), uintptr(flags), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func utimes(path string, times *[2]Timeval) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_UTIMES, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(times)), 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func utimensat(dirfd int, path string, times *[2]Timespec) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_UTIMENSAT, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(times)))
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func futimesat(dirfd int, path *byte, times *[2]Timeval) (err error) {
	_, _, e1 := Syscall(SYS_FUTIMESAT, uintptr(dirfd), uintptr(unsafe.Pointer(path)), uintptr(unsafe.Pointer(times)))
	if e1 != 0 {
		err = e1
	}
	return
}


func Getcwd(buf []byte) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(buf) > 0 {
		_p0 = unsafe.Pointer(&buf[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall(SYS_GETCWD, uintptr(_p0), uintptr(len(buf)), 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func wait4(pid int, wstatus *_C_int, options int, rusage *Rusage) (wpid int, err error) {
	r0, _, e1 := Syscall6(SYS_WAIT4, uintptr(pid), uintptr(unsafe.Pointer(wstatus)), uintptr(options), uintptr(unsafe.Pointer(rusage)), 0, 0)
	wpid = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func ptrace(request int, pid int, addr uintptr, data uintptr) (err error) {
	_, _, e1 := Syscall6(SYS_PTRACE, uintptr(request), uintptr(pid), uintptr(addr), uintptr(data), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func reboot(magic1 uint, magic2 uint, cmd int, arg string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(arg)
	if err != nil {
		return
	}
	_, _, e1 := Syscall6(SYS_REBOOT, uintptr(magic1), uintptr(magic2), uintptr(cmd), uintptr(unsafe.Pointer(_p0)), 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func mount(source string, target string, fstype string, flags uintptr, data *byte) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(source)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(target)
	if err != nil {
		return
	}
	var _p2 *byte
	_p2, err = BytePtrFromString(fstype)
	if err != nil {
		return
	}
	_, _, e1 := Syscall6(SYS_MOUNT, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), uintptr(unsafe.Pointer(_p2)), uintptr(flags), uintptr(unsafe.Pointer(data)), 0)
	use(unsafe.Pointer(_p0))
	use(unsafe.Pointer(_p1))
	use(unsafe.Pointer(_p2))
	if e1 != 0 {
		err = e1
	}
	return
}


func Access(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_ACCESS, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func Acct(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_ACCT, uintptr(unsafe.Pointer(_p0)), 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func Adjtimex(buf *Timex) (state int, err error) {
	r0, _, e1 := Syscall(SYS_ADJTIMEX, uintptr(unsafe.Pointer(buf)), 0, 0)
	state = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Chdir(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_CHDIR, uintptr(unsafe.Pointer(_p0)), 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func Chmod(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_CHMOD, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func Chroot(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_CHROOT, uintptr(unsafe.Pointer(_p0)), 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func Close(fd int) (err error) {
	_, _, e1 := Syscall(SYS_CLOSE, uintptr(fd), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Creat(path string, mode uint32) (fd int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := Syscall(SYS_CREAT, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0)
	use(unsafe.Pointer(_p0))
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Dup(oldfd int) (fd int, err error) {
	r0, _, e1 := RawSyscall(SYS_DUP, uintptr(oldfd), 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Dup2(oldfd int, newfd int) (err error) {
	_, _, e1 := RawSyscall(SYS_DUP2, uintptr(oldfd), uintptr(newfd), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Dup3(oldfd int, newfd int, flags int) (err error) {
	_, _, e1 := RawSyscall(SYS_DUP3, uintptr(oldfd), uintptr(newfd), uintptr(flags))
	if e1 != 0 {
		err = e1
	}
	return
}


func EpollCreate(size int) (fd int, err error) {
	r0, _, e1 := RawSyscall(SYS_EPOLL_CREATE, uintptr(size), 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func EpollCreate1(flag int) (fd int, err error) {
	r0, _, e1 := RawSyscall(SYS_EPOLL_CREATE1, uintptr(flag), 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func EpollCtl(epfd int, op int, fd int, event *EpollEvent) (err error) {
	_, _, e1 := RawSyscall6(SYS_EPOLL_CTL, uintptr(epfd), uintptr(op), uintptr(fd), uintptr(unsafe.Pointer(event)), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(events) > 0 {
		_p0 = unsafe.Pointer(&events[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall6(SYS_EPOLL_WAIT, uintptr(epfd), uintptr(_p0), uintptr(len(events)), uintptr(msec), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Exit(code int) {
	Syscall(SYS_EXIT_GROUP, uintptr(code), 0, 0)
	return
}


func Faccessat(dirfd int, path string, mode uint32, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall6(SYS_FACCESSAT, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(flags), 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func Fallocate(fd int, mode uint32, off int64, len int64) (err error) {
	_, _, e1 := Syscall6(SYS_FALLOCATE, uintptr(fd), uintptr(mode), uintptr(off>>32), uintptr(off), uintptr(len>>32), uintptr(len))
	if e1 != 0 {
		err = e1
	}
	return
}


func Fchdir(fd int) (err error) {
	_, _, e1 := Syscall(SYS_FCHDIR, uintptr(fd), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Fchmod(fd int, mode uint32) (err error) {
	_, _, e1 := Syscall(SYS_FCHMOD, uintptr(fd), uintptr(mode), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Fchmodat(dirfd int, path string, mode uint32, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall6(SYS_FCHMODAT, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(flags), 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func Fchownat(dirfd int, path string, uid int, gid int, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall6(SYS_FCHOWNAT, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid), uintptr(flags), 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func fcntl(fd int, cmd int, arg int) (val int, err error) {
	r0, _, e1 := Syscall(SYS_FCNTL, uintptr(fd), uintptr(cmd), uintptr(arg))
	val = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Fdatasync(fd int) (err error) {
	_, _, e1 := Syscall(SYS_FDATASYNC, uintptr(fd), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Flock(fd int, how int) (err error) {
	_, _, e1 := Syscall(SYS_FLOCK, uintptr(fd), uintptr(how), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Fsync(fd int) (err error) {
	_, _, e1 := Syscall(SYS_FSYNC, uintptr(fd), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Getdents(fd int, buf []byte) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(buf) > 0 {
		_p0 = unsafe.Pointer(&buf[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall(SYS_GETDENTS64, uintptr(fd), uintptr(_p0), uintptr(len(buf)))
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Getpgid(pid int) (pgid int, err error) {
	r0, _, e1 := RawSyscall(SYS_GETPGID, uintptr(pid), 0, 0)
	pgid = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Getpgrp() (pid int) {
	r0, _, _ := RawSyscall(SYS_GETPGRP, 0, 0, 0)
	pid = int(r0)
	return
}


func Getpid() (pid int) {
	r0, _, _ := RawSyscall(SYS_GETPID, 0, 0, 0)
	pid = int(r0)
	return
}


func Getppid() (ppid int) {
	r0, _, _ := RawSyscall(SYS_GETPPID, 0, 0, 0)
	ppid = int(r0)
	return
}


func Getpriority(which int, who int) (prio int, err error) {
	r0, _, e1 := Syscall(SYS_GETPRIORITY, uintptr(which), uintptr(who), 0)
	prio = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Getrusage(who int, rusage *Rusage) (err error) {
	_, _, e1 := RawSyscall(SYS_GETRUSAGE, uintptr(who), uintptr(unsafe.Pointer(rusage)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Gettid() (tid int) {
	r0, _, _ := RawSyscall(SYS_GETTID, 0, 0, 0)
	tid = int(r0)
	return
}


func Getxattr(path string, attr string, dest []byte) (sz int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(attr)
	if err != nil {
		return
	}
	var _p2 unsafe.Pointer
	if len(dest) > 0 {
		_p2 = unsafe.Pointer(&dest[0])
	} else {
		_p2 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall6(SYS_GETXATTR, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), uintptr(_p2), uintptr(len(dest)), 0, 0)
	use(unsafe.Pointer(_p0))
	use(unsafe.Pointer(_p1))
	sz = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func InotifyAddWatch(fd int, pathname string, mask uint32) (watchdesc int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(pathname)
	if err != nil {
		return
	}
	r0, _, e1 := Syscall(SYS_INOTIFY_ADD_WATCH, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(mask))
	use(unsafe.Pointer(_p0))
	watchdesc = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func InotifyInit() (fd int, err error) {
	r0, _, e1 := RawSyscall(SYS_INOTIFY_INIT, 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func InotifyInit1(flags int) (fd int, err error) {
	r0, _, e1 := RawSyscall(SYS_INOTIFY_INIT1, uintptr(flags), 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func InotifyRmWatch(fd int, watchdesc uint32) (success int, err error) {
	r0, _, e1 := RawSyscall(SYS_INOTIFY_RM_WATCH, uintptr(fd), uintptr(watchdesc), 0)
	success = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Kill(pid int, sig Signal) (err error) {
	_, _, e1 := RawSyscall(SYS_KILL, uintptr(pid), uintptr(sig), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Klogctl(typ int, buf []byte) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(buf) > 0 {
		_p0 = unsafe.Pointer(&buf[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall(SYS_SYSLOG, uintptr(typ), uintptr(_p0), uintptr(len(buf)))
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Link(oldpath string, newpath string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(oldpath)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(newpath)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_LINK, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0)
	use(unsafe.Pointer(_p0))
	use(unsafe.Pointer(_p1))
	if e1 != 0 {
		err = e1
	}
	return
}


func Listxattr(path string, dest []byte) (sz int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 unsafe.Pointer
	if len(dest) > 0 {
		_p1 = unsafe.Pointer(&dest[0])
	} else {
		_p1 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall(SYS_LISTXATTR, uintptr(unsafe.Pointer(_p0)), uintptr(_p1), uintptr(len(dest)))
	use(unsafe.Pointer(_p0))
	sz = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Mkdir(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_MKDIR, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func Mkdirat(dirfd int, path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_MKDIRAT, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode))
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func Mknod(path string, mode uint32, dev int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_MKNOD, uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(dev))
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func Mknodat(dirfd int, path string, mode uint32, dev int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall6(SYS_MKNODAT, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(dev), 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func Nanosleep(time *Timespec, leftover *Timespec) (err error) {
	_, _, e1 := Syscall(SYS_NANOSLEEP, uintptr(unsafe.Pointer(time)), uintptr(unsafe.Pointer(leftover)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Pause() (err error) {
	_, _, e1 := Syscall(SYS_PAUSE, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func PivotRoot(newroot string, putold string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(newroot)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(putold)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_PIVOT_ROOT, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0)
	use(unsafe.Pointer(_p0))
	use(unsafe.Pointer(_p1))
	if e1 != 0 {
		err = e1
	}
	return
}


func prlimit(pid int, resource int, old *Rlimit, newlimit *Rlimit) (err error) {
	_, _, e1 := RawSyscall6(SYS_PRLIMIT64, uintptr(pid), uintptr(resource), uintptr(unsafe.Pointer(old)), uintptr(unsafe.Pointer(newlimit)), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func read(fd int, p []byte) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(p) > 0 {
		_p0 = unsafe.Pointer(&p[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall(SYS_READ, uintptr(fd), uintptr(_p0), uintptr(len(p)))
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Readlink(path string, buf []byte) (n int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 unsafe.Pointer
	if len(buf) > 0 {
		_p1 = unsafe.Pointer(&buf[0])
	} else {
		_p1 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall(SYS_READLINK, uintptr(unsafe.Pointer(_p0)), uintptr(_p1), uintptr(len(buf)))
	use(unsafe.Pointer(_p0))
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Removexattr(path string, attr string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(attr)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_REMOVEXATTR, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0)
	use(unsafe.Pointer(_p0))
	use(unsafe.Pointer(_p1))
	if e1 != 0 {
		err = e1
	}
	return
}


func Rename(oldpath string, newpath string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(oldpath)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(newpath)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_RENAME, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0)
	use(unsafe.Pointer(_p0))
	use(unsafe.Pointer(_p1))
	if e1 != 0 {
		err = e1
	}
	return
}


func Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(oldpath)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(newpath)
	if err != nil {
		return
	}
	_, _, e1 := Syscall6(SYS_RENAMEAT, uintptr(olddirfd), uintptr(unsafe.Pointer(_p0)), uintptr(newdirfd), uintptr(unsafe.Pointer(_p1)), 0, 0)
	use(unsafe.Pointer(_p0))
	use(unsafe.Pointer(_p1))
	if e1 != 0 {
		err = e1
	}
	return
}


func Rmdir(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_RMDIR, uintptr(unsafe.Pointer(_p0)), 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func Setdomainname(p []byte) (err error) {
	var _p0 unsafe.Pointer
	if len(p) > 0 {
		_p0 = unsafe.Pointer(&p[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := Syscall(SYS_SETDOMAINNAME, uintptr(_p0), uintptr(len(p)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Sethostname(p []byte) (err error) {
	var _p0 unsafe.Pointer
	if len(p) > 0 {
		_p0 = unsafe.Pointer(&p[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := Syscall(SYS_SETHOSTNAME, uintptr(_p0), uintptr(len(p)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Setpgid(pid int, pgid int) (err error) {
	_, _, e1 := RawSyscall(SYS_SETPGID, uintptr(pid), uintptr(pgid), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Setsid() (pid int, err error) {
	r0, _, e1 := RawSyscall(SYS_SETSID, 0, 0, 0)
	pid = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Settimeofday(tv *Timeval) (err error) {
	_, _, e1 := RawSyscall(SYS_SETTIMEOFDAY, uintptr(unsafe.Pointer(tv)), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Setpriority(which int, who int, prio int) (err error) {
	_, _, e1 := Syscall(SYS_SETPRIORITY, uintptr(which), uintptr(who), uintptr(prio))
	if e1 != 0 {
		err = e1
	}
	return
}


func Setxattr(path string, attr string, data []byte, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(attr)
	if err != nil {
		return
	}
	var _p2 unsafe.Pointer
	if len(data) > 0 {
		_p2 = unsafe.Pointer(&data[0])
	} else {
		_p2 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := Syscall6(SYS_SETXATTR, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), uintptr(_p2), uintptr(len(data)), uintptr(flags), 0)
	use(unsafe.Pointer(_p0))
	use(unsafe.Pointer(_p1))
	if e1 != 0 {
		err = e1
	}
	return
}


func Symlink(oldpath string, newpath string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(oldpath)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(newpath)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_SYMLINK, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0)
	use(unsafe.Pointer(_p0))
	use(unsafe.Pointer(_p1))
	if e1 != 0 {
		err = e1
	}
	return
}


func Sync() {
	Syscall(SYS_SYNC, 0, 0, 0)
	return
}


func Sysinfo(info *Sysinfo_t) (err error) {
	_, _, e1 := RawSyscall(SYS_SYSINFO, uintptr(unsafe.Pointer(info)), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Tee(rfd int, wfd int, len int, flags int) (n int64, err error) {
	r0, r1, e1 := Syscall6(SYS_TEE, uintptr(rfd), uintptr(wfd), uintptr(len), uintptr(flags), 0, 0)
	n = int64(int64(r0)<<32 | int64(r1))
	if e1 != 0 {
		err = e1
	}
	return
}


func Tgkill(tgid int, tid int, sig Signal) (err error) {
	_, _, e1 := RawSyscall(SYS_TGKILL, uintptr(tgid), uintptr(tid), uintptr(sig))
	if e1 != 0 {
		err = e1
	}
	return
}


func Times(tms *Tms) (ticks uintptr, err error) {
	r0, _, e1 := RawSyscall(SYS_TIMES, uintptr(unsafe.Pointer(tms)), 0, 0)
	ticks = uintptr(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Umask(mask int) (oldmask int) {
	r0, _, _ := RawSyscall(SYS_UMASK, uintptr(mask), 0, 0)
	oldmask = int(r0)
	return
}


func Uname(buf *Utsname) (err error) {
	_, _, e1 := RawSyscall(SYS_UNAME, uintptr(unsafe.Pointer(buf)), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Unlink(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_UNLINK, uintptr(unsafe.Pointer(_p0)), 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func Unlinkat(dirfd int, path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_UNLINKAT, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func Unmount(target string, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(target)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_UMOUNT2, uintptr(unsafe.Pointer(_p0)), uintptr(flags), 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func Unshare(flags int) (err error) {
	_, _, e1 := Syscall(SYS_UNSHARE, uintptr(flags), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Ustat(dev int, ubuf *Ustat_t) (err error) {
	_, _, e1 := Syscall(SYS_USTAT, uintptr(dev), uintptr(unsafe.Pointer(ubuf)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Utime(path string, buf *Utimbuf) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_UTIME, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(buf)), 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func write(fd int, p []byte) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(p) > 0 {
		_p0 = unsafe.Pointer(&p[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall(SYS_WRITE, uintptr(fd), uintptr(_p0), uintptr(len(p)))
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func exitThread(code int) (err error) {
	_, _, e1 := Syscall(SYS_EXIT, uintptr(code), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func readlen(fd int, p *byte, np int) (n int, err error) {
	r0, _, e1 := Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(p)), uintptr(np))
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func writelen(fd int, p *byte, np int) (n int, err error) {
	r0, _, e1 := Syscall(SYS_WRITE, uintptr(fd), uintptr(unsafe.Pointer(p)), uintptr(np))
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func munmap(addr uintptr, length uintptr) (err error) {
	_, _, e1 := Syscall(SYS_MUNMAP, uintptr(addr), uintptr(length), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Madvise(b []byte, advice int) (err error) {
	var _p0 unsafe.Pointer
	if len(b) > 0 {
		_p0 = unsafe.Pointer(&b[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := Syscall(SYS_MADVISE, uintptr(_p0), uintptr(len(b)), uintptr(advice))
	if e1 != 0 {
		err = e1
	}
	return
}


func Mprotect(b []byte, prot int) (err error) {
	var _p0 unsafe.Pointer
	if len(b) > 0 {
		_p0 = unsafe.Pointer(&b[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := Syscall(SYS_MPROTECT, uintptr(_p0), uintptr(len(b)), uintptr(prot))
	if e1 != 0 {
		err = e1
	}
	return
}


func Mlock(b []byte) (err error) {
	var _p0 unsafe.Pointer
	if len(b) > 0 {
		_p0 = unsafe.Pointer(&b[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := Syscall(SYS_MLOCK, uintptr(_p0), uintptr(len(b)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Munlock(b []byte) (err error) {
	var _p0 unsafe.Pointer
	if len(b) > 0 {
		_p0 = unsafe.Pointer(&b[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := Syscall(SYS_MUNLOCK, uintptr(_p0), uintptr(len(b)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Mlockall(flags int) (err error) {
	_, _, e1 := Syscall(SYS_MLOCKALL, uintptr(flags), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Munlockall() (err error) {
	_, _, e1 := Syscall(SYS_MUNLOCKALL, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, err error) {
	r0, _, e1 := Syscall(SYS_ACCEPT, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)))
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error) {
	r0, _, e1 := Syscall6(SYS_ACCEPT4, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), uintptr(flags), 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, _, e1 := Syscall(SYS_BIND, uintptr(s), uintptr(addr), uintptr(addrlen))
	if e1 != 0 {
		err = e1
	}
	return
}


func connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, _, e1 := Syscall(SYS_CONNECT, uintptr(s), uintptr(addr), uintptr(addrlen))
	if e1 != 0 {
		err = e1
	}
	return
}


func getgroups(n int, list *_Gid_t) (nn int, err error) {
	r0, _, e1 := RawSyscall(SYS_GETGROUPS, uintptr(n), uintptr(unsafe.Pointer(list)), 0)
	nn = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func setgroups(n int, list *_Gid_t) (err error) {
	_, _, e1 := RawSyscall(SYS_SETGROUPS, uintptr(n), uintptr(unsafe.Pointer(list)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error) {
	_, _, e1 := Syscall6(SYS_GETSOCKOPT, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(unsafe.Pointer(vallen)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error) {
	_, _, e1 := Syscall6(SYS_SETSOCKOPT, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(vallen), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func socket(domain int, typ int, proto int) (fd int, err error) {
	r0, _, e1 := RawSyscall(SYS_SOCKET, uintptr(domain), uintptr(typ), uintptr(proto))
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, _, e1 := RawSyscall(SYS_GETPEERNAME, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)))
	if e1 != 0 {
		err = e1
	}
	return
}


func getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, _, e1 := RawSyscall(SYS_GETSOCKNAME, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)))
	if e1 != 0 {
		err = e1
	}
	return
}


func recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(p) > 0 {
		_p0 = unsafe.Pointer(&p[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall6(SYS_RECVFROM, uintptr(fd), uintptr(_p0), uintptr(len(p)), uintptr(flags), uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(fromlen)))
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error) {
	var _p0 unsafe.Pointer
	if len(buf) > 0 {
		_p0 = unsafe.Pointer(&buf[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := Syscall6(SYS_SENDTO, uintptr(s), uintptr(_p0), uintptr(len(buf)), uintptr(flags), uintptr(to), uintptr(addrlen))
	if e1 != 0 {
		err = e1
	}
	return
}


func socketpair(domain int, typ int, flags int, fd *[2]int32) (err error) {
	_, _, e1 := RawSyscall6(SYS_SOCKETPAIR, uintptr(domain), uintptr(typ), uintptr(flags), uintptr(unsafe.Pointer(fd)), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func recvmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, _, e1 := Syscall(SYS_RECVMSG, uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags))
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func sendmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, _, e1 := Syscall(SYS_SENDMSG, uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags))
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Chown(path string, uid int, gid int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_CHOWN, uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid))
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func Fchown(fd int, uid int, gid int) (err error) {
	_, _, e1 := Syscall(SYS_FCHOWN, uintptr(fd), uintptr(uid), uintptr(gid))
	if e1 != 0 {
		err = e1
	}
	return
}


func Fstat(fd int, stat *Stat_t) (err error) {
	_, _, e1 := Syscall(SYS_FSTAT64, uintptr(fd), uintptr(unsafe.Pointer(stat)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Getegid() (egid int) {
	r0, _, _ := RawSyscall(SYS_GETEGID, 0, 0, 0)
	egid = int(r0)
	return
}


func Geteuid() (euid int) {
	r0, _, _ := RawSyscall(SYS_GETEUID, 0, 0, 0)
	euid = int(r0)
	return
}


func Getgid() (gid int) {
	r0, _, _ := RawSyscall(SYS_GETGID, 0, 0, 0)
	gid = int(r0)
	return
}


func Getuid() (uid int) {
	r0, _, _ := RawSyscall(SYS_GETUID, 0, 0, 0)
	uid = int(r0)
	return
}


func Lchown(path string, uid int, gid int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_LCHOWN, uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid))
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func Listen(s int, n int) (err error) {
	_, _, e1 := Syscall(SYS_LISTEN, uintptr(s), uintptr(n), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Lstat(path string, stat *Stat_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_LSTAT64, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)), 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) {
	r0, _, e1 := Syscall6(SYS_SENDFILE64, uintptr(outfd), uintptr(infd), uintptr(unsafe.Pointer(offset)), uintptr(count), 0, 0)
	written = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error) {
	r0, _, e1 := Syscall6(SYS__NEWSELECT, uintptr(nfd), uintptr(unsafe.Pointer(r)), uintptr(unsafe.Pointer(w)), uintptr(unsafe.Pointer(e)), uintptr(unsafe.Pointer(timeout)), 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Setfsgid(gid int) (err error) {
	_, _, e1 := Syscall(SYS_SETFSGID, uintptr(gid), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Setfsuid(uid int) (err error) {
	_, _, e1 := Syscall(SYS_SETFSUID, uintptr(uid), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Setregid(rgid int, egid int) (err error) {
	_, _, e1 := RawSyscall(SYS_SETREGID, uintptr(rgid), uintptr(egid), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Setresgid(rgid int, egid int, sgid int) (err error) {
	_, _, e1 := RawSyscall(SYS_SETRESGID, uintptr(rgid), uintptr(egid), uintptr(sgid))
	if e1 != 0 {
		err = e1
	}
	return
}


func Setresuid(ruid int, euid int, suid int) (err error) {
	_, _, e1 := RawSyscall(SYS_SETRESUID, uintptr(ruid), uintptr(euid), uintptr(suid))
	if e1 != 0 {
		err = e1
	}
	return
}


func Setreuid(ruid int, euid int) (err error) {
	_, _, e1 := RawSyscall(SYS_SETREUID, uintptr(ruid), uintptr(euid), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Shutdown(fd int, how int) (err error) {
	_, _, e1 := Syscall(SYS_SHUTDOWN, uintptr(fd), uintptr(how), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int, err error) {
	r0, _, e1 := Syscall6(SYS_SPLICE, uintptr(rfd), uintptr(unsafe.Pointer(roff)), uintptr(wfd), uintptr(unsafe.Pointer(woff)), uintptr(len), uintptr(flags))
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Stat(path string, stat *Stat_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_STAT64, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)), 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func Gettimeofday(tv *Timeval) (err error) {
	_, _, e1 := RawSyscall(SYS_GETTIMEOFDAY, uintptr(unsafe.Pointer(tv)), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Time(t *Time_t) (tt Time_t, err error) {
	r0, _, e1 := RawSyscall(SYS_TIME, uintptr(unsafe.Pointer(t)), 0, 0)
	tt = Time_t(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Pread(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(p) > 0 {
		_p0 = unsafe.Pointer(&p[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall6(SYS_PREAD64, uintptr(fd), uintptr(_p0), uintptr(len(p)), 0, uintptr(offset>>32), uintptr(offset))
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Pwrite(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(p) > 0 {
		_p0 = unsafe.Pointer(&p[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall6(SYS_PWRITE64, uintptr(fd), uintptr(_p0), uintptr(len(p)), 0, uintptr(offset>>32), uintptr(offset))
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func Truncate(path string, length int64) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall6(SYS_TRUNCATE64, uintptr(unsafe.Pointer(_p0)), 0, uintptr(length>>32), uintptr(length), 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}


func Ftruncate(fd int, length int64) (err error) {
	_, _, e1 := Syscall6(SYS_FTRUNCATE64, uintptr(fd), 0, uintptr(length>>32), uintptr(length), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func mmap2(addr uintptr, length uintptr, prot int, flags int, fd int, pageOffset uintptr) (xaddr uintptr, err error) {
	r0, _, e1 := Syscall6(SYS_MMAP2, uintptr(addr), uintptr(length), uintptr(prot), uintptr(flags), uintptr(fd), uintptr(pageOffset))
	xaddr = uintptr(r0)
	if e1 != 0 {
		err = e1
	}
	return
}


func getrlimit(resource int, rlim *rlimit32) (err error) {
	_, _, e1 := RawSyscall(SYS_GETRLIMIT, uintptr(resource), uintptr(unsafe.Pointer(rlim)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}


func setrlimit(resource int, rlim *rlimit32) (err error) {
	_, _, e1 := RawSyscall(SYS_SETRLIMIT, uintptr(resource), uintptr(unsafe.Pointer(rlim)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}
