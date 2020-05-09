create table Disks (
  ID int NOT NULL,
  Accele tinyint,
  BlastV tinyint,
  BlastH tinyint,
  Charge tinyint,
  CONSTRAINT PK_Type PRIMARY KEY CLUSTERED(
    ID ASC
  ) WITH (PAD_INDEX = OFF, IGNORE_DUP_KEY = OFF)
);
GO

insert into SystemVersion values('Disks',0,0,0);
GO
