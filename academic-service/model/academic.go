package model

type Program struct {
    ID          int32
    Name        string
    Description string
}

type Course struct {
    ID        int32
    Title     string
    Code      string
    ProgramID int32
}
