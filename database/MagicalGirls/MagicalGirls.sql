create table MagicalGirls (
  ID int IDENTITY(1,1) NOT NULL,
  Name nvarchar(100) NOT NULL,
  Version nvarchar(50) NOT NULL,
  RomanName nvarchar(100) NOT NULL,
  AttributeID int NOT NULL,
  TypeID int NOT NULL,
  DiskID int NOT NULL,
  ConnectID int NOT NULL,
  
  CONSTRAINT PK_MagicalGirl PRIMARY KEY NONCLUSTERED(
    ID ASC
  ) WITH (PAD_INDEX = OFF, IGNORE_DUP_KEY=OFF)
)
GO
CREATE CLUSTERED INDEX IDX_RomanName ON MagicalGirls(
  RomanName ASC
)WITH (PAD_INDEX = OFF, IGNORE_DUP_KEY=OFF)
GO
insert into SystemVersion values('MagicalGirls',0,0,0);
GO