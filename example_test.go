package progress_test

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
