[![Build Status](https://travis-ci.org/pocke/goevent.svg?branch=master)](https://travis-ci.org/pocke/goevent)
[![Coverage Status](https://img.shields.io/coveralls/pocke/goevent.svg)](https://coveralls.io/r/pocke/goevent)
[![GoDoc](https://godoc.org/github.com/pocke/goevent?status.svg)](https://godoc.org/github.com/pocke/goevent)


goevent
===============

goevent is event dispatcher written by golang.


example
===========

listen for event
-----------------

```go
e := goevent.New()
e.On(func(i int, s string){
  fmt.Printf("%d: %s\n", i, s)
})
```

Trigger
----------

```go
e.Trigger(1, "foo")
```


Use event table
----------------

```go
table := goevent.NewTable()
table.On("foo", func(i int){
  fmt.Printf("foo: %d\n", i)
})
table.On("bar", func(s string){
  fmt.Printf("bar: %s\n", s)
})

table.Trigger("foo", 1)
table.Trigger("bar", "hoge")
table.Trigger("bar", 38)    // retrun error
```


LICENSE
-----------

Copyright &copy; 2015 pocke
Licensed [MIT][mit]
[MIT]: http://www.opensource.org/licenses/mit-license.php
