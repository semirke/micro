package buffer
/*
In this module I've created a TrieNode structure that holds the children nodes and the end data structure. Data gets inserted into TrieNode via the Insert function, and the end TrieNode holds your data structure based on tag name.
Then a TagMap object holds a map of the Trie for each source file. The `insertRecord` function is used to insert records into their respective Trie, where the key is the source file and value is the tag name.
The `find` function is available in both TagMap and TrieNode for searching records by tag name or prefix from the respective Trie.
The `parseRecord` function parses the record string and returns a Record struct. This function can be modified to handle the data as per the input string format.
In the `NewTrieFromString` function, we split the data string into individual records, parse each record, and insert them into their respective Trie.
The `InsertFromString` function is similar, but it uses an already existing TagMap.
*/

import (
	"fmt"
	"strings"
)

// Tags record structure
type Record struct {
	tagName      string
	sourceFile   string
	definition   string
	tag_type     string
	startLine    string
	scope        string
	reffile      string
	accessMethod string
}

// Node in th tree. End node holds our structure
type TrieNode struct {
	children map[rune]*TrieNode
	end      bool
	data     []*Record
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

// We
type TagMap struct {
	tries map[string]*TrieNode
}

// Returns an empty TagMap
func NewTagMap() *TagMap {
	return &TagMap{tries: make(map[string]*TrieNode)}
}


func (t *TrieNode) insert(record *Record) {
	node := t
	for _, r := range record.tagName {
		if _, ok := node.children[r]; !ok {
			node.children[r] = NewTrieNode()
		}
		node = node.children[r]
	}
	node.end = true
	node.data=append(node.data, record)
}

// Inserts a record to the tree identified by `key` at record.tagName
// Creates new tree if it did not exist
func (t *TagMap) InsertRecord(key string, record *Record) {
	root, ok := t.tries[key]
	if !ok {
		root = NewTrieNode()
		t.tries[key] = root
	}
	root.insert(record)
}

func (t *TagMap) find(key string, prefix string) []*Record {
	if root, ok := t.tries[key]; ok {
		return root.find(prefix, 0)
	}
	return []*Record{}
}

func (t *TrieNode) find(prefix string, index int) []*Record {
	if index == len(prefix) {
		if t.end == true {
			return t.data
		} else {
			var results []*Record
			for _, childNode := range t.children {
				results = append(results, childNode.find(prefix, index)...)
			}
			return results
		}
	}
	char := rune(prefix[index])
	if childNode, ok := t.children[char]; ok {
		return childNode.find(prefix, index+1)
	}
	return []*Record{}
}

func parseRecord(line string) *Record {
	parts := strings.Split(line, "\t")
	return &Record{
		tagName:      parts[0],
		sourceFile:   parts[1],
		definition:   parts[2],
		tag_type:        parts[3],
		startLine:    parts[4],
		scope:        parts[5],
		reffile:       parts[6],
		accessMethod: parts[7],
	}
}

func NewTrieFromString(data string) (*TagMap){

	lines := strings.Split(data, "\n")
	tagMap := NewTagMap()

	for i := 1; i < len(lines); i++ {
		record := parseRecord(lines[i])
		tagMap.InsertRecord(record.sourceFile, record.tagName, record)
	}

	return tagMap
}

func (tagMap *TagMap) InsertFromString(data string) {

	lines := strings.Split(data, "\n")

	for i := 1; i < len(lines); i++ {
		record := parseRecord(lines[i])
		tagMap.InsertRecord(record.sourceFile, record.tagName, record)
	}
}
