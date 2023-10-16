# godeline

> Extract substrings within delimiters.

## Usage

Import the needed packages:

```go
import (
  "github.com/DavidEsdrs/godeline"
  "github.com/DavidEsdrs/godeline/editnode"
)
```

Instantiate a new prefix tree:

```go
tree := editnode.NewEditTree()
```

Set the delimiters that you want in the tree. First argument is the opening tag,
the second is the closing tag:

```go
tree.AddDelimiterType("[[", "]]")
tree.AddDelimiterType("{{", "}}")
tree.AddDelimiterType("{!", "!}")
```

> **Hint**: The delimiters may be equals. Such as an opening "!!" and a closing also "!!"

Instantiate the text processor giving the tree that you created:

```go
proc := godeline.NewProcessor(&tree, 1<<12) // 1<<12 = 2^12 = 4096
```

> **Hint**: The second argument defines the furthest distance the package will look to find the closing tag. If the closing tag is not within this distance, the search is terminated to avoid excessive processing.

Call the Tokenize function:

```go
result, err := proc.Tokenize("this is an [[input]] with the {{delimiters}} given in the tree {!tree!}!", false)
```

> **Hint**: The second argument indicates if you want to sanitize the result or not. If you give false: The InnerText in the resulting tokens will be "[[input]]", "{{delimiters}}" and "{!tree!}". If you give true the resulting will be "input", "delimiters" and "tree".

`result` holds the result of the search - of type TextResult - with an array of `tokens`:

```go
type TextResult struct {
	tokens        []*token.Token
	TokenQuantity int
}

// tokens is:

type Token struct {
  // Holds the InnerText
	InnerText string
	// Position holds the starting position of the token
	Position position.Position
	Length   int
	EditNode *editnode.EditNode // by now, has no use (always nil)
  // Holds the Opening and Closing tags
	Tag      tags.Tag
}
```

to get the tokens found, use the following:

```go
for _, t := range result.Tokens() {
  // process the tokens
  fmt.Printf("%v", t)
}

// output:
//  &{InnerText:[[input]] Position:{Ln:0 Col:16 Index:11} Length:9 EditNode:<nil> Tag:{Opening:[[ Closing:]]}}
//  &{InnerText:{{delimiters}} Position:{Ln:0 Col:40 Index:30} Length:14 EditNode:<nil> Tag:{Opening:{{ Closing:}}}}
//  &{InnerText:{!tree!} Position:{Ln:0 Col:67 Index:63} Length:8 EditNode:<nil> Tag:{Opening:{! Closing:!}}}
```