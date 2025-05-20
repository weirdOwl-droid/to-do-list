package utl

// ANSI escape codes untuk warna dan gaya
const (
	reset     = "\033[0m"
	green     = "\033[32m"
	yellow    = "\033[33m"
	cyan      = "\033[36m"
	bold      = "\033[1m"
	boldGreen = "\033[1;32m"
)

// Struktur data Task
type Task struct {
	Id        int
	Deskripsi string
	Status    bool
}

// Variabel global
var TaskList []Task
var IdNext int = 1
