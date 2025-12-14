package safety

var AllowedCommands = map[string]bool{
	"ls":    true,
	"cat":   true,
	"echo":  true,
	"mkdir": true,
	"grep":  true,
	"pwd":   true,
	"go":    true,
	"dir":   true, // Windows
}

func IsAllowedRequest(cmd string) bool {
	return AllowedCommands[cmd]
}
