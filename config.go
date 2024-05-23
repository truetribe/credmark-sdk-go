package credmark

type Config struct {
	ApiKey       string `validate:"required"`
	RetryWaitMin int    // Second
	RetryWaitMax int    // Second
	RetryMax     int    // Time. e.g. 10 for 10 times
}
