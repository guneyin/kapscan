package general

import (
	"github.com/guneyin/kapscan/internal/model"
	"github.com/guneyin/kapscan/util"
	"time"
)

type Service struct{}

func NewGeneralService() *Service {
	return &Service{}
}

func (s *Service) Status() model.GeneralStatusResponse {
	uptime := time.Now().Sub(util.GetLastRun())
	version := util.GetVersion()

	res := model.GeneralStatusResponse{}
	res.Status = model.ServiceStatusOnline
	res.Version.Version = version.Version
	res.Version.CommitHash = version.CommitHash
	res.Version.BuildTime = version.BuildTime
	res.Env = model.EnvProduction
	res.Uptime = uptime.String()

	return res
}
