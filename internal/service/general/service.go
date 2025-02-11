package general

import (
	"time"

	"github.com/guneyin/kapscan/internal/dto"
	"github.com/guneyin/kapscan/util"
)

type Service struct{}

func NewGeneralService() *Service {
	return &Service{}
}

func (s *Service) Status() dto.GeneralStatusResponse {
	uptime := time.Since(util.GetLastRun())
	version := util.GetVersion()

	res := dto.GeneralStatusResponse{}
	res.Status = dto.ServiceStatusOnline
	res.Version.Version = version.Version
	res.Version.CommitHash = version.CommitHash
	res.Version.BuildTime = version.BuildTime
	res.Env = dto.EnvProduction
	res.Uptime = uptime.String()

	return res
}
