package mysql

import "github.com/ozkar99/logger"

func init() {
	logger.SetCreateTableSQL(
		`create table if not exists logs (
			ID int not null auto_increment primary key,
			Level varchar(10) not null,
			Message varchar(250) not null,
			CreatedAt timestamp(6) default current_timestamp(6)
		)`)
}
