package main

import (
	"bufio"
	"errors"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/tianrandailove/x-appender/src/level"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	appName          = "XAPPENDER"
	logNameFormatEnv = appName + "_LOG_NAME_FORMAT"
	logRootPath      = appName + "_LOG_ROOT_PATH"
	logMaxAge        = appName + "_LOG_MAX_AGE"
	logRotation      = appName + "LOG_ROTATION"
)

var logMap = make(map[string]*rotatelogs.RotateLogs)

func main() {
	path := os.Getenv(logRootPath)
	if path == "" {
		path = "./logs"
	}

	defer func() {
		for _, log := range logMap {
			log.Close()
		}
	}()

	inputScanner := bufio.NewScanner(os.Stdin)
	for inputScanner.Scan() {
		str := inputScanner.Text() + "\n"
		//stdout
		fmt.Print(str)
		//get current log level
		logLevel := level.GetLevelState(str)
		if log, ok := logMap[logLevel]; ok && log != nil {
			log.Write([]byte(str))
		} else {
			log, err := createLog(logLevel, path)
			if err != nil {
				panic(appName + " create log failed:" + err.Error())
			}
			logMap[logLevel] = log
			log.Write([]byte(str))
		}
	}
}

// create a log persistence object
func createLog(logLevel string, rootPath string) (*rotatelogs.RotateLogs, error) {
	logFileNameFormat := os.Getenv(logNameFormatEnv)
	logMaxAgeTime := os.Getenv(logMaxAge)
	logRotationTime := os.Getenv(logRotation)

	var maxAge time.Duration
	if logMaxAgeTime == "" {
		maxAge = 30 * 24 * time.Hour
	} else {
		if day, err := strconv.Atoi(logMaxAgeTime); err == nil {
			maxAge = time.Duration(day) * 24 * time.Hour
		} else {
			maxAge = 30 * 24 * time.Hour
		}
	}
	var rotationTime time.Duration
	if logRotationTime == "" {
		rotationTime = 24 * time.Hour
	} else {
		if hour, err := strconv.Atoi(logRotationTime); err == nil {
			rotationTime = time.Duration(hour) * time.Hour
		} else {
			rotationTime = 24 * time.Hour
		}
	}
	logFilePath := filepath.Join(rootPath, strings.ToLower(logLevel), strings.ToLower(logLevel)+"-"+logFileNameFormat+".log")
	linkPath := filepath.Join(rootPath, strings.ToLower(logLevel)+".log")
	log, err := rotatelogs.New(
		logFilePath,
		rotatelogs.WithLinkName(linkPath),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil || log == nil {
		return nil, errors.New("create log failed:" + err.Error())
	}
	return log, nil
}
