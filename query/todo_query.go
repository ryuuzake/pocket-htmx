package query

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/ryuuzake/pocket-htmx/model"
)

func TodoQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&model.Todo{})
}

func GetListTodos(dao *daos.Dao) ([]*model.Todo, error) {
	todos := []*model.Todo{}

	err := TodoQuery(dao).
		Limit(10).
		OrderBy("created desc").
		All(&todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}
