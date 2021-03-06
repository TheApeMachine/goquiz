package dbclient

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/theapemachine/goquiz/quiz-service/model"
)

// IBoltClient : Sets up an interface for the bolt embedded key/value store.
type IBoltClient interface {
	OpenBoltDb()
	QueryQuizzes() ([]model.Quiz, error)
	QueryQuiz(quizID string) (model.Quiz, error)
	Seed()
	Check() bool
}

// BoltClient : Implements the bolt interface.
type BoltClient struct {
	boltDB *bolt.DB
}

// OpenBoltDb : Opens a connection to our embedded key/value store.
func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("quizzes.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func (bc *BoltClient) initializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("QuizBucket"))

		if err != nil {
			return fmt.Errorf("Create bucket failed: %s", err)
		}

		return nil
	})
}

func (bc *BoltClient) seedQuizzes() {
	total := 10

	for i := 0; i < total; i++ {
		key := strconv.Itoa(10000 + i)

		var options []model.Option
		var steps []model.Step

		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				option := model.Option{
					ID:      strconv.Itoa(10000 + k),
					Name:    "Option_" + strconv.Itoa(k),
					Correct: false,
				}

				options = append(options, option)
			}

			step := model.Step{
				ID:       strconv.Itoa(10000 + j),
				Position: j + 1,
				Options:  options,
			}

			steps = append(steps, step)
		}

		quiz := model.Quiz{
			ID:    key,
			Name:  "Quiz_" + strconv.Itoa(i),
			Steps: steps,
		}

		jsonBytes, _ := json.Marshal(quiz)

		bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("QuizBucket"))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}

	fmt.Printf("Seeded %v fake quizzes...\n", total)
}

// Seed : Seed the DB with some data.
func (bc *BoltClient) Seed() {
	bc.initializeBucket()
	bc.seedQuizzes()
}

// QueryQuizzes : Implement retrieval of a list of quizzes from the datastore.
func (bc *BoltClient) QueryQuizzes() ([]model.Quiz, error) {
	quizzes := []model.Quiz{}

	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("QuizBucket"))

		c := b.Cursor()

		if c == nil {
			return fmt.Errorf("No quizzes found")
		}

		for k, v := c.First(); k != nil; k, v = c.Next() {
			quizzes = append(quizzes, model.Quiz{
				ID:   string(k),
				Name: string(v),
			})
		}

		return nil
	})

	if err != nil {
		return []model.Quiz{}, err
	}

	return quizzes, nil
}

// QueryQuiz : Implements a way to retrieve an Quiz by ID.
func (bc *BoltClient) QueryQuiz(quizID string) (model.Quiz, error) {
	quiz := model.Quiz{}

	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("QuizBucket"))

		quizBytes := b.Get([]byte(quizID))

		if quizBytes == nil {
			return fmt.Errorf("No quiz found for " + quizID)
		}

		json.Unmarshal(quizBytes, &quiz)

		return nil
	})

	if err != nil {
		return model.Quiz{}, err
	}

	return quiz, nil
}

// Check : Naive healthcheck.
func (bc *BoltClient) Check() bool {
	return bc.boltDB != nil
}
