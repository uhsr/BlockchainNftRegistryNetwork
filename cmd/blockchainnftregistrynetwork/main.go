// cmd/blockchainnftregistrynetwork/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftregistrynetwork/internal/blockchainnftregistrynetwork"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftregistrynetwork.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
