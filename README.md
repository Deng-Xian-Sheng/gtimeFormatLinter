# gtimeFormatLinter

## What is it？ 

This is a plugin for golangci-lint

gtime refers to gtime.Time is the time type of GoFrame. The formal parameters of its Format method are different from those of the Format method of time.Time in the standard library.

It is possible to make mistakes and will not report errors after making mistakes, so it is difficult to detect them.

The purpose of this letter is to check such errors.

## How to use it？

I have tried to submit it to golangci-lint. If golangci-lint accepts it, you can directly use golangci-lint to use it.

If you do not accept it, you can compile golangci-lint and add the linter.