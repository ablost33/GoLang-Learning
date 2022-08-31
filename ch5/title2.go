package ch5

import "./html"

/*
	This function is intended to report if an HTML document contains
	multiple <title> elements. The goal here was to experiment with handling panics,
	specifically by using a distinct type for panic values and testing whether that value
	was returned by recover

	REMINDER:
		- recover() simply ends the current state of panic and returns the panic value
*/
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)
		}
	}()

	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}
