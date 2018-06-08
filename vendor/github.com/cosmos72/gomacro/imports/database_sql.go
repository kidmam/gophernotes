// this file was generated by gomacro command: import _b "database/sql"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"database/sql"
)

// reflection: allow interpreted code to import "database/sql"
func init() {
	Packages["database/sql"] = Package{
	Binds: map[string]Value{
		"Drivers":	ValueOf(sql.Drivers),
		"ErrConnDone":	ValueOf(&sql.ErrConnDone).Elem(),
		"ErrNoRows":	ValueOf(&sql.ErrNoRows).Elem(),
		"ErrTxDone":	ValueOf(&sql.ErrTxDone).Elem(),
		"LevelDefault":	ValueOf(sql.LevelDefault),
		"LevelLinearizable":	ValueOf(sql.LevelLinearizable),
		"LevelReadCommitted":	ValueOf(sql.LevelReadCommitted),
		"LevelReadUncommitted":	ValueOf(sql.LevelReadUncommitted),
		"LevelRepeatableRead":	ValueOf(sql.LevelRepeatableRead),
		"LevelSerializable":	ValueOf(sql.LevelSerializable),
		"LevelSnapshot":	ValueOf(sql.LevelSnapshot),
		"LevelWriteCommitted":	ValueOf(sql.LevelWriteCommitted),
		"Named":	ValueOf(sql.Named),
		"Open":	ValueOf(sql.Open),
		"Register":	ValueOf(sql.Register),
	}, Types: map[string]Type{
		"ColumnType":	TypeOf((*sql.ColumnType)(nil)).Elem(),
		"Conn":	TypeOf((*sql.Conn)(nil)).Elem(),
		"DB":	TypeOf((*sql.DB)(nil)).Elem(),
		"DBStats":	TypeOf((*sql.DBStats)(nil)).Elem(),
		"IsolationLevel":	TypeOf((*sql.IsolationLevel)(nil)).Elem(),
		"NamedArg":	TypeOf((*sql.NamedArg)(nil)).Elem(),
		"NullBool":	TypeOf((*sql.NullBool)(nil)).Elem(),
		"NullFloat64":	TypeOf((*sql.NullFloat64)(nil)).Elem(),
		"NullInt64":	TypeOf((*sql.NullInt64)(nil)).Elem(),
		"NullString":	TypeOf((*sql.NullString)(nil)).Elem(),
		"Out":	TypeOf((*sql.Out)(nil)).Elem(),
		"RawBytes":	TypeOf((*sql.RawBytes)(nil)).Elem(),
		"Result":	TypeOf((*sql.Result)(nil)).Elem(),
		"Row":	TypeOf((*sql.Row)(nil)).Elem(),
		"Rows":	TypeOf((*sql.Rows)(nil)).Elem(),
		"Scanner":	TypeOf((*sql.Scanner)(nil)).Elem(),
		"Stmt":	TypeOf((*sql.Stmt)(nil)).Elem(),
		"Tx":	TypeOf((*sql.Tx)(nil)).Elem(),
		"TxOptions":	TypeOf((*sql.TxOptions)(nil)).Elem(),
	}, Proxies: map[string]Type{
		"Result":	TypeOf((*P_database_sql_Result)(nil)).Elem(),
		"Scanner":	TypeOf((*P_database_sql_Scanner)(nil)).Elem(),
	}, 
	}
}

// --------------- proxy for database/sql.Result ---------------
type P_database_sql_Result struct {
	Object	interface{}
	LastInsertId_	func(interface{}) (int64, error)
	RowsAffected_	func(interface{}) (int64, error)
}
func (P *P_database_sql_Result) LastInsertId() (int64, error) {
	return P.LastInsertId_(P.Object)
}
func (P *P_database_sql_Result) RowsAffected() (int64, error) {
	return P.RowsAffected_(P.Object)
}

// --------------- proxy for database/sql.Scanner ---------------
type P_database_sql_Scanner struct {
	Object	interface{}
	Scan_	func(_proxy_obj_ interface{}, src interface{}) error
}
func (P *P_database_sql_Scanner) Scan(src interface{}) error {
	return P.Scan_(P.Object, src)
}