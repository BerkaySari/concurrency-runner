package concurrencyrunner

type ConcurrencyRunner struct {
	Results []Result
}

type Result struct {
	Result any
	Error  error
}
