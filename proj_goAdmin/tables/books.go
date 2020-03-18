package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetBooksTable(ctx *context.Context) table.Table {

	booksTable := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := booksTable.GetInfo()

	info.AddField("Id", "id", db.Integer).FieldFilterable()
	info.AddField("Title", "title", db.Char)
	info.AddField("Publisher", "publisher", db.Char)
	info.AddField("Author", "author", db.Char)
	info.AddField("Created_at", "created_at", db.Timestamp)
	//info.AddField("Created_at", "created_at", db.Timestamp).FieldFilterable(types.FilterType{FormType: form.DatetimeRange})
	info.AddField("Updated_at", "updated_at", db.Timestamp)
	//info.AddField("Updated_at", "updated_at", db.Timestamp).FieldEditAble(editType.Datetime)

	info.SetTable("books").SetTitle("Books").SetDescription("Books")

	formList := booksTable.GetForm()

	formList.AddField("Id", "id", db.Integer, form.Default).FieldNotAllowAdd()
	formList.AddField("Title", "title", db.Char, form.Text)
	formList.AddField("Publisher", "publisher", db.Char, form.Text)
	formList.AddField("Author", "author", db.Char, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).FieldNotAllowAdd()
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime).FieldNotAllowAdd()

	formList.SetTable("books").SetTitle("Books").SetDescription("Books")

	return booksTable
}
