package main
import (
	"github.com/cheggaaa/pb"
	"time"
)

func main() {
	count := 6000
	bar := pb.StartNew(count)
	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.ShowBar = true
	bar.ShowSpeed = true
	bar.FinishPrint("The End!")
}