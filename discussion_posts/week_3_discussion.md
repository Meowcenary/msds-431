### Notes From: "Golang Maps vs Structs, Which Method Should You Use to Parse JSON"
---
[Article URL](https://medium.com/swlh/golang-maps-vs-structs-which-method-should-you-use-to-parse-json-16c7ad7be171)
[Archived URL](https://web.archive.org/web/20210416203757/https://bencane.com/2020/12/08/maps-vs-structs-for-json/)

#### Parsing Data Into a Map
This is the example JSON that all examples will be parsing:
```
    {
        "name": "example",
        "numbers": [
            1, 2, 3, 4
        ],
        "nested": {
            "isit": true,
            "description": "a nested json"
        }
    }
```

Here is an initial attempt at parsing the data that will be improved on:
```
    package main

    import (
        "encoding/json"
        "fmt"
    )

    func main() {
        // Create a map to parse the JSON
        var data map[string]interface{}

        // Define a JSON string
        j := `{"name":"example","numbers":[1,2,3,4],"nested":{"isit":true,"description":"a nested json"}}`

        // Parse our JSON string
        err := json.Unmarshal([]byte(j), &data)
        if err != nil {
            fmt.Printf("Error parsing JSON string - %s", err)
        }

        // Print out one of our JSON values
        fmt.Printf("Name is %s", data["name"].(string))
    }
```
Parsing the JSON is straightforward enough, but the first issue lies in how the values are printed at the end. The
assumption is that data["name"] will always be defined (i.e that field will always be included in our data), but what if
it is omitted or removed?

The "comma ok" pattern can be used to check for the presence of a field before continuing when working with  JSON that
may have unknown values:
```
	// Print out one of our JSON values
	n, ok := data["name"]
	if !ok {
		// access it another way
		n = "default"
	}
	fmt.Printf("Name is %s", n.(string))
```

This pattern will helps with data that does not exist, but it assumes the value data["name"] will always be a string
type. If something else, like an intger type were passed in, the code block would fail. Using the "comma ok" pattern
again this issue can handled like below:
```
	// Print out one of our JSON values
	n, ok := data["name"]
	if !ok {
		// access it another way
		n = "default"
	}
	v, ok := n.(string)
	if !ok {
		// figure out type another way
		v = "default"
	}
	fmt.Printf("Name is %s", v)
```

Here is a code block for parsing the "numbers" field of the example JSON:
```
	// Print out the JSON Numbers
	var nums []int
	i, ok := data["numbers"].([]interface{})
	if ok {
		for _, v := range i {
			x, ok := v.(float64)
			if !ok {
				// set to default
				nums = []int{}
				break
			}
			nums = append(nums, int(x))

		}
	}
	fmt.Printf("Numbers are")
	for _, v := range nums {
		fmt.Printf(" %d", v)
	}
	fmt.Printf("\n")
```
Any array is parsed and placed into the map as an array of interfaces which is why the "comma ok" check compares to the
type []interface{}. Additionally, each value within the "numbers" field needs to be checked to ensure that it can be
cast to a type float64. If not, the for loop is broken out of and the numbers to be printed set to a blank array.

The disadvantage of working with maps is that you need to account for many variations in the data and if you fail to the
program will crash.

#### Parsing Data Into a Struct
Here is an eample of parsing the same example JSON from the beginning using a Golang struct
```
package main

import (
	"encoding/json"
	"fmt"
)

// Example is our main data structure used for JSON parsing
type Example struct {
	Name    string `json:"name"`
	Numbers []int  `json:"numbers"`
	Nested  Nested `json:"nested"`
}

// Nested is an embedded structure within Example
type Nested struct {
	IsIt        bool   `json:"isit"`
	Description string `json:"description"`
}

func main() {
	// Define a JSON string
	j := `{"name":"example","numbers":[1,2,3,4],"nested":{"isit":true,"description":"a nested json"}}`

	// Parse JSON string into data object
	var data Example
	err := json.Unmarshal([]byte(j), &data)
	if err != nil {
		fmt.Printf("Error parsing JSON string - %s", err)
	}

	// Print the name
	fmt.Printf("Name is %s", data.Name)
}
```
When using structs to parse JSON, the JSON parser itself will handle the type assertations that were done manually with
using Golang maps. The method `Unmarshal` will throw errors if the fields hold data types other than what is specified
by the definend struct.

Another example of using structs, but this time for the more complicated "Numbers" field:
```
    // Data has had Unmarshal called on it
	// Print the name
	fmt.Printf("Name is %s\n", data.Name)

	// Print the Numbers
	fmt.Printf("Numbers include")
	for _, v := range data.Numbers {
		fmt.Printf(" %d", v)
	}
	fmt.Printf("\n")

	// Print the Description
	fmt.Printf("Description is %s", data.Nested.Description)
```

#### When to Use Map versus Struct Parsing?
Using structs to parse data is generally the preferred way to parse JSON data, but in some rare instances the structure
of the data may not be known in which case a map is needed. In other words, maps are necessary for dealing with an
unknown or changing JSON structure, but otherwise use structs.

### Notes From: "A Complete Guide to JSON in Golang (With Examples)"
[Article URL](https://www.sohamkamani.com/golang/json/#structured-data-decoding-json-into-structs)

#### Structured vs Unstructured JSON
When you know the format of the JSON beforehand it is called "structured JSON". Opposite of this, if you don't know the
format of the JSON beforehand is is "unstructured JSON".

### Explain the difference between structured and unstructured JSON data as it relates to Go structures
---
In short, when you know the format of the JSON you are receiving it is considered to be structured JSON. Opposite of
this, if you are unsure of how the JSON you are receiving, for example data types may be inconsistent for a field or
fields could be occasionally ommitted is formatted it is considered unstructured JSON. Fittingly, when working with JSON
in Golang the standard approach for structured JSON is to define a Golang struct that defines the expected fields and
types which is then used by Golang's json.Unmarshal method to parse (decode) the data into the struct and json.Marshal
to encode the struct to JSON.

Give an example of either structured or unstructured JSON data and how you would implement it using Go structures
---
To give an example of structured data, suppose a baseball database with an API interface that returns JSON from a set of
publicly available endpoints. We'll consider the endpoint "exampleBaseballApi.com/team_records/:season/:team_name" where
season is an integer (e.g 2023) and team name is a string (e.g "chicago_white_sox") that returns JSON formatted like so:
```
[
    {
        wins: Integer,
        losses: Integer,
        win_loss_percentage: Float64,
    }
]
```

Since this is structured JSON (structure of the data is known beforehand) a Golang struct is the best option for parsing
this data:

```
type SeasonRecord struct {
    Wins Int,
    Losses Int,
    Win_loss_percentage Float64
}
```

Here is an example of parsing this data on [Go Playground](https://go.dev/play/p/QAAlgwscmDY)
