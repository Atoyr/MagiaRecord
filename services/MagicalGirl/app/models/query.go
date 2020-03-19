package models

const selectAttributeQuery = `
  select 
    ID,
    Name
  from Magia.dbo.Attribute
`

const selectTypeQuery = `
  select 
    ID,
    Name
  from Magia.dbo.Type
`

const selectMagicalGirlQuery = `
  select 
    ID,
    Name,
    Version,
    RomanName,
    AttributeID,
    TypeID
  from Magia.dbo.MagicalGirl
`
