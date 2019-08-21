package errors

import "regexp"

var (
	pipelineAlreadyExistsRegexp = regexp.MustCompile(`.*A pipeline with name .* already exists.*`)
)

// IsPipelineAlreadyExists returns true if the error indicates that a pipeline
// already exists.
func IsPipelineAlreadyExists(err error) bool {
	if err == nil {
		return false
	}

	return pipelineAlreadyExistsRegexp.MatchString(err.Error())
}
