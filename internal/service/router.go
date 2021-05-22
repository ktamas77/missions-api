package service

import (
    "github.com/go-chi/chi"
    "github.com/redcuckoo/bsc-checker-events/internal/config"
    "github.com/redcuckoo/bsc-checker-events/internal/data/pg"
    "github.com/redcuckoo/bsc-checker-events/internal/service/handlers"
    "github.com/redcuckoo/bsc-checker-events/internal/service/helpers"
    "gitlab.com/distributed_lab/ape"
)

func (s *service) router(cfg config.Config) chi.Router {
    r := chi.NewRouter()

    r.Use(
        ape.RecoverMiddleware(s.log),
        ape.LoganMiddleware(s.log), // this line may cause compilation error but in general case `dep ensure -v` will fix it
        ape.CtxMiddleware(
            helpers.CtxLog(s.log),
            helpers.CtxMission(pg.NewMissionQ(cfg.DB())),
        ),
    )

    // configure endpoints here
    r.Route("/missions", func (r chi.Router){
        //r.Get("/{missionId}", handlers.getMissionById)
        r.Get("/", handlers.GetListMissions)
    })

    //r.Route("/explorers", func(r chi.Router){
    //    //r.Get("{explorerId}", handlers.getMissionsByExplorerId)
    //})

    return r
}
