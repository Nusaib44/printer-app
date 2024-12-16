package platform

// import (
// 	"syscall"
// 	"unsafe"

// 	"golang.org/x/sys/windows"
// )

// // DiscoverPrintersWindows lists printers using Windows APIs.
// func DiscoverPrintersWindows() ([]string, error) {
// 	lib, err := syscall.LoadLibrary("winspool.drv")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer syscall.FreeLibrary(lib)

// 	proc, err := syscall.GetProcAddress(lib, "EnumPrintersW")
// 	if err != nil {
// 		return nil, err
// 	}

// 	type PRINTER_INFO_2 struct {
// 		pPrinterName *uint16
// 		// Other fields omitted for brevity
// 	}

// 	var needed, returned uint32
// 	windows.EnumPrinters(windows.PRINTER_ENUM_LOCAL, nil, 2, nil, 0, &needed, &returned)
// 	buffer := make([]byte, needed)
// 	windows.EnumPrinters(windows.PRINTER_ENUM_LOCAL, nil, 2, &buffer[0], needed, &needed, &returned)

// 	var printers []string
// 	count := int(returned)
// 	offset := uintptr(unsafe.Pointer(&buffer[0]))

// 	for i := 0; i < count; i++ {
// 		info := (*PRINTER_INFO_2)(unsafe.Pointer(offset))
// 		printers = append(printers, windows.UTF16PtrToString(info.pPrinterName))
// 		offset += unsafe.Sizeof(*info)
// 	}
// 	return printers, nil
// }
