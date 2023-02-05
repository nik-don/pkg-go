# pkg-go
collection of go packages/functions ready to be imported

Currently this package contains these Go libraries
- filereadwrite: that provide functionality for reading and writing files
- urlread: reading data from URLs


## filereadwrite
The filereadwrite library provides three functions for reading and writing files:

Usage
Here's an example of how you can use the filereadwrite library:

```Go
// ReadLines: reads a whole file into memory and returns a slice of its lines.
lines, err := filereadwrite.ReadLines("file.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Println(lines)

// WriteLines: writes the lines to the given file.
err = filereadwrite.WriteLines(lines, "newfile.txt")
if err != nil {
    log.Fatal(err)
}

// AppendLines: appends a line of text to the given file. It creates the file if it doesn't already exist.
err = filereadwrite.AppendLines(lines, "appendfile.txt")
if err != nil {
    log.Fatal(err)
}
```

## urlread
The urlread library provides two functions for reading data from URLs:

Usage

Here's an example of how you can use the urlread library:

```Go
GetFromURL: returns a slice of string of in the form of lines from a URL.
lines, err := urlread.GetFromURL("https://www.example.com")
if err != nil {
    log.Fatal(err)
}
fmt.Println(lines)
```

License
This code is released under the MIT License. A copy of the license is available in the LICENSE file.

```
The MIT License (MIT)

Copyright (c) 2023 Nik Don

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
