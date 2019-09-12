package errlog

import "log"

var (
	//DefaultLoggerPrintFunc is log.Printf without return values
	DefaultLoggerPrintFunc = func(format string, data ...interface{}) {
		log.Printf(format+"\n", data...)
	}

	//DefaultLoggerPrintFunc is log.Printf without return values
	DefaultLoggerPrintlnFunc = func(data ...interface{}) {
		log.Println(data...)
	}

	//DefaultLogger logger implements default configuration for a logger
	DefaultLogger = &logger{
		config: &Config{
			PrintFunc:          DefaultLoggerPrintFunc,
			PrintlnFunc:        DefaultLoggerPrintlnFunc,
			LinesBefore:        4,
			LinesAfter:         2,
			PrintStack:         false,
			PrintSource:        true,
			PrintError:         true,
			ExitOnDebugSuccess: false,
		},
	}
)
