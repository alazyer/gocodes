package main

func main() {
    m := make(map[string]string)
    m["age"] = "19"
    delete(m, "name")
}
