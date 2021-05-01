// IRIS Endpoint-Server (EPS)
// Copyright (C) 2021-2021 The IRIS Endpoint-Server Authors (see AUTHORS.md)
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"github.com/iris-gateway/eps"
	"github.com/iris-gateway/eps/cmd/helpers"
	"github.com/iris-gateway/eps/definitions"
)

func main() {
	if settings, err := helpers.Settings(&definitions.Default); err != nil {
		eps.Log.Error(err)
		return
	} else {
		helpers.CLI(&definitions.Default, settings)
	}
}
