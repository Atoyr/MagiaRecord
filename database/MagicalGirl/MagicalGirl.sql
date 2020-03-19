use Magia
create table MagicalGirl (
  ID int IDENTITY(1,1) NOT NULL,
  Name nvarchar(100) NOT NULL,
  Version nvarchar(50) NOT NULL,
  RomanName nvarchar(100) NOT NULL,
  AttributeID nvarchar(5) NOT NULL,
  TypeID nvarchar(10) NOT NULL,
  CONSTRAINT PK_MagicalGirl PRIMARY KEY NONCLUSTERED(
    ID ASC
  ) WITH (PAD_INDEX = OFF, IGNORE_DUP_KEY=OFF)
)
GO
CREATE CLUSTERED INDEX IDX_RomanName ON MagicalGirl(
  RomanName ASC
)WITH (PAD_INDEX = OFF, IGNORE_DUP_KEY=OFF)
GO

