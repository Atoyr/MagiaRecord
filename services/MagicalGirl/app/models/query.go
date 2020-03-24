package models

const selectAttributeQuery = `
  select 
    ID,
    Name
  from Attribute
`

const selectTypeQuery = `
  select 
    ID,
    Name
  from Type
`

const selectMagicalGirlQuery = `
  select 
    ID,
    Name,
    Version,
    RomanName,
    AttributeID,
    TypeID
  from MagicalGirl
`
