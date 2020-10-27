package godom

var console = global.Get("console")

// Log .
func Log(v ...interface{}) {
	console.Call("log", v...)
}

// LogInfo .
func LogInfo(v ...interface{}) {
	console.Call("info", v...)
}

// LogWarn .
func LogWarn(v ...interface{}) {
	console.Call("warn", v...)
}

// LogDebug .
func LogDebug(v ...interface{}) {
	console.Call("debug", v...)
}

// LogError .
func LogError(v ...interface{}) {
	console.Call("error", v...)
}

// LogTrace .
func LogTrace() {
	console.Call("trace")
}

// LogGroup .
func LogGroup(v interface{}) {
	console.Call("group", v)
}

// LogGroupEnd .
func LogGroupEnd() {
	console.Call("groupEnd")
}
