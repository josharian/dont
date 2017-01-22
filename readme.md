### Don't: template-based, decentralized static analysis for Go

dont is a static analysis tool. It complements go vet, staticcheck, and regular tests.

To use dont, write a template. It looks like a regular function and describes something you shouldn't do:

```go
// If you attempt to Close a file when opening it failed, the Close call will panic.
func DontCloseFileWhenOpenFailed(path string) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return
	}
}
```

Then invoke dont on the package containing your template and any packages you want to check:

```bash
$ dont path/to/template/package my/code/...
```

You can also include the templates that ship with dont:

```bash
$ dont github.com/josharian/dont/builtin/... path/to/my/templates/... my/code/...
```

You can also include any templates that have been included in any packages you depend on:

```bash
$ dont -deps my/code/...
```

### Don't, vet, staticcheck, tests

vet and staticcheck contain careful, hand-written static analysis checks.
They're great, but they're hard to extend without effort or background knowledge.
They're also centralized.

dont is uncentralized. You can easily add your own templates with minimal ramp-up,
and no need to ask anyone's permission.

dont checks are generally less powerful and general than vet and staticcheck tests,
but in a great many cases that suffices.
And they are easier to add, to maintain, and to reason about.
dont complements vet and staticcheck.

dont also augments regular tests. Regular tests help your package function correctly.
dont helps others use your package correctly. dont functions should sit alongside
tests and examples and serve as enforceable documentation.

### Status

I've written a few dont functions to convince myself that I like how they look and
that they are expressive enough.

The actual matching engine remains to be written. :)
I expect it to be non-trivial.

It might also not live at github.com/josharian/dont. The "dont" name is subject to change,
and I might want it to be at github.com/commaok or foggy.co instead. TBD.

### License

Private for now. It'll probably be Go-like BSD and/or MIT, but I haven't fully decided yet.

