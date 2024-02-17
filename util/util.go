package util

import "strconv"


func ValidateFlags(flags []string) error {
	_, err := strconv.Atoi(flags[0])
	if err != nil {
		return err
	}

	return nil
}
