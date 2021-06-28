# Cells Linter

This tool is a custom GO linter for [Cells](https://github.com/pydio/cells) code, based on the go/analysis package. 
It must be passing before builds are performed. 

## Usage

Compile the linter and run `./cells-linter github.com/pydio/cells/...` (or any other package, or a specific go file).

The execution output is 0 if no warnings were found, 3 otherwise.

## Analyzers

### addcheck

A sample boilerplate based on [Using go/analysis to write a custom linter](https://arslan.io/2019/06/13/using-go-analysis-to-write-a-custom-linter/) blog post.

### zapslices

An analyzer looking for zap.Any() calls that would pass a "slice" as argument to logger. 
Real-life examples show that passing huge slices to the logger can totally stick a service.
When writing code for Cells, developers must use log.DangerouslyZapSmallSlices() instead, indicating that they are sure that the logged slice will never grow huge.

