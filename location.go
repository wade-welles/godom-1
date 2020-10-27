package godom

import "syscall/js"

// Location .
type Location struct {
	val js.Value
}

// Hash .
func (l *Location) Hash() string {
	return l.val.Get("hash").String()
}

/*

Properties
Location.href
Is a stringifier that returns a USVString containing the entire URL. If changed, the associated document navigates to the new page. It can be set from a different origin than the associated document.
Location.protocol
Is a USVString containing the protocol scheme of the URL, including the final ':'.
Location.host
Is a USVString containing the host, that is the hostname, a ':', and the port of the URL.
Location.hostname
Is a USVString containing the domain of the URL.
Location.port
Is a USVString containing the port number of the URL.
Location.pathname
Is a USVString containing an initial '/' followed by the path of the URL.
Location.search
Is a USVString containing a '?' followed by the parameters or "querystring" of the URL. Modern browsers provide URLSearchParams and URL.searchParams to make it easy to parse out the parameters from the querystring.
Location.hash
Is a USVString containing a '#' followed by the fragment identifier of the URL.
Location.username
Is a USVString containing the username specified before the domain name.
Location.password
Is a USVString containing the password specified before the domain name.
Location.origin Read only
Returns a USVString containing the canonical form of the origin of the specific location.
Methods
Location.assign()
Loads the resource at the URL provided in parameter.
Location.reload()
Reloads the resource from the current URL. Its optional unique parameter is a Boolean, which, when it is true, causes the page to always be reloaded from the server. If it is false or not specified, the browser may reload the page from its cache.
Location.replace()
Replaces the current resource with the one at the provided URL. The difference from the assign() method is that after using replace() the current page will not be saved in session History, meaning the user won't be able to use the back button to navigate to it.
Location.toString()
Returns a USVString containing the whole URL. It is a synonym for HTMLHyperlinkElementUtils.href, though it can't be used to modify the value.

*/
