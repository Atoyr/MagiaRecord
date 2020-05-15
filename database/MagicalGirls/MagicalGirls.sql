create table MagicalGirls (
  ID int NOT NULL,
  Name nvarchar(128) NOT NULL,
  Version nvarchar(64) NOT NULL,
  RomanName nvarchar(128) NOT NULL,
  MinRare tinyint NOT NULL,
  MaxRare tinyint NOT NULL,
  AttributeID int NOT NULL,
  TypeID int NOT NULL,
  Accele tinyint NOT NULL,
  BlastV tinyint NOT NULL,
  BlastH tinyint NOT NULL,
  Charge tinyint NOT NULL,
  ConnectName nvarchar(128),
  MagiaName nvarchar(64),
  DoppleName nvarchar(64),
  
  CONSTRAINT PK_MagicalGirls PRIMARY KEY NONCLUSTERED(
    ID ASC
  ) WITH (PAD_INDEX = OFF, IGNORE_DUP_KEY=OFF)
)

CREATE CLUSTERED INDEX IDX_Name ON MagicalGirls(
  Name ASC
)WITH (PAD_INDEX = OFF, IGNORE_DUP_KEY=OFF)

insert into SystemVersion values('MagicalGirls',0,0,0);
