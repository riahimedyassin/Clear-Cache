package main

import (
	"sync"

	"github.com/riahimedyassin/Clear-Cache/lib"
)

func init() {

}

func main() {
	c := lib.NewConfigLoader()
	c.InitConfig()
	wg := sync.WaitGroup{}
	go lib.DeleteDirectoryFiles(c.Config.CACHE, &wg)
	go lib.DeleteDirectoryFiles(c.Config.ROAMING, &wg)
}
