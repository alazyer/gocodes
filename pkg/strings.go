package pkg

import (
    "fmt"
    "strings"
)

func CommonStrings() {
    fmt.Println("Compare:    ", strings.Compare("ABC", "abc"))  // -1
    fmt.Println("Contains:   ", strings.Contains("test", "es"))  // true
    fmt.Println("Contains:   ", strings.Contains("Shell-12541", "1-2-9")) // false
    fmt.Println("ContainsAny:", strings.ContainsAny("Shell-12541", "1-2-9")) // true
    fmt.Println("Count:      ", strings.Count("test", "t"))  // 2
    fmt.Println("Fields:     ", strings.Fields("test t"))  // [test t]
    fmt.Println("HasPrefix:  ", strings.HasPrefix("test", "te"))  // true
    fmt.Println("HasSuffix:  ", strings.HasSuffix("test", "st"))  // true
    fmt.Println("Index:      ", strings.Index("test", "e"))  // 1
    fmt.Println("LastIndex:  ", strings.LastIndex("testeea", "e")) // 5
    fmt.Println("Join:       ", strings.Join([]string{"a", "b"}, "-"))  // a-b
    fmt.Println("Repeat:     ", strings.Repeat("a", 5))  // aaaaa
    fmt.Println("Replace:    ", strings.Replace("foo", "o", "0", -1))  // f00
    fmt.Println("Replace:    ", strings.Replace("foo", "o", "0", 1))  // f0o
    fmt.Println("Split:      ", strings.Split("a-b-c-d-e", "-")) // [a b c d e]
    fmt.Println("SplitAfter: ", strings.SplitAfter("a-b-c-d-e", "-"))  // [a- b- c- d- e]
    fmt.Println("Tittle:     ", strings.Title("TEST"))  // Test
    fmt.Println("ToLower:    ", strings.ToLower("TEST"))  // test
    fmt.Println("ToUpper:    ", strings.ToUpper("test"))  // TEST
    fmt.Println("Trim:       ", strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))  // Hello, Gophers

    fmt.Println()

    fmt.Println("Len:        ", len("hello"))  // 5
    fmt.Println("Char:       ", "hello"[1])  // 101
}
