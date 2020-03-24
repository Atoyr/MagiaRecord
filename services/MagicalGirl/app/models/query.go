package models

const selectMagicalGirlQuery = `
  select 
     ID
    ,Name
    ,Version
    ,RomanName
    ,AttributeID
    ,TypeID
    ,DiskID
    ,ConnectID
    ,MagiaID
    ,DoppelID
  from MagicalGirl
`
