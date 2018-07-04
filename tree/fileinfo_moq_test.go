// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package tree

import (
	"os"
	"sync"
	"time"
)

var (
	lockosFileInfoMockIsDir   sync.RWMutex
	lockosFileInfoMockModTime sync.RWMutex
	lockosFileInfoMockMode    sync.RWMutex
	lockosFileInfoMockName    sync.RWMutex
	lockosFileInfoMockSize    sync.RWMutex
	lockosFileInfoMockSys     sync.RWMutex
)

// osFileInfoMock is a mock implementation of osFileInfo.
//
//     func TestSomethingThatUsesosFileInfo(t *testing.T) {
//
//         // make and configure a mocked osFileInfo
//         mockedosFileInfo := &osFileInfoMock{
//             IsDirFunc: func() bool {
// 	               panic("TODO: mock out the IsDir method")
//             },
//             ModTimeFunc: func() time.Time {
// 	               panic("TODO: mock out the ModTime method")
//             },
//             ModeFunc: func() os.FileMode {
// 	               panic("TODO: mock out the Mode method")
//             },
//             NameFunc: func() string {
// 	               panic("TODO: mock out the Name method")
//             },
//             SizeFunc: func() int64 {
// 	               panic("TODO: mock out the Size method")
//             },
//             SysFunc: func() interface{} {
// 	               panic("TODO: mock out the Sys method")
//             },
//         }
//
//         // TODO: use mockedosFileInfo in code that requires osFileInfo
//         //       and then make assertions.
//
//     }
type osFileInfoMock struct {
	// IsDirFunc mocks the IsDir method.
	IsDirFunc func() bool

	// ModTimeFunc mocks the ModTime method.
	ModTimeFunc func() time.Time

	// ModeFunc mocks the Mode method.
	ModeFunc func() os.FileMode

	// NameFunc mocks the Name method.
	NameFunc func() string

	// SizeFunc mocks the Size method.
	SizeFunc func() int64

	// SysFunc mocks the Sys method.
	SysFunc func() interface{}

	// calls tracks calls to the methods.
	calls struct {
		// IsDir holds details about calls to the IsDir method.
		IsDir []struct {
		}
		// ModTime holds details about calls to the ModTime method.
		ModTime []struct {
		}
		// Mode holds details about calls to the Mode method.
		Mode []struct {
		}
		// Name holds details about calls to the Name method.
		Name []struct {
		}
		// Size holds details about calls to the Size method.
		Size []struct {
		}
		// Sys holds details about calls to the Sys method.
		Sys []struct {
		}
	}
}

// IsDir calls IsDirFunc.
func (mock *osFileInfoMock) IsDir() bool {
	if mock.IsDirFunc == nil {
		panic("osFileInfoMock.IsDirFunc: method is nil but osFileInfo.IsDir was just called")
	}
	callInfo := struct {
	}{}
	lockosFileInfoMockIsDir.Lock()
	mock.calls.IsDir = append(mock.calls.IsDir, callInfo)
	lockosFileInfoMockIsDir.Unlock()
	return mock.IsDirFunc()
}

// IsDirCalls gets all the calls that were made to IsDir.
// Check the length with:
//     len(mockedosFileInfo.IsDirCalls())
func (mock *osFileInfoMock) IsDirCalls() []struct {
} {
	var calls []struct {
	}
	lockosFileInfoMockIsDir.RLock()
	calls = mock.calls.IsDir
	lockosFileInfoMockIsDir.RUnlock()
	return calls
}

// ModTime calls ModTimeFunc.
func (mock *osFileInfoMock) ModTime() time.Time {
	if mock.ModTimeFunc == nil {
		panic("osFileInfoMock.ModTimeFunc: method is nil but osFileInfo.ModTime was just called")
	}
	callInfo := struct {
	}{}
	lockosFileInfoMockModTime.Lock()
	mock.calls.ModTime = append(mock.calls.ModTime, callInfo)
	lockosFileInfoMockModTime.Unlock()
	return mock.ModTimeFunc()
}

