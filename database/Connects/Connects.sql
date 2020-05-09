create table Connects (
  ID int NOT NULL,
  Name nvarchar(50),
  MagicalGirlID int NOT NULL,
  CONSTRAINT PK_Connects PRIMARY KEY CLUSTERED(
    ID ASC
  ) WITH (PAD_INDEX = OFF, IGNORE_DUP_KEY = OFF)
);
GO

insert into SystemVersion values('Connects',0,0,0);
GO
