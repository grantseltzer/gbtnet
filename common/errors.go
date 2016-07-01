package common

import "logrus"

// ErrorCheck checks errors
func ErrorCheck(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}

// ListenErrorCheck is an error check for listening microservices
func ListenErrorCheck(err error, projectName string) {
	logrus.Fatal("Interuptted Listen and Serve:  ", projectName, err)
}
