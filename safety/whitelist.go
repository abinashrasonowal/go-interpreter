package safety

var AllowedCommands = map[string]bool{
	"ls":    true,
	"cat":   true,
	"echo":  true,
	"mkdir": true,
	"grep":  true,
	"pwd":   true,
	"touch": true,
	"rm":    true,
	"cp":    true,
	"mv":    true,
	"find":  true,
	"wc":    true,
	"head":  true,
	"tail":  true,

	// Windows commands
	"dir":           true,
	"Get-ChildItem": true,
	"Get-Content":   true,

	// Cross-platform
	"go": true,
}

func IsAllowedRequest(cmd string) bool {
	return AllowedCommands[cmd]
}
