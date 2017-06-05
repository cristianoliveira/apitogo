package store

import (
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/document"
	"time"
)

type Store struct {
	db bleve.Index
}

func (s *Store) Indexer() bleve.Index {
	return s.db
}

func (s *Store) Index(id string, data interface{}) error {
	return s.db.Index(id, data)
}

func (s *Store) IndexCollections(colName string, collections []interface{}) error {
	for _, c := range collections {
		col := c.(map[string]interface{})
		indexID := fmt.Sprint(colName, "-", col["id"])
		err := s.db.Index(indexID, col)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Store) All() (*bleve.SearchResult, error) {
	query := bleve.NewMatchAllQuery()
	searchRequest := bleve.NewSearchRequest(query)
	return s.db.Search(searchRequest)
}

func (s *Store) QueryById(id string) (*bleve.SearchResult, error) {
	query := bleve.NewDocIDQuery([]string{id})
	search := bleve.NewSearchRequest(query)
	return s.db.Search(search)
}

func (s *Store) QueryMatch(queryString string) (*bleve.SearchResult, error) {
	query := bleve.NewMatchQuery(queryString)
	search := bleve.NewSearchRequest(query)
	return s.db.Search(search)
}

func (s *Store) QueryMatchAsDocs(queryString string) ([]interface{}, error) {
	query := bleve.NewMatchQuery(queryString)
	search := bleve.NewSearchRequest(query)
	res, err := s.db.Search(search)
	if err != nil {
		return nil, err
	}

	var documents []interface{}
	for _, hit := range res.Hits {
		document, err := s.Document(hit.ID)
		if err != nil {
			return nil, err
		}
		documents = append(documents, document)
	}

	return documents, nil
}

func (s *Store) Document(id string) (map[string]interface{}, error) {
	doc, err := s.db.Document(id)
	if err != nil {
		return nil, err
	}
	return DocToJson(doc), nil
}

func DocToJson(doc *document.Document) map[string]interface{} {
	rv := make(map[string]interface{})
	if doc == nil {
		return rv
	}

	for _, field := range doc.Fields {
		var newval interface{}
		switch field := field.(type) {
		case *document.TextField:
			newval = string(field.Value())
		case *document.NumericField:
			n, err := field.Number()
			if err == nil {
				newval = n
			}
		case *document.DateTimeField:
			d, err := field.DateTime()
			if err == nil {
				newval = d.Format(time.RFC3339Nano)
			}
		}
		existing, existed := rv[field.Name()]
		if existed {
			switch existing := existing.(type) {
			case []interface{}:
				rv[field.Name()] = append(existing, newval)
			case interface{}:
				arr := make([]interface{}, 2)
				arr[0] = existing
				arr[1] = newval
				rv[field.Name()] = arr
			}
		} else {
			rv[field.Name()] = newval
		}
	}

	return rv
}

func NewStore() *Store {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.NewMemOnly(mapping)
	if err != nil {
		panic(err)
	}

	return &Store{db: index}
}
