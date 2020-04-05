package Controllers

import "net/url"


func IsValidUrl(url_ string) bool {
	_, err := url.ParseRequestURI(url_)
	if err != nil {
		return false
	}
	return true
}