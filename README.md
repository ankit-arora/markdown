[![LinuxUnitTest](https://github.com/go-spectest/markdown/actions/workflows/linux_test.yml/badge.svg)](https://github.com/go-spectest/markdown/actions/workflows/linux_test.yml)
[![MacUnitTest](https://github.com/go-spectest/markdown/actions/workflows/mac_test.yml/badge.svg)](https://github.com/go-spectest/markdown/actions/workflows/mac_test.yml)
[![WindowsUnitTest](https://github.com/go-spectest/markdown/actions/workflows/windows_test.yml/badge.svg)](https://github.com/go-spectest/markdown/actions/workflows/windows_test.yml)
[![reviewdog](https://github.com/go-spectest/markdown/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/go-spectest/markdown/actions/workflows/reviewdog.yml)
[![Gosec](https://github.com/go-spectest/markdown/actions/workflows/gosec.yml/badge.svg)](https://github.com/go-spectest/markdown/actions/workflows/gosec.yml)
![Coverage](https://raw.githubusercontent.com/go-spectest/octocovs-central-repo/main/badges/go-spectest/markdown/coverage.svg)
# What is markdown package
The Package markdown is a simple markdown builder from Go code. This library assembles Markdown using method chaining, not uses a template engine like [html/template](https://pkg.go.dev/html/template). 
  
This library was initially developed to display test results in [go-spectest/spectest](https://github.com/go-spectest/spectest). Therefore, it implements the features required by spectest, but there are no plans to add additional functionalities unless requested by someone.
  
Additionally, complex code that increases the complexity of the library, such as generating nested lists, will not be added. I want to keep this library as simple as possible.
  
## Supported OS and go version
- OS: Linux, macOS, Windows
- Go: 1.18 or later
  
## Supported Markdown features
- [x] Heading; H1, H2, H3, H4, H5, H6
- [x] Blockquote 
- [x] Bullet list
- [x] Ordered list
- [x] Checkbox list 
- [x] Code blocks
- [x] Horizontal rule 
- [x] Table
- [x] Text formatting; bold, italic, code, strikethrough, bold italic
- [x] text with link
- [x] text with image
- [x] plain text
- [x] details 
  
## Example
```go
package main

import (
	"os"

	md "github.com/go-spectest/markdown"
)

func main() {
	md.NewMarkdown(os.Stdout).
		H1("This is H1").
		PlainText("This is plain text").
		H2f("This is %s with text format", "H2").
		PlainTextf("Text formatting, such as %s and %s, %s styles.",
			md.Bold("bold"), md.Italic("italic"), md.Code("code")).
		H2("Code Block").
		CodeBlocks(md.SyntaxHighlightGo,
			`package main
import "fmt"

func main() {
	fmt.Println("Hello, World!")
}`).
		H2("List").
		BulletList("Bullet Item 1", "Bullet Item 2", "Bullet Item 3").
		OrderedList("Ordered Item 1", "Ordered Item 2", "Ordered Item 3").
		H2("CheckBox").
		CheckBox([]md.CheckBoxSet{
			{Checked: false, Text: md.Code("sample code")},
			{Checked: true, Text: md.Link("Go", "https://golang.org")},
			{Checked: false, Text: md.Strikethrough("strikethrough")},
		}).
		H2("Blockquote").
		Blockquote("If you can dream it, you can do it.").
		H3("Horizontal Rule").
		HorizontalRule().
		H2("Table").
		Table(md.TableSet{
			Header: []string{"Name", "Age", "Country"},
			Rows: [][]string{
				{"David", "23", "USA"},
				{"John", "30", "UK"},
				{"Bob", "25", "Canada"},
			},
		}).
		H2("Image").
		PlainTextf(md.Image("sample_image", "./sample.png")).
		Build()
}
```

Output:
````
# This is H1
This is plain text
  
## This is H2 with text format
Text formatting, such as **bold** and *italic*, `code` styles.
  
## Code Block
```go
package main
import "fmt"

func main() {
        fmt.Println("Hello, World!")
}
```
  
## List
- Bullet Item 1
- Bullet Item 2
- Bullet Item 3
1. Ordered Item 1
2. Ordered Item 2
3. Ordered Item 3
  
## CheckBox
- [ ] `sample code`
- [x] [Go](https://golang.org)
- [ ] ~~strikethrough~~
  
## Blockquote
> If you can dream it, you can do it.
  
### Horizontal Rule
---
  
## Table
| NAME  | AGE | COUNTRY |
|-------|-----|---------|
| David |  23 | USA     |
| John  |  30 | UK      |
| Bob   |  25 | Canada  |

## Image
![sample_image](./sample.png)
````

If you want to see how it looks in Markdown, please refer to the following link.
- [sample.md](./doc/generated_example.md)

## License
[MIT License](./LICENSE)