// ModTimeCalls gets all the calls that were made to ModTime.
// Check the length with:
//     len(mockedosFileInfo.ModTimeCalls())
func (mock *osFileInfoMock) ModTimeCalls() []struct {
} {
	var calls []struct {
	}
	lockosFileInfoMockModTime.RLock()
	calls = mock.calls.ModTime
	lockosFileInfoMockModTime.RUnlock()
	return calls
}

// Mode calls ModeFunc.
func (mock *osFileInfoMock) Mode() os.FileMode {
	if mock.ModeFunc == nil {
		panic("osFileInfoMock.ModeFunc: method is nil but osFileInfo.Mode was just called")
	}
	callInfo := struct {
	}{}
	lockosFileInfoMockMode.Lock()
	mock.calls.Mode = append(mock.calls.Mode, callInfo)
	lockosFileInfoMockMode.Unlock()
	return mock.ModeFunc()
}

// ModeCalls gets all the calls that were made to Mode.
// Check the length with:
//     len(mockedosFileInfo.ModeCalls())
func (mock *osFileInfoMock) ModeCalls() []struct {
} {
	var calls []struct {
	}
	lockosFileInfoMockMode.RLock()
	calls = mock.calls.Mode
	lockosFileInfoMockMode.RUnlock()
	return calls
}

// Name calls NameFunc.
func (mock *osFileInfoMock) Name() string {
	if mock.NameFunc == nil {
		panic("osFileInfoMock.NameFunc: method is nil but osFileInfo.Name was just called")
	}
	callInfo := struct {
	}{}
	lockosFileInfoMockName.Lock()
	mock.calls.Name = append(mock.calls.Name, callInfo)
	lockosFileInfoMockName.Unlock()
	return mock.NameFunc()
}

// NameCalls gets all the calls that were made to Name.
// Check the length with:
//     len(mockedosFileInfo.NameCalls())
func (mock *osFileInfoMock) NameCalls() []struct {
} {
	var calls []struct {
	}
	lockosFileInfoMockName.RLock()
	calls = mock.calls.Name
	lockosFileInfoMockName.RUnlock()
	return calls
}

// Size calls SizeFunc.
func (mock *osFileInfoMock) Size() int64 {
	if mock.SizeFunc == nil {
		panic("osFileInfoMock.SizeFunc: method is nil but osFileInfo.Size was just called")
	}
	callInfo := struct {
	}{}
	lockosFileInfoMockSize.Lock()
	mock.calls.Size = append(mock.calls.Size, callInfo)
	lockosFileInfoMockSize.Unlock()
	return mock.SizeFunc()
}

// SizeCalls gets all the calls that were made to Size.
// Check the length with:
//     len(mockedosFileInfo.SizeCalls())
func (mock *osFileInfoMock) SizeCalls() []struct {
} {
	var calls []struct {
	}
	lockosFileInfoMockSize.RLock()
	calls = mock.calls.Size
	lockosFileInfoMockSize.RUnlock()
	return calls
}

// Sys calls SysFunc.
func (mock *osFileInfoMock) Sys() interface{} {
	if mock.SysFunc == nil {
		panic("osFileInfoMock.SysFunc: method is nil but osFileInfo.Sys was just called")
	}
	callInfo := struct {
	}{}
	lockosFileInfoMockSys.Lock()
	mock.calls.Sys = append(mock.calls.Sys, callInfo)
	lockosFileInfoMockSys.Unlock()
	return mock.SysFunc()
}

// SysCalls gets all the calls that were made to Sys.
// Check the length with:
//     len(mockedosFileInfo.SysCalls())
func (mock *osFileInfoMock) SysCalls() []struct {
} {
	var calls []struct {
	}
	lockosFileInfoMockSys.RLock()
	calls = mock.calls.Sys
	lockosFileInfoMockSys.RUnlock()
	return calls
}