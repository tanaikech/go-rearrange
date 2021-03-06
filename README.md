go-rearrange
=====

[![Build Status](https://travis-ci.org/tanaikech/go-rearrange.svg?branch=master)](https://travis-ci.org/tanaikech/go-rearrange)
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
result, history, err := rearrange.Do(data, step, selectmode, indexmode)
~~~

#### Input
- **data ([]string)** : Data for rearranging. **This has to be 1 dimensional string array.** Each element is rearranged.
- **step (int)** : Number of steps for PageUp, PageDown.
- **selectmode (bool)** : If this is true, it's used as select mode. In this case, users only select a value. The selected values are output.
- **indexmode (bool)** : If this is true, the rearranged result is output as the change of index for the source data. For example, if the source data and rearranged data are ``["a", "b", "c"]`` and ``["c", "b", "a"]``, respectively. The output will become ``[2, 1, 0]``.

#### Output
- **result ([]string)** : Rearranged data. Data is returned as ``[]string``.
- **history ([]string)** : History of rearranged data. When the data is rearranged, the selected data is saved as a history like ``[{Index:9 Value:sample10} {Index:1 Value:sample1}]``.
- **err (error)** : Error.

### Keys for rearranging
Use up, down, page up, page down, home, end, enter, back space, Ctrl + c and escape keys.

| Key | Effect |
|:-----------|:------------|
| **Up**, **Down** | Moving one line |
| **Page up**, **Page down** | Moving several lines |
| **Home**, **End** | Moving top and bottom of data |
| **Enter** | Selecting a value to rearrange |
| **Back space** or **Space** | Reset the rearranged data |
| **Ctrl + c** or **Escape** | Finishing rearranging |

# Applications
- [gorearrange](https://github.com/tanaikech/gorearrange) : This is a CLI tool to interactively rearrange a text data on a terminal.
- [ggsrun](https://github.com/tanaikech/ggsrun/blob/master/help/README.md#rearrangescripts) : This was used for rearranging scripts in a GAS project.

<a name="Update_History"></a>
# Update History
* v1.0.0 (October 15, 2017)

    Initial release.

* v1.0.1 (October 16, 2017)

    - As one of outputs, **indexmode (bool)** was added. If this is true, the rearranged result is output as the change of index for the source data. For example, if the source data and rearranged data are ``["a", "b", "c"]`` and ``["c", "b", "a"]``, respectively. The output will become ``[2, 1, 0]``.

* v1.0.2 (October 18, 2017)

    - From this version, data included multi-bytes characters can be used. At Linux, it works fine. At Windows DOS, rearranging and selecting data can be done. But the displayed data is shifted. Although this may be a bug of termbox-go, I don't know the reason. I'm sorry. On the other hand, data with only single-byte characters works fine. About MAC, I don't have it. If someone can confirm and tell me it, I'm glad.


[TOP](#TOP)
