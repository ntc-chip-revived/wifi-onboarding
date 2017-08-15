package main

import (
  "io"
  "log"
  "os"
)

var (
  // Set the debug mode
  debugMode = true

  Debug *log.Logger
  Info *log.Logger
  Warning *log.Logger
  Error *log.Logger
)

func initLoggers(
  debugHandle io.Writer,
  infoHandle io.Writer,
  warningHandle io.Writer,
  errorHandle io.Writer) {

  if (debugMode) {
    debugHandle = os.Stdout
  }

  Debug = log.New(debugHandle,
    "DEBUG: ",
    log.Ltime|log.Lshortfile)

  Info = log.New(infoHandle,
    "INFO: ",
    log.Ltime|log.Lshortfile)

  Warning = log.New(warningHandle,
    "WARNING: ",
    log.Ltime|log.Lshortfile)

  Error = log.New(errorHandle,
    "ERROR: ",
    log.Ltime|log.Lshortfile)

  Debug.Println("Initializing loggers")
}
