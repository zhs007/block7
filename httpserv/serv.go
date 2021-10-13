package block7serv

import (
	"strconv"
	"strings"

	realip "github.com/Ferluci/fast-realip"
	"github.com/valyala/fasthttp"
	block7http "github.com/zhs007/block7/http"
	goutils "github.com/zhs007/goutils"
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

	s.RegHandle(goutils.AppendString(BasicURL, "login"),
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

			ipaddr := realip.FromRequest(ctx)
			params.IPAddr = ipaddr

			goutils.Debug("block7serv.Serv.login:ParseBody",
				goutils.JSON("params", params))

			ret, err := s.Service.Login(params)
			if err != nil {
				goutils.Warn("block7serv.Serv.login:Login",
					zap.Error(err))

				s.SetHTTPStatus(ctx, fasthttp.StatusInternalServerError)

				return
			}

			if cfg.IsDebugMode {
				goutils.Debug("block7serv.Serv.login",
					goutils.JSON("result", ret))
			}

			s.SetResponse(ctx, ret)
		})

	s.RegHandle(goutils.AppendString(BasicURL, "mission"),
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
						goutils.Warn("block7serv.Serv.mission:VisitAll:missionid",
							zap.Error(err))
					} else {
						params.MissionID = i
					}
				} else if string(k) == "mission" {
					i64, err := strconv.ParseInt(string(v), 10, 64)
					if err != nil {
						goutils.Warn("block7serv.Serv.mission:VisitAll:mission",
							zap.Error(err))
					} else {
						params.SceneID = i64
					}
				} else if string(k) == "history" {
					i64, err := strconv.ParseInt(string(v), 10, 64)
					if err != nil {
						goutils.Warn("block7serv.Serv.mission:VisitAll:history",
							zap.Error(err))
					} else {
						params.HistoryID = i64
					}
				}
			})

			goutils.Debug("block7serv.Serv.mission:ParseBody",
				goutils.JSON("params", params))

			// if params.HistoryID > 0 && params.MissionID == 0 {

			// }

			// if params.MissionID <= 0 {
			// 	goutils.Warn("block7serv.Serv.mission:ParseBody",
			// 		zap.Int("missionid", params.MissionID))

			// 	s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

			// 	return
			// }

			ret, err := s.Service.Mission(params)
			if err != nil {
				goutils.Warn("block7serv.Serv.mission:Mission",
					zap.Error(err))

				s.SetHTTPStatus(ctx, fasthttp.StatusInternalServerError)

				return
			}

			if cfg.IsDebugMode {
				goutils.Debug("block7serv.Serv.mission",
					goutils.JSON("result", ret))
			}

			s.SetResponse(ctx, ret)
		})

	s.RegHandle(goutils.AppendString(BasicURL, "missiondata"),
		func(ctx *fasthttp.RequestCtx, serv *block7http.Serv) {
			if !ctx.Request.Header.IsPost() {
				s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

				return
			}

			params, err := parseMissionDataParams(ctx.PostBody())
			if err != nil {
				goutils.Warn("block7serv.Serv.missiondata:parseMissionDataParams",
					zap.Error(err))

				s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

				return
			}

			// params := &MissionDataParams{}
			// err := s.ParseBody(ctx, params)
			// if err != nil {
			// 	goutils.Warn("block7serv.Serv.missiondata:ParseBody",
			// 		zap.Error(err))

			// 	s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

			// 	return
			// }

			goutils.Debug("block7serv.Serv.missiondata:ParseBody",
				goutils.JSON("params", params))

			ret, err := s.Service.MissionData(params)
			if err != nil {
				goutils.Warn("block7serv.Serv.missiondata:MissionData",
					zap.Error(err))

				s.SetHTTPStatus(ctx, fasthttp.StatusInternalServerError)

				return
			}

			if cfg.IsDebugMode {
				goutils.Debug("block7serv.Serv.missiondata",
					goutils.JSON("result", ret))
			}

			s.SetResponse(ctx, ret)
		})

	s.RegHandle(goutils.AppendString(BasicURL, "userdata"),
		func(ctx *fasthttp.RequestCtx, serv *block7http.Serv) {
			if !ctx.Request.Header.IsGet() {
				s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

				return
			}

			params := &UserDataParams{}
			ctx.QueryArgs().VisitAll(func(k []byte, v []byte) {
				if string(k) == "name" {
					params.Name = strings.TrimSpace(string(v))
				} else if string(k) == "platform" {
					params.Platform = strings.TrimSpace(string(v))
				}
			})

			// goutils.Debug("block7serv.Serv.userdata:ParseBody",
			// 	goutils.JSON("params", params))

			if params.Name == "" || params.Platform == "" {
				goutils.Warn("block7serv.Serv.userdata:ParseBody",
					zap.String("name", params.Name),
					zap.String("platform", params.Platform))

				s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

				return
			}

			ret, err := s.Service.GetUserData(params)
			if err != nil {
				goutils.Warn("block7serv.Serv.userdata:GetUserData",
					zap.Error(err))

				s.SetHTTPStatus(ctx, fasthttp.StatusInternalServerError)

				return
			}

			// if cfg.IsDebugMode {
			// 	goutils.Debug("block7serv.Serv.userdata",
			// 		goutils.JSON("result", ret))
			// }

			s.SetResponse(ctx, ret)
		})

	s.RegHandle(goutils.AppendString(BasicURL, "upduserdata"),
		func(ctx *fasthttp.RequestCtx, serv *block7http.Serv) {
			if !ctx.Request.Header.IsPost() {
				s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

				return
			}

			ud, uds, err := parseUpdUserDataParams(ctx.PostBody())
			if err != nil {
				goutils.Warn("block7serv.Serv.upduserdata:parseUpdUserDataParams",
					zap.Error(err))

				s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

				return
			}

			// goutils.Debug("block7serv.Serv.upduserdata:ParseBody",
			// 	goutils.JSON("params", ud))

			ret, err := s.Service.UpdUserData(ud, uds)
			if err != nil {
				goutils.Warn("block7serv.Serv.upduserdata:MissionData",
					zap.Error(err))

				s.SetHTTPStatus(ctx, fasthttp.StatusInternalServerError)

				return
			}

			// if cfg.IsDebugMode {
			// 	goutils.Debug("block7serv.Serv.upduserdata",
			// 		goutils.JSON("result", ret))
			// }

			s.SetResponse(ctx, ret)
		})

	s.RegHandle(goutils.AppendString(BasicURL, "stats"),
		func(ctx *fasthttp.RequestCtx, serv *block7http.Serv) {
			if !ctx.Request.Header.IsGet() {
				s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

				return
			}

			params := &StatsParams{}
			ctx.QueryArgs().VisitAll(func(k []byte, v []byte) {
				if string(k) == "token" {
					params.Token = strings.TrimSpace(string(v))
				}
			})

			goutils.Debug("block7serv.Serv.stats:ParseBody",
				goutils.JSON("params", params))

			if params.Token == "" {
				goutils.Warn("block7serv.Serv.stats:Token:nil",
					zap.Error(ErrInvalidToken))

				s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

				return
			}

			ret, err := s.Service.Stats(params)
			if err != nil {
				goutils.Warn("block7serv.Serv.stats:GetUserData",
					zap.Error(err))

				s.SetHTTPStatus(ctx, fasthttp.StatusInternalServerError)

				return
			}

			if cfg.IsDebugMode {
				goutils.Debug("block7serv.Serv.stats",
					goutils.JSON("result", ret))
			}

			s.SetResponse(ctx, ret)
		})

	s.RegHandle(goutils.AppendString(BasicURL, "userstats"),
		func(ctx *fasthttp.RequestCtx, serv *block7http.Serv) {
			if !ctx.Request.Header.IsGet() {
				s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

				return
			}

			params := &UserStatsParams{}
			ctx.QueryArgs().VisitAll(func(k []byte, v []byte) {
				if string(k) == "token" {
					params.Token = strings.TrimSpace(string(v))
				} else if string(k) == "uid" {
					uid, err := goutils.String2Int64(strings.TrimSpace(string(v)))
					if err != nil {
						goutils.Warn("block7serv.Serv.userstats:ParseBody",
							zap.String("uid", string(v)),
							goutils.JSON("params", params))

						return
					}

					params.UserID = uid
				} else if string(k) == "userHash" {
					params.UserHash = strings.TrimSpace(string(v))
				}
			})

			goutils.Debug("block7serv.Serv.userstats:ParseBody",
				goutils.JSON("params", params))

			if params.Token == "" {
				goutils.Warn("block7serv.Serv.userstats:Token:nil",
					zap.Error(ErrInvalidToken))

				s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

				return
			}

			ret, err := s.Service.UserStats(params)
			if err != nil {
				goutils.Warn("block7serv.Serv.userstats:GetUserData",
					zap.Error(err))

				s.SetHTTPStatus(ctx, fasthttp.StatusInternalServerError)

				return
			}

			if cfg.IsDebugMode {
				goutils.Debug("block7serv.Serv.userstats",
					goutils.JSON("result", ret))
			}

			s.SetResponse(ctx, ret)
		})

	return s
}

// Stop - stop a server
func (s *Serv) Stop() error {
	s.Service.Stop()

	return s.Serv.Stop()
}

// Start - start a server
func (s *Serv) Start() error {
	s.Service.Start()

	return s.Serv.Start()
}
