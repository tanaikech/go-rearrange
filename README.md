go-rearrange
=====

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENCE)

<a name="TOP"></a>
# Overview
This is a Golang library to interactively rearrange a text data on a terminal.

# Description
When a text data is rearranged, there are many applications for automatically sorting text data. But there are a little CLI applications for manually rearranging it using. Furthermore, I just had to create an application for manually and interactively rearranging data. So I created this.

## Features
- Data can be interactively rearranged on your terminal as a CLI tool.
- Output rearranged data as ``[]string``.
- Retrieve selected values and select history.

# Install
You can install this using ``go get`` as follows.

~~~bash
$ go get -u github.com/tanaikech/go-rearrange
~~~

This library uses [termbox-go](https://github.com/nsf/termbox-go).

# Usage
~~~
result, history, err := rearrange.Do(data, step, selectmode)
~~~

#### Input
- **data ([]string)** : Data for rearranging. **This has to be 1 dimensional string array.** Each element is rearranged.
- **step (int)** : Number of steps for PageUp, PageDown.
- **selectmode (bool)** : If this is true, it's used as select mode. In this case, users only select a value. The selected values are output.

#### Output
- **result ([]string)** : Rearranged data. Data is returned as ``[]string``.
- **history ([]string)** : History of rearranged data. When the data is rearranged, the selected data is saved as a history like ``[{Index:9 Value:sample10} {Index:1 Value:sample1}]``.
- **err (error)** : Error.

### Keys for rearranging
Use up, down, page up, page down, home, end, enter, back space, Ctrl + c and escape keys.

- **Up** and **Down** are used for moving one line.
- **Page up** and **Page down** are used for moving several lines.
- **Home** and **End** are used for moving top and bottom of data.
- **Enter** is used for selecting a value to rearrange.
- **Back space** is used for reset the rearranged data.
- **Ctrl + c** and **Escape** are used for finishing rearranging.

# Applications
- [gorearrange](https://github.com/tanaikech/gorearrange) : This is a CLI tool to interactively rearrange a text data on a terminal.

<a name="Update_History"></a>
# Update History
* v1.0.0 (October 15, 2017)

    Initial release.

[TOP](#TOP)
