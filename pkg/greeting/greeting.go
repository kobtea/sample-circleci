package greeting

import "fmt"

func Say(name string) string {
	return fmt.Sprintf("hello %s", name)
}
