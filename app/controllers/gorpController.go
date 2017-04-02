package controllers

import (
    "github.com/go-gorp/gorp"
    "database/sql"
    "github.com/revel/revel"
)

var (
	DBm *gorp.DbMap
)

type gorpController struct {
    *revel.Controller
    Txn *gorp.Transaction
}

func (c *gorpController) Begin() revel.Result {
    txn, err := Dbm.Begin()
    if err != nil {
        panic(err)
    }
    c.Txn = txn
    return nil
}

func (c *gorpController) Commit() revel.Result {
    if c.Txn == nil {
        return nil
    }
    if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
        panic(err)
    }
    c.Txn = nil
    return nil
}

func (c *gorpController) Rollback() revel.Result {
    if c.Txn == nil {
        return nil
    }
    if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
        panic(err)
    }
    c.Txn = nil
    return nil
}