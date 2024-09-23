package syllabus

import (
    "embed"
    "net/http"
)

//go:embed syllabus.json
var syllabusData embed.FS

func SyllabusHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        data, err := syllabusData.ReadFile("syllabus.json")
        if err != nil {
            http.Error(w, "Could not read syllabus", http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(data)
    } else if r.Method == http.MethodDelete {
        w.Write([]byte("deleted – stubbed"))
    } else if r.Method == http.MethodPost {
        w.Write([]byte("create-stubbed"))
    } else if r.Method == http.MethodPut {
        w.Write([]byte("update-stubbed"))
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
