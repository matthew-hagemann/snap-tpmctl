package log

import "maps"

// SetLevelHandler allows to define the default handler function for a given level.
func SetLevelHandler(level Level, handler Handler) {
	if handler == nil {
		h, ok := defaultHandlers[level]
		if !ok {
			return
		}
		handler = h
	}
	handlers[level] = handler
}

// SetHandler allows to define the default handler function for all log levels.
func SetHandler(handler Handler) {
	if handler == nil {
		handlers = maps.Clone(defaultHandlers)
		return
	}
	for _, level := range allLevels {
		handlers[level] = handler
	}
}
