package timescale

type TimeScale string

const (
	HOURS        TimeScale = "hours"
	MINUTES      TimeScale = "minutes"
	SECONDS      TimeScale = "seconds"
	MILLISECONDS TimeScale = "milliseconds"
	NANOSECONDS  TimeScale = "nanoseconds"
)


type ITimeScale interface {
	toString() func(str TimeScale) string
}
