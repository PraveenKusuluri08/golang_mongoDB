package helpers

type Error struct {
	IsError bool
	Message string
}

func ErrorHandler(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}
