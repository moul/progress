package progress_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"moul.io/progress"
)

func TestFlow(t *testing.T) {
	// initialize a new progress
	prog := progress.New()
	{
		require.NotEmpty(t, prog)
		require.Empty(t, prog.Steps)
		require.NotZero(t, prog.CreatedAt)
		require.True(t, prog.CreatedAt.Before(time.Now()))
		snapshot := prog.Snapshot()
		require.Equal(t, progress.StateNotStarted, snapshot.State)
		require.Equal(t, 0, snapshot.Total)
		require.Equal(t, 0, snapshot.Completed)
		require.Equal(t, 0, snapshot.NotStarted)
		require.Equal(t, 0, snapshot.InProgress)
		require.Equal(t, float64(0), snapshot.Progress)
		require.Equal(t, snapshot.Progress, prog.Progress())
		require.Nil(t, prog.Get("step1"))
	}

	// add a first step
	{
		prog.AddStep("step1")
		require.NotEmpty(t, prog.Steps)
		require.Len(t, prog.Steps, 1)
		require.True(t, prog.CreatedAt.Before(time.Now()))

		snapshot := prog.Snapshot()
		require.Equal(t, "", snapshot.Doing)
		require.Equal(t, progress.StateNotStarted, snapshot.State)
		require.Equal(t, 1, snapshot.Total)
		require.Equal(t, 0, snapshot.Completed)
		require.Equal(t, 1, snapshot.NotStarted)
		require.Equal(t, 0, snapshot.InProgress)
		require.NotNil(t, prog.Get("step1"))
		require.Equal(t, float64(0), snapshot.Progress)
		require.Equal(t, snapshot.Progress, prog.Progress())

		step1 := prog.Get("step1")
		require.NotNil(t, step1)
		require.Equal(t, step1.State, progress.StateNotStarted)
		require.Empty(t, step1.Description)
		step1.SetDescription("hello")
		require.Equal(t, "hello", step1.Description)
	}

	// add a second step
	{
		prog.AddStep("step2")
		require.NotEmpty(t, prog.Steps)
		require.Len(t, prog.Steps, 2)

		snapshot := prog.Snapshot()
		require.Equal(t, "", snapshot.Doing)
		require.Equal(t, progress.StateNotStarted, snapshot.State)
		require.Equal(t, 2, snapshot.Total)
		require.Equal(t, 0, snapshot.Completed)
		require.Equal(t, 2, snapshot.NotStarted)
		require.Equal(t, 0, snapshot.InProgress)
		require.NotNil(t, prog.Get("step2"))
		require.Equal(t, float64(0), snapshot.Progress)
		require.Equal(t, snapshot.Progress, prog.Progress())
	}

	// start the first step
	{
		step1 := prog.Get("step1")
		step1.Start()

		require.NotEmpty(t, prog.Steps)
		require.Len(t, prog.Steps, 2)

		snapshot := prog.Snapshot()
		require.Equal(t, "hello", snapshot.Doing)
		require.Equal(t, progress.StateInProgress, snapshot.State)
		require.Equal(t, 2, snapshot.Total)
		require.Equal(t, 0, snapshot.Completed)
		require.Equal(t, 1, snapshot.NotStarted)
		require.Equal(t, 1, snapshot.InProgress)
		require.Equal(t, float64(0.25), snapshot.Progress)
		require.Equal(t, snapshot.Progress, prog.Progress())
	}

	// mark the first step as done
	{
		time.Sleep(200 * time.Millisecond)
		step1 := prog.Get("step1")
		step1.Done()
		require.Equal(t, progress.StateDone, step1.State)

		require.NotEmpty(t, prog.Steps)
		require.Len(t, prog.Steps, 2)

		snapshot := prog.Snapshot()
		require.Equal(t, "", snapshot.Doing)
		require.Equal(t, progress.StateInProgress, snapshot.State)
		require.Equal(t, 2, snapshot.Total)
		require.Equal(t, 1, snapshot.Completed)
		require.Equal(t, 1, snapshot.NotStarted)
		require.Equal(t, 0, snapshot.InProgress)
		require.Equal(t, float64(0.5), snapshot.Progress)
		require.Equal(t, snapshot.Progress, prog.Progress())
	}

	// mark the second step as done without starting it first
	{
		step2 := prog.Get("step2")
		step2.Done()
		require.Equal(t, progress.StateDone, step2.State)

		require.NotEmpty(t, prog.Steps)
		require.Len(t, prog.Steps, 2)

		snapshot := prog.Snapshot()
		require.Equal(t, "", snapshot.Doing)
		require.Equal(t, progress.StateDone, snapshot.State)
		require.Equal(t, 2, snapshot.Total)
		require.Equal(t, 2, snapshot.Completed)
		require.Equal(t, 0, snapshot.NotStarted)
		require.Equal(t, 0, snapshot.InProgress)
		require.Equal(t, float64(1), snapshot.Progress)
		require.Equal(t, snapshot.Progress, prog.Progress())
	}

	// add a third step
	{
		prog.AddStep("step3")
		require.NotEmpty(t, prog.Steps)
		require.Len(t, prog.Steps, 3)
		require.NotNil(t, prog.Get("step3"))

		snapshot := prog.Snapshot()
		require.Equal(t, "", snapshot.Doing)
		require.Equal(t, progress.StateInProgress, snapshot.State)
		require.Equal(t, 3, snapshot.Total)
		require.Equal(t, 2, snapshot.Completed)
		require.Equal(t, 1, snapshot.NotStarted)
		require.Equal(t, 0, snapshot.InProgress)
		require.Equal(t, 66, int(snapshot.Progress*100))
		require.Equal(t, snapshot.Progress, prog.Progress())
	}

	// add a fourth step
	{
		prog.AddStep("step4")
		require.NotEmpty(t, prog.Steps)
		require.Len(t, prog.Steps, 4)
		require.NotNil(t, prog.Get("step4"))

		snapshot := prog.Snapshot()
		require.Equal(t, "", snapshot.Doing)
		require.Equal(t, progress.StateInProgress, snapshot.State)
		require.Equal(t, 4, snapshot.Total)
		require.Equal(t, 2, snapshot.Completed)
		require.Equal(t, 2, snapshot.NotStarted)
		require.Equal(t, 0, snapshot.InProgress)
		require.Equal(t, float64(0.5), snapshot.Progress)
		require.Equal(t, snapshot.Progress, prog.Progress())
	}

	// start step3 and step4 at the same time
	{
		step3 := prog.Get("step3")
		step4 := prog.Get("step4")
		step3.Start()
		step4.Start()

		snapshot := prog.Snapshot()
		require.Equal(t, "step3, step4", snapshot.Doing)
		require.Equal(t, progress.StateInProgress, snapshot.State)
		require.Equal(t, 4, snapshot.Total)
		require.Equal(t, 2, snapshot.Completed)
		require.Equal(t, 0, snapshot.NotStarted)
		require.Equal(t, 2, snapshot.InProgress)
		require.Equal(t, float64(0.75), snapshot.Progress)
		require.Equal(t, snapshot.Progress, prog.Progress())
	}

	// mark step3 and step4 as done at the same time
	{
		time.Sleep(200 * time.Millisecond)
		step1 := prog.Get("step1")
		step2 := prog.Get("step2")
		step3 := prog.Get("step3")
		step4 := prog.Get("step4")
		step3.Done()
		step4.Done()

		snapshot := prog.Snapshot()
		require.Equal(t, "", snapshot.Doing)
		require.Equal(t, progress.StateDone, snapshot.State)
		require.Equal(t, 4, snapshot.Total)
		require.Equal(t, 4, snapshot.Completed)
		require.Equal(t, 0, snapshot.NotStarted)
		require.Equal(t, 0, snapshot.InProgress)
		require.Equal(t, float64(1), snapshot.Progress)
		require.Equal(t, snapshot.Progress, prog.Progress())

		require.True(t, step1.Duration() > 200*time.Millisecond && step1.Duration() < 400*time.Millisecond)
		require.Zero(t, step2.Duration())
		require.True(t, step3.Duration() > 200*time.Millisecond && step3.Duration() < 400*time.Millisecond)
		require.True(t, step4.Duration() > 200*time.Millisecond && step4.Duration() < 400*time.Millisecond)
		require.True(t, snapshot.TotalDuration > 400*time.Millisecond && snapshot.TotalDuration < 600*time.Millisecond)
	}

	// create a new step and use SetProgress instead of Start
	{
		prog.AddStep("step5")
		require.NotEmpty(t, prog.Steps)
		require.Len(t, prog.Steps, 5)
		require.NotNil(t, prog.Get("step5"))

		snapshot := prog.Snapshot()
		require.Equal(t, "", snapshot.Doing)
		require.Equal(t, progress.StateInProgress, snapshot.State)
		require.Equal(t, 5, snapshot.Total)
		require.Equal(t, 4, snapshot.Completed)
		require.Equal(t, 1, snapshot.NotStarted)
		require.Equal(t, 0, snapshot.InProgress)
		require.Equal(t, float64(0.8), snapshot.Progress)
		require.Equal(t, snapshot.Progress, prog.Progress())

		prog.Get("step5").SetProgress(0)
		snapshot = prog.Snapshot()
		require.Equal(t, "", snapshot.Doing)
		require.Equal(t, progress.StateInProgress, snapshot.State)
		require.Equal(t, 5, snapshot.Total)
		require.Equal(t, 4, snapshot.Completed)
		require.Equal(t, 1, snapshot.NotStarted)
		require.Equal(t, 0, snapshot.InProgress)
		require.Equal(t, float64(0.8), snapshot.Progress)
		require.Equal(t, snapshot.Progress, prog.Progress())

		prog.Get("step5").SetProgress(0.2)
		snapshot = prog.Snapshot()
		require.Equal(t, "step5", snapshot.Doing)
		require.Equal(t, progress.StateInProgress, snapshot.State)
		require.Equal(t, 5, snapshot.Total)
		require.Equal(t, 4, snapshot.Completed)
		require.Equal(t, 0, snapshot.NotStarted)
		require.Equal(t, 1, snapshot.InProgress)
		require.Equal(t, 84, int(snapshot.Progress*100))
		require.Equal(t, snapshot.Progress, prog.Progress())

		prog.Get("step5").SetProgress(0.8)
		snapshot = prog.Snapshot()
		require.Equal(t, "step5", snapshot.Doing)
		require.Equal(t, progress.StateInProgress, snapshot.State)
		require.Equal(t, 5, snapshot.Total)
		require.Equal(t, 4, snapshot.Completed)
		require.Equal(t, 0, snapshot.NotStarted)
		require.Equal(t, 1, snapshot.InProgress)
		require.Equal(t, 96, int(snapshot.Progress*100))
		require.Equal(t, snapshot.Progress, prog.Progress())

		prog.Get("step5").SetProgress(1.0)
		snapshot = prog.Snapshot()
		require.Equal(t, "", snapshot.Doing)
		require.Equal(t, progress.StateDone, snapshot.State)
		require.Equal(t, 5, snapshot.Total)
		require.Equal(t, 5, snapshot.Completed)
		require.Equal(t, 0, snapshot.NotStarted)
		require.Equal(t, 0, snapshot.InProgress)
		require.Equal(t, 1.0, snapshot.Progress)
		require.Equal(t, snapshot.Progress, prog.Progress())
	}

	// create 3 new steps to test the Step.SetAsCurrent() helper
	{
		prog.AddStep("step10")
		prog.AddStep("step11")
		prog.AddStep("step12")
		require.Equal(t, "", prog.Snapshot().Doing)
		prog.Get("step11").SetAsCurrent()
		require.Equal(t, "step11", prog.Snapshot().Doing)
		prog.Get("step10").SetAsCurrent()
		require.Equal(t, "step10", prog.Snapshot().Doing)
		prog.Get("step12").SetAsCurrent()
		require.Equal(t, "step12", prog.Snapshot().Doing)
		prog.Get("step12").Done()
		require.Equal(t, "", prog.Snapshot().Doing)
	}

	// debug
	// fmt.Println(u.PrettyJSON(prog))
}

func TestSubscribe(t *testing.T) {
	prog := progress.New()
	done := make(chan bool)
	ch := make(chan *progress.Step, 0)
	prog.Subscribe(ch)
	seen := 0
	go func() {
		for step := range ch {
			if step == nil {
				break
			}
			seen++
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
	prog.Get("step3").Done()
	// fmt.Println(u.PrettyJSON(prog))
	<-done
	close(ch)
	require.Equal(t, seen, 9)
}
