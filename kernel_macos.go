// Copyright 2018 Mikhail Klementev. All rights reserved.
// Use of this source code is governed by a AGPLv3 license
// (or later) that can be found in the LICENSE file.

// +build darwin

package main

import (
	"errors"

	"code.dumpstack.io/tools/out-of-tree/config"
)

func genHostKernels(download bool) (kcfg config.KernelConfig, err error) {
	err = errors.New("generate host kernels for macOS is not supported")
	return
}
