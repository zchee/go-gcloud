// Copyright 2019 The go-gcloud Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/zchee/go-gcloud/pkg/gcloud"
)

var (
	file string
)

func init() {
	flag.StringVar(&file, "file", "", "json file")
}

func main() {
	flag.Parse()

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)

	var t gcloud.Gcloud
	if err := dec.Decode(&t); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("t.Commands[`access-context-manager`]: %s\n", spew.Sdump(t.Commands[`access-context-manager`]))
}
