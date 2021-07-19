package block7serv

import (
	"strconv"

	"github.com/valyala/fasthttp"
	block7http "github.com/zhs007/block7/http"
	block7utils "github.com/zhs007/block7/utils"
	"go.uber.org/zap"
)

// BasicURL - basic url
const BasicURL = "/v1/games/"

// Serv -
type Serv struct {
	*block7http.Serv
	Service IService
}

// NewServ - new a serv
func NewServ(service IService) *Serv {
	cfg := service.GetConfig()

	s := &Serv{
		block7http.NewServ(cfg.BindAddr, cfg.IsDebugMode),
		service,
	}

	s.RegHandle(block7utils.AppendString(BasicURL, "login"),
		func(ctx *fasthttp.RequestCtx, serv *block7http.Serv) {
			if !ctx.Request.Header.IsGet() {
				s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

				return
			}

			params := &LoginParams{}
			ctx.QueryArgs().VisitAll(func(k []byte, v []byte) {
				if string(k) == "userHash" {
					params.UserHash = string(v)
				} else if string(k) == "game" {
					params.Game = string(v)
				} else if string(k) == "platform" {
					params.Platform = string(v)
				} else if string(k) == "adid" {
					params.ADID = string(v)
				} else if string(k) == "guid" {
					params.GUID = string(v)
				} else if string(k) == "platformInfo" {
					params.PlatformInfo = string(v)
				} else if string(k) == "gameVersion" {
					params.GameVersion = string(v)
				} else if string(k) == "resVersion" {
					params.ResourceVersion = string(v)
				} else if string(k) == "deviceInfo" {
					params.DeviceInfo = string(v)
				}
			})

			block7utils.Debug("block7serv.Serv.login:ParseBody",
				block7utils.JSON("params", params))

			ret, err := s.Service.Login(params)
			if err != nil {
				block7utils.Warn("block7serv.Serv.login:Login",
					zap.Error(err))

				s.SetHTTPStatus(ctx, fasthttp.StatusInternalServerError)

				return
			}

			if cfg.IsDebugMode {
				block7utils.Debug("block7serv.Serv.login",
					block7utils.JSON("result", ret))
			}

			s.SetResponse(ctx, ret)
		})

	s.RegHandle(block7utils.AppendString(BasicURL, "mission"),
		func(ctx *fasthttp.RequestCtx, serv *block7http.Serv) {
			if !ctx.Request.Header.IsGet() {
				s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

				return
			}

			params := &MissionParams{}
			ctx.QueryArgs().VisitAll(func(k []byte, v []byte) {
				if string(k) == "userHash" {
					params.UserHash = string(v)
				} else if string(k) == "missionid" {
					i, err := strconv.Atoi(string(v))
					if err != nil {
						block7utils.Warn("block7serv.Serv.mission:VisitAll:missionid",
							zap.Error(err))
					} else {
						params.MissionID = i
					}
				}
			})

			block7utils.Debug("block7serv.Serv.mission:ParseBody",
				block7utils.JSON("params", params))

			if params.MissionID <= 0 {
				block7utils.Warn("block7serv.Serv.mission:ParseBody",
					zap.Int("missionid", params.MissionID))

				s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

				return
			}

			ret, err := s.Service.Mission(params)
			if err != nil {
				block7utils.Warn("block7serv.Serv.mission:Mission",
					zap.Error(err))

				s.SetHTTPStatus(ctx, fasthttp.StatusInternalServerError)

				return
			}

			if cfg.IsDebugMode {
				block7utils.Debug("block7serv.Serv.mission",
					block7utils.JSON("result", ret))
			}

			s.SetResponse(ctx, ret)
		})

	s.RegHandle(block7utils.AppendString(BasicURL, "missiondata"),
		func(ctx *fasthttp.RequestCtx, serv *block7http.Serv) {
			if !ctx.Request.Header.IsPost() {
				s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

				return
			}

			params := &MissionDataParams{}
			err := s.ParseBody(ctx, params)
			if err != nil {
				block7utils.Warn("block7serv.Serv.missiondata:ParseBody",
					zap.Error(err))

				s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

				return
			}

			block7utils.Debug("block7serv.Serv.missiondata:ParseBody",
				block7utils.JSON("params", params))

			ret, err := s.Service.MissionData(params)
			if err != nil {
				block7utils.Warn("block7serv.Serv.missiondata:MissionData",
					zap.Error(err))

				s.SetHTTPStatus(ctx, fasthttp.StatusInternalServerError)

				return
			}

			if cfg.IsDebugMode {
				block7utils.Debug("block7serv.Serv.missiondata",
					block7utils.JSON("result", ret))
			}

			s.SetResponse(ctx, ret)
		})

	return s
}
