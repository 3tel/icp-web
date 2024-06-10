// Copyright 2024 3tel. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package static

import (
	"embed"
)

//go:embed css/* js/*
var FS embed.FS
