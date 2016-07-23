package core

import (
	"testing"
)

func TestBasicSortedIntIdSet(t *testing.T) {
	idSets := NewSortedIntSet()
	idSets.Add(23)
	if len(idSets.ids.data) != 1 || idSets.ids.data[0] != 23 {
		t.Errorf("error insert id into idsets %v", idSets.ids.data)
	}
	idSets.Add(34)
	if len(idSets.ids.data) != 2 || idSets.ids.data[0] != 23 || idSets.ids.data[1] != 34 {
		t.Errorf("error in insert second items %v", idSets.ids.data)
	}
}

func TestAdvancedSortedIntIdSet(t *testing.T) {
	idSets := NewSortedIntSet()
	maxLen := DEFAUL_ID_NUM + 1
	for i := 0; i < maxLen; i++ {
		idSets.Add(DocId(i * 2))
		idSets.Add(DocId(i*2 + 1))
	}

	if len(idSets.ids.data) != maxLen*2 {
		t.Errorf("wrong length in results")
	}

	for i := 0; i < maxLen*2; i++ {
		if idSets.ids.data[i] != DocId(i) {
			t.Errorf("wrong in elements in docidset %v != %v", idSets.ids.data[i], i)
		}
	}
}
