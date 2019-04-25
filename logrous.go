package main

import (
  "github.com/sirupsen/logrus"
)


func main() {

  log := logrus.New()

  log.SetLevel(logrus.TraceLevel)

  log.Trace("I love lu")
  log.Debug("I love lu")
  log.Info("I love lu")
  log.Warn("I love lu")


}