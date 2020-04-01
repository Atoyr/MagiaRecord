create table Disk (
  ID int NOT NULL,
  Accele tinyint,
  BlastH tinyint,
  BlastV tinyint,
  Charge tinyint,
  CONSTRAINT PK_Type PRIMARY KEY CLUSTERED(
    ID ASC
  ) WITH (PAD_INDEX = OFF, IGNORE_DUP_KEY = OFF)
);
GO

insert into SystemVersion values('Disk',0,0,0);
GO
