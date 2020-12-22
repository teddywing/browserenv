package browserenv

import "fmt"

const errorPrefix = "browserenv: "

// CopyError represents a failure to copy data.
type CopyError struct {
	err error
}

func (e *CopyError) Error() string {
	return fmt.Sprintf(errorPrefix+"can't copy from reader: %v", e.err)
}

func (e *CopyError) Unwrap() error { return e.err }

// ExecError results from executing an external command.
type ExecError struct {
	command string
	err     error
}

func (e *ExecError) Error() string {
	return fmt.Sprintf(
		errorPrefix+"failed to run BROWSER command %q: %v",
		e.command,
		e.err,
	)
}

func (e *ExecError) Unwrap() error { return e.err }

// ExecZeroError results from executing an external command that produces an
// error with an exit code of 0.
type ExecZeroError struct {
	command string
	err     error
}

func (e *ExecZeroError) Error() string {
	return fmt.Sprintf(
		errorPrefix+"error running command %q: %v",
		e.command,
		e.err,
	)
}

func (e *ExecZeroError) Unwrap() error { return e.err }

// PathResolutionError means the given path couldn't be resolved.
type PathResolutionError struct {
	path string
	err  error
}

func (e *PathResolutionError) Error() string {
	return fmt.Sprintf(
		errorPrefix+"can't resolve path for %q: %v",
		e.path,
		e.err,
	)
}

func (e *PathResolutionError) Unwrap() error { return e.err }

// TempFileError corresponds to an error while creating a temporary file.
type TempFileError struct {
	err error
}

func (e *TempFileError) Error() string {
	return fmt.Sprintf(errorPrefix+"can't create temporary file: %v", e.err)
}

func (e *TempFileError) Unwrap() error { return e.err }
