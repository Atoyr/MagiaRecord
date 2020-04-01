create table SystemVersion (
  TableName nvarchar(100),
  Mejor tinyint,
  Minor tinyint,
  Revison tinyint,
  CONSTRAINT PK_SystemVersion PRIMARY KEY CLUSTERED(
    TableName ASC
  ) WITH (PAD_INDEX = OFF, IGNORE_DUP_KEY = OFF)
);

