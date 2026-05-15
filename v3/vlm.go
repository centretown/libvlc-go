package vlc

/*
#cgo LDFLAGS: -lvlc
#include <vlc/vlc.h>
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

// var (
// 	ErrVlmBroadcast = errors.New("vlm could not broadcast")
// 	ErrVlmVOD       = errors.New("vlm could not provide video on demand")
// 	ErrVlmPlayMedia = errors.New("vlm could not play media")
// )

func makeOptions(args ...string) (argc int, argv []*C.char) {
	argc = len(args)
	argv = make([]*C.char, argc)
	for i, arg := range args {
		argv[i] = C.CString(arg)
	}
	return
}

func freeCStrings(argv ...*C.char) {
	for i := range argv {
		C.free(unsafe.Pointer(argv[i]))
	}
}

// use for branch free conditional
// hopefully this gets inlined (so far it does)
func booltoint(val bool) (i C.int) {
	if val {
		i = C.int(1)
	} else {
		i = C.int(0)
	}
	return i
}

func VlmRelease() {
	C.libvlc_vlm_release(inst.handle)
}

func VlmAddBroadcast(name string, input string, output string, enabled bool, loop bool, options ...string) error {
	optionc, optionv := makeOptions(options...)
	pszName, pszInput, pszOutput := C.CString(name), C.CString(input), C.CString(output)
	defer func() {
		freeCStrings(optionv...)
		freeCStrings(pszName, pszInput, pszOutput)
	}()

	if 0 != int(C.libvlc_vlm_add_broadcast(inst.handle, pszName, pszInput, pszOutput,
		C.int(optionc), *(***C.char)(unsafe.Pointer(&optionv)), booltoint(enabled), booltoint(loop))) {
		return getError()
	}
	return nil
}

func VlmAddVod(name string, input string, output string, enabled bool, mux string, options ...string) error {
	optionc, optionv := makeOptions(options...)
	pszName, pszInput, pszMux := C.CString(name), C.CString(output), C.CString(mux)
	defer func() {
		freeCStrings(optionv...)
		freeCStrings(pszName, pszInput, pszMux)
	}()

	if 0 != int(C.libvlc_vlm_add_vod(inst.handle, pszName, pszInput, C.int(optionc), *(***C.char)(unsafe.Pointer(&optionv)), booltoint(enabled), pszMux)) {
		return getError()
	}
	return nil
}

func VlmDeleteMedia(name string) error {
	pszName := C.CString(name)
	defer freeCStrings(pszName)
	if 0 != int(C.libvlc_vlm_del_media(inst.handle, pszName)) {
		return getError()
	}
	return nil
}

func VlmSetEnabled(name string, enabled bool) error {
	pszName := C.CString(name)
	defer freeCStrings(pszName)
	if 0 != int(C.libvlc_vlm_set_enabled(inst.handle, pszName, booltoint(enabled))) {
		return getError()
	}
	return nil
}

func VlmSetOutput(name string, output string) error {
	pszName, pszOutput := C.CString(name), C.CString(output)
	defer freeCStrings(pszName, pszOutput)
	if 0 != int(C.libvlc_vlm_set_output(inst.handle, pszName, pszOutput)) {
		return getError()
	}
	return nil
}

func VlmSetInput(name string, input string) error {
	pszName, pszInput := C.CString(name), C.CString(input)
	defer freeCStrings(pszName, pszInput)
	if 0 != int(C.libvlc_vlm_set_input(inst.handle, pszName, pszInput)) {
		return getError()
	}
	return nil
}

func VlmAddInput(name string, input string) error {
	pszName, pszInput := C.CString(name), C.CString(input)
	defer freeCStrings(pszName, pszInput)
	if 0 != int(C.libvlc_vlm_add_input(inst.handle, pszName, pszInput)) {
		return getError()
	}
	return nil
}

func VlmSetLoop(name string, loop bool) error {
	pszName := C.CString(name)
	defer freeCStrings(pszName)
	if 0 != int(C.libvlc_vlm_set_loop(inst.handle, pszName, booltoint(loop))) {
		return getError()
	}
	return nil
}

func VlmSetMux(name string, mux string) error {
	pszName, pszMux := C.CString(name), C.CString(mux)
	defer freeCStrings(pszName, pszMux)
	if 0 != int(C.libvlc_vlm_set_mux(inst.handle, pszName, pszMux)) {
		return getError()
	}
	return nil
}

func VlmChangeMedia(name string, input string, output string, enabled bool, loop bool, options ...string) error {
	optionc, optionv := makeOptions(options...)
	pszName, pszInput, pszOutput := C.CString(name), C.CString(input), C.CString(output)
	defer func() {
		freeCStrings(optionv...)
		freeCStrings(pszName, pszInput, pszOutput)
	}()

	if 0 != C.libvlc_vlm_change_media(inst.handle, pszName, pszInput, pszOutput, C.int(optionc), *(***C.char)(unsafe.Pointer(&optionv)), booltoint(enabled), booltoint(loop)) {
		return getError()
	}
	return nil
}

func VlmPlayMedia(name string) error {
	pszName := C.CString(name)
	defer freeCStrings(pszName)
	if 0 != int(C.libvlc_vlm_play_media(inst.handle, pszName)) {
		return getError()
	}
	return nil
}

func VlmStopMedia(name string) error {
	pszName := C.CString(name)
	defer freeCStrings(pszName)
	if 0 != int(C.libvlc_vlm_stop_media(inst.handle, pszName)) {
		return getError()
	}
	return nil
}

func VlmPauseMedia(name string) error {
	pszName := C.CString(name)
	defer freeCStrings(pszName)
	if 0 != int(C.libvlc_vlm_pause_media(inst.handle, pszName)) {
		return getError()
	}
	return nil
}

func VlmSeekMedia(name string, percent float32) error {
	pszName := C.CString(name)
	defer freeCStrings(pszName)
	if 0 != int(C.libvlc_vlm_seek_media(inst.handle, pszName, C.float(percent))) {
		return getError()
	}
	return nil
}

func VlmShowMedia(name string) (info string, err error) {
	pszName := C.CString(name)
	pszInfo := C.libvlc_vlm_show_media(inst.handle, pszName)
	defer freeCStrings(pszName, pszInfo)
	info = C.GoString(pszInfo)
	if len(info) == 0 {
		err = getError()
	}
	return
}

// Get vlm_media instance position by name or instance id
func VlmGetMediaInstancePosition(name string, media_instance int) (pos float32, err error) {
	pszName := C.CString(name)
	defer freeCStrings(pszName)
	pos = float32(C.libvlc_vlm_get_media_instance_position(inst.handle, pszName, C.int(media_instance)))
	if pos < 0.0 {
		err = getError()
	}
	return
}

func VlmGetMediaInstanceTime(name string, media_instance int) (media_time int, err error) {
	pszName := C.CString(name)
	defer freeCStrings(pszName)
	media_time = int(C.libvlc_vlm_get_media_instance_time(inst.handle, pszName, C.int(media_instance)))
	if media_time < 0 {
		err = getError()
	}
	return
}

func VlmGetMediaInstanceLength(name string, media_instance int) (media_length int, err error) {
	pszName := C.CString(name)
	defer freeCStrings(pszName)
	media_length = int(C.libvlc_vlm_get_media_instance_length(inst.handle, pszName, C.int(media_instance)))
	if media_length < 0 {
		err = getError()
	}
	return
}

func VlmGetMediaInstanceRate(name string, media_instance int) (media_rate int, err error) {
	pszName := C.CString(name)
	defer freeCStrings(pszName)
	media_rate = int(C.libvlc_vlm_get_media_instance_rate(inst.handle, pszName, C.int(media_instance)))
	if media_rate < 0 {
		err = getError()
	}
	return
}

func VlmGetEventManager() (mgr *EventManager) {
	event_manager := C.libvlc_vlm_get_event_manager(inst.handle)
	mgr = newEventManager(event_manager)
	return
}
