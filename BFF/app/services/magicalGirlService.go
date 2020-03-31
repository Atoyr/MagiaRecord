package services

import (
  "github.com/atoyr/MagiaRecord/BFF/app"
  "context"
 )


type MagicalGirlService struct {

}



func(mgs *MagicalGirlService) GetMagcalGirl(context.Context, *app.GetMagicalGirlParam) (*app.GetMagicalGirlResponce, error) {
  mg := app.MagicalGirl{ Id : 0 ,Name : "hoge"}
  mgr := &app.GetMagicalGirlResponce {}
  mgr.MagicalGirls = append(mgr.MagicalGirls, &mg)
  return mgr,nil
}
