// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Contains the metrics collected by the fetcher.

package fetcher

import (
	"github.com/yooba-team/yooba/metrics"
)

var (
	propAnnounceInMeter   = metrics.NewRegisteredMeter("yoo/fetcher/prop/announces/in",nil)
	propAnnounceOutTimer  = metrics.NewRegisteredTimer("yoo/fetcher/prop/announces/out",nil)
	propAnnounceDropMeter = metrics.NewRegisteredMeter("yoo/fetcher/prop/announces/drop",nil)
	propAnnounceDOSMeter  = metrics.NewRegisteredMeter("yoo/fetcher/prop/announces/dos",nil)
	propBroadcastInMeter   = metrics.NewRegisteredMeter("yoo/fetcher/prop/broadcasts/in",nil)
	propBroadcastOutTimer  = metrics.NewRegisteredTimer("yoo/fetcher/prop/broadcasts/out",nil)
	propBroadcastDropMeter = metrics.NewRegisteredMeter("yoo/fetcher/prop/broadcasts/drop",nil)
	propBroadcastDOSMeter  = metrics.NewRegisteredMeter("yoo/fetcher/prop/broadcasts/dos",nil)
	headerFetchMeter = metrics.NewRegisteredMeter("yoo/fetcher/fetch/headers",nil)
	bodyFetchMeter   = metrics.NewRegisteredMeter("yoo/fetcher/fetch/bodies",nil)
	headerFilterInMeter  = metrics.NewRegisteredMeter("yoo/fetcher/filter/headers/in",nil)
	headerFilterOutMeter = metrics.NewRegisteredMeter("yoo/fetcher/filter/headers/out",nil)
	bodyFilterInMeter    = metrics.NewRegisteredMeter("yoo/fetcher/filter/bodies/in",nil)
	bodyFilterOutMeter   = metrics.NewRegisteredMeter("yoo/fetcher/filter/bodies/out",nil)


)
