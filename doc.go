/*
Package rearrange (go-rearrange.go) :
This is a Golang library to interactively rearrange a text data on a terminal.

When a text data is rearranged, there are many applications for automatically sorting text data. But there are a little CLI applications for manually rearranging it using. Furthermore, I just had to create an application for manually and interactively rearranging data. So I created this.

# Features of "go-rearrange" are as follows.

1. Data can be interactively rearranged on your terminal as a CLI tool.

2. Output rearranged data as ``[]string``.

3. Retrieve selected values as a history.

https://github.com/tanaikech/go-rearrange/

You can read the detail information there.


---------------------------------------------------------------

# Usage

result, history, err := rearrange.Do(data, step, selectmode, indexmode)

# Input
- data ([]string) : Data for rearranging. This has to be 1 dimensional string array. Each element is rearranged.
- step (int) : Number of steps for PageUp, PageDown.
- selectmode (bool) : If this is true, it's used as select mode. In this case, users only select a value. The selected values are output.
- indexmode (bool) : If this is true, the rearranged result is output as the change of index for the source data. For example, if the source data and rearranged data are ["a", "b", "c"] and ["c", "b", "a"], respectively. The output will become [2, 1, 0].

# Output
- result ([]string) : Rearranged data. Data is returned as []string.
- history ([]string) : History of rearranged data. When the data is rearranged, the selected data is saved as a history like [{Index:9 Value:sample10} {Index:1 Value:sample1}].
- err (error) : Error.

---------------------------------------------------------------
*/
package rearrange
