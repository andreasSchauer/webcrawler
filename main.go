package main

import (

)

func main() {

	htmlBody := `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
	`

	rawBaseURL := "https://blog.boot.dev"
	getURLsFromHTML(htmlBody, rawBaseURL)
}