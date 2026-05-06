package main
import "github.com/gdamore/tcell/v2"
func main() { s, _ := tcell.NewScreen(); s.Init(); s.Fini() }
