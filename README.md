About t32
===

Go wrapper for Lauterbach Trace32 Remote API.

It is based on the API shipped with Trace32 S.2012.0y.000038005 Build 38005

Requirements:
--------------
* C version of Trace32 Remote API to compile this module. For copyright issue, the files not included.
* MinGW (32bit / 64bit depends on your Go version) installed, needed by cgo to compile C code.

Install
--------------
* Assume your Go is installed at **c:\\go**, Download t32.go and put into folder:
**c:\\go\\src\\pkg\\github.com\\yongzhy\\t32**

* Copy from c:\\T32\\demo\\api\\capi\\src the following three files, put them in the same folder as t32.go

        hlinknet.c
        hremote.c
        t32.h

* From MinGW command prompt, issue following command :

        go install github.com/yongzhy/t32


Example
--------------
Start Trace32 with following configuration:

    ; Remote API Access
    RCL=NETASSIST
    PACKLEN=1024
    PORT=2000

Example Code:

	if err := t32.Config("NODE=", "localhost"); err != nil {
		panic(err)
	}
	if err := t32.Config("PACKLEN=", "1024"); err != nil {
		panic(err)
	}
	if err := t32.Config("PORT=", "2000"); err != nil {
		panic(err)
	}
	if err := t32.Init(); err != nil {
		panic(err)
	}
	if err := t32.Attach(t32.T32_DEV_ICD); err != nil {
		panic(err)
	}
    if err := t32.Cmd("Print \"Hello\"); err != nil {
        panic(err)
    }
    t32.Exit()