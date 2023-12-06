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
	"strings"
	"strconv"
)

// Tags record structure
type Record struct {
	TagName      string
	SourceFile   string
	Definition   string
	Scope        string
	// dynamic fields
	StartLine    int
	Member       string
	Access       string
	Module       string
	Function     string
	File         string
	Typeref      string
	Signature    string
	Class        string
	Nameref      string
	Package      string
	Import       string
	Namespace    string
	End          int
	Data         string
}

// Node in th tree. End node holds our structure
type TrieNode struct {
	children map[rune]*TrieNode
	end      bool
	data     []*Record
}

type FileNode struct {
	children map[string][]*Record
}


func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode,3)}
}

func NewFileNode() *FileNode {
	return &FileNode{children: make(map[string][]*Record,3)}
}

// We
type TagMap struct {
	tries    map[string]*TrieNode
	filemaps map[string]*FileNode
}

// Returns an empty TagMap
func NewTagMap() *TagMap {
	return &TagMap{tries: make(map[string]*TrieNode, 10),
			filemaps: make(map[string]*FileNode, 10)}
}


func (t *TrieNode) insert(record *Record) {
	node := t
	for _, r := range record.TagName {
		if _, ok := node.children[r]; !ok {
			node.children[r] = NewTrieNode()
		}
		node = node.children[r]
	}
	node.end = true
	node.data=append(node.data, record)
}

// Inserts a record to the tree identified by `key` at record.TagName
// Creates new tree if it did not exist
func (t *TagMap) InsertRecord(key string, record *Record) {
	// first record of file identified by `key`
	newfile, ok := t.filemaps[key]
	if !ok {
		newfile = NewFileNode()
		newfile.children[record.TagName]  =append(newfile.children[record.TagName], record)
		t.filemaps[key] = newfile
	}

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

func (t *TagMap) findAll(prefix string) []*Record {
	recs := []*Record{}
	for _, root := range(t.tries) {
		ret := root.find(prefix, 0)
		recs = append(recs, ret...)
	}
	return recs
}

func (t *TrieNode) find(prefix string, index int) []*Record {
	if len(prefix) == 0 {
		return []*Record{}
	}

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
	if (len(parts) < 4) {
		return nil
	}

	rec := Record{
		TagName:      string(parts[0]),
		SourceFile:   parts[1],
		Definition:   parts[2],
		Scope:        parts[3],
		Data:         line,
	}


	for k:=5; k < len(parts); k++ {
		valp := strings.Split(parts[k],":")
		if len(valp) < 2 {
			continue
		}

		partType := valp[0]
		valparts := strings.Join(valp[1:], ":")
		switch(partType) {
			case "member":
				rec.Member = valparts
			case "access":
				rec.Access = valparts
			case "module":
				rec.Module = valparts
			case "function":
				rec.Function = valparts
			case "file":
				rec.File = valparts
			case "package":
				rec.Package = valparts
			case "import":
				rec.Import = valparts
			case "namespace":
				rec.Namespace = valparts
			case "typeref":
				rec.Typeref = valparts
			case "signature":
				rec.Signature = valparts
			case "class":
				rec.Class = valparts
			case "nameref":
				rec.Nameref = valparts
			case "line":
				// Luckily ctags output changes based on record type
				rec.StartLine = -1
				i, err := strconv.Atoi(valparts)
				if err == nil {
					rec.StartLine = i
				}
			case "end":
				i, err := strconv.Atoi(valparts)
				if err == nil {
					rec.End = i
				}
		}
	}

	return &rec
}

func NewTrieFromString(data string) (*TagMap){

	lines := strings.Split(data, "\n")
	tagMap := NewTagMap()

	for i := 1; i < len(lines); i++ {
		if len(lines[i]) == 0 || byte(lines[i][0]) == byte('!') {
			continue
		}
		record := parseRecord(lines[i])
		if record != nil {
			tagMap.InsertRecord(record.SourceFile, record)
		}
	}

	return tagMap
}

func (tagMap *TagMap) InsertFromString(data string) {

	lines := strings.Split(data, "\n")

	for i := 1; i < len(lines); i++ {
		if len(lines[i]) == 0 || byte(lines[i][0]) == byte('!') {
			continue
		}
		record := parseRecord(lines[i])
		if record != nil {
			tagMap.InsertRecord(record.SourceFile, record)
		}
	}
}
