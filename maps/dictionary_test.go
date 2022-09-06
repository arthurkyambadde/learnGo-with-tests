package dictionary

import (
	"testing"
)

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q ", got, want)
	}
}

func TestDictionary(t *testing.T) {

	dictionary := Dictionary{
		"test": "This is a test statement",
	}

	t.Run("Known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")

		want := "This is a test statement"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {

		_, err := dictionary.Search("unknown word")
		want := "word not in the dictionary"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), want)

	})

	t.Run("Add a word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "This is a test statement")

		got, err := dictionary.Search("test")

		want := "This is a test statement"

		if err != nil {
			t.Fatal("should find added word:", err)
		}

		assertStrings(t, got, want)
	})

	t.Run("Existing word", func(t *testing.T) {
		definition := "this is just a test"
		word := "test"
		dictionary := Dictionary{word: definition}

		err1 := dictionary.Add(word, "new test")

		got, err := dictionary.Search(word)

		if err != nil {
			t.Fatal("should find added word:", err)
		}

		if definition != got {
			t.Errorf("got %q want %q", got, definition)
		}

		if err1 != ErrWordExists {
			t.Errorf("got %q want %q", err, ErrWordExists)
		}

	})

}
