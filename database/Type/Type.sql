create table Type (
  ID int NOT NULL,
  Name nvarchar(25),
  CONSTRATION PK_Type PRIMARY KEY CLUSTERD(
    ID ASC
  ) WITH (PAD_INDEX = OFF, IGNORE_DUP_KEY = OFF)
)
