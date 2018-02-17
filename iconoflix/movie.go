package iconoflix

import (
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/go-yaml/yaml"
)

var corpus Movies

// Seeds randomizer and movie db
func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	var err error
	corpus, err = LoadMem()
	if err != nil {
		panic(err)
	}
}

type (
	// Icon an emoji representation
	Icon struct {
		Name string `yaml:"name"`
	}

	// Movie an Iconoflix movie representation
	Movie struct {
		Name  string `yaml:"name"`
		Icons []Icon `yaml':",inline"`
	}

	// Movies a collection of Iconoflix movies
	Movies struct {
		Movies []Movie `yaml:"movies"`
	}
)

// RandMovie picks a random movie...
func RandMovie() Movie {
	return corpus.Movies[rand.Intn(len(corpus.Movies))]
}

// LoadMem the movie database from memory
func LoadMem() (d Movies, err error) {
	err = yaml.Unmarshal([]byte(MOVIES), &d)
	return
}

// LoadFile the movie database from file
func LoadFile(path string) (Movies, error) {
	var movies Movies

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return movies, err
	}

	err = yaml.Unmarshal(data, &movies)
	if err != nil {
		return movies, err
	}

	return movies, nil
}

// ----------------------------------------------------------------------------

// MOVIES in memory movie corpus
const MOVIES = `
movies:
- name: Home Alone
  icons:
  - name: ! "🏡"
  - name: ! "🍣"
- name: Terminator
  icons:
  - name: ! "🏡"
  - name: ! "👻"
- name: Ghost
  icons:
  - name: ! "🏡"
  - name: ! "👻"
`
