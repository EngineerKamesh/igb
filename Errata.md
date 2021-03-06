# Isomorphic Go Errata List

This page contains the errata list for the Isomorphic Go book. If you come across a typo or an issue with a code example that's not included on the list below, feel free to [open a new issue](https://github.com/EngineerKamesh/igb/issues/new) and tell us about it.

## Known Issues

* The code bundled with the book, provided by the publisher, is stale. It is a snapshot of the source code from when the book was released. You should obtain and use the code from the [book's official repository](https://github.com/EngineerKamesh/igb) which will have the latest updates and bug fixes.

* There is an inconsistent spelling of the terms "server-side" (sometimes spelt "server side") and "client-side" (sometimes spelt "client side"). For the record, the author prefers the usage of "server-side" and "client-side".

* Code refactoring performed on the `github.com/james-bowman/nlp` package has caused some import calls in the Chapter 8 code examples to break. This has been resolved by vendoring the version of the nlp and sparse packages that the book utilizes, with the igb code bundle. More information on this issue can be found by reading [this ticket](https://github.com/EngineerKamesh/igb/issues/3).

## Acknowledgments

* In the "Thanks to all family members" sub-section, the initials of Sri. P.K.C. Krishnan (Financial Advisor, Chartered Accountant, Pune) are printed incorrectly. His name should be written as Sri. P.K.C. Krishnan.

## Chapter 1

* The link to Charlie Robbin's blog post, *Scaling Isomorphic JavaScript Code*, provided on page 35, no longer works since the Nodejitsu website has been shut down. An archived version of the article can be found here: [http://archive.is/ZrVMc](http://archive.is/ZrVMc)

## Chapter 3

On page 83, there is a typo in the second line of code for the GopherJS example to change the CSS style property of an element. It should be:

`element := js.Get("document").Call("getElementById", "primaryContent")`

## Chapter 9

* Page 353, 1st paragraph, it should read "Figure 9.8 is a screenshot of the live clock cogs displayed on the homepage."

