# progress

:smile: progress

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/moul.io/progress)
[![License](https://img.shields.io/badge/license-Apache--2.0%20%2F%20MIT-%2397ca00.svg)](https://github.com/moul/progress/blob/master/COPYRIGHT)
[![GitHub release](https://img.shields.io/github/release/moul/progress.svg)](https://github.com/moul/progress/releases)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)

[![Go](https://github.com/moul/progress/workflows/Go/badge.svg)](https://github.com/moul/progress/actions?query=workflow%3AGo)
[![PR](https://github.com/moul/progress/workflows/PR/badge.svg)](https://github.com/moul/progress/actions?query=workflow%3APR)
[![GolangCI](https://golangci.com/badges/github.com/moul/progress.svg)](https://golangci.com/r/github.com/moul/progress)
[![codecov](https://codecov.io/gh/moul/progress/branch/master/graph/badge.svg)](https://codecov.io/gh/moul/progress)
[![Go Report Card](https://goreportcard.com/badge/moul.io/progress)](https://goreportcard.com/report/moul.io/progress)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/progress/badge)](https://www.codefactor.io/repository/github/moul/progress)

[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/moul/progress)

## Usage

[embedmd]:# (example_test.go /import\ / $)
```go
import (
	"fmt"

	"moul.io/progress"
	"moul.io/u"
)

func Example() {
	// initialize a new progress.Progress
	prog := progress.New()
	prog.AddStep("init").SetDescription("initialize")
	prog.AddStep("step1").SetDescription("step 1")
	prog.AddStep("step2").SetData([]string{"hello", "world"}).SetDescription("step 2")
	prog.AddStep("step3")
	prog.AddStep("finish")

	// automatically mark the last step as done when the function quit
	defer prog.Get("finish").Done()

	// mark init as Done
	prog.Get("init").Done()

	// mark step1 as started
	prog.Get("step1").SetData(42).Start()

	// then, mark it as done + attach custom data
	prog.Get("step1").SetData(1337).Done()

	// mark step2 as started
	prog.Get("step2").Start()

	fmt.Println(u.PrettyJSON(prog))

	// outputs something like this:
	// {
	//  "steps": [
	//    {
	//      "id": "init",
	//      "description": "initialize",
	//      "started_at": "2020-12-22T20:26:05.717427484+01:00",
	//      "done_at": "2020-12-22T20:26:05.717427484+01:00",
	//      "state": "done"
	//    },
	//    {
	//      "id": "step1",
	//      "description": "step 1",
	//      "started_at": "2020-12-22T20:26:05.71742797+01:00",
	//      "done_at": "2020-12-22T20:26:05.717428258+01:00",
	//      "state": "done",
	//      "data": 1337,
	//      "duration": 286
	//    },
	//    {
	//      "id": "step2",
	//      "description": "step 2",
	//      "started_at": "2020-12-22T20:26:05.71742865+01:00",
	//      "state": "in progress",
	//      "data": [
	//        "hello",
	//        "world"
	//      ],
	//      "duration": 496251
	//    },
	//    {
	//      "id": "step3"
	//    },
	//    {
	//      "id": "finish"
	//    }
	//  ],
	//  "created_at": "2020-12-22T20:26:05.717423018+01:00",
	//  "snapshot": {
	//    "state": "in progress",
	//    "doing": "step 2",
	//    "not_started": 2,
	//    "in_progress": 1,
	//    "completed": 2,
	//    "total": 5,
	//    "percent": 50,
	//    "total_duration": 25935,
	//    "started_at": "2020-12-22T20:26:05.717427484+01:00"
	//  }
	//}
}
```

## Install

### Using go

```sh
go get moul.io/progress
```

## Contribute

![Contribute <3](https://raw.githubusercontent.com/moul/moul/master/contribute.gif)

I really welcome contributions.
Your input is the most precious material.
I'm well aware of that and I thank you in advance.
Everyone is encouraged to look at what they can do on their own scale;
no effort is too small.

Everything on contribution is sum up here: [CONTRIBUTING.md](./CONTRIBUTING.md)

### Contributors ✨

<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg)](#contributors)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="http://manfred.life"><img src="https://avatars1.githubusercontent.com/u/94029?v=4" width="100px;" alt=""/><br /><sub><b>Manfred Touron</b></sub></a><br /><a href="#maintenance-moul" title="Maintenance">🚧</a> <a href="https://github.com/moul/progress/commits?author=moul" title="Documentation">📖</a> <a href="https://github.com/moul/progress/commits?author=moul" title="Tests">⚠️</a> <a href="https://github.com/moul/progress/commits?author=moul" title="Code">💻</a></td>
    <td align="center"><a href="https://manfred.life/moul-bot"><img src="https://avatars1.githubusercontent.com/u/41326314?v=4" width="100px;" alt=""/><br /><sub><b>moul-bot</b></sub></a><br /><a href="#maintenance-moul-bot" title="Maintenance">🚧</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors)
specification. Contributions of any kind welcome!

### Stargazers over time

[![Stargazers over time](https://starchart.cc/moul/progress.svg)](https://starchart.cc/moul/progress)

## License

© 2020 [Manfred Touron](https://manfred.life)

Licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)
([`LICENSE-APACHE`](LICENSE-APACHE)) or the [MIT license](https://opensource.org/licenses/MIT)
([`LICENSE-MIT`](LICENSE-MIT)), at your option.
See the [`COPYRIGHT`](COPYRIGHT) file for more details.

`SPDX-License-Identifier: (Apache-2.0 OR MIT)`
