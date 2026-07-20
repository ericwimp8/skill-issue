package catalog

import (
	"os"

	"example.invalid/catalog/internal/schema"
)

func Run() {
	encoded := schema.Encode(os.Args[1:])
	_, _ = os.Stdout.Write(encoded)
}

