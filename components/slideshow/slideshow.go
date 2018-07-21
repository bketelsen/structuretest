package slideshow

import (
	"fmt"

	"github.com/factorapp/structure/core"
	dom "github.com/gowasm/go-js-dom"
)

type SlideshowController struct {
	core.BasicController
	Index int
}

func (s *SlideshowController) currentSlide() int {
	return s.Index + 1
}

func (s *SlideshowController) showSlide() {
	for i, el := range s.Targets()["slide"] {
		fmt.Println(el.TagName())
		if i == s.Index-1 {
			el.Class().SetString("slide.slide-current")
		} else {
			el.Class().SetString("slide")
		}
	}

}
func (s *SlideshowController) Next(event dom.Event) {
	el := event.Target()
	fmt.Println("Target:", el)
	fmt.Println("Tag:", el.TagName())
	tlen := len(s.Targets()["slide"])
	if s.Index >= tlen {
		s.Index = 0
	}
	s.Index++
	fmt.Println("INDEX:", s.Index)
	s.showSlide()
}

func (s *SlideshowController) Previous(event dom.Event) {

	s.Index--

	if s.Index <= 0 {
		s.Index = len(s.Targets()["slide"])
		fmt.Println("INDEX:", s.Index)
		return
	}
	fmt.Println("INDEX:", s.Index)

	s.showSlide()
}
