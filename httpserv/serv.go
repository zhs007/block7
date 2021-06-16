package block7serv

import (
	"strconv"

	"github.com/valyala/fasthttp"
	block7 "github.com/zhs007/block7"
	block7http "github.com/zhs007/block7/http"
	"go.uber.org/zap"
)

// BasicURL - basic url
const BasicURL = "/v1/games/"

// Serv -
type Serv struct {
	*block7http.Serv
	Service IService
	cfg     *Config
}

// NewServ - new a serv
func NewServ(service IService, cfg *Config) *Serv {
	s := &Serv{
		block7http.NewServ(cfg.BindAddr, cfg.IsDebugMode),
		service,
		cfg,
	}

	s.RegHandle(block7.AppendString(BasicURL, "mission"),
		func(ctx *fasthttp.RequestCtx, serv *block7http.Serv) {
			if !ctx.Request.Header.IsGet() {
				s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

				return
			}

			params := &MissionParams{}
			ctx.QueryArgs().VisitAll(func(k []byte, v []byte) {
				if string(k) == "missionid" {
					i, err := strconv.Atoi(string(v))
					if err != nil {
						block7.Warn("block7serv.Serv.mission:VisitAll:missionid",
							zap.Error(err))
					} else {
						params.MissionID = i
					}
				}
			})

			if params.MissionID <= 0 {
				block7.Warn("block7serv.Serv.mission:ParseBody",
					zap.Int("missionid", params.MissionID))

				s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

				return
			}

			ret, err := s.Service.Mission(params)
			if err != nil {
				block7.Warn("block7serv.Serv.mission:Mission",
					zap.Error(err))

				s.SetHTTPStatus(ctx, fasthttp.StatusInternalServerError)

				return
			}

			s.SetResponse(ctx, ret)
		})

	s.RegHandle(block7.AppendString(BasicURL, "/missiondata"),
		func(ctx *fasthttp.RequestCtx, serv *block7http.Serv) {
			if !ctx.Request.Header.IsPost() {
				s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

				return
			}

			params := &MissionParams{}
			err := s.ParseBody(ctx, params)
			if err != nil {
				block7.Warn("block7serv.Serv.mission:ParseBody",
					zap.Error(err))

				s.SetHTTPStatus(ctx, fasthttp.StatusBadRequest)

				return
			}

			ret, err := s.Service.Mission(params)
			if err != nil {
				block7.Warn("block7serv.Serv.mission:Mission",
					zap.Error(err))

				s.SetHTTPStatus(ctx, fasthttp.StatusInternalServerError)

				return
			}

			s.SetResponse(ctx, ret)
		})

	return s
}
