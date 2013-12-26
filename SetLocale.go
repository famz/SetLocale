package SetLocale

// #include <locale.h>
import (
    "C"
)

const (
    LC_ALL      = 0
    LC_COLLATE  = 1
    LC_CTYPE    = 2
    LC_MONETARY = 3
    LC_NUMERIC  = 5
    LC_TIME     = 6
)

func SetLocale(lc C.int, locale string) string {
    param := C.CString(locale)
    ret := C.setlocale(lc, param)
    return C.GoString(ret)
}
