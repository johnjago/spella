# spella

spella provides basic English spelling correction.

You give it a word and it'll suggest a correction if it's spelled wrong.

```sh
$ go run main.go testerr
Did you mean tester?
```

If it's spelled right, it will tell you that as well.

```sh
$ go run main.go tester
Correct
```

## Why

I wanted to implement this [same example](https://github.com/johnjago/spellchecker)
in Go to see how different it would be from the Clojure implementation. The two
biggest takeaways are that it's about twice the number of lines (since I had
to make my own functions to mimick Clojure's [map](https://clojuredocs.org/clojure.core/map)
and [set](https://clojuredocs.org/clojure.core/set)), and that the types made
it much easier to recall what a function does, only by looking at its signature,
after I returned to this small project after over a week. I'll be writing a blog
post soon with more detailed observations.
