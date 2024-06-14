package reflection

func walk(x interface{}, fn func(input string)) {
    fn("I still can't believe Germany put five past Scotland, while England failed to score.")
}

/* Reflection resources: https://go.dev/blog/laws-of-reflection */
