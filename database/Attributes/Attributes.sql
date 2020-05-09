create table Attributes (
  ID int NOT NULL,
  SystemName nvarchar(25),
  Name nvarchar(25),
  CONSTRAINT PK_Attributes PRIMARY KEY CLUSTERED(
    ID ASC
  ) WITH (PAD_INDEX = OFF, IGNORE_DUP_KEY = OFF)
);
GO

insert into SystemVersion values('Attributes',0,0,0);
GO

