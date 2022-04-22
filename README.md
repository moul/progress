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

[embedmd]:# (.tmp/godoc.txt txt /TYPES/ $)
```txt
TYPES

type Progress struct {
	Steps     []*Step   `json:"steps,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Has unexported fields.
}
    Progress is the top-level object of the 'progress' library.

func New() *Progress
    New creates and returns a new Progress.

func (p *Progress) AddStep(id string) *Step
    AddStep creates and returns a new Step with the provided 'id'. A non-empty,
    unique 'id' is required, else it will panic.

func (p *Progress) Close()
    Close cleans up the allocated ressources.

func (p *Progress) Get(id string) *Step
    Get retrieves a Step by its 'id'. A non-empty 'id' is required, else it will
    panic. If 'id' does not match an existing step, nil is returned.

func (p *Progress) MarshalJSON() ([]byte, error)
    MarshalJSON is a custom JSON marshaler that automatically computes and
    append the current snapshot.

func (p *Progress) Progress() float64
    Progress returns the current completion rate, it's a faster alternative to
    Progress.Snapshot().Progress. The returned value is between 0.0 and 1.0.

func (p *Progress) SafeAddStep(id string) (*Step, error)
    SafeAddStep is equivalent to AddStep with but returns error instead of
    panicking.

func (p *Progress) Snapshot() Snapshot
    Snapshot computes and returns the current stats of the Progress.

func (p *Progress) Subscribe() chan *Step
    Subscribe registers the provided chan as a target called each time a step is
    changed.

type Snapshot struct {
	State              State         `json:"state,omitempty"`
	Doing              string        `json:"doing,omitempty"`
	NotStarted         int           `json:"not_started,omitempty"`
	InProgress         int           `json:"in_progress,omitempty"`
	Completed          int           `json:"completed,omitempty"`
	Total              int           `json:"total,omitempty"`
	Progress           float64       `json:"progress,omitempty"`
	TotalDuration      time.Duration `json:"total_duration,omitempty"`
	StepDuration       time.Duration `json:"step_duration,omitempty"`
	CompletionEstimate time.Duration `json:"completion_estimate,omitempty"`
	DoneAt             *time.Time    `json:"done_at,omitempty"`
	StartedAt          *time.Time    `json:"started_at,omitempty"`
}
    Snapshot represents info and stats about a progress at a given time.

type State string

const (
	StateNotStarted State = "not started"
	StateInProgress State = "in progress"
	StateDone       State = "done"
	StateStopped    State = "stopped"
)
type Step struct {
	ID          string      `json:"id,omitempty"`
	Description string      `json:"description,omitempty"`
	StartedAt   *time.Time  `json:"started_at,omitempty"`
	DoneAt      *time.Time  `json:"done_at,omitempty"`
	State       State       `json:"state,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	Progress    float64     `json:"progress,omitempty"`
	Child       *Progress   `json:"child,omitempty"`

	// Has unexported fields.
}
    Step represents a progress step. It always have an 'id' and can be
    customized using helpers.

func (s *Step) Done() *Step
    Done marks a step as done. If the step was already done, it panics.

func (s *Step) Duration() time.Duration
    Duration computes the step duration.

func (s *Step) MarshalJSON() ([]byte, error)
    MarshalJSON is a custom JSON marshaler that automatically computes and
    append some runtime metadata.

func (s *Step) SetAsCurrent() *Step
    SetAsCurrent stops all in-progress steps and start this one.

func (s *Step) SetChild(prog *Progress) *Step
    SetChild configures a dedicated Progress on the Step

func (s *Step) SetData(data interface{}) *Step
    SetData sets a custom step data. It returns itself (*Step) for chaining.

func (s *Step) SetDescription(desc string) *Step
    SetDescription sets a custom step description. It returns itself (*Step) for
    chaining.

func (s *Step) SetProgress(progress float64) *Step
    SetProgress sets the current step progress rate. It may also update the
    current Step.State depending on the passed progress. The value should be
    something between 0.0 and 1.0.

func (s *Step) Start() *Step
    Start marks a step as started. If a step was already InProgress or Done, it
    panics.

```

## Example

[embedmd]:# (example_test.go /import\ / $)
```go
import (
	"fmt"
	"time"

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

func ExampleProgressSubscribe() {
	prog := progress.New()
	defer prog.Close()
	done := make(chan bool)
	ch := prog.Subscribe()

	go func() {
		idx := 0
		for step := range ch {
			if step == nil {
				break
			}
			fmt.Println(idx, step.ID, step.State)
			idx++
		}
		done <- true
	}()
	time.Sleep(10 * time.Millisecond)
	prog.AddStep("step1").SetDescription("hello")
	prog.AddStep("step2")
	prog.Get("step1").Start()
	prog.Get("step2").Done()
	prog.AddStep("step3")
	prog.Get("step3").Start()
	prog.Get("step1").Done()
	prog.AddStep("step4")
	prog.Get("step3").Done()
	prog.Get("step4").SetAsCurrent()
	prog.Get("step4").Done()
	// fmt.Println(u.PrettyJSON(prog))
	<-done

	// Output:
	// 0 step1 not started
	// 1 step1 not started
	// 2 step2 not started
	// 3 step1 in progress
	// 4 step2 done
	// 5 step3 not started
	// 6 step3 in progress
	// 7 step1 done
	// 8 step4 not started
	// 9 step3 done
	// 10 step4 in progress
	// 11 step4 done
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
