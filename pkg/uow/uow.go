package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// function to return a repository with the transaction assigned
type RepositoryFactory func(tx *sql.Tx) interface{}

type UowInterface interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(uow UowInterface) error) error
	CommitOrRollback() error
	Rollback() error
	Unregister(name string)
}

type Uow struct {
	Db           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory
}

func NewUow(db *sql.DB) *Uow {
	return &Uow{
		Db:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}

func (u *Uow) Register(name string, fc RepositoryFactory) {
	u.Repositories[name] = fc
}

func (u *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {

	if repo, ok := u.Repositories[name]; ok {

		if u.Tx == nil {
			tx, err := u.Db.BeginTx(ctx, nil)
			if err != nil {
				return nil, err
			}

			u.Tx = tx
		}

		return repo(u.Tx), nil
	}

	return nil, errors.New("repository not found")
}

func (u *Uow) Do(ctx context.Context, fn func(uow *Uow) error) error {

	if u.Tx != nil {
		return fmt.Errorf("transaction already started")
	}

	tx, err := u.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	u.Tx = tx

	err = fn(u)
	if err != nil {
		if errRb := u.Rollback(); errRb != nil {
			return errors.New(fmt.Errorf("error rolling back transaction: %v", errRb.Error()).Error())
		}
		return err
	}

	return u.CommitOrRollback()
}

func (u *Uow) CommitOrRollback() error {
	if u.Tx == nil {
		return fmt.Errorf("no transaction to commit")
	}

	if err := u.Tx.Commit(); err != nil {
		if errRb := u.Rollback(); errRb != nil {
			return errors.New(fmt.Errorf("error rolling back transaction: %v", errRb.Error()).Error())
		}
		return err
	}

	u.Tx = nil

	return nil
}

func (u *Uow) Rollback() error {
	if u.Tx == nil {
		return fmt.Errorf("no transaction to rollback")
	}

	if err := u.Tx.Rollback(); err != nil {
		return err
	}

	u.Tx = nil

	return nil
}

func (u *Uow) Unregister(name string) {
	delete(u.Repositories, name)
}
