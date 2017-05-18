// https://www.youtube.com/watch?v=iIztjjNTSjs

package main02

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":4000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `
        <!DOCTYPE html>
        <html>
            <head>
                <meta charset="UTF-8">
                <title>Title</title>
            </head>
            <style>
                html, body, h1 {
                    padding: 0;
                    border: 0;
                    margin: 0;
                    box-sizing: border-box;
                }

                body {
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    background-image: linear-gradient(red, yellow, blue);
                }

                h1 {
                    font-size: 8rem;
                    color: white;
                }

            </style>
            <body>
                <h1>hello world</h1>
            </body>
        </html>
        `)
}
