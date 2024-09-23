package help

import (
    "net/http"
)

func HelpHandler(w http.ResponseWriter, r *http.Request) {
    helpText := `
    Available APIs:
    - GET /hello-world: Responds with plain text
    - GET /hello-world-html: Responds with HTML
    - GET /hello-world-json: Responds with JSON
    - GET /syllabus: Responds with the embedded JSON syllabus
    - DELETE /syllabus: Responds with a stubbed delete message
    - POST /syllabus: Responds with a stubbed create message
    - PUT /syllabus: Responds with a stubbed update message
    `
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte(helpText))
}
