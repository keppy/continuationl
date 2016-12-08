package continuation

import "fmt"
import "net/url"

/*
* continuationl lets you describe continuations via an interface that takes
* some sort of continuation notation that you parse in order to generate the NextFunction
*
* You need to satisfy the Continuation interface with whichever type fits your domain.
* Right now the only available Continuation type is Pathlike, but you could implement your own.
*
* The bare usage is to create a Pathlike in your own code and satisfy the NextFunction() in the
* Continuation interface; it should return a func() you specified in your path.
*
 */

type Continuation interface {
	NextFunction() func()
	Parse()
}

type Pathlike struct {
	Path *url.URL
}

func (p Pathlike) Parse(s string) {
	u, err := url.Parse(s)
	if err != nil {
		fmt.Errorf("error parsing url: %d", err)
	}
	p.Path = u
}

func (p Pathlike) Print() {
	fmt.Println(p.Path)
}
