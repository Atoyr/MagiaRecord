create table Type (
  ID int NOT NULL,
  SystemName nvarchar(25),
  Name nvarchar(25),
  CONSTRAINT PK_Type PRIMARY KEY CLUSTERED(
    ID ASC
  ) WITH (PAD_INDEX = OFF, IGNORE_DUP_KEY = OFF)
);
GO
insert into SystemVersion values('Type',0,0,0);
GO
