package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/kardianos/service"
	"github.com/locug/sil"
)

// Program structures.
//  Define Start and Stop methods.
type program struct {
	exit   chan struct{}  // the exit channel, used to gracefully shut down
	logger service.Logger // logger service held in the type rather than global

	active map[int]*sil.SIL // the sil files that are active with int map as id passed to client
}

func main() {
	// program flags
	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()

	svcConfig := &service.Config{
		Name:        "SILAPI",
		DisplayName: "SIL API endpoint",
		Description: "Make SIL files by hitting an API.",
		Dependencies: []string{
			"Requires=network.target",
			"After=network-online.target syslog.target"},
	}

	p := &program{}
	s, err := service.New(p, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	errs := make(chan error, 5)
	p.logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	// make the active map
	p.active = make(map[int]*sil.SIL)

	go func() {
		for {
			var err = <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()

	if len(*svcFlag) != 0 {
		err = service.Control(s, *svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}
	err = s.Run()
	if err != nil {
		_ = p.logger.Error(err)
	}
}

func (p *program) Start(s service.Service) error {
	if service.Interactive() {
		_ = p.logger.Info("Running in terminal.")
	} else {
		_ = p.logger.Info("Running under service manager.")
	}
	p.exit = make(chan struct{})

	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *program) run() {
	_ = p.logger.Infof("Start SIL API Service")

	router := p.NewRouter()

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
	// ticker := time.NewTicker(2 * time.Second)
	// for {
	// 	select {
	// 	case tm := <-ticker.C:
	// 		p.logger.Infof("Still running at %v...", tm)
	// 	case <-p.exit:
	// 		ticker.Stop()
	// 		return nil
	// 	}
	// }
}

func (p *program) Stop(s service.Service) error {
	// Any work in Stop should be quick, usually a few seconds at most.
	_ = p.logger.Info("I'm Stopping!")
	close(p.exit)
	return nil
}
