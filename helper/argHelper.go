package arghelper

import "errors"

func GetArgValue(args []string, argumentName string) (string, error) {
	var value string
	var searchString = "-" + argumentName
	for pos, arg := range args {
		if arg == searchString {
			value = args[pos+1]
			break
		}
	}
	if value == "" {
		var errMsg = "Argument '" + argumentName + "' not found."
		return "", errors.New(errMsg)
	}
	return value, nil
}
