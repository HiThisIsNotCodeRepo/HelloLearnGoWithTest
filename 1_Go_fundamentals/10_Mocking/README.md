# Mocking

## Testing scenario

- Print "3","2","1" and "Go!", between each print wait for 1 second.

## How to write mock?

Use dependency injection to give `Countdown` general interface. In testing we use `bytes.Buffer` and in real application
we use `os.Stdout`.

```
func Countdown(out io.Writer)
```

## How to mock 1 second interval?

Use dependency injection in `sleep` , first define dependency as an interface. So that we can use real Sleeper in `main`
and mock sleeper in test.

```
type Sleeper interface {
	Sleep()
}

type MockSleeper struct {
	Calls int
}

func (s *MockSleeper) Sleep() {
	s.Calls++
}

```

When testing the mock sleeper will be called 4 times. But it won't test if the time interval is in the right
order. `CountdownOperationsSpy` can fix it.

```
type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}
```

To ensure `CountdownOperationsSpy` can be accepted by the both arguments
of `func Countdown(out io.Writer, sleeper Sleeper)` we need it to implement 2 interfaces.

## How to write configurable sleeper for dependency injection

```
type ConfigurableSleeper struct {
	Duration  time.Duration
	SleepFunc func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.SleepFunc(c.Duration)
}
```

In real implementation, the signature of `time.Sleep` is same as `SleepFunc`

```
ConfigurableSleeper{Duration: 1 * time.Second, SleepFunc: time.Sleep}
```

Mock implementation

```
type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

ConfigurableSleeper{sleepTime, spyTime.Sleep}
```

To test just compare `SpyTime.durationSlept` with `sleepTime`

```
sleepTime := 5 * time.Second
....
if spyTime.durationSlept != sleepTime {
    t.Errorf("should have sletp for %v but slept for %v", sleepTime, spyTime.durationSlept)
}
